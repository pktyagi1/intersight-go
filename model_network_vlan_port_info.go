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

// NetworkVlanPortInfo Vlan Port information of a Fabric Interconnect.
type NetworkVlanPortInfo struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of available VLAN access ports on a Fabric Interconnect.
	AccessVlanPortCount *int64 `json:"AccessVlanPortCount,omitempty"`
	// The number of available VLAN border ports on a Fabric Interconnect.
	BorderVlanPortCount *int64 `json:"BorderVlanPortCount,omitempty"`
	// The number of compressed VLAN Group count on a Fabric Interconnect calculated by VLAN port group library.
	CompressedOptimizationSetsValue *int64 `json:"CompressedOptimizationSetsValue,omitempty"`
	// The number of compressed VLAN ports on a Fabric Interconnect.
	// Deprecated
	CompressedVlanPortCount *string `json:"CompressedVlanPortCount,omitempty"`
	// The number of compressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library.
	CompressedVlanPortCountValue *int64 `json:"CompressedVlanPortCountValue,omitempty"`
	// The total number of VLAN ports on a Fabric Interconnect.
	TotalVlanPortCount *int64 `json:"TotalVlanPortCount,omitempty"`
	// The number of uncompressed VLAN ports on a Fabric Interconnect.
	// Deprecated
	UncompressedVlanPortCount *string `json:"UncompressedVlanPortCount,omitempty"`
	// The number of uncompressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library.
	UncompressedVlanPortCountValue *int64 `json:"UncompressedVlanPortCountValue,omitempty"`
	// The maximum number of VLAN ports allowed on a Fabric Interconnect.
	VlanPortLimit *int64 `json:"VlanPortLimit,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	NetworkElement *NetworkElementRelationship `json:"NetworkElement,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkVlanPortInfo NetworkVlanPortInfo

// NewNetworkVlanPortInfo instantiates a new NetworkVlanPortInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkVlanPortInfo(classId string, objectType string) *NetworkVlanPortInfo {
	this := NetworkVlanPortInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkVlanPortInfoWithDefaults instantiates a new NetworkVlanPortInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkVlanPortInfoWithDefaults() *NetworkVlanPortInfo {
	this := NetworkVlanPortInfo{}
	var classId string = "network.VlanPortInfo"
	this.ClassId = classId
	var objectType string = "network.VlanPortInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkVlanPortInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfo) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkVlanPortInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkVlanPortInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkVlanPortInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessVlanPortCount returns the AccessVlanPortCount field value if set, zero value otherwise.
func (o *NetworkVlanPortInfo) GetAccessVlanPortCount() int64 {
	if o == nil || o.AccessVlanPortCount == nil {
		var ret int64
		return ret
	}
	return *o.AccessVlanPortCount
}

// GetAccessVlanPortCountOk returns a tuple with the AccessVlanPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfo) GetAccessVlanPortCountOk() (*int64, bool) {
	if o == nil || o.AccessVlanPortCount == nil {
		return nil, false
	}
	return o.AccessVlanPortCount, true
}

// HasAccessVlanPortCount returns a boolean if a field has been set.
func (o *NetworkVlanPortInfo) HasAccessVlanPortCount() bool {
	if o != nil && o.AccessVlanPortCount != nil {
		return true
	}

	return false
}

// SetAccessVlanPortCount gets a reference to the given int64 and assigns it to the AccessVlanPortCount field.
func (o *NetworkVlanPortInfo) SetAccessVlanPortCount(v int64) {
	o.AccessVlanPortCount = &v
}

// GetBorderVlanPortCount returns the BorderVlanPortCount field value if set, zero value otherwise.
func (o *NetworkVlanPortInfo) GetBorderVlanPortCount() int64 {
	if o == nil || o.BorderVlanPortCount == nil {
		var ret int64
		return ret
	}
	return *o.BorderVlanPortCount
}

// GetBorderVlanPortCountOk returns a tuple with the BorderVlanPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfo) GetBorderVlanPortCountOk() (*int64, bool) {
	if o == nil || o.BorderVlanPortCount == nil {
		return nil, false
	}
	return o.BorderVlanPortCount, true
}

