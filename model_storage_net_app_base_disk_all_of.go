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

// StorageNetAppBaseDiskAllOf Definition of the list of properties defined in 'storage.NetAppBaseDisk', excluding properties defined in parent classes.
type StorageNetAppBaseDiskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The NetApp base disk model.
	BaseDiskModel *string `json:"BaseDiskModel,omitempty"`
	// Supported container type for NetApp disk. * `Unknown` - Default container type is currently unknown. * `Aggregate` - Disk is used as a physical disk in an aggregate. * `Broken` - Disk is in a broken pool. * `Label Maintenance` - Disk is in online label maintenance list. * `Foreign` - Array LUN has been marked foreign. * `Maintenance` - Disk is in maintenance center. * `Mediator` - A mediator disk is a disk used on non-shared HA systems hosted by an external node which is used to communicate the viability of the storage failover between non-shared HA nodes. * `Shared` - Disk is partitioned or in a storage pool. * `Remote` - Disk belongs to a remote cluster. * `Spare` - The disk is a spare disk. * `Unassigned` - Disk ownership has not been assigned. * `Unsupported` - The disk is not supported.
	ContainerType *string `json:"ContainerType,omitempty"`
	// The type of the NetApp disk. * `Unknown` - Default unknown disk type. * `SSDNVM` - Solid state disk with Non-Volatile Memory Express protocol enabled. * `ATA` - Advanced Technology Attachment is a type of disk drive that integrates the drive controller directly on the drive itself. * `FCAL` - For the FC-AL disk connection type, disk shelves are connected to the controller in a loop. * `BSAS` - Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf. * `FSAS` - Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, and rotational speed of traditional enterprise-class SATA drives with the fully capable SAS interface typical for classic SAS drives. * `LUN` - Logical Unit Number refers to a logical disk. * `SAS` - Storage disk with serial attached SCSI. * `MSATA` - SATA disk in multi-disk carrier storage shelf. * `SSD` - Storage disk with Solid state disk. * `VMDISK` - Virtual machine Data Disk.
	DiskType *string `json:"DiskType,omitempty"`
	// Current state of the NetApp disk. * `Present` - Storage disk state type is present. * `Copy` - Storage disk state type is copy. * `Broken` - Storage disk state type is broken. * `Maintenance` - Storage disk state type is maintenance. * `Partner` - Storage disk state type is partner. * `Pending` - Storage disk state type is pending. * `Reconstructing` - Storage disk state type is reconstructing. * `Removed` - Storage disk state type is removed. * `Spare` - Storage disk state type is spare. * `Unfail` - Storage disk state type is unfail. * `Zeroing` - Storage disk state type is zeroing.
	State *string `json:"State,omitempty"`
	// Universally unique identifier of the NetApp Disk.
	Uuid *string `json:"Uuid,omitempty"`
	Array *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	ArrayController *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	// An array of relationships to storageNetAppAggregate resources.
	DiskPool []StorageNetAppAggregateRelationship `json:"DiskPool,omitempty"`
	// An array of relationships to storageNetAppDiskEvent resources.
	Events []StorageNetAppDiskEventRelationship `json:"Events,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppBaseDiskAllOf StorageNetAppBaseDiskAllOf

// NewStorageNetAppBaseDiskAllOf instantiates a new StorageNetAppBaseDiskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppBaseDiskAllOf(classId string, objectType string) *StorageNetAppBaseDiskAllOf {
	this := StorageNetAppBaseDiskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppBaseDiskAllOfWithDefaults instantiates a new StorageNetAppBaseDiskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppBaseDiskAllOfWithDefaults() *StorageNetAppBaseDiskAllOf {
	this := StorageNetAppBaseDiskAllOf{}
	var classId string = "storage.NetAppBaseDisk"
	this.ClassId = classId
	var objectType string = "storage.NetAppBaseDisk"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppBaseDiskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDiskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppBaseDiskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppBaseDiskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDiskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppBaseDiskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBaseDiskModel returns the BaseDiskModel field value if set, zero value otherwise.
func (o *StorageNetAppBaseDiskAllOf) GetBaseDiskModel() string {
	if o == nil || o.BaseDiskModel == nil {
		var ret string
		return ret
	}
	return *o.BaseDiskModel
}

// GetBaseDiskModelOk returns a tuple with the BaseDiskModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDiskAllOf) GetBaseDiskModelOk() (*string, bool) {
	if o == nil || o.BaseDiskModel == nil {
		return nil, false
	}
	return o.BaseDiskModel, true
}

// HasBaseDiskModel returns a boolean if a field has been set.
func (o *StorageNetAppBaseDiskAllOf) HasBaseDiskModel() bool {
	if o != nil && o.BaseDiskModel != nil {
		return true
	}

	return false
}

// SetBaseDiskModel gets a reference to the given string and assigns it to the BaseDiskModel field.
func (o *StorageNetAppBaseDiskAllOf) SetBaseDiskModel(v string) {
	o.BaseDiskModel = &v
}

// GetContainerType returns the ContainerType field value if set, zero value otherwise.
func (o *StorageNetAppBaseDiskAllOf) GetContainerType() string {
	if o == nil || o.ContainerType == nil {
		var ret string
		return ret
	}
	return *o.ContainerType
}

// GetContainerTypeOk returns a tuple with the ContainerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDiskAllOf) GetContainerTypeOk() (*string, bool) {
	if o == nil || o.ContainerType == nil {
		return nil, false
	}
	return o.ContainerType, true
}

// HasContainerType returns a boolean if a field has been set.
func (o *StorageNetAppBaseDiskAllOf) HasContainerType() bool {
	if o != nil && o.ContainerType != nil {
		return true
	}

	return false
}

// SetContainerType gets a reference to the given string and assigns it to the ContainerType field.
func (o *StorageNetAppBaseDiskAllOf) SetContainerType(v string) {
	o.ContainerType = &v
}

// GetDiskType returns the DiskType field value if set, zero value otherwise.
func (o *StorageNetAppBaseDiskAllOf) GetDiskType() string {
	if o == nil || o.DiskType == nil {
		var ret string
		return ret
	}
	return *o.DiskType
}

// GetDiskTypeOk returns a tuple with the DiskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDiskAllOf) GetDiskTypeOk() (*string, bool) {
	if o == nil || o.DiskType == nil {
		return nil, false
	}
	return o.DiskType, true
}

// HasDiskType returns a boolean if a field has been set.
func (o *StorageNetAppBaseDiskAllOf) HasDiskType() bool {
	if o != nil && o.DiskType != nil {
		return true
	}

	return false
}

// SetDiskType gets a reference to the given string and assigns it to the DiskType field.
func (o *StorageNetAppBaseDiskAllOf) SetDiskType(v string) {
	o.DiskType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppBaseDiskAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDiskAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppBaseDiskAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppBaseDiskAllOf) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppBaseDiskAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDiskAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppBaseDiskAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppBaseDiskAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppBaseDiskAllOf) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDiskAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppBaseDiskAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppBaseDiskAllOf) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise.
func (o *StorageNetAppBaseDiskAllOf) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || o.ArrayController == nil {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDiskAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil || o.ArrayController == nil {
		return nil, false
	}
	return o.ArrayController, true
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppBaseDiskAllOf) HasArrayController() bool {
	if o != nil && o.ArrayController != nil {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given StorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppBaseDiskAllOf) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController = &v
}

// GetDiskPool returns the DiskPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppBaseDiskAllOf) GetDiskPool() []StorageNetAppAggregateRelationship {
	if o == nil  {
		var ret []StorageNetAppAggregateRelationship
		return ret
	}
	return o.DiskPool
}

// GetDiskPoolOk returns a tuple with the DiskPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppBaseDiskAllOf) GetDiskPoolOk() (*[]StorageNetAppAggregateRelationship, bool) {
	if o == nil || o.DiskPool == nil {
		return nil, false
	}
	return &o.DiskPool, true
}

// HasDiskPool returns a boolean if a field has been set.
func (o *StorageNetAppBaseDiskAllOf) HasDiskPool() bool {
	if o != nil && o.DiskPool != nil {
		return true
	}

	return false
}

// SetDiskPool gets a reference to the given []StorageNetAppAggregateRelationship and assigns it to the DiskPool field.
func (o *StorageNetAppBaseDiskAllOf) SetDiskPool(v []StorageNetAppAggregateRelationship) {
	o.DiskPool = v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppBaseDiskAllOf) GetEvents() []StorageNetAppDiskEventRelationship {
	if o == nil  {
		var ret []StorageNetAppDiskEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppBaseDiskAllOf) GetEventsOk() (*[]StorageNetAppDiskEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppBaseDiskAllOf) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppDiskEventRelationship and assigns it to the Events field.
func (o *StorageNetAppBaseDiskAllOf) SetEvents(v []StorageNetAppDiskEventRelationship) {
	o.Events = v
}

func (o StorageNetAppBaseDiskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BaseDiskModel != nil {
		toSerialize["BaseDiskModel"] = o.BaseDiskModel
	}
	if o.ContainerType != nil {
		toSerialize["ContainerType"] = o.ContainerType
	}
	if o.DiskType != nil {
		toSerialize["DiskType"] = o.DiskType
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.ArrayController != nil {
		toSerialize["ArrayController"] = o.ArrayController
	}
	if o.DiskPool != nil {
		toSerialize["DiskPool"] = o.DiskPool
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppBaseDiskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppBaseDiskAllOf := _StorageNetAppBaseDiskAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppBaseDiskAllOf); err == nil {
		*o = StorageNetAppBaseDiskAllOf(varStorageNetAppBaseDiskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BaseDiskModel")
		delete(additionalProperties, "ContainerType")
		delete(additionalProperties, "DiskType")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ArrayController")
		delete(additionalProperties, "DiskPool")
		delete(additionalProperties, "Events")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppBaseDiskAllOf struct {
	value *StorageNetAppBaseDiskAllOf
	isSet bool
}

func (v NullableStorageNetAppBaseDiskAllOf) Get() *StorageNetAppBaseDiskAllOf {
	return v.value
}

func (v *NullableStorageNetAppBaseDiskAllOf) Set(val *StorageNetAppBaseDiskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppBaseDiskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppBaseDiskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppBaseDiskAllOf(val *StorageNetAppBaseDiskAllOf) *NullableStorageNetAppBaseDiskAllOf {
	return &NullableStorageNetAppBaseDiskAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppBaseDiskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppBaseDiskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


