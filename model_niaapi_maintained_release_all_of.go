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

// NiaapiMaintainedReleaseAllOf Definition of the list of properties defined in 'niaapi.MaintainedRelease', excluding properties defined in parent classes.
type NiaapiMaintainedReleaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Lastest maintained release.
	MaintainedRelease *string `json:"MaintainedRelease,omitempty"`
	// Software release version string.
	SoftwareRelease *string `json:"SoftwareRelease,omitempty"`
	// Long lived version or short lived version.
	VersionTag *string `json:"VersionTag,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiMaintainedReleaseAllOf NiaapiMaintainedReleaseAllOf

// NewNiaapiMaintainedReleaseAllOf instantiates a new NiaapiMaintainedReleaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiMaintainedReleaseAllOf(classId string, objectType string) *NiaapiMaintainedReleaseAllOf {
	this := NiaapiMaintainedReleaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiMaintainedReleaseAllOfWithDefaults instantiates a new NiaapiMaintainedReleaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiMaintainedReleaseAllOfWithDefaults() *NiaapiMaintainedReleaseAllOf {
	this := NiaapiMaintainedReleaseAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiMaintainedReleaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiMaintainedReleaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiMaintainedReleaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiMaintainedReleaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiMaintainedReleaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiMaintainedReleaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMaintainedRelease returns the MaintainedRelease field value if set, zero value otherwise.
func (o *NiaapiMaintainedReleaseAllOf) GetMaintainedRelease() string {
	if o == nil || o.MaintainedRelease == nil {
		var ret string
		return ret
	}
	return *o.MaintainedRelease
}

// GetMaintainedReleaseOk returns a tuple with the MaintainedRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiMaintainedReleaseAllOf) GetMaintainedReleaseOk() (*string, bool) {
	if o == nil || o.MaintainedRelease == nil {
		return nil, false
	}
	return o.MaintainedRelease, true
}

// HasMaintainedRelease returns a boolean if a field has been set.
func (o *NiaapiMaintainedReleaseAllOf) HasMaintainedRelease() bool {
	if o != nil && o.MaintainedRelease != nil {
		return true
	}

	return false
}

// SetMaintainedRelease gets a reference to the given string and assigns it to the MaintainedRelease field.
func (o *NiaapiMaintainedReleaseAllOf) SetMaintainedRelease(v string) {
	o.MaintainedRelease = &v
}

// GetSoftwareRelease returns the SoftwareRelease field value if set, zero value otherwise.
func (o *NiaapiMaintainedReleaseAllOf) GetSoftwareRelease() string {
	if o == nil || o.SoftwareRelease == nil {
		var ret string
		return ret
	}
	return *o.SoftwareRelease
}

// GetSoftwareReleaseOk returns a tuple with the SoftwareRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiMaintainedReleaseAllOf) GetSoftwareReleaseOk() (*string, bool) {
	if o == nil || o.SoftwareRelease == nil {
		return nil, false
	}
	return o.SoftwareRelease, true
}

// HasSoftwareRelease returns a boolean if a field has been set.
func (o *NiaapiMaintainedReleaseAllOf) HasSoftwareRelease() bool {
	if o != nil && o.SoftwareRelease != nil {
		return true
	}

	return false
}

// SetSoftwareRelease gets a reference to the given string and assigns it to the SoftwareRelease field.
func (o *NiaapiMaintainedReleaseAllOf) SetSoftwareRelease(v string) {
	o.SoftwareRelease = &v
}

// GetVersionTag returns the VersionTag field value if set, zero value otherwise.
func (o *NiaapiMaintainedReleaseAllOf) GetVersionTag() string {
	if o == nil || o.VersionTag == nil {
		var ret string
		return ret
	}
	return *o.VersionTag
}

// GetVersionTagOk returns a tuple with the VersionTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiMaintainedReleaseAllOf) GetVersionTagOk() (*string, bool) {
	if o == nil || o.VersionTag == nil {
		return nil, false
	}
	return o.VersionTag, true
}

// HasVersionTag returns a boolean if a field has been set.
func (o *NiaapiMaintainedReleaseAllOf) HasVersionTag() bool {
	if o != nil && o.VersionTag != nil {
		return true
	}

	return false
}

// SetVersionTag gets a reference to the given string and assigns it to the VersionTag field.
func (o *NiaapiMaintainedReleaseAllOf) SetVersionTag(v string) {
	o.VersionTag = &v
}

func (o NiaapiMaintainedReleaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MaintainedRelease != nil {
		toSerialize["MaintainedRelease"] = o.MaintainedRelease
	}
	if o.SoftwareRelease != nil {
		toSerialize["SoftwareRelease"] = o.SoftwareRelease
	}
	if o.VersionTag != nil {
		toSerialize["VersionTag"] = o.VersionTag
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiMaintainedReleaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiaapiMaintainedReleaseAllOf := _NiaapiMaintainedReleaseAllOf{}

	if err = json.Unmarshal(bytes, &varNiaapiMaintainedReleaseAllOf); err == nil {
		*o = NiaapiMaintainedReleaseAllOf(varNiaapiMaintainedReleaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MaintainedRelease")
		delete(additionalProperties, "SoftwareRelease")
		delete(additionalProperties, "VersionTag")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiaapiMaintainedReleaseAllOf struct {
	value *NiaapiMaintainedReleaseAllOf
	isSet bool
}

func (v NullableNiaapiMaintainedReleaseAllOf) Get() *NiaapiMaintainedReleaseAllOf {
	return v.value
}

func (v *NullableNiaapiMaintainedReleaseAllOf) Set(val *NiaapiMaintainedReleaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiMaintainedReleaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiMaintainedReleaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiMaintainedReleaseAllOf(val *NiaapiMaintainedReleaseAllOf) *NullableNiaapiMaintainedReleaseAllOf {
	return &NullableNiaapiMaintainedReleaseAllOf{value: val, isSet: true}
}

func (v NullableNiaapiMaintainedReleaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiMaintainedReleaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


