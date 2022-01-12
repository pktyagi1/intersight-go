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

// WorkflowInternalProperties Internal properties for a task definition which are not editable by the user.
type WorkflowInternalProperties struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field will hold the base task type like HttpBaseTask or RemoteAnsibleBaseTask.
	BaseTaskType *string `json:"BaseTaskType,omitempty"`
	Constraints NullableWorkflowTaskConstraints `json:"Constraints,omitempty"`
	// Denotes this is an internal task. Internal tasks will be hidden from the UI when executing a workflow.
	Internal *bool `json:"Internal,omitempty"`
	// The service that owns and is responsible for execution of the task.
	Owner *string `json:"Owner,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowInternalProperties WorkflowInternalProperties

// NewWorkflowInternalProperties instantiates a new WorkflowInternalProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowInternalProperties(classId string, objectType string) *WorkflowInternalProperties {
	this := WorkflowInternalProperties{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowInternalPropertiesWithDefaults instantiates a new WorkflowInternalProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowInternalPropertiesWithDefaults() *WorkflowInternalProperties {
	this := WorkflowInternalProperties{}
	var classId string = "workflow.InternalProperties"
	this.ClassId = classId
	var objectType string = "workflow.InternalProperties"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowInternalProperties) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowInternalProperties) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowInternalProperties) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowInternalProperties) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowInternalProperties) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowInternalProperties) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBaseTaskType returns the BaseTaskType field value if set, zero value otherwise.
func (o *WorkflowInternalProperties) GetBaseTaskType() string {
	if o == nil || o.BaseTaskType == nil {
		var ret string
		return ret
	}
	return *o.BaseTaskType
}

// GetBaseTaskTypeOk returns a tuple with the BaseTaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInternalProperties) GetBaseTaskTypeOk() (*string, bool) {
	if o == nil || o.BaseTaskType == nil {
		return nil, false
	}
	return o.BaseTaskType, true
}

// HasBaseTaskType returns a boolean if a field has been set.
func (o *WorkflowInternalProperties) HasBaseTaskType() bool {
	if o != nil && o.BaseTaskType != nil {
		return true
	}

	return false
}

// SetBaseTaskType gets a reference to the given string and assigns it to the BaseTaskType field.
func (o *WorkflowInternalProperties) SetBaseTaskType(v string) {
	o.BaseTaskType = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowInternalProperties) GetConstraints() WorkflowTaskConstraints {
	if o == nil || o.Constraints.Get() == nil {
		var ret WorkflowTaskConstraints
		return ret
	}
	return *o.Constraints.Get()
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowInternalProperties) GetConstraintsOk() (*WorkflowTaskConstraints, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Constraints.Get(), o.Constraints.IsSet()
}

// HasConstraints returns a boolean if a field has been set.
func (o *WorkflowInternalProperties) HasConstraints() bool {
	if o != nil && o.Constraints.IsSet() {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given NullableWorkflowTaskConstraints and assigns it to the Constraints field.
func (o *WorkflowInternalProperties) SetConstraints(v WorkflowTaskConstraints) {
	o.Constraints.Set(&v)
}
// SetConstraintsNil sets the value for Constraints to be an explicit nil
func (o *WorkflowInternalProperties) SetConstraintsNil() {
	o.Constraints.Set(nil)
}

// UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
func (o *WorkflowInternalProperties) UnsetConstraints() {
	o.Constraints.Unset()
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *WorkflowInternalProperties) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInternalProperties) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *WorkflowInternalProperties) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *WorkflowInternalProperties) SetInternal(v bool) {
	o.Internal = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *WorkflowInternalProperties) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInternalProperties) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *WorkflowInternalProperties) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *WorkflowInternalProperties) SetOwner(v string) {
	o.Owner = &v
}

func (o WorkflowInternalProperties) MarshalJSON() ([]byte, error) {
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
	if o.BaseTaskType != nil {
		toSerialize["BaseTaskType"] = o.BaseTaskType
	}
	if o.Constraints.IsSet() {
		toSerialize["Constraints"] = o.Constraints.Get()
	}
	if o.Internal != nil {
		toSerialize["Internal"] = o.Internal
	}
	if o.Owner != nil {
		toSerialize["Owner"] = o.Owner
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowInternalProperties) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowInternalPropertiesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This field will hold the base task type like HttpBaseTask or RemoteAnsibleBaseTask.
		BaseTaskType *string `json:"BaseTaskType,omitempty"`
		Constraints NullableWorkflowTaskConstraints `json:"Constraints,omitempty"`
		// Denotes this is an internal task. Internal tasks will be hidden from the UI when executing a workflow.
		Internal *bool `json:"Internal,omitempty"`
		// The service that owns and is responsible for execution of the task.
		Owner *string `json:"Owner,omitempty"`
	}

	varWorkflowInternalPropertiesWithoutEmbeddedStruct := WorkflowInternalPropertiesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowInternalPropertiesWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowInternalProperties := _WorkflowInternalProperties{}
		varWorkflowInternalProperties.ClassId = varWorkflowInternalPropertiesWithoutEmbeddedStruct.ClassId
		varWorkflowInternalProperties.ObjectType = varWorkflowInternalPropertiesWithoutEmbeddedStruct.ObjectType
		varWorkflowInternalProperties.BaseTaskType = varWorkflowInternalPropertiesWithoutEmbeddedStruct.BaseTaskType
		varWorkflowInternalProperties.Constraints = varWorkflowInternalPropertiesWithoutEmbeddedStruct.Constraints
		varWorkflowInternalProperties.Internal = varWorkflowInternalPropertiesWithoutEmbeddedStruct.Internal
		varWorkflowInternalProperties.Owner = varWorkflowInternalPropertiesWithoutEmbeddedStruct.Owner
		*o = WorkflowInternalProperties(varWorkflowInternalProperties)
	} else {
		return err
	}

	varWorkflowInternalProperties := _WorkflowInternalProperties{}

	err = json.Unmarshal(bytes, &varWorkflowInternalProperties)
	if err == nil {
		o.MoBaseComplexType = varWorkflowInternalProperties.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BaseTaskType")
		delete(additionalProperties, "Constraints")
		delete(additionalProperties, "Internal")
		delete(additionalProperties, "Owner")

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

type NullableWorkflowInternalProperties struct {
	value *WorkflowInternalProperties
	isSet bool
}

func (v NullableWorkflowInternalProperties) Get() *WorkflowInternalProperties {
	return v.value
}

func (v *NullableWorkflowInternalProperties) Set(val *WorkflowInternalProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowInternalProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowInternalProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowInternalProperties(val *WorkflowInternalProperties) *NullableWorkflowInternalProperties {
	return &NullableWorkflowInternalProperties{value: val, isSet: true}
}

func (v NullableWorkflowInternalProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowInternalProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


