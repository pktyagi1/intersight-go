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

// SessionAbstractSubSession A base abstract class for all sub-sessions.
type SessionAbstractSubSession struct {
	SessionAbstractSession
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Name of target on which session is initiated.
	TargetName *string `json:"TargetName,omitempty"`
	Session *SessionAbstractSessionRelationship `json:"Session,omitempty"`
	Target *MoBaseMoRelationship `json:"Target,omitempty"`
	User *IamUserRelationship `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionAbstractSubSession SessionAbstractSubSession

// NewSessionAbstractSubSession instantiates a new SessionAbstractSubSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionAbstractSubSession(classId string, objectType string) *SessionAbstractSubSession {
	this := SessionAbstractSubSession{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Active"
	this.Status = &status
	return &this
}

// NewSessionAbstractSubSessionWithDefaults instantiates a new SessionAbstractSubSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionAbstractSubSessionWithDefaults() *SessionAbstractSubSession {
	this := SessionAbstractSubSession{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *SessionAbstractSubSession) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SessionAbstractSubSession) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SessionAbstractSubSession) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SessionAbstractSubSession) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SessionAbstractSubSession) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SessionAbstractSubSession) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *SessionAbstractSubSession) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionAbstractSubSession) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *SessionAbstractSubSession) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *SessionAbstractSubSession) SetTargetName(v string) {
	o.TargetName = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *SessionAbstractSubSession) GetSession() SessionAbstractSessionRelationship {
	if o == nil || o.Session == nil {
		var ret SessionAbstractSessionRelationship
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionAbstractSubSession) GetSessionOk() (*SessionAbstractSessionRelationship, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *SessionAbstractSubSession) HasSession() bool {
	if o != nil && o.Session != nil {
		return true
	}

	return false
}

// SetSession gets a reference to the given SessionAbstractSessionRelationship and assigns it to the Session field.
func (o *SessionAbstractSubSession) SetSession(v SessionAbstractSessionRelationship) {
	o.Session = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SessionAbstractSubSession) GetTarget() MoBaseMoRelationship {
	if o == nil || o.Target == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionAbstractSubSession) GetTargetOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SessionAbstractSubSession) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given MoBaseMoRelationship and assigns it to the Target field.
func (o *SessionAbstractSubSession) SetTarget(v MoBaseMoRelationship) {
	o.Target = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *SessionAbstractSubSession) GetUser() IamUserRelationship {
	if o == nil || o.User == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionAbstractSubSession) GetUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *SessionAbstractSubSession) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IamUserRelationship and assigns it to the User field.
func (o *SessionAbstractSubSession) SetUser(v IamUserRelationship) {
	o.User = &v
}

func (o SessionAbstractSubSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSessionAbstractSession, errSessionAbstractSession := json.Marshal(o.SessionAbstractSession)
	if errSessionAbstractSession != nil {
		return []byte{}, errSessionAbstractSession
	}
	errSessionAbstractSession = json.Unmarshal([]byte(serializedSessionAbstractSession), &toSerialize)
	if errSessionAbstractSession != nil {
		return []byte{}, errSessionAbstractSession
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TargetName != nil {
		toSerialize["TargetName"] = o.TargetName
	}
	if o.Session != nil {
		toSerialize["Session"] = o.Session
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SessionAbstractSubSession) UnmarshalJSON(bytes []byte) (err error) {
	type SessionAbstractSubSessionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Name of target on which session is initiated.
		TargetName *string `json:"TargetName,omitempty"`
		Session *SessionAbstractSessionRelationship `json:"Session,omitempty"`
		Target *MoBaseMoRelationship `json:"Target,omitempty"`
		User *IamUserRelationship `json:"User,omitempty"`
	}

	varSessionAbstractSubSessionWithoutEmbeddedStruct := SessionAbstractSubSessionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSessionAbstractSubSessionWithoutEmbeddedStruct)
	if err == nil {
		varSessionAbstractSubSession := _SessionAbstractSubSession{}
		varSessionAbstractSubSession.ClassId = varSessionAbstractSubSessionWithoutEmbeddedStruct.ClassId
		varSessionAbstractSubSession.ObjectType = varSessionAbstractSubSessionWithoutEmbeddedStruct.ObjectType
		varSessionAbstractSubSession.TargetName = varSessionAbstractSubSessionWithoutEmbeddedStruct.TargetName
		varSessionAbstractSubSession.Session = varSessionAbstractSubSessionWithoutEmbeddedStruct.Session
		varSessionAbstractSubSession.Target = varSessionAbstractSubSessionWithoutEmbeddedStruct.Target
		varSessionAbstractSubSession.User = varSessionAbstractSubSessionWithoutEmbeddedStruct.User
		*o = SessionAbstractSubSession(varSessionAbstractSubSession)
	} else {
		return err
	}

	varSessionAbstractSubSession := _SessionAbstractSubSession{}

	err = json.Unmarshal(bytes, &varSessionAbstractSubSession)
	if err == nil {
		o.SessionAbstractSession = varSessionAbstractSubSession.SessionAbstractSession
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetName")
		delete(additionalProperties, "Session")
		delete(additionalProperties, "Target")
		delete(additionalProperties, "User")

		// remove fields from embedded structs
		reflectSessionAbstractSession := reflect.ValueOf(o.SessionAbstractSession)
		for i := 0; i < reflectSessionAbstractSession.Type().NumField(); i++ {
			t := reflectSessionAbstractSession.Type().Field(i)

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

type NullableSessionAbstractSubSession struct {
	value *SessionAbstractSubSession
	isSet bool
}

func (v NullableSessionAbstractSubSession) Get() *SessionAbstractSubSession {
	return v.value
}

func (v *NullableSessionAbstractSubSession) Set(val *SessionAbstractSubSession) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionAbstractSubSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionAbstractSubSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionAbstractSubSession(val *SessionAbstractSubSession) *NullableSessionAbstractSubSession {
	return &NullableSessionAbstractSubSession{value: val, isSet: true}
}

func (v NullableSessionAbstractSubSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionAbstractSubSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


