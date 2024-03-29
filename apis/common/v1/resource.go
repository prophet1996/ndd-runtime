/*
Copyright 2021 NDD.

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

package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

// LabelKeyNetworkNodeName is added to NetworkNodeUsages to relate them to their
// NetworkNode.
const LabelKeyNetworkNodeName = "ndd.yndd.io/networknode"

// A Reference to a named object.
type Reference struct {
	// Name of the referenced object.
	Name string `json:"name"`
}

// A TypedReference refers to an object by Name, Kind, and APIVersion. It is
// commonly used to reference cluster-scoped objects or objects where the
// namespace is already known.
type TypedReference struct {
	// APIVersion of the referenced object.
	APIVersion string `json:"apiVersion"`

	// Kind of the referenced object.
	Kind string `json:"kind"`

	// Name of the referenced object.
	Name string `json:"name"`

	// UID of the referenced object.
	// +optional
	UID types.UID `json:"uid,omitempty"`
}

// A Selector selects an object.
type Selector struct {
	// MatchLabels ensures an object with matching labels is selected.
	MatchLabels map[string]string `json:"matchLabels,omitempty"`

	// MatchControllerRef ensures an object with the same controller reference
	// as the selecting object is selected.
	MatchControllerRef *bool `json:"matchControllerRef,omitempty"`
}

// SetGroupVersionKind sets the Kind and APIVersion of a TypedReference.
func (obj *TypedReference) SetGroupVersionKind(gvk schema.GroupVersionKind) {
	obj.APIVersion, obj.Kind = gvk.ToAPIVersionAndKind()
}

// GroupVersionKind gets the GroupVersionKind of a TypedReference.
func (obj *TypedReference) GroupVersionKind() schema.GroupVersionKind {
	return schema.FromAPIVersionAndKind(obj.APIVersion, obj.Kind)
}

// GetObjectKind get the ObjectKind of a TypedReference.
func (obj *TypedReference) GetObjectKind() schema.ObjectKind { return obj }

// A ResourceSpec defines the desired state of a managed resource.
type ResourceSpec struct {
	// Active specifies if the managed resource is active or not
	// +kubebuilder:default=Active
	DeploymentPolicy DeploymentPolicy `json:"deploymentPolicy,omitempty"`

	// NetworkNodeReference specifies which network node will be used to
	// create, observe, update, and delete this managed resource
	NetworkNodeReference *Reference `json:"networkNodeRef"`

	// DeletionPolicy specifies what will happen to the underlying external
	// when this managed resource is deleted - either "Delete" or "Orphan" the
	// external resource.
	// +optional
	// +kubebuilder:default=Delete
	DeletionPolicy DeletionPolicy `json:"deletionPolicy,omitempty"`
}

// ResourceStatus represents the observed state of a managed resource.
type ResourceStatus struct {
	// the condition status
	ConditionedStatus `json:",inline"`
	// rootPaths define the rootPaths of the cr, used to monitor the resource status
	RootPaths []string `json:"rootPaths,omitempty"`
	// HierPaths tracks the hierarchical paths of the CR based on adjacent resources
	HierPaths map[string][]string `json:"hierPaths,omitempty"`
}

// A NetworkNodeStatus defines the observed status of a NetworkNode.
type NetworkNodeStatus struct {
	ConditionedStatus `json:",inline"`

	// Users of this network node configuration.
	Users int64 `json:"users,omitempty"`
}

type TransactionResourceStatus struct {
	// the condition status
	ConditionedStatus `json:",inline"`

	// conditionsStatus per device
	Device map[string]ConditionedStatus `json:"device,omitempty"`
}

/*
type TransactionDeviceStatus struct {
	//DeviceName *string `json:"deviceName,omitempty"`
	// the condition status
	ConditionedStatus `json:",inline"`
}
*/

// A NetworkNodeUsage is a record that a particular managed resource is using
// a particular network node configuration.
type NetworkNodeUsage struct {
	// NetworkNodeReference to the network node config being used.
	NetworkNodeReference Reference `json:"NetworkNodeRef"`

	// ResourceReference to the managed resource using the network node config.
	ResourceReference TypedReference `json:"resourceRef"`
}

type ResourceName struct {
	Name string `json:"name,omitempty"`
}
