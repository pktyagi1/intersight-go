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

// MetaPropDefinitionAllOf Definition of the list of properties defined in 'meta.PropDefinition', excluding properties defined in parent classes.
type MetaPropDefinitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// API access control for the property. Examples are NoAccess, ReadOnly, CreateOnly etc. * `NoAccess` - The property is not accessible from the API. * `ReadOnly` - The value of the property is read-only.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa ReadOnly property. * `CreateOnly` - The value of the property can be set when the REST resource is created. It cannot be changed after object creation.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa CreateOnly property.CreateOnly properties are returned in the response body of HTTP GET requests. * `ReadWrite` - The property has read/write access. * `WriteOnly` - The value of the property can be set but it is never returned in the response body of supported HTTP methods.This settings is used for sensitive properties such as passwords. * `ReadOnCreate` - The value of the property is returned in the HTTP POST response body when the REST resource is created.The property is not writeable and cannot be queried through a GET request after the resource has been created.
	ApiAccess *string `json:"ApiAccess,omitempty"`
	// The default value of the property. Not applicable when IsComplexType is True.
	Default interface{} `json:"Default,omitempty"`
	// Indicates whether the property is a collection (i.e. a JSON array), otherwise it is a single value.
	IsCollection *bool `json:"IsCollection,omitempty"`
	// Indicates whether the property is a complex type, otherwise it is a basic scalar type.
	IsComplexType *bool `json:"IsComplexType,omitempty"`
	// The kind of the property. * `Unknown` - The property kind is unknown. * `Boolean` - The 'Boolean' kind of property. * `String` - The 'String' kind of property. * `Integer` - The 'Integer' kind of property. * `Int32` - The 'Int32' kind of property. * `Int64` - The 'Int64' kind of property. * `Float` - The 'Float' kind of property. * `Double` - The 'Double' kind of property. * `Date` - The 'Date' kind of property. * `Duration` - The 'Duration' kind of property. * `Time` - The 'Time' kind of property. * `Json` - The 'JSON' kind of property. * `Binary` - The 'Binary' kind of property. * `EnumString` - The 'EnumString' kind of property. * `EnumInteger` - The 'EnumInteger' kind of property. * `ComplexType` - The 'ComplexType' kind of property.
	Kind *string `json:"Kind,omitempty"`
	// The name of the property.
	Name *string `json:"Name,omitempty"`
	// The data-at-rest security protection applied to this property when a Managed Object is persisted. For example, Encrypted or Cleartext. * `ClearText` - Data at rest is not encrypted using account specific keys.Note that data is always protected using volume encryption. ClearText propertiesare queryable and searchable. * `Encrypted` - The value of the property is encrypted with account-specific cryptographic keys.This setting is used for properties that contain sensitive data. Encrypted propertiesare not queryable and searchable. * `Pbkdf2` - The value of the property is hashed using the pbkdf2 key derivation functions thattakes a password, a salt, and a cost factor as inputs then generates a password hash.Its purpose is to make each password guessing trial by an attacker who has obtaineda password hash file expensive and therefore the cost of a guessing attack high or prohibitive. * `Bcrypt` - The value of the property is hashed using the bcrypt key derivation function. * `Sha512crypt` - The value of the property is hashed using the sha512crypt key derivation function. * `Argon2id` - The value of the property is hashed using the argon2id key derivation function.
	OpSecurity *string `json:"OpSecurity,omitempty"`
	// Enables the property to be searchable from global search. A value of 0 means this property is not globally searchable.
	SearchWeight *float32 `json:"SearchWeight,omitempty"`
	// The type of the property. In case of collection properties the type is that of individual elements.
	Type *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetaPropDefinitionAllOf MetaPropDefinitionAllOf

// NewMetaPropDefinitionAllOf instantiates a new MetaPropDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaPropDefinitionAllOf(classId string, objectType string) *MetaPropDefinitionAllOf {
	this := MetaPropDefinitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMetaPropDefinitionAllOfWithDefaults instantiates a new MetaPropDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaPropDefinitionAllOfWithDefaults() *MetaPropDefinitionAllOf {
	this := MetaPropDefinitionAllOf{}
	var classId string = "meta.PropDefinition"
	this.ClassId = classId
	var objectType string = "meta.PropDefinition"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MetaPropDefinitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MetaPropDefinitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MetaPropDefinitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MetaPropDefinitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MetaPropDefinitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MetaPropDefinitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetApiAccess returns the ApiAccess field value if set, zero value otherwise.
func (o *MetaPropDefinitionAllOf) GetApiAccess() string {
	if o == nil || o.ApiAccess == nil {
		var ret string
		return ret
	}
	return *o.ApiAccess
}

// GetApiAccessOk returns a tuple with the ApiAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaPropDefinitionAllOf) GetApiAccessOk() (*string, bool) {
	if o == nil || o.ApiAccess == nil {
		return nil, false
	}
	return o.ApiAccess, true
}

// HasApiAccess returns a boolean if a field has been set.
func (o *MetaPropDefinitionAllOf) HasApiAccess() bool {
	if o != nil && o.ApiAccess != nil {
		return true
	}

	return false
}

// SetApiAccess gets a reference to the given string and assigns it to the ApiAccess field.
func (o *MetaPropDefinitionAllOf) SetApiAccess(v string) {
	o.ApiAccess = &v
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetaPropDefinitionAllOf) GetDefault() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPropDefinitionAllOf) GetDefaultOk() (*interface{}, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return &o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *MetaPropDefinitionAllOf) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given interface{} and assigns it to the Default field.
func (o *MetaPropDefinitionAllOf) SetDefault(v interface{}) {
	o.Default = v
}

// GetIsCollection returns the IsCollection field value if set, zero value otherwise.
func (o *MetaPropDefinitionAllOf) GetIsCollection() bool {
	if o == nil || o.IsCollection == nil {
		var ret bool
		return ret
	}
	return *o.IsCollection
}

// GetIsCollectionOk returns a tuple with the IsCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaPropDefinitionAllOf) GetIsCollectionOk() (*bool, bool) {
	if o == nil || o.IsCollection == nil {
		return nil, false
	}
	return o.IsCollection, true
}

// HasIsCollection returns a boolean if a field has been set.
func (o *MetaPropDefinitionAllOf) HasIsCollection() bool {
	if o != nil && o.IsCollection != nil {
		return true
	}

	return false
}

// SetIsCollection gets a reference to the given bool and assigns it to the IsCollection field.
func (o *MetaPropDefinitionAllOf) SetIsCollection(v bool) {
	o.IsCollection = &v
}

// GetIsComplexType returns the IsComplexType field value if set, zero value otherwise.
func (o *MetaPropDefinitionAllOf) GetIsComplexType() bool {
	if o == nil || o.IsComplexType == nil {
		var ret bool
		return ret
	}
	return *o.IsComplexType
}

// GetIsComplexTypeOk returns a tuple with the IsComplexType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaPropDefinitionAllOf) GetIsComplexTypeOk() (*bool, bool) {
	if o == nil || o.IsComplexType == nil {
		return nil, false
	}
	return o.IsComplexType, true
}

// HasIsComplexType returns a boolean if a field has been set.
func (o *MetaPropDefinitionAllOf) HasIsComplexType() bool {
	if o != nil && o.IsComplexType != nil {
		return true
	}

	return false
}

// SetIsComplexType gets a reference to the given bool and assigns it to the IsComplexType field.
func (o *MetaPropDefinitionAllOf) SetIsComplexType(v bool) {
	o.IsComplexType = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *MetaPropDefinitionAllOf) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaPropDefinitionAllOf) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *MetaPropDefinitionAllOf) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *MetaPropDefinitionAllOf) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetaPropDefinitionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaPropDefinitionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetaPropDefinitionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MetaPropDefinitionAllOf) SetName(v string) {
	o.Name = &v
}

