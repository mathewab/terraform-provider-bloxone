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

// checks if the Resolver type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Resolver{}

// Resolver struct for Resolver
type Resolver struct {
	// address that can be used as resolver
	Address *string `json:"address,omitempty"`
	// Mark it true to set default DNS resolvers that will be used in case if the BloxOne Cloud is unreachable.
	IsFallback *bool `json:"is_fallback,omitempty"`
	// Mark it true to set internal or local DNS servers' IPv4 or IPv6 addresses that are used as DNS resolvers
	IsLocal *bool `json:"is_local,omitempty"`
	// The list of DNS resolver communication protocols.
	Protocols []DNSProtocol `json:"protocols,omitempty"`
}

// NewResolver instantiates a new Resolver object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolver() *Resolver {
	this := Resolver{}
	return &this
}

// NewResolverWithDefaults instantiates a new Resolver object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolverWithDefaults() *Resolver {
	this := Resolver{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Resolver) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resolver) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Resolver) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *Resolver) SetAddress(v string) {
	o.Address = &v
}

// GetIsFallback returns the IsFallback field value if set, zero value otherwise.
func (o *Resolver) GetIsFallback() bool {
	if o == nil || IsNil(o.IsFallback) {
		var ret bool
		return ret
	}
	return *o.IsFallback
}

// GetIsFallbackOk returns a tuple with the IsFallback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resolver) GetIsFallbackOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFallback) {
		return nil, false
	}
	return o.IsFallback, true
}

// HasIsFallback returns a boolean if a field has been set.
func (o *Resolver) HasIsFallback() bool {
	if o != nil && !IsNil(o.IsFallback) {
		return true
	}

	return false
}

// SetIsFallback gets a reference to the given bool and assigns it to the IsFallback field.
func (o *Resolver) SetIsFallback(v bool) {
	o.IsFallback = &v
}

// GetIsLocal returns the IsLocal field value if set, zero value otherwise.
func (o *Resolver) GetIsLocal() bool {
	if o == nil || IsNil(o.IsLocal) {
		var ret bool
		return ret
	}
	return *o.IsLocal
}

// GetIsLocalOk returns a tuple with the IsLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resolver) GetIsLocalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocal) {
		return nil, false
	}
	return o.IsLocal, true
}

// HasIsLocal returns a boolean if a field has been set.
func (o *Resolver) HasIsLocal() bool {
	if o != nil && !IsNil(o.IsLocal) {
		return true
	}

	return false
}

// SetIsLocal gets a reference to the given bool and assigns it to the IsLocal field.
func (o *Resolver) SetIsLocal(v bool) {
	o.IsLocal = &v
}

// GetProtocols returns the Protocols field value if set, zero value otherwise.
func (o *Resolver) GetProtocols() []DNSProtocol {
	if o == nil || IsNil(o.Protocols) {
		var ret []DNSProtocol
		return ret
	}
	return o.Protocols
}

// GetProtocolsOk returns a tuple with the Protocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resolver) GetProtocolsOk() ([]DNSProtocol, bool) {
	if o == nil || IsNil(o.Protocols) {
		return nil, false
	}
	return o.Protocols, true
}

// HasProtocols returns a boolean if a field has been set.
func (o *Resolver) HasProtocols() bool {
	if o != nil && !IsNil(o.Protocols) {
		return true
	}

	return false
}

// SetProtocols gets a reference to the given []DNSProtocol and assigns it to the Protocols field.
func (o *Resolver) SetProtocols(v []DNSProtocol) {
	o.Protocols = v
}

func (o Resolver) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Resolver) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.IsFallback) {
		toSerialize["is_fallback"] = o.IsFallback
	}
	if !IsNil(o.IsLocal) {
		toSerialize["is_local"] = o.IsLocal
	}
	if !IsNil(o.Protocols) {
		toSerialize["protocols"] = o.Protocols
	}
	return toSerialize, nil
}

type NullableResolver struct {
	value *Resolver
	isSet bool
}

func (v NullableResolver) Get() *Resolver {
	return v.value
}

func (v *NullableResolver) Set(val *Resolver) {
	v.value = val
	v.isSet = true
}

func (v NullableResolver) IsSet() bool {
	return v.isSet
}

func (v *NullableResolver) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolver(val *Resolver) *NullableResolver {
	return &NullableResolver{value: val, isSet: true}
}

func (v NullableResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolver) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
