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

// AssetServiceNowCredential Credentials for authenticating with a ServiceNow target. Management of a ServiceNow target leverages multiple management interfaces and this credential object contains all necessary credentials.
type AssetServiceNowCredential struct {
	AssetCredential
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	OauthAuthenticationCredential *AssetOauthClientIdSecretCredential `json:"OauthAuthenticationCredential,omitempty"`
	UserPasswordCredential *AssetUsernamePasswordCredential `json:"UserPasswordCredential,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetServiceNowCredential AssetServiceNowCredential

// NewAssetServiceNowCredential instantiates a new AssetServiceNowCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetServiceNowCredential(classId string, objectType string) *AssetServiceNowCredential {
	this := AssetServiceNowCredential{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetServiceNowCredentialWithDefaults instantiates a new AssetServiceNowCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetServiceNowCredentialWithDefaults() *AssetServiceNowCredential {
	this := AssetServiceNowCredential{}
	var classId string = "asset.ServiceNowCredential"
	this.ClassId = classId
	var objectType string = "asset.ServiceNowCredential"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetServiceNowCredential) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetServiceNowCredential) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetServiceNowCredential) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetServiceNowCredential) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetServiceNowCredential) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetServiceNowCredential) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOauthAuthenticationCredential returns the OauthAuthenticationCredential field value if set, zero value otherwise.
func (o *AssetServiceNowCredential) GetOauthAuthenticationCredential() AssetOauthClientIdSecretCredential {
	if o == nil || o.OauthAuthenticationCredential == nil {
		var ret AssetOauthClientIdSecretCredential
		return ret
	}
	return *o.OauthAuthenticationCredential
}

// GetOauthAuthenticationCredentialOk returns a tuple with the OauthAuthenticationCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetServiceNowCredential) GetOauthAuthenticationCredentialOk() (*AssetOauthClientIdSecretCredential, bool) {
	if o == nil || o.OauthAuthenticationCredential == nil {
		return nil, false
	}
	return o.OauthAuthenticationCredential, true
}

// HasOauthAuthenticationCredential returns a boolean if a field has been set.
func (o *AssetServiceNowCredential) HasOauthAuthenticationCredential() bool {
	if o != nil && o.OauthAuthenticationCredential != nil {
		return true
	}

	return false
}

// SetOauthAuthenticationCredential gets a reference to the given AssetOauthClientIdSecretCredential and assigns it to the OauthAuthenticationCredential field.
func (o *AssetServiceNowCredential) SetOauthAuthenticationCredential(v AssetOauthClientIdSecretCredential) {
	o.OauthAuthenticationCredential = &v
}

// GetUserPasswordCredential returns the UserPasswordCredential field value if set, zero value otherwise.
func (o *AssetServiceNowCredential) GetUserPasswordCredential() AssetUsernamePasswordCredential {
	if o == nil || o.UserPasswordCredential == nil {
		var ret AssetUsernamePasswordCredential
		return ret
	}
	return *o.UserPasswordCredential
}

// GetUserPasswordCredentialOk returns a tuple with the UserPasswordCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetServiceNowCredential) GetUserPasswordCredentialOk() (*AssetUsernamePasswordCredential, bool) {
	if o == nil || o.UserPasswordCredential == nil {
		return nil, false
	}
	return o.UserPasswordCredential, true
}

// HasUserPasswordCredential returns a boolean if a field has been set.
func (o *AssetServiceNowCredential) HasUserPasswordCredential() bool {
	if o != nil && o.UserPasswordCredential != nil {
		return true
	}

	return false
}

// SetUserPasswordCredential gets a reference to the given AssetUsernamePasswordCredential and assigns it to the UserPasswordCredential field.
func (o *AssetServiceNowCredential) SetUserPasswordCredential(v AssetUsernamePasswordCredential) {
	o.UserPasswordCredential = &v
}

func (o AssetServiceNowCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetCredential, errAssetCredential := json.Marshal(o.AssetCredential)
	if errAssetCredential != nil {
		return []byte{}, errAssetCredential
	}
	errAssetCredential = json.Unmarshal([]byte(serializedAssetCredential), &toSerialize)
	if errAssetCredential != nil {
		return []byte{}, errAssetCredential
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OauthAuthenticationCredential != nil {
		toSerialize["OauthAuthenticationCredential"] = o.OauthAuthenticationCredential
	}
	if o.UserPasswordCredential != nil {
		toSerialize["UserPasswordCredential"] = o.UserPasswordCredential
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetServiceNowCredential) UnmarshalJSON(bytes []byte) (err error) {
	type AssetServiceNowCredentialWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		OauthAuthenticationCredential *AssetOauthClientIdSecretCredential `json:"OauthAuthenticationCredential,omitempty"`
		UserPasswordCredential *AssetUsernamePasswordCredential `json:"UserPasswordCredential,omitempty"`
	}

	varAssetServiceNowCredentialWithoutEmbeddedStruct := AssetServiceNowCredentialWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetServiceNowCredentialWithoutEmbeddedStruct)
	if err == nil {
		varAssetServiceNowCredential := _AssetServiceNowCredential{}
		varAssetServiceNowCredential.ClassId = varAssetServiceNowCredentialWithoutEmbeddedStruct.ClassId
		varAssetServiceNowCredential.ObjectType = varAssetServiceNowCredentialWithoutEmbeddedStruct.ObjectType
		varAssetServiceNowCredential.OauthAuthenticationCredential = varAssetServiceNowCredentialWithoutEmbeddedStruct.OauthAuthenticationCredential
		varAssetServiceNowCredential.UserPasswordCredential = varAssetServiceNowCredentialWithoutEmbeddedStruct.UserPasswordCredential
		*o = AssetServiceNowCredential(varAssetServiceNowCredential)
	} else {
		return err
	}

	varAssetServiceNowCredential := _AssetServiceNowCredential{}

	err = json.Unmarshal(bytes, &varAssetServiceNowCredential)
	if err == nil {
		o.AssetCredential = varAssetServiceNowCredential.AssetCredential
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OauthAuthenticationCredential")
		delete(additionalProperties, "UserPasswordCredential")

		// remove fields from embedded structs
		reflectAssetCredential := reflect.ValueOf(o.AssetCredential)
		for i := 0; i < reflectAssetCredential.Type().NumField(); i++ {
			t := reflectAssetCredential.Type().Field(i)

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

type NullableAssetServiceNowCredential struct {
	value *AssetServiceNowCredential
	isSet bool
}

func (v NullableAssetServiceNowCredential) Get() *AssetServiceNowCredential {
	return v.value
}

func (v *NullableAssetServiceNowCredential) Set(val *AssetServiceNowCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetServiceNowCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetServiceNowCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetServiceNowCredential(val *AssetServiceNowCredential) *NullableAssetServiceNowCredential {
	return &NullableAssetServiceNowCredential{value: val, isSet: true}
}

func (v NullableAssetServiceNowCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetServiceNowCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


