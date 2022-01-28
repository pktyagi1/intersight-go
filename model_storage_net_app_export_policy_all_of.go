/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// StorageNetAppExportPolicyAllOf Definition of the list of properties defined in 'storage.NetAppExportPolicy', excluding properties defined in parent classes.
type StorageNetAppExportPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique identity of the device.
	ClusterUuid *string `json:"ClusterUuid,omitempty"`
	NetAppExportPolicyRule []StorageNetAppExportPolicyRule `json:"NetAppExportPolicyRule,omitempty"`
	// ID for the Export Policy.
	PolicyId *int64 `json:"PolicyId,omitempty"`
	Array *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	Tenant *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppExportPolicyAllOf StorageNetAppExportPolicyAllOf

// NewStorageNetAppExportPolicyAllOf instantiates a new StorageNetAppExportPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppExportPolicyAllOf(classId string, objectType string) *StorageNetAppExportPolicyAllOf {
	this := StorageNetAppExportPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppExportPolicyAllOfWithDefaults instantiates a new StorageNetAppExportPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppExportPolicyAllOfWithDefaults() *StorageNetAppExportPolicyAllOf {
	this := StorageNetAppExportPolicyAllOf{}
	var classId string = "storage.NetAppExportPolicy"
	this.ClassId = classId
	var objectType string = "storage.NetAppExportPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppExportPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppExportPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppExportPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppExportPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *StorageNetAppExportPolicyAllOf) GetClusterUuid() string {
	if o == nil || o.ClusterUuid == nil {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicyAllOf) GetClusterUuidOk() (*string, bool) {
	if o == nil || o.ClusterUuid == nil {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicyAllOf) HasClusterUuid() bool {
	if o != nil && o.ClusterUuid != nil {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *StorageNetAppExportPolicyAllOf) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetNetAppExportPolicyRule returns the NetAppExportPolicyRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppExportPolicyAllOf) GetNetAppExportPolicyRule() []StorageNetAppExportPolicyRule {
	if o == nil  {
		var ret []StorageNetAppExportPolicyRule
		return ret
	}
	return o.NetAppExportPolicyRule
}

// GetNetAppExportPolicyRuleOk returns a tuple with the NetAppExportPolicyRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppExportPolicyAllOf) GetNetAppExportPolicyRuleOk() (*[]StorageNetAppExportPolicyRule, bool) {
	if o == nil || o.NetAppExportPolicyRule == nil {
		return nil, false
	}
	return &o.NetAppExportPolicyRule, true
}

// HasNetAppExportPolicyRule returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicyAllOf) HasNetAppExportPolicyRule() bool {
	if o != nil && o.NetAppExportPolicyRule != nil {
		return true
	}

	return false
}

// SetNetAppExportPolicyRule gets a reference to the given []StorageNetAppExportPolicyRule and assigns it to the NetAppExportPolicyRule field.
func (o *StorageNetAppExportPolicyAllOf) SetNetAppExportPolicyRule(v []StorageNetAppExportPolicyRule) {
	o.NetAppExportPolicyRule = v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *StorageNetAppExportPolicyAllOf) GetPolicyId() int64 {
	if o == nil || o.PolicyId == nil {
		var ret int64
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicyAllOf) GetPolicyIdOk() (*int64, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicyAllOf) HasPolicyId() bool {
	if o != nil && o.PolicyId != nil {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given int64 and assigns it to the PolicyId field.
func (o *StorageNetAppExportPolicyAllOf) SetPolicyId(v int64) {
	o.PolicyId = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppExportPolicyAllOf) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicyAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicyAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppExportPolicyAllOf) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppExportPolicyAllOf) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicyAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicyAllOf) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppExportPolicyAllOf) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppExportPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterUuid != nil {
		toSerialize["ClusterUuid"] = o.ClusterUuid
	}
	if o.NetAppExportPolicyRule != nil {
		toSerialize["NetAppExportPolicyRule"] = o.NetAppExportPolicyRule
	}
	if o.PolicyId != nil {
		toSerialize["PolicyId"] = o.PolicyId
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppExportPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppExportPolicyAllOf := _StorageNetAppExportPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppExportPolicyAllOf); err == nil {
		*o = StorageNetAppExportPolicyAllOf(varStorageNetAppExportPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterUuid")
		delete(additionalProperties, "NetAppExportPolicyRule")
		delete(additionalProperties, "PolicyId")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "Tenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppExportPolicyAllOf struct {
	value *StorageNetAppExportPolicyAllOf
	isSet bool
}

func (v NullableStorageNetAppExportPolicyAllOf) Get() *StorageNetAppExportPolicyAllOf {
	return v.value
}

func (v *NullableStorageNetAppExportPolicyAllOf) Set(val *StorageNetAppExportPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppExportPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppExportPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppExportPolicyAllOf(val *StorageNetAppExportPolicyAllOf) *NullableStorageNetAppExportPolicyAllOf {
	return &NullableStorageNetAppExportPolicyAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppExportPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppExportPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


