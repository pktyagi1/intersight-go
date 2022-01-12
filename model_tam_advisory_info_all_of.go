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

// TamAdvisoryInfoAllOf Definition of the list of properties defined in 'tam.AdvisoryInfo', excluding properties defined in parent classes.
type TamAdvisoryInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Current state of the advisory for the owner. Indicates if the user is interested in getting updates for the advisory. * `active` - Advisory is currently active and the user wants to receive updates for this advisory. * `acknowledged` - Advisory is seen and acknowledged by the user and she no longer wants to recieve updates.
	State *string `json:"State,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	Advisory *TamBaseAdvisoryRelationship `json:"Advisory,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamAdvisoryInfoAllOf TamAdvisoryInfoAllOf

// NewTamAdvisoryInfoAllOf instantiates a new TamAdvisoryInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamAdvisoryInfoAllOf(classId string, objectType string) *TamAdvisoryInfoAllOf {
	this := TamAdvisoryInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var state string = "active"
	this.State = &state
	return &this
}

// NewTamAdvisoryInfoAllOfWithDefaults instantiates a new TamAdvisoryInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamAdvisoryInfoAllOfWithDefaults() *TamAdvisoryInfoAllOf {
	this := TamAdvisoryInfoAllOf{}
	var classId string = "tam.AdvisoryInfo"
	this.ClassId = classId
	var objectType string = "tam.AdvisoryInfo"
	this.ObjectType = objectType
	var state string = "active"
	this.State = &state
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamAdvisoryInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamAdvisoryInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamAdvisoryInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamAdvisoryInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TamAdvisoryInfoAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInfoAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TamAdvisoryInfoAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TamAdvisoryInfoAllOf) SetState(v string) {
	o.State = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *TamAdvisoryInfoAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInfoAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *TamAdvisoryInfoAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *TamAdvisoryInfoAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetAdvisory returns the Advisory field value if set, zero value otherwise.
func (o *TamAdvisoryInfoAllOf) GetAdvisory() TamBaseAdvisoryRelationship {
	if o == nil || o.Advisory == nil {
		var ret TamBaseAdvisoryRelationship
		return ret
	}
	return *o.Advisory
}

// GetAdvisoryOk returns a tuple with the Advisory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInfoAllOf) GetAdvisoryOk() (*TamBaseAdvisoryRelationship, bool) {
	if o == nil || o.Advisory == nil {
		return nil, false
	}
	return o.Advisory, true
}

// HasAdvisory returns a boolean if a field has been set.
func (o *TamAdvisoryInfoAllOf) HasAdvisory() bool {
	if o != nil && o.Advisory != nil {
		return true
	}

	return false
}

// SetAdvisory gets a reference to the given TamBaseAdvisoryRelationship and assigns it to the Advisory field.
func (o *TamAdvisoryInfoAllOf) SetAdvisory(v TamBaseAdvisoryRelationship) {
	o.Advisory = &v
}

func (o TamAdvisoryInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Advisory != nil {
		toSerialize["Advisory"] = o.Advisory
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamAdvisoryInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTamAdvisoryInfoAllOf := _TamAdvisoryInfoAllOf{}

	if err = json.Unmarshal(bytes, &varTamAdvisoryInfoAllOf); err == nil {
		*o = TamAdvisoryInfoAllOf(varTamAdvisoryInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Advisory")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTamAdvisoryInfoAllOf struct {
	value *TamAdvisoryInfoAllOf
	isSet bool
}

func (v NullableTamAdvisoryInfoAllOf) Get() *TamAdvisoryInfoAllOf {
	return v.value
}

func (v *NullableTamAdvisoryInfoAllOf) Set(val *TamAdvisoryInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamAdvisoryInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamAdvisoryInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamAdvisoryInfoAllOf(val *TamAdvisoryInfoAllOf) *NullableTamAdvisoryInfoAllOf {
	return &NullableTamAdvisoryInfoAllOf{value: val, isSet: true}
}

func (v NullableTamAdvisoryInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamAdvisoryInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


