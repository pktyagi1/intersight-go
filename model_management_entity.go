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

// ManagementEntity Logical representation that captures the role of each Fabric Interconnect in UCS Manager.
type ManagementEntity struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Cluster link state between the Fabric Interconnects.
	ClusterLinkState *string `json:"ClusterLinkState,omitempty"`
	// Cluster readiness of the Fabric Interconnect.
	ClusterReadiness *string `json:"ClusterReadiness,omitempty"`
	// Cluster state of the Fabric Interconnect.
	ClusterState *string `json:"ClusterState,omitempty"`
	// Identity of the Fabric Interconnect - A/B.
	EntityId *string `json:"EntityId,omitempty"`
	// Role (Primary / Subordinate) of the Fabric Interconnect.
	Leadership *string `json:"Leadership,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	NetworkElement *NetworkElementRelationship `json:"NetworkElement,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManagementEntity ManagementEntity

// NewManagementEntity instantiates a new ManagementEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementEntity(classId string, objectType string) *ManagementEntity {
	this := ManagementEntity{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewManagementEntityWithDefaults instantiates a new ManagementEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementEntityWithDefaults() *ManagementEntity {
	this := ManagementEntity{}
	var classId string = "management.Entity"
	this.ClassId = classId
	var objectType string = "management.Entity"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ManagementEntity) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ManagementEntity) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ManagementEntity) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ManagementEntity) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterLinkState returns the ClusterLinkState field value if set, zero value otherwise.
func (o *ManagementEntity) GetClusterLinkState() string {
	if o == nil || o.ClusterLinkState == nil {
		var ret string
		return ret
	}
	return *o.ClusterLinkState
}

// GetClusterLinkStateOk returns a tuple with the ClusterLinkState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetClusterLinkStateOk() (*string, bool) {
	if o == nil || o.ClusterLinkState == nil {
		return nil, false
	}
	return o.ClusterLinkState, true
}

// HasClusterLinkState returns a boolean if a field has been set.
func (o *ManagementEntity) HasClusterLinkState() bool {
	if o != nil && o.ClusterLinkState != nil {
		return true
	}

	return false
}

// SetClusterLinkState gets a reference to the given string and assigns it to the ClusterLinkState field.
func (o *ManagementEntity) SetClusterLinkState(v string) {
	o.ClusterLinkState = &v
}

// GetClusterReadiness returns the ClusterReadiness field value if set, zero value otherwise.
func (o *ManagementEntity) GetClusterReadiness() string {
	if o == nil || o.ClusterReadiness == nil {
		var ret string
		return ret
	}
	return *o.ClusterReadiness
}

// GetClusterReadinessOk returns a tuple with the ClusterReadiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetClusterReadinessOk() (*string, bool) {
	if o == nil || o.ClusterReadiness == nil {
		return nil, false
	}
	return o.ClusterReadiness, true
}

// HasClusterReadiness returns a boolean if a field has been set.
func (o *ManagementEntity) HasClusterReadiness() bool {
	if o != nil && o.ClusterReadiness != nil {
		return true
	}

	return false
}

// SetClusterReadiness gets a reference to the given string and assigns it to the ClusterReadiness field.
func (o *ManagementEntity) SetClusterReadiness(v string) {
	o.ClusterReadiness = &v
}

// GetClusterState returns the ClusterState field value if set, zero value otherwise.
func (o *ManagementEntity) GetClusterState() string {
	if o == nil || o.ClusterState == nil {
		var ret string
		return ret
	}
	return *o.ClusterState
}

// GetClusterStateOk returns a tuple with the ClusterState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetClusterStateOk() (*string, bool) {
	if o == nil || o.ClusterState == nil {
		return nil, false
	}
	return o.ClusterState, true
}

// HasClusterState returns a boolean if a field has been set.
func (o *ManagementEntity) HasClusterState() bool {
	if o != nil && o.ClusterState != nil {
		return true
	}

	return false
}

// SetClusterState gets a reference to the given string and assigns it to the ClusterState field.
func (o *ManagementEntity) SetClusterState(v string) {
	o.ClusterState = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *ManagementEntity) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *ManagementEntity) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *ManagementEntity) SetEntityId(v string) {
	o.EntityId = &v
}

// GetLeadership returns the Leadership field value if set, zero value otherwise.
func (o *ManagementEntity) GetLeadership() string {
	if o == nil || o.Leadership == nil {
		var ret string
		return ret
	}
	return *o.Leadership
}

