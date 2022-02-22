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

// StorageLocalKeySetting Models the local key configurarion required for disks encryptions.
type StorageLocalKeySetting struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Existing key which is already configured on the server.
	ExistingKey *string `json:"ExistingKey,omitempty"`
	// Indicates whether the value of the 'existingKey' property has been set.
	IsExistingKeySet *bool `json:"IsExistingKeySet,omitempty"`
	// Indicates whether the value of the 'newKey' property has been set.
	IsNewKeySet *bool `json:"IsNewKeySet,omitempty"`
	// New key to be configured on the controller.
	NewKey *string `json:"NewKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageLocalKeySetting StorageLocalKeySetting

// NewStorageLocalKeySetting instantiates a new StorageLocalKeySetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageLocalKeySetting(classId string, objectType string) *StorageLocalKeySetting {
	this := StorageLocalKeySetting{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageLocalKeySettingWithDefaults instantiates a new StorageLocalKeySetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageLocalKeySettingWithDefaults() *StorageLocalKeySetting {
	this := StorageLocalKeySetting{}
	var classId string = "storage.LocalKeySetting"
	this.ClassId = classId
	var objectType string = "storage.LocalKeySetting"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageLocalKeySetting) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageLocalKeySetting) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageLocalKeySetting) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageLocalKeySetting) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageLocalKeySetting) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageLocalKeySetting) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExistingKey returns the ExistingKey field value if set, zero value otherwise.
func (o *StorageLocalKeySetting) GetExistingKey() string {
	if o == nil || o.ExistingKey == nil {
		var ret string
		return ret
	}
	return *o.ExistingKey
}

// GetExistingKeyOk returns a tuple with the ExistingKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageLocalKeySetting) GetExistingKeyOk() (*string, bool) {
	if o == nil || o.ExistingKey == nil {
		return nil, false
	}
	return o.ExistingKey, true
}

// HasExistingKey returns a boolean if a field has been set.
func (o *StorageLocalKeySetting) HasExistingKey() bool {
	if o != nil && o.ExistingKey != nil {
		return true
	}

	return false
}

// SetExistingKey gets a reference to the given string and assigns it to the ExistingKey field.
func (o *StorageLocalKeySetting) SetExistingKey(v string) {
	o.ExistingKey = &v
}

// GetIsExistingKeySet returns the IsExistingKeySet field value if set, zero value otherwise.
func (o *StorageLocalKeySetting) GetIsExistingKeySet() bool {
	if o == nil || o.IsExistingKeySet == nil {
		var ret bool
		return ret
	}
	return *o.IsExistingKeySet
}

// GetIsExistingKeySetOk returns a tuple with the IsExistingKeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageLocalKeySetting) GetIsExistingKeySetOk() (*bool, bool) {
	if o == nil || o.IsExistingKeySet == nil {
		return nil, false
	}
	return o.IsExistingKeySet, true
}

// HasIsExistingKeySet returns a boolean if a field has been set.
func (o *StorageLocalKeySetting) HasIsExistingKeySet() bool {
	if o != nil && o.IsExistingKeySet != nil {
		return true
	}

	return false
}

// SetIsExistingKeySet gets a reference to the given bool and assigns it to the IsExistingKeySet field.
func (o *StorageLocalKeySetting) SetIsExistingKeySet(v bool) {
	o.IsExistingKeySet = &v
}

// GetIsNewKeySet returns the IsNewKeySet field value if set, zero value otherwise.
func (o *StorageLocalKeySetting) GetIsNewKeySet() bool {
	if o == nil || o.IsNewKeySet == nil {
		var ret bool
		return ret
	}
	return *o.IsNewKeySet
}

// GetIsNewKeySetOk returns a tuple with the IsNewKeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageLocalKeySetting) GetIsNewKeySetOk() (*bool, bool) {
	if o == nil || o.IsNewKeySet == nil {
		return nil, false
	}
	return o.IsNewKeySet, true
}

// HasIsNewKeySet returns a boolean if a field has been set.
func (o *StorageLocalKeySetting) HasIsNewKeySet() bool {
	if o != nil && o.IsNewKeySet != nil {
		return true
	}

	return false
}

// SetIsNewKeySet gets a reference to the given bool and assigns it to the IsNewKeySet field.
func (o *StorageLocalKeySetting) SetIsNewKeySet(v bool) {
	o.IsNewKeySet = &v
}

// GetNewKey returns the NewKey field value if set, zero value otherwise.
func (o *StorageLocalKeySetting) GetNewKey() string {
	if o == nil || o.NewKey == nil {
		var ret string
		return ret
	}
	return *o.NewKey
}

// GetNewKeyOk returns a tuple with the NewKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageLocalKeySetting) GetNewKeyOk() (*string, bool) {
	if o == nil || o.NewKey == nil {
		return nil, false
	}
	return o.NewKey, true
}

// HasNewKey returns a boolean if a field has been set.
func (o *StorageLocalKeySetting) HasNewKey() bool {
	if o != nil && o.NewKey != nil {
		return true
	}

	return false
}

// SetNewKey gets a reference to the given string and assigns it to the NewKey field.
func (o *StorageLocalKeySetting) SetNewKey(v string) {
	o.NewKey = &v
}

func (o StorageLocalKeySetting) MarshalJSON() ([]byte, error) {
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
	if o.ExistingKey != nil {
		toSerialize["ExistingKey"] = o.ExistingKey
	}
	if o.IsExistingKeySet != nil {
		toSerialize["IsExistingKeySet"] = o.IsExistingKeySet
	}
	if o.IsNewKeySet != nil {
		toSerialize["IsNewKeySet"] = o.IsNewKeySet
	}
	if o.NewKey != nil {
		toSerialize["NewKey"] = o.NewKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageLocalKeySetting) UnmarshalJSON(bytes []byte) (err error) {
	type StorageLocalKeySettingWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Existing key which is already configured on the server.
		ExistingKey *string `json:"ExistingKey,omitempty"`
		// Indicates whether the value of the 'existingKey' property has been set.
		IsExistingKeySet *bool `json:"IsExistingKeySet,omitempty"`
		// Indicates whether the value of the 'newKey' property has been set.
		IsNewKeySet *bool `json:"IsNewKeySet,omitempty"`
		// New key to be configured on the controller.
		NewKey *string `json:"NewKey,omitempty"`
	}

	varStorageLocalKeySettingWithoutEmbeddedStruct := StorageLocalKeySettingWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageLocalKeySettingWithoutEmbeddedStruct)
	if err == nil {
		varStorageLocalKeySetting := _StorageLocalKeySetting{}
		varStorageLocalKeySetting.ClassId = varStorageLocalKeySettingWithoutEmbeddedStruct.ClassId
		varStorageLocalKeySetting.ObjectType = varStorageLocalKeySettingWithoutEmbeddedStruct.ObjectType
		varStorageLocalKeySetting.ExistingKey = varStorageLocalKeySettingWithoutEmbeddedStruct.ExistingKey
		varStorageLocalKeySetting.IsExistingKeySet = varStorageLocalKeySettingWithoutEmbeddedStruct.IsExistingKeySet
		varStorageLocalKeySetting.IsNewKeySet = varStorageLocalKeySettingWithoutEmbeddedStruct.IsNewKeySet
		varStorageLocalKeySetting.NewKey = varStorageLocalKeySettingWithoutEmbeddedStruct.NewKey
		*o = StorageLocalKeySetting(varStorageLocalKeySetting)
	} else {
		return err
	}

	varStorageLocalKeySetting := _StorageLocalKeySetting{}

	err = json.Unmarshal(bytes, &varStorageLocalKeySetting)
	if err == nil {
		o.MoBaseComplexType = varStorageLocalKeySetting.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExistingKey")
		delete(additionalProperties, "IsExistingKeySet")
		delete(additionalProperties, "IsNewKeySet")
		delete(additionalProperties, "NewKey")

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

type NullableStorageLocalKeySetting struct {
	value *StorageLocalKeySetting
	isSet bool
}

func (v NullableStorageLocalKeySetting) Get() *StorageLocalKeySetting {
	return v.value
}

func (v *NullableStorageLocalKeySetting) Set(val *StorageLocalKeySetting) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageLocalKeySetting) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageLocalKeySetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageLocalKeySetting(val *StorageLocalKeySetting) *NullableStorageLocalKeySetting {
	return &NullableStorageLocalKeySetting{value: val, isSet: true}
}

func (v NullableStorageLocalKeySetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageLocalKeySetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


