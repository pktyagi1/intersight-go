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
	"reflect"
	"strings"
)

// HclSupportedDriverName Supported driver names for a given product for the given operating system.
type HclSupportedDriverName struct {
	MoBaseMo
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

type _HclSupportedDriverName HclSupportedDriverName

// NewHclSupportedDriverName instantiates a new HclSupportedDriverName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclSupportedDriverName(classId string, objectType string) *HclSupportedDriverName {
	this := HclSupportedDriverName{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHclSupportedDriverNameWithDefaults instantiates a new HclSupportedDriverName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclSupportedDriverNameWithDefaults() *HclSupportedDriverName {
	this := HclSupportedDriverName{}
	var classId string = "hcl.SupportedDriverName"
	this.ClassId = classId
	var objectType string = "hcl.SupportedDriverName"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HclSupportedDriverName) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HclSupportedDriverName) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HclSupportedDriverName) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HclSupportedDriverName) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HclSupportedDriverName) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HclSupportedDriverName) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOsVendor returns the OsVendor field value if set, zero value otherwise.
func (o *HclSupportedDriverName) GetOsVendor() string {
	if o == nil || o.OsVendor == nil {
		var ret string
		return ret
	}
	return *o.OsVendor
}

// GetOsVendorOk returns a tuple with the OsVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclSupportedDriverName) GetOsVendorOk() (*string, bool) {
	if o == nil || o.OsVendor == nil {
		return nil, false
	}
	return o.OsVendor, true
}

// HasOsVendor returns a boolean if a field has been set.
func (o *HclSupportedDriverName) HasOsVendor() bool {
	if o != nil && o.OsVendor != nil {
		return true
	}

	return false
}

// SetOsVendor gets a reference to the given string and assigns it to the OsVendor field.
func (o *HclSupportedDriverName) SetOsVendor(v string) {
	o.OsVendor = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *HclSupportedDriverName) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclSupportedDriverName) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *HclSupportedDriverName) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *HclSupportedDriverName) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetProductList returns the ProductList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HclSupportedDriverName) GetProductList() []HclProduct {
	if o == nil  {
		var ret []HclProduct
		return ret
	}
	return o.ProductList
}

// GetProductListOk returns a tuple with the ProductList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HclSupportedDriverName) GetProductListOk() (*[]HclProduct, bool) {
	if o == nil || o.ProductList == nil {
		return nil, false
	}
	return &o.ProductList, true
}

// HasProductList returns a boolean if a field has been set.
func (o *HclSupportedDriverName) HasProductList() bool {
	if o != nil && o.ProductList != nil {
		return true
	}

	return false
}

// SetProductList gets a reference to the given []HclProduct and assigns it to the ProductList field.
func (o *HclSupportedDriverName) SetProductList(v []HclProduct) {
	o.ProductList = v
}

func (o HclSupportedDriverName) MarshalJSON() ([]byte, error) {
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

func (o *HclSupportedDriverName) UnmarshalJSON(bytes []byte) (err error) {
	type HclSupportedDriverNameWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Vendor distributing the Operating System.
		OsVendor *string `json:"OsVendor,omitempty"`
		// Version of the Operating System.
		OsVersion *string `json:"OsVersion,omitempty"`
		ProductList []HclProduct `json:"ProductList,omitempty"`
	}

	varHclSupportedDriverNameWithoutEmbeddedStruct := HclSupportedDriverNameWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHclSupportedDriverNameWithoutEmbeddedStruct)
	if err == nil {
		varHclSupportedDriverName := _HclSupportedDriverName{}
		varHclSupportedDriverName.ClassId = varHclSupportedDriverNameWithoutEmbeddedStruct.ClassId
		varHclSupportedDriverName.ObjectType = varHclSupportedDriverNameWithoutEmbeddedStruct.ObjectType
		varHclSupportedDriverName.OsVendor = varHclSupportedDriverNameWithoutEmbeddedStruct.OsVendor
		varHclSupportedDriverName.OsVersion = varHclSupportedDriverNameWithoutEmbeddedStruct.OsVersion
		varHclSupportedDriverName.ProductList = varHclSupportedDriverNameWithoutEmbeddedStruct.ProductList
		*o = HclSupportedDriverName(varHclSupportedDriverName)
	} else {
		return err
	}

	varHclSupportedDriverName := _HclSupportedDriverName{}

	err = json.Unmarshal(bytes, &varHclSupportedDriverName)
	if err == nil {
		o.MoBaseMo = varHclSupportedDriverName.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OsVendor")
		delete(additionalProperties, "OsVersion")
		delete(additionalProperties, "ProductList")

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

type NullableHclSupportedDriverName struct {
	value *HclSupportedDriverName
	isSet bool
}

func (v NullableHclSupportedDriverName) Get() *HclSupportedDriverName {
	return v.value
}

func (v *NullableHclSupportedDriverName) Set(val *HclSupportedDriverName) {
	v.value = val
	v.isSet = true
}

func (v NullableHclSupportedDriverName) IsSet() bool {
	return v.isSet
}

func (v *NullableHclSupportedDriverName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclSupportedDriverName(val *HclSupportedDriverName) *NullableHclSupportedDriverName {
	return &NullableHclSupportedDriverName{value: val, isSet: true}
}

func (v NullableHclSupportedDriverName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclSupportedDriverName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