// HasBorderVlanPortCount returns a boolean if a field has been set.
func (o *NetworkVlanPortInfo) HasBorderVlanPortCount() bool {
	if o != nil && o.BorderVlanPortCount != nil {
		return true
	}

	return false
}

// SetBorderVlanPortCount gets a reference to the given int64 and assigns it to the BorderVlanPortCount field.
func (o *NetworkVlanPortInfo) SetBorderVlanPortCount(v int64) {
	o.BorderVlanPortCount = &v
}

// GetCompressedOptimizationSetsValue returns the CompressedOptimizationSetsValue field value if set, zero value otherwise.
func (o *NetworkVlanPortInfo) GetCompressedOptimizationSetsValue() int64 {
	if o == nil || o.CompressedOptimizationSetsValue == nil {
		var ret int64
		return ret
	}
	return *o.CompressedOptimizationSetsValue
}

// GetCompressedOptimizationSetsValueOk returns a tuple with the CompressedOptimizationSetsValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfo) GetCompressedOptimizationSetsValueOk() (*int64, bool) {
	if o == nil || o.CompressedOptimizationSetsValue == nil {
		return nil, false
	}
	return o.CompressedOptimizationSetsValue, true
}

// HasCompressedOptimizationSetsValue returns a boolean if a field has been set.
func (o *NetworkVlanPortInfo) HasCompressedOptimizationSetsValue() bool {
	if o != nil && o.CompressedOptimizationSetsValue != nil {
		return true
	}

	return false
}

// SetCompressedOptimizationSetsValue gets a reference to the given int64 and assigns it to the CompressedOptimizationSetsValue field.
func (o *NetworkVlanPortInfo) SetCompressedOptimizationSetsValue(v int64) {
	o.CompressedOptimizationSetsValue = &v
}

// GetCompressedVlanPortCount returns the CompressedVlanPortCount field value if set, zero value otherwise.
// Deprecated
func (o *NetworkVlanPortInfo) GetCompressedVlanPortCount() string {
	if o == nil || o.CompressedVlanPortCount == nil {
		var ret string
		return ret
	}
	return *o.CompressedVlanPortCount
}

// GetCompressedVlanPortCountOk returns a tuple with the CompressedVlanPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *NetworkVlanPortInfo) GetCompressedVlanPortCountOk() (*string, bool) {
	if o == nil || o.CompressedVlanPortCount == nil {
		return nil, false
	}
	return o.CompressedVlanPortCount, true
}

// HasCompressedVlanPortCount returns a boolean if a field has been set.
func (o *NetworkVlanPortInfo) HasCompressedVlanPortCount() bool {
	if o != nil && o.CompressedVlanPortCount != nil {
		return true
	}

	return false
}

// SetCompressedVlanPortCount gets a reference to the given string and assigns it to the CompressedVlanPortCount field.
// Deprecated
func (o *NetworkVlanPortInfo) SetCompressedVlanPortCount(v string) {
	o.CompressedVlanPortCount = &v
}

// GetCompressedVlanPortCountValue returns the CompressedVlanPortCountValue field value if set, zero value otherwise.
func (o *NetworkVlanPortInfo) GetCompressedVlanPortCountValue() int64 {
	if o == nil || o.CompressedVlanPortCountValue == nil {
		var ret int64
		return ret
	}
	return *o.CompressedVlanPortCountValue
}

// GetCompressedVlanPortCountValueOk returns a tuple with the CompressedVlanPortCountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfo) GetCompressedVlanPortCountValueOk() (*int64, bool) {
	if o == nil || o.CompressedVlanPortCountValue == nil {
		return nil, false
	}
	return o.CompressedVlanPortCountValue, true
}

// HasCompressedVlanPortCountValue returns a boolean if a field has been set.
func (o *NetworkVlanPortInfo) HasCompressedVlanPortCountValue() bool {
	if o != nil && o.CompressedVlanPortCountValue != nil {
		return true
	}

	return false
}

