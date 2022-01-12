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

// NiatelemetryNexusDashboardControllerDetailsAllOf Definition of the list of properties defined in 'niatelemetry.NexusDashboardControllerDetails', excluding properties defined in parent classes.
type NiatelemetryNexusDashboardControllerDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Health of the site serviced by ND.
	SiteHealth *int64 `json:"SiteHealth,omitempty"`
	// Name of fabric domain of the controller.
	SiteName *string `json:"SiteName,omitempty"`
	// Version of the controller serviced by ND.
	VersionOfController *string `json:"VersionOfController,omitempty"`
	NexusDashboard *NiatelemetryNexusDashboardsRelationship `json:"NexusDashboard,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNexusDashboardControllerDetailsAllOf NiatelemetryNexusDashboardControllerDetailsAllOf

// NewNiatelemetryNexusDashboardControllerDetailsAllOf instantiates a new NiatelemetryNexusDashboardControllerDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNexusDashboardControllerDetailsAllOf(classId string, objectType string) *NiatelemetryNexusDashboardControllerDetailsAllOf {
	this := NiatelemetryNexusDashboardControllerDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNexusDashboardControllerDetailsAllOfWithDefaults instantiates a new NiatelemetryNexusDashboardControllerDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNexusDashboardControllerDetailsAllOfWithDefaults() *NiatelemetryNexusDashboardControllerDetailsAllOf {
	this := NiatelemetryNexusDashboardControllerDetailsAllOf{}
	var classId string = "niatelemetry.NexusDashboardControllerDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.NexusDashboardControllerDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSiteHealth returns the SiteHealth field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetSiteHealth() int64 {
	if o == nil || o.SiteHealth == nil {
		var ret int64
		return ret
	}
	return *o.SiteHealth
}

// GetSiteHealthOk returns a tuple with the SiteHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetSiteHealthOk() (*int64, bool) {
	if o == nil || o.SiteHealth == nil {
		return nil, false
	}
	return o.SiteHealth, true
}

// HasSiteHealth returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) HasSiteHealth() bool {
	if o != nil && o.SiteHealth != nil {
		return true
	}

	return false
}

// SetSiteHealth gets a reference to the given int64 and assigns it to the SiteHealth field.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) SetSiteHealth(v int64) {
	o.SiteHealth = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetVersionOfController returns the VersionOfController field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetVersionOfController() string {
	if o == nil || o.VersionOfController == nil {
		var ret string
		return ret
	}
	return *o.VersionOfController
}

// GetVersionOfControllerOk returns a tuple with the VersionOfController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetVersionOfControllerOk() (*string, bool) {
	if o == nil || o.VersionOfController == nil {
		return nil, false
	}
	return o.VersionOfController, true
}

// HasVersionOfController returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) HasVersionOfController() bool {
	if o != nil && o.VersionOfController != nil {
		return true
	}

	return false
}

// SetVersionOfController gets a reference to the given string and assigns it to the VersionOfController field.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) SetVersionOfController(v string) {
	o.VersionOfController = &v
}

// GetNexusDashboard returns the NexusDashboard field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetNexusDashboard() NiatelemetryNexusDashboardsRelationship {
	if o == nil || o.NexusDashboard == nil {
		var ret NiatelemetryNexusDashboardsRelationship
		return ret
	}
	return *o.NexusDashboard
}

// GetNexusDashboardOk returns a tuple with the NexusDashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetNexusDashboardOk() (*NiatelemetryNexusDashboardsRelationship, bool) {
	if o == nil || o.NexusDashboard == nil {
		return nil, false
	}
	return o.NexusDashboard, true
}

// HasNexusDashboard returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) HasNexusDashboard() bool {
	if o != nil && o.NexusDashboard != nil {
		return true
	}

	return false
}

// SetNexusDashboard gets a reference to the given NiatelemetryNexusDashboardsRelationship and assigns it to the NexusDashboard field.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) SetNexusDashboard(v NiatelemetryNexusDashboardsRelationship) {
	o.NexusDashboard = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryNexusDashboardControllerDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.SiteHealth != nil {
		toSerialize["SiteHealth"] = o.SiteHealth
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.VersionOfController != nil {
		toSerialize["VersionOfController"] = o.VersionOfController
	}
	if o.NexusDashboard != nil {
		toSerialize["NexusDashboard"] = o.NexusDashboard
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNexusDashboardControllerDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryNexusDashboardControllerDetailsAllOf := _NiatelemetryNexusDashboardControllerDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryNexusDashboardControllerDetailsAllOf); err == nil {
		*o = NiatelemetryNexusDashboardControllerDetailsAllOf(varNiatelemetryNexusDashboardControllerDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SiteHealth")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "VersionOfController")
		delete(additionalProperties, "NexusDashboard")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryNexusDashboardControllerDetailsAllOf struct {
	value *NiatelemetryNexusDashboardControllerDetailsAllOf
	isSet bool
}

func (v NullableNiatelemetryNexusDashboardControllerDetailsAllOf) Get() *NiatelemetryNexusDashboardControllerDetailsAllOf {
	return v.value
}

func (v *NullableNiatelemetryNexusDashboardControllerDetailsAllOf) Set(val *NiatelemetryNexusDashboardControllerDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNexusDashboardControllerDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNexusDashboardControllerDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNexusDashboardControllerDetailsAllOf(val *NiatelemetryNexusDashboardControllerDetailsAllOf) *NullableNiatelemetryNexusDashboardControllerDetailsAllOf {
	return &NullableNiatelemetryNexusDashboardControllerDetailsAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNexusDashboardControllerDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNexusDashboardControllerDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


