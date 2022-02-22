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

// CapabilityEquipmentPhysicalDefAllOf Definition of the list of properties defined in 'capability.EquipmentPhysicalDef', excluding properties defined in parent classes.
type CapabilityEquipmentPhysicalDefAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Depth information for a Switch/Fabric-Interconnect.
	Depth *float32 `json:"Depth,omitempty"`
	// Height information for a Switch/Fabric-Interconnect.
	Height *float32 `json:"Height,omitempty"`
	// Max Power information for a Switch/Fabric-Interconnect.
	MaxPower *int64 `json:"MaxPower,omitempty"`
	// Min Power information for a Switch/Fabric-Interconnect.
	MinPower *int64 `json:"MinPower,omitempty"`
	// Nominal Power information for a Switch/Fabric-Interconnect.
	NominalPower *int64 `json:"NominalPower,omitempty"`
	// Weight information for a Switch/Fabric-Interconnect.
	Weight *float32 `json:"Weight,omitempty"`
	// Width information for a Switch/Fabric-Interconnect.
	Width *float32 `json:"Width,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityEquipmentPhysicalDefAllOf CapabilityEquipmentPhysicalDefAllOf

// NewCapabilityEquipmentPhysicalDefAllOf instantiates a new CapabilityEquipmentPhysicalDefAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityEquipmentPhysicalDefAllOf(classId string, objectType string) *CapabilityEquipmentPhysicalDefAllOf {
	this := CapabilityEquipmentPhysicalDefAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityEquipmentPhysicalDefAllOfWithDefaults instantiates a new CapabilityEquipmentPhysicalDefAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityEquipmentPhysicalDefAllOfWithDefaults() *CapabilityEquipmentPhysicalDefAllOf {
	this := CapabilityEquipmentPhysicalDefAllOf{}
	var classId string = "capability.EquipmentPhysicalDef"
	this.ClassId = classId
	var objectType string = "capability.EquipmentPhysicalDef"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityEquipmentPhysicalDefAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityEquipmentPhysicalDefAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityEquipmentPhysicalDefAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityEquipmentPhysicalDefAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDepth returns the Depth field value if set, zero value otherwise.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetDepth() float32 {
	if o == nil || o.Depth == nil {
		var ret float32
		return ret
	}
	return *o.Depth
}

// GetDepthOk returns a tuple with the Depth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetDepthOk() (*float32, bool) {
	if o == nil || o.Depth == nil {
		return nil, false
	}
	return o.Depth, true
}

// HasDepth returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) HasDepth() bool {
	if o != nil && o.Depth != nil {
		return true
	}

	return false
}

// SetDepth gets a reference to the given float32 and assigns it to the Depth field.
func (o *CapabilityEquipmentPhysicalDefAllOf) SetDepth(v float32) {
	o.Depth = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetHeight() float32 {
	if o == nil || o.Height == nil {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetHeightOk() (*float32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *CapabilityEquipmentPhysicalDefAllOf) SetHeight(v float32) {
	o.Height = &v
}

// GetMaxPower returns the MaxPower field value if set, zero value otherwise.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetMaxPower() int64 {
	if o == nil || o.MaxPower == nil {
		var ret int64
		return ret
	}
	return *o.MaxPower
}

// GetMaxPowerOk returns a tuple with the MaxPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetMaxPowerOk() (*int64, bool) {
	if o == nil || o.MaxPower == nil {
		return nil, false
	}
	return o.MaxPower, true
}

// HasMaxPower returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) HasMaxPower() bool {
	if o != nil && o.MaxPower != nil {
		return true
	}

	return false
}

// SetMaxPower gets a reference to the given int64 and assigns it to the MaxPower field.
func (o *CapabilityEquipmentPhysicalDefAllOf) SetMaxPower(v int64) {
	o.MaxPower = &v
}

// GetMinPower returns the MinPower field value if set, zero value otherwise.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetMinPower() int64 {
	if o == nil || o.MinPower == nil {
		var ret int64
		return ret
	}
	return *o.MinPower
}

// GetMinPowerOk returns a tuple with the MinPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetMinPowerOk() (*int64, bool) {
	if o == nil || o.MinPower == nil {
		return nil, false
	}
	return o.MinPower, true
}

// HasMinPower returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) HasMinPower() bool {
	if o != nil && o.MinPower != nil {
		return true
	}

	return false
}

// SetMinPower gets a reference to the given int64 and assigns it to the MinPower field.
func (o *CapabilityEquipmentPhysicalDefAllOf) SetMinPower(v int64) {
	o.MinPower = &v
}

// GetNominalPower returns the NominalPower field value if set, zero value otherwise.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetNominalPower() int64 {
	if o == nil || o.NominalPower == nil {
		var ret int64
		return ret
	}
	return *o.NominalPower
}

// GetNominalPowerOk returns a tuple with the NominalPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetNominalPowerOk() (*int64, bool) {
	if o == nil || o.NominalPower == nil {
		return nil, false
	}
	return o.NominalPower, true
}

// HasNominalPower returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) HasNominalPower() bool {
	if o != nil && o.NominalPower != nil {
		return true
	}

	return false
}

// SetNominalPower gets a reference to the given int64 and assigns it to the NominalPower field.
func (o *CapabilityEquipmentPhysicalDefAllOf) SetNominalPower(v int64) {
	o.NominalPower = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetWeight() float32 {
	if o == nil || o.Weight == nil {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetWeightOk() (*float32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *CapabilityEquipmentPhysicalDefAllOf) SetWeight(v float32) {
	o.Weight = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetWidth() float32 {
	if o == nil || o.Width == nil {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) GetWidthOk() (*float32, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDefAllOf) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *CapabilityEquipmentPhysicalDefAllOf) SetWidth(v float32) {
	o.Width = &v
}

func (o CapabilityEquipmentPhysicalDefAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Depth != nil {
		toSerialize["Depth"] = o.Depth
	}
	if o.Height != nil {
		toSerialize["Height"] = o.Height
	}
	if o.MaxPower != nil {
		toSerialize["MaxPower"] = o.MaxPower
	}
	if o.MinPower != nil {
		toSerialize["MinPower"] = o.MinPower
	}
	if o.NominalPower != nil {
		toSerialize["NominalPower"] = o.NominalPower
	}
	if o.Weight != nil {
		toSerialize["Weight"] = o.Weight
	}
	if o.Width != nil {
		toSerialize["Width"] = o.Width
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityEquipmentPhysicalDefAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilityEquipmentPhysicalDefAllOf := _CapabilityEquipmentPhysicalDefAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilityEquipmentPhysicalDefAllOf); err == nil {
		*o = CapabilityEquipmentPhysicalDefAllOf(varCapabilityEquipmentPhysicalDefAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Depth")
		delete(additionalProperties, "Height")
		delete(additionalProperties, "MaxPower")
		delete(additionalProperties, "MinPower")
		delete(additionalProperties, "NominalPower")
		delete(additionalProperties, "Weight")
		delete(additionalProperties, "Width")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilityEquipmentPhysicalDefAllOf struct {
	value *CapabilityEquipmentPhysicalDefAllOf
	isSet bool
}

func (v NullableCapabilityEquipmentPhysicalDefAllOf) Get() *CapabilityEquipmentPhysicalDefAllOf {
	return v.value
}

func (v *NullableCapabilityEquipmentPhysicalDefAllOf) Set(val *CapabilityEquipmentPhysicalDefAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityEquipmentPhysicalDefAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityEquipmentPhysicalDefAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityEquipmentPhysicalDefAllOf(val *CapabilityEquipmentPhysicalDefAllOf) *NullableCapabilityEquipmentPhysicalDefAllOf {
	return &NullableCapabilityEquipmentPhysicalDefAllOf{value: val, isSet: true}
}

func (v NullableCapabilityEquipmentPhysicalDefAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityEquipmentPhysicalDefAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


