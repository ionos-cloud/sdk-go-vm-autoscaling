# ReplicaPropertiesPostAllOf



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Volumes** | Pointer to [**[]ReplicaVolumePost**](ReplicaVolumePost.md) | List of volumes associated with this Replica. Only a single volume is currently supported. | [optional] |

## Methods


### GetVolumes

`func (o *ReplicaPropertiesPostAllOf) GetVolumes() []ReplicaVolumePost`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ReplicaPropertiesPostAllOf) GetVolumesOk() (*[]ReplicaVolumePost, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ReplicaPropertiesPostAllOf) SetVolumes(v []ReplicaVolumePost)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ReplicaPropertiesPostAllOf) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.



