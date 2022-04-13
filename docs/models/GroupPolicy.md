# GroupPolicy

Specifies the behavior of this autoscaling group. A policy consists of Triggers and Actions, whereby an Action is some kind of automated behavior, and the Trigger defines the circumstances, under which the Action is triggered. Currently, two separate Actions, namely Scaling In and Out are supported, triggered through the thresholds, defined for a given Metric.


## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metric** | [**Metric**](Metric.md) |  | |
|**Range** | Pointer to **string** | Defines the time range, for which the samples will be aggregated. | [optional] [default to "120s"]|
|**ScaleInAction** | [**GroupPolicyScaleInAction**](GroupPolicyScaleInAction.md) |  | |
|**ScaleInThreshold** | **float32** | The lower threshold for the value of the &#x60;metric&#x60;. Will be used with &#x60;less than&#x60; (&lt;) operator. Exceeding this will start a Scale-In action as specified by the &#x60;scaleInAction&#x60; property. The value must have a higher minimum delta to the &#x60;scaleOutThreshold&#x60; depending on the &#x60;metric&#x60; to avoid competitive actions at the same time. | |
|**ScaleOutAction** | [**GroupPolicyScaleOutAction**](GroupPolicyScaleOutAction.md) |  | |
|**ScaleOutThreshold** | **float32** | The upper threshold for the value of the &#x60;metric&#x60;.  Will be used with &#x60;greater than&#x60; (&gt;) operator. Exceeding this will start a Scale-Out action as specified by the &#x60;scaleOutAction&#x60; property. The value must have a lower minimum delta to the &#x60;scaleInThreshold&#x60; depending on the &#x60;metric&#x60; to avoid competitive actions at the same time. | |
|**Unit** | [**QueryUnit**](QueryUnit.md) |  | |

## Methods


### GetMetric

`func (o *GroupPolicy) GetMetric() Metric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *GroupPolicy) GetMetricOk() (*Metric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *GroupPolicy) SetMetric(v Metric)`

SetMetric sets Metric field to given value.


### GetRange

`func (o *GroupPolicy) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *GroupPolicy) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *GroupPolicy) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *GroupPolicy) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScaleInAction

`func (o *GroupPolicy) GetScaleInAction() GroupPolicyScaleInAction`

GetScaleInAction returns the ScaleInAction field if non-nil, zero value otherwise.

### GetScaleInActionOk

`func (o *GroupPolicy) GetScaleInActionOk() (*GroupPolicyScaleInAction, bool)`

GetScaleInActionOk returns a tuple with the ScaleInAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleInAction

`func (o *GroupPolicy) SetScaleInAction(v GroupPolicyScaleInAction)`

SetScaleInAction sets ScaleInAction field to given value.


### GetScaleInThreshold

`func (o *GroupPolicy) GetScaleInThreshold() float32`

GetScaleInThreshold returns the ScaleInThreshold field if non-nil, zero value otherwise.

### GetScaleInThresholdOk

`func (o *GroupPolicy) GetScaleInThresholdOk() (*float32, bool)`

GetScaleInThresholdOk returns a tuple with the ScaleInThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleInThreshold

`func (o *GroupPolicy) SetScaleInThreshold(v float32)`

SetScaleInThreshold sets ScaleInThreshold field to given value.


### GetScaleOutAction

`func (o *GroupPolicy) GetScaleOutAction() GroupPolicyScaleOutAction`

GetScaleOutAction returns the ScaleOutAction field if non-nil, zero value otherwise.

### GetScaleOutActionOk

`func (o *GroupPolicy) GetScaleOutActionOk() (*GroupPolicyScaleOutAction, bool)`

GetScaleOutActionOk returns a tuple with the ScaleOutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleOutAction

`func (o *GroupPolicy) SetScaleOutAction(v GroupPolicyScaleOutAction)`

SetScaleOutAction sets ScaleOutAction field to given value.


### GetScaleOutThreshold

`func (o *GroupPolicy) GetScaleOutThreshold() float32`

GetScaleOutThreshold returns the ScaleOutThreshold field if non-nil, zero value otherwise.

### GetScaleOutThresholdOk

`func (o *GroupPolicy) GetScaleOutThresholdOk() (*float32, bool)`

GetScaleOutThresholdOk returns a tuple with the ScaleOutThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleOutThreshold

`func (o *GroupPolicy) SetScaleOutThreshold(v float32)`

SetScaleOutThreshold sets ScaleOutThreshold field to given value.


### GetUnit

`func (o *GroupPolicy) GetUnit() QueryUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GroupPolicy) GetUnitOk() (*QueryUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GroupPolicy) SetUnit(v QueryUnit)`

SetUnit sets Unit field to given value.




