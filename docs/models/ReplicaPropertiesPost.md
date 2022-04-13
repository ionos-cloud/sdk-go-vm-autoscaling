# ReplicaPropertiesPost



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) |  | |
|**Cores** | **int32** | The total number of cores for the VMs. | |
|**CpuFamily** | Pointer to [**CpuFamily**](CpuFamily.md) |  | [optional] |
|**Nics** | Pointer to [**[]ReplicaNic**](ReplicaNic.md) | List of NICs associated with this Replica. | [optional] |
|**Ram** | **int32** | The amount of memory for the VMs in MB, e.g. 2048. Size must be specified in multiples of 256 MB with a minimum of 256 MB; however, if you set ramHotPlug to TRUE then you must use a minimum of 1024 MB. If you set the RAM size more than 240GB, then ramHotPlug will be set to FALSE and can not be set to TRUE unless RAM size not set to less than 240GB. | |
|**Volumes** | Pointer to [**[]ReplicaVolumePost**](ReplicaVolumePost.md) | List of volumes associated with this Replica. Only a single volume is currently supported. | [optional] |

## Methods


### GetAvailabilityZone

`func (o *ReplicaPropertiesPost) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ReplicaPropertiesPost) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ReplicaPropertiesPost) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetCores

`func (o *ReplicaPropertiesPost) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *ReplicaPropertiesPost) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *ReplicaPropertiesPost) SetCores(v int32)`

SetCores sets Cores field to given value.


### GetCpuFamily

`func (o *ReplicaPropertiesPost) GetCpuFamily() CpuFamily`

GetCpuFamily returns the CpuFamily field if non-nil, zero value otherwise.

### GetCpuFamilyOk

`func (o *ReplicaPropertiesPost) GetCpuFamilyOk() (*CpuFamily, bool)`

GetCpuFamilyOk returns a tuple with the CpuFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFamily

`func (o *ReplicaPropertiesPost) SetCpuFamily(v CpuFamily)`

SetCpuFamily sets CpuFamily field to given value.

### HasCpuFamily

`func (o *ReplicaPropertiesPost) HasCpuFamily() bool`

HasCpuFamily returns a boolean if a field has been set.

### GetNics

`func (o *ReplicaPropertiesPost) GetNics() []ReplicaNic`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *ReplicaPropertiesPost) GetNicsOk() (*[]ReplicaNic, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *ReplicaPropertiesPost) SetNics(v []ReplicaNic)`

SetNics sets Nics field to given value.

### HasNics

`func (o *ReplicaPropertiesPost) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetRam

`func (o *ReplicaPropertiesPost) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ReplicaPropertiesPost) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ReplicaPropertiesPost) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetVolumes

`func (o *ReplicaPropertiesPost) GetVolumes() []ReplicaVolumePost`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ReplicaPropertiesPost) GetVolumesOk() (*[]ReplicaVolumePost, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ReplicaPropertiesPost) SetVolumes(v []ReplicaVolumePost)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ReplicaPropertiesPost) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.



