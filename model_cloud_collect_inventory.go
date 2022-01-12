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

// CloudCollectInventory REST API endpoint for inventory collection. It is invoked asynchronously after TFC creation. The new target is sent to inventory collector to fetch organizations, agent pools and workspaces.
type CloudCollectInventory struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The id of the new Terraform cloud asset which was created.
	TargetId *string `json:"TargetId,omitempty"`
	Target *AssetTargetRelationship `json:"Target,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudCollectInventory CloudCollectInventory

// NewCloudCollectInventory instantiates a new CloudCollectInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudCollectInventory(classId string, objectType string) *CloudCollectInventory {
	this := CloudCollectInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudCollectInventoryWithDefaults instantiates a new CloudCollectInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudCollectInventoryWithDefaults() *CloudCollectInventory {
	this := CloudCollectInventory{}
	var classId string = "cloud.CollectInventory"
	this.ClassId = classId
	var objectType string = "cloud.CollectInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudCollectInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudCollectInventory) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudCollectInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudCollectInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudCollectInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudCollectInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *CloudCollectInventory) GetTargetId() string {
	if o == nil || o.TargetId == nil {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCollectInventory) GetTargetIdOk() (*string, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *CloudCollectInventory) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *CloudCollectInventory) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *CloudCollectInventory) GetTarget() AssetTargetRelationship {
	if o == nil || o.Target == nil {
		var ret AssetTargetRelationship
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCollectInventory) GetTargetOk() (*AssetTargetRelationship, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *CloudCollectInventory) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AssetTargetRelationship and assigns it to the Target field.
func (o *CloudCollectInventory) SetTarget(v AssetTargetRelationship) {
	o.Target = &v
}

func (o CloudCollectInventory) MarshalJSON() ([]byte, error) {
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
	if o.TargetId != nil {
		toSerialize["TargetId"] = o.TargetId
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudCollectInventory) UnmarshalJSON(bytes []byte) (err error) {
	type CloudCollectInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The id of the new Terraform cloud asset which was created.
		TargetId *string `json:"TargetId,omitempty"`
		Target *AssetTargetRelationship `json:"Target,omitempty"`
	}

	varCloudCollectInventoryWithoutEmbeddedStruct := CloudCollectInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudCollectInventoryWithoutEmbeddedStruct)
	if err == nil {
		varCloudCollectInventory := _CloudCollectInventory{}
		varCloudCollectInventory.ClassId = varCloudCollectInventoryWithoutEmbeddedStruct.ClassId
		varCloudCollectInventory.ObjectType = varCloudCollectInventoryWithoutEmbeddedStruct.ObjectType
		varCloudCollectInventory.TargetId = varCloudCollectInventoryWithoutEmbeddedStruct.TargetId
		varCloudCollectInventory.Target = varCloudCollectInventoryWithoutEmbeddedStruct.Target
		*o = CloudCollectInventory(varCloudCollectInventory)
	} else {
		return err
	}

	varCloudCollectInventory := _CloudCollectInventory{}

	err = json.Unmarshal(bytes, &varCloudCollectInventory)
	if err == nil {
		o.MoBaseMo = varCloudCollectInventory.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetId")
		delete(additionalProperties, "Target")

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

type NullableCloudCollectInventory struct {
	value *CloudCollectInventory
	isSet bool
}

func (v NullableCloudCollectInventory) Get() *CloudCollectInventory {
	return v.value
}

func (v *NullableCloudCollectInventory) Set(val *CloudCollectInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudCollectInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudCollectInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudCollectInventory(val *CloudCollectInventory) *NullableCloudCollectInventory {
	return &NullableCloudCollectInventory{value: val, isSet: true}
}

func (v NullableCloudCollectInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudCollectInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

