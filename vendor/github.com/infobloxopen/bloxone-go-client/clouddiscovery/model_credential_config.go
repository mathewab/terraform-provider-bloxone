/*
Discovery Configuration API V2

The Discovery configuration service is a BloxOne Service that provides configuration for accessing and syncing the Cloud assets   Base Paths:  1. provider: **_/api/cloud_discovery/v2/_**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clouddiscovery

import (
	"encoding/json"
)

// checks if the CredentialConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialConfig{}

// CredentialConfig struct for CredentialConfig
type CredentialConfig struct {
	AccessIdentifier     *string `json:"access_identifier,omitempty"`
	Enclave              *string `json:"enclave,omitempty"`
	Region               *string `json:"region,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CredentialConfig CredentialConfig

// NewCredentialConfig instantiates a new CredentialConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialConfig() *CredentialConfig {
	this := CredentialConfig{}
	return &this
}

// NewCredentialConfigWithDefaults instantiates a new CredentialConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialConfigWithDefaults() *CredentialConfig {
	this := CredentialConfig{}
	return &this
}

// GetAccessIdentifier returns the AccessIdentifier field value if set, zero value otherwise.
func (o *CredentialConfig) GetAccessIdentifier() string {
	if o == nil || IsNil(o.AccessIdentifier) {
		var ret string
		return ret
	}
	return *o.AccessIdentifier
}

// GetAccessIdentifierOk returns a tuple with the AccessIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialConfig) GetAccessIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.AccessIdentifier) {
		return nil, false
	}
	return o.AccessIdentifier, true
}

// HasAccessIdentifier returns a boolean if a field has been set.
func (o *CredentialConfig) HasAccessIdentifier() bool {
	if o != nil && !IsNil(o.AccessIdentifier) {
		return true
	}

	return false
}

// SetAccessIdentifier gets a reference to the given string and assigns it to the AccessIdentifier field.
func (o *CredentialConfig) SetAccessIdentifier(v string) {
	o.AccessIdentifier = &v
}

// GetEnclave returns the Enclave field value if set, zero value otherwise.
func (o *CredentialConfig) GetEnclave() string {
	if o == nil || IsNil(o.Enclave) {
		var ret string
		return ret
	}
	return *o.Enclave
}

// GetEnclaveOk returns a tuple with the Enclave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialConfig) GetEnclaveOk() (*string, bool) {
	if o == nil || IsNil(o.Enclave) {
		return nil, false
	}
	return o.Enclave, true
}

// HasEnclave returns a boolean if a field has been set.
func (o *CredentialConfig) HasEnclave() bool {
	if o != nil && !IsNil(o.Enclave) {
		return true
	}

	return false
}

// SetEnclave gets a reference to the given string and assigns it to the Enclave field.
func (o *CredentialConfig) SetEnclave(v string) {
	o.Enclave = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CredentialConfig) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialConfig) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CredentialConfig) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CredentialConfig) SetRegion(v string) {
	o.Region = &v
}

func (o CredentialConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessIdentifier) {
		toSerialize["access_identifier"] = o.AccessIdentifier
	}
	if !IsNil(o.Enclave) {
		toSerialize["enclave"] = o.Enclave
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CredentialConfig) UnmarshalJSON(data []byte) (err error) {
	varCredentialConfig := _CredentialConfig{}

	err = json.Unmarshal(data, &varCredentialConfig)

	if err != nil {
		return err
	}

	*o = CredentialConfig(varCredentialConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access_identifier")
		delete(additionalProperties, "enclave")
		delete(additionalProperties, "region")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCredentialConfig struct {
	value *CredentialConfig
	isSet bool
}

func (v NullableCredentialConfig) Get() *CredentialConfig {
	return v.value
}

func (v *NullableCredentialConfig) Set(val *CredentialConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialConfig(val *CredentialConfig) *NullableCredentialConfig {
	return &NullableCredentialConfig{value: val, isSet: true}
}

func (v NullableCredentialConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
