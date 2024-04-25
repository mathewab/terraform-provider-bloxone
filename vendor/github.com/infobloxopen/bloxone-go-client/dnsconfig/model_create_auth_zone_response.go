/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"encoding/json"
)

// checks if the CreateAuthZoneResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAuthZoneResponse{}

// CreateAuthZoneResponse The Authoritative Zone object create response format.
type CreateAuthZoneResponse struct {
	Result *AuthZone `json:"result,omitempty"`
}

// NewCreateAuthZoneResponse instantiates a new CreateAuthZoneResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthZoneResponse() *CreateAuthZoneResponse {
	this := CreateAuthZoneResponse{}
	return &this
}

// NewCreateAuthZoneResponseWithDefaults instantiates a new CreateAuthZoneResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthZoneResponseWithDefaults() *CreateAuthZoneResponse {
	this := CreateAuthZoneResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateAuthZoneResponse) GetResult() AuthZone {
	if o == nil || IsNil(o.Result) {
		var ret AuthZone
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthZoneResponse) GetResultOk() (*AuthZone, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateAuthZoneResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given AuthZone and assigns it to the Result field.
func (o *CreateAuthZoneResponse) SetResult(v AuthZone) {
	o.Result = &v
}

func (o CreateAuthZoneResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAuthZoneResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableCreateAuthZoneResponse struct {
	value *CreateAuthZoneResponse
	isSet bool
}

func (v NullableCreateAuthZoneResponse) Get() *CreateAuthZoneResponse {
	return v.value
}

func (v *NullableCreateAuthZoneResponse) Set(val *CreateAuthZoneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthZoneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthZoneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthZoneResponse(val *CreateAuthZoneResponse) *NullableCreateAuthZoneResponse {
	return &NullableCreateAuthZoneResponse{value: val, isSet: true}
}

func (v NullableCreateAuthZoneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthZoneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
