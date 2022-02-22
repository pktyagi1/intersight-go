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
)

// VirtualizationBaseVirtualDiskAllOf Definition of the list of properties defined in 'virtualization.BaseVirtualDisk', excluding properties defined in parent classes.
type VirtualizationBaseVirtualDiskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Name of the storage disk. Name must be unique within a Datastore.
	Name *string `json:"Name,omitempty"`
	// Disk size represented in bytes.
	Size *int64 `json:"Size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseVirtualDiskAllOf VirtualizationBaseVirtualDiskAllOf

// NewVirtualizationBaseVirtualDiskAllOf instantiates a new VirtualizationBaseVirtualDiskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseVirtualDiskAllOf(classId string, objectType string) *VirtualizationBaseVirtualDiskAllOf {
	this := VirtualizationBaseVirtualDiskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationBaseVirtualDiskAllOfWithDefaults instantiates a new VirtualizationBaseVirtualDiskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseVirtualDiskAllOfWithDefaults() *VirtualizationBaseVirtualDiskAllOf {
	this := VirtualizationBaseVirtualDiskAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBaseVirtualDiskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualDiskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBaseVirtualDiskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBaseVirtualDiskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualDiskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBaseVirtualDiskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualDiskAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualDiskAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualDiskAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseVirtualDiskAllOf) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualDiskAllOf) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualDiskAllOf) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualDiskAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *VirtualizationBaseVirtualDiskAllOf) SetSize(v int64) {
	o.Size = &v
}

func (o VirtualizationBaseVirtualDiskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBaseVirtualDiskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationBaseVirtualDiskAllOf := _VirtualizationBaseVirtualDiskAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationBaseVirtualDiskAllOf); err == nil {
		*o = VirtualizationBaseVirtualDiskAllOf(varVirtualizationBaseVirtualDiskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Size")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationBaseVirtualDiskAllOf struct {
	value *VirtualizationBaseVirtualDiskAllOf
	isSet bool
}

func (v NullableVirtualizationBaseVirtualDiskAllOf) Get() *VirtualizationBaseVirtualDiskAllOf {
	return v.value
}

func (v *NullableVirtualizationBaseVirtualDiskAllOf) Set(val *VirtualizationBaseVirtualDiskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseVirtualDiskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseVirtualDiskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseVirtualDiskAllOf(val *VirtualizationBaseVirtualDiskAllOf) *NullableVirtualizationBaseVirtualDiskAllOf {
	return &NullableVirtualizationBaseVirtualDiskAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationBaseVirtualDiskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseVirtualDiskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


