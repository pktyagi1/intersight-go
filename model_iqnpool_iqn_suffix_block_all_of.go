/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// IqnpoolIqnSuffixBlockAllOf Definition of the list of properties defined in 'iqnpool.IqnSuffixBlock', excluding properties defined in parent classes.
type IqnpoolIqnSuffixBlockAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The first suffix number in the block.
	From *int64 `json:"From,omitempty"`
	// The suffix for this block of IQNs.
	Suffix *string `json:"Suffix,omitempty"`
	// The last suffix number in the block.
	To *int64 `json:"To,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IqnpoolIqnSuffixBlockAllOf IqnpoolIqnSuffixBlockAllOf

// NewIqnpoolIqnSuffixBlockAllOf instantiates a new IqnpoolIqnSuffixBlockAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIqnpoolIqnSuffixBlockAllOf(classId string, objectType string) *IqnpoolIqnSuffixBlockAllOf {
	this := IqnpoolIqnSuffixBlockAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIqnpoolIqnSuffixBlockAllOfWithDefaults instantiates a new IqnpoolIqnSuffixBlockAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIqnpoolIqnSuffixBlockAllOfWithDefaults() *IqnpoolIqnSuffixBlockAllOf {
	this := IqnpoolIqnSuffixBlockAllOf{}
	var classId string = "iqnpool.IqnSuffixBlock"
	this.ClassId = classId
	var objectType string = "iqnpool.IqnSuffixBlock"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IqnpoolIqnSuffixBlockAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IqnpoolIqnSuffixBlockAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IqnpoolIqnSuffixBlockAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IqnpoolIqnSuffixBlockAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IqnpoolIqnSuffixBlockAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IqnpoolIqnSuffixBlockAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *IqnpoolIqnSuffixBlockAllOf) GetFrom() int64 {
	if o == nil || o.From == nil {
		var ret int64
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolIqnSuffixBlockAllOf) GetFromOk() (*int64, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *IqnpoolIqnSuffixBlockAllOf) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int64 and assigns it to the From field.
func (o *IqnpoolIqnSuffixBlockAllOf) SetFrom(v int64) {
	o.From = &v
}

// GetSuffix returns the Suffix field value if set, zero value otherwise.
func (o *IqnpoolIqnSuffixBlockAllOf) GetSuffix() string {
	if o == nil || o.Suffix == nil {
		var ret string
		return ret
	}
	return *o.Suffix
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolIqnSuffixBlockAllOf) GetSuffixOk() (*string, bool) {
	if o == nil || o.Suffix == nil {
		return nil, false
	}
	return o.Suffix, true
}

// HasSuffix returns a boolean if a field has been set.
func (o *IqnpoolIqnSuffixBlockAllOf) HasSuffix() bool {
	if o != nil && o.Suffix != nil {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given string and assigns it to the Suffix field.
func (o *IqnpoolIqnSuffixBlockAllOf) SetSuffix(v string) {
	o.Suffix = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *IqnpoolIqnSuffixBlockAllOf) GetTo() int64 {
	if o == nil || o.To == nil {
		var ret int64
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolIqnSuffixBlockAllOf) GetToOk() (*int64, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *IqnpoolIqnSuffixBlockAllOf) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given int64 and assigns it to the To field.
func (o *IqnpoolIqnSuffixBlockAllOf) SetTo(v int64) {
	o.To = &v
}

func (o IqnpoolIqnSuffixBlockAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.From != nil {
		toSerialize["From"] = o.From
	}
	if o.Suffix != nil {
		toSerialize["Suffix"] = o.Suffix
	}
	if o.To != nil {
		toSerialize["To"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IqnpoolIqnSuffixBlockAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIqnpoolIqnSuffixBlockAllOf := _IqnpoolIqnSuffixBlockAllOf{}

	if err = json.Unmarshal(bytes, &varIqnpoolIqnSuffixBlockAllOf); err == nil {
		*o = IqnpoolIqnSuffixBlockAllOf(varIqnpoolIqnSuffixBlockAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "From")
		delete(additionalProperties, "Suffix")
		delete(additionalProperties, "To")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIqnpoolIqnSuffixBlockAllOf struct {
	value *IqnpoolIqnSuffixBlockAllOf
	isSet bool
}

func (v NullableIqnpoolIqnSuffixBlockAllOf) Get() *IqnpoolIqnSuffixBlockAllOf {
	return v.value
}

func (v *NullableIqnpoolIqnSuffixBlockAllOf) Set(val *IqnpoolIqnSuffixBlockAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIqnpoolIqnSuffixBlockAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIqnpoolIqnSuffixBlockAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIqnpoolIqnSuffixBlockAllOf(val *IqnpoolIqnSuffixBlockAllOf) *NullableIqnpoolIqnSuffixBlockAllOf {
	return &NullableIqnpoolIqnSuffixBlockAllOf{value: val, isSet: true}
}

func (v NullableIqnpoolIqnSuffixBlockAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIqnpoolIqnSuffixBlockAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


