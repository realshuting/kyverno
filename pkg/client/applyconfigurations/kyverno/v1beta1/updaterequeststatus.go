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

package v1beta1

import (
	v1beta1 "github.com/kyverno/kyverno/api/kyverno/v1beta1"
	v1 "github.com/kyverno/kyverno/pkg/client/applyconfigurations/kyverno/v1"
)

// UpdateRequestStatusApplyConfiguration represents an declarative configuration of the UpdateRequestStatus type for use
// with apply.
type UpdateRequestStatusApplyConfiguration struct {
	Handler            *string                             `json:"handler,omitempty"`
	State              *v1beta1.UpdateRequestState         `json:"state,omitempty"`
	Message            *string                             `json:"message,omitempty"`
	GeneratedResources []v1.ResourceSpecApplyConfiguration `json:"generatedResources,omitempty"`
}

// UpdateRequestStatusApplyConfiguration constructs an declarative configuration of the UpdateRequestStatus type for use with
// apply.
func UpdateRequestStatus() *UpdateRequestStatusApplyConfiguration {
	return &UpdateRequestStatusApplyConfiguration{}
}

// WithHandler sets the Handler field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Handler field is set to the value of the last call.
func (b *UpdateRequestStatusApplyConfiguration) WithHandler(value string) *UpdateRequestStatusApplyConfiguration {
	b.Handler = &value
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *UpdateRequestStatusApplyConfiguration) WithState(value v1beta1.UpdateRequestState) *UpdateRequestStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *UpdateRequestStatusApplyConfiguration) WithMessage(value string) *UpdateRequestStatusApplyConfiguration {
	b.Message = &value
	return b
}

// WithGeneratedResources adds the given value to the GeneratedResources field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the GeneratedResources field.
func (b *UpdateRequestStatusApplyConfiguration) WithGeneratedResources(values ...*v1.ResourceSpecApplyConfiguration) *UpdateRequestStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithGeneratedResources")
		}
		b.GeneratedResources = append(b.GeneratedResources, *values[i])
	}
	return b
}
