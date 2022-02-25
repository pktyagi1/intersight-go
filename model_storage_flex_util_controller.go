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

// StorageFlexUtilController Storage Flex Util Adapter.
type StorageFlexUtilController struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the Flex Util Controller.
	ControllerName *string `json:"ControllerName,omitempty"`
	// The current status of the controller.
	ControllerStatus *string `json:"ControllerStatus,omitempty"`
	// Identifier for the Storage Flex Util Controller.
	FfControllerId *string `json:"FfControllerId,omitempty"`
	// The internal state of the controller.
	InternalState *string `json:"InternalState,omitempty"`
	ComputeBoard *ComputeBoardRelationship `json:"ComputeBoard,omitempty"`
	// An array of relationships to storageFlexUtilPhysicalDrive resources.
	FlexUtilPhysicalDrives []StorageFlexUtilPhysicalDriveRelationship `json:"FlexUtilPhysicalDrives,omitempty"`
	// An array of relationships to storageFlexUtilVirtualDrive resources.
	FlexUtilVirtualDrives []StorageFlexUtilVirtualDriveRelationship `json:"FlexUtilVirtualDrives,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageFlexUtilController StorageFlexUtilController

// NewStorageFlexUtilController instantiates a new StorageFlexUtilController object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageFlexUtilController(classId string, objectType string) *StorageFlexUtilController {
	this := StorageFlexUtilController{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageFlexUtilControllerWithDefaults instantiates a new StorageFlexUtilController object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageFlexUtilControllerWithDefaults() *StorageFlexUtilController {
	this := StorageFlexUtilController{}
	var classId string = "storage.FlexUtilController"
	this.ClassId = classId
	var objectType string = "storage.FlexUtilController"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageFlexUtilController) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilController) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageFlexUtilController) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageFlexUtilController) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilController) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageFlexUtilController) SetObjectType(v string) {
	o.ObjectType = v
}

// GetControllerName returns the ControllerName field value if set, zero value otherwise.
func (o *StorageFlexUtilController) GetControllerName() string {
	if o == nil || o.ControllerName == nil {
		var ret string
		return ret
	}
	return *o.ControllerName
}

// GetControllerNameOk returns a tuple with the ControllerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilController) GetControllerNameOk() (*string, bool) {
	if o == nil || o.ControllerName == nil {
		return nil, false
	}
	return o.ControllerName, true
}

// HasControllerName returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasControllerName() bool {
	if o != nil && o.ControllerName != nil {
		return true
	}

	return false
}

// SetControllerName gets a reference to the given string and assigns it to the ControllerName field.
func (o *StorageFlexUtilController) SetControllerName(v string) {
	o.ControllerName = &v
}

// GetControllerStatus returns the ControllerStatus field value if set, zero value otherwise.
func (o *StorageFlexUtilController) GetControllerStatus() string {
	if o == nil || o.ControllerStatus == nil {
		var ret string
		return ret
	}
	return *o.ControllerStatus
}

// GetControllerStatusOk returns a tuple with the ControllerStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilController) GetControllerStatusOk() (*string, bool) {
	if o == nil || o.ControllerStatus == nil {
		return nil, false
	}
	return o.ControllerStatus, true
}

// HasControllerStatus returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasControllerStatus() bool {
	if o != nil && o.ControllerStatus != nil {
		return true
	}

	return false
}

// SetControllerStatus gets a reference to the given string and assigns it to the ControllerStatus field.
func (o *StorageFlexUtilController) SetControllerStatus(v string) {
	o.ControllerStatus = &v
}

// GetFfControllerId returns the FfControllerId field value if set, zero value otherwise.
func (o *StorageFlexUtilController) GetFfControllerId() string {
	if o == nil || o.FfControllerId == nil {
		var ret string
		return ret
	}
	return *o.FfControllerId
}

// GetFfControllerIdOk returns a tuple with the FfControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilController) GetFfControllerIdOk() (*string, bool) {
	if o == nil || o.FfControllerId == nil {
		return nil, false
	}
	return o.FfControllerId, true
}

// HasFfControllerId returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasFfControllerId() bool {
	if o != nil && o.FfControllerId != nil {
		return true
	}

	return false
}

// SetFfControllerId gets a reference to the given string and assigns it to the FfControllerId field.
func (o *StorageFlexUtilController) SetFfControllerId(v string) {
	o.FfControllerId = &v
}

// GetInternalState returns the InternalState field value if set, zero value otherwise.
func (o *StorageFlexUtilController) GetInternalState() string {
	if o == nil || o.InternalState == nil {
		var ret string
		return ret
	}
	return *o.InternalState
}

// GetInternalStateOk returns a tuple with the InternalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilController) GetInternalStateOk() (*string, bool) {
	if o == nil || o.InternalState == nil {
		return nil, false
	}
	return o.InternalState, true
}

// HasInternalState returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasInternalState() bool {
	if o != nil && o.InternalState != nil {
		return true
	}

	return false
}

// SetInternalState gets a reference to the given string and assigns it to the InternalState field.
func (o *StorageFlexUtilController) SetInternalState(v string) {
	o.InternalState = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise.
func (o *StorageFlexUtilController) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || o.ComputeBoard == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilController) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.ComputeBoard == nil {
		return nil, false
	}
	return o.ComputeBoard, true
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard != nil {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given ComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *StorageFlexUtilController) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard = &v
}

// GetFlexUtilPhysicalDrives returns the FlexUtilPhysicalDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexUtilController) GetFlexUtilPhysicalDrives() []StorageFlexUtilPhysicalDriveRelationship {
	if o == nil  {
		var ret []StorageFlexUtilPhysicalDriveRelationship
		return ret
	}
	return o.FlexUtilPhysicalDrives
}

// GetFlexUtilPhysicalDrivesOk returns a tuple with the FlexUtilPhysicalDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexUtilController) GetFlexUtilPhysicalDrivesOk() (*[]StorageFlexUtilPhysicalDriveRelationship, bool) {
	if o == nil || o.FlexUtilPhysicalDrives == nil {
		return nil, false
	}
	return &o.FlexUtilPhysicalDrives, true
}

// HasFlexUtilPhysicalDrives returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasFlexUtilPhysicalDrives() bool {
	if o != nil && o.FlexUtilPhysicalDrives != nil {
		return true
	}

	return false
}

// SetFlexUtilPhysicalDrives gets a reference to the given []StorageFlexUtilPhysicalDriveRelationship and assigns it to the FlexUtilPhysicalDrives field.
func (o *StorageFlexUtilController) SetFlexUtilPhysicalDrives(v []StorageFlexUtilPhysicalDriveRelationship) {
	o.FlexUtilPhysicalDrives = v
}

// GetFlexUtilVirtualDrives returns the FlexUtilVirtualDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexUtilController) GetFlexUtilVirtualDrives() []StorageFlexUtilVirtualDriveRelationship {
	if o == nil  {
		var ret []StorageFlexUtilVirtualDriveRelationship
		return ret
	}
	return o.FlexUtilVirtualDrives
}

// GetFlexUtilVirtualDrivesOk returns a tuple with the FlexUtilVirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexUtilController) GetFlexUtilVirtualDrivesOk() (*[]StorageFlexUtilVirtualDriveRelationship, bool) {
	if o == nil || o.FlexUtilVirtualDrives == nil {
		return nil, false
	}
	return &o.FlexUtilVirtualDrives, true
}

// HasFlexUtilVirtualDrives returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasFlexUtilVirtualDrives() bool {
	if o != nil && o.FlexUtilVirtualDrives != nil {
		return true
	}

	return false
}

// SetFlexUtilVirtualDrives gets a reference to the given []StorageFlexUtilVirtualDriveRelationship and assigns it to the FlexUtilVirtualDrives field.
func (o *StorageFlexUtilController) SetFlexUtilVirtualDrives(v []StorageFlexUtilVirtualDriveRelationship) {
	o.FlexUtilVirtualDrives = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageFlexUtilController) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilController) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageFlexUtilController) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageFlexUtilController) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilController) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageFlexUtilController) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageFlexUtilController) MarshalJSON() ([]byte, error) {
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
	if o.ControllerName != nil {
		toSerialize["ControllerName"] = o.ControllerName
	}
	if o.ControllerStatus != nil {
		toSerialize["ControllerStatus"] = o.ControllerStatus
	}
	if o.FfControllerId != nil {
		toSerialize["FfControllerId"] = o.FfControllerId
	}
	if o.InternalState != nil {
		toSerialize["InternalState"] = o.InternalState
	}
	if o.ComputeBoard != nil {
		toSerialize["ComputeBoard"] = o.ComputeBoard
	}
	if o.FlexUtilPhysicalDrives != nil {
		toSerialize["FlexUtilPhysicalDrives"] = o.FlexUtilPhysicalDrives
	}
	if o.FlexUtilVirtualDrives != nil {
		toSerialize["FlexUtilVirtualDrives"] = o.FlexUtilVirtualDrives
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageFlexUtilController) UnmarshalJSON(bytes []byte) (err error) {
	type StorageFlexUtilControllerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the Flex Util Controller.
		ControllerName *string `json:"ControllerName,omitempty"`
		// The current status of the controller.
		ControllerStatus *string `json:"ControllerStatus,omitempty"`
		// Identifier for the Storage Flex Util Controller.
		FfControllerId *string `json:"FfControllerId,omitempty"`
		// The internal state of the controller.
		InternalState *string `json:"InternalState,omitempty"`
		ComputeBoard *ComputeBoardRelationship `json:"ComputeBoard,omitempty"`
		// An array of relationships to storageFlexUtilPhysicalDrive resources.
		FlexUtilPhysicalDrives []StorageFlexUtilPhysicalDriveRelationship `json:"FlexUtilPhysicalDrives,omitempty"`
		// An array of relationships to storageFlexUtilVirtualDrive resources.
		FlexUtilVirtualDrives []StorageFlexUtilVirtualDriveRelationship `json:"FlexUtilVirtualDrives,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStorageFlexUtilControllerWithoutEmbeddedStruct := StorageFlexUtilControllerWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageFlexUtilControllerWithoutEmbeddedStruct)
	if err == nil {
		varStorageFlexUtilController := _StorageFlexUtilController{}
		varStorageFlexUtilController.ClassId = varStorageFlexUtilControllerWithoutEmbeddedStruct.ClassId
		varStorageFlexUtilController.ObjectType = varStorageFlexUtilControllerWithoutEmbeddedStruct.ObjectType
		varStorageFlexUtilController.ControllerName = varStorageFlexUtilControllerWithoutEmbeddedStruct.ControllerName
		varStorageFlexUtilController.ControllerStatus = varStorageFlexUtilControllerWithoutEmbeddedStruct.ControllerStatus
		varStorageFlexUtilController.FfControllerId = varStorageFlexUtilControllerWithoutEmbeddedStruct.FfControllerId
		varStorageFlexUtilController.InternalState = varStorageFlexUtilControllerWithoutEmbeddedStruct.InternalState
		varStorageFlexUtilController.ComputeBoard = varStorageFlexUtilControllerWithoutEmbeddedStruct.ComputeBoard
		varStorageFlexUtilController.FlexUtilPhysicalDrives = varStorageFlexUtilControllerWithoutEmbeddedStruct.FlexUtilPhysicalDrives
		varStorageFlexUtilController.FlexUtilVirtualDrives = varStorageFlexUtilControllerWithoutEmbeddedStruct.FlexUtilVirtualDrives
		varStorageFlexUtilController.InventoryDeviceInfo = varStorageFlexUtilControllerWithoutEmbeddedStruct.InventoryDeviceInfo
		varStorageFlexUtilController.RegisteredDevice = varStorageFlexUtilControllerWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageFlexUtilController(varStorageFlexUtilController)
	} else {
		return err
	}

	varStorageFlexUtilController := _StorageFlexUtilController{}

	err = json.Unmarshal(bytes, &varStorageFlexUtilController)
	if err == nil {
		o.InventoryBase = varStorageFlexUtilController.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ControllerName")
		delete(additionalProperties, "ControllerStatus")
		delete(additionalProperties, "FfControllerId")
		delete(additionalProperties, "InternalState")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "FlexUtilPhysicalDrives")
		delete(additionalProperties, "FlexUtilVirtualDrives")
		delete(additionalProperties, "InventoryDeviceInfo")
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

type NullableStorageFlexUtilController struct {
	value *StorageFlexUtilController
	isSet bool
}

func (v NullableStorageFlexUtilController) Get() *StorageFlexUtilController {
	return v.value
}

func (v *NullableStorageFlexUtilController) Set(val *StorageFlexUtilController) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageFlexUtilController) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageFlexUtilController) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageFlexUtilController(val *StorageFlexUtilController) *NullableStorageFlexUtilController {
	return &NullableStorageFlexUtilController{value: val, isSet: true}
}

func (v NullableStorageFlexUtilController) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageFlexUtilController) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


