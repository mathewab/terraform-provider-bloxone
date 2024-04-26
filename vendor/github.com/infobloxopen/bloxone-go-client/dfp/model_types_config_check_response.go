/*
DFP API

BloxOne Cloud is a SaaS offering designed to provide protection to devices on and off-premises, including roaming, remote, and branch offices. It provides visibility into infected and compromised devices, prevents DNS-based data exfiltration, and automatically stops device communications with command-and-control servers (C&Cs) and botnets, in addition to providing recursive DNS services in the cloud. You can access the services by deploying the BloxOne Endpoint agent or the DNS forwarding proxy.  For remote office deployments or in cases where installing an endpoint agent is not desirable or possible, you can use the DNS forwarding proxy. It is a software that runs on bare-metal, VM infrastructures, or Infoblox NIOS appliances; and it embeds the client IPs in DNS queries before forwarding them to BloxOne Cloud. The communications are encrypted and client visibility is maintained. The proxy also provides DNS resolution to local DNS zones when you configure local resolvers. Once you set up a DNS forwarding proxy, it becomes the main DNS server for your remote site. It will also cache responses to speed resolution of future queries.  By implementing the DNS forwarding proxy, you can rest assured that BloxOne Cloud effectively enforces DNS client-based security policies at your remote sites. On-premises devices that send DNS queries reveal their actual client IP addresses (instead of their NAT IP address), which allows BloxOne Cloud to apply the security policies applicable to the respective endpoints and identify infected clients.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfp

import (
	"encoding/json"
)

// checks if the TypesConfigCheckResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypesConfigCheckResponse{}

// TypesConfigCheckResponse The Config Check response.
type TypesConfigCheckResponse struct {
	// The list of check result.
	Results              []TypesConfigCheckResult `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TypesConfigCheckResponse TypesConfigCheckResponse

// NewTypesConfigCheckResponse instantiates a new TypesConfigCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesConfigCheckResponse() *TypesConfigCheckResponse {
	this := TypesConfigCheckResponse{}
	return &this
}

// NewTypesConfigCheckResponseWithDefaults instantiates a new TypesConfigCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesConfigCheckResponseWithDefaults() *TypesConfigCheckResponse {
	this := TypesConfigCheckResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *TypesConfigCheckResponse) GetResults() []TypesConfigCheckResult {
	if o == nil || IsNil(o.Results) {
		var ret []TypesConfigCheckResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesConfigCheckResponse) GetResultsOk() ([]TypesConfigCheckResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *TypesConfigCheckResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []TypesConfigCheckResult and assigns it to the Results field.
func (o *TypesConfigCheckResponse) SetResults(v []TypesConfigCheckResult) {
	o.Results = v
}

func (o TypesConfigCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypesConfigCheckResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TypesConfigCheckResponse) UnmarshalJSON(data []byte) (err error) {
	varTypesConfigCheckResponse := _TypesConfigCheckResponse{}

	err = json.Unmarshal(data, &varTypesConfigCheckResponse)

	if err != nil {
		return err
	}

	*o = TypesConfigCheckResponse(varTypesConfigCheckResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTypesConfigCheckResponse struct {
	value *TypesConfigCheckResponse
	isSet bool
}

func (v NullableTypesConfigCheckResponse) Get() *TypesConfigCheckResponse {
	return v.value
}

func (v *NullableTypesConfigCheckResponse) Set(val *TypesConfigCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesConfigCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesConfigCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesConfigCheckResponse(val *TypesConfigCheckResponse) *NullableTypesConfigCheckResponse {
	return &NullableTypesConfigCheckResponse{value: val, isSet: true}
}

func (v NullableTypesConfigCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesConfigCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
