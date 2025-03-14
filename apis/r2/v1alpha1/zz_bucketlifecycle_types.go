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

type AbortMultipartUploadsTransitionInitParameters struct {

	// (Attributes) Condition for lifecycle transitions to apply after an object reaches an age in seconds (see below for nested schema)
	Condition *ConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`
}

type AbortMultipartUploadsTransitionObservation struct {

	// (Attributes) Condition for lifecycle transitions to apply after an object reaches an age in seconds (see below for nested schema)
	Condition *ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`
}

type AbortMultipartUploadsTransitionParameters struct {

	// (Attributes) Condition for lifecycle transitions to apply after an object reaches an age in seconds (see below for nested schema)
	// +kubebuilder:validation:Optional
	Condition *ConditionParameters `json:"condition" tf:"condition,omitempty"`
}

type BucketLifecycleInitParameters struct {

	// (String) Account ID
	// Account ID
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) Name of the bucket
	// Name of the bucket
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// (Attributes List) (see below for nested schema)
	Rules []BucketLifecycleRulesInitParameters `json:"rules,omitempty" tf:"rules,omitempty"`
}

type BucketLifecycleObservation struct {

	// (String) Account ID
	// Account ID
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) Name of the bucket
	// Name of the bucket
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// (String) Unique identifier for this rule
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Attributes List) (see below for nested schema)
	Rules []BucketLifecycleRulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`
}

type BucketLifecycleParameters struct {

	// (String) Account ID
	// Account ID
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) Name of the bucket
	// Name of the bucket
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// (Attributes List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Rules []BucketLifecycleRulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`
}

type BucketLifecycleRulesInitParameters struct {

	// (Attributes) Transition to abort ongoing multipart uploads (see below for nested schema)
	AbortMultipartUploadsTransition *AbortMultipartUploadsTransitionInitParameters `json:"abortMultipartUploadsTransition,omitempty" tf:"abort_multipart_uploads_transition,omitempty"`

	// (Attributes) Conditions that apply to all transitions of this rule (see below for nested schema)
	Conditions *ConditionsInitParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// (Attributes) Transition to delete objects (see below for nested schema)
	DeleteObjectsTransition *DeleteObjectsTransitionInitParameters `json:"deleteObjectsTransition,omitempty" tf:"delete_objects_transition,omitempty"`

	// (Boolean) Whether or not this rule is in effect
	// Whether or not this rule is in effect
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Unique identifier for this rule
	// Unique identifier for this rule
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Attributes List) Transitions to change the storage class of objects (see below for nested schema)
	StorageClassTransitions []StorageClassTransitionsInitParameters `json:"storageClassTransitions,omitempty" tf:"storage_class_transitions,omitempty"`
}

type BucketLifecycleRulesObservation struct {

	// (Attributes) Transition to abort ongoing multipart uploads (see below for nested schema)
	AbortMultipartUploadsTransition *AbortMultipartUploadsTransitionObservation `json:"abortMultipartUploadsTransition,omitempty" tf:"abort_multipart_uploads_transition,omitempty"`

	// (Attributes) Conditions that apply to all transitions of this rule (see below for nested schema)
	Conditions *ConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// (Attributes) Transition to delete objects (see below for nested schema)
	DeleteObjectsTransition *DeleteObjectsTransitionObservation `json:"deleteObjectsTransition,omitempty" tf:"delete_objects_transition,omitempty"`

	// (Boolean) Whether or not this rule is in effect
	// Whether or not this rule is in effect
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Unique identifier for this rule
	// Unique identifier for this rule
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Attributes List) Transitions to change the storage class of objects (see below for nested schema)
	StorageClassTransitions []StorageClassTransitionsObservation `json:"storageClassTransitions,omitempty" tf:"storage_class_transitions,omitempty"`
}

