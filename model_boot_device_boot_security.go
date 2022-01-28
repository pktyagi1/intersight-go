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
	"reflect"
	"strings"
)

// BootDeviceBootSecurity Boot Security of the devices that BIOS uses to boot them.
type BootDeviceBootSecurity struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The user desired BIOS secure boot as configured in the boot policy.
	SecureBoot *string `json:"SecureBoot,omitempty"`
	ComputePhysical *ComputePhysicalRelationship `json:"ComputePhysical,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootDeviceBootSecurity BootDeviceBootSecurity

// NewBootDeviceBootSecurity instantiates a new BootDeviceBootSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootDeviceBootSecurity(classId string, objectType string) *BootDeviceBootSecurity {
	this := BootDeviceBootSecurity{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBootDeviceBootSecurityWithDefaults instantiates a new BootDeviceBootSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootDeviceBootSecurityWithDefaults() *BootDeviceBootSecurity {
	this := BootDeviceBootSecurity{}
	var classId string = "boot.DeviceBootSecurity"
	this.ClassId = classId
	var objectType string = "boot.DeviceBootSecurity"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootDeviceBootSecurity) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootDeviceBootSecurity) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootDeviceBootSecurity) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootDeviceBootSecurity) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootDeviceBootSecurity) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootDeviceBootSecurity) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSecureBoot returns the SecureBoot field value if set, zero value otherwise.
func (o *BootDeviceBootSecurity) GetSecureBoot() string {
	if o == nil || o.SecureBoot == nil {
		var ret string
		return ret
	}
	return *o.SecureBoot
}

// GetSecureBootOk returns a tuple with the SecureBoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootDeviceBootSecurity) GetSecureBootOk() (*string, bool) {
	if o == nil || o.SecureBoot == nil {
		return nil, false
	}
	return o.SecureBoot, true
}

// HasSecureBoot returns a boolean if a field has been set.
func (o *BootDeviceBootSecurity) HasSecureBoot() bool {
	if o != nil && o.SecureBoot != nil {
		return true
	}

	return false
}

// SetSecureBoot gets a reference to the given string and assigns it to the SecureBoot field.
func (o *BootDeviceBootSecurity) SetSecureBoot(v string) {
	o.SecureBoot = &v
}

// GetComputePhysical returns the ComputePhysical field value if set, zero value otherwise.
func (o *BootDeviceBootSecurity) GetComputePhysical() ComputePhysicalRelationship {
	if o == nil || o.ComputePhysical == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.ComputePhysical
}

// GetComputePhysicalOk returns a tuple with the ComputePhysical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootDeviceBootSecurity) GetComputePhysicalOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.ComputePhysical == nil {
		return nil, false
	}
	return o.ComputePhysical, true
}

// HasComputePhysical returns a boolean if a field has been set.
func (o *BootDeviceBootSecurity) HasComputePhysical() bool {
	if o != nil && o.ComputePhysical != nil {
		return true
	}

	return false
}

// SetComputePhysical gets a reference to the given ComputePhysicalRelationship and assigns it to the ComputePhysical field.
func (o *BootDeviceBootSecurity) SetComputePhysical(v ComputePhysicalRelationship) {
	o.ComputePhysical = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *BootDeviceBootSecurity) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootDeviceBootSecurity) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *BootDeviceBootSecurity) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *BootDeviceBootSecurity) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *BootDeviceBootSecurity) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootDeviceBootSecurity) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BootDeviceBootSecurity) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BootDeviceBootSecurity) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o BootDeviceBootSecurity) MarshalJSON() ([]byte, error) {
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
	if o.SecureBoot != nil {
		toSerialize["SecureBoot"] = o.SecureBoot
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

func (o *BootDeviceBootSecurity) UnmarshalJSON(bytes []byte) (err error) {
	type BootDeviceBootSecurityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The user desired BIOS secure boot as configured in the boot policy.
		SecureBoot *string `json:"SecureBoot,omitempty"`
		ComputePhysical *ComputePhysicalRelationship `json:"ComputePhysical,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varBootDeviceBootSecurityWithoutEmbeddedStruct := BootDeviceBootSecurityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBootDeviceBootSecurityWithoutEmbeddedStruct)
	if err == nil {
		varBootDeviceBootSecurity := _BootDeviceBootSecurity{}
		varBootDeviceBootSecurity.ClassId = varBootDeviceBootSecurityWithoutEmbeddedStruct.ClassId
		varBootDeviceBootSecurity.ObjectType = varBootDeviceBootSecurityWithoutEmbeddedStruct.ObjectType
		varBootDeviceBootSecurity.SecureBoot = varBootDeviceBootSecurityWithoutEmbeddedStruct.SecureBoot
		varBootDeviceBootSecurity.ComputePhysical = varBootDeviceBootSecurityWithoutEmbeddedStruct.ComputePhysical
		varBootDeviceBootSecurity.InventoryDeviceInfo = varBootDeviceBootSecurityWithoutEmbeddedStruct.InventoryDeviceInfo
		varBootDeviceBootSecurity.RegisteredDevice = varBootDeviceBootSecurityWithoutEmbeddedStruct.RegisteredDevice
		*o = BootDeviceBootSecurity(varBootDeviceBootSecurity)
	} else {
		return err
	}

	varBootDeviceBootSecurity := _BootDeviceBootSecurity{}

	err = json.Unmarshal(bytes, &varBootDeviceBootSecurity)
	if err == nil {
		o.InventoryBase = varBootDeviceBootSecurity.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SecureBoot")
		delete(additionalProperties, "ComputePhysical")
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

type NullableBootDeviceBootSecurity struct {
	value *BootDeviceBootSecurity
	isSet bool
}

func (v NullableBootDeviceBootSecurity) Get() *BootDeviceBootSecurity {
	return v.value
}

func (v *NullableBootDeviceBootSecurity) Set(val *BootDeviceBootSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableBootDeviceBootSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableBootDeviceBootSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootDeviceBootSecurity(val *BootDeviceBootSecurity) *NullableBootDeviceBootSecurity {
	return &NullableBootDeviceBootSecurity{value: val, isSet: true}
}

func (v NullableBootDeviceBootSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootDeviceBootSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


