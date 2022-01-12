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
)

// VirtualizationBasePlacementAllOf Definition of the list of properties defined in 'virtualization.BasePlacement', excluding properties defined in parent classes.
type VirtualizationBasePlacementAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Name of the virtual machine placement. It is the name of the VPC (Virtual Private Cloud) in case of AWS virtual machine, and datacenter name in case of VMware virtual machine.
	Name *string `json:"Name,omitempty"`
	// The uuid of this placement. The uuid is internally generated and not user specified.
	Uuid *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBasePlacementAllOf VirtualizationBasePlacementAllOf

// NewVirtualizationBasePlacementAllOf instantiates a new VirtualizationBasePlacementAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBasePlacementAllOf(classId string, objectType string) *VirtualizationBasePlacementAllOf {
	this := VirtualizationBasePlacementAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationBasePlacementAllOfWithDefaults instantiates a new VirtualizationBasePlacementAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBasePlacementAllOfWithDefaults() *VirtualizationBasePlacementAllOf {
	this := VirtualizationBasePlacementAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBasePlacementAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBasePlacementAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBasePlacementAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBasePlacementAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBasePlacementAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBasePlacementAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBasePlacementAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBasePlacementAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBasePlacementAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBasePlacementAllOf) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *VirtualizationBasePlacementAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBasePlacementAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *VirtualizationBasePlacementAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *VirtualizationBasePlacementAllOf) SetUuid(v string) {
	o.Uuid = &v
}

func (o VirtualizationBasePlacementAllOf) MarshalJSON() ([]byte, error) {
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
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBasePlacementAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationBasePlacementAllOf := _VirtualizationBasePlacementAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationBasePlacementAllOf); err == nil {
		*o = VirtualizationBasePlacementAllOf(varVirtualizationBasePlacementAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationBasePlacementAllOf struct {
	value *VirtualizationBasePlacementAllOf
	isSet bool
}

func (v NullableVirtualizationBasePlacementAllOf) Get() *VirtualizationBasePlacementAllOf {
	return v.value
}

func (v *NullableVirtualizationBasePlacementAllOf) Set(val *VirtualizationBasePlacementAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBasePlacementAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBasePlacementAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBasePlacementAllOf(val *VirtualizationBasePlacementAllOf) *NullableVirtualizationBasePlacementAllOf {
	return &NullableVirtualizationBasePlacementAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationBasePlacementAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBasePlacementAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


