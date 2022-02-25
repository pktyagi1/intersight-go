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
	"reflect"
	"strings"
)

// CloudCloudRegion Geographic location where your public cloud resources are located.
type CloudCloudRegion struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of geographic location where your public cloud resources are located.
	Name *string `json:"Name,omitempty"`
	// The ID of geographic location where your public cloud resources are located.
	RegionId *string `json:"RegionId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudCloudRegion CloudCloudRegion

// NewCloudCloudRegion instantiates a new CloudCloudRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudCloudRegion(classId string, objectType string) *CloudCloudRegion {
	this := CloudCloudRegion{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudCloudRegionWithDefaults instantiates a new CloudCloudRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudCloudRegionWithDefaults() *CloudCloudRegion {
	this := CloudCloudRegion{}
	var classId string = "cloud.CloudRegion"
	this.ClassId = classId
	var objectType string = "cloud.CloudRegion"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudCloudRegion) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudCloudRegion) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudCloudRegion) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudCloudRegion) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudCloudRegion) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudCloudRegion) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudCloudRegion) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCloudRegion) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudCloudRegion) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudCloudRegion) SetName(v string) {
	o.Name = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *CloudCloudRegion) GetRegionId() string {
	if o == nil || o.RegionId == nil {
		var ret string
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCloudRegion) GetRegionIdOk() (*string, bool) {
	if o == nil || o.RegionId == nil {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *CloudCloudRegion) HasRegionId() bool {
	if o != nil && o.RegionId != nil {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given string and assigns it to the RegionId field.
func (o *CloudCloudRegion) SetRegionId(v string) {
	o.RegionId = &v
}

func (o CloudCloudRegion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.RegionId != nil {
		toSerialize["RegionId"] = o.RegionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudCloudRegion) UnmarshalJSON(bytes []byte) (err error) {
	type CloudCloudRegionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of geographic location where your public cloud resources are located.
		Name *string `json:"Name,omitempty"`
		// The ID of geographic location where your public cloud resources are located.
		RegionId *string `json:"RegionId,omitempty"`
	}

	varCloudCloudRegionWithoutEmbeddedStruct := CloudCloudRegionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudCloudRegionWithoutEmbeddedStruct)
	if err == nil {
		varCloudCloudRegion := _CloudCloudRegion{}
		varCloudCloudRegion.ClassId = varCloudCloudRegionWithoutEmbeddedStruct.ClassId
		varCloudCloudRegion.ObjectType = varCloudCloudRegionWithoutEmbeddedStruct.ObjectType
		varCloudCloudRegion.Name = varCloudCloudRegionWithoutEmbeddedStruct.Name
		varCloudCloudRegion.RegionId = varCloudCloudRegionWithoutEmbeddedStruct.RegionId
		*o = CloudCloudRegion(varCloudCloudRegion)
	} else {
		return err
	}

	varCloudCloudRegion := _CloudCloudRegion{}

	err = json.Unmarshal(bytes, &varCloudCloudRegion)
	if err == nil {
		o.MoBaseComplexType = varCloudCloudRegion.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RegionId")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableCloudCloudRegion struct {
	value *CloudCloudRegion
	isSet bool
}

func (v NullableCloudCloudRegion) Get() *CloudCloudRegion {
	return v.value
}

func (v *NullableCloudCloudRegion) Set(val *CloudCloudRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudCloudRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudCloudRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudCloudRegion(val *CloudCloudRegion) *NullableCloudCloudRegion {
	return &NullableCloudCloudRegion{value: val, isSet: true}
}

func (v NullableCloudCloudRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudCloudRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


