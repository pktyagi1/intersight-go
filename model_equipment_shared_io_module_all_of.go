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

// EquipmentSharedIoModuleAllOf Definition of the list of properties defined in 'equipment.SharedIoModule', excluding properties defined in parent classes.
type EquipmentSharedIoModuleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field identifies the configuration state for this SIOM Unit.
	ConfigState *string `json:"ConfigState,omitempty"`
	// This field identifies the discovery state of SIOM.
	Discovery *string `json:"Discovery,omitempty"`
	// This field identifies the MAC of IOM-A side.
	MacOfSharedIomAside *string `json:"MacOfSharedIomAside,omitempty"`
	// This field identifies the MAC of IOM-B side.
	MacOfSharedIomBside *string `json:"MacOfSharedIomBside,omitempty"`
	// This field identifies the SIOM operational state.
	OperState *string `json:"OperState,omitempty"`
	// This field identifies the Part Number for this SIOM Unit.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field identifies the reachability to FI-A and B side.
	Reachability *string `json:"Reachability,omitempty"`
	// User label configured for the SIOM.
	UsrLbl *string `json:"UsrLbl,omitempty"`
	// This field identifies the vendor id for this SIOM Unit.
	Vid *string `json:"Vid,omitempty"`
	Controller *ManagementControllerRelationship `json:"Controller,omitempty"`
	EquipmentSystemIoController *EquipmentSystemIoControllerRelationship `json:"EquipmentSystemIoController,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to portGroup resources.
	PortGroups []PortGroupRelationship `json:"PortGroups,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentSharedIoModuleAllOf EquipmentSharedIoModuleAllOf

// NewEquipmentSharedIoModuleAllOf instantiates a new EquipmentSharedIoModuleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentSharedIoModuleAllOf(classId string, objectType string) *EquipmentSharedIoModuleAllOf {
	this := EquipmentSharedIoModuleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentSharedIoModuleAllOfWithDefaults instantiates a new EquipmentSharedIoModuleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentSharedIoModuleAllOfWithDefaults() *EquipmentSharedIoModuleAllOf {
	this := EquipmentSharedIoModuleAllOf{}
	var classId string = "equipment.SharedIoModule"
	this.ClassId = classId
	var objectType string = "equipment.SharedIoModule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentSharedIoModuleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModuleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentSharedIoModuleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentSharedIoModuleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModuleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentSharedIoModuleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *EquipmentSharedIoModuleAllOf) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModuleAllOf) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *EquipmentSharedIoModuleAllOf) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *EquipmentSharedIoModuleAllOf) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetDiscovery returns the Discovery field value if set, zero value otherwise.
func (o *EquipmentSharedIoModuleAllOf) GetDiscovery() string {
	if o == nil || o.Discovery == nil {
		var ret string
		return ret
	}
	return *o.Discovery
}

// GetDiscoveryOk returns a tuple with the Discovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModuleAllOf) GetDiscoveryOk() (*string, bool) {
	if o == nil || o.Discovery == nil {
		return nil, false
	}
	return o.Discovery, true
}

// HasDiscovery returns a boolean if a field has been set.
func (o *EquipmentSharedIoModuleAllOf) HasDiscovery() bool {
	if o != nil && o.Discovery != nil {
		return true
	}

	return false
}

// SetDiscovery gets a reference to the given string and assigns it to the Discovery field.
func (o *EquipmentSharedIoModuleAllOf) SetDiscovery(v string) {
	o.Discovery = &v
}

// GetMacOfSharedIomAside returns the MacOfSharedIomAside field value if set, zero value otherwise.
func (o *EquipmentSharedIoModuleAllOf) GetMacOfSharedIomAside() string {
	if o == nil || o.MacOfSharedIomAside == nil {
		var ret string
		return ret
	}
	return *o.MacOfSharedIomAside
}

// GetMacOfSharedIomAsideOk returns a tuple with the MacOfSharedIomAside field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModuleAllOf) GetMacOfSharedIomAsideOk() (*string, bool) {
	if o == nil || o.MacOfSharedIomAside == nil {
		return nil, false
	}
	return o.MacOfSharedIomAside, true
}

// HasMacOfSharedIomAside returns a boolean if a field has been set.
func (o *EquipmentSharedIoModuleAllOf) HasMacOfSharedIomAside() bool {
	if o != nil && o.MacOfSharedIomAside != nil {
		return true
	}

	return false
}

