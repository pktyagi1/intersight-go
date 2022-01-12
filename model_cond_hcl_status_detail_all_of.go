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

// CondHclStatusDetailAllOf Definition of the list of properties defined in 'cond.HclStatusDetail', excluding properties defined in parent classes.
type CondHclStatusDetailAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The model is considered as part of the hardware profile for the component. This will provide the HCL validation status for the hardware profile. The reasons can be one of the following \"Incompatible-Server-With-Component\" - the server model and component combination is not listed in HCL \"Incompatible-Firmware\" - The server's firmware is not listed for this component's hardware profile \"Incompatible-Component\" - the component's model is not listed in the HCL \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Not-Evaluated\" - the hardware profile was not evaulated for the component because the server's hw/sw status is not listed or server is exempted. \"Compatible\" - this component's hardware profile is listed in the HCL. * `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Incompatible-Server-With-Component` - The validation failed for this component because he server model and component model combination was not found in the HCL. * `Incompatible-Processor` - The validation failed because the given processor was not found for the given server PID. * `Incompatible-Os-Info` - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination. * `Incompatible-Component-Model` - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version. * `Incompatible-Firmware` - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model. * `Incompatible-Driver` - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware. * `Incompatible-Firmware-Driver` - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware. * `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * `Service-Error` - The validation has failed because the HCL data service has return a service error or unrecognized result. * `Unrecognized-Protocol` - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values. * `Not-Evaluated` - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated. * `Compatible` - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version.
	HardwareStatus *string `json:"HardwareStatus,omitempty"`
	// The current CIMC version for the server normalized for querying HCL data.
	HclCimcVersion *string `json:"HclCimcVersion,omitempty"`
	// The current driver name of the component we are validating normalized for querying HCL data.
	HclDriverName *string `json:"HclDriverName,omitempty"`
	// The current driver version of the component we are validating normalized for querying HCL data.
	HclDriverVersion *string `json:"HclDriverVersion,omitempty"`
	// The current firmware version of the component model normalized for querying HCL data.
	HclFirmwareVersion *string `json:"HclFirmwareVersion,omitempty"`
	// The component model we are trying to validate normalized for querying HCL data.
	HclModel *string `json:"HclModel,omitempty"`
	// The current CIMC version for the server as received from inventory.
	InvCimcVersion *string `json:"InvCimcVersion,omitempty"`
	// The current driver name of the component we are validating as received from inventory.
	InvDriverName *string `json:"InvDriverName,omitempty"`
	// The current driver version of the component we are validating as received from inventory.
	InvDriverVersion *string `json:"InvDriverVersion,omitempty"`
	// The current firmware version of the component model as received from inventory.
	InvFirmwareVersion *string `json:"InvFirmwareVersion,omitempty"`
	// The component model we are trying to validate as received from inventory.
	InvModel *string `json:"InvModel,omitempty"`
	// The reason for the status. The reason can be one of \"Incompatible-Server-With-Component\" - HCL validation has failed because the server model is not validated with this component \"Incompatible-Processor\" - HCL validation has failed because the processor is not validated with this server \"Incompatible-Os-Info\" - HCL validation has failed because the os vendor and version is not validated with this server \"Incompatible-Component-Model\" - HCL validation has failed because the component model is not validated \"Incompatible-Firmware\" - HCL validation has failed because the component or server firmware version is not validated \"Incompatible-Driver\" - HCL validation has failed because the driver version is not validated \"Incompatible-Firmware-Driver\" - HCL validation has failed because the firmware version and driver version is not validated \"Missing-Os-Driver-Info\" - HCL validation was not performed because we are missing os driver information form the inventory \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Service-Error\" - HCL data service is available but an error occured when making the request or parsing the response \"Unrecognized-Protocol\" - This service does not recognize the reason code in the response from the HCL data service \"Compatible\" - this component's inventory data has \"Validated\" status with the HCL. \"Not-Evaluated\" - The component is not evaluated against the HCL because the server is exempted. * `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Incompatible-Server-With-Component` - The validation failed for this component because he server model and component model combination was not found in the HCL. * `Incompatible-Processor` - The validation failed because the given processor was not found for the given server PID. * `Incompatible-Os-Info` - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination. * `Incompatible-Component-Model` - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version. * `Incompatible-Firmware` - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model. * `Incompatible-Driver` - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware. * `Incompatible-Firmware-Driver` - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware. * `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * `Service-Error` - The validation has failed because the HCL data service has return a service error or unrecognized result. * `Unrecognized-Protocol` - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values. * `Not-Evaluated` - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated. * `Compatible` - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version.
	Reason *string `json:"Reason,omitempty"`
	// The firmware, driver name and driver version are considered as part of the software profile for the component. This will provide the HCL validation status for the software profile. The reasons can be one of the following \"Incompatible-Firmware\" - the component's firmware is not listed under the server's hardware and software profile and the component's hardware profile \"Incompatible-Driver\" - the component's driver is not listed under the server's hardware and software profile and the component's hardware profile \"Incompatible-Firmware-Driver\" - the component's firmware and driver are not listed under the server's hardware and software profile and the component's hardware profile \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Not-Evaluated\" - the component's hardware status was not evaluated because the server's hardware or software profile is not listed or server is exempted. \"Compatible\" - this component's hardware profile is listed in the HCL. * `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Incompatible-Server-With-Component` - The validation failed for this component because he server model and component model combination was not found in the HCL. * `Incompatible-Processor` - The validation failed because the given processor was not found for the given server PID. * `Incompatible-Os-Info` - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination. * `Incompatible-Component-Model` - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version. * `Incompatible-Firmware` - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model. * `Incompatible-Driver` - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware. * `Incompatible-Firmware-Driver` - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware. * `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * `Service-Error` - The validation has failed because the HCL data service has return a service error or unrecognized result. * `Unrecognized-Protocol` - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values. * `Not-Evaluated` - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated. * `Compatible` - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version.
	SoftwareStatus *string `json:"SoftwareStatus,omitempty"`
	// The status for the component model, firmware version, driver name, and driver version after validating against the HCL. The status can be one of the following \"Unknown\" - we do not have enough information to evaluate against the HCL data \"Validated\" - we have validated this component against the HCL and it has \"Validated\" status \"Not-Validated\" - we have validated this component against the HCL and it has \"Not-Validated\" status. \"Not-Evaluated\" - The component is not evaluated against the HCL because the server is exempted. * `Incomplete` - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL. * `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL. * `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL. * `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.
	Status *string `json:"Status,omitempty"`
	Component *InventoryBaseRelationship `json:"Component,omitempty"`
	HclStatus *CondHclStatusRelationship `json:"HclStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CondHclStatusDetailAllOf CondHclStatusDetailAllOf

// NewCondHclStatusDetailAllOf instantiates a new CondHclStatusDetailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondHclStatusDetailAllOf(classId string, objectType string) *CondHclStatusDetailAllOf {
	this := CondHclStatusDetailAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hardwareStatus string = "Missing-Os-Driver-Info"
	this.HardwareStatus = &hardwareStatus
	var reason string = "Missing-Os-Driver-Info"
	this.Reason = &reason
	var softwareStatus string = "Missing-Os-Driver-Info"
	this.SoftwareStatus = &softwareStatus
	var status string = "Incomplete"
	this.Status = &status
	return &this
}

// NewCondHclStatusDetailAllOfWithDefaults instantiates a new CondHclStatusDetailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCondHclStatusDetailAllOfWithDefaults() *CondHclStatusDetailAllOf {
	this := CondHclStatusDetailAllOf{}
	var classId string = "cond.HclStatusDetail"
	this.ClassId = classId
	var objectType string = "cond.HclStatusDetail"
	this.ObjectType = objectType
	var hardwareStatus string = "Missing-Os-Driver-Info"
	this.HardwareStatus = &hardwareStatus
	var reason string = "Missing-Os-Driver-Info"
	this.Reason = &reason
	var softwareStatus string = "Missing-Os-Driver-Info"
	this.SoftwareStatus = &softwareStatus
	var status string = "Incomplete"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *CondHclStatusDetailAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CondHclStatusDetailAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CondHclStatusDetailAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CondHclStatusDetailAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHardwareStatus returns the HardwareStatus field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetHardwareStatus() string {
	if o == nil || o.HardwareStatus == nil {
		var ret string
		return ret
	}
	return *o.HardwareStatus
}

// GetHardwareStatusOk returns a tuple with the HardwareStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetHardwareStatusOk() (*string, bool) {
	if o == nil || o.HardwareStatus == nil {
		return nil, false
	}
	return o.HardwareStatus, true
}

// HasHardwareStatus returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasHardwareStatus() bool {
	if o != nil && o.HardwareStatus != nil {
		return true
	}

	return false
}

// SetHardwareStatus gets a reference to the given string and assigns it to the HardwareStatus field.
func (o *CondHclStatusDetailAllOf) SetHardwareStatus(v string) {
	o.HardwareStatus = &v
}

// GetHclCimcVersion returns the HclCimcVersion field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetHclCimcVersion() string {
	if o == nil || o.HclCimcVersion == nil {
		var ret string
		return ret
	}
	return *o.HclCimcVersion
}

// GetHclCimcVersionOk returns a tuple with the HclCimcVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetHclCimcVersionOk() (*string, bool) {
	if o == nil || o.HclCimcVersion == nil {
		return nil, false
	}
	return o.HclCimcVersion, true
}

// HasHclCimcVersion returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasHclCimcVersion() bool {
	if o != nil && o.HclCimcVersion != nil {
		return true
	}

	return false
}

// SetHclCimcVersion gets a reference to the given string and assigns it to the HclCimcVersion field.
func (o *CondHclStatusDetailAllOf) SetHclCimcVersion(v string) {
	o.HclCimcVersion = &v
}

// GetHclDriverName returns the HclDriverName field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetHclDriverName() string {
	if o == nil || o.HclDriverName == nil {
		var ret string
		return ret
	}
	return *o.HclDriverName
}

// GetHclDriverNameOk returns a tuple with the HclDriverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetHclDriverNameOk() (*string, bool) {
	if o == nil || o.HclDriverName == nil {
		return nil, false
	}
	return o.HclDriverName, true
}

// HasHclDriverName returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasHclDriverName() bool {
	if o != nil && o.HclDriverName != nil {
		return true
	}

	return false
}

// SetHclDriverName gets a reference to the given string and assigns it to the HclDriverName field.
func (o *CondHclStatusDetailAllOf) SetHclDriverName(v string) {
	o.HclDriverName = &v
}

// GetHclDriverVersion returns the HclDriverVersion field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetHclDriverVersion() string {
	if o == nil || o.HclDriverVersion == nil {
		var ret string
		return ret
	}
	return *o.HclDriverVersion
}

// GetHclDriverVersionOk returns a tuple with the HclDriverVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetHclDriverVersionOk() (*string, bool) {
	if o == nil || o.HclDriverVersion == nil {
		return nil, false
	}
	return o.HclDriverVersion, true
}

// HasHclDriverVersion returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasHclDriverVersion() bool {
	if o != nil && o.HclDriverVersion != nil {
		return true
	}

	return false
}

// SetHclDriverVersion gets a reference to the given string and assigns it to the HclDriverVersion field.
func (o *CondHclStatusDetailAllOf) SetHclDriverVersion(v string) {
	o.HclDriverVersion = &v
}

// GetHclFirmwareVersion returns the HclFirmwareVersion field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetHclFirmwareVersion() string {
	if o == nil || o.HclFirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.HclFirmwareVersion
}

// GetHclFirmwareVersionOk returns a tuple with the HclFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetHclFirmwareVersionOk() (*string, bool) {
	if o == nil || o.HclFirmwareVersion == nil {
		return nil, false
	}
	return o.HclFirmwareVersion, true
}

// HasHclFirmwareVersion returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasHclFirmwareVersion() bool {
	if o != nil && o.HclFirmwareVersion != nil {
		return true
	}

	return false
}

// SetHclFirmwareVersion gets a reference to the given string and assigns it to the HclFirmwareVersion field.
func (o *CondHclStatusDetailAllOf) SetHclFirmwareVersion(v string) {
	o.HclFirmwareVersion = &v
}

// GetHclModel returns the HclModel field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetHclModel() string {
	if o == nil || o.HclModel == nil {
		var ret string
		return ret
	}
	return *o.HclModel
}

// GetHclModelOk returns a tuple with the HclModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetHclModelOk() (*string, bool) {
	if o == nil || o.HclModel == nil {
		return nil, false
	}
	return o.HclModel, true
}

// HasHclModel returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasHclModel() bool {
	if o != nil && o.HclModel != nil {
		return true
	}

	return false
}

// SetHclModel gets a reference to the given string and assigns it to the HclModel field.
func (o *CondHclStatusDetailAllOf) SetHclModel(v string) {
	o.HclModel = &v
}

// GetInvCimcVersion returns the InvCimcVersion field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetInvCimcVersion() string {
	if o == nil || o.InvCimcVersion == nil {
		var ret string
		return ret
	}
	return *o.InvCimcVersion
}

// GetInvCimcVersionOk returns a tuple with the InvCimcVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetInvCimcVersionOk() (*string, bool) {
	if o == nil || o.InvCimcVersion == nil {
		return nil, false
	}
	return o.InvCimcVersion, true
}

// HasInvCimcVersion returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasInvCimcVersion() bool {
	if o != nil && o.InvCimcVersion != nil {
		return true
	}

	return false
}

// SetInvCimcVersion gets a reference to the given string and assigns it to the InvCimcVersion field.
func (o *CondHclStatusDetailAllOf) SetInvCimcVersion(v string) {
	o.InvCimcVersion = &v
}

// GetInvDriverName returns the InvDriverName field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetInvDriverName() string {
	if o == nil || o.InvDriverName == nil {
		var ret string
		return ret
	}
	return *o.InvDriverName
}

// GetInvDriverNameOk returns a tuple with the InvDriverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetInvDriverNameOk() (*string, bool) {
	if o == nil || o.InvDriverName == nil {
		return nil, false
	}
	return o.InvDriverName, true
}

// HasInvDriverName returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasInvDriverName() bool {
	if o != nil && o.InvDriverName != nil {
		return true
	}

	return false
}

// SetInvDriverName gets a reference to the given string and assigns it to the InvDriverName field.
func (o *CondHclStatusDetailAllOf) SetInvDriverName(v string) {
	o.InvDriverName = &v
}

// GetInvDriverVersion returns the InvDriverVersion field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetInvDriverVersion() string {
	if o == nil || o.InvDriverVersion == nil {
		var ret string
		return ret
	}
	return *o.InvDriverVersion
}

// GetInvDriverVersionOk returns a tuple with the InvDriverVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetInvDriverVersionOk() (*string, bool) {
	if o == nil || o.InvDriverVersion == nil {
		return nil, false
	}
	return o.InvDriverVersion, true
}

// HasInvDriverVersion returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasInvDriverVersion() bool {
	if o != nil && o.InvDriverVersion != nil {
		return true
	}

	return false
}

// SetInvDriverVersion gets a reference to the given string and assigns it to the InvDriverVersion field.
func (o *CondHclStatusDetailAllOf) SetInvDriverVersion(v string) {
	o.InvDriverVersion = &v
}

// GetInvFirmwareVersion returns the InvFirmwareVersion field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetInvFirmwareVersion() string {
	if o == nil || o.InvFirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.InvFirmwareVersion
}

// GetInvFirmwareVersionOk returns a tuple with the InvFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetInvFirmwareVersionOk() (*string, bool) {
	if o == nil || o.InvFirmwareVersion == nil {
		return nil, false
	}
	return o.InvFirmwareVersion, true
}

// HasInvFirmwareVersion returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasInvFirmwareVersion() bool {
	if o != nil && o.InvFirmwareVersion != nil {
		return true
	}

	return false
}

// SetInvFirmwareVersion gets a reference to the given string and assigns it to the InvFirmwareVersion field.
func (o *CondHclStatusDetailAllOf) SetInvFirmwareVersion(v string) {
	o.InvFirmwareVersion = &v
}

// GetInvModel returns the InvModel field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetInvModel() string {
	if o == nil || o.InvModel == nil {
		var ret string
		return ret
	}
	return *o.InvModel
}

// GetInvModelOk returns a tuple with the InvModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetInvModelOk() (*string, bool) {
	if o == nil || o.InvModel == nil {
		return nil, false
	}
	return o.InvModel, true
}

// HasInvModel returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasInvModel() bool {
	if o != nil && o.InvModel != nil {
		return true
	}

	return false
}

// SetInvModel gets a reference to the given string and assigns it to the InvModel field.
func (o *CondHclStatusDetailAllOf) SetInvModel(v string) {
	o.InvModel = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *CondHclStatusDetailAllOf) SetReason(v string) {
	o.Reason = &v
}

// GetSoftwareStatus returns the SoftwareStatus field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetSoftwareStatus() string {
	if o == nil || o.SoftwareStatus == nil {
		var ret string
		return ret
	}
	return *o.SoftwareStatus
}

// GetSoftwareStatusOk returns a tuple with the SoftwareStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetSoftwareStatusOk() (*string, bool) {
	if o == nil || o.SoftwareStatus == nil {
		return nil, false
	}
	return o.SoftwareStatus, true
}

// HasSoftwareStatus returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasSoftwareStatus() bool {
	if o != nil && o.SoftwareStatus != nil {
		return true
	}

	return false
}

// SetSoftwareStatus gets a reference to the given string and assigns it to the SoftwareStatus field.
func (o *CondHclStatusDetailAllOf) SetSoftwareStatus(v string) {
	o.SoftwareStatus = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CondHclStatusDetailAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetComponent() InventoryBaseRelationship {
	if o == nil || o.Component == nil {
		var ret InventoryBaseRelationship
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetComponentOk() (*InventoryBaseRelationship, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given InventoryBaseRelationship and assigns it to the Component field.
func (o *CondHclStatusDetailAllOf) SetComponent(v InventoryBaseRelationship) {
	o.Component = &v
}

// GetHclStatus returns the HclStatus field value if set, zero value otherwise.
func (o *CondHclStatusDetailAllOf) GetHclStatus() CondHclStatusRelationship {
	if o == nil || o.HclStatus == nil {
		var ret CondHclStatusRelationship
		return ret
	}
	return *o.HclStatus
}

// GetHclStatusOk returns a tuple with the HclStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetailAllOf) GetHclStatusOk() (*CondHclStatusRelationship, bool) {
	if o == nil || o.HclStatus == nil {
		return nil, false
	}
	return o.HclStatus, true
}

// HasHclStatus returns a boolean if a field has been set.
func (o *CondHclStatusDetailAllOf) HasHclStatus() bool {
	if o != nil && o.HclStatus != nil {
		return true
	}

	return false
}

// SetHclStatus gets a reference to the given CondHclStatusRelationship and assigns it to the HclStatus field.
func (o *CondHclStatusDetailAllOf) SetHclStatus(v CondHclStatusRelationship) {
	o.HclStatus = &v
}

func (o CondHclStatusDetailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.HardwareStatus != nil {
		toSerialize["HardwareStatus"] = o.HardwareStatus
	}
	if o.HclCimcVersion != nil {
		toSerialize["HclCimcVersion"] = o.HclCimcVersion
	}
	if o.HclDriverName != nil {
		toSerialize["HclDriverName"] = o.HclDriverName
	}
	if o.HclDriverVersion != nil {
		toSerialize["HclDriverVersion"] = o.HclDriverVersion
	}
	if o.HclFirmwareVersion != nil {
		toSerialize["HclFirmwareVersion"] = o.HclFirmwareVersion
	}
	if o.HclModel != nil {
		toSerialize["HclModel"] = o.HclModel
	}
	if o.InvCimcVersion != nil {
		toSerialize["InvCimcVersion"] = o.InvCimcVersion
	}
	if o.InvDriverName != nil {
		toSerialize["InvDriverName"] = o.InvDriverName
	}
	if o.InvDriverVersion != nil {
		toSerialize["InvDriverVersion"] = o.InvDriverVersion
	}
	if o.InvFirmwareVersion != nil {
		toSerialize["InvFirmwareVersion"] = o.InvFirmwareVersion
	}
	if o.InvModel != nil {
		toSerialize["InvModel"] = o.InvModel
	}
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}
	if o.SoftwareStatus != nil {
		toSerialize["SoftwareStatus"] = o.SoftwareStatus
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Component != nil {
		toSerialize["Component"] = o.Component
	}
	if o.HclStatus != nil {
		toSerialize["HclStatus"] = o.HclStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CondHclStatusDetailAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCondHclStatusDetailAllOf := _CondHclStatusDetailAllOf{}

	if err = json.Unmarshal(bytes, &varCondHclStatusDetailAllOf); err == nil {
		*o = CondHclStatusDetailAllOf(varCondHclStatusDetailAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HardwareStatus")
		delete(additionalProperties, "HclCimcVersion")
		delete(additionalProperties, "HclDriverName")
		delete(additionalProperties, "HclDriverVersion")
		delete(additionalProperties, "HclFirmwareVersion")
		delete(additionalProperties, "HclModel")
		delete(additionalProperties, "InvCimcVersion")
		delete(additionalProperties, "InvDriverName")
		delete(additionalProperties, "InvDriverVersion")
		delete(additionalProperties, "InvFirmwareVersion")
		delete(additionalProperties, "InvModel")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "SoftwareStatus")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Component")
		delete(additionalProperties, "HclStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCondHclStatusDetailAllOf struct {
	value *CondHclStatusDetailAllOf
	isSet bool
}

func (v NullableCondHclStatusDetailAllOf) Get() *CondHclStatusDetailAllOf {
	return v.value
}

func (v *NullableCondHclStatusDetailAllOf) Set(val *CondHclStatusDetailAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCondHclStatusDetailAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCondHclStatusDetailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondHclStatusDetailAllOf(val *CondHclStatusDetailAllOf) *NullableCondHclStatusDetailAllOf {
	return &NullableCondHclStatusDetailAllOf{value: val, isSet: true}
}

func (v NullableCondHclStatusDetailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondHclStatusDetailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


