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

// ApplianceDiagSetting DiagSetting model is used for changing the password of the operating system's diagnostic user account. The diagnostic user account can be used to login to the Intersight Appliance virtual machine. The diagnostic user account is protected by two separate authentication mechanisms: user's password and Cisco CT-engine generated key. Only the Intersight Appliance's local account administrator has the privileges to use this REST API.
type ApplianceDiagSetting struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Status message of the password change operation.
	Message *string `json:"Message,omitempty"`
	// Password of the Intersight Appliance's OS diagnostic user account.
	Password *string `json:"Password,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceDiagSetting ApplianceDiagSetting

// NewApplianceDiagSetting instantiates a new ApplianceDiagSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceDiagSetting(classId string, objectType string) *ApplianceDiagSetting {
	this := ApplianceDiagSetting{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceDiagSettingWithDefaults instantiates a new ApplianceDiagSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceDiagSettingWithDefaults() *ApplianceDiagSetting {
	this := ApplianceDiagSetting{}
	var classId string = "appliance.DiagSetting"
	this.ClassId = classId
	var objectType string = "appliance.DiagSetting"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceDiagSetting) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceDiagSetting) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceDiagSetting) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceDiagSetting) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceDiagSetting) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceDiagSetting) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *ApplianceDiagSetting) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDiagSetting) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *ApplianceDiagSetting) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *ApplianceDiagSetting) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApplianceDiagSetting) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDiagSetting) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApplianceDiagSetting) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApplianceDiagSetting) SetMessage(v string) {
	o.Message = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApplianceDiagSetting) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDiagSetting) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApplianceDiagSetting) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApplianceDiagSetting) SetPassword(v string) {
	o.Password = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceDiagSetting) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDiagSetting) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceDiagSetting) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceDiagSetting) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o ApplianceDiagSetting) MarshalJSON() ([]byte, error) {
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
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceDiagSetting) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceDiagSettingWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// Status message of the password change operation.
		Message *string `json:"Message,omitempty"`
		// Password of the Intersight Appliance's OS diagnostic user account.
		Password *string `json:"Password,omitempty"`
		Account *IamAccountRelationship `json:"Account,omitempty"`
	}

	varApplianceDiagSettingWithoutEmbeddedStruct := ApplianceDiagSettingWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceDiagSettingWithoutEmbeddedStruct)
	if err == nil {
		varApplianceDiagSetting := _ApplianceDiagSetting{}
		varApplianceDiagSetting.ClassId = varApplianceDiagSettingWithoutEmbeddedStruct.ClassId
		varApplianceDiagSetting.ObjectType = varApplianceDiagSettingWithoutEmbeddedStruct.ObjectType
		varApplianceDiagSetting.IsPasswordSet = varApplianceDiagSettingWithoutEmbeddedStruct.IsPasswordSet
		varApplianceDiagSetting.Message = varApplianceDiagSettingWithoutEmbeddedStruct.Message
		varApplianceDiagSetting.Password = varApplianceDiagSettingWithoutEmbeddedStruct.Password
		varApplianceDiagSetting.Account = varApplianceDiagSettingWithoutEmbeddedStruct.Account
		*o = ApplianceDiagSetting(varApplianceDiagSetting)
	} else {
		return err
	}

	varApplianceDiagSetting := _ApplianceDiagSetting{}

	err = json.Unmarshal(bytes, &varApplianceDiagSetting)
	if err == nil {
		o.MoBaseMo = varApplianceDiagSetting.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Password")
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

type NullableApplianceDiagSetting struct {
	value *ApplianceDiagSetting
	isSet bool
}

func (v NullableApplianceDiagSetting) Get() *ApplianceDiagSetting {
	return v.value
}

func (v *NullableApplianceDiagSetting) Set(val *ApplianceDiagSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceDiagSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceDiagSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceDiagSetting(val *ApplianceDiagSetting) *NullableApplianceDiagSetting {
	return &NullableApplianceDiagSetting{value: val, isSet: true}
}

func (v NullableApplianceDiagSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceDiagSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

