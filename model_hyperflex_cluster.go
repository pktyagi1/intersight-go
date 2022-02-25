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

// HyperflexCluster A HyperFlex cluster. Contains inventory information concerning the health, software versions, storage, and nodes of the cluster.
type HyperflexCluster struct {
	HyperflexBaseCluster
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The storage type of this cluster (All Flash or Hybrid).
	// Deprecated
	ClusterType *int64 `json:"ClusterType,omitempty"`
	// The unique identifier for this HyperFlex cluster.
	ClusterUuid *string `json:"ClusterUuid,omitempty"`
	// The unique identifier of the device registration that represents this HyperFlex cluster's connection to Intersight.
	DeviceId *string `json:"DeviceId,omitempty"`
	DnsServers []string `json:"DnsServers,omitempty"`
	// This captures the encryption status for a HyperFlex cluster. Currently it will have the status if HXA-CLU-0020 alarm is raised. In the future it can capture other details.
	EncryptionStatus *string `json:"EncryptionStatus,omitempty"`
	// The number of yellow (warning) and red (critical) alarms stored as an aggregate. The first 16 bits indicate the number of red alarms, and the last 16 bits contain the number of yellow alarms.
	// Deprecated
	FltAggr *int64 `json:"FltAggr,omitempty"`
	// The version and build number of the HyperFlex Data Platform for this cluster. After a cluster upgrade, this version string will be updated on the next inventory cycle to reflect the newly installed version.
	HxdpBuildVersion *string `json:"HxdpBuildVersion,omitempty"`
	NtpServers []string `json:"NtpServers,omitempty"`
	Summary NullableHyperflexSummary `json:"Summary,omitempty"`
	// The upgrade status of the HyperFlex cluster. * `Unknown` - The upgrade status of the HyperFlex cluster could not be determined. * `Ok` - The upgrade of the HyperFlex cluster is complete. * `InProgress` - The upgrade of the HyperFlex cluster is in-progress. * `Failed` - The upgrade of the HyperFlex cluster has failed. * `Waiting` - The upgrade of the HyperFlex cluster is waiting to continue execution.
	UpgradeStatus *string `json:"UpgradeStatus,omitempty"`
	// The number of virtual machines present on this cluster.
	VmCount *int64 `json:"VmCount,omitempty"`
	// An array of relationships to hyperflexAlarm resources.
	Alarm []HyperflexAlarmRelationship `json:"Alarm,omitempty"`
	Encryption *HyperflexEncryptionRelationship `json:"Encryption,omitempty"`
	Health *HyperflexHealthRelationship `json:"Health,omitempty"`
	License *HyperflexLicenseRelationship `json:"License,omitempty"`
	// An array of relationships to hyperflexNode resources.
	Nodes []HyperflexNodeRelationship `json:"Nodes,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to ippoolPool resources.
	StorageClientIpPools []IppoolPoolRelationship `json:"StorageClientIpPools,omitempty"`
	StorageClientVrf *VrfVrfRelationship `json:"StorageClientVrf,omitempty"`
	// An array of relationships to storageHyperFlexStorageContainer resources.
	StorageContainers []StorageHyperFlexStorageContainerRelationship `json:"StorageContainers,omitempty"`
	// An array of relationships to storageHyperFlexVolume resources.
	Volumes []StorageHyperFlexVolumeRelationship `json:"Volumes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexCluster HyperflexCluster

// NewHyperflexCluster instantiates a new HyperflexCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexCluster(classId string, objectType string) *HyperflexCluster {
	this := HyperflexCluster{}
	this.ClassId = classId
	this.ObjectType = objectType
	var clusterPurpose string = "Storage"
	this.ClusterPurpose = &clusterPurpose
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	return &this
}

// NewHyperflexClusterWithDefaults instantiates a new HyperflexCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterWithDefaults() *HyperflexCluster {
	this := HyperflexCluster{}
	var classId string = "hyperflex.Cluster"
	this.ClassId = classId
	var objectType string = "hyperflex.Cluster"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexCluster) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexCluster) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexCluster) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexCluster) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterType returns the ClusterType field value if set, zero value otherwise.
