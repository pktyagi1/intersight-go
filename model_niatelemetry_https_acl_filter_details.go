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
	"reflect"
	"strings"
)

// NiatelemetryHttpsAclFilterDetails Object to capture the HTTPS ACL EPGs filter details in APIC.
type NiatelemetryHttpsAclFilterDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Destination From Port HTTPS ACL EPGs filter MO for APIC.
	DestFromPort *string `json:"DestFromPort,omitempty"`
	// Destination To Port HTTPS ACL EPGs filter MO for APIC.
	DestToPort *string `json:"DestToPort,omitempty"`
	// Dn of the HTTPS ACL EPGs filter MO for APIC.
	Dn *string `json:"Dn,omitempty"`
	// Name of HTTPS ACL filter for APIC.
	FilterName *string `json:"FilterName,omitempty"`
	// Prot of the HTTPS ACL EPGs filter MO for APIC.
	Prot *string `json:"Prot,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// Source From Port HTTPS ACL EPGs filter MO for APIC.
	SrcFromPort *string `json:"SrcFromPort,omitempty"`
	// Source To Port HTTPS ACL EPGs filter MO for APIC.
	SrcToPort *string `json:"SrcToPort,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryHttpsAclFilterDetails NiatelemetryHttpsAclFilterDetails

// NewNiatelemetryHttpsAclFilterDetails instantiates a new NiatelemetryHttpsAclFilterDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryHttpsAclFilterDetails(classId string, objectType string) *NiatelemetryHttpsAclFilterDetails {
	this := NiatelemetryHttpsAclFilterDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryHttpsAclFilterDetailsWithDefaults instantiates a new NiatelemetryHttpsAclFilterDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryHttpsAclFilterDetailsWithDefaults() *NiatelemetryHttpsAclFilterDetails {
	this := NiatelemetryHttpsAclFilterDetails{}
	var classId string = "niatelemetry.HttpsAclFilterDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.HttpsAclFilterDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryHttpsAclFilterDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetails) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryHttpsAclFilterDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryHttpsAclFilterDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryHttpsAclFilterDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDestFromPort returns the DestFromPort field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetails) GetDestFromPort() string {
	if o == nil || o.DestFromPort == nil {
		var ret string
		return ret
	}
	return *o.DestFromPort
}

// GetDestFromPortOk returns a tuple with the DestFromPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetails) GetDestFromPortOk() (*string, bool) {
	if o == nil || o.DestFromPort == nil {
		return nil, false
	}
	return o.DestFromPort, true
}

// HasDestFromPort returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetails) HasDestFromPort() bool {
	if o != nil && o.DestFromPort != nil {
		return true
	}

	return false
}

// SetDestFromPort gets a reference to the given string and assigns it to the DestFromPort field.
func (o *NiatelemetryHttpsAclFilterDetails) SetDestFromPort(v string) {
	o.DestFromPort = &v
}

// GetDestToPort returns the DestToPort field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetails) GetDestToPort() string {
	if o == nil || o.DestToPort == nil {
		var ret string
		return ret
	}
	return *o.DestToPort
}

// GetDestToPortOk returns a tuple with the DestToPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetails) GetDestToPortOk() (*string, bool) {
	if o == nil || o.DestToPort == nil {
		return nil, false
	}
	return o.DestToPort, true
}

// HasDestToPort returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetails) HasDestToPort() bool {
	if o != nil && o.DestToPort != nil {
		return true
	}

	return false
}

// SetDestToPort gets a reference to the given string and assigns it to the DestToPort field.
func (o *NiatelemetryHttpsAclFilterDetails) SetDestToPort(v string) {
	o.DestToPort = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetails) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetails) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetails) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryHttpsAclFilterDetails) SetDn(v string) {
	o.Dn = &v
}

// GetFilterName returns the FilterName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetails) GetFilterName() string {
	if o == nil || o.FilterName == nil {
		var ret string
		return ret
	}
	return *o.FilterName
}

// GetFilterNameOk returns a tuple with the FilterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetails) GetFilterNameOk() (*string, bool) {
	if o == nil || o.FilterName == nil {
		return nil, false
	}
	return o.FilterName, true
}

// HasFilterName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetails) HasFilterName() bool {
	if o != nil && o.FilterName != nil {
		return true
	}

	return false
}

// SetFilterName gets a reference to the given string and assigns it to the FilterName field.
func (o *NiatelemetryHttpsAclFilterDetails) SetFilterName(v string) {
	o.FilterName = &v
}

// GetProt returns the Prot field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetails) GetProt() string {
	if o == nil || o.Prot == nil {
		var ret string
		return ret
	}
	return *o.Prot
}

// GetProtOk returns a tuple with the Prot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetails) GetProtOk() (*string, bool) {
	if o == nil || o.Prot == nil {
		return nil, false
	}
	return o.Prot, true
}

// HasProt returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetails) HasProt() bool {
	if o != nil && o.Prot != nil {
		return true
	}

	return false
}

// SetProt gets a reference to the given string and assigns it to the Prot field.
func (o *NiatelemetryHttpsAclFilterDetails) SetProt(v string) {
	o.Prot = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetails) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetails) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryHttpsAclFilterDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetails) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetails) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryHttpsAclFilterDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetails) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetails) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryHttpsAclFilterDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSrcFromPort returns the SrcFromPort field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetails) GetSrcFromPort() string {
	if o == nil || o.SrcFromPort == nil {
		var ret string
		return ret
	}
	return *o.SrcFromPort
}

