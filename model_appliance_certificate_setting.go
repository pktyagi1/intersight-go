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
	"reflect"
	"strings"
)

// ApplianceCertificateSetting Certificate the appliance uses for browser traffic.
type ApplianceCertificateSetting struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	Certificate *IamCertificateRelationship `json:"Certificate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceCertificateSetting ApplianceCertificateSetting

// NewApplianceCertificateSetting instantiates a new ApplianceCertificateSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceCertificateSetting(classId string, objectType string) *ApplianceCertificateSetting {
	this := ApplianceCertificateSetting{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceCertificateSettingWithDefaults instantiates a new ApplianceCertificateSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceCertificateSettingWithDefaults() *ApplianceCertificateSetting {
	this := ApplianceCertificateSetting{}
	var classId string = "appliance.CertificateSetting"
	this.ClassId = classId
	var objectType string = "appliance.CertificateSetting"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceCertificateSetting) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceCertificateSetting) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceCertificateSetting) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceCertificateSetting) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceCertificateSetting) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceCertificateSetting) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceCertificateSetting) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceCertificateSetting) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceCertificateSetting) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceCertificateSetting) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *ApplianceCertificateSetting) GetCertificate() IamCertificateRelationship {
	if o == nil || o.Certificate == nil {
		var ret IamCertificateRelationship
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceCertificateSetting) GetCertificateOk() (*IamCertificateRelationship, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *ApplianceCertificateSetting) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given IamCertificateRelationship and assigns it to the Certificate field.
func (o *ApplianceCertificateSetting) SetCertificate(v IamCertificateRelationship) {
	o.Certificate = &v
}

func (o ApplianceCertificateSetting) MarshalJSON() ([]byte, error) {
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
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Certificate != nil {
		toSerialize["Certificate"] = o.Certificate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceCertificateSetting) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceCertificateSettingWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		Account *IamAccountRelationship `json:"Account,omitempty"`
		Certificate *IamCertificateRelationship `json:"Certificate,omitempty"`
	}

	varApplianceCertificateSettingWithoutEmbeddedStruct := ApplianceCertificateSettingWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceCertificateSettingWithoutEmbeddedStruct)
	if err == nil {
		varApplianceCertificateSetting := _ApplianceCertificateSetting{}
		varApplianceCertificateSetting.ClassId = varApplianceCertificateSettingWithoutEmbeddedStruct.ClassId
		varApplianceCertificateSetting.ObjectType = varApplianceCertificateSettingWithoutEmbeddedStruct.ObjectType
		varApplianceCertificateSetting.Account = varApplianceCertificateSettingWithoutEmbeddedStruct.Account
		varApplianceCertificateSetting.Certificate = varApplianceCertificateSettingWithoutEmbeddedStruct.Certificate
		*o = ApplianceCertificateSetting(varApplianceCertificateSetting)
	} else {
		return err
	}

	varApplianceCertificateSetting := _ApplianceCertificateSetting{}

	err = json.Unmarshal(bytes, &varApplianceCertificateSetting)
	if err == nil {
		o.MoBaseMo = varApplianceCertificateSetting.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Certificate")

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

type NullableApplianceCertificateSetting struct {
	value *ApplianceCertificateSetting
	isSet bool
}

func (v NullableApplianceCertificateSetting) Get() *ApplianceCertificateSetting {
	return v.value
}

func (v *NullableApplianceCertificateSetting) Set(val *ApplianceCertificateSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceCertificateSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceCertificateSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceCertificateSetting(val *ApplianceCertificateSetting) *NullableApplianceCertificateSetting {
	return &NullableApplianceCertificateSetting{value: val, isSet: true}
}

func (v NullableApplianceCertificateSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceCertificateSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


