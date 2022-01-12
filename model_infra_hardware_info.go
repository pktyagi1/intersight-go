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

// InfraHardwareInfo Information about the hardware platform (cpu, memory).
type InfraHardwareInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of cpu cores on this hardware platform.
	CpuCores *int64 `json:"CpuCores,omitempty"`
	// Speed of cpu in MHz. Usually cpu speeds are reported for modern cpus in GHz but MHz makes it more precise.
	CpuSpeed *int64 `json:"CpuSpeed,omitempty"`
	// The amount of memory allocated (bytes) to this hardware platform.
	MemorySize *int64 `json:"MemorySize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InfraHardwareInfo InfraHardwareInfo

// NewInfraHardwareInfo instantiates a new InfraHardwareInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraHardwareInfo(classId string, objectType string) *InfraHardwareInfo {
	this := InfraHardwareInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInfraHardwareInfoWithDefaults instantiates a new InfraHardwareInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraHardwareInfoWithDefaults() *InfraHardwareInfo {
	this := InfraHardwareInfo{}
	var classId string = "infra.HardwareInfo"
	this.ClassId = classId
	var objectType string = "infra.HardwareInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *InfraHardwareInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InfraHardwareInfo) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InfraHardwareInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *InfraHardwareInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InfraHardwareInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InfraHardwareInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCpuCores returns the CpuCores field value if set, zero value otherwise.
func (o *InfraHardwareInfo) GetCpuCores() int64 {
	if o == nil || o.CpuCores == nil {
		var ret int64
		return ret
	}
	return *o.CpuCores
}

// GetCpuCoresOk returns a tuple with the CpuCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHardwareInfo) GetCpuCoresOk() (*int64, bool) {
	if o == nil || o.CpuCores == nil {
		return nil, false
	}
	return o.CpuCores, true
}

// HasCpuCores returns a boolean if a field has been set.
func (o *InfraHardwareInfo) HasCpuCores() bool {
	if o != nil && o.CpuCores != nil {
		return true
	}

	return false
}

// SetCpuCores gets a reference to the given int64 and assigns it to the CpuCores field.
func (o *InfraHardwareInfo) SetCpuCores(v int64) {
	o.CpuCores = &v
}

// GetCpuSpeed returns the CpuSpeed field value if set, zero value otherwise.
func (o *InfraHardwareInfo) GetCpuSpeed() int64 {
	if o == nil || o.CpuSpeed == nil {
		var ret int64
		return ret
	}
	return *o.CpuSpeed
}

// GetCpuSpeedOk returns a tuple with the CpuSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHardwareInfo) GetCpuSpeedOk() (*int64, bool) {
	if o == nil || o.CpuSpeed == nil {
		return nil, false
	}
	return o.CpuSpeed, true
}

// HasCpuSpeed returns a boolean if a field has been set.
func (o *InfraHardwareInfo) HasCpuSpeed() bool {
	if o != nil && o.CpuSpeed != nil {
		return true
	}

	return false
}

// SetCpuSpeed gets a reference to the given int64 and assigns it to the CpuSpeed field.
func (o *InfraHardwareInfo) SetCpuSpeed(v int64) {
	o.CpuSpeed = &v
}

// GetMemorySize returns the MemorySize field value if set, zero value otherwise.
func (o *InfraHardwareInfo) GetMemorySize() int64 {
	if o == nil || o.MemorySize == nil {
		var ret int64
		return ret
	}
	return *o.MemorySize
}

// GetMemorySizeOk returns a tuple with the MemorySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHardwareInfo) GetMemorySizeOk() (*int64, bool) {
	if o == nil || o.MemorySize == nil {
		return nil, false
	}
	return o.MemorySize, true
}

// HasMemorySize returns a boolean if a field has been set.
func (o *InfraHardwareInfo) HasMemorySize() bool {
	if o != nil && o.MemorySize != nil {
		return true
	}

	return false
}

// SetMemorySize gets a reference to the given int64 and assigns it to the MemorySize field.
func (o *InfraHardwareInfo) SetMemorySize(v int64) {
	o.MemorySize = &v
}

func (o InfraHardwareInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CpuCores != nil {
		toSerialize["CpuCores"] = o.CpuCores
	}
	if o.CpuSpeed != nil {
		toSerialize["CpuSpeed"] = o.CpuSpeed
	}
	if o.MemorySize != nil {
		toSerialize["MemorySize"] = o.MemorySize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InfraHardwareInfo) UnmarshalJSON(bytes []byte) (err error) {
	type InfraHardwareInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The number of cpu cores on this hardware platform.
		CpuCores *int64 `json:"CpuCores,omitempty"`
		// Speed of cpu in MHz. Usually cpu speeds are reported for modern cpus in GHz but MHz makes it more precise.
		CpuSpeed *int64 `json:"CpuSpeed,omitempty"`
		// The amount of memory allocated (bytes) to this hardware platform.
		MemorySize *int64 `json:"MemorySize,omitempty"`
	}

	varInfraHardwareInfoWithoutEmbeddedStruct := InfraHardwareInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varInfraHardwareInfoWithoutEmbeddedStruct)
	if err == nil {
		varInfraHardwareInfo := _InfraHardwareInfo{}
		varInfraHardwareInfo.ClassId = varInfraHardwareInfoWithoutEmbeddedStruct.ClassId
		varInfraHardwareInfo.ObjectType = varInfraHardwareInfoWithoutEmbeddedStruct.ObjectType
		varInfraHardwareInfo.CpuCores = varInfraHardwareInfoWithoutEmbeddedStruct.CpuCores
		varInfraHardwareInfo.CpuSpeed = varInfraHardwareInfoWithoutEmbeddedStruct.CpuSpeed
		varInfraHardwareInfo.MemorySize = varInfraHardwareInfoWithoutEmbeddedStruct.MemorySize
		*o = InfraHardwareInfo(varInfraHardwareInfo)
	} else {
		return err
	}

	varInfraHardwareInfo := _InfraHardwareInfo{}

	err = json.Unmarshal(bytes, &varInfraHardwareInfo)
	if err == nil {
		o.MoBaseComplexType = varInfraHardwareInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CpuCores")
		delete(additionalProperties, "CpuSpeed")
		delete(additionalProperties, "MemorySize")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableInfraHardwareInfo struct {
	value *InfraHardwareInfo
	isSet bool
}

func (v NullableInfraHardwareInfo) Get() *InfraHardwareInfo {
	return v.value
}

func (v *NullableInfraHardwareInfo) Set(val *InfraHardwareInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraHardwareInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraHardwareInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraHardwareInfo(val *InfraHardwareInfo) *NullableInfraHardwareInfo {
	return &NullableInfraHardwareInfo{value: val, isSet: true}
}

func (v NullableInfraHardwareInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraHardwareInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

