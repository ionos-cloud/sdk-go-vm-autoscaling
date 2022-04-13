# StartedActions



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**StartedActions** | Pointer to [**[]ActionResource**](ActionResource.md) | Any background activity caused by this request. This allows the caller to track the progress of such activity. | [optional] [readonly] |

## Methods


### GetStartedActions

`func (o *StartedActions) GetStartedActions() []ActionResource`

GetStartedActions returns the StartedActions field if non-nil, zero value otherwise.

### GetStartedActionsOk

`func (o *StartedActions) GetStartedActionsOk() (*[]ActionResource, bool)`

GetStartedActionsOk returns a tuple with the StartedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedActions

`func (o *StartedActions) SetStartedActions(v []ActionResource)`

SetStartedActions sets StartedActions field to given value.

### HasStartedActions

`func (o *StartedActions) HasStartedActions() bool`

HasStartedActions returns a boolean if a field has been set.



