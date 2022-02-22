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

// FabricUdldSettings Type to represent UDLD settings for the port and port-channel.
type FabricUdldSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin configured UDLD State for this port. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
	AdminState *string `json:"AdminState,omitempty"`
	// Admin configured UDLD Mode for this port. * `normal` - Admin configured 'normal' UDLD mode. * `aggressive` - Admin configured 'aggressive' UDLD mode.
	Mode *string `json:"Mode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricUdldSettings FabricUdldSettings

// NewFabricUdldSettings instantiates a new FabricUdldSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricUdldSettings(classId string, objectType string) *FabricUdldSettings {
	this := FabricUdldSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminState string = "Disabled"
	this.AdminState = &adminState
	var mode string = "normal"
	this.Mode = &mode
	return &this
}

// NewFabricUdldSettingsWithDefaults instantiates a new FabricUdldSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricUdldSettingsWithDefaults() *FabricUdldSettings {
	this := FabricUdldSettings{}
	var classId string = "fabric.UdldSettings"
	this.ClassId = classId
	var objectType string = "fabric.UdldSettings"
	this.ObjectType = objectType
	var adminState string = "Disabled"
	this.AdminState = &adminState
	var mode string = "normal"
	this.Mode = &mode
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricUdldSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricUdldSettings) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricUdldSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricUdldSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricUdldSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricUdldSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *FabricUdldSettings) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricUdldSettings) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *FabricUdldSettings) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *FabricUdldSettings) SetAdminState(v string) {
	o.AdminState = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *FabricUdldSettings) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricUdldSettings) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *FabricUdldSettings) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *FabricUdldSettings) SetMode(v string) {
	o.Mode = &v
}

func (o FabricUdldSettings) MarshalJSON() ([]byte, error) {
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
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricUdldSettings) UnmarshalJSON(bytes []byte) (err error) {
	type FabricUdldSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin configured UDLD State for this port. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
		AdminState *string `json:"AdminState,omitempty"`
		// Admin configured UDLD Mode for this port. * `normal` - Admin configured 'normal' UDLD mode. * `aggressive` - Admin configured 'aggressive' UDLD mode.
		Mode *string `json:"Mode,omitempty"`
	}

	varFabricUdldSettingsWithoutEmbeddedStruct := FabricUdldSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricUdldSettingsWithoutEmbeddedStruct)
	if err == nil {
		varFabricUdldSettings := _FabricUdldSettings{}
		varFabricUdldSettings.ClassId = varFabricUdldSettingsWithoutEmbeddedStruct.ClassId
		varFabricUdldSettings.ObjectType = varFabricUdldSettingsWithoutEmbeddedStruct.ObjectType
		varFabricUdldSettings.AdminState = varFabricUdldSettingsWithoutEmbeddedStruct.AdminState
		varFabricUdldSettings.Mode = varFabricUdldSettingsWithoutEmbeddedStruct.Mode
		*o = FabricUdldSettings(varFabricUdldSettings)
	} else {
		return err
	}

	varFabricUdldSettings := _FabricUdldSettings{}

	err = json.Unmarshal(bytes, &varFabricUdldSettings)
	if err == nil {
		o.MoBaseComplexType = varFabricUdldSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "Mode")

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

type NullableFabricUdldSettings struct {
	value *FabricUdldSettings
	isSet bool
}

func (v NullableFabricUdldSettings) Get() *FabricUdldSettings {
	return v.value
}

func (v *NullableFabricUdldSettings) Set(val *FabricUdldSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricUdldSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricUdldSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricUdldSettings(val *FabricUdldSettings) *NullableFabricUdldSettings {
	return &NullableFabricUdldSettings{value: val, isSet: true}
}

func (v NullableFabricUdldSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricUdldSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


