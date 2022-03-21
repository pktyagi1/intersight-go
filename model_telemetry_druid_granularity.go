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
	"fmt"
)

// TelemetryDruidGranularity - The granularity field determines how data gets bucketed across the time dimension, or how it gets aggregated by hour, day, minute, etc. It can be specified either as a string for simple granularities or as an object for arbitrary granularities. See [Granularities](https://druid.apache.org/docs/latest/querying/granularities.html).
type TelemetryDruidGranularity struct {
	TelemetryDruidDurationGranularity *TelemetryDruidDurationGranularity
	TelemetryDruidPeriodGranularity *TelemetryDruidPeriodGranularity
}

// TelemetryDruidDurationGranularityAsTelemetryDruidGranularity is a convenience function that returns TelemetryDruidDurationGranularity wrapped in TelemetryDruidGranularity
func TelemetryDruidDurationGranularityAsTelemetryDruidGranularity(v *TelemetryDruidDurationGranularity) TelemetryDruidGranularity {
	return TelemetryDruidGranularity{ TelemetryDruidDurationGranularity: v}
}

// TelemetryDruidPeriodGranularityAsTelemetryDruidGranularity is a convenience function that returns TelemetryDruidPeriodGranularity wrapped in TelemetryDruidGranularity
func TelemetryDruidPeriodGranularityAsTelemetryDruidGranularity(v *TelemetryDruidPeriodGranularity) TelemetryDruidGranularity {
	return TelemetryDruidGranularity{ TelemetryDruidPeriodGranularity: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TelemetryDruidGranularity) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'duration'
	if jsonDict["type"] == "duration" {
		// try to unmarshal JSON data into TelemetryDruidDurationGranularity
		err = json.Unmarshal(data, &dst.TelemetryDruidDurationGranularity)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidDurationGranularity, return on the first match
		} else {
			dst.TelemetryDruidDurationGranularity = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidGranularity as TelemetryDruidDurationGranularity: %s", err.Error())
		}
	}

	// check if the discriminator value is 'period'
	if jsonDict["type"] == "period" {
		// try to unmarshal JSON data into TelemetryDruidPeriodGranularity
		err = json.Unmarshal(data, &dst.TelemetryDruidPeriodGranularity)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidPeriodGranularity, return on the first match
		} else {
			dst.TelemetryDruidPeriodGranularity = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidGranularity as TelemetryDruidPeriodGranularity: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidDurationGranularity'
	if jsonDict["type"] == "telemetry.DruidDurationGranularity" {
		// try to unmarshal JSON data into TelemetryDruidDurationGranularity
		err = json.Unmarshal(data, &dst.TelemetryDruidDurationGranularity)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidDurationGranularity, return on the first match
		} else {
			dst.TelemetryDruidDurationGranularity = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidGranularity as TelemetryDruidDurationGranularity: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidPeriodGranularity'
	if jsonDict["type"] == "telemetry.DruidPeriodGranularity" {
		// try to unmarshal JSON data into TelemetryDruidPeriodGranularity
		err = json.Unmarshal(data, &dst.TelemetryDruidPeriodGranularity)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidPeriodGranularity, return on the first match
		} else {
			dst.TelemetryDruidPeriodGranularity = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidGranularity as TelemetryDruidPeriodGranularity: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TelemetryDruidGranularity) MarshalJSON() ([]byte, error) {
	if src.TelemetryDruidDurationGranularity != nil {
		return json.Marshal(&src.TelemetryDruidDurationGranularity)
	}

	if src.TelemetryDruidPeriodGranularity != nil {
		return json.Marshal(&src.TelemetryDruidPeriodGranularity)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TelemetryDruidGranularity) GetActualInstance() (interface{}) {
	if obj.TelemetryDruidDurationGranularity != nil {
		return obj.TelemetryDruidDurationGranularity
	}

	if obj.TelemetryDruidPeriodGranularity != nil {
		return obj.TelemetryDruidPeriodGranularity
	}

	// all schemas are nil
	return nil
}

type NullableTelemetryDruidGranularity struct {
	value *TelemetryDruidGranularity
	isSet bool
}

func (v NullableTelemetryDruidGranularity) Get() *TelemetryDruidGranularity {
	return v.value
}

func (v *NullableTelemetryDruidGranularity) Set(val *TelemetryDruidGranularity) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidGranularity) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidGranularity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidGranularity(val *TelemetryDruidGranularity) *NullableTelemetryDruidGranularity {
	return &NullableTelemetryDruidGranularity{value: val, isSet: true}
}

func (v NullableTelemetryDruidGranularity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidGranularity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


