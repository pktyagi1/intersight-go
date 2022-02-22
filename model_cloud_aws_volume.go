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

// CloudAwsVolume AWS volume object is represented here.It depicts configuration used to create volume and its identity.
type CloudAwsVolume struct {
	CloudBaseVolume
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	AwsBillingUnit *CloudAwsBillingUnitRelationship `json:"AwsBillingUnit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudAwsVolume CloudAwsVolume

// NewCloudAwsVolume instantiates a new CloudAwsVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAwsVolume(classId string, objectType string) *CloudAwsVolume {
	this := CloudAwsVolume{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudAwsVolumeWithDefaults instantiates a new CloudAwsVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAwsVolumeWithDefaults() *CloudAwsVolume {
	this := CloudAwsVolume{}
	var classId string = "cloud.AwsVolume"
	this.ClassId = classId
	var objectType string = "cloud.AwsVolume"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudAwsVolume) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudAwsVolume) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudAwsVolume) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudAwsVolume) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudAwsVolume) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudAwsVolume) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAwsBillingUnit returns the AwsBillingUnit field value if set, zero value otherwise.
func (o *CloudAwsVolume) GetAwsBillingUnit() CloudAwsBillingUnitRelationship {
	if o == nil || o.AwsBillingUnit == nil {
		var ret CloudAwsBillingUnitRelationship
		return ret
	}
	return *o.AwsBillingUnit
}

// GetAwsBillingUnitOk returns a tuple with the AwsBillingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVolume) GetAwsBillingUnitOk() (*CloudAwsBillingUnitRelationship, bool) {
	if o == nil || o.AwsBillingUnit == nil {
		return nil, false
	}
	return o.AwsBillingUnit, true
}

// HasAwsBillingUnit returns a boolean if a field has been set.
func (o *CloudAwsVolume) HasAwsBillingUnit() bool {
	if o != nil && o.AwsBillingUnit != nil {
		return true
	}

	return false
}

// SetAwsBillingUnit gets a reference to the given CloudAwsBillingUnitRelationship and assigns it to the AwsBillingUnit field.
func (o *CloudAwsVolume) SetAwsBillingUnit(v CloudAwsBillingUnitRelationship) {
	o.AwsBillingUnit = &v
}

func (o CloudAwsVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVolume, errCloudBaseVolume := json.Marshal(o.CloudBaseVolume)
	if errCloudBaseVolume != nil {
		return []byte{}, errCloudBaseVolume
	}
	errCloudBaseVolume = json.Unmarshal([]byte(serializedCloudBaseVolume), &toSerialize)
	if errCloudBaseVolume != nil {
		return []byte{}, errCloudBaseVolume
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AwsBillingUnit != nil {
		toSerialize["AwsBillingUnit"] = o.AwsBillingUnit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudAwsVolume) UnmarshalJSON(bytes []byte) (err error) {
	type CloudAwsVolumeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		AwsBillingUnit *CloudAwsBillingUnitRelationship `json:"AwsBillingUnit,omitempty"`
	}

	varCloudAwsVolumeWithoutEmbeddedStruct := CloudAwsVolumeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudAwsVolumeWithoutEmbeddedStruct)
	if err == nil {
		varCloudAwsVolume := _CloudAwsVolume{}
		varCloudAwsVolume.ClassId = varCloudAwsVolumeWithoutEmbeddedStruct.ClassId
		varCloudAwsVolume.ObjectType = varCloudAwsVolumeWithoutEmbeddedStruct.ObjectType
		varCloudAwsVolume.AwsBillingUnit = varCloudAwsVolumeWithoutEmbeddedStruct.AwsBillingUnit
		*o = CloudAwsVolume(varCloudAwsVolume)
	} else {
		return err
	}

	varCloudAwsVolume := _CloudAwsVolume{}

	err = json.Unmarshal(bytes, &varCloudAwsVolume)
	if err == nil {
		o.CloudBaseVolume = varCloudAwsVolume.CloudBaseVolume
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AwsBillingUnit")

		// remove fields from embedded structs
		reflectCloudBaseVolume := reflect.ValueOf(o.CloudBaseVolume)
		for i := 0; i < reflectCloudBaseVolume.Type().NumField(); i++ {
			t := reflectCloudBaseVolume.Type().Field(i)

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

type NullableCloudAwsVolume struct {
	value *CloudAwsVolume
	isSet bool
}

func (v NullableCloudAwsVolume) Get() *CloudAwsVolume {
	return v.value
}

func (v *NullableCloudAwsVolume) Set(val *CloudAwsVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAwsVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAwsVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAwsVolume(val *CloudAwsVolume) *NullableCloudAwsVolume {
	return &NullableCloudAwsVolume{value: val, isSet: true}
}

func (v NullableCloudAwsVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAwsVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


