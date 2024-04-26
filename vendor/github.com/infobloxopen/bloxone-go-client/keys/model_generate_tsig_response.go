/*
DDI Keys API

The DDI Keys application is a BloxOne DDI service for managing TSIG keys and GSS-TSIG (Kerberos) keys which are used by other BloxOne DDI applications. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keys

import (
	"encoding/json"
)

// checks if the GenerateTSIGResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateTSIGResponse{}

// GenerateTSIGResponse The response format to generate the TSIG key.
type GenerateTSIGResponse struct {
	Result               *GenerateTSIGResult `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GenerateTSIGResponse GenerateTSIGResponse

// NewGenerateTSIGResponse instantiates a new GenerateTSIGResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateTSIGResponse() *GenerateTSIGResponse {
	this := GenerateTSIGResponse{}
	return &this
}

// NewGenerateTSIGResponseWithDefaults instantiates a new GenerateTSIGResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateTSIGResponseWithDefaults() *GenerateTSIGResponse {
	this := GenerateTSIGResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GenerateTSIGResponse) GetResult() GenerateTSIGResult {
	if o == nil || IsNil(o.Result) {
		var ret GenerateTSIGResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateTSIGResponse) GetResultOk() (*GenerateTSIGResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GenerateTSIGResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given GenerateTSIGResult and assigns it to the Result field.
func (o *GenerateTSIGResponse) SetResult(v GenerateTSIGResult) {
	o.Result = &v
}

func (o GenerateTSIGResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateTSIGResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GenerateTSIGResponse) UnmarshalJSON(data []byte) (err error) {
	varGenerateTSIGResponse := _GenerateTSIGResponse{}

	err = json.Unmarshal(data, &varGenerateTSIGResponse)

	if err != nil {
		return err
	}

	*o = GenerateTSIGResponse(varGenerateTSIGResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenerateTSIGResponse struct {
	value *GenerateTSIGResponse
	isSet bool
}

func (v NullableGenerateTSIGResponse) Get() *GenerateTSIGResponse {
	return v.value
}

func (v *NullableGenerateTSIGResponse) Set(val *GenerateTSIGResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateTSIGResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateTSIGResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateTSIGResponse(val *GenerateTSIGResponse) *NullableGenerateTSIGResponse {
	return &NullableGenerateTSIGResponse{value: val, isSet: true}
}

func (v NullableGenerateTSIGResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateTSIGResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
