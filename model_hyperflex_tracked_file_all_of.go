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

// HyperflexTrackedFileAllOf Definition of the list of properties defined in 'hyperflex.TrackedFile', excluding properties defined in parent classes.
type HyperflexTrackedFileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	FilePath NullableHyperflexFilePath `json:"FilePath,omitempty"`
	// File type for the tracked file.
	FileType *string `json:"FileType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexTrackedFileAllOf HyperflexTrackedFileAllOf

// NewHyperflexTrackedFileAllOf instantiates a new HyperflexTrackedFileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexTrackedFileAllOf(classId string, objectType string) *HyperflexTrackedFileAllOf {
	this := HyperflexTrackedFileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexTrackedFileAllOfWithDefaults instantiates a new HyperflexTrackedFileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexTrackedFileAllOfWithDefaults() *HyperflexTrackedFileAllOf {
	this := HyperflexTrackedFileAllOf{}
	var classId string = "hyperflex.TrackedFile"
	this.ClassId = classId
	var objectType string = "hyperflex.TrackedFile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexTrackedFileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexTrackedFileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexTrackedFileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexTrackedFileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexTrackedFileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexTrackedFileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexTrackedFileAllOf) GetFilePath() HyperflexFilePath {
	if o == nil || o.FilePath.Get() == nil {
		var ret HyperflexFilePath
		return ret
	}
	return *o.FilePath.Get()
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexTrackedFileAllOf) GetFilePathOk() (*HyperflexFilePath, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FilePath.Get(), o.FilePath.IsSet()
}

// HasFilePath returns a boolean if a field has been set.
func (o *HyperflexTrackedFileAllOf) HasFilePath() bool {
	if o != nil && o.FilePath.IsSet() {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given NullableHyperflexFilePath and assigns it to the FilePath field.
func (o *HyperflexTrackedFileAllOf) SetFilePath(v HyperflexFilePath) {
	o.FilePath.Set(&v)
}
// SetFilePathNil sets the value for FilePath to be an explicit nil
func (o *HyperflexTrackedFileAllOf) SetFilePathNil() {
	o.FilePath.Set(nil)
}

// UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
func (o *HyperflexTrackedFileAllOf) UnsetFilePath() {
	o.FilePath.Unset()
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *HyperflexTrackedFileAllOf) GetFileType() string {
	if o == nil || o.FileType == nil {
		var ret string
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexTrackedFileAllOf) GetFileTypeOk() (*string, bool) {
	if o == nil || o.FileType == nil {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *HyperflexTrackedFileAllOf) HasFileType() bool {
	if o != nil && o.FileType != nil {
		return true
	}

	return false
}

// SetFileType gets a reference to the given string and assigns it to the FileType field.
func (o *HyperflexTrackedFileAllOf) SetFileType(v string) {
	o.FileType = &v
}

func (o HyperflexTrackedFileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FilePath.IsSet() {
		toSerialize["FilePath"] = o.FilePath.Get()
	}
	if o.FileType != nil {
		toSerialize["FileType"] = o.FileType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexTrackedFileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexTrackedFileAllOf := _HyperflexTrackedFileAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexTrackedFileAllOf); err == nil {
		*o = HyperflexTrackedFileAllOf(varHyperflexTrackedFileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FilePath")
		delete(additionalProperties, "FileType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexTrackedFileAllOf struct {
	value *HyperflexTrackedFileAllOf
	isSet bool
}

func (v NullableHyperflexTrackedFileAllOf) Get() *HyperflexTrackedFileAllOf {
	return v.value
}

func (v *NullableHyperflexTrackedFileAllOf) Set(val *HyperflexTrackedFileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexTrackedFileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexTrackedFileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexTrackedFileAllOf(val *HyperflexTrackedFileAllOf) *NullableHyperflexTrackedFileAllOf {
	return &NullableHyperflexTrackedFileAllOf{value: val, isSet: true}
}

func (v NullableHyperflexTrackedFileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexTrackedFileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


