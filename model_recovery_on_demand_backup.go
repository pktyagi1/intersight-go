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

// RecoveryOnDemandBackup Handles requests for on demand backup for a given endpoint.
type RecoveryOnDemandBackup struct {
	RecoveryAbstractBackupConfig
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	ConfigResult *RecoveryConfigResultRelationship `json:"ConfigResult,omitempty"`
	DeviceId *AssetDeviceRegistrationRelationship `json:"DeviceId,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryOnDemandBackup RecoveryOnDemandBackup

// NewRecoveryOnDemandBackup instantiates a new RecoveryOnDemandBackup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryOnDemandBackup(classId string, objectType string) *RecoveryOnDemandBackup {
	this := RecoveryOnDemandBackup{}
	this.ClassId = classId
	this.ObjectType = objectType
	var locationType string = "Network Share"
	this.LocationType = &locationType
	var protocol string = "SCP"
	this.Protocol = &protocol
	var retentionCount int64 = 10
	this.RetentionCount = &retentionCount
	return &this
}

// NewRecoveryOnDemandBackupWithDefaults instantiates a new RecoveryOnDemandBackup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryOnDemandBackupWithDefaults() *RecoveryOnDemandBackup {
	this := RecoveryOnDemandBackup{}
	var classId string = "recovery.OnDemandBackup"
	this.ClassId = classId
	var objectType string = "recovery.OnDemandBackup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecoveryOnDemandBackup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecoveryOnDemandBackup) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecoveryOnDemandBackup) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecoveryOnDemandBackup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecoveryOnDemandBackup) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecoveryOnDemandBackup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *RecoveryOnDemandBackup) GetConfigResult() RecoveryConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret RecoveryConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryOnDemandBackup) GetConfigResultOk() (*RecoveryConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *RecoveryOnDemandBackup) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given RecoveryConfigResultRelationship and assigns it to the ConfigResult field.
func (o *RecoveryOnDemandBackup) SetConfigResult(v RecoveryConfigResultRelationship) {
	o.ConfigResult = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *RecoveryOnDemandBackup) GetDeviceId() AssetDeviceRegistrationRelationship {
	if o == nil || o.DeviceId == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryOnDemandBackup) GetDeviceIdOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *RecoveryOnDemandBackup) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the DeviceId field.
func (o *RecoveryOnDemandBackup) SetDeviceId(v AssetDeviceRegistrationRelationship) {
	o.DeviceId = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *RecoveryOnDemandBackup) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryOnDemandBackup) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *RecoveryOnDemandBackup) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *RecoveryOnDemandBackup) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o RecoveryOnDemandBackup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRecoveryAbstractBackupConfig, errRecoveryAbstractBackupConfig := json.Marshal(o.RecoveryAbstractBackupConfig)
	if errRecoveryAbstractBackupConfig != nil {
		return []byte{}, errRecoveryAbstractBackupConfig
	}
	errRecoveryAbstractBackupConfig = json.Unmarshal([]byte(serializedRecoveryAbstractBackupConfig), &toSerialize)
	if errRecoveryAbstractBackupConfig != nil {
		return []byte{}, errRecoveryAbstractBackupConfig
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecoveryOnDemandBackup) UnmarshalJSON(bytes []byte) (err error) {
	type RecoveryOnDemandBackupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		ConfigResult *RecoveryConfigResultRelationship `json:"ConfigResult,omitempty"`
		DeviceId *AssetDeviceRegistrationRelationship `json:"DeviceId,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varRecoveryOnDemandBackupWithoutEmbeddedStruct := RecoveryOnDemandBackupWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varRecoveryOnDemandBackupWithoutEmbeddedStruct)
	if err == nil {
		varRecoveryOnDemandBackup := _RecoveryOnDemandBackup{}
		varRecoveryOnDemandBackup.ClassId = varRecoveryOnDemandBackupWithoutEmbeddedStruct.ClassId
		varRecoveryOnDemandBackup.ObjectType = varRecoveryOnDemandBackupWithoutEmbeddedStruct.ObjectType
		varRecoveryOnDemandBackup.ConfigResult = varRecoveryOnDemandBackupWithoutEmbeddedStruct.ConfigResult
		varRecoveryOnDemandBackup.DeviceId = varRecoveryOnDemandBackupWithoutEmbeddedStruct.DeviceId
		varRecoveryOnDemandBackup.Organization = varRecoveryOnDemandBackupWithoutEmbeddedStruct.Organization
		*o = RecoveryOnDemandBackup(varRecoveryOnDemandBackup)
	} else {
		return err
	}

	varRecoveryOnDemandBackup := _RecoveryOnDemandBackup{}

	err = json.Unmarshal(bytes, &varRecoveryOnDemandBackup)
	if err == nil {
		o.RecoveryAbstractBackupConfig = varRecoveryOnDemandBackup.RecoveryAbstractBackupConfig
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigResult")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectRecoveryAbstractBackupConfig := reflect.ValueOf(o.RecoveryAbstractBackupConfig)
		for i := 0; i < reflectRecoveryAbstractBackupConfig.Type().NumField(); i++ {
			t := reflectRecoveryAbstractBackupConfig.Type().Field(i)

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

type NullableRecoveryOnDemandBackup struct {
	value *RecoveryOnDemandBackup
	isSet bool
}

func (v NullableRecoveryOnDemandBackup) Get() *RecoveryOnDemandBackup {
	return v.value
}

func (v *NullableRecoveryOnDemandBackup) Set(val *RecoveryOnDemandBackup) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryOnDemandBackup) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryOnDemandBackup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryOnDemandBackup(val *RecoveryOnDemandBackup) *NullableRecoveryOnDemandBackup {
	return &NullableRecoveryOnDemandBackup{value: val, isSet: true}
}

func (v NullableRecoveryOnDemandBackup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryOnDemandBackup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


