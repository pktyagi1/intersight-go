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
)

// EquipmentBaseSensorAllOf Definition of the list of properties defined in 'equipment.BaseSensor', excluding properties defined in parent classes.
type EquipmentBaseSensorAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The name of the storage array controller that the sensor applies to.
	ControllerName *string `json:"ControllerName,omitempty"`
	// The name of a specific sensor.
	Name *string `json:"Name,omitempty"`
	// The state of the specified sensor.
	State *string `json:"State,omitempty"`
	// The type of the specified sensor.
	Type *string `json:"Type,omitempty"`
	// The units that correspond to the value of the sensor, if applicable.
	Units *string `json:"Units,omitempty"`
	// The value of the specified sensor.
	Value *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentBaseSensorAllOf EquipmentBaseSensorAllOf

// NewEquipmentBaseSensorAllOf instantiates a new EquipmentBaseSensorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentBaseSensorAllOf(classId string, objectType string) *EquipmentBaseSensorAllOf {
	this := EquipmentBaseSensorAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentBaseSensorAllOfWithDefaults instantiates a new EquipmentBaseSensorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentBaseSensorAllOfWithDefaults() *EquipmentBaseSensorAllOf {
	this := EquipmentBaseSensorAllOf{}
	var classId string = "storage.NetAppSensor"
	this.ClassId = classId
	var objectType string = "storage.NetAppSensor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentBaseSensorAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensorAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentBaseSensorAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentBaseSensorAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensorAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentBaseSensorAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetControllerName returns the ControllerName field value if set, zero value otherwise.
func (o *EquipmentBaseSensorAllOf) GetControllerName() string {
	if o == nil || o.ControllerName == nil {
		var ret string
		return ret
	}
	return *o.ControllerName
}

// GetControllerNameOk returns a tuple with the ControllerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensorAllOf) GetControllerNameOk() (*string, bool) {
	if o == nil || o.ControllerName == nil {
		return nil, false
	}
	return o.ControllerName, true
}

// HasControllerName returns a boolean if a field has been set.
func (o *EquipmentBaseSensorAllOf) HasControllerName() bool {
	if o != nil && o.ControllerName != nil {
		return true
	}

	return false
}

// SetControllerName gets a reference to the given string and assigns it to the ControllerName field.
func (o *EquipmentBaseSensorAllOf) SetControllerName(v string) {
	o.ControllerName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EquipmentBaseSensorAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensorAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EquipmentBaseSensorAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EquipmentBaseSensorAllOf) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *EquipmentBaseSensorAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensorAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *EquipmentBaseSensorAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *EquipmentBaseSensorAllOf) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EquipmentBaseSensorAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensorAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EquipmentBaseSensorAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EquipmentBaseSensorAllOf) SetType(v string) {
	o.Type = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *EquipmentBaseSensorAllOf) GetUnits() string {
	if o == nil || o.Units == nil {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensorAllOf) GetUnitsOk() (*string, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *EquipmentBaseSensorAllOf) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *EquipmentBaseSensorAllOf) SetUnits(v string) {
	o.Units = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EquipmentBaseSensorAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensorAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EquipmentBaseSensorAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EquipmentBaseSensorAllOf) SetValue(v string) {
	o.Value = &v
}

func (o EquipmentBaseSensorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ControllerName != nil {
		toSerialize["ControllerName"] = o.ControllerName
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Units != nil {
		toSerialize["Units"] = o.Units
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentBaseSensorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentBaseSensorAllOf := _EquipmentBaseSensorAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentBaseSensorAllOf); err == nil {
		*o = EquipmentBaseSensorAllOf(varEquipmentBaseSensorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ControllerName")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Units")
		delete(additionalProperties, "Value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentBaseSensorAllOf struct {
	value *EquipmentBaseSensorAllOf
	isSet bool
}

func (v NullableEquipmentBaseSensorAllOf) Get() *EquipmentBaseSensorAllOf {
	return v.value
}

func (v *NullableEquipmentBaseSensorAllOf) Set(val *EquipmentBaseSensorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentBaseSensorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentBaseSensorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentBaseSensorAllOf(val *EquipmentBaseSensorAllOf) *NullableEquipmentBaseSensorAllOf {
	return &NullableEquipmentBaseSensorAllOf{value: val, isSet: true}
}

func (v NullableEquipmentBaseSensorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentBaseSensorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


