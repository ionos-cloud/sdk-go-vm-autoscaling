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

// ParseErrorAllOf struct for ParseErrorAllOf
type ParseErrorAllOf struct {
	HttpStatus *int                 `json:"httpStatus,omitempty"`
	Messages   *[]ErrorMessageParse `json:"messages,omitempty"`
}

// NewParseErrorAllOf instantiates a new ParseErrorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParseErrorAllOf() *ParseErrorAllOf {
	this := ParseErrorAllOf{}

	return &this
}

// NewParseErrorAllOfWithDefaults instantiates a new ParseErrorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParseErrorAllOfWithDefaults() *ParseErrorAllOf {
	this := ParseErrorAllOf{}
	return &this
}

// GetHttpStatus returns the HttpStatus field value
// If the value is explicit nil, the zero value for int will be returned
func (o *ParseErrorAllOf) GetHttpStatus() *int {
	if o == nil {
		return nil
	}

	return o.HttpStatus

}

// GetHttpStatusOk returns a tuple with the HttpStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParseErrorAllOf) GetHttpStatusOk() (*int, bool) {
	if o == nil {
		return nil, false
	}

	return o.HttpStatus, true
}

// SetHttpStatus sets field value
func (o *ParseErrorAllOf) SetHttpStatus(v int) {

	o.HttpStatus = &v

}

// HasHttpStatus returns a boolean if a field has been set.
func (o *ParseErrorAllOf) HasHttpStatus() bool {
	if o != nil && o.HttpStatus != nil {
		return true
	}

	return false
}

// GetMessages returns the Messages field value
// If the value is explicit nil, the zero value for []ErrorMessageParse will be returned
func (o *ParseErrorAllOf) GetMessages() *[]ErrorMessageParse {
	if o == nil {
		return nil
	}

	return o.Messages

}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParseErrorAllOf) GetMessagesOk() (*[]ErrorMessageParse, bool) {
	if o == nil {
		return nil, false
	}

	return o.Messages, true
}

// SetMessages sets field value
func (o *ParseErrorAllOf) SetMessages(v []ErrorMessageParse) {

	o.Messages = &v

}

// HasMessages returns a boolean if a field has been set.
func (o *ParseErrorAllOf) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

func (o ParseErrorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HttpStatus != nil {
		toSerialize["httpStatus"] = o.HttpStatus
	}

	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}

	return json.Marshal(toSerialize)
}

type NullableParseErrorAllOf struct {
	value *ParseErrorAllOf
	isSet bool
}

func (v NullableParseErrorAllOf) Get() *ParseErrorAllOf {
	return v.value
}

func (v *NullableParseErrorAllOf) Set(val *ParseErrorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableParseErrorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableParseErrorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParseErrorAllOf(val *ParseErrorAllOf) *NullableParseErrorAllOf {
	return &NullableParseErrorAllOf{value: val, isSet: true}
}

func (v NullableParseErrorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParseErrorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
