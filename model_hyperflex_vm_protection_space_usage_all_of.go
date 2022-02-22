/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5313
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// HyperflexVmProtectionSpaceUsageAllOf Definition of the list of properties defined in 'hyperflex.VmProtectionSpaceUsage', excluding properties defined in parent classes.
type HyperflexVmProtectionSpaceUsageAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Space usage of the Virtual Machine from StDataServiceManager.
	SpaceUsage *int64 `json:"SpaceUsage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexVmProtectionSpaceUsageAllOf HyperflexVmProtectionSpaceUsageAllOf

// NewHyperflexVmProtectionSpaceUsageAllOf instantiates a new HyperflexVmProtectionSpaceUsageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVmProtectionSpaceUsageAllOf(classId string, objectType string) *HyperflexVmProtectionSpaceUsageAllOf {
	this := HyperflexVmProtectionSpaceUsageAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexVmProtectionSpaceUsageAllOfWithDefaults instantiates a new HyperflexVmProtectionSpaceUsageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVmProtectionSpaceUsageAllOfWithDefaults() *HyperflexVmProtectionSpaceUsageAllOf {
	this := HyperflexVmProtectionSpaceUsageAllOf{}
	var classId string = "hyperflex.VmProtectionSpaceUsage"
	this.ClassId = classId
	var objectType string = "hyperflex.VmProtectionSpaceUsage"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexVmProtectionSpaceUsageAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexVmProtectionSpaceUsageAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexVmProtectionSpaceUsageAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexVmProtectionSpaceUsageAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexVmProtectionSpaceUsageAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexVmProtectionSpaceUsageAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSpaceUsage returns the SpaceUsage field value if set, zero value otherwise.
func (o *HyperflexVmProtectionSpaceUsageAllOf) GetSpaceUsage() int64 {
	if o == nil || o.SpaceUsage == nil {
		var ret int64
		return ret
	}
	return *o.SpaceUsage
}

// GetSpaceUsageOk returns a tuple with the SpaceUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmProtectionSpaceUsageAllOf) GetSpaceUsageOk() (*int64, bool) {
	if o == nil || o.SpaceUsage == nil {
		return nil, false
	}
	return o.SpaceUsage, true
}

// HasSpaceUsage returns a boolean if a field has been set.
func (o *HyperflexVmProtectionSpaceUsageAllOf) HasSpaceUsage() bool {
	if o != nil && o.SpaceUsage != nil {
		return true
	}

	return false
}

// SetSpaceUsage gets a reference to the given int64 and assigns it to the SpaceUsage field.
func (o *HyperflexVmProtectionSpaceUsageAllOf) SetSpaceUsage(v int64) {
	o.SpaceUsage = &v
}

func (o HyperflexVmProtectionSpaceUsageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.SpaceUsage != nil {
		toSerialize["SpaceUsage"] = o.SpaceUsage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexVmProtectionSpaceUsageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexVmProtectionSpaceUsageAllOf := _HyperflexVmProtectionSpaceUsageAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexVmProtectionSpaceUsageAllOf); err == nil {
		*o = HyperflexVmProtectionSpaceUsageAllOf(varHyperflexVmProtectionSpaceUsageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SpaceUsage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexVmProtectionSpaceUsageAllOf struct {
	value *HyperflexVmProtectionSpaceUsageAllOf
	isSet bool
}

func (v NullableHyperflexVmProtectionSpaceUsageAllOf) Get() *HyperflexVmProtectionSpaceUsageAllOf {
	return v.value
}

func (v *NullableHyperflexVmProtectionSpaceUsageAllOf) Set(val *HyperflexVmProtectionSpaceUsageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVmProtectionSpaceUsageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVmProtectionSpaceUsageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVmProtectionSpaceUsageAllOf(val *HyperflexVmProtectionSpaceUsageAllOf) *NullableHyperflexVmProtectionSpaceUsageAllOf {
	return &NullableHyperflexVmProtectionSpaceUsageAllOf{value: val, isSet: true}
}

func (v NullableHyperflexVmProtectionSpaceUsageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVmProtectionSpaceUsageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


