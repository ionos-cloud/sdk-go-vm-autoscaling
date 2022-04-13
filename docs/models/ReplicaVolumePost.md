# ReplicaVolumePost



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Image** | **string** | The image installed on the volume. Only the UUID of the image is presently supported. | |
|**Name** | **string** | Name of the replica volume. | |
|**Size** | **int32** | User-defined size for this replica volume in GB. | |
|**SshKeys** | Pointer to **[]string** | Ssh keys that has access to the volume. | [optional] |
|**Type** | [**VolumeHwType**](VolumeHwType.md) |  | |
|**UserData** | Pointer to **string** | user-data (Cloud Init) for this replica volume. | [optional] |
|**Bus** | Pointer to [**BusType**](BusType.md) |  | [optional] |
|**ImagePassword** | Pointer to **string** | Image password for this replica volume. | [optional] |

## Methods


### GetImage

`func (o *ReplicaVolumePost) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ReplicaVolumePost) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ReplicaVolumePost) SetImage(v string)`

SetImage sets Image field to given value.


### GetName

`func (o *ReplicaVolumePost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReplicaVolumePost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReplicaVolumePost) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *ReplicaVolumePost) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ReplicaVolumePost) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ReplicaVolumePost) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSshKeys

`func (o *ReplicaVolumePost) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *ReplicaVolumePost) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *ReplicaVolumePost) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *ReplicaVolumePost) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetType

`func (o *ReplicaVolumePost) GetType() VolumeHwType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReplicaVolumePost) GetTypeOk() (*VolumeHwType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReplicaVolumePost) SetType(v VolumeHwType)`

SetType sets Type field to given value.


### GetUserData

`func (o *ReplicaVolumePost) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *ReplicaVolumePost) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *ReplicaVolumePost) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *ReplicaVolumePost) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetBus

`func (o *ReplicaVolumePost) GetBus() BusType`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *ReplicaVolumePost) GetBusOk() (*BusType, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *ReplicaVolumePost) SetBus(v BusType)`

SetBus sets Bus field to given value.

### HasBus

`func (o *ReplicaVolumePost) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetImagePassword

`func (o *ReplicaVolumePost) GetImagePassword() string`

GetImagePassword returns the ImagePassword field if non-nil, zero value otherwise.

### GetImagePasswordOk

`func (o *ReplicaVolumePost) GetImagePasswordOk() (*string, bool)`

GetImagePasswordOk returns a tuple with the ImagePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePassword

`func (o *ReplicaVolumePost) SetImagePassword(v string)`

SetImagePassword sets ImagePassword field to given value.

### HasImagePassword

`func (o *ReplicaVolumePost) HasImagePassword() bool`

HasImagePassword returns a boolean if a field has been set.



