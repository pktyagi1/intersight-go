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
)

// IamIdpReferenceAllOf Definition of the list of properties defined in 'iam.IdpReference', excluding properties defined in parent classes.
type IamIdpReferenceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The email domain name for this IdP of the user. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication.
	DomainName *string `json:"DomainName,omitempty"`
	// Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP/Service Provider.
	IdpEntityId *string `json:"IdpEntityId,omitempty"`
	// The flag represents if the second factor of authentication is required for Cisco IdP users.
	MultiFactorAuthentication *bool `json:"MultiFactorAuthentication,omitempty"`
	// Cisco IdP reference in an account.
	Name *string `json:"Name,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	Idp *IamIdpRelationship `json:"Idp,omitempty"`
	// An array of relationships to iamUserPreference resources.
	UserPreferences []IamUserPreferenceRelationship `json:"UserPreferences,omitempty"`
	// An array of relationships to iamUserGroup resources.
	Usergroups []IamUserGroupRelationship `json:"Usergroups,omitempty"`
	// An array of relationships to iamUser resources.
	Users []IamUserRelationship `json:"Users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamIdpReferenceAllOf IamIdpReferenceAllOf

// NewIamIdpReferenceAllOf instantiates a new IamIdpReferenceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamIdpReferenceAllOf(classId string, objectType string) *IamIdpReferenceAllOf {
	this := IamIdpReferenceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var multiFactorAuthentication bool = false
	this.MultiFactorAuthentication = &multiFactorAuthentication
	return &this
}

// NewIamIdpReferenceAllOfWithDefaults instantiates a new IamIdpReferenceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamIdpReferenceAllOfWithDefaults() *IamIdpReferenceAllOf {
	this := IamIdpReferenceAllOf{}
	var classId string = "iam.IdpReference"
	this.ClassId = classId
	var objectType string = "iam.IdpReference"
	this.ObjectType = objectType
	var multiFactorAuthentication bool = false
	this.MultiFactorAuthentication = &multiFactorAuthentication
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamIdpReferenceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamIdpReferenceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamIdpReferenceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamIdpReferenceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamIdpReferenceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamIdpReferenceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *IamIdpReferenceAllOf) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdpReferenceAllOf) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *IamIdpReferenceAllOf) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *IamIdpReferenceAllOf) SetDomainName(v string) {
	o.DomainName = &v
}

// GetIdpEntityId returns the IdpEntityId field value if set, zero value otherwise.
func (o *IamIdpReferenceAllOf) GetIdpEntityId() string {
	if o == nil || o.IdpEntityId == nil {
		var ret string
		return ret
	}
	return *o.IdpEntityId
}

// GetIdpEntityIdOk returns a tuple with the IdpEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdpReferenceAllOf) GetIdpEntityIdOk() (*string, bool) {
	if o == nil || o.IdpEntityId == nil {
		return nil, false
	}
	return o.IdpEntityId, true
}

// HasIdpEntityId returns a boolean if a field has been set.
func (o *IamIdpReferenceAllOf) HasIdpEntityId() bool {
	if o != nil && o.IdpEntityId != nil {
		return true
	}

	return false
}

// SetIdpEntityId gets a reference to the given string and assigns it to the IdpEntityId field.
func (o *IamIdpReferenceAllOf) SetIdpEntityId(v string) {
	o.IdpEntityId = &v
}

// GetMultiFactorAuthentication returns the MultiFactorAuthentication field value if set, zero value otherwise.
func (o *IamIdpReferenceAllOf) GetMultiFactorAuthentication() bool {
	if o == nil || o.MultiFactorAuthentication == nil {
		var ret bool
		return ret
	}
	return *o.MultiFactorAuthentication
}

// GetMultiFactorAuthenticationOk returns a tuple with the MultiFactorAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdpReferenceAllOf) GetMultiFactorAuthenticationOk() (*bool, bool) {
	if o == nil || o.MultiFactorAuthentication == nil {
		return nil, false
	}
	return o.MultiFactorAuthentication, true
}

// HasMultiFactorAuthentication returns a boolean if a field has been set.
func (o *IamIdpReferenceAllOf) HasMultiFactorAuthentication() bool {
	if o != nil && o.MultiFactorAuthentication != nil {
		return true
	}

	return false
}