// GetOpSecurity returns the OpSecurity field value if set, zero value otherwise.
func (o *MetaPropDefinitionAllOf) GetOpSecurity() string {
	if o == nil || o.OpSecurity == nil {
		var ret string
		return ret
	}
	return *o.OpSecurity
}

// GetOpSecurityOk returns a tuple with the OpSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaPropDefinitionAllOf) GetOpSecurityOk() (*string, bool) {
	if o == nil || o.OpSecurity == nil {
		return nil, false
	}
	return o.OpSecurity, true
}

// HasOpSecurity returns a boolean if a field has been set.
func (o *MetaPropDefinitionAllOf) HasOpSecurity() bool {
	if o != nil && o.OpSecurity != nil {
		return true
	}

	return false
}

// SetOpSecurity gets a reference to the given string and assigns it to the OpSecurity field.
func (o *MetaPropDefinitionAllOf) SetOpSecurity(v string) {
	o.OpSecurity = &v
}

// GetSearchWeight returns the SearchWeight field value if set, zero value otherwise.
func (o *MetaPropDefinitionAllOf) GetSearchWeight() float32 {
	if o == nil || o.SearchWeight == nil {
		var ret float32
		return ret
	}
	return *o.SearchWeight
}

// GetSearchWeightOk returns a tuple with the SearchWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaPropDefinitionAllOf) GetSearchWeightOk() (*float32, bool) {
	if o == nil || o.SearchWeight == nil {
		return nil, false
	}
	return o.SearchWeight, true
}

// HasSearchWeight returns a boolean if a field has been set.
func (o *MetaPropDefinitionAllOf) HasSearchWeight() bool {
	if o != nil && o.SearchWeight != nil {
		return true
	}

	return false
}

// SetSearchWeight gets a reference to the given float32 and assigns it to the SearchWeight field.
func (o *MetaPropDefinitionAllOf) SetSearchWeight(v float32) {
	o.SearchWeight = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MetaPropDefinitionAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaPropDefinitionAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetaPropDefinitionAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MetaPropDefinitionAllOf) SetType(v string) {
	o.Type = &v
}

func (o MetaPropDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ApiAccess != nil {
		toSerialize["ApiAccess"] = o.ApiAccess
	}
	if o.Default != nil {
		toSerialize["Default"] = o.Default
	}
	if o.IsCollection != nil {
		toSerialize["IsCollection"] = o.IsCollection
	}
	if o.IsComplexType != nil {
		toSerialize["IsComplexType"] = o.IsComplexType
	}
	if o.Kind != nil {
		toSerialize["Kind"] = o.Kind
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OpSecurity != nil {
		toSerialize["OpSecurity"] = o.OpSecurity
	}
	if o.SearchWeight != nil {
		toSerialize["SearchWeight"] = o.SearchWeight
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MetaPropDefinitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMetaPropDefinitionAllOf := _MetaPropDefinitionAllOf{}

	if err = json.Unmarshal(bytes, &varMetaPropDefinitionAllOf); err == nil {
		*o = MetaPropDefinitionAllOf(varMetaPropDefinitionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApiAccess")
		delete(additionalProperties, "Default")
		delete(additionalProperties, "IsCollection")
		delete(additionalProperties, "IsComplexType")
		delete(additionalProperties, "Kind")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OpSecurity")
		delete(additionalProperties, "SearchWeight")
		delete(additionalProperties, "Type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetaPropDefinitionAllOf struct {
	value *MetaPropDefinitionAllOf
	isSet bool
}

func (v NullableMetaPropDefinitionAllOf) Get() *MetaPropDefinitionAllOf {
	return v.value
}

func (v *NullableMetaPropDefinitionAllOf) Set(val *MetaPropDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaPropDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaPropDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaPropDefinitionAllOf(val *MetaPropDefinitionAllOf) *NullableMetaPropDefinitionAllOf {
	return &NullableMetaPropDefinitionAllOf{value: val, isSet: true}
}

func (v NullableMetaPropDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaPropDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

