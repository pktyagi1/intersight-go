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

// ApplianceAutoRmaPolicy Auto rma policy to decide whether rma data should be collected.
type ApplianceAutoRmaPolicy struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of the data collection mode. If the value is 'true', then data collection is enabled.
	Enable *bool `json:"Enable,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceAutoRmaPolicy ApplianceAutoRmaPolicy

// NewApplianceAutoRmaPolicy instantiates a new ApplianceAutoRmaPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAutoRmaPolicy(classId string, objectType string) *ApplianceAutoRmaPolicy {
	this := ApplianceAutoRmaPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceAutoRmaPolicyWithDefaults instantiates a new ApplianceAutoRmaPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAutoRmaPolicyWithDefaults() *ApplianceAutoRmaPolicy {
	this := ApplianceAutoRmaPolicy{}
	var classId string = "appliance.AutoRmaPolicy"
	this.ClassId = classId
	var objectType string = "appliance.AutoRmaPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceAutoRmaPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceAutoRmaPolicy) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceAutoRmaPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceAutoRmaPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceAutoRmaPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceAutoRmaPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *ApplianceAutoRmaPolicy) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAutoRmaPolicy) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *ApplianceAutoRmaPolicy) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *ApplianceAutoRmaPolicy) SetEnable(v bool) {
	o.Enable = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ApplianceAutoRmaPolicy) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAutoRmaPolicy) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ApplianceAutoRmaPolicy) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ApplianceAutoRmaPolicy) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o ApplianceAutoRmaPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enable != nil {
		toSerialize["Enable"] = o.Enable
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceAutoRmaPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceAutoRmaPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Status of the data collection mode. If the value is 'true', then data collection is enabled.
		Enable *bool `json:"Enable,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varApplianceAutoRmaPolicyWithoutEmbeddedStruct := ApplianceAutoRmaPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceAutoRmaPolicyWithoutEmbeddedStruct)
	if err == nil {
		varApplianceAutoRmaPolicy := _ApplianceAutoRmaPolicy{}
		varApplianceAutoRmaPolicy.ClassId = varApplianceAutoRmaPolicyWithoutEmbeddedStruct.ClassId
		varApplianceAutoRmaPolicy.ObjectType = varApplianceAutoRmaPolicyWithoutEmbeddedStruct.ObjectType
		varApplianceAutoRmaPolicy.Enable = varApplianceAutoRmaPolicyWithoutEmbeddedStruct.Enable
		varApplianceAutoRmaPolicy.RegisteredDevice = varApplianceAutoRmaPolicyWithoutEmbeddedStruct.RegisteredDevice
		*o = ApplianceAutoRmaPolicy(varApplianceAutoRmaPolicy)
	} else {
		return err
	}

	varApplianceAutoRmaPolicy := _ApplianceAutoRmaPolicy{}

	err = json.Unmarshal(bytes, &varApplianceAutoRmaPolicy)
	if err == nil {
		o.MoBaseMo = varApplianceAutoRmaPolicy.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enable")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableApplianceAutoRmaPolicy struct {
	value *ApplianceAutoRmaPolicy
	isSet bool
}

func (v NullableApplianceAutoRmaPolicy) Get() *ApplianceAutoRmaPolicy {
	return v.value
}

func (v *NullableApplianceAutoRmaPolicy) Set(val *ApplianceAutoRmaPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAutoRmaPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAutoRmaPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAutoRmaPolicy(val *ApplianceAutoRmaPolicy) *NullableApplianceAutoRmaPolicy {
	return &NullableApplianceAutoRmaPolicy{value: val, isSet: true}
}

func (v NullableApplianceAutoRmaPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAutoRmaPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


