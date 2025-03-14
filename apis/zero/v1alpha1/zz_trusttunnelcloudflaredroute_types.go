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

type TrustTunnelCloudflaredRouteInitParameters struct {

	// (String) Cloudflare account ID
	// Cloudflare account ID
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) Optional remark describing the route.
	// Optional remark describing the route.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// (String) The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// (String) UUID of the tunnel.
	// UUID of the tunnel.
	TunnelID *string `json:"tunnelId,omitempty" tf:"tunnel_id,omitempty"`

	// (String) UUID of the virtual network.
	// UUID of the virtual network.
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`
}

type TrustTunnelCloudflaredRouteObservation struct {

	// (String) Cloudflare account ID
	// Cloudflare account ID
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) Optional remark describing the route.
	// Optional remark describing the route.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// (String) Timestamp of when the resource was created.
	// Timestamp of when the resource was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (String) Timestamp of when the resource was deleted. If null, the resource has not been deleted.
	// Timestamp of when the resource was deleted. If `null`, the resource has not been deleted.
	DeletedAt *string `json:"deletedAt,omitempty" tf:"deleted_at,omitempty"`

	// (String) UUID of the route.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// (String) UUID of the tunnel.
	// UUID of the tunnel.
	TunnelID *string `json:"tunnelId,omitempty" tf:"tunnel_id,omitempty"`

	// (String) UUID of the virtual network.
	// UUID of the virtual network.
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`
}

type TrustTunnelCloudflaredRouteParameters struct {

	// (String) Cloudflare account ID
	// Cloudflare account ID
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) Optional remark describing the route.
	// Optional remark describing the route.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// (String) The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// (String) UUID of the tunnel.
	// UUID of the tunnel.
	// +kubebuilder:validation:Optional
	TunnelID *string `json:"tunnelId,omitempty" tf:"tunnel_id,omitempty"`

	// (String) UUID of the virtual network.
	// UUID of the virtual network.
	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`
}

// TrustTunnelCloudflaredRouteSpec defines the desired state of TrustTunnelCloudflaredRoute
type TrustTunnelCloudflaredRouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrustTunnelCloudflaredRouteParameters `json:"forProvider"`
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
	InitProvider TrustTunnelCloudflaredRouteInitParameters `json:"initProvider,omitempty"`
}

// TrustTunnelCloudflaredRouteStatus defines the observed state of TrustTunnelCloudflaredRoute.
type TrustTunnelCloudflaredRouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrustTunnelCloudflaredRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TrustTunnelCloudflaredRoute is the Schema for the TrustTunnelCloudflaredRoutes API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type TrustTunnelCloudflaredRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accountId) || (has(self.initProvider) && has(self.initProvider.accountId))",message="spec.forProvider.accountId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.network) || (has(self.initProvider) && has(self.initProvider.network))",message="spec.forProvider.network is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tunnelId) || (has(self.initProvider) && has(self.initProvider.tunnelId))",message="spec.forProvider.tunnelId is a required parameter"
	Spec   TrustTunnelCloudflaredRouteSpec   `json:"spec"`
	Status TrustTunnelCloudflaredRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrustTunnelCloudflaredRouteList contains a list of TrustTunnelCloudflaredRoutes
type TrustTunnelCloudflaredRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrustTunnelCloudflaredRoute `json:"items"`
}

// Repository type metadata.
var (
	TrustTunnelCloudflaredRoute_Kind             = "TrustTunnelCloudflaredRoute"
	TrustTunnelCloudflaredRoute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrustTunnelCloudflaredRoute_Kind}.String()
	TrustTunnelCloudflaredRoute_KindAPIVersion   = TrustTunnelCloudflaredRoute_Kind + "." + CRDGroupVersion.String()
	TrustTunnelCloudflaredRoute_GroupVersionKind = CRDGroupVersion.WithKind(TrustTunnelCloudflaredRoute_Kind)
)

func init() {
	SchemeBuilder.Register(&TrustTunnelCloudflaredRoute{}, &TrustTunnelCloudflaredRouteList{})
}
