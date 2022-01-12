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

// KubernetesStatefulSet StatefulSet is the workload API object used to manage stateful applications. It manages the deployment and scaling of a set of Pods, and provides guarantees about the ordering and uniqueness of these Pods.
type KubernetesStatefulSet struct {
	KubernetesAbstractStatefulSet
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	Status NullableKubernetesStatefulSetStatus `json:"Status,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesStatefulSet KubernetesStatefulSet

// NewKubernetesStatefulSet instantiates a new KubernetesStatefulSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesStatefulSet(classId string, objectType string) *KubernetesStatefulSet {
	this := KubernetesStatefulSet{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesStatefulSetWithDefaults instantiates a new KubernetesStatefulSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesStatefulSetWithDefaults() *KubernetesStatefulSet {
	this := KubernetesStatefulSet{}
	var classId string = "kubernetes.StatefulSet"
	this.ClassId = classId
	var objectType string = "kubernetes.StatefulSet"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesStatefulSet) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesStatefulSet) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesStatefulSet) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesStatefulSet) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesStatefulSet) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesStatefulSet) SetObjectType(v string) {
	o.ObjectType = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesStatefulSet) GetStatus() KubernetesStatefulSetStatus {
	if o == nil || o.Status.Get() == nil {
		var ret KubernetesStatefulSetStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesStatefulSet) GetStatusOk() (*KubernetesStatefulSetStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *KubernetesStatefulSet) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableKubernetesStatefulSetStatus and assigns it to the Status field.
func (o *KubernetesStatefulSet) SetStatus(v KubernetesStatefulSetStatus) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *KubernetesStatefulSet) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *KubernetesStatefulSet) UnsetStatus() {
	o.Status.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *KubernetesStatefulSet) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesStatefulSet) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *KubernetesStatefulSet) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *KubernetesStatefulSet) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o KubernetesStatefulSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesAbstractStatefulSet, errKubernetesAbstractStatefulSet := json.Marshal(o.KubernetesAbstractStatefulSet)
	if errKubernetesAbstractStatefulSet != nil {
		return []byte{}, errKubernetesAbstractStatefulSet
	}
	errKubernetesAbstractStatefulSet = json.Unmarshal([]byte(serializedKubernetesAbstractStatefulSet), &toSerialize)
	if errKubernetesAbstractStatefulSet != nil {
		return []byte{}, errKubernetesAbstractStatefulSet
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Status.IsSet() {
		toSerialize["Status"] = o.Status.Get()
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesStatefulSet) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesStatefulSetWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		Status NullableKubernetesStatefulSetStatus `json:"Status,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varKubernetesStatefulSetWithoutEmbeddedStruct := KubernetesStatefulSetWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesStatefulSetWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesStatefulSet := _KubernetesStatefulSet{}
		varKubernetesStatefulSet.ClassId = varKubernetesStatefulSetWithoutEmbeddedStruct.ClassId
		varKubernetesStatefulSet.ObjectType = varKubernetesStatefulSetWithoutEmbeddedStruct.ObjectType
		varKubernetesStatefulSet.Status = varKubernetesStatefulSetWithoutEmbeddedStruct.Status
		varKubernetesStatefulSet.RegisteredDevice = varKubernetesStatefulSetWithoutEmbeddedStruct.RegisteredDevice
		*o = KubernetesStatefulSet(varKubernetesStatefulSet)
	} else {
		return err
	}

	varKubernetesStatefulSet := _KubernetesStatefulSet{}

	err = json.Unmarshal(bytes, &varKubernetesStatefulSet)
	if err == nil {
		o.KubernetesAbstractStatefulSet = varKubernetesStatefulSet.KubernetesAbstractStatefulSet
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectKubernetesAbstractStatefulSet := reflect.ValueOf(o.KubernetesAbstractStatefulSet)
		for i := 0; i < reflectKubernetesAbstractStatefulSet.Type().NumField(); i++ {
			t := reflectKubernetesAbstractStatefulSet.Type().Field(i)

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

type NullableKubernetesStatefulSet struct {
	value *KubernetesStatefulSet
	isSet bool
}

func (v NullableKubernetesStatefulSet) Get() *KubernetesStatefulSet {
	return v.value
}

func (v *NullableKubernetesStatefulSet) Set(val *KubernetesStatefulSet) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesStatefulSet) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesStatefulSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesStatefulSet(val *KubernetesStatefulSet) *NullableKubernetesStatefulSet {
	return &NullableKubernetesStatefulSet{value: val, isSet: true}
}

func (v NullableKubernetesStatefulSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesStatefulSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