// SetMultiFactorAuthentication gets a reference to the given bool and assigns it to the MultiFactorAuthentication field.
func (o *IamIdpReferenceAllOf) SetMultiFactorAuthentication(v bool) {
	o.MultiFactorAuthentication = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamIdpReferenceAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdpReferenceAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamIdpReferenceAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamIdpReferenceAllOf) SetName(v string) {
	o.Name = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamIdpReferenceAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdpReferenceAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamIdpReferenceAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamIdpReferenceAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *IamIdpReferenceAllOf) GetIdp() IamIdpRelationship {
	if o == nil || o.Idp == nil {
		var ret IamIdpRelationship
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdpReferenceAllOf) GetIdpOk() (*IamIdpRelationship, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *IamIdpReferenceAllOf) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given IamIdpRelationship and assigns it to the Idp field.
func (o *IamIdpReferenceAllOf) SetIdp(v IamIdpRelationship) {
	o.Idp = &v
}

// GetUserPreferences returns the UserPreferences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdpReferenceAllOf) GetUserPreferences() []IamUserPreferenceRelationship {
	if o == nil  {
		var ret []IamUserPreferenceRelationship
		return ret
	}
	return o.UserPreferences
}

// GetUserPreferencesOk returns a tuple with the UserPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdpReferenceAllOf) GetUserPreferencesOk() (*[]IamUserPreferenceRelationship, bool) {
	if o == nil || o.UserPreferences == nil {
		return nil, false
	}
	return &o.UserPreferences, true
}

// HasUserPreferences returns a boolean if a field has been set.
func (o *IamIdpReferenceAllOf) HasUserPreferences() bool {
	if o != nil && o.UserPreferences != nil {
		return true
	}

	return false
}

// SetUserPreferences gets a reference to the given []IamUserPreferenceRelationship and assigns it to the UserPreferences field.
func (o *IamIdpReferenceAllOf) SetUserPreferences(v []IamUserPreferenceRelationship) {
	o.UserPreferences = v
}

// GetUsergroups returns the Usergroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdpReferenceAllOf) GetUsergroups() []IamUserGroupRelationship {
	if o == nil  {
		var ret []IamUserGroupRelationship
		return ret
	}
	return o.Usergroups
}

// GetUsergroupsOk returns a tuple with the Usergroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdpReferenceAllOf) GetUsergroupsOk() (*[]IamUserGroupRelationship, bool) {
	if o == nil || o.Usergroups == nil {
		return nil, false
	}
	return &o.Usergroups, true
}

// HasUsergroups returns a boolean if a field has been set.
func (o *IamIdpReferenceAllOf) HasUsergroups() bool {
	if o != nil && o.Usergroups != nil {
		return true
	}

	return false
}

// SetUsergroups gets a reference to the given []IamUserGroupRelationship and assigns it to the Usergroups field.
func (o *IamIdpReferenceAllOf) SetUsergroups(v []IamUserGroupRelationship) {
	o.Usergroups = v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdpReferenceAllOf) GetUsers() []IamUserRelationship {
	if o == nil  {
		var ret []IamUserRelationship
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdpReferenceAllOf) GetUsersOk() (*[]IamUserRelationship, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *IamIdpReferenceAllOf) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []IamUserRelationship and assigns it to the Users field.
func (o *IamIdpReferenceAllOf) SetUsers(v []IamUserRelationship) {
	o.Users = v
}

func (o IamIdpReferenceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DomainName != nil {
		toSerialize["DomainName"] = o.DomainName
	}
	if o.IdpEntityId != nil {
		toSerialize["IdpEntityId"] = o.IdpEntityId
	}
	if o.MultiFactorAuthentication != nil {
		toSerialize["MultiFactorAuthentication"] = o.MultiFactorAuthentication
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Idp != nil {
		toSerialize["Idp"] = o.Idp
	}
	if o.UserPreferences != nil {
		toSerialize["UserPreferences"] = o.UserPreferences
	}
	if o.Usergroups != nil {
		toSerialize["Usergroups"] = o.Usergroups
	}
	if o.Users != nil {
		toSerialize["Users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamIdpReferenceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamIdpReferenceAllOf := _IamIdpReferenceAllOf{}

	if err = json.Unmarshal(bytes, &varIamIdpReferenceAllOf); err == nil {
		*o = IamIdpReferenceAllOf(varIamIdpReferenceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DomainName")
		delete(additionalProperties, "IdpEntityId")
		delete(additionalProperties, "MultiFactorAuthentication")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Idp")
		delete(additionalProperties, "UserPreferences")
		delete(additionalProperties, "Usergroups")
		delete(additionalProperties, "Users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamIdpReferenceAllOf struct {
	value *IamIdpReferenceAllOf
	isSet bool
}

func (v NullableIamIdpReferenceAllOf) Get() *IamIdpReferenceAllOf {
	return v.value
}

func (v *NullableIamIdpReferenceAllOf) Set(val *IamIdpReferenceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamIdpReferenceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamIdpReferenceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamIdpReferenceAllOf(val *IamIdpReferenceAllOf) *NullableIamIdpReferenceAllOf {
	return &NullableIamIdpReferenceAllOf{value: val, isSet: true}
}

func (v NullableIamIdpReferenceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamIdpReferenceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


