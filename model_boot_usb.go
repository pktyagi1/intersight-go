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

// BootUsb Device type used when booting from USB device.
type BootUsb struct {
	BootDeviceBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The subtype for the selected device type. * `None` - No sub type for USB boot device. * `usb-cd` - Use of Compact Disk (CD) as sub-type for the USB boot device. * `usb-fdd` - Use of Floppy Disk Drive (FDD) as sub-type for the USB boot device. * `usb-hdd` - Use of Hard Disk Drive (HDD) as sub-type for the USB boot device.
	Subtype *string `json:"Subtype,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootUsb BootUsb

// NewBootUsb instantiates a new BootUsb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootUsb(classId string, objectType string) *BootUsb {
	this := BootUsb{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	var subtype string = "None"
	this.Subtype = &subtype
	return &this
}

// NewBootUsbWithDefaults instantiates a new BootUsb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootUsbWithDefaults() *BootUsb {
	this := BootUsb{}
	var classId string = "boot.Usb"
	this.ClassId = classId
	var objectType string = "boot.Usb"
	this.ObjectType = objectType
	var subtype string = "None"
	this.Subtype = &subtype
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootUsb) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootUsb) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootUsb) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootUsb) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootUsb) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootUsb) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *BootUsb) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootUsb) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *BootUsb) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *BootUsb) SetSubtype(v string) {
	o.Subtype = &v
}

func (o BootUsb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBootDeviceBase, errBootDeviceBase := json.Marshal(o.BootDeviceBase)
	if errBootDeviceBase != nil {
		return []byte{}, errBootDeviceBase
	}
	errBootDeviceBase = json.Unmarshal([]byte(serializedBootDeviceBase), &toSerialize)
	if errBootDeviceBase != nil {
		return []byte{}, errBootDeviceBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Subtype != nil {
		toSerialize["Subtype"] = o.Subtype
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootUsb) UnmarshalJSON(bytes []byte) (err error) {
	type BootUsbWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The subtype for the selected device type. * `None` - No sub type for USB boot device. * `usb-cd` - Use of Compact Disk (CD) as sub-type for the USB boot device. * `usb-fdd` - Use of Floppy Disk Drive (FDD) as sub-type for the USB boot device. * `usb-hdd` - Use of Hard Disk Drive (HDD) as sub-type for the USB boot device.
		Subtype *string `json:"Subtype,omitempty"`
	}

	varBootUsbWithoutEmbeddedStruct := BootUsbWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBootUsbWithoutEmbeddedStruct)
	if err == nil {
		varBootUsb := _BootUsb{}
		varBootUsb.ClassId = varBootUsbWithoutEmbeddedStruct.ClassId
		varBootUsb.ObjectType = varBootUsbWithoutEmbeddedStruct.ObjectType
		varBootUsb.Subtype = varBootUsbWithoutEmbeddedStruct.Subtype
		*o = BootUsb(varBootUsb)
	} else {
		return err
	}

	varBootUsb := _BootUsb{}

	err = json.Unmarshal(bytes, &varBootUsb)
	if err == nil {
		o.BootDeviceBase = varBootUsb.BootDeviceBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Subtype")

		// remove fields from embedded structs
		reflectBootDeviceBase := reflect.ValueOf(o.BootDeviceBase)
		for i := 0; i < reflectBootDeviceBase.Type().NumField(); i++ {
			t := reflectBootDeviceBase.Type().Field(i)

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

type NullableBootUsb struct {
	value *BootUsb
	isSet bool
}

func (v NullableBootUsb) Get() *BootUsb {
	return v.value
}

func (v *NullableBootUsb) Set(val *BootUsb) {
	v.value = val
	v.isSet = true
}

func (v NullableBootUsb) IsSet() bool {
	return v.isSet
}

func (v *NullableBootUsb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootUsb(val *BootUsb) *NullableBootUsb {
	return &NullableBootUsb{value: val, isSet: true}
}

func (v NullableBootUsb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootUsb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


