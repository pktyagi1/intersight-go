/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5517
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// KubernetesCniConfigAllOf Definition of the list of properties defined in 'kubernetes.CniConfig', excluding properties defined in parent classes.
type KubernetesCniConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// CNI Image registry location.
	Registry *string `json:"Registry,omitempty"`
	// Container newtork interface plugin configuration.
	Version *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesCniConfigAllOf KubernetesCniConfigAllOf

// NewKubernetesCniConfigAllOf instantiates a new KubernetesCniConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesCniConfigAllOf(classId string, objectType string) *KubernetesCniConfigAllOf {
	this := KubernetesCniConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesCniConfigAllOfWithDefaults instantiates a new KubernetesCniConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesCniConfigAllOfWithDefaults() *KubernetesCniConfigAllOf {
	this := KubernetesCniConfigAllOf{}
	var classId string = "kubernetes.CalicoConfig"
	this.ClassId = classId
	var objectType string = "kubernetes.CalicoConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesCniConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesCniConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesCniConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesCniConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesCniConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesCniConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *KubernetesCniConfigAllOf) GetRegistry() string {
	if o == nil || o.Registry == nil {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCniConfigAllOf) GetRegistryOk() (*string, bool) {
	if o == nil || o.Registry == nil {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *KubernetesCniConfigAllOf) HasRegistry() bool {
	if o != nil && o.Registry != nil {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *KubernetesCniConfigAllOf) SetRegistry(v string) {
	o.Registry = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KubernetesCniConfigAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCniConfigAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KubernetesCniConfigAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *KubernetesCniConfigAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o KubernetesCniConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Registry != nil {
		toSerialize["Registry"] = o.Registry
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesCniConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesCniConfigAllOf := _KubernetesCniConfigAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesCniConfigAllOf); err == nil {
		*o = KubernetesCniConfigAllOf(varKubernetesCniConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Registry")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesCniConfigAllOf struct {
	value *KubernetesCniConfigAllOf
	isSet bool
}

func (v NullableKubernetesCniConfigAllOf) Get() *KubernetesCniConfigAllOf {
	return v.value
}

func (v *NullableKubernetesCniConfigAllOf) Set(val *KubernetesCniConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesCniConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesCniConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesCniConfigAllOf(val *KubernetesCniConfigAllOf) *NullableKubernetesCniConfigAllOf {
	return &NullableKubernetesCniConfigAllOf{value: val, isSet: true}
}

func (v NullableKubernetesCniConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesCniConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


