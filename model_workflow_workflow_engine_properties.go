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

// WorkflowWorkflowEngineProperties Properties for performing workflow action on workflow engine.
type WorkflowWorkflowEngineProperties struct {
	MoBaseComplexType
	AdditionalProperties map[string]interface{}
}

type _WorkflowWorkflowEngineProperties WorkflowWorkflowEngineProperties

// NewWorkflowWorkflowEngineProperties instantiates a new WorkflowWorkflowEngineProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWorkflowEngineProperties(classId string, objectType string) *WorkflowWorkflowEngineProperties {
	this := WorkflowWorkflowEngineProperties{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowWorkflowEnginePropertiesWithDefaults instantiates a new WorkflowWorkflowEngineProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWorkflowEnginePropertiesWithDefaults() *WorkflowWorkflowEngineProperties {
	this := WorkflowWorkflowEngineProperties{}
	return &this
}

func (o WorkflowWorkflowEngineProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWorkflowEngineProperties) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowWorkflowEnginePropertiesWithoutEmbeddedStruct struct {
	}

	varWorkflowWorkflowEnginePropertiesWithoutEmbeddedStruct := WorkflowWorkflowEnginePropertiesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowWorkflowEnginePropertiesWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowWorkflowEngineProperties := _WorkflowWorkflowEngineProperties{}
		*o = WorkflowWorkflowEngineProperties(varWorkflowWorkflowEngineProperties)
	} else {
		return err
	}

	varWorkflowWorkflowEngineProperties := _WorkflowWorkflowEngineProperties{}

	err = json.Unmarshal(bytes, &varWorkflowWorkflowEngineProperties)
	if err == nil {
		o.MoBaseComplexType = varWorkflowWorkflowEngineProperties.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

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

type NullableWorkflowWorkflowEngineProperties struct {
	value *WorkflowWorkflowEngineProperties
	isSet bool
}

func (v NullableWorkflowWorkflowEngineProperties) Get() *WorkflowWorkflowEngineProperties {
	return v.value
}

func (v *NullableWorkflowWorkflowEngineProperties) Set(val *WorkflowWorkflowEngineProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWorkflowEngineProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWorkflowEngineProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWorkflowEngineProperties(val *WorkflowWorkflowEngineProperties) *NullableWorkflowWorkflowEngineProperties {
	return &NullableWorkflowWorkflowEngineProperties{value: val, isSet: true}
}

func (v NullableWorkflowWorkflowEngineProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWorkflowEngineProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


