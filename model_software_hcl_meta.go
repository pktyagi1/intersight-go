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
	"reflect"
	"strings"
)

// SoftwareHclMeta A JSON file wth HCL metadata uploaded for consumption by the HCL service.
type SoftwareHclMeta struct {
	FirmwareBaseDistributable
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The type of content that the Json file holds (Incremental or full dump). * `Full` - Indicates that the JSON File does have full content for HCL metadata. * `Incremental` - Indicates that the JSON File does have only the diff of the Hcl meta to be uploaded.
	ContentType *string `json:"ContentType,omitempty"`
	Catalog *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwareHclMeta SoftwareHclMeta

// NewSoftwareHclMeta instantiates a new SoftwareHclMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareHclMeta(classId string, objectType string) *SoftwareHclMeta {
	this := SoftwareHclMeta{}
	this.ClassId = classId
	this.ObjectType = objectType
	var importAction string = "None"
	this.ImportAction = &importAction
	var vendor string = "Cisco"
	this.Vendor = &vendor
	var contentType string = "Full"
	this.ContentType = &contentType
	return &this
}

// NewSoftwareHclMetaWithDefaults instantiates a new SoftwareHclMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareHclMetaWithDefaults() *SoftwareHclMeta {
	this := SoftwareHclMeta{}
	var classId string = "software.HclMeta"
	this.ClassId = classId
	var objectType string = "software.HclMeta"
	this.ObjectType = objectType
	var contentType string = "Full"
	this.ContentType = &contentType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwareHclMeta) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwareHclMeta) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwareHclMeta) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwareHclMeta) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwareHclMeta) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwareHclMeta) SetObjectType(v string) {
	o.ObjectType = v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *SoftwareHclMeta) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareHclMeta) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *SoftwareHclMeta) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *SoftwareHclMeta) SetContentType(v string) {
	o.ContentType = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *SoftwareHclMeta) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareHclMeta) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwareHclMeta) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwareHclMeta) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

func (o SoftwareHclMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareBaseDistributable, errFirmwareBaseDistributable := json.Marshal(o.FirmwareBaseDistributable)
	if errFirmwareBaseDistributable != nil {
		return []byte{}, errFirmwareBaseDistributable
	}
	errFirmwareBaseDistributable = json.Unmarshal([]byte(serializedFirmwareBaseDistributable), &toSerialize)
	if errFirmwareBaseDistributable != nil {
		return []byte{}, errFirmwareBaseDistributable
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ContentType != nil {
		toSerialize["ContentType"] = o.ContentType
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwareHclMeta) UnmarshalJSON(bytes []byte) (err error) {
	type SoftwareHclMetaWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The type of content that the Json file holds (Incremental or full dump). * `Full` - Indicates that the JSON File does have full content for HCL metadata. * `Incremental` - Indicates that the JSON File does have only the diff of the Hcl meta to be uploaded.
		ContentType *string `json:"ContentType,omitempty"`
		Catalog *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	}

	varSoftwareHclMetaWithoutEmbeddedStruct := SoftwareHclMetaWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSoftwareHclMetaWithoutEmbeddedStruct)
	if err == nil {
		varSoftwareHclMeta := _SoftwareHclMeta{}
		varSoftwareHclMeta.ClassId = varSoftwareHclMetaWithoutEmbeddedStruct.ClassId
		varSoftwareHclMeta.ObjectType = varSoftwareHclMetaWithoutEmbeddedStruct.ObjectType
		varSoftwareHclMeta.ContentType = varSoftwareHclMetaWithoutEmbeddedStruct.ContentType
		varSoftwareHclMeta.Catalog = varSoftwareHclMetaWithoutEmbeddedStruct.Catalog
		*o = SoftwareHclMeta(varSoftwareHclMeta)
	} else {
		return err
	}

	varSoftwareHclMeta := _SoftwareHclMeta{}

	err = json.Unmarshal(bytes, &varSoftwareHclMeta)
	if err == nil {
		o.FirmwareBaseDistributable = varSoftwareHclMeta.FirmwareBaseDistributable
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ContentType")
		delete(additionalProperties, "Catalog")

		// remove fields from embedded structs
		reflectFirmwareBaseDistributable := reflect.ValueOf(o.FirmwareBaseDistributable)
		for i := 0; i < reflectFirmwareBaseDistributable.Type().NumField(); i++ {
			t := reflectFirmwareBaseDistributable.Type().Field(i)

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

type NullableSoftwareHclMeta struct {
	value *SoftwareHclMeta
	isSet bool
}

func (v NullableSoftwareHclMeta) Get() *SoftwareHclMeta {
	return v.value
}

func (v *NullableSoftwareHclMeta) Set(val *SoftwareHclMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareHclMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareHclMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareHclMeta(val *SoftwareHclMeta) *NullableSoftwareHclMeta {
	return &NullableSoftwareHclMeta{value: val, isSet: true}
}

func (v NullableSoftwareHclMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareHclMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


