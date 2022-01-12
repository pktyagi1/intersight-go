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
	"time"
	"reflect"
	"strings"
)

// ApplianceBackupPolicy BackupPolicy stores the Intersight Appliance's backup policy. There will be only one BackupPolicy managed object in the Intersight Appliance. Default backup policy managed object is created during the Intersight Appliance setup, and it is configured in the manual backup mode.
type ApplianceBackupPolicy struct {
	ApplianceBackupBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The next backup time set by the backup scheduler. Backup scheduler calculates the next backup time with the user-defined schedule set in the Schedule field.
	BackupTime *time.Time `json:"BackupTime,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Backup mode of the appliance. Automatic backups of the appliance are not initiated if this property is set to 'true' and the backup schedule field is ignored.
	ManualBackup *bool `json:"ManualBackup,omitempty"`
	// Password to authenticate the file server.
	Password *string `json:"Password,omitempty"`
	Schedule NullableOnpremSchedule `json:"Schedule,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceBackupPolicy ApplianceBackupPolicy

// NewApplianceBackupPolicy instantiates a new ApplianceBackupPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceBackupPolicy(classId string, objectType string) *ApplianceBackupPolicy {
	this := ApplianceBackupPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var protocol string = "scp"
	this.Protocol = &protocol
	return &this
}

// NewApplianceBackupPolicyWithDefaults instantiates a new ApplianceBackupPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceBackupPolicyWithDefaults() *ApplianceBackupPolicy {
	this := ApplianceBackupPolicy{}
	var classId string = "appliance.BackupPolicy"
	this.ClassId = classId
	var objectType string = "appliance.BackupPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceBackupPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceBackupPolicy) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceBackupPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceBackupPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceBackupPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceBackupPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBackupTime returns the BackupTime field value if set, zero value otherwise.
func (o *ApplianceBackupPolicy) GetBackupTime() time.Time {
	if o == nil || o.BackupTime == nil {
		var ret time.Time
		return ret
	}
	return *o.BackupTime
}

// GetBackupTimeOk returns a tuple with the BackupTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupPolicy) GetBackupTimeOk() (*time.Time, bool) {
	if o == nil || o.BackupTime == nil {
		return nil, false
	}
	return o.BackupTime, true
}

// HasBackupTime returns a boolean if a field has been set.
func (o *ApplianceBackupPolicy) HasBackupTime() bool {
	if o != nil && o.BackupTime != nil {
		return true
	}

	return false
}

// SetBackupTime gets a reference to the given time.Time and assigns it to the BackupTime field.
func (o *ApplianceBackupPolicy) SetBackupTime(v time.Time) {
	o.BackupTime = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *ApplianceBackupPolicy) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupPolicy) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *ApplianceBackupPolicy) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *ApplianceBackupPolicy) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetManualBackup returns the ManualBackup field value if set, zero value otherwise.
func (o *ApplianceBackupPolicy) GetManualBackup() bool {
	if o == nil || o.ManualBackup == nil {
		var ret bool
		return ret
	}
	return *o.ManualBackup
}

// GetManualBackupOk returns a tuple with the ManualBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupPolicy) GetManualBackupOk() (*bool, bool) {
	if o == nil || o.ManualBackup == nil {
		return nil, false
	}
	return o.ManualBackup, true
}

// HasManualBackup returns a boolean if a field has been set.
func (o *ApplianceBackupPolicy) HasManualBackup() bool {
	if o != nil && o.ManualBackup != nil {
		return true
	}

	return false
}

