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

// ErrorReplicaValidateMessage struct for ErrorReplicaValidateMessage
type ErrorReplicaValidateMessage struct {
	ErrorCode *string `json:"errorCode,omitempty"`
	Message   *string `json:"message,omitempty"`
}

// NewErrorReplicaValidateMessage instantiates a new ErrorReplicaValidateMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorReplicaValidateMessage() *ErrorReplicaValidateMessage {
	this := ErrorReplicaValidateMessage{}

	return &this
}

// NewErrorReplicaValidateMessageWithDefaults instantiates a new ErrorReplicaValidateMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorReplicaValidateMessageWithDefaults() *ErrorReplicaValidateMessage {
	this := ErrorReplicaValidateMessage{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ErrorReplicaValidateMessage) GetErrorCode() *string {
	if o == nil {
		return nil
	}

	return o.ErrorCode

}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorReplicaValidateMessage) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *ErrorReplicaValidateMessage) SetErrorCode(v string) {

	o.ErrorCode = &v

}

// HasErrorCode returns a boolean if a field has been set.
func (o *ErrorReplicaValidateMessage) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ErrorReplicaValidateMessage) GetMessage() *string {
	if o == nil {
		return nil
	}

	return o.Message

}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorReplicaValidateMessage) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Message, true
}

// SetMessage sets field value
func (o *ErrorReplicaValidateMessage) SetMessage(v string) {

	o.Message = &v

}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorReplicaValidateMessage) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

func (o ErrorReplicaValidateMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}

	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	return json.Marshal(toSerialize)
}

type NullableErrorReplicaValidateMessage struct {
	value *ErrorReplicaValidateMessage
	isSet bool
}

func (v NullableErrorReplicaValidateMessage) Get() *ErrorReplicaValidateMessage {
	return v.value
}

func (v *NullableErrorReplicaValidateMessage) Set(val *ErrorReplicaValidateMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorReplicaValidateMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorReplicaValidateMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorReplicaValidateMessage(val *ErrorReplicaValidateMessage) *NullableErrorReplicaValidateMessage {
	return &NullableErrorReplicaValidateMessage{value: val, isSet: true}
}

func (v NullableErrorReplicaValidateMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorReplicaValidateMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
