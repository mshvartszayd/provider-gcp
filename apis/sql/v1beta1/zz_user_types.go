/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PasswordPolicyObservation struct {
	Status []StatusObservation `json:"status,omitempty" tf:"status,omitempty"`
}

type PasswordPolicyParameters struct {

	// Number of failed attempts allowed before the user get locked.
	// +kubebuilder:validation:Optional
	AllowedFailedAttempts *float64 `json:"allowedFailedAttempts,omitempty" tf:"allowed_failed_attempts,omitempty"`

	// If true, the check that will lock user after too many failed login attempts will be enabled.
	// +kubebuilder:validation:Optional
	EnableFailedAttemptsCheck *bool `json:"enableFailedAttemptsCheck,omitempty" tf:"enable_failed_attempts_check,omitempty"`

	// If true, the user must specify the current password before changing the password. This flag is supported only for MySQL.
	// +kubebuilder:validation:Optional
	EnablePasswordVerification *bool `json:"enablePasswordVerification,omitempty" tf:"enable_password_verification,omitempty"`

	// Password expiration duration with one week grace period.
	// +kubebuilder:validation:Optional
	PasswordExpirationDuration *string `json:"passwordExpirationDuration,omitempty" tf:"password_expiration_duration,omitempty"`
}

type SQLServerUserDetailsObservation struct {
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	ServerRoles []*string `json:"serverRoles,omitempty" tf:"server_roles,omitempty"`
}

type SQLServerUserDetailsParameters struct {
}

type StatusObservation struct {

	// (read only) If true, user does not have login privileges.
	Locked *bool `json:"locked,omitempty" tf:"locked,omitempty"`

	// (read only) Password expiration duration with one week grace period.
	PasswordExpirationTime *string `json:"passwordExpirationTime,omitempty" tf:"password_expiration_time,omitempty"`
}

type StatusParameters struct {
}

type UserObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordPolicy []PasswordPolicyObservation `json:"passwordPolicy,omitempty" tf:"password_policy,omitempty"`

	SQLServerUserDetails []SQLServerUserDetailsObservation `json:"sqlServerUserDetails,omitempty" tf:"sql_server_user_details,omitempty"`
}

type UserParameters struct {

	// The deletion policy for the user.
	// Setting ABANDON allows the resource to be abandoned rather than deleted. This is useful
	// for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
	// +kubebuilder:validation:Optional
	DeletionPolicy *string `json:"deletionPolicy,omitempty" tf:"deletion_policy,omitempty"`

	// The host the user can connect from. This is only supported
	// for MySQL instances. Don't set this field for PostgreSQL instances.
	// Can be an IP address. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	// +crossplane:generate:reference:type=DatabaseInstance
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// Reference to a DatabaseInstance to populate instance.
	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// Selector for a DatabaseInstance to populate instance.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PasswordPolicy []PasswordPolicyParameters `json:"passwordPolicy,omitempty" tf:"password_policy,omitempty"`

	// The password for the user. Can be updated. For Postgres
	// instances this is a Required field, unless type is set to either CLOUD_IAM_USER
	// or CLOUD_IAM_SERVICE_ACCOUNT.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The user type. It determines the method to authenticate the
	// user during login. The default is the database's built-in user type. Flags
	// include "BUILT_IN", "CLOUD_IAM_USER", or "CLOUD_IAM_SERVICE_ACCOUNT".
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// User is the Schema for the Users API. Creates a new SQL user in Google Cloud SQL.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserSpec   `json:"spec"`
	Status            UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
