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
)

// AssetWorkloadOptimizerHypervOptionsAllOf Definition of the list of properties defined in 'asset.WorkloadOptimizerHypervOptions', excluding properties defined in parent classes.
type AssetWorkloadOptimizerHypervOptionsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When enabled, all Hyper-V hosts in the same cluster of specified Hyper-V host target will be discovered Each server must still be configured to allow remote management (for example, configuring WinRM using a GPO).
	DiscoverHostCluster *bool `json:"DiscoverHostCluster,omitempty"`
	// Fully qualified domain name of the Hyper-V target. It is used to get the name of the Hyper-V cluster node and do Active Directory authentication process through Kerberos scheme.
	FullDomainName *string `json:"FullDomainName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetWorkloadOptimizerHypervOptionsAllOf AssetWorkloadOptimizerHypervOptionsAllOf

// NewAssetWorkloadOptimizerHypervOptionsAllOf instantiates a new AssetWorkloadOptimizerHypervOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerHypervOptionsAllOf(classId string, objectType string) *AssetWorkloadOptimizerHypervOptionsAllOf {
	this := AssetWorkloadOptimizerHypervOptionsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerHypervOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerHypervOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerHypervOptionsAllOfWithDefaults() *AssetWorkloadOptimizerHypervOptionsAllOf {
	this := AssetWorkloadOptimizerHypervOptionsAllOf{}
	var classId string = "asset.WorkloadOptimizerHypervOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerHypervOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerHypervOptionsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerHypervOptionsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerHypervOptionsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerHypervOptionsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerHypervOptionsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerHypervOptionsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDiscoverHostCluster returns the DiscoverHostCluster field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerHypervOptionsAllOf) GetDiscoverHostCluster() bool {
	if o == nil || o.DiscoverHostCluster == nil {
		var ret bool
		return ret
	}
	return *o.DiscoverHostCluster
}

// GetDiscoverHostClusterOk returns a tuple with the DiscoverHostCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerHypervOptionsAllOf) GetDiscoverHostClusterOk() (*bool, bool) {
	if o == nil || o.DiscoverHostCluster == nil {
		return nil, false
	}
	return o.DiscoverHostCluster, true
}

// HasDiscoverHostCluster returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerHypervOptionsAllOf) HasDiscoverHostCluster() bool {
	if o != nil && o.DiscoverHostCluster != nil {
		return true
	}

	return false
}

// SetDiscoverHostCluster gets a reference to the given bool and assigns it to the DiscoverHostCluster field.
func (o *AssetWorkloadOptimizerHypervOptionsAllOf) SetDiscoverHostCluster(v bool) {
	o.DiscoverHostCluster = &v
}

// GetFullDomainName returns the FullDomainName field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerHypervOptionsAllOf) GetFullDomainName() string {
	if o == nil || o.FullDomainName == nil {
		var ret string
		return ret
	}
	return *o.FullDomainName
}

// GetFullDomainNameOk returns a tuple with the FullDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerHypervOptionsAllOf) GetFullDomainNameOk() (*string, bool) {
	if o == nil || o.FullDomainName == nil {
		return nil, false
	}
	return o.FullDomainName, true
}

// HasFullDomainName returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerHypervOptionsAllOf) HasFullDomainName() bool {
	if o != nil && o.FullDomainName != nil {
		return true
	}

	return false
}

// SetFullDomainName gets a reference to the given string and assigns it to the FullDomainName field.
func (o *AssetWorkloadOptimizerHypervOptionsAllOf) SetFullDomainName(v string) {
	o.FullDomainName = &v
}

func (o AssetWorkloadOptimizerHypervOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DiscoverHostCluster != nil {
		toSerialize["DiscoverHostCluster"] = o.DiscoverHostCluster
	}
	if o.FullDomainName != nil {
		toSerialize["FullDomainName"] = o.FullDomainName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerHypervOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetWorkloadOptimizerHypervOptionsAllOf := _AssetWorkloadOptimizerHypervOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerHypervOptionsAllOf); err == nil {
		*o = AssetWorkloadOptimizerHypervOptionsAllOf(varAssetWorkloadOptimizerHypervOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DiscoverHostCluster")
		delete(additionalProperties, "FullDomainName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetWorkloadOptimizerHypervOptionsAllOf struct {
	value *AssetWorkloadOptimizerHypervOptionsAllOf
	isSet bool
}

func (v NullableAssetWorkloadOptimizerHypervOptionsAllOf) Get() *AssetWorkloadOptimizerHypervOptionsAllOf {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerHypervOptionsAllOf) Set(val *AssetWorkloadOptimizerHypervOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerHypervOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerHypervOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerHypervOptionsAllOf(val *AssetWorkloadOptimizerHypervOptionsAllOf) *NullableAssetWorkloadOptimizerHypervOptionsAllOf {
	return &NullableAssetWorkloadOptimizerHypervOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerHypervOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerHypervOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


