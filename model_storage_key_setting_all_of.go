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

// StorageKeySettingAllOf Definition of the list of properties defined in 'storage.KeySetting', excluding properties defined in parent classes.
type StorageKeySettingAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Method to be used for fetching the encryption key. * `None` - Drive encryption not configured. * `Manual` - Drive encryption using manual key. * `Kmip` - Remote encryption using KMIP.
	KeyType *string `json:"KeyType,omitempty"`
	ManualKey NullableStorageLocalKeySetting `json:"ManualKey,omitempty"`
	RemoteKey NullableStorageRemoteKeySetting `json:"RemoteKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageKeySettingAllOf StorageKeySettingAllOf

// NewStorageKeySettingAllOf instantiates a new StorageKeySettingAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageKeySettingAllOf(classId string, objectType string) *StorageKeySettingAllOf {
	this := StorageKeySettingAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var keyType string = "None"
	this.KeyType = &keyType
	return &this
}

// NewStorageKeySettingAllOfWithDefaults instantiates a new StorageKeySettingAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageKeySettingAllOfWithDefaults() *StorageKeySettingAllOf {
	this := StorageKeySettingAllOf{}
	var classId string = "storage.KeySetting"
	this.ClassId = classId
	var objectType string = "storage.KeySetting"
	this.ObjectType = objectType
	var keyType string = "None"
	this.KeyType = &keyType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageKeySettingAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageKeySettingAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageKeySettingAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageKeySettingAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageKeySettingAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageKeySettingAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *StorageKeySettingAllOf) GetKeyType() string {
	if o == nil || o.KeyType == nil {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageKeySettingAllOf) GetKeyTypeOk() (*string, bool) {
	if o == nil || o.KeyType == nil {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *StorageKeySettingAllOf) HasKeyType() bool {
	if o != nil && o.KeyType != nil {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *StorageKeySettingAllOf) SetKeyType(v string) {
	o.KeyType = &v
}

// GetManualKey returns the ManualKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageKeySettingAllOf) GetManualKey() StorageLocalKeySetting {
	if o == nil || o.ManualKey.Get() == nil {
		var ret StorageLocalKeySetting
		return ret
	}
	return *o.ManualKey.Get()
}

// GetManualKeyOk returns a tuple with the ManualKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageKeySettingAllOf) GetManualKeyOk() (*StorageLocalKeySetting, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ManualKey.Get(), o.ManualKey.IsSet()
}

// HasManualKey returns a boolean if a field has been set.
func (o *StorageKeySettingAllOf) HasManualKey() bool {
	if o != nil && o.ManualKey.IsSet() {
		return true
	}

	return false
}

// SetManualKey gets a reference to the given NullableStorageLocalKeySetting and assigns it to the ManualKey field.
func (o *StorageKeySettingAllOf) SetManualKey(v StorageLocalKeySetting) {
	o.ManualKey.Set(&v)
}
// SetManualKeyNil sets the value for ManualKey to be an explicit nil
func (o *StorageKeySettingAllOf) SetManualKeyNil() {
	o.ManualKey.Set(nil)
}

// UnsetManualKey ensures that no value is present for ManualKey, not even an explicit nil
func (o *StorageKeySettingAllOf) UnsetManualKey() {
	o.ManualKey.Unset()
}

// GetRemoteKey returns the RemoteKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageKeySettingAllOf) GetRemoteKey() StorageRemoteKeySetting {
	if o == nil || o.RemoteKey.Get() == nil {
		var ret StorageRemoteKeySetting
		return ret
	}
	return *o.RemoteKey.Get()
}

// GetRemoteKeyOk returns a tuple with the RemoteKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageKeySettingAllOf) GetRemoteKeyOk() (*StorageRemoteKeySetting, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteKey.Get(), o.RemoteKey.IsSet()
}

// HasRemoteKey returns a boolean if a field has been set.
func (o *StorageKeySettingAllOf) HasRemoteKey() bool {
	if o != nil && o.RemoteKey.IsSet() {
		return true
	}

	return false
}

// SetRemoteKey gets a reference to the given NullableStorageRemoteKeySetting and assigns it to the RemoteKey field.
func (o *StorageKeySettingAllOf) SetRemoteKey(v StorageRemoteKeySetting) {
	o.RemoteKey.Set(&v)
}
// SetRemoteKeyNil sets the value for RemoteKey to be an explicit nil
func (o *StorageKeySettingAllOf) SetRemoteKeyNil() {
	o.RemoteKey.Set(nil)
}

// UnsetRemoteKey ensures that no value is present for RemoteKey, not even an explicit nil
func (o *StorageKeySettingAllOf) UnsetRemoteKey() {
	o.RemoteKey.Unset()
}

func (o StorageKeySettingAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.KeyType != nil {
		toSerialize["KeyType"] = o.KeyType
	}
	if o.ManualKey.IsSet() {
		toSerialize["ManualKey"] = o.ManualKey.Get()
	}
	if o.RemoteKey.IsSet() {
		toSerialize["RemoteKey"] = o.RemoteKey.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageKeySettingAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageKeySettingAllOf := _StorageKeySettingAllOf{}

	if err = json.Unmarshal(bytes, &varStorageKeySettingAllOf); err == nil {
		*o = StorageKeySettingAllOf(varStorageKeySettingAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "KeyType")
		delete(additionalProperties, "ManualKey")
		delete(additionalProperties, "RemoteKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageKeySettingAllOf struct {
	value *StorageKeySettingAllOf
	isSet bool
}

func (v NullableStorageKeySettingAllOf) Get() *StorageKeySettingAllOf {
	return v.value
}

func (v *NullableStorageKeySettingAllOf) Set(val *StorageKeySettingAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageKeySettingAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageKeySettingAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageKeySettingAllOf(val *StorageKeySettingAllOf) *NullableStorageKeySettingAllOf {
	return &NullableStorageKeySettingAllOf{value: val, isSet: true}
}

func (v NullableStorageKeySettingAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageKeySettingAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


