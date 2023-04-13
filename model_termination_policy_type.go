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
	"fmt"
)

// TerminationPolicyType The type of termination policy for the Auto Scaling group to follow a specific pattern for scaling-in replicas. The default termination policy is 'OLDEST_SERVER_FIRST'.
type TerminationPolicyType string

// List of TerminationPolicyType
const (
	OLDEST_SERVER_FIRST TerminationPolicyType = "OLDEST_SERVER_FIRST"
	NEWEST_SERVER_FIRST TerminationPolicyType = "NEWEST_SERVER_FIRST"
	RANDOM              TerminationPolicyType = "RANDOM"
)

func (v *TerminationPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TerminationPolicyType(value)
	for _, existing := range []TerminationPolicyType{"OLDEST_SERVER_FIRST", "NEWEST_SERVER_FIRST", "RANDOM"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TerminationPolicyType", value)
}

// Ptr returns reference to TerminationPolicyType value
func (v TerminationPolicyType) Ptr() *TerminationPolicyType {
	return &v
}

type NullableTerminationPolicyType struct {
	value *TerminationPolicyType
	isSet bool
}

func (v NullableTerminationPolicyType) Get() *TerminationPolicyType {
	return v.value
}

func (v *NullableTerminationPolicyType) Set(val *TerminationPolicyType) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminationPolicyType) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminationPolicyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminationPolicyType(val *TerminationPolicyType) *NullableTerminationPolicyType {
	return &NullableTerminationPolicyType{value: val, isSet: true}
}

func (v NullableTerminationPolicyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminationPolicyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
