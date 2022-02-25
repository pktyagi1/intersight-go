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

// StorageVirtualDriveContainerAllOf Definition of the list of properties defined in 'storage.VirtualDriveContainer', excluding properties defined in parent classes.
type StorageVirtualDriveContainerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The identifier for this container.
	ContainerId *int64 `json:"ContainerId,omitempty"`
	EquipmentChassis *EquipmentChassisRelationship `json:"EquipmentChassis,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to storageVirtualDrive resources.
	VirtualDrive []StorageVirtualDriveRelationship `json:"VirtualDrive,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageVirtualDriveContainerAllOf StorageVirtualDriveContainerAllOf

// NewStorageVirtualDriveContainerAllOf instantiates a new StorageVirtualDriveContainerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVirtualDriveContainerAllOf(classId string, objectType string) *StorageVirtualDriveContainerAllOf {
	this := StorageVirtualDriveContainerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageVirtualDriveContainerAllOfWithDefaults instantiates a new StorageVirtualDriveContainerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVirtualDriveContainerAllOfWithDefaults() *StorageVirtualDriveContainerAllOf {
	this := StorageVirtualDriveContainerAllOf{}
	var classId string = "storage.VirtualDriveContainer"
	this.ClassId = classId
	var objectType string = "storage.VirtualDriveContainer"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageVirtualDriveContainerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveContainerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageVirtualDriveContainerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageVirtualDriveContainerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveContainerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageVirtualDriveContainerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise.
func (o *StorageVirtualDriveContainerAllOf) GetContainerId() int64 {
	if o == nil || o.ContainerId == nil {
		var ret int64
		return ret
	}
	return *o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveContainerAllOf) GetContainerIdOk() (*int64, bool) {
	if o == nil || o.ContainerId == nil {
		return nil, false
	}
	return o.ContainerId, true
}

// HasContainerId returns a boolean if a field has been set.
func (o *StorageVirtualDriveContainerAllOf) HasContainerId() bool {
	if o != nil && o.ContainerId != nil {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given int64 and assigns it to the ContainerId field.
func (o *StorageVirtualDriveContainerAllOf) SetContainerId(v int64) {
	o.ContainerId = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *StorageVirtualDriveContainerAllOf) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveContainerAllOf) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *StorageVirtualDriveContainerAllOf) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *StorageVirtualDriveContainerAllOf) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageVirtualDriveContainerAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveContainerAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageVirtualDriveContainerAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageVirtualDriveContainerAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageVirtualDriveContainerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveContainerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageVirtualDriveContainerAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageVirtualDriveContainerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetVirtualDrive returns the VirtualDrive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageVirtualDriveContainerAllOf) GetVirtualDrive() []StorageVirtualDriveRelationship {
	if o == nil  {
		var ret []StorageVirtualDriveRelationship
		return ret
	}
	return o.VirtualDrive
}

// GetVirtualDriveOk returns a tuple with the VirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageVirtualDriveContainerAllOf) GetVirtualDriveOk() (*[]StorageVirtualDriveRelationship, bool) {
	if o == nil || o.VirtualDrive == nil {
		return nil, false
	}
	return &o.VirtualDrive, true
}

// HasVirtualDrive returns a boolean if a field has been set.
func (o *StorageVirtualDriveContainerAllOf) HasVirtualDrive() bool {
	if o != nil && o.VirtualDrive != nil {
		return true
	}

	return false
}

// SetVirtualDrive gets a reference to the given []StorageVirtualDriveRelationship and assigns it to the VirtualDrive field.
func (o *StorageVirtualDriveContainerAllOf) SetVirtualDrive(v []StorageVirtualDriveRelationship) {
	o.VirtualDrive = v
}

func (o StorageVirtualDriveContainerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ContainerId != nil {
		toSerialize["ContainerId"] = o.ContainerId
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.VirtualDrive != nil {
		toSerialize["VirtualDrive"] = o.VirtualDrive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageVirtualDriveContainerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageVirtualDriveContainerAllOf := _StorageVirtualDriveContainerAllOf{}

	if err = json.Unmarshal(bytes, &varStorageVirtualDriveContainerAllOf); err == nil {
		*o = StorageVirtualDriveContainerAllOf(varStorageVirtualDriveContainerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ContainerId")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "VirtualDrive")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageVirtualDriveContainerAllOf struct {
	value *StorageVirtualDriveContainerAllOf
	isSet bool
}

func (v NullableStorageVirtualDriveContainerAllOf) Get() *StorageVirtualDriveContainerAllOf {
	return v.value
}

func (v *NullableStorageVirtualDriveContainerAllOf) Set(val *StorageVirtualDriveContainerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVirtualDriveContainerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVirtualDriveContainerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVirtualDriveContainerAllOf(val *StorageVirtualDriveContainerAllOf) *NullableStorageVirtualDriveContainerAllOf {
	return &NullableStorageVirtualDriveContainerAllOf{value: val, isSet: true}
}

func (v NullableStorageVirtualDriveContainerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVirtualDriveContainerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


