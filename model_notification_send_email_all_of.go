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
)

// NotificationSendEmailAllOf Definition of the list of properties defined in 'notification.SendEmail', excluding properties defined in parent classes.
type NotificationSendEmailAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Email of the recipient, who expects to be notified about the event that happens in Intersight.
	Email *string `json:"Email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationSendEmailAllOf NotificationSendEmailAllOf

// NewNotificationSendEmailAllOf instantiates a new NotificationSendEmailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSendEmailAllOf(classId string, objectType string) *NotificationSendEmailAllOf {
	this := NotificationSendEmailAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNotificationSendEmailAllOfWithDefaults instantiates a new NotificationSendEmailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSendEmailAllOfWithDefaults() *NotificationSendEmailAllOf {
	this := NotificationSendEmailAllOf{}
	var classId string = "notification.SendEmail"
	this.ClassId = classId
	var objectType string = "notification.SendEmail"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NotificationSendEmailAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NotificationSendEmailAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NotificationSendEmailAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NotificationSendEmailAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NotificationSendEmailAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NotificationSendEmailAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *NotificationSendEmailAllOf) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSendEmailAllOf) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *NotificationSendEmailAllOf) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *NotificationSendEmailAllOf) SetEmail(v string) {
	o.Email = &v
}

func (o NotificationSendEmailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Email != nil {
		toSerialize["Email"] = o.Email
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NotificationSendEmailAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNotificationSendEmailAllOf := _NotificationSendEmailAllOf{}

	if err = json.Unmarshal(bytes, &varNotificationSendEmailAllOf); err == nil {
		*o = NotificationSendEmailAllOf(varNotificationSendEmailAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationSendEmailAllOf struct {
	value *NotificationSendEmailAllOf
	isSet bool
}

func (v NullableNotificationSendEmailAllOf) Get() *NotificationSendEmailAllOf {
	return v.value
}

func (v *NullableNotificationSendEmailAllOf) Set(val *NotificationSendEmailAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSendEmailAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSendEmailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSendEmailAllOf(val *NotificationSendEmailAllOf) *NullableNotificationSendEmailAllOf {
	return &NullableNotificationSendEmailAllOf{value: val, isSet: true}
}

func (v NullableNotificationSendEmailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSendEmailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


