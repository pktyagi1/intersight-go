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

// InfraMetaDataAllOf Definition of the list of properties defined in 'infra.MetaData', excluding properties defined in parent classes.
type InfraMetaDataAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the meta property which identifies a specific resource.
	Name *string `json:"Name,omitempty"`
	// Value of the meta property which identifies a specific resource.
	Value *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InfraMetaDataAllOf InfraMetaDataAllOf

// NewInfraMetaDataAllOf instantiates a new InfraMetaDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraMetaDataAllOf(classId string, objectType string) *InfraMetaDataAllOf {
	this := InfraMetaDataAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInfraMetaDataAllOfWithDefaults instantiates a new InfraMetaDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraMetaDataAllOfWithDefaults() *InfraMetaDataAllOf {
	this := InfraMetaDataAllOf{}
	var classId string = "infra.MetaData"
	this.ClassId = classId
	var objectType string = "infra.MetaData"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *InfraMetaDataAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InfraMetaDataAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InfraMetaDataAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *InfraMetaDataAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InfraMetaDataAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InfraMetaDataAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InfraMetaDataAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraMetaDataAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InfraMetaDataAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InfraMetaDataAllOf) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InfraMetaDataAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraMetaDataAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InfraMetaDataAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InfraMetaDataAllOf) SetValue(v string) {
	o.Value = &v
}

func (o InfraMetaDataAllOf) MarshalJSON() ([]byte, error) {
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
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InfraMetaDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varInfraMetaDataAllOf := _InfraMetaDataAllOf{}

	if err = json.Unmarshal(bytes, &varInfraMetaDataAllOf); err == nil {
		*o = InfraMetaDataAllOf(varInfraMetaDataAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInfraMetaDataAllOf struct {
	value *InfraMetaDataAllOf
	isSet bool
}

func (v NullableInfraMetaDataAllOf) Get() *InfraMetaDataAllOf {
	return v.value
}

func (v *NullableInfraMetaDataAllOf) Set(val *InfraMetaDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraMetaDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraMetaDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraMetaDataAllOf(val *InfraMetaDataAllOf) *NullableInfraMetaDataAllOf {
	return &NullableInfraMetaDataAllOf{value: val, isSet: true}
}

func (v NullableInfraMetaDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraMetaDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


