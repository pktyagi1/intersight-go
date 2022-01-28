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

// StorageStoragePolicyAllOf Definition of the list of properties defined in 'storage.StoragePolicy', excluding properties defined in parent classes.
type StorageStoragePolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A collection of disks that is to be used as hot spares, globally, for all the RAID groups. Allowed value is a number range separated by a comma or a hyphen.
	GlobalHotSpares *string `json:"GlobalHotSpares,omitempty"`
	M2VirtualDrive NullableStorageM2VirtualDriveConfig `json:"M2VirtualDrive,omitempty"`
	Raid0Drive NullableStorageR0Drive `json:"Raid0Drive,omitempty"`
	// State to which disks, not used in this policy, are to be moved. NoChange will not change the drive state. * `NoChange` - Drive state will not be modified by Storage Policy. * `UnconfiguredGood` - Unconfigured good state -ready to be added in a RAID group. * `Jbod` - JBOD state where the disks start showing up to Host OS.
	UnusedDisksState *string `json:"UnusedDisksState,omitempty"`
	// Disks in JBOD State are used to create virtual drives.
	UseJbodForVdCreation *bool `json:"UseJbodForVdCreation,omitempty"`
	// An array of relationships to storageDriveGroup resources.
	DriveGroup []StorageDriveGroupRelationship `json:"DriveGroup,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageStoragePolicyAllOf StorageStoragePolicyAllOf

// NewStorageStoragePolicyAllOf instantiates a new StorageStoragePolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageStoragePolicyAllOf(classId string, objectType string) *StorageStoragePolicyAllOf {
	this := StorageStoragePolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var unusedDisksState string = "NoChange"
	this.UnusedDisksState = &unusedDisksState
	return &this
}

// NewStorageStoragePolicyAllOfWithDefaults instantiates a new StorageStoragePolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageStoragePolicyAllOfWithDefaults() *StorageStoragePolicyAllOf {
	this := StorageStoragePolicyAllOf{}
	var classId string = "storage.StoragePolicy"
	this.ClassId = classId
	var objectType string = "storage.StoragePolicy"
	this.ObjectType = objectType
	var unusedDisksState string = "NoChange"
	this.UnusedDisksState = &unusedDisksState
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageStoragePolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageStoragePolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageStoragePolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageStoragePolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGlobalHotSpares returns the GlobalHotSpares field value if set, zero value otherwise.
func (o *StorageStoragePolicyAllOf) GetGlobalHotSpares() string {
	if o == nil || o.GlobalHotSpares == nil {
		var ret string
		return ret
	}
	return *o.GlobalHotSpares
}

// GetGlobalHotSparesOk returns a tuple with the GlobalHotSpares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicyAllOf) GetGlobalHotSparesOk() (*string, bool) {
	if o == nil || o.GlobalHotSpares == nil {
		return nil, false
	}
	return o.GlobalHotSpares, true
}

// HasGlobalHotSpares returns a boolean if a field has been set.
func (o *StorageStoragePolicyAllOf) HasGlobalHotSpares() bool {
	if o != nil && o.GlobalHotSpares != nil {
		return true
	}

	return false
}

// SetGlobalHotSpares gets a reference to the given string and assigns it to the GlobalHotSpares field.
func (o *StorageStoragePolicyAllOf) SetGlobalHotSpares(v string) {
	o.GlobalHotSpares = &v
}

// GetM2VirtualDrive returns the M2VirtualDrive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageStoragePolicyAllOf) GetM2VirtualDrive() StorageM2VirtualDriveConfig {
	if o == nil || o.M2VirtualDrive.Get() == nil {
		var ret StorageM2VirtualDriveConfig
		return ret
	}
	return *o.M2VirtualDrive.Get()
}

// GetM2VirtualDriveOk returns a tuple with the M2VirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageStoragePolicyAllOf) GetM2VirtualDriveOk() (*StorageM2VirtualDriveConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.M2VirtualDrive.Get(), o.M2VirtualDrive.IsSet()
}

// HasM2VirtualDrive returns a boolean if a field has been set.
func (o *StorageStoragePolicyAllOf) HasM2VirtualDrive() bool {
	if o != nil && o.M2VirtualDrive.IsSet() {
		return true
	}

	return false
}

// SetM2VirtualDrive gets a reference to the given NullableStorageM2VirtualDriveConfig and assigns it to the M2VirtualDrive field.
func (o *StorageStoragePolicyAllOf) SetM2VirtualDrive(v StorageM2VirtualDriveConfig) {
	o.M2VirtualDrive.Set(&v)
}
// SetM2VirtualDriveNil sets the value for M2VirtualDrive to be an explicit nil
func (o *StorageStoragePolicyAllOf) SetM2VirtualDriveNil() {
	o.M2VirtualDrive.Set(nil)
}

// UnsetM2VirtualDrive ensures that no value is present for M2VirtualDrive, not even an explicit nil
func (o *StorageStoragePolicyAllOf) UnsetM2VirtualDrive() {
	o.M2VirtualDrive.Unset()
}

// GetRaid0Drive returns the Raid0Drive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageStoragePolicyAllOf) GetRaid0Drive() StorageR0Drive {
	if o == nil || o.Raid0Drive.Get() == nil {
		var ret StorageR0Drive
		return ret
	}
	return *o.Raid0Drive.Get()
}

// GetRaid0DriveOk returns a tuple with the Raid0Drive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageStoragePolicyAllOf) GetRaid0DriveOk() (*StorageR0Drive, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Raid0Drive.Get(), o.Raid0Drive.IsSet()
}

// HasRaid0Drive returns a boolean if a field has been set.
func (o *StorageStoragePolicyAllOf) HasRaid0Drive() bool {
	if o != nil && o.Raid0Drive.IsSet() {
		return true
	}

	return false
}

// SetRaid0Drive gets a reference to the given NullableStorageR0Drive and assigns it to the Raid0Drive field.
func (o *StorageStoragePolicyAllOf) SetRaid0Drive(v StorageR0Drive) {
	o.Raid0Drive.Set(&v)
}
// SetRaid0DriveNil sets the value for Raid0Drive to be an explicit nil
func (o *StorageStoragePolicyAllOf) SetRaid0DriveNil() {
	o.Raid0Drive.Set(nil)
}

// UnsetRaid0Drive ensures that no value is present for Raid0Drive, not even an explicit nil
func (o *StorageStoragePolicyAllOf) UnsetRaid0Drive() {
	o.Raid0Drive.Unset()
}

// GetUnusedDisksState returns the UnusedDisksState field value if set, zero value otherwise.
func (o *StorageStoragePolicyAllOf) GetUnusedDisksState() string {
	if o == nil || o.UnusedDisksState == nil {
		var ret string
		return ret
	}
	return *o.UnusedDisksState
}

// GetUnusedDisksStateOk returns a tuple with the UnusedDisksState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicyAllOf) GetUnusedDisksStateOk() (*string, bool) {
	if o == nil || o.UnusedDisksState == nil {
		return nil, false
	}
	return o.UnusedDisksState, true
}

// HasUnusedDisksState returns a boolean if a field has been set.
func (o *StorageStoragePolicyAllOf) HasUnusedDisksState() bool {
	if o != nil && o.UnusedDisksState != nil {
		return true
	}

	return false
}

// SetUnusedDisksState gets a reference to the given string and assigns it to the UnusedDisksState field.
func (o *StorageStoragePolicyAllOf) SetUnusedDisksState(v string) {
	o.UnusedDisksState = &v
}

// GetUseJbodForVdCreation returns the UseJbodForVdCreation field value if set, zero value otherwise.
func (o *StorageStoragePolicyAllOf) GetUseJbodForVdCreation() bool {
	if o == nil || o.UseJbodForVdCreation == nil {
		var ret bool
		return ret
	}
	return *o.UseJbodForVdCreation
}

// GetUseJbodForVdCreationOk returns a tuple with the UseJbodForVdCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicyAllOf) GetUseJbodForVdCreationOk() (*bool, bool) {
	if o == nil || o.UseJbodForVdCreation == nil {
		return nil, false
	}
	return o.UseJbodForVdCreation, true
}

// HasUseJbodForVdCreation returns a boolean if a field has been set.
func (o *StorageStoragePolicyAllOf) HasUseJbodForVdCreation() bool {
	if o != nil && o.UseJbodForVdCreation != nil {
		return true
	}

	return false
}

// SetUseJbodForVdCreation gets a reference to the given bool and assigns it to the UseJbodForVdCreation field.
func (o *StorageStoragePolicyAllOf) SetUseJbodForVdCreation(v bool) {
	o.UseJbodForVdCreation = &v
}

// GetDriveGroup returns the DriveGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageStoragePolicyAllOf) GetDriveGroup() []StorageDriveGroupRelationship {
	if o == nil  {
		var ret []StorageDriveGroupRelationship
		return ret
	}
	return o.DriveGroup
}

// GetDriveGroupOk returns a tuple with the DriveGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageStoragePolicyAllOf) GetDriveGroupOk() (*[]StorageDriveGroupRelationship, bool) {
	if o == nil || o.DriveGroup == nil {
		return nil, false
	}
	return &o.DriveGroup, true
}

// HasDriveGroup returns a boolean if a field has been set.
func (o *StorageStoragePolicyAllOf) HasDriveGroup() bool {
	if o != nil && o.DriveGroup != nil {
		return true
	}

	return false
}

// SetDriveGroup gets a reference to the given []StorageDriveGroupRelationship and assigns it to the DriveGroup field.
func (o *StorageStoragePolicyAllOf) SetDriveGroup(v []StorageDriveGroupRelationship) {
	o.DriveGroup = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *StorageStoragePolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *StorageStoragePolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *StorageStoragePolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageStoragePolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil  {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageStoragePolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *StorageStoragePolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *StorageStoragePolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o StorageStoragePolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.GlobalHotSpares != nil {
		toSerialize["GlobalHotSpares"] = o.GlobalHotSpares
	}
	if o.M2VirtualDrive.IsSet() {
		toSerialize["M2VirtualDrive"] = o.M2VirtualDrive.Get()
	}
	if o.Raid0Drive.IsSet() {
		toSerialize["Raid0Drive"] = o.Raid0Drive.Get()
	}
	if o.UnusedDisksState != nil {
		toSerialize["UnusedDisksState"] = o.UnusedDisksState
	}
	if o.UseJbodForVdCreation != nil {
		toSerialize["UseJbodForVdCreation"] = o.UseJbodForVdCreation
	}
	if o.DriveGroup != nil {
		toSerialize["DriveGroup"] = o.DriveGroup
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageStoragePolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageStoragePolicyAllOf := _StorageStoragePolicyAllOf{}

	if err = json.Unmarshal(bytes, &varStorageStoragePolicyAllOf); err == nil {
		*o = StorageStoragePolicyAllOf(varStorageStoragePolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "GlobalHotSpares")
		delete(additionalProperties, "M2VirtualDrive")
		delete(additionalProperties, "Raid0Drive")
		delete(additionalProperties, "UnusedDisksState")
		delete(additionalProperties, "UseJbodForVdCreation")
		delete(additionalProperties, "DriveGroup")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageStoragePolicyAllOf struct {
	value *StorageStoragePolicyAllOf
	isSet bool
}

func (v NullableStorageStoragePolicyAllOf) Get() *StorageStoragePolicyAllOf {
	return v.value
}

func (v *NullableStorageStoragePolicyAllOf) Set(val *StorageStoragePolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageStoragePolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageStoragePolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageStoragePolicyAllOf(val *StorageStoragePolicyAllOf) *NullableStorageStoragePolicyAllOf {
	return &NullableStorageStoragePolicyAllOf{value: val, isSet: true}
}

func (v NullableStorageStoragePolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageStoragePolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


