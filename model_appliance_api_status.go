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

// ApplianceApiStatus Status of an API. APIs are all get type and from ApiDocs.
type ApplianceApiStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The elapsed time for query in seconds.
	ElapsedTime *float32 `json:"ElapsedTime,omitempty"`
	// Name to identify the API.
	ObjectName *string `json:"ObjectName,omitempty"`
	// Reason to address why the API call failed, if API call was successed, reason would be null.
	Reason *string `json:"Reason,omitempty"`
	// Response code of the API call.
	Response *int64 `json:"Response,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceApiStatus ApplianceApiStatus

// NewApplianceApiStatus instantiates a new ApplianceApiStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceApiStatus(classId string, objectType string) *ApplianceApiStatus {
	this := ApplianceApiStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceApiStatusWithDefaults instantiates a new ApplianceApiStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceApiStatusWithDefaults() *ApplianceApiStatus {
	this := ApplianceApiStatus{}
	var classId string = "appliance.ApiStatus"
	this.ClassId = classId
	var objectType string = "appliance.ApiStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceApiStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceApiStatus) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceApiStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceApiStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceApiStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceApiStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetElapsedTime returns the ElapsedTime field value if set, zero value otherwise.
func (o *ApplianceApiStatus) GetElapsedTime() float32 {
	if o == nil || o.ElapsedTime == nil {
		var ret float32
		return ret
	}
	return *o.ElapsedTime
}

// GetElapsedTimeOk returns a tuple with the ElapsedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceApiStatus) GetElapsedTimeOk() (*float32, bool) {
	if o == nil || o.ElapsedTime == nil {
		return nil, false
	}
	return o.ElapsedTime, true
}

// HasElapsedTime returns a boolean if a field has been set.
func (o *ApplianceApiStatus) HasElapsedTime() bool {
	if o != nil && o.ElapsedTime != nil {
		return true
	}

	return false
}

// SetElapsedTime gets a reference to the given float32 and assigns it to the ElapsedTime field.
func (o *ApplianceApiStatus) SetElapsedTime(v float32) {
	o.ElapsedTime = &v
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *ApplianceApiStatus) GetObjectName() string {
	if o == nil || o.ObjectName == nil {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceApiStatus) GetObjectNameOk() (*string, bool) {
	if o == nil || o.ObjectName == nil {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *ApplianceApiStatus) HasObjectName() bool {
	if o != nil && o.ObjectName != nil {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *ApplianceApiStatus) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ApplianceApiStatus) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceApiStatus) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ApplianceApiStatus) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ApplianceApiStatus) SetReason(v string) {
	o.Reason = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ApplianceApiStatus) GetResponse() int64 {
	if o == nil || o.Response == nil {
		var ret int64
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceApiStatus) GetResponseOk() (*int64, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ApplianceApiStatus) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given int64 and assigns it to the Response field.
func (o *ApplianceApiStatus) SetResponse(v int64) {
	o.Response = &v
}

func (o ApplianceApiStatus) MarshalJSON() ([]byte, error) {
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
	if o.ElapsedTime != nil {
		toSerialize["ElapsedTime"] = o.ElapsedTime
	}
	if o.ObjectName != nil {
		toSerialize["ObjectName"] = o.ObjectName
	}
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}
	if o.Response != nil {
		toSerialize["Response"] = o.Response
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceApiStatus) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceApiStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The elapsed time for query in seconds.
		ElapsedTime *float32 `json:"ElapsedTime,omitempty"`
		// Name to identify the API.
		ObjectName *string `json:"ObjectName,omitempty"`
		// Reason to address why the API call failed, if API call was successed, reason would be null.
		Reason *string `json:"Reason,omitempty"`
		// Response code of the API call.
		Response *int64 `json:"Response,omitempty"`
	}

	varApplianceApiStatusWithoutEmbeddedStruct := ApplianceApiStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceApiStatusWithoutEmbeddedStruct)
	if err == nil {
		varApplianceApiStatus := _ApplianceApiStatus{}
		varApplianceApiStatus.ClassId = varApplianceApiStatusWithoutEmbeddedStruct.ClassId
		varApplianceApiStatus.ObjectType = varApplianceApiStatusWithoutEmbeddedStruct.ObjectType
		varApplianceApiStatus.ElapsedTime = varApplianceApiStatusWithoutEmbeddedStruct.ElapsedTime
		varApplianceApiStatus.ObjectName = varApplianceApiStatusWithoutEmbeddedStruct.ObjectName
		varApplianceApiStatus.Reason = varApplianceApiStatusWithoutEmbeddedStruct.Reason
		varApplianceApiStatus.Response = varApplianceApiStatusWithoutEmbeddedStruct.Response
		*o = ApplianceApiStatus(varApplianceApiStatus)
	} else {
		return err
	}

	varApplianceApiStatus := _ApplianceApiStatus{}

	err = json.Unmarshal(bytes, &varApplianceApiStatus)
	if err == nil {
		o.MoBaseComplexType = varApplianceApiStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ElapsedTime")
		delete(additionalProperties, "ObjectName")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "Response")

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

type NullableApplianceApiStatus struct {
	value *ApplianceApiStatus
	isSet bool
}

func (v NullableApplianceApiStatus) Get() *ApplianceApiStatus {
	return v.value
}

func (v *NullableApplianceApiStatus) Set(val *ApplianceApiStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceApiStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceApiStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceApiStatus(val *ApplianceApiStatus) *NullableApplianceApiStatus {
	return &NullableApplianceApiStatus{value: val, isSet: true}
}

func (v NullableApplianceApiStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceApiStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

