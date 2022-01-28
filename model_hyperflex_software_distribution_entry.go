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

// HyperflexSoftwareDistributionEntry A HyperFlex Software Distribution.
type HyperflexSoftwareDistributionEntry struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The HyperFlex Software Distribution type.
	DistributionType *string `json:"DistributionType,omitempty"`
	AppCatalog *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	// An array of relationships to hyperflexSoftwareDistributionVersion resources.
	DistributionVersions []HyperflexSoftwareDistributionVersionRelationship `json:"DistributionVersions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexSoftwareDistributionEntry HyperflexSoftwareDistributionEntry

// NewHyperflexSoftwareDistributionEntry instantiates a new HyperflexSoftwareDistributionEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSoftwareDistributionEntry(classId string, objectType string) *HyperflexSoftwareDistributionEntry {
	this := HyperflexSoftwareDistributionEntry{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexSoftwareDistributionEntryWithDefaults instantiates a new HyperflexSoftwareDistributionEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSoftwareDistributionEntryWithDefaults() *HyperflexSoftwareDistributionEntry {
	this := HyperflexSoftwareDistributionEntry{}
	var classId string = "hyperflex.SoftwareDistributionEntry"
	this.ClassId = classId
	var objectType string = "hyperflex.SoftwareDistributionEntry"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexSoftwareDistributionEntry) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionEntry) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexSoftwareDistributionEntry) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexSoftwareDistributionEntry) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionEntry) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexSoftwareDistributionEntry) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDistributionType returns the DistributionType field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionEntry) GetDistributionType() string {
	if o == nil || o.DistributionType == nil {
		var ret string
		return ret
	}
	return *o.DistributionType
}

// GetDistributionTypeOk returns a tuple with the DistributionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionEntry) GetDistributionTypeOk() (*string, bool) {
	if o == nil || o.DistributionType == nil {
		return nil, false
	}
	return o.DistributionType, true
}

// HasDistributionType returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionEntry) HasDistributionType() bool {
	if o != nil && o.DistributionType != nil {
		return true
	}

	return false
}

// SetDistributionType gets a reference to the given string and assigns it to the DistributionType field.
func (o *HyperflexSoftwareDistributionEntry) SetDistributionType(v string) {
	o.DistributionType = &v
}

// GetAppCatalog returns the AppCatalog field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionEntry) GetAppCatalog() HyperflexAppCatalogRelationship {
	if o == nil || o.AppCatalog == nil {
		var ret HyperflexAppCatalogRelationship
		return ret
	}
	return *o.AppCatalog
}

// GetAppCatalogOk returns a tuple with the AppCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionEntry) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool) {
	if o == nil || o.AppCatalog == nil {
		return nil, false
	}
	return o.AppCatalog, true
}

// HasAppCatalog returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionEntry) HasAppCatalog() bool {
	if o != nil && o.AppCatalog != nil {
		return true
	}

	return false
}

// SetAppCatalog gets a reference to the given HyperflexAppCatalogRelationship and assigns it to the AppCatalog field.
func (o *HyperflexSoftwareDistributionEntry) SetAppCatalog(v HyperflexAppCatalogRelationship) {
	o.AppCatalog = &v
}

// GetDistributionVersions returns the DistributionVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSoftwareDistributionEntry) GetDistributionVersions() []HyperflexSoftwareDistributionVersionRelationship {
	if o == nil  {
		var ret []HyperflexSoftwareDistributionVersionRelationship
		return ret
	}
	return o.DistributionVersions
}

// GetDistributionVersionsOk returns a tuple with the DistributionVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSoftwareDistributionEntry) GetDistributionVersionsOk() (*[]HyperflexSoftwareDistributionVersionRelationship, bool) {
	if o == nil || o.DistributionVersions == nil {
		return nil, false
	}
	return &o.DistributionVersions, true
}

// HasDistributionVersions returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionEntry) HasDistributionVersions() bool {
	if o != nil && o.DistributionVersions != nil {
		return true
	}

	return false
}

// SetDistributionVersions gets a reference to the given []HyperflexSoftwareDistributionVersionRelationship and assigns it to the DistributionVersions field.
func (o *HyperflexSoftwareDistributionEntry) SetDistributionVersions(v []HyperflexSoftwareDistributionVersionRelationship) {
	o.DistributionVersions = v
}

func (o HyperflexSoftwareDistributionEntry) MarshalJSON() ([]byte, error) {
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
	if o.DistributionType != nil {
		toSerialize["DistributionType"] = o.DistributionType
	}
	if o.AppCatalog != nil {
		toSerialize["AppCatalog"] = o.AppCatalog
	}
	if o.DistributionVersions != nil {
		toSerialize["DistributionVersions"] = o.DistributionVersions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexSoftwareDistributionEntry) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexSoftwareDistributionEntryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The HyperFlex Software Distribution type.
		DistributionType *string `json:"DistributionType,omitempty"`
		AppCatalog *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
		// An array of relationships to hyperflexSoftwareDistributionVersion resources.
		DistributionVersions []HyperflexSoftwareDistributionVersionRelationship `json:"DistributionVersions,omitempty"`
	}

	varHyperflexSoftwareDistributionEntryWithoutEmbeddedStruct := HyperflexSoftwareDistributionEntryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexSoftwareDistributionEntryWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexSoftwareDistributionEntry := _HyperflexSoftwareDistributionEntry{}
		varHyperflexSoftwareDistributionEntry.ClassId = varHyperflexSoftwareDistributionEntryWithoutEmbeddedStruct.ClassId
		varHyperflexSoftwareDistributionEntry.ObjectType = varHyperflexSoftwareDistributionEntryWithoutEmbeddedStruct.ObjectType
		varHyperflexSoftwareDistributionEntry.DistributionType = varHyperflexSoftwareDistributionEntryWithoutEmbeddedStruct.DistributionType
		varHyperflexSoftwareDistributionEntry.AppCatalog = varHyperflexSoftwareDistributionEntryWithoutEmbeddedStruct.AppCatalog
		varHyperflexSoftwareDistributionEntry.DistributionVersions = varHyperflexSoftwareDistributionEntryWithoutEmbeddedStruct.DistributionVersions
		*o = HyperflexSoftwareDistributionEntry(varHyperflexSoftwareDistributionEntry)
	} else {
		return err
	}

	varHyperflexSoftwareDistributionEntry := _HyperflexSoftwareDistributionEntry{}

	err = json.Unmarshal(bytes, &varHyperflexSoftwareDistributionEntry)
	if err == nil {
		o.MoBaseMo = varHyperflexSoftwareDistributionEntry.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DistributionType")
		delete(additionalProperties, "AppCatalog")
		delete(additionalProperties, "DistributionVersions")

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

type NullableHyperflexSoftwareDistributionEntry struct {
	value *HyperflexSoftwareDistributionEntry
	isSet bool
}

func (v NullableHyperflexSoftwareDistributionEntry) Get() *HyperflexSoftwareDistributionEntry {
	return v.value
}

func (v *NullableHyperflexSoftwareDistributionEntry) Set(val *HyperflexSoftwareDistributionEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSoftwareDistributionEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSoftwareDistributionEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSoftwareDistributionEntry(val *HyperflexSoftwareDistributionEntry) *NullableHyperflexSoftwareDistributionEntry {
	return &NullableHyperflexSoftwareDistributionEntry{value: val, isSet: true}
}

func (v NullableHyperflexSoftwareDistributionEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSoftwareDistributionEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


