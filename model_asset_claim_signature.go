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
	"time"
	"reflect"
	"strings"
)

// AssetClaimSignature A cryptographic signature that can be used to auto-claim a device connector into a user's account.
type AssetClaimSignature struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The result of signing the deviceId appended with the timeStamp fields with the devices private key.
	Signature *string `json:"Signature,omitempty"`
	// The time at which the signature was generated. Date is accurate to Intersights clock. Used to expire the signature.
	TimeStamp *time.Time `json:"TimeStamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetClaimSignature AssetClaimSignature

// NewAssetClaimSignature instantiates a new AssetClaimSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetClaimSignature(classId string, objectType string) *AssetClaimSignature {
	this := AssetClaimSignature{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetClaimSignatureWithDefaults instantiates a new AssetClaimSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetClaimSignatureWithDefaults() *AssetClaimSignature {
	this := AssetClaimSignature{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetClaimSignature) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetClaimSignature) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetClaimSignature) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetClaimSignature) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetClaimSignature) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetClaimSignature) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *AssetClaimSignature) GetSignature() string {
	if o == nil || o.Signature == nil {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetClaimSignature) GetSignatureOk() (*string, bool) {
	if o == nil || o.Signature == nil {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *AssetClaimSignature) HasSignature() bool {
	if o != nil && o.Signature != nil {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *AssetClaimSignature) SetSignature(v string) {
	o.Signature = &v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *AssetClaimSignature) GetTimeStamp() time.Time {
	if o == nil || o.TimeStamp == nil {
		var ret time.Time
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetClaimSignature) GetTimeStampOk() (*time.Time, bool) {
	if o == nil || o.TimeStamp == nil {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *AssetClaimSignature) HasTimeStamp() bool {
	if o != nil && o.TimeStamp != nil {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given time.Time and assigns it to the TimeStamp field.
func (o *AssetClaimSignature) SetTimeStamp(v time.Time) {
	o.TimeStamp = &v
}

func (o AssetClaimSignature) MarshalJSON() ([]byte, error) {
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
	if o.Signature != nil {
		toSerialize["Signature"] = o.Signature
	}
	if o.TimeStamp != nil {
		toSerialize["TimeStamp"] = o.TimeStamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetClaimSignature) UnmarshalJSON(bytes []byte) (err error) {
	type AssetClaimSignatureWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The result of signing the deviceId appended with the timeStamp fields with the devices private key.
		Signature *string `json:"Signature,omitempty"`
		// The time at which the signature was generated. Date is accurate to Intersights clock. Used to expire the signature.
		TimeStamp *time.Time `json:"TimeStamp,omitempty"`
	}

	varAssetClaimSignatureWithoutEmbeddedStruct := AssetClaimSignatureWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetClaimSignatureWithoutEmbeddedStruct)
	if err == nil {
		varAssetClaimSignature := _AssetClaimSignature{}
		varAssetClaimSignature.ClassId = varAssetClaimSignatureWithoutEmbeddedStruct.ClassId
		varAssetClaimSignature.ObjectType = varAssetClaimSignatureWithoutEmbeddedStruct.ObjectType
		varAssetClaimSignature.Signature = varAssetClaimSignatureWithoutEmbeddedStruct.Signature
		varAssetClaimSignature.TimeStamp = varAssetClaimSignatureWithoutEmbeddedStruct.TimeStamp
		*o = AssetClaimSignature(varAssetClaimSignature)
	} else {
		return err
	}

	varAssetClaimSignature := _AssetClaimSignature{}

	err = json.Unmarshal(bytes, &varAssetClaimSignature)
	if err == nil {
		o.MoBaseComplexType = varAssetClaimSignature.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Signature")
		delete(additionalProperties, "TimeStamp")

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

type NullableAssetClaimSignature struct {
	value *AssetClaimSignature
	isSet bool
}

func (v NullableAssetClaimSignature) Get() *AssetClaimSignature {
	return v.value
}

func (v *NullableAssetClaimSignature) Set(val *AssetClaimSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetClaimSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetClaimSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetClaimSignature(val *AssetClaimSignature) *NullableAssetClaimSignature {
	return &NullableAssetClaimSignature{value: val, isSet: true}
}

func (v NullableAssetClaimSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetClaimSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


