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

// NiatelemetryHttpsAclEpgContractMap Object to capture the HTTPS ACL contract EPG mappings in APIC.
type NiatelemetryHttpsAclEpgContractMap struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of HTTPS ACL contract for APIC.
	ContractName *string `json:"ContractName,omitempty"`
	// Dn of the HTTPS ACL contract provider for APIC.
	Dn *string `json:"Dn,omitempty"`
	// Name of EPGs of the HTTPS ACL contract for APIC.
	EpgName *string `json:"EpgName,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// TDn of the HTTPS ACL contract provider for APIC.
	TargetDn *string `json:"TargetDn,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryHttpsAclEpgContractMap NiatelemetryHttpsAclEpgContractMap

// NewNiatelemetryHttpsAclEpgContractMap instantiates a new NiatelemetryHttpsAclEpgContractMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryHttpsAclEpgContractMap(classId string, objectType string) *NiatelemetryHttpsAclEpgContractMap {
	this := NiatelemetryHttpsAclEpgContractMap{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryHttpsAclEpgContractMapWithDefaults instantiates a new NiatelemetryHttpsAclEpgContractMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryHttpsAclEpgContractMapWithDefaults() *NiatelemetryHttpsAclEpgContractMap {
	this := NiatelemetryHttpsAclEpgContractMap{}
	var classId string = "niatelemetry.HttpsAclEpgContractMap"
	this.ClassId = classId
	var objectType string = "niatelemetry.HttpsAclEpgContractMap"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryHttpsAclEpgContractMap) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryHttpsAclEpgContractMap) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryHttpsAclEpgContractMap) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryHttpsAclEpgContractMap) SetObjectType(v string) {
	o.ObjectType = v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMap) GetContractName() string {
	if o == nil || o.ContractName == nil {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) GetContractNameOk() (*string, bool) {
	if o == nil || o.ContractName == nil {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) HasContractName() bool {
	if o != nil && o.ContractName != nil {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *NiatelemetryHttpsAclEpgContractMap) SetContractName(v string) {
	o.ContractName = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMap) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryHttpsAclEpgContractMap) SetDn(v string) {
	o.Dn = &v
}

// GetEpgName returns the EpgName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMap) GetEpgName() string {
	if o == nil || o.EpgName == nil {
		var ret string
		return ret
	}
	return *o.EpgName
}

// GetEpgNameOk returns a tuple with the EpgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) GetEpgNameOk() (*string, bool) {
	if o == nil || o.EpgName == nil {
		return nil, false
	}
	return o.EpgName, true
}

// HasEpgName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) HasEpgName() bool {
	if o != nil && o.EpgName != nil {
		return true
	}

	return false
}

// SetEpgName gets a reference to the given string and assigns it to the EpgName field.
func (o *NiatelemetryHttpsAclEpgContractMap) SetEpgName(v string) {
	o.EpgName = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMap) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryHttpsAclEpgContractMap) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMap) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryHttpsAclEpgContractMap) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMap) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryHttpsAclEpgContractMap) SetSiteName(v string) {
	o.SiteName = &v
}

// GetTargetDn returns the TargetDn field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMap) GetTargetDn() string {
	if o == nil || o.TargetDn == nil {
		var ret string
		return ret
	}
	return *o.TargetDn
}

// GetTargetDnOk returns a tuple with the TargetDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) GetTargetDnOk() (*string, bool) {
	if o == nil || o.TargetDn == nil {
		return nil, false
	}
	return o.TargetDn, true
}

// HasTargetDn returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) HasTargetDn() bool {
	if o != nil && o.TargetDn != nil {
		return true
	}

	return false
}

