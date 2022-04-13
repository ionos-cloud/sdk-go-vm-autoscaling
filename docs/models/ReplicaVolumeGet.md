# ReplicaVolumeGet



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

## Methods


### GetImage

`func (o *ReplicaVolumeGet) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ReplicaVolumeGet) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ReplicaVolumeGet) SetImage(v string)`

SetImage sets Image field to given value.


### GetName

`func (o *ReplicaVolumeGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReplicaVolumeGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReplicaVolumeGet) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *ReplicaVolumeGet) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ReplicaVolumeGet) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ReplicaVolumeGet) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSshKeys

`func (o *ReplicaVolumeGet) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *ReplicaVolumeGet) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *ReplicaVolumeGet) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *ReplicaVolumeGet) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetType

`func (o *ReplicaVolumeGet) GetType() VolumeHwType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReplicaVolumeGet) GetTypeOk() (*VolumeHwType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReplicaVolumeGet) SetType(v VolumeHwType)`

SetType sets Type field to given value.


### GetUserData

`func (o *ReplicaVolumeGet) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *ReplicaVolumeGet) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *ReplicaVolumeGet) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *ReplicaVolumeGet) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetBus

`func (o *ReplicaVolumeGet) GetBus() BusType`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *ReplicaVolumeGet) GetBusOk() (*BusType, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *ReplicaVolumeGet) SetBus(v BusType)`

SetBus sets Bus field to given value.

### HasBus

`func (o *ReplicaVolumeGet) HasBus() bool`

HasBus returns a boolean if a field has been set.