// Deprecated
func (o *HyperflexCluster) GetClusterType() int64 {
	if o == nil || o.ClusterType == nil {
		var ret int64
		return ret
	}
	return *o.ClusterType
}

// GetClusterTypeOk returns a tuple with the ClusterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *HyperflexCluster) GetClusterTypeOk() (*int64, bool) {
	if o == nil || o.ClusterType == nil {
		return nil, false
	}
	return o.ClusterType, true
}

// HasClusterType returns a boolean if a field has been set.
func (o *HyperflexCluster) HasClusterType() bool {
	if o != nil && o.ClusterType != nil {
		return true
	}

	return false
}

// SetClusterType gets a reference to the given int64 and assigns it to the ClusterType field.
// Deprecated
func (o *HyperflexCluster) SetClusterType(v int64) {
	o.ClusterType = &v
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *HyperflexCluster) GetClusterUuid() string {
	if o == nil || o.ClusterUuid == nil {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetClusterUuidOk() (*string, bool) {
	if o == nil || o.ClusterUuid == nil {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *HyperflexCluster) HasClusterUuid() bool {
	if o != nil && o.ClusterUuid != nil {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *HyperflexCluster) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *HyperflexCluster) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *HyperflexCluster) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *HyperflexCluster) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexCluster) GetDnsServers() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexCluster) GetDnsServersOk() (*[]string, bool) {
	if o == nil || o.DnsServers == nil {
		return nil, false
	}
	return &o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *HyperflexCluster) HasDnsServers() bool {
	if o != nil && o.DnsServers != nil {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
func (o *HyperflexCluster) SetDnsServers(v []string) {
	o.DnsServers = v
}

// GetEncryptionStatus returns the EncryptionStatus field value if set, zero value otherwise.
func (o *HyperflexCluster) GetEncryptionStatus() string {
	if o == nil || o.EncryptionStatus == nil {
		var ret string
		return ret
	}
	return *o.EncryptionStatus
}

// GetEncryptionStatusOk returns a tuple with the EncryptionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetEncryptionStatusOk() (*string, bool) {
	if o == nil || o.EncryptionStatus == nil {
		return nil, false
	}
	return o.EncryptionStatus, true
}

// HasEncryptionStatus returns a boolean if a field has been set.
func (o *HyperflexCluster) HasEncryptionStatus() bool {
	if o != nil && o.EncryptionStatus != nil {
		return true
	}

	return false
}

// SetEncryptionStatus gets a reference to the given string and assigns it to the EncryptionStatus field.
func (o *HyperflexCluster) SetEncryptionStatus(v string) {
	o.EncryptionStatus = &v
}

// GetFltAggr returns the FltAggr field value if set, zero value otherwise.
// Deprecated
func (o *HyperflexCluster) GetFltAggr() int64 {
	if o == nil || o.FltAggr == nil {
		var ret int64
		return ret
	}
	return *o.FltAggr
}

// GetFltAggrOk returns a tuple with the FltAggr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *HyperflexCluster) GetFltAggrOk() (*int64, bool) {
	if o == nil || o.FltAggr == nil {
		return nil, false
	}
	return o.FltAggr, true
}

// HasFltAggr returns a boolean if a field has been set.
func (o *HyperflexCluster) HasFltAggr() bool {
	if o != nil && o.FltAggr != nil {
		return true
	}

	return false
}

// SetFltAggr gets a reference to the given int64 and assigns it to the FltAggr field.
// Deprecated
func (o *HyperflexCluster) SetFltAggr(v int64) {
	o.FltAggr = &v
}

// GetHxdpBuildVersion returns the HxdpBuildVersion field value if set, zero value otherwise.
func (o *HyperflexCluster) GetHxdpBuildVersion() string {
	if o == nil || o.HxdpBuildVersion == nil {
		var ret string
		return ret
	}
	return *o.HxdpBuildVersion
}

// GetHxdpBuildVersionOk returns a tuple with the HxdpBuildVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetHxdpBuildVersionOk() (*string, bool) {
	if o == nil || o.HxdpBuildVersion == nil {
		return nil, false
	}
	return o.HxdpBuildVersion, true
}

