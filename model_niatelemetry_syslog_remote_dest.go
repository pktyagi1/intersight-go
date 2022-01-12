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

// NiatelemetrySyslogRemoteDest Object to capture Syslog remote dest details.
type NiatelemetrySyslogRemoteDest struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin state of Syslog remote dest in APIC.
	AdminState *string `json:"AdminState,omitempty"`
	// Parent common policy for syslog src in APIC.
	CommonPolicy *string `json:"CommonPolicy,omitempty"`
	// Dn of the Syslog remote dest in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Host of Syslog remote dest in APIC.
	Host *string `json:"Host,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// Dn of sys log src dest grp in APIC.
	SyslogRsDestGrp *string `json:"SyslogRsDestGrp,omitempty"`
	// Dn of parent syslog src for the syslog dest grp in APIC.
	SyslogSrc *string `json:"SyslogSrc,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetrySyslogRemoteDest NiatelemetrySyslogRemoteDest

// NewNiatelemetrySyslogRemoteDest instantiates a new NiatelemetrySyslogRemoteDest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetrySyslogRemoteDest(classId string, objectType string) *NiatelemetrySyslogRemoteDest {
	this := NiatelemetrySyslogRemoteDest{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetrySyslogRemoteDestWithDefaults instantiates a new NiatelemetrySyslogRemoteDest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetrySyslogRemoteDestWithDefaults() *NiatelemetrySyslogRemoteDest {
	this := NiatelemetrySyslogRemoteDest{}
	var classId string = "niatelemetry.SyslogRemoteDest"
	this.ClassId = classId
	var objectType string = "niatelemetry.SyslogRemoteDest"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetrySyslogRemoteDest) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDest) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetrySyslogRemoteDest) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetrySyslogRemoteDest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDest) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetrySyslogRemoteDest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDest) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDest) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDest) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *NiatelemetrySyslogRemoteDest) SetAdminState(v string) {
	o.AdminState = &v
}

// GetCommonPolicy returns the CommonPolicy field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDest) GetCommonPolicy() string {
	if o == nil || o.CommonPolicy == nil {
		var ret string
		return ret
	}
	return *o.CommonPolicy
}

// GetCommonPolicyOk returns a tuple with the CommonPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDest) GetCommonPolicyOk() (*string, bool) {
	if o == nil || o.CommonPolicy == nil {
		return nil, false
	}
	return o.CommonPolicy, true
}

// HasCommonPolicy returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDest) HasCommonPolicy() bool {
	if o != nil && o.CommonPolicy != nil {
		return true
	}

	return false
}

// SetCommonPolicy gets a reference to the given string and assigns it to the CommonPolicy field.
func (o *NiatelemetrySyslogRemoteDest) SetCommonPolicy(v string) {
	o.CommonPolicy = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDest) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDest) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDest) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetrySyslogRemoteDest) SetDn(v string) {
	o.Dn = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDest) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDest) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDest) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *NiatelemetrySyslogRemoteDest) SetHost(v string) {
	o.Host = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDest) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDest) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDest) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetrySyslogRemoteDest) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDest) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDest) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDest) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetrySyslogRemoteDest) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDest) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDest) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDest) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetrySyslogRemoteDest) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSyslogRsDestGrp returns the SyslogRsDestGrp field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDest) GetSyslogRsDestGrp() string {
	if o == nil || o.SyslogRsDestGrp == nil {
		var ret string
		return ret
	}
	return *o.SyslogRsDestGrp
}

// GetSyslogRsDestGrpOk returns a tuple with the SyslogRsDestGrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDest) GetSyslogRsDestGrpOk() (*string, bool) {
	if o == nil || o.SyslogRsDestGrp == nil {
		return nil, false
	}
	return o.SyslogRsDestGrp, true
}

// HasSyslogRsDestGrp returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDest) HasSyslogRsDestGrp() bool {
	if o != nil && o.SyslogRsDestGrp != nil {
		return true
	}

	return false
}

// SetSyslogRsDestGrp gets a reference to the given string and assigns it to the SyslogRsDestGrp field.
func (o *NiatelemetrySyslogRemoteDest) SetSyslogRsDestGrp(v string) {
	o.SyslogRsDestGrp = &v
}

// GetSyslogSrc returns the SyslogSrc field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDest) GetSyslogSrc() string {
	if o == nil || o.SyslogSrc == nil {
		var ret string
		return ret
	}
	return *o.SyslogSrc
}

// GetSyslogSrcOk returns a tuple with the SyslogSrc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDest) GetSyslogSrcOk() (*string, bool) {
	if o == nil || o.SyslogSrc == nil {
		return nil, false
	}
	return o.SyslogSrc, true
}

// HasSyslogSrc returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDest) HasSyslogSrc() bool {
	if o != nil && o.SyslogSrc != nil {
		return true
	}

	return false
}

