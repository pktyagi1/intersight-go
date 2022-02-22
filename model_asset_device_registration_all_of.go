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
	"time"
)

// AssetDeviceRegistrationAllOf Definition of the list of properties defined in 'asset.DeviceRegistration', excluding properties defined in parent classes.
type AssetDeviceRegistrationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An identifier for the credential used by the device connector to authenticate with the Intersight web socket gateway.
	AccessKeyId *string `json:"AccessKeyId,omitempty"`
	// The name of the user who claimed the device for the account.
	ClaimedByUserName *string `json:"ClaimedByUserName,omitempty"`
	// The date and time at which the device was claimed to this account.
	ClaimedTime *time.Time `json:"ClaimedTime,omitempty"`
	DeviceHostname []string `json:"DeviceHostname,omitempty"`
	DeviceIpAddress []string `json:"DeviceIpAddress,omitempty"`
	// Indicates if the platform is an actual device or an emulated device for testing, demos, etc. Permitted values are [Normal, Emulator, ContainerEmulator]. * `` - The device reported an empty or unrecognized executionMode. * `Normal` - The device connector is running in normal mode, i.e. it is not a simulation. * `Emulator` - The device connector is running in simulation mode inside an emulated device. * `ContainerEmulator` - The device connector is running in simulation mode inside a containerized emulated device.
	ExecutionMode *string `json:"ExecutionMode,omitempty"`
	ParentSignature NullableAssetClaimSignature `json:"ParentSignature,omitempty"`
	Pid []string `json:"Pid,omitempty"`
	// The platform type on which device connector is executing. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `IMCRack` - A standalone UCS M6 and above server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `IWE` - An Intersight Workload Engine. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance. * `IntersightAssist` - A Cisco Intersight Assist. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `NexusDevice` - A generic platform type to support Nexus Network Device. This can also be extended to support all network devices later on. * `MDSDevice` - A platform type to support MDS devices. * `UCSC890` - A standalone Cisco UCSC890 server. * `NetAppOntap` - A NetApp ONTAP storage system. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `NewRelic` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * `DellCompellent` - A Dell Compellent storage system. * `HPE3Par` - A HPE 3PAR storage system. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * `IMCBlade` - An Intersight managed UCS Blade Server. * `TerraformCloud` - A Terraform Cloud account. * `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `AnsibleEndpoint` - An external endpoint added as Target that can be accessed through Ansible in Intersight Cloud Orchestrator automation workflow. * `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic, Bearer Token. * `SSHEndpoint` - An external endpoint added as Target that can be accessed through SSH in Intersight Cloud Orchestrator automation workflow. * `CiscoCatalyst` - A Cisco Catalyst networking switch device. * `PowerShellEndpoint` - A Windows machine on which PowerShell scripts can be executed remotely.
	PlatformType *string `json:"PlatformType,omitempty"`
	// The device connector's public key used by Intersight to authenticate a connection from the device connector. The public key is used to verify that the signature a device connector sends on connect has been signed by the connector's private key stored on the device's filesystem. Must be a PEM encoded RSA public key string.
	PublicAccessKey *string `json:"PublicAccessKey,omitempty"`
	// Flag reported by devices to indicate an administrator of the device has disabled management operations of the device connector and only monitoring is permitted.
	ReadOnly *bool `json:"ReadOnly,omitempty"`
	Serial []string `json:"Serial,omitempty"`
	// The vendor of the managed device.
	Vendor *string `json:"Vendor,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	ClaimedByUser *IamUserRelationship `json:"ClaimedByUser,omitempty"`
	// An array of relationships to assetClusterMember resources.
	ClusterMembers []AssetClusterMemberRelationship `json:"ClusterMembers,omitempty"`
	DeviceClaim *AssetDeviceClaimRelationship `json:"DeviceClaim,omitempty"`
	DeviceConfiguration *AssetDeviceConfigurationRelationship `json:"DeviceConfiguration,omitempty"`
	DomainGroup *IamDomainGroupRelationship `json:"DomainGroup,omitempty"`
	ParentConnection *AssetDeviceRegistrationRelationship `json:"ParentConnection,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetDeviceRegistrationAllOf AssetDeviceRegistrationAllOf

// NewAssetDeviceRegistrationAllOf instantiates a new AssetDeviceRegistrationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeviceRegistrationAllOf(classId string, objectType string) *AssetDeviceRegistrationAllOf {
	this := AssetDeviceRegistrationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var executionMode string = ""
	this.ExecutionMode = &executionMode
	var platformType string = ""
	this.PlatformType = &platformType
	return &this
}

