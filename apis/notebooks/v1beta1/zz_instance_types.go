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

type AcceleratorConfigObservation struct {
}

type AcceleratorConfigParameters struct {

	// Count of cores of this accelerator.
	// +kubebuilder:validation:Required
	CoreCount *float64 `json:"coreCount" tf:"core_count,omitempty"`

	// Type of this accelerator.
	// Possible values are ACCELERATOR_TYPE_UNSPECIFIED, NVIDIA_TESLA_K80, NVIDIA_TESLA_P100, NVIDIA_TESLA_V100, NVIDIA_TESLA_P4, NVIDIA_TESLA_T4, NVIDIA_TESLA_T4_VWS, NVIDIA_TESLA_P100_VWS, NVIDIA_TESLA_P4_VWS, NVIDIA_TESLA_A100, TPU_V2, and TPU_V3.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type InstanceContainerImageObservation struct {
}

type InstanceContainerImageParameters struct {

	// The path to the container image repository.
	// For example: gcr.io/{project_id}/{imageName}
	// +kubebuilder:validation:Required
	Repository *string `json:"repository" tf:"repository,omitempty"`

	// The tag of the container image. If not specified, this defaults to the latest tag.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type InstanceObservation struct {

	// Instance creation time
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/instances/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The proxy endpoint that is used to access the Jupyter notebook.
	ProxyURI *string `json:"proxyUri,omitempty" tf:"proxy_uri,omitempty"`

	// The state of this instance.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Instance update time.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type InstanceParameters struct {

	// The hardware accelerator used on this instance. If you use accelerators,
	// make sure that your configuration has enough vCPUs and memory to support the
	// machineType you have selected.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AcceleratorConfig []AcceleratorConfigParameters `json:"acceleratorConfig,omitempty" tf:"accelerator_config,omitempty"`

	// The size of the boot disk in GB attached to this instance,
	// up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB.
	// If not specified, this defaults to 100.
	// +kubebuilder:validation:Optional
	BootDiskSizeGb *float64 `json:"bootDiskSizeGb,omitempty" tf:"boot_disk_size_gb,omitempty"`

	// Possible disk types for notebook instances.
	// Possible values are DISK_TYPE_UNSPECIFIED, PD_STANDARD, PD_SSD, PD_BALANCED, and PD_EXTREME.
	// +kubebuilder:validation:Optional
	BootDiskType *string `json:"bootDiskType,omitempty" tf:"boot_disk_type,omitempty"`

	// Use a container image to start the notebook instance.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ContainerImage []InstanceContainerImageParameters `json:"containerImage,omitempty" tf:"container_image,omitempty"`

	// Specify a custom Cloud Storage path where the GPU driver is stored.
	// If not specified, we'll automatically choose from official GPU drivers.
	// +kubebuilder:validation:Optional
	CustomGpuDriverPath *string `json:"customGpuDriverPath,omitempty" tf:"custom_gpu_driver_path,omitempty"`

	// The size of the data disk in GB attached to this instance,
	// up to a maximum of 64000 GB (64 TB).
	// You can choose the size of the data disk based on how big your notebooks and data are.
	// If not specified, this defaults to 100.
	// +kubebuilder:validation:Optional
	DataDiskSizeGb *float64 `json:"dataDiskSizeGb,omitempty" tf:"data_disk_size_gb,omitempty"`

	// Possible disk types for notebook instances.
	// Possible values are DISK_TYPE_UNSPECIFIED, PD_STANDARD, PD_SSD, PD_BALANCED, and PD_EXTREME.
	// +kubebuilder:validation:Optional
	DataDiskType *string `json:"dataDiskType,omitempty" tf:"data_disk_type,omitempty"`

	// Disk encryption method used on the boot and data disks, defaults to GMEK.
	// Possible values are DISK_ENCRYPTION_UNSPECIFIED, GMEK, and CMEK.
	// +kubebuilder:validation:Optional
	DiskEncryption *string `json:"diskEncryption,omitempty" tf:"disk_encryption,omitempty"`

	// Whether the end user authorizes Google Cloud to install GPU driver
	// on this instance. If this field is empty or set to false, the GPU driver
	// won't be installed. Only applicable to instances with GPUs.
	// +kubebuilder:validation:Optional
	InstallGpuDriver *bool `json:"installGpuDriver,omitempty" tf:"install_gpu_driver,omitempty"`

	// The list of owners of this instance after creation.
	// Format: alias@example.com.
	// Currently supports one owner only.
	// If not specified, all of the service account users of
	// your VM instance's service account can use the instance.
	// +kubebuilder:validation:Optional
	InstanceOwners []*string `json:"instanceOwners,omitempty" tf:"instance_owners,omitempty"`

	// The KMS key used to encrypt the disks, only applicable if diskEncryption is CMEK.
	// Format: projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}
	// +kubebuilder:validation:Optional
	KMSKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`

	// Labels to apply to this instance. These can be later modified by the setLabels method.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A reference to the zone where the machine resides.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// A reference to a machine type which defines VM kind.
	// +kubebuilder:validation:Required
	MachineType *string `json:"machineType" tf:"machine_type,omitempty"`

	// Custom metadata to apply to this instance.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the VPC that this instance is in.
	// Format: projects/{project_id}/global/networks/{network_id}
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The type of vNIC driver.
	// Possible values are UNSPECIFIED_NIC_TYPE, VIRTIO_NET, and GVNIC.
	// +kubebuilder:validation:Optional
	NicType *string `json:"nicType,omitempty" tf:"nic_type,omitempty"`

	// The notebook instance will not register with the proxy..
	// +kubebuilder:validation:Optional
	NoProxyAccess *bool `json:"noProxyAccess,omitempty" tf:"no_proxy_access,omitempty"`

	// No public IP will be assigned to this instance.
	// +kubebuilder:validation:Optional
	NoPublicIP *bool `json:"noPublicIp,omitempty" tf:"no_public_ip,omitempty"`

	// If true, the data disk will not be auto deleted when deleting the instance.
	// +kubebuilder:validation:Optional
	NoRemoveDataDisk *bool `json:"noRemoveDataDisk,omitempty" tf:"no_remove_data_disk,omitempty"`

	// Path to a Bash script that automatically runs after a
	// notebook instance fully boots up. The path must be a URL
	// or Cloud Storage path (gs://path-to-file/file-name).
	// +kubebuilder:validation:Optional
	PostStartupScript *string `json:"postStartupScript,omitempty" tf:"post_startup_script,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reservation Affinity for consuming Zonal reservation.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ReservationAffinity []ReservationAffinityParameters `json:"reservationAffinity,omitempty" tf:"reservation_affinity,omitempty"`

	// The service account on this instance, giving access to other
	// Google Cloud services. You can use any service account within
	// the same project, but you must have the service account user
	// permission to use the instance. If not specified,
	// the Compute Engine default service account is used.
	// +kubebuilder:validation:Optional
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// Optional. The URIs of service account scopes to be included in Compute Engine instances.
	// If not specified, the following scopes are defined:
	// +kubebuilder:validation:Optional
	ServiceAccountScopes []*string `json:"serviceAccountScopes,omitempty" tf:"service_account_scopes,omitempty"`

	// A set of Shielded Instance options. Check [Images using supported Shielded VM features]
	// Not all combinations are valid
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ShieldedInstanceConfig []ShieldedInstanceConfigParameters `json:"shieldedInstanceConfig,omitempty" tf:"shielded_instance_config,omitempty"`

	// The name of the subnet that this instance is in.
	// Format: projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}
	// +kubebuilder:validation:Optional
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// The Compute Engine tags to add to instance.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	VMImage []InstanceVMImageParameters `json:"vmImage,omitempty" tf:"vm_image,omitempty"`
}

