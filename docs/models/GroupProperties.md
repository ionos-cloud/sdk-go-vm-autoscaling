# GroupProperties



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
|**Location** | **string** | Location of the data center. | [readonly] |

## Methods


### GetMaxReplicaCount

`func (o *GroupProperties) GetMaxReplicaCount() int64`

GetMaxReplicaCount returns the MaxReplicaCount field if non-nil, zero value otherwise.

### GetMaxReplicaCountOk

`func (o *GroupProperties) GetMaxReplicaCountOk() (*int64, bool)`

GetMaxReplicaCountOk returns a tuple with the MaxReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicaCount

`func (o *GroupProperties) SetMaxReplicaCount(v int64)`

SetMaxReplicaCount sets MaxReplicaCount field to given value.

### HasMaxReplicaCount

`func (o *GroupProperties) HasMaxReplicaCount() bool`

HasMaxReplicaCount returns a boolean if a field has been set.

### GetMinReplicaCount

`func (o *GroupProperties) GetMinReplicaCount() int64`

GetMinReplicaCount returns the MinReplicaCount field if non-nil, zero value otherwise.

### GetMinReplicaCountOk

`func (o *GroupProperties) GetMinReplicaCountOk() (*int64, bool)`

GetMinReplicaCountOk returns a tuple with the MinReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicaCount

`func (o *GroupProperties) SetMinReplicaCount(v int64)`

SetMinReplicaCount sets MinReplicaCount field to given value.

### HasMinReplicaCount

`func (o *GroupProperties) HasMinReplicaCount() bool`

HasMinReplicaCount returns a boolean if a field has been set.

### GetTargetReplicaCount

`func (o *GroupProperties) GetTargetReplicaCount() int64`

GetTargetReplicaCount returns the TargetReplicaCount field if non-nil, zero value otherwise.

### GetTargetReplicaCountOk

`func (o *GroupProperties) GetTargetReplicaCountOk() (*int64, bool)`

GetTargetReplicaCountOk returns a tuple with the TargetReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetReplicaCount

`func (o *GroupProperties) SetTargetReplicaCount(v int64)`

SetTargetReplicaCount sets TargetReplicaCount field to given value.

### HasTargetReplicaCount

`func (o *GroupProperties) HasTargetReplicaCount() bool`

HasTargetReplicaCount returns a boolean if a field has been set.

### GetName

`func (o *GroupProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicy

`func (o *GroupProperties) GetPolicy() GroupPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GroupProperties) GetPolicyOk() (*GroupPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GroupProperties) SetPolicy(v GroupPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GroupProperties) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetReplicaConfiguration

`func (o *GroupProperties) GetReplicaConfiguration() ReplicaPropertiesPost`

GetReplicaConfiguration returns the ReplicaConfiguration field if non-nil, zero value otherwise.

### GetReplicaConfigurationOk

`func (o *GroupProperties) GetReplicaConfigurationOk() (*ReplicaPropertiesPost, bool)`

GetReplicaConfigurationOk returns a tuple with the ReplicaConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaConfiguration

`func (o *GroupProperties) SetReplicaConfiguration(v ReplicaPropertiesPost)`

SetReplicaConfiguration sets ReplicaConfiguration field to given value.

### HasReplicaConfiguration

`func (o *GroupProperties) HasReplicaConfiguration() bool`

HasReplicaConfiguration returns a boolean if a field has been set.

### GetDatacenter

`func (o *GroupProperties) GetDatacenter() Resource`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *GroupProperties) GetDatacenterOk() (*Resource, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *GroupProperties) SetDatacenter(v Resource)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *GroupProperties) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetLocation

`func (o *GroupProperties) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GroupProperties) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GroupProperties) SetLocation(v string)`

SetLocation sets Location field to given value.




