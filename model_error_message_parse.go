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

// ErrorMessageParse struct for ErrorMessageParse
type ErrorMessageParse struct {
	ErrorCode *int    `json:"errorCode,omitempty"`
	Message   *string `json:"message,omitempty"`
}

// NewErrorMessageParse instantiates a new ErrorMessageParse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorMessageParse() *ErrorMessageParse {
	this := ErrorMessageParse{}

	return &this
}

// NewErrorMessageParseWithDefaults instantiates a new ErrorMessageParse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorMessageParseWithDefaults() *ErrorMessageParse {
	this := ErrorMessageParse{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
// If the value is explicit nil, the zero value for int will be returned
func (o *ErrorMessageParse) GetErrorCode() *int {
	if o == nil {
		return nil
	}

	return o.ErrorCode

}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorMessageParse) GetErrorCodeOk() (*int, bool) {
	if o == nil {
		return nil, false
	}

	return o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *ErrorMessageParse) SetErrorCode(v int) {

	o.ErrorCode = &v

}

// HasErrorCode returns a boolean if a field has been set.
func (o *ErrorMessageParse) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ErrorMessageParse) GetMessage() *string {
	if o == nil {
		return nil
	}

	return o.Message

}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorMessageParse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Message, true
}

// SetMessage sets field value
func (o *ErrorMessageParse) SetMessage(v string) {

	o.Message = &v

}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorMessageParse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

func (o ErrorMessageParse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}

	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	return json.Marshal(toSerialize)
}

type NullableErrorMessageParse struct {
	value *ErrorMessageParse
	isSet bool
}

func (v NullableErrorMessageParse) Get() *ErrorMessageParse {
	return v.value
}

func (v *NullableErrorMessageParse) Set(val *ErrorMessageParse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorMessageParse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorMessageParse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorMessageParse(val *ErrorMessageParse) *NullableErrorMessageParse {
	return &NullableErrorMessageParse{value: val, isSet: true}
}

func (v NullableErrorMessageParse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorMessageParse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