// NewAssetDeviceRegistrationAllOfWithDefaults instantiates a new AssetDeviceRegistrationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeviceRegistrationAllOfWithDefaults() *AssetDeviceRegistrationAllOf {
	this := AssetDeviceRegistrationAllOf{}
	var classId string = "asset.DeviceRegistration"
	this.ClassId = classId
	var objectType string = "asset.DeviceRegistration"
	this.ObjectType = objectType
	var executionMode string = ""
	this.ExecutionMode = &executionMode
	var platformType string = ""
	this.PlatformType = &platformType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetDeviceRegistrationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetDeviceRegistrationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetDeviceRegistrationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetDeviceRegistrationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *AssetDeviceRegistrationAllOf) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasAccessKeyId() bool {
	if o != nil && o.AccessKeyId != nil {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *AssetDeviceRegistrationAllOf) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetClaimedByUserName returns the ClaimedByUserName field value if set, zero value otherwise.
func (o *AssetDeviceRegistrationAllOf) GetClaimedByUserName() string {
	if o == nil || o.ClaimedByUserName == nil {
		var ret string
		return ret
	}
	return *o.ClaimedByUserName
}

// GetClaimedByUserNameOk returns a tuple with the ClaimedByUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetClaimedByUserNameOk() (*string, bool) {
	if o == nil || o.ClaimedByUserName == nil {
		return nil, false
	}
	return o.ClaimedByUserName, true
}

// HasClaimedByUserName returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasClaimedByUserName() bool {
	if o != nil && o.ClaimedByUserName != nil {
		return true
	}

	return false
}

// SetClaimedByUserName gets a reference to the given string and assigns it to the ClaimedByUserName field.
func (o *AssetDeviceRegistrationAllOf) SetClaimedByUserName(v string) {
	o.ClaimedByUserName = &v
}

// GetClaimedTime returns the ClaimedTime field value if set, zero value otherwise.
func (o *AssetDeviceRegistrationAllOf) GetClaimedTime() time.Time {
	if o == nil || o.ClaimedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ClaimedTime
}

// GetClaimedTimeOk returns a tuple with the ClaimedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetClaimedTimeOk() (*time.Time, bool) {
	if o == nil || o.ClaimedTime == nil {
		return nil, false
	}
	return o.ClaimedTime, true
}

// HasClaimedTime returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasClaimedTime() bool {
	if o != nil && o.ClaimedTime != nil {
		return true
	}

	return false
}

// SetClaimedTime gets a reference to the given time.Time and assigns it to the ClaimedTime field.
func (o *AssetDeviceRegistrationAllOf) SetClaimedTime(v time.Time) {
	o.ClaimedTime = &v
}

// GetDeviceHostname returns the DeviceHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceRegistrationAllOf) GetDeviceHostname() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.DeviceHostname
}

// GetDeviceHostnameOk returns a tuple with the DeviceHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceRegistrationAllOf) GetDeviceHostnameOk() (*[]string, bool) {
	if o == nil || o.DeviceHostname == nil {
		return nil, false
	}
	return &o.DeviceHostname, true
}

// HasDeviceHostname returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasDeviceHostname() bool {
	if o != nil && o.DeviceHostname != nil {
		return true
	}

	return false
}

// SetDeviceHostname gets a reference to the given []string and assigns it to the DeviceHostname field.
func (o *AssetDeviceRegistrationAllOf) SetDeviceHostname(v []string) {
	o.DeviceHostname = v
}

// GetDeviceIpAddress returns the DeviceIpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceRegistrationAllOf) GetDeviceIpAddress() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.DeviceIpAddress
}

// GetDeviceIpAddressOk returns a tuple with the DeviceIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceRegistrationAllOf) GetDeviceIpAddressOk() (*[]string, bool) {
	if o == nil || o.DeviceIpAddress == nil {
		return nil, false
	}
	return &o.DeviceIpAddress, true
}

// HasDeviceIpAddress returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasDeviceIpAddress() bool {
	if o != nil && o.DeviceIpAddress != nil {
		return true
	}

	return false
}

// SetDeviceIpAddress gets a reference to the given []string and assigns it to the DeviceIpAddress field.
func (o *AssetDeviceRegistrationAllOf) SetDeviceIpAddress(v []string) {
	o.DeviceIpAddress = v
}

// GetExecutionMode returns the ExecutionMode field value if set, zero value otherwise.
func (o *AssetDeviceRegistrationAllOf) GetExecutionMode() string {
	if o == nil || o.ExecutionMode == nil {
		var ret string
		return ret
	}
	return *o.ExecutionMode
}