// HasHxdpBuildVersion returns a boolean if a field has been set.
func (o *HyperflexCluster) HasHxdpBuildVersion() bool {
	if o != nil && o.HxdpBuildVersion != nil {
		return true
	}

	return false
}

// SetHxdpBuildVersion gets a reference to the given string and assigns it to the HxdpBuildVersion field.
func (o *HyperflexCluster) SetHxdpBuildVersion(v string) {
	o.HxdpBuildVersion = &v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexCluster) GetNtpServers() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexCluster) GetNtpServersOk() (*[]string, bool) {
	if o == nil || o.NtpServers == nil {
		return nil, false
	}
	return &o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *HyperflexCluster) HasNtpServers() bool {
	if o != nil && o.NtpServers != nil {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *HyperflexCluster) SetNtpServers(v []string) {
	o.NtpServers = v
}

// GetSummary returns the Summary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexCluster) GetSummary() HyperflexSummary {
	if o == nil || o.Summary.Get() == nil {
		var ret HyperflexSummary
		return ret
	}
	return *o.Summary.Get()
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexCluster) GetSummaryOk() (*HyperflexSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Summary.Get(), o.Summary.IsSet()
}

// HasSummary returns a boolean if a field has been set.
func (o *HyperflexCluster) HasSummary() bool {
	if o != nil && o.Summary.IsSet() {
		return true
	}

	return false
}

// SetSummary gets a reference to the given NullableHyperflexSummary and assigns it to the Summary field.
func (o *HyperflexCluster) SetSummary(v HyperflexSummary) {
	o.Summary.Set(&v)
}
// SetSummaryNil sets the value for Summary to be an explicit nil
func (o *HyperflexCluster) SetSummaryNil() {
	o.Summary.Set(nil)
}

// UnsetSummary ensures that no value is present for Summary, not even an explicit nil
func (o *HyperflexCluster) UnsetSummary() {
	o.Summary.Unset()
}

// GetUpgradeStatus returns the UpgradeStatus field value if set, zero value otherwise.
func (o *HyperflexCluster) GetUpgradeStatus() string {
	if o == nil || o.UpgradeStatus == nil {
		var ret string
		return ret
	}
	return *o.UpgradeStatus
}

// GetUpgradeStatusOk returns a tuple with the UpgradeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetUpgradeStatusOk() (*string, bool) {
	if o == nil || o.UpgradeStatus == nil {
		return nil, false
	}
	return o.UpgradeStatus, true
}

// HasUpgradeStatus returns a boolean if a field has been set.
func (o *HyperflexCluster) HasUpgradeStatus() bool {
	if o != nil && o.UpgradeStatus != nil {
		return true
	}

	return false
}

// SetUpgradeStatus gets a reference to the given string and assigns it to the UpgradeStatus field.
func (o *HyperflexCluster) SetUpgradeStatus(v string) {
	o.UpgradeStatus = &v
}

// GetVmCount returns the VmCount field value if set, zero value otherwise.
func (o *HyperflexCluster) GetVmCount() int64 {
	if o == nil || o.VmCount == nil {
		var ret int64
		return ret
	}
	return *o.VmCount
}

// GetVmCountOk returns a tuple with the VmCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetVmCountOk() (*int64, bool) {
	if o == nil || o.VmCount == nil {
		return nil, false
	}
	return o.VmCount, true
}

// HasVmCount returns a boolean if a field has been set.
func (o *HyperflexCluster) HasVmCount() bool {
	if o != nil && o.VmCount != nil {
		return true
	}

	return false
}

// SetVmCount gets a reference to the given int64 and assigns it to the VmCount field.
func (o *HyperflexCluster) SetVmCount(v int64) {
	o.VmCount = &v
}

// GetAlarm returns the Alarm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexCluster) GetAlarm() []HyperflexAlarmRelationship {
	if o == nil  {
		var ret []HyperflexAlarmRelationship
		return ret
	}
	return o.Alarm
}

