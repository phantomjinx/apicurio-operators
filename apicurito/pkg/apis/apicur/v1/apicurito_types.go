/*
 * Copyright (C) 2020 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ApicuritoSpec defines the desired state of Apicurito
// +k8s:openapi-gen=true
type ApicuritoSpec struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Size"
	// +operator-sdk:csv:customresourcedefinitions:type=status,xDescriptors="urn:alm:descriptor:com.tectonic.ui:number"
	Size int32 `json:"size"`

	// The external hostname to access the UI service
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Application Route Host Name"
	// +operator-sdk:csv:customresourcedefinitions:type=status,xDescriptors="urn:alm:descriptor:text"
	UIRouteHostname string `json:"uiRouteHostname,omitempty"`

	// The external hostname to access the generator service
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Generator Route Host Name"
	// +operator-sdk:csv:customresourcedefinitions:type=status,xDescriptors="urn:alm:descriptor:text"
	GeneratorRouteHostname string `json:"generatorRouteHostname,omitempty"`
}

// ApicuritoStatus defines the observed state of Apicurito
// +k8s:openapi-gen=true
type ApicuritoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Apicurito is the Schema for the apicuritoes API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:resource:path=apicuritoes,scope=Namespaced
// +operator-sdk:csv:customresourcedefinitions:displayName="Apicurito"
// +operator-sdk:csv:customresourcedefinitions:resources={{ServiceAccount,v1},{ClusterRole,rbac.authorization.k8s.io/v1},{Role,rbac.authorization.k8s.io/v1},{Deployment,apps/v1}}
type Apicurito struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApicuritoSpec   `json:"spec,omitempty"`
	Status ApicuritoStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApicuritoList contains a list of Apicurito
type ApicuritoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apicurito `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Apicurito{}, &ApicuritoList{})
}