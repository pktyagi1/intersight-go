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
	"reflect"
	"strings"
)

// ConnectorScopedInventory Abstract object defined to handle scoped inventory for a set of objects identified by the query parameters defined in the properties. All vendor specific scoped inventory objects will inherit from this object and the scoped inventory framework will address the inventory collection.
type ConnectorScopedInventory struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// A property that uniquely identifies the object to be inventoried as a part of the scoped inventory.
	NamingProperty *string `json:"NamingProperty,omitempty"`
	// Set of queries to identify objects to be inventoried as part of this scoped inventory action.
	Queries interface{} `json:"Queries,omitempty"`
	// Type of the object for which scoped inventory needs to be run.
	Type *string `json:"Type,omitempty"`
	Values []string `json:"Values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorScopedInventory ConnectorScopedInventory

// NewConnectorScopedInventory instantiates a new ConnectorScopedInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorScopedInventory(classId string, objectType string) *ConnectorScopedInventory {
	this := ConnectorScopedInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorScopedInventoryWithDefaults instantiates a new ConnectorScopedInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorScopedInventoryWithDefaults() *ConnectorScopedInventory {
	this := ConnectorScopedInventory{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorScopedInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorScopedInventory) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorScopedInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorScopedInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorScopedInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorScopedInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNamingProperty returns the NamingProperty field value if set, zero value otherwise.
func (o *ConnectorScopedInventory) GetNamingProperty() string {
	if o == nil || o.NamingProperty == nil {
		var ret string
		return ret
	}
	return *o.NamingProperty
}

// GetNamingPropertyOk returns a tuple with the NamingProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorScopedInventory) GetNamingPropertyOk() (*string, bool) {
	if o == nil || o.NamingProperty == nil {
		return nil, false
	}
	return o.NamingProperty, true
}

// HasNamingProperty returns a boolean if a field has been set.
func (o *ConnectorScopedInventory) HasNamingProperty() bool {
	if o != nil && o.NamingProperty != nil {
		return true
	}

	return false
}

// SetNamingProperty gets a reference to the given string and assigns it to the NamingProperty field.
func (o *ConnectorScopedInventory) SetNamingProperty(v string) {
	o.NamingProperty = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorScopedInventory) GetQueries() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorScopedInventory) GetQueriesOk() (*interface{}, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return &o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *ConnectorScopedInventory) HasQueries() bool {
	if o != nil && o.Queries != nil {
		return true
	}

	return false
}

// SetQueries gets a reference to the given interface{} and assigns it to the Queries field.
func (o *ConnectorScopedInventory) SetQueries(v interface{}) {
	o.Queries = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConnectorScopedInventory) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorScopedInventory) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConnectorScopedInventory) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConnectorScopedInventory) SetType(v string) {
	o.Type = &v
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorScopedInventory) GetValues() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorScopedInventory) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ConnectorScopedInventory) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *ConnectorScopedInventory) SetValues(v []string) {
	o.Values = v
}

func (o ConnectorScopedInventory) MarshalJSON() ([]byte, error) {
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
	if o.NamingProperty != nil {
		toSerialize["NamingProperty"] = o.NamingProperty
	}
	if o.Queries != nil {
		toSerialize["Queries"] = o.Queries
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Values != nil {
		toSerialize["Values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorScopedInventory) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorScopedInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// A property that uniquely identifies the object to be inventoried as a part of the scoped inventory.
		NamingProperty *string `json:"NamingProperty,omitempty"`
		// Set of queries to identify objects to be inventoried as part of this scoped inventory action.
		Queries interface{} `json:"Queries,omitempty"`
		// Type of the object for which scoped inventory needs to be run.
		Type *string `json:"Type,omitempty"`
		Values []string `json:"Values,omitempty"`
	}

	varConnectorScopedInventoryWithoutEmbeddedStruct := ConnectorScopedInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorScopedInventoryWithoutEmbeddedStruct)
	if err == nil {
		varConnectorScopedInventory := _ConnectorScopedInventory{}
		varConnectorScopedInventory.ClassId = varConnectorScopedInventoryWithoutEmbeddedStruct.ClassId
		varConnectorScopedInventory.ObjectType = varConnectorScopedInventoryWithoutEmbeddedStruct.ObjectType
		varConnectorScopedInventory.NamingProperty = varConnectorScopedInventoryWithoutEmbeddedStruct.NamingProperty
		varConnectorScopedInventory.Queries = varConnectorScopedInventoryWithoutEmbeddedStruct.Queries
		varConnectorScopedInventory.Type = varConnectorScopedInventoryWithoutEmbeddedStruct.Type
		varConnectorScopedInventory.Values = varConnectorScopedInventoryWithoutEmbeddedStruct.Values
		*o = ConnectorScopedInventory(varConnectorScopedInventory)
	} else {
		return err
	}

	varConnectorScopedInventory := _ConnectorScopedInventory{}

	err = json.Unmarshal(bytes, &varConnectorScopedInventory)
	if err == nil {
		o.MoBaseMo = varConnectorScopedInventory.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NamingProperty")
		delete(additionalProperties, "Queries")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Values")

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

type NullableConnectorScopedInventory struct {
	value *ConnectorScopedInventory
	isSet bool
}

func (v NullableConnectorScopedInventory) Get() *ConnectorScopedInventory {
	return v.value
}

func (v *NullableConnectorScopedInventory) Set(val *ConnectorScopedInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorScopedInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorScopedInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorScopedInventory(val *ConnectorScopedInventory) *NullableConnectorScopedInventory {
	return &NullableConnectorScopedInventory{value: val, isSet: true}
}

func (v NullableConnectorScopedInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorScopedInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