// GetExecutionModeOk returns a tuple with the ExecutionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetExecutionModeOk() (*string, bool) {
	if o == nil || o.ExecutionMode == nil {
		return nil, false
	}
	return o.ExecutionMode, true
}

// HasExecutionMode returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasExecutionMode() bool {
	if o != nil && o.ExecutionMode != nil {
		return true
	}

	return false
}

// SetExecutionMode gets a reference to the given string and assigns it to the ExecutionMode field.
func (o *AssetDeviceRegistrationAllOf) SetExecutionMode(v string) {
	o.ExecutionMode = &v
}

// GetParentSignature returns the ParentSignature field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceRegistrationAllOf) GetParentSignature() AssetClaimSignature {
	if o == nil || o.ParentSignature.Get() == nil {
		var ret AssetClaimSignature
		return ret
	}
	return *o.ParentSignature.Get()
}

// GetParentSignatureOk returns a tuple with the ParentSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceRegistrationAllOf) GetParentSignatureOk() (*AssetClaimSignature, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentSignature.Get(), o.ParentSignature.IsSet()
}

// HasParentSignature returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasParentSignature() bool {
	if o != nil && o.ParentSignature.IsSet() {
		return true
	}

	return false
}

// SetParentSignature gets a reference to the given NullableAssetClaimSignature and assigns it to the ParentSignature field.
func (o *AssetDeviceRegistrationAllOf) SetParentSignature(v AssetClaimSignature) {
	o.ParentSignature.Set(&v)
}
// SetParentSignatureNil sets the value for ParentSignature to be an explicit nil
func (o *AssetDeviceRegistrationAllOf) SetParentSignatureNil() {
	o.ParentSignature.Set(nil)
}

// UnsetParentSignature ensures that no value is present for ParentSignature, not even an explicit nil
func (o *AssetDeviceRegistrationAllOf) UnsetParentSignature() {
	o.ParentSignature.Unset()
}

// GetPid returns the Pid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceRegistrationAllOf) GetPid() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceRegistrationAllOf) GetPidOk() (*[]string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return &o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given []string and assigns it to the Pid field.
func (o *AssetDeviceRegistrationAllOf) SetPid(v []string) {
	o.Pid = v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *AssetDeviceRegistrationAllOf) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *AssetDeviceRegistrationAllOf) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetPublicAccessKey returns the PublicAccessKey field value if set, zero value otherwise.
func (o *AssetDeviceRegistrationAllOf) GetPublicAccessKey() string {
	if o == nil || o.PublicAccessKey == nil {
		var ret string
		return ret
	}
	return *o.PublicAccessKey
}

// GetPublicAccessKeyOk returns a tuple with the PublicAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetPublicAccessKeyOk() (*string, bool) {
	if o == nil || o.PublicAccessKey == nil {
		return nil, false
	}
	return o.PublicAccessKey, true
}

// HasPublicAccessKey returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasPublicAccessKey() bool {
	if o != nil && o.PublicAccessKey != nil {
		return true
	}

	return false
}

// SetPublicAccessKey gets a reference to the given string and assigns it to the PublicAccessKey field.
func (o *AssetDeviceRegistrationAllOf) SetPublicAccessKey(v string) {
	o.PublicAccessKey = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *AssetDeviceRegistrationAllOf) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *AssetDeviceRegistrationAllOf) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceRegistrationAllOf) GetSerial() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceRegistrationAllOf) GetSerialOk() (*[]string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return &o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given []string and assigns it to the Serial field.
func (o *AssetDeviceRegistrationAllOf) SetSerial(v []string) {
	o.Serial = v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *AssetDeviceRegistrationAllOf) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *AssetDeviceRegistrationAllOf) SetVendor(v string) {
	o.Vendor = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AssetDeviceRegistrationAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *AssetDeviceRegistrationAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetClaimedByUser returns the ClaimedByUser field value if set, zero value otherwise.
func (o *AssetDeviceRegistrationAllOf) GetClaimedByUser() IamUserRelationship {
	if o == nil || o.ClaimedByUser == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.ClaimedByUser
}

// GetClaimedByUserOk returns a tuple with the ClaimedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetClaimedByUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.ClaimedByUser == nil {
		return nil, false
	}
	return o.ClaimedByUser, true
}

// HasClaimedByUser returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasClaimedByUser() bool {
	if o != nil && o.ClaimedByUser != nil {
		return true
	}

	return false
}

