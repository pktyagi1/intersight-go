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
	"reflect"
	"strings"
)

// AdapterDceInterfaceSettings Settings applicable for the Data Center Ethernet (DCE) interfaces on the adapter card.
type AdapterDceInterfaceSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Forward Error Correction (FEC) mode setting for the DCE interfaces of the adapter. FEC mode setting is supported only for Cisco VIC 14xx adapters. FEC mode 'cl74' is unsupported for Cisco VIC 1495/1497. This setting will be ignored for unsupported adapters and for unavailable DCE interfaces. * `cl91` - Use cl91 standard as FEC mode setting. 'Clause 91' aka RS-FEC ('ReedSolomon' FEC) offers better error protection against bursty and random errors but adds latency. * `cl74` - Use cl74 standard as FEC mode setting. 'Clause 74' aka FC-FEC ('FireCode' FEC) offers simple, low-latency protection against 1 burst/sparse bit error, but it is not good for random errors. * `Off` - Disable FEC mode on the DCE Interface.
	FecMode *string `json:"FecMode,omitempty"`
	// DCE interface id on which settings needs to be configured. Supported values are (0-3).
	InterfaceId *int64 `json:"InterfaceId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdapterDceInterfaceSettings AdapterDceInterfaceSettings

// NewAdapterDceInterfaceSettings instantiates a new AdapterDceInterfaceSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterDceInterfaceSettings(classId string, objectType string) *AdapterDceInterfaceSettings {
	this := AdapterDceInterfaceSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var fecMode string = "cl91"
	this.FecMode = &fecMode
	return &this
}

// NewAdapterDceInterfaceSettingsWithDefaults instantiates a new AdapterDceInterfaceSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterDceInterfaceSettingsWithDefaults() *AdapterDceInterfaceSettings {
	this := AdapterDceInterfaceSettings{}
	var classId string = "adapter.DceInterfaceSettings"
	this.ClassId = classId
	var objectType string = "adapter.DceInterfaceSettings"
	this.ObjectType = objectType
	var fecMode string = "cl91"
	this.FecMode = &fecMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *AdapterDceInterfaceSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AdapterDceInterfaceSettings) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AdapterDceInterfaceSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AdapterDceInterfaceSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AdapterDceInterfaceSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AdapterDceInterfaceSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFecMode returns the FecMode field value if set, zero value otherwise.
func (o *AdapterDceInterfaceSettings) GetFecMode() string {
	if o == nil || o.FecMode == nil {
		var ret string
		return ret
	}
	return *o.FecMode
}

// GetFecModeOk returns a tuple with the FecMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterDceInterfaceSettings) GetFecModeOk() (*string, bool) {
	if o == nil || o.FecMode == nil {
		return nil, false
	}
	return o.FecMode, true
}

// HasFecMode returns a boolean if a field has been set.
func (o *AdapterDceInterfaceSettings) HasFecMode() bool {
	if o != nil && o.FecMode != nil {
		return true
	}

	return false
}

// SetFecMode gets a reference to the given string and assigns it to the FecMode field.
func (o *AdapterDceInterfaceSettings) SetFecMode(v string) {
	o.FecMode = &v
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *AdapterDceInterfaceSettings) GetInterfaceId() int64 {
	if o == nil || o.InterfaceId == nil {
		var ret int64
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterDceInterfaceSettings) GetInterfaceIdOk() (*int64, bool) {
	if o == nil || o.InterfaceId == nil {
		return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *AdapterDceInterfaceSettings) HasInterfaceId() bool {
	if o != nil && o.InterfaceId != nil {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given int64 and assigns it to the InterfaceId field.
func (o *AdapterDceInterfaceSettings) SetInterfaceId(v int64) {
	o.InterfaceId = &v
}

func (o AdapterDceInterfaceSettings) MarshalJSON() ([]byte, error) {
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
	if o.FecMode != nil {
		toSerialize["FecMode"] = o.FecMode
	}
	if o.InterfaceId != nil {
		toSerialize["InterfaceId"] = o.InterfaceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AdapterDceInterfaceSettings) UnmarshalJSON(bytes []byte) (err error) {
	type AdapterDceInterfaceSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Forward Error Correction (FEC) mode setting for the DCE interfaces of the adapter. FEC mode setting is supported only for Cisco VIC 14xx adapters. FEC mode 'cl74' is unsupported for Cisco VIC 1495/1497. This setting will be ignored for unsupported adapters and for unavailable DCE interfaces. * `cl91` - Use cl91 standard as FEC mode setting. 'Clause 91' aka RS-FEC ('ReedSolomon' FEC) offers better error protection against bursty and random errors but adds latency. * `cl74` - Use cl74 standard as FEC mode setting. 'Clause 74' aka FC-FEC ('FireCode' FEC) offers simple, low-latency protection against 1 burst/sparse bit error, but it is not good for random errors. * `Off` - Disable FEC mode on the DCE Interface.
		FecMode *string `json:"FecMode,omitempty"`
		// DCE interface id on which settings needs to be configured. Supported values are (0-3).
		InterfaceId *int64 `json:"InterfaceId,omitempty"`
	}

	varAdapterDceInterfaceSettingsWithoutEmbeddedStruct := AdapterDceInterfaceSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAdapterDceInterfaceSettingsWithoutEmbeddedStruct)
	if err == nil {
		varAdapterDceInterfaceSettings := _AdapterDceInterfaceSettings{}
		varAdapterDceInterfaceSettings.ClassId = varAdapterDceInterfaceSettingsWithoutEmbeddedStruct.ClassId
		varAdapterDceInterfaceSettings.ObjectType = varAdapterDceInterfaceSettingsWithoutEmbeddedStruct.ObjectType
		varAdapterDceInterfaceSettings.FecMode = varAdapterDceInterfaceSettingsWithoutEmbeddedStruct.FecMode
		varAdapterDceInterfaceSettings.InterfaceId = varAdapterDceInterfaceSettingsWithoutEmbeddedStruct.InterfaceId
		*o = AdapterDceInterfaceSettings(varAdapterDceInterfaceSettings)
	} else {
		return err
	}

	varAdapterDceInterfaceSettings := _AdapterDceInterfaceSettings{}

	err = json.Unmarshal(bytes, &varAdapterDceInterfaceSettings)
	if err == nil {
		o.MoBaseComplexType = varAdapterDceInterfaceSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FecMode")
		delete(additionalProperties, "InterfaceId")

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

type NullableAdapterDceInterfaceSettings struct {
	value *AdapterDceInterfaceSettings
	isSet bool
}

func (v NullableAdapterDceInterfaceSettings) Get() *AdapterDceInterfaceSettings {
	return v.value
}

func (v *NullableAdapterDceInterfaceSettings) Set(val *AdapterDceInterfaceSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterDceInterfaceSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterDceInterfaceSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterDceInterfaceSettings(val *AdapterDceInterfaceSettings) *NullableAdapterDceInterfaceSettings {
	return &NullableAdapterDceInterfaceSettings{value: val, isSet: true}
}

func (v NullableAdapterDceInterfaceSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterDceInterfaceSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


