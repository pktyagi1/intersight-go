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

// KubernetesConfigResult The profile configuration (deploy, validation) results with the overall state and detailed result messages.
type KubernetesConfigResult struct {
	PolicyAbstractConfigResult
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	Profile *KubernetesNodeProfileRelationship `json:"Profile,omitempty"`
	// An array of relationships to kubernetesConfigResultEntry resources.
	ResultEntries []KubernetesConfigResultEntryRelationship `json:"ResultEntries,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesConfigResult KubernetesConfigResult

// NewKubernetesConfigResult instantiates a new KubernetesConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesConfigResult(classId string, objectType string) *KubernetesConfigResult {
	this := KubernetesConfigResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesConfigResultWithDefaults instantiates a new KubernetesConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesConfigResultWithDefaults() *KubernetesConfigResult {
	this := KubernetesConfigResult{}
	var classId string = "kubernetes.ConfigResult"
	this.ClassId = classId
	var objectType string = "kubernetes.ConfigResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesConfigResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesConfigResult) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesConfigResult) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesConfigResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesConfigResult) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesConfigResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *KubernetesConfigResult) GetProfile() KubernetesNodeProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret KubernetesNodeProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesConfigResult) GetProfileOk() (*KubernetesNodeProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *KubernetesConfigResult) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given KubernetesNodeProfileRelationship and assigns it to the Profile field.
func (o *KubernetesConfigResult) SetProfile(v KubernetesNodeProfileRelationship) {
	o.Profile = &v
}

// GetResultEntries returns the ResultEntries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesConfigResult) GetResultEntries() []KubernetesConfigResultEntryRelationship {
	if o == nil  {
		var ret []KubernetesConfigResultEntryRelationship
		return ret
	}
	return o.ResultEntries
}

// GetResultEntriesOk returns a tuple with the ResultEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesConfigResult) GetResultEntriesOk() (*[]KubernetesConfigResultEntryRelationship, bool) {
	if o == nil || o.ResultEntries == nil {
		return nil, false
	}
	return &o.ResultEntries, true
}

// HasResultEntries returns a boolean if a field has been set.
func (o *KubernetesConfigResult) HasResultEntries() bool {
	if o != nil && o.ResultEntries != nil {
		return true
	}

	return false
}

// SetResultEntries gets a reference to the given []KubernetesConfigResultEntryRelationship and assigns it to the ResultEntries field.
func (o *KubernetesConfigResult) SetResultEntries(v []KubernetesConfigResultEntryRelationship) {
	o.ResultEntries = v
}

func (o KubernetesConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigResult, errPolicyAbstractConfigResult := json.Marshal(o.PolicyAbstractConfigResult)
	if errPolicyAbstractConfigResult != nil {
		return []byte{}, errPolicyAbstractConfigResult
	}
	errPolicyAbstractConfigResult = json.Unmarshal([]byte(serializedPolicyAbstractConfigResult), &toSerialize)
	if errPolicyAbstractConfigResult != nil {
		return []byte{}, errPolicyAbstractConfigResult
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}
	if o.ResultEntries != nil {
		toSerialize["ResultEntries"] = o.ResultEntries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesConfigResult) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesConfigResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		Profile *KubernetesNodeProfileRelationship `json:"Profile,omitempty"`
		// An array of relationships to kubernetesConfigResultEntry resources.
		ResultEntries []KubernetesConfigResultEntryRelationship `json:"ResultEntries,omitempty"`
	}

	varKubernetesConfigResultWithoutEmbeddedStruct := KubernetesConfigResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesConfigResultWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesConfigResult := _KubernetesConfigResult{}
		varKubernetesConfigResult.ClassId = varKubernetesConfigResultWithoutEmbeddedStruct.ClassId
		varKubernetesConfigResult.ObjectType = varKubernetesConfigResultWithoutEmbeddedStruct.ObjectType
		varKubernetesConfigResult.Profile = varKubernetesConfigResultWithoutEmbeddedStruct.Profile
		varKubernetesConfigResult.ResultEntries = varKubernetesConfigResultWithoutEmbeddedStruct.ResultEntries
		*o = KubernetesConfigResult(varKubernetesConfigResult)
	} else {
		return err
	}

	varKubernetesConfigResult := _KubernetesConfigResult{}

	err = json.Unmarshal(bytes, &varKubernetesConfigResult)
	if err == nil {
		o.PolicyAbstractConfigResult = varKubernetesConfigResult.PolicyAbstractConfigResult
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Profile")
		delete(additionalProperties, "ResultEntries")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigResult := reflect.ValueOf(o.PolicyAbstractConfigResult)
		for i := 0; i < reflectPolicyAbstractConfigResult.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigResult.Type().Field(i)

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

type NullableKubernetesConfigResult struct {
	value *KubernetesConfigResult
	isSet bool
}

func (v NullableKubernetesConfigResult) Get() *KubernetesConfigResult {
	return v.value
}

func (v *NullableKubernetesConfigResult) Set(val *KubernetesConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesConfigResult(val *KubernetesConfigResult) *NullableKubernetesConfigResult {
	return &NullableKubernetesConfigResult{value: val, isSet: true}
}

func (v NullableKubernetesConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


