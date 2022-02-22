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

// ConnectorExpectPrompt This models a single expect and answer prompt of the interactive command.
type ConnectorExpectPrompt struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The regex of the expect prompt of the interactive command.
	Expect *string `json:"Expect,omitempty"`
	// The timeout for the expect prompt while executing interactive command. If timeout is not set a default of 60 seconds will be used.
	ExpectTimeout *int64 `json:"ExpectTimeout,omitempty"`
	// The answer string to the expect prompt.
	Send *string `json:"Send,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorExpectPrompt ConnectorExpectPrompt

// NewConnectorExpectPrompt instantiates a new ConnectorExpectPrompt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorExpectPrompt(classId string, objectType string) *ConnectorExpectPrompt {
	this := ConnectorExpectPrompt{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorExpectPromptWithDefaults instantiates a new ConnectorExpectPrompt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorExpectPromptWithDefaults() *ConnectorExpectPrompt {
	this := ConnectorExpectPrompt{}
	var classId string = "connector.ExpectPrompt"
	this.ClassId = classId
	var objectType string = "connector.ExpectPrompt"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorExpectPrompt) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorExpectPrompt) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorExpectPrompt) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorExpectPrompt) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorExpectPrompt) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorExpectPrompt) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExpect returns the Expect field value if set, zero value otherwise.
func (o *ConnectorExpectPrompt) GetExpect() string {
	if o == nil || o.Expect == nil {
		var ret string
		return ret
	}
	return *o.Expect
}

// GetExpectOk returns a tuple with the Expect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorExpectPrompt) GetExpectOk() (*string, bool) {
	if o == nil || o.Expect == nil {
		return nil, false
	}
	return o.Expect, true
}

// HasExpect returns a boolean if a field has been set.
func (o *ConnectorExpectPrompt) HasExpect() bool {
	if o != nil && o.Expect != nil {
		return true
	}

	return false
}

// SetExpect gets a reference to the given string and assigns it to the Expect field.
func (o *ConnectorExpectPrompt) SetExpect(v string) {
	o.Expect = &v
}

// GetExpectTimeout returns the ExpectTimeout field value if set, zero value otherwise.
func (o *ConnectorExpectPrompt) GetExpectTimeout() int64 {
	if o == nil || o.ExpectTimeout == nil {
		var ret int64
		return ret
	}
	return *o.ExpectTimeout
}

// GetExpectTimeoutOk returns a tuple with the ExpectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorExpectPrompt) GetExpectTimeoutOk() (*int64, bool) {
	if o == nil || o.ExpectTimeout == nil {
		return nil, false
	}
	return o.ExpectTimeout, true
}

// HasExpectTimeout returns a boolean if a field has been set.
func (o *ConnectorExpectPrompt) HasExpectTimeout() bool {
	if o != nil && o.ExpectTimeout != nil {
		return true
	}

	return false
}

// SetExpectTimeout gets a reference to the given int64 and assigns it to the ExpectTimeout field.
func (o *ConnectorExpectPrompt) SetExpectTimeout(v int64) {
	o.ExpectTimeout = &v
}

// GetSend returns the Send field value if set, zero value otherwise.
func (o *ConnectorExpectPrompt) GetSend() string {
	if o == nil || o.Send == nil {
		var ret string
		return ret
	}
	return *o.Send
}

// GetSendOk returns a tuple with the Send field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorExpectPrompt) GetSendOk() (*string, bool) {
	if o == nil || o.Send == nil {
		return nil, false
	}
	return o.Send, true
}

// HasSend returns a boolean if a field has been set.
func (o *ConnectorExpectPrompt) HasSend() bool {
	if o != nil && o.Send != nil {
		return true
	}

	return false
}

// SetSend gets a reference to the given string and assigns it to the Send field.
func (o *ConnectorExpectPrompt) SetSend(v string) {
	o.Send = &v
}

func (o ConnectorExpectPrompt) MarshalJSON() ([]byte, error) {
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
	if o.Expect != nil {
		toSerialize["Expect"] = o.Expect
	}
	if o.ExpectTimeout != nil {
		toSerialize["ExpectTimeout"] = o.ExpectTimeout
	}
	if o.Send != nil {
		toSerialize["Send"] = o.Send
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorExpectPrompt) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorExpectPromptWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The regex of the expect prompt of the interactive command.
		Expect *string `json:"Expect,omitempty"`
		// The timeout for the expect prompt while executing interactive command. If timeout is not set a default of 60 seconds will be used.
		ExpectTimeout *int64 `json:"ExpectTimeout,omitempty"`
		// The answer string to the expect prompt.
		Send *string `json:"Send,omitempty"`
	}

	varConnectorExpectPromptWithoutEmbeddedStruct := ConnectorExpectPromptWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorExpectPromptWithoutEmbeddedStruct)
	if err == nil {
		varConnectorExpectPrompt := _ConnectorExpectPrompt{}
		varConnectorExpectPrompt.ClassId = varConnectorExpectPromptWithoutEmbeddedStruct.ClassId
		varConnectorExpectPrompt.ObjectType = varConnectorExpectPromptWithoutEmbeddedStruct.ObjectType
		varConnectorExpectPrompt.Expect = varConnectorExpectPromptWithoutEmbeddedStruct.Expect
		varConnectorExpectPrompt.ExpectTimeout = varConnectorExpectPromptWithoutEmbeddedStruct.ExpectTimeout
		varConnectorExpectPrompt.Send = varConnectorExpectPromptWithoutEmbeddedStruct.Send
		*o = ConnectorExpectPrompt(varConnectorExpectPrompt)
	} else {
		return err
	}

	varConnectorExpectPrompt := _ConnectorExpectPrompt{}

	err = json.Unmarshal(bytes, &varConnectorExpectPrompt)
	if err == nil {
		o.MoBaseComplexType = varConnectorExpectPrompt.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Expect")
		delete(additionalProperties, "ExpectTimeout")
		delete(additionalProperties, "Send")

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

type NullableConnectorExpectPrompt struct {
	value *ConnectorExpectPrompt
	isSet bool
}

func (v NullableConnectorExpectPrompt) Get() *ConnectorExpectPrompt {
	return v.value
}

func (v *NullableConnectorExpectPrompt) Set(val *ConnectorExpectPrompt) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorExpectPrompt) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorExpectPrompt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorExpectPrompt(val *ConnectorExpectPrompt) *NullableConnectorExpectPrompt {
	return &NullableConnectorExpectPrompt{value: val, isSet: true}
}

func (v NullableConnectorExpectPrompt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorExpectPrompt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


