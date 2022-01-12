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
)

// IaasLicenseInfoAllOf Definition of the list of properties defined in 'iaas.LicenseInfo', excluding properties defined in parent classes.
type IaasLicenseInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// UCS Director license expiration date.
	LicenseExpirationDate *string `json:"LicenseExpirationDate,omitempty"`
	LicenseKeysInfo []IaasLicenseKeysInfo `json:"LicenseKeysInfo,omitempty"`
	// License type of UCSD whether it is EVAL/Permanent/Subscription..
	LicenseType *string `json:"LicenseType,omitempty"`
	LicenseUtilizationInfo []IaasLicenseUtilizationInfo `json:"LicenseUtilizationInfo,omitempty"`
	Guid *IaasUcsdInfoRelationship `json:"Guid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasLicenseInfoAllOf IaasLicenseInfoAllOf

// NewIaasLicenseInfoAllOf instantiates a new IaasLicenseInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasLicenseInfoAllOf(classId string, objectType string) *IaasLicenseInfoAllOf {
	this := IaasLicenseInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasLicenseInfoAllOfWithDefaults instantiates a new IaasLicenseInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasLicenseInfoAllOfWithDefaults() *IaasLicenseInfoAllOf {
	this := IaasLicenseInfoAllOf{}
	var classId string = "iaas.LicenseInfo"
	this.ClassId = classId
	var objectType string = "iaas.LicenseInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasLicenseInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasLicenseInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasLicenseInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IaasLicenseInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasLicenseInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasLicenseInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLicenseExpirationDate returns the LicenseExpirationDate field value if set, zero value otherwise.
func (o *IaasLicenseInfoAllOf) GetLicenseExpirationDate() string {
	if o == nil || o.LicenseExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.LicenseExpirationDate
}

// GetLicenseExpirationDateOk returns a tuple with the LicenseExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseInfoAllOf) GetLicenseExpirationDateOk() (*string, bool) {
	if o == nil || o.LicenseExpirationDate == nil {
		return nil, false
	}
	return o.LicenseExpirationDate, true
}

// HasLicenseExpirationDate returns a boolean if a field has been set.
func (o *IaasLicenseInfoAllOf) HasLicenseExpirationDate() bool {
	if o != nil && o.LicenseExpirationDate != nil {
		return true
	}

	return false
}

// SetLicenseExpirationDate gets a reference to the given string and assigns it to the LicenseExpirationDate field.
func (o *IaasLicenseInfoAllOf) SetLicenseExpirationDate(v string) {
	o.LicenseExpirationDate = &v
}

// GetLicenseKeysInfo returns the LicenseKeysInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasLicenseInfoAllOf) GetLicenseKeysInfo() []IaasLicenseKeysInfo {
	if o == nil  {
		var ret []IaasLicenseKeysInfo
		return ret
	}
	return o.LicenseKeysInfo
}

// GetLicenseKeysInfoOk returns a tuple with the LicenseKeysInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasLicenseInfoAllOf) GetLicenseKeysInfoOk() (*[]IaasLicenseKeysInfo, bool) {
	if o == nil || o.LicenseKeysInfo == nil {
		return nil, false
	}
	return &o.LicenseKeysInfo, true
}

// HasLicenseKeysInfo returns a boolean if a field has been set.
func (o *IaasLicenseInfoAllOf) HasLicenseKeysInfo() bool {
	if o != nil && o.LicenseKeysInfo != nil {
		return true
	}

	return false
}

// SetLicenseKeysInfo gets a reference to the given []IaasLicenseKeysInfo and assigns it to the LicenseKeysInfo field.
func (o *IaasLicenseInfoAllOf) SetLicenseKeysInfo(v []IaasLicenseKeysInfo) {
	o.LicenseKeysInfo = v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *IaasLicenseInfoAllOf) GetLicenseType() string {
	if o == nil || o.LicenseType == nil {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseInfoAllOf) GetLicenseTypeOk() (*string, bool) {
	if o == nil || o.LicenseType == nil {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *IaasLicenseInfoAllOf) HasLicenseType() bool {
	if o != nil && o.LicenseType != nil {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *IaasLicenseInfoAllOf) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetLicenseUtilizationInfo returns the LicenseUtilizationInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasLicenseInfoAllOf) GetLicenseUtilizationInfo() []IaasLicenseUtilizationInfo {
	if o == nil  {
		var ret []IaasLicenseUtilizationInfo
		return ret
	}
	return o.LicenseUtilizationInfo
}

// GetLicenseUtilizationInfoOk returns a tuple with the LicenseUtilizationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasLicenseInfoAllOf) GetLicenseUtilizationInfoOk() (*[]IaasLicenseUtilizationInfo, bool) {
	if o == nil || o.LicenseUtilizationInfo == nil {
		return nil, false
	}
	return &o.LicenseUtilizationInfo, true
}

// HasLicenseUtilizationInfo returns a boolean if a field has been set.
func (o *IaasLicenseInfoAllOf) HasLicenseUtilizationInfo() bool {
	if o != nil && o.LicenseUtilizationInfo != nil {
		return true
	}

	return false
}

// SetLicenseUtilizationInfo gets a reference to the given []IaasLicenseUtilizationInfo and assigns it to the LicenseUtilizationInfo field.
func (o *IaasLicenseInfoAllOf) SetLicenseUtilizationInfo(v []IaasLicenseUtilizationInfo) {
	o.LicenseUtilizationInfo = v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *IaasLicenseInfoAllOf) GetGuid() IaasUcsdInfoRelationship {
	if o == nil || o.Guid == nil {
		var ret IaasUcsdInfoRelationship
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseInfoAllOf) GetGuidOk() (*IaasUcsdInfoRelationship, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *IaasLicenseInfoAllOf) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given IaasUcsdInfoRelationship and assigns it to the Guid field.
func (o *IaasLicenseInfoAllOf) SetGuid(v IaasUcsdInfoRelationship) {
	o.Guid = &v
}

func (o IaasLicenseInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LicenseExpirationDate != nil {
		toSerialize["LicenseExpirationDate"] = o.LicenseExpirationDate
	}
	if o.LicenseKeysInfo != nil {
		toSerialize["LicenseKeysInfo"] = o.LicenseKeysInfo
	}
	if o.LicenseType != nil {
		toSerialize["LicenseType"] = o.LicenseType
	}
	if o.LicenseUtilizationInfo != nil {
		toSerialize["LicenseUtilizationInfo"] = o.LicenseUtilizationInfo
	}
	if o.Guid != nil {
		toSerialize["Guid"] = o.Guid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IaasLicenseInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIaasLicenseInfoAllOf := _IaasLicenseInfoAllOf{}

	if err = json.Unmarshal(bytes, &varIaasLicenseInfoAllOf); err == nil {
		*o = IaasLicenseInfoAllOf(varIaasLicenseInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LicenseExpirationDate")
		delete(additionalProperties, "LicenseKeysInfo")
		delete(additionalProperties, "LicenseType")
		delete(additionalProperties, "LicenseUtilizationInfo")
		delete(additionalProperties, "Guid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIaasLicenseInfoAllOf struct {
	value *IaasLicenseInfoAllOf
	isSet bool
}

func (v NullableIaasLicenseInfoAllOf) Get() *IaasLicenseInfoAllOf {
	return v.value
}

func (v *NullableIaasLicenseInfoAllOf) Set(val *IaasLicenseInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasLicenseInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasLicenseInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasLicenseInfoAllOf(val *IaasLicenseInfoAllOf) *NullableIaasLicenseInfoAllOf {
	return &NullableIaasLicenseInfoAllOf{value: val, isSet: true}
}

func (v NullableIaasLicenseInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasLicenseInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

