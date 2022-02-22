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

// FcpoolBlock Block of WWN identifiers.
type FcpoolBlock struct {
	PoolAbstractBlockType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Starting WWN identifier of the block must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN's in the SAN fabric, you are strongly encouraged to use the following WWN prefix; 20:00:00:25:B5:xx:xx:xx.
	From *string `json:"From,omitempty"`
	// Ending WWN identifier of the block must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.
	To *string `json:"To,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FcpoolBlock FcpoolBlock

// NewFcpoolBlock instantiates a new FcpoolBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcpoolBlock(classId string, objectType string) *FcpoolBlock {
	this := FcpoolBlock{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFcpoolBlockWithDefaults instantiates a new FcpoolBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcpoolBlockWithDefaults() *FcpoolBlock {
	this := FcpoolBlock{}
	var classId string = "fcpool.Block"
	this.ClassId = classId
	var objectType string = "fcpool.Block"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FcpoolBlock) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FcpoolBlock) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FcpoolBlock) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FcpoolBlock) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FcpoolBlock) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FcpoolBlock) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *FcpoolBlock) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolBlock) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *FcpoolBlock) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *FcpoolBlock) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *FcpoolBlock) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolBlock) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *FcpoolBlock) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *FcpoolBlock) SetTo(v string) {
	o.To = &v
}

func (o FcpoolBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractBlockType, errPoolAbstractBlockType := json.Marshal(o.PoolAbstractBlockType)
	if errPoolAbstractBlockType != nil {
		return []byte{}, errPoolAbstractBlockType
	}
	errPoolAbstractBlockType = json.Unmarshal([]byte(serializedPoolAbstractBlockType), &toSerialize)
	if errPoolAbstractBlockType != nil {
		return []byte{}, errPoolAbstractBlockType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.From != nil {
		toSerialize["From"] = o.From
	}
	if o.To != nil {
		toSerialize["To"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FcpoolBlock) UnmarshalJSON(bytes []byte) (err error) {
	type FcpoolBlockWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Starting WWN identifier of the block must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN's in the SAN fabric, you are strongly encouraged to use the following WWN prefix; 20:00:00:25:B5:xx:xx:xx.
		From *string `json:"From,omitempty"`
		// Ending WWN identifier of the block must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.
		To *string `json:"To,omitempty"`
	}

	varFcpoolBlockWithoutEmbeddedStruct := FcpoolBlockWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFcpoolBlockWithoutEmbeddedStruct)
	if err == nil {
		varFcpoolBlock := _FcpoolBlock{}
		varFcpoolBlock.ClassId = varFcpoolBlockWithoutEmbeddedStruct.ClassId
		varFcpoolBlock.ObjectType = varFcpoolBlockWithoutEmbeddedStruct.ObjectType
		varFcpoolBlock.From = varFcpoolBlockWithoutEmbeddedStruct.From
		varFcpoolBlock.To = varFcpoolBlockWithoutEmbeddedStruct.To
		*o = FcpoolBlock(varFcpoolBlock)
	} else {
		return err
	}

	varFcpoolBlock := _FcpoolBlock{}

	err = json.Unmarshal(bytes, &varFcpoolBlock)
	if err == nil {
		o.PoolAbstractBlockType = varFcpoolBlock.PoolAbstractBlockType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "From")
		delete(additionalProperties, "To")

		// remove fields from embedded structs
		reflectPoolAbstractBlockType := reflect.ValueOf(o.PoolAbstractBlockType)
		for i := 0; i < reflectPoolAbstractBlockType.Type().NumField(); i++ {
			t := reflectPoolAbstractBlockType.Type().Field(i)

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

type NullableFcpoolBlock struct {
	value *FcpoolBlock
	isSet bool
}

func (v NullableFcpoolBlock) Get() *FcpoolBlock {
	return v.value
}

func (v *NullableFcpoolBlock) Set(val *FcpoolBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpoolBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpoolBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpoolBlock(val *FcpoolBlock) *NullableFcpoolBlock {
	return &NullableFcpoolBlock{value: val, isSet: true}
}

func (v NullableFcpoolBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcpoolBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


