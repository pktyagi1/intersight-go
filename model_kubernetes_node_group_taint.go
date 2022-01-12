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

// KubernetesNodeGroupTaint Taint to be added to all the nodes in a node group. The taint has key, value and effect. Kubernetes will enforce the effect while scheduling pods to match the key/value pairs.
type KubernetesNodeGroupTaint struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The effect to enforce when the key does not match the value.
	Effect *string `json:"Effect,omitempty"`
	// The key for a Kubernetes taint.
	Key *string `json:"Key,omitempty"`
	// If the key does not match this value, the specified effect is enforced.
	Value *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNodeGroupTaint KubernetesNodeGroupTaint

// NewKubernetesNodeGroupTaint instantiates a new KubernetesNodeGroupTaint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeGroupTaint(classId string, objectType string) *KubernetesNodeGroupTaint {
	this := KubernetesNodeGroupTaint{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesNodeGroupTaintWithDefaults instantiates a new KubernetesNodeGroupTaint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeGroupTaintWithDefaults() *KubernetesNodeGroupTaint {
	this := KubernetesNodeGroupTaint{}
	var classId string = "kubernetes.NodeGroupTaint"
	this.ClassId = classId
	var objectType string = "kubernetes.NodeGroupTaint"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNodeGroupTaint) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupTaint) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNodeGroupTaint) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNodeGroupTaint) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupTaint) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNodeGroupTaint) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEffect returns the Effect field value if set, zero value otherwise.
func (o *KubernetesNodeGroupTaint) GetEffect() string {
	if o == nil || o.Effect == nil {
		var ret string
		return ret
	}
	return *o.Effect
}

// GetEffectOk returns a tuple with the Effect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupTaint) GetEffectOk() (*string, bool) {
	if o == nil || o.Effect == nil {
		return nil, false
	}
	return o.Effect, true
}

// HasEffect returns a boolean if a field has been set.
func (o *KubernetesNodeGroupTaint) HasEffect() bool {
	if o != nil && o.Effect != nil {
		return true
	}

	return false
}

// SetEffect gets a reference to the given string and assigns it to the Effect field.
func (o *KubernetesNodeGroupTaint) SetEffect(v string) {
	o.Effect = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *KubernetesNodeGroupTaint) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupTaint) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *KubernetesNodeGroupTaint) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *KubernetesNodeGroupTaint) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *KubernetesNodeGroupTaint) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupTaint) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *KubernetesNodeGroupTaint) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *KubernetesNodeGroupTaint) SetValue(v string) {
	o.Value = &v
}

func (o KubernetesNodeGroupTaint) MarshalJSON() ([]byte, error) {
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
	if o.Effect != nil {
		toSerialize["Effect"] = o.Effect
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesNodeGroupTaint) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesNodeGroupTaintWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The effect to enforce when the key does not match the value.
		Effect *string `json:"Effect,omitempty"`
		// The key for a Kubernetes taint.
		Key *string `json:"Key,omitempty"`
		// If the key does not match this value, the specified effect is enforced.
		Value *string `json:"Value,omitempty"`
	}

	varKubernetesNodeGroupTaintWithoutEmbeddedStruct := KubernetesNodeGroupTaintWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesNodeGroupTaintWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesNodeGroupTaint := _KubernetesNodeGroupTaint{}
		varKubernetesNodeGroupTaint.ClassId = varKubernetesNodeGroupTaintWithoutEmbeddedStruct.ClassId
		varKubernetesNodeGroupTaint.ObjectType = varKubernetesNodeGroupTaintWithoutEmbeddedStruct.ObjectType
		varKubernetesNodeGroupTaint.Effect = varKubernetesNodeGroupTaintWithoutEmbeddedStruct.Effect
		varKubernetesNodeGroupTaint.Key = varKubernetesNodeGroupTaintWithoutEmbeddedStruct.Key
		varKubernetesNodeGroupTaint.Value = varKubernetesNodeGroupTaintWithoutEmbeddedStruct.Value
		*o = KubernetesNodeGroupTaint(varKubernetesNodeGroupTaint)
	} else {
		return err
	}

	varKubernetesNodeGroupTaint := _KubernetesNodeGroupTaint{}

	err = json.Unmarshal(bytes, &varKubernetesNodeGroupTaint)
	if err == nil {
		o.MoBaseComplexType = varKubernetesNodeGroupTaint.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Effect")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Value")

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

type NullableKubernetesNodeGroupTaint struct {
	value *KubernetesNodeGroupTaint
	isSet bool
}

func (v NullableKubernetesNodeGroupTaint) Get() *KubernetesNodeGroupTaint {
	return v.value
}

func (v *NullableKubernetesNodeGroupTaint) Set(val *KubernetesNodeGroupTaint) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeGroupTaint) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeGroupTaint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeGroupTaint(val *KubernetesNodeGroupTaint) *NullableKubernetesNodeGroupTaint {
	return &NullableKubernetesNodeGroupTaint{value: val, isSet: true}
}

func (v NullableKubernetesNodeGroupTaint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeGroupTaint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


