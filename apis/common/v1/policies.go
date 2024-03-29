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

// A DeploymentPolicy determines what should happen to the underlying external
// resource when a managed resource is deployed.
// +kubebuilder:validation:Enum=Active;Planned
type DeploymentPolicy string

const (
	// DeploymentActive means the external resource will deployed
	DeploymentActive DeploymentPolicy = "Active"

	// DeploymentPlanned means the resource identifier will be allocated but not deployed
	DeploymentPlanned DeploymentPolicy = "Planned"
)

// A DeletionPolicy determines what should happen to the underlying external
// resource when a managed resource is deleted.
// +kubebuilder:validation:Enum=Orphan;Delete
type DeletionPolicy string

const (
	// DeletionOrphan means the external resource will orphaned when its managed
	// resource is deleted.
	DeletionOrphan DeletionPolicy = "Orphan"

	// DeletionDelete means both the  external resource will be deleted when its
	// managed resource is deleted.
	DeletionDelete DeletionPolicy = "Delete"
)