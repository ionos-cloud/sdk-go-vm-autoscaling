# GroupEntities



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Actions** | Pointer to [**ActionCollection**](ActionCollection.md) |  | [optional] |
|**Servers** | Pointer to [**ServerCollection**](ServerCollection.md) |  | [optional] |

## Methods


### GetActions

`func (o *GroupEntities) GetActions() ActionCollection`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GroupEntities) GetActionsOk() (*ActionCollection, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GroupEntities) SetActions(v ActionCollection)`

SetActions sets Actions field to given value.

### HasActions

`func (o *GroupEntities) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetServers

`func (o *GroupEntities) GetServers() ServerCollection`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GroupEntities) GetServersOk() (*ServerCollection, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GroupEntities) SetServers(v ServerCollection)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GroupEntities) HasServers() bool`

HasServers returns a boolean if a field has been set.



