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

// BootPrecisionPolicy Boot order policy models a reusable boot order configuration that can be applied to multiple servers via profile association. It supports advanced boot order configuration on Cisco CIMC servers.
type BootPrecisionPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	BootDevices []BootDeviceBase `json:"BootDevices,omitempty"`
	// Sets the BIOS boot mode. UEFI uses the GUID Partition Table (GPT) whereas Legacy mode uses the Master Boot Record (MBR) partitioning scheme. To apply this setting, Please reboot the server. * `Legacy` - Legacy mode refers to the traditional process of booting from BIOS. Legacy mode uses the Master Boot Record (MBR) to locate the bootloader. * `Uefi` - UEFI mode uses the GUID Partition Table (GPT) to locate EFI Service Partitions to boot from.
	ConfiguredBootMode *string `json:"ConfiguredBootMode,omitempty"`
	// If UEFI secure boot is enabled, the boot mode is set to UEFI by default. Secure boot enforces that device boots using only software that is trusted by the Original Equipment Manufacturer (OEM).
	EnforceUefiSecureBoot *bool `json:"EnforceUefiSecureBoot,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootPrecisionPolicy BootPrecisionPolicy

// NewBootPrecisionPolicy instantiates a new BootPrecisionPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootPrecisionPolicy(classId string, objectType string) *BootPrecisionPolicy {
	this := BootPrecisionPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var configuredBootMode string = "Legacy"
	this.ConfiguredBootMode = &configuredBootMode
	var enforceUefiSecureBoot bool = false
	this.EnforceUefiSecureBoot = &enforceUefiSecureBoot
	return &this
}

// NewBootPrecisionPolicyWithDefaults instantiates a new BootPrecisionPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootPrecisionPolicyWithDefaults() *BootPrecisionPolicy {
	this := BootPrecisionPolicy{}
	var classId string = "boot.PrecisionPolicy"
	this.ClassId = classId
	var objectType string = "boot.PrecisionPolicy"
	this.ObjectType = objectType
	var configuredBootMode string = "Legacy"
	this.ConfiguredBootMode = &configuredBootMode
	var enforceUefiSecureBoot bool = false
	this.EnforceUefiSecureBoot = &enforceUefiSecureBoot
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootPrecisionPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootPrecisionPolicy) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootPrecisionPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootPrecisionPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootPrecisionPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootPrecisionPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootDevices returns the BootDevices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootPrecisionPolicy) GetBootDevices() []BootDeviceBase {
	if o == nil  {
		var ret []BootDeviceBase
		return ret
	}
	return o.BootDevices
}

// GetBootDevicesOk returns a tuple with the BootDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootPrecisionPolicy) GetBootDevicesOk() (*[]BootDeviceBase, bool) {
	if o == nil || o.BootDevices == nil {
		return nil, false
	}
	return &o.BootDevices, true
}

// HasBootDevices returns a boolean if a field has been set.
func (o *BootPrecisionPolicy) HasBootDevices() bool {
	if o != nil && o.BootDevices != nil {
		return true
	}

	return false
}

// SetBootDevices gets a reference to the given []BootDeviceBase and assigns it to the BootDevices field.
func (o *BootPrecisionPolicy) SetBootDevices(v []BootDeviceBase) {
	o.BootDevices = v
}

// GetConfiguredBootMode returns the ConfiguredBootMode field value if set, zero value otherwise.
func (o *BootPrecisionPolicy) GetConfiguredBootMode() string {
	if o == nil || o.ConfiguredBootMode == nil {
		var ret string
		return ret
	}
	return *o.ConfiguredBootMode
}

// GetConfiguredBootModeOk returns a tuple with the ConfiguredBootMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootPrecisionPolicy) GetConfiguredBootModeOk() (*string, bool) {
	if o == nil || o.ConfiguredBootMode == nil {
		return nil, false
	}
	return o.ConfiguredBootMode, true
}

// HasConfiguredBootMode returns a boolean if a field has been set.
func (o *BootPrecisionPolicy) HasConfiguredBootMode() bool {
	if o != nil && o.ConfiguredBootMode != nil {
		return true
	}

	return false
}

// SetConfiguredBootMode gets a reference to the given string and assigns it to the ConfiguredBootMode field.
func (o *BootPrecisionPolicy) SetConfiguredBootMode(v string) {
	o.ConfiguredBootMode = &v
}

// GetEnforceUefiSecureBoot returns the EnforceUefiSecureBoot field value if set, zero value otherwise.
func (o *BootPrecisionPolicy) GetEnforceUefiSecureBoot() bool {
	if o == nil || o.EnforceUefiSecureBoot == nil {
		var ret bool
		return ret
	}
	return *o.EnforceUefiSecureBoot
}

// GetEnforceUefiSecureBootOk returns a tuple with the EnforceUefiSecureBoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootPrecisionPolicy) GetEnforceUefiSecureBootOk() (*bool, bool) {
	if o == nil || o.EnforceUefiSecureBoot == nil {
		return nil, false
	}
	return o.EnforceUefiSecureBoot, true
}

// HasEnforceUefiSecureBoot returns a boolean if a field has been set.
func (o *BootPrecisionPolicy) HasEnforceUefiSecureBoot() bool {
	if o != nil && o.EnforceUefiSecureBoot != nil {
		return true
	}

	return false
}

// SetEnforceUefiSecureBoot gets a reference to the given bool and assigns it to the EnforceUefiSecureBoot field.
func (o *BootPrecisionPolicy) SetEnforceUefiSecureBoot(v bool) {
	o.EnforceUefiSecureBoot = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *BootPrecisionPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootPrecisionPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *BootPrecisionPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *BootPrecisionPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootPrecisionPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil  {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootPrecisionPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *BootPrecisionPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *BootPrecisionPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o BootPrecisionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BootDevices != nil {
		toSerialize["BootDevices"] = o.BootDevices
	}
	if o.ConfiguredBootMode != nil {
		toSerialize["ConfiguredBootMode"] = o.ConfiguredBootMode
	}
	if o.EnforceUefiSecureBoot != nil {
		toSerialize["EnforceUefiSecureBoot"] = o.EnforceUefiSecureBoot
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

func (o *BootPrecisionPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type BootPrecisionPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		BootDevices []BootDeviceBase `json:"BootDevices,omitempty"`
		// Sets the BIOS boot mode. UEFI uses the GUID Partition Table (GPT) whereas Legacy mode uses the Master Boot Record (MBR) partitioning scheme. To apply this setting, Please reboot the server. * `Legacy` - Legacy mode refers to the traditional process of booting from BIOS. Legacy mode uses the Master Boot Record (MBR) to locate the bootloader. * `Uefi` - UEFI mode uses the GUID Partition Table (GPT) to locate EFI Service Partitions to boot from.
		ConfiguredBootMode *string `json:"ConfiguredBootMode,omitempty"`
		// If UEFI secure boot is enabled, the boot mode is set to UEFI by default. Secure boot enforces that device boots using only software that is trusted by the Original Equipment Manufacturer (OEM).
		EnforceUefiSecureBoot *bool `json:"EnforceUefiSecureBoot,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varBootPrecisionPolicyWithoutEmbeddedStruct := BootPrecisionPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBootPrecisionPolicyWithoutEmbeddedStruct)
	if err == nil {
		varBootPrecisionPolicy := _BootPrecisionPolicy{}
		varBootPrecisionPolicy.ClassId = varBootPrecisionPolicyWithoutEmbeddedStruct.ClassId
		varBootPrecisionPolicy.ObjectType = varBootPrecisionPolicyWithoutEmbeddedStruct.ObjectType
		varBootPrecisionPolicy.BootDevices = varBootPrecisionPolicyWithoutEmbeddedStruct.BootDevices
		varBootPrecisionPolicy.ConfiguredBootMode = varBootPrecisionPolicyWithoutEmbeddedStruct.ConfiguredBootMode
		varBootPrecisionPolicy.EnforceUefiSecureBoot = varBootPrecisionPolicyWithoutEmbeddedStruct.EnforceUefiSecureBoot
		varBootPrecisionPolicy.Organization = varBootPrecisionPolicyWithoutEmbeddedStruct.Organization
		varBootPrecisionPolicy.Profiles = varBootPrecisionPolicyWithoutEmbeddedStruct.Profiles
		*o = BootPrecisionPolicy(varBootPrecisionPolicy)
	} else {
		return err
	}

	varBootPrecisionPolicy := _BootPrecisionPolicy{}

	err = json.Unmarshal(bytes, &varBootPrecisionPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varBootPrecisionPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BootDevices")
		delete(additionalProperties, "ConfiguredBootMode")
		delete(additionalProperties, "EnforceUefiSecureBoot")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableBootPrecisionPolicy struct {
	value *BootPrecisionPolicy
	isSet bool
}

func (v NullableBootPrecisionPolicy) Get() *BootPrecisionPolicy {
	return v.value
}

func (v *NullableBootPrecisionPolicy) Set(val *BootPrecisionPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableBootPrecisionPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableBootPrecisionPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootPrecisionPolicy(val *BootPrecisionPolicy) *NullableBootPrecisionPolicy {
	return &NullableBootPrecisionPolicy{value: val, isSet: true}
}

func (v NullableBootPrecisionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootPrecisionPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


