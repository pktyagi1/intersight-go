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

// EtherPhysicalPortBaseAllOf Definition of the list of properties defined in 'ether.PhysicalPortBase', excluding properties defined in parent classes.
type EtherPhysicalPortBaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Mac Address of a port in the Fabric Interconnect.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Operating mode of this port.
	Mode *string `json:"Mode,omitempty"`
	// Current Operational speed for this port.
	OperSpeed *string `json:"OperSpeed,omitempty"`
	// PeerDn for ethernet physical port.
	PeerDn *string `json:"PeerDn,omitempty"`
	// Port channel id for port channel created on FI switch.
	PortChannelId *int64 `json:"PortChannelId,omitempty"`
	// Defines the transport type for this port (ethernet OR fc).
	PortType *string `json:"PortType,omitempty"`
	// Transceiver model attached to a port in the Fabric Interconnect.
	TransceiverType *string `json:"TransceiverType,omitempty"`
	AcknowledgedPeerInterface *PortInterfaceBaseRelationship `json:"AcknowledgedPeerInterface,omitempty"`
	PeerInterface *PortInterfaceBaseRelationship `json:"PeerInterface,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EtherPhysicalPortBaseAllOf EtherPhysicalPortBaseAllOf

// NewEtherPhysicalPortBaseAllOf instantiates a new EtherPhysicalPortBaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEtherPhysicalPortBaseAllOf(classId string, objectType string) *EtherPhysicalPortBaseAllOf {
	this := EtherPhysicalPortBaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEtherPhysicalPortBaseAllOfWithDefaults instantiates a new EtherPhysicalPortBaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEtherPhysicalPortBaseAllOfWithDefaults() *EtherPhysicalPortBaseAllOf {
	this := EtherPhysicalPortBaseAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *EtherPhysicalPortBaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EtherPhysicalPortBaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EtherPhysicalPortBaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EtherPhysicalPortBaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *EtherPhysicalPortBaseAllOf) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBaseAllOf) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *EtherPhysicalPortBaseAllOf) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *EtherPhysicalPortBaseAllOf) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *EtherPhysicalPortBaseAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBaseAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *EtherPhysicalPortBaseAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *EtherPhysicalPortBaseAllOf) SetMode(v string) {
	o.Mode = &v
}

// GetOperSpeed returns the OperSpeed field value if set, zero value otherwise.
func (o *EtherPhysicalPortBaseAllOf) GetOperSpeed() string {
	if o == nil || o.OperSpeed == nil {
		var ret string
		return ret
	}
	return *o.OperSpeed
}

// GetOperSpeedOk returns a tuple with the OperSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBaseAllOf) GetOperSpeedOk() (*string, bool) {
	if o == nil || o.OperSpeed == nil {
		return nil, false
	}
	return o.OperSpeed, true
}

// HasOperSpeed returns a boolean if a field has been set.
func (o *EtherPhysicalPortBaseAllOf) HasOperSpeed() bool {
	if o != nil && o.OperSpeed != nil {
		return true
	}

	return false
}

// SetOperSpeed gets a reference to the given string and assigns it to the OperSpeed field.
func (o *EtherPhysicalPortBaseAllOf) SetOperSpeed(v string) {
	o.OperSpeed = &v
}

// GetPeerDn returns the PeerDn field value if set, zero value otherwise.
func (o *EtherPhysicalPortBaseAllOf) GetPeerDn() string {
	if o == nil || o.PeerDn == nil {
		var ret string
		return ret
	}
	return *o.PeerDn
}

// GetPeerDnOk returns a tuple with the PeerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBaseAllOf) GetPeerDnOk() (*string, bool) {
	if o == nil || o.PeerDn == nil {
		return nil, false
	}
	return o.PeerDn, true
}

// HasPeerDn returns a boolean if a field has been set.
func (o *EtherPhysicalPortBaseAllOf) HasPeerDn() bool {
	if o != nil && o.PeerDn != nil {
		return true
	}

	return false
}

// SetPeerDn gets a reference to the given string and assigns it to the PeerDn field.
func (o *EtherPhysicalPortBaseAllOf) SetPeerDn(v string) {
	o.PeerDn = &v
}

// GetPortChannelId returns the PortChannelId field value if set, zero value otherwise.
func (o *EtherPhysicalPortBaseAllOf) GetPortChannelId() int64 {
	if o == nil || o.PortChannelId == nil {
		var ret int64
		return ret
	}
	return *o.PortChannelId
}

// GetPortChannelIdOk returns a tuple with the PortChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBaseAllOf) GetPortChannelIdOk() (*int64, bool) {
	if o == nil || o.PortChannelId == nil {
		return nil, false
	}
	return o.PortChannelId, true
}

// HasPortChannelId returns a boolean if a field has been set.
func (o *EtherPhysicalPortBaseAllOf) HasPortChannelId() bool {
	if o != nil && o.PortChannelId != nil {
		return true
	}

	return false
}

// SetPortChannelId gets a reference to the given int64 and assigns it to the PortChannelId field.
func (o *EtherPhysicalPortBaseAllOf) SetPortChannelId(v int64) {
	o.PortChannelId = &v
}

// GetPortType returns the PortType field value if set, zero value otherwise.
func (o *EtherPhysicalPortBaseAllOf) GetPortType() string {
	if o == nil || o.PortType == nil {
		var ret string
		return ret
	}
	return *o.PortType
}

// GetPortTypeOk returns a tuple with the PortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBaseAllOf) GetPortTypeOk() (*string, bool) {
	if o == nil || o.PortType == nil {
		return nil, false
	}
	return o.PortType, true
}

// HasPortType returns a boolean if a field has been set.
func (o *EtherPhysicalPortBaseAllOf) HasPortType() bool {
	if o != nil && o.PortType != nil {
		return true
	}

	return false
}

// SetPortType gets a reference to the given string and assigns it to the PortType field.
func (o *EtherPhysicalPortBaseAllOf) SetPortType(v string) {
	o.PortType = &v
}

// GetTransceiverType returns the TransceiverType field value if set, zero value otherwise.
func (o *EtherPhysicalPortBaseAllOf) GetTransceiverType() string {
	if o == nil || o.TransceiverType == nil {
		var ret string
		return ret
	}
	return *o.TransceiverType
}

// GetTransceiverTypeOk returns a tuple with the TransceiverType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBaseAllOf) GetTransceiverTypeOk() (*string, bool) {
	if o == nil || o.TransceiverType == nil {
		return nil, false
	}
	return o.TransceiverType, true
}

// HasTransceiverType returns a boolean if a field has been set.
func (o *EtherPhysicalPortBaseAllOf) HasTransceiverType() bool {
	if o != nil && o.TransceiverType != nil {
		return true
	}

	return false
}

// SetTransceiverType gets a reference to the given string and assigns it to the TransceiverType field.
func (o *EtherPhysicalPortBaseAllOf) SetTransceiverType(v string) {
	o.TransceiverType = &v
}

// GetAcknowledgedPeerInterface returns the AcknowledgedPeerInterface field value if set, zero value otherwise.
func (o *EtherPhysicalPortBaseAllOf) GetAcknowledgedPeerInterface() PortInterfaceBaseRelationship {
	if o == nil || o.AcknowledgedPeerInterface == nil {
		var ret PortInterfaceBaseRelationship
		return ret
	}
	return *o.AcknowledgedPeerInterface
}

// GetAcknowledgedPeerInterfaceOk returns a tuple with the AcknowledgedPeerInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBaseAllOf) GetAcknowledgedPeerInterfaceOk() (*PortInterfaceBaseRelationship, bool) {
	if o == nil || o.AcknowledgedPeerInterface == nil {
		return nil, false
	}
	return o.AcknowledgedPeerInterface, true
}

// HasAcknowledgedPeerInterface returns a boolean if a field has been set.
func (o *EtherPhysicalPortBaseAllOf) HasAcknowledgedPeerInterface() bool {
	if o != nil && o.AcknowledgedPeerInterface != nil {
		return true
	}

	return false
}

// SetAcknowledgedPeerInterface gets a reference to the given PortInterfaceBaseRelationship and assigns it to the AcknowledgedPeerInterface field.
func (o *EtherPhysicalPortBaseAllOf) SetAcknowledgedPeerInterface(v PortInterfaceBaseRelationship) {
	o.AcknowledgedPeerInterface = &v
}

// GetPeerInterface returns the PeerInterface field value if set, zero value otherwise.
func (o *EtherPhysicalPortBaseAllOf) GetPeerInterface() PortInterfaceBaseRelationship {
	if o == nil || o.PeerInterface == nil {
		var ret PortInterfaceBaseRelationship
		return ret
	}
	return *o.PeerInterface
}

// GetPeerInterfaceOk returns a tuple with the PeerInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBaseAllOf) GetPeerInterfaceOk() (*PortInterfaceBaseRelationship, bool) {
	if o == nil || o.PeerInterface == nil {
		return nil, false
	}
	return o.PeerInterface, true
}

// HasPeerInterface returns a boolean if a field has been set.
func (o *EtherPhysicalPortBaseAllOf) HasPeerInterface() bool {
	if o != nil && o.PeerInterface != nil {
		return true
	}

	return false
}

// SetPeerInterface gets a reference to the given PortInterfaceBaseRelationship and assigns it to the PeerInterface field.
func (o *EtherPhysicalPortBaseAllOf) SetPeerInterface(v PortInterfaceBaseRelationship) {
	o.PeerInterface = &v
}

func (o EtherPhysicalPortBaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.OperSpeed != nil {
		toSerialize["OperSpeed"] = o.OperSpeed
	}
	if o.PeerDn != nil {
		toSerialize["PeerDn"] = o.PeerDn
	}
	if o.PortChannelId != nil {
		toSerialize["PortChannelId"] = o.PortChannelId
	}
	if o.PortType != nil {
		toSerialize["PortType"] = o.PortType
	}
	if o.TransceiverType != nil {
		toSerialize["TransceiverType"] = o.TransceiverType
	}
	if o.AcknowledgedPeerInterface != nil {
		toSerialize["AcknowledgedPeerInterface"] = o.AcknowledgedPeerInterface
	}
	if o.PeerInterface != nil {
		toSerialize["PeerInterface"] = o.PeerInterface
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EtherPhysicalPortBaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEtherPhysicalPortBaseAllOf := _EtherPhysicalPortBaseAllOf{}

	if err = json.Unmarshal(bytes, &varEtherPhysicalPortBaseAllOf); err == nil {
		*o = EtherPhysicalPortBaseAllOf(varEtherPhysicalPortBaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "OperSpeed")
		delete(additionalProperties, "PeerDn")
		delete(additionalProperties, "PortChannelId")
		delete(additionalProperties, "PortType")
		delete(additionalProperties, "TransceiverType")
		delete(additionalProperties, "AcknowledgedPeerInterface")
		delete(additionalProperties, "PeerInterface")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEtherPhysicalPortBaseAllOf struct {
	value *EtherPhysicalPortBaseAllOf
	isSet bool
}

func (v NullableEtherPhysicalPortBaseAllOf) Get() *EtherPhysicalPortBaseAllOf {
	return v.value
}

func (v *NullableEtherPhysicalPortBaseAllOf) Set(val *EtherPhysicalPortBaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEtherPhysicalPortBaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEtherPhysicalPortBaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEtherPhysicalPortBaseAllOf(val *EtherPhysicalPortBaseAllOf) *NullableEtherPhysicalPortBaseAllOf {
	return &NullableEtherPhysicalPortBaseAllOf{value: val, isSet: true}
}

func (v NullableEtherPhysicalPortBaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEtherPhysicalPortBaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


