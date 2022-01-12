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

// NiaapiApicSweol The software end of life notice for APIC.
type NiaapiApicSweol struct {
	NiaapiSoftwareEol
	AdditionalProperties map[string]interface{}
}

type _NiaapiApicSweol NiaapiApicSweol

// NewNiaapiApicSweol instantiates a new NiaapiApicSweol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiApicSweol(classId string, objectType string) *NiaapiApicSweol {
	this := NiaapiApicSweol{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiApicSweolWithDefaults instantiates a new NiaapiApicSweol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiApicSweolWithDefaults() *NiaapiApicSweol {
	this := NiaapiApicSweol{}
	return &this
}

func (o NiaapiApicSweol) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedNiaapiSoftwareEol, errNiaapiSoftwareEol := json.Marshal(o.NiaapiSoftwareEol)
	if errNiaapiSoftwareEol != nil {
		return []byte{}, errNiaapiSoftwareEol
	}
	errNiaapiSoftwareEol = json.Unmarshal([]byte(serializedNiaapiSoftwareEol), &toSerialize)
	if errNiaapiSoftwareEol != nil {
		return []byte{}, errNiaapiSoftwareEol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiApicSweol) UnmarshalJSON(bytes []byte) (err error) {
	type NiaapiApicSweolWithoutEmbeddedStruct struct {
	}

	varNiaapiApicSweolWithoutEmbeddedStruct := NiaapiApicSweolWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiaapiApicSweolWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiApicSweol := _NiaapiApicSweol{}
		*o = NiaapiApicSweol(varNiaapiApicSweol)
	} else {
		return err
	}

	varNiaapiApicSweol := _NiaapiApicSweol{}

	err = json.Unmarshal(bytes, &varNiaapiApicSweol)
	if err == nil {
		o.NiaapiSoftwareEol = varNiaapiApicSweol.NiaapiSoftwareEol
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectNiaapiSoftwareEol := reflect.ValueOf(o.NiaapiSoftwareEol)
		for i := 0; i < reflectNiaapiSoftwareEol.Type().NumField(); i++ {
			t := reflectNiaapiSoftwareEol.Type().Field(i)

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

type NullableNiaapiApicSweol struct {
	value *NiaapiApicSweol
	isSet bool
}

func (v NullableNiaapiApicSweol) Get() *NiaapiApicSweol {
	return v.value
}

func (v *NullableNiaapiApicSweol) Set(val *NiaapiApicSweol) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiApicSweol) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiApicSweol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiApicSweol(val *NiaapiApicSweol) *NullableNiaapiApicSweol {
	return &NullableNiaapiApicSweol{value: val, isSet: true}
}

func (v NullableNiaapiApicSweol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiApicSweol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

