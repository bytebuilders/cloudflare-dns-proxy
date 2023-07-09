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

package v1alpha1

// QueuingConfigurationApplyConfiguration represents an declarative configuration of the QueuingConfiguration type for use
// with apply.
type QueuingConfigurationApplyConfiguration struct {
	Queues           *int32 `json:"queues,omitempty"`
	HandSize         *int32 `json:"handSize,omitempty"`
	QueueLengthLimit *int32 `json:"queueLengthLimit,omitempty"`
}

// QueuingConfigurationApplyConfiguration constructs an declarative configuration of the QueuingConfiguration type for use with
// apply.
func QueuingConfiguration() *QueuingConfigurationApplyConfiguration {
	return &QueuingConfigurationApplyConfiguration{}
}

// WithQueues sets the Queues field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Queues field is set to the value of the last call.
func (b *QueuingConfigurationApplyConfiguration) WithQueues(value int32) *QueuingConfigurationApplyConfiguration {
	b.Queues = &value
	return b
}

// WithHandSize sets the HandSize field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HandSize field is set to the value of the last call.
func (b *QueuingConfigurationApplyConfiguration) WithHandSize(value int32) *QueuingConfigurationApplyConfiguration {
	b.HandSize = &value
	return b
}

// WithQueueLengthLimit sets the QueueLengthLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the QueueLengthLimit field is set to the value of the last call.
func (b *QueuingConfigurationApplyConfiguration) WithQueueLengthLimit(value int32) *QueuingConfigurationApplyConfiguration {
	b.QueueLengthLimit = &value
	return b
}
