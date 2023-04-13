/*
 * VM Auto Scaling API
 *
 * The VM Auto Scaling Service enables IONOS clients to horizontally scale the number of VM replicas based on configured rules. You can use Auto Scaling to ensure that you have a sufficient number of replicas to handle your application loads at all times.  For this purpose, create an Auto Scaling group that contains the server replicas. The VM Auto Scaling Service ensures that the number of replicas in the group is always within the defined limits. For example, if the number of target replicas is specified, Auto Scaling maintains the specified number of replicas.   When scaling policies are set, Auto Scaling creates or deletes replicas according to the requirements of your applications. For each policy, specified 'scale-in' and 'scale-out' actions are performed when the corresponding thresholds are reached.
 *
 * API version: 1-SDK.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud_vm_autoscaling

import (
	"encoding/json"
)

// ReplicaVolumeGet struct for ReplicaVolumeGet
type ReplicaVolumeGet struct {
	// The image installed on the disk. Currently, only the UUID of the image is supported.  >Note that either 'image' or 'imageAlias' must be specified, but not both.
	Image *string `json:"image,omitempty"`
	// The image installed on the volume. Must be an 'imageAlias' as specified via the images API. Note that one of 'image' or 'imageAlias' must be set, but not both.
	ImageAlias *string `json:"imageAlias,omitempty"`
	// The replica volume name.
	Name *string `json:"name"`
	// The size of this replica volume in GB.
	Size *int32 `json:"size"`
	// The SSH keys of this volume.
	SshKeys *[]string     `json:"sshKeys,omitempty"`
	Type    *VolumeHwType `json:"type"`
	// The user data (Cloud Init) for this replica volume.
	UserData *string  `json:"userData,omitempty"`
	Bus      *BusType `json:"bus,omitempty"`
	// The ID of the backup unit that the user has access to. The property is immutable and is only allowed to be set on creation of a new a volume. It is mandatory to provide either 'public image' or 'imageAlias' in conjunction with this property.
	BackupunitId *string `json:"backupunitId,omitempty"`
	// Determines whether the volume will be used as a boot volume. Set to NONE, the volume will not be used as boot volume. Set to PRIMARY, the volume will be used as boot volume and set to AUTO will delegate the decision to the provisioning engine to decide whether to use the voluem as boot volume. Notice that exactly one volume can be set to PRIMARY or all of them set to AUTO.
	BootOrder *string `json:"bootOrder"`
}

// NewReplicaVolumeGet instantiates a new ReplicaVolumeGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicaVolumeGet(name string, size int32, type_ VolumeHwType, bootOrder string) *ReplicaVolumeGet {
	this := ReplicaVolumeGet{}

	this.Name = &name

	this.Size = &size

	this.Type = &type_

	var bus BusType = VIRTIO
	this.Bus = &bus

	this.BootOrder = &bootOrder

	return &this
}

// NewReplicaVolumeGetWithDefaults instantiates a new ReplicaVolumeGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicaVolumeGetWithDefaults() *ReplicaVolumeGet {
	this := ReplicaVolumeGet{}
	var bus BusType = VIRTIO
	this.Bus = &bus
	return &this
}

// GetImage returns the Image field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ReplicaVolumeGet) GetImage() *string {
	if o == nil {
		return nil
	}

	return o.Image

}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicaVolumeGet) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Image, true
}

// SetImage sets field value
func (o *ReplicaVolumeGet) SetImage(v string) {

	o.Image = &v

}

// HasImage returns a boolean if a field has been set.
func (o *ReplicaVolumeGet) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// GetImageAlias returns the ImageAlias field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ReplicaVolumeGet) GetImageAlias() *string {
	if o == nil {
		return nil
	}

	return o.ImageAlias

}

// GetImageAliasOk returns a tuple with the ImageAlias field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicaVolumeGet) GetImageAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ImageAlias, true
}

// SetImageAlias sets field value
func (o *ReplicaVolumeGet) SetImageAlias(v string) {

	o.ImageAlias = &v

}

// HasImageAlias returns a boolean if a field has been set.
func (o *ReplicaVolumeGet) HasImageAlias() bool {
	if o != nil && o.ImageAlias != nil {
		return true
	}

	return false
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ReplicaVolumeGet) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicaVolumeGet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *ReplicaVolumeGet) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *ReplicaVolumeGet) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetSize returns the Size field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ReplicaVolumeGet) GetSize() *int32 {
	if o == nil {
		return nil
	}

	return o.Size

}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicaVolumeGet) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Size, true
}

// SetSize sets field value
func (o *ReplicaVolumeGet) SetSize(v int32) {

	o.Size = &v

}

// HasSize returns a boolean if a field has been set.
func (o *ReplicaVolumeGet) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// GetSshKeys returns the SshKeys field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ReplicaVolumeGet) GetSshKeys() *[]string {
	if o == nil {
		return nil
	}

	return o.SshKeys

}

// GetSshKeysOk returns a tuple with the SshKeys field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicaVolumeGet) GetSshKeysOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.SshKeys, true
}

// SetSshKeys sets field value
func (o *ReplicaVolumeGet) SetSshKeys(v []string) {

	o.SshKeys = &v

}

// HasSshKeys returns a boolean if a field has been set.
func (o *ReplicaVolumeGet) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for VolumeHwType will be returned
func (o *ReplicaVolumeGet) GetType() *VolumeHwType {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicaVolumeGet) GetTypeOk() (*VolumeHwType, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *ReplicaVolumeGet) SetType(v VolumeHwType) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *ReplicaVolumeGet) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetUserData returns the UserData field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ReplicaVolumeGet) GetUserData() *string {
	if o == nil {
		return nil
	}

	return o.UserData

}

// GetUserDataOk returns a tuple with the UserData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicaVolumeGet) GetUserDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.UserData, true
}

// SetUserData sets field value
func (o *ReplicaVolumeGet) SetUserData(v string) {

	o.UserData = &v

}

// HasUserData returns a boolean if a field has been set.
func (o *ReplicaVolumeGet) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

// GetBus returns the Bus field value
// If the value is explicit nil, the zero value for BusType will be returned
func (o *ReplicaVolumeGet) GetBus() *BusType {
	if o == nil {
		return nil
	}

	return o.Bus

}

// GetBusOk returns a tuple with the Bus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicaVolumeGet) GetBusOk() (*BusType, bool) {
	if o == nil {
		return nil, false
	}

	return o.Bus, true
}

// SetBus sets field value
func (o *ReplicaVolumeGet) SetBus(v BusType) {

	o.Bus = &v

}

// HasBus returns a boolean if a field has been set.
func (o *ReplicaVolumeGet) HasBus() bool {
	if o != nil && o.Bus != nil {
		return true
	}

	return false
}

// GetBackupunitId returns the BackupunitId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ReplicaVolumeGet) GetBackupunitId() *string {
	if o == nil {
		return nil
	}

	return o.BackupunitId

}

// GetBackupunitIdOk returns a tuple with the BackupunitId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicaVolumeGet) GetBackupunitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.BackupunitId, true
}

// SetBackupunitId sets field value
func (o *ReplicaVolumeGet) SetBackupunitId(v string) {

	o.BackupunitId = &v

}

// HasBackupunitId returns a boolean if a field has been set.
func (o *ReplicaVolumeGet) HasBackupunitId() bool {
	if o != nil && o.BackupunitId != nil {
		return true
	}

	return false
}

// GetBootOrder returns the BootOrder field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ReplicaVolumeGet) GetBootOrder() *string {
	if o == nil {
		return nil
	}

	return o.BootOrder

}

// GetBootOrderOk returns a tuple with the BootOrder field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicaVolumeGet) GetBootOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.BootOrder, true
}

// SetBootOrder sets field value
func (o *ReplicaVolumeGet) SetBootOrder(v string) {

	o.BootOrder = &v

}

// HasBootOrder returns a boolean if a field has been set.
func (o *ReplicaVolumeGet) HasBootOrder() bool {
	if o != nil && o.BootOrder != nil {
		return true
	}

	return false
}

func (o ReplicaVolumeGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image

	toSerialize["imageAlias"] = o.ImageAlias

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	if o.SshKeys != nil {
		toSerialize["sshKeys"] = o.SshKeys
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.UserData != nil {
		toSerialize["userData"] = o.UserData
	}

	if o.Bus != nil {
		toSerialize["bus"] = o.Bus
	}

	if o.BackupunitId != nil {
		toSerialize["backupunitId"] = o.BackupunitId
	}

	if o.BootOrder != nil {
		toSerialize["bootOrder"] = o.BootOrder
	}

	return json.Marshal(toSerialize)
}

type NullableReplicaVolumeGet struct {
	value *ReplicaVolumeGet
	isSet bool
}

func (v NullableReplicaVolumeGet) Get() *ReplicaVolumeGet {
	return v.value
}

func (v *NullableReplicaVolumeGet) Set(val *ReplicaVolumeGet) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicaVolumeGet) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicaVolumeGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicaVolumeGet(val *ReplicaVolumeGet) *NullableReplicaVolumeGet {
	return &NullableReplicaVolumeGet{value: val, isSet: true}
}

func (v NullableReplicaVolumeGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicaVolumeGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
