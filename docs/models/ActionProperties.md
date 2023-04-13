# ActionProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ActionStatus** | [**ActionStatus**](ActionStatus.md) |  | |
|**ActionType** | [**ActionType**](ActionType.md) |  | |
|**TargetReplicaCount** | **int64** | This action is considered successful when the associated autoscaling group reaches this replica count. | |

## Methods

### NewActionProperties

`func NewActionProperties(actionStatus ActionStatus, actionType ActionType, targetReplicaCount int64, ) *ActionProperties`

NewActionProperties instantiates a new ActionProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionPropertiesWithDefaults

`func NewActionPropertiesWithDefaults() *ActionProperties`

NewActionPropertiesWithDefaults instantiates a new ActionProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *ActionProperties) GetActionStatus() ActionStatus`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *ActionProperties) GetActionStatusOk() (*ActionStatus, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *ActionProperties) SetActionStatus(v ActionStatus)`

SetActionStatus sets ActionStatus field to given value.


### GetActionType

`func (o *ActionProperties) GetActionType() ActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *ActionProperties) GetActionTypeOk() (*ActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *ActionProperties) SetActionType(v ActionType)`

SetActionType sets ActionType field to given value.


### GetTargetReplicaCount

`func (o *ActionProperties) GetTargetReplicaCount() int64`

GetTargetReplicaCount returns the TargetReplicaCount field if non-nil, zero value otherwise.

### GetTargetReplicaCountOk

`func (o *ActionProperties) GetTargetReplicaCountOk() (*int64, bool)`

GetTargetReplicaCountOk returns a tuple with the TargetReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetReplicaCount

`func (o *ActionProperties) SetTargetReplicaCount(v int64)`

SetTargetReplicaCount sets TargetReplicaCount field to given value.




