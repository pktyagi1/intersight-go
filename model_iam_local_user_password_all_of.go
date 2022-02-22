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
)

// IamLocalUserPasswordAllOf Definition of the list of properties defined in 'iam.LocalUserPassword', excluding properties defined in parent classes.
type IamLocalUserPasswordAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// User-entered passsord to be compared to password for change password function.
	CurrentPassword *string `json:"CurrentPassword,omitempty"`
	// Indicates whether the value of the 'currentPassword' property has been set.
	IsCurrentPasswordSet *bool `json:"IsCurrentPasswordSet,omitempty"`
	// Indicates whether the value of the 'newPassword' property has been set.
	IsNewPasswordSet *bool `json:"IsNewPasswordSet,omitempty"`
	// New password that the user's password should be changed to.
	NewPassword *string `json:"NewPassword,omitempty"`
	// User's current valid passsord.
	Password *string `json:"Password,omitempty"`
	User *IamUserRelationship `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamLocalUserPasswordAllOf IamLocalUserPasswordAllOf

// NewIamLocalUserPasswordAllOf instantiates a new IamLocalUserPasswordAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamLocalUserPasswordAllOf(classId string, objectType string) *IamLocalUserPasswordAllOf {
	this := IamLocalUserPasswordAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamLocalUserPasswordAllOfWithDefaults instantiates a new IamLocalUserPasswordAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamLocalUserPasswordAllOfWithDefaults() *IamLocalUserPasswordAllOf {
	this := IamLocalUserPasswordAllOf{}
	var classId string = "iam.LocalUserPassword"
	this.ClassId = classId
	var objectType string = "iam.LocalUserPassword"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamLocalUserPasswordAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamLocalUserPasswordAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamLocalUserPasswordAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamLocalUserPasswordAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCurrentPassword returns the CurrentPassword field value if set, zero value otherwise.
func (o *IamLocalUserPasswordAllOf) GetCurrentPassword() string {
	if o == nil || o.CurrentPassword == nil {
		var ret string
		return ret
	}
	return *o.CurrentPassword
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordAllOf) GetCurrentPasswordOk() (*string, bool) {
	if o == nil || o.CurrentPassword == nil {
		return nil, false
	}
	return o.CurrentPassword, true
}

// HasCurrentPassword returns a boolean if a field has been set.
func (o *IamLocalUserPasswordAllOf) HasCurrentPassword() bool {
	if o != nil && o.CurrentPassword != nil {
		return true
	}

	return false
}

// SetCurrentPassword gets a reference to the given string and assigns it to the CurrentPassword field.
func (o *IamLocalUserPasswordAllOf) SetCurrentPassword(v string) {
	o.CurrentPassword = &v
}

// GetIsCurrentPasswordSet returns the IsCurrentPasswordSet field value if set, zero value otherwise.
func (o *IamLocalUserPasswordAllOf) GetIsCurrentPasswordSet() bool {
	if o == nil || o.IsCurrentPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsCurrentPasswordSet
}

// GetIsCurrentPasswordSetOk returns a tuple with the IsCurrentPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordAllOf) GetIsCurrentPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsCurrentPasswordSet == nil {
		return nil, false
	}
	return o.IsCurrentPasswordSet, true
}

// HasIsCurrentPasswordSet returns a boolean if a field has been set.
func (o *IamLocalUserPasswordAllOf) HasIsCurrentPasswordSet() bool {
	if o != nil && o.IsCurrentPasswordSet != nil {
		return true
	}

	return false
}

// SetIsCurrentPasswordSet gets a reference to the given bool and assigns it to the IsCurrentPasswordSet field.
func (o *IamLocalUserPasswordAllOf) SetIsCurrentPasswordSet(v bool) {
	o.IsCurrentPasswordSet = &v
}

// GetIsNewPasswordSet returns the IsNewPasswordSet field value if set, zero value otherwise.
func (o *IamLocalUserPasswordAllOf) GetIsNewPasswordSet() bool {
	if o == nil || o.IsNewPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsNewPasswordSet
}

// GetIsNewPasswordSetOk returns a tuple with the IsNewPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordAllOf) GetIsNewPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsNewPasswordSet == nil {
		return nil, false
	}
	return o.IsNewPasswordSet, true
}

// HasIsNewPasswordSet returns a boolean if a field has been set.
func (o *IamLocalUserPasswordAllOf) HasIsNewPasswordSet() bool {
	if o != nil && o.IsNewPasswordSet != nil {
		return true
	}

	return false
}

// SetIsNewPasswordSet gets a reference to the given bool and assigns it to the IsNewPasswordSet field.
func (o *IamLocalUserPasswordAllOf) SetIsNewPasswordSet(v bool) {
	o.IsNewPasswordSet = &v
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise.
func (o *IamLocalUserPasswordAllOf) GetNewPassword() string {
	if o == nil || o.NewPassword == nil {
		var ret string
		return ret
	}
	return *o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordAllOf) GetNewPasswordOk() (*string, bool) {
	if o == nil || o.NewPassword == nil {
		return nil, false
	}
	return o.NewPassword, true
}

// HasNewPassword returns a boolean if a field has been set.
func (o *IamLocalUserPasswordAllOf) HasNewPassword() bool {
	if o != nil && o.NewPassword != nil {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given string and assigns it to the NewPassword field.
func (o *IamLocalUserPasswordAllOf) SetNewPassword(v string) {
	o.NewPassword = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IamLocalUserPasswordAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IamLocalUserPasswordAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IamLocalUserPasswordAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IamLocalUserPasswordAllOf) GetUser() IamUserRelationship {
	if o == nil || o.User == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordAllOf) GetUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IamLocalUserPasswordAllOf) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IamUserRelationship and assigns it to the User field.
func (o *IamLocalUserPasswordAllOf) SetUser(v IamUserRelationship) {
	o.User = &v
}

func (o IamLocalUserPasswordAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CurrentPassword != nil {
		toSerialize["CurrentPassword"] = o.CurrentPassword
	}
	if o.IsCurrentPasswordSet != nil {
		toSerialize["IsCurrentPasswordSet"] = o.IsCurrentPasswordSet
	}
	if o.IsNewPasswordSet != nil {
		toSerialize["IsNewPasswordSet"] = o.IsNewPasswordSet
	}
	if o.NewPassword != nil {
		toSerialize["NewPassword"] = o.NewPassword
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamLocalUserPasswordAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamLocalUserPasswordAllOf := _IamLocalUserPasswordAllOf{}

	if err = json.Unmarshal(bytes, &varIamLocalUserPasswordAllOf); err == nil {
		*o = IamLocalUserPasswordAllOf(varIamLocalUserPasswordAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CurrentPassword")
		delete(additionalProperties, "IsCurrentPasswordSet")
		delete(additionalProperties, "IsNewPasswordSet")
		delete(additionalProperties, "NewPassword")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "User")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamLocalUserPasswordAllOf struct {
	value *IamLocalUserPasswordAllOf
	isSet bool
}

func (v NullableIamLocalUserPasswordAllOf) Get() *IamLocalUserPasswordAllOf {
	return v.value
}

func (v *NullableIamLocalUserPasswordAllOf) Set(val *IamLocalUserPasswordAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamLocalUserPasswordAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamLocalUserPasswordAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamLocalUserPasswordAllOf(val *IamLocalUserPasswordAllOf) *NullableIamLocalUserPasswordAllOf {
	return &NullableIamLocalUserPasswordAllOf{value: val, isSet: true}
}

func (v NullableIamLocalUserPasswordAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamLocalUserPasswordAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


