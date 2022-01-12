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

// StorageItem FI Local Storage information.
type StorageItem struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The alarmType of the Local storage in FI.
	AlarmType *string `json:"AlarmType,omitempty"`
	// The name of the Local storage in FI.
	Name *string `json:"Name,omitempty"`
	// The operState of the Local storage in FI.
	OperState *string `json:"OperState,omitempty"`
	// The size (MB) of the Local storage in FI.
	Size *string `json:"Size,omitempty"`
	// The used percent of the Local storage in FI.
	Used *string `json:"Used,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	NetworkElement *NetworkElementRelationship `json:"NetworkElement,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageItem StorageItem

// NewStorageItem instantiates a new StorageItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageItem(classId string, objectType string) *StorageItem {
	this := StorageItem{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageItemWithDefaults instantiates a new StorageItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageItemWithDefaults() *StorageItem {
	this := StorageItem{}
	var classId string = "storage.Item"
	this.ClassId = classId
	var objectType string = "storage.Item"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageItem) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageItem) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageItem) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageItem) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageItem) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageItem) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlarmType returns the AlarmType field value if set, zero value otherwise.
func (o *StorageItem) GetAlarmType() string {
	if o == nil || o.AlarmType == nil {
		var ret string
		return ret
	}
	return *o.AlarmType
}

// GetAlarmTypeOk returns a tuple with the AlarmType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageItem) GetAlarmTypeOk() (*string, bool) {
	if o == nil || o.AlarmType == nil {
		return nil, false
	}
	return o.AlarmType, true
}

// HasAlarmType returns a boolean if a field has been set.
func (o *StorageItem) HasAlarmType() bool {
	if o != nil && o.AlarmType != nil {
		return true
	}

	return false
}

// SetAlarmType gets a reference to the given string and assigns it to the AlarmType field.
func (o *StorageItem) SetAlarmType(v string) {
	o.AlarmType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageItem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageItem) SetName(v string) {
	o.Name = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *StorageItem) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageItem) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *StorageItem) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *StorageItem) SetOperState(v string) {
	o.OperState = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageItem) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageItem) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageItem) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *StorageItem) SetSize(v string) {
	o.Size = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *StorageItem) GetUsed() string {
	if o == nil || o.Used == nil {
		var ret string
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageItem) GetUsedOk() (*string, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *StorageItem) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given string and assigns it to the Used field.
func (o *StorageItem) SetUsed(v string) {
	o.Used = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageItem) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageItem) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageItem) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageItem) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *StorageItem) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageItem) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *StorageItem) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *StorageItem) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageItem) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageItem) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageItem) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageItem) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AlarmType != nil {
		toSerialize["AlarmType"] = o.AlarmType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.Used != nil {
		toSerialize["Used"] = o.Used
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageItem) UnmarshalJSON(bytes []byte) (err error) {
	type StorageItemWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The alarmType of the Local storage in FI.
		AlarmType *string `json:"AlarmType,omitempty"`
		// The name of the Local storage in FI.
		Name *string `json:"Name,omitempty"`
		// The operState of the Local storage in FI.
		OperState *string `json:"OperState,omitempty"`
		// The size (MB) of the Local storage in FI.
		Size *string `json:"Size,omitempty"`
		// The used percent of the Local storage in FI.
		Used *string `json:"Used,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
		NetworkElement *NetworkElementRelationship `json:"NetworkElement,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStorageItemWithoutEmbeddedStruct := StorageItemWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageItemWithoutEmbeddedStruct)
	if err == nil {
		varStorageItem := _StorageItem{}
		varStorageItem.ClassId = varStorageItemWithoutEmbeddedStruct.ClassId
		varStorageItem.ObjectType = varStorageItemWithoutEmbeddedStruct.ObjectType
		varStorageItem.AlarmType = varStorageItemWithoutEmbeddedStruct.AlarmType
		varStorageItem.Name = varStorageItemWithoutEmbeddedStruct.Name
		varStorageItem.OperState = varStorageItemWithoutEmbeddedStruct.OperState
		varStorageItem.Size = varStorageItemWithoutEmbeddedStruct.Size
		varStorageItem.Used = varStorageItemWithoutEmbeddedStruct.Used
		varStorageItem.InventoryDeviceInfo = varStorageItemWithoutEmbeddedStruct.InventoryDeviceInfo
		varStorageItem.NetworkElement = varStorageItemWithoutEmbeddedStruct.NetworkElement
		varStorageItem.RegisteredDevice = varStorageItemWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageItem(varStorageItem)
	} else {
		return err
	}

	varStorageItem := _StorageItem{}

	err = json.Unmarshal(bytes, &varStorageItem)
	if err == nil {
		o.InventoryBase = varStorageItem.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlarmType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "Used")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableStorageItem struct {
	value *StorageItem
	isSet bool
}

func (v NullableStorageItem) Get() *StorageItem {
	return v.value
}

func (v *NullableStorageItem) Set(val *StorageItem) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageItem) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageItem(val *StorageItem) *NullableStorageItem {
	return &NullableStorageItem{value: val, isSet: true}
}

func (v NullableStorageItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

