// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package patch

import (
	"fmt"
	"reflect"
	"testing"

	webhook2 "github.com/apache/dubbo-kubernetes/pkg/config/webhook"

	"github.com/apache/dubbo-kubernetes/pkg/core/client/webhook"

	dubbo_cp "github.com/apache/dubbo-kubernetes/pkg/config/app/dubbo-cp"
	"github.com/apache/dubbo-kubernetes/pkg/config/kube"
	"github.com/apache/dubbo-kubernetes/pkg/config/security"
	"github.com/apache/dubbo-kubernetes/pkg/config/server"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type fakeKubeClient struct {
	webhook.Client
}

func (f *fakeKubeClient) GetNamespaceLabels(namespace string) map[string]string {
	if namespace == "matched" {
		return map[string]string{
			"dubbo-ca.inject":        "true",
			RegistryInjectNacosLabel: Labeled,
		}
	} else {
		return map[string]string{}
	}
}

func (f *fakeKubeClient) ListServices(namespace string, listOptions metav1.ListOptions) *v1.ServiceList {
	if namespace != "matched" {
		return nil
	}

	for _, registry := range registryInjectLabelPriorities {
		if listOptions.LabelSelector == fmt.Sprintf("%s=%s", registry, Labeled) {
			if registry == RegistryInjectK8sLabel { // k8s registry
				return nil
			}

			return &v1.ServiceList{
				Items: []v1.Service{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      fmt.Sprintf("%s-registry", registrySchemas[registry]),
							Namespace: namespace,
						},
					},
				},
			}
		}
	}

	if listOptions.LabelSelector == fmt.Sprintf("%s=%s", "dubbo.apache.org/prometheus", Labeled) {
		return &v1.ServiceList{
			Items: []v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: namespace,
					},
				},
			},
		}
	}

	return nil
}

func TestEmpty(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{}

	newPod, _ := sdk.NewPodWithDubboCa(pod)

	if !reflect.DeepEqual(newPod, pod) {
		t.Error("should be equal")
	}
}

func TestInjectFromLabel(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{}

	pod.Labels = make(map[string]string)
	pod.Labels["dubbo-ca.inject"] = "true"

	newPod, _ := sdk.NewPodWithDubboCa(pod)

	if reflect.DeepEqual(newPod, pod) {
		t.Error("should not be equal")
	}
}

func TestInjectFromNs(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{}

	pod.Namespace = "matched"

	newPod, _ := sdk.NewPodWithDubboCa(pod)

	if reflect.DeepEqual(newPod, pod) {
		t.Error("should not be equal")
	}
}

func TestInjectVolumes(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{}

	pod.Namespace = "matched"

	newPod, _ := sdk.NewPodWithDubboCa(pod)

	if reflect.DeepEqual(newPod, pod) {
		t.Error("should not be equal")
	}

	if len(newPod.Spec.Volumes) != 2 {
		t.Error("should have 1 volume")
	}

	if newPod.Spec.Volumes[0].Name != "dubbo-ca-token" {
		t.Error("should have dubbo-ca-token volume")
	}

	if len(newPod.Spec.Volumes[0].Projected.Sources) != 1 {
		t.Error("should have 1 projected source")
	}

	if newPod.Spec.Volumes[0].Projected.Sources[0].ServiceAccountToken.Path != "token" {
		t.Error("should have token path")
	}

	if newPod.Spec.Volumes[0].Projected.Sources[0].ServiceAccountToken.Audience != "dubbo-ca" {
		t.Error("should have dubbo-ca audience")
	}

	if *newPod.Spec.Volumes[0].Projected.Sources[0].ServiceAccountToken.ExpirationSeconds != 1800 {
		t.Error("should have 1800 expiration seconds")
	}

	if newPod.Spec.Volumes[1].Name != "dubbo-ca-cert" {
		t.Error("should have dubbo-ca-cert volume")
	}

	if len(newPod.Spec.Volumes[1].Projected.Sources) != 1 {
		t.Error("should have 1 projected source")
	}

	if newPod.Spec.Volumes[1].Projected.Sources[0].ConfigMap.Name != "dubbo-ca-cert" {
		t.Error("should have dubbo-ca-cert configmap")
	}

	if len(newPod.Spec.Volumes[1].Projected.Sources[0].ConfigMap.Items) != 1 {
		t.Error("should have 1 item")
	}

	if newPod.Spec.Volumes[1].Projected.Sources[0].ConfigMap.Items[0].Key != "ca.crt" {
		t.Error("should have ca.crt key")
	}

	if newPod.Spec.Volumes[1].Projected.Sources[0].ConfigMap.Items[0].Path != "ca.crt" {
		t.Error("should have ca.crt path")
	}
}

