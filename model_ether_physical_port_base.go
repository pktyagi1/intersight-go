/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4950
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// EtherPhysicalPortBase Abstract class for ether physical port and host port.
type EtherPhysicalPortBase struct {
	PortPhysical
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

type _EtherPhysicalPortBase EtherPhysicalPortBase

// NewEtherPhysicalPortBase instantiates a new EtherPhysicalPortBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEtherPhysicalPortBase(classId string, objectType string) *EtherPhysicalPortBase {
	this := EtherPhysicalPortBase{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEtherPhysicalPortBaseWithDefaults instantiates a new EtherPhysicalPortBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEtherPhysicalPortBaseWithDefaults() *EtherPhysicalPortBase {
	this := EtherPhysicalPortBase{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *EtherPhysicalPortBase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBase) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EtherPhysicalPortBase) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EtherPhysicalPortBase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBase) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EtherPhysicalPortBase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *EtherPhysicalPortBase) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBase) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *EtherPhysicalPortBase) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *EtherPhysicalPortBase) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *EtherPhysicalPortBase) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBase) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *EtherPhysicalPortBase) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *EtherPhysicalPortBase) SetMode(v string) {
	o.Mode = &v
}

// GetOperSpeed returns the OperSpeed field value if set, zero value otherwise.
func (o *EtherPhysicalPortBase) GetOperSpeed() string {
	if o == nil || o.OperSpeed == nil {
		var ret string
		return ret
	}
	return *o.OperSpeed
}

// GetOperSpeedOk returns a tuple with the OperSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBase) GetOperSpeedOk() (*string, bool) {
	if o == nil || o.OperSpeed == nil {
		return nil, false
	}
	return o.OperSpeed, true
}

// HasOperSpeed returns a boolean if a field has been set.
func (o *EtherPhysicalPortBase) HasOperSpeed() bool {
	if o != nil && o.OperSpeed != nil {
		return true
	}

	return false
}

// SetOperSpeed gets a reference to the given string and assigns it to the OperSpeed field.
func (o *EtherPhysicalPortBase) SetOperSpeed(v string) {
	o.OperSpeed = &v
}

// GetPeerDn returns the PeerDn field value if set, zero value otherwise.
func (o *EtherPhysicalPortBase) GetPeerDn() string {
	if o == nil || o.PeerDn == nil {
		var ret string
		return ret
	}
	return *o.PeerDn
}

// GetPeerDnOk returns a tuple with the PeerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBase) GetPeerDnOk() (*string, bool) {
	if o == nil || o.PeerDn == nil {
		return nil, false
	}
	return o.PeerDn, true
}

// HasPeerDn returns a boolean if a field has been set.
func (o *EtherPhysicalPortBase) HasPeerDn() bool {
	if o != nil && o.PeerDn != nil {
		return true
	}

	return false
}

// SetPeerDn gets a reference to the given string and assigns it to the PeerDn field.
func (o *EtherPhysicalPortBase) SetPeerDn(v string) {
	o.PeerDn = &v
}

// GetPortChannelId returns the PortChannelId field value if set, zero value otherwise.
func (o *EtherPhysicalPortBase) GetPortChannelId() int64 {
	if o == nil || o.PortChannelId == nil {
		var ret int64
		return ret
	}
	return *o.PortChannelId
}

// GetPortChannelIdOk returns a tuple with the PortChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBase) GetPortChannelIdOk() (*int64, bool) {
	if o == nil || o.PortChannelId == nil {
		return nil, false
	}
	return o.PortChannelId, true
}

// HasPortChannelId returns a boolean if a field has been set.
func (o *EtherPhysicalPortBase) HasPortChannelId() bool {
	if o != nil && o.PortChannelId != nil {
		return true
	}

	return false
}

// SetPortChannelId gets a reference to the given int64 and assigns it to the PortChannelId field.
func (o *EtherPhysicalPortBase) SetPortChannelId(v int64) {
	o.PortChannelId = &v
}

