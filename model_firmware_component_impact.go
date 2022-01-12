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

// FirmwareComponentImpact Impact for the components such as CIMC, BIOS, and disk during the upgrade operation.
type FirmwareComponentImpact struct {
	FirmwareBaseImpact
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Impact on the component after the upgrade. * `ALL` - This represents all the components. * `ALL,HDD` - This represents all the components plus the HDDs. * `Drive-U.2` - This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category). * `Storage` - This represents the storage controller components. * `None` - This represents none of the components. * `NXOS` - This represents NXOS components. * `IOM` - This represents IOM components. * `PSU` - This represents PSU components. * `CIMC` - This represents CIMC components. * `BIOS` - This represents BIOS components. * `PCIE` - This represents PCIE components. * `Drive` - This represents Drive components. * `DIMM` - This represents DIMM components. * `BoardController` - This represents Board Controller components. * `StorageController` - This represents Storage Controller components. * `Storage-Sasexpander` - This represents Storage Sas-Expander components. * `Storage-U.2` - This represents U2 Storage Controller components. * `HBA` - This represents HBA components. * `GPU` - This represents GPU components. * `SasExpander` - This represents SasExpander components. * `MSwitch` - This represents mSwitch components. * `CMC` - This represents CMC components.
	Component *string `json:"Component,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareComponentImpact FirmwareComponentImpact

// NewFirmwareComponentImpact instantiates a new FirmwareComponentImpact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareComponentImpact(classId string, objectType string) *FirmwareComponentImpact {
	this := FirmwareComponentImpact{}
	this.ClassId = classId
	this.ObjectType = objectType
	var computationStatusDetail string = "Inprogress"
	this.ComputationStatusDetail = &computationStatusDetail
	var impactType string = "NoReboot"
	this.ImpactType = &impactType
	var versionImpact string = "None"
	this.VersionImpact = &versionImpact
	var component string = "ALL"
	this.Component = &component
	return &this
}

// NewFirmwareComponentImpactWithDefaults instantiates a new FirmwareComponentImpact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareComponentImpactWithDefaults() *FirmwareComponentImpact {
	this := FirmwareComponentImpact{}
	var classId string = "firmware.ComponentImpact"
	this.ClassId = classId
	var objectType string = "firmware.ComponentImpact"
	this.ObjectType = objectType
	var component string = "ALL"
	this.Component = &component
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareComponentImpact) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareComponentImpact) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareComponentImpact) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareComponentImpact) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareComponentImpact) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareComponentImpact) SetObjectType(v string) {
	o.ObjectType = v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *FirmwareComponentImpact) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentImpact) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *FirmwareComponentImpact) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *FirmwareComponentImpact) SetComponent(v string) {
	o.Component = &v
}

func (o FirmwareComponentImpact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareBaseImpact, errFirmwareBaseImpact := json.Marshal(o.FirmwareBaseImpact)
	if errFirmwareBaseImpact != nil {
		return []byte{}, errFirmwareBaseImpact
	}
	errFirmwareBaseImpact = json.Unmarshal([]byte(serializedFirmwareBaseImpact), &toSerialize)
	if errFirmwareBaseImpact != nil {
		return []byte{}, errFirmwareBaseImpact
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Component != nil {
		toSerialize["Component"] = o.Component
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareComponentImpact) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareComponentImpactWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Impact on the component after the upgrade. * `ALL` - This represents all the components. * `ALL,HDD` - This represents all the components plus the HDDs. * `Drive-U.2` - This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category). * `Storage` - This represents the storage controller components. * `None` - This represents none of the components. * `NXOS` - This represents NXOS components. * `IOM` - This represents IOM components. * `PSU` - This represents PSU components. * `CIMC` - This represents CIMC components. * `BIOS` - This represents BIOS components. * `PCIE` - This represents PCIE components. * `Drive` - This represents Drive components. * `DIMM` - This represents DIMM components. * `BoardController` - This represents Board Controller components. * `StorageController` - This represents Storage Controller components. * `Storage-Sasexpander` - This represents Storage Sas-Expander components. * `Storage-U.2` - This represents U2 Storage Controller components. * `HBA` - This represents HBA components. * `GPU` - This represents GPU components. * `SasExpander` - This represents SasExpander components. * `MSwitch` - This represents mSwitch components. * `CMC` - This represents CMC components.
		Component *string `json:"Component,omitempty"`
	}

	varFirmwareComponentImpactWithoutEmbeddedStruct := FirmwareComponentImpactWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareComponentImpactWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareComponentImpact := _FirmwareComponentImpact{}
		varFirmwareComponentImpact.ClassId = varFirmwareComponentImpactWithoutEmbeddedStruct.ClassId
		varFirmwareComponentImpact.ObjectType = varFirmwareComponentImpactWithoutEmbeddedStruct.ObjectType
		varFirmwareComponentImpact.Component = varFirmwareComponentImpactWithoutEmbeddedStruct.Component
		*o = FirmwareComponentImpact(varFirmwareComponentImpact)
	} else {
		return err
	}

	varFirmwareComponentImpact := _FirmwareComponentImpact{}

	err = json.Unmarshal(bytes, &varFirmwareComponentImpact)
	if err == nil {
		o.FirmwareBaseImpact = varFirmwareComponentImpact.FirmwareBaseImpact
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Component")

		// remove fields from embedded structs
		reflectFirmwareBaseImpact := reflect.ValueOf(o.FirmwareBaseImpact)
		for i := 0; i < reflectFirmwareBaseImpact.Type().NumField(); i++ {
			t := reflectFirmwareBaseImpact.Type().Field(i)

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

type NullableFirmwareComponentImpact struct {
	value *FirmwareComponentImpact
	isSet bool
}

func (v NullableFirmwareComponentImpact) Get() *FirmwareComponentImpact {
	return v.value
}

func (v *NullableFirmwareComponentImpact) Set(val *FirmwareComponentImpact) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareComponentImpact) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareComponentImpact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareComponentImpact(val *FirmwareComponentImpact) *NullableFirmwareComponentImpact {
	return &NullableFirmwareComponentImpact{value: val, isSet: true}
}

func (v NullableFirmwareComponentImpact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareComponentImpact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


