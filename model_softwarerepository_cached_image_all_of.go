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
	"time"
)

// SoftwarerepositoryCachedImageAllOf Definition of the list of properties defined in 'softwarerepository.CachedImage', excluding properties defined in parent classes.
type SoftwarerepositoryCachedImageAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in endpoint. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source. * `None` - No action should be taken on the imported file. * `GeneratePreSignedUploadUrl` - Generate pre signed URL of file for importing into the repository. * `GeneratePreSignedDownloadUrl` - Generate pre signed URL of file in the repository to download. * `CompleteImportProcess` - Mark that the import process of the file into the repository is complete. * `MarkImportFailed` - Mark to indicate that the import process of the file into the repository failed. * `PreCache` - Cache the file into the Intersight Appliance. * `Cancel` - The cancel import process for the file into the repository. * `Extract` - The action to extract the file in the external repository. * `Evict` - Evict the cached file from the Intersight Appliance.
	Action *string `json:"Action,omitempty"`
	// The state  of this file in the endpoint The importState is updated during the cache operation and as part of the cache monitoring process. * `ReadyForImport` - The image is ready to be imported into the repository. * `Importing` - The image is being imported into the repository. * `Imported` - The image has been extracted and imported into the repository. * `PendingExtraction` - Indicates that the image has been imported but not extracted in the repository. * `Extracting` - Indicates that the image is being extracted into the repository. * `Extracted` - Indicates that the image has been extracted into the repository. * `Failed` - The image import from an external source to the repository has failed. * `MetaOnly` - The image is present in an external repository. * `ReadyForCache` - The image is ready to be cached into the Intersight Appliance. * `Caching` - Indicates that the image is being cached into the Intersight Appliance or endpoint cache. * `Cached` - Indicates that the image has been cached into the Intersight Appliance or endpoint cache. * `CachingFailed` - Indicates that the image caching into the Intersight Appliance failed or endpoint cache. * `Corrupted` - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached. * `Evicted` - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space. * `Invalid` - Indicates that the corresponding distributable MO has been removed from the backend. This can be due to unpublishing of an image.
	CacheState *string `json:"CacheState,omitempty"`
	// The time when the image or file was cached into the FI storage.
	CachedTime *time.Time `json:"CachedTime,omitempty"`
	// Used by the cache monitoring process to determine the files that are to be evicted from the cache.
	LastAccessTime *time.Time `json:"LastAccessTime,omitempty"`
	// The MD5 sum of the firmware image that will be used by the endpoint to validate the integrity of the image.
	Md5sum *string `json:"Md5sum,omitempty"`
	// The actual sha512sum of the cached image.
	OriginalSha512sum *string `json:"OriginalSha512sum,omitempty"`
	// The absolute path of the imported file in the endpoint.
	Path *string `json:"Path,omitempty"`
	RegisteredWorkflows []string `json:"RegisteredWorkflows,omitempty"`
	// The number of times this file has been used to copy or upgrade or install actions. Used by the cache monitoring process to determine the files to be evicted from the cache.
	UsedCount *int64 `json:"UsedCount,omitempty"`
	File *SoftwarerepositoryFileRelationship `json:"File,omitempty"`
	NetworkElement *NetworkElementRelationship `json:"NetworkElement,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryCachedImageAllOf SoftwarerepositoryCachedImageAllOf

// NewSoftwarerepositoryCachedImageAllOf instantiates a new SoftwarerepositoryCachedImageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryCachedImageAllOf(classId string, objectType string) *SoftwarerepositoryCachedImageAllOf {
	this := SoftwarerepositoryCachedImageAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// NewSoftwarerepositoryCachedImageAllOfWithDefaults instantiates a new SoftwarerepositoryCachedImageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryCachedImageAllOfWithDefaults() *SoftwarerepositoryCachedImageAllOf {
	this := SoftwarerepositoryCachedImageAllOf{}
	var classId string = "softwarerepository.CachedImage"
	this.ClassId = classId
	var objectType string = "softwarerepository.CachedImage"
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryCachedImageAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCachedImageAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryCachedImageAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryCachedImageAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCachedImageAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryCachedImageAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *SoftwarerepositoryCachedImageAllOf) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCachedImageAllOf) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *SoftwarerepositoryCachedImageAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *SoftwarerepositoryCachedImageAllOf) SetAction(v string) {
	o.Action = &v
}

// GetCacheState returns the CacheState field value if set, zero value otherwise.
func (o *SoftwarerepositoryCachedImageAllOf) GetCacheState() string {
	if o == nil || o.CacheState == nil {
		var ret string
		return ret
	}
	return *o.CacheState
}

// GetCacheStateOk returns a tuple with the CacheState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCachedImageAllOf) GetCacheStateOk() (*string, bool) {
	if o == nil || o.CacheState == nil {
		return nil, false
	}
	return o.CacheState, true
}

// HasCacheState returns a boolean if a field has been set.
func (o *SoftwarerepositoryCachedImageAllOf) HasCacheState() bool {
	if o != nil && o.CacheState != nil {
		return true
	}

	return false
}

// SetCacheState gets a reference to the given string and assigns it to the CacheState field.
func (o *SoftwarerepositoryCachedImageAllOf) SetCacheState(v string) {
	o.CacheState = &v
}

// GetCachedTime returns the CachedTime field value if set, zero value otherwise.
func (o *SoftwarerepositoryCachedImageAllOf) GetCachedTime() time.Time {
	if o == nil || o.CachedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CachedTime
}

// GetCachedTimeOk returns a tuple with the CachedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCachedImageAllOf) GetCachedTimeOk() (*time.Time, bool) {
	if o == nil || o.CachedTime == nil {
		return nil, false
	}
	return o.CachedTime, true
}

// HasCachedTime returns a boolean if a field has been set.
func (o *SoftwarerepositoryCachedImageAllOf) HasCachedTime() bool {
	if o != nil && o.CachedTime != nil {
		return true
	}

	return false
}

// SetCachedTime gets a reference to the given time.Time and assigns it to the CachedTime field.
func (o *SoftwarerepositoryCachedImageAllOf) SetCachedTime(v time.Time) {
	o.CachedTime = &v
}

// GetLastAccessTime returns the LastAccessTime field value if set, zero value otherwise.
func (o *SoftwarerepositoryCachedImageAllOf) GetLastAccessTime() time.Time {
	if o == nil || o.LastAccessTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessTime
}

// GetLastAccessTimeOk returns a tuple with the LastAccessTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCachedImageAllOf) GetLastAccessTimeOk() (*time.Time, bool) {
	if o == nil || o.LastAccessTime == nil {
		return nil, false
	}
	return o.LastAccessTime, true
}

// HasLastAccessTime returns a boolean if a field has been set.
func (o *SoftwarerepositoryCachedImageAllOf) HasLastAccessTime() bool {
	if o != nil && o.LastAccessTime != nil {
		return true
	}

	return false
}

// SetLastAccessTime gets a reference to the given time.Time and assigns it to the LastAccessTime field.
func (o *SoftwarerepositoryCachedImageAllOf) SetLastAccessTime(v time.Time) {
	o.LastAccessTime = &v
}

// GetMd5sum returns the Md5sum field value if set, zero value otherwise.
func (o *SoftwarerepositoryCachedImageAllOf) GetMd5sum() string {
	if o == nil || o.Md5sum == nil {
		var ret string
		return ret
	}
	return *o.Md5sum
}

// GetMd5sumOk returns a tuple with the Md5sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCachedImageAllOf) GetMd5sumOk() (*string, bool) {
	if o == nil || o.Md5sum == nil {
		return nil, false
	}
	return o.Md5sum, true
}

// HasMd5sum returns a boolean if a field has been set.
func (o *SoftwarerepositoryCachedImageAllOf) HasMd5sum() bool {
	if o != nil && o.Md5sum != nil {
		return true
	}

	return false
}

// SetMd5sum gets a reference to the given string and assigns it to the Md5sum field.
func (o *SoftwarerepositoryCachedImageAllOf) SetMd5sum(v string) {
	o.Md5sum = &v
}

// GetOriginalSha512sum returns the OriginalSha512sum field value if set, zero value otherwise.
func (o *SoftwarerepositoryCachedImageAllOf) GetOriginalSha512sum() string {
	if o == nil || o.OriginalSha512sum == nil {
		var ret string
		return ret
	}
	return *o.OriginalSha512sum
}

// GetOriginalSha512sumOk returns a tuple with the OriginalSha512sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCachedImageAllOf) GetOriginalSha512sumOk() (*string, bool) {
	if o == nil || o.OriginalSha512sum == nil {
		return nil, false
	}
	return o.OriginalSha512sum, true
}

// HasOriginalSha512sum returns a boolean if a field has been set.
func (o *SoftwarerepositoryCachedImageAllOf) HasOriginalSha512sum() bool {
	if o != nil && o.OriginalSha512sum != nil {
		return true
	}

	return false
}

// SetOriginalSha512sum gets a reference to the given string and assigns it to the OriginalSha512sum field.
func (o *SoftwarerepositoryCachedImageAllOf) SetOriginalSha512sum(v string) {
	o.OriginalSha512sum = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SoftwarerepositoryCachedImageAllOf) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCachedImageAllOf) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SoftwarerepositoryCachedImageAllOf) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SoftwarerepositoryCachedImageAllOf) SetPath(v string) {
	o.Path = &v
}

// GetRegisteredWorkflows returns the RegisteredWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryCachedImageAllOf) GetRegisteredWorkflows() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.RegisteredWorkflows
}

// GetRegisteredWorkflowsOk returns a tuple with the RegisteredWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryCachedImageAllOf) GetRegisteredWorkflowsOk() (*[]string, bool) {
	if o == nil || o.RegisteredWorkflows == nil {
		return nil, false
	}
	return &o.RegisteredWorkflows, true
}

// HasRegisteredWorkflows returns a boolean if a field has been set.
func (o *SoftwarerepositoryCachedImageAllOf) HasRegisteredWorkflows() bool {
	if o != nil && o.RegisteredWorkflows != nil {
		return true
	}

	return false
}

// SetRegisteredWorkflows gets a reference to the given []string and assigns it to the RegisteredWorkflows field.
func (o *SoftwarerepositoryCachedImageAllOf) SetRegisteredWorkflows(v []string) {
	o.RegisteredWorkflows = v
}

// GetUsedCount returns the UsedCount field value if set, zero value otherwise.
func (o *SoftwarerepositoryCachedImageAllOf) GetUsedCount() int64 {
	if o == nil || o.UsedCount == nil {
		var ret int64
		return ret
	}
	return *o.UsedCount
}

// GetUsedCountOk returns a tuple with the UsedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCachedImageAllOf) GetUsedCountOk() (*int64, bool) {
	if o == nil || o.UsedCount == nil {
		return nil, false
	}
	return o.UsedCount, true
}

// HasUsedCount returns a boolean if a field has been set.
func (o *SoftwarerepositoryCachedImageAllOf) HasUsedCount() bool {
	if o != nil && o.UsedCount != nil {
		return true
	}

	return false
}

// SetUsedCount gets a reference to the given int64 and assigns it to the UsedCount field.
func (o *SoftwarerepositoryCachedImageAllOf) SetUsedCount(v int64) {
	o.UsedCount = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *SoftwarerepositoryCachedImageAllOf) GetFile() SoftwarerepositoryFileRelationship {
	if o == nil || o.File == nil {
		var ret SoftwarerepositoryFileRelationship
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCachedImageAllOf) GetFileOk() (*SoftwarerepositoryFileRelationship, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *SoftwarerepositoryCachedImageAllOf) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given SoftwarerepositoryFileRelationship and assigns it to the File field.
func (o *SoftwarerepositoryCachedImageAllOf) SetFile(v SoftwarerepositoryFileRelationship) {
	o.File = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *SoftwarerepositoryCachedImageAllOf) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCachedImageAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *SoftwarerepositoryCachedImageAllOf) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *SoftwarerepositoryCachedImageAllOf) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

func (o SoftwarerepositoryCachedImageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.CacheState != nil {
		toSerialize["CacheState"] = o.CacheState
	}
	if o.CachedTime != nil {
		toSerialize["CachedTime"] = o.CachedTime
	}
	if o.LastAccessTime != nil {
		toSerialize["LastAccessTime"] = o.LastAccessTime
	}
	if o.Md5sum != nil {
		toSerialize["Md5sum"] = o.Md5sum
	}
	if o.OriginalSha512sum != nil {
		toSerialize["OriginalSha512sum"] = o.OriginalSha512sum
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.RegisteredWorkflows != nil {
		toSerialize["RegisteredWorkflows"] = o.RegisteredWorkflows
	}
	if o.UsedCount != nil {
		toSerialize["UsedCount"] = o.UsedCount
	}
	if o.File != nil {
		toSerialize["File"] = o.File
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryCachedImageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwarerepositoryCachedImageAllOf := _SoftwarerepositoryCachedImageAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwarerepositoryCachedImageAllOf); err == nil {
		*o = SoftwarerepositoryCachedImageAllOf(varSoftwarerepositoryCachedImageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "CacheState")
		delete(additionalProperties, "CachedTime")
		delete(additionalProperties, "LastAccessTime")
		delete(additionalProperties, "Md5sum")
		delete(additionalProperties, "OriginalSha512sum")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "RegisteredWorkflows")
		delete(additionalProperties, "UsedCount")
		delete(additionalProperties, "File")
		delete(additionalProperties, "NetworkElement")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwarerepositoryCachedImageAllOf struct {
	value *SoftwarerepositoryCachedImageAllOf
	isSet bool
}

func (v NullableSoftwarerepositoryCachedImageAllOf) Get() *SoftwarerepositoryCachedImageAllOf {
	return v.value
}

func (v *NullableSoftwarerepositoryCachedImageAllOf) Set(val *SoftwarerepositoryCachedImageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryCachedImageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryCachedImageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryCachedImageAllOf(val *SoftwarerepositoryCachedImageAllOf) *NullableSoftwarerepositoryCachedImageAllOf {
	return &NullableSoftwarerepositoryCachedImageAllOf{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryCachedImageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryCachedImageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


