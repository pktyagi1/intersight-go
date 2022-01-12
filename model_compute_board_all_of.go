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
)

// ComputeBoardAllOf Definition of the list of properties defined in 'compute.Board', excluding properties defined in parent classes.
type ComputeBoardAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique identifier of the mother board present in the server.
	BoardId *int64 `json:"BoardId,omitempty"`
	// The type of central processing unit on the mother board.
	CpuTypeController *string `json:"CpuTypeController,omitempty"`
	// Current power state of the mother board of the server.
	OperPowerState *string `json:"OperPowerState,omitempty"`
	OperReason []string `json:"OperReason,omitempty"`
	ComputeBlade *ComputeBladeRelationship `json:"ComputeBlade,omitempty"`
	ComputeRackUnit *ComputeRackUnitRelationship `json:"ComputeRackUnit,omitempty"`
	// An array of relationships to equipmentTpm resources.
	EquipmentTpms []EquipmentTpmRelationship `json:"EquipmentTpms,omitempty"`
	// An array of relationships to graphicsCard resources.
	GraphicsCards []GraphicsCardRelationship `json:"GraphicsCards,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to memoryArray resources.
	MemoryArrays []MemoryArrayRelationship `json:"MemoryArrays,omitempty"`
	// An array of relationships to pciCoprocessorCard resources.
	PciCoprocessorCards []PciCoprocessorCardRelationship `json:"PciCoprocessorCards,omitempty"`
	// An array of relationships to pciSwitch resources.
	PciSwitch []PciSwitchRelationship `json:"PciSwitch,omitempty"`
	PersistentMemoryConfiguration *MemoryPersistentMemoryConfigurationRelationship `json:"PersistentMemoryConfiguration,omitempty"`
	// An array of relationships to processorUnit resources.
	Processors []ProcessorUnitRelationship `json:"Processors,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to securityUnit resources.
	SecurityUnits []SecurityUnitRelationship `json:"SecurityUnits,omitempty"`
	// An array of relationships to storageController resources.
	StorageControllers []StorageControllerRelationship `json:"StorageControllers,omitempty"`
	// An array of relationships to storageFlexFlashController resources.
	StorageFlexFlashControllers []StorageFlexFlashControllerRelationship `json:"StorageFlexFlashControllers,omitempty"`
	// An array of relationships to storageFlexUtilController resources.
	StorageFlexUtilControllers []StorageFlexUtilControllerRelationship `json:"StorageFlexUtilControllers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeBoardAllOf ComputeBoardAllOf

// NewComputeBoardAllOf instantiates a new ComputeBoardAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeBoardAllOf(classId string, objectType string) *ComputeBoardAllOf {
	this := ComputeBoardAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewComputeBoardAllOfWithDefaults instantiates a new ComputeBoardAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeBoardAllOfWithDefaults() *ComputeBoardAllOf {
	this := ComputeBoardAllOf{}
	var classId string = "compute.Board"
	this.ClassId = classId
	var objectType string = "compute.Board"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeBoardAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeBoardAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeBoardAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeBoardAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeBoardAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeBoardAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBoardId returns the BoardId field value if set, zero value otherwise.
func (o *ComputeBoardAllOf) GetBoardId() int64 {
	if o == nil || o.BoardId == nil {
		var ret int64
		return ret
	}
	return *o.BoardId
}

// GetBoardIdOk returns a tuple with the BoardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBoardAllOf) GetBoardIdOk() (*int64, bool) {
	if o == nil || o.BoardId == nil {
		return nil, false
	}
	return o.BoardId, true
}

// HasBoardId returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasBoardId() bool {
	if o != nil && o.BoardId != nil {
		return true
	}

	return false
}

// SetBoardId gets a reference to the given int64 and assigns it to the BoardId field.
func (o *ComputeBoardAllOf) SetBoardId(v int64) {
	o.BoardId = &v
}

// GetCpuTypeController returns the CpuTypeController field value if set, zero value otherwise.
func (o *ComputeBoardAllOf) GetCpuTypeController() string {
	if o == nil || o.CpuTypeController == nil {
		var ret string
		return ret
	}
	return *o.CpuTypeController
}

// GetCpuTypeControllerOk returns a tuple with the CpuTypeController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBoardAllOf) GetCpuTypeControllerOk() (*string, bool) {
	if o == nil || o.CpuTypeController == nil {
		return nil, false
	}
	return o.CpuTypeController, true
}

// HasCpuTypeController returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasCpuTypeController() bool {
	if o != nil && o.CpuTypeController != nil {
		return true
	}

	return false
}

// SetCpuTypeController gets a reference to the given string and assigns it to the CpuTypeController field.
func (o *ComputeBoardAllOf) SetCpuTypeController(v string) {
	o.CpuTypeController = &v
}

// GetOperPowerState returns the OperPowerState field value if set, zero value otherwise.
func (o *ComputeBoardAllOf) GetOperPowerState() string {
	if o == nil || o.OperPowerState == nil {
		var ret string
		return ret
	}
	return *o.OperPowerState
}

// GetOperPowerStateOk returns a tuple with the OperPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBoardAllOf) GetOperPowerStateOk() (*string, bool) {
	if o == nil || o.OperPowerState == nil {
		return nil, false
	}
	return o.OperPowerState, true
}

// HasOperPowerState returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasOperPowerState() bool {
	if o != nil && o.OperPowerState != nil {
		return true
	}

	return false
}

// SetOperPowerState gets a reference to the given string and assigns it to the OperPowerState field.
func (o *ComputeBoardAllOf) SetOperPowerState(v string) {
	o.OperPowerState = &v
}

// GetOperReason returns the OperReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBoardAllOf) GetOperReason() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.OperReason
}

// GetOperReasonOk returns a tuple with the OperReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBoardAllOf) GetOperReasonOk() (*[]string, bool) {
	if o == nil || o.OperReason == nil {
		return nil, false
	}
	return &o.OperReason, true
}

// HasOperReason returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasOperReason() bool {
	if o != nil && o.OperReason != nil {
		return true
	}

	return false
}

// SetOperReason gets a reference to the given []string and assigns it to the OperReason field.
func (o *ComputeBoardAllOf) SetOperReason(v []string) {
	o.OperReason = v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *ComputeBoardAllOf) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBoardAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *ComputeBoardAllOf) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *ComputeBoardAllOf) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBoardAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *ComputeBoardAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetEquipmentTpms returns the EquipmentTpms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBoardAllOf) GetEquipmentTpms() []EquipmentTpmRelationship {
	if o == nil  {
		var ret []EquipmentTpmRelationship
		return ret
	}
	return o.EquipmentTpms
}

// GetEquipmentTpmsOk returns a tuple with the EquipmentTpms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBoardAllOf) GetEquipmentTpmsOk() (*[]EquipmentTpmRelationship, bool) {
	if o == nil || o.EquipmentTpms == nil {
		return nil, false
	}
	return &o.EquipmentTpms, true
}

// HasEquipmentTpms returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasEquipmentTpms() bool {
	if o != nil && o.EquipmentTpms != nil {
		return true
	}

	return false
}

// SetEquipmentTpms gets a reference to the given []EquipmentTpmRelationship and assigns it to the EquipmentTpms field.
func (o *ComputeBoardAllOf) SetEquipmentTpms(v []EquipmentTpmRelationship) {
	o.EquipmentTpms = v
}

// GetGraphicsCards returns the GraphicsCards field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBoardAllOf) GetGraphicsCards() []GraphicsCardRelationship {
	if o == nil  {
		var ret []GraphicsCardRelationship
		return ret
	}
	return o.GraphicsCards
}

// GetGraphicsCardsOk returns a tuple with the GraphicsCards field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBoardAllOf) GetGraphicsCardsOk() (*[]GraphicsCardRelationship, bool) {
	if o == nil || o.GraphicsCards == nil {
		return nil, false
	}
	return &o.GraphicsCards, true
}

// HasGraphicsCards returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasGraphicsCards() bool {
	if o != nil && o.GraphicsCards != nil {
		return true
	}

	return false
}

// SetGraphicsCards gets a reference to the given []GraphicsCardRelationship and assigns it to the GraphicsCards field.
func (o *ComputeBoardAllOf) SetGraphicsCards(v []GraphicsCardRelationship) {
	o.GraphicsCards = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *ComputeBoardAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBoardAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *ComputeBoardAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetMemoryArrays returns the MemoryArrays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBoardAllOf) GetMemoryArrays() []MemoryArrayRelationship {
	if o == nil  {
		var ret []MemoryArrayRelationship
		return ret
	}
	return o.MemoryArrays
}

// GetMemoryArraysOk returns a tuple with the MemoryArrays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBoardAllOf) GetMemoryArraysOk() (*[]MemoryArrayRelationship, bool) {
	if o == nil || o.MemoryArrays == nil {
		return nil, false
	}
	return &o.MemoryArrays, true
}

// HasMemoryArrays returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasMemoryArrays() bool {
	if o != nil && o.MemoryArrays != nil {
		return true
	}

	return false
}

// SetMemoryArrays gets a reference to the given []MemoryArrayRelationship and assigns it to the MemoryArrays field.
func (o *ComputeBoardAllOf) SetMemoryArrays(v []MemoryArrayRelationship) {
	o.MemoryArrays = v
}

// GetPciCoprocessorCards returns the PciCoprocessorCards field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBoardAllOf) GetPciCoprocessorCards() []PciCoprocessorCardRelationship {
	if o == nil  {
		var ret []PciCoprocessorCardRelationship
		return ret
	}
	return o.PciCoprocessorCards
}

// GetPciCoprocessorCardsOk returns a tuple with the PciCoprocessorCards field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBoardAllOf) GetPciCoprocessorCardsOk() (*[]PciCoprocessorCardRelationship, bool) {
	if o == nil || o.PciCoprocessorCards == nil {
		return nil, false
	}
	return &o.PciCoprocessorCards, true
}

// HasPciCoprocessorCards returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasPciCoprocessorCards() bool {
	if o != nil && o.PciCoprocessorCards != nil {
		return true
	}

	return false
}

// SetPciCoprocessorCards gets a reference to the given []PciCoprocessorCardRelationship and assigns it to the PciCoprocessorCards field.
func (o *ComputeBoardAllOf) SetPciCoprocessorCards(v []PciCoprocessorCardRelationship) {
	o.PciCoprocessorCards = v
}

// GetPciSwitch returns the PciSwitch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBoardAllOf) GetPciSwitch() []PciSwitchRelationship {
	if o == nil  {
		var ret []PciSwitchRelationship
		return ret
	}
	return o.PciSwitch
}

// GetPciSwitchOk returns a tuple with the PciSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBoardAllOf) GetPciSwitchOk() (*[]PciSwitchRelationship, bool) {
	if o == nil || o.PciSwitch == nil {
		return nil, false
	}
	return &o.PciSwitch, true
}

// HasPciSwitch returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasPciSwitch() bool {
	if o != nil && o.PciSwitch != nil {
		return true
	}

	return false
}

// SetPciSwitch gets a reference to the given []PciSwitchRelationship and assigns it to the PciSwitch field.
func (o *ComputeBoardAllOf) SetPciSwitch(v []PciSwitchRelationship) {
	o.PciSwitch = v
}

// GetPersistentMemoryConfiguration returns the PersistentMemoryConfiguration field value if set, zero value otherwise.
func (o *ComputeBoardAllOf) GetPersistentMemoryConfiguration() MemoryPersistentMemoryConfigurationRelationship {
	if o == nil || o.PersistentMemoryConfiguration == nil {
		var ret MemoryPersistentMemoryConfigurationRelationship
		return ret
	}
	return *o.PersistentMemoryConfiguration
}

// GetPersistentMemoryConfigurationOk returns a tuple with the PersistentMemoryConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBoardAllOf) GetPersistentMemoryConfigurationOk() (*MemoryPersistentMemoryConfigurationRelationship, bool) {
	if o == nil || o.PersistentMemoryConfiguration == nil {
		return nil, false
	}
	return o.PersistentMemoryConfiguration, true
}

// HasPersistentMemoryConfiguration returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasPersistentMemoryConfiguration() bool {
	if o != nil && o.PersistentMemoryConfiguration != nil {
		return true
	}

	return false
}

// SetPersistentMemoryConfiguration gets a reference to the given MemoryPersistentMemoryConfigurationRelationship and assigns it to the PersistentMemoryConfiguration field.
func (o *ComputeBoardAllOf) SetPersistentMemoryConfiguration(v MemoryPersistentMemoryConfigurationRelationship) {
	o.PersistentMemoryConfiguration = &v
}

// GetProcessors returns the Processors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBoardAllOf) GetProcessors() []ProcessorUnitRelationship {
	if o == nil  {
		var ret []ProcessorUnitRelationship
		return ret
	}
	return o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBoardAllOf) GetProcessorsOk() (*[]ProcessorUnitRelationship, bool) {
	if o == nil || o.Processors == nil {
		return nil, false
	}
	return &o.Processors, true
}

// HasProcessors returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasProcessors() bool {
	if o != nil && o.Processors != nil {
		return true
	}

	return false
}

// SetProcessors gets a reference to the given []ProcessorUnitRelationship and assigns it to the Processors field.
func (o *ComputeBoardAllOf) SetProcessors(v []ProcessorUnitRelationship) {
	o.Processors = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ComputeBoardAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBoardAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ComputeBoardAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetSecurityUnits returns the SecurityUnits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBoardAllOf) GetSecurityUnits() []SecurityUnitRelationship {
	if o == nil  {
		var ret []SecurityUnitRelationship
		return ret
	}
	return o.SecurityUnits
}

// GetSecurityUnitsOk returns a tuple with the SecurityUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBoardAllOf) GetSecurityUnitsOk() (*[]SecurityUnitRelationship, bool) {
	if o == nil || o.SecurityUnits == nil {
		return nil, false
	}
	return &o.SecurityUnits, true
}

// HasSecurityUnits returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasSecurityUnits() bool {
	if o != nil && o.SecurityUnits != nil {
		return true
	}

	return false
}

// SetSecurityUnits gets a reference to the given []SecurityUnitRelationship and assigns it to the SecurityUnits field.
func (o *ComputeBoardAllOf) SetSecurityUnits(v []SecurityUnitRelationship) {
	o.SecurityUnits = v
}

// GetStorageControllers returns the StorageControllers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBoardAllOf) GetStorageControllers() []StorageControllerRelationship {
	if o == nil  {
		var ret []StorageControllerRelationship
		return ret
	}
	return o.StorageControllers
}

// GetStorageControllersOk returns a tuple with the StorageControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBoardAllOf) GetStorageControllersOk() (*[]StorageControllerRelationship, bool) {
	if o == nil || o.StorageControllers == nil {
		return nil, false
	}
	return &o.StorageControllers, true
}

// HasStorageControllers returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasStorageControllers() bool {
	if o != nil && o.StorageControllers != nil {
		return true
	}

	return false
}

// SetStorageControllers gets a reference to the given []StorageControllerRelationship and assigns it to the StorageControllers field.
func (o *ComputeBoardAllOf) SetStorageControllers(v []StorageControllerRelationship) {
	o.StorageControllers = v
}

// GetStorageFlexFlashControllers returns the StorageFlexFlashControllers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBoardAllOf) GetStorageFlexFlashControllers() []StorageFlexFlashControllerRelationship {
	if o == nil  {
		var ret []StorageFlexFlashControllerRelationship
		return ret
	}
	return o.StorageFlexFlashControllers
}

// GetStorageFlexFlashControllersOk returns a tuple with the StorageFlexFlashControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBoardAllOf) GetStorageFlexFlashControllersOk() (*[]StorageFlexFlashControllerRelationship, bool) {
	if o == nil || o.StorageFlexFlashControllers == nil {
		return nil, false
	}
	return &o.StorageFlexFlashControllers, true
}

// HasStorageFlexFlashControllers returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasStorageFlexFlashControllers() bool {
	if o != nil && o.StorageFlexFlashControllers != nil {
		return true
	}

	return false
}

// SetStorageFlexFlashControllers gets a reference to the given []StorageFlexFlashControllerRelationship and assigns it to the StorageFlexFlashControllers field.
func (o *ComputeBoardAllOf) SetStorageFlexFlashControllers(v []StorageFlexFlashControllerRelationship) {
	o.StorageFlexFlashControllers = v
}

// GetStorageFlexUtilControllers returns the StorageFlexUtilControllers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBoardAllOf) GetStorageFlexUtilControllers() []StorageFlexUtilControllerRelationship {
	if o == nil  {
		var ret []StorageFlexUtilControllerRelationship
		return ret
	}
	return o.StorageFlexUtilControllers
}

// GetStorageFlexUtilControllersOk returns a tuple with the StorageFlexUtilControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBoardAllOf) GetStorageFlexUtilControllersOk() (*[]StorageFlexUtilControllerRelationship, bool) {
	if o == nil || o.StorageFlexUtilControllers == nil {
		return nil, false
	}
	return &o.StorageFlexUtilControllers, true
}

// HasStorageFlexUtilControllers returns a boolean if a field has been set.
func (o *ComputeBoardAllOf) HasStorageFlexUtilControllers() bool {
	if o != nil && o.StorageFlexUtilControllers != nil {
		return true
	}

	return false
}

// SetStorageFlexUtilControllers gets a reference to the given []StorageFlexUtilControllerRelationship and assigns it to the StorageFlexUtilControllers field.
func (o *ComputeBoardAllOf) SetStorageFlexUtilControllers(v []StorageFlexUtilControllerRelationship) {
	o.StorageFlexUtilControllers = v
}

func (o ComputeBoardAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BoardId != nil {
		toSerialize["BoardId"] = o.BoardId
	}
	if o.CpuTypeController != nil {
		toSerialize["CpuTypeController"] = o.CpuTypeController
	}
	if o.OperPowerState != nil {
		toSerialize["OperPowerState"] = o.OperPowerState
	}
	if o.OperReason != nil {
		toSerialize["OperReason"] = o.OperReason
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.EquipmentTpms != nil {
		toSerialize["EquipmentTpms"] = o.EquipmentTpms
	}
	if o.GraphicsCards != nil {
		toSerialize["GraphicsCards"] = o.GraphicsCards
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.MemoryArrays != nil {
		toSerialize["MemoryArrays"] = o.MemoryArrays
	}
	if o.PciCoprocessorCards != nil {
		toSerialize["PciCoprocessorCards"] = o.PciCoprocessorCards
	}
	if o.PciSwitch != nil {
		toSerialize["PciSwitch"] = o.PciSwitch
	}
	if o.PersistentMemoryConfiguration != nil {
		toSerialize["PersistentMemoryConfiguration"] = o.PersistentMemoryConfiguration
	}
	if o.Processors != nil {
		toSerialize["Processors"] = o.Processors
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.SecurityUnits != nil {
		toSerialize["SecurityUnits"] = o.SecurityUnits
	}
	if o.StorageControllers != nil {
		toSerialize["StorageControllers"] = o.StorageControllers
	}
	if o.StorageFlexFlashControllers != nil {
		toSerialize["StorageFlexFlashControllers"] = o.StorageFlexFlashControllers
	}
	if o.StorageFlexUtilControllers != nil {
		toSerialize["StorageFlexUtilControllers"] = o.StorageFlexUtilControllers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeBoardAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputeBoardAllOf := _ComputeBoardAllOf{}

	if err = json.Unmarshal(bytes, &varComputeBoardAllOf); err == nil {
		*o = ComputeBoardAllOf(varComputeBoardAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BoardId")
		delete(additionalProperties, "CpuTypeController")
		delete(additionalProperties, "OperPowerState")
		delete(additionalProperties, "OperReason")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "EquipmentTpms")
		delete(additionalProperties, "GraphicsCards")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "MemoryArrays")
		delete(additionalProperties, "PciCoprocessorCards")
		delete(additionalProperties, "PciSwitch")
		delete(additionalProperties, "PersistentMemoryConfiguration")
		delete(additionalProperties, "Processors")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "SecurityUnits")
		delete(additionalProperties, "StorageControllers")
		delete(additionalProperties, "StorageFlexFlashControllers")
		delete(additionalProperties, "StorageFlexUtilControllers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeBoardAllOf struct {
	value *ComputeBoardAllOf
	isSet bool
}

func (v NullableComputeBoardAllOf) Get() *ComputeBoardAllOf {
	return v.value
}

func (v *NullableComputeBoardAllOf) Set(val *ComputeBoardAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeBoardAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeBoardAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeBoardAllOf(val *ComputeBoardAllOf) *NullableComputeBoardAllOf {
	return &NullableComputeBoardAllOf{value: val, isSet: true}
}

func (v NullableComputeBoardAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeBoardAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


