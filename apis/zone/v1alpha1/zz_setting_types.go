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

type SettingInitParameters struct {

	// (String) Setting name
	// Setting name
	SettingID *string `json:"settingId,omitempty" tf:"setting_id,omitempty"`

	// (Dynamic) Current value of the zone setting.
	// Current value of the zone setting.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// (String) Identifier
	// Identifier
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type SettingObservation struct {

	// (Boolean) Whether or not this setting can be modified for this zone (based on your Cloudflare plan level).
	// Whether or not this setting can be modified for this zone (based on your Cloudflare plan level).
	Editable *bool `json:"editable,omitempty" tf:"editable,omitempty"`

	// (String) ID of the zone setting.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) last time this setting was modified.
	// last time this setting was modified.
	ModifiedOn *string `json:"modifiedOn,omitempty" tf:"modified_on,omitempty"`

	// (String) Setting name
	// Setting name
	SettingID *string `json:"settingId,omitempty" tf:"setting_id,omitempty"`

	// (Number) Value of the zone setting.
	// Notes: The interval (in seconds) from when development mode expires (positive integer) or last expired (negative integer) for the domain. If development mode has never been enabled, this value is false.
	// Value of the zone setting.
	// Notes: The interval (in seconds) from when development mode expires (positive integer) or last expired (negative integer) for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining *float64 `json:"timeRemaining,omitempty" tf:"time_remaining,omitempty"`

	// (Dynamic) Current value of the zone setting.
	// Current value of the zone setting.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// (String) Identifier
	// Identifier
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type SettingParameters struct {

	// (String) Setting name
	// Setting name
	// +kubebuilder:validation:Optional
	SettingID *string `json:"settingId,omitempty" tf:"setting_id,omitempty"`

	// (Dynamic) Current value of the zone setting.
	// Current value of the zone setting.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// (String) Identifier
	// Identifier
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

// SettingSpec defines the desired state of Setting
type SettingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SettingParameters `json:"forProvider"`
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
	InitProvider SettingInitParameters `json:"initProvider,omitempty"`
}

// SettingStatus defines the observed state of Setting.
type SettingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SettingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Setting is the Schema for the Settings API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Setting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.settingId) || (has(self.initProvider) && has(self.initProvider.settingId))",message="spec.forProvider.settingId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || (has(self.initProvider) && has(self.initProvider.value))",message="spec.forProvider.value is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zoneId) || (has(self.initProvider) && has(self.initProvider.zoneId))",message="spec.forProvider.zoneId is a required parameter"
	Spec   SettingSpec   `json:"spec"`
	Status SettingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SettingList contains a list of Settings
type SettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Setting `json:"items"`
}

// Repository type metadata.
var (
	Setting_Kind             = "Setting"
	Setting_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Setting_Kind}.String()
	Setting_KindAPIVersion   = Setting_Kind + "." + CRDGroupVersion.String()
	Setting_GroupVersionKind = CRDGroupVersion.WithKind(Setting_Kind)
)

func init() {
	SchemeBuilder.Register(&Setting{}, &SettingList{})
}
