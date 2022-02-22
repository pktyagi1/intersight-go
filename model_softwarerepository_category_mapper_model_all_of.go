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

// SoftwarerepositoryCategoryMapperModelAllOf Definition of the list of properties defined in 'softwarerepository.CategoryMapperModel', excluding properties defined in parent classes.
type SoftwarerepositoryCategoryMapperModelAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The category of the model series.
	Category *string `json:"Category,omitempty"`
	// The distributable tag value of the model series.
	DistTag *string `json:"DistTag,omitempty"`
	// The type of image based on the endpoint it can upgrade. For example, ucs-bundle-6400-infra.4.1.2a.bin can upgrade ucs managed fabric interconnects, so the image type is UCS Managed Fabric Interconnect.
	ImageType *string `json:"ImageType,omitempty"`
	// The regex that all images of this model follow.
	RegexPattern *string `json:"RegexPattern,omitempty"`
	// Cisco hardware model series.
	SeriesId *string `json:"SeriesId,omitempty"`
	SupportedModels []string `json:"SupportedModels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryCategoryMapperModelAllOf SoftwarerepositoryCategoryMapperModelAllOf

// NewSoftwarerepositoryCategoryMapperModelAllOf instantiates a new SoftwarerepositoryCategoryMapperModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryCategoryMapperModelAllOf(classId string, objectType string) *SoftwarerepositoryCategoryMapperModelAllOf {
	this := SoftwarerepositoryCategoryMapperModelAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwarerepositoryCategoryMapperModelAllOfWithDefaults instantiates a new SoftwarerepositoryCategoryMapperModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryCategoryMapperModelAllOfWithDefaults() *SoftwarerepositoryCategoryMapperModelAllOf {
	this := SoftwarerepositoryCategoryMapperModelAllOf{}
	var classId string = "softwarerepository.CategoryMapperModel"
	this.ClassId = classId
	var objectType string = "softwarerepository.CategoryMapperModel"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetCategory(v string) {
	o.Category = &v
}

// GetDistTag returns the DistTag field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetDistTag() string {
	if o == nil || o.DistTag == nil {
		var ret string
		return ret
	}
	return *o.DistTag
}

// GetDistTagOk returns a tuple with the DistTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetDistTagOk() (*string, bool) {
	if o == nil || o.DistTag == nil {
		return nil, false
	}
	return o.DistTag, true
}

// HasDistTag returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) HasDistTag() bool {
	if o != nil && o.DistTag != nil {
		return true
	}

	return false
}

// SetDistTag gets a reference to the given string and assigns it to the DistTag field.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetDistTag(v string) {
	o.DistTag = &v
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetImageType() string {
	if o == nil || o.ImageType == nil {
		var ret string
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetImageTypeOk() (*string, bool) {
	if o == nil || o.ImageType == nil {
		return nil, false
	}
	return o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) HasImageType() bool {
	if o != nil && o.ImageType != nil {
		return true
	}

	return false
}

// SetImageType gets a reference to the given string and assigns it to the ImageType field.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetImageType(v string) {
	o.ImageType = &v
}

// GetRegexPattern returns the RegexPattern field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetRegexPattern() string {
	if o == nil || o.RegexPattern == nil {
		var ret string
		return ret
	}
	return *o.RegexPattern
}

// GetRegexPatternOk returns a tuple with the RegexPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetRegexPatternOk() (*string, bool) {
	if o == nil || o.RegexPattern == nil {
		return nil, false
	}
	return o.RegexPattern, true
}

// HasRegexPattern returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) HasRegexPattern() bool {
	if o != nil && o.RegexPattern != nil {
		return true
	}

	return false
}

// SetRegexPattern gets a reference to the given string and assigns it to the RegexPattern field.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetRegexPattern(v string) {
	o.RegexPattern = &v
}

// GetSeriesId returns the SeriesId field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetSeriesId() string {
	if o == nil || o.SeriesId == nil {
		var ret string
		return ret
	}
	return *o.SeriesId
}

// GetSeriesIdOk returns a tuple with the SeriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetSeriesIdOk() (*string, bool) {
	if o == nil || o.SeriesId == nil {
		return nil, false
	}
	return o.SeriesId, true
}

// HasSeriesId returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) HasSeriesId() bool {
	if o != nil && o.SeriesId != nil {
		return true
	}

	return false
}

// SetSeriesId gets a reference to the given string and assigns it to the SeriesId field.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetSeriesId(v string) {
	o.SeriesId = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetSupportedModels() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetSupportedModelsOk() (*[]string, bool) {
	if o == nil || o.SupportedModels == nil {
		return nil, false
	}
	return &o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) HasSupportedModels() bool {
	if o != nil && o.SupportedModels != nil {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

func (o SoftwarerepositoryCategoryMapperModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.DistTag != nil {
		toSerialize["DistTag"] = o.DistTag
	}
	if o.ImageType != nil {
		toSerialize["ImageType"] = o.ImageType
	}
	if o.RegexPattern != nil {
		toSerialize["RegexPattern"] = o.RegexPattern
	}
	if o.SeriesId != nil {
		toSerialize["SeriesId"] = o.SeriesId
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryCategoryMapperModelAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwarerepositoryCategoryMapperModelAllOf := _SoftwarerepositoryCategoryMapperModelAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwarerepositoryCategoryMapperModelAllOf); err == nil {
		*o = SoftwarerepositoryCategoryMapperModelAllOf(varSoftwarerepositoryCategoryMapperModelAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "DistTag")
		delete(additionalProperties, "ImageType")
		delete(additionalProperties, "RegexPattern")
		delete(additionalProperties, "SeriesId")
		delete(additionalProperties, "SupportedModels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwarerepositoryCategoryMapperModelAllOf struct {
	value *SoftwarerepositoryCategoryMapperModelAllOf
	isSet bool
}

func (v NullableSoftwarerepositoryCategoryMapperModelAllOf) Get() *SoftwarerepositoryCategoryMapperModelAllOf {
	return v.value
}

func (v *NullableSoftwarerepositoryCategoryMapperModelAllOf) Set(val *SoftwarerepositoryCategoryMapperModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryCategoryMapperModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryCategoryMapperModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryCategoryMapperModelAllOf(val *SoftwarerepositoryCategoryMapperModelAllOf) *NullableSoftwarerepositoryCategoryMapperModelAllOf {
	return &NullableSoftwarerepositoryCategoryMapperModelAllOf{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryCategoryMapperModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryCategoryMapperModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


