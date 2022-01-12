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

// ApplianceDataExportPolicyAllOf Definition of the list of properties defined in 'appliance.DataExportPolicy', excluding properties defined in parent classes.
type ApplianceDataExportPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of the data collection mode. If the value is 'true', then data collection is enabled.
	Enable *bool `json:"Enable,omitempty"`
	// Name of the Data Export Policy.
	Name *string `json:"Name,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	ParentConfig *ApplianceDataExportPolicyRelationship `json:"ParentConfig,omitempty"`
	// An array of relationships to applianceDataExportPolicy resources.
	SubConfigs []ApplianceDataExportPolicyRelationship `json:"SubConfigs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceDataExportPolicyAllOf ApplianceDataExportPolicyAllOf

// NewApplianceDataExportPolicyAllOf instantiates a new ApplianceDataExportPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceDataExportPolicyAllOf(classId string, objectType string) *ApplianceDataExportPolicyAllOf {
	this := ApplianceDataExportPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceDataExportPolicyAllOfWithDefaults instantiates a new ApplianceDataExportPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceDataExportPolicyAllOfWithDefaults() *ApplianceDataExportPolicyAllOf {
	this := ApplianceDataExportPolicyAllOf{}
	var classId string = "appliance.DataExportPolicy"
	this.ClassId = classId
	var objectType string = "appliance.DataExportPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceDataExportPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceDataExportPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceDataExportPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceDataExportPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceDataExportPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceDataExportPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *ApplianceDataExportPolicyAllOf) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDataExportPolicyAllOf) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *ApplianceDataExportPolicyAllOf) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *ApplianceDataExportPolicyAllOf) SetEnable(v bool) {
	o.Enable = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplianceDataExportPolicyAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDataExportPolicyAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplianceDataExportPolicyAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplianceDataExportPolicyAllOf) SetName(v string) {
	o.Name = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceDataExportPolicyAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDataExportPolicyAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceDataExportPolicyAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceDataExportPolicyAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetParentConfig returns the ParentConfig field value if set, zero value otherwise.
func (o *ApplianceDataExportPolicyAllOf) GetParentConfig() ApplianceDataExportPolicyRelationship {
	if o == nil || o.ParentConfig == nil {
		var ret ApplianceDataExportPolicyRelationship
		return ret
	}
	return *o.ParentConfig
}

// GetParentConfigOk returns a tuple with the ParentConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDataExportPolicyAllOf) GetParentConfigOk() (*ApplianceDataExportPolicyRelationship, bool) {
	if o == nil || o.ParentConfig == nil {
		return nil, false
	}
	return o.ParentConfig, true
}

// HasParentConfig returns a boolean if a field has been set.
func (o *ApplianceDataExportPolicyAllOf) HasParentConfig() bool {
	if o != nil && o.ParentConfig != nil {
		return true
	}

	return false
}

// SetParentConfig gets a reference to the given ApplianceDataExportPolicyRelationship and assigns it to the ParentConfig field.
func (o *ApplianceDataExportPolicyAllOf) SetParentConfig(v ApplianceDataExportPolicyRelationship) {
	o.ParentConfig = &v
}

// GetSubConfigs returns the SubConfigs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceDataExportPolicyAllOf) GetSubConfigs() []ApplianceDataExportPolicyRelationship {
	if o == nil  {
		var ret []ApplianceDataExportPolicyRelationship
		return ret
	}
	return o.SubConfigs
}

// GetSubConfigsOk returns a tuple with the SubConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceDataExportPolicyAllOf) GetSubConfigsOk() (*[]ApplianceDataExportPolicyRelationship, bool) {
	if o == nil || o.SubConfigs == nil {
		return nil, false
	}
	return &o.SubConfigs, true
}

// HasSubConfigs returns a boolean if a field has been set.
func (o *ApplianceDataExportPolicyAllOf) HasSubConfigs() bool {
	if o != nil && o.SubConfigs != nil {
		return true
	}

	return false
}

// SetSubConfigs gets a reference to the given []ApplianceDataExportPolicyRelationship and assigns it to the SubConfigs field.
func (o *ApplianceDataExportPolicyAllOf) SetSubConfigs(v []ApplianceDataExportPolicyRelationship) {
	o.SubConfigs = v
}

func (o ApplianceDataExportPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enable != nil {
		toSerialize["Enable"] = o.Enable
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.ParentConfig != nil {
		toSerialize["ParentConfig"] = o.ParentConfig
	}
	if o.SubConfigs != nil {
		toSerialize["SubConfigs"] = o.SubConfigs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceDataExportPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceDataExportPolicyAllOf := _ApplianceDataExportPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceDataExportPolicyAllOf); err == nil {
		*o = ApplianceDataExportPolicyAllOf(varApplianceDataExportPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enable")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "ParentConfig")
		delete(additionalProperties, "SubConfigs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceDataExportPolicyAllOf struct {
	value *ApplianceDataExportPolicyAllOf
	isSet bool
}

func (v NullableApplianceDataExportPolicyAllOf) Get() *ApplianceDataExportPolicyAllOf {
	return v.value
}

func (v *NullableApplianceDataExportPolicyAllOf) Set(val *ApplianceDataExportPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceDataExportPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceDataExportPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceDataExportPolicyAllOf(val *ApplianceDataExportPolicyAllOf) *NullableApplianceDataExportPolicyAllOf {
	return &NullableApplianceDataExportPolicyAllOf{value: val, isSet: true}
}

func (v NullableApplianceDataExportPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceDataExportPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

