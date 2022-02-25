/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5517
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ContentComplexType If the given API/device response is a collection of items, each item and its properties can be modeled as a complex type. The types are uniquely named within the grammar and provides the list of parameters to be extracted from each item. Name of the complex type can be used as the type of parameter that represents the complex value.
type ContentComplexType struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The unique name of this complex type within the grammar specification.
	Name *string `json:"Name,omitempty"`
	Parameters []ContentBaseParameter `json:"Parameters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContentComplexType ContentComplexType

// NewContentComplexType instantiates a new ContentComplexType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentComplexType(classId string, objectType string) *ContentComplexType {
	this := ContentComplexType{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewContentComplexTypeWithDefaults instantiates a new ContentComplexType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentComplexTypeWithDefaults() *ContentComplexType {
	this := ContentComplexType{}
	var classId string = "content.ComplexType"
	this.ClassId = classId
	var objectType string = "content.ComplexType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ContentComplexType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ContentComplexType) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ContentComplexType) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ContentComplexType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ContentComplexType) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ContentComplexType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContentComplexType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentComplexType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContentComplexType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContentComplexType) SetName(v string) {
	o.Name = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentComplexType) GetParameters() []ContentBaseParameter {
	if o == nil  {
		var ret []ContentBaseParameter
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentComplexType) GetParametersOk() (*[]ContentBaseParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ContentComplexType) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ContentBaseParameter and assigns it to the Parameters field.
func (o *ContentComplexType) SetParameters(v []ContentBaseParameter) {
	o.Parameters = v
}

func (o ContentComplexType) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Parameters != nil {
		toSerialize["Parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContentComplexType) UnmarshalJSON(bytes []byte) (err error) {
	type ContentComplexTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The unique name of this complex type within the grammar specification.
		Name *string `json:"Name,omitempty"`
		Parameters []ContentBaseParameter `json:"Parameters,omitempty"`
	}

	varContentComplexTypeWithoutEmbeddedStruct := ContentComplexTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varContentComplexTypeWithoutEmbeddedStruct)
	if err == nil {
		varContentComplexType := _ContentComplexType{}
		varContentComplexType.ClassId = varContentComplexTypeWithoutEmbeddedStruct.ClassId
		varContentComplexType.ObjectType = varContentComplexTypeWithoutEmbeddedStruct.ObjectType
		varContentComplexType.Name = varContentComplexTypeWithoutEmbeddedStruct.Name
		varContentComplexType.Parameters = varContentComplexTypeWithoutEmbeddedStruct.Parameters
		*o = ContentComplexType(varContentComplexType)
	} else {
		return err
	}

	varContentComplexType := _ContentComplexType{}

	err = json.Unmarshal(bytes, &varContentComplexType)
	if err == nil {
		o.MoBaseComplexType = varContentComplexType.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Parameters")

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

type NullableContentComplexType struct {
	value *ContentComplexType
	isSet bool
}

func (v NullableContentComplexType) Get() *ContentComplexType {
	return v.value
}

func (v *NullableContentComplexType) Set(val *ContentComplexType) {
	v.value = val
	v.isSet = true
}

func (v NullableContentComplexType) IsSet() bool {
	return v.isSet
}

func (v *NullableContentComplexType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentComplexType(val *ContentComplexType) *NullableContentComplexType {
	return &NullableContentComplexType{value: val, isSet: true}
}

func (v NullableContentComplexType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentComplexType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


