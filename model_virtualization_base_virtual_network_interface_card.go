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
	"reflect"
	"strings"
)

// VirtualizationBaseVirtualNetworkInterfaceCard Common attributes of any virtual network interface card managed by a hypervisor. Serves as a base class for all concrete virtual network interface card types.
type VirtualizationBaseVirtualNetworkInterfaceCard struct {
	VirtualizationBaseSourceDevice
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Type or model of the virtual network interface card. * `Unknown` - The type of the network adaptor type is unknown. * `E1000` - Emulated version of the Intel 82545EM Gigabit Ethernet NIC. * `SRIOV` - Representation of a virtual function (VF) on a physical NIC with SR-IOV support. * `VMXNET2` - VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later. * `VMXNET3` - VMXNET 3 offers all the features available in VMXNET 2 and adds several new features. * `E1000E` - E1000E – emulates a newer real network adapter, the 1 Gbit Intel 82574, and is available for Windows 2012 and later. The E1000E needs virtual machine hardware version 8 or later. * `NE2K_PCI` - The Ne2000 network card uses two ring buffers for packet handling. These are circular buffers made of 256-byte pages that the chip's DMA logic will use to store received packets or to get received packets. * `PCnet` - The PCnet-PCI II is a PCI network adapter. It has built-in support for CRC checks and can automatically pad short packets to the minimum Ethernet length. * `RTL8139` - The RTL8139 is a fast Ethernet card that operates at 10/100 Mbps. It is compliant with PCI version 2.0/2.1 and it is known for reliability and superior performance. * `VirtIO` - VirtIO is a standardized interface which allows virtual machines access to simplified \"virtual\" devices, such as block devices, network adapters and consoles. Accessing devices through VirtIO on a guest VM improves performance over more traditional \"emulated\" devices, as VirtIO devices require only the bare minimum setup and configuration needed to send and receive data, while the host machine handles the majority of the setup and maintenance of the actual physical hardware. * `` - Default network adaptor type supported by the hypervisor.
	AdapterType *string `json:"AdapterType,omitempty"`
	// MAC address assigned to the virtual network interface card.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Name of the virtual network interface card.
	Name *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseVirtualNetworkInterfaceCard VirtualizationBaseVirtualNetworkInterfaceCard

// NewVirtualizationBaseVirtualNetworkInterfaceCard instantiates a new VirtualizationBaseVirtualNetworkInterfaceCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseVirtualNetworkInterfaceCard(classId string, objectType string) *VirtualizationBaseVirtualNetworkInterfaceCard {
	this := VirtualizationBaseVirtualNetworkInterfaceCard{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adapterType string = "Unknown"
	this.AdapterType = &adapterType
	return &this
}

// NewVirtualizationBaseVirtualNetworkInterfaceCardWithDefaults instantiates a new VirtualizationBaseVirtualNetworkInterfaceCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseVirtualNetworkInterfaceCardWithDefaults() *VirtualizationBaseVirtualNetworkInterfaceCard {
	this := VirtualizationBaseVirtualNetworkInterfaceCard{}
	var adapterType string = "Unknown"
	this.AdapterType = &adapterType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdapterType returns the AdapterType field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) GetAdapterType() string {
	if o == nil || o.AdapterType == nil {
		var ret string
		return ret
	}
	return *o.AdapterType
}

// GetAdapterTypeOk returns a tuple with the AdapterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) GetAdapterTypeOk() (*string, bool) {
	if o == nil || o.AdapterType == nil {
		return nil, false
	}
	return o.AdapterType, true
}

// HasAdapterType returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) HasAdapterType() bool {
	if o != nil && o.AdapterType != nil {
		return true
	}

	return false
}

// SetAdapterType gets a reference to the given string and assigns it to the AdapterType field.
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) SetAdapterType(v string) {
	o.AdapterType = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseVirtualNetworkInterfaceCard) SetName(v string) {
	o.Name = &v
}

