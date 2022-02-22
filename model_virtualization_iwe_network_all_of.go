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

// VirtualizationIweNetworkAllOf Definition of the list of properties defined in 'virtualization.IweNetwork', excluding properties defined in parent classes.
type VirtualizationIweNetworkAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A flag to distinguish if a network belongs to a HxAp infrastructure network or a user defined network that guest workloads can use.
	InfraNetwork *bool `json:"InfraNetwork,omitempty"`
	// The MTU size of the network.
	Mtu *int64 `json:"Mtu,omitempty"`
	// Network attachment type, only \"L2\" is available now. * `unknown` - This network is of an unknown network type. * `L2` - A Layer 2 switching network type.
	NetworkType *string `json:"NetworkType,omitempty"`
	Trunk []string `json:"Trunk,omitempty"`
	// A vlan id set on the network attachment point.
	Vlan *int64 `json:"Vlan,omitempty"`
	Cluster *VirtualizationIweClusterRelationship `json:"Cluster,omitempty"`
	Dvswitch *VirtualizationIweDvswitchRelationship `json:"Dvswitch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationIweNetworkAllOf VirtualizationIweNetworkAllOf

// NewVirtualizationIweNetworkAllOf instantiates a new VirtualizationIweNetworkAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationIweNetworkAllOf(classId string, objectType string) *VirtualizationIweNetworkAllOf {
	this := VirtualizationIweNetworkAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var networkType string = "unknown"
	this.NetworkType = &networkType
	return &this
}

// NewVirtualizationIweNetworkAllOfWithDefaults instantiates a new VirtualizationIweNetworkAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationIweNetworkAllOfWithDefaults() *VirtualizationIweNetworkAllOf {
	this := VirtualizationIweNetworkAllOf{}
	var classId string = "virtualization.IweNetwork"
	this.ClassId = classId
	var objectType string = "virtualization.IweNetwork"
	this.ObjectType = objectType
	var networkType string = "unknown"
	this.NetworkType = &networkType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationIweNetworkAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweNetworkAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationIweNetworkAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationIweNetworkAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweNetworkAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationIweNetworkAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInfraNetwork returns the InfraNetwork field value if set, zero value otherwise.
func (o *VirtualizationIweNetworkAllOf) GetInfraNetwork() bool {
	if o == nil || o.InfraNetwork == nil {
		var ret bool
		return ret
	}
	return *o.InfraNetwork
}

// GetInfraNetworkOk returns a tuple with the InfraNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweNetworkAllOf) GetInfraNetworkOk() (*bool, bool) {
	if o == nil || o.InfraNetwork == nil {
		return nil, false
	}
	return o.InfraNetwork, true
}

// HasInfraNetwork returns a boolean if a field has been set.
func (o *VirtualizationIweNetworkAllOf) HasInfraNetwork() bool {
	if o != nil && o.InfraNetwork != nil {
		return true
	}

	return false
}

// SetInfraNetwork gets a reference to the given bool and assigns it to the InfraNetwork field.
func (o *VirtualizationIweNetworkAllOf) SetInfraNetwork(v bool) {
	o.InfraNetwork = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *VirtualizationIweNetworkAllOf) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweNetworkAllOf) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *VirtualizationIweNetworkAllOf) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *VirtualizationIweNetworkAllOf) SetMtu(v int64) {
	o.Mtu = &v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *VirtualizationIweNetworkAllOf) GetNetworkType() string {
	if o == nil || o.NetworkType == nil {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweNetworkAllOf) GetNetworkTypeOk() (*string, bool) {
	if o == nil || o.NetworkType == nil {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *VirtualizationIweNetworkAllOf) HasNetworkType() bool {
	if o != nil && o.NetworkType != nil {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *VirtualizationIweNetworkAllOf) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetTrunk returns the Trunk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweNetworkAllOf) GetTrunk() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Trunk
}

// GetTrunkOk returns a tuple with the Trunk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweNetworkAllOf) GetTrunkOk() (*[]string, bool) {
	if o == nil || o.Trunk == nil {
		return nil, false
	}
	return &o.Trunk, true
}

// HasTrunk returns a boolean if a field has been set.
func (o *VirtualizationIweNetworkAllOf) HasTrunk() bool {
	if o != nil && o.Trunk != nil {
		return true
	}

	return false
}

// SetTrunk gets a reference to the given []string and assigns it to the Trunk field.
func (o *VirtualizationIweNetworkAllOf) SetTrunk(v []string) {
	o.Trunk = v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *VirtualizationIweNetworkAllOf) GetVlan() int64 {
	if o == nil || o.Vlan == nil {
		var ret int64
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweNetworkAllOf) GetVlanOk() (*int64, bool) {
	if o == nil || o.Vlan == nil {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *VirtualizationIweNetworkAllOf) HasVlan() bool {
	if o != nil && o.Vlan != nil {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int64 and assigns it to the Vlan field.
func (o *VirtualizationIweNetworkAllOf) SetVlan(v int64) {
	o.Vlan = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VirtualizationIweNetworkAllOf) GetCluster() VirtualizationIweClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret VirtualizationIweClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweNetworkAllOf) GetClusterOk() (*VirtualizationIweClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationIweNetworkAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given VirtualizationIweClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationIweNetworkAllOf) SetCluster(v VirtualizationIweClusterRelationship) {
	o.Cluster = &v
}

// GetDvswitch returns the Dvswitch field value if set, zero value otherwise.
func (o *VirtualizationIweNetworkAllOf) GetDvswitch() VirtualizationIweDvswitchRelationship {
	if o == nil || o.Dvswitch == nil {
		var ret VirtualizationIweDvswitchRelationship
		return ret
	}
	return *o.Dvswitch
}

// GetDvswitchOk returns a tuple with the Dvswitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweNetworkAllOf) GetDvswitchOk() (*VirtualizationIweDvswitchRelationship, bool) {
	if o == nil || o.Dvswitch == nil {
		return nil, false
	}
	return o.Dvswitch, true
}

// HasDvswitch returns a boolean if a field has been set.
func (o *VirtualizationIweNetworkAllOf) HasDvswitch() bool {
	if o != nil && o.Dvswitch != nil {
		return true
	}

	return false
}

// SetDvswitch gets a reference to the given VirtualizationIweDvswitchRelationship and assigns it to the Dvswitch field.
func (o *VirtualizationIweNetworkAllOf) SetDvswitch(v VirtualizationIweDvswitchRelationship) {
	o.Dvswitch = &v
}

func (o VirtualizationIweNetworkAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InfraNetwork != nil {
		toSerialize["InfraNetwork"] = o.InfraNetwork
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.NetworkType != nil {
		toSerialize["NetworkType"] = o.NetworkType
	}
	if o.Trunk != nil {
		toSerialize["Trunk"] = o.Trunk
	}
	if o.Vlan != nil {
		toSerialize["Vlan"] = o.Vlan
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Dvswitch != nil {
		toSerialize["Dvswitch"] = o.Dvswitch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationIweNetworkAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationIweNetworkAllOf := _VirtualizationIweNetworkAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationIweNetworkAllOf); err == nil {
		*o = VirtualizationIweNetworkAllOf(varVirtualizationIweNetworkAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InfraNetwork")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "NetworkType")
		delete(additionalProperties, "Trunk")
		delete(additionalProperties, "Vlan")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Dvswitch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationIweNetworkAllOf struct {
	value *VirtualizationIweNetworkAllOf
	isSet bool
}

func (v NullableVirtualizationIweNetworkAllOf) Get() *VirtualizationIweNetworkAllOf {
	return v.value
}

func (v *NullableVirtualizationIweNetworkAllOf) Set(val *VirtualizationIweNetworkAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationIweNetworkAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationIweNetworkAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationIweNetworkAllOf(val *VirtualizationIweNetworkAllOf) *NullableVirtualizationIweNetworkAllOf {
	return &NullableVirtualizationIweNetworkAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationIweNetworkAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationIweNetworkAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


