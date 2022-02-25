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

// VirtualizationIweHostInterface A Intersight Workload Engine compute host interface entity that can be connected by IweHostVswitch.
type VirtualizationIweHostInterface struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	BondState NullableVirtualizationBondState `json:"BondState,omitempty"`
	// The UUID of the host to which this interface belongs to.
	HostName *string `json:"HostName,omitempty"`
	// The UUID of the host to which this interface belongs to.
	HostUuid *string `json:"HostUuid,omitempty"`
	// A hint of the interface type, such as \"ovs-bond\", \"device\", \"openvswitch\".
	IfType *string `json:"IfType,omitempty"`
	IpAddresses []string `json:"IpAddresses,omitempty"`
	// Link state information such as \"up\", \"down\". * `unknown` - The interface line is unknown. * `up` - The interface line is up. * `down` - The interface line is down. * `degraded` - For a bond/team interface, not all member interface is up.
	LinkState *string `json:"LinkState,omitempty"`
	// The MAC address of the interface.
	Mac *string `json:"Mac,omitempty"`
	// The MTU size of the interface.
	Mtu *int64 `json:"Mtu,omitempty"`
	// The name of the host to which this interface belongs to.
	Name *string `json:"Name,omitempty"`
	// A list of vlans allowed on this interface.
	Vlans *string `json:"Vlans,omitempty"`
	Cluster *VirtualizationIweClusterRelationship `json:"Cluster,omitempty"`
	DvUplink *VirtualizationIweDvUplinkRelationship `json:"DvUplink,omitempty"`
	Host *VirtualizationIweHostRelationship `json:"Host,omitempty"`
	Network *VirtualizationIweNetworkRelationship `json:"Network,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationIweHostInterface VirtualizationIweHostInterface

// NewVirtualizationIweHostInterface instantiates a new VirtualizationIweHostInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationIweHostInterface(classId string, objectType string) *VirtualizationIweHostInterface {
	this := VirtualizationIweHostInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	var linkState string = "unknown"
	this.LinkState = &linkState
	return &this
}

// NewVirtualizationIweHostInterfaceWithDefaults instantiates a new VirtualizationIweHostInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationIweHostInterfaceWithDefaults() *VirtualizationIweHostInterface {
	this := VirtualizationIweHostInterface{}
	var classId string = "virtualization.IweHostInterface"
	this.ClassId = classId
	var objectType string = "virtualization.IweHostInterface"
	this.ObjectType = objectType
	var linkState string = "unknown"
	this.LinkState = &linkState
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationIweHostInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostInterface) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationIweHostInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationIweHostInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationIweHostInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBondState returns the BondState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweHostInterface) GetBondState() VirtualizationBondState {
	if o == nil || o.BondState.Get() == nil {
		var ret VirtualizationBondState
		return ret
	}
	return *o.BondState.Get()
}

// GetBondStateOk returns a tuple with the BondState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweHostInterface) GetBondStateOk() (*VirtualizationBondState, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BondState.Get(), o.BondState.IsSet()
}

// HasBondState returns a boolean if a field has been set.
func (o *VirtualizationIweHostInterface) HasBondState() bool {
	if o != nil && o.BondState.IsSet() {
		return true
	}

	return false
}

// SetBondState gets a reference to the given NullableVirtualizationBondState and assigns it to the BondState field.
func (o *VirtualizationIweHostInterface) SetBondState(v VirtualizationBondState) {
	o.BondState.Set(&v)
}
// SetBondStateNil sets the value for BondState to be an explicit nil
func (o *VirtualizationIweHostInterface) SetBondStateNil() {
	o.BondState.Set(nil)
}

// UnsetBondState ensures that no value is present for BondState, not even an explicit nil
func (o *VirtualizationIweHostInterface) UnsetBondState() {
	o.BondState.Unset()
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *VirtualizationIweHostInterface) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostInterface) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *VirtualizationIweHostInterface) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *VirtualizationIweHostInterface) SetHostName(v string) {
	o.HostName = &v
}

// GetHostUuid returns the HostUuid field value if set, zero value otherwise.
func (o *VirtualizationIweHostInterface) GetHostUuid() string {
	if o == nil || o.HostUuid == nil {
		var ret string
		return ret
	}
	return *o.HostUuid
}

// GetHostUuidOk returns a tuple with the HostUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostInterface) GetHostUuidOk() (*string, bool) {
	if o == nil || o.HostUuid == nil {
		return nil, false
	}
	return o.HostUuid, true
}

// HasHostUuid returns a boolean if a field has been set.
func (o *VirtualizationIweHostInterface) HasHostUuid() bool {
	if o != nil && o.HostUuid != nil {
		return true
	}

	return false
}

// SetHostUuid gets a reference to the given string and assigns it to the HostUuid field.
func (o *VirtualizationIweHostInterface) SetHostUuid(v string) {
	o.HostUuid = &v
}

// GetIfType returns the IfType field value if set, zero value otherwise.
func (o *VirtualizationIweHostInterface) GetIfType() string {
	if o == nil || o.IfType == nil {
		var ret string
		return ret
	}
	return *o.IfType
}

// GetIfTypeOk returns a tuple with the IfType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostInterface) GetIfTypeOk() (*string, bool) {
	if o == nil || o.IfType == nil {
		return nil, false
	}
	return o.IfType, true
}

// HasIfType returns a boolean if a field has been set.
func (o *VirtualizationIweHostInterface) HasIfType() bool {
	if o != nil && o.IfType != nil {
		return true
	}

	return false
}

// SetIfType gets a reference to the given string and assigns it to the IfType field.
func (o *VirtualizationIweHostInterface) SetIfType(v string) {
	o.IfType = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweHostInterface) GetIpAddresses() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweHostInterface) GetIpAddressesOk() (*[]string, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return &o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *VirtualizationIweHostInterface) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *VirtualizationIweHostInterface) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetLinkState returns the LinkState field value if set, zero value otherwise.
func (o *VirtualizationIweHostInterface) GetLinkState() string {
	if o == nil || o.LinkState == nil {
		var ret string
		return ret
	}
	return *o.LinkState
}

// GetLinkStateOk returns a tuple with the LinkState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostInterface) GetLinkStateOk() (*string, bool) {
	if o == nil || o.LinkState == nil {
		return nil, false
	}
	return o.LinkState, true
}

// HasLinkState returns a boolean if a field has been set.
func (o *VirtualizationIweHostInterface) HasLinkState() bool {
	if o != nil && o.LinkState != nil {
		return true
	}

	return false
}

// SetLinkState gets a reference to the given string and assigns it to the LinkState field.
func (o *VirtualizationIweHostInterface) SetLinkState(v string) {
	o.LinkState = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *VirtualizationIweHostInterface) GetMac() string {
	if o == nil || o.Mac == nil {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostInterface) GetMacOk() (*string, bool) {
	if o == nil || o.Mac == nil {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *VirtualizationIweHostInterface) HasMac() bool {
	if o != nil && o.Mac != nil {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *VirtualizationIweHostInterface) SetMac(v string) {
	o.Mac = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *VirtualizationIweHostInterface) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostInterface) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *VirtualizationIweHostInterface) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *VirtualizationIweHostInterface) SetMtu(v int64) {
	o.Mtu = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationIweHostInterface) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostInterface) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationIweHostInterface) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationIweHostInterface) SetName(v string) {
	o.Name = &v
}

// GetVlans returns the Vlans field value if set, zero value otherwise.
func (o *VirtualizationIweHostInterface) GetVlans() string {
	if o == nil || o.Vlans == nil {
		var ret string
		return ret
	}
	return *o.Vlans
}

// GetVlansOk returns a tuple with the Vlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostInterface) GetVlansOk() (*string, bool) {
	if o == nil || o.Vlans == nil {
		return nil, false
	}
	return o.Vlans, true
}

// HasVlans returns a boolean if a field has been set.
func (o *VirtualizationIweHostInterface) HasVlans() bool {
	if o != nil && o.Vlans != nil {
		return true
	}

	return false
}

// SetVlans gets a reference to the given string and assigns it to the Vlans field.
func (o *VirtualizationIweHostInterface) SetVlans(v string) {
	o.Vlans = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VirtualizationIweHostInterface) GetCluster() VirtualizationIweClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret VirtualizationIweClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostInterface) GetClusterOk() (*VirtualizationIweClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationIweHostInterface) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given VirtualizationIweClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationIweHostInterface) SetCluster(v VirtualizationIweClusterRelationship) {
	o.Cluster = &v
}

// GetDvUplink returns the DvUplink field value if set, zero value otherwise.
func (o *VirtualizationIweHostInterface) GetDvUplink() VirtualizationIweDvUplinkRelationship {
	if o == nil || o.DvUplink == nil {
		var ret VirtualizationIweDvUplinkRelationship
		return ret
	}
	return *o.DvUplink
}

// GetDvUplinkOk returns a tuple with the DvUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostInterface) GetDvUplinkOk() (*VirtualizationIweDvUplinkRelationship, bool) {
	if o == nil || o.DvUplink == nil {
		return nil, false
	}
	return o.DvUplink, true
}

// HasDvUplink returns a boolean if a field has been set.
func (o *VirtualizationIweHostInterface) HasDvUplink() bool {
	if o != nil && o.DvUplink != nil {
		return true
	}

	return false
}

// SetDvUplink gets a reference to the given VirtualizationIweDvUplinkRelationship and assigns it to the DvUplink field.
func (o *VirtualizationIweHostInterface) SetDvUplink(v VirtualizationIweDvUplinkRelationship) {
	o.DvUplink = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *VirtualizationIweHostInterface) GetHost() VirtualizationIweHostRelationship {
	if o == nil || o.Host == nil {
		var ret VirtualizationIweHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostInterface) GetHostOk() (*VirtualizationIweHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *VirtualizationIweHostInterface) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given VirtualizationIweHostRelationship and assigns it to the Host field.
func (o *VirtualizationIweHostInterface) SetHost(v VirtualizationIweHostRelationship) {
	o.Host = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *VirtualizationIweHostInterface) GetNetwork() VirtualizationIweNetworkRelationship {
	if o == nil || o.Network == nil {
		var ret VirtualizationIweNetworkRelationship
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostInterface) GetNetworkOk() (*VirtualizationIweNetworkRelationship, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *VirtualizationIweHostInterface) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given VirtualizationIweNetworkRelationship and assigns it to the Network field.
func (o *VirtualizationIweHostInterface) SetNetwork(v VirtualizationIweNetworkRelationship) {
	o.Network = &v
}

func (o VirtualizationIweHostInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BondState.IsSet() {
		toSerialize["BondState"] = o.BondState.Get()
	}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}
	if o.HostUuid != nil {
		toSerialize["HostUuid"] = o.HostUuid
	}
	if o.IfType != nil {
		toSerialize["IfType"] = o.IfType
	}
	if o.IpAddresses != nil {
		toSerialize["IpAddresses"] = o.IpAddresses
	}
	if o.LinkState != nil {
		toSerialize["LinkState"] = o.LinkState
	}
	if o.Mac != nil {
		toSerialize["Mac"] = o.Mac
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Vlans != nil {
		toSerialize["Vlans"] = o.Vlans
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.DvUplink != nil {
		toSerialize["DvUplink"] = o.DvUplink
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}
	if o.Network != nil {
		toSerialize["Network"] = o.Network
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationIweHostInterface) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationIweHostInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		BondState NullableVirtualizationBondState `json:"BondState,omitempty"`
		// The UUID of the host to which this interface belongs to.
		HostName *string `json:"HostName,omitempty"`
		// The UUID of the host to which this interface belongs to.
		HostUuid *string `json:"HostUuid,omitempty"`
		// A hint of the interface type, such as \"ovs-bond\", \"device\", \"openvswitch\".
		IfType *string `json:"IfType,omitempty"`
		IpAddresses []string `json:"IpAddresses,omitempty"`
		// Link state information such as \"up\", \"down\". * `unknown` - The interface line is unknown. * `up` - The interface line is up. * `down` - The interface line is down. * `degraded` - For a bond/team interface, not all member interface is up.
		LinkState *string `json:"LinkState,omitempty"`
		// The MAC address of the interface.
		Mac *string `json:"Mac,omitempty"`
		// The MTU size of the interface.
		Mtu *int64 `json:"Mtu,omitempty"`
		// The name of the host to which this interface belongs to.
		Name *string `json:"Name,omitempty"`
		// A list of vlans allowed on this interface.
		Vlans *string `json:"Vlans,omitempty"`
		Cluster *VirtualizationIweClusterRelationship `json:"Cluster,omitempty"`
		DvUplink *VirtualizationIweDvUplinkRelationship `json:"DvUplink,omitempty"`
		Host *VirtualizationIweHostRelationship `json:"Host,omitempty"`
		Network *VirtualizationIweNetworkRelationship `json:"Network,omitempty"`
	}

	varVirtualizationIweHostInterfaceWithoutEmbeddedStruct := VirtualizationIweHostInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationIweHostInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationIweHostInterface := _VirtualizationIweHostInterface{}
		varVirtualizationIweHostInterface.ClassId = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.ClassId
		varVirtualizationIweHostInterface.ObjectType = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.ObjectType
		varVirtualizationIweHostInterface.BondState = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.BondState
		varVirtualizationIweHostInterface.HostName = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.HostName
		varVirtualizationIweHostInterface.HostUuid = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.HostUuid
		varVirtualizationIweHostInterface.IfType = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.IfType
		varVirtualizationIweHostInterface.IpAddresses = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.IpAddresses
		varVirtualizationIweHostInterface.LinkState = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.LinkState
		varVirtualizationIweHostInterface.Mac = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.Mac
		varVirtualizationIweHostInterface.Mtu = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.Mtu
		varVirtualizationIweHostInterface.Name = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.Name
		varVirtualizationIweHostInterface.Vlans = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.Vlans
		varVirtualizationIweHostInterface.Cluster = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.Cluster
		varVirtualizationIweHostInterface.DvUplink = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.DvUplink
		varVirtualizationIweHostInterface.Host = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.Host
		varVirtualizationIweHostInterface.Network = varVirtualizationIweHostInterfaceWithoutEmbeddedStruct.Network
		*o = VirtualizationIweHostInterface(varVirtualizationIweHostInterface)
	} else {
		return err
	}

	varVirtualizationIweHostInterface := _VirtualizationIweHostInterface{}

	err = json.Unmarshal(bytes, &varVirtualizationIweHostInterface)
	if err == nil {
		o.MoBaseMo = varVirtualizationIweHostInterface.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BondState")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "HostUuid")
		delete(additionalProperties, "IfType")
		delete(additionalProperties, "IpAddresses")
		delete(additionalProperties, "LinkState")
		delete(additionalProperties, "Mac")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Vlans")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "DvUplink")
		delete(additionalProperties, "Host")
		delete(additionalProperties, "Network")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableVirtualizationIweHostInterface struct {
	value *VirtualizationIweHostInterface
	isSet bool
}

func (v NullableVirtualizationIweHostInterface) Get() *VirtualizationIweHostInterface {
	return v.value
}

func (v *NullableVirtualizationIweHostInterface) Set(val *VirtualizationIweHostInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationIweHostInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationIweHostInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationIweHostInterface(val *VirtualizationIweHostInterface) *NullableVirtualizationIweHostInterface {
	return &NullableVirtualizationIweHostInterface{value: val, isSet: true}
}

func (v NullableVirtualizationIweHostInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationIweHostInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


