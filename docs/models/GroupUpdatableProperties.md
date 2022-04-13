# GroupUpdatableProperties



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**MaxReplicaCount** | Pointer to **int64** | Maximum replica count value for &#x60;targetReplicaCount&#x60;. Will be enforced for both automatic and manual changes. | [optional] |
|**MinReplicaCount** | Pointer to **int64** | Minimum replica count value for &#x60;targetReplicaCount&#x60;. Will be enforced for both automatic and manual changes. | [optional] |
|**TargetReplicaCount** | Pointer to **int64** | The target number of VMs in this Group. Depending on the scaling policy, this number will be adjusted automatically. VMs will be created or destroyed automatically in order to adjust the actual number of VMs to this number. If targetReplicaCount is given in the request body then it must be &gt;&#x3D; minReplicaCount and &lt;&#x3D; maxReplicaCount. | [optional] |
|**Name** | Pointer to **string** | User-defined name for the autoscaling group. | [optional] |
|**Policy** | Pointer to [**GroupPolicy**](GroupPolicy.md) |  | [optional] |
|**ReplicaConfiguration** | Pointer to [**ReplicaPropertiesPost**](ReplicaPropertiesPost.md) |  | [optional] |
|**Datacenter** | Pointer to [**Resource**](Resource.md) |  | [optional] |

## Methods


### GetMaxReplicaCount

`func (o *GroupUpdatableProperties) GetMaxReplicaCount() int64`

GetMaxReplicaCount returns the MaxReplicaCount field if non-nil, zero value otherwise.

### GetMaxReplicaCountOk

`func (o *GroupUpdatableProperties) GetMaxReplicaCountOk() (*int64, bool)`

GetMaxReplicaCountOk returns a tuple with the MaxReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicaCount

`func (o *GroupUpdatableProperties) SetMaxReplicaCount(v int64)`

SetMaxReplicaCount sets MaxReplicaCount field to given value.

### HasMaxReplicaCount

`func (o *GroupUpdatableProperties) HasMaxReplicaCount() bool`

HasMaxReplicaCount returns a boolean if a field has been set.

### GetMinReplicaCount

`func (o *GroupUpdatableProperties) GetMinReplicaCount() int64`

GetMinReplicaCount returns the MinReplicaCount field if non-nil, zero value otherwise.

### GetMinReplicaCountOk

`func (o *GroupUpdatableProperties) GetMinReplicaCountOk() (*int64, bool)`

GetMinReplicaCountOk returns a tuple with the MinReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicaCount

`func (o *GroupUpdatableProperties) SetMinReplicaCount(v int64)`

SetMinReplicaCount sets MinReplicaCount field to given value.

### HasMinReplicaCount

`func (o *GroupUpdatableProperties) HasMinReplicaCount() bool`

HasMinReplicaCount returns a boolean if a field has been set.

### GetTargetReplicaCount

`func (o *GroupUpdatableProperties) GetTargetReplicaCount() int64`

GetTargetReplicaCount returns the TargetReplicaCount field if non-nil, zero value otherwise.

### GetTargetReplicaCountOk

`func (o *GroupUpdatableProperties) GetTargetReplicaCountOk() (*int64, bool)`

GetTargetReplicaCountOk returns a tuple with the TargetReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetReplicaCount

`func (o *GroupUpdatableProperties) SetTargetReplicaCount(v int64)`

SetTargetReplicaCount sets TargetReplicaCount field to given value.

### HasTargetReplicaCount

`func (o *GroupUpdatableProperties) HasTargetReplicaCount() bool`

HasTargetReplicaCount returns a boolean if a field has been set.

### GetName

`func (o *GroupUpdatableProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupUpdatableProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupUpdatableProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupUpdatableProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicy

`func (o *GroupUpdatableProperties) GetPolicy() GroupPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GroupUpdatableProperties) GetPolicyOk() (*GroupPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GroupUpdatableProperties) SetPolicy(v GroupPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GroupUpdatableProperties) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetReplicaConfiguration

`func (o *GroupUpdatableProperties) GetReplicaConfiguration() ReplicaPropertiesPost`

GetReplicaConfiguration returns the ReplicaConfiguration field if non-nil, zero value otherwise.

### GetReplicaConfigurationOk

`func (o *GroupUpdatableProperties) GetReplicaConfigurationOk() (*ReplicaPropertiesPost, bool)`

GetReplicaConfigurationOk returns a tuple with the ReplicaConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaConfiguration

`func (o *GroupUpdatableProperties) SetReplicaConfiguration(v ReplicaPropertiesPost)`

SetReplicaConfiguration sets ReplicaConfiguration field to given value.

### HasReplicaConfiguration

`func (o *GroupUpdatableProperties) HasReplicaConfiguration() bool`

HasReplicaConfiguration returns a boolean if a field has been set.

### GetDatacenter

`func (o *GroupUpdatableProperties) GetDatacenter() Resource`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *GroupUpdatableProperties) GetDatacenterOk() (*Resource, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *GroupUpdatableProperties) SetDatacenter(v Resource)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *GroupUpdatableProperties) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.



