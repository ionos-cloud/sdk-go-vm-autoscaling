# ReplicaProperties



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) |  | |
|**Cores** | **int32** | The total number of cores for the VMs. | |
|**CpuFamily** | Pointer to [**CpuFamily**](CpuFamily.md) |  | [optional] |
|**Nics** | Pointer to [**[]ReplicaNic**](ReplicaNic.md) | List of NICs associated with this Replica. | [optional] |
|**Ram** | **int32** | The amount of memory for the VMs in MB, e.g. 2048. Size must be specified in multiples of 256 MB with a minimum of 256 MB; however, if you set ramHotPlug to TRUE then you must use a minimum of 1024 MB. If you set the RAM size more than 240GB, then ramHotPlug will be set to FALSE and can not be set to TRUE unless RAM size not set to less than 240GB. | |

## Methods


### GetAvailabilityZone

`func (o *ReplicaProperties) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ReplicaProperties) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ReplicaProperties) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetCores

`func (o *ReplicaProperties) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *ReplicaProperties) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *ReplicaProperties) SetCores(v int32)`

SetCores sets Cores field to given value.


### GetCpuFamily

`func (o *ReplicaProperties) GetCpuFamily() CpuFamily`

GetCpuFamily returns the CpuFamily field if non-nil, zero value otherwise.

### GetCpuFamilyOk

`func (o *ReplicaProperties) GetCpuFamilyOk() (*CpuFamily, bool)`

GetCpuFamilyOk returns a tuple with the CpuFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFamily

`func (o *ReplicaProperties) SetCpuFamily(v CpuFamily)`

SetCpuFamily sets CpuFamily field to given value.

### HasCpuFamily

`func (o *ReplicaProperties) HasCpuFamily() bool`

HasCpuFamily returns a boolean if a field has been set.

### GetNics

`func (o *ReplicaProperties) GetNics() []ReplicaNic`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *ReplicaProperties) GetNicsOk() (*[]ReplicaNic, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *ReplicaProperties) SetNics(v []ReplicaNic)`

SetNics sets Nics field to given value.

### HasNics

`func (o *ReplicaProperties) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetRam

`func (o *ReplicaProperties) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ReplicaProperties) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ReplicaProperties) SetRam(v int32)`

SetRam sets Ram field to given value.




