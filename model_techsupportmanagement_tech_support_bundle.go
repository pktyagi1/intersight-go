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

// TechsupportmanagementTechSupportBundle A request to collect techsupport and upload it to Intersight Storage Service. The serial number, PID and/or relationship to the target resource provided by the user is used to determine the device type for techsupport collection. If the serial number, PID and target resource are specified in the request, the values must match. Valid values of device types are network.Element for fabric interconnect, compute.Blade for blade server, compute.RackUnit for rack server, equipment.Chassis for chassis, equipment.IoCard for IO Module, equipment.FEX for fabric extender and adapter.Unit for network adapter. UCSM techsupport is collected for device type network.Element. Chassis techsupport is collected for compute.Blade, equipment.Chassis, equipment.IoCard, and blade adapter.Unit. Server techsupport is collected for compute.RackUnit and rack adapter.Unit. Fabric extender techsupport is collected for device type equipment.FEX. Hyper Flex node level techsupport is collected when the request specifies the platform type (HX) and the device type is Hyperflex.Node.
type TechsupportmanagementTechSupportBundle struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The device identifier used to uniquely identify an individual device.
	DeviceIdentifier *string `json:"DeviceIdentifier,omitempty"`
	// The device type obtained from the inventory.
	DeviceType *string `json:"DeviceType,omitempty"`
	// Product identification of the device.
	Pid *string `json:"Pid,omitempty"`
	PlatformParam NullableConnectorPlatformParamBase `json:"PlatformParam,omitempty"`
	// The platform type of the device. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `IMCRack` - A standalone UCS M6 and above server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `IWE` - An Intersight Workload Engine. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance. * `IntersightAssist` - A Cisco Intersight Assist. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `NexusDevice` - A generic platform type to support Nexus Network Device. This can also be extended to support all network devices later on. * `MDSDevice` - A platform type to support MDS devices. * `UCSC890` - A standalone Cisco UCSC890 server. * `NetAppOntap` - A NetApp ONTAP storage system. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `NewRelic` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * `DellCompellent` - A Dell Compellent storage system. * `HPE3Par` - A HPE 3PAR storage system. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * `IMCBlade` - An Intersight managed UCS Blade Server. * `TerraformCloud` - A Terraform Cloud account. * `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `AnsibleEndpoint` - An external endpoint added as Target that can be accessed through Ansible in Intersight Cloud Orchestrator automation workflow. * `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic, Bearer Token. * `SSHEndpoint` - An external endpoint added as Target that can be accessed through SSH in Intersight Cloud Orchestrator automation workflow. * `CiscoCatalyst` - A Cisco Catalyst networking switch device. * `PowerShellEndpoint` - A Windows machine on which PowerShell scripts can be executed remotely.
	PlatformType *string `json:"PlatformType,omitempty"`
	// Serial number of the device.
	Serial *string `json:"Serial,omitempty"`
	DeviceRegistration *AssetDeviceRegistrationRelationship `json:"DeviceRegistration,omitempty"`
	TargetResource *MoBaseMoRelationship `json:"TargetResource,omitempty"`
	TechSupportStatus *TechsupportmanagementTechSupportStatusRelationship `json:"TechSupportStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TechsupportmanagementTechSupportBundle TechsupportmanagementTechSupportBundle

// NewTechsupportmanagementTechSupportBundle instantiates a new TechsupportmanagementTechSupportBundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechsupportmanagementTechSupportBundle(classId string, objectType string) *TechsupportmanagementTechSupportBundle {
	this := TechsupportmanagementTechSupportBundle{}
	this.ClassId = classId
	this.ObjectType = objectType
	var platformType string = ""
	this.PlatformType = &platformType
	return &this
}

// NewTechsupportmanagementTechSupportBundleWithDefaults instantiates a new TechsupportmanagementTechSupportBundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechsupportmanagementTechSupportBundleWithDefaults() *TechsupportmanagementTechSupportBundle {
	this := TechsupportmanagementTechSupportBundle{}
	var classId string = "techsupportmanagement.TechSupportBundle"
	this.ClassId = classId
	var objectType string = "techsupportmanagement.TechSupportBundle"
	this.ObjectType = objectType
	var platformType string = ""
	this.PlatformType = &platformType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TechsupportmanagementTechSupportBundle) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundle) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TechsupportmanagementTechSupportBundle) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TechsupportmanagementTechSupportBundle) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundle) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TechsupportmanagementTechSupportBundle) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceIdentifier returns the DeviceIdentifier field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundle) GetDeviceIdentifier() string {
	if o == nil || o.DeviceIdentifier == nil {
		var ret string
		return ret
	}
	return *o.DeviceIdentifier
}

// GetDeviceIdentifierOk returns a tuple with the DeviceIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundle) GetDeviceIdentifierOk() (*string, bool) {
	if o == nil || o.DeviceIdentifier == nil {
		return nil, false
	}
	return o.DeviceIdentifier, true
}

// HasDeviceIdentifier returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundle) HasDeviceIdentifier() bool {
	if o != nil && o.DeviceIdentifier != nil {
		return true
	}

	return false
}

// SetDeviceIdentifier gets a reference to the given string and assigns it to the DeviceIdentifier field.
func (o *TechsupportmanagementTechSupportBundle) SetDeviceIdentifier(v string) {
	o.DeviceIdentifier = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundle) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundle) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundle) HasDeviceType() bool {
	if o != nil && o.DeviceType != nil {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *TechsupportmanagementTechSupportBundle) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundle) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundle) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundle) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *TechsupportmanagementTechSupportBundle) SetPid(v string) {
	o.Pid = &v
}

// GetPlatformParam returns the PlatformParam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TechsupportmanagementTechSupportBundle) GetPlatformParam() ConnectorPlatformParamBase {
	if o == nil || o.PlatformParam.Get() == nil {
		var ret ConnectorPlatformParamBase
		return ret
	}
	return *o.PlatformParam.Get()
}

// GetPlatformParamOk returns a tuple with the PlatformParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TechsupportmanagementTechSupportBundle) GetPlatformParamOk() (*ConnectorPlatformParamBase, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PlatformParam.Get(), o.PlatformParam.IsSet()
}

// HasPlatformParam returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundle) HasPlatformParam() bool {
	if o != nil && o.PlatformParam.IsSet() {
		return true
	}

	return false
}

// SetPlatformParam gets a reference to the given NullableConnectorPlatformParamBase and assigns it to the PlatformParam field.
func (o *TechsupportmanagementTechSupportBundle) SetPlatformParam(v ConnectorPlatformParamBase) {
	o.PlatformParam.Set(&v)
}
// SetPlatformParamNil sets the value for PlatformParam to be an explicit nil
func (o *TechsupportmanagementTechSupportBundle) SetPlatformParamNil() {
	o.PlatformParam.Set(nil)
}

// UnsetPlatformParam ensures that no value is present for PlatformParam, not even an explicit nil
func (o *TechsupportmanagementTechSupportBundle) UnsetPlatformParam() {
	o.PlatformParam.Unset()
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundle) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundle) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundle) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *TechsupportmanagementTechSupportBundle) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundle) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundle) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundle) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *TechsupportmanagementTechSupportBundle) SetSerial(v string) {
	o.Serial = &v
}

// GetDeviceRegistration returns the DeviceRegistration field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundle) GetDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || o.DeviceRegistration == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceRegistration
}

// GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundle) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.DeviceRegistration == nil {
		return nil, false
	}
	return o.DeviceRegistration, true
}

// HasDeviceRegistration returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundle) HasDeviceRegistration() bool {
	if o != nil && o.DeviceRegistration != nil {
		return true
	}

	return false
}

// SetDeviceRegistration gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the DeviceRegistration field.
func (o *TechsupportmanagementTechSupportBundle) SetDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.DeviceRegistration = &v
}

// GetTargetResource returns the TargetResource field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundle) GetTargetResource() MoBaseMoRelationship {
	if o == nil || o.TargetResource == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetResource
}

// GetTargetResourceOk returns a tuple with the TargetResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundle) GetTargetResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetResource == nil {
		return nil, false
	}
	return o.TargetResource, true
}

// HasTargetResource returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundle) HasTargetResource() bool {
	if o != nil && o.TargetResource != nil {
		return true
	}

	return false
}

// SetTargetResource gets a reference to the given MoBaseMoRelationship and assigns it to the TargetResource field.
func (o *TechsupportmanagementTechSupportBundle) SetTargetResource(v MoBaseMoRelationship) {
	o.TargetResource = &v
}

// GetTechSupportStatus returns the TechSupportStatus field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundle) GetTechSupportStatus() TechsupportmanagementTechSupportStatusRelationship {
	if o == nil || o.TechSupportStatus == nil {
		var ret TechsupportmanagementTechSupportStatusRelationship
		return ret
	}
	return *o.TechSupportStatus
}

// GetTechSupportStatusOk returns a tuple with the TechSupportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundle) GetTechSupportStatusOk() (*TechsupportmanagementTechSupportStatusRelationship, bool) {
	if o == nil || o.TechSupportStatus == nil {
		return nil, false
	}
	return o.TechSupportStatus, true
}

// HasTechSupportStatus returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundle) HasTechSupportStatus() bool {
	if o != nil && o.TechSupportStatus != nil {
		return true
	}

	return false
}

// SetTechSupportStatus gets a reference to the given TechsupportmanagementTechSupportStatusRelationship and assigns it to the TechSupportStatus field.
func (o *TechsupportmanagementTechSupportBundle) SetTechSupportStatus(v TechsupportmanagementTechSupportStatusRelationship) {
	o.TechSupportStatus = &v
}

func (o TechsupportmanagementTechSupportBundle) MarshalJSON() ([]byte, error) {
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
	if o.DeviceIdentifier != nil {
		toSerialize["DeviceIdentifier"] = o.DeviceIdentifier
	}
	if o.DeviceType != nil {
		toSerialize["DeviceType"] = o.DeviceType
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.PlatformParam.IsSet() {
		toSerialize["PlatformParam"] = o.PlatformParam.Get()
	}
	if o.PlatformType != nil {
		toSerialize["PlatformType"] = o.PlatformType
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.DeviceRegistration != nil {
		toSerialize["DeviceRegistration"] = o.DeviceRegistration
	}
	if o.TargetResource != nil {
		toSerialize["TargetResource"] = o.TargetResource
	}
	if o.TechSupportStatus != nil {
		toSerialize["TechSupportStatus"] = o.TechSupportStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TechsupportmanagementTechSupportBundle) UnmarshalJSON(bytes []byte) (err error) {
	type TechsupportmanagementTechSupportBundleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The device identifier used to uniquely identify an individual device.
		DeviceIdentifier *string `json:"DeviceIdentifier,omitempty"`
		// The device type obtained from the inventory.
		DeviceType *string `json:"DeviceType,omitempty"`
		// Product identification of the device.
		Pid *string `json:"Pid,omitempty"`
		PlatformParam NullableConnectorPlatformParamBase `json:"PlatformParam,omitempty"`
		// The platform type of the device. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `IMCRack` - A standalone UCS M6 and above server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `IWE` - An Intersight Workload Engine. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance. * `IntersightAssist` - A Cisco Intersight Assist. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `NexusDevice` - A generic platform type to support Nexus Network Device. This can also be extended to support all network devices later on. * `MDSDevice` - A platform type to support MDS devices. * `UCSC890` - A standalone Cisco UCSC890 server. * `NetAppOntap` - A NetApp ONTAP storage system. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `NewRelic` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * `DellCompellent` - A Dell Compellent storage system. * `HPE3Par` - A HPE 3PAR storage system. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * `IMCBlade` - An Intersight managed UCS Blade Server. * `TerraformCloud` - A Terraform Cloud account. * `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `AnsibleEndpoint` - An external endpoint added as Target that can be accessed through Ansible in Intersight Cloud Orchestrator automation workflow. * `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic, Bearer Token. * `SSHEndpoint` - An external endpoint added as Target that can be accessed through SSH in Intersight Cloud Orchestrator automation workflow. * `CiscoCatalyst` - A Cisco Catalyst networking switch device. * `PowerShellEndpoint` - A Windows machine on which PowerShell scripts can be executed remotely.
		PlatformType *string `json:"PlatformType,omitempty"`
		// Serial number of the device.
		Serial *string `json:"Serial,omitempty"`
		DeviceRegistration *AssetDeviceRegistrationRelationship `json:"DeviceRegistration,omitempty"`
		TargetResource *MoBaseMoRelationship `json:"TargetResource,omitempty"`
		TechSupportStatus *TechsupportmanagementTechSupportStatusRelationship `json:"TechSupportStatus,omitempty"`
	}

	varTechsupportmanagementTechSupportBundleWithoutEmbeddedStruct := TechsupportmanagementTechSupportBundleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTechsupportmanagementTechSupportBundleWithoutEmbeddedStruct)
	if err == nil {
		varTechsupportmanagementTechSupportBundle := _TechsupportmanagementTechSupportBundle{}
		varTechsupportmanagementTechSupportBundle.ClassId = varTechsupportmanagementTechSupportBundleWithoutEmbeddedStruct.ClassId
		varTechsupportmanagementTechSupportBundle.ObjectType = varTechsupportmanagementTechSupportBundleWithoutEmbeddedStruct.ObjectType
		varTechsupportmanagementTechSupportBundle.DeviceIdentifier = varTechsupportmanagementTechSupportBundleWithoutEmbeddedStruct.DeviceIdentifier
		varTechsupportmanagementTechSupportBundle.DeviceType = varTechsupportmanagementTechSupportBundleWithoutEmbeddedStruct.DeviceType
		varTechsupportmanagementTechSupportBundle.Pid = varTechsupportmanagementTechSupportBundleWithoutEmbeddedStruct.Pid
		varTechsupportmanagementTechSupportBundle.PlatformParam = varTechsupportmanagementTechSupportBundleWithoutEmbeddedStruct.PlatformParam
		varTechsupportmanagementTechSupportBundle.PlatformType = varTechsupportmanagementTechSupportBundleWithoutEmbeddedStruct.PlatformType
		varTechsupportmanagementTechSupportBundle.Serial = varTechsupportmanagementTechSupportBundleWithoutEmbeddedStruct.Serial
		varTechsupportmanagementTechSupportBundle.DeviceRegistration = varTechsupportmanagementTechSupportBundleWithoutEmbeddedStruct.DeviceRegistration
		varTechsupportmanagementTechSupportBundle.TargetResource = varTechsupportmanagementTechSupportBundleWithoutEmbeddedStruct.TargetResource
		varTechsupportmanagementTechSupportBundle.TechSupportStatus = varTechsupportmanagementTechSupportBundleWithoutEmbeddedStruct.TechSupportStatus
		*o = TechsupportmanagementTechSupportBundle(varTechsupportmanagementTechSupportBundle)
	} else {
		return err
	}

	varTechsupportmanagementTechSupportBundle := _TechsupportmanagementTechSupportBundle{}

	err = json.Unmarshal(bytes, &varTechsupportmanagementTechSupportBundle)
	if err == nil {
		o.MoBaseMo = varTechsupportmanagementTechSupportBundle.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceIdentifier")
		delete(additionalProperties, "DeviceType")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "PlatformParam")
		delete(additionalProperties, "PlatformType")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "DeviceRegistration")
		delete(additionalProperties, "TargetResource")
		delete(additionalProperties, "TechSupportStatus")

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

type NullableTechsupportmanagementTechSupportBundle struct {
	value *TechsupportmanagementTechSupportBundle
	isSet bool
}

func (v NullableTechsupportmanagementTechSupportBundle) Get() *TechsupportmanagementTechSupportBundle {
	return v.value
}

func (v *NullableTechsupportmanagementTechSupportBundle) Set(val *TechsupportmanagementTechSupportBundle) {
	v.value = val
	v.isSet = true
}

func (v NullableTechsupportmanagementTechSupportBundle) IsSet() bool {
	return v.isSet
}

func (v *NullableTechsupportmanagementTechSupportBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechsupportmanagementTechSupportBundle(val *TechsupportmanagementTechSupportBundle) *NullableTechsupportmanagementTechSupportBundle {
	return &NullableTechsupportmanagementTechSupportBundle{value: val, isSet: true}
}

func (v NullableTechsupportmanagementTechSupportBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechsupportmanagementTechSupportBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