// SetCompressedVlanPortCountValue gets a reference to the given int64 and assigns it to the CompressedVlanPortCountValue field.
func (o *NetworkVlanPortInfo) SetCompressedVlanPortCountValue(v int64) {
	o.CompressedVlanPortCountValue = &v
}

// GetTotalVlanPortCount returns the TotalVlanPortCount field value if set, zero value otherwise.
func (o *NetworkVlanPortInfo) GetTotalVlanPortCount() int64 {
	if o == nil || o.TotalVlanPortCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalVlanPortCount
}

// GetTotalVlanPortCountOk returns a tuple with the TotalVlanPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfo) GetTotalVlanPortCountOk() (*int64, bool) {
	if o == nil || o.TotalVlanPortCount == nil {
		return nil, false
	}
	return o.TotalVlanPortCount, true
}

// HasTotalVlanPortCount returns a boolean if a field has been set.
func (o *NetworkVlanPortInfo) HasTotalVlanPortCount() bool {
	if o != nil && o.TotalVlanPortCount != nil {
		return true
	}

	return false
}

// SetTotalVlanPortCount gets a reference to the given int64 and assigns it to the TotalVlanPortCount field.
func (o *NetworkVlanPortInfo) SetTotalVlanPortCount(v int64) {
	o.TotalVlanPortCount = &v
}

// GetUncompressedVlanPortCount returns the UncompressedVlanPortCount field value if set, zero value otherwise.
// Deprecated
func (o *NetworkVlanPortInfo) GetUncompressedVlanPortCount() string {
	if o == nil || o.UncompressedVlanPortCount == nil {
		var ret string
		return ret
	}
	return *o.UncompressedVlanPortCount
}

// GetUncompressedVlanPortCountOk returns a tuple with the UncompressedVlanPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *NetworkVlanPortInfo) GetUncompressedVlanPortCountOk() (*string, bool) {
	if o == nil || o.UncompressedVlanPortCount == nil {
		return nil, false
	}
	return o.UncompressedVlanPortCount, true
}

// HasUncompressedVlanPortCount returns a boolean if a field has been set.
func (o *NetworkVlanPortInfo) HasUncompressedVlanPortCount() bool {
	if o != nil && o.UncompressedVlanPortCount != nil {
		return true
	}

	return false
}

// SetUncompressedVlanPortCount gets a reference to the given string and assigns it to the UncompressedVlanPortCount field.
// Deprecated
func (o *NetworkVlanPortInfo) SetUncompressedVlanPortCount(v string) {
	o.UncompressedVlanPortCount = &v
}

// GetUncompressedVlanPortCountValue returns the UncompressedVlanPortCountValue field value if set, zero value otherwise.
func (o *NetworkVlanPortInfo) GetUncompressedVlanPortCountValue() int64 {
	if o == nil || o.UncompressedVlanPortCountValue == nil {
		var ret int64
		return ret
	}
	return *o.UncompressedVlanPortCountValue
}

// GetUncompressedVlanPortCountValueOk returns a tuple with the UncompressedVlanPortCountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfo) GetUncompressedVlanPortCountValueOk() (*int64, bool) {
	if o == nil || o.UncompressedVlanPortCountValue == nil {
		return nil, false
	}
	return o.UncompressedVlanPortCountValue, true
}

// HasUncompressedVlanPortCountValue returns a boolean if a field has been set.
func (o *NetworkVlanPortInfo) HasUncompressedVlanPortCountValue() bool {
	if o != nil && o.UncompressedVlanPortCountValue != nil {
		return true
	}

	return false
}

// SetUncompressedVlanPortCountValue gets a reference to the given int64 and assigns it to the UncompressedVlanPortCountValue field.
func (o *NetworkVlanPortInfo) SetUncompressedVlanPortCountValue(v int64) {
	o.UncompressedVlanPortCountValue = &v
}

// GetVlanPortLimit returns the VlanPortLimit field value if set, zero value otherwise.
func (o *NetworkVlanPortInfo) GetVlanPortLimit() int64 {
	if o == nil || o.VlanPortLimit == nil {
		var ret int64
		return ret
	}
	return *o.VlanPortLimit
}

