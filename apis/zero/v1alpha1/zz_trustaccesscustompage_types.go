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

type TrustAccessCustomPageInitParameters struct {

	// (String) Identifier
	// Identifier
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Number) Number of apps the custom page is assigned to.
	// Number of apps the custom page is assigned to.
	AppCount *float64 `json:"appCount,omitempty" tf:"app_count,omitempty"`

	// (String) Custom page HTML.
	// Custom page HTML.
	CustomHTML *string `json:"customHtml,omitempty" tf:"custom_html,omitempty"`

	// (String) Custom page type.
	// Custom page type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TrustAccessCustomPageObservation struct {

	// (String) Identifier
	// Identifier
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Number) Number of apps the custom page is assigned to.
	// Number of apps the custom page is assigned to.
	AppCount *float64 `json:"appCount,omitempty" tf:"app_count,omitempty"`

	// (String)
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (String) Custom page HTML.
	// Custom page HTML.
	CustomHTML *string `json:"customHtml,omitempty" tf:"custom_html,omitempty"`

	// (String) UUID
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Custom page type.
	// Custom page type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) UUID
	// UUID
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// (String)
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type TrustAccessCustomPageParameters struct {

	// (String) Identifier
	// Identifier
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Number) Number of apps the custom page is assigned to.
	// Number of apps the custom page is assigned to.
	// +kubebuilder:validation:Optional
	AppCount *float64 `json:"appCount,omitempty" tf:"app_count,omitempty"`

	// (String) Custom page HTML.
	// Custom page HTML.
	// +kubebuilder:validation:Optional
	CustomHTML *string `json:"customHtml,omitempty" tf:"custom_html,omitempty"`

	// (String) Custom page type.
	// Custom page type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// TrustAccessCustomPageSpec defines the desired state of TrustAccessCustomPage
type TrustAccessCustomPageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrustAccessCustomPageParameters `json:"forProvider"`
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
	InitProvider TrustAccessCustomPageInitParameters `json:"initProvider,omitempty"`
}

// TrustAccessCustomPageStatus defines the observed state of TrustAccessCustomPage.
type TrustAccessCustomPageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrustAccessCustomPageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TrustAccessCustomPage is the Schema for the TrustAccessCustomPages API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type TrustAccessCustomPage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accountId) || (has(self.initProvider) && has(self.initProvider.accountId))",message="spec.forProvider.accountId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.customHtml) || (has(self.initProvider) && has(self.initProvider.customHtml))",message="spec.forProvider.customHtml is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   TrustAccessCustomPageSpec   `json:"spec"`
	Status TrustAccessCustomPageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrustAccessCustomPageList contains a list of TrustAccessCustomPages
type TrustAccessCustomPageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrustAccessCustomPage `json:"items"`
}

// Repository type metadata.
var (
	TrustAccessCustomPage_Kind             = "TrustAccessCustomPage"
	TrustAccessCustomPage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrustAccessCustomPage_Kind}.String()
	TrustAccessCustomPage_KindAPIVersion   = TrustAccessCustomPage_Kind + "." + CRDGroupVersion.String()
	TrustAccessCustomPage_GroupVersionKind = CRDGroupVersion.WithKind(TrustAccessCustomPage_Kind)
)

func init() {
	SchemeBuilder.Register(&TrustAccessCustomPage{}, &TrustAccessCustomPageList{})
}