// GetSrcFromPortOk returns a tuple with the SrcFromPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetails) GetSrcFromPortOk() (*string, bool) {
	if o == nil || o.SrcFromPort == nil {
		return nil, false
	}
	return o.SrcFromPort, true
}

// HasSrcFromPort returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetails) HasSrcFromPort() bool {
	if o != nil && o.SrcFromPort != nil {
		return true
	}

	return false
}

// SetSrcFromPort gets a reference to the given string and assigns it to the SrcFromPort field.
func (o *NiatelemetryHttpsAclFilterDetails) SetSrcFromPort(v string) {
	o.SrcFromPort = &v
}

// GetSrcToPort returns the SrcToPort field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetails) GetSrcToPort() string {
	if o == nil || o.SrcToPort == nil {
		var ret string
		return ret
	}
	return *o.SrcToPort
}

// GetSrcToPortOk returns a tuple with the SrcToPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetails) GetSrcToPortOk() (*string, bool) {
	if o == nil || o.SrcToPort == nil {
		return nil, false
	}
	return o.SrcToPort, true
}

// HasSrcToPort returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetails) HasSrcToPort() bool {
	if o != nil && o.SrcToPort != nil {
		return true
	}

	return false
}

// SetSrcToPort gets a reference to the given string and assigns it to the SrcToPort field.
func (o *NiatelemetryHttpsAclFilterDetails) SetSrcToPort(v string) {
	o.SrcToPort = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryHttpsAclFilterDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryHttpsAclFilterDetails) MarshalJSON() ([]byte, error) {
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
	if o.DestFromPort != nil {
		toSerialize["DestFromPort"] = o.DestFromPort
	}
	if o.DestToPort != nil {
		toSerialize["DestToPort"] = o.DestToPort
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.FilterName != nil {
		toSerialize["FilterName"] = o.FilterName
	}
	if o.Prot != nil {
		toSerialize["Prot"] = o.Prot
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
	if o.SrcFromPort != nil {
		toSerialize["SrcFromPort"] = o.SrcFromPort
	}
	if o.SrcToPort != nil {
		toSerialize["SrcToPort"] = o.SrcToPort
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryHttpsAclFilterDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Destination From Port HTTPS ACL EPGs filter MO for APIC.
		DestFromPort *string `json:"DestFromPort,omitempty"`
		// Destination To Port HTTPS ACL EPGs filter MO for APIC.
		DestToPort *string `json:"DestToPort,omitempty"`
		// Dn of the HTTPS ACL EPGs filter MO for APIC.
		Dn *string `json:"Dn,omitempty"`
		// Name of HTTPS ACL filter for APIC.
		FilterName *string `json:"FilterName,omitempty"`
		// Prot of the HTTPS ACL EPGs filter MO for APIC.
		Prot *string `json:"Prot,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName *string `json:"SiteName,omitempty"`
		// Source From Port HTTPS ACL EPGs filter MO for APIC.
		SrcFromPort *string `json:"SrcFromPort,omitempty"`
		// Source To Port HTTPS ACL EPGs filter MO for APIC.
		SrcToPort *string `json:"SrcToPort,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct := NiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryHttpsAclFilterDetails := _NiatelemetryHttpsAclFilterDetails{}
		varNiatelemetryHttpsAclFilterDetails.ClassId = varNiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryHttpsAclFilterDetails.ObjectType = varNiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryHttpsAclFilterDetails.DestFromPort = varNiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct.DestFromPort
		varNiatelemetryHttpsAclFilterDetails.DestToPort = varNiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct.DestToPort
		varNiatelemetryHttpsAclFilterDetails.Dn = varNiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct.Dn
		varNiatelemetryHttpsAclFilterDetails.FilterName = varNiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct.FilterName
		varNiatelemetryHttpsAclFilterDetails.Prot = varNiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct.Prot
		varNiatelemetryHttpsAclFilterDetails.RecordType = varNiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryHttpsAclFilterDetails.RecordVersion = varNiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryHttpsAclFilterDetails.SiteName = varNiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryHttpsAclFilterDetails.SrcFromPort = varNiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct.SrcFromPort
		varNiatelemetryHttpsAclFilterDetails.SrcToPort = varNiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct.SrcToPort
		varNiatelemetryHttpsAclFilterDetails.RegisteredDevice = varNiatelemetryHttpsAclFilterDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryHttpsAclFilterDetails(varNiatelemetryHttpsAclFilterDetails)
	} else {
		return err
	}

	varNiatelemetryHttpsAclFilterDetails := _NiatelemetryHttpsAclFilterDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryHttpsAclFilterDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryHttpsAclFilterDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DestFromPort")
		delete(additionalProperties, "DestToPort")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "FilterName")
		delete(additionalProperties, "Prot")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SrcFromPort")
		delete(additionalProperties, "SrcToPort")
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

type NullableNiatelemetryHttpsAclFilterDetails struct {
	value *NiatelemetryHttpsAclFilterDetails
	isSet bool
}

func (v NullableNiatelemetryHttpsAclFilterDetails) Get() *NiatelemetryHttpsAclFilterDetails {
	return v.value
}

func (v *NullableNiatelemetryHttpsAclFilterDetails) Set(val *NiatelemetryHttpsAclFilterDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryHttpsAclFilterDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryHttpsAclFilterDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryHttpsAclFilterDetails(val *NiatelemetryHttpsAclFilterDetails) *NullableNiatelemetryHttpsAclFilterDetails {
	return &NullableNiatelemetryHttpsAclFilterDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryHttpsAclFilterDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryHttpsAclFilterDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


