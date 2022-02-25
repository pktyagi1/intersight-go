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

// TerraformCloudResourceAllOf Definition of the list of properties defined in 'terraform.CloudResource', excluding properties defined in parent classes.
type TerraformCloudResourceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Currentstatus of the resource if applicable on the cloud.
	CurrentStatus *string `json:"CurrentStatus,omitempty"`
	// Desiredstatus of the resource if applicable on the cloud.
	DesiredStatus *string `json:"DesiredStatus,omitempty"`
	// Unique id of the resource from the cloud provider.
	ResourceId *string `json:"ResourceId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TerraformCloudResourceAllOf TerraformCloudResourceAllOf

// NewTerraformCloudResourceAllOf instantiates a new TerraformCloudResourceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerraformCloudResourceAllOf(classId string, objectType string) *TerraformCloudResourceAllOf {
	this := TerraformCloudResourceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTerraformCloudResourceAllOfWithDefaults instantiates a new TerraformCloudResourceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerraformCloudResourceAllOfWithDefaults() *TerraformCloudResourceAllOf {
	this := TerraformCloudResourceAllOf{}
	var classId string = "terraform.CloudResource"
	this.ClassId = classId
	var objectType string = "terraform.CloudResource"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TerraformCloudResourceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TerraformCloudResourceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TerraformCloudResourceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TerraformCloudResourceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TerraformCloudResourceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TerraformCloudResourceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCurrentStatus returns the CurrentStatus field value if set, zero value otherwise.
func (o *TerraformCloudResourceAllOf) GetCurrentStatus() string {
	if o == nil || o.CurrentStatus == nil {
		var ret string
		return ret
	}
	return *o.CurrentStatus
}

// GetCurrentStatusOk returns a tuple with the CurrentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformCloudResourceAllOf) GetCurrentStatusOk() (*string, bool) {
	if o == nil || o.CurrentStatus == nil {
		return nil, false
	}
	return o.CurrentStatus, true
}

// HasCurrentStatus returns a boolean if a field has been set.
func (o *TerraformCloudResourceAllOf) HasCurrentStatus() bool {
	if o != nil && o.CurrentStatus != nil {
		return true
	}

	return false
}

// SetCurrentStatus gets a reference to the given string and assigns it to the CurrentStatus field.
func (o *TerraformCloudResourceAllOf) SetCurrentStatus(v string) {
	o.CurrentStatus = &v
}

// GetDesiredStatus returns the DesiredStatus field value if set, zero value otherwise.
func (o *TerraformCloudResourceAllOf) GetDesiredStatus() string {
	if o == nil || o.DesiredStatus == nil {
		var ret string
		return ret
	}
	return *o.DesiredStatus
}

// GetDesiredStatusOk returns a tuple with the DesiredStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformCloudResourceAllOf) GetDesiredStatusOk() (*string, bool) {
	if o == nil || o.DesiredStatus == nil {
		return nil, false
	}
	return o.DesiredStatus, true
}

// HasDesiredStatus returns a boolean if a field has been set.
func (o *TerraformCloudResourceAllOf) HasDesiredStatus() bool {
	if o != nil && o.DesiredStatus != nil {
		return true
	}

	return false
}

// SetDesiredStatus gets a reference to the given string and assigns it to the DesiredStatus field.
func (o *TerraformCloudResourceAllOf) SetDesiredStatus(v string) {
	o.DesiredStatus = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *TerraformCloudResourceAllOf) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformCloudResourceAllOf) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *TerraformCloudResourceAllOf) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *TerraformCloudResourceAllOf) SetResourceId(v string) {
	o.ResourceId = &v
}

func (o TerraformCloudResourceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CurrentStatus != nil {
		toSerialize["CurrentStatus"] = o.CurrentStatus
	}
	if o.DesiredStatus != nil {
		toSerialize["DesiredStatus"] = o.DesiredStatus
	}
	if o.ResourceId != nil {
		toSerialize["ResourceId"] = o.ResourceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TerraformCloudResourceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTerraformCloudResourceAllOf := _TerraformCloudResourceAllOf{}

	if err = json.Unmarshal(bytes, &varTerraformCloudResourceAllOf); err == nil {
		*o = TerraformCloudResourceAllOf(varTerraformCloudResourceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CurrentStatus")
		delete(additionalProperties, "DesiredStatus")
		delete(additionalProperties, "ResourceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTerraformCloudResourceAllOf struct {
	value *TerraformCloudResourceAllOf
	isSet bool
}

func (v NullableTerraformCloudResourceAllOf) Get() *TerraformCloudResourceAllOf {
	return v.value
}

func (v *NullableTerraformCloudResourceAllOf) Set(val *TerraformCloudResourceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTerraformCloudResourceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTerraformCloudResourceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerraformCloudResourceAllOf(val *TerraformCloudResourceAllOf) *NullableTerraformCloudResourceAllOf {
	return &NullableTerraformCloudResourceAllOf{value: val, isSet: true}
}

func (v NullableTerraformCloudResourceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerraformCloudResourceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


