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

// CapabilitySiocModuleDescriptor Descriptor that uniquely identifies an SIOC module.
type CapabilitySiocModuleDescriptor struct {
	CapabilityHardwareDescriptor
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Revision for the SIOC module.
	Revision *string `json:"Revision,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitySiocModuleDescriptor CapabilitySiocModuleDescriptor

// NewCapabilitySiocModuleDescriptor instantiates a new CapabilitySiocModuleDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySiocModuleDescriptor(classId string, objectType string) *CapabilitySiocModuleDescriptor {
	this := CapabilitySiocModuleDescriptor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilitySiocModuleDescriptorWithDefaults instantiates a new CapabilitySiocModuleDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySiocModuleDescriptorWithDefaults() *CapabilitySiocModuleDescriptor {
	this := CapabilitySiocModuleDescriptor{}
	var classId string = "capability.SiocModuleDescriptor"
	this.ClassId = classId
	var objectType string = "capability.SiocModuleDescriptor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilitySiocModuleDescriptor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilitySiocModuleDescriptor) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilitySiocModuleDescriptor) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilitySiocModuleDescriptor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilitySiocModuleDescriptor) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilitySiocModuleDescriptor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *CapabilitySiocModuleDescriptor) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySiocModuleDescriptor) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *CapabilitySiocModuleDescriptor) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *CapabilitySiocModuleDescriptor) SetRevision(v string) {
	o.Revision = &v
}

func (o CapabilitySiocModuleDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityHardwareDescriptor, errCapabilityHardwareDescriptor := json.Marshal(o.CapabilityHardwareDescriptor)
	if errCapabilityHardwareDescriptor != nil {
		return []byte{}, errCapabilityHardwareDescriptor
	}
	errCapabilityHardwareDescriptor = json.Unmarshal([]byte(serializedCapabilityHardwareDescriptor), &toSerialize)
	if errCapabilityHardwareDescriptor != nil {
		return []byte{}, errCapabilityHardwareDescriptor
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Revision != nil {
		toSerialize["Revision"] = o.Revision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitySiocModuleDescriptor) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilitySiocModuleDescriptorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Revision for the SIOC module.
		Revision *string `json:"Revision,omitempty"`
	}

	varCapabilitySiocModuleDescriptorWithoutEmbeddedStruct := CapabilitySiocModuleDescriptorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilitySiocModuleDescriptorWithoutEmbeddedStruct)
	if err == nil {
		varCapabilitySiocModuleDescriptor := _CapabilitySiocModuleDescriptor{}
		varCapabilitySiocModuleDescriptor.ClassId = varCapabilitySiocModuleDescriptorWithoutEmbeddedStruct.ClassId
		varCapabilitySiocModuleDescriptor.ObjectType = varCapabilitySiocModuleDescriptorWithoutEmbeddedStruct.ObjectType
		varCapabilitySiocModuleDescriptor.Revision = varCapabilitySiocModuleDescriptorWithoutEmbeddedStruct.Revision
		*o = CapabilitySiocModuleDescriptor(varCapabilitySiocModuleDescriptor)
	} else {
		return err
	}

	varCapabilitySiocModuleDescriptor := _CapabilitySiocModuleDescriptor{}

	err = json.Unmarshal(bytes, &varCapabilitySiocModuleDescriptor)
	if err == nil {
		o.CapabilityHardwareDescriptor = varCapabilitySiocModuleDescriptor.CapabilityHardwareDescriptor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Revision")

		// remove fields from embedded structs
		reflectCapabilityHardwareDescriptor := reflect.ValueOf(o.CapabilityHardwareDescriptor)
		for i := 0; i < reflectCapabilityHardwareDescriptor.Type().NumField(); i++ {
			t := reflectCapabilityHardwareDescriptor.Type().Field(i)

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

type NullableCapabilitySiocModuleDescriptor struct {
	value *CapabilitySiocModuleDescriptor
	isSet bool
}

func (v NullableCapabilitySiocModuleDescriptor) Get() *CapabilitySiocModuleDescriptor {
	return v.value
}

func (v *NullableCapabilitySiocModuleDescriptor) Set(val *CapabilitySiocModuleDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySiocModuleDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySiocModuleDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySiocModuleDescriptor(val *CapabilitySiocModuleDescriptor) *NullableCapabilitySiocModuleDescriptor {
	return &NullableCapabilitySiocModuleDescriptor{value: val, isSet: true}
}

func (v NullableCapabilitySiocModuleDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySiocModuleDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


