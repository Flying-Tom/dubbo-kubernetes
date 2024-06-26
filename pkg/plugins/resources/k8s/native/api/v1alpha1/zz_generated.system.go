// Generated by tools/resource-gen
// Run "make generate" to update this file.

// nolint:whitespace
package v1alpha1

import (
	"fmt"
)

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

import (
	system_proto "github.com/apache/dubbo-kubernetes/api/system/v1alpha1"
	core_model "github.com/apache/dubbo-kubernetes/pkg/core/resources/model"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/resources/k8s/native/pkg/model"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/resources/k8s/native/pkg/registry"
	util_proto "github.com/apache/dubbo-kubernetes/pkg/util/proto"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Cluster
type DataSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo DataSource resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type DataSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataSource{}, &DataSourceList{})
}

func (cb *DataSource) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *DataSource) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *DataSource) GetMesh() string {
	return cb.Mesh
}

func (cb *DataSource) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *DataSource) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := system_proto.DataSource{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *DataSource) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*system_proto.DataSource)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *DataSource) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *DataSourceList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&system_proto.DataSource{}, &DataSource{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "DataSource",
		},
	})
	registry.RegisterListType(&system_proto.DataSource{}, &DataSourceList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "DataSourceList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Cluster
type Secret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo Secret resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type SecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Secret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Secret{}, &SecretList{})
}

func (cb *Secret) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *Secret) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *Secret) GetMesh() string {
	return cb.Mesh
}

func (cb *Secret) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *Secret) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := system_proto.Secret{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *Secret) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*system_proto.Secret)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *Secret) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *SecretList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&system_proto.Secret{}, &Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "Secret",
		},
	})
	registry.RegisterListType(&system_proto.Secret{}, &SecretList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "SecretList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Cluster
type Zone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo Zone resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Zone `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Zone{}, &ZoneList{})
}

func (cb *Zone) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *Zone) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *Zone) GetMesh() string {
	return cb.Mesh
}

func (cb *Zone) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *Zone) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := system_proto.Zone{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *Zone) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*system_proto.Zone)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *Zone) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *ZoneList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&system_proto.Zone{}, &Zone{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "Zone",
		},
	})
	registry.RegisterListType(&system_proto.Zone{}, &ZoneList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ZoneList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Cluster
type ZoneInsight struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo ZoneInsight resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type ZoneInsightList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZoneInsight `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ZoneInsight{}, &ZoneInsightList{})
}

func (cb *ZoneInsight) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *ZoneInsight) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *ZoneInsight) GetMesh() string {
	return cb.Mesh
}

func (cb *ZoneInsight) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *ZoneInsight) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := system_proto.ZoneInsight{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *ZoneInsight) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*system_proto.ZoneInsight)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *ZoneInsight) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *ZoneInsightList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&system_proto.ZoneInsight{}, &ZoneInsight{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ZoneInsight",
		},
	})
	registry.RegisterListType(&system_proto.ZoneInsight{}, &ZoneInsightList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ZoneInsightList",
		},
	})
}
