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

// TelemetryDruidColumnComparisonFilterAllOf struct for TelemetryDruidColumnComparisonFilterAllOf
type TelemetryDruidColumnComparisonFilterAllOf struct {
	// A list of DimensionSpecs, making it possible to apply an extraction function if needed.
	Dimensions []TelemetryDruidDimensionSpec `json:"dimensions"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidColumnComparisonFilterAllOf TelemetryDruidColumnComparisonFilterAllOf

// NewTelemetryDruidColumnComparisonFilterAllOf instantiates a new TelemetryDruidColumnComparisonFilterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidColumnComparisonFilterAllOf(dimensions []TelemetryDruidDimensionSpec) *TelemetryDruidColumnComparisonFilterAllOf {
	this := TelemetryDruidColumnComparisonFilterAllOf{}
	this.Dimensions = dimensions
	return &this
}

// NewTelemetryDruidColumnComparisonFilterAllOfWithDefaults instantiates a new TelemetryDruidColumnComparisonFilterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidColumnComparisonFilterAllOfWithDefaults() *TelemetryDruidColumnComparisonFilterAllOf {
	this := TelemetryDruidColumnComparisonFilterAllOf{}
	return &this
}

// GetDimensions returns the Dimensions field value
func (o *TelemetryDruidColumnComparisonFilterAllOf) GetDimensions() []TelemetryDruidDimensionSpec {
	if o == nil {
		var ret []TelemetryDruidDimensionSpec
		return ret
	}

	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidColumnComparisonFilterAllOf) GetDimensionsOk() (*[]TelemetryDruidDimensionSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value
func (o *TelemetryDruidColumnComparisonFilterAllOf) SetDimensions(v []TelemetryDruidDimensionSpec) {
	o.Dimensions = v
}

func (o TelemetryDruidColumnComparisonFilterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dimensions"] = o.Dimensions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidColumnComparisonFilterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidColumnComparisonFilterAllOf := _TelemetryDruidColumnComparisonFilterAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidColumnComparisonFilterAllOf); err == nil {
		*o = TelemetryDruidColumnComparisonFilterAllOf(varTelemetryDruidColumnComparisonFilterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dimensions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidColumnComparisonFilterAllOf struct {
	value *TelemetryDruidColumnComparisonFilterAllOf
	isSet bool
}

func (v NullableTelemetryDruidColumnComparisonFilterAllOf) Get() *TelemetryDruidColumnComparisonFilterAllOf {
	return v.value
}

func (v *NullableTelemetryDruidColumnComparisonFilterAllOf) Set(val *TelemetryDruidColumnComparisonFilterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidColumnComparisonFilterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidColumnComparisonFilterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidColumnComparisonFilterAllOf(val *TelemetryDruidColumnComparisonFilterAllOf) *NullableTelemetryDruidColumnComparisonFilterAllOf {
	return &NullableTelemetryDruidColumnComparisonFilterAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidColumnComparisonFilterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidColumnComparisonFilterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