// GetAlarmOk returns a tuple with the Alarm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexCluster) GetAlarmOk() (*[]HyperflexAlarmRelationship, bool) {
	if o == nil || o.Alarm == nil {
		return nil, false
	}
	return &o.Alarm, true
}

// HasAlarm returns a boolean if a field has been set.
func (o *HyperflexCluster) HasAlarm() bool {
	if o != nil && o.Alarm != nil {
		return true
	}

	return false
}

// SetAlarm gets a reference to the given []HyperflexAlarmRelationship and assigns it to the Alarm field.
func (o *HyperflexCluster) SetAlarm(v []HyperflexAlarmRelationship) {
	o.Alarm = v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *HyperflexCluster) GetEncryption() HyperflexEncryptionRelationship {
	if o == nil || o.Encryption == nil {
		var ret HyperflexEncryptionRelationship
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetEncryptionOk() (*HyperflexEncryptionRelationship, bool) {
	if o == nil || o.Encryption == nil {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *HyperflexCluster) HasEncryption() bool {
	if o != nil && o.Encryption != nil {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given HyperflexEncryptionRelationship and assigns it to the Encryption field.
func (o *HyperflexCluster) SetEncryption(v HyperflexEncryptionRelationship) {
	o.Encryption = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *HyperflexCluster) GetHealth() HyperflexHealthRelationship {
	if o == nil || o.Health == nil {
		var ret HyperflexHealthRelationship
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetHealthOk() (*HyperflexHealthRelationship, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *HyperflexCluster) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given HyperflexHealthRelationship and assigns it to the Health field.
func (o *HyperflexCluster) SetHealth(v HyperflexHealthRelationship) {
	o.Health = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *HyperflexCluster) GetLicense() HyperflexLicenseRelationship {
	if o == nil || o.License == nil {
		var ret HyperflexLicenseRelationship
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetLicenseOk() (*HyperflexLicenseRelationship, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *HyperflexCluster) HasLicense() bool {
	if o != nil && o.License != nil {
		return true
	}

	return false
}

// SetLicense gets a reference to the given HyperflexLicenseRelationship and assigns it to the License field.
func (o *HyperflexCluster) SetLicense(v HyperflexLicenseRelationship) {
	o.License = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexCluster) GetNodes() []HyperflexNodeRelationship {
	if o == nil  {
		var ret []HyperflexNodeRelationship
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexCluster) GetNodesOk() (*[]HyperflexNodeRelationship, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *HyperflexCluster) HasNodes() bool {
	if o != nil && o.Nodes != nil {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []HyperflexNodeRelationship and assigns it to the Nodes field.
func (o *HyperflexCluster) SetNodes(v []HyperflexNodeRelationship) {
	o.Nodes = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *HyperflexCluster) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *HyperflexCluster) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *HyperflexCluster) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageClientIpPools returns the StorageClientIpPools field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexCluster) GetStorageClientIpPools() []IppoolPoolRelationship {
	if o == nil  {
		var ret []IppoolPoolRelationship
		return ret
	}
	return o.StorageClientIpPools
}

// GetStorageClientIpPoolsOk returns a tuple with the StorageClientIpPools field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexCluster) GetStorageClientIpPoolsOk() (*[]IppoolPoolRelationship, bool) {
	if o == nil || o.StorageClientIpPools == nil {
		return nil, false
	}
	return &o.StorageClientIpPools, true
}

// HasStorageClientIpPools returns a boolean if a field has been set.
func (o *HyperflexCluster) HasStorageClientIpPools() bool {
	if o != nil && o.StorageClientIpPools != nil {
		return true
	}

	return false
}

// SetStorageClientIpPools gets a reference to the given []IppoolPoolRelationship and assigns it to the StorageClientIpPools field.
func (o *HyperflexCluster) SetStorageClientIpPools(v []IppoolPoolRelationship) {
	o.StorageClientIpPools = v
}

// GetStorageClientVrf returns the StorageClientVrf field value if set, zero value otherwise.
func (o *HyperflexCluster) GetStorageClientVrf() VrfVrfRelationship {
	if o == nil || o.StorageClientVrf == nil {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.StorageClientVrf
}

// GetStorageClientVrfOk returns a tuple with the StorageClientVrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetStorageClientVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil || o.StorageClientVrf == nil {
		return nil, false
	}
	return o.StorageClientVrf, true
}

// HasStorageClientVrf returns a boolean if a field has been set.
func (o *HyperflexCluster) HasStorageClientVrf() bool {
	if o != nil && o.StorageClientVrf != nil {
		return true
	}

	return false
}

// SetStorageClientVrf gets a reference to the given VrfVrfRelationship and assigns it to the StorageClientVrf field.
func (o *HyperflexCluster) SetStorageClientVrf(v VrfVrfRelationship) {
	o.StorageClientVrf = &v
}

// GetStorageContainers returns the StorageContainers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexCluster) GetStorageContainers() []StorageHyperFlexStorageContainerRelationship {
	if o == nil  {
		var ret []StorageHyperFlexStorageContainerRelationship
		return ret
	}
	return o.StorageContainers
}

// GetStorageContainersOk returns a tuple with the StorageContainers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexCluster) GetStorageContainersOk() (*[]StorageHyperFlexStorageContainerRelationship, bool) {
	if o == nil || o.StorageContainers == nil {
		return nil, false
	}
	return &o.StorageContainers, true
}

// HasStorageContainers returns a boolean if a field has been set.
func (o *HyperflexCluster) HasStorageContainers() bool {
	if o != nil && o.StorageContainers != nil {
		return true
	}

	return false
}

// SetStorageContainers gets a reference to the given []StorageHyperFlexStorageContainerRelationship and assigns it to the StorageContainers field.
func (o *HyperflexCluster) SetStorageContainers(v []StorageHyperFlexStorageContainerRelationship) {
	o.StorageContainers = v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexCluster) GetVolumes() []StorageHyperFlexVolumeRelationship {
	if o == nil  {
		var ret []StorageHyperFlexVolumeRelationship
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexCluster) GetVolumesOk() (*[]StorageHyperFlexVolumeRelationship, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *HyperflexCluster) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []StorageHyperFlexVolumeRelationship and assigns it to the Volumes field.
func (o *HyperflexCluster) SetVolumes(v []StorageHyperFlexVolumeRelationship) {
	o.Volumes = v
}

func (o HyperflexCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedHyperflexBaseCluster, errHyperflexBaseCluster := json.Marshal(o.HyperflexBaseCluster)
	if errHyperflexBaseCluster != nil {
		return []byte{}, errHyperflexBaseCluster
	}
	errHyperflexBaseCluster = json.Unmarshal([]byte(serializedHyperflexBaseCluster), &toSerialize)
	if errHyperflexBaseCluster != nil {
		return []byte{}, errHyperflexBaseCluster
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterType != nil {
		toSerialize["ClusterType"] = o.ClusterType
	}
	if o.ClusterUuid != nil {
		toSerialize["ClusterUuid"] = o.ClusterUuid
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.DnsServers != nil {
		toSerialize["DnsServers"] = o.DnsServers
	}
	if o.EncryptionStatus != nil {
		toSerialize["EncryptionStatus"] = o.EncryptionStatus
	}
	if o.FltAggr != nil {
		toSerialize["FltAggr"] = o.FltAggr
	}
	if o.HxdpBuildVersion != nil {
		toSerialize["HxdpBuildVersion"] = o.HxdpBuildVersion
	}
	if o.NtpServers != nil {
		toSerialize["NtpServers"] = o.NtpServers
	}
	if o.Summary.IsSet() {
		toSerialize["Summary"] = o.Summary.Get()
	}
	if o.UpgradeStatus != nil {
		toSerialize["UpgradeStatus"] = o.UpgradeStatus
	}
	if o.VmCount != nil {
		toSerialize["VmCount"] = o.VmCount
	}
	if o.Alarm != nil {
		toSerialize["Alarm"] = o.Alarm
	}
	if o.Encryption != nil {
		toSerialize["Encryption"] = o.Encryption
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.License != nil {
		toSerialize["License"] = o.License
	}
	if o.Nodes != nil {
		toSerialize["Nodes"] = o.Nodes
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageClientIpPools != nil {
		toSerialize["StorageClientIpPools"] = o.StorageClientIpPools
	}
	if o.StorageClientVrf != nil {
		toSerialize["StorageClientVrf"] = o.StorageClientVrf
	}
	if o.StorageContainers != nil {
		toSerialize["StorageContainers"] = o.StorageContainers
	}
	if o.Volumes != nil {
		toSerialize["Volumes"] = o.Volumes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexCluster) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexClusterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The storage type of this cluster (All Flash or Hybrid).
		// Deprecated
		ClusterType *int64 `json:"ClusterType,omitempty"`
		// The unique identifier for this HyperFlex cluster.
		ClusterUuid *string `json:"ClusterUuid,omitempty"`
		// The unique identifier of the device registration that represents this HyperFlex cluster's connection to Intersight.
		DeviceId *string `json:"DeviceId,omitempty"`
		DnsServers []string `json:"DnsServers,omitempty"`
		// This captures the encryption status for a HyperFlex cluster. Currently it will have the status if HXA-CLU-0020 alarm is raised. In the future it can capture other details.
		EncryptionStatus *string `json:"EncryptionStatus,omitempty"`
		// The number of yellow (warning) and red (critical) alarms stored as an aggregate. The first 16 bits indicate the number of red alarms, and the last 16 bits contain the number of yellow alarms.
		// Deprecated
		FltAggr *int64 `json:"FltAggr,omitempty"`
		// The version and build number of the HyperFlex Data Platform for this cluster. After a cluster upgrade, this version string will be updated on the next inventory cycle to reflect the newly installed version.
		HxdpBuildVersion *string `json:"HxdpBuildVersion,omitempty"`
		NtpServers []string `json:"NtpServers,omitempty"`
		Summary NullableHyperflexSummary `json:"Summary,omitempty"`
		// The upgrade status of the HyperFlex cluster. * `Unknown` - The upgrade status of the HyperFlex cluster could not be determined. * `Ok` - The upgrade of the HyperFlex cluster is complete. * `InProgress` - The upgrade of the HyperFlex cluster is in-progress. * `Failed` - The upgrade of the HyperFlex cluster has failed. * `Waiting` - The upgrade of the HyperFlex cluster is waiting to continue execution.
		UpgradeStatus *string `json:"UpgradeStatus,omitempty"`
		// The number of virtual machines present on this cluster.
		VmCount *int64 `json:"VmCount,omitempty"`
		// An array of relationships to hyperflexAlarm resources.
		Alarm []HyperflexAlarmRelationship `json:"Alarm,omitempty"`
		Encryption *HyperflexEncryptionRelationship `json:"Encryption,omitempty"`
		Health *HyperflexHealthRelationship `json:"Health,omitempty"`
		License *HyperflexLicenseRelationship `json:"License,omitempty"`
		// An array of relationships to hyperflexNode resources.
		Nodes []HyperflexNodeRelationship `json:"Nodes,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		// An array of relationships to ippoolPool resources.
		StorageClientIpPools []IppoolPoolRelationship `json:"StorageClientIpPools,omitempty"`
		StorageClientVrf *VrfVrfRelationship `json:"StorageClientVrf,omitempty"`
		// An array of relationships to storageHyperFlexStorageContainer resources.
		StorageContainers []StorageHyperFlexStorageContainerRelationship `json:"StorageContainers,omitempty"`
		// An array of relationships to storageHyperFlexVolume resources.
		Volumes []StorageHyperFlexVolumeRelationship `json:"Volumes,omitempty"`
	}

	varHyperflexClusterWithoutEmbeddedStruct := HyperflexClusterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexClusterWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexCluster := _HyperflexCluster{}
		varHyperflexCluster.ClassId = varHyperflexClusterWithoutEmbeddedStruct.ClassId
		varHyperflexCluster.ObjectType = varHyperflexClusterWithoutEmbeddedStruct.ObjectType
		varHyperflexCluster.ClusterType = varHyperflexClusterWithoutEmbeddedStruct.ClusterType
		varHyperflexCluster.ClusterUuid = varHyperflexClusterWithoutEmbeddedStruct.ClusterUuid
		varHyperflexCluster.DeviceId = varHyperflexClusterWithoutEmbeddedStruct.DeviceId
		varHyperflexCluster.DnsServers = varHyperflexClusterWithoutEmbeddedStruct.DnsServers
		varHyperflexCluster.EncryptionStatus = varHyperflexClusterWithoutEmbeddedStruct.EncryptionStatus
		varHyperflexCluster.FltAggr = varHyperflexClusterWithoutEmbeddedStruct.FltAggr
		varHyperflexCluster.HxdpBuildVersion = varHyperflexClusterWithoutEmbeddedStruct.HxdpBuildVersion
		varHyperflexCluster.NtpServers = varHyperflexClusterWithoutEmbeddedStruct.NtpServers
		varHyperflexCluster.Summary = varHyperflexClusterWithoutEmbeddedStruct.Summary
		varHyperflexCluster.UpgradeStatus = varHyperflexClusterWithoutEmbeddedStruct.UpgradeStatus
		varHyperflexCluster.VmCount = varHyperflexClusterWithoutEmbeddedStruct.VmCount
		varHyperflexCluster.Alarm = varHyperflexClusterWithoutEmbeddedStruct.Alarm
		varHyperflexCluster.Encryption = varHyperflexClusterWithoutEmbeddedStruct.Encryption
		varHyperflexCluster.Health = varHyperflexClusterWithoutEmbeddedStruct.Health
		varHyperflexCluster.License = varHyperflexClusterWithoutEmbeddedStruct.License
		varHyperflexCluster.Nodes = varHyperflexClusterWithoutEmbeddedStruct.Nodes
		varHyperflexCluster.RegisteredDevice = varHyperflexClusterWithoutEmbeddedStruct.RegisteredDevice
		varHyperflexCluster.StorageClientIpPools = varHyperflexClusterWithoutEmbeddedStruct.StorageClientIpPools
		varHyperflexCluster.StorageClientVrf = varHyperflexClusterWithoutEmbeddedStruct.StorageClientVrf
		varHyperflexCluster.StorageContainers = varHyperflexClusterWithoutEmbeddedStruct.StorageContainers
		varHyperflexCluster.Volumes = varHyperflexClusterWithoutEmbeddedStruct.Volumes
		*o = HyperflexCluster(varHyperflexCluster)
	} else {
		return err
	}

	varHyperflexCluster := _HyperflexCluster{}

	err = json.Unmarshal(bytes, &varHyperflexCluster)
	if err == nil {
		o.HyperflexBaseCluster = varHyperflexCluster.HyperflexBaseCluster
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterType")
		delete(additionalProperties, "ClusterUuid")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "DnsServers")
		delete(additionalProperties, "EncryptionStatus")
		delete(additionalProperties, "FltAggr")
		delete(additionalProperties, "HxdpBuildVersion")
		delete(additionalProperties, "NtpServers")
		delete(additionalProperties, "Summary")
		delete(additionalProperties, "UpgradeStatus")
		delete(additionalProperties, "VmCount")
		delete(additionalProperties, "Alarm")
		delete(additionalProperties, "Encryption")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "License")
		delete(additionalProperties, "Nodes")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageClientIpPools")
		delete(additionalProperties, "StorageClientVrf")
		delete(additionalProperties, "StorageContainers")
		delete(additionalProperties, "Volumes")

		// remove fields from embedded structs
		reflectHyperflexBaseCluster := reflect.ValueOf(o.HyperflexBaseCluster)
		for i := 0; i < reflectHyperflexBaseCluster.Type().NumField(); i++ {
			t := reflectHyperflexBaseCluster.Type().Field(i)

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

type NullableHyperflexCluster struct {
	value *HyperflexCluster
	isSet bool
}

func (v NullableHyperflexCluster) Get() *HyperflexCluster {
	return v.value
}

func (v *NullableHyperflexCluster) Set(val *HyperflexCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexCluster(val *HyperflexCluster) *NullableHyperflexCluster {
	return &NullableHyperflexCluster{value: val, isSet: true}
}

func (v NullableHyperflexCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


