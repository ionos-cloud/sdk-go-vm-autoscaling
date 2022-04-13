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

// CpuFamily CPU family for the VMs created using this configuration. If null, the VM will be created with the default CPU family for the assigned location.
type CpuFamily string

// List of CpuFamily
const (
	AMD_OPTERON   CpuFamily = "AMD_OPTERON"
	INTEL_SKYLAKE CpuFamily = "INTEL_SKYLAKE"
	INTEL_XEON    CpuFamily = "INTEL_XEON"
)

func (v *CpuFamily) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CpuFamily(value)
	for _, existing := range []CpuFamily{"AMD_OPTERON", "INTEL_SKYLAKE", "INTEL_XEON"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CpuFamily", value)
}

// Ptr returns reference to CpuFamily value
func (v CpuFamily) Ptr() *CpuFamily {
	return &v
}

type NullableCpuFamily struct {
	value *CpuFamily
	isSet bool
}

func (v NullableCpuFamily) Get() *CpuFamily {
	return v.value
}

func (v *NullableCpuFamily) Set(val *CpuFamily) {
	v.value = val
	v.isSet = true
}

func (v NullableCpuFamily) IsSet() bool {
	return v.isSet
}

func (v *NullableCpuFamily) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpuFamily(val *CpuFamily) *NullableCpuFamily {
	return &NullableCpuFamily{value: val, isSet: true}
}

func (v NullableCpuFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpuFamily) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
