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

// FabricUdldGlobalSettingsAllOf Definition of the list of properties defined in 'fabric.UdldGlobalSettings', excluding properties defined in parent classes.
type FabricUdldGlobalSettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Configures the time between UDLD probe messages on ports that are in advertisement mode and are currently determined to be bidirectional. Valid values are from 7 to 90 seconds.
	MessageInterval *int64 `json:"MessageInterval,omitempty"`
	// UDLD recovery when enabled, attempts to bring an UDLD error-disabled port out of reset. * `none` - The standard 4th generation UCS Fabric Interconnect with 54 ports. * `reset` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.
	RecoveryAction *string `json:"RecoveryAction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricUdldGlobalSettingsAllOf FabricUdldGlobalSettingsAllOf

// NewFabricUdldGlobalSettingsAllOf instantiates a new FabricUdldGlobalSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricUdldGlobalSettingsAllOf(classId string, objectType string) *FabricUdldGlobalSettingsAllOf {
	this := FabricUdldGlobalSettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var messageInterval int64 = 15
	this.MessageInterval = &messageInterval
	var recoveryAction string = "none"
	this.RecoveryAction = &recoveryAction
	return &this
}

// NewFabricUdldGlobalSettingsAllOfWithDefaults instantiates a new FabricUdldGlobalSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricUdldGlobalSettingsAllOfWithDefaults() *FabricUdldGlobalSettingsAllOf {
	this := FabricUdldGlobalSettingsAllOf{}
	var classId string = "fabric.UdldGlobalSettings"
	this.ClassId = classId
	var objectType string = "fabric.UdldGlobalSettings"
	this.ObjectType = objectType
	var messageInterval int64 = 15
	this.MessageInterval = &messageInterval
	var recoveryAction string = "none"
	this.RecoveryAction = &recoveryAction
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricUdldGlobalSettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricUdldGlobalSettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricUdldGlobalSettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricUdldGlobalSettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricUdldGlobalSettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricUdldGlobalSettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMessageInterval returns the MessageInterval field value if set, zero value otherwise.
func (o *FabricUdldGlobalSettingsAllOf) GetMessageInterval() int64 {
	if o == nil || o.MessageInterval == nil {
		var ret int64
		return ret
	}
	return *o.MessageInterval
}

// GetMessageIntervalOk returns a tuple with the MessageInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricUdldGlobalSettingsAllOf) GetMessageIntervalOk() (*int64, bool) {
	if o == nil || o.MessageInterval == nil {
		return nil, false
	}
	return o.MessageInterval, true
}

// HasMessageInterval returns a boolean if a field has been set.
func (o *FabricUdldGlobalSettingsAllOf) HasMessageInterval() bool {
	if o != nil && o.MessageInterval != nil {
		return true
	}

	return false
}

// SetMessageInterval gets a reference to the given int64 and assigns it to the MessageInterval field.
func (o *FabricUdldGlobalSettingsAllOf) SetMessageInterval(v int64) {
	o.MessageInterval = &v
}

// GetRecoveryAction returns the RecoveryAction field value if set, zero value otherwise.
func (o *FabricUdldGlobalSettingsAllOf) GetRecoveryAction() string {
	if o == nil || o.RecoveryAction == nil {
		var ret string
		return ret
	}
	return *o.RecoveryAction
}

// GetRecoveryActionOk returns a tuple with the RecoveryAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricUdldGlobalSettingsAllOf) GetRecoveryActionOk() (*string, bool) {
	if o == nil || o.RecoveryAction == nil {
		return nil, false
	}
	return o.RecoveryAction, true
}

// HasRecoveryAction returns a boolean if a field has been set.
func (o *FabricUdldGlobalSettingsAllOf) HasRecoveryAction() bool {
	if o != nil && o.RecoveryAction != nil {
		return true
	}

	return false
}

// SetRecoveryAction gets a reference to the given string and assigns it to the RecoveryAction field.
func (o *FabricUdldGlobalSettingsAllOf) SetRecoveryAction(v string) {
	o.RecoveryAction = &v
}

func (o FabricUdldGlobalSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MessageInterval != nil {
		toSerialize["MessageInterval"] = o.MessageInterval
	}
	if o.RecoveryAction != nil {
		toSerialize["RecoveryAction"] = o.RecoveryAction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricUdldGlobalSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricUdldGlobalSettingsAllOf := _FabricUdldGlobalSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varFabricUdldGlobalSettingsAllOf); err == nil {
		*o = FabricUdldGlobalSettingsAllOf(varFabricUdldGlobalSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MessageInterval")
		delete(additionalProperties, "RecoveryAction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricUdldGlobalSettingsAllOf struct {
	value *FabricUdldGlobalSettingsAllOf
	isSet bool
}

func (v NullableFabricUdldGlobalSettingsAllOf) Get() *FabricUdldGlobalSettingsAllOf {
	return v.value
}

func (v *NullableFabricUdldGlobalSettingsAllOf) Set(val *FabricUdldGlobalSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricUdldGlobalSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricUdldGlobalSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricUdldGlobalSettingsAllOf(val *FabricUdldGlobalSettingsAllOf) *NullableFabricUdldGlobalSettingsAllOf {
	return &NullableFabricUdldGlobalSettingsAllOf{value: val, isSet: true}
}

func (v NullableFabricUdldGlobalSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricUdldGlobalSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


