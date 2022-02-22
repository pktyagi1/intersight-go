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
)

// ResourcePerTypeCombinedSelectorAllOf Definition of the list of properties defined in 'resource.PerTypeCombinedSelector', excluding properties defined in parent classes.
type ResourcePerTypeCombinedSelectorAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A single filter expression created by OR'ing the $filter criteria of the 'selectors'. Used to efficiently maintain the membership of the Group.
	CombinedSelector *string `json:"CombinedSelector,omitempty"`
	// If true, then resources are added using just object type without filter.
	EmptyFilter *bool `json:"EmptyFilter,omitempty"`
	// The ObjectType on which the selectors are defined. Used to efficiently query resource groups for a given ObjectType.
	SelectorObjectType *string `json:"SelectorObjectType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourcePerTypeCombinedSelectorAllOf ResourcePerTypeCombinedSelectorAllOf

// NewResourcePerTypeCombinedSelectorAllOf instantiates a new ResourcePerTypeCombinedSelectorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcePerTypeCombinedSelectorAllOf(classId string, objectType string) *ResourcePerTypeCombinedSelectorAllOf {
	this := ResourcePerTypeCombinedSelectorAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewResourcePerTypeCombinedSelectorAllOfWithDefaults instantiates a new ResourcePerTypeCombinedSelectorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcePerTypeCombinedSelectorAllOfWithDefaults() *ResourcePerTypeCombinedSelectorAllOf {
	this := ResourcePerTypeCombinedSelectorAllOf{}
	var classId string = "resource.PerTypeCombinedSelector"
	this.ClassId = classId
	var objectType string = "resource.PerTypeCombinedSelector"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourcePerTypeCombinedSelectorAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourcePerTypeCombinedSelectorAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourcePerTypeCombinedSelectorAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourcePerTypeCombinedSelectorAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourcePerTypeCombinedSelectorAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourcePerTypeCombinedSelectorAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCombinedSelector returns the CombinedSelector field value if set, zero value otherwise.
func (o *ResourcePerTypeCombinedSelectorAllOf) GetCombinedSelector() string {
	if o == nil || o.CombinedSelector == nil {
		var ret string
		return ret
	}
	return *o.CombinedSelector
}

// GetCombinedSelectorOk returns a tuple with the CombinedSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePerTypeCombinedSelectorAllOf) GetCombinedSelectorOk() (*string, bool) {
	if o == nil || o.CombinedSelector == nil {
		return nil, false
	}
	return o.CombinedSelector, true
}

// HasCombinedSelector returns a boolean if a field has been set.
func (o *ResourcePerTypeCombinedSelectorAllOf) HasCombinedSelector() bool {
	if o != nil && o.CombinedSelector != nil {
		return true
	}

	return false
}

// SetCombinedSelector gets a reference to the given string and assigns it to the CombinedSelector field.
func (o *ResourcePerTypeCombinedSelectorAllOf) SetCombinedSelector(v string) {
	o.CombinedSelector = &v
}

// GetEmptyFilter returns the EmptyFilter field value if set, zero value otherwise.
func (o *ResourcePerTypeCombinedSelectorAllOf) GetEmptyFilter() bool {
	if o == nil || o.EmptyFilter == nil {
		var ret bool
		return ret
	}
	return *o.EmptyFilter
}

// GetEmptyFilterOk returns a tuple with the EmptyFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePerTypeCombinedSelectorAllOf) GetEmptyFilterOk() (*bool, bool) {
	if o == nil || o.EmptyFilter == nil {
		return nil, false
	}
	return o.EmptyFilter, true
}

// HasEmptyFilter returns a boolean if a field has been set.
func (o *ResourcePerTypeCombinedSelectorAllOf) HasEmptyFilter() bool {
	if o != nil && o.EmptyFilter != nil {
		return true
	}

	return false
}

// SetEmptyFilter gets a reference to the given bool and assigns it to the EmptyFilter field.
func (o *ResourcePerTypeCombinedSelectorAllOf) SetEmptyFilter(v bool) {
	o.EmptyFilter = &v
}

// GetSelectorObjectType returns the SelectorObjectType field value if set, zero value otherwise.
func (o *ResourcePerTypeCombinedSelectorAllOf) GetSelectorObjectType() string {
	if o == nil || o.SelectorObjectType == nil {
		var ret string
		return ret
	}
	return *o.SelectorObjectType
}

// GetSelectorObjectTypeOk returns a tuple with the SelectorObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePerTypeCombinedSelectorAllOf) GetSelectorObjectTypeOk() (*string, bool) {
	if o == nil || o.SelectorObjectType == nil {
		return nil, false
	}
	return o.SelectorObjectType, true
}

// HasSelectorObjectType returns a boolean if a field has been set.
func (o *ResourcePerTypeCombinedSelectorAllOf) HasSelectorObjectType() bool {
	if o != nil && o.SelectorObjectType != nil {
		return true
	}

	return false
}

// SetSelectorObjectType gets a reference to the given string and assigns it to the SelectorObjectType field.
func (o *ResourcePerTypeCombinedSelectorAllOf) SetSelectorObjectType(v string) {
	o.SelectorObjectType = &v
}

func (o ResourcePerTypeCombinedSelectorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CombinedSelector != nil {
		toSerialize["CombinedSelector"] = o.CombinedSelector
	}
	if o.EmptyFilter != nil {
		toSerialize["EmptyFilter"] = o.EmptyFilter
	}
	if o.SelectorObjectType != nil {
		toSerialize["SelectorObjectType"] = o.SelectorObjectType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourcePerTypeCombinedSelectorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varResourcePerTypeCombinedSelectorAllOf := _ResourcePerTypeCombinedSelectorAllOf{}

	if err = json.Unmarshal(bytes, &varResourcePerTypeCombinedSelectorAllOf); err == nil {
		*o = ResourcePerTypeCombinedSelectorAllOf(varResourcePerTypeCombinedSelectorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CombinedSelector")
		delete(additionalProperties, "EmptyFilter")
		delete(additionalProperties, "SelectorObjectType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourcePerTypeCombinedSelectorAllOf struct {
	value *ResourcePerTypeCombinedSelectorAllOf
	isSet bool
}

func (v NullableResourcePerTypeCombinedSelectorAllOf) Get() *ResourcePerTypeCombinedSelectorAllOf {
	return v.value
}

func (v *NullableResourcePerTypeCombinedSelectorAllOf) Set(val *ResourcePerTypeCombinedSelectorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcePerTypeCombinedSelectorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcePerTypeCombinedSelectorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcePerTypeCombinedSelectorAllOf(val *ResourcePerTypeCombinedSelectorAllOf) *NullableResourcePerTypeCombinedSelectorAllOf {
	return &NullableResourcePerTypeCombinedSelectorAllOf{value: val, isSet: true}
}

func (v NullableResourcePerTypeCombinedSelectorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcePerTypeCombinedSelectorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


