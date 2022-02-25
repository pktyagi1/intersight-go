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

// TelemetryDruidRegexFilterAllOf struct for TelemetryDruidRegexFilterAllOf
type TelemetryDruidRegexFilterAllOf struct {
	// null
	Dimension string `json:"dimension"`
	// null
	Pattern string `json:"pattern"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidRegexFilterAllOf TelemetryDruidRegexFilterAllOf

// NewTelemetryDruidRegexFilterAllOf instantiates a new TelemetryDruidRegexFilterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidRegexFilterAllOf(dimension string, pattern string) *TelemetryDruidRegexFilterAllOf {
	this := TelemetryDruidRegexFilterAllOf{}
	this.Dimension = dimension
	this.Pattern = pattern
	return &this
}

// NewTelemetryDruidRegexFilterAllOfWithDefaults instantiates a new TelemetryDruidRegexFilterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidRegexFilterAllOfWithDefaults() *TelemetryDruidRegexFilterAllOf {
	this := TelemetryDruidRegexFilterAllOf{}
	return &this
}

// GetDimension returns the Dimension field value
func (o *TelemetryDruidRegexFilterAllOf) GetDimension() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidRegexFilterAllOf) GetDimensionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value
func (o *TelemetryDruidRegexFilterAllOf) SetDimension(v string) {
	o.Dimension = v
}

// GetPattern returns the Pattern field value
func (o *TelemetryDruidRegexFilterAllOf) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidRegexFilterAllOf) GetPatternOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *TelemetryDruidRegexFilterAllOf) SetPattern(v string) {
	o.Pattern = v
}

func (o TelemetryDruidRegexFilterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dimension"] = o.Dimension
	}
	if true {
		toSerialize["pattern"] = o.Pattern
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidRegexFilterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidRegexFilterAllOf := _TelemetryDruidRegexFilterAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidRegexFilterAllOf); err == nil {
		*o = TelemetryDruidRegexFilterAllOf(varTelemetryDruidRegexFilterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dimension")
		delete(additionalProperties, "pattern")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidRegexFilterAllOf struct {
	value *TelemetryDruidRegexFilterAllOf
	isSet bool
}

func (v NullableTelemetryDruidRegexFilterAllOf) Get() *TelemetryDruidRegexFilterAllOf {
	return v.value
}

func (v *NullableTelemetryDruidRegexFilterAllOf) Set(val *TelemetryDruidRegexFilterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidRegexFilterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidRegexFilterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidRegexFilterAllOf(val *TelemetryDruidRegexFilterAllOf) *NullableTelemetryDruidRegexFilterAllOf {
	return &NullableTelemetryDruidRegexFilterAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidRegexFilterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidRegexFilterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


