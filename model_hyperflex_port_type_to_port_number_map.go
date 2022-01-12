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
	"reflect"
	"strings"
)

// HyperflexPortTypeToPortNumberMap Replication Ports Key = Control || Data Value = port number.
type HyperflexPortTypeToPortNumberMap struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Integer describing port type to port number map.
	I16 *int64 `json:"I16,omitempty"`
	// String describing port type to port number map.
	String *string `json:"String,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexPortTypeToPortNumberMap HyperflexPortTypeToPortNumberMap

// NewHyperflexPortTypeToPortNumberMap instantiates a new HyperflexPortTypeToPortNumberMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexPortTypeToPortNumberMap(classId string, objectType string) *HyperflexPortTypeToPortNumberMap {
	this := HyperflexPortTypeToPortNumberMap{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexPortTypeToPortNumberMapWithDefaults instantiates a new HyperflexPortTypeToPortNumberMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexPortTypeToPortNumberMapWithDefaults() *HyperflexPortTypeToPortNumberMap {
	this := HyperflexPortTypeToPortNumberMap{}
	var classId string = "hyperflex.PortTypeToPortNumberMap"
	this.ClassId = classId
	var objectType string = "hyperflex.PortTypeToPortNumberMap"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexPortTypeToPortNumberMap) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexPortTypeToPortNumberMap) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexPortTypeToPortNumberMap) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexPortTypeToPortNumberMap) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexPortTypeToPortNumberMap) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexPortTypeToPortNumberMap) SetObjectType(v string) {
	o.ObjectType = v
}

// GetI16 returns the I16 field value if set, zero value otherwise.
func (o *HyperflexPortTypeToPortNumberMap) GetI16() int64 {
	if o == nil || o.I16 == nil {
		var ret int64
		return ret
	}
	return *o.I16
}

// GetI16Ok returns a tuple with the I16 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexPortTypeToPortNumberMap) GetI16Ok() (*int64, bool) {
	if o == nil || o.I16 == nil {
		return nil, false
	}
	return o.I16, true
}

// HasI16 returns a boolean if a field has been set.
func (o *HyperflexPortTypeToPortNumberMap) HasI16() bool {
	if o != nil && o.I16 != nil {
		return true
	}

	return false
}

// SetI16 gets a reference to the given int64 and assigns it to the I16 field.
func (o *HyperflexPortTypeToPortNumberMap) SetI16(v int64) {
	o.I16 = &v
}

// GetString returns the String field value if set, zero value otherwise.
func (o *HyperflexPortTypeToPortNumberMap) GetString() string {
	if o == nil || o.String == nil {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexPortTypeToPortNumberMap) GetStringOk() (*string, bool) {
	if o == nil || o.String == nil {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *HyperflexPortTypeToPortNumberMap) HasString() bool {
	if o != nil && o.String != nil {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *HyperflexPortTypeToPortNumberMap) SetString(v string) {
	o.String = &v
}

func (o HyperflexPortTypeToPortNumberMap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.I16 != nil {
		toSerialize["I16"] = o.I16
	}
	if o.String != nil {
		toSerialize["String"] = o.String
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexPortTypeToPortNumberMap) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexPortTypeToPortNumberMapWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Integer describing port type to port number map.
		I16 *int64 `json:"I16,omitempty"`
		// String describing port type to port number map.
		String *string `json:"String,omitempty"`
	}

	varHyperflexPortTypeToPortNumberMapWithoutEmbeddedStruct := HyperflexPortTypeToPortNumberMapWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexPortTypeToPortNumberMapWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexPortTypeToPortNumberMap := _HyperflexPortTypeToPortNumberMap{}
		varHyperflexPortTypeToPortNumberMap.ClassId = varHyperflexPortTypeToPortNumberMapWithoutEmbeddedStruct.ClassId
		varHyperflexPortTypeToPortNumberMap.ObjectType = varHyperflexPortTypeToPortNumberMapWithoutEmbeddedStruct.ObjectType
		varHyperflexPortTypeToPortNumberMap.I16 = varHyperflexPortTypeToPortNumberMapWithoutEmbeddedStruct.I16
		varHyperflexPortTypeToPortNumberMap.String = varHyperflexPortTypeToPortNumberMapWithoutEmbeddedStruct.String
		*o = HyperflexPortTypeToPortNumberMap(varHyperflexPortTypeToPortNumberMap)
	} else {
		return err
	}

	varHyperflexPortTypeToPortNumberMap := _HyperflexPortTypeToPortNumberMap{}

	err = json.Unmarshal(bytes, &varHyperflexPortTypeToPortNumberMap)
	if err == nil {
		o.MoBaseComplexType = varHyperflexPortTypeToPortNumberMap.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "I16")
		delete(additionalProperties, "String")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexPortTypeToPortNumberMap struct {
	value *HyperflexPortTypeToPortNumberMap
	isSet bool
}

func (v NullableHyperflexPortTypeToPortNumberMap) Get() *HyperflexPortTypeToPortNumberMap {
	return v.value
}

func (v *NullableHyperflexPortTypeToPortNumberMap) Set(val *HyperflexPortTypeToPortNumberMap) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexPortTypeToPortNumberMap) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexPortTypeToPortNumberMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexPortTypeToPortNumberMap(val *HyperflexPortTypeToPortNumberMap) *NullableHyperflexPortTypeToPortNumberMap {
	return &NullableHyperflexPortTypeToPortNumberMap{value: val, isSet: true}
}

func (v NullableHyperflexPortTypeToPortNumberMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexPortTypeToPortNumberMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


