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
	"fmt"
)

// MetadataState the model 'MetadataState'
type MetadataState string

// List of MetadataState
const (
	AVAILABLE MetadataState = "AVAILABLE"
	BUSY      MetadataState = "BUSY"
	INACTIVE  MetadataState = "INACTIVE"
	SUSPENDED MetadataState = "SUSPENDED"
)

func (v *MetadataState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetadataState(value)
	for _, existing := range []MetadataState{"AVAILABLE", "BUSY", "INACTIVE", "SUSPENDED"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetadataState", value)
}

// Ptr returns reference to MetadataState value
func (v MetadataState) Ptr() *MetadataState {
	return &v
}

type NullableMetadataState struct {
	value *MetadataState
	isSet bool
}

func (v NullableMetadataState) Get() *MetadataState {
	return v.value
}

func (v *NullableMetadataState) Set(val *MetadataState) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataState) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataState(val *MetadataState) *NullableMetadataState {
	return &NullableMetadataState{value: val, isSet: true}
}

func (v NullableMetadataState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