type BucketLifecycleRulesParameters struct {

	// (Attributes) Transition to abort ongoing multipart uploads (see below for nested schema)
	// +kubebuilder:validation:Optional
	AbortMultipartUploadsTransition *AbortMultipartUploadsTransitionParameters `json:"abortMultipartUploadsTransition,omitempty" tf:"abort_multipart_uploads_transition,omitempty"`

	// (Attributes) Conditions that apply to all transitions of this rule (see below for nested schema)
	// +kubebuilder:validation:Optional
	Conditions *ConditionsParameters `json:"conditions" tf:"conditions,omitempty"`

	// (Attributes) Transition to delete objects (see below for nested schema)
	// +kubebuilder:validation:Optional
	DeleteObjectsTransition *DeleteObjectsTransitionParameters `json:"deleteObjectsTransition,omitempty" tf:"delete_objects_transition,omitempty"`

	// (Boolean) Whether or not this rule is in effect
	// Whether or not this rule is in effect
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// (String) Unique identifier for this rule
	// Unique identifier for this rule
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`

	// (Attributes List) Transitions to change the storage class of objects (see below for nested schema)
	// +kubebuilder:validation:Optional
	StorageClassTransitions []StorageClassTransitionsParameters `json:"storageClassTransitions" tf:"storage_class_transitions,omitempty"`
}

type ConditionInitParameters struct {

	// (Number)
	MaxAge *float64 `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// (String)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConditionObservation struct {

	// (Number)
	MaxAge *float64 `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// (String)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConditionParameters struct {

	// (Number)
	// +kubebuilder:validation:Optional
	MaxAge *float64 `json:"maxAge" tf:"max_age,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type ConditionsInitParameters struct {

	// (String) Transitions will only apply to objects/uploads in the bucket that start with the given prefix, an empty prefix can be provided to scope rule to all objects/uploads
	// Transitions will only apply to objects/uploads in the bucket that start with the given prefix, an empty prefix can be provided to scope rule to all objects/uploads
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type ConditionsObservation struct {

	// (String) Transitions will only apply to objects/uploads in the bucket that start with the given prefix, an empty prefix can be provided to scope rule to all objects/uploads
	// Transitions will only apply to objects/uploads in the bucket that start with the given prefix, an empty prefix can be provided to scope rule to all objects/uploads
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type ConditionsParameters struct {

	// (String) Transitions will only apply to objects/uploads in the bucket that start with the given prefix, an empty prefix can be provided to scope rule to all objects/uploads
	// Transitions will only apply to objects/uploads in the bucket that start with the given prefix, an empty prefix can be provided to scope rule to all objects/uploads
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix" tf:"prefix,omitempty"`
}

type DeleteObjectsTransitionConditionInitParameters struct {

	// (String)
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// (Number)
	MaxAge *float64 `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// (String)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DeleteObjectsTransitionConditionObservation struct {

	// (String)
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// (Number)
	MaxAge *float64 `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// (String)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DeleteObjectsTransitionConditionParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	MaxAge *float64 `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type DeleteObjectsTransitionInitParameters struct {

	// (Attributes) Condition for lifecycle transitions to apply after an object reaches an age in seconds (see below for nested schema)
	Condition *DeleteObjectsTransitionConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`
}

type DeleteObjectsTransitionObservation struct {

	// (Attributes) Condition for lifecycle transitions to apply after an object reaches an age in seconds (see below for nested schema)
	Condition *DeleteObjectsTransitionConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`
}

type DeleteObjectsTransitionParameters struct {

	// (Attributes) Condition for lifecycle transitions to apply after an object reaches an age in seconds (see below for nested schema)
	// +kubebuilder:validation:Optional
	Condition *DeleteObjectsTransitionConditionParameters `json:"condition" tf:"condition,omitempty"`
}

type StorageClassTransitionsConditionInitParameters struct {

	// (String)
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// (Number)
	MaxAge *float64 `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// (String)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StorageClassTransitionsConditionObservation struct {

	// (String)
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// (Number)
	MaxAge *float64 `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// (String)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StorageClassTransitionsConditionParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	MaxAge *float64 `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type StorageClassTransitionsInitParameters struct {

	// (Attributes) Condition for lifecycle transitions to apply after an object reaches an age in seconds (see below for nested schema)
	Condition *StorageClassTransitionsConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// (String)
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type StorageClassTransitionsObservation struct {

	// (Attributes) Condition for lifecycle transitions to apply after an object reaches an age in seconds (see below for nested schema)
	Condition *StorageClassTransitionsConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// (String)
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type StorageClassTransitionsParameters struct {

	// (Attributes) Condition for lifecycle transitions to apply after an object reaches an age in seconds (see below for nested schema)
	// +kubebuilder:validation:Optional
	Condition *StorageClassTransitionsConditionParameters `json:"condition" tf:"condition,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

// BucketLifecycleSpec defines the desired state of BucketLifecycle
type BucketLifecycleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketLifecycleParameters `json:"forProvider"`
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
	InitProvider BucketLifecycleInitParameters `json:"initProvider,omitempty"`
}

// BucketLifecycleStatus defines the observed state of BucketLifecycle.
type BucketLifecycleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketLifecycleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BucketLifecycle is the Schema for the BucketLifecycles API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type BucketLifecycle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accountId) || (has(self.initProvider) && has(self.initProvider.accountId))",message="spec.forProvider.accountId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bucketName) || (has(self.initProvider) && has(self.initProvider.bucketName))",message="spec.forProvider.bucketName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rules) || (has(self.initProvider) && has(self.initProvider.rules))",message="spec.forProvider.rules is a required parameter"
	Spec   BucketLifecycleSpec   `json:"spec"`
	Status BucketLifecycleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketLifecycleList contains a list of BucketLifecycles
type BucketLifecycleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketLifecycle `json:"items"`
}

// Repository type metadata.
var (
	BucketLifecycle_Kind             = "BucketLifecycle"
	BucketLifecycle_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketLifecycle_Kind}.String()
	BucketLifecycle_KindAPIVersion   = BucketLifecycle_Kind + "." + CRDGroupVersion.String()
	BucketLifecycle_GroupVersionKind = CRDGroupVersion.WithKind(BucketLifecycle_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketLifecycle{}, &BucketLifecycleList{})
}