// SetTargetDn gets a reference to the given string and assigns it to the TargetDn field.
func (o *NiatelemetryHttpsAclEpgContractMap) SetTargetDn(v string) {
	o.TargetDn = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMap) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMap) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryHttpsAclEpgContractMap) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryHttpsAclEpgContractMap) MarshalJSON() ([]byte, error) {
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
	if o.ContractName != nil {
		toSerialize["ContractName"] = o.ContractName
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.EpgName != nil {
		toSerialize["EpgName"] = o.EpgName
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.TargetDn != nil {
		toSerialize["TargetDn"] = o.TargetDn
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryHttpsAclEpgContractMap) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryHttpsAclEpgContractMapWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of HTTPS ACL contract for APIC.
		ContractName *string `json:"ContractName,omitempty"`
		// Dn of the HTTPS ACL contract provider for APIC.
		Dn *string `json:"Dn,omitempty"`
		// Name of EPGs of the HTTPS ACL contract for APIC.
		EpgName *string `json:"EpgName,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName *string `json:"SiteName,omitempty"`
		// TDn of the HTTPS ACL contract provider for APIC.
		TargetDn *string `json:"TargetDn,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryHttpsAclEpgContractMapWithoutEmbeddedStruct := NiatelemetryHttpsAclEpgContractMapWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryHttpsAclEpgContractMapWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryHttpsAclEpgContractMap := _NiatelemetryHttpsAclEpgContractMap{}
		varNiatelemetryHttpsAclEpgContractMap.ClassId = varNiatelemetryHttpsAclEpgContractMapWithoutEmbeddedStruct.ClassId
		varNiatelemetryHttpsAclEpgContractMap.ObjectType = varNiatelemetryHttpsAclEpgContractMapWithoutEmbeddedStruct.ObjectType
		varNiatelemetryHttpsAclEpgContractMap.ContractName = varNiatelemetryHttpsAclEpgContractMapWithoutEmbeddedStruct.ContractName
		varNiatelemetryHttpsAclEpgContractMap.Dn = varNiatelemetryHttpsAclEpgContractMapWithoutEmbeddedStruct.Dn
		varNiatelemetryHttpsAclEpgContractMap.EpgName = varNiatelemetryHttpsAclEpgContractMapWithoutEmbeddedStruct.EpgName
		varNiatelemetryHttpsAclEpgContractMap.RecordType = varNiatelemetryHttpsAclEpgContractMapWithoutEmbeddedStruct.RecordType
		varNiatelemetryHttpsAclEpgContractMap.RecordVersion = varNiatelemetryHttpsAclEpgContractMapWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryHttpsAclEpgContractMap.SiteName = varNiatelemetryHttpsAclEpgContractMapWithoutEmbeddedStruct.SiteName
		varNiatelemetryHttpsAclEpgContractMap.TargetDn = varNiatelemetryHttpsAclEpgContractMapWithoutEmbeddedStruct.TargetDn
		varNiatelemetryHttpsAclEpgContractMap.RegisteredDevice = varNiatelemetryHttpsAclEpgContractMapWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryHttpsAclEpgContractMap(varNiatelemetryHttpsAclEpgContractMap)
	} else {
		return err
	}

	varNiatelemetryHttpsAclEpgContractMap := _NiatelemetryHttpsAclEpgContractMap{}

	err = json.Unmarshal(bytes, &varNiatelemetryHttpsAclEpgContractMap)
	if err == nil {
		o.MoBaseMo = varNiatelemetryHttpsAclEpgContractMap.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ContractName")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "EpgName")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "TargetDn")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableNiatelemetryHttpsAclEpgContractMap struct {
	value *NiatelemetryHttpsAclEpgContractMap
	isSet bool
}

func (v NullableNiatelemetryHttpsAclEpgContractMap) Get() *NiatelemetryHttpsAclEpgContractMap {
	return v.value
}

func (v *NullableNiatelemetryHttpsAclEpgContractMap) Set(val *NiatelemetryHttpsAclEpgContractMap) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryHttpsAclEpgContractMap) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryHttpsAclEpgContractMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryHttpsAclEpgContractMap(val *NiatelemetryHttpsAclEpgContractMap) *NullableNiatelemetryHttpsAclEpgContractMap {
	return &NullableNiatelemetryHttpsAclEpgContractMap{value: val, isSet: true}
}

func (v NullableNiatelemetryHttpsAclEpgContractMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryHttpsAclEpgContractMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


