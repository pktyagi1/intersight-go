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

// ConvergedinfraHealthCheckDefinition Pod health check definition metadata.
type ConvergedinfraHealthCheckDefinition struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Category that the health check belongs to.
	Category *string `json:"Category,omitempty"`
	// Static information detailing the common causes for the health check failure.
	CommonCauses *string `json:"CommonCauses,omitempty"`
	// Description of the health check definition.
	Description *string `json:"Description,omitempty"`
	// Execution mode of the health check on converged infrastructure pod. * `OnDemand` - Execute the health check on-demand. * `Periodic` - Execute the health check on a periodic basis.
	ExecutionMode *string `json:"ExecutionMode,omitempty"`
	// Label for the health check definition that is displayed on UI.
	Label *string `json:"Label,omitempty"`
	// Name of the health check definition.
	Name *string `json:"Name,omitempty"`
	// Static information detailing the possible remediation actions that can be taken to remedy the health check failure.
	SuggestedResolution *string `json:"SuggestedResolution,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraHealthCheckDefinition ConvergedinfraHealthCheckDefinition

// NewConvergedinfraHealthCheckDefinition instantiates a new ConvergedinfraHealthCheckDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraHealthCheckDefinition(classId string, objectType string) *ConvergedinfraHealthCheckDefinition {
	this := ConvergedinfraHealthCheckDefinition{}
	this.ClassId = classId
	this.ObjectType = objectType
	var executionMode string = "OnDemand"
	this.ExecutionMode = &executionMode
	return &this
}

// NewConvergedinfraHealthCheckDefinitionWithDefaults instantiates a new ConvergedinfraHealthCheckDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraHealthCheckDefinitionWithDefaults() *ConvergedinfraHealthCheckDefinition {
	this := ConvergedinfraHealthCheckDefinition{}
	var classId string = "convergedinfra.HealthCheckDefinition"
	this.ClassId = classId
	var objectType string = "convergedinfra.HealthCheckDefinition"
	this.ObjectType = objectType
	var executionMode string = "OnDemand"
	this.ExecutionMode = &executionMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraHealthCheckDefinition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckDefinition) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraHealthCheckDefinition) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraHealthCheckDefinition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckDefinition) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraHealthCheckDefinition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckDefinition) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckDefinition) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckDefinition) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ConvergedinfraHealthCheckDefinition) SetCategory(v string) {
	o.Category = &v
}

// GetCommonCauses returns the CommonCauses field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckDefinition) GetCommonCauses() string {
	if o == nil || o.CommonCauses == nil {
		var ret string
		return ret
	}
	return *o.CommonCauses
}

// GetCommonCausesOk returns a tuple with the CommonCauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckDefinition) GetCommonCausesOk() (*string, bool) {
	if o == nil || o.CommonCauses == nil {
		return nil, false
	}
	return o.CommonCauses, true
}

// HasCommonCauses returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckDefinition) HasCommonCauses() bool {
	if o != nil && o.CommonCauses != nil {
		return true
	}

	return false
}

// SetCommonCauses gets a reference to the given string and assigns it to the CommonCauses field.
func (o *ConvergedinfraHealthCheckDefinition) SetCommonCauses(v string) {
	o.CommonCauses = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckDefinition) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConvergedinfraHealthCheckDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetExecutionMode returns the ExecutionMode field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckDefinition) GetExecutionMode() string {
	if o == nil || o.ExecutionMode == nil {
		var ret string
		return ret
	}
	return *o.ExecutionMode
}

// GetExecutionModeOk returns a tuple with the ExecutionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckDefinition) GetExecutionModeOk() (*string, bool) {
	if o == nil || o.ExecutionMode == nil {
		return nil, false
	}
	return o.ExecutionMode, true
}

// HasExecutionMode returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckDefinition) HasExecutionMode() bool {
	if o != nil && o.ExecutionMode != nil {
		return true
	}

	return false
}

// SetExecutionMode gets a reference to the given string and assigns it to the ExecutionMode field.
func (o *ConvergedinfraHealthCheckDefinition) SetExecutionMode(v string) {
	o.ExecutionMode = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckDefinition) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckDefinition) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckDefinition) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ConvergedinfraHealthCheckDefinition) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckDefinition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConvergedinfraHealthCheckDefinition) SetName(v string) {
	o.Name = &v
}

// GetSuggestedResolution returns the SuggestedResolution field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckDefinition) GetSuggestedResolution() string {
	if o == nil || o.SuggestedResolution == nil {
		var ret string
		return ret
	}
	return *o.SuggestedResolution
}

// GetSuggestedResolutionOk returns a tuple with the SuggestedResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckDefinition) GetSuggestedResolutionOk() (*string, bool) {
	if o == nil || o.SuggestedResolution == nil {
		return nil, false
	}
	return o.SuggestedResolution, true
}

// HasSuggestedResolution returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckDefinition) HasSuggestedResolution() bool {
	if o != nil && o.SuggestedResolution != nil {
		return true
	}

	return false
}

// SetSuggestedResolution gets a reference to the given string and assigns it to the SuggestedResolution field.
func (o *ConvergedinfraHealthCheckDefinition) SetSuggestedResolution(v string) {
	o.SuggestedResolution = &v
}

func (o ConvergedinfraHealthCheckDefinition) MarshalJSON() ([]byte, error) {
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
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.CommonCauses != nil {
		toSerialize["CommonCauses"] = o.CommonCauses
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.ExecutionMode != nil {
		toSerialize["ExecutionMode"] = o.ExecutionMode
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SuggestedResolution != nil {
		toSerialize["SuggestedResolution"] = o.SuggestedResolution
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraHealthCheckDefinition) UnmarshalJSON(bytes []byte) (err error) {
	type ConvergedinfraHealthCheckDefinitionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Category that the health check belongs to.
		Category *string `json:"Category,omitempty"`
		// Static information detailing the common causes for the health check failure.
		CommonCauses *string `json:"CommonCauses,omitempty"`
		// Description of the health check definition.
		Description *string `json:"Description,omitempty"`
		// Execution mode of the health check on converged infrastructure pod. * `OnDemand` - Execute the health check on-demand. * `Periodic` - Execute the health check on a periodic basis.
		ExecutionMode *string `json:"ExecutionMode,omitempty"`
		// Label for the health check definition that is displayed on UI.
		Label *string `json:"Label,omitempty"`
		// Name of the health check definition.
		Name *string `json:"Name,omitempty"`
		// Static information detailing the possible remediation actions that can be taken to remedy the health check failure.
		SuggestedResolution *string `json:"SuggestedResolution,omitempty"`
	}

	varConvergedinfraHealthCheckDefinitionWithoutEmbeddedStruct := ConvergedinfraHealthCheckDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConvergedinfraHealthCheckDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varConvergedinfraHealthCheckDefinition := _ConvergedinfraHealthCheckDefinition{}
		varConvergedinfraHealthCheckDefinition.ClassId = varConvergedinfraHealthCheckDefinitionWithoutEmbeddedStruct.ClassId
		varConvergedinfraHealthCheckDefinition.ObjectType = varConvergedinfraHealthCheckDefinitionWithoutEmbeddedStruct.ObjectType
		varConvergedinfraHealthCheckDefinition.Category = varConvergedinfraHealthCheckDefinitionWithoutEmbeddedStruct.Category
		varConvergedinfraHealthCheckDefinition.CommonCauses = varConvergedinfraHealthCheckDefinitionWithoutEmbeddedStruct.CommonCauses
		varConvergedinfraHealthCheckDefinition.Description = varConvergedinfraHealthCheckDefinitionWithoutEmbeddedStruct.Description
		varConvergedinfraHealthCheckDefinition.ExecutionMode = varConvergedinfraHealthCheckDefinitionWithoutEmbeddedStruct.ExecutionMode
		varConvergedinfraHealthCheckDefinition.Label = varConvergedinfraHealthCheckDefinitionWithoutEmbeddedStruct.Label
		varConvergedinfraHealthCheckDefinition.Name = varConvergedinfraHealthCheckDefinitionWithoutEmbeddedStruct.Name
		varConvergedinfraHealthCheckDefinition.SuggestedResolution = varConvergedinfraHealthCheckDefinitionWithoutEmbeddedStruct.SuggestedResolution
		*o = ConvergedinfraHealthCheckDefinition(varConvergedinfraHealthCheckDefinition)
	} else {
		return err
	}

	varConvergedinfraHealthCheckDefinition := _ConvergedinfraHealthCheckDefinition{}

	err = json.Unmarshal(bytes, &varConvergedinfraHealthCheckDefinition)
	if err == nil {
		o.MoBaseMo = varConvergedinfraHealthCheckDefinition.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "CommonCauses")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "ExecutionMode")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SuggestedResolution")

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

type NullableConvergedinfraHealthCheckDefinition struct {
	value *ConvergedinfraHealthCheckDefinition
	isSet bool
}

func (v NullableConvergedinfraHealthCheckDefinition) Get() *ConvergedinfraHealthCheckDefinition {
	return v.value
}

func (v *NullableConvergedinfraHealthCheckDefinition) Set(val *ConvergedinfraHealthCheckDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraHealthCheckDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraHealthCheckDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraHealthCheckDefinition(val *ConvergedinfraHealthCheckDefinition) *NullableConvergedinfraHealthCheckDefinition {
	return &NullableConvergedinfraHealthCheckDefinition{value: val, isSet: true}
}

func (v NullableConvergedinfraHealthCheckDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraHealthCheckDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


