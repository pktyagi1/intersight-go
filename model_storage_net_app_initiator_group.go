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

// StorageNetAppInitiatorGroup NetApp Initiator Group specifies host access to LUNs on the storage system.
type StorageNetAppInitiatorGroup struct {
	StorageBaseHost
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Initiator group protocol. * `FCP` - Fibre channel initiator type which contains WWN of an HBA on the host. * `iSCSI` - An iSCSI initiator type used by the host. * `mixed` - For systems using both FC and iSCSI connections to the same LUN, create two igroups, one for FC and one for iSCSI. Then map the LUN to both igroups.
	Protocol *string `json:"Protocol,omitempty"`
	// Universally unique identifier of the LUN.
	Uuid *string `json:"Uuid,omitempty"`
	Tenant *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppInitiatorGroup StorageNetAppInitiatorGroup

// NewStorageNetAppInitiatorGroup instantiates a new StorageNetAppInitiatorGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppInitiatorGroup(classId string, objectType string) *StorageNetAppInitiatorGroup {
	this := StorageNetAppInitiatorGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppInitiatorGroupWithDefaults instantiates a new StorageNetAppInitiatorGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppInitiatorGroupWithDefaults() *StorageNetAppInitiatorGroup {
	this := StorageNetAppInitiatorGroup{}
	var classId string = "storage.NetAppInitiatorGroup"
	this.ClassId = classId
	var objectType string = "storage.NetAppInitiatorGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppInitiatorGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppInitiatorGroup) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppInitiatorGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppInitiatorGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppInitiatorGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppInitiatorGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *StorageNetAppInitiatorGroup) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppInitiatorGroup) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *StorageNetAppInitiatorGroup) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *StorageNetAppInitiatorGroup) SetProtocol(v string) {
	o.Protocol = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppInitiatorGroup) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppInitiatorGroup) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppInitiatorGroup) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppInitiatorGroup) SetUuid(v string) {
	o.Uuid = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppInitiatorGroup) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppInitiatorGroup) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppInitiatorGroup) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppInitiatorGroup) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppInitiatorGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseHost, errStorageBaseHost := json.Marshal(o.StorageBaseHost)
	if errStorageBaseHost != nil {
		return []byte{}, errStorageBaseHost
	}
	errStorageBaseHost = json.Unmarshal([]byte(serializedStorageBaseHost), &toSerialize)
	if errStorageBaseHost != nil {
		return []byte{}, errStorageBaseHost
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppInitiatorGroup) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppInitiatorGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Initiator group protocol. * `FCP` - Fibre channel initiator type which contains WWN of an HBA on the host. * `iSCSI` - An iSCSI initiator type used by the host. * `mixed` - For systems using both FC and iSCSI connections to the same LUN, create two igroups, one for FC and one for iSCSI. Then map the LUN to both igroups.
		Protocol *string `json:"Protocol,omitempty"`
		// Universally unique identifier of the LUN.
		Uuid *string `json:"Uuid,omitempty"`
		Tenant *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	}

	varStorageNetAppInitiatorGroupWithoutEmbeddedStruct := StorageNetAppInitiatorGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppInitiatorGroupWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppInitiatorGroup := _StorageNetAppInitiatorGroup{}
		varStorageNetAppInitiatorGroup.ClassId = varStorageNetAppInitiatorGroupWithoutEmbeddedStruct.ClassId
		varStorageNetAppInitiatorGroup.ObjectType = varStorageNetAppInitiatorGroupWithoutEmbeddedStruct.ObjectType
		varStorageNetAppInitiatorGroup.Protocol = varStorageNetAppInitiatorGroupWithoutEmbeddedStruct.Protocol
		varStorageNetAppInitiatorGroup.Uuid = varStorageNetAppInitiatorGroupWithoutEmbeddedStruct.Uuid
		varStorageNetAppInitiatorGroup.Tenant = varStorageNetAppInitiatorGroupWithoutEmbeddedStruct.Tenant
		*o = StorageNetAppInitiatorGroup(varStorageNetAppInitiatorGroup)
	} else {
		return err
	}

	varStorageNetAppInitiatorGroup := _StorageNetAppInitiatorGroup{}

	err = json.Unmarshal(bytes, &varStorageNetAppInitiatorGroup)
	if err == nil {
		o.StorageBaseHost = varStorageNetAppInitiatorGroup.StorageBaseHost
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Tenant")

		// remove fields from embedded structs
		reflectStorageBaseHost := reflect.ValueOf(o.StorageBaseHost)
		for i := 0; i < reflectStorageBaseHost.Type().NumField(); i++ {
			t := reflectStorageBaseHost.Type().Field(i)

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

type NullableStorageNetAppInitiatorGroup struct {
	value *StorageNetAppInitiatorGroup
	isSet bool
}

func (v NullableStorageNetAppInitiatorGroup) Get() *StorageNetAppInitiatorGroup {
	return v.value
}

func (v *NullableStorageNetAppInitiatorGroup) Set(val *StorageNetAppInitiatorGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppInitiatorGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppInitiatorGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppInitiatorGroup(val *StorageNetAppInitiatorGroup) *NullableStorageNetAppInitiatorGroup {
	return &NullableStorageNetAppInitiatorGroup{value: val, isSet: true}
}

func (v NullableStorageNetAppInitiatorGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppInitiatorGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


