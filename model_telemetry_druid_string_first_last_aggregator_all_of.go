/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TelemetryDruidStringFirstLastAggregatorAllOf struct for TelemetryDruidStringFirstLastAggregatorAllOf
type TelemetryDruidStringFirstLastAggregatorAllOf struct {
	// Output name for the minimum/maximum timestamp value.
	Name string `json:"name"`
	// Name of the metric column.
	FieldName string `json:"fieldName"`
	// null
	MaxStringBytes *int32 `json:"maxStringBytes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidStringFirstLastAggregatorAllOf TelemetryDruidStringFirstLastAggregatorAllOf

// NewTelemetryDruidStringFirstLastAggregatorAllOf instantiates a new TelemetryDruidStringFirstLastAggregatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidStringFirstLastAggregatorAllOf(name string, fieldName string) *TelemetryDruidStringFirstLastAggregatorAllOf {
	this := TelemetryDruidStringFirstLastAggregatorAllOf{}
	this.Name = name
	this.FieldName = fieldName
	var maxStringBytes int32 = 1024
	this.MaxStringBytes = &maxStringBytes
	return &this
}

// NewTelemetryDruidStringFirstLastAggregatorAllOfWithDefaults instantiates a new TelemetryDruidStringFirstLastAggregatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidStringFirstLastAggregatorAllOfWithDefaults() *TelemetryDruidStringFirstLastAggregatorAllOf {
	this := TelemetryDruidStringFirstLastAggregatorAllOf{}
	var maxStringBytes int32 = 1024
	this.MaxStringBytes = &maxStringBytes
	return &this
}

// GetName returns the Name field value
func (o *TelemetryDruidStringFirstLastAggregatorAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidStringFirstLastAggregatorAllOf) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TelemetryDruidStringFirstLastAggregatorAllOf) SetName(v string) {
	o.Name = v
}

// GetFieldName returns the FieldName field value
func (o *TelemetryDruidStringFirstLastAggregatorAllOf) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidStringFirstLastAggregatorAllOf) GetFieldNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *TelemetryDruidStringFirstLastAggregatorAllOf) SetFieldName(v string) {
	o.FieldName = v
}

// GetMaxStringBytes returns the MaxStringBytes field value if set, zero value otherwise.
func (o *TelemetryDruidStringFirstLastAggregatorAllOf) GetMaxStringBytes() int32 {
	if o == nil || o.MaxStringBytes == nil {
		var ret int32
		return ret
	}
	return *o.MaxStringBytes
}

// GetMaxStringBytesOk returns a tuple with the MaxStringBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidStringFirstLastAggregatorAllOf) GetMaxStringBytesOk() (*int32, bool) {
	if o == nil || o.MaxStringBytes == nil {
		return nil, false
	}
	return o.MaxStringBytes, true
}

// HasMaxStringBytes returns a boolean if a field has been set.
func (o *TelemetryDruidStringFirstLastAggregatorAllOf) HasMaxStringBytes() bool {
	if o != nil && o.MaxStringBytes != nil {
		return true
	}

	return false
}

// SetMaxStringBytes gets a reference to the given int32 and assigns it to the MaxStringBytes field.
func (o *TelemetryDruidStringFirstLastAggregatorAllOf) SetMaxStringBytes(v int32) {
	o.MaxStringBytes = &v
}

func (o TelemetryDruidStringFirstLastAggregatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["fieldName"] = o.FieldName
	}
	if o.MaxStringBytes != nil {
		toSerialize["maxStringBytes"] = o.MaxStringBytes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidStringFirstLastAggregatorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidStringFirstLastAggregatorAllOf := _TelemetryDruidStringFirstLastAggregatorAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidStringFirstLastAggregatorAllOf); err == nil {
		*o = TelemetryDruidStringFirstLastAggregatorAllOf(varTelemetryDruidStringFirstLastAggregatorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "fieldName")
		delete(additionalProperties, "maxStringBytes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidStringFirstLastAggregatorAllOf struct {
	value *TelemetryDruidStringFirstLastAggregatorAllOf
	isSet bool
}

func (v NullableTelemetryDruidStringFirstLastAggregatorAllOf) Get() *TelemetryDruidStringFirstLastAggregatorAllOf {
	return v.value
}

func (v *NullableTelemetryDruidStringFirstLastAggregatorAllOf) Set(val *TelemetryDruidStringFirstLastAggregatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidStringFirstLastAggregatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidStringFirstLastAggregatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidStringFirstLastAggregatorAllOf(val *TelemetryDruidStringFirstLastAggregatorAllOf) *NullableTelemetryDruidStringFirstLastAggregatorAllOf {
	return &NullableTelemetryDruidStringFirstLastAggregatorAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidStringFirstLastAggregatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidStringFirstLastAggregatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


