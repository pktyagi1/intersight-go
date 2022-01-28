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

// NiatelemetryNode Stores information related to golden image on dcnm controller.
type NiatelemetryNode struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns hostname of the node.
	Hostname *string `json:"Hostname,omitempty"`
	// Returns management IP of the node.
	ManagementtIp *string `json:"ManagementtIp,omitempty"`
	// Returns out of band IP of the node.
	OutofbandIp *string `json:"OutofbandIp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNode NiatelemetryNode

// NewNiatelemetryNode instantiates a new NiatelemetryNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNode(classId string, objectType string) *NiatelemetryNode {
	this := NiatelemetryNode{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNodeWithDefaults instantiates a new NiatelemetryNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNodeWithDefaults() *NiatelemetryNode {
	this := NiatelemetryNode{}
	var classId string = "niatelemetry.Node"
	this.ClassId = classId
	var objectType string = "niatelemetry.Node"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNode) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNode) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNode) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNode) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNode) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNode) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *NiatelemetryNode) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNode) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *NiatelemetryNode) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *NiatelemetryNode) SetHostname(v string) {
	o.Hostname = &v
}

// GetManagementtIp returns the ManagementtIp field value if set, zero value otherwise.
func (o *NiatelemetryNode) GetManagementtIp() string {
	if o == nil || o.ManagementtIp == nil {
		var ret string
		return ret
	}
	return *o.ManagementtIp
}

// GetManagementtIpOk returns a tuple with the ManagementtIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNode) GetManagementtIpOk() (*string, bool) {
	if o == nil || o.ManagementtIp == nil {
		return nil, false
	}
	return o.ManagementtIp, true
}

// HasManagementtIp returns a boolean if a field has been set.
func (o *NiatelemetryNode) HasManagementtIp() bool {
	if o != nil && o.ManagementtIp != nil {
		return true
	}

	return false
}

// SetManagementtIp gets a reference to the given string and assigns it to the ManagementtIp field.
func (o *NiatelemetryNode) SetManagementtIp(v string) {
	o.ManagementtIp = &v
}

// GetOutofbandIp returns the OutofbandIp field value if set, zero value otherwise.
func (o *NiatelemetryNode) GetOutofbandIp() string {
	if o == nil || o.OutofbandIp == nil {
		var ret string
		return ret
	}
	return *o.OutofbandIp
}

// GetOutofbandIpOk returns a tuple with the OutofbandIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNode) GetOutofbandIpOk() (*string, bool) {
	if o == nil || o.OutofbandIp == nil {
		return nil, false
	}
	return o.OutofbandIp, true
}

// HasOutofbandIp returns a boolean if a field has been set.
func (o *NiatelemetryNode) HasOutofbandIp() bool {
	if o != nil && o.OutofbandIp != nil {
		return true
	}

	return false
}

// SetOutofbandIp gets a reference to the given string and assigns it to the OutofbandIp field.
func (o *NiatelemetryNode) SetOutofbandIp(v string) {
	o.OutofbandIp = &v
}

func (o NiatelemetryNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.ManagementtIp != nil {
		toSerialize["ManagementtIp"] = o.ManagementtIp
	}
	if o.OutofbandIp != nil {
		toSerialize["OutofbandIp"] = o.OutofbandIp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNode) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryNodeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Returns hostname of the node.
		Hostname *string `json:"Hostname,omitempty"`
		// Returns management IP of the node.
		ManagementtIp *string `json:"ManagementtIp,omitempty"`
		// Returns out of band IP of the node.
		OutofbandIp *string `json:"OutofbandIp,omitempty"`
	}

	varNiatelemetryNodeWithoutEmbeddedStruct := NiatelemetryNodeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryNodeWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryNode := _NiatelemetryNode{}
		varNiatelemetryNode.ClassId = varNiatelemetryNodeWithoutEmbeddedStruct.ClassId
		varNiatelemetryNode.ObjectType = varNiatelemetryNodeWithoutEmbeddedStruct.ObjectType
		varNiatelemetryNode.Hostname = varNiatelemetryNodeWithoutEmbeddedStruct.Hostname
		varNiatelemetryNode.ManagementtIp = varNiatelemetryNodeWithoutEmbeddedStruct.ManagementtIp
		varNiatelemetryNode.OutofbandIp = varNiatelemetryNodeWithoutEmbeddedStruct.OutofbandIp
		*o = NiatelemetryNode(varNiatelemetryNode)
	} else {
		return err
	}

	varNiatelemetryNode := _NiatelemetryNode{}

	err = json.Unmarshal(bytes, &varNiatelemetryNode)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryNode.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "ManagementtIp")
		delete(additionalProperties, "OutofbandIp")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableNiatelemetryNode struct {
	value *NiatelemetryNode
	isSet bool
}

func (v NullableNiatelemetryNode) Get() *NiatelemetryNode {
	return v.value
}

func (v *NullableNiatelemetryNode) Set(val *NiatelemetryNode) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNode) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNode(val *NiatelemetryNode) *NullableNiatelemetryNode {
	return &NullableNiatelemetryNode{value: val, isSet: true}
}

func (v NullableNiatelemetryNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


