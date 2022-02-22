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

// KubernetesAddonAllOf Definition of the list of properties defined in 'kubernetes.Addon', excluding properties defined in parent classes.
type KubernetesAddonAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	AddonConfiguration NullableKubernetesAddonConfiguration `json:"AddonConfiguration,omitempty"`
	AddonPolicy *MoMoRef `json:"AddonPolicy,omitempty"`
	// Name of addon to be installed on a Kubernetes cluster.
	Name *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAddonAllOf KubernetesAddonAllOf

// NewKubernetesAddonAllOf instantiates a new KubernetesAddonAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAddonAllOf(classId string, objectType string) *KubernetesAddonAllOf {
	this := KubernetesAddonAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesAddonAllOfWithDefaults instantiates a new KubernetesAddonAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAddonAllOfWithDefaults() *KubernetesAddonAllOf {
	this := KubernetesAddonAllOf{}
	var classId string = "kubernetes.Addon"
	this.ClassId = classId
	var objectType string = "kubernetes.Addon"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAddonAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAddonAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAddonAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAddonAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddonConfiguration returns the AddonConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAddonAllOf) GetAddonConfiguration() KubernetesAddonConfiguration {
	if o == nil || o.AddonConfiguration.Get() == nil {
		var ret KubernetesAddonConfiguration
		return ret
	}
	return *o.AddonConfiguration.Get()
}

// GetAddonConfigurationOk returns a tuple with the AddonConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAddonAllOf) GetAddonConfigurationOk() (*KubernetesAddonConfiguration, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AddonConfiguration.Get(), o.AddonConfiguration.IsSet()
}

// HasAddonConfiguration returns a boolean if a field has been set.
func (o *KubernetesAddonAllOf) HasAddonConfiguration() bool {
	if o != nil && o.AddonConfiguration.IsSet() {
		return true
	}

	return false
}

// SetAddonConfiguration gets a reference to the given NullableKubernetesAddonConfiguration and assigns it to the AddonConfiguration field.
func (o *KubernetesAddonAllOf) SetAddonConfiguration(v KubernetesAddonConfiguration) {
	o.AddonConfiguration.Set(&v)
}
// SetAddonConfigurationNil sets the value for AddonConfiguration to be an explicit nil
func (o *KubernetesAddonAllOf) SetAddonConfigurationNil() {
	o.AddonConfiguration.Set(nil)
}

// UnsetAddonConfiguration ensures that no value is present for AddonConfiguration, not even an explicit nil
func (o *KubernetesAddonAllOf) UnsetAddonConfiguration() {
	o.AddonConfiguration.Unset()
}

// GetAddonPolicy returns the AddonPolicy field value if set, zero value otherwise.
func (o *KubernetesAddonAllOf) GetAddonPolicy() MoMoRef {
	if o == nil || o.AddonPolicy == nil {
		var ret MoMoRef
		return ret
	}
	return *o.AddonPolicy
}

// GetAddonPolicyOk returns a tuple with the AddonPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonAllOf) GetAddonPolicyOk() (*MoMoRef, bool) {
	if o == nil || o.AddonPolicy == nil {
		return nil, false
	}
	return o.AddonPolicy, true
}

// HasAddonPolicy returns a boolean if a field has been set.
func (o *KubernetesAddonAllOf) HasAddonPolicy() bool {
	if o != nil && o.AddonPolicy != nil {
		return true
	}

	return false
}

// SetAddonPolicy gets a reference to the given MoMoRef and assigns it to the AddonPolicy field.
func (o *KubernetesAddonAllOf) SetAddonPolicy(v MoMoRef) {
	o.AddonPolicy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesAddonAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesAddonAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesAddonAllOf) SetName(v string) {
	o.Name = &v
}

func (o KubernetesAddonAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AddonConfiguration.IsSet() {
		toSerialize["AddonConfiguration"] = o.AddonConfiguration.Get()
	}
	if o.AddonPolicy != nil {
		toSerialize["AddonPolicy"] = o.AddonPolicy
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesAddonAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesAddonAllOf := _KubernetesAddonAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesAddonAllOf); err == nil {
		*o = KubernetesAddonAllOf(varKubernetesAddonAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AddonConfiguration")
		delete(additionalProperties, "AddonPolicy")
		delete(additionalProperties, "Name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesAddonAllOf struct {
	value *KubernetesAddonAllOf
	isSet bool
}

func (v NullableKubernetesAddonAllOf) Get() *KubernetesAddonAllOf {
	return v.value
}

func (v *NullableKubernetesAddonAllOf) Set(val *KubernetesAddonAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAddonAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAddonAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAddonAllOf(val *KubernetesAddonAllOf) *NullableKubernetesAddonAllOf {
	return &NullableKubernetesAddonAllOf{value: val, isSet: true}
}

func (v NullableKubernetesAddonAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAddonAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


