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

// HyperflexBackupPolicySettings HyperFlex Backup Policy settings definition.
type HyperflexBackupPolicySettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Backup datastore name used during auto creation of the datastore. All Virtual Machines created in this datastore will be automatically backed up.
	BackupDataStoreName *string `json:"BackupDataStoreName,omitempty"`
	// Capacity of Backup source datastore.
	BackupDataStoreSize *int64 `json:"BackupDataStoreSize,omitempty"`
	// Unit of backupDataStoreSize.  Allowable values are \"GB\" (Gigabytes) or \"TB\" (Terabytes). * `GB` - BackupDataStoreSize is specified in Gigabytes. * `TB` - BackupDataStoreSize is specified in Terabytes.
	BackupDataStoreSizeUnit *string `json:"BackupDataStoreSizeUnit,omitempty"`
	// Whether the datastore is encrypted or not.
	DataStoreEncryptionEnabled *bool `json:"DataStoreEncryptionEnabled,omitempty"`
	// Number of snapshots that will be retained.
	LocalSnapshotRetentionCount *int64 `json:"LocalSnapshotRetentionCount,omitempty"`
	// Snapshot replication interval.
	ReplicationIntervalInMinutes *int64 `json:"ReplicationIntervalInMinutes,omitempty"`
	// Replication cluster pairing name prefix.
	ReplicationPairNamePrefix *string `json:"ReplicationPairNamePrefix,omitempty"`
	// Number of snapshots that will be retained.
	SnapshotRetentionCount *int64 `json:"SnapshotRetentionCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexBackupPolicySettings HyperflexBackupPolicySettings

// NewHyperflexBackupPolicySettings instantiates a new HyperflexBackupPolicySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexBackupPolicySettings(classId string, objectType string) *HyperflexBackupPolicySettings {
	this := HyperflexBackupPolicySettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexBackupPolicySettingsWithDefaults instantiates a new HyperflexBackupPolicySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexBackupPolicySettingsWithDefaults() *HyperflexBackupPolicySettings {
	this := HyperflexBackupPolicySettings{}
	var classId string = "hyperflex.BackupPolicySettings"
	this.ClassId = classId
	var objectType string = "hyperflex.BackupPolicySettings"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexBackupPolicySettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexBackupPolicySettings) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexBackupPolicySettings) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexBackupPolicySettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexBackupPolicySettings) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexBackupPolicySettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBackupDataStoreName returns the BackupDataStoreName field value if set, zero value otherwise.
func (o *HyperflexBackupPolicySettings) GetBackupDataStoreName() string {
	if o == nil || o.BackupDataStoreName == nil {
		var ret string
		return ret
	}
	return *o.BackupDataStoreName
}

// GetBackupDataStoreNameOk returns a tuple with the BackupDataStoreName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBackupPolicySettings) GetBackupDataStoreNameOk() (*string, bool) {
	if o == nil || o.BackupDataStoreName == nil {
		return nil, false
	}
	return o.BackupDataStoreName, true
}

// HasBackupDataStoreName returns a boolean if a field has been set.
func (o *HyperflexBackupPolicySettings) HasBackupDataStoreName() bool {
	if o != nil && o.BackupDataStoreName != nil {
		return true
	}

	return false
}

// SetBackupDataStoreName gets a reference to the given string and assigns it to the BackupDataStoreName field.
func (o *HyperflexBackupPolicySettings) SetBackupDataStoreName(v string) {
	o.BackupDataStoreName = &v
}

// GetBackupDataStoreSize returns the BackupDataStoreSize field value if set, zero value otherwise.
func (o *HyperflexBackupPolicySettings) GetBackupDataStoreSize() int64 {
	if o == nil || o.BackupDataStoreSize == nil {
		var ret int64
		return ret
	}
	return *o.BackupDataStoreSize
}

// GetBackupDataStoreSizeOk returns a tuple with the BackupDataStoreSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBackupPolicySettings) GetBackupDataStoreSizeOk() (*int64, bool) {
	if o == nil || o.BackupDataStoreSize == nil {
		return nil, false
	}
	return o.BackupDataStoreSize, true
}

// HasBackupDataStoreSize returns a boolean if a field has been set.
func (o *HyperflexBackupPolicySettings) HasBackupDataStoreSize() bool {
	if o != nil && o.BackupDataStoreSize != nil {
		return true
	}

	return false
}

// SetBackupDataStoreSize gets a reference to the given int64 and assigns it to the BackupDataStoreSize field.
func (o *HyperflexBackupPolicySettings) SetBackupDataStoreSize(v int64) {
	o.BackupDataStoreSize = &v
}

// GetBackupDataStoreSizeUnit returns the BackupDataStoreSizeUnit field value if set, zero value otherwise.
func (o *HyperflexBackupPolicySettings) GetBackupDataStoreSizeUnit() string {
	if o == nil || o.BackupDataStoreSizeUnit == nil {
		var ret string
		return ret
	}
	return *o.BackupDataStoreSizeUnit
}

// GetBackupDataStoreSizeUnitOk returns a tuple with the BackupDataStoreSizeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBackupPolicySettings) GetBackupDataStoreSizeUnitOk() (*string, bool) {
	if o == nil || o.BackupDataStoreSizeUnit == nil {
		return nil, false
	}
	return o.BackupDataStoreSizeUnit, true
}

// HasBackupDataStoreSizeUnit returns a boolean if a field has been set.
func (o *HyperflexBackupPolicySettings) HasBackupDataStoreSizeUnit() bool {
	if o != nil && o.BackupDataStoreSizeUnit != nil {
		return true
	}

	return false
}

// SetBackupDataStoreSizeUnit gets a reference to the given string and assigns it to the BackupDataStoreSizeUnit field.
func (o *HyperflexBackupPolicySettings) SetBackupDataStoreSizeUnit(v string) {
	o.BackupDataStoreSizeUnit = &v
}

// GetDataStoreEncryptionEnabled returns the DataStoreEncryptionEnabled field value if set, zero value otherwise.
func (o *HyperflexBackupPolicySettings) GetDataStoreEncryptionEnabled() bool {
	if o == nil || o.DataStoreEncryptionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DataStoreEncryptionEnabled
}

// GetDataStoreEncryptionEnabledOk returns a tuple with the DataStoreEncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBackupPolicySettings) GetDataStoreEncryptionEnabledOk() (*bool, bool) {
	if o == nil || o.DataStoreEncryptionEnabled == nil {
		return nil, false
	}
	return o.DataStoreEncryptionEnabled, true
}

// HasDataStoreEncryptionEnabled returns a boolean if a field has been set.
func (o *HyperflexBackupPolicySettings) HasDataStoreEncryptionEnabled() bool {
	if o != nil && o.DataStoreEncryptionEnabled != nil {
		return true
	}

	return false
}

// SetDataStoreEncryptionEnabled gets a reference to the given bool and assigns it to the DataStoreEncryptionEnabled field.
func (o *HyperflexBackupPolicySettings) SetDataStoreEncryptionEnabled(v bool) {
	o.DataStoreEncryptionEnabled = &v
}

// GetLocalSnapshotRetentionCount returns the LocalSnapshotRetentionCount field value if set, zero value otherwise.
func (o *HyperflexBackupPolicySettings) GetLocalSnapshotRetentionCount() int64 {
	if o == nil || o.LocalSnapshotRetentionCount == nil {
		var ret int64
		return ret
	}
	return *o.LocalSnapshotRetentionCount
}

// GetLocalSnapshotRetentionCountOk returns a tuple with the LocalSnapshotRetentionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBackupPolicySettings) GetLocalSnapshotRetentionCountOk() (*int64, bool) {
	if o == nil || o.LocalSnapshotRetentionCount == nil {
		return nil, false
	}
	return o.LocalSnapshotRetentionCount, true
}

// HasLocalSnapshotRetentionCount returns a boolean if a field has been set.
func (o *HyperflexBackupPolicySettings) HasLocalSnapshotRetentionCount() bool {
	if o != nil && o.LocalSnapshotRetentionCount != nil {
		return true
	}

	return false
}

// SetLocalSnapshotRetentionCount gets a reference to the given int64 and assigns it to the LocalSnapshotRetentionCount field.
func (o *HyperflexBackupPolicySettings) SetLocalSnapshotRetentionCount(v int64) {
	o.LocalSnapshotRetentionCount = &v
}

// GetReplicationIntervalInMinutes returns the ReplicationIntervalInMinutes field value if set, zero value otherwise.
func (o *HyperflexBackupPolicySettings) GetReplicationIntervalInMinutes() int64 {
	if o == nil || o.ReplicationIntervalInMinutes == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationIntervalInMinutes
}

// GetReplicationIntervalInMinutesOk returns a tuple with the ReplicationIntervalInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBackupPolicySettings) GetReplicationIntervalInMinutesOk() (*int64, bool) {
	if o == nil || o.ReplicationIntervalInMinutes == nil {
		return nil, false
	}
	return o.ReplicationIntervalInMinutes, true
}

// HasReplicationIntervalInMinutes returns a boolean if a field has been set.
func (o *HyperflexBackupPolicySettings) HasReplicationIntervalInMinutes() bool {
	if o != nil && o.ReplicationIntervalInMinutes != nil {
		return true
	}

	return false
}

// SetReplicationIntervalInMinutes gets a reference to the given int64 and assigns it to the ReplicationIntervalInMinutes field.
func (o *HyperflexBackupPolicySettings) SetReplicationIntervalInMinutes(v int64) {
	o.ReplicationIntervalInMinutes = &v
}

// GetReplicationPairNamePrefix returns the ReplicationPairNamePrefix field value if set, zero value otherwise.
func (o *HyperflexBackupPolicySettings) GetReplicationPairNamePrefix() string {
	if o == nil || o.ReplicationPairNamePrefix == nil {
		var ret string
		return ret
	}
	return *o.ReplicationPairNamePrefix
}

// GetReplicationPairNamePrefixOk returns a tuple with the ReplicationPairNamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBackupPolicySettings) GetReplicationPairNamePrefixOk() (*string, bool) {
	if o == nil || o.ReplicationPairNamePrefix == nil {
		return nil, false
	}
	return o.ReplicationPairNamePrefix, true
}

// HasReplicationPairNamePrefix returns a boolean if a field has been set.
func (o *HyperflexBackupPolicySettings) HasReplicationPairNamePrefix() bool {
	if o != nil && o.ReplicationPairNamePrefix != nil {
		return true
	}

	return false
}

// SetReplicationPairNamePrefix gets a reference to the given string and assigns it to the ReplicationPairNamePrefix field.
func (o *HyperflexBackupPolicySettings) SetReplicationPairNamePrefix(v string) {
	o.ReplicationPairNamePrefix = &v
}

// GetSnapshotRetentionCount returns the SnapshotRetentionCount field value if set, zero value otherwise.
func (o *HyperflexBackupPolicySettings) GetSnapshotRetentionCount() int64 {
	if o == nil || o.SnapshotRetentionCount == nil {
		var ret int64
		return ret
	}
	return *o.SnapshotRetentionCount
}

// GetSnapshotRetentionCountOk returns a tuple with the SnapshotRetentionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBackupPolicySettings) GetSnapshotRetentionCountOk() (*int64, bool) {
	if o == nil || o.SnapshotRetentionCount == nil {
		return nil, false
	}
	return o.SnapshotRetentionCount, true
}

// HasSnapshotRetentionCount returns a boolean if a field has been set.
func (o *HyperflexBackupPolicySettings) HasSnapshotRetentionCount() bool {
	if o != nil && o.SnapshotRetentionCount != nil {
		return true
	}

	return false
}

// SetSnapshotRetentionCount gets a reference to the given int64 and assigns it to the SnapshotRetentionCount field.
func (o *HyperflexBackupPolicySettings) SetSnapshotRetentionCount(v int64) {
	o.SnapshotRetentionCount = &v
}

func (o HyperflexBackupPolicySettings) MarshalJSON() ([]byte, error) {
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
	if o.BackupDataStoreName != nil {
		toSerialize["BackupDataStoreName"] = o.BackupDataStoreName
	}
	if o.BackupDataStoreSize != nil {
		toSerialize["BackupDataStoreSize"] = o.BackupDataStoreSize
	}
	if o.BackupDataStoreSizeUnit != nil {
		toSerialize["BackupDataStoreSizeUnit"] = o.BackupDataStoreSizeUnit
	}
	if o.DataStoreEncryptionEnabled != nil {
		toSerialize["DataStoreEncryptionEnabled"] = o.DataStoreEncryptionEnabled
	}
	if o.LocalSnapshotRetentionCount != nil {
		toSerialize["LocalSnapshotRetentionCount"] = o.LocalSnapshotRetentionCount
	}
	if o.ReplicationIntervalInMinutes != nil {
		toSerialize["ReplicationIntervalInMinutes"] = o.ReplicationIntervalInMinutes
	}
	if o.ReplicationPairNamePrefix != nil {
		toSerialize["ReplicationPairNamePrefix"] = o.ReplicationPairNamePrefix
	}
	if o.SnapshotRetentionCount != nil {
		toSerialize["SnapshotRetentionCount"] = o.SnapshotRetentionCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexBackupPolicySettings) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexBackupPolicySettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Backup datastore name used during auto creation of the datastore. All Virtual Machines created in this datastore will be automatically backed up.
		BackupDataStoreName *string `json:"BackupDataStoreName,omitempty"`
		// Capacity of Backup source datastore.
		BackupDataStoreSize *int64 `json:"BackupDataStoreSize,omitempty"`
		// Unit of backupDataStoreSize.  Allowable values are \"GB\" (Gigabytes) or \"TB\" (Terabytes). * `GB` - BackupDataStoreSize is specified in Gigabytes. * `TB` - BackupDataStoreSize is specified in Terabytes.
		BackupDataStoreSizeUnit *string `json:"BackupDataStoreSizeUnit,omitempty"`
		// Whether the datastore is encrypted or not.
		DataStoreEncryptionEnabled *bool `json:"DataStoreEncryptionEnabled,omitempty"`
		// Number of snapshots that will be retained.
		LocalSnapshotRetentionCount *int64 `json:"LocalSnapshotRetentionCount,omitempty"`
		// Snapshot replication interval.
		ReplicationIntervalInMinutes *int64 `json:"ReplicationIntervalInMinutes,omitempty"`
		// Replication cluster pairing name prefix.
		ReplicationPairNamePrefix *string `json:"ReplicationPairNamePrefix,omitempty"`
		// Number of snapshots that will be retained.
		SnapshotRetentionCount *int64 `json:"SnapshotRetentionCount,omitempty"`
	}

	varHyperflexBackupPolicySettingsWithoutEmbeddedStruct := HyperflexBackupPolicySettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexBackupPolicySettingsWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexBackupPolicySettings := _HyperflexBackupPolicySettings{}
		varHyperflexBackupPolicySettings.ClassId = varHyperflexBackupPolicySettingsWithoutEmbeddedStruct.ClassId
		varHyperflexBackupPolicySettings.ObjectType = varHyperflexBackupPolicySettingsWithoutEmbeddedStruct.ObjectType
		varHyperflexBackupPolicySettings.BackupDataStoreName = varHyperflexBackupPolicySettingsWithoutEmbeddedStruct.BackupDataStoreName
		varHyperflexBackupPolicySettings.BackupDataStoreSize = varHyperflexBackupPolicySettingsWithoutEmbeddedStruct.BackupDataStoreSize
		varHyperflexBackupPolicySettings.BackupDataStoreSizeUnit = varHyperflexBackupPolicySettingsWithoutEmbeddedStruct.BackupDataStoreSizeUnit
		varHyperflexBackupPolicySettings.DataStoreEncryptionEnabled = varHyperflexBackupPolicySettingsWithoutEmbeddedStruct.DataStoreEncryptionEnabled
		varHyperflexBackupPolicySettings.LocalSnapshotRetentionCount = varHyperflexBackupPolicySettingsWithoutEmbeddedStruct.LocalSnapshotRetentionCount
		varHyperflexBackupPolicySettings.ReplicationIntervalInMinutes = varHyperflexBackupPolicySettingsWithoutEmbeddedStruct.ReplicationIntervalInMinutes
		varHyperflexBackupPolicySettings.ReplicationPairNamePrefix = varHyperflexBackupPolicySettingsWithoutEmbeddedStruct.ReplicationPairNamePrefix
		varHyperflexBackupPolicySettings.SnapshotRetentionCount = varHyperflexBackupPolicySettingsWithoutEmbeddedStruct.SnapshotRetentionCount
		*o = HyperflexBackupPolicySettings(varHyperflexBackupPolicySettings)
	} else {
		return err
	}

	varHyperflexBackupPolicySettings := _HyperflexBackupPolicySettings{}

	err = json.Unmarshal(bytes, &varHyperflexBackupPolicySettings)
	if err == nil {
		o.MoBaseComplexType = varHyperflexBackupPolicySettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BackupDataStoreName")
		delete(additionalProperties, "BackupDataStoreSize")
		delete(additionalProperties, "BackupDataStoreSizeUnit")
		delete(additionalProperties, "DataStoreEncryptionEnabled")
		delete(additionalProperties, "LocalSnapshotRetentionCount")
		delete(additionalProperties, "ReplicationIntervalInMinutes")
		delete(additionalProperties, "ReplicationPairNamePrefix")
		delete(additionalProperties, "SnapshotRetentionCount")

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

type NullableHyperflexBackupPolicySettings struct {
	value *HyperflexBackupPolicySettings
	isSet bool
}

func (v NullableHyperflexBackupPolicySettings) Get() *HyperflexBackupPolicySettings {
	return v.value
}

func (v *NullableHyperflexBackupPolicySettings) Set(val *HyperflexBackupPolicySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexBackupPolicySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexBackupPolicySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexBackupPolicySettings(val *HyperflexBackupPolicySettings) *NullableHyperflexBackupPolicySettings {
	return &NullableHyperflexBackupPolicySettings{value: val, isSet: true}
}

func (v NullableHyperflexBackupPolicySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexBackupPolicySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


