/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5517
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// KvmSessionAllOf Definition of the list of properties defined in 'kvm.Session', excluding properties defined in parent classes.
type KvmSessionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// One time URL that is used to launch the KVM console.
	KvmLaunchUrlPath *string `json:"KvmLaunchUrlPath,omitempty"`
	// Unique ID of the KVM Session URI.
	KvmSessionId *string `json:"KvmSessionId,omitempty"`
	// Temporary one-time password for vKVM access.
	OneTimePassword *string `json:"OneTimePassword,omitempty"`
	// Indicates if vKVM SSO is supported on the server.
	SsoSupported *bool `json:"SsoSupported,omitempty"`
	// Username used for vKVM access.
	Username *string `json:"Username,omitempty"`
	Device *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	Server *ComputePhysicalRelationship `json:"Server,omitempty"`
	Tunnel *KvmTunnelRelationship `json:"Tunnel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KvmSessionAllOf KvmSessionAllOf

// NewKvmSessionAllOf instantiates a new KvmSessionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmSessionAllOf(classId string, objectType string) *KvmSessionAllOf {
	this := KvmSessionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKvmSessionAllOfWithDefaults instantiates a new KvmSessionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmSessionAllOfWithDefaults() *KvmSessionAllOf {
	this := KvmSessionAllOf{}
	var classId string = "kvm.Session"
	this.ClassId = classId
	var objectType string = "kvm.Session"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KvmSessionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KvmSessionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KvmSessionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KvmSessionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KvmSessionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KvmSessionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKvmLaunchUrlPath returns the KvmLaunchUrlPath field value if set, zero value otherwise.
func (o *KvmSessionAllOf) GetKvmLaunchUrlPath() string {
	if o == nil || o.KvmLaunchUrlPath == nil {
		var ret string
		return ret
	}
	return *o.KvmLaunchUrlPath
}

// GetKvmLaunchUrlPathOk returns a tuple with the KvmLaunchUrlPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSessionAllOf) GetKvmLaunchUrlPathOk() (*string, bool) {
	if o == nil || o.KvmLaunchUrlPath == nil {
		return nil, false
	}
	return o.KvmLaunchUrlPath, true
}

// HasKvmLaunchUrlPath returns a boolean if a field has been set.
func (o *KvmSessionAllOf) HasKvmLaunchUrlPath() bool {
	if o != nil && o.KvmLaunchUrlPath != nil {
		return true
	}

	return false
}

// SetKvmLaunchUrlPath gets a reference to the given string and assigns it to the KvmLaunchUrlPath field.
func (o *KvmSessionAllOf) SetKvmLaunchUrlPath(v string) {
	o.KvmLaunchUrlPath = &v
}

// GetKvmSessionId returns the KvmSessionId field value if set, zero value otherwise.
func (o *KvmSessionAllOf) GetKvmSessionId() string {
	if o == nil || o.KvmSessionId == nil {
		var ret string
		return ret
	}
	return *o.KvmSessionId
}

// GetKvmSessionIdOk returns a tuple with the KvmSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSessionAllOf) GetKvmSessionIdOk() (*string, bool) {
	if o == nil || o.KvmSessionId == nil {
		return nil, false
	}
	return o.KvmSessionId, true
}

// HasKvmSessionId returns a boolean if a field has been set.
func (o *KvmSessionAllOf) HasKvmSessionId() bool {
	if o != nil && o.KvmSessionId != nil {
		return true
	}

	return false
}

// SetKvmSessionId gets a reference to the given string and assigns it to the KvmSessionId field.
func (o *KvmSessionAllOf) SetKvmSessionId(v string) {
	o.KvmSessionId = &v
}

// GetOneTimePassword returns the OneTimePassword field value if set, zero value otherwise.
func (o *KvmSessionAllOf) GetOneTimePassword() string {
	if o == nil || o.OneTimePassword == nil {
		var ret string
		return ret
	}
	return *o.OneTimePassword
}

// GetOneTimePasswordOk returns a tuple with the OneTimePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSessionAllOf) GetOneTimePasswordOk() (*string, bool) {
	if o == nil || o.OneTimePassword == nil {
		return nil, false
	}
	return o.OneTimePassword, true
}

// HasOneTimePassword returns a boolean if a field has been set.
func (o *KvmSessionAllOf) HasOneTimePassword() bool {
	if o != nil && o.OneTimePassword != nil {
		return true
	}

	return false
}

// SetOneTimePassword gets a reference to the given string and assigns it to the OneTimePassword field.
func (o *KvmSessionAllOf) SetOneTimePassword(v string) {
	o.OneTimePassword = &v
}

// GetSsoSupported returns the SsoSupported field value if set, zero value otherwise.
func (o *KvmSessionAllOf) GetSsoSupported() bool {
	if o == nil || o.SsoSupported == nil {
		var ret bool
		return ret
	}
	return *o.SsoSupported
}

// GetSsoSupportedOk returns a tuple with the SsoSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSessionAllOf) GetSsoSupportedOk() (*bool, bool) {
	if o == nil || o.SsoSupported == nil {
		return nil, false
	}
	return o.SsoSupported, true
}

// HasSsoSupported returns a boolean if a field has been set.
func (o *KvmSessionAllOf) HasSsoSupported() bool {
	if o != nil && o.SsoSupported != nil {
		return true
	}

	return false
}

// SetSsoSupported gets a reference to the given bool and assigns it to the SsoSupported field.
func (o *KvmSessionAllOf) SetSsoSupported(v bool) {
	o.SsoSupported = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *KvmSessionAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSessionAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *KvmSessionAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *KvmSessionAllOf) SetUsername(v string) {
	o.Username = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *KvmSessionAllOf) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSessionAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *KvmSessionAllOf) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *KvmSessionAllOf) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *KvmSessionAllOf) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSessionAllOf) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *KvmSessionAllOf) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *KvmSessionAllOf) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

// GetTunnel returns the Tunnel field value if set, zero value otherwise.
func (o *KvmSessionAllOf) GetTunnel() KvmTunnelRelationship {
	if o == nil || o.Tunnel == nil {
		var ret KvmTunnelRelationship
		return ret
	}
	return *o.Tunnel
}

// GetTunnelOk returns a tuple with the Tunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSessionAllOf) GetTunnelOk() (*KvmTunnelRelationship, bool) {
	if o == nil || o.Tunnel == nil {
		return nil, false
	}
	return o.Tunnel, true
}

// HasTunnel returns a boolean if a field has been set.
func (o *KvmSessionAllOf) HasTunnel() bool {
	if o != nil && o.Tunnel != nil {
		return true
	}

	return false
}

// SetTunnel gets a reference to the given KvmTunnelRelationship and assigns it to the Tunnel field.
func (o *KvmSessionAllOf) SetTunnel(v KvmTunnelRelationship) {
	o.Tunnel = &v
}

func (o KvmSessionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.KvmLaunchUrlPath != nil {
		toSerialize["KvmLaunchUrlPath"] = o.KvmLaunchUrlPath
	}
	if o.KvmSessionId != nil {
		toSerialize["KvmSessionId"] = o.KvmSessionId
	}
	if o.OneTimePassword != nil {
		toSerialize["OneTimePassword"] = o.OneTimePassword
	}
	if o.SsoSupported != nil {
		toSerialize["SsoSupported"] = o.SsoSupported
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}
	if o.Tunnel != nil {
		toSerialize["Tunnel"] = o.Tunnel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KvmSessionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKvmSessionAllOf := _KvmSessionAllOf{}

	if err = json.Unmarshal(bytes, &varKvmSessionAllOf); err == nil {
		*o = KvmSessionAllOf(varKvmSessionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "KvmLaunchUrlPath")
		delete(additionalProperties, "KvmSessionId")
		delete(additionalProperties, "OneTimePassword")
		delete(additionalProperties, "SsoSupported")
		delete(additionalProperties, "Username")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "Server")
		delete(additionalProperties, "Tunnel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKvmSessionAllOf struct {
	value *KvmSessionAllOf
	isSet bool
}

func (v NullableKvmSessionAllOf) Get() *KvmSessionAllOf {
	return v.value
}

func (v *NullableKvmSessionAllOf) Set(val *KvmSessionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmSessionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmSessionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmSessionAllOf(val *KvmSessionAllOf) *NullableKvmSessionAllOf {
	return &NullableKvmSessionAllOf{value: val, isSet: true}
}

func (v NullableKvmSessionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmSessionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


