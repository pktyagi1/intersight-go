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

// SoftwarerepositoryDownloadSpecAllOf Definition of the list of properties defined in 'softwarerepository.DownloadSpec', excluding properties defined in parent classes.
type SoftwarerepositoryDownloadSpecAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The OAuth2 token that will be used during image download by the endpoint to authenticate with file server.
	AuthToken *string `json:"AuthToken,omitempty"`
	// The certificate of file server that will be used by the endpoint to validate the server before starting the file download.
	Certificate *string `json:"Certificate,omitempty"`
	// The name of the firmware image.
	Filename *string `json:"Filename,omitempty"`
	// MD5 sum of the firmware image that will be used by the endpoint to validate the integrity of the image, post download.
	Md5sum *string `json:"Md5sum,omitempty"`
	// The size (in bytes) of the firmware image.
	Size *int64 `json:"Size,omitempty"`
	// The URL of this file in file server. The endpoint uses this URL to download the file from the file server.
	Url *string `json:"Url,omitempty"`
	File *SoftwarerepositoryFileRelationship `json:"File,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryDownloadSpecAllOf SoftwarerepositoryDownloadSpecAllOf

// NewSoftwarerepositoryDownloadSpecAllOf instantiates a new SoftwarerepositoryDownloadSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryDownloadSpecAllOf(classId string, objectType string) *SoftwarerepositoryDownloadSpecAllOf {
	this := SoftwarerepositoryDownloadSpecAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwarerepositoryDownloadSpecAllOfWithDefaults instantiates a new SoftwarerepositoryDownloadSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryDownloadSpecAllOfWithDefaults() *SoftwarerepositoryDownloadSpecAllOf {
	this := SoftwarerepositoryDownloadSpecAllOf{}
	var classId string = "softwarerepository.DownloadSpec"
	this.ClassId = classId
	var objectType string = "softwarerepository.DownloadSpec"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryDownloadSpecAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryDownloadSpecAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryDownloadSpecAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryDownloadSpecAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAuthToken returns the AuthToken field value if set, zero value otherwise.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetAuthToken() string {
	if o == nil || o.AuthToken == nil {
		var ret string
		return ret
	}
	return *o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetAuthTokenOk() (*string, bool) {
	if o == nil || o.AuthToken == nil {
		return nil, false
	}
	return o.AuthToken, true
}

// HasAuthToken returns a boolean if a field has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) HasAuthToken() bool {
	if o != nil && o.AuthToken != nil {
		return true
	}

	return false
}

// SetAuthToken gets a reference to the given string and assigns it to the AuthToken field.
func (o *SoftwarerepositoryDownloadSpecAllOf) SetAuthToken(v string) {
	o.AuthToken = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *SoftwarerepositoryDownloadSpecAllOf) SetCertificate(v string) {
	o.Certificate = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *SoftwarerepositoryDownloadSpecAllOf) SetFilename(v string) {
	o.Filename = &v
}

// GetMd5sum returns the Md5sum field value if set, zero value otherwise.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetMd5sum() string {
	if o == nil || o.Md5sum == nil {
		var ret string
		return ret
	}
	return *o.Md5sum
}

// GetMd5sumOk returns a tuple with the Md5sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetMd5sumOk() (*string, bool) {
	if o == nil || o.Md5sum == nil {
		return nil, false
	}
	return o.Md5sum, true
}

// HasMd5sum returns a boolean if a field has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) HasMd5sum() bool {
	if o != nil && o.Md5sum != nil {
		return true
	}

	return false
}

// SetMd5sum gets a reference to the given string and assigns it to the Md5sum field.
func (o *SoftwarerepositoryDownloadSpecAllOf) SetMd5sum(v string) {
	o.Md5sum = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *SoftwarerepositoryDownloadSpecAllOf) SetSize(v int64) {
	o.Size = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SoftwarerepositoryDownloadSpecAllOf) SetUrl(v string) {
	o.Url = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetFile() SoftwarerepositoryFileRelationship {
	if o == nil || o.File == nil {
		var ret SoftwarerepositoryFileRelationship
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) GetFileOk() (*SoftwarerepositoryFileRelationship, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *SoftwarerepositoryDownloadSpecAllOf) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given SoftwarerepositoryFileRelationship and assigns it to the File field.
func (o *SoftwarerepositoryDownloadSpecAllOf) SetFile(v SoftwarerepositoryFileRelationship) {
	o.File = &v
}

func (o SoftwarerepositoryDownloadSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AuthToken != nil {
		toSerialize["AuthToken"] = o.AuthToken
	}
	if o.Certificate != nil {
		toSerialize["Certificate"] = o.Certificate
	}
	if o.Filename != nil {
		toSerialize["Filename"] = o.Filename
	}
	if o.Md5sum != nil {
		toSerialize["Md5sum"] = o.Md5sum
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.Url != nil {
		toSerialize["Url"] = o.Url
	}
	if o.File != nil {
		toSerialize["File"] = o.File
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryDownloadSpecAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwarerepositoryDownloadSpecAllOf := _SoftwarerepositoryDownloadSpecAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwarerepositoryDownloadSpecAllOf); err == nil {
		*o = SoftwarerepositoryDownloadSpecAllOf(varSoftwarerepositoryDownloadSpecAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthToken")
		delete(additionalProperties, "Certificate")
		delete(additionalProperties, "Filename")
		delete(additionalProperties, "Md5sum")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "Url")
		delete(additionalProperties, "File")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwarerepositoryDownloadSpecAllOf struct {
	value *SoftwarerepositoryDownloadSpecAllOf
	isSet bool
}

func (v NullableSoftwarerepositoryDownloadSpecAllOf) Get() *SoftwarerepositoryDownloadSpecAllOf {
	return v.value
}

func (v *NullableSoftwarerepositoryDownloadSpecAllOf) Set(val *SoftwarerepositoryDownloadSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryDownloadSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryDownloadSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryDownloadSpecAllOf(val *SoftwarerepositoryDownloadSpecAllOf) *NullableSoftwarerepositoryDownloadSpecAllOf {
	return &NullableSoftwarerepositoryDownloadSpecAllOf{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryDownloadSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryDownloadSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


