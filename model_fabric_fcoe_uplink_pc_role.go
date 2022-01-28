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
	"reflect"
	"strings"
)

// FabricFcoeUplinkPcRole Object sent by user to configure a fcoe uplink port-channel on the collection of ports.
type FabricFcoeUplinkPcRole struct {
	FabricPortChannelRole
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin configured speed for the port. * `Auto` - Admin configurable speed AUTO ( default ). * `1Gbps` - Admin configurable speed 1Gbps. * `10Gbps` - Admin configurable speed 10Gbps. * `25Gbps` - Admin configurable speed 25Gbps. * `40Gbps` - Admin configurable speed 40Gbps. * `100Gbps` - Admin configurable speed 100Gbps.
	AdminSpeed *string `json:"AdminSpeed,omitempty"`
	LinkAggregationPolicy *FabricLinkAggregationPolicyRelationship `json:"LinkAggregationPolicy,omitempty"`
	LinkControlPolicy *FabricLinkControlPolicyRelationship `json:"LinkControlPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricFcoeUplinkPcRole FabricFcoeUplinkPcRole

// NewFabricFcoeUplinkPcRole instantiates a new FabricFcoeUplinkPcRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricFcoeUplinkPcRole(classId string, objectType string) *FabricFcoeUplinkPcRole {
	this := FabricFcoeUplinkPcRole{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	return &this
}

// NewFabricFcoeUplinkPcRoleWithDefaults instantiates a new FabricFcoeUplinkPcRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricFcoeUplinkPcRoleWithDefaults() *FabricFcoeUplinkPcRole {
	this := FabricFcoeUplinkPcRole{}
	var classId string = "fabric.FcoeUplinkPcRole"
	this.ClassId = classId
	var objectType string = "fabric.FcoeUplinkPcRole"
	this.ObjectType = objectType
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricFcoeUplinkPcRole) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricFcoeUplinkPcRole) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricFcoeUplinkPcRole) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricFcoeUplinkPcRole) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricFcoeUplinkPcRole) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricFcoeUplinkPcRole) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminSpeed returns the AdminSpeed field value if set, zero value otherwise.
func (o *FabricFcoeUplinkPcRole) GetAdminSpeed() string {
	if o == nil || o.AdminSpeed == nil {
		var ret string
		return ret
	}
	return *o.AdminSpeed
}

// GetAdminSpeedOk returns a tuple with the AdminSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcoeUplinkPcRole) GetAdminSpeedOk() (*string, bool) {
	if o == nil || o.AdminSpeed == nil {
		return nil, false
	}
	return o.AdminSpeed, true
}

// HasAdminSpeed returns a boolean if a field has been set.
func (o *FabricFcoeUplinkPcRole) HasAdminSpeed() bool {
	if o != nil && o.AdminSpeed != nil {
		return true
	}

	return false
}

// SetAdminSpeed gets a reference to the given string and assigns it to the AdminSpeed field.
func (o *FabricFcoeUplinkPcRole) SetAdminSpeed(v string) {
	o.AdminSpeed = &v
}

// GetLinkAggregationPolicy returns the LinkAggregationPolicy field value if set, zero value otherwise.
func (o *FabricFcoeUplinkPcRole) GetLinkAggregationPolicy() FabricLinkAggregationPolicyRelationship {
	if o == nil || o.LinkAggregationPolicy == nil {
		var ret FabricLinkAggregationPolicyRelationship
		return ret
	}
	return *o.LinkAggregationPolicy
}

// GetLinkAggregationPolicyOk returns a tuple with the LinkAggregationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcoeUplinkPcRole) GetLinkAggregationPolicyOk() (*FabricLinkAggregationPolicyRelationship, bool) {
	if o == nil || o.LinkAggregationPolicy == nil {
		return nil, false
	}
	return o.LinkAggregationPolicy, true
}

// HasLinkAggregationPolicy returns a boolean if a field has been set.
func (o *FabricFcoeUplinkPcRole) HasLinkAggregationPolicy() bool {
	if o != nil && o.LinkAggregationPolicy != nil {
		return true
	}

	return false
}

// SetLinkAggregationPolicy gets a reference to the given FabricLinkAggregationPolicyRelationship and assigns it to the LinkAggregationPolicy field.
func (o *FabricFcoeUplinkPcRole) SetLinkAggregationPolicy(v FabricLinkAggregationPolicyRelationship) {
	o.LinkAggregationPolicy = &v
}

// GetLinkControlPolicy returns the LinkControlPolicy field value if set, zero value otherwise.
func (o *FabricFcoeUplinkPcRole) GetLinkControlPolicy() FabricLinkControlPolicyRelationship {
	if o == nil || o.LinkControlPolicy == nil {
		var ret FabricLinkControlPolicyRelationship
		return ret
	}
	return *o.LinkControlPolicy
}

// GetLinkControlPolicyOk returns a tuple with the LinkControlPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcoeUplinkPcRole) GetLinkControlPolicyOk() (*FabricLinkControlPolicyRelationship, bool) {
	if o == nil || o.LinkControlPolicy == nil {
		return nil, false
	}
	return o.LinkControlPolicy, true
}

// HasLinkControlPolicy returns a boolean if a field has been set.
func (o *FabricFcoeUplinkPcRole) HasLinkControlPolicy() bool {
	if o != nil && o.LinkControlPolicy != nil {
		return true
	}

	return false
}

// SetLinkControlPolicy gets a reference to the given FabricLinkControlPolicyRelationship and assigns it to the LinkControlPolicy field.
func (o *FabricFcoeUplinkPcRole) SetLinkControlPolicy(v FabricLinkControlPolicyRelationship) {
	o.LinkControlPolicy = &v
}

func (o FabricFcoeUplinkPcRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricPortChannelRole, errFabricPortChannelRole := json.Marshal(o.FabricPortChannelRole)
	if errFabricPortChannelRole != nil {
		return []byte{}, errFabricPortChannelRole
	}
	errFabricPortChannelRole = json.Unmarshal([]byte(serializedFabricPortChannelRole), &toSerialize)
	if errFabricPortChannelRole != nil {
		return []byte{}, errFabricPortChannelRole
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminSpeed != nil {
		toSerialize["AdminSpeed"] = o.AdminSpeed
	}
	if o.LinkAggregationPolicy != nil {
		toSerialize["LinkAggregationPolicy"] = o.LinkAggregationPolicy
	}
	if o.LinkControlPolicy != nil {
		toSerialize["LinkControlPolicy"] = o.LinkControlPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricFcoeUplinkPcRole) UnmarshalJSON(bytes []byte) (err error) {
	type FabricFcoeUplinkPcRoleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin configured speed for the port. * `Auto` - Admin configurable speed AUTO ( default ). * `1Gbps` - Admin configurable speed 1Gbps. * `10Gbps` - Admin configurable speed 10Gbps. * `25Gbps` - Admin configurable speed 25Gbps. * `40Gbps` - Admin configurable speed 40Gbps. * `100Gbps` - Admin configurable speed 100Gbps.
		AdminSpeed *string `json:"AdminSpeed,omitempty"`
		LinkAggregationPolicy *FabricLinkAggregationPolicyRelationship `json:"LinkAggregationPolicy,omitempty"`
		LinkControlPolicy *FabricLinkControlPolicyRelationship `json:"LinkControlPolicy,omitempty"`
	}

	varFabricFcoeUplinkPcRoleWithoutEmbeddedStruct := FabricFcoeUplinkPcRoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricFcoeUplinkPcRoleWithoutEmbeddedStruct)
	if err == nil {
		varFabricFcoeUplinkPcRole := _FabricFcoeUplinkPcRole{}
		varFabricFcoeUplinkPcRole.ClassId = varFabricFcoeUplinkPcRoleWithoutEmbeddedStruct.ClassId
		varFabricFcoeUplinkPcRole.ObjectType = varFabricFcoeUplinkPcRoleWithoutEmbeddedStruct.ObjectType
		varFabricFcoeUplinkPcRole.AdminSpeed = varFabricFcoeUplinkPcRoleWithoutEmbeddedStruct.AdminSpeed
		varFabricFcoeUplinkPcRole.LinkAggregationPolicy = varFabricFcoeUplinkPcRoleWithoutEmbeddedStruct.LinkAggregationPolicy
		varFabricFcoeUplinkPcRole.LinkControlPolicy = varFabricFcoeUplinkPcRoleWithoutEmbeddedStruct.LinkControlPolicy
		*o = FabricFcoeUplinkPcRole(varFabricFcoeUplinkPcRole)
	} else {
		return err
	}

	varFabricFcoeUplinkPcRole := _FabricFcoeUplinkPcRole{}

	err = json.Unmarshal(bytes, &varFabricFcoeUplinkPcRole)
	if err == nil {
		o.FabricPortChannelRole = varFabricFcoeUplinkPcRole.FabricPortChannelRole
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSpeed")
		delete(additionalProperties, "LinkAggregationPolicy")
		delete(additionalProperties, "LinkControlPolicy")

		// remove fields from embedded structs
		reflectFabricPortChannelRole := reflect.ValueOf(o.FabricPortChannelRole)
		for i := 0; i < reflectFabricPortChannelRole.Type().NumField(); i++ {
			t := reflectFabricPortChannelRole.Type().Field(i)

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

type NullableFabricFcoeUplinkPcRole struct {
	value *FabricFcoeUplinkPcRole
	isSet bool
}

func (v NullableFabricFcoeUplinkPcRole) Get() *FabricFcoeUplinkPcRole {
	return v.value
}

func (v *NullableFabricFcoeUplinkPcRole) Set(val *FabricFcoeUplinkPcRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricFcoeUplinkPcRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricFcoeUplinkPcRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricFcoeUplinkPcRole(val *FabricFcoeUplinkPcRole) *NullableFabricFcoeUplinkPcRole {
	return &NullableFabricFcoeUplinkPcRole{value: val, isSet: true}
}

func (v NullableFabricFcoeUplinkPcRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricFcoeUplinkPcRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


