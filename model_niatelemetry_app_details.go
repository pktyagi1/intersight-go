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

// NiatelemetryAppDetails Details of apps installed on Nexus Dashboard.
type NiatelemetryAppDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Names of apps running on ND.
	AppName *string `json:"AppName,omitempty"`
	// Status of apps running on ND.
	AppStatus *string `json:"AppStatus,omitempty"`
	// Versions of apps running on ND.
	AppVersion *string `json:"AppVersion,omitempty"`
	// Clustername on which apps are running on ND.
	NexusDashboard *string `json:"NexusDashboard,omitempty"`
	// Number of sites on which particular app installed on ND.
	NumberOfSitesOnboarded *int64 `json:"NumberOfSitesOnboarded,omitempty"`
	// Type of apps running on ND.
	Type *string `json:"Type,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryAppDetails NiatelemetryAppDetails

// NewNiatelemetryAppDetails instantiates a new NiatelemetryAppDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryAppDetails(classId string, objectType string) *NiatelemetryAppDetails {
	this := NiatelemetryAppDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryAppDetailsWithDefaults instantiates a new NiatelemetryAppDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryAppDetailsWithDefaults() *NiatelemetryAppDetails {
	this := NiatelemetryAppDetails{}
	var classId string = "niatelemetry.AppDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.AppDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryAppDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryAppDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryAppDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryAppDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *NiatelemetryAppDetails) GetAppName() string {
	if o == nil || o.AppName == nil {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetAppNameOk() (*string, bool) {
	if o == nil || o.AppName == nil {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *NiatelemetryAppDetails) HasAppName() bool {
	if o != nil && o.AppName != nil {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *NiatelemetryAppDetails) SetAppName(v string) {
	o.AppName = &v
}

// GetAppStatus returns the AppStatus field value if set, zero value otherwise.
func (o *NiatelemetryAppDetails) GetAppStatus() string {
	if o == nil || o.AppStatus == nil {
		var ret string
		return ret
	}
	return *o.AppStatus
}

// GetAppStatusOk returns a tuple with the AppStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetAppStatusOk() (*string, bool) {
	if o == nil || o.AppStatus == nil {
		return nil, false
	}
	return o.AppStatus, true
}

// HasAppStatus returns a boolean if a field has been set.
func (o *NiatelemetryAppDetails) HasAppStatus() bool {
	if o != nil && o.AppStatus != nil {
		return true
	}

	return false
}

// SetAppStatus gets a reference to the given string and assigns it to the AppStatus field.
func (o *NiatelemetryAppDetails) SetAppStatus(v string) {
	o.AppStatus = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *NiatelemetryAppDetails) GetAppVersion() string {
	if o == nil || o.AppVersion == nil {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetAppVersionOk() (*string, bool) {
	if o == nil || o.AppVersion == nil {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *NiatelemetryAppDetails) HasAppVersion() bool {
	if o != nil && o.AppVersion != nil {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *NiatelemetryAppDetails) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetNexusDashboard returns the NexusDashboard field value if set, zero value otherwise.
func (o *NiatelemetryAppDetails) GetNexusDashboard() string {
	if o == nil || o.NexusDashboard == nil {
		var ret string
		return ret
	}
	return *o.NexusDashboard
}

// GetNexusDashboardOk returns a tuple with the NexusDashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetNexusDashboardOk() (*string, bool) {
	if o == nil || o.NexusDashboard == nil {
		return nil, false
	}
	return o.NexusDashboard, true
}

// HasNexusDashboard returns a boolean if a field has been set.
func (o *NiatelemetryAppDetails) HasNexusDashboard() bool {
	if o != nil && o.NexusDashboard != nil {
		return true
	}

	return false
}

// SetNexusDashboard gets a reference to the given string and assigns it to the NexusDashboard field.
func (o *NiatelemetryAppDetails) SetNexusDashboard(v string) {
	o.NexusDashboard = &v
}

// GetNumberOfSitesOnboarded returns the NumberOfSitesOnboarded field value if set, zero value otherwise.
func (o *NiatelemetryAppDetails) GetNumberOfSitesOnboarded() int64 {
	if o == nil || o.NumberOfSitesOnboarded == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfSitesOnboarded
}

// GetNumberOfSitesOnboardedOk returns a tuple with the NumberOfSitesOnboarded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetNumberOfSitesOnboardedOk() (*int64, bool) {
	if o == nil || o.NumberOfSitesOnboarded == nil {
		return nil, false
	}
	return o.NumberOfSitesOnboarded, true
}

// HasNumberOfSitesOnboarded returns a boolean if a field has been set.
func (o *NiatelemetryAppDetails) HasNumberOfSitesOnboarded() bool {
	if o != nil && o.NumberOfSitesOnboarded != nil {
		return true
	}

	return false
}

// SetNumberOfSitesOnboarded gets a reference to the given int64 and assigns it to the NumberOfSitesOnboarded field.
func (o *NiatelemetryAppDetails) SetNumberOfSitesOnboarded(v int64) {
	o.NumberOfSitesOnboarded = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NiatelemetryAppDetails) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NiatelemetryAppDetails) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NiatelemetryAppDetails) SetType(v string) {
	o.Type = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryAppDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryAppDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryAppDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryAppDetails) MarshalJSON() ([]byte, error) {
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
	if o.AppName != nil {
		toSerialize["AppName"] = o.AppName
	}
	if o.AppStatus != nil {
		toSerialize["AppStatus"] = o.AppStatus
	}
	if o.AppVersion != nil {
		toSerialize["AppVersion"] = o.AppVersion
	}
	if o.NexusDashboard != nil {
		toSerialize["NexusDashboard"] = o.NexusDashboard
	}
	if o.NumberOfSitesOnboarded != nil {
		toSerialize["NumberOfSitesOnboarded"] = o.NumberOfSitesOnboarded
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryAppDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryAppDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Names of apps running on ND.
		AppName *string `json:"AppName,omitempty"`
		// Status of apps running on ND.
		AppStatus *string `json:"AppStatus,omitempty"`
		// Versions of apps running on ND.
		AppVersion *string `json:"AppVersion,omitempty"`
		// Clustername on which apps are running on ND.
		NexusDashboard *string `json:"NexusDashboard,omitempty"`
		// Number of sites on which particular app installed on ND.
		NumberOfSitesOnboarded *int64 `json:"NumberOfSitesOnboarded,omitempty"`
		// Type of apps running on ND.
		Type *string `json:"Type,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryAppDetailsWithoutEmbeddedStruct := NiatelemetryAppDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryAppDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryAppDetails := _NiatelemetryAppDetails{}
		varNiatelemetryAppDetails.ClassId = varNiatelemetryAppDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryAppDetails.ObjectType = varNiatelemetryAppDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryAppDetails.AppName = varNiatelemetryAppDetailsWithoutEmbeddedStruct.AppName
		varNiatelemetryAppDetails.AppStatus = varNiatelemetryAppDetailsWithoutEmbeddedStruct.AppStatus
		varNiatelemetryAppDetails.AppVersion = varNiatelemetryAppDetailsWithoutEmbeddedStruct.AppVersion
		varNiatelemetryAppDetails.NexusDashboard = varNiatelemetryAppDetailsWithoutEmbeddedStruct.NexusDashboard
		varNiatelemetryAppDetails.NumberOfSitesOnboarded = varNiatelemetryAppDetailsWithoutEmbeddedStruct.NumberOfSitesOnboarded
		varNiatelemetryAppDetails.Type = varNiatelemetryAppDetailsWithoutEmbeddedStruct.Type
		varNiatelemetryAppDetails.RegisteredDevice = varNiatelemetryAppDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryAppDetails(varNiatelemetryAppDetails)
	} else {
		return err
	}

	varNiatelemetryAppDetails := _NiatelemetryAppDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryAppDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryAppDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AppName")
		delete(additionalProperties, "AppStatus")
		delete(additionalProperties, "AppVersion")
		delete(additionalProperties, "NexusDashboard")
		delete(additionalProperties, "NumberOfSitesOnboarded")
		delete(additionalProperties, "Type")
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

type NullableNiatelemetryAppDetails struct {
	value *NiatelemetryAppDetails
	isSet bool
}

func (v NullableNiatelemetryAppDetails) Get() *NiatelemetryAppDetails {
	return v.value
}

func (v *NullableNiatelemetryAppDetails) Set(val *NiatelemetryAppDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryAppDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryAppDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryAppDetails(val *NiatelemetryAppDetails) *NullableNiatelemetryAppDetails {
	return &NullableNiatelemetryAppDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryAppDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryAppDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


