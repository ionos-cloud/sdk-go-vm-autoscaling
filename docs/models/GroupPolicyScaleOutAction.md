# GroupPolicyScaleOutAction



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Amount** | **float32** | When &#x60;amountType &#x3D;&#x3D; ABSOLUTE&#x60;, this is the number of VMs added or removed in one step. When &#x60;amountType &#x3D;&#x3D; PERCENTAGE&#x60;, this is a percentage value, which will be applied to the autoscaling group&#39;s current &#x60;targetReplicaCount&#x60; in order to derive the number of VMs that will be added or removed in one step. There will always be at least one VM added or removed.   For SCALE_IN operation now volumes are NOT deleted after the server deletion. | |
|**AmountType** | [**ActionAmount**](ActionAmount.md) |  | |
|**CooldownPeriod** | Pointer to **NullableString** | Minimum time to pass after this Scaling action has started, until the next Scaling action will be started. Additionally, if a Scaling action is currently in progress, no second Scaling action will be started for the same autoscaling group. Instead, the Metric will be re-evaluated after the current Scaling action is completed (either successfully or with failures). This is validated with a minimum value of 2 minutes and a maximum of 24 hours currently. Default value is 5 minutes if not given. | [optional] [default to "5m"]|

## Methods


### GetAmount

`func (o *GroupPolicyScaleOutAction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GroupPolicyScaleOutAction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GroupPolicyScaleOutAction) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetAmountType

`func (o *GroupPolicyScaleOutAction) GetAmountType() ActionAmount`

GetAmountType returns the AmountType field if non-nil, zero value otherwise.

### GetAmountTypeOk

`func (o *GroupPolicyScaleOutAction) GetAmountTypeOk() (*ActionAmount, bool)`

GetAmountTypeOk returns a tuple with the AmountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountType

`func (o *GroupPolicyScaleOutAction) SetAmountType(v ActionAmount)`

SetAmountType sets AmountType field to given value.


### GetCooldownPeriod

`func (o *GroupPolicyScaleOutAction) GetCooldownPeriod() string`

GetCooldownPeriod returns the CooldownPeriod field if non-nil, zero value otherwise.

### GetCooldownPeriodOk

`func (o *GroupPolicyScaleOutAction) GetCooldownPeriodOk() (*string, bool)`

GetCooldownPeriodOk returns a tuple with the CooldownPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownPeriod

`func (o *GroupPolicyScaleOutAction) SetCooldownPeriod(v string)`

SetCooldownPeriod sets CooldownPeriod field to given value.

### HasCooldownPeriod

`func (o *GroupPolicyScaleOutAction) HasCooldownPeriod() bool`

HasCooldownPeriod returns a boolean if a field has been set.

### SetCooldownPeriodNil

`func (o *GroupPolicyScaleOutAction) SetCooldownPeriodNil(b bool)`

 SetCooldownPeriodNil sets the value for CooldownPeriod to be an explicit nil

### UnsetCooldownPeriod
`func (o *GroupPolicyScaleOutAction) UnsetCooldownPeriod()`

UnsetCooldownPeriod ensures that no value is present for CooldownPeriod, not even an explicit nil


