/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/apache/dubbo-admin/pkg/core/gen/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AuthenticationPolicies returns a AuthenticationPolicyInformer.
	AuthenticationPolicies() AuthenticationPolicyInformer
	// AuthorizationPolicies returns a AuthorizationPolicyInformer.
	AuthorizationPolicies() AuthorizationPolicyInformer
	// ConditionRoutes returns a ConditionRouteInformer.
	ConditionRoutes() ConditionRouteInformer
	// DynamicConfigs returns a DynamicConfigInformer.
	DynamicConfigs() DynamicConfigInformer
	// ServiceNameMappings returns a ServiceNameMappingInformer.
	ServiceNameMappings() ServiceNameMappingInformer
	// TagRoutes returns a TagRouteInformer.
	TagRoutes() TagRouteInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AuthenticationPolicies returns a AuthenticationPolicyInformer.
func (v *version) AuthenticationPolicies() AuthenticationPolicyInformer {
	return &authenticationPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AuthorizationPolicies returns a AuthorizationPolicyInformer.
func (v *version) AuthorizationPolicies() AuthorizationPolicyInformer {
	return &authorizationPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConditionRoutes returns a ConditionRouteInformer.
func (v *version) ConditionRoutes() ConditionRouteInformer {
	return &conditionRouteInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DynamicConfigs returns a DynamicConfigInformer.
func (v *version) DynamicConfigs() DynamicConfigInformer {
	return &dynamicConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceNameMappings returns a ServiceNameMappingInformer.
func (v *version) ServiceNameMappings() ServiceNameMappingInformer {
	return &serviceNameMappingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TagRoutes returns a TagRouteInformer.
func (v *version) TagRoutes() TagRouteInformer {
	return &tagRouteInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}