/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

It is a controller's job to ensure that, for any given object, the actual state
of the world(both the cluster state, and potentially external state llike running containers
for Kubeleet for loadbalacners for a cloud provider) matches the desired state in the object.
Each controller focuses on one root Kind, but may interact with other kinds.

We call this process reconciling.

In controller-runtime, the logic that implements the reconciling for a specific kind is called a
Reconciler. A reconciler takes the name of an object, and retruns whether or not we need to ry again.

We have some package-level markers that denote that there are Kubernetes
objects in this package, and that this package represetns the group batch.tutorial.kubebuilder.io.
The object generator makes use of the former, while the latter is ued by the CRD generator to
generate the right metadatga for the CRDs it creates from this package.
*/

// Package v1 contains API Schema definitions for the batch v1 API group
// +kubebuilder:object:generate=true
// +groupName=batch.tutorial.kubebuilder.io
package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "batch.tutorial.kubebuilder.io", Version: "v1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