type InstanceVMImageObservation struct {
}

type InstanceVMImageParameters struct {

	// Use this VM image family to find the image; the newest image in this family will be used.
	// +kubebuilder:validation:Optional
	ImageFamily *string `json:"imageFamily,omitempty" tf:"image_family,omitempty"`

	// Use VM image name to find the image.
	// +kubebuilder:validation:Optional
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`
}

type ReservationAffinityObservation struct {
}

type ReservationAffinityParameters struct {

	// The type of Compute Reservation.
	// Possible values are NO_RESERVATION, ANY_RESERVATION, and SPECIFIC_RESERVATION.
	// +kubebuilder:validation:Required
	ConsumeReservationType *string `json:"consumeReservationType" tf:"consume_reservation_type,omitempty"`

	// Corresponds to the label key of reservation resource.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Corresponds to the label values of reservation resource.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ShieldedInstanceConfigObservation struct {
}

type ShieldedInstanceConfigParameters struct {

	// Defines whether the instance has integrity monitoring enabled. Enables monitoring and attestation of the
	// boot integrity of the instance. The attestation is performed against the integrity policy baseline.
	// This baseline is initially derived from the implicitly trusted boot image when the instance is created.
	// Enabled by default.
	// +kubebuilder:validation:Optional
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty" tf:"enable_integrity_monitoring,omitempty"`

	// Defines whether the instance has Secure Boot enabled. Secure Boot helps ensure that the system only runs
	// authentic software by verifying the digital signature of all boot components, and halting the boot process
	// if signature verification fails.
	// Disabled by default.
	// +kubebuilder:validation:Optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty" tf:"enable_secure_boot,omitempty"`

	// Defines whether the instance has the vTPM enabled.
	// Enabled by default.
	// +kubebuilder:validation:Optional
	EnableVtpm *bool `json:"enableVtpm,omitempty" tf:"enable_vtpm,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. A Cloud AI Platform Notebook instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
