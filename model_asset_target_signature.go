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
	"reflect"
	"strings"
)

// AssetTargetSignature A signature generated by a Intersight on creation of a target.
type AssetTargetSignature struct {
	AssetClaimSignature
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The moid of the associated target.
	TargetId *string `json:"TargetId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetTargetSignature AssetTargetSignature

// NewAssetTargetSignature instantiates a new AssetTargetSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetTargetSignature(classId string, objectType string) *AssetTargetSignature {
	this := AssetTargetSignature{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetTargetSignatureWithDefaults instantiates a new AssetTargetSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTargetSignatureWithDefaults() *AssetTargetSignature {
	this := AssetTargetSignature{}
	var classId string = "asset.TargetSignature"
	this.ClassId = classId
	var objectType string = "asset.TargetSignature"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetTargetSignature) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetTargetSignature) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetTargetSignature) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetTargetSignature) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetTargetSignature) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetTargetSignature) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *AssetTargetSignature) GetTargetId() string {
	if o == nil || o.TargetId == nil {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetSignature) GetTargetIdOk() (*string, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *AssetTargetSignature) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *AssetTargetSignature) SetTargetId(v string) {
	o.TargetId = &v
}

func (o AssetTargetSignature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetClaimSignature, errAssetClaimSignature := json.Marshal(o.AssetClaimSignature)
	if errAssetClaimSignature != nil {
		return []byte{}, errAssetClaimSignature
	}
	errAssetClaimSignature = json.Unmarshal([]byte(serializedAssetClaimSignature), &toSerialize)
	if errAssetClaimSignature != nil {
		return []byte{}, errAssetClaimSignature
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TargetId != nil {
		toSerialize["TargetId"] = o.TargetId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetTargetSignature) UnmarshalJSON(bytes []byte) (err error) {
	type AssetTargetSignatureWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The moid of the associated target.
		TargetId *string `json:"TargetId,omitempty"`
	}

	varAssetTargetSignatureWithoutEmbeddedStruct := AssetTargetSignatureWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetTargetSignatureWithoutEmbeddedStruct)
	if err == nil {
		varAssetTargetSignature := _AssetTargetSignature{}
		varAssetTargetSignature.ClassId = varAssetTargetSignatureWithoutEmbeddedStruct.ClassId
		varAssetTargetSignature.ObjectType = varAssetTargetSignatureWithoutEmbeddedStruct.ObjectType
		varAssetTargetSignature.TargetId = varAssetTargetSignatureWithoutEmbeddedStruct.TargetId
		*o = AssetTargetSignature(varAssetTargetSignature)
	} else {
		return err
	}

	varAssetTargetSignature := _AssetTargetSignature{}

	err = json.Unmarshal(bytes, &varAssetTargetSignature)
	if err == nil {
		o.AssetClaimSignature = varAssetTargetSignature.AssetClaimSignature
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetId")

		// remove fields from embedded structs
		reflectAssetClaimSignature := reflect.ValueOf(o.AssetClaimSignature)
		for i := 0; i < reflectAssetClaimSignature.Type().NumField(); i++ {
			t := reflectAssetClaimSignature.Type().Field(i)

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

type NullableAssetTargetSignature struct {
	value *AssetTargetSignature
	isSet bool
}

func (v NullableAssetTargetSignature) Get() *AssetTargetSignature {
	return v.value
}

func (v *NullableAssetTargetSignature) Set(val *AssetTargetSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTargetSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTargetSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTargetSignature(val *AssetTargetSignature) *NullableAssetTargetSignature {
	return &NullableAssetTargetSignature{value: val, isSet: true}
}

func (v NullableAssetTargetSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTargetSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