func TestInjectOneContainer(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{}

	pod.Namespace = "matched"

	pod.Spec.Containers = make([]v1.Container, 1)
	pod.Spec.Containers[0].Name = "test"

	newPod, _ := sdk.NewPodWithDubboCa(pod)

	if reflect.DeepEqual(newPod, pod) {
		t.Error("should not be equal")
	}

	if len(newPod.Spec.Containers) != 1 {
		t.Error("should have 1 container")
	}

	container := newPod.Spec.Containers[0]
	checkContainer(t, container)
}

func TestInjectTwoContainer(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{}

	pod.Namespace = "matched"

	pod.Spec.Containers = make([]v1.Container, 2)
	pod.Spec.Containers[0].Name = "test"
	pod.Spec.Containers[1].Name = "test"

	newPod, _ := sdk.NewPodWithDubboCa(pod)

	if reflect.DeepEqual(newPod, pod) {
		t.Error("should not be equal")
	}

	if len(newPod.Spec.Containers) != 2 {
		t.Error("should have 2 container")
	}

	container := newPod.Spec.Containers[0]
	checkContainer(t, container)

	container = newPod.Spec.Containers[1]
	checkContainer(t, container)
}

func checkContainer(t *testing.T, container v1.Container) {
	if container.Name != "test" {
		t.Error("should have test container")
	}

	if len(container.Env) != 4 {
		t.Error("should have 3 env")
	}

	if container.Env[0].Name != "DUBBO_CA_ADDRESS" {
		t.Error("should have DUBBO_CA_ADDRESS env")
	}

	if container.Env[0].Value != "dubbo-ca.dubbo-system.svc:30062" {
		t.Error("should have dubbo-ca.dubbo-system.svc:30062 value")
	}

	if container.Env[1].Name != "DUBBO_CA_CERT_PATH" {
		t.Error("should have DUBBO_CA_TOKEN_PATH env")
	}

	if container.Env[1].Value != "/var/run/secrets/dubbo-ca-cert/ca.crt" {
		t.Error("should have /var/run/secrets/dubbo-ca-cert/ca.crt value")
	}

	if container.Env[2].Name != "DUBBO_OIDC_TOKEN" {
		t.Error("should have DUBBO_OIDC_TOKEN env")
	}

	if container.Env[2].Value != "/var/run/secrets/dubbo-ca-token/token" {
		t.Error("should have /var/run/secrets/dubbo-ca-token/token value")
	}

	if container.Env[3].Name != "DUBBO_OIDC_TOKEN_TYPE" {
		t.Error("should have DUBBO_OIDC_TOKEN_TYPE env")
	}

	if container.Env[3].Value != "dubbo-ca-token" {
		t.Error("should have dubbo-ca-token value")
	}

	if len(container.VolumeMounts) != 2 {
		t.Error("should have 2 volume mounts")
	}

	if container.VolumeMounts[0].Name != "dubbo-ca-token" {
		t.Error("should have dubbo-ca-token volume mount")
	}

	if container.VolumeMounts[0].MountPath != "/var/run/secrets/dubbo-ca-token" {
		t.Error("should have /var/run/secrets/dubbo-ca-token mount path")
	}

	if container.VolumeMounts[1].Name != "dubbo-ca-cert" {
		t.Error("should have dubbo-ca-cert volume mount")
	}

	if container.VolumeMounts[1].MountPath != "/var/run/secrets/dubbo-ca-cert" {
		t.Error("should have /var/run/secrets/dubbo-ca-cert mount path")
	}
}

func TestCheckVolume1(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{}

	pod.Namespace = "matched"

	pod.Spec.Containers = make([]v1.Container, 1)
	pod.Spec.Containers[0].Name = "test"

	pod.Spec.Volumes = make([]v1.Volume, 1)
	pod.Spec.Volumes[0].Name = "dubbo-ca-token"

	newPod, _ := sdk.NewPodWithDubboCa(pod)

	if !reflect.DeepEqual(newPod, pod) {
		t.Error("should be equal")
	}
}

