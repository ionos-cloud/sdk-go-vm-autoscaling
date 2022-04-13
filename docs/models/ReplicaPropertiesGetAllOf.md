# ReplicaPropertiesGetAllOf



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Volumes** | Pointer to [**[]ReplicaVolumeGet**](ReplicaVolumeGet.md) | List of volumes associated with this Replica. Only a single volume is currently supported. | [optional] |

## Methods


### GetVolumes

`func (o *ReplicaPropertiesGetAllOf) GetVolumes() []ReplicaVolumeGet`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ReplicaPropertiesGetAllOf) GetVolumesOk() (*[]ReplicaVolumeGet, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ReplicaPropertiesGetAllOf) SetVolumes(v []ReplicaVolumeGet)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ReplicaPropertiesGetAllOf) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.



