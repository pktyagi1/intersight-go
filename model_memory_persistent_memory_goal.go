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

// MemoryPersistentMemoryGoal A Persistent Memory Goal that needs to be applied on the associated servers through the policy. This would result in the creation of regions and allocation of volatile memory on the server.
type MemoryPersistentMemoryGoal struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Volatile memory percentage.
	MemoryModePercentage *int64 `json:"MemoryModePercentage,omitempty"`
	// Type of the Persistent Memory configuration where the Persistent Memory Modules are combined in an interleaved set or not. * `app-direct` - The App Direct interleaved Persistent Memory type. * `app-direct-non-interleaved` - The App Direct non-interleaved Persistent Memory type.
	PersistentMemoryType *string `json:"PersistentMemoryType,omitempty"`
	// CPU Socket ID to which this goal will be applied. * `All Sockets` - All the CPU socket IDs in a server.
	SocketId *string `json:"SocketId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemoryPersistentMemoryGoal MemoryPersistentMemoryGoal

// NewMemoryPersistentMemoryGoal instantiates a new MemoryPersistentMemoryGoal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPersistentMemoryGoal(classId string, objectType string) *MemoryPersistentMemoryGoal {
	this := MemoryPersistentMemoryGoal{}
	this.ClassId = classId
	this.ObjectType = objectType
	var persistentMemoryType string = "app-direct"
	this.PersistentMemoryType = &persistentMemoryType
	var socketId string = "All Sockets"
	this.SocketId = &socketId
	return &this
}

// NewMemoryPersistentMemoryGoalWithDefaults instantiates a new MemoryPersistentMemoryGoal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPersistentMemoryGoalWithDefaults() *MemoryPersistentMemoryGoal {
	this := MemoryPersistentMemoryGoal{}
	var classId string = "memory.PersistentMemoryGoal"
	this.ClassId = classId
	var objectType string = "memory.PersistentMemoryGoal"
	this.ObjectType = objectType
	var persistentMemoryType string = "app-direct"
	this.PersistentMemoryType = &persistentMemoryType
	var socketId string = "All Sockets"
	this.SocketId = &socketId
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryPersistentMemoryGoal) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryGoal) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryPersistentMemoryGoal) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MemoryPersistentMemoryGoal) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryGoal) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryPersistentMemoryGoal) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMemoryModePercentage returns the MemoryModePercentage field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryGoal) GetMemoryModePercentage() int64 {
	if o == nil || o.MemoryModePercentage == nil {
		var ret int64
		return ret
	}
	return *o.MemoryModePercentage
}

// GetMemoryModePercentageOk returns a tuple with the MemoryModePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryGoal) GetMemoryModePercentageOk() (*int64, bool) {
	if o == nil || o.MemoryModePercentage == nil {
		return nil, false
	}
	return o.MemoryModePercentage, true
}

// HasMemoryModePercentage returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryGoal) HasMemoryModePercentage() bool {
	if o != nil && o.MemoryModePercentage != nil {
		return true
	}

	return false
}

// SetMemoryModePercentage gets a reference to the given int64 and assigns it to the MemoryModePercentage field.
func (o *MemoryPersistentMemoryGoal) SetMemoryModePercentage(v int64) {
	o.MemoryModePercentage = &v
}

// GetPersistentMemoryType returns the PersistentMemoryType field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryGoal) GetPersistentMemoryType() string {
	if o == nil || o.PersistentMemoryType == nil {
		var ret string
		return ret
	}
	return *o.PersistentMemoryType
}

// GetPersistentMemoryTypeOk returns a tuple with the PersistentMemoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryGoal) GetPersistentMemoryTypeOk() (*string, bool) {
	if o == nil || o.PersistentMemoryType == nil {
		return nil, false
	}
	return o.PersistentMemoryType, true
}

// HasPersistentMemoryType returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryGoal) HasPersistentMemoryType() bool {
	if o != nil && o.PersistentMemoryType != nil {
		return true
	}

	return false
}

// SetPersistentMemoryType gets a reference to the given string and assigns it to the PersistentMemoryType field.
func (o *MemoryPersistentMemoryGoal) SetPersistentMemoryType(v string) {
	o.PersistentMemoryType = &v
}

// GetSocketId returns the SocketId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryGoal) GetSocketId() string {
	if o == nil || o.SocketId == nil {
		var ret string
		return ret
	}
	return *o.SocketId
}

// GetSocketIdOk returns a tuple with the SocketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryGoal) GetSocketIdOk() (*string, bool) {
	if o == nil || o.SocketId == nil {
		return nil, false
	}
	return o.SocketId, true
}

// HasSocketId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryGoal) HasSocketId() bool {
	if o != nil && o.SocketId != nil {
		return true
	}

	return false
}

// SetSocketId gets a reference to the given string and assigns it to the SocketId field.
func (o *MemoryPersistentMemoryGoal) SetSocketId(v string) {
	o.SocketId = &v
}

func (o MemoryPersistentMemoryGoal) MarshalJSON() ([]byte, error) {
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
	if o.MemoryModePercentage != nil {
		toSerialize["MemoryModePercentage"] = o.MemoryModePercentage
	}
	if o.PersistentMemoryType != nil {
		toSerialize["PersistentMemoryType"] = o.PersistentMemoryType
	}
	if o.SocketId != nil {
		toSerialize["SocketId"] = o.SocketId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MemoryPersistentMemoryGoal) UnmarshalJSON(bytes []byte) (err error) {
	type MemoryPersistentMemoryGoalWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Volatile memory percentage.
		MemoryModePercentage *int64 `json:"MemoryModePercentage,omitempty"`
		// Type of the Persistent Memory configuration where the Persistent Memory Modules are combined in an interleaved set or not. * `app-direct` - The App Direct interleaved Persistent Memory type. * `app-direct-non-interleaved` - The App Direct non-interleaved Persistent Memory type.
		PersistentMemoryType *string `json:"PersistentMemoryType,omitempty"`
		// CPU Socket ID to which this goal will be applied. * `All Sockets` - All the CPU socket IDs in a server.
		SocketId *string `json:"SocketId,omitempty"`
	}

	varMemoryPersistentMemoryGoalWithoutEmbeddedStruct := MemoryPersistentMemoryGoalWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMemoryPersistentMemoryGoalWithoutEmbeddedStruct)
	if err == nil {
		varMemoryPersistentMemoryGoal := _MemoryPersistentMemoryGoal{}
		varMemoryPersistentMemoryGoal.ClassId = varMemoryPersistentMemoryGoalWithoutEmbeddedStruct.ClassId
		varMemoryPersistentMemoryGoal.ObjectType = varMemoryPersistentMemoryGoalWithoutEmbeddedStruct.ObjectType
		varMemoryPersistentMemoryGoal.MemoryModePercentage = varMemoryPersistentMemoryGoalWithoutEmbeddedStruct.MemoryModePercentage
		varMemoryPersistentMemoryGoal.PersistentMemoryType = varMemoryPersistentMemoryGoalWithoutEmbeddedStruct.PersistentMemoryType
		varMemoryPersistentMemoryGoal.SocketId = varMemoryPersistentMemoryGoalWithoutEmbeddedStruct.SocketId
		*o = MemoryPersistentMemoryGoal(varMemoryPersistentMemoryGoal)
	} else {
		return err
	}

	varMemoryPersistentMemoryGoal := _MemoryPersistentMemoryGoal{}

	err = json.Unmarshal(bytes, &varMemoryPersistentMemoryGoal)
	if err == nil {
		o.MoBaseComplexType = varMemoryPersistentMemoryGoal.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MemoryModePercentage")
		delete(additionalProperties, "PersistentMemoryType")
		delete(additionalProperties, "SocketId")

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

type NullableMemoryPersistentMemoryGoal struct {
	value *MemoryPersistentMemoryGoal
	isSet bool
}

func (v NullableMemoryPersistentMemoryGoal) Get() *MemoryPersistentMemoryGoal {
	return v.value
}

func (v *NullableMemoryPersistentMemoryGoal) Set(val *MemoryPersistentMemoryGoal) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPersistentMemoryGoal) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPersistentMemoryGoal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPersistentMemoryGoal(val *MemoryPersistentMemoryGoal) *NullableMemoryPersistentMemoryGoal {
	return &NullableMemoryPersistentMemoryGoal{value: val, isSet: true}
}

func (v NullableMemoryPersistentMemoryGoal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPersistentMemoryGoal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