func TestCheckVolume2(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{}

	pod.Namespace = "matched"

	pod.Spec.Containers = make([]v1.Container, 1)
	pod.Spec.Containers[0].Name = "test"

	pod.Spec.Volumes = make([]v1.Volume, 1)
	pod.Spec.Volumes[0].Name = "dubbo-ca-cert"

	newPod, _ := sdk.NewPodWithDubboCa(pod)

	if !reflect.DeepEqual(newPod, pod) {
		t.Error("should be equal")
	}
}

func TestCheckEnv1(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{}

	pod.Namespace = "matched"

	pod.Spec.Containers = make([]v1.Container, 1)
	pod.Spec.Containers[0].Name = "test"

	pod.Spec.Containers[0].Env = make([]v1.EnvVar, 1)
	pod.Spec.Containers[0].Env[0].Name = "DUBBO_CA_ADDRESS"

	newPod, _ := sdk.NewPodWithDubboCa(pod)

	if !reflect.DeepEqual(newPod, pod) {
		t.Error("should be equal")
	}
}

func TestCheckEnv2(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{}

	pod.Namespace = "matched"

	pod.Spec.Containers = make([]v1.Container, 1)
	pod.Spec.Containers[0].Name = "test"

	pod.Spec.Containers[0].Env = make([]v1.EnvVar, 1)
	pod.Spec.Containers[0].Env[0].Name = "DUBBO_CA_CERT_PATH"

	newPod, _ := sdk.NewPodWithDubboCa(pod)

	if !reflect.DeepEqual(newPod, pod) {
		t.Error("should be equal")
	}
}

func TestCheckEnv3(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{}

	pod.Namespace = "matched"

	pod.Spec.Containers = make([]v1.Container, 1)
	pod.Spec.Containers[0].Name = "test"

	pod.Spec.Containers[0].Env = make([]v1.EnvVar, 1)
	pod.Spec.Containers[0].Env[0].Name = "DUBBO_OIDC_TOKEN"

	newPod, _ := sdk.NewPodWithDubboCa(pod)

	if !reflect.DeepEqual(newPod, pod) {
		t.Error("should be equal")
	}
}

func TestCheckEnv4(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{}

	pod.Namespace = "matched"

	pod.Spec.Containers = make([]v1.Container, 2)
	pod.Spec.Containers[0].Name = "test"
	pod.Spec.Containers[1].Name = "test"

	pod.Spec.Containers[1].Env = make([]v1.EnvVar, 1)
	pod.Spec.Containers[1].Env[0].Name = "DUBBO_OIDC_TOKEN"

	newPod, _ := sdk.NewPodWithDubboCa(pod)

	if !reflect.DeepEqual(newPod, pod) {
		t.Error("should be equal")
	}
}

func TestCheckContainerVolume1(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{}

	pod.Namespace = "matched"

	pod.Spec.Containers = make([]v1.Container, 1)
	pod.Spec.Containers[0].Name = "test"

	pod.Spec.Containers[0].VolumeMounts = make([]v1.VolumeMount, 1)
	pod.Spec.Containers[0].VolumeMounts[0].Name = "dubbo-ca-token"

	newPod, _ := sdk.NewPodWithDubboCa(pod)

	if !reflect.DeepEqual(newPod, pod) {
		t.Error("should be equal")
	}
}

func TestCheckContainerVolume2(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{}

	pod.Namespace = "matched"

	pod.Spec.Containers = make([]v1.Container, 1)
	pod.Spec.Containers[0].Name = "test"

	pod.Spec.Containers[0].VolumeMounts = make([]v1.VolumeMount, 1)
	pod.Spec.Containers[0].VolumeMounts[0].Name = "dubbo-ca-cert"

	newPod, _ := sdk.NewPodWithDubboCa(pod)

	if !reflect.DeepEqual(newPod, pod) {
		t.Error("should be equal")
	}
}

