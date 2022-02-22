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

// EquipmentTpm TPM security chip on server board.
type EquipmentTpm struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Identifies the activation status of the TPM.
	ActivationStatus *string `json:"ActivationStatus,omitempty"`
	// Identifies the admin configured state of the TPM.
	AdminState *string `json:"AdminState,omitempty"`
	// Firmware Version of the Trusted Platform Module.
	FirmwareVersion *string `json:"FirmwareVersion,omitempty"`
	// Identifies the ownership information of the TPM.
	Ownership *string `json:"Ownership,omitempty"`
	// Enter  the ID of the trusted platform module.
	TpmId *int64 `json:"TpmId,omitempty"`
	// Identifies the version of the Trusted Platform Module.
	Version *string `json:"Version,omitempty"`
	ComputeBoard *ComputeBoardRelationship `json:"ComputeBoard,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentTpm EquipmentTpm

// NewEquipmentTpm instantiates a new EquipmentTpm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentTpm(classId string, objectType string) *EquipmentTpm {
	this := EquipmentTpm{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentTpmWithDefaults instantiates a new EquipmentTpm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentTpmWithDefaults() *EquipmentTpm {
	this := EquipmentTpm{}
	var classId string = "equipment.Tpm"
	this.ClassId = classId
	var objectType string = "equipment.Tpm"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentTpm) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentTpm) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentTpm) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentTpm) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentTpm) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentTpm) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActivationStatus returns the ActivationStatus field value if set, zero value otherwise.
func (o *EquipmentTpm) GetActivationStatus() string {
	if o == nil || o.ActivationStatus == nil {
		var ret string
		return ret
	}
	return *o.ActivationStatus
}

// GetActivationStatusOk returns a tuple with the ActivationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpm) GetActivationStatusOk() (*string, bool) {
	if o == nil || o.ActivationStatus == nil {
		return nil, false
	}
	return o.ActivationStatus, true
}

// HasActivationStatus returns a boolean if a field has been set.
func (o *EquipmentTpm) HasActivationStatus() bool {
	if o != nil && o.ActivationStatus != nil {
		return true
	}

	return false
}

// SetActivationStatus gets a reference to the given string and assigns it to the ActivationStatus field.
func (o *EquipmentTpm) SetActivationStatus(v string) {
	o.ActivationStatus = &v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *EquipmentTpm) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpm) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *EquipmentTpm) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *EquipmentTpm) SetAdminState(v string) {
	o.AdminState = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *EquipmentTpm) GetFirmwareVersion() string {
	if o == nil || o.FirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpm) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || o.FirmwareVersion == nil {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *EquipmentTpm) HasFirmwareVersion() bool {
	if o != nil && o.FirmwareVersion != nil {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *EquipmentTpm) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetOwnership returns the Ownership field value if set, zero value otherwise.
func (o *EquipmentTpm) GetOwnership() string {
	if o == nil || o.Ownership == nil {
		var ret string
		return ret
	}
	return *o.Ownership
}

// GetOwnershipOk returns a tuple with the Ownership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpm) GetOwnershipOk() (*string, bool) {
	if o == nil || o.Ownership == nil {
		return nil, false
	}
	return o.Ownership, true
}

// HasOwnership returns a boolean if a field has been set.
func (o *EquipmentTpm) HasOwnership() bool {
	if o != nil && o.Ownership != nil {
		return true
	}

	return false
}

// SetOwnership gets a reference to the given string and assigns it to the Ownership field.
func (o *EquipmentTpm) SetOwnership(v string) {
	o.Ownership = &v
}

// GetTpmId returns the TpmId field value if set, zero value otherwise.
func (o *EquipmentTpm) GetTpmId() int64 {
	if o == nil || o.TpmId == nil {
		var ret int64
		return ret
	}
	return *o.TpmId
}

// GetTpmIdOk returns a tuple with the TpmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpm) GetTpmIdOk() (*int64, bool) {
	if o == nil || o.TpmId == nil {
		return nil, false
	}
	return o.TpmId, true
}

// HasTpmId returns a boolean if a field has been set.
func (o *EquipmentTpm) HasTpmId() bool {
	if o != nil && o.TpmId != nil {
		return true
	}

	return false
}

// SetTpmId gets a reference to the given int64 and assigns it to the TpmId field.
func (o *EquipmentTpm) SetTpmId(v int64) {
	o.TpmId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EquipmentTpm) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpm) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EquipmentTpm) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EquipmentTpm) SetVersion(v string) {
	o.Version = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise.
func (o *EquipmentTpm) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || o.ComputeBoard == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpm) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.ComputeBoard == nil {
		return nil, false
	}
	return o.ComputeBoard, true
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *EquipmentTpm) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard != nil {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given ComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *EquipmentTpm) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentTpm) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpm) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentTpm) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentTpm) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentTpm) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpm) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentTpm) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentTpm) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentTpm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ActivationStatus != nil {
		toSerialize["ActivationStatus"] = o.ActivationStatus
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.FirmwareVersion != nil {
		toSerialize["FirmwareVersion"] = o.FirmwareVersion
	}
	if o.Ownership != nil {
		toSerialize["Ownership"] = o.Ownership
	}
	if o.TpmId != nil {
		toSerialize["TpmId"] = o.TpmId
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.ComputeBoard != nil {
		toSerialize["ComputeBoard"] = o.ComputeBoard
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

func (o *EquipmentTpm) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentTpmWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Identifies the activation status of the TPM.
		ActivationStatus *string `json:"ActivationStatus,omitempty"`
		// Identifies the admin configured state of the TPM.
		AdminState *string `json:"AdminState,omitempty"`
		// Firmware Version of the Trusted Platform Module.
		FirmwareVersion *string `json:"FirmwareVersion,omitempty"`
		// Identifies the ownership information of the TPM.
		Ownership *string `json:"Ownership,omitempty"`
		// Enter  the ID of the trusted platform module.
		TpmId *int64 `json:"TpmId,omitempty"`
		// Identifies the version of the Trusted Platform Module.
		Version *string `json:"Version,omitempty"`
		ComputeBoard *ComputeBoardRelationship `json:"ComputeBoard,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEquipmentTpmWithoutEmbeddedStruct := EquipmentTpmWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentTpmWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentTpm := _EquipmentTpm{}
		varEquipmentTpm.ClassId = varEquipmentTpmWithoutEmbeddedStruct.ClassId
		varEquipmentTpm.ObjectType = varEquipmentTpmWithoutEmbeddedStruct.ObjectType
		varEquipmentTpm.ActivationStatus = varEquipmentTpmWithoutEmbeddedStruct.ActivationStatus
		varEquipmentTpm.AdminState = varEquipmentTpmWithoutEmbeddedStruct.AdminState
		varEquipmentTpm.FirmwareVersion = varEquipmentTpmWithoutEmbeddedStruct.FirmwareVersion
		varEquipmentTpm.Ownership = varEquipmentTpmWithoutEmbeddedStruct.Ownership
		varEquipmentTpm.TpmId = varEquipmentTpmWithoutEmbeddedStruct.TpmId
		varEquipmentTpm.Version = varEquipmentTpmWithoutEmbeddedStruct.Version
		varEquipmentTpm.ComputeBoard = varEquipmentTpmWithoutEmbeddedStruct.ComputeBoard
		varEquipmentTpm.InventoryDeviceInfo = varEquipmentTpmWithoutEmbeddedStruct.InventoryDeviceInfo
		varEquipmentTpm.RegisteredDevice = varEquipmentTpmWithoutEmbeddedStruct.RegisteredDevice
		*o = EquipmentTpm(varEquipmentTpm)
	} else {
		return err
	}

	varEquipmentTpm := _EquipmentTpm{}

	err = json.Unmarshal(bytes, &varEquipmentTpm)
	if err == nil {
		o.EquipmentBase = varEquipmentTpm.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActivationStatus")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "FirmwareVersion")
		delete(additionalProperties, "Ownership")
		delete(additionalProperties, "TpmId")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableEquipmentTpm struct {
	value *EquipmentTpm
	isSet bool
}

func (v NullableEquipmentTpm) Get() *EquipmentTpm {
	return v.value
}

func (v *NullableEquipmentTpm) Set(val *EquipmentTpm) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentTpm) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentTpm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentTpm(val *EquipmentTpm) *NullableEquipmentTpm {
	return &NullableEquipmentTpm{value: val, isSet: true}
}

func (v NullableEquipmentTpm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentTpm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


