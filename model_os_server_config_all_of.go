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

// OsServerConfigAllOf Definition of the list of properties defined in 'os.ServerConfig', excluding properties defined in parent classes.
type OsServerConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	AdditionalParameters []OsPlaceHolder `json:"AdditionalParameters,omitempty"`
	Answers NullableOsAnswers `json:"Answers,omitempty"`
	ErrorMsgs []string `json:"ErrorMsgs,omitempty"`
	// The target in which OS installation triggered, the value represented is StorageControllerID follwed by PhysicalDisk SerialNumber in case of Physcial disk or VirtualDriveId for virtual drive.
	InstallTarget *string `json:"InstallTarget,omitempty"`
	OperatingSystemParameters NullableOsOperatingSystemParameters `json:"OperatingSystemParameters,omitempty"`
	ProcessedInstallTarget NullableOsInstallTarget `json:"ProcessedInstallTarget,omitempty"`
	// The Serial Number of the server.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsServerConfigAllOf OsServerConfigAllOf

// NewOsServerConfigAllOf instantiates a new OsServerConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsServerConfigAllOf(classId string, objectType string) *OsServerConfigAllOf {
	this := OsServerConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsServerConfigAllOfWithDefaults instantiates a new OsServerConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsServerConfigAllOfWithDefaults() *OsServerConfigAllOf {
	this := OsServerConfigAllOf{}
	var classId string = "os.ServerConfig"
	this.ClassId = classId
	var objectType string = "os.ServerConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsServerConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsServerConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsServerConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsServerConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsServerConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsServerConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdditionalParameters returns the AdditionalParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsServerConfigAllOf) GetAdditionalParameters() []OsPlaceHolder {
	if o == nil  {
		var ret []OsPlaceHolder
		return ret
	}
	return o.AdditionalParameters
}

// GetAdditionalParametersOk returns a tuple with the AdditionalParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsServerConfigAllOf) GetAdditionalParametersOk() (*[]OsPlaceHolder, bool) {
	if o == nil || o.AdditionalParameters == nil {
		return nil, false
	}
	return &o.AdditionalParameters, true
}

// HasAdditionalParameters returns a boolean if a field has been set.
func (o *OsServerConfigAllOf) HasAdditionalParameters() bool {
	if o != nil && o.AdditionalParameters != nil {
		return true
	}

	return false
}

// SetAdditionalParameters gets a reference to the given []OsPlaceHolder and assigns it to the AdditionalParameters field.
func (o *OsServerConfigAllOf) SetAdditionalParameters(v []OsPlaceHolder) {
	o.AdditionalParameters = v
}

// GetAnswers returns the Answers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsServerConfigAllOf) GetAnswers() OsAnswers {
	if o == nil || o.Answers.Get() == nil {
		var ret OsAnswers
		return ret
	}
	return *o.Answers.Get()
}

// GetAnswersOk returns a tuple with the Answers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsServerConfigAllOf) GetAnswersOk() (*OsAnswers, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Answers.Get(), o.Answers.IsSet()
}

// HasAnswers returns a boolean if a field has been set.
func (o *OsServerConfigAllOf) HasAnswers() bool {
	if o != nil && o.Answers.IsSet() {
		return true
	}

	return false
}

// SetAnswers gets a reference to the given NullableOsAnswers and assigns it to the Answers field.
func (o *OsServerConfigAllOf) SetAnswers(v OsAnswers) {
	o.Answers.Set(&v)
}
// SetAnswersNil sets the value for Answers to be an explicit nil
func (o *OsServerConfigAllOf) SetAnswersNil() {
	o.Answers.Set(nil)
}

// UnsetAnswers ensures that no value is present for Answers, not even an explicit nil
func (o *OsServerConfigAllOf) UnsetAnswers() {
	o.Answers.Unset()
}

// GetErrorMsgs returns the ErrorMsgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsServerConfigAllOf) GetErrorMsgs() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.ErrorMsgs
}

// GetErrorMsgsOk returns a tuple with the ErrorMsgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsServerConfigAllOf) GetErrorMsgsOk() (*[]string, bool) {
	if o == nil || o.ErrorMsgs == nil {
		return nil, false
	}
	return &o.ErrorMsgs, true
}

// HasErrorMsgs returns a boolean if a field has been set.
func (o *OsServerConfigAllOf) HasErrorMsgs() bool {
	if o != nil && o.ErrorMsgs != nil {
		return true
	}

	return false
}

// SetErrorMsgs gets a reference to the given []string and assigns it to the ErrorMsgs field.
func (o *OsServerConfigAllOf) SetErrorMsgs(v []string) {
	o.ErrorMsgs = v
}

// GetInstallTarget returns the InstallTarget field value if set, zero value otherwise.
func (o *OsServerConfigAllOf) GetInstallTarget() string {
	if o == nil || o.InstallTarget == nil {
		var ret string
		return ret
	}
	return *o.InstallTarget
}

// GetInstallTargetOk returns a tuple with the InstallTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsServerConfigAllOf) GetInstallTargetOk() (*string, bool) {
	if o == nil || o.InstallTarget == nil {
		return nil, false
	}
	return o.InstallTarget, true
}

// HasInstallTarget returns a boolean if a field has been set.
func (o *OsServerConfigAllOf) HasInstallTarget() bool {
	if o != nil && o.InstallTarget != nil {
		return true
	}

	return false
}

