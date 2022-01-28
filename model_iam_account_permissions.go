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

// IamAccountPermissions Users can log in through the base URL (https://intersight.com) or account-specific URLs. When the Intersight user logs in through the base URL, Intersight identifies the accounts and permissions within each account which the user has access to. In case multiple permissions are identified, the user and session objects are created in the onboarding-user account, and the session object is updated with account and permission information. Intersight GUI uses this information to show available accounts and permissions for the user to select.
type IamAccountPermissions struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// MOID of the account which a user can select after authentication.
	AccountIdentifier *string `json:"AccountIdentifier,omitempty"`
	// Name of the account which a user can select after authentication.
	AccountName *string `json:"AccountName,omitempty"`
	// Status of the account. Account remains inactive until a device is claimed to the account.
	AccountStatus *string `json:"AccountStatus,omitempty"`
	Permissions []IamPermissionReference `json:"Permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamAccountPermissions IamAccountPermissions

// NewIamAccountPermissions instantiates a new IamAccountPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamAccountPermissions(classId string, objectType string) *IamAccountPermissions {
	this := IamAccountPermissions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamAccountPermissionsWithDefaults instantiates a new IamAccountPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamAccountPermissionsWithDefaults() *IamAccountPermissions {
	this := IamAccountPermissions{}
	var classId string = "iam.AccountPermissions"
	this.ClassId = classId
	var objectType string = "iam.AccountPermissions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamAccountPermissions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamAccountPermissions) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamAccountPermissions) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamAccountPermissions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamAccountPermissions) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamAccountPermissions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccountIdentifier returns the AccountIdentifier field value if set, zero value otherwise.
func (o *IamAccountPermissions) GetAccountIdentifier() string {
	if o == nil || o.AccountIdentifier == nil {
		var ret string
		return ret
	}
	return *o.AccountIdentifier
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAccountPermissions) GetAccountIdentifierOk() (*string, bool) {
	if o == nil || o.AccountIdentifier == nil {
		return nil, false
	}
	return o.AccountIdentifier, true
}

// HasAccountIdentifier returns a boolean if a field has been set.
func (o *IamAccountPermissions) HasAccountIdentifier() bool {
	if o != nil && o.AccountIdentifier != nil {
		return true
	}

	return false
}

// SetAccountIdentifier gets a reference to the given string and assigns it to the AccountIdentifier field.
func (o *IamAccountPermissions) SetAccountIdentifier(v string) {
	o.AccountIdentifier = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *IamAccountPermissions) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAccountPermissions) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *IamAccountPermissions) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *IamAccountPermissions) SetAccountName(v string) {
	o.AccountName = &v
}

// GetAccountStatus returns the AccountStatus field value if set, zero value otherwise.
func (o *IamAccountPermissions) GetAccountStatus() string {
	if o == nil || o.AccountStatus == nil {
		var ret string
		return ret
	}
	return *o.AccountStatus
}

// GetAccountStatusOk returns a tuple with the AccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAccountPermissions) GetAccountStatusOk() (*string, bool) {
	if o == nil || o.AccountStatus == nil {
		return nil, false
	}
	return o.AccountStatus, true
}

// HasAccountStatus returns a boolean if a field has been set.
func (o *IamAccountPermissions) HasAccountStatus() bool {
	if o != nil && o.AccountStatus != nil {
		return true
	}

	return false
}

// SetAccountStatus gets a reference to the given string and assigns it to the AccountStatus field.
func (o *IamAccountPermissions) SetAccountStatus(v string) {
	o.AccountStatus = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccountPermissions) GetPermissions() []IamPermissionReference {
	if o == nil  {
		var ret []IamPermissionReference
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccountPermissions) GetPermissionsOk() (*[]IamPermissionReference, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *IamAccountPermissions) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []IamPermissionReference and assigns it to the Permissions field.
func (o *IamAccountPermissions) SetPermissions(v []IamPermissionReference) {
	o.Permissions = v
}

func (o IamAccountPermissions) MarshalJSON() ([]byte, error) {
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
	if o.AccountIdentifier != nil {
		toSerialize["AccountIdentifier"] = o.AccountIdentifier
	}
	if o.AccountName != nil {
		toSerialize["AccountName"] = o.AccountName
	}
	if o.AccountStatus != nil {
		toSerialize["AccountStatus"] = o.AccountStatus
	}
	if o.Permissions != nil {
		toSerialize["Permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamAccountPermissions) UnmarshalJSON(bytes []byte) (err error) {
	type IamAccountPermissionsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// MOID of the account which a user can select after authentication.
		AccountIdentifier *string `json:"AccountIdentifier,omitempty"`
		// Name of the account which a user can select after authentication.
		AccountName *string `json:"AccountName,omitempty"`
		// Status of the account. Account remains inactive until a device is claimed to the account.
		AccountStatus *string `json:"AccountStatus,omitempty"`
		Permissions []IamPermissionReference `json:"Permissions,omitempty"`
	}

	varIamAccountPermissionsWithoutEmbeddedStruct := IamAccountPermissionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamAccountPermissionsWithoutEmbeddedStruct)
	if err == nil {
		varIamAccountPermissions := _IamAccountPermissions{}
		varIamAccountPermissions.ClassId = varIamAccountPermissionsWithoutEmbeddedStruct.ClassId
		varIamAccountPermissions.ObjectType = varIamAccountPermissionsWithoutEmbeddedStruct.ObjectType
		varIamAccountPermissions.AccountIdentifier = varIamAccountPermissionsWithoutEmbeddedStruct.AccountIdentifier
		varIamAccountPermissions.AccountName = varIamAccountPermissionsWithoutEmbeddedStruct.AccountName
		varIamAccountPermissions.AccountStatus = varIamAccountPermissionsWithoutEmbeddedStruct.AccountStatus
		varIamAccountPermissions.Permissions = varIamAccountPermissionsWithoutEmbeddedStruct.Permissions
		*o = IamAccountPermissions(varIamAccountPermissions)
	} else {
		return err
	}

	varIamAccountPermissions := _IamAccountPermissions{}

	err = json.Unmarshal(bytes, &varIamAccountPermissions)
	if err == nil {
		o.MoBaseComplexType = varIamAccountPermissions.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccountIdentifier")
		delete(additionalProperties, "AccountName")
		delete(additionalProperties, "AccountStatus")
		delete(additionalProperties, "Permissions")

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

type NullableIamAccountPermissions struct {
	value *IamAccountPermissions
	isSet bool
}

func (v NullableIamAccountPermissions) Get() *IamAccountPermissions {
	return v.value
}

func (v *NullableIamAccountPermissions) Set(val *IamAccountPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableIamAccountPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableIamAccountPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamAccountPermissions(val *IamAccountPermissions) *NullableIamAccountPermissions {
	return &NullableIamAccountPermissions{value: val, isSet: true}
}

func (v NullableIamAccountPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamAccountPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


