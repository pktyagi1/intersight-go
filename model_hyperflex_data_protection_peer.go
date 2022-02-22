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

// HyperflexDataProtectionPeer Data Protection Peer describing the cluster pair from the HyperFlex backend.
type HyperflexDataProtectionPeer struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	Er NullableHyperflexEntityReference `json:"Er,omitempty"`
	PeerInfo NullableHyperflexReplicationPeerInfo `json:"PeerInfo,omitempty"`
	SrcCluster *HyperflexClusterRelationship `json:"SrcCluster,omitempty"`
	TgtCluster *HyperflexClusterRelationship `json:"TgtCluster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexDataProtectionPeer HyperflexDataProtectionPeer

// NewHyperflexDataProtectionPeer instantiates a new HyperflexDataProtectionPeer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexDataProtectionPeer(classId string, objectType string) *HyperflexDataProtectionPeer {
	this := HyperflexDataProtectionPeer{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexDataProtectionPeerWithDefaults instantiates a new HyperflexDataProtectionPeer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexDataProtectionPeerWithDefaults() *HyperflexDataProtectionPeer {
	this := HyperflexDataProtectionPeer{}
	var classId string = "hyperflex.DataProtectionPeer"
	this.ClassId = classId
	var objectType string = "hyperflex.DataProtectionPeer"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexDataProtectionPeer) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexDataProtectionPeer) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexDataProtectionPeer) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexDataProtectionPeer) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexDataProtectionPeer) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexDataProtectionPeer) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEr returns the Er field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexDataProtectionPeer) GetEr() HyperflexEntityReference {
	if o == nil || o.Er.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.Er.Get()
}

// GetErOk returns a tuple with the Er field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexDataProtectionPeer) GetErOk() (*HyperflexEntityReference, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Er.Get(), o.Er.IsSet()
}

// HasEr returns a boolean if a field has been set.
func (o *HyperflexDataProtectionPeer) HasEr() bool {
	if o != nil && o.Er.IsSet() {
		return true
	}

	return false
}

// SetEr gets a reference to the given NullableHyperflexEntityReference and assigns it to the Er field.
func (o *HyperflexDataProtectionPeer) SetEr(v HyperflexEntityReference) {
	o.Er.Set(&v)
}
// SetErNil sets the value for Er to be an explicit nil
func (o *HyperflexDataProtectionPeer) SetErNil() {
	o.Er.Set(nil)
}

// UnsetEr ensures that no value is present for Er, not even an explicit nil
func (o *HyperflexDataProtectionPeer) UnsetEr() {
	o.Er.Unset()
}

// GetPeerInfo returns the PeerInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexDataProtectionPeer) GetPeerInfo() HyperflexReplicationPeerInfo {
	if o == nil || o.PeerInfo.Get() == nil {
		var ret HyperflexReplicationPeerInfo
		return ret
	}
	return *o.PeerInfo.Get()
}

// GetPeerInfoOk returns a tuple with the PeerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexDataProtectionPeer) GetPeerInfoOk() (*HyperflexReplicationPeerInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PeerInfo.Get(), o.PeerInfo.IsSet()
}

// HasPeerInfo returns a boolean if a field has been set.
func (o *HyperflexDataProtectionPeer) HasPeerInfo() bool {
	if o != nil && o.PeerInfo.IsSet() {
		return true
	}

	return false
}

// SetPeerInfo gets a reference to the given NullableHyperflexReplicationPeerInfo and assigns it to the PeerInfo field.
func (o *HyperflexDataProtectionPeer) SetPeerInfo(v HyperflexReplicationPeerInfo) {
	o.PeerInfo.Set(&v)
}
// SetPeerInfoNil sets the value for PeerInfo to be an explicit nil
func (o *HyperflexDataProtectionPeer) SetPeerInfoNil() {
	o.PeerInfo.Set(nil)
}

// UnsetPeerInfo ensures that no value is present for PeerInfo, not even an explicit nil
func (o *HyperflexDataProtectionPeer) UnsetPeerInfo() {
	o.PeerInfo.Unset()
}

// GetSrcCluster returns the SrcCluster field value if set, zero value otherwise.
func (o *HyperflexDataProtectionPeer) GetSrcCluster() HyperflexClusterRelationship {
	if o == nil || o.SrcCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.SrcCluster
}

// GetSrcClusterOk returns a tuple with the SrcCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDataProtectionPeer) GetSrcClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.SrcCluster == nil {
		return nil, false
	}
	return o.SrcCluster, true
}

// HasSrcCluster returns a boolean if a field has been set.
func (o *HyperflexDataProtectionPeer) HasSrcCluster() bool {
	if o != nil && o.SrcCluster != nil {
		return true
	}

	return false
}

// SetSrcCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the SrcCluster field.
func (o *HyperflexDataProtectionPeer) SetSrcCluster(v HyperflexClusterRelationship) {
	o.SrcCluster = &v
}

// GetTgtCluster returns the TgtCluster field value if set, zero value otherwise.
func (o *HyperflexDataProtectionPeer) GetTgtCluster() HyperflexClusterRelationship {
	if o == nil || o.TgtCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.TgtCluster
}

// GetTgtClusterOk returns a tuple with the TgtCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDataProtectionPeer) GetTgtClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.TgtCluster == nil {
		return nil, false
	}
	return o.TgtCluster, true
}

// HasTgtCluster returns a boolean if a field has been set.
func (o *HyperflexDataProtectionPeer) HasTgtCluster() bool {
	if o != nil && o.TgtCluster != nil {
		return true
	}

	return false
}

// SetTgtCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the TgtCluster field.
func (o *HyperflexDataProtectionPeer) SetTgtCluster(v HyperflexClusterRelationship) {
	o.TgtCluster = &v
}

func (o HyperflexDataProtectionPeer) MarshalJSON() ([]byte, error) {
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
	if o.Er.IsSet() {
		toSerialize["Er"] = o.Er.Get()
	}
	if o.PeerInfo.IsSet() {
		toSerialize["PeerInfo"] = o.PeerInfo.Get()
	}
	if o.SrcCluster != nil {
		toSerialize["SrcCluster"] = o.SrcCluster
	}
	if o.TgtCluster != nil {
		toSerialize["TgtCluster"] = o.TgtCluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexDataProtectionPeer) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexDataProtectionPeerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		Er NullableHyperflexEntityReference `json:"Er,omitempty"`
		PeerInfo NullableHyperflexReplicationPeerInfo `json:"PeerInfo,omitempty"`
		SrcCluster *HyperflexClusterRelationship `json:"SrcCluster,omitempty"`
		TgtCluster *HyperflexClusterRelationship `json:"TgtCluster,omitempty"`
	}

	varHyperflexDataProtectionPeerWithoutEmbeddedStruct := HyperflexDataProtectionPeerWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexDataProtectionPeerWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexDataProtectionPeer := _HyperflexDataProtectionPeer{}
		varHyperflexDataProtectionPeer.ClassId = varHyperflexDataProtectionPeerWithoutEmbeddedStruct.ClassId
		varHyperflexDataProtectionPeer.ObjectType = varHyperflexDataProtectionPeerWithoutEmbeddedStruct.ObjectType
		varHyperflexDataProtectionPeer.Er = varHyperflexDataProtectionPeerWithoutEmbeddedStruct.Er
		varHyperflexDataProtectionPeer.PeerInfo = varHyperflexDataProtectionPeerWithoutEmbeddedStruct.PeerInfo
		varHyperflexDataProtectionPeer.SrcCluster = varHyperflexDataProtectionPeerWithoutEmbeddedStruct.SrcCluster
		varHyperflexDataProtectionPeer.TgtCluster = varHyperflexDataProtectionPeerWithoutEmbeddedStruct.TgtCluster
		*o = HyperflexDataProtectionPeer(varHyperflexDataProtectionPeer)
	} else {
		return err
	}

	varHyperflexDataProtectionPeer := _HyperflexDataProtectionPeer{}

	err = json.Unmarshal(bytes, &varHyperflexDataProtectionPeer)
	if err == nil {
		o.MoBaseMo = varHyperflexDataProtectionPeer.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Er")
		delete(additionalProperties, "PeerInfo")
		delete(additionalProperties, "SrcCluster")
		delete(additionalProperties, "TgtCluster")

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

type NullableHyperflexDataProtectionPeer struct {
	value *HyperflexDataProtectionPeer
	isSet bool
}

func (v NullableHyperflexDataProtectionPeer) Get() *HyperflexDataProtectionPeer {
	return v.value
}

func (v *NullableHyperflexDataProtectionPeer) Set(val *HyperflexDataProtectionPeer) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexDataProtectionPeer) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexDataProtectionPeer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexDataProtectionPeer(val *HyperflexDataProtectionPeer) *NullableHyperflexDataProtectionPeer {
	return &NullableHyperflexDataProtectionPeer{value: val, isSet: true}
}

func (v NullableHyperflexDataProtectionPeer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexDataProtectionPeer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