func (o VirtualizationBaseVirtualNetworkInterfaceCard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseSourceDevice, errVirtualizationBaseSourceDevice := json.Marshal(o.VirtualizationBaseSourceDevice)
	if errVirtualizationBaseSourceDevice != nil {
		return []byte{}, errVirtualizationBaseSourceDevice
	}
	errVirtualizationBaseSourceDevice = json.Unmarshal([]byte(serializedVirtualizationBaseSourceDevice), &toSerialize)
	if errVirtualizationBaseSourceDevice != nil {
		return []byte{}, errVirtualizationBaseSourceDevice
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdapterType != nil {
		toSerialize["AdapterType"] = o.AdapterType
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBaseVirtualNetworkInterfaceCard) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationBaseVirtualNetworkInterfaceCardWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Type or model of the virtual network interface card. * `Unknown` - The type of the network adaptor type is unknown. * `E1000` - Emulated version of the Intel 82545EM Gigabit Ethernet NIC. * `SRIOV` - Representation of a virtual function (VF) on a physical NIC with SR-IOV support. * `VMXNET2` - VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later. * `VMXNET3` - VMXNET 3 offers all the features available in VMXNET 2 and adds several new features. * `E1000E` - E1000E – emulates a newer real network adapter, the 1 Gbit Intel 82574, and is available for Windows 2012 and later. The E1000E needs virtual machine hardware version 8 or later. * `NE2K_PCI` - The Ne2000 network card uses two ring buffers for packet handling. These are circular buffers made of 256-byte pages that the chip's DMA logic will use to store received packets or to get received packets. * `PCnet` - The PCnet-PCI II is a PCI network adapter. It has built-in support for CRC checks and can automatically pad short packets to the minimum Ethernet length. * `RTL8139` - The RTL8139 is a fast Ethernet card that operates at 10/100 Mbps. It is compliant with PCI version 2.0/2.1 and it is known for reliability and superior performance. * `VirtIO` - VirtIO is a standardized interface which allows virtual machines access to simplified \"virtual\" devices, such as block devices, network adapters and consoles. Accessing devices through VirtIO on a guest VM improves performance over more traditional \"emulated\" devices, as VirtIO devices require only the bare minimum setup and configuration needed to send and receive data, while the host machine handles the majority of the setup and maintenance of the actual physical hardware. * `` - Default network adaptor type supported by the hypervisor.
		AdapterType *string `json:"AdapterType,omitempty"`
		// MAC address assigned to the virtual network interface card.
		MacAddress *string `json:"MacAddress,omitempty"`
		// Name of the virtual network interface card.
		Name *string `json:"Name,omitempty"`
	}

	varVirtualizationBaseVirtualNetworkInterfaceCardWithoutEmbeddedStruct := VirtualizationBaseVirtualNetworkInterfaceCardWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseVirtualNetworkInterfaceCardWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationBaseVirtualNetworkInterfaceCard := _VirtualizationBaseVirtualNetworkInterfaceCard{}
		varVirtualizationBaseVirtualNetworkInterfaceCard.ClassId = varVirtualizationBaseVirtualNetworkInterfaceCardWithoutEmbeddedStruct.ClassId
		varVirtualizationBaseVirtualNetworkInterfaceCard.ObjectType = varVirtualizationBaseVirtualNetworkInterfaceCardWithoutEmbeddedStruct.ObjectType
		varVirtualizationBaseVirtualNetworkInterfaceCard.AdapterType = varVirtualizationBaseVirtualNetworkInterfaceCardWithoutEmbeddedStruct.AdapterType
		varVirtualizationBaseVirtualNetworkInterfaceCard.MacAddress = varVirtualizationBaseVirtualNetworkInterfaceCardWithoutEmbeddedStruct.MacAddress
		varVirtualizationBaseVirtualNetworkInterfaceCard.Name = varVirtualizationBaseVirtualNetworkInterfaceCardWithoutEmbeddedStruct.Name
		*o = VirtualizationBaseVirtualNetworkInterfaceCard(varVirtualizationBaseVirtualNetworkInterfaceCard)
	} else {
		return err
	}

	varVirtualizationBaseVirtualNetworkInterfaceCard := _VirtualizationBaseVirtualNetworkInterfaceCard{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseVirtualNetworkInterfaceCard)
	if err == nil {
		o.VirtualizationBaseSourceDevice = varVirtualizationBaseVirtualNetworkInterfaceCard.VirtualizationBaseSourceDevice
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdapterType")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Name")

		// remove fields from embedded structs
		reflectVirtualizationBaseSourceDevice := reflect.ValueOf(o.VirtualizationBaseSourceDevice)
		for i := 0; i < reflectVirtualizationBaseSourceDevice.Type().NumField(); i++ {
			t := reflectVirtualizationBaseSourceDevice.Type().Field(i)

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

type NullableVirtualizationBaseVirtualNetworkInterfaceCard struct {
	value *VirtualizationBaseVirtualNetworkInterfaceCard
	isSet bool
}

func (v NullableVirtualizationBaseVirtualNetworkInterfaceCard) Get() *VirtualizationBaseVirtualNetworkInterfaceCard {
	return v.value
}

func (v *NullableVirtualizationBaseVirtualNetworkInterfaceCard) Set(val *VirtualizationBaseVirtualNetworkInterfaceCard) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseVirtualNetworkInterfaceCard) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseVirtualNetworkInterfaceCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseVirtualNetworkInterfaceCard(val *VirtualizationBaseVirtualNetworkInterfaceCard) *NullableVirtualizationBaseVirtualNetworkInterfaceCard {
	return &NullableVirtualizationBaseVirtualNetworkInterfaceCard{value: val, isSet: true}
}

func (v NullableVirtualizationBaseVirtualNetworkInterfaceCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseVirtualNetworkInterfaceCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


