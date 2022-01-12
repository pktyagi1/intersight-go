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
	"fmt"
)

// TopSystemRelationship - A relationship to the 'top.System' resource, or the expanded 'top.System' resource, or the 'null' value.
type TopSystemRelationship struct {
	MoMoRef *MoMoRef
	TopSystem *TopSystem
}

// MoMoRefAsTopSystemRelationship is a convenience function that returns MoMoRef wrapped in TopSystemRelationship
func MoMoRefAsTopSystemRelationship(v *MoMoRef) TopSystemRelationship {
	return TopSystemRelationship{ MoMoRef: v}
}

// TopSystemAsTopSystemRelationship is a convenience function that returns TopSystem wrapped in TopSystemRelationship
func TopSystemAsTopSystemRelationship(v *TopSystem) TopSystemRelationship {
	return TopSystemRelationship{ TopSystem: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TopSystemRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("Failed to unmarshal TopSystemRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'top.System'
	if jsonDict["ClassId"] == "top.System" {
		// try to unmarshal JSON data into TopSystem
		err = json.Unmarshal(data, &dst.TopSystem)
		if err == nil {
			return nil // data stored in dst.TopSystem, return on the first match
		} else {
			dst.TopSystem = nil
			return fmt.Errorf("Failed to unmarshal TopSystemRelationship as TopSystem: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TopSystemRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.TopSystem != nil {
		return json.Marshal(&src.TopSystem)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TopSystemRelationship) GetActualInstance() (interface{}) {
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.TopSystem != nil {
		return obj.TopSystem
	}

	// all schemas are nil
	return nil
}

type NullableTopSystemRelationship struct {
	value *TopSystemRelationship
	isSet bool
}

func (v NullableTopSystemRelationship) Get() *TopSystemRelationship {
	return v.value
}

func (v *NullableTopSystemRelationship) Set(val *TopSystemRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableTopSystemRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableTopSystemRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopSystemRelationship(val *TopSystemRelationship) *NullableTopSystemRelationship {
	return &NullableTopSystemRelationship{value: val, isSet: true}
}

func (v NullableTopSystemRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopSystemRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


