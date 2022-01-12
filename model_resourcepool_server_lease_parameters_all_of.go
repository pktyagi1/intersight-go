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

// ResourcepoolServerLeaseParametersAllOf Definition of the list of properties defined in 'resourcepool.ServerLeaseParameters', excluding properties defined in parent classes.
type ResourcepoolServerLeaseParametersAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	LifeCycleStates []string `json:"LifeCycleStates,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourcepoolServerLeaseParametersAllOf ResourcepoolServerLeaseParametersAllOf

// NewResourcepoolServerLeaseParametersAllOf instantiates a new ResourcepoolServerLeaseParametersAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcepoolServerLeaseParametersAllOf(classId string, objectType string) *ResourcepoolServerLeaseParametersAllOf {
	this := ResourcepoolServerLeaseParametersAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewResourcepoolServerLeaseParametersAllOfWithDefaults instantiates a new ResourcepoolServerLeaseParametersAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcepoolServerLeaseParametersAllOfWithDefaults() *ResourcepoolServerLeaseParametersAllOf {
	this := ResourcepoolServerLeaseParametersAllOf{}
	var classId string = "resourcepool.ServerLeaseParameters"
	this.ClassId = classId
	var objectType string = "resourcepool.ServerLeaseParameters"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourcepoolServerLeaseParametersAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolServerLeaseParametersAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourcepoolServerLeaseParametersAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourcepoolServerLeaseParametersAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolServerLeaseParametersAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourcepoolServerLeaseParametersAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLifeCycleStates returns the LifeCycleStates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcepoolServerLeaseParametersAllOf) GetLifeCycleStates() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.LifeCycleStates
}

// GetLifeCycleStatesOk returns a tuple with the LifeCycleStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcepoolServerLeaseParametersAllOf) GetLifeCycleStatesOk() (*[]string, bool) {
	if o == nil || o.LifeCycleStates == nil {
		return nil, false
	}
	return &o.LifeCycleStates, true
}

// HasLifeCycleStates returns a boolean if a field has been set.
func (o *ResourcepoolServerLeaseParametersAllOf) HasLifeCycleStates() bool {
	if o != nil && o.LifeCycleStates != nil {
		return true
	}

	return false
}

// SetLifeCycleStates gets a reference to the given []string and assigns it to the LifeCycleStates field.
func (o *ResourcepoolServerLeaseParametersAllOf) SetLifeCycleStates(v []string) {
	o.LifeCycleStates = v
}

func (o ResourcepoolServerLeaseParametersAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LifeCycleStates != nil {
		toSerialize["LifeCycleStates"] = o.LifeCycleStates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourcepoolServerLeaseParametersAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varResourcepoolServerLeaseParametersAllOf := _ResourcepoolServerLeaseParametersAllOf{}

	if err = json.Unmarshal(bytes, &varResourcepoolServerLeaseParametersAllOf); err == nil {
		*o = ResourcepoolServerLeaseParametersAllOf(varResourcepoolServerLeaseParametersAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LifeCycleStates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourcepoolServerLeaseParametersAllOf struct {
	value *ResourcepoolServerLeaseParametersAllOf
	isSet bool
}

func (v NullableResourcepoolServerLeaseParametersAllOf) Get() *ResourcepoolServerLeaseParametersAllOf {
	return v.value
}

func (v *NullableResourcepoolServerLeaseParametersAllOf) Set(val *ResourcepoolServerLeaseParametersAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcepoolServerLeaseParametersAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcepoolServerLeaseParametersAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcepoolServerLeaseParametersAllOf(val *ResourcepoolServerLeaseParametersAllOf) *NullableResourcepoolServerLeaseParametersAllOf {
	return &NullableResourcepoolServerLeaseParametersAllOf{value: val, isSet: true}
}

func (v NullableResourcepoolServerLeaseParametersAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcepoolServerLeaseParametersAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