// SetClaimedByUser gets a reference to the given IamUserRelationship and assigns it to the ClaimedByUser field.
func (o *AssetDeviceRegistrationAllOf) SetClaimedByUser(v IamUserRelationship) {
	o.ClaimedByUser = &v
}

// GetClusterMembers returns the ClusterMembers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceRegistrationAllOf) GetClusterMembers() []AssetClusterMemberRelationship {
	if o == nil  {
		var ret []AssetClusterMemberRelationship
		return ret
	}
	return o.ClusterMembers
}

// GetClusterMembersOk returns a tuple with the ClusterMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceRegistrationAllOf) GetClusterMembersOk() (*[]AssetClusterMemberRelationship, bool) {
	if o == nil || o.ClusterMembers == nil {
		return nil, false
	}
	return &o.ClusterMembers, true
}

// HasClusterMembers returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasClusterMembers() bool {
	if o != nil && o.ClusterMembers != nil {
		return true
	}

	return false
}

// SetClusterMembers gets a reference to the given []AssetClusterMemberRelationship and assigns it to the ClusterMembers field.
func (o *AssetDeviceRegistrationAllOf) SetClusterMembers(v []AssetClusterMemberRelationship) {
	o.ClusterMembers = v
}

// GetDeviceClaim returns the DeviceClaim field value if set, zero value otherwise.
func (o *AssetDeviceRegistrationAllOf) GetDeviceClaim() AssetDeviceClaimRelationship {
	if o == nil || o.DeviceClaim == nil {
		var ret AssetDeviceClaimRelationship
		return ret
	}
	return *o.DeviceClaim
}

// GetDeviceClaimOk returns a tuple with the DeviceClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetDeviceClaimOk() (*AssetDeviceClaimRelationship, bool) {
	if o == nil || o.DeviceClaim == nil {
		return nil, false
	}
	return o.DeviceClaim, true
}

// HasDeviceClaim returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasDeviceClaim() bool {
	if o != nil && o.DeviceClaim != nil {
		return true
	}

	return false
}

// SetDeviceClaim gets a reference to the given AssetDeviceClaimRelationship and assigns it to the DeviceClaim field.
func (o *AssetDeviceRegistrationAllOf) SetDeviceClaim(v AssetDeviceClaimRelationship) {
	o.DeviceClaim = &v
}

// GetDeviceConfiguration returns the DeviceConfiguration field value if set, zero value otherwise.
func (o *AssetDeviceRegistrationAllOf) GetDeviceConfiguration() AssetDeviceConfigurationRelationship {
	if o == nil || o.DeviceConfiguration == nil {
		var ret AssetDeviceConfigurationRelationship
		return ret
	}
	return *o.DeviceConfiguration
}

// GetDeviceConfigurationOk returns a tuple with the DeviceConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetDeviceConfigurationOk() (*AssetDeviceConfigurationRelationship, bool) {
	if o == nil || o.DeviceConfiguration == nil {
		return nil, false
	}
	return o.DeviceConfiguration, true
}

// HasDeviceConfiguration returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasDeviceConfiguration() bool {
	if o != nil && o.DeviceConfiguration != nil {
		return true
	}

	return false
}

// SetDeviceConfiguration gets a reference to the given AssetDeviceConfigurationRelationship and assigns it to the DeviceConfiguration field.
func (o *AssetDeviceRegistrationAllOf) SetDeviceConfiguration(v AssetDeviceConfigurationRelationship) {
	o.DeviceConfiguration = &v
}

// GetDomainGroup returns the DomainGroup field value if set, zero value otherwise.
func (o *AssetDeviceRegistrationAllOf) GetDomainGroup() IamDomainGroupRelationship {
	if o == nil || o.DomainGroup == nil {
		var ret IamDomainGroupRelationship
		return ret
	}
	return *o.DomainGroup
}

// GetDomainGroupOk returns a tuple with the DomainGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetDomainGroupOk() (*IamDomainGroupRelationship, bool) {
	if o == nil || o.DomainGroup == nil {
		return nil, false
	}
	return o.DomainGroup, true
}

// HasDomainGroup returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasDomainGroup() bool {
	if o != nil && o.DomainGroup != nil {
		return true
	}

	return false
}

// SetDomainGroup gets a reference to the given IamDomainGroupRelationship and assigns it to the DomainGroup field.
func (o *AssetDeviceRegistrationAllOf) SetDomainGroup(v IamDomainGroupRelationship) {
	o.DomainGroup = &v
}