// SetManualBackup gets a reference to the given bool and assigns it to the ManualBackup field.
func (o *ApplianceBackupPolicy) SetManualBackup(v bool) {
	o.ManualBackup = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApplianceBackupPolicy) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupPolicy) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApplianceBackupPolicy) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApplianceBackupPolicy) SetPassword(v string) {
	o.Password = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceBackupPolicy) GetSchedule() OnpremSchedule {
	if o == nil || o.Schedule.Get() == nil {
		var ret OnpremSchedule
		return ret
	}
	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceBackupPolicy) GetScheduleOk() (*OnpremSchedule, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// HasSchedule returns a boolean if a field has been set.
func (o *ApplianceBackupPolicy) HasSchedule() bool {
	if o != nil && o.Schedule.IsSet() {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NullableOnpremSchedule and assigns it to the Schedule field.
func (o *ApplianceBackupPolicy) SetSchedule(v OnpremSchedule) {
	o.Schedule.Set(&v)
}
// SetScheduleNil sets the value for Schedule to be an explicit nil
func (o *ApplianceBackupPolicy) SetScheduleNil() {
	o.Schedule.Set(nil)
}

// UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
func (o *ApplianceBackupPolicy) UnsetSchedule() {
	o.Schedule.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceBackupPolicy) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupPolicy) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceBackupPolicy) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceBackupPolicy) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o ApplianceBackupPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedApplianceBackupBase, errApplianceBackupBase := json.Marshal(o.ApplianceBackupBase)
	if errApplianceBackupBase != nil {
		return []byte{}, errApplianceBackupBase
	}
	errApplianceBackupBase = json.Unmarshal([]byte(serializedApplianceBackupBase), &toSerialize)
	if errApplianceBackupBase != nil {
		return []byte{}, errApplianceBackupBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BackupTime != nil {
		toSerialize["BackupTime"] = o.BackupTime
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.ManualBackup != nil {
		toSerialize["ManualBackup"] = o.ManualBackup
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Schedule.IsSet() {
		toSerialize["Schedule"] = o.Schedule.Get()
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceBackupPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceBackupPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The next backup time set by the backup scheduler. Backup scheduler calculates the next backup time with the user-defined schedule set in the Schedule field.
		BackupTime *time.Time `json:"BackupTime,omitempty"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// Backup mode of the appliance. Automatic backups of the appliance are not initiated if this property is set to 'true' and the backup schedule field is ignored.
		ManualBackup *bool `json:"ManualBackup,omitempty"`
		// Password to authenticate the file server.
		Password *string `json:"Password,omitempty"`
		Schedule NullableOnpremSchedule `json:"Schedule,omitempty"`
		Account *IamAccountRelationship `json:"Account,omitempty"`
	}

	varApplianceBackupPolicyWithoutEmbeddedStruct := ApplianceBackupPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceBackupPolicyWithoutEmbeddedStruct)
	if err == nil {
		varApplianceBackupPolicy := _ApplianceBackupPolicy{}
		varApplianceBackupPolicy.ClassId = varApplianceBackupPolicyWithoutEmbeddedStruct.ClassId
		varApplianceBackupPolicy.ObjectType = varApplianceBackupPolicyWithoutEmbeddedStruct.ObjectType
		varApplianceBackupPolicy.BackupTime = varApplianceBackupPolicyWithoutEmbeddedStruct.BackupTime
		varApplianceBackupPolicy.IsPasswordSet = varApplianceBackupPolicyWithoutEmbeddedStruct.IsPasswordSet
		varApplianceBackupPolicy.ManualBackup = varApplianceBackupPolicyWithoutEmbeddedStruct.ManualBackup
		varApplianceBackupPolicy.Password = varApplianceBackupPolicyWithoutEmbeddedStruct.Password
		varApplianceBackupPolicy.Schedule = varApplianceBackupPolicyWithoutEmbeddedStruct.Schedule
		varApplianceBackupPolicy.Account = varApplianceBackupPolicyWithoutEmbeddedStruct.Account
		*o = ApplianceBackupPolicy(varApplianceBackupPolicy)
	} else {
		return err
	}

	varApplianceBackupPolicy := _ApplianceBackupPolicy{}

	err = json.Unmarshal(bytes, &varApplianceBackupPolicy)
	if err == nil {
		o.ApplianceBackupBase = varApplianceBackupPolicy.ApplianceBackupBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BackupTime")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "ManualBackup")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Schedule")
		delete(additionalProperties, "Account")

		// remove fields from embedded structs
		reflectApplianceBackupBase := reflect.ValueOf(o.ApplianceBackupBase)
		for i := 0; i < reflectApplianceBackupBase.Type().NumField(); i++ {
			t := reflectApplianceBackupBase.Type().Field(i)

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

type NullableApplianceBackupPolicy struct {
	value *ApplianceBackupPolicy
	isSet bool
}

func (v NullableApplianceBackupPolicy) Get() *ApplianceBackupPolicy {
	return v.value
}

func (v *NullableApplianceBackupPolicy) Set(val *ApplianceBackupPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceBackupPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceBackupPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceBackupPolicy(val *ApplianceBackupPolicy) *NullableApplianceBackupPolicy {
	return &NullableApplianceBackupPolicy{value: val, isSet: true}
}

func (v NullableApplianceBackupPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceBackupPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

