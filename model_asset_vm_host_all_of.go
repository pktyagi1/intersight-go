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

// AssetVmHostAllOf Definition of the list of properties defined in 'asset.VmHost', excluding properties defined in parent classes.
type AssetVmHostAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Reference to virtualization target account ID.
	AccountMoid *string `json:"AccountMoid,omitempty"`
	// Reference to virtualization cluster identity.
	ClusterIdentity *string `json:"ClusterIdentity,omitempty"`
	// Reference to virtualization cluster ID.
	ClusterMoid *string `json:"ClusterMoid,omitempty"`
	// Reference to virtualization cluster name.
	ClusterName *string `json:"ClusterName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetVmHostAllOf AssetVmHostAllOf

// NewAssetVmHostAllOf instantiates a new AssetVmHostAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetVmHostAllOf(classId string, objectType string) *AssetVmHostAllOf {
	this := AssetVmHostAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetVmHostAllOfWithDefaults instantiates a new AssetVmHostAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetVmHostAllOfWithDefaults() *AssetVmHostAllOf {
	this := AssetVmHostAllOf{}
	var classId string = "asset.VmHost"
	this.ClassId = classId
	var objectType string = "asset.VmHost"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetVmHostAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetVmHostAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetVmHostAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetVmHostAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetVmHostAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetVmHostAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccountMoid returns the AccountMoid field value if set, zero value otherwise.
func (o *AssetVmHostAllOf) GetAccountMoid() string {
	if o == nil || o.AccountMoid == nil {
		var ret string
		return ret
	}
	return *o.AccountMoid
}

// GetAccountMoidOk returns a tuple with the AccountMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetVmHostAllOf) GetAccountMoidOk() (*string, bool) {
	if o == nil || o.AccountMoid == nil {
		return nil, false
	}
	return o.AccountMoid, true
}

// HasAccountMoid returns a boolean if a field has been set.
func (o *AssetVmHostAllOf) HasAccountMoid() bool {
	if o != nil && o.AccountMoid != nil {
		return true
	}

	return false
}

// SetAccountMoid gets a reference to the given string and assigns it to the AccountMoid field.
func (o *AssetVmHostAllOf) SetAccountMoid(v string) {
	o.AccountMoid = &v
}

// GetClusterIdentity returns the ClusterIdentity field value if set, zero value otherwise.
func (o *AssetVmHostAllOf) GetClusterIdentity() string {
	if o == nil || o.ClusterIdentity == nil {
		var ret string
		return ret
	}
	return *o.ClusterIdentity
}

// GetClusterIdentityOk returns a tuple with the ClusterIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetVmHostAllOf) GetClusterIdentityOk() (*string, bool) {
	if o == nil || o.ClusterIdentity == nil {
		return nil, false
	}
	return o.ClusterIdentity, true
}

// HasClusterIdentity returns a boolean if a field has been set.
func (o *AssetVmHostAllOf) HasClusterIdentity() bool {
	if o != nil && o.ClusterIdentity != nil {
		return true
	}

	return false
}

// SetClusterIdentity gets a reference to the given string and assigns it to the ClusterIdentity field.
func (o *AssetVmHostAllOf) SetClusterIdentity(v string) {
	o.ClusterIdentity = &v
}

// GetClusterMoid returns the ClusterMoid field value if set, zero value otherwise.
func (o *AssetVmHostAllOf) GetClusterMoid() string {
	if o == nil || o.ClusterMoid == nil {
		var ret string
		return ret
	}
	return *o.ClusterMoid
}

// GetClusterMoidOk returns a tuple with the ClusterMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetVmHostAllOf) GetClusterMoidOk() (*string, bool) {
	if o == nil || o.ClusterMoid == nil {
		return nil, false
	}
	return o.ClusterMoid, true
}

// HasClusterMoid returns a boolean if a field has been set.
func (o *AssetVmHostAllOf) HasClusterMoid() bool {
	if o != nil && o.ClusterMoid != nil {
		return true
	}

	return false
}

// SetClusterMoid gets a reference to the given string and assigns it to the ClusterMoid field.
func (o *AssetVmHostAllOf) SetClusterMoid(v string) {
	o.ClusterMoid = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *AssetVmHostAllOf) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetVmHostAllOf) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *AssetVmHostAllOf) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *AssetVmHostAllOf) SetClusterName(v string) {
	o.ClusterName = &v
}

func (o AssetVmHostAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccountMoid != nil {
		toSerialize["AccountMoid"] = o.AccountMoid
	}
	if o.ClusterIdentity != nil {
		toSerialize["ClusterIdentity"] = o.ClusterIdentity
	}
	if o.ClusterMoid != nil {
		toSerialize["ClusterMoid"] = o.ClusterMoid
	}
	if o.ClusterName != nil {
		toSerialize["ClusterName"] = o.ClusterName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetVmHostAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetVmHostAllOf := _AssetVmHostAllOf{}

	if err = json.Unmarshal(bytes, &varAssetVmHostAllOf); err == nil {
		*o = AssetVmHostAllOf(varAssetVmHostAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccountMoid")
		delete(additionalProperties, "ClusterIdentity")
		delete(additionalProperties, "ClusterMoid")
		delete(additionalProperties, "ClusterName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetVmHostAllOf struct {
	value *AssetVmHostAllOf
	isSet bool
}

func (v NullableAssetVmHostAllOf) Get() *AssetVmHostAllOf {
	return v.value
}

func (v *NullableAssetVmHostAllOf) Set(val *AssetVmHostAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetVmHostAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetVmHostAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetVmHostAllOf(val *AssetVmHostAllOf) *NullableAssetVmHostAllOf {
	return &NullableAssetVmHostAllOf{value: val, isSet: true}
}

func (v NullableAssetVmHostAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetVmHostAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


