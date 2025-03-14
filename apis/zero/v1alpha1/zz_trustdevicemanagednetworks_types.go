// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TrustDeviceManagedNetworksConfigInitParameters struct {

	// 256 hash of the TLS certificate presented by the host found at tls_sockaddr. If absent, regular certificate verification (trusted roots, valid timestamp, etc) will be used to validate the certificate.
	// The SHA-256 hash of the TLS certificate presented by the host found at tls_sockaddr. If absent, regular certificate verification (trusted roots, valid timestamp, etc) will be used to validate the certificate.
	Sha256 *string `json:"sha256,omitempty" tf:"sha256,omitempty"`

	// (String) A network address of the form "host:port" that the WARP client will use to detect the presence of a TLS host.
	// A network address of the form "host:port" that the WARP client will use to detect the presence of a TLS host.
	TLSSockaddr *string `json:"tlsSockaddr,omitempty" tf:"tls_sockaddr,omitempty"`
}

type TrustDeviceManagedNetworksConfigObservation struct {

	// 256 hash of the TLS certificate presented by the host found at tls_sockaddr. If absent, regular certificate verification (trusted roots, valid timestamp, etc) will be used to validate the certificate.
	// The SHA-256 hash of the TLS certificate presented by the host found at tls_sockaddr. If absent, regular certificate verification (trusted roots, valid timestamp, etc) will be used to validate the certificate.
	Sha256 *string `json:"sha256,omitempty" tf:"sha256,omitempty"`

	// (String) A network address of the form "host:port" that the WARP client will use to detect the presence of a TLS host.
	// A network address of the form "host:port" that the WARP client will use to detect the presence of a TLS host.
	TLSSockaddr *string `json:"tlsSockaddr,omitempty" tf:"tls_sockaddr,omitempty"`
}

type TrustDeviceManagedNetworksConfigParameters struct {

	// 256 hash of the TLS certificate presented by the host found at tls_sockaddr. If absent, regular certificate verification (trusted roots, valid timestamp, etc) will be used to validate the certificate.
	// The SHA-256 hash of the TLS certificate presented by the host found at tls_sockaddr. If absent, regular certificate verification (trusted roots, valid timestamp, etc) will be used to validate the certificate.
	// +kubebuilder:validation:Optional
	Sha256 *string `json:"sha256,omitempty" tf:"sha256,omitempty"`

	// (String) A network address of the form "host:port" that the WARP client will use to detect the presence of a TLS host.
	// A network address of the form "host:port" that the WARP client will use to detect the presence of a TLS host.
	// +kubebuilder:validation:Optional
	TLSSockaddr *string `json:"tlsSockaddr" tf:"tls_sockaddr,omitempty"`
}

type TrustDeviceManagedNetworksInitParameters struct {

	// (String)
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Attributes) The configuration object containing information for the WARP client to detect the managed network. (see below for nested schema)
	Config *TrustDeviceManagedNetworksConfigInitParameters `json:"config,omitempty" tf:"config,omitempty"`

	// (String) The type of device managed network.
	// The type of device managed network.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TrustDeviceManagedNetworksObservation struct {

	// (String)
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Attributes) The configuration object containing information for the WARP client to detect the managed network. (see below for nested schema)
	Config *TrustDeviceManagedNetworksConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	// (String) API UUID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) API UUID.
	// API UUID.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// (String) The type of device managed network.
	// The type of device managed network.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TrustDeviceManagedNetworksParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Attributes) The configuration object containing information for the WARP client to detect the managed network. (see below for nested schema)
	// +kubebuilder:validation:Optional
	Config *TrustDeviceManagedNetworksConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// (String) The type of device managed network.
	// The type of device managed network.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// TrustDeviceManagedNetworksSpec defines the desired state of TrustDeviceManagedNetworks
type TrustDeviceManagedNetworksSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrustDeviceManagedNetworksParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider TrustDeviceManagedNetworksInitParameters `json:"initProvider,omitempty"`
}

// TrustDeviceManagedNetworksStatus defines the observed state of TrustDeviceManagedNetworks.
type TrustDeviceManagedNetworksStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrustDeviceManagedNetworksObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TrustDeviceManagedNetworks is the Schema for the TrustDeviceManagedNetworkss API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type TrustDeviceManagedNetworks struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accountId) || (has(self.initProvider) && has(self.initProvider.accountId))",message="spec.forProvider.accountId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.config) || (has(self.initProvider) && has(self.initProvider.config))",message="spec.forProvider.config is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   TrustDeviceManagedNetworksSpec   `json:"spec"`
	Status TrustDeviceManagedNetworksStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrustDeviceManagedNetworksList contains a list of TrustDeviceManagedNetworkss
type TrustDeviceManagedNetworksList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrustDeviceManagedNetworks `json:"items"`
}

// Repository type metadata.
var (
	TrustDeviceManagedNetworks_Kind             = "TrustDeviceManagedNetworks"
	TrustDeviceManagedNetworks_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrustDeviceManagedNetworks_Kind}.String()
	TrustDeviceManagedNetworks_KindAPIVersion   = TrustDeviceManagedNetworks_Kind + "." + CRDGroupVersion.String()
	TrustDeviceManagedNetworks_GroupVersionKind = CRDGroupVersion.WithKind(TrustDeviceManagedNetworks_Kind)
)

func init() {
	SchemeBuilder.Register(&TrustDeviceManagedNetworks{}, &TrustDeviceManagedNetworksList{})
}
