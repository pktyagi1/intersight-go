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

// CloudSkuRegionRateCardsAllOf Definition of the list of properties defined in 'cloud.SkuRegionRateCards', excluding properties defined in parent classes.
type CloudSkuRegionRateCardsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The currency code used for the price. For e.g. USD or EUR. * `USD` - The currency code for United states dollar. * `EUR` - The currency code for European Union.
	Currency *string `json:"Currency,omitempty"`
	CustomAttributes []CloudCustomAttributes `json:"CustomAttributes,omitempty"`
	// The OS distribution running on this instance type.
	DistributionType *string `json:"DistributionType,omitempty"`
	// Applicable for private cloud where price is set by the user.
	IsUserDefined *bool `json:"IsUserDefined,omitempty"`
	// The platformType for this SKU. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `IMCRack` - A standalone UCS M6 and above server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `IWE` - An Intersight Workload Engine. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance. * `IntersightAssist` - A Cisco Intersight Assist. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `NexusDevice` - A generic platform type to support Nexus Network Device. This can also be extended to support all network devices later on. * `MDSDevice` - A platform type to support MDS devices. * `UCSC890` - A standalone Cisco UCSC890 server. * `NetAppOntap` - A NetApp ONTAP storage system. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `NewRelic` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * `DellCompellent` - A Dell Compellent storage system. * `HPE3Par` - A HPE 3PAR storage system. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * `IMCBlade` - An Intersight managed UCS Blade Server. * `TerraformCloud` - A Terraform Cloud account. * `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `AnsibleEndpoint` - An external endpoint added as Target that can be accessed through Ansible in Intersight Cloud Orchestrator automation workflow. * `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic, Bearer Token. * `SSHEndpoint` - An external endpoint added as Target that can be accessed through SSH in Intersight Cloud Orchestrator automation workflow. * `CiscoCatalyst` - A Cisco Catalyst networking switch device. * `PowerShellEndpoint` - A Windows machine on which PowerShell scripts can be executed remotely.
	PlatformType *string `json:"PlatformType,omitempty"`
	// Price of the instance type.
	Price *float32 `json:"Price,omitempty"`
	// The regionId associated with this rate card.
	RegionId *string `json:"RegionId,omitempty"`
	// Indicates if this sku belongs to Compute, Storage, Database or Network category. * `Compute` - Compute service offered by cloud provider. * `Storage` - Storage service offered by cloud provider. * `Database` - Database service offered by cloud provider. * `Network` - Network service offered by cloud provider.
	ServiceCategory *string `json:"ServiceCategory,omitempty"`
	// The sku name associated with this rate card.
	SkuName *string `json:"SkuName,omitempty"`
	// The billing unit to use for computing the price. For e.g. when serviceCategory is Compute the unit will be \"Hrs\", for Storage it will be \"GB-Mo\".
	Unit *string `json:"Unit,omitempty"`
	// The epoch start time from which the price will be applied.
	ValidFrom *int64 `json:"ValidFrom,omitempty"`
	// The epoch end time of the current price.
	ValidTo *int64 `json:"ValidTo,omitempty"`
	Region *CloudRegionsRelationship `json:"Region,omitempty"`
	Sku *CloudBaseSkuRelationship `json:"Sku,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudSkuRegionRateCardsAllOf CloudSkuRegionRateCardsAllOf

// NewCloudSkuRegionRateCardsAllOf instantiates a new CloudSkuRegionRateCardsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSkuRegionRateCardsAllOf(classId string, objectType string) *CloudSkuRegionRateCardsAllOf {
	this := CloudSkuRegionRateCardsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var currency string = "USD"
	this.Currency = &currency
	var platformType string = ""
	this.PlatformType = &platformType
	var serviceCategory string = "Compute"
	this.ServiceCategory = &serviceCategory
	return &this
}

// NewCloudSkuRegionRateCardsAllOfWithDefaults instantiates a new CloudSkuRegionRateCardsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSkuRegionRateCardsAllOfWithDefaults() *CloudSkuRegionRateCardsAllOf {
	this := CloudSkuRegionRateCardsAllOf{}
	var classId string = "cloud.SkuRegionRateCards"
	this.ClassId = classId
	var objectType string = "cloud.SkuRegionRateCards"
	this.ObjectType = objectType
	var currency string = "USD"
	this.Currency = &currency
	var platformType string = ""
	this.PlatformType = &platformType
	var serviceCategory string = "Compute"
	this.ServiceCategory = &serviceCategory
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudSkuRegionRateCardsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudSkuRegionRateCardsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudSkuRegionRateCardsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudSkuRegionRateCardsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudSkuRegionRateCardsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudSkuRegionRateCardsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CloudSkuRegionRateCardsAllOf) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuRegionRateCardsAllOf) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CloudSkuRegionRateCardsAllOf) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CloudSkuRegionRateCardsAllOf) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomAttributes returns the CustomAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudSkuRegionRateCardsAllOf) GetCustomAttributes() []CloudCustomAttributes {
	if o == nil  {
		var ret []CloudCustomAttributes
		return ret
	}
	return o.CustomAttributes
}

// GetCustomAttributesOk returns a tuple with the CustomAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudSkuRegionRateCardsAllOf) GetCustomAttributesOk() (*[]CloudCustomAttributes, bool) {
	if o == nil || o.CustomAttributes == nil {
		return nil, false
	}
	return &o.CustomAttributes, true
}

// HasCustomAttributes returns a boolean if a field has been set.
func (o *CloudSkuRegionRateCardsAllOf) HasCustomAttributes() bool {
	if o != nil && o.CustomAttributes != nil {
		return true
	}

	return false
}

// SetCustomAttributes gets a reference to the given []CloudCustomAttributes and assigns it to the CustomAttributes field.
func (o *CloudSkuRegionRateCardsAllOf) SetCustomAttributes(v []CloudCustomAttributes) {
	o.CustomAttributes = v
}

// GetDistributionType returns the DistributionType field value if set, zero value otherwise.
func (o *CloudSkuRegionRateCardsAllOf) GetDistributionType() string {
	if o == nil || o.DistributionType == nil {
		var ret string
		return ret
	}
	return *o.DistributionType
}

// GetDistributionTypeOk returns a tuple with the DistributionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuRegionRateCardsAllOf) GetDistributionTypeOk() (*string, bool) {
	if o == nil || o.DistributionType == nil {
		return nil, false
	}
	return o.DistributionType, true
}

// HasDistributionType returns a boolean if a field has been set.
func (o *CloudSkuRegionRateCardsAllOf) HasDistributionType() bool {
	if o != nil && o.DistributionType != nil {
		return true
	}

	return false
}

// SetDistributionType gets a reference to the given string and assigns it to the DistributionType field.
func (o *CloudSkuRegionRateCardsAllOf) SetDistributionType(v string) {
	o.DistributionType = &v
}

// GetIsUserDefined returns the IsUserDefined field value if set, zero value otherwise.
func (o *CloudSkuRegionRateCardsAllOf) GetIsUserDefined() bool {
	if o == nil || o.IsUserDefined == nil {
		var ret bool
		return ret
	}
	return *o.IsUserDefined
}

// GetIsUserDefinedOk returns a tuple with the IsUserDefined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuRegionRateCardsAllOf) GetIsUserDefinedOk() (*bool, bool) {
	if o == nil || o.IsUserDefined == nil {
		return nil, false
	}
	return o.IsUserDefined, true
}

// HasIsUserDefined returns a boolean if a field has been set.
func (o *CloudSkuRegionRateCardsAllOf) HasIsUserDefined() bool {
	if o != nil && o.IsUserDefined != nil {
		return true
	}

	return false
}

// SetIsUserDefined gets a reference to the given bool and assigns it to the IsUserDefined field.
func (o *CloudSkuRegionRateCardsAllOf) SetIsUserDefined(v bool) {
	o.IsUserDefined = &v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *CloudSkuRegionRateCardsAllOf) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuRegionRateCardsAllOf) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *CloudSkuRegionRateCardsAllOf) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *CloudSkuRegionRateCardsAllOf) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CloudSkuRegionRateCardsAllOf) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuRegionRateCardsAllOf) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CloudSkuRegionRateCardsAllOf) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *CloudSkuRegionRateCardsAllOf) SetPrice(v float32) {
	o.Price = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *CloudSkuRegionRateCardsAllOf) GetRegionId() string {
	if o == nil || o.RegionId == nil {
		var ret string
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuRegionRateCardsAllOf) GetRegionIdOk() (*string, bool) {
	if o == nil || o.RegionId == nil {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *CloudSkuRegionRateCardsAllOf) HasRegionId() bool {
	if o != nil && o.RegionId != nil {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given string and assigns it to the RegionId field.
func (o *CloudSkuRegionRateCardsAllOf) SetRegionId(v string) {
	o.RegionId = &v
}

// GetServiceCategory returns the ServiceCategory field value if set, zero value otherwise.
func (o *CloudSkuRegionRateCardsAllOf) GetServiceCategory() string {
	if o == nil || o.ServiceCategory == nil {
		var ret string
		return ret
	}
	return *o.ServiceCategory
}

// GetServiceCategoryOk returns a tuple with the ServiceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuRegionRateCardsAllOf) GetServiceCategoryOk() (*string, bool) {
	if o == nil || o.ServiceCategory == nil {
		return nil, false
	}
	return o.ServiceCategory, true
}

// HasServiceCategory returns a boolean if a field has been set.
func (o *CloudSkuRegionRateCardsAllOf) HasServiceCategory() bool {
	if o != nil && o.ServiceCategory != nil {
		return true
	}

	return false
}

// SetServiceCategory gets a reference to the given string and assigns it to the ServiceCategory field.
func (o *CloudSkuRegionRateCardsAllOf) SetServiceCategory(v string) {
	o.ServiceCategory = &v
}

// GetSkuName returns the SkuName field value if set, zero value otherwise.
func (o *CloudSkuRegionRateCardsAllOf) GetSkuName() string {
	if o == nil || o.SkuName == nil {
		var ret string
		return ret
	}
	return *o.SkuName
}

// GetSkuNameOk returns a tuple with the SkuName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuRegionRateCardsAllOf) GetSkuNameOk() (*string, bool) {
	if o == nil || o.SkuName == nil {
		return nil, false
	}
	return o.SkuName, true
}

// HasSkuName returns a boolean if a field has been set.
func (o *CloudSkuRegionRateCardsAllOf) HasSkuName() bool {
	if o != nil && o.SkuName != nil {
		return true
	}

	return false
}

// SetSkuName gets a reference to the given string and assigns it to the SkuName field.
func (o *CloudSkuRegionRateCardsAllOf) SetSkuName(v string) {
	o.SkuName = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *CloudSkuRegionRateCardsAllOf) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuRegionRateCardsAllOf) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *CloudSkuRegionRateCardsAllOf) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *CloudSkuRegionRateCardsAllOf) SetUnit(v string) {
	o.Unit = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *CloudSkuRegionRateCardsAllOf) GetValidFrom() int64 {
	if o == nil || o.ValidFrom == nil {
		var ret int64
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuRegionRateCardsAllOf) GetValidFromOk() (*int64, bool) {
	if o == nil || o.ValidFrom == nil {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *CloudSkuRegionRateCardsAllOf) HasValidFrom() bool {
	if o != nil && o.ValidFrom != nil {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given int64 and assigns it to the ValidFrom field.
func (o *CloudSkuRegionRateCardsAllOf) SetValidFrom(v int64) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *CloudSkuRegionRateCardsAllOf) GetValidTo() int64 {
	if o == nil || o.ValidTo == nil {
		var ret int64
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuRegionRateCardsAllOf) GetValidToOk() (*int64, bool) {
	if o == nil || o.ValidTo == nil {
		return nil, false
	}
	return o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *CloudSkuRegionRateCardsAllOf) HasValidTo() bool {
	if o != nil && o.ValidTo != nil {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given int64 and assigns it to the ValidTo field.
func (o *CloudSkuRegionRateCardsAllOf) SetValidTo(v int64) {
	o.ValidTo = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CloudSkuRegionRateCardsAllOf) GetRegion() CloudRegionsRelationship {
	if o == nil || o.Region == nil {
		var ret CloudRegionsRelationship
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuRegionRateCardsAllOf) GetRegionOk() (*CloudRegionsRelationship, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CloudSkuRegionRateCardsAllOf) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given CloudRegionsRelationship and assigns it to the Region field.
func (o *CloudSkuRegionRateCardsAllOf) SetRegion(v CloudRegionsRelationship) {
	o.Region = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *CloudSkuRegionRateCardsAllOf) GetSku() CloudBaseSkuRelationship {
	if o == nil || o.Sku == nil {
		var ret CloudBaseSkuRelationship
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuRegionRateCardsAllOf) GetSkuOk() (*CloudBaseSkuRelationship, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *CloudSkuRegionRateCardsAllOf) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given CloudBaseSkuRelationship and assigns it to the Sku field.
func (o *CloudSkuRegionRateCardsAllOf) SetSku(v CloudBaseSkuRelationship) {
	o.Sku = &v
}

func (o CloudSkuRegionRateCardsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Currency != nil {
		toSerialize["Currency"] = o.Currency
	}
	if o.CustomAttributes != nil {
		toSerialize["CustomAttributes"] = o.CustomAttributes
	}
	if o.DistributionType != nil {
		toSerialize["DistributionType"] = o.DistributionType
	}
	if o.IsUserDefined != nil {
		toSerialize["IsUserDefined"] = o.IsUserDefined
	}
	if o.PlatformType != nil {
		toSerialize["PlatformType"] = o.PlatformType
	}
	if o.Price != nil {
		toSerialize["Price"] = o.Price
	}
	if o.RegionId != nil {
		toSerialize["RegionId"] = o.RegionId
	}
	if o.ServiceCategory != nil {
		toSerialize["ServiceCategory"] = o.ServiceCategory
	}
	if o.SkuName != nil {
		toSerialize["SkuName"] = o.SkuName
	}
	if o.Unit != nil {
		toSerialize["Unit"] = o.Unit
	}
	if o.ValidFrom != nil {
		toSerialize["ValidFrom"] = o.ValidFrom
	}
	if o.ValidTo != nil {
		toSerialize["ValidTo"] = o.ValidTo
	}
	if o.Region != nil {
		toSerialize["Region"] = o.Region
	}
	if o.Sku != nil {
		toSerialize["Sku"] = o.Sku
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudSkuRegionRateCardsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudSkuRegionRateCardsAllOf := _CloudSkuRegionRateCardsAllOf{}

	if err = json.Unmarshal(bytes, &varCloudSkuRegionRateCardsAllOf); err == nil {
		*o = CloudSkuRegionRateCardsAllOf(varCloudSkuRegionRateCardsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Currency")
		delete(additionalProperties, "CustomAttributes")
		delete(additionalProperties, "DistributionType")
		delete(additionalProperties, "IsUserDefined")
		delete(additionalProperties, "PlatformType")
		delete(additionalProperties, "Price")
		delete(additionalProperties, "RegionId")
		delete(additionalProperties, "ServiceCategory")
		delete(additionalProperties, "SkuName")
		delete(additionalProperties, "Unit")
		delete(additionalProperties, "ValidFrom")
		delete(additionalProperties, "ValidTo")
		delete(additionalProperties, "Region")
		delete(additionalProperties, "Sku")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudSkuRegionRateCardsAllOf struct {
	value *CloudSkuRegionRateCardsAllOf
	isSet bool
}

func (v NullableCloudSkuRegionRateCardsAllOf) Get() *CloudSkuRegionRateCardsAllOf {
	return v.value
}

func (v *NullableCloudSkuRegionRateCardsAllOf) Set(val *CloudSkuRegionRateCardsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSkuRegionRateCardsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSkuRegionRateCardsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSkuRegionRateCardsAllOf(val *CloudSkuRegionRateCardsAllOf) *NullableCloudSkuRegionRateCardsAllOf {
	return &NullableCloudSkuRegionRateCardsAllOf{value: val, isSet: true}
}

func (v NullableCloudSkuRegionRateCardsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSkuRegionRateCardsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


