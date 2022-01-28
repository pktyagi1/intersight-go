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
)

// VnicVifStatusAllOf Definition of the list of properties defined in 'vnic.VifStatus', excluding properties defined in parent classes.
type VnicVifStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the vNIC for which the status is reported.
	Name *string `json:"Name,omitempty"`
	// The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure.
	Reason *string `json:"Reason,omitempty"`
	// Indicates if the vNIC / vHBA is ready for deploy or not. * `ok` - No issues with the LCP/SCP/VIF. * `error` - The LCP/SCP/VIF cannot be deployed due to error. * `validating` - Validation in progress for the LCP.
	Status *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicVifStatusAllOf VnicVifStatusAllOf

// NewVnicVifStatusAllOf instantiates a new VnicVifStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicVifStatusAllOf(classId string, objectType string) *VnicVifStatusAllOf {
	this := VnicVifStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "ok"
	this.Status = &status
	return &this
}

// NewVnicVifStatusAllOfWithDefaults instantiates a new VnicVifStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicVifStatusAllOfWithDefaults() *VnicVifStatusAllOf {
	this := VnicVifStatusAllOf{}
	var classId string = "vnic.VifStatus"
	this.ClassId = classId
	var objectType string = "vnic.VifStatus"
	this.ObjectType = objectType
	var status string = "ok"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicVifStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicVifStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicVifStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicVifStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicVifStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicVifStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VnicVifStatusAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVifStatusAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VnicVifStatusAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VnicVifStatusAllOf) SetName(v string) {
	o.Name = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *VnicVifStatusAllOf) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVifStatusAllOf) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *VnicVifStatusAllOf) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *VnicVifStatusAllOf) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VnicVifStatusAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVifStatusAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VnicVifStatusAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VnicVifStatusAllOf) SetStatus(v string) {
	o.Status = &v
}

func (o VnicVifStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicVifStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicVifStatusAllOf := _VnicVifStatusAllOf{}

	if err = json.Unmarshal(bytes, &varVnicVifStatusAllOf); err == nil {
		*o = VnicVifStatusAllOf(varVnicVifStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "Status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicVifStatusAllOf struct {
	value *VnicVifStatusAllOf
	isSet bool
}

func (v NullableVnicVifStatusAllOf) Get() *VnicVifStatusAllOf {
	return v.value
}

func (v *NullableVnicVifStatusAllOf) Set(val *VnicVifStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicVifStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicVifStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicVifStatusAllOf(val *VnicVifStatusAllOf) *NullableVnicVifStatusAllOf {
	return &NullableVnicVifStatusAllOf{value: val, isSet: true}
}

func (v NullableVnicVifStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicVifStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


