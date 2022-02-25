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

// CloudNetworkAccessConfigAllOf Definition of the list of properties defined in 'cloud.NetworkAccessConfig', excluding properties defined in parent classes.
type CloudNetworkAccessConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	ExternalIps []CloudNetworkAddress `json:"ExternalIps,omitempty"`
	// Private DNS name assigned to the network interface.
	PrivateDns *string `json:"PrivateDns,omitempty"`
	// Public DNS name assigned to the network interface.
	PublicDns *string `json:"PublicDns,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudNetworkAccessConfigAllOf CloudNetworkAccessConfigAllOf

// NewCloudNetworkAccessConfigAllOf instantiates a new CloudNetworkAccessConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudNetworkAccessConfigAllOf(classId string, objectType string) *CloudNetworkAccessConfigAllOf {
	this := CloudNetworkAccessConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudNetworkAccessConfigAllOfWithDefaults instantiates a new CloudNetworkAccessConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudNetworkAccessConfigAllOfWithDefaults() *CloudNetworkAccessConfigAllOf {
	this := CloudNetworkAccessConfigAllOf{}
	var classId string = "cloud.NetworkAccessConfig"
	this.ClassId = classId
	var objectType string = "cloud.NetworkAccessConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudNetworkAccessConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudNetworkAccessConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudNetworkAccessConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudNetworkAccessConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudNetworkAccessConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudNetworkAccessConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExternalIps returns the ExternalIps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudNetworkAccessConfigAllOf) GetExternalIps() []CloudNetworkAddress {
	if o == nil  {
		var ret []CloudNetworkAddress
		return ret
	}
	return o.ExternalIps
}

// GetExternalIpsOk returns a tuple with the ExternalIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudNetworkAccessConfigAllOf) GetExternalIpsOk() (*[]CloudNetworkAddress, bool) {
	if o == nil || o.ExternalIps == nil {
		return nil, false
	}
	return &o.ExternalIps, true
}

// HasExternalIps returns a boolean if a field has been set.
func (o *CloudNetworkAccessConfigAllOf) HasExternalIps() bool {
	if o != nil && o.ExternalIps != nil {
		return true
	}

	return false
}

// SetExternalIps gets a reference to the given []CloudNetworkAddress and assigns it to the ExternalIps field.
func (o *CloudNetworkAccessConfigAllOf) SetExternalIps(v []CloudNetworkAddress) {
	o.ExternalIps = v
}

// GetPrivateDns returns the PrivateDns field value if set, zero value otherwise.
func (o *CloudNetworkAccessConfigAllOf) GetPrivateDns() string {
	if o == nil || o.PrivateDns == nil {
		var ret string
		return ret
	}
	return *o.PrivateDns
}

// GetPrivateDnsOk returns a tuple with the PrivateDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkAccessConfigAllOf) GetPrivateDnsOk() (*string, bool) {
	if o == nil || o.PrivateDns == nil {
		return nil, false
	}
	return o.PrivateDns, true
}

// HasPrivateDns returns a boolean if a field has been set.
func (o *CloudNetworkAccessConfigAllOf) HasPrivateDns() bool {
	if o != nil && o.PrivateDns != nil {
		return true
	}

	return false
}

// SetPrivateDns gets a reference to the given string and assigns it to the PrivateDns field.
func (o *CloudNetworkAccessConfigAllOf) SetPrivateDns(v string) {
	o.PrivateDns = &v
}

// GetPublicDns returns the PublicDns field value if set, zero value otherwise.
func (o *CloudNetworkAccessConfigAllOf) GetPublicDns() string {
	if o == nil || o.PublicDns == nil {
		var ret string
		return ret
	}
	return *o.PublicDns
}

// GetPublicDnsOk returns a tuple with the PublicDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkAccessConfigAllOf) GetPublicDnsOk() (*string, bool) {
	if o == nil || o.PublicDns == nil {
		return nil, false
	}
	return o.PublicDns, true
}

// HasPublicDns returns a boolean if a field has been set.
func (o *CloudNetworkAccessConfigAllOf) HasPublicDns() bool {
	if o != nil && o.PublicDns != nil {
		return true
	}

	return false
}

// SetPublicDns gets a reference to the given string and assigns it to the PublicDns field.
func (o *CloudNetworkAccessConfigAllOf) SetPublicDns(v string) {
	o.PublicDns = &v
}

func (o CloudNetworkAccessConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExternalIps != nil {
		toSerialize["ExternalIps"] = o.ExternalIps
	}
	if o.PrivateDns != nil {
		toSerialize["PrivateDns"] = o.PrivateDns
	}
	if o.PublicDns != nil {
		toSerialize["PublicDns"] = o.PublicDns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudNetworkAccessConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudNetworkAccessConfigAllOf := _CloudNetworkAccessConfigAllOf{}

	if err = json.Unmarshal(bytes, &varCloudNetworkAccessConfigAllOf); err == nil {
		*o = CloudNetworkAccessConfigAllOf(varCloudNetworkAccessConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExternalIps")
		delete(additionalProperties, "PrivateDns")
		delete(additionalProperties, "PublicDns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudNetworkAccessConfigAllOf struct {
	value *CloudNetworkAccessConfigAllOf
	isSet bool
}

func (v NullableCloudNetworkAccessConfigAllOf) Get() *CloudNetworkAccessConfigAllOf {
	return v.value
}

func (v *NullableCloudNetworkAccessConfigAllOf) Set(val *CloudNetworkAccessConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudNetworkAccessConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudNetworkAccessConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudNetworkAccessConfigAllOf(val *CloudNetworkAccessConfigAllOf) *NullableCloudNetworkAccessConfigAllOf {
	return &NullableCloudNetworkAccessConfigAllOf{value: val, isSet: true}
}

func (v NullableCloudNetworkAccessConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudNetworkAccessConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


