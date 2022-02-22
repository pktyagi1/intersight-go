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
	"reflect"
	"strings"
)

// HyperflexServerFirmwareVersionEntry An entry specifying supported server firmware version in regex format.
type HyperflexServerFirmwareVersionEntry struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	Constraint NullableHyperflexAppSettingConstraint `json:"Constraint,omitempty"`
	// The server platform type that is applicable for the server firmware bundle version. * `M5` - M5 generation of UCS server. * `M3` - M3 generation of UCS server. * `M4` - M4 generation of UCS server. * `M6` - M6 generation of UCS server.
	ServerPlatform *string `json:"ServerPlatform,omitempty"`
	// The server firmware bundle version.
	Version *string `json:"Version,omitempty"`
	ServerFirmwareVersion *HyperflexServerFirmwareVersionRelationship `json:"ServerFirmwareVersion,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexServerFirmwareVersionEntry HyperflexServerFirmwareVersionEntry

// NewHyperflexServerFirmwareVersionEntry instantiates a new HyperflexServerFirmwareVersionEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexServerFirmwareVersionEntry(classId string, objectType string) *HyperflexServerFirmwareVersionEntry {
	this := HyperflexServerFirmwareVersionEntry{}
	this.ClassId = classId
	this.ObjectType = objectType
	var serverPlatform string = "M5"
	this.ServerPlatform = &serverPlatform
	return &this
}

// NewHyperflexServerFirmwareVersionEntryWithDefaults instantiates a new HyperflexServerFirmwareVersionEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexServerFirmwareVersionEntryWithDefaults() *HyperflexServerFirmwareVersionEntry {
	this := HyperflexServerFirmwareVersionEntry{}
	var classId string = "hyperflex.ServerFirmwareVersionEntry"
	this.ClassId = classId
	var objectType string = "hyperflex.ServerFirmwareVersionEntry"
	this.ObjectType = objectType
	var serverPlatform string = "M5"
	this.ServerPlatform = &serverPlatform
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexServerFirmwareVersionEntry) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionEntry) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexServerFirmwareVersionEntry) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexServerFirmwareVersionEntry) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionEntry) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexServerFirmwareVersionEntry) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConstraint returns the Constraint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexServerFirmwareVersionEntry) GetConstraint() HyperflexAppSettingConstraint {
	if o == nil || o.Constraint.Get() == nil {
		var ret HyperflexAppSettingConstraint
		return ret
	}
	return *o.Constraint.Get()
}

// GetConstraintOk returns a tuple with the Constraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexServerFirmwareVersionEntry) GetConstraintOk() (*HyperflexAppSettingConstraint, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Constraint.Get(), o.Constraint.IsSet()
}

// HasConstraint returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersionEntry) HasConstraint() bool {
	if o != nil && o.Constraint.IsSet() {
		return true
	}

	return false
}

// SetConstraint gets a reference to the given NullableHyperflexAppSettingConstraint and assigns it to the Constraint field.
func (o *HyperflexServerFirmwareVersionEntry) SetConstraint(v HyperflexAppSettingConstraint) {
	o.Constraint.Set(&v)
}
// SetConstraintNil sets the value for Constraint to be an explicit nil
func (o *HyperflexServerFirmwareVersionEntry) SetConstraintNil() {
	o.Constraint.Set(nil)
}

// UnsetConstraint ensures that no value is present for Constraint, not even an explicit nil
func (o *HyperflexServerFirmwareVersionEntry) UnsetConstraint() {
	o.Constraint.Unset()
}

// GetServerPlatform returns the ServerPlatform field value if set, zero value otherwise.
func (o *HyperflexServerFirmwareVersionEntry) GetServerPlatform() string {
	if o == nil || o.ServerPlatform == nil {
		var ret string
		return ret
	}
	return *o.ServerPlatform
}

// GetServerPlatformOk returns a tuple with the ServerPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionEntry) GetServerPlatformOk() (*string, bool) {
	if o == nil || o.ServerPlatform == nil {
		return nil, false
	}
	return o.ServerPlatform, true
}

// HasServerPlatform returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersionEntry) HasServerPlatform() bool {
	if o != nil && o.ServerPlatform != nil {
		return true
	}

	return false
}

// SetServerPlatform gets a reference to the given string and assigns it to the ServerPlatform field.
func (o *HyperflexServerFirmwareVersionEntry) SetServerPlatform(v string) {
	o.ServerPlatform = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexServerFirmwareVersionEntry) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionEntry) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersionEntry) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexServerFirmwareVersionEntry) SetVersion(v string) {
	o.Version = &v
}

// GetServerFirmwareVersion returns the ServerFirmwareVersion field value if set, zero value otherwise.
func (o *HyperflexServerFirmwareVersionEntry) GetServerFirmwareVersion() HyperflexServerFirmwareVersionRelationship {
	if o == nil || o.ServerFirmwareVersion == nil {
		var ret HyperflexServerFirmwareVersionRelationship
		return ret
	}
	return *o.ServerFirmwareVersion
}

// GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionEntry) GetServerFirmwareVersionOk() (*HyperflexServerFirmwareVersionRelationship, bool) {
	if o == nil || o.ServerFirmwareVersion == nil {
		return nil, false
	}
	return o.ServerFirmwareVersion, true
}

// HasServerFirmwareVersion returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersionEntry) HasServerFirmwareVersion() bool {
	if o != nil && o.ServerFirmwareVersion != nil {
		return true
	}

	return false
}

// SetServerFirmwareVersion gets a reference to the given HyperflexServerFirmwareVersionRelationship and assigns it to the ServerFirmwareVersion field.
func (o *HyperflexServerFirmwareVersionEntry) SetServerFirmwareVersion(v HyperflexServerFirmwareVersionRelationship) {
	o.ServerFirmwareVersion = &v
}

func (o HyperflexServerFirmwareVersionEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Constraint.IsSet() {
		toSerialize["Constraint"] = o.Constraint.Get()
	}
	if o.ServerPlatform != nil {
		toSerialize["ServerPlatform"] = o.ServerPlatform
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.ServerFirmwareVersion != nil {
		toSerialize["ServerFirmwareVersion"] = o.ServerFirmwareVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexServerFirmwareVersionEntry) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexServerFirmwareVersionEntryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		Constraint NullableHyperflexAppSettingConstraint `json:"Constraint,omitempty"`
		// The server platform type that is applicable for the server firmware bundle version. * `M5` - M5 generation of UCS server. * `M3` - M3 generation of UCS server. * `M4` - M4 generation of UCS server. * `M6` - M6 generation of UCS server.
		ServerPlatform *string `json:"ServerPlatform,omitempty"`
		// The server firmware bundle version.
		Version *string `json:"Version,omitempty"`
		ServerFirmwareVersion *HyperflexServerFirmwareVersionRelationship `json:"ServerFirmwareVersion,omitempty"`
	}

	varHyperflexServerFirmwareVersionEntryWithoutEmbeddedStruct := HyperflexServerFirmwareVersionEntryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexServerFirmwareVersionEntryWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexServerFirmwareVersionEntry := _HyperflexServerFirmwareVersionEntry{}
		varHyperflexServerFirmwareVersionEntry.ClassId = varHyperflexServerFirmwareVersionEntryWithoutEmbeddedStruct.ClassId
		varHyperflexServerFirmwareVersionEntry.ObjectType = varHyperflexServerFirmwareVersionEntryWithoutEmbeddedStruct.ObjectType
		varHyperflexServerFirmwareVersionEntry.Constraint = varHyperflexServerFirmwareVersionEntryWithoutEmbeddedStruct.Constraint
		varHyperflexServerFirmwareVersionEntry.ServerPlatform = varHyperflexServerFirmwareVersionEntryWithoutEmbeddedStruct.ServerPlatform
		varHyperflexServerFirmwareVersionEntry.Version = varHyperflexServerFirmwareVersionEntryWithoutEmbeddedStruct.Version
		varHyperflexServerFirmwareVersionEntry.ServerFirmwareVersion = varHyperflexServerFirmwareVersionEntryWithoutEmbeddedStruct.ServerFirmwareVersion
		*o = HyperflexServerFirmwareVersionEntry(varHyperflexServerFirmwareVersionEntry)
	} else {
		return err
	}

	varHyperflexServerFirmwareVersionEntry := _HyperflexServerFirmwareVersionEntry{}

	err = json.Unmarshal(bytes, &varHyperflexServerFirmwareVersionEntry)
	if err == nil {
		o.MoBaseMo = varHyperflexServerFirmwareVersionEntry.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Constraint")
		delete(additionalProperties, "ServerPlatform")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "ServerFirmwareVersion")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableHyperflexServerFirmwareVersionEntry struct {
	value *HyperflexServerFirmwareVersionEntry
	isSet bool
}

func (v NullableHyperflexServerFirmwareVersionEntry) Get() *HyperflexServerFirmwareVersionEntry {
	return v.value
}

func (v *NullableHyperflexServerFirmwareVersionEntry) Set(val *HyperflexServerFirmwareVersionEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexServerFirmwareVersionEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexServerFirmwareVersionEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexServerFirmwareVersionEntry(val *HyperflexServerFirmwareVersionEntry) *NullableHyperflexServerFirmwareVersionEntry {
	return &NullableHyperflexServerFirmwareVersionEntry{value: val, isSet: true}
}

func (v NullableHyperflexServerFirmwareVersionEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexServerFirmwareVersionEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