// GetVlanPortLimitOk returns a tuple with the VlanPortLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfo) GetVlanPortLimitOk() (*int64, bool) {
	if o == nil || o.VlanPortLimit == nil {
		return nil, false
	}
	return o.VlanPortLimit, true
}

// HasVlanPortLimit returns a boolean if a field has been set.
func (o *NetworkVlanPortInfo) HasVlanPortLimit() bool {
	if o != nil && o.VlanPortLimit != nil {
		return true
	}

	return false
}

// SetVlanPortLimit gets a reference to the given int64 and assigns it to the VlanPortLimit field.
func (o *NetworkVlanPortInfo) SetVlanPortLimit(v int64) {
	o.VlanPortLimit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *NetworkVlanPortInfo) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfo) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *NetworkVlanPortInfo) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *NetworkVlanPortInfo) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *NetworkVlanPortInfo) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfo) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NetworkVlanPortInfo) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NetworkVlanPortInfo) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NetworkVlanPortInfo) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfo) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkVlanPortInfo) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkVlanPortInfo) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NetworkVlanPortInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccessVlanPortCount != nil {
		toSerialize["AccessVlanPortCount"] = o.AccessVlanPortCount
	}
	if o.BorderVlanPortCount != nil {
		toSerialize["BorderVlanPortCount"] = o.BorderVlanPortCount
	}
	if o.CompressedOptimizationSetsValue != nil {
		toSerialize["CompressedOptimizationSetsValue"] = o.CompressedOptimizationSetsValue
	}
	if o.CompressedVlanPortCount != nil {
		toSerialize["CompressedVlanPortCount"] = o.CompressedVlanPortCount
	}
	if o.CompressedVlanPortCountValue != nil {
		toSerialize["CompressedVlanPortCountValue"] = o.CompressedVlanPortCountValue
	}
	if o.TotalVlanPortCount != nil {
		toSerialize["TotalVlanPortCount"] = o.TotalVlanPortCount
	}
	if o.UncompressedVlanPortCount != nil {
		toSerialize["UncompressedVlanPortCount"] = o.UncompressedVlanPortCount
	}
	if o.UncompressedVlanPortCountValue != nil {
		toSerialize["UncompressedVlanPortCountValue"] = o.UncompressedVlanPortCountValue
	}
	if o.VlanPortLimit != nil {
		toSerialize["VlanPortLimit"] = o.VlanPortLimit
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkVlanPortInfo) UnmarshalJSON(bytes []byte) (err error) {
	type NetworkVlanPortInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The number of available VLAN access ports on a Fabric Interconnect.
		AccessVlanPortCount *int64 `json:"AccessVlanPortCount,omitempty"`
		// The number of available VLAN border ports on a Fabric Interconnect.
		BorderVlanPortCount *int64 `json:"BorderVlanPortCount,omitempty"`
		// The number of compressed VLAN Group count on a Fabric Interconnect calculated by VLAN port group library.
		CompressedOptimizationSetsValue *int64 `json:"CompressedOptimizationSetsValue,omitempty"`
		// The number of compressed VLAN ports on a Fabric Interconnect.
		// Deprecated
		CompressedVlanPortCount *string `json:"CompressedVlanPortCount,omitempty"`
		// The number of compressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library.
		CompressedVlanPortCountValue *int64 `json:"CompressedVlanPortCountValue,omitempty"`
		// The total number of VLAN ports on a Fabric Interconnect.
		TotalVlanPortCount *int64 `json:"TotalVlanPortCount,omitempty"`
		// The number of uncompressed VLAN ports on a Fabric Interconnect.
		// Deprecated
		UncompressedVlanPortCount *string `json:"UncompressedVlanPortCount,omitempty"`
		// The number of uncompressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library.
		UncompressedVlanPortCountValue *int64 `json:"UncompressedVlanPortCountValue,omitempty"`
		// The maximum number of VLAN ports allowed on a Fabric Interconnect.
		VlanPortLimit *int64 `json:"VlanPortLimit,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
		NetworkElement *NetworkElementRelationship `json:"NetworkElement,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNetworkVlanPortInfoWithoutEmbeddedStruct := NetworkVlanPortInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNetworkVlanPortInfoWithoutEmbeddedStruct)
	if err == nil {
		varNetworkVlanPortInfo := _NetworkVlanPortInfo{}
		varNetworkVlanPortInfo.ClassId = varNetworkVlanPortInfoWithoutEmbeddedStruct.ClassId
		varNetworkVlanPortInfo.ObjectType = varNetworkVlanPortInfoWithoutEmbeddedStruct.ObjectType
		varNetworkVlanPortInfo.AccessVlanPortCount = varNetworkVlanPortInfoWithoutEmbeddedStruct.AccessVlanPortCount
		varNetworkVlanPortInfo.BorderVlanPortCount = varNetworkVlanPortInfoWithoutEmbeddedStruct.BorderVlanPortCount
		varNetworkVlanPortInfo.CompressedOptimizationSetsValue = varNetworkVlanPortInfoWithoutEmbeddedStruct.CompressedOptimizationSetsValue
		varNetworkVlanPortInfo.CompressedVlanPortCount = varNetworkVlanPortInfoWithoutEmbeddedStruct.CompressedVlanPortCount
		varNetworkVlanPortInfo.CompressedVlanPortCountValue = varNetworkVlanPortInfoWithoutEmbeddedStruct.CompressedVlanPortCountValue
		varNetworkVlanPortInfo.TotalVlanPortCount = varNetworkVlanPortInfoWithoutEmbeddedStruct.TotalVlanPortCount
		varNetworkVlanPortInfo.UncompressedVlanPortCount = varNetworkVlanPortInfoWithoutEmbeddedStruct.UncompressedVlanPortCount
		varNetworkVlanPortInfo.UncompressedVlanPortCountValue = varNetworkVlanPortInfoWithoutEmbeddedStruct.UncompressedVlanPortCountValue
		varNetworkVlanPortInfo.VlanPortLimit = varNetworkVlanPortInfoWithoutEmbeddedStruct.VlanPortLimit
		varNetworkVlanPortInfo.InventoryDeviceInfo = varNetworkVlanPortInfoWithoutEmbeddedStruct.InventoryDeviceInfo
		varNetworkVlanPortInfo.NetworkElement = varNetworkVlanPortInfoWithoutEmbeddedStruct.NetworkElement
		varNetworkVlanPortInfo.RegisteredDevice = varNetworkVlanPortInfoWithoutEmbeddedStruct.RegisteredDevice
		*o = NetworkVlanPortInfo(varNetworkVlanPortInfo)
	} else {
		return err
	}

	varNetworkVlanPortInfo := _NetworkVlanPortInfo{}

	err = json.Unmarshal(bytes, &varNetworkVlanPortInfo)
	if err == nil {
		o.InventoryBase = varNetworkVlanPortInfo.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessVlanPortCount")
		delete(additionalProperties, "BorderVlanPortCount")
		delete(additionalProperties, "CompressedOptimizationSetsValue")
		delete(additionalProperties, "CompressedVlanPortCount")
		delete(additionalProperties, "CompressedVlanPortCountValue")
		delete(additionalProperties, "TotalVlanPortCount")
		delete(additionalProperties, "UncompressedVlanPortCount")
		delete(additionalProperties, "UncompressedVlanPortCountValue")
		delete(additionalProperties, "VlanPortLimit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableNetworkVlanPortInfo struct {
	value *NetworkVlanPortInfo
	isSet bool
}

func (v NullableNetworkVlanPortInfo) Get() *NetworkVlanPortInfo {
	return v.value
}

func (v *NullableNetworkVlanPortInfo) Set(val *NetworkVlanPortInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkVlanPortInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkVlanPortInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkVlanPortInfo(val *NetworkVlanPortInfo) *NullableNetworkVlanPortInfo {
	return &NullableNetworkVlanPortInfo{value: val, isSet: true}
}

func (v NullableNetworkVlanPortInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkVlanPortInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


