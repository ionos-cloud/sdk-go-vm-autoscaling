/*
 * VM Auto Scaling service (CloudAPI)
 *
 * VM Auto Scaling service enables IONOS clients to horizontally scale the number of VM instances, based on configured rules. Use Auto Scaling to ensure you will have a sufficient number of instances to handle your application loads at all times.  Create an Auto Scaling group that contains the server instances; Auto Scaling service will ensure that the number of instances in the group is always within these limits.  When target replica count is specified, Auto Scaling will maintain the set number on instances.  When scaling policies are specified, Auto Scaling will create or delete instances based on the demands of your applications. For each policy, specified scale-in and scale-out actions are performed whenever the corresponding thresholds are met.
 *
 * API version: 1.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud_vm_autoscaling

import (
	"encoding/json"
)

// ErrorUuid struct for ErrorUuid
type ErrorUuid struct {
	ErrorUuid *string `json:"errorUuid,omitempty"`
}

// NewErrorUuid instantiates a new ErrorUuid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorUuid() *ErrorUuid {
	this := ErrorUuid{}

	return &this
}

// NewErrorUuidWithDefaults instantiates a new ErrorUuid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorUuidWithDefaults() *ErrorUuid {
	this := ErrorUuid{}
	return &this
}

// GetErrorUuid returns the ErrorUuid field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ErrorUuid) GetErrorUuid() *string {
	if o == nil {
		return nil
	}

	return o.ErrorUuid

}

// GetErrorUuidOk returns a tuple with the ErrorUuid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorUuid) GetErrorUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ErrorUuid, true
}

// SetErrorUuid sets field value
func (o *ErrorUuid) SetErrorUuid(v string) {

	o.ErrorUuid = &v

}

// HasErrorUuid returns a boolean if a field has been set.
func (o *ErrorUuid) HasErrorUuid() bool {
	if o != nil && o.ErrorUuid != nil {
		return true
	}

	return false
}

func (o ErrorUuid) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorUuid != nil {
		toSerialize["errorUuid"] = o.ErrorUuid
	}

	return json.Marshal(toSerialize)
}

type NullableErrorUuid struct {
	value *ErrorUuid
	isSet bool
}

func (v NullableErrorUuid) Get() *ErrorUuid {
	return v.value
}

func (v *NullableErrorUuid) Set(val *ErrorUuid) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorUuid) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorUuid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorUuid(val *ErrorUuid) *NullableErrorUuid {
	return &NullableErrorUuid{value: val, isSet: true}
}

func (v NullableErrorUuid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorUuid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
