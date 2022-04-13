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

// DatacenterServerAllOf Links to the actual server entity in the data center.
type DatacenterServerAllOf struct {
	// URL to the object representation (absolute path)
	Href *string `json:"href,omitempty"`
	// The resource's unique identifier
	Id *string `json:"id,omitempty"`
	// The type of object that has been created
	Type *string `json:"type,omitempty"`
}

// NewDatacenterServerAllOf instantiates a new DatacenterServerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatacenterServerAllOf() *DatacenterServerAllOf {
	this := DatacenterServerAllOf{}

	return &this
}

// NewDatacenterServerAllOfWithDefaults instantiates a new DatacenterServerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatacenterServerAllOfWithDefaults() *DatacenterServerAllOf {
	this := DatacenterServerAllOf{}
	return &this
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DatacenterServerAllOf) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatacenterServerAllOf) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *DatacenterServerAllOf) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *DatacenterServerAllOf) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DatacenterServerAllOf) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatacenterServerAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *DatacenterServerAllOf) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *DatacenterServerAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DatacenterServerAllOf) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatacenterServerAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *DatacenterServerAllOf) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *DatacenterServerAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

func (o DatacenterServerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	return json.Marshal(toSerialize)
}

type NullableDatacenterServerAllOf struct {
	value *DatacenterServerAllOf
	isSet bool
}

func (v NullableDatacenterServerAllOf) Get() *DatacenterServerAllOf {
	return v.value
}

func (v *NullableDatacenterServerAllOf) Set(val *DatacenterServerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDatacenterServerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDatacenterServerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatacenterServerAllOf(val *DatacenterServerAllOf) *NullableDatacenterServerAllOf {
	return &NullableDatacenterServerAllOf{value: val, isSet: true}
}

func (v NullableDatacenterServerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatacenterServerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