// SetInstallTarget gets a reference to the given string and assigns it to the InstallTarget field.
func (o *OsServerConfigAllOf) SetInstallTarget(v string) {
	o.InstallTarget = &v
}

// GetOperatingSystemParameters returns the OperatingSystemParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsServerConfigAllOf) GetOperatingSystemParameters() OsOperatingSystemParameters {
	if o == nil || o.OperatingSystemParameters.Get() == nil {
		var ret OsOperatingSystemParameters
		return ret
	}
	return *o.OperatingSystemParameters.Get()
}

// GetOperatingSystemParametersOk returns a tuple with the OperatingSystemParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsServerConfigAllOf) GetOperatingSystemParametersOk() (*OsOperatingSystemParameters, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OperatingSystemParameters.Get(), o.OperatingSystemParameters.IsSet()
}

// HasOperatingSystemParameters returns a boolean if a field has been set.
func (o *OsServerConfigAllOf) HasOperatingSystemParameters() bool {
	if o != nil && o.OperatingSystemParameters.IsSet() {
		return true
	}

	return false
}

// SetOperatingSystemParameters gets a reference to the given NullableOsOperatingSystemParameters and assigns it to the OperatingSystemParameters field.
func (o *OsServerConfigAllOf) SetOperatingSystemParameters(v OsOperatingSystemParameters) {
	o.OperatingSystemParameters.Set(&v)
}
// SetOperatingSystemParametersNil sets the value for OperatingSystemParameters to be an explicit nil
func (o *OsServerConfigAllOf) SetOperatingSystemParametersNil() {
	o.OperatingSystemParameters.Set(nil)
}

// UnsetOperatingSystemParameters ensures that no value is present for OperatingSystemParameters, not even an explicit nil
func (o *OsServerConfigAllOf) UnsetOperatingSystemParameters() {
	o.OperatingSystemParameters.Unset()
}

// GetProcessedInstallTarget returns the ProcessedInstallTarget field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsServerConfigAllOf) GetProcessedInstallTarget() OsInstallTarget {
	if o == nil || o.ProcessedInstallTarget.Get() == nil {
		var ret OsInstallTarget
		return ret
	}
	return *o.ProcessedInstallTarget.Get()
}

// GetProcessedInstallTargetOk returns a tuple with the ProcessedInstallTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsServerConfigAllOf) GetProcessedInstallTargetOk() (*OsInstallTarget, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProcessedInstallTarget.Get(), o.ProcessedInstallTarget.IsSet()
}

// HasProcessedInstallTarget returns a boolean if a field has been set.
func (o *OsServerConfigAllOf) HasProcessedInstallTarget() bool {
	if o != nil && o.ProcessedInstallTarget.IsSet() {
		return true
	}

	return false
}

// SetProcessedInstallTarget gets a reference to the given NullableOsInstallTarget and assigns it to the ProcessedInstallTarget field.
func (o *OsServerConfigAllOf) SetProcessedInstallTarget(v OsInstallTarget) {
	o.ProcessedInstallTarget.Set(&v)
}
// SetProcessedInstallTargetNil sets the value for ProcessedInstallTarget to be an explicit nil
func (o *OsServerConfigAllOf) SetProcessedInstallTargetNil() {
	o.ProcessedInstallTarget.Set(nil)
}

// UnsetProcessedInstallTarget ensures that no value is present for ProcessedInstallTarget, not even an explicit nil
func (o *OsServerConfigAllOf) UnsetProcessedInstallTarget() {
	o.ProcessedInstallTarget.Unset()
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *OsServerConfigAllOf) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsServerConfigAllOf) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *OsServerConfigAllOf) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *OsServerConfigAllOf) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

func (o OsServerConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdditionalParameters != nil {
		toSerialize["AdditionalParameters"] = o.AdditionalParameters
	}
	if o.Answers.IsSet() {
		toSerialize["Answers"] = o.Answers.Get()
	}
	if o.ErrorMsgs != nil {
		toSerialize["ErrorMsgs"] = o.ErrorMsgs
	}
	if o.InstallTarget != nil {
		toSerialize["InstallTarget"] = o.InstallTarget
	}
	if o.OperatingSystemParameters.IsSet() {
		toSerialize["OperatingSystemParameters"] = o.OperatingSystemParameters.Get()
	}
	if o.ProcessedInstallTarget.IsSet() {
		toSerialize["ProcessedInstallTarget"] = o.ProcessedInstallTarget.Get()
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsServerConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOsServerConfigAllOf := _OsServerConfigAllOf{}

	if err = json.Unmarshal(bytes, &varOsServerConfigAllOf); err == nil {
		*o = OsServerConfigAllOf(varOsServerConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdditionalParameters")
		delete(additionalProperties, "Answers")
		delete(additionalProperties, "ErrorMsgs")
		delete(additionalProperties, "InstallTarget")
		delete(additionalProperties, "OperatingSystemParameters")
		delete(additionalProperties, "ProcessedInstallTarget")
		delete(additionalProperties, "SerialNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOsServerConfigAllOf struct {
	value *OsServerConfigAllOf
	isSet bool
}

func (v NullableOsServerConfigAllOf) Get() *OsServerConfigAllOf {
	return v.value
}

func (v *NullableOsServerConfigAllOf) Set(val *OsServerConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsServerConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsServerConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsServerConfigAllOf(val *OsServerConfigAllOf) *NullableOsServerConfigAllOf {
	return &NullableOsServerConfigAllOf{value: val, isSet: true}
}

func (v NullableOsServerConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsServerConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


