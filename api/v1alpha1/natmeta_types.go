/*
Copyright 2021. Netris, Inc.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NatMetaSpec defines the desired state of NatMeta
type NatMetaSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Imported        bool   `json:"imported"`
	Reclaim         bool   `json:"reclaimPolicy"`
	NatCRGeneration int64  `json:"natGeneration"`
	ID              int    `json:"id"`
	NatName         string `json:"natName"`

	Comment    string `json:"comment,omitempty"`
	State      string `json:"state,omitempty"`
	SiteID     int    `json:"siteID"`
	Action     string `json:"action"`
	Protocol   string `json:"protocol"`
	SrcAddress string `json:"srcAddress"`
	SrcPort    string `json:"srcPort,omitempty"`
	DstAddress string `json:"dstAddress"`
	DstPort    string `json:"dstPort,omitempty"`
	SnatToIP   string `json:"snatToIp,omitempty"`
	SnatToPool string `json:"snatToPool,omitempty"`
	DnatToIP   string `json:"dnatToIp,omitempty"`
	DnatToPort int    `json:"dnatToPort,omitempty"`
}

// NatMetaStatus defines the observed state of NatMeta
type NatMetaStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NatMeta is the Schema for the natmeta API
type NatMeta struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NatMetaSpec   `json:"spec,omitempty"`
	Status NatMetaStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NatMetaList contains a list of NatMeta
type NatMetaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NatMeta `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NatMeta{}, &NatMetaList{})
}
