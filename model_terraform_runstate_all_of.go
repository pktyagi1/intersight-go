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

// TerraformRunstateAllOf Definition of the list of properties defined in 'terraform.Runstate', excluding properties defined in parent classes.
type TerraformRunstateAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Run identifier for every terraform execution.
	RunId *string `json:"RunId,omitempty"`
	// StateFile identifier of terraform execution.
	StateFile *string `json:"StateFile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TerraformRunstateAllOf TerraformRunstateAllOf

// NewTerraformRunstateAllOf instantiates a new TerraformRunstateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerraformRunstateAllOf(classId string, objectType string) *TerraformRunstateAllOf {
	this := TerraformRunstateAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTerraformRunstateAllOfWithDefaults instantiates a new TerraformRunstateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerraformRunstateAllOfWithDefaults() *TerraformRunstateAllOf {
	this := TerraformRunstateAllOf{}
	var classId string = "terraform.Runstate"
	this.ClassId = classId
	var objectType string = "terraform.Runstate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TerraformRunstateAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TerraformRunstateAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TerraformRunstateAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TerraformRunstateAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TerraformRunstateAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TerraformRunstateAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRunId returns the RunId field value if set, zero value otherwise.
func (o *TerraformRunstateAllOf) GetRunId() string {
	if o == nil || o.RunId == nil {
		var ret string
		return ret
	}
	return *o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformRunstateAllOf) GetRunIdOk() (*string, bool) {
	if o == nil || o.RunId == nil {
		return nil, false
	}
	return o.RunId, true
}

// HasRunId returns a boolean if a field has been set.
func (o *TerraformRunstateAllOf) HasRunId() bool {
	if o != nil && o.RunId != nil {
		return true
	}

	return false
}

// SetRunId gets a reference to the given string and assigns it to the RunId field.
func (o *TerraformRunstateAllOf) SetRunId(v string) {
	o.RunId = &v
}

// GetStateFile returns the StateFile field value if set, zero value otherwise.
func (o *TerraformRunstateAllOf) GetStateFile() string {
	if o == nil || o.StateFile == nil {
		var ret string
		return ret
	}
	return *o.StateFile
}

// GetStateFileOk returns a tuple with the StateFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformRunstateAllOf) GetStateFileOk() (*string, bool) {
	if o == nil || o.StateFile == nil {
		return nil, false
	}
	return o.StateFile, true
}

// HasStateFile returns a boolean if a field has been set.
func (o *TerraformRunstateAllOf) HasStateFile() bool {
	if o != nil && o.StateFile != nil {
		return true
	}

	return false
}

// SetStateFile gets a reference to the given string and assigns it to the StateFile field.
func (o *TerraformRunstateAllOf) SetStateFile(v string) {
	o.StateFile = &v
}

func (o TerraformRunstateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.RunId != nil {
		toSerialize["RunId"] = o.RunId
	}
	if o.StateFile != nil {
		toSerialize["StateFile"] = o.StateFile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TerraformRunstateAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTerraformRunstateAllOf := _TerraformRunstateAllOf{}

	if err = json.Unmarshal(bytes, &varTerraformRunstateAllOf); err == nil {
		*o = TerraformRunstateAllOf(varTerraformRunstateAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "RunId")
		delete(additionalProperties, "StateFile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTerraformRunstateAllOf struct {
	value *TerraformRunstateAllOf
	isSet bool
}

func (v NullableTerraformRunstateAllOf) Get() *TerraformRunstateAllOf {
	return v.value
}

func (v *NullableTerraformRunstateAllOf) Set(val *TerraformRunstateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTerraformRunstateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTerraformRunstateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerraformRunstateAllOf(val *TerraformRunstateAllOf) *NullableTerraformRunstateAllOf {
	return &NullableTerraformRunstateAllOf{value: val, isSet: true}
}

func (v NullableTerraformRunstateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerraformRunstateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

