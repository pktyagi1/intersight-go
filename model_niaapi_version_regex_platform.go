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
	"reflect"
	"strings"
)

// NiaapiVersionRegexPlatform The regular expression to parse the software version strings for specific platform.
type NiaapiVersionRegexPlatform struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// All long live version Regex pattern.
	Anyllregex *string `json:"Anyllregex,omitempty"`
	Currentlltrain NullableNiaapiSoftwareRegex `json:"Currentlltrain,omitempty"`
	Latestsltrain NullableNiaapiSoftwareRegex `json:"Latestsltrain,omitempty"`
	Sltrain []NiaapiSoftwareRegex `json:"Sltrain,omitempty"`
	Upcominglltrain NullableNiaapiSoftwareRegex `json:"Upcominglltrain,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiVersionRegexPlatform NiaapiVersionRegexPlatform

// NewNiaapiVersionRegexPlatform instantiates a new NiaapiVersionRegexPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiVersionRegexPlatform(classId string, objectType string) *NiaapiVersionRegexPlatform {
	this := NiaapiVersionRegexPlatform{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiVersionRegexPlatformWithDefaults instantiates a new NiaapiVersionRegexPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiVersionRegexPlatformWithDefaults() *NiaapiVersionRegexPlatform {
	this := NiaapiVersionRegexPlatform{}
	var classId string = "niaapi.VersionRegexPlatform"
	this.ClassId = classId
	var objectType string = "niaapi.VersionRegexPlatform"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiVersionRegexPlatform) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatform) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiVersionRegexPlatform) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiVersionRegexPlatform) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatform) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiVersionRegexPlatform) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAnyllregex returns the Anyllregex field value if set, zero value otherwise.
func (o *NiaapiVersionRegexPlatform) GetAnyllregex() string {
	if o == nil || o.Anyllregex == nil {
		var ret string
		return ret
	}
	return *o.Anyllregex
}

// GetAnyllregexOk returns a tuple with the Anyllregex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatform) GetAnyllregexOk() (*string, bool) {
	if o == nil || o.Anyllregex == nil {
		return nil, false
	}
	return o.Anyllregex, true
}

// HasAnyllregex returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatform) HasAnyllregex() bool {
	if o != nil && o.Anyllregex != nil {
		return true
	}

	return false
}

// SetAnyllregex gets a reference to the given string and assigns it to the Anyllregex field.
func (o *NiaapiVersionRegexPlatform) SetAnyllregex(v string) {
	o.Anyllregex = &v
}

// GetCurrentlltrain returns the Currentlltrain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiVersionRegexPlatform) GetCurrentlltrain() NiaapiSoftwareRegex {
	if o == nil || o.Currentlltrain.Get() == nil {
		var ret NiaapiSoftwareRegex
		return ret
	}
	return *o.Currentlltrain.Get()
}

// GetCurrentlltrainOk returns a tuple with the Currentlltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiVersionRegexPlatform) GetCurrentlltrainOk() (*NiaapiSoftwareRegex, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currentlltrain.Get(), o.Currentlltrain.IsSet()
}

// HasCurrentlltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatform) HasCurrentlltrain() bool {
	if o != nil && o.Currentlltrain.IsSet() {
		return true
	}

	return false
}

// SetCurrentlltrain gets a reference to the given NullableNiaapiSoftwareRegex and assigns it to the Currentlltrain field.
func (o *NiaapiVersionRegexPlatform) SetCurrentlltrain(v NiaapiSoftwareRegex) {
	o.Currentlltrain.Set(&v)
}
// SetCurrentlltrainNil sets the value for Currentlltrain to be an explicit nil
func (o *NiaapiVersionRegexPlatform) SetCurrentlltrainNil() {
	o.Currentlltrain.Set(nil)
}

// UnsetCurrentlltrain ensures that no value is present for Currentlltrain, not even an explicit nil
func (o *NiaapiVersionRegexPlatform) UnsetCurrentlltrain() {
	o.Currentlltrain.Unset()
}

// GetLatestsltrain returns the Latestsltrain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiVersionRegexPlatform) GetLatestsltrain() NiaapiSoftwareRegex {
	if o == nil || o.Latestsltrain.Get() == nil {
		var ret NiaapiSoftwareRegex
		return ret
	}
	return *o.Latestsltrain.Get()
}

// GetLatestsltrainOk returns a tuple with the Latestsltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiVersionRegexPlatform) GetLatestsltrainOk() (*NiaapiSoftwareRegex, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Latestsltrain.Get(), o.Latestsltrain.IsSet()
}

// HasLatestsltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatform) HasLatestsltrain() bool {
	if o != nil && o.Latestsltrain.IsSet() {
		return true
	}

	return false
}

// SetLatestsltrain gets a reference to the given NullableNiaapiSoftwareRegex and assigns it to the Latestsltrain field.
func (o *NiaapiVersionRegexPlatform) SetLatestsltrain(v NiaapiSoftwareRegex) {
	o.Latestsltrain.Set(&v)
}
// SetLatestsltrainNil sets the value for Latestsltrain to be an explicit nil
func (o *NiaapiVersionRegexPlatform) SetLatestsltrainNil() {
	o.Latestsltrain.Set(nil)
}

// UnsetLatestsltrain ensures that no value is present for Latestsltrain, not even an explicit nil
func (o *NiaapiVersionRegexPlatform) UnsetLatestsltrain() {
	o.Latestsltrain.Unset()
}

// GetSltrain returns the Sltrain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiVersionRegexPlatform) GetSltrain() []NiaapiSoftwareRegex {
	if o == nil  {
		var ret []NiaapiSoftwareRegex
		return ret
	}
	return o.Sltrain
}

// GetSltrainOk returns a tuple with the Sltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiVersionRegexPlatform) GetSltrainOk() (*[]NiaapiSoftwareRegex, bool) {
	if o == nil || o.Sltrain == nil {
		return nil, false
	}
	return &o.Sltrain, true
}

// HasSltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatform) HasSltrain() bool {
	if o != nil && o.Sltrain != nil {
		return true
	}

	return false
}

// SetSltrain gets a reference to the given []NiaapiSoftwareRegex and assigns it to the Sltrain field.
func (o *NiaapiVersionRegexPlatform) SetSltrain(v []NiaapiSoftwareRegex) {
	o.Sltrain = v
}

// GetUpcominglltrain returns the Upcominglltrain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiVersionRegexPlatform) GetUpcominglltrain() NiaapiSoftwareRegex {
	if o == nil || o.Upcominglltrain.Get() == nil {
		var ret NiaapiSoftwareRegex
		return ret
	}
	return *o.Upcominglltrain.Get()
}

// GetUpcominglltrainOk returns a tuple with the Upcominglltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiVersionRegexPlatform) GetUpcominglltrainOk() (*NiaapiSoftwareRegex, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Upcominglltrain.Get(), o.Upcominglltrain.IsSet()
}

// HasUpcominglltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatform) HasUpcominglltrain() bool {
	if o != nil && o.Upcominglltrain.IsSet() {
		return true
	}

	return false
}

// SetUpcominglltrain gets a reference to the given NullableNiaapiSoftwareRegex and assigns it to the Upcominglltrain field.
func (o *NiaapiVersionRegexPlatform) SetUpcominglltrain(v NiaapiSoftwareRegex) {
	o.Upcominglltrain.Set(&v)
}
// SetUpcominglltrainNil sets the value for Upcominglltrain to be an explicit nil
func (o *NiaapiVersionRegexPlatform) SetUpcominglltrainNil() {
	o.Upcominglltrain.Set(nil)
}

// UnsetUpcominglltrain ensures that no value is present for Upcominglltrain, not even an explicit nil
func (o *NiaapiVersionRegexPlatform) UnsetUpcominglltrain() {
	o.Upcominglltrain.Unset()
}

func (o NiaapiVersionRegexPlatform) MarshalJSON() ([]byte, error) {
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
	if o.Anyllregex != nil {
		toSerialize["Anyllregex"] = o.Anyllregex
	}
	if o.Currentlltrain.IsSet() {
		toSerialize["Currentlltrain"] = o.Currentlltrain.Get()
	}
	if o.Latestsltrain.IsSet() {
		toSerialize["Latestsltrain"] = o.Latestsltrain.Get()
	}
	if o.Sltrain != nil {
		toSerialize["Sltrain"] = o.Sltrain
	}
	if o.Upcominglltrain.IsSet() {
		toSerialize["Upcominglltrain"] = o.Upcominglltrain.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiVersionRegexPlatform) UnmarshalJSON(bytes []byte) (err error) {
	type NiaapiVersionRegexPlatformWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// All long live version Regex pattern.
		Anyllregex *string `json:"Anyllregex,omitempty"`
		Currentlltrain NullableNiaapiSoftwareRegex `json:"Currentlltrain,omitempty"`
		Latestsltrain NullableNiaapiSoftwareRegex `json:"Latestsltrain,omitempty"`
		Sltrain []NiaapiSoftwareRegex `json:"Sltrain,omitempty"`
		Upcominglltrain NullableNiaapiSoftwareRegex `json:"Upcominglltrain,omitempty"`
	}

	varNiaapiVersionRegexPlatformWithoutEmbeddedStruct := NiaapiVersionRegexPlatformWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiaapiVersionRegexPlatformWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiVersionRegexPlatform := _NiaapiVersionRegexPlatform{}
		varNiaapiVersionRegexPlatform.ClassId = varNiaapiVersionRegexPlatformWithoutEmbeddedStruct.ClassId
		varNiaapiVersionRegexPlatform.ObjectType = varNiaapiVersionRegexPlatformWithoutEmbeddedStruct.ObjectType
		varNiaapiVersionRegexPlatform.Anyllregex = varNiaapiVersionRegexPlatformWithoutEmbeddedStruct.Anyllregex
		varNiaapiVersionRegexPlatform.Currentlltrain = varNiaapiVersionRegexPlatformWithoutEmbeddedStruct.Currentlltrain
		varNiaapiVersionRegexPlatform.Latestsltrain = varNiaapiVersionRegexPlatformWithoutEmbeddedStruct.Latestsltrain
		varNiaapiVersionRegexPlatform.Sltrain = varNiaapiVersionRegexPlatformWithoutEmbeddedStruct.Sltrain
		varNiaapiVersionRegexPlatform.Upcominglltrain = varNiaapiVersionRegexPlatformWithoutEmbeddedStruct.Upcominglltrain
		*o = NiaapiVersionRegexPlatform(varNiaapiVersionRegexPlatform)
	} else {
		return err
	}

	varNiaapiVersionRegexPlatform := _NiaapiVersionRegexPlatform{}

	err = json.Unmarshal(bytes, &varNiaapiVersionRegexPlatform)
	if err == nil {
		o.MoBaseComplexType = varNiaapiVersionRegexPlatform.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Anyllregex")
		delete(additionalProperties, "Currentlltrain")
		delete(additionalProperties, "Latestsltrain")
		delete(additionalProperties, "Sltrain")
		delete(additionalProperties, "Upcominglltrain")

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

type NullableNiaapiVersionRegexPlatform struct {
	value *NiaapiVersionRegexPlatform
	isSet bool
}

func (v NullableNiaapiVersionRegexPlatform) Get() *NiaapiVersionRegexPlatform {
	return v.value
}

func (v *NullableNiaapiVersionRegexPlatform) Set(val *NiaapiVersionRegexPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiVersionRegexPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiVersionRegexPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiVersionRegexPlatform(val *NiaapiVersionRegexPlatform) *NullableNiaapiVersionRegexPlatform {
	return &NullableNiaapiVersionRegexPlatform{value: val, isSet: true}
}

func (v NullableNiaapiVersionRegexPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiVersionRegexPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


