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

// CloudRegions The geographic location where a clouds resources are located. It has details such as cloud name, region name, region identifier, list of zones, region endpoint etc.
type CloudRegions struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	AlternateNames []string `json:"AlternateNames,omitempty"`
	// The default zone for this region.
	DefaultZone *string `json:"DefaultZone,omitempty"`
	// Property to identify regions in same category which shares credentials. For e.g. AWS Govcloud Vs AWS Global Vs AWS China.
	Group *string `json:"Group,omitempty"`
	// Flag to indicate of this region is active or not.
	IsActive *bool `json:"IsActive,omitempty"`
	// Property to pick regions for orchestration.
	IsBillingOnly *bool `json:"IsBillingOnly,omitempty"`
	// The display name of the region.
	Name *string `json:"Name,omitempty"`
	// The platform type for this region. For e.g. AmazonWebService. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `IMCRack` - A standalone UCS M6 and above server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `IWE` - An Intersight Workload Engine. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance. * `IntersightAssist` - A Cisco Intersight Assist. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `NexusDevice` - A generic platform type to support Nexus Network Device. This can also be extended to support all network devices later on. * `UCSC890` - A standalone Cisco UCSC890 server. * `NetAppOntap` - A NetApp ONTAP storage system. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `NewRelic` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * `DellCompellent` - A Dell Compellent storage system. * `HPE3Par` - A HPE 3PAR storage system. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * `IMCBlade` - An Intersight managed UCS Blade Server. * `TerraformCloud` - A Terraform Cloud account. * `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `AnsibleEndpoint` - An external endpoint added as Target that can be accessed through Ansible in Intersight Cloud Orchestrator automation workflow. * `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic, Bearer Token. * `SSHEndpoint` - An external endpoint added as Target that can be accessed through SSH in Intersight Cloud Orchestrator automation workflow. * `CiscoCatalyst` - A Cisco Catalyst networking switch device. * `PowerShellEndpoint` - A Windows machine on which PowerShell scripts can be executed remotely.
	PlatformType *string `json:"PlatformType,omitempty"`
	// HTTP endpoint of the region. For example https://ec2.us-east-2.amazonaws.com.
	RegionEndPoint *string `json:"RegionEndPoint,omitempty"`
	// The region Id which is assigned by the cloud provider. For e.g. us-east-1.
	RegionId *string `json:"RegionId,omitempty"`
	Zones []string `json:"Zones,omitempty"`
	Target *AssetTargetRelationship `json:"Target,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudRegions CloudRegions

// NewCloudRegions instantiates a new CloudRegions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRegions(classId string, objectType string) *CloudRegions {
	this := CloudRegions{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isActive bool = true
	this.IsActive = &isActive
	var platformType string = ""
	this.PlatformType = &platformType
	return &this
}

// NewCloudRegionsWithDefaults instantiates a new CloudRegions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRegionsWithDefaults() *CloudRegions {
	this := CloudRegions{}
	var classId string = "cloud.Regions"
	this.ClassId = classId
	var objectType string = "cloud.Regions"
	this.ObjectType = objectType
	var isActive bool = true
	this.IsActive = &isActive
	var platformType string = ""
	this.PlatformType = &platformType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudRegions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudRegions) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudRegions) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudRegions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudRegions) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudRegions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlternateNames returns the AlternateNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudRegions) GetAlternateNames() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.AlternateNames
}

// GetAlternateNamesOk returns a tuple with the AlternateNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudRegions) GetAlternateNamesOk() (*[]string, bool) {
	if o == nil || o.AlternateNames == nil {
		return nil, false
	}
	return &o.AlternateNames, true
}

// HasAlternateNames returns a boolean if a field has been set.
func (o *CloudRegions) HasAlternateNames() bool {
	if o != nil && o.AlternateNames != nil {
		return true
	}

	return false
}

// SetAlternateNames gets a reference to the given []string and assigns it to the AlternateNames field.
func (o *CloudRegions) SetAlternateNames(v []string) {
	o.AlternateNames = v
}

// GetDefaultZone returns the DefaultZone field value if set, zero value otherwise.
func (o *CloudRegions) GetDefaultZone() string {
	if o == nil || o.DefaultZone == nil {
		var ret string
		return ret
	}
	return *o.DefaultZone
}

// GetDefaultZoneOk returns a tuple with the DefaultZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegions) GetDefaultZoneOk() (*string, bool) {
	if o == nil || o.DefaultZone == nil {
		return nil, false
	}
	return o.DefaultZone, true
}

// HasDefaultZone returns a boolean if a field has been set.
func (o *CloudRegions) HasDefaultZone() bool {
	if o != nil && o.DefaultZone != nil {
		return true
	}

	return false
}

// SetDefaultZone gets a reference to the given string and assigns it to the DefaultZone field.
func (o *CloudRegions) SetDefaultZone(v string) {
	o.DefaultZone = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *CloudRegions) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegions) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *CloudRegions) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *CloudRegions) SetGroup(v string) {
	o.Group = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *CloudRegions) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegions) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *CloudRegions) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *CloudRegions) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsBillingOnly returns the IsBillingOnly field value if set, zero value otherwise.
func (o *CloudRegions) GetIsBillingOnly() bool {
	if o == nil || o.IsBillingOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsBillingOnly
}

// GetIsBillingOnlyOk returns a tuple with the IsBillingOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegions) GetIsBillingOnlyOk() (*bool, bool) {
	if o == nil || o.IsBillingOnly == nil {
		return nil, false
	}
	return o.IsBillingOnly, true
}

// HasIsBillingOnly returns a boolean if a field has been set.
func (o *CloudRegions) HasIsBillingOnly() bool {
	if o != nil && o.IsBillingOnly != nil {
		return true
	}

	return false
}

// SetIsBillingOnly gets a reference to the given bool and assigns it to the IsBillingOnly field.
func (o *CloudRegions) SetIsBillingOnly(v bool) {
	o.IsBillingOnly = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudRegions) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegions) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudRegions) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudRegions) SetName(v string) {
	o.Name = &v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *CloudRegions) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegions) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *CloudRegions) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *CloudRegions) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetRegionEndPoint returns the RegionEndPoint field value if set, zero value otherwise.
func (o *CloudRegions) GetRegionEndPoint() string {
	if o == nil || o.RegionEndPoint == nil {
		var ret string
		return ret
	}
	return *o.RegionEndPoint
}

// GetRegionEndPointOk returns a tuple with the RegionEndPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegions) GetRegionEndPointOk() (*string, bool) {
	if o == nil || o.RegionEndPoint == nil {
		return nil, false
	}
	return o.RegionEndPoint, true
}

// HasRegionEndPoint returns a boolean if a field has been set.
func (o *CloudRegions) HasRegionEndPoint() bool {
	if o != nil && o.RegionEndPoint != nil {
		return true
	}

	return false
}

// SetRegionEndPoint gets a reference to the given string and assigns it to the RegionEndPoint field.
func (o *CloudRegions) SetRegionEndPoint(v string) {
	o.RegionEndPoint = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *CloudRegions) GetRegionId() string {
	if o == nil || o.RegionId == nil {
		var ret string
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegions) GetRegionIdOk() (*string, bool) {
	if o == nil || o.RegionId == nil {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *CloudRegions) HasRegionId() bool {
	if o != nil && o.RegionId != nil {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given string and assigns it to the RegionId field.
func (o *CloudRegions) SetRegionId(v string) {
	o.RegionId = &v
}

// GetZones returns the Zones field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudRegions) GetZones() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudRegions) GetZonesOk() (*[]string, bool) {
	if o == nil || o.Zones == nil {
		return nil, false
	}
	return &o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *CloudRegions) HasZones() bool {
	if o != nil && o.Zones != nil {
		return true
	}

	return false
}

// SetZones gets a reference to the given []string and assigns it to the Zones field.
func (o *CloudRegions) SetZones(v []string) {
	o.Zones = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *CloudRegions) GetTarget() AssetTargetRelationship {
	if o == nil || o.Target == nil {
		var ret AssetTargetRelationship
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegions) GetTargetOk() (*AssetTargetRelationship, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *CloudRegions) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AssetTargetRelationship and assigns it to the Target field.
func (o *CloudRegions) SetTarget(v AssetTargetRelationship) {
	o.Target = &v
}

func (o CloudRegions) MarshalJSON() ([]byte, error) {
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
	if o.AlternateNames != nil {
		toSerialize["AlternateNames"] = o.AlternateNames
	}
	if o.DefaultZone != nil {
		toSerialize["DefaultZone"] = o.DefaultZone
	}
	if o.Group != nil {
		toSerialize["Group"] = o.Group
	}
	if o.IsActive != nil {
		toSerialize["IsActive"] = o.IsActive
	}
	if o.IsBillingOnly != nil {
		toSerialize["IsBillingOnly"] = o.IsBillingOnly
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PlatformType != nil {
		toSerialize["PlatformType"] = o.PlatformType
	}
	if o.RegionEndPoint != nil {
		toSerialize["RegionEndPoint"] = o.RegionEndPoint
	}
	if o.RegionId != nil {
		toSerialize["RegionId"] = o.RegionId
	}
	if o.Zones != nil {
		toSerialize["Zones"] = o.Zones
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudRegions) UnmarshalJSON(bytes []byte) (err error) {
	type CloudRegionsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		AlternateNames []string `json:"AlternateNames,omitempty"`
		// The default zone for this region.
		DefaultZone *string `json:"DefaultZone,omitempty"`
		// Property to identify regions in same category which shares credentials. For e.g. AWS Govcloud Vs AWS Global Vs AWS China.
		Group *string `json:"Group,omitempty"`
		// Flag to indicate of this region is active or not.
		IsActive *bool `json:"IsActive,omitempty"`
		// Property to pick regions for orchestration.
		IsBillingOnly *bool `json:"IsBillingOnly,omitempty"`
		// The display name of the region.
		Name *string `json:"Name,omitempty"`
		// The platform type for this region. For e.g. AmazonWebService. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `IMCRack` - A standalone UCS M6 and above server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `IWE` - An Intersight Workload Engine. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance. * `IntersightAssist` - A Cisco Intersight Assist. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `NexusDevice` - A generic platform type to support Nexus Network Device. This can also be extended to support all network devices later on. * `UCSC890` - A standalone Cisco UCSC890 server. * `NetAppOntap` - A NetApp ONTAP storage system. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `NewRelic` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * `DellCompellent` - A Dell Compellent storage system. * `HPE3Par` - A HPE 3PAR storage system. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * `IMCBlade` - An Intersight managed UCS Blade Server. * `TerraformCloud` - A Terraform Cloud account. * `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `AnsibleEndpoint` - An external endpoint added as Target that can be accessed through Ansible in Intersight Cloud Orchestrator automation workflow. * `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic, Bearer Token. * `SSHEndpoint` - An external endpoint added as Target that can be accessed through SSH in Intersight Cloud Orchestrator automation workflow. * `CiscoCatalyst` - A Cisco Catalyst networking switch device. * `PowerShellEndpoint` - A Windows machine on which PowerShell scripts can be executed remotely.
		PlatformType *string `json:"PlatformType,omitempty"`
		// HTTP endpoint of the region. For example https://ec2.us-east-2.amazonaws.com.
		RegionEndPoint *string `json:"RegionEndPoint,omitempty"`
		// The region Id which is assigned by the cloud provider. For e.g. us-east-1.
		RegionId *string `json:"RegionId,omitempty"`
		Zones []string `json:"Zones,omitempty"`
		Target *AssetTargetRelationship `json:"Target,omitempty"`
	}

	varCloudRegionsWithoutEmbeddedStruct := CloudRegionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudRegionsWithoutEmbeddedStruct)
	if err == nil {
		varCloudRegions := _CloudRegions{}
		varCloudRegions.ClassId = varCloudRegionsWithoutEmbeddedStruct.ClassId
		varCloudRegions.ObjectType = varCloudRegionsWithoutEmbeddedStruct.ObjectType
		varCloudRegions.AlternateNames = varCloudRegionsWithoutEmbeddedStruct.AlternateNames
		varCloudRegions.DefaultZone = varCloudRegionsWithoutEmbeddedStruct.DefaultZone
		varCloudRegions.Group = varCloudRegionsWithoutEmbeddedStruct.Group
		varCloudRegions.IsActive = varCloudRegionsWithoutEmbeddedStruct.IsActive
		varCloudRegions.IsBillingOnly = varCloudRegionsWithoutEmbeddedStruct.IsBillingOnly
		varCloudRegions.Name = varCloudRegionsWithoutEmbeddedStruct.Name
		varCloudRegions.PlatformType = varCloudRegionsWithoutEmbeddedStruct.PlatformType
		varCloudRegions.RegionEndPoint = varCloudRegionsWithoutEmbeddedStruct.RegionEndPoint
		varCloudRegions.RegionId = varCloudRegionsWithoutEmbeddedStruct.RegionId
		varCloudRegions.Zones = varCloudRegionsWithoutEmbeddedStruct.Zones
		varCloudRegions.Target = varCloudRegionsWithoutEmbeddedStruct.Target
		*o = CloudRegions(varCloudRegions)
	} else {
		return err
	}

	varCloudRegions := _CloudRegions{}

	err = json.Unmarshal(bytes, &varCloudRegions)
	if err == nil {
		o.MoBaseMo = varCloudRegions.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlternateNames")
		delete(additionalProperties, "DefaultZone")
		delete(additionalProperties, "Group")
		delete(additionalProperties, "IsActive")
		delete(additionalProperties, "IsBillingOnly")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PlatformType")
		delete(additionalProperties, "RegionEndPoint")
		delete(additionalProperties, "RegionId")
		delete(additionalProperties, "Zones")
		delete(additionalProperties, "Target")

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

type NullableCloudRegions struct {
	value *CloudRegions
	isSet bool
}

func (v NullableCloudRegions) Get() *CloudRegions {
	return v.value
}

func (v *NullableCloudRegions) Set(val *CloudRegions) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRegions) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRegions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRegions(val *CloudRegions) *NullableCloudRegions {
	return &NullableCloudRegions{value: val, isSet: true}
}

func (v NullableCloudRegions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRegions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


