/*
Copyright 2025 The Crossplane Authors.

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

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// VPCParameters are the configurable fields of a VPC.
type VPCParameters struct {
	Name string `json:"name"`
}

// VPCObservation are the observable fields of a VPC.
type VPCObservation struct {
	Name string `json:"name"`
}

// A VPCSpec defines the desired state of a VPC.
type VPCSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VPCParameters `json:"forProvider"`
}

// A VPCStatus represents the observed state of a VPC.
type VPCStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VPCObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A VPC is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex}
type VPC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VPCSpec   `json:"spec"`
	Status VPCStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCList contains a list of VPC
type VPCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPC `json:"items"`
}

// VPC type metadata.
var (
	VPCKind             = reflect.TypeOf(VPC{}).Name()
	VPCGroupKind        = schema.GroupKind{Group: Group, Kind: VPCKind}.String()
	VPCKindAPIVersion   = VPCKind + "." + SchemeGroupVersion.String()
	VPCGroupVersionKind = SchemeGroupVersion.WithKind(VPCKind)
)

func init() {
	SchemeBuilder.Register(&VPC{}, &VPCList{})
}
