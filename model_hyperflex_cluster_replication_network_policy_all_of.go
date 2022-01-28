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

// HyperflexClusterReplicationNetworkPolicyAllOf Definition of the list of properties defined in 'hyperflex.ClusterReplicationNetworkPolicy', excluding properties defined in parent classes.
type HyperflexClusterReplicationNetworkPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Bandwidth for the Replication network in Mbps.
	ReplicationBandwidthMbps *int64 `json:"ReplicationBandwidthMbps,omitempty"`
	ReplicationIpranges []HyperflexIpAddrRange `json:"ReplicationIpranges,omitempty"`
	// MTU for the Replication network.
	ReplicationMtu *int64 `json:"ReplicationMtu,omitempty"`
	ReplicationVlan NullableHyperflexNamedVlan `json:"ReplicationVlan,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexClusterReplicationNetworkPolicyAllOf HyperflexClusterReplicationNetworkPolicyAllOf

// NewHyperflexClusterReplicationNetworkPolicyAllOf instantiates a new HyperflexClusterReplicationNetworkPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexClusterReplicationNetworkPolicyAllOf(classId string, objectType string) *HyperflexClusterReplicationNetworkPolicyAllOf {
	this := HyperflexClusterReplicationNetworkPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var replicationBandwidthMbps int64 = 0
	this.ReplicationBandwidthMbps = &replicationBandwidthMbps
	var replicationMtu int64 = 1500
	this.ReplicationMtu = &replicationMtu
	return &this
}

// NewHyperflexClusterReplicationNetworkPolicyAllOfWithDefaults instantiates a new HyperflexClusterReplicationNetworkPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterReplicationNetworkPolicyAllOfWithDefaults() *HyperflexClusterReplicationNetworkPolicyAllOf {
	this := HyperflexClusterReplicationNetworkPolicyAllOf{}
	var classId string = "hyperflex.ClusterReplicationNetworkPolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.ClusterReplicationNetworkPolicy"
	this.ObjectType = objectType
	var replicationBandwidthMbps int64 = 0
	this.ReplicationBandwidthMbps = &replicationBandwidthMbps
	var replicationMtu int64 = 1500
	this.ReplicationMtu = &replicationMtu
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetReplicationBandwidthMbps returns the ReplicationBandwidthMbps field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationBandwidthMbps() int64 {
	if o == nil || o.ReplicationBandwidthMbps == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationBandwidthMbps
}

// GetReplicationBandwidthMbpsOk returns a tuple with the ReplicationBandwidthMbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationBandwidthMbpsOk() (*int64, bool) {
	if o == nil || o.ReplicationBandwidthMbps == nil {
		return nil, false
	}
	return o.ReplicationBandwidthMbps, true
}

// HasReplicationBandwidthMbps returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) HasReplicationBandwidthMbps() bool {
	if o != nil && o.ReplicationBandwidthMbps != nil {
		return true
	}

	return false
}

// SetReplicationBandwidthMbps gets a reference to the given int64 and assigns it to the ReplicationBandwidthMbps field.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetReplicationBandwidthMbps(v int64) {
	o.ReplicationBandwidthMbps = &v
}

// GetReplicationIpranges returns the ReplicationIpranges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationIpranges() []HyperflexIpAddrRange {
	if o == nil  {
		var ret []HyperflexIpAddrRange
		return ret
	}
	return o.ReplicationIpranges
}

// GetReplicationIprangesOk returns a tuple with the ReplicationIpranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationIprangesOk() (*[]HyperflexIpAddrRange, bool) {
	if o == nil || o.ReplicationIpranges == nil {
		return nil, false
	}
	return &o.ReplicationIpranges, true
}

// HasReplicationIpranges returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) HasReplicationIpranges() bool {
	if o != nil && o.ReplicationIpranges != nil {
		return true
	}

	return false
}

// SetReplicationIpranges gets a reference to the given []HyperflexIpAddrRange and assigns it to the ReplicationIpranges field.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetReplicationIpranges(v []HyperflexIpAddrRange) {
	o.ReplicationIpranges = v
}

// GetReplicationMtu returns the ReplicationMtu field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationMtu() int64 {
	if o == nil || o.ReplicationMtu == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationMtu
}

// GetReplicationMtuOk returns a tuple with the ReplicationMtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationMtuOk() (*int64, bool) {
	if o == nil || o.ReplicationMtu == nil {
		return nil, false
	}
	return o.ReplicationMtu, true
}

// HasReplicationMtu returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) HasReplicationMtu() bool {
	if o != nil && o.ReplicationMtu != nil {
		return true
	}

	return false
}

// SetReplicationMtu gets a reference to the given int64 and assigns it to the ReplicationMtu field.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetReplicationMtu(v int64) {
	o.ReplicationMtu = &v
}

// GetReplicationVlan returns the ReplicationVlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationVlan() HyperflexNamedVlan {
	if o == nil || o.ReplicationVlan.Get() == nil {
		var ret HyperflexNamedVlan
		return ret
	}
	return *o.ReplicationVlan.Get()
}

// GetReplicationVlanOk returns a tuple with the ReplicationVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationVlanOk() (*HyperflexNamedVlan, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReplicationVlan.Get(), o.ReplicationVlan.IsSet()
}

// HasReplicationVlan returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) HasReplicationVlan() bool {
	if o != nil && o.ReplicationVlan.IsSet() {
		return true
	}

	return false
}

// SetReplicationVlan gets a reference to the given NullableHyperflexNamedVlan and assigns it to the ReplicationVlan field.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetReplicationVlan(v HyperflexNamedVlan) {
	o.ReplicationVlan.Set(&v)
}
// SetReplicationVlanNil sets the value for ReplicationVlan to be an explicit nil
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetReplicationVlanNil() {
	o.ReplicationVlan.Set(nil)
}

// UnsetReplicationVlan ensures that no value is present for ReplicationVlan, not even an explicit nil
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) UnsetReplicationVlan() {
	o.ReplicationVlan.Unset()
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil  {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexClusterReplicationNetworkPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ReplicationBandwidthMbps != nil {
		toSerialize["ReplicationBandwidthMbps"] = o.ReplicationBandwidthMbps
	}
	if o.ReplicationIpranges != nil {
		toSerialize["ReplicationIpranges"] = o.ReplicationIpranges
	}
	if o.ReplicationMtu != nil {
		toSerialize["ReplicationMtu"] = o.ReplicationMtu
	}
	if o.ReplicationVlan.IsSet() {
		toSerialize["ReplicationVlan"] = o.ReplicationVlan.Get()
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexClusterReplicationNetworkPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexClusterReplicationNetworkPolicyAllOf := _HyperflexClusterReplicationNetworkPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexClusterReplicationNetworkPolicyAllOf); err == nil {
		*o = HyperflexClusterReplicationNetworkPolicyAllOf(varHyperflexClusterReplicationNetworkPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ReplicationBandwidthMbps")
		delete(additionalProperties, "ReplicationIpranges")
		delete(additionalProperties, "ReplicationMtu")
		delete(additionalProperties, "ReplicationVlan")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexClusterReplicationNetworkPolicyAllOf struct {
	value *HyperflexClusterReplicationNetworkPolicyAllOf
	isSet bool
}

func (v NullableHyperflexClusterReplicationNetworkPolicyAllOf) Get() *HyperflexClusterReplicationNetworkPolicyAllOf {
	return v.value
}

func (v *NullableHyperflexClusterReplicationNetworkPolicyAllOf) Set(val *HyperflexClusterReplicationNetworkPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterReplicationNetworkPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterReplicationNetworkPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterReplicationNetworkPolicyAllOf(val *HyperflexClusterReplicationNetworkPolicyAllOf) *NullableHyperflexClusterReplicationNetworkPolicyAllOf {
	return &NullableHyperflexClusterReplicationNetworkPolicyAllOf{value: val, isSet: true}
}

func (v NullableHyperflexClusterReplicationNetworkPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterReplicationNetworkPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