// SetSyslogSrc gets a reference to the given string and assigns it to the SyslogSrc field.
func (o *NiatelemetrySyslogRemoteDest) SetSyslogSrc(v string) {
	o.SyslogSrc = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDest) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDest) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDest) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetrySyslogRemoteDest) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetrySyslogRemoteDest) MarshalJSON() ([]byte, error) {
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
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.CommonPolicy != nil {
		toSerialize["CommonPolicy"] = o.CommonPolicy
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
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
	if o.SyslogRsDestGrp != nil {
		toSerialize["SyslogRsDestGrp"] = o.SyslogRsDestGrp
	}
	if o.SyslogSrc != nil {
		toSerialize["SyslogSrc"] = o.SyslogSrc
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetrySyslogRemoteDest) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetrySyslogRemoteDestWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin state of Syslog remote dest in APIC.
		AdminState *string `json:"AdminState,omitempty"`
		// Parent common policy for syslog src in APIC.
		CommonPolicy *string `json:"CommonPolicy,omitempty"`
		// Dn of the Syslog remote dest in APIC.
		Dn *string `json:"Dn,omitempty"`
		// Host of Syslog remote dest in APIC.
		Host *string `json:"Host,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName *string `json:"SiteName,omitempty"`
		// Dn of sys log src dest grp in APIC.
		SyslogRsDestGrp *string `json:"SyslogRsDestGrp,omitempty"`
		// Dn of parent syslog src for the syslog dest grp in APIC.
		SyslogSrc *string `json:"SyslogSrc,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetrySyslogRemoteDestWithoutEmbeddedStruct := NiatelemetrySyslogRemoteDestWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetrySyslogRemoteDestWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetrySyslogRemoteDest := _NiatelemetrySyslogRemoteDest{}
		varNiatelemetrySyslogRemoteDest.ClassId = varNiatelemetrySyslogRemoteDestWithoutEmbeddedStruct.ClassId
		varNiatelemetrySyslogRemoteDest.ObjectType = varNiatelemetrySyslogRemoteDestWithoutEmbeddedStruct.ObjectType
		varNiatelemetrySyslogRemoteDest.AdminState = varNiatelemetrySyslogRemoteDestWithoutEmbeddedStruct.AdminState
		varNiatelemetrySyslogRemoteDest.CommonPolicy = varNiatelemetrySyslogRemoteDestWithoutEmbeddedStruct.CommonPolicy
		varNiatelemetrySyslogRemoteDest.Dn = varNiatelemetrySyslogRemoteDestWithoutEmbeddedStruct.Dn
		varNiatelemetrySyslogRemoteDest.Host = varNiatelemetrySyslogRemoteDestWithoutEmbeddedStruct.Host
		varNiatelemetrySyslogRemoteDest.RecordType = varNiatelemetrySyslogRemoteDestWithoutEmbeddedStruct.RecordType
		varNiatelemetrySyslogRemoteDest.RecordVersion = varNiatelemetrySyslogRemoteDestWithoutEmbeddedStruct.RecordVersion
		varNiatelemetrySyslogRemoteDest.SiteName = varNiatelemetrySyslogRemoteDestWithoutEmbeddedStruct.SiteName
		varNiatelemetrySyslogRemoteDest.SyslogRsDestGrp = varNiatelemetrySyslogRemoteDestWithoutEmbeddedStruct.SyslogRsDestGrp
		varNiatelemetrySyslogRemoteDest.SyslogSrc = varNiatelemetrySyslogRemoteDestWithoutEmbeddedStruct.SyslogSrc
		varNiatelemetrySyslogRemoteDest.RegisteredDevice = varNiatelemetrySyslogRemoteDestWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetrySyslogRemoteDest(varNiatelemetrySyslogRemoteDest)
	} else {
		return err
	}

	varNiatelemetrySyslogRemoteDest := _NiatelemetrySyslogRemoteDest{}

	err = json.Unmarshal(bytes, &varNiatelemetrySyslogRemoteDest)
	if err == nil {
		o.MoBaseMo = varNiatelemetrySyslogRemoteDest.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "CommonPolicy")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "Host")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SyslogRsDestGrp")
		delete(additionalProperties, "SyslogSrc")
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

type NullableNiatelemetrySyslogRemoteDest struct {
	value *NiatelemetrySyslogRemoteDest
	isSet bool
}

func (v NullableNiatelemetrySyslogRemoteDest) Get() *NiatelemetrySyslogRemoteDest {
	return v.value
}

func (v *NullableNiatelemetrySyslogRemoteDest) Set(val *NiatelemetrySyslogRemoteDest) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetrySyslogRemoteDest) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetrySyslogRemoteDest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetrySyslogRemoteDest(val *NiatelemetrySyslogRemoteDest) *NullableNiatelemetrySyslogRemoteDest {
	return &NullableNiatelemetrySyslogRemoteDest{value: val, isSet: true}
}

func (v NullableNiatelemetrySyslogRemoteDest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetrySyslogRemoteDest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