// SetMacOfSharedIomAside gets a reference to the given string and assigns it to the MacOfSharedIomAside field.
func (o *EquipmentSharedIoModuleAllOf) SetMacOfSharedIomAside(v string) {
	o.MacOfSharedIomAside = &v
}

// GetMacOfSharedIomBside returns the MacOfSharedIomBside field value if set, zero value otherwise.
func (o *EquipmentSharedIoModuleAllOf) GetMacOfSharedIomBside() string {
	if o == nil || o.MacOfSharedIomBside == nil {
		var ret string
		return ret
	}
	return *o.MacOfSharedIomBside
}

// GetMacOfSharedIomBsideOk returns a tuple with the MacOfSharedIomBside field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModuleAllOf) GetMacOfSharedIomBsideOk() (*string, bool) {
	if o == nil || o.MacOfSharedIomBside == nil {
		return nil, false
	}
	return o.MacOfSharedIomBside, true
}

// HasMacOfSharedIomBside returns a boolean if a field has been set.
func (o *EquipmentSharedIoModuleAllOf) HasMacOfSharedIomBside() bool {
	if o != nil && o.MacOfSharedIomBside != nil {
		return true
	}

	return false
}

// SetMacOfSharedIomBside gets a reference to the given string and assigns it to the MacOfSharedIomBside field.
func (o *EquipmentSharedIoModuleAllOf) SetMacOfSharedIomBside(v string) {
	o.MacOfSharedIomBside = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentSharedIoModuleAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModuleAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentSharedIoModuleAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentSharedIoModuleAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentSharedIoModuleAllOf) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModuleAllOf) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentSharedIoModuleAllOf) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentSharedIoModuleAllOf) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetReachability returns the Reachability field value if set, zero value otherwise.
func (o *EquipmentSharedIoModuleAllOf) GetReachability() string {
	if o == nil || o.Reachability == nil {
		var ret string
		return ret
	}
	return *o.Reachability
}

// GetReachabilityOk returns a tuple with the Reachability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModuleAllOf) GetReachabilityOk() (*string, bool) {
	if o == nil || o.Reachability == nil {
		return nil, false
	}
	return o.Reachability, true
}

// HasReachability returns a boolean if a field has been set.
func (o *EquipmentSharedIoModuleAllOf) HasReachability() bool {
	if o != nil && o.Reachability != nil {
		return true
	}

	return false
}

// SetReachability gets a reference to the given string and assigns it to the Reachability field.
func (o *EquipmentSharedIoModuleAllOf) SetReachability(v string) {
	o.Reachability = &v
}

// GetUsrLbl returns the UsrLbl field value if set, zero value otherwise.
func (o *EquipmentSharedIoModuleAllOf) GetUsrLbl() string {
	if o == nil || o.UsrLbl == nil {
		var ret string
		return ret
	}
	return *o.UsrLbl
}

// GetUsrLblOk returns a tuple with the UsrLbl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModuleAllOf) GetUsrLblOk() (*string, bool) {
	if o == nil || o.UsrLbl == nil {
		return nil, false
	}
	return o.UsrLbl, true
}

// HasUsrLbl returns a boolean if a field has been set.
func (o *EquipmentSharedIoModuleAllOf) HasUsrLbl() bool {
	if o != nil && o.UsrLbl != nil {
		return true
	}

	return false
}

// SetUsrLbl gets a reference to the given string and assigns it to the UsrLbl field.
func (o *EquipmentSharedIoModuleAllOf) SetUsrLbl(v string) {
	o.UsrLbl = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *EquipmentSharedIoModuleAllOf) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModuleAllOf) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *EquipmentSharedIoModuleAllOf) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *EquipmentSharedIoModuleAllOf) SetVid(v string) {
	o.Vid = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *EquipmentSharedIoModuleAllOf) GetController() ManagementControllerRelationship {
	if o == nil || o.Controller == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModuleAllOf) GetControllerOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *EquipmentSharedIoModuleAllOf) HasController() bool {
	if o != nil && o.Controller != nil {
		return true
	}

	return false
}

// SetController gets a reference to the given ManagementControllerRelationship and assigns it to the Controller field.
func (o *EquipmentSharedIoModuleAllOf) SetController(v ManagementControllerRelationship) {
	o.Controller = &v
}

// GetEquipmentSystemIoController returns the EquipmentSystemIoController field value if set, zero value otherwise.
func (o *EquipmentSharedIoModuleAllOf) GetEquipmentSystemIoController() EquipmentSystemIoControllerRelationship {
	if o == nil || o.EquipmentSystemIoController == nil {
		var ret EquipmentSystemIoControllerRelationship
		return ret
	}
	return *o.EquipmentSystemIoController
}

