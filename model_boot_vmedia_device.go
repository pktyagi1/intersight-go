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

// BootVmediaDevice Virtual Media Boot Device configured on the server.
type BootVmediaDevice struct {
	BootConfiguredDevice
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	ComputePhysical *ComputePhysicalRelationship `json:"ComputePhysical,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootVmediaDevice BootVmediaDevice

// NewBootVmediaDevice instantiates a new BootVmediaDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootVmediaDevice(classId string, objectType string) *BootVmediaDevice {
	this := BootVmediaDevice{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBootVmediaDeviceWithDefaults instantiates a new BootVmediaDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootVmediaDeviceWithDefaults() *BootVmediaDevice {
	this := BootVmediaDevice{}
	var classId string = "boot.VmediaDevice"
	this.ClassId = classId
	var objectType string = "boot.VmediaDevice"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootVmediaDevice) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootVmediaDevice) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootVmediaDevice) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootVmediaDevice) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootVmediaDevice) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootVmediaDevice) SetObjectType(v string) {
	o.ObjectType = v
}

// GetComputePhysical returns the ComputePhysical field value if set, zero value otherwise.
func (o *BootVmediaDevice) GetComputePhysical() ComputePhysicalRelationship {
	if o == nil || o.ComputePhysical == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.ComputePhysical
}

// GetComputePhysicalOk returns a tuple with the ComputePhysical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootVmediaDevice) GetComputePhysicalOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.ComputePhysical == nil {
		return nil, false
	}
	return o.ComputePhysical, true
}

// HasComputePhysical returns a boolean if a field has been set.
func (o *BootVmediaDevice) HasComputePhysical() bool {
	if o != nil && o.ComputePhysical != nil {
		return true
	}

	return false
}

// SetComputePhysical gets a reference to the given ComputePhysicalRelationship and assigns it to the ComputePhysical field.
func (o *BootVmediaDevice) SetComputePhysical(v ComputePhysicalRelationship) {
	o.ComputePhysical = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *BootVmediaDevice) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootVmediaDevice) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *BootVmediaDevice) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *BootVmediaDevice) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *BootVmediaDevice) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootVmediaDevice) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BootVmediaDevice) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BootVmediaDevice) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o BootVmediaDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBootConfiguredDevice, errBootConfiguredDevice := json.Marshal(o.BootConfiguredDevice)
	if errBootConfiguredDevice != nil {
		return []byte{}, errBootConfiguredDevice
	}
	errBootConfiguredDevice = json.Unmarshal([]byte(serializedBootConfiguredDevice), &toSerialize)
	if errBootConfiguredDevice != nil {
		return []byte{}, errBootConfiguredDevice
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ComputePhysical != nil {
		toSerialize["ComputePhysical"] = o.ComputePhysical
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

func (o *BootVmediaDevice) UnmarshalJSON(bytes []byte) (err error) {
	type BootVmediaDeviceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		ComputePhysical *ComputePhysicalRelationship `json:"ComputePhysical,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varBootVmediaDeviceWithoutEmbeddedStruct := BootVmediaDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBootVmediaDeviceWithoutEmbeddedStruct)
	if err == nil {
		varBootVmediaDevice := _BootVmediaDevice{}
		varBootVmediaDevice.ClassId = varBootVmediaDeviceWithoutEmbeddedStruct.ClassId
		varBootVmediaDevice.ObjectType = varBootVmediaDeviceWithoutEmbeddedStruct.ObjectType
		varBootVmediaDevice.ComputePhysical = varBootVmediaDeviceWithoutEmbeddedStruct.ComputePhysical
		varBootVmediaDevice.InventoryDeviceInfo = varBootVmediaDeviceWithoutEmbeddedStruct.InventoryDeviceInfo
		varBootVmediaDevice.RegisteredDevice = varBootVmediaDeviceWithoutEmbeddedStruct.RegisteredDevice
		*o = BootVmediaDevice(varBootVmediaDevice)
	} else {
		return err
	}

	varBootVmediaDevice := _BootVmediaDevice{}

	err = json.Unmarshal(bytes, &varBootVmediaDevice)
	if err == nil {
		o.BootConfiguredDevice = varBootVmediaDevice.BootConfiguredDevice
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ComputePhysical")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectBootConfiguredDevice := reflect.ValueOf(o.BootConfiguredDevice)
		for i := 0; i < reflectBootConfiguredDevice.Type().NumField(); i++ {
			t := reflectBootConfiguredDevice.Type().Field(i)

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

type NullableBootVmediaDevice struct {
	value *BootVmediaDevice
	isSet bool
}

func (v NullableBootVmediaDevice) Get() *BootVmediaDevice {
	return v.value
}

func (v *NullableBootVmediaDevice) Set(val *BootVmediaDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableBootVmediaDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableBootVmediaDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootVmediaDevice(val *BootVmediaDevice) *NullableBootVmediaDevice {
	return &NullableBootVmediaDevice{value: val, isSet: true}
}

func (v NullableBootVmediaDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootVmediaDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


