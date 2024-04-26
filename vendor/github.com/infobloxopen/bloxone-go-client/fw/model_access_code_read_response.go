/*
BloxOne FW API

BloxOne Threat Defense Cloud is an extension of the BloxOne Suite that provides visibility into infected and compromised off-premises devices, roaming users, remote sites, and branch offices. You can subscribe to BloxOne Cloud and use its functionality to mitigate and control malware as well as provide unprecedented insight into your network security posture and enable timely action. BloxOne Cloud also offers unified policy management, reporting, and threat analytics across the entire spectrum. Using automated and high-quality threat intelligence feeds and unique behavioral analytics, it automatically stops device communications with C&Cs/botnets and prevents DNS based data exfiltration.  The mission-critical DNS infrastructure can become a vulnerable component in your network when it is inadequately protected by traditional security solutions and consequently used as an attack surface. Compromised DNS services can result in catastrophic network and system failures. To fully protect your network in today’s cyber security threat environment, Infoblox sets a new DNS security standard by offering scalable, enterprise-grade, and integrated protection for your DNS infrastructure.  Through the Infoblox Cloud Services Portal, you can view the status of your subscription and threat intelligence feeds, manage your network scope and roaming end users, and learn more about threats on your networks through the Infoblox Threat Lookup tool and predefined reports.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fw

import (
	"encoding/json"
)

// checks if the AccessCodeReadResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessCodeReadResponse{}

// AccessCodeReadResponse struct for AccessCodeReadResponse
type AccessCodeReadResponse struct {
	Results              *AccessCode `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessCodeReadResponse AccessCodeReadResponse

// NewAccessCodeReadResponse instantiates a new AccessCodeReadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessCodeReadResponse() *AccessCodeReadResponse {
	this := AccessCodeReadResponse{}
	return &this
}

// NewAccessCodeReadResponseWithDefaults instantiates a new AccessCodeReadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessCodeReadResponseWithDefaults() *AccessCodeReadResponse {
	this := AccessCodeReadResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *AccessCodeReadResponse) GetResults() AccessCode {
	if o == nil || IsNil(o.Results) {
		var ret AccessCode
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCodeReadResponse) GetResultsOk() (*AccessCode, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *AccessCodeReadResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given AccessCode and assigns it to the Results field.
func (o *AccessCodeReadResponse) SetResults(v AccessCode) {
	o.Results = &v
}

func (o AccessCodeReadResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessCodeReadResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessCodeReadResponse) UnmarshalJSON(data []byte) (err error) {
	varAccessCodeReadResponse := _AccessCodeReadResponse{}

	err = json.Unmarshal(data, &varAccessCodeReadResponse)

	if err != nil {
		return err
	}

	*o = AccessCodeReadResponse(varAccessCodeReadResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessCodeReadResponse struct {
	value *AccessCodeReadResponse
	isSet bool
}

func (v NullableAccessCodeReadResponse) Get() *AccessCodeReadResponse {
	return v.value
}

func (v *NullableAccessCodeReadResponse) Set(val *AccessCodeReadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessCodeReadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessCodeReadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessCodeReadResponse(val *AccessCodeReadResponse) *NullableAccessCodeReadResponse {
	return &NullableAccessCodeReadResponse{value: val, isSet: true}
}

func (v NullableAccessCodeReadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessCodeReadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