// GetEquipmentSystemIoControllerOk returns a tuple with the EquipmentSystemIoController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModuleAllOf) GetEquipmentSystemIoControllerOk() (*EquipmentSystemIoControllerRelationship, bool) {
	if o == nil || o.EquipmentSystemIoController == nil {
		return nil, false
	}
	return o.EquipmentSystemIoController, true
}

// HasEquipmentSystemIoController returns a boolean if a field has been set.
func (o *EquipmentSharedIoModuleAllOf) HasEquipmentSystemIoController() bool {
	if o != nil && o.EquipmentSystemIoController != nil {
		return true
	}

	return false
}

// SetEquipmentSystemIoController gets a reference to the given EquipmentSystemIoControllerRelationship and assigns it to the EquipmentSystemIoController field.
func (o *EquipmentSharedIoModuleAllOf) SetEquipmentSystemIoController(v EquipmentSystemIoControllerRelationship) {
	o.EquipmentSystemIoController = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentSharedIoModuleAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModuleAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentSharedIoModuleAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentSharedIoModuleAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPortGroups returns the PortGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentSharedIoModuleAllOf) GetPortGroups() []PortGroupRelationship {
	if o == nil  {
		var ret []PortGroupRelationship
		return ret
	}
	return o.PortGroups
}

// GetPortGroupsOk returns a tuple with the PortGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentSharedIoModuleAllOf) GetPortGroupsOk() (*[]PortGroupRelationship, bool) {
	if o == nil || o.PortGroups == nil {
		return nil, false
	}
	return &o.PortGroups, true
}

// HasPortGroups returns a boolean if a field has been set.
func (o *EquipmentSharedIoModuleAllOf) HasPortGroups() bool {
	if o != nil && o.PortGroups != nil {
		return true
	}

	return false
}

// SetPortGroups gets a reference to the given []PortGroupRelationship and assigns it to the PortGroups field.
func (o *EquipmentSharedIoModuleAllOf) SetPortGroups(v []PortGroupRelationship) {
	o.PortGroups = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentSharedIoModuleAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModuleAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentSharedIoModuleAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentSharedIoModuleAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentSharedIoModuleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.Discovery != nil {
		toSerialize["Discovery"] = o.Discovery
	}
	if o.MacOfSharedIomAside != nil {
		toSerialize["MacOfSharedIomAside"] = o.MacOfSharedIomAside
	}
	if o.MacOfSharedIomBside != nil {
		toSerialize["MacOfSharedIomBside"] = o.MacOfSharedIomBside
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.Reachability != nil {
		toSerialize["Reachability"] = o.Reachability
	}
	if o.UsrLbl != nil {
		toSerialize["UsrLbl"] = o.UsrLbl
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}
	if o.Controller != nil {
		toSerialize["Controller"] = o.Controller
	}
	if o.EquipmentSystemIoController != nil {
		toSerialize["EquipmentSystemIoController"] = o.EquipmentSystemIoController
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.PortGroups != nil {
		toSerialize["PortGroups"] = o.PortGroups
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentSharedIoModuleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentSharedIoModuleAllOf := _EquipmentSharedIoModuleAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentSharedIoModuleAllOf); err == nil {
		*o = EquipmentSharedIoModuleAllOf(varEquipmentSharedIoModuleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "Discovery")
		delete(additionalProperties, "MacOfSharedIomAside")
		delete(additionalProperties, "MacOfSharedIomBside")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Reachability")
		delete(additionalProperties, "UsrLbl")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "Controller")
		delete(additionalProperties, "EquipmentSystemIoController")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PortGroups")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentSharedIoModuleAllOf struct {
	value *EquipmentSharedIoModuleAllOf
	isSet bool
}

func (v NullableEquipmentSharedIoModuleAllOf) Get() *EquipmentSharedIoModuleAllOf {
	return v.value
}

func (v *NullableEquipmentSharedIoModuleAllOf) Set(val *EquipmentSharedIoModuleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentSharedIoModuleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentSharedIoModuleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentSharedIoModuleAllOf(val *EquipmentSharedIoModuleAllOf) *NullableEquipmentSharedIoModuleAllOf {
	return &NullableEquipmentSharedIoModuleAllOf{value: val, isSet: true}
}

func (v NullableEquipmentSharedIoModuleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentSharedIoModuleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