// GetLeadershipOk returns a tuple with the Leadership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetLeadershipOk() (*string, bool) {
	if o == nil || o.Leadership == nil {
		return nil, false
	}
	return o.Leadership, true
}

// HasLeadership returns a boolean if a field has been set.
func (o *ManagementEntity) HasLeadership() bool {
	if o != nil && o.Leadership != nil {
		return true
	}

	return false
}

// SetLeadership gets a reference to the given string and assigns it to the Leadership field.
func (o *ManagementEntity) SetLeadership(v string) {
	o.Leadership = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *ManagementEntity) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *ManagementEntity) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *ManagementEntity) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *ManagementEntity) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *ManagementEntity) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *ManagementEntity) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ManagementEntity) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ManagementEntity) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ManagementEntity) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o ManagementEntity) MarshalJSON() ([]byte, error) {
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
	if o.ClusterLinkState != nil {
		toSerialize["ClusterLinkState"] = o.ClusterLinkState
	}
	if o.ClusterReadiness != nil {
		toSerialize["ClusterReadiness"] = o.ClusterReadiness
	}
	if o.ClusterState != nil {
		toSerialize["ClusterState"] = o.ClusterState
	}
	if o.EntityId != nil {
		toSerialize["EntityId"] = o.EntityId
	}
	if o.Leadership != nil {
		toSerialize["Leadership"] = o.Leadership
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

func (o *ManagementEntity) UnmarshalJSON(bytes []byte) (err error) {
	type ManagementEntityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Cluster link state between the Fabric Interconnects.
		ClusterLinkState *string `json:"ClusterLinkState,omitempty"`
		// Cluster readiness of the Fabric Interconnect.
		ClusterReadiness *string `json:"ClusterReadiness,omitempty"`
		// Cluster state of the Fabric Interconnect.
		ClusterState *string `json:"ClusterState,omitempty"`
		// Identity of the Fabric Interconnect - A/B.
		EntityId *string `json:"EntityId,omitempty"`
		// Role (Primary / Subordinate) of the Fabric Interconnect.
		Leadership *string `json:"Leadership,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
		NetworkElement *NetworkElementRelationship `json:"NetworkElement,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varManagementEntityWithoutEmbeddedStruct := ManagementEntityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varManagementEntityWithoutEmbeddedStruct)
	if err == nil {
		varManagementEntity := _ManagementEntity{}
		varManagementEntity.ClassId = varManagementEntityWithoutEmbeddedStruct.ClassId
		varManagementEntity.ObjectType = varManagementEntityWithoutEmbeddedStruct.ObjectType
		varManagementEntity.ClusterLinkState = varManagementEntityWithoutEmbeddedStruct.ClusterLinkState
		varManagementEntity.ClusterReadiness = varManagementEntityWithoutEmbeddedStruct.ClusterReadiness
		varManagementEntity.ClusterState = varManagementEntityWithoutEmbeddedStruct.ClusterState
		varManagementEntity.EntityId = varManagementEntityWithoutEmbeddedStruct.EntityId
		varManagementEntity.Leadership = varManagementEntityWithoutEmbeddedStruct.Leadership
		varManagementEntity.InventoryDeviceInfo = varManagementEntityWithoutEmbeddedStruct.InventoryDeviceInfo
		varManagementEntity.NetworkElement = varManagementEntityWithoutEmbeddedStruct.NetworkElement
		varManagementEntity.RegisteredDevice = varManagementEntityWithoutEmbeddedStruct.RegisteredDevice
		*o = ManagementEntity(varManagementEntity)
	} else {
		return err
	}

	varManagementEntity := _ManagementEntity{}

	err = json.Unmarshal(bytes, &varManagementEntity)
	if err == nil {
		o.InventoryBase = varManagementEntity.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterLinkState")
		delete(additionalProperties, "ClusterReadiness")
		delete(additionalProperties, "ClusterState")
		delete(additionalProperties, "EntityId")
		delete(additionalProperties, "Leadership")
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

type NullableManagementEntity struct {
	value *ManagementEntity
	isSet bool
}

func (v NullableManagementEntity) Get() *ManagementEntity {
	return v.value
}

func (v *NullableManagementEntity) Set(val *ManagementEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementEntity(val *ManagementEntity) *NullableManagementEntity {
	return &NullableManagementEntity{value: val, isSet: true}
}

func (v NullableManagementEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


