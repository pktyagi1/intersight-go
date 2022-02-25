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

// WorkflowWorkflowInfoProperties Properties for a workflowInfo.
type WorkflowWorkflowInfoProperties struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When true, this workflow can be retried if has not been modified for more than a period of 2 weeks.
	Retryable *bool `json:"Retryable,omitempty"`
	// Status of rollback for this workflow instance. The rollback action of the workflow can be enabled, disabled, completed. * `Disabled` - Status of the rollback action when workflow is disabled for rollback. * `Enabled` - Status of the rollback action when workflow is enabled for rollback. * `Completed` - Status of the rollback action once workflow completes the rollback for all eligible tasks.
	RollbackAction *string `json:"RollbackAction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWorkflowInfoProperties WorkflowWorkflowInfoProperties

// NewWorkflowWorkflowInfoProperties instantiates a new WorkflowWorkflowInfoProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWorkflowInfoProperties(classId string, objectType string) *WorkflowWorkflowInfoProperties {
	this := WorkflowWorkflowInfoProperties{}
	this.ClassId = classId
	this.ObjectType = objectType
	var retryable bool = false
	this.Retryable = &retryable
	return &this
}

// NewWorkflowWorkflowInfoPropertiesWithDefaults instantiates a new WorkflowWorkflowInfoProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWorkflowInfoPropertiesWithDefaults() *WorkflowWorkflowInfoProperties {
	this := WorkflowWorkflowInfoProperties{}
	var classId string = "workflow.WorkflowInfoProperties"
	this.ClassId = classId
	var objectType string = "workflow.WorkflowInfoProperties"
	this.ObjectType = objectType
	var retryable bool = false
	this.Retryable = &retryable
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowWorkflowInfoProperties) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowInfoProperties) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowWorkflowInfoProperties) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowWorkflowInfoProperties) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowInfoProperties) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowWorkflowInfoProperties) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRetryable returns the Retryable field value if set, zero value otherwise.
func (o *WorkflowWorkflowInfoProperties) GetRetryable() bool {
	if o == nil || o.Retryable == nil {
		var ret bool
		return ret
	}
	return *o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowInfoProperties) GetRetryableOk() (*bool, bool) {
	if o == nil || o.Retryable == nil {
		return nil, false
	}
	return o.Retryable, true
}

// HasRetryable returns a boolean if a field has been set.
func (o *WorkflowWorkflowInfoProperties) HasRetryable() bool {
	if o != nil && o.Retryable != nil {
		return true
	}

	return false
}

// SetRetryable gets a reference to the given bool and assigns it to the Retryable field.
func (o *WorkflowWorkflowInfoProperties) SetRetryable(v bool) {
	o.Retryable = &v
}

// GetRollbackAction returns the RollbackAction field value if set, zero value otherwise.
func (o *WorkflowWorkflowInfoProperties) GetRollbackAction() string {
	if o == nil || o.RollbackAction == nil {
		var ret string
		return ret
	}
	return *o.RollbackAction
}

// GetRollbackActionOk returns a tuple with the RollbackAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowInfoProperties) GetRollbackActionOk() (*string, bool) {
	if o == nil || o.RollbackAction == nil {
		return nil, false
	}
	return o.RollbackAction, true
}

// HasRollbackAction returns a boolean if a field has been set.
func (o *WorkflowWorkflowInfoProperties) HasRollbackAction() bool {
	if o != nil && o.RollbackAction != nil {
		return true
	}

	return false
}

// SetRollbackAction gets a reference to the given string and assigns it to the RollbackAction field.
func (o *WorkflowWorkflowInfoProperties) SetRollbackAction(v string) {
	o.RollbackAction = &v
}

func (o WorkflowWorkflowInfoProperties) MarshalJSON() ([]byte, error) {
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
	if o.Retryable != nil {
		toSerialize["Retryable"] = o.Retryable
	}
	if o.RollbackAction != nil {
		toSerialize["RollbackAction"] = o.RollbackAction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWorkflowInfoProperties) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowWorkflowInfoPropertiesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// When true, this workflow can be retried if has not been modified for more than a period of 2 weeks.
		Retryable *bool `json:"Retryable,omitempty"`
		// Status of rollback for this workflow instance. The rollback action of the workflow can be enabled, disabled, completed. * `Disabled` - Status of the rollback action when workflow is disabled for rollback. * `Enabled` - Status of the rollback action when workflow is enabled for rollback. * `Completed` - Status of the rollback action once workflow completes the rollback for all eligible tasks.
		RollbackAction *string `json:"RollbackAction,omitempty"`
	}

	varWorkflowWorkflowInfoPropertiesWithoutEmbeddedStruct := WorkflowWorkflowInfoPropertiesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowWorkflowInfoPropertiesWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowWorkflowInfoProperties := _WorkflowWorkflowInfoProperties{}
		varWorkflowWorkflowInfoProperties.ClassId = varWorkflowWorkflowInfoPropertiesWithoutEmbeddedStruct.ClassId
		varWorkflowWorkflowInfoProperties.ObjectType = varWorkflowWorkflowInfoPropertiesWithoutEmbeddedStruct.ObjectType
		varWorkflowWorkflowInfoProperties.Retryable = varWorkflowWorkflowInfoPropertiesWithoutEmbeddedStruct.Retryable
		varWorkflowWorkflowInfoProperties.RollbackAction = varWorkflowWorkflowInfoPropertiesWithoutEmbeddedStruct.RollbackAction
		*o = WorkflowWorkflowInfoProperties(varWorkflowWorkflowInfoProperties)
	} else {
		return err
	}

	varWorkflowWorkflowInfoProperties := _WorkflowWorkflowInfoProperties{}

	err = json.Unmarshal(bytes, &varWorkflowWorkflowInfoProperties)
	if err == nil {
		o.MoBaseComplexType = varWorkflowWorkflowInfoProperties.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Retryable")
		delete(additionalProperties, "RollbackAction")

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

type NullableWorkflowWorkflowInfoProperties struct {
	value *WorkflowWorkflowInfoProperties
	isSet bool
}

func (v NullableWorkflowWorkflowInfoProperties) Get() *WorkflowWorkflowInfoProperties {
	return v.value
}

func (v *NullableWorkflowWorkflowInfoProperties) Set(val *WorkflowWorkflowInfoProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWorkflowInfoProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWorkflowInfoProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWorkflowInfoProperties(val *WorkflowWorkflowInfoProperties) *NullableWorkflowWorkflowInfoProperties {
	return &NullableWorkflowWorkflowInfoProperties{value: val, isSet: true}
}

func (v NullableWorkflowWorkflowInfoProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWorkflowInfoProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


