# ReplicaNic



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Lan** | **int32** | Lan ID for this replica Nic. | |
|**Name** | **string** | Name for this replica NIC. | |
|**Dhcp** | Pointer to **NullableBool** | Dhcp flag for this replica Nic. This is an optional attribute with default value of &#39;true&#39; if not given in the request payload or given as null. | [optional] |

## Methods


### GetLan

`func (o *ReplicaNic) GetLan() int32`

GetLan returns the Lan field if non-nil, zero value otherwise.

### GetLanOk

`func (o *ReplicaNic) GetLanOk() (*int32, bool)`

GetLanOk returns a tuple with the Lan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan

`func (o *ReplicaNic) SetLan(v int32)`

SetLan sets Lan field to given value.


### GetName

`func (o *ReplicaNic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReplicaNic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReplicaNic) SetName(v string)`

SetName sets Name field to given value.


### GetDhcp

`func (o *ReplicaNic) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *ReplicaNic) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *ReplicaNic) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *ReplicaNic) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### SetDhcpNil

`func (o *ReplicaNic) SetDhcpNil(b bool)`

 SetDhcpNil sets the value for Dhcp to be an explicit nil

### UnsetDhcp
`func (o *ReplicaNic) UnsetDhcp()`

UnsetDhcp ensures that no value is present for Dhcp, not even an explicit nil


