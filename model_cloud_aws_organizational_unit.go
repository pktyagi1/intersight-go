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

// CloudAwsOrganizationalUnit AWS organization unit is represented here. It is a Logical grouping of accounts in a AWS organization.
type CloudAwsOrganizationalUnit struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The identity of this organization. This entity is not manipulated by users. It aids in uniquely identifying the organization object.
	Identity *string `json:"Identity,omitempty"`
	// Name of the organizational unit.
	Name *string `json:"Name,omitempty"`
	ParentOrg *CloudAwsOrganizationalUnitRelationship `json:"ParentOrg,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudAwsOrganizationalUnit CloudAwsOrganizationalUnit

// NewCloudAwsOrganizationalUnit instantiates a new CloudAwsOrganizationalUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAwsOrganizationalUnit(classId string, objectType string) *CloudAwsOrganizationalUnit {
	this := CloudAwsOrganizationalUnit{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudAwsOrganizationalUnitWithDefaults instantiates a new CloudAwsOrganizationalUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAwsOrganizationalUnitWithDefaults() *CloudAwsOrganizationalUnit {
	this := CloudAwsOrganizationalUnit{}
	var classId string = "cloud.AwsOrganizationalUnit"
	this.ClassId = classId
	var objectType string = "cloud.AwsOrganizationalUnit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudAwsOrganizationalUnit) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudAwsOrganizationalUnit) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudAwsOrganizationalUnit) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudAwsOrganizationalUnit) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudAwsOrganizationalUnit) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudAwsOrganizationalUnit) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CloudAwsOrganizationalUnit) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsOrganizationalUnit) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CloudAwsOrganizationalUnit) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *CloudAwsOrganizationalUnit) SetIdentity(v string) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudAwsOrganizationalUnit) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsOrganizationalUnit) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudAwsOrganizationalUnit) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudAwsOrganizationalUnit) SetName(v string) {
	o.Name = &v
}

// GetParentOrg returns the ParentOrg field value if set, zero value otherwise.
func (o *CloudAwsOrganizationalUnit) GetParentOrg() CloudAwsOrganizationalUnitRelationship {
	if o == nil || o.ParentOrg == nil {
		var ret CloudAwsOrganizationalUnitRelationship
		return ret
	}
	return *o.ParentOrg
}

// GetParentOrgOk returns a tuple with the ParentOrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsOrganizationalUnit) GetParentOrgOk() (*CloudAwsOrganizationalUnitRelationship, bool) {
	if o == nil || o.ParentOrg == nil {
		return nil, false
	}
	return o.ParentOrg, true
}

// HasParentOrg returns a boolean if a field has been set.
func (o *CloudAwsOrganizationalUnit) HasParentOrg() bool {
	if o != nil && o.ParentOrg != nil {
		return true
	}

	return false
}

// SetParentOrg gets a reference to the given CloudAwsOrganizationalUnitRelationship and assigns it to the ParentOrg field.
func (o *CloudAwsOrganizationalUnit) SetParentOrg(v CloudAwsOrganizationalUnitRelationship) {
	o.ParentOrg = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *CloudAwsOrganizationalUnit) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsOrganizationalUnit) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *CloudAwsOrganizationalUnit) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *CloudAwsOrganizationalUnit) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o CloudAwsOrganizationalUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ParentOrg != nil {
		toSerialize["ParentOrg"] = o.ParentOrg
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudAwsOrganizationalUnit) UnmarshalJSON(bytes []byte) (err error) {
	type CloudAwsOrganizationalUnitWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The identity of this organization. This entity is not manipulated by users. It aids in uniquely identifying the organization object.
		Identity *string `json:"Identity,omitempty"`
		// Name of the organizational unit.
		Name *string `json:"Name,omitempty"`
		ParentOrg *CloudAwsOrganizationalUnitRelationship `json:"ParentOrg,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varCloudAwsOrganizationalUnitWithoutEmbeddedStruct := CloudAwsOrganizationalUnitWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudAwsOrganizationalUnitWithoutEmbeddedStruct)
	if err == nil {
		varCloudAwsOrganizationalUnit := _CloudAwsOrganizationalUnit{}
		varCloudAwsOrganizationalUnit.ClassId = varCloudAwsOrganizationalUnitWithoutEmbeddedStruct.ClassId
		varCloudAwsOrganizationalUnit.ObjectType = varCloudAwsOrganizationalUnitWithoutEmbeddedStruct.ObjectType
		varCloudAwsOrganizationalUnit.Identity = varCloudAwsOrganizationalUnitWithoutEmbeddedStruct.Identity
		varCloudAwsOrganizationalUnit.Name = varCloudAwsOrganizationalUnitWithoutEmbeddedStruct.Name
		varCloudAwsOrganizationalUnit.ParentOrg = varCloudAwsOrganizationalUnitWithoutEmbeddedStruct.ParentOrg
		varCloudAwsOrganizationalUnit.RegisteredDevice = varCloudAwsOrganizationalUnitWithoutEmbeddedStruct.RegisteredDevice
		*o = CloudAwsOrganizationalUnit(varCloudAwsOrganizationalUnit)
	} else {
		return err
	}

	varCloudAwsOrganizationalUnit := _CloudAwsOrganizationalUnit{}

	err = json.Unmarshal(bytes, &varCloudAwsOrganizationalUnit)
	if err == nil {
		o.MoBaseMo = varCloudAwsOrganizationalUnit.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ParentOrg")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableCloudAwsOrganizationalUnit struct {
	value *CloudAwsOrganizationalUnit
	isSet bool
}

func (v NullableCloudAwsOrganizationalUnit) Get() *CloudAwsOrganizationalUnit {
	return v.value
}

func (v *NullableCloudAwsOrganizationalUnit) Set(val *CloudAwsOrganizationalUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAwsOrganizationalUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAwsOrganizationalUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAwsOrganizationalUnit(val *CloudAwsOrganizationalUnit) *NullableCloudAwsOrganizationalUnit {
	return &NullableCloudAwsOrganizationalUnit{value: val, isSet: true}
}

func (v NullableCloudAwsOrganizationalUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAwsOrganizationalUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