func TestCheckContainerVolume3(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{}

	pod.Namespace = "matched"

	pod.Spec.Containers = make([]v1.Container, 2)
	pod.Spec.Containers[0].Name = "test"
	pod.Spec.Containers[1].Name = "test"

	pod.Spec.Containers[1].VolumeMounts = make([]v1.VolumeMount, 1)
	pod.Spec.Containers[1].VolumeMounts[0].Name = "dubbo-ca-cert"

	newPod, _ := sdk.NewPodWithDubboCa(pod)

	if !reflect.DeepEqual(newPod, pod) {
		t.Error("should be equal")
	}
}

func TestZkRegistryInjectFromLabel(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "matched",
			Labels: map[string]string{
				RegistryInjectZookeeperLabel: Labeled,
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{},
			},
		},
	}

	newPod, err := sdk.NewPodWithDubboRegistryInject(pod)
	if err != nil {
		t.Error(err.Error())
	}
	if !checkExpectedEnv(newPod, EnvDubboRegistryAddress, "zookeeper://zookeeper-registry.matched.svc") {
		t.Error("registry should be injected")
	}
}

func TestNacosRegistryInjectFromLabel(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "matched",
			Labels: map[string]string{
				RegistryInjectNacosLabel: Labeled,
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{},
			},
		},
	}

	newPod, err := sdk.NewPodWithDubboRegistryInject(pod)
	if err != nil {
		t.Error(err.Error())
	}
	if !checkExpectedEnv(newPod, EnvDubboRegistryAddress, "nacos://nacos-registry.matched.svc") {
		t.Error("registry should be injected")
	}
}

func TestK8sRegistryInjectFromLabel(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "matched",
			Labels: map[string]string{
				RegistryInjectK8sLabel: Labeled,
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{},
			},
		},
	}

	newPod, err := sdk.NewPodWithDubboRegistryInject(pod)
	if err != nil {
		t.Error(err.Error())
	}
	if !checkExpectedEnv(newPod, EnvDubboRegistryAddress, DefaultK8sRegistryAddress) {
		t.Error("registry should be injected")
	}
}

func TestRegistryNotInjectFromLabel(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}

	userSpecifiedAddress := "some address"
	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "matched",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Env: []v1.EnvVar{
						{
							Name:  EnvDubboRegistryAddress,
							Value: userSpecifiedAddress,
						},
					},
				},
			},
		},
	}

	newPod, err := sdk.NewPodWithDubboRegistryInject(pod)
	if err != nil {
		t.Error(err.Error())
	}
	if !checkExpectedEnv(newPod, EnvDubboRegistryAddress, userSpecifiedAddress) {
		t.Error("registry should not be injected")
	}
}

func TestRegistryInjectFromNs(t *testing.T) {
	t.Parallel()

	options := &dubbo_cp.Config{
		KubeConfig: kube.KubeConfig{
			IsKubernetesConnected: false,
			Namespace:             "dubbo-system",
			ServiceName:           "dubbo-ca",
		},
		Security: security.SecurityConfig{
			CaValidity:   30 * 24 * 60 * 60 * 1000, // 30 day
			CertValidity: 1 * 60 * 60 * 1000,       // 1 hour
		},
		Webhook: webhook2.Webhook{
			Port:       30080,
			AllowOnErr: false,
		},
		GrpcServer: server.ServerConfig{
			PlainServerPort:  30060,
			SecureServerPort: 30062,
			DebugPort:        30070,
		},
	}
	sdk := NewDubboSdk(options, &fakeKubeClient{}, nil)
	pod := &v1.Pod{
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{},
			},
		},
	}

	pod.Namespace = "matched"

	newPod, err := sdk.NewPodWithDubboRegistryInject(pod)
	if err != nil {
		t.Error(err.Error())
	}
	if !checkExpectedEnv(newPod, EnvDubboRegistryAddress, "nacos://nacos-registry.matched.svc") {
		t.Error("registry should be injected")
	}
}

func checkExpectedEnv(pod *v1.Pod, expectedEnvName, expectedEnvValue string) bool {
	if len(pod.Spec.Containers) <= 0 || len(pod.Spec.Containers[0].Env) <= 0 {
		return false
	}

	for _, env := range pod.Spec.Containers[0].Env {
		if env.Name == expectedEnvName {
			if env.Value == expectedEnvValue {
				return true
			}
		}
	}

	return false
}