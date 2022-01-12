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
	"time"
	"reflect"
	"strings"
)

// NiaapiNibMetadata Contains the latest metadata available for download from server.
type NiaapiNibMetadata struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	Content []NiaapiDetail `json:"Content,omitempty"`
	// The date when the package was generated.
	Date *time.Time `json:"Date,omitempty"`
	// Chksum used to check the integrity of the metadata file downloaded.
	MetadataChksum *string `json:"MetadataChksum,omitempty"`
	// The filename of the metadata package.
	MetadataFilename *string `json:"MetadataFilename,omitempty"`
	// The version number of the metadata package.
	Version *int64 `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiNibMetadata NiaapiNibMetadata

// NewNiaapiNibMetadata instantiates a new NiaapiNibMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiNibMetadata(classId string, objectType string) *NiaapiNibMetadata {
	this := NiaapiNibMetadata{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiNibMetadataWithDefaults instantiates a new NiaapiNibMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiNibMetadataWithDefaults() *NiaapiNibMetadata {
	this := NiaapiNibMetadata{}
	var classId string = "niaapi.NibMetadata"
	this.ClassId = classId
	var objectType string = "niaapi.NibMetadata"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiNibMetadata) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiNibMetadata) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiNibMetadata) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiNibMetadata) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiNibMetadata) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiNibMetadata) SetObjectType(v string) {
	o.ObjectType = v
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiNibMetadata) GetContent() []NiaapiDetail {
	if o == nil  {
		var ret []NiaapiDetail
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiNibMetadata) GetContentOk() (*[]NiaapiDetail, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return &o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *NiaapiNibMetadata) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []NiaapiDetail and assigns it to the Content field.
func (o *NiaapiNibMetadata) SetContent(v []NiaapiDetail) {
	o.Content = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *NiaapiNibMetadata) GetDate() time.Time {
	if o == nil || o.Date == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNibMetadata) GetDateOk() (*time.Time, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *NiaapiNibMetadata) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *NiaapiNibMetadata) SetDate(v time.Time) {
	o.Date = &v
}

// GetMetadataChksum returns the MetadataChksum field value if set, zero value otherwise.
func (o *NiaapiNibMetadata) GetMetadataChksum() string {
	if o == nil || o.MetadataChksum == nil {
		var ret string
		return ret
	}
	return *o.MetadataChksum
}

// GetMetadataChksumOk returns a tuple with the MetadataChksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNibMetadata) GetMetadataChksumOk() (*string, bool) {
	if o == nil || o.MetadataChksum == nil {
		return nil, false
	}
	return o.MetadataChksum, true
}

// HasMetadataChksum returns a boolean if a field has been set.
func (o *NiaapiNibMetadata) HasMetadataChksum() bool {
	if o != nil && o.MetadataChksum != nil {
		return true
	}

	return false
}

// SetMetadataChksum gets a reference to the given string and assigns it to the MetadataChksum field.
func (o *NiaapiNibMetadata) SetMetadataChksum(v string) {
	o.MetadataChksum = &v
}

// GetMetadataFilename returns the MetadataFilename field value if set, zero value otherwise.
func (o *NiaapiNibMetadata) GetMetadataFilename() string {
	if o == nil || o.MetadataFilename == nil {
		var ret string
		return ret
	}
	return *o.MetadataFilename
}

// GetMetadataFilenameOk returns a tuple with the MetadataFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNibMetadata) GetMetadataFilenameOk() (*string, bool) {
	if o == nil || o.MetadataFilename == nil {
		return nil, false
	}
	return o.MetadataFilename, true
}

// HasMetadataFilename returns a boolean if a field has been set.
func (o *NiaapiNibMetadata) HasMetadataFilename() bool {
	if o != nil && o.MetadataFilename != nil {
		return true
	}

	return false
}

// SetMetadataFilename gets a reference to the given string and assigns it to the MetadataFilename field.
func (o *NiaapiNibMetadata) SetMetadataFilename(v string) {
	o.MetadataFilename = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NiaapiNibMetadata) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNibMetadata) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NiaapiNibMetadata) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *NiaapiNibMetadata) SetVersion(v int64) {
	o.Version = &v
}

func (o NiaapiNibMetadata) MarshalJSON() ([]byte, error) {
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
	if o.Content != nil {
		toSerialize["Content"] = o.Content
	}
	if o.Date != nil {
		toSerialize["Date"] = o.Date
	}
	if o.MetadataChksum != nil {
		toSerialize["MetadataChksum"] = o.MetadataChksum
	}
	if o.MetadataFilename != nil {
		toSerialize["MetadataFilename"] = o.MetadataFilename
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiNibMetadata) UnmarshalJSON(bytes []byte) (err error) {
	type NiaapiNibMetadataWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		Content []NiaapiDetail `json:"Content,omitempty"`
		// The date when the package was generated.
		Date *time.Time `json:"Date,omitempty"`
		// Chksum used to check the integrity of the metadata file downloaded.
		MetadataChksum *string `json:"MetadataChksum,omitempty"`
		// The filename of the metadata package.
		MetadataFilename *string `json:"MetadataFilename,omitempty"`
		// The version number of the metadata package.
		Version *int64 `json:"Version,omitempty"`
	}

	varNiaapiNibMetadataWithoutEmbeddedStruct := NiaapiNibMetadataWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiaapiNibMetadataWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiNibMetadata := _NiaapiNibMetadata{}
		varNiaapiNibMetadata.ClassId = varNiaapiNibMetadataWithoutEmbeddedStruct.ClassId
		varNiaapiNibMetadata.ObjectType = varNiaapiNibMetadataWithoutEmbeddedStruct.ObjectType
		varNiaapiNibMetadata.Content = varNiaapiNibMetadataWithoutEmbeddedStruct.Content
		varNiaapiNibMetadata.Date = varNiaapiNibMetadataWithoutEmbeddedStruct.Date
		varNiaapiNibMetadata.MetadataChksum = varNiaapiNibMetadataWithoutEmbeddedStruct.MetadataChksum
		varNiaapiNibMetadata.MetadataFilename = varNiaapiNibMetadataWithoutEmbeddedStruct.MetadataFilename
		varNiaapiNibMetadata.Version = varNiaapiNibMetadataWithoutEmbeddedStruct.Version
		*o = NiaapiNibMetadata(varNiaapiNibMetadata)
	} else {
		return err
	}

	varNiaapiNibMetadata := _NiaapiNibMetadata{}

	err = json.Unmarshal(bytes, &varNiaapiNibMetadata)
	if err == nil {
		o.MoBaseMo = varNiaapiNibMetadata.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Content")
		delete(additionalProperties, "Date")
		delete(additionalProperties, "MetadataChksum")
		delete(additionalProperties, "MetadataFilename")
		delete(additionalProperties, "Version")

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

type NullableNiaapiNibMetadata struct {
	value *NiaapiNibMetadata
	isSet bool
}

func (v NullableNiaapiNibMetadata) Get() *NiaapiNibMetadata {
	return v.value
}

func (v *NullableNiaapiNibMetadata) Set(val *NiaapiNibMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiNibMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiNibMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiNibMetadata(val *NiaapiNibMetadata) *NullableNiaapiNibMetadata {
	return &NullableNiaapiNibMetadata{value: val, isSet: true}
}

func (v NullableNiaapiNibMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiNibMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


