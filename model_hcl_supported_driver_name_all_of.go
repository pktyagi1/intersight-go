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

// HclSupportedDriverNameAllOf Definition of the list of properties defined in 'hcl.SupportedDriverName', excluding properties defined in parent classes.
type HclSupportedDriverNameAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Vendor distributing the Operating System.
	OsVendor *string `json:"OsVendor,omitempty"`
	// Version of the Operating System.
	OsVersion *string `json:"OsVersion,omitempty"`
	ProductList []HclProduct `json:"ProductList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclSupportedDriverNameAllOf HclSupportedDriverNameAllOf

// NewHclSupportedDriverNameAllOf instantiates a new HclSupportedDriverNameAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclSupportedDriverNameAllOf(classId string, objectType string) *HclSupportedDriverNameAllOf {
	this := HclSupportedDriverNameAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHclSupportedDriverNameAllOfWithDefaults instantiates a new HclSupportedDriverNameAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclSupportedDriverNameAllOfWithDefaults() *HclSupportedDriverNameAllOf {
	this := HclSupportedDriverNameAllOf{}
	var classId string = "hcl.SupportedDriverName"
	this.ClassId = classId
	var objectType string = "hcl.SupportedDriverName"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HclSupportedDriverNameAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HclSupportedDriverNameAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HclSupportedDriverNameAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HclSupportedDriverNameAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HclSupportedDriverNameAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HclSupportedDriverNameAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOsVendor returns the OsVendor field value if set, zero value otherwise.
func (o *HclSupportedDriverNameAllOf) GetOsVendor() string {
	if o == nil || o.OsVendor == nil {
		var ret string
		return ret
	}
	return *o.OsVendor
}

// GetOsVendorOk returns a tuple with the OsVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclSupportedDriverNameAllOf) GetOsVendorOk() (*string, bool) {
	if o == nil || o.OsVendor == nil {
		return nil, false
	}
	return o.OsVendor, true
}

// HasOsVendor returns a boolean if a field has been set.
func (o *HclSupportedDriverNameAllOf) HasOsVendor() bool {
	if o != nil && o.OsVendor != nil {
		return true
	}

	return false
}

// SetOsVendor gets a reference to the given string and assigns it to the OsVendor field.
func (o *HclSupportedDriverNameAllOf) SetOsVendor(v string) {
	o.OsVendor = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *HclSupportedDriverNameAllOf) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclSupportedDriverNameAllOf) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *HclSupportedDriverNameAllOf) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *HclSupportedDriverNameAllOf) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetProductList returns the ProductList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HclSupportedDriverNameAllOf) GetProductList() []HclProduct {
	if o == nil  {
		var ret []HclProduct
		return ret
	}
	return o.ProductList
}

// GetProductListOk returns a tuple with the ProductList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HclSupportedDriverNameAllOf) GetProductListOk() (*[]HclProduct, bool) {
	if o == nil || o.ProductList == nil {
		return nil, false
	}
	return &o.ProductList, true
}

// HasProductList returns a boolean if a field has been set.
func (o *HclSupportedDriverNameAllOf) HasProductList() bool {
	if o != nil && o.ProductList != nil {
		return true
	}

	return false
}

// SetProductList gets a reference to the given []HclProduct and assigns it to the ProductList field.
func (o *HclSupportedDriverNameAllOf) SetProductList(v []HclProduct) {
	o.ProductList = v
}

func (o HclSupportedDriverNameAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OsVendor != nil {
		toSerialize["OsVendor"] = o.OsVendor
	}
	if o.OsVersion != nil {
		toSerialize["OsVersion"] = o.OsVersion
	}
	if o.ProductList != nil {
		toSerialize["ProductList"] = o.ProductList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HclSupportedDriverNameAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHclSupportedDriverNameAllOf := _HclSupportedDriverNameAllOf{}

	if err = json.Unmarshal(bytes, &varHclSupportedDriverNameAllOf); err == nil {
		*o = HclSupportedDriverNameAllOf(varHclSupportedDriverNameAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OsVendor")
		delete(additionalProperties, "OsVersion")
		delete(additionalProperties, "ProductList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHclSupportedDriverNameAllOf struct {
	value *HclSupportedDriverNameAllOf
	isSet bool
}

func (v NullableHclSupportedDriverNameAllOf) Get() *HclSupportedDriverNameAllOf {
	return v.value
}

func (v *NullableHclSupportedDriverNameAllOf) Set(val *HclSupportedDriverNameAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHclSupportedDriverNameAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHclSupportedDriverNameAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclSupportedDriverNameAllOf(val *HclSupportedDriverNameAllOf) *NullableHclSupportedDriverNameAllOf {
	return &NullableHclSupportedDriverNameAllOf{value: val, isSet: true}
}

func (v NullableHclSupportedDriverNameAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclSupportedDriverNameAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


