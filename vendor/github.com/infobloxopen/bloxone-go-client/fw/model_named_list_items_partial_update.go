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

// checks if the NamedListItemsPartialUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamedListItemsPartialUpdate{}

// NamedListItemsPartialUpdate struct for NamedListItemsPartialUpdate
type NamedListItemsPartialUpdate struct {
	// The List of ItemStructs structure which contains the item and its description
	DeletedItemsDescribed []ItemStructs `json:"deleted_items_described,omitempty"`
	// The Named List object identifier.
	Id *int32 `json:"id,omitempty"`
	// The List of ItemStructs structure which contains the item and its description
	InsertedItemsDescribed []ItemStructs `json:"inserted_items_described,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _NamedListItemsPartialUpdate NamedListItemsPartialUpdate

// NewNamedListItemsPartialUpdate instantiates a new NamedListItemsPartialUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamedListItemsPartialUpdate() *NamedListItemsPartialUpdate {
	this := NamedListItemsPartialUpdate{}
	return &this
}

// NewNamedListItemsPartialUpdateWithDefaults instantiates a new NamedListItemsPartialUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamedListItemsPartialUpdateWithDefaults() *NamedListItemsPartialUpdate {
	this := NamedListItemsPartialUpdate{}
	return &this
}

// GetDeletedItemsDescribed returns the DeletedItemsDescribed field value if set, zero value otherwise.
func (o *NamedListItemsPartialUpdate) GetDeletedItemsDescribed() []ItemStructs {
	if o == nil || IsNil(o.DeletedItemsDescribed) {
		var ret []ItemStructs
		return ret
	}
	return o.DeletedItemsDescribed
}

// GetDeletedItemsDescribedOk returns a tuple with the DeletedItemsDescribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamedListItemsPartialUpdate) GetDeletedItemsDescribedOk() ([]ItemStructs, bool) {
	if o == nil || IsNil(o.DeletedItemsDescribed) {
		return nil, false
	}
	return o.DeletedItemsDescribed, true
}

// HasDeletedItemsDescribed returns a boolean if a field has been set.
func (o *NamedListItemsPartialUpdate) HasDeletedItemsDescribed() bool {
	if o != nil && !IsNil(o.DeletedItemsDescribed) {
		return true
	}

	return false
}

// SetDeletedItemsDescribed gets a reference to the given []ItemStructs and assigns it to the DeletedItemsDescribed field.
func (o *NamedListItemsPartialUpdate) SetDeletedItemsDescribed(v []ItemStructs) {
	o.DeletedItemsDescribed = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NamedListItemsPartialUpdate) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamedListItemsPartialUpdate) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NamedListItemsPartialUpdate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NamedListItemsPartialUpdate) SetId(v int32) {
	o.Id = &v
}

// GetInsertedItemsDescribed returns the InsertedItemsDescribed field value if set, zero value otherwise.
func (o *NamedListItemsPartialUpdate) GetInsertedItemsDescribed() []ItemStructs {
	if o == nil || IsNil(o.InsertedItemsDescribed) {
		var ret []ItemStructs
		return ret
	}
	return o.InsertedItemsDescribed
}

// GetInsertedItemsDescribedOk returns a tuple with the InsertedItemsDescribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamedListItemsPartialUpdate) GetInsertedItemsDescribedOk() ([]ItemStructs, bool) {
	if o == nil || IsNil(o.InsertedItemsDescribed) {
		return nil, false
	}
	return o.InsertedItemsDescribed, true
}

// HasInsertedItemsDescribed returns a boolean if a field has been set.
func (o *NamedListItemsPartialUpdate) HasInsertedItemsDescribed() bool {
	if o != nil && !IsNil(o.InsertedItemsDescribed) {
		return true
	}

	return false
}

// SetInsertedItemsDescribed gets a reference to the given []ItemStructs and assigns it to the InsertedItemsDescribed field.
func (o *NamedListItemsPartialUpdate) SetInsertedItemsDescribed(v []ItemStructs) {
	o.InsertedItemsDescribed = v
}

func (o NamedListItemsPartialUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamedListItemsPartialUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeletedItemsDescribed) {
		toSerialize["deleted_items_described"] = o.DeletedItemsDescribed
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InsertedItemsDescribed) {
		toSerialize["inserted_items_described"] = o.InsertedItemsDescribed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NamedListItemsPartialUpdate) UnmarshalJSON(data []byte) (err error) {
	varNamedListItemsPartialUpdate := _NamedListItemsPartialUpdate{}

	err = json.Unmarshal(data, &varNamedListItemsPartialUpdate)

	if err != nil {
		return err
	}

	*o = NamedListItemsPartialUpdate(varNamedListItemsPartialUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deleted_items_described")
		delete(additionalProperties, "id")
		delete(additionalProperties, "inserted_items_described")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNamedListItemsPartialUpdate struct {
	value *NamedListItemsPartialUpdate
	isSet bool
}

func (v NullableNamedListItemsPartialUpdate) Get() *NamedListItemsPartialUpdate {
	return v.value
}

func (v *NullableNamedListItemsPartialUpdate) Set(val *NamedListItemsPartialUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableNamedListItemsPartialUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableNamedListItemsPartialUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamedListItemsPartialUpdate(val *NamedListItemsPartialUpdate) *NullableNamedListItemsPartialUpdate {
	return &NullableNamedListItemsPartialUpdate{value: val, isSet: true}
}

func (v NullableNamedListItemsPartialUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamedListItemsPartialUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
