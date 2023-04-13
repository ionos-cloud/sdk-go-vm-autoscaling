/*
 * VM Auto Scaling API
 *
 * The VM Auto Scaling Service enables IONOS clients to horizontally scale the number of VM replicas based on configured rules. You can use Auto Scaling to ensure that you have a sufficient number of replicas to handle your application loads at all times.  For this purpose, create an Auto Scaling group that contains the server replicas. The VM Auto Scaling Service ensures that the number of replicas in the group is always within the defined limits. For example, if the number of target replicas is specified, Auto Scaling maintains the specified number of replicas.   When scaling policies are set, Auto Scaling creates or deletes replicas according to the requirements of your applications. For each policy, specified 'scale-in' and 'scale-out' actions are performed when the corresponding thresholds are reached.
 *
 * API version: 1-SDK.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud_vm_autoscaling

import (
	"encoding/json"
)

// GroupPolicyScaleOutAction Defines the action to be performed when the 'scaleOutThreshold' is exceeded. Here, scaling is always about adding new VMs to this Auto Scaling group.
type GroupPolicyScaleOutAction struct {
	// 'amountType=ABSOLUTE' specifies the absolute number of VMs that are added or removed. The value must be between 1 to 10.   'amountType=PERCENTAGE' specifies the percentage value that is applied to the current 'targetReplicaCount' of the autoscaling group. The value must be between 1 to 200.   At least one VM is always added or removed.   Note that for 'SCALE_IN' operations, volumes are not deleted after the server is deleted.
	Amount     *float32      `json:"amount"`
	AmountType *ActionAmount `json:"amountType"`
	// The minimum time that elapses after the start of this scaling action until the following scaling action is started. While a scaling action is in progress, no second action is initiated for the same Auto Scaling group. Instead, the metric is re-evaluated after the current scaling action completes (either successfully or with errors). This is currently validated with a minimum value of 2 minutes and a maximum of 24 hours. The default value is 5 minutes if not specified.
	CooldownPeriod *string `json:"cooldownPeriod,omitempty"`
}

// NewGroupPolicyScaleOutAction instantiates a new GroupPolicyScaleOutAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupPolicyScaleOutAction(amount float32, amountType ActionAmount) *GroupPolicyScaleOutAction {
	this := GroupPolicyScaleOutAction{}

	this.Amount = &amount

	this.AmountType = &amountType

	var cooldownPeriod string = "5m"
	this.CooldownPeriod = &cooldownPeriod

	return &this
}

// NewGroupPolicyScaleOutActionWithDefaults instantiates a new GroupPolicyScaleOutAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPolicyScaleOutActionWithDefaults() *GroupPolicyScaleOutAction {
	this := GroupPolicyScaleOutAction{}
	var cooldownPeriod string = "5m"
	this.CooldownPeriod = &cooldownPeriod
	return &this
}

// GetAmount returns the Amount field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *GroupPolicyScaleOutAction) GetAmount() *float32 {
	if o == nil {
		return nil
	}

	return o.Amount

}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupPolicyScaleOutAction) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Amount, true
}

// SetAmount sets field value
func (o *GroupPolicyScaleOutAction) SetAmount(v float32) {

	o.Amount = &v

}

// HasAmount returns a boolean if a field has been set.
func (o *GroupPolicyScaleOutAction) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// GetAmountType returns the AmountType field value
// If the value is explicit nil, the zero value for ActionAmount will be returned
func (o *GroupPolicyScaleOutAction) GetAmountType() *ActionAmount {
	if o == nil {
		return nil
	}

	return o.AmountType

}

// GetAmountTypeOk returns a tuple with the AmountType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupPolicyScaleOutAction) GetAmountTypeOk() (*ActionAmount, bool) {
	if o == nil {
		return nil, false
	}

	return o.AmountType, true
}

// SetAmountType sets field value
func (o *GroupPolicyScaleOutAction) SetAmountType(v ActionAmount) {

	o.AmountType = &v

}

// HasAmountType returns a boolean if a field has been set.
func (o *GroupPolicyScaleOutAction) HasAmountType() bool {
	if o != nil && o.AmountType != nil {
		return true
	}

	return false
}

// GetCooldownPeriod returns the CooldownPeriod field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupPolicyScaleOutAction) GetCooldownPeriod() *string {
	if o == nil {
		return nil
	}

	return o.CooldownPeriod

}

// GetCooldownPeriodOk returns a tuple with the CooldownPeriod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupPolicyScaleOutAction) GetCooldownPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.CooldownPeriod, true
}

// SetCooldownPeriod sets field value
func (o *GroupPolicyScaleOutAction) SetCooldownPeriod(v string) {

	o.CooldownPeriod = &v

}

// HasCooldownPeriod returns a boolean if a field has been set.
func (o *GroupPolicyScaleOutAction) HasCooldownPeriod() bool {
	if o != nil && o.CooldownPeriod != nil {
		return true
	}

	return false
}

func (o GroupPolicyScaleOutAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}

	if o.AmountType != nil {
		toSerialize["amountType"] = o.AmountType
	}

	toSerialize["cooldownPeriod"] = o.CooldownPeriod

	return json.Marshal(toSerialize)
}

type NullableGroupPolicyScaleOutAction struct {
	value *GroupPolicyScaleOutAction
	isSet bool
}

func (v NullableGroupPolicyScaleOutAction) Get() *GroupPolicyScaleOutAction {
	return v.value
}

func (v *NullableGroupPolicyScaleOutAction) Set(val *GroupPolicyScaleOutAction) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupPolicyScaleOutAction) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupPolicyScaleOutAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupPolicyScaleOutAction(val *GroupPolicyScaleOutAction) *NullableGroupPolicyScaleOutAction {
	return &NullableGroupPolicyScaleOutAction{value: val, isSet: true}
}

func (v NullableGroupPolicyScaleOutAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupPolicyScaleOutAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
