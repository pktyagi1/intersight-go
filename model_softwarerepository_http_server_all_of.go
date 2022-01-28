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

// SoftwarerepositoryHttpServerAllOf Definition of the list of properties defined in 'softwarerepository.HttpServer', excluding properties defined in parent classes.
type SoftwarerepositoryHttpServerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// HTTP/HTTPS link to the image. Accepted formats are HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image.
	LocationLink *string `json:"LocationLink,omitempty"`
	// Password as configured on the HTTP[S] server for user authentication. It is generally required to authenticate user provided HTTP[S] based software repositories.
	Password *string `json:"Password,omitempty"`
	// Username as configured on the HTTP[S] server for user authentication. It is generally required to authenticate user provided HTTP[S] based software repositories.
	Username *string `json:"Username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryHttpServerAllOf SoftwarerepositoryHttpServerAllOf

// NewSoftwarerepositoryHttpServerAllOf instantiates a new SoftwarerepositoryHttpServerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryHttpServerAllOf(classId string, objectType string) *SoftwarerepositoryHttpServerAllOf {
	this := SoftwarerepositoryHttpServerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwarerepositoryHttpServerAllOfWithDefaults instantiates a new SoftwarerepositoryHttpServerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryHttpServerAllOfWithDefaults() *SoftwarerepositoryHttpServerAllOf {
	this := SoftwarerepositoryHttpServerAllOf{}
	var classId string = "softwarerepository.HttpServer"
	this.ClassId = classId
	var objectType string = "softwarerepository.HttpServer"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryHttpServerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryHttpServerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryHttpServerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryHttpServerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryHttpServerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryHttpServerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *SoftwarerepositoryHttpServerAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryHttpServerAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *SoftwarerepositoryHttpServerAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *SoftwarerepositoryHttpServerAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetLocationLink returns the LocationLink field value if set, zero value otherwise.
func (o *SoftwarerepositoryHttpServerAllOf) GetLocationLink() string {
	if o == nil || o.LocationLink == nil {
		var ret string
		return ret
	}
	return *o.LocationLink
}

// GetLocationLinkOk returns a tuple with the LocationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryHttpServerAllOf) GetLocationLinkOk() (*string, bool) {
	if o == nil || o.LocationLink == nil {
		return nil, false
	}
	return o.LocationLink, true
}

// HasLocationLink returns a boolean if a field has been set.
func (o *SoftwarerepositoryHttpServerAllOf) HasLocationLink() bool {
	if o != nil && o.LocationLink != nil {
		return true
	}

	return false
}

// SetLocationLink gets a reference to the given string and assigns it to the LocationLink field.
func (o *SoftwarerepositoryHttpServerAllOf) SetLocationLink(v string) {
	o.LocationLink = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SoftwarerepositoryHttpServerAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryHttpServerAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SoftwarerepositoryHttpServerAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SoftwarerepositoryHttpServerAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SoftwarerepositoryHttpServerAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryHttpServerAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SoftwarerepositoryHttpServerAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SoftwarerepositoryHttpServerAllOf) SetUsername(v string) {
	o.Username = &v
}

func (o SoftwarerepositoryHttpServerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.LocationLink != nil {
		toSerialize["LocationLink"] = o.LocationLink
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryHttpServerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwarerepositoryHttpServerAllOf := _SoftwarerepositoryHttpServerAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwarerepositoryHttpServerAllOf); err == nil {
		*o = SoftwarerepositoryHttpServerAllOf(varSoftwarerepositoryHttpServerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "LocationLink")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwarerepositoryHttpServerAllOf struct {
	value *SoftwarerepositoryHttpServerAllOf
	isSet bool
}

func (v NullableSoftwarerepositoryHttpServerAllOf) Get() *SoftwarerepositoryHttpServerAllOf {
	return v.value
}

func (v *NullableSoftwarerepositoryHttpServerAllOf) Set(val *SoftwarerepositoryHttpServerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryHttpServerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryHttpServerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryHttpServerAllOf(val *SoftwarerepositoryHttpServerAllOf) *NullableSoftwarerepositoryHttpServerAllOf {
	return &NullableSoftwarerepositoryHttpServerAllOf{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryHttpServerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryHttpServerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


