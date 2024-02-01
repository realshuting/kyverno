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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v2alpha1

// GlobalContextEntrySpecApplyConfiguration represents an declarative configuration of the GlobalContextEntrySpec type for use
// with apply.
type GlobalContextEntrySpecApplyConfiguration struct {
	KubernetesResource *KubernetesResourceApplyConfiguration `json:"kubernetesResource,omitempty"`
	APICall            *ExternalAPICallApplyConfiguration    `json:"apiCall,omitempty"`
}

// GlobalContextEntrySpecApplyConfiguration constructs an declarative configuration of the GlobalContextEntrySpec type for use with
// apply.
func GlobalContextEntrySpec() *GlobalContextEntrySpecApplyConfiguration {
	return &GlobalContextEntrySpecApplyConfiguration{}
}

// WithKubernetesResource sets the KubernetesResource field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KubernetesResource field is set to the value of the last call.
func (b *GlobalContextEntrySpecApplyConfiguration) WithKubernetesResource(value *KubernetesResourceApplyConfiguration) *GlobalContextEntrySpecApplyConfiguration {
	b.KubernetesResource = value
	return b
}

// WithAPICall sets the APICall field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APICall field is set to the value of the last call.
func (b *GlobalContextEntrySpecApplyConfiguration) WithAPICall(value *ExternalAPICallApplyConfiguration) *GlobalContextEntrySpecApplyConfiguration {
	b.APICall = value
	return b
}
