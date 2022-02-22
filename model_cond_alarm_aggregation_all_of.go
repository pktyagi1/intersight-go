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

// CondAlarmAggregationAllOf Definition of the list of properties defined in 'cond.AlarmAggregation', excluding properties defined in parent classes.
type CondAlarmAggregationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	AlarmSummary NullableCondAlarmSummary `json:"AlarmSummary,omitempty"`
	// Count of all alarms with severity Critical, irrespective of acknowledgement status.
	CriticalAlarmsCount *int64 `json:"CriticalAlarmsCount,omitempty"`
	// Health of the managed end point. The highest severity computed from alarmSummary property is set as the health. * `None` - The Enum value None represents that there is no severity. * `Info` - The Enum value Info represents the Informational level of severity. * `Critical` - The Enum value Critical represents the Critical level of severity. * `Warning` - The Enum value Warning represents the Warning level of severity. * `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.
	Health *string `json:"Health,omitempty"`
	// Count of all alarms with severity Info, irrespective of acknowledgement status.
	InfoAlarmsCount *int64 `json:"InfoAlarmsCount,omitempty"`
	// Managed object type. For example, FI managed object type will be network.Element.
	MoType *string `json:"MoType,omitempty"`
	// Count of all alarms with severity Warning, irrespective of acknowledgement status.
	WarningAlarmsCount *int64 `json:"WarningAlarmsCount,omitempty"`
	AlarmAggregationSource *MoBaseMoRelationship `json:"AlarmAggregationSource,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CondAlarmAggregationAllOf CondAlarmAggregationAllOf

// NewCondAlarmAggregationAllOf instantiates a new CondAlarmAggregationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondAlarmAggregationAllOf(classId string, objectType string) *CondAlarmAggregationAllOf {
	this := CondAlarmAggregationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var health string = "None"
	this.Health = &health
	return &this
}

// NewCondAlarmAggregationAllOfWithDefaults instantiates a new CondAlarmAggregationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCondAlarmAggregationAllOfWithDefaults() *CondAlarmAggregationAllOf {
	this := CondAlarmAggregationAllOf{}
	var classId string = "cond.AlarmAggregation"
	this.ClassId = classId
	var objectType string = "cond.AlarmAggregation"
	this.ObjectType = objectType
	var health string = "None"
	this.Health = &health
	return &this
}

// GetClassId returns the ClassId field value
func (o *CondAlarmAggregationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CondAlarmAggregationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CondAlarmAggregationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CondAlarmAggregationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CondAlarmAggregationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CondAlarmAggregationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlarmSummary returns the AlarmSummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CondAlarmAggregationAllOf) GetAlarmSummary() CondAlarmSummary {
	if o == nil || o.AlarmSummary.Get() == nil {
		var ret CondAlarmSummary
		return ret
	}
	return *o.AlarmSummary.Get()
}

// GetAlarmSummaryOk returns a tuple with the AlarmSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CondAlarmAggregationAllOf) GetAlarmSummaryOk() (*CondAlarmSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AlarmSummary.Get(), o.AlarmSummary.IsSet()
}

// HasAlarmSummary returns a boolean if a field has been set.
func (o *CondAlarmAggregationAllOf) HasAlarmSummary() bool {
	if o != nil && o.AlarmSummary.IsSet() {
		return true
	}

	return false
}

// SetAlarmSummary gets a reference to the given NullableCondAlarmSummary and assigns it to the AlarmSummary field.
func (o *CondAlarmAggregationAllOf) SetAlarmSummary(v CondAlarmSummary) {
	o.AlarmSummary.Set(&v)
}
// SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil
func (o *CondAlarmAggregationAllOf) SetAlarmSummaryNil() {
	o.AlarmSummary.Set(nil)
}

// UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
func (o *CondAlarmAggregationAllOf) UnsetAlarmSummary() {
	o.AlarmSummary.Unset()
}

// GetCriticalAlarmsCount returns the CriticalAlarmsCount field value if set, zero value otherwise.
func (o *CondAlarmAggregationAllOf) GetCriticalAlarmsCount() int64 {
	if o == nil || o.CriticalAlarmsCount == nil {
		var ret int64
		return ret
	}
	return *o.CriticalAlarmsCount
}

// GetCriticalAlarmsCountOk returns a tuple with the CriticalAlarmsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAggregationAllOf) GetCriticalAlarmsCountOk() (*int64, bool) {
	if o == nil || o.CriticalAlarmsCount == nil {
		return nil, false
	}
	return o.CriticalAlarmsCount, true
}

// HasCriticalAlarmsCount returns a boolean if a field has been set.
func (o *CondAlarmAggregationAllOf) HasCriticalAlarmsCount() bool {
	if o != nil && o.CriticalAlarmsCount != nil {
		return true
	}

	return false
}

// SetCriticalAlarmsCount gets a reference to the given int64 and assigns it to the CriticalAlarmsCount field.
func (o *CondAlarmAggregationAllOf) SetCriticalAlarmsCount(v int64) {
	o.CriticalAlarmsCount = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *CondAlarmAggregationAllOf) GetHealth() string {
	if o == nil || o.Health == nil {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAggregationAllOf) GetHealthOk() (*string, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *CondAlarmAggregationAllOf) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *CondAlarmAggregationAllOf) SetHealth(v string) {
	o.Health = &v
}

// GetInfoAlarmsCount returns the InfoAlarmsCount field value if set, zero value otherwise.
func (o *CondAlarmAggregationAllOf) GetInfoAlarmsCount() int64 {
	if o == nil || o.InfoAlarmsCount == nil {
		var ret int64
		return ret
	}
	return *o.InfoAlarmsCount
}

// GetInfoAlarmsCountOk returns a tuple with the InfoAlarmsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAggregationAllOf) GetInfoAlarmsCountOk() (*int64, bool) {
	if o == nil || o.InfoAlarmsCount == nil {
		return nil, false
	}
	return o.InfoAlarmsCount, true
}

// HasInfoAlarmsCount returns a boolean if a field has been set.
func (o *CondAlarmAggregationAllOf) HasInfoAlarmsCount() bool {
	if o != nil && o.InfoAlarmsCount != nil {
		return true
	}

	return false
}

// SetInfoAlarmsCount gets a reference to the given int64 and assigns it to the InfoAlarmsCount field.
func (o *CondAlarmAggregationAllOf) SetInfoAlarmsCount(v int64) {
	o.InfoAlarmsCount = &v
}

// GetMoType returns the MoType field value if set, zero value otherwise.
func (o *CondAlarmAggregationAllOf) GetMoType() string {
	if o == nil || o.MoType == nil {
		var ret string
		return ret
	}
	return *o.MoType
}

// GetMoTypeOk returns a tuple with the MoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAggregationAllOf) GetMoTypeOk() (*string, bool) {
	if o == nil || o.MoType == nil {
		return nil, false
	}
	return o.MoType, true
}

// HasMoType returns a boolean if a field has been set.
func (o *CondAlarmAggregationAllOf) HasMoType() bool {
	if o != nil && o.MoType != nil {
		return true
	}

	return false
}

// SetMoType gets a reference to the given string and assigns it to the MoType field.
func (o *CondAlarmAggregationAllOf) SetMoType(v string) {
	o.MoType = &v
}

// GetWarningAlarmsCount returns the WarningAlarmsCount field value if set, zero value otherwise.
func (o *CondAlarmAggregationAllOf) GetWarningAlarmsCount() int64 {
	if o == nil || o.WarningAlarmsCount == nil {
		var ret int64
		return ret
	}
	return *o.WarningAlarmsCount
}

// GetWarningAlarmsCountOk returns a tuple with the WarningAlarmsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAggregationAllOf) GetWarningAlarmsCountOk() (*int64, bool) {
	if o == nil || o.WarningAlarmsCount == nil {
		return nil, false
	}
	return o.WarningAlarmsCount, true
}

// HasWarningAlarmsCount returns a boolean if a field has been set.
func (o *CondAlarmAggregationAllOf) HasWarningAlarmsCount() bool {
	if o != nil && o.WarningAlarmsCount != nil {
		return true
	}

	return false
}

// SetWarningAlarmsCount gets a reference to the given int64 and assigns it to the WarningAlarmsCount field.
func (o *CondAlarmAggregationAllOf) SetWarningAlarmsCount(v int64) {
	o.WarningAlarmsCount = &v
}

// GetAlarmAggregationSource returns the AlarmAggregationSource field value if set, zero value otherwise.
func (o *CondAlarmAggregationAllOf) GetAlarmAggregationSource() MoBaseMoRelationship {
	if o == nil || o.AlarmAggregationSource == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AlarmAggregationSource
}

// GetAlarmAggregationSourceOk returns a tuple with the AlarmAggregationSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAggregationAllOf) GetAlarmAggregationSourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AlarmAggregationSource == nil {
		return nil, false
	}
	return o.AlarmAggregationSource, true
}

// HasAlarmAggregationSource returns a boolean if a field has been set.
func (o *CondAlarmAggregationAllOf) HasAlarmAggregationSource() bool {
	if o != nil && o.AlarmAggregationSource != nil {
		return true
	}

	return false
}

// SetAlarmAggregationSource gets a reference to the given MoBaseMoRelationship and assigns it to the AlarmAggregationSource field.
func (o *CondAlarmAggregationAllOf) SetAlarmAggregationSource(v MoBaseMoRelationship) {
	o.AlarmAggregationSource = &v
}

func (o CondAlarmAggregationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AlarmSummary.IsSet() {
		toSerialize["AlarmSummary"] = o.AlarmSummary.Get()
	}
	if o.CriticalAlarmsCount != nil {
		toSerialize["CriticalAlarmsCount"] = o.CriticalAlarmsCount
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.InfoAlarmsCount != nil {
		toSerialize["InfoAlarmsCount"] = o.InfoAlarmsCount
	}
	if o.MoType != nil {
		toSerialize["MoType"] = o.MoType
	}
	if o.WarningAlarmsCount != nil {
		toSerialize["WarningAlarmsCount"] = o.WarningAlarmsCount
	}
	if o.AlarmAggregationSource != nil {
		toSerialize["AlarmAggregationSource"] = o.AlarmAggregationSource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CondAlarmAggregationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCondAlarmAggregationAllOf := _CondAlarmAggregationAllOf{}

	if err = json.Unmarshal(bytes, &varCondAlarmAggregationAllOf); err == nil {
		*o = CondAlarmAggregationAllOf(varCondAlarmAggregationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlarmSummary")
		delete(additionalProperties, "CriticalAlarmsCount")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "InfoAlarmsCount")
		delete(additionalProperties, "MoType")
		delete(additionalProperties, "WarningAlarmsCount")
		delete(additionalProperties, "AlarmAggregationSource")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCondAlarmAggregationAllOf struct {
	value *CondAlarmAggregationAllOf
	isSet bool
}

func (v NullableCondAlarmAggregationAllOf) Get() *CondAlarmAggregationAllOf {
	return v.value
}

func (v *NullableCondAlarmAggregationAllOf) Set(val *CondAlarmAggregationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCondAlarmAggregationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCondAlarmAggregationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondAlarmAggregationAllOf(val *CondAlarmAggregationAllOf) *NullableCondAlarmAggregationAllOf {
	return &NullableCondAlarmAggregationAllOf{value: val, isSet: true}
}

func (v NullableCondAlarmAggregationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondAlarmAggregationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


