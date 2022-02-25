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

// ApplianceGroupStatus Status of a group of applications.
type ApplianceGroupStatus struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of the group.
	Description *string `json:"Description,omitempty"`
	// The name of group, which includes Identity Management, Device Connector Service, Core Service, Analytics, Internal and Appliance.
	GroupName *string `json:"GroupName,omitempty"`
	// The overall API status from this group's applications.
	OverallStatus *string `json:"OverallStatus,omitempty"`
	// An array of relationships to applianceAppStatus resources.
	Apps []ApplianceAppStatusRelationship `json:"Apps,omitempty"`
	SystemStatus *ApplianceSystemStatusRelationship `json:"SystemStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceGroupStatus ApplianceGroupStatus

// NewApplianceGroupStatus instantiates a new ApplianceGroupStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceGroupStatus(classId string, objectType string) *ApplianceGroupStatus {
	this := ApplianceGroupStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceGroupStatusWithDefaults instantiates a new ApplianceGroupStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceGroupStatusWithDefaults() *ApplianceGroupStatus {
	this := ApplianceGroupStatus{}
	var classId string = "appliance.GroupStatus"
	this.ClassId = classId
	var objectType string = "appliance.GroupStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceGroupStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceGroupStatus) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceGroupStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceGroupStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceGroupStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceGroupStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplianceGroupStatus) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceGroupStatus) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplianceGroupStatus) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplianceGroupStatus) SetDescription(v string) {
	o.Description = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *ApplianceGroupStatus) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceGroupStatus) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *ApplianceGroupStatus) HasGroupName() bool {
	if o != nil && o.GroupName != nil {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *ApplianceGroupStatus) SetGroupName(v string) {
	o.GroupName = &v
}

// GetOverallStatus returns the OverallStatus field value if set, zero value otherwise.
func (o *ApplianceGroupStatus) GetOverallStatus() string {
	if o == nil || o.OverallStatus == nil {
		var ret string
		return ret
	}
	return *o.OverallStatus
}

// GetOverallStatusOk returns a tuple with the OverallStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceGroupStatus) GetOverallStatusOk() (*string, bool) {
	if o == nil || o.OverallStatus == nil {
		return nil, false
	}
	return o.OverallStatus, true
}

// HasOverallStatus returns a boolean if a field has been set.
func (o *ApplianceGroupStatus) HasOverallStatus() bool {
	if o != nil && o.OverallStatus != nil {
		return true
	}

	return false
}

// SetOverallStatus gets a reference to the given string and assigns it to the OverallStatus field.
func (o *ApplianceGroupStatus) SetOverallStatus(v string) {
	o.OverallStatus = &v
}

// GetApps returns the Apps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceGroupStatus) GetApps() []ApplianceAppStatusRelationship {
	if o == nil  {
		var ret []ApplianceAppStatusRelationship
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceGroupStatus) GetAppsOk() (*[]ApplianceAppStatusRelationship, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return &o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *ApplianceGroupStatus) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []ApplianceAppStatusRelationship and assigns it to the Apps field.
func (o *ApplianceGroupStatus) SetApps(v []ApplianceAppStatusRelationship) {
	o.Apps = v
}

// GetSystemStatus returns the SystemStatus field value if set, zero value otherwise.
func (o *ApplianceGroupStatus) GetSystemStatus() ApplianceSystemStatusRelationship {
	if o == nil || o.SystemStatus == nil {
		var ret ApplianceSystemStatusRelationship
		return ret
	}
	return *o.SystemStatus
}

// GetSystemStatusOk returns a tuple with the SystemStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceGroupStatus) GetSystemStatusOk() (*ApplianceSystemStatusRelationship, bool) {
	if o == nil || o.SystemStatus == nil {
		return nil, false
	}
	return o.SystemStatus, true
}

// HasSystemStatus returns a boolean if a field has been set.
func (o *ApplianceGroupStatus) HasSystemStatus() bool {
	if o != nil && o.SystemStatus != nil {
		return true
	}

	return false
}

// SetSystemStatus gets a reference to the given ApplianceSystemStatusRelationship and assigns it to the SystemStatus field.
func (o *ApplianceGroupStatus) SetSystemStatus(v ApplianceSystemStatusRelationship) {
	o.SystemStatus = &v
}

func (o ApplianceGroupStatus) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.GroupName != nil {
		toSerialize["GroupName"] = o.GroupName
	}
	if o.OverallStatus != nil {
		toSerialize["OverallStatus"] = o.OverallStatus
	}
	if o.Apps != nil {
		toSerialize["Apps"] = o.Apps
	}
	if o.SystemStatus != nil {
		toSerialize["SystemStatus"] = o.SystemStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceGroupStatus) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceGroupStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description of the group.
		Description *string `json:"Description,omitempty"`
		// The name of group, which includes Identity Management, Device Connector Service, Core Service, Analytics, Internal and Appliance.
		GroupName *string `json:"GroupName,omitempty"`
		// The overall API status from this group's applications.
		OverallStatus *string `json:"OverallStatus,omitempty"`
		// An array of relationships to applianceAppStatus resources.
		Apps []ApplianceAppStatusRelationship `json:"Apps,omitempty"`
		SystemStatus *ApplianceSystemStatusRelationship `json:"SystemStatus,omitempty"`
	}

	varApplianceGroupStatusWithoutEmbeddedStruct := ApplianceGroupStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceGroupStatusWithoutEmbeddedStruct)
	if err == nil {
		varApplianceGroupStatus := _ApplianceGroupStatus{}
		varApplianceGroupStatus.ClassId = varApplianceGroupStatusWithoutEmbeddedStruct.ClassId
		varApplianceGroupStatus.ObjectType = varApplianceGroupStatusWithoutEmbeddedStruct.ObjectType
		varApplianceGroupStatus.Description = varApplianceGroupStatusWithoutEmbeddedStruct.Description
		varApplianceGroupStatus.GroupName = varApplianceGroupStatusWithoutEmbeddedStruct.GroupName
		varApplianceGroupStatus.OverallStatus = varApplianceGroupStatusWithoutEmbeddedStruct.OverallStatus
		varApplianceGroupStatus.Apps = varApplianceGroupStatusWithoutEmbeddedStruct.Apps
		varApplianceGroupStatus.SystemStatus = varApplianceGroupStatusWithoutEmbeddedStruct.SystemStatus
		*o = ApplianceGroupStatus(varApplianceGroupStatus)
	} else {
		return err
	}

	varApplianceGroupStatus := _ApplianceGroupStatus{}

	err = json.Unmarshal(bytes, &varApplianceGroupStatus)
	if err == nil {
		o.MoBaseMo = varApplianceGroupStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "GroupName")
		delete(additionalProperties, "OverallStatus")
		delete(additionalProperties, "Apps")
		delete(additionalProperties, "SystemStatus")

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

type NullableApplianceGroupStatus struct {
	value *ApplianceGroupStatus
	isSet bool
}

func (v NullableApplianceGroupStatus) Get() *ApplianceGroupStatus {
	return v.value
}

func (v *NullableApplianceGroupStatus) Set(val *ApplianceGroupStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceGroupStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceGroupStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceGroupStatus(val *ApplianceGroupStatus) *NullableApplianceGroupStatus {
	return &NullableApplianceGroupStatus{value: val, isSet: true}
}

func (v NullableApplianceGroupStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceGroupStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


