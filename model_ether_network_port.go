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

// EtherNetworkPort Model contains the details of the ethernet port connected to the FI side.
type EtherNetworkPort struct {
	PortInterfaceBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Febric extender identifier for this port.
	ModuleId *int64 `json:"ModuleId,omitempty"`
	// Peer DN for network host port of fabric extender.
	PeerDn *string `json:"PeerDn,omitempty"`
	// Switch physical port identifier.
	PortId *int64 `json:"PortId,omitempty"`
	// Switch expansion slot module identifier.
	SlotId *int64 `json:"SlotId,omitempty"`
	// Network Port Speed of IO card or fabric extender.
	Speed *string `json:"Speed,omitempty"`
	// Switch Identifier that is local to a cluster.
	SwitchId *string `json:"SwitchId,omitempty"`
	EquipmentIoCardBase *EquipmentIoCardBaseRelationship `json:"EquipmentIoCardBase,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EtherNetworkPort EtherNetworkPort

// NewEtherNetworkPort instantiates a new EtherNetworkPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEtherNetworkPort(classId string, objectType string) *EtherNetworkPort {
	this := EtherNetworkPort{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEtherNetworkPortWithDefaults instantiates a new EtherNetworkPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEtherNetworkPortWithDefaults() *EtherNetworkPort {
	this := EtherNetworkPort{}
	var classId string = "ether.NetworkPort"
	this.ClassId = classId
	var objectType string = "ether.NetworkPort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EtherNetworkPort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EtherNetworkPort) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EtherNetworkPort) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EtherNetworkPort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EtherNetworkPort) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EtherNetworkPort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *EtherNetworkPort) GetModuleId() int64 {
	if o == nil || o.ModuleId == nil {
		var ret int64
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherNetworkPort) GetModuleIdOk() (*int64, bool) {
	if o == nil || o.ModuleId == nil {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *EtherNetworkPort) HasModuleId() bool {
	if o != nil && o.ModuleId != nil {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given int64 and assigns it to the ModuleId field.
func (o *EtherNetworkPort) SetModuleId(v int64) {
	o.ModuleId = &v
}

// GetPeerDn returns the PeerDn field value if set, zero value otherwise.
func (o *EtherNetworkPort) GetPeerDn() string {
	if o == nil || o.PeerDn == nil {
		var ret string
		return ret
	}
	return *o.PeerDn
}

// GetPeerDnOk returns a tuple with the PeerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherNetworkPort) GetPeerDnOk() (*string, bool) {
	if o == nil || o.PeerDn == nil {
		return nil, false
	}
	return o.PeerDn, true
}

// HasPeerDn returns a boolean if a field has been set.
func (o *EtherNetworkPort) HasPeerDn() bool {
	if o != nil && o.PeerDn != nil {
		return true
	}

	return false
}

// SetPeerDn gets a reference to the given string and assigns it to the PeerDn field.
func (o *EtherNetworkPort) SetPeerDn(v string) {
	o.PeerDn = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *EtherNetworkPort) GetPortId() int64 {
	if o == nil || o.PortId == nil {
		var ret int64
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherNetworkPort) GetPortIdOk() (*int64, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *EtherNetworkPort) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given int64 and assigns it to the PortId field.
func (o *EtherNetworkPort) SetPortId(v int64) {
	o.PortId = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *EtherNetworkPort) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherNetworkPort) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *EtherNetworkPort) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *EtherNetworkPort) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *EtherNetworkPort) GetSpeed() string {
	if o == nil || o.Speed == nil {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherNetworkPort) GetSpeedOk() (*string, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *EtherNetworkPort) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *EtherNetworkPort) SetSpeed(v string) {
	o.Speed = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *EtherNetworkPort) GetSwitchId() string {
	if o == nil || o.SwitchId == nil {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherNetworkPort) GetSwitchIdOk() (*string, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *EtherNetworkPort) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *EtherNetworkPort) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetEquipmentIoCardBase returns the EquipmentIoCardBase field value if set, zero value otherwise.
func (o *EtherNetworkPort) GetEquipmentIoCardBase() EquipmentIoCardBaseRelationship {
	if o == nil || o.EquipmentIoCardBase == nil {
		var ret EquipmentIoCardBaseRelationship
		return ret
	}
	return *o.EquipmentIoCardBase
}

// GetEquipmentIoCardBaseOk returns a tuple with the EquipmentIoCardBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherNetworkPort) GetEquipmentIoCardBaseOk() (*EquipmentIoCardBaseRelationship, bool) {
	if o == nil || o.EquipmentIoCardBase == nil {
		return nil, false
	}
	return o.EquipmentIoCardBase, true
}

// HasEquipmentIoCardBase returns a boolean if a field has been set.
func (o *EtherNetworkPort) HasEquipmentIoCardBase() bool {
	if o != nil && o.EquipmentIoCardBase != nil {
		return true
	}

	return false
}

// SetEquipmentIoCardBase gets a reference to the given EquipmentIoCardBaseRelationship and assigns it to the EquipmentIoCardBase field.
func (o *EtherNetworkPort) SetEquipmentIoCardBase(v EquipmentIoCardBaseRelationship) {
	o.EquipmentIoCardBase = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EtherNetworkPort) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherNetworkPort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EtherNetworkPort) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EtherNetworkPort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EtherNetworkPort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPortInterfaceBase, errPortInterfaceBase := json.Marshal(o.PortInterfaceBase)
	if errPortInterfaceBase != nil {
		return []byte{}, errPortInterfaceBase
	}
	errPortInterfaceBase = json.Unmarshal([]byte(serializedPortInterfaceBase), &toSerialize)
	if errPortInterfaceBase != nil {
		return []byte{}, errPortInterfaceBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ModuleId != nil {
		toSerialize["ModuleId"] = o.ModuleId
	}
	if o.PeerDn != nil {
		toSerialize["PeerDn"] = o.PeerDn
	}
	if o.PortId != nil {
		toSerialize["PortId"] = o.PortId
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}
	if o.Speed != nil {
		toSerialize["Speed"] = o.Speed
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.EquipmentIoCardBase != nil {
		toSerialize["EquipmentIoCardBase"] = o.EquipmentIoCardBase
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EtherNetworkPort) UnmarshalJSON(bytes []byte) (err error) {
	type EtherNetworkPortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Febric extender identifier for this port.
		ModuleId *int64 `json:"ModuleId,omitempty"`
		// Peer DN for network host port of fabric extender.
		PeerDn *string `json:"PeerDn,omitempty"`
		// Switch physical port identifier.
		PortId *int64 `json:"PortId,omitempty"`
		// Switch expansion slot module identifier.
		SlotId *int64 `json:"SlotId,omitempty"`
		// Network Port Speed of IO card or fabric extender.
		Speed *string `json:"Speed,omitempty"`
		// Switch Identifier that is local to a cluster.
		SwitchId *string `json:"SwitchId,omitempty"`
		EquipmentIoCardBase *EquipmentIoCardBaseRelationship `json:"EquipmentIoCardBase,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEtherNetworkPortWithoutEmbeddedStruct := EtherNetworkPortWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEtherNetworkPortWithoutEmbeddedStruct)
	if err == nil {
		varEtherNetworkPort := _EtherNetworkPort{}
		varEtherNetworkPort.ClassId = varEtherNetworkPortWithoutEmbeddedStruct.ClassId
		varEtherNetworkPort.ObjectType = varEtherNetworkPortWithoutEmbeddedStruct.ObjectType
		varEtherNetworkPort.ModuleId = varEtherNetworkPortWithoutEmbeddedStruct.ModuleId
		varEtherNetworkPort.PeerDn = varEtherNetworkPortWithoutEmbeddedStruct.PeerDn
		varEtherNetworkPort.PortId = varEtherNetworkPortWithoutEmbeddedStruct.PortId
		varEtherNetworkPort.SlotId = varEtherNetworkPortWithoutEmbeddedStruct.SlotId
		varEtherNetworkPort.Speed = varEtherNetworkPortWithoutEmbeddedStruct.Speed
		varEtherNetworkPort.SwitchId = varEtherNetworkPortWithoutEmbeddedStruct.SwitchId
		varEtherNetworkPort.EquipmentIoCardBase = varEtherNetworkPortWithoutEmbeddedStruct.EquipmentIoCardBase
		varEtherNetworkPort.RegisteredDevice = varEtherNetworkPortWithoutEmbeddedStruct.RegisteredDevice
		*o = EtherNetworkPort(varEtherNetworkPort)
	} else {
		return err
	}

	varEtherNetworkPort := _EtherNetworkPort{}

	err = json.Unmarshal(bytes, &varEtherNetworkPort)
	if err == nil {
		o.PortInterfaceBase = varEtherNetworkPort.PortInterfaceBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ModuleId")
		delete(additionalProperties, "PeerDn")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "SlotId")
		delete(additionalProperties, "Speed")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "EquipmentIoCardBase")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectPortInterfaceBase := reflect.ValueOf(o.PortInterfaceBase)
		for i := 0; i < reflectPortInterfaceBase.Type().NumField(); i++ {
			t := reflectPortInterfaceBase.Type().Field(i)

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

type NullableEtherNetworkPort struct {
	value *EtherNetworkPort
	isSet bool
}

func (v NullableEtherNetworkPort) Get() *EtherNetworkPort {
	return v.value
}

func (v *NullableEtherNetworkPort) Set(val *EtherNetworkPort) {
	v.value = val
	v.isSet = true
}

func (v NullableEtherNetworkPort) IsSet() bool {
	return v.isSet
}

func (v *NullableEtherNetworkPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEtherNetworkPort(val *EtherNetworkPort) *NullableEtherNetworkPort {
	return &NullableEtherNetworkPort{value: val, isSet: true}
}

func (v NullableEtherNetworkPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEtherNetworkPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

