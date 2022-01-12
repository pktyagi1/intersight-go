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

// VnicPlogiSettings Fibre Channel Plogi Settings.
type VnicPlogiSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of times that the system tries to log in to a port after the first failure.
	Retries *int64 `json:"Retries,omitempty"`
	// The number of milliseconds that the system waits before it tries to log in again.
	Timeout *int64 `json:"Timeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicPlogiSettings VnicPlogiSettings

// NewVnicPlogiSettings instantiates a new VnicPlogiSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicPlogiSettings(classId string, objectType string) *VnicPlogiSettings {
	this := VnicPlogiSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var retries int64 = 8
	this.Retries = &retries
	var timeout int64 = 20000
	this.Timeout = &timeout
	return &this
}

// NewVnicPlogiSettingsWithDefaults instantiates a new VnicPlogiSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicPlogiSettingsWithDefaults() *VnicPlogiSettings {
	this := VnicPlogiSettings{}
	var classId string = "vnic.PlogiSettings"
	this.ClassId = classId
	var objectType string = "vnic.PlogiSettings"
	this.ObjectType = objectType
	var retries int64 = 8
	this.Retries = &retries
	var timeout int64 = 20000
	this.Timeout = &timeout
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicPlogiSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicPlogiSettings) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicPlogiSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicPlogiSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicPlogiSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicPlogiSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *VnicPlogiSettings) GetRetries() int64 {
	if o == nil || o.Retries == nil {
		var ret int64
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlogiSettings) GetRetriesOk() (*int64, bool) {
	if o == nil || o.Retries == nil {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *VnicPlogiSettings) HasRetries() bool {
	if o != nil && o.Retries != nil {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int64 and assigns it to the Retries field.
func (o *VnicPlogiSettings) SetRetries(v int64) {
	o.Retries = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *VnicPlogiSettings) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlogiSettings) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *VnicPlogiSettings) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *VnicPlogiSettings) SetTimeout(v int64) {
	o.Timeout = &v
}

func (o VnicPlogiSettings) MarshalJSON() ([]byte, error) {
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
	if o.Retries != nil {
		toSerialize["Retries"] = o.Retries
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicPlogiSettings) UnmarshalJSON(bytes []byte) (err error) {
	type VnicPlogiSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The number of times that the system tries to log in to a port after the first failure.
		Retries *int64 `json:"Retries,omitempty"`
		// The number of milliseconds that the system waits before it tries to log in again.
		Timeout *int64 `json:"Timeout,omitempty"`
	}

	varVnicPlogiSettingsWithoutEmbeddedStruct := VnicPlogiSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicPlogiSettingsWithoutEmbeddedStruct)
	if err == nil {
		varVnicPlogiSettings := _VnicPlogiSettings{}
		varVnicPlogiSettings.ClassId = varVnicPlogiSettingsWithoutEmbeddedStruct.ClassId
		varVnicPlogiSettings.ObjectType = varVnicPlogiSettingsWithoutEmbeddedStruct.ObjectType
		varVnicPlogiSettings.Retries = varVnicPlogiSettingsWithoutEmbeddedStruct.Retries
		varVnicPlogiSettings.Timeout = varVnicPlogiSettingsWithoutEmbeddedStruct.Timeout
		*o = VnicPlogiSettings(varVnicPlogiSettings)
	} else {
		return err
	}

	varVnicPlogiSettings := _VnicPlogiSettings{}

	err = json.Unmarshal(bytes, &varVnicPlogiSettings)
	if err == nil {
		o.MoBaseComplexType = varVnicPlogiSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Retries")
		delete(additionalProperties, "Timeout")

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

type NullableVnicPlogiSettings struct {
	value *VnicPlogiSettings
	isSet bool
}

func (v NullableVnicPlogiSettings) Get() *VnicPlogiSettings {
	return v.value
}

func (v *NullableVnicPlogiSettings) Set(val *VnicPlogiSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicPlogiSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicPlogiSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicPlogiSettings(val *VnicPlogiSettings) *NullableVnicPlogiSettings {
	return &NullableVnicPlogiSettings{value: val, isSet: true}
}

func (v NullableVnicPlogiSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicPlogiSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


