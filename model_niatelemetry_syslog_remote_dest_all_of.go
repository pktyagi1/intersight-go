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
)

// NiatelemetrySyslogRemoteDestAllOf Definition of the list of properties defined in 'niatelemetry.SyslogRemoteDest', excluding properties defined in parent classes.
type NiatelemetrySyslogRemoteDestAllOf struct {
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

type _NiatelemetrySyslogRemoteDestAllOf NiatelemetrySyslogRemoteDestAllOf

// NewNiatelemetrySyslogRemoteDestAllOf instantiates a new NiatelemetrySyslogRemoteDestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetrySyslogRemoteDestAllOf(classId string, objectType string) *NiatelemetrySyslogRemoteDestAllOf {
	this := NiatelemetrySyslogRemoteDestAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetrySyslogRemoteDestAllOfWithDefaults instantiates a new NiatelemetrySyslogRemoteDestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetrySyslogRemoteDestAllOfWithDefaults() *NiatelemetrySyslogRemoteDestAllOf {
	this := NiatelemetrySyslogRemoteDestAllOf{}
	var classId string = "niatelemetry.SyslogRemoteDest"
	this.ClassId = classId
	var objectType string = "niatelemetry.SyslogRemoteDest"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetrySyslogRemoteDestAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetrySyslogRemoteDestAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetrySyslogRemoteDestAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetrySyslogRemoteDestAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *NiatelemetrySyslogRemoteDestAllOf) SetAdminState(v string) {
	o.AdminState = &v
}

// GetCommonPolicy returns the CommonPolicy field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetCommonPolicy() string {
	if o == nil || o.CommonPolicy == nil {
		var ret string
		return ret
	}
	return *o.CommonPolicy
}

// GetCommonPolicyOk returns a tuple with the CommonPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetCommonPolicyOk() (*string, bool) {
	if o == nil || o.CommonPolicy == nil {
		return nil, false
	}
	return o.CommonPolicy, true
}

// HasCommonPolicy returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) HasCommonPolicy() bool {
	if o != nil && o.CommonPolicy != nil {
		return true
	}

	return false
}

// SetCommonPolicy gets a reference to the given string and assigns it to the CommonPolicy field.
func (o *NiatelemetrySyslogRemoteDestAllOf) SetCommonPolicy(v string) {
	o.CommonPolicy = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetrySyslogRemoteDestAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *NiatelemetrySyslogRemoteDestAllOf) SetHost(v string) {
	o.Host = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetrySyslogRemoteDestAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetrySyslogRemoteDestAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetrySyslogRemoteDestAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSyslogRsDestGrp returns the SyslogRsDestGrp field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetSyslogRsDestGrp() string {
	if o == nil || o.SyslogRsDestGrp == nil {
		var ret string
		return ret
	}
	return *o.SyslogRsDestGrp
}

// GetSyslogRsDestGrpOk returns a tuple with the SyslogRsDestGrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetSyslogRsDestGrpOk() (*string, bool) {
	if o == nil || o.SyslogRsDestGrp == nil {
		return nil, false
	}
	return o.SyslogRsDestGrp, true
}

// HasSyslogRsDestGrp returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) HasSyslogRsDestGrp() bool {
	if o != nil && o.SyslogRsDestGrp != nil {
		return true
	}

	return false
}

// SetSyslogRsDestGrp gets a reference to the given string and assigns it to the SyslogRsDestGrp field.
func (o *NiatelemetrySyslogRemoteDestAllOf) SetSyslogRsDestGrp(v string) {
	o.SyslogRsDestGrp = &v
}

// GetSyslogSrc returns the SyslogSrc field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetSyslogSrc() string {
	if o == nil || o.SyslogSrc == nil {
		var ret string
		return ret
	}
	return *o.SyslogSrc
}

// GetSyslogSrcOk returns a tuple with the SyslogSrc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetSyslogSrcOk() (*string, bool) {
	if o == nil || o.SyslogSrc == nil {
		return nil, false
	}
	return o.SyslogSrc, true
}

// HasSyslogSrc returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) HasSyslogSrc() bool {
	if o != nil && o.SyslogSrc != nil {
		return true
	}

	return false
}

// SetSyslogSrc gets a reference to the given string and assigns it to the SyslogSrc field.
func (o *NiatelemetrySyslogRemoteDestAllOf) SetSyslogSrc(v string) {
	o.SyslogSrc = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetrySyslogRemoteDestAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetrySyslogRemoteDestAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetrySyslogRemoteDestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *NiatelemetrySyslogRemoteDestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetrySyslogRemoteDestAllOf := _NiatelemetrySyslogRemoteDestAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetrySyslogRemoteDestAllOf); err == nil {
		*o = NiatelemetrySyslogRemoteDestAllOf(varNiatelemetrySyslogRemoteDestAllOf)
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetrySyslogRemoteDestAllOf struct {
	value *NiatelemetrySyslogRemoteDestAllOf
	isSet bool
}

func (v NullableNiatelemetrySyslogRemoteDestAllOf) Get() *NiatelemetrySyslogRemoteDestAllOf {
	return v.value
}

func (v *NullableNiatelemetrySyslogRemoteDestAllOf) Set(val *NiatelemetrySyslogRemoteDestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetrySyslogRemoteDestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetrySyslogRemoteDestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetrySyslogRemoteDestAllOf(val *NiatelemetrySyslogRemoteDestAllOf) *NullableNiatelemetrySyslogRemoteDestAllOf {
	return &NullableNiatelemetrySyslogRemoteDestAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetrySyslogRemoteDestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetrySyslogRemoteDestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