// GetPortType returns the PortType field value if set, zero value otherwise.
func (o *EtherPhysicalPortBase) GetPortType() string {
	if o == nil || o.PortType == nil {
		var ret string
		return ret
	}
	return *o.PortType
}

// GetPortTypeOk returns a tuple with the PortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBase) GetPortTypeOk() (*string, bool) {
	if o == nil || o.PortType == nil {
		return nil, false
	}
	return o.PortType, true
}

// HasPortType returns a boolean if a field has been set.
func (o *EtherPhysicalPortBase) HasPortType() bool {
	if o != nil && o.PortType != nil {
		return true
	}

	return false
}

// SetPortType gets a reference to the given string and assigns it to the PortType field.
func (o *EtherPhysicalPortBase) SetPortType(v string) {
	o.PortType = &v
}

// GetTransceiverType returns the TransceiverType field value if set, zero value otherwise.
func (o *EtherPhysicalPortBase) GetTransceiverType() string {
	if o == nil || o.TransceiverType == nil {
		var ret string
		return ret
	}
	return *o.TransceiverType
}

// GetTransceiverTypeOk returns a tuple with the TransceiverType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBase) GetTransceiverTypeOk() (*string, bool) {
	if o == nil || o.TransceiverType == nil {
		return nil, false
	}
	return o.TransceiverType, true
}

// HasTransceiverType returns a boolean if a field has been set.
func (o *EtherPhysicalPortBase) HasTransceiverType() bool {
	if o != nil && o.TransceiverType != nil {
		return true
	}

	return false
}

// SetTransceiverType gets a reference to the given string and assigns it to the TransceiverType field.
func (o *EtherPhysicalPortBase) SetTransceiverType(v string) {
	o.TransceiverType = &v
}

// GetAcknowledgedPeerInterface returns the AcknowledgedPeerInterface field value if set, zero value otherwise.
func (o *EtherPhysicalPortBase) GetAcknowledgedPeerInterface() PortInterfaceBaseRelationship {
	if o == nil || o.AcknowledgedPeerInterface == nil {
		var ret PortInterfaceBaseRelationship
		return ret
	}
	return *o.AcknowledgedPeerInterface
}

// GetAcknowledgedPeerInterfaceOk returns a tuple with the AcknowledgedPeerInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBase) GetAcknowledgedPeerInterfaceOk() (*PortInterfaceBaseRelationship, bool) {
	if o == nil || o.AcknowledgedPeerInterface == nil {
		return nil, false
	}
	return o.AcknowledgedPeerInterface, true
}

// HasAcknowledgedPeerInterface returns a boolean if a field has been set.
func (o *EtherPhysicalPortBase) HasAcknowledgedPeerInterface() bool {
	if o != nil && o.AcknowledgedPeerInterface != nil {
		return true
	}

	return false
}

// SetAcknowledgedPeerInterface gets a reference to the given PortInterfaceBaseRelationship and assigns it to the AcknowledgedPeerInterface field.
func (o *EtherPhysicalPortBase) SetAcknowledgedPeerInterface(v PortInterfaceBaseRelationship) {
	o.AcknowledgedPeerInterface = &v
}

// GetPeerInterface returns the PeerInterface field value if set, zero value otherwise.
func (o *EtherPhysicalPortBase) GetPeerInterface() PortInterfaceBaseRelationship {
	if o == nil || o.PeerInterface == nil {
		var ret PortInterfaceBaseRelationship
		return ret
	}
	return *o.PeerInterface
}

// GetPeerInterfaceOk returns a tuple with the PeerInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPhysicalPortBase) GetPeerInterfaceOk() (*PortInterfaceBaseRelationship, bool) {
	if o == nil || o.PeerInterface == nil {
		return nil, false
	}
	return o.PeerInterface, true
}

// HasPeerInterface returns a boolean if a field has been set.
func (o *EtherPhysicalPortBase) HasPeerInterface() bool {
	if o != nil && o.PeerInterface != nil {
		return true
	}

	return false
}

