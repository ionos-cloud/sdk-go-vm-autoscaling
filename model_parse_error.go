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

// ParseError struct for ParseError
type ParseError struct {
	HttpStatus *int                 `json:"httpStatus,omitempty"`
	Messages   *[]ErrorMessageParse `json:"messages,omitempty"`
	ErrorUuid  *string              `json:"errorUuid,omitempty"`
}

// NewParseError instantiates a new ParseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParseError() *ParseError {
	this := ParseError{}

	return &this
}

// NewParseErrorWithDefaults instantiates a new ParseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParseErrorWithDefaults() *ParseError {
	this := ParseError{}
	return &this
}

// GetHttpStatus returns the HttpStatus field value
// If the value is explicit nil, the zero value for int will be returned
func (o *ParseError) GetHttpStatus() *int {
	if o == nil {
		return nil
	}

	return o.HttpStatus

}

// GetHttpStatusOk returns a tuple with the HttpStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParseError) GetHttpStatusOk() (*int, bool) {
	if o == nil {
		return nil, false
	}

	return o.HttpStatus, true
}

// SetHttpStatus sets field value
func (o *ParseError) SetHttpStatus(v int) {

	o.HttpStatus = &v

}

// HasHttpStatus returns a boolean if a field has been set.
func (o *ParseError) HasHttpStatus() bool {
	if o != nil && o.HttpStatus != nil {
		return true
	}

	return false
}

// GetMessages returns the Messages field value
// If the value is explicit nil, the zero value for []ErrorMessageParse will be returned
func (o *ParseError) GetMessages() *[]ErrorMessageParse {
	if o == nil {
		return nil
	}

	return o.Messages

}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParseError) GetMessagesOk() (*[]ErrorMessageParse, bool) {
	if o == nil {
		return nil, false
	}

	return o.Messages, true
}

// SetMessages sets field value
func (o *ParseError) SetMessages(v []ErrorMessageParse) {

	o.Messages = &v

}

// HasMessages returns a boolean if a field has been set.
func (o *ParseError) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// GetErrorUuid returns the ErrorUuid field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ParseError) GetErrorUuid() *string {
	if o == nil {
		return nil
	}

	return o.ErrorUuid

}

// GetErrorUuidOk returns a tuple with the ErrorUuid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParseError) GetErrorUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ErrorUuid, true
}

// SetErrorUuid sets field value
func (o *ParseError) SetErrorUuid(v string) {

	o.ErrorUuid = &v

}

// HasErrorUuid returns a boolean if a field has been set.
func (o *ParseError) HasErrorUuid() bool {
	if o != nil && o.ErrorUuid != nil {
		return true
	}

	return false
}

func (o ParseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HttpStatus != nil {
		toSerialize["httpStatus"] = o.HttpStatus
	}

	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}

	if o.ErrorUuid != nil {
		toSerialize["errorUuid"] = o.ErrorUuid
	}

	return json.Marshal(toSerialize)
}

type NullableParseError struct {
	value *ParseError
	isSet bool
}

func (v NullableParseError) Get() *ParseError {
	return v.value
}

func (v *NullableParseError) Set(val *ParseError) {
	v.value = val
	v.isSet = true
}

func (v NullableParseError) IsSet() bool {
	return v.isSet
}

func (v *NullableParseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParseError(val *ParseError) *NullableParseError {
	return &NullableParseError{value: val, isSet: true}
}

func (v NullableParseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
