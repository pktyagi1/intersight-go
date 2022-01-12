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

// NiatelemetryHttpsAclEpgDetailsAllOf Definition of the list of properties defined in 'niatelemetry.HttpsAclEpgDetails', excluding properties defined in parent classes.
type NiatelemetryHttpsAclEpgDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of the HTTPS ACL EPG for APIC.
	Dn *string `json:"Dn,omitempty"`
	// Name of EPGs of the HTTPS ACL EPG for APIC.
	EpgName *string `json:"EpgName,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryHttpsAclEpgDetailsAllOf NiatelemetryHttpsAclEpgDetailsAllOf

// NewNiatelemetryHttpsAclEpgDetailsAllOf instantiates a new NiatelemetryHttpsAclEpgDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryHttpsAclEpgDetailsAllOf(classId string, objectType string) *NiatelemetryHttpsAclEpgDetailsAllOf {
	this := NiatelemetryHttpsAclEpgDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryHttpsAclEpgDetailsAllOfWithDefaults instantiates a new NiatelemetryHttpsAclEpgDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryHttpsAclEpgDetailsAllOfWithDefaults() *NiatelemetryHttpsAclEpgDetailsAllOf {
	this := NiatelemetryHttpsAclEpgDetailsAllOf{}
	var classId string = "niatelemetry.HttpsAclEpgDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.HttpsAclEpgDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetEpgName returns the EpgName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetEpgName() string {
	if o == nil || o.EpgName == nil {
		var ret string
		return ret
	}
	return *o.EpgName
}

// GetEpgNameOk returns a tuple with the EpgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetEpgNameOk() (*string, bool) {
	if o == nil || o.EpgName == nil {
		return nil, false
	}
	return o.EpgName, true
}

// HasEpgName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) HasEpgName() bool {
	if o != nil && o.EpgName != nil {
		return true
	}

	return false
}

// SetEpgName gets a reference to the given string and assigns it to the EpgName field.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) SetEpgName(v string) {
	o.EpgName = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryHttpsAclEpgDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryHttpsAclEpgDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
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
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryHttpsAclEpgDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryHttpsAclEpgDetailsAllOf := _NiatelemetryHttpsAclEpgDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryHttpsAclEpgDetailsAllOf); err == nil {
		*o = NiatelemetryHttpsAclEpgDetailsAllOf(varNiatelemetryHttpsAclEpgDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "EpgName")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryHttpsAclEpgDetailsAllOf struct {
	value *NiatelemetryHttpsAclEpgDetailsAllOf
	isSet bool
}

func (v NullableNiatelemetryHttpsAclEpgDetailsAllOf) Get() *NiatelemetryHttpsAclEpgDetailsAllOf {
	return v.value
}

func (v *NullableNiatelemetryHttpsAclEpgDetailsAllOf) Set(val *NiatelemetryHttpsAclEpgDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryHttpsAclEpgDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryHttpsAclEpgDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryHttpsAclEpgDetailsAllOf(val *NiatelemetryHttpsAclEpgDetailsAllOf) *NullableNiatelemetryHttpsAclEpgDetailsAllOf {
	return &NullableNiatelemetryHttpsAclEpgDetailsAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryHttpsAclEpgDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryHttpsAclEpgDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