// SetPeerInterface gets a reference to the given PortInterfaceBaseRelationship and assigns it to the PeerInterface field.
func (o *EtherPhysicalPortBase) SetPeerInterface(v PortInterfaceBaseRelationship) {
	o.PeerInterface = &v
}

func (o EtherPhysicalPortBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPortPhysical, errPortPhysical := json.Marshal(o.PortPhysical)
	if errPortPhysical != nil {
		return []byte{}, errPortPhysical
	}
	errPortPhysical = json.Unmarshal([]byte(serializedPortPhysical), &toSerialize)
	if errPortPhysical != nil {
		return []byte{}, errPortPhysical
	}
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

func (o *EtherPhysicalPortBase) UnmarshalJSON(bytes []byte) (err error) {
	type EtherPhysicalPortBaseWithoutEmbeddedStruct struct {
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
	}

	varEtherPhysicalPortBaseWithoutEmbeddedStruct := EtherPhysicalPortBaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEtherPhysicalPortBaseWithoutEmbeddedStruct)
	if err == nil {
		varEtherPhysicalPortBase := _EtherPhysicalPortBase{}
		varEtherPhysicalPortBase.ClassId = varEtherPhysicalPortBaseWithoutEmbeddedStruct.ClassId
		varEtherPhysicalPortBase.ObjectType = varEtherPhysicalPortBaseWithoutEmbeddedStruct.ObjectType
		varEtherPhysicalPortBase.MacAddress = varEtherPhysicalPortBaseWithoutEmbeddedStruct.MacAddress
		varEtherPhysicalPortBase.Mode = varEtherPhysicalPortBaseWithoutEmbeddedStruct.Mode
		varEtherPhysicalPortBase.OperSpeed = varEtherPhysicalPortBaseWithoutEmbeddedStruct.OperSpeed
		varEtherPhysicalPortBase.PeerDn = varEtherPhysicalPortBaseWithoutEmbeddedStruct.PeerDn
		varEtherPhysicalPortBase.PortChannelId = varEtherPhysicalPortBaseWithoutEmbeddedStruct.PortChannelId
		varEtherPhysicalPortBase.PortType = varEtherPhysicalPortBaseWithoutEmbeddedStruct.PortType
		varEtherPhysicalPortBase.TransceiverType = varEtherPhysicalPortBaseWithoutEmbeddedStruct.TransceiverType
		varEtherPhysicalPortBase.AcknowledgedPeerInterface = varEtherPhysicalPortBaseWithoutEmbeddedStruct.AcknowledgedPeerInterface
		varEtherPhysicalPortBase.PeerInterface = varEtherPhysicalPortBaseWithoutEmbeddedStruct.PeerInterface
		*o = EtherPhysicalPortBase(varEtherPhysicalPortBase)
	} else {
		return err
	}

	varEtherPhysicalPortBase := _EtherPhysicalPortBase{}

	err = json.Unmarshal(bytes, &varEtherPhysicalPortBase)
	if err == nil {
		o.PortPhysical = varEtherPhysicalPortBase.PortPhysical
	} else {
		return err
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

		// remove fields from embedded structs
		reflectPortPhysical := reflect.ValueOf(o.PortPhysical)
		for i := 0; i < reflectPortPhysical.Type().NumField(); i++ {
			t := reflectPortPhysical.Type().Field(i)

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

type NullableEtherPhysicalPortBase struct {
	value *EtherPhysicalPortBase
	isSet bool
}

func (v NullableEtherPhysicalPortBase) Get() *EtherPhysicalPortBase {
	return v.value
}

func (v *NullableEtherPhysicalPortBase) Set(val *EtherPhysicalPortBase) {
	v.value = val
	v.isSet = true
}

func (v NullableEtherPhysicalPortBase) IsSet() bool {
	return v.isSet
}

func (v *NullableEtherPhysicalPortBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEtherPhysicalPortBase(val *EtherPhysicalPortBase) *NullableEtherPhysicalPortBase {
	return &NullableEtherPhysicalPortBase{value: val, isSet: true}
}

func (v NullableEtherPhysicalPortBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEtherPhysicalPortBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

