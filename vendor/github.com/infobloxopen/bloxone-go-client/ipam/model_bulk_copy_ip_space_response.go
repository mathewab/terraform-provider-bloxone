/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the BulkCopyIPSpaceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkCopyIPSpaceResponse{}

// BulkCopyIPSpaceResponse struct for BulkCopyIPSpaceResponse
type BulkCopyIPSpaceResponse struct {
	Errors               []BulkCopyError `json:"errors,omitempty"`
	Results              []CopyResponse  `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkCopyIPSpaceResponse BulkCopyIPSpaceResponse

// NewBulkCopyIPSpaceResponse instantiates a new BulkCopyIPSpaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkCopyIPSpaceResponse() *BulkCopyIPSpaceResponse {
	this := BulkCopyIPSpaceResponse{}
	return &this
}

// NewBulkCopyIPSpaceResponseWithDefaults instantiates a new BulkCopyIPSpaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkCopyIPSpaceResponseWithDefaults() *BulkCopyIPSpaceResponse {
	this := BulkCopyIPSpaceResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BulkCopyIPSpaceResponse) GetErrors() []BulkCopyError {
	if o == nil || IsNil(o.Errors) {
		var ret []BulkCopyError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCopyIPSpaceResponse) GetErrorsOk() ([]BulkCopyError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BulkCopyIPSpaceResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []BulkCopyError and assigns it to the Errors field.
func (o *BulkCopyIPSpaceResponse) SetErrors(v []BulkCopyError) {
	o.Errors = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *BulkCopyIPSpaceResponse) GetResults() []CopyResponse {
	if o == nil || IsNil(o.Results) {
		var ret []CopyResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCopyIPSpaceResponse) GetResultsOk() ([]CopyResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BulkCopyIPSpaceResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CopyResponse and assigns it to the Results field.
func (o *BulkCopyIPSpaceResponse) SetResults(v []CopyResponse) {
	o.Results = v
}

func (o BulkCopyIPSpaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkCopyIPSpaceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkCopyIPSpaceResponse) UnmarshalJSON(data []byte) (err error) {
	varBulkCopyIPSpaceResponse := _BulkCopyIPSpaceResponse{}

	err = json.Unmarshal(data, &varBulkCopyIPSpaceResponse)

	if err != nil {
		return err
	}

	*o = BulkCopyIPSpaceResponse(varBulkCopyIPSpaceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errors")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkCopyIPSpaceResponse struct {
	value *BulkCopyIPSpaceResponse
	isSet bool
}

func (v NullableBulkCopyIPSpaceResponse) Get() *BulkCopyIPSpaceResponse {
	return v.value
}

func (v *NullableBulkCopyIPSpaceResponse) Set(val *BulkCopyIPSpaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkCopyIPSpaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkCopyIPSpaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkCopyIPSpaceResponse(val *BulkCopyIPSpaceResponse) *NullableBulkCopyIPSpaceResponse {
	return &NullableBulkCopyIPSpaceResponse{value: val, isSet: true}
}

func (v NullableBulkCopyIPSpaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkCopyIPSpaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
