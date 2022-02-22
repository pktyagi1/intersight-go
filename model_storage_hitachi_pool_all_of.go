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

// StorageHitachiPoolAllOf Definition of the list of properties defined in 'storage.HitachiPool', excluding properties defined in parent classes.
type StorageHitachiPoolAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Setting the protection function for a virtual volume. When the DP pool is blockade, whether the read and write operations can be performed for the DP volume that uses the target DP pool is output. Yes, read and write operations are not possible. No, read and write operations are possible. -, Thin Image pool or not available.
	BlockingModeBlockade *string `json:"BlockingModeBlockade,omitempty"`
	// Setting the protection function for a virtual volume. When the DP pool is full, whether the read and write operations can be performed for the DP volume that uses the target DP pool is output. Yes, read and write operations are not possible. No, read and write operations are possible. -, Thin Image pool or not available.
	BlockingModeFull *string `json:"BlockingModeFull,omitempty"`
	// The depletion threshold set for the pool (%).
	DepletionThreshold *string `json:"DepletionThreshold,omitempty"`
	// Whether the pool is shrinking is output.
	IsShrinking *bool `json:"IsShrinking,omitempty"`
	// Performance monitoring execution mode (monitor mode). * `N/A` - Performance monitoring is not available. * `Period mode` - Period mode is the default setting. If Period mode is enabled, tier range values and page relocations are determined based solely on the monitoring data from the last complete cycle. * `Continuous mode` - When Continuous mode is enabled, the weighted average efficiency is calculated using the latest monitoring information and the collected monitoring information in the past cycles. Page relocations are determined using this weighted average efficiency.
	MonitoringMode *string `json:"MonitoringMode,omitempty"`
	// Status of monitor information.
	MonitoringStatus *string `json:"MonitoringStatus,omitempty"`
	// Execution mode for the pool. * `N/A` - Execution Mode is not available for the pool. * `Auto` - The mode in which the monitor is started or stopped at the specified time, and the Tier range is specified by automatic calculation of the DKC (specified by using Storage Navigator). * `Manual` - The mode in which the monitor is started or stopped by instructions from the REST API server, and the Tier range is specified by automatic calculation of the DKC.
	PoolActionMode *string `json:"PoolActionMode,omitempty"`
	// Displays the status of the tier relocation processing.
	ProgressOfReplacing *string `json:"ProgressOfReplacing,omitempty"`
	// Total capacity of the reserved page (bytes) of the DP volume that is related to the DP pool.
	TotalReservedCapacity *int64 `json:"TotalReservedCapacity,omitempty"`
	// The warning threshold set for the pool (%).
	WarningThreshold *int64 `json:"WarningThreshold,omitempty"`
	Array *StorageHitachiArrayRelationship `json:"Array,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiPoolAllOf StorageHitachiPoolAllOf

// NewStorageHitachiPoolAllOf instantiates a new StorageHitachiPoolAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiPoolAllOf(classId string, objectType string) *StorageHitachiPoolAllOf {
	this := StorageHitachiPoolAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiPoolAllOfWithDefaults instantiates a new StorageHitachiPoolAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiPoolAllOfWithDefaults() *StorageHitachiPoolAllOf {
	this := StorageHitachiPoolAllOf{}
	var classId string = "storage.HitachiPool"
	this.ClassId = classId
	var objectType string = "storage.HitachiPool"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiPoolAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiPoolAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiPoolAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiPoolAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiPoolAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiPoolAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBlockingModeBlockade returns the BlockingModeBlockade field value if set, zero value otherwise.
func (o *StorageHitachiPoolAllOf) GetBlockingModeBlockade() string {
	if o == nil || o.BlockingModeBlockade == nil {
		var ret string
		return ret
	}
	return *o.BlockingModeBlockade
}

// GetBlockingModeBlockadeOk returns a tuple with the BlockingModeBlockade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPoolAllOf) GetBlockingModeBlockadeOk() (*string, bool) {
	if o == nil || o.BlockingModeBlockade == nil {
		return nil, false
	}
	return o.BlockingModeBlockade, true
}

// HasBlockingModeBlockade returns a boolean if a field has been set.
func (o *StorageHitachiPoolAllOf) HasBlockingModeBlockade() bool {
	if o != nil && o.BlockingModeBlockade != nil {
		return true
	}

	return false
}

// SetBlockingModeBlockade gets a reference to the given string and assigns it to the BlockingModeBlockade field.
func (o *StorageHitachiPoolAllOf) SetBlockingModeBlockade(v string) {
	o.BlockingModeBlockade = &v
}

// GetBlockingModeFull returns the BlockingModeFull field value if set, zero value otherwise.
func (o *StorageHitachiPoolAllOf) GetBlockingModeFull() string {
	if o == nil || o.BlockingModeFull == nil {
		var ret string
		return ret
	}
	return *o.BlockingModeFull
}

// GetBlockingModeFullOk returns a tuple with the BlockingModeFull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPoolAllOf) GetBlockingModeFullOk() (*string, bool) {
	if o == nil || o.BlockingModeFull == nil {
		return nil, false
	}
	return o.BlockingModeFull, true
}

// HasBlockingModeFull returns a boolean if a field has been set.
func (o *StorageHitachiPoolAllOf) HasBlockingModeFull() bool {
	if o != nil && o.BlockingModeFull != nil {
		return true
	}

	return false
}

// SetBlockingModeFull gets a reference to the given string and assigns it to the BlockingModeFull field.
func (o *StorageHitachiPoolAllOf) SetBlockingModeFull(v string) {
	o.BlockingModeFull = &v
}

// GetDepletionThreshold returns the DepletionThreshold field value if set, zero value otherwise.
func (o *StorageHitachiPoolAllOf) GetDepletionThreshold() string {
	if o == nil || o.DepletionThreshold == nil {
		var ret string
		return ret
	}
	return *o.DepletionThreshold
}

// GetDepletionThresholdOk returns a tuple with the DepletionThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPoolAllOf) GetDepletionThresholdOk() (*string, bool) {
	if o == nil || o.DepletionThreshold == nil {
		return nil, false
	}
	return o.DepletionThreshold, true
}

// HasDepletionThreshold returns a boolean if a field has been set.
func (o *StorageHitachiPoolAllOf) HasDepletionThreshold() bool {
	if o != nil && o.DepletionThreshold != nil {
		return true
	}

	return false
}

// SetDepletionThreshold gets a reference to the given string and assigns it to the DepletionThreshold field.
func (o *StorageHitachiPoolAllOf) SetDepletionThreshold(v string) {
	o.DepletionThreshold = &v
}

// GetIsShrinking returns the IsShrinking field value if set, zero value otherwise.
func (o *StorageHitachiPoolAllOf) GetIsShrinking() bool {
	if o == nil || o.IsShrinking == nil {
		var ret bool
		return ret
	}
	return *o.IsShrinking
}

// GetIsShrinkingOk returns a tuple with the IsShrinking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPoolAllOf) GetIsShrinkingOk() (*bool, bool) {
	if o == nil || o.IsShrinking == nil {
		return nil, false
	}
	return o.IsShrinking, true
}

// HasIsShrinking returns a boolean if a field has been set.
func (o *StorageHitachiPoolAllOf) HasIsShrinking() bool {
	if o != nil && o.IsShrinking != nil {
		return true
	}

	return false
}

// SetIsShrinking gets a reference to the given bool and assigns it to the IsShrinking field.
func (o *StorageHitachiPoolAllOf) SetIsShrinking(v bool) {
	o.IsShrinking = &v
}

// GetMonitoringMode returns the MonitoringMode field value if set, zero value otherwise.
func (o *StorageHitachiPoolAllOf) GetMonitoringMode() string {
	if o == nil || o.MonitoringMode == nil {
		var ret string
		return ret
	}
	return *o.MonitoringMode
}

// GetMonitoringModeOk returns a tuple with the MonitoringMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPoolAllOf) GetMonitoringModeOk() (*string, bool) {
	if o == nil || o.MonitoringMode == nil {
		return nil, false
	}
	return o.MonitoringMode, true
}

// HasMonitoringMode returns a boolean if a field has been set.
func (o *StorageHitachiPoolAllOf) HasMonitoringMode() bool {
	if o != nil && o.MonitoringMode != nil {
		return true
	}

	return false
}

// SetMonitoringMode gets a reference to the given string and assigns it to the MonitoringMode field.
func (o *StorageHitachiPoolAllOf) SetMonitoringMode(v string) {
	o.MonitoringMode = &v
}

// GetMonitoringStatus returns the MonitoringStatus field value if set, zero value otherwise.
func (o *StorageHitachiPoolAllOf) GetMonitoringStatus() string {
	if o == nil || o.MonitoringStatus == nil {
		var ret string
		return ret
	}
	return *o.MonitoringStatus
}

// GetMonitoringStatusOk returns a tuple with the MonitoringStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPoolAllOf) GetMonitoringStatusOk() (*string, bool) {
	if o == nil || o.MonitoringStatus == nil {
		return nil, false
	}
	return o.MonitoringStatus, true
}

// HasMonitoringStatus returns a boolean if a field has been set.
func (o *StorageHitachiPoolAllOf) HasMonitoringStatus() bool {
	if o != nil && o.MonitoringStatus != nil {
		return true
	}

	return false
}

// SetMonitoringStatus gets a reference to the given string and assigns it to the MonitoringStatus field.
func (o *StorageHitachiPoolAllOf) SetMonitoringStatus(v string) {
	o.MonitoringStatus = &v
}

// GetPoolActionMode returns the PoolActionMode field value if set, zero value otherwise.
func (o *StorageHitachiPoolAllOf) GetPoolActionMode() string {
	if o == nil || o.PoolActionMode == nil {
		var ret string
		return ret
	}
	return *o.PoolActionMode
}

// GetPoolActionModeOk returns a tuple with the PoolActionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPoolAllOf) GetPoolActionModeOk() (*string, bool) {
	if o == nil || o.PoolActionMode == nil {
		return nil, false
	}
	return o.PoolActionMode, true
}

// HasPoolActionMode returns a boolean if a field has been set.
func (o *StorageHitachiPoolAllOf) HasPoolActionMode() bool {
	if o != nil && o.PoolActionMode != nil {
		return true
	}

	return false
}

// SetPoolActionMode gets a reference to the given string and assigns it to the PoolActionMode field.
func (o *StorageHitachiPoolAllOf) SetPoolActionMode(v string) {
	o.PoolActionMode = &v
}

// GetProgressOfReplacing returns the ProgressOfReplacing field value if set, zero value otherwise.
func (o *StorageHitachiPoolAllOf) GetProgressOfReplacing() string {
	if o == nil || o.ProgressOfReplacing == nil {
		var ret string
		return ret
	}
	return *o.ProgressOfReplacing
}

// GetProgressOfReplacingOk returns a tuple with the ProgressOfReplacing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPoolAllOf) GetProgressOfReplacingOk() (*string, bool) {
	if o == nil || o.ProgressOfReplacing == nil {
		return nil, false
	}
	return o.ProgressOfReplacing, true
}

// HasProgressOfReplacing returns a boolean if a field has been set.
func (o *StorageHitachiPoolAllOf) HasProgressOfReplacing() bool {
	if o != nil && o.ProgressOfReplacing != nil {
		return true
	}

	return false
}

// SetProgressOfReplacing gets a reference to the given string and assigns it to the ProgressOfReplacing field.
func (o *StorageHitachiPoolAllOf) SetProgressOfReplacing(v string) {
	o.ProgressOfReplacing = &v
}

// GetTotalReservedCapacity returns the TotalReservedCapacity field value if set, zero value otherwise.
func (o *StorageHitachiPoolAllOf) GetTotalReservedCapacity() int64 {
	if o == nil || o.TotalReservedCapacity == nil {
		var ret int64
		return ret
	}
	return *o.TotalReservedCapacity
}

// GetTotalReservedCapacityOk returns a tuple with the TotalReservedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPoolAllOf) GetTotalReservedCapacityOk() (*int64, bool) {
	if o == nil || o.TotalReservedCapacity == nil {
		return nil, false
	}
	return o.TotalReservedCapacity, true
}

// HasTotalReservedCapacity returns a boolean if a field has been set.
func (o *StorageHitachiPoolAllOf) HasTotalReservedCapacity() bool {
	if o != nil && o.TotalReservedCapacity != nil {
		return true
	}

	return false
}

// SetTotalReservedCapacity gets a reference to the given int64 and assigns it to the TotalReservedCapacity field.
func (o *StorageHitachiPoolAllOf) SetTotalReservedCapacity(v int64) {
	o.TotalReservedCapacity = &v
}

// GetWarningThreshold returns the WarningThreshold field value if set, zero value otherwise.
func (o *StorageHitachiPoolAllOf) GetWarningThreshold() int64 {
	if o == nil || o.WarningThreshold == nil {
		var ret int64
		return ret
	}
	return *o.WarningThreshold
}

// GetWarningThresholdOk returns a tuple with the WarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPoolAllOf) GetWarningThresholdOk() (*int64, bool) {
	if o == nil || o.WarningThreshold == nil {
		return nil, false
	}
	return o.WarningThreshold, true
}

// HasWarningThreshold returns a boolean if a field has been set.
func (o *StorageHitachiPoolAllOf) HasWarningThreshold() bool {
	if o != nil && o.WarningThreshold != nil {
		return true
	}

	return false
}

// SetWarningThreshold gets a reference to the given int64 and assigns it to the WarningThreshold field.
func (o *StorageHitachiPoolAllOf) SetWarningThreshold(v int64) {
	o.WarningThreshold = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageHitachiPoolAllOf) GetArray() StorageHitachiArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPoolAllOf) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiPoolAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiPoolAllOf) SetArray(v StorageHitachiArrayRelationship) {
	o.Array = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHitachiPoolAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPoolAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiPoolAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiPoolAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageHitachiPoolAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BlockingModeBlockade != nil {
		toSerialize["BlockingModeBlockade"] = o.BlockingModeBlockade
	}
	if o.BlockingModeFull != nil {
		toSerialize["BlockingModeFull"] = o.BlockingModeFull
	}
	if o.DepletionThreshold != nil {
		toSerialize["DepletionThreshold"] = o.DepletionThreshold
	}
	if o.IsShrinking != nil {
		toSerialize["IsShrinking"] = o.IsShrinking
	}
	if o.MonitoringMode != nil {
		toSerialize["MonitoringMode"] = o.MonitoringMode
	}
	if o.MonitoringStatus != nil {
		toSerialize["MonitoringStatus"] = o.MonitoringStatus
	}
	if o.PoolActionMode != nil {
		toSerialize["PoolActionMode"] = o.PoolActionMode
	}
	if o.ProgressOfReplacing != nil {
		toSerialize["ProgressOfReplacing"] = o.ProgressOfReplacing
	}
	if o.TotalReservedCapacity != nil {
		toSerialize["TotalReservedCapacity"] = o.TotalReservedCapacity
	}
	if o.WarningThreshold != nil {
		toSerialize["WarningThreshold"] = o.WarningThreshold
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHitachiPoolAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageHitachiPoolAllOf := _StorageHitachiPoolAllOf{}

	if err = json.Unmarshal(bytes, &varStorageHitachiPoolAllOf); err == nil {
		*o = StorageHitachiPoolAllOf(varStorageHitachiPoolAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BlockingModeBlockade")
		delete(additionalProperties, "BlockingModeFull")
		delete(additionalProperties, "DepletionThreshold")
		delete(additionalProperties, "IsShrinking")
		delete(additionalProperties, "MonitoringMode")
		delete(additionalProperties, "MonitoringStatus")
		delete(additionalProperties, "PoolActionMode")
		delete(additionalProperties, "ProgressOfReplacing")
		delete(additionalProperties, "TotalReservedCapacity")
		delete(additionalProperties, "WarningThreshold")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageHitachiPoolAllOf struct {
	value *StorageHitachiPoolAllOf
	isSet bool
}

func (v NullableStorageHitachiPoolAllOf) Get() *StorageHitachiPoolAllOf {
	return v.value
}

func (v *NullableStorageHitachiPoolAllOf) Set(val *StorageHitachiPoolAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiPoolAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiPoolAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiPoolAllOf(val *StorageHitachiPoolAllOf) *NullableStorageHitachiPoolAllOf {
	return &NullableStorageHitachiPoolAllOf{value: val, isSet: true}
}

func (v NullableStorageHitachiPoolAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiPoolAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