// GetParentConnection returns the ParentConnection field value if set, zero value otherwise.
func (o *AssetDeviceRegistrationAllOf) GetParentConnection() AssetDeviceRegistrationRelationship {
	if o == nil || o.ParentConnection == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.ParentConnection
}

// GetParentConnectionOk returns a tuple with the ParentConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceRegistrationAllOf) GetParentConnectionOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.ParentConnection == nil {
		return nil, false
	}
	return o.ParentConnection, true
}

// HasParentConnection returns a boolean if a field has been set.
func (o *AssetDeviceRegistrationAllOf) HasParentConnection() bool {
	if o != nil && o.ParentConnection != nil {
		return true
	}

	return false
}

// SetParentConnection gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the ParentConnection field.
func (o *AssetDeviceRegistrationAllOf) SetParentConnection(v AssetDeviceRegistrationRelationship) {
	o.ParentConnection = &v
}

func (o AssetDeviceRegistrationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccessKeyId != nil {
		toSerialize["AccessKeyId"] = o.AccessKeyId
	}
	if o.ClaimedByUserName != nil {
		toSerialize["ClaimedByUserName"] = o.ClaimedByUserName
	}
	if o.ClaimedTime != nil {
		toSerialize["ClaimedTime"] = o.ClaimedTime
	}
	if o.DeviceHostname != nil {
		toSerialize["DeviceHostname"] = o.DeviceHostname
	}
	if o.DeviceIpAddress != nil {
		toSerialize["DeviceIpAddress"] = o.DeviceIpAddress
	}
	if o.ExecutionMode != nil {
		toSerialize["ExecutionMode"] = o.ExecutionMode
	}
	if o.ParentSignature.IsSet() {
		toSerialize["ParentSignature"] = o.ParentSignature.Get()
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.PlatformType != nil {
		toSerialize["PlatformType"] = o.PlatformType
	}
	if o.PublicAccessKey != nil {
		toSerialize["PublicAccessKey"] = o.PublicAccessKey
	}
	if o.ReadOnly != nil {
		toSerialize["ReadOnly"] = o.ReadOnly
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.ClaimedByUser != nil {
		toSerialize["ClaimedByUser"] = o.ClaimedByUser
	}
	if o.ClusterMembers != nil {
		toSerialize["ClusterMembers"] = o.ClusterMembers
	}
	if o.DeviceClaim != nil {
		toSerialize["DeviceClaim"] = o.DeviceClaim
	}
	if o.DeviceConfiguration != nil {
		toSerialize["DeviceConfiguration"] = o.DeviceConfiguration
	}
	if o.DomainGroup != nil {
		toSerialize["DomainGroup"] = o.DomainGroup
	}
	if o.ParentConnection != nil {
		toSerialize["ParentConnection"] = o.ParentConnection
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetDeviceRegistrationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetDeviceRegistrationAllOf := _AssetDeviceRegistrationAllOf{}

	if err = json.Unmarshal(bytes, &varAssetDeviceRegistrationAllOf); err == nil {
		*o = AssetDeviceRegistrationAllOf(varAssetDeviceRegistrationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessKeyId")
		delete(additionalProperties, "ClaimedByUserName")
		delete(additionalProperties, "ClaimedTime")
		delete(additionalProperties, "DeviceHostname")
		delete(additionalProperties, "DeviceIpAddress")
		delete(additionalProperties, "ExecutionMode")
		delete(additionalProperties, "ParentSignature")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "PlatformType")
		delete(additionalProperties, "PublicAccessKey")
		delete(additionalProperties, "ReadOnly")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "Vendor")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "ClaimedByUser")
		delete(additionalProperties, "ClusterMembers")
		delete(additionalProperties, "DeviceClaim")
		delete(additionalProperties, "DeviceConfiguration")
		delete(additionalProperties, "DomainGroup")
		delete(additionalProperties, "ParentConnection")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetDeviceRegistrationAllOf struct {
	value *AssetDeviceRegistrationAllOf
	isSet bool
}

func (v NullableAssetDeviceRegistrationAllOf) Get() *AssetDeviceRegistrationAllOf {
	return v.value
}

func (v *NullableAssetDeviceRegistrationAllOf) Set(val *AssetDeviceRegistrationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeviceRegistrationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeviceRegistrationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeviceRegistrationAllOf(val *AssetDeviceRegistrationAllOf) *NullableAssetDeviceRegistrationAllOf {
	return &NullableAssetDeviceRegistrationAllOf{value: val, isSet: true}
}

func (v NullableAssetDeviceRegistrationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeviceRegistrationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


