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

// KubernetesNodeAddressAllOf Definition of the list of properties defined in 'kubernetes.NodeAddress', excluding properties defined in parent classes.
type KubernetesNodeAddressAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The address of type specified by the type field.
	Address *string `json:"Address,omitempty"`
	// The address type of the Node. The usage of the IP address depending on the cloud provider or bare metal configuration. HostName - The hostname as reported by the node's kernel. ExternalIP - Typically the IP address of the node that is externally routable (available from outside the cluster) InternalIP - Typically the IP address of the node that is routable only within the cluster.
	Type *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNodeAddressAllOf KubernetesNodeAddressAllOf

// NewKubernetesNodeAddressAllOf instantiates a new KubernetesNodeAddressAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeAddressAllOf(classId string, objectType string) *KubernetesNodeAddressAllOf {
	this := KubernetesNodeAddressAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesNodeAddressAllOfWithDefaults instantiates a new KubernetesNodeAddressAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeAddressAllOfWithDefaults() *KubernetesNodeAddressAllOf {
	this := KubernetesNodeAddressAllOf{}
	var classId string = "kubernetes.NodeAddress"
	this.ClassId = classId
	var objectType string = "kubernetes.NodeAddress"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNodeAddressAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeAddressAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNodeAddressAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNodeAddressAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeAddressAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNodeAddressAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *KubernetesNodeAddressAllOf) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeAddressAllOf) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *KubernetesNodeAddressAllOf) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *KubernetesNodeAddressAllOf) SetAddress(v string) {
	o.Address = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KubernetesNodeAddressAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeAddressAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KubernetesNodeAddressAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KubernetesNodeAddressAllOf) SetType(v string) {
	o.Type = &v
}

func (o KubernetesNodeAddressAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Address != nil {
		toSerialize["Address"] = o.Address
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesNodeAddressAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesNodeAddressAllOf := _KubernetesNodeAddressAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesNodeAddressAllOf); err == nil {
		*o = KubernetesNodeAddressAllOf(varKubernetesNodeAddressAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Address")
		delete(additionalProperties, "Type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesNodeAddressAllOf struct {
	value *KubernetesNodeAddressAllOf
	isSet bool
}

func (v NullableKubernetesNodeAddressAllOf) Get() *KubernetesNodeAddressAllOf {
	return v.value
}

func (v *NullableKubernetesNodeAddressAllOf) Set(val *KubernetesNodeAddressAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeAddressAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeAddressAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeAddressAllOf(val *KubernetesNodeAddressAllOf) *NullableKubernetesNodeAddressAllOf {
	return &NullableKubernetesNodeAddressAllOf{value: val, isSet: true}
}

func (v NullableKubernetesNodeAddressAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeAddressAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


