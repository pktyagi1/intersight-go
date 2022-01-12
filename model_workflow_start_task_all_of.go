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

// WorkflowStartTaskAllOf Definition of the list of properties defined in 'workflow.StartTask', excluding properties defined in parent classes.
type WorkflowStartTaskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the next task (Task names unique within workflow) to run.  In a graph model, denotes an edge to another Task Node.
	NextTask *string `json:"NextTask,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowStartTaskAllOf WorkflowStartTaskAllOf

// NewWorkflowStartTaskAllOf instantiates a new WorkflowStartTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStartTaskAllOf(classId string, objectType string) *WorkflowStartTaskAllOf {
	this := WorkflowStartTaskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowStartTaskAllOfWithDefaults instantiates a new WorkflowStartTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStartTaskAllOfWithDefaults() *WorkflowStartTaskAllOf {
	this := WorkflowStartTaskAllOf{}
	var classId string = "workflow.StartTask"
	this.ClassId = classId
	var objectType string = "workflow.StartTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowStartTaskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowStartTaskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowStartTaskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowStartTaskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowStartTaskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowStartTaskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNextTask returns the NextTask field value if set, zero value otherwise.
func (o *WorkflowStartTaskAllOf) GetNextTask() string {
	if o == nil || o.NextTask == nil {
		var ret string
		return ret
	}
	return *o.NextTask
}

// GetNextTaskOk returns a tuple with the NextTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStartTaskAllOf) GetNextTaskOk() (*string, bool) {
	if o == nil || o.NextTask == nil {
		return nil, false
	}
	return o.NextTask, true
}

// HasNextTask returns a boolean if a field has been set.
func (o *WorkflowStartTaskAllOf) HasNextTask() bool {
	if o != nil && o.NextTask != nil {
		return true
	}

	return false
}

// SetNextTask gets a reference to the given string and assigns it to the NextTask field.
func (o *WorkflowStartTaskAllOf) SetNextTask(v string) {
	o.NextTask = &v
}

func (o WorkflowStartTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.NextTask != nil {
		toSerialize["NextTask"] = o.NextTask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowStartTaskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowStartTaskAllOf := _WorkflowStartTaskAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowStartTaskAllOf); err == nil {
		*o = WorkflowStartTaskAllOf(varWorkflowStartTaskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NextTask")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowStartTaskAllOf struct {
	value *WorkflowStartTaskAllOf
	isSet bool
}

func (v NullableWorkflowStartTaskAllOf) Get() *WorkflowStartTaskAllOf {
	return v.value
}

func (v *NullableWorkflowStartTaskAllOf) Set(val *WorkflowStartTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStartTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStartTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStartTaskAllOf(val *WorkflowStartTaskAllOf) *NullableWorkflowStartTaskAllOf {
	return &NullableWorkflowStartTaskAllOf{value: val, isSet: true}
}

func (v NullableWorkflowStartTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStartTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

