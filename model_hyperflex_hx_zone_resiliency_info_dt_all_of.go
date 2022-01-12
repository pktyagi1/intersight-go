/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4950
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// HyperflexHxZoneResiliencyInfoDtAllOf Definition of the list of properties defined in 'hyperflex.HxZoneResiliencyInfoDt', excluding properties defined in parent classes.
type HyperflexHxZoneResiliencyInfoDtAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the availability zone.
	Name *string `json:"Name,omitempty"`
	ResiliencyInfo NullableHyperflexHxResiliencyInfoDt `json:"ResiliencyInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxZoneResiliencyInfoDtAllOf HyperflexHxZoneResiliencyInfoDtAllOf

// NewHyperflexHxZoneResiliencyInfoDtAllOf instantiates a new HyperflexHxZoneResiliencyInfoDtAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxZoneResiliencyInfoDtAllOf(classId string, objectType string) *HyperflexHxZoneResiliencyInfoDtAllOf {
	this := HyperflexHxZoneResiliencyInfoDtAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxZoneResiliencyInfoDtAllOfWithDefaults instantiates a new HyperflexHxZoneResiliencyInfoDtAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxZoneResiliencyInfoDtAllOfWithDefaults() *HyperflexHxZoneResiliencyInfoDtAllOf {
	this := HyperflexHxZoneResiliencyInfoDtAllOf{}
	var classId string = "hyperflex.HxZoneResiliencyInfoDt"
	this.ClassId = classId
	var objectType string = "hyperflex.HxZoneResiliencyInfoDt"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) SetName(v string) {
	o.Name = &v
}

// GetResiliencyInfo returns the ResiliencyInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetResiliencyInfo() HyperflexHxResiliencyInfoDt {
	if o == nil || o.ResiliencyInfo.Get() == nil {
		var ret HyperflexHxResiliencyInfoDt
		return ret
	}
	return *o.ResiliencyInfo.Get()
}

// GetResiliencyInfoOk returns a tuple with the ResiliencyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) GetResiliencyInfoOk() (*HyperflexHxResiliencyInfoDt, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResiliencyInfo.Get(), o.ResiliencyInfo.IsSet()
}

// HasResiliencyInfo returns a boolean if a field has been set.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) HasResiliencyInfo() bool {
	if o != nil && o.ResiliencyInfo.IsSet() {
		return true
	}

	return false
}

// SetResiliencyInfo gets a reference to the given NullableHyperflexHxResiliencyInfoDt and assigns it to the ResiliencyInfo field.
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) SetResiliencyInfo(v HyperflexHxResiliencyInfoDt) {
	o.ResiliencyInfo.Set(&v)
}
// SetResiliencyInfoNil sets the value for ResiliencyInfo to be an explicit nil
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) SetResiliencyInfoNil() {
	o.ResiliencyInfo.Set(nil)
}

// UnsetResiliencyInfo ensures that no value is present for ResiliencyInfo, not even an explicit nil
func (o *HyperflexHxZoneResiliencyInfoDtAllOf) UnsetResiliencyInfo() {
	o.ResiliencyInfo.Unset()
}

func (o HyperflexHxZoneResiliencyInfoDtAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ResiliencyInfo.IsSet() {
		toSerialize["ResiliencyInfo"] = o.ResiliencyInfo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxZoneResiliencyInfoDtAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxZoneResiliencyInfoDtAllOf := _HyperflexHxZoneResiliencyInfoDtAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHxZoneResiliencyInfoDtAllOf); err == nil {
		*o = HyperflexHxZoneResiliencyInfoDtAllOf(varHyperflexHxZoneResiliencyInfoDtAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ResiliencyInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxZoneResiliencyInfoDtAllOf struct {
	value *HyperflexHxZoneResiliencyInfoDtAllOf
	isSet bool
}

func (v NullableHyperflexHxZoneResiliencyInfoDtAllOf) Get() *HyperflexHxZoneResiliencyInfoDtAllOf {
	return v.value
}

func (v *NullableHyperflexHxZoneResiliencyInfoDtAllOf) Set(val *HyperflexHxZoneResiliencyInfoDtAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxZoneResiliencyInfoDtAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxZoneResiliencyInfoDtAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxZoneResiliencyInfoDtAllOf(val *HyperflexHxZoneResiliencyInfoDtAllOf) *NullableHyperflexHxZoneResiliencyInfoDtAllOf {
	return &NullableHyperflexHxZoneResiliencyInfoDtAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxZoneResiliencyInfoDtAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxZoneResiliencyInfoDtAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


