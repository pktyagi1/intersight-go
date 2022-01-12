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

// TamAdvisoryCount Total number of advisories currently affecting a given Account.
type TamAdvisoryCount struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Total number of advisories affecting the account.
	AdvisoryCount *int64 `json:"AdvisoryCount,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamAdvisoryCount TamAdvisoryCount

// NewTamAdvisoryCount instantiates a new TamAdvisoryCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamAdvisoryCount(classId string, objectType string) *TamAdvisoryCount {
	this := TamAdvisoryCount{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTamAdvisoryCountWithDefaults instantiates a new TamAdvisoryCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamAdvisoryCountWithDefaults() *TamAdvisoryCount {
	this := TamAdvisoryCount{}
	var classId string = "tam.AdvisoryCount"
	this.ClassId = classId
	var objectType string = "tam.AdvisoryCount"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamAdvisoryCount) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamAdvisoryCount) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamAdvisoryCount) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamAdvisoryCount) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamAdvisoryCount) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamAdvisoryCount) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdvisoryCount returns the AdvisoryCount field value if set, zero value otherwise.
func (o *TamAdvisoryCount) GetAdvisoryCount() int64 {
	if o == nil || o.AdvisoryCount == nil {
		var ret int64
		return ret
	}
	return *o.AdvisoryCount
}

// GetAdvisoryCountOk returns a tuple with the AdvisoryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryCount) GetAdvisoryCountOk() (*int64, bool) {
	if o == nil || o.AdvisoryCount == nil {
		return nil, false
	}
	return o.AdvisoryCount, true
}

// HasAdvisoryCount returns a boolean if a field has been set.
func (o *TamAdvisoryCount) HasAdvisoryCount() bool {
	if o != nil && o.AdvisoryCount != nil {
		return true
	}

	return false
}

// SetAdvisoryCount gets a reference to the given int64 and assigns it to the AdvisoryCount field.
func (o *TamAdvisoryCount) SetAdvisoryCount(v int64) {
	o.AdvisoryCount = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *TamAdvisoryCount) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryCount) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *TamAdvisoryCount) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *TamAdvisoryCount) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o TamAdvisoryCount) MarshalJSON() ([]byte, error) {
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
	if o.AdvisoryCount != nil {
		toSerialize["AdvisoryCount"] = o.AdvisoryCount
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamAdvisoryCount) UnmarshalJSON(bytes []byte) (err error) {
	type TamAdvisoryCountWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Total number of advisories affecting the account.
		AdvisoryCount *int64 `json:"AdvisoryCount,omitempty"`
		Account *IamAccountRelationship `json:"Account,omitempty"`
	}

	varTamAdvisoryCountWithoutEmbeddedStruct := TamAdvisoryCountWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTamAdvisoryCountWithoutEmbeddedStruct)
	if err == nil {
		varTamAdvisoryCount := _TamAdvisoryCount{}
		varTamAdvisoryCount.ClassId = varTamAdvisoryCountWithoutEmbeddedStruct.ClassId
		varTamAdvisoryCount.ObjectType = varTamAdvisoryCountWithoutEmbeddedStruct.ObjectType
		varTamAdvisoryCount.AdvisoryCount = varTamAdvisoryCountWithoutEmbeddedStruct.AdvisoryCount
		varTamAdvisoryCount.Account = varTamAdvisoryCountWithoutEmbeddedStruct.Account
		*o = TamAdvisoryCount(varTamAdvisoryCount)
	} else {
		return err
	}

	varTamAdvisoryCount := _TamAdvisoryCount{}

	err = json.Unmarshal(bytes, &varTamAdvisoryCount)
	if err == nil {
		o.MoBaseMo = varTamAdvisoryCount.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdvisoryCount")
		delete(additionalProperties, "Account")

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

type NullableTamAdvisoryCount struct {
	value *TamAdvisoryCount
	isSet bool
}

func (v NullableTamAdvisoryCount) Get() *TamAdvisoryCount {
	return v.value
}

func (v *NullableTamAdvisoryCount) Set(val *TamAdvisoryCount) {
	v.value = val
	v.isSet = true
}

func (v NullableTamAdvisoryCount) IsSet() bool {
	return v.isSet
}

func (v *NullableTamAdvisoryCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamAdvisoryCount(val *TamAdvisoryCount) *NullableTamAdvisoryCount {
	return &NullableTamAdvisoryCount{value: val, isSet: true}
}

func (v NullableTamAdvisoryCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamAdvisoryCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

