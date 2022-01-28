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
	"time"
)

// IwotenantTenantStatusAllOf Definition of the list of properties defined in 'iwotenant.TenantStatus', excluding properties defined in parent classes.
type IwotenantTenantStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The deployStatus provides the current status of this deployment. * `NotStarted` - The workflow to deploy the tenant cluster has not yet started. * `InProgress` - The workflow to deploy the tenant cluster in progress. All the tasks required for thesuccessful tenant creation are running. * `Completed` - The workflow to deploy the tenant cluster has completed and health checks have passed. * `Failed` - The workflow to deploy the tenant cluster has failed. Detailed reason for the failure isprovided from Tenant.deployStatus.
	DeployStatus *string `json:"DeployStatus,omitempty"`
	// The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId.
	IwoId *string `json:"IwoId,omitempty"`
	// During IWO tenant upgrade (or reconfiguration), deployStatus is set to InProgress and referenceTime set to current time. When tenant upgrade (or reconfiguration) does not complete within a pre-defined time using this as reference, deployStatus is set as Failed.
	ReferenceTime *time.Time `json:"ReferenceTime,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IwotenantTenantStatusAllOf IwotenantTenantStatusAllOf

// NewIwotenantTenantStatusAllOf instantiates a new IwotenantTenantStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIwotenantTenantStatusAllOf(classId string, objectType string) *IwotenantTenantStatusAllOf {
	this := IwotenantTenantStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIwotenantTenantStatusAllOfWithDefaults instantiates a new IwotenantTenantStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIwotenantTenantStatusAllOfWithDefaults() *IwotenantTenantStatusAllOf {
	this := IwotenantTenantStatusAllOf{}
	var classId string = "iwotenant.TenantStatus"
	this.ClassId = classId
	var objectType string = "iwotenant.TenantStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IwotenantTenantStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IwotenantTenantStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IwotenantTenantStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IwotenantTenantStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IwotenantTenantStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IwotenantTenantStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeployStatus returns the DeployStatus field value if set, zero value otherwise.
func (o *IwotenantTenantStatusAllOf) GetDeployStatus() string {
	if o == nil || o.DeployStatus == nil {
		var ret string
		return ret
	}
	return *o.DeployStatus
}

// GetDeployStatusOk returns a tuple with the DeployStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantTenantStatusAllOf) GetDeployStatusOk() (*string, bool) {
	if o == nil || o.DeployStatus == nil {
		return nil, false
	}
	return o.DeployStatus, true
}

// HasDeployStatus returns a boolean if a field has been set.
func (o *IwotenantTenantStatusAllOf) HasDeployStatus() bool {
	if o != nil && o.DeployStatus != nil {
		return true
	}

	return false
}

// SetDeployStatus gets a reference to the given string and assigns it to the DeployStatus field.
func (o *IwotenantTenantStatusAllOf) SetDeployStatus(v string) {
	o.DeployStatus = &v
}

// GetIwoId returns the IwoId field value if set, zero value otherwise.
func (o *IwotenantTenantStatusAllOf) GetIwoId() string {
	if o == nil || o.IwoId == nil {
		var ret string
		return ret
	}
	return *o.IwoId
}

// GetIwoIdOk returns a tuple with the IwoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantTenantStatusAllOf) GetIwoIdOk() (*string, bool) {
	if o == nil || o.IwoId == nil {
		return nil, false
	}
	return o.IwoId, true
}

// HasIwoId returns a boolean if a field has been set.
func (o *IwotenantTenantStatusAllOf) HasIwoId() bool {
	if o != nil && o.IwoId != nil {
		return true
	}

	return false
}

// SetIwoId gets a reference to the given string and assigns it to the IwoId field.
func (o *IwotenantTenantStatusAllOf) SetIwoId(v string) {
	o.IwoId = &v
}

// GetReferenceTime returns the ReferenceTime field value if set, zero value otherwise.
func (o *IwotenantTenantStatusAllOf) GetReferenceTime() time.Time {
	if o == nil || o.ReferenceTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ReferenceTime
}

// GetReferenceTimeOk returns a tuple with the ReferenceTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantTenantStatusAllOf) GetReferenceTimeOk() (*time.Time, bool) {
	if o == nil || o.ReferenceTime == nil {
		return nil, false
	}
	return o.ReferenceTime, true
}

// HasReferenceTime returns a boolean if a field has been set.
func (o *IwotenantTenantStatusAllOf) HasReferenceTime() bool {
	if o != nil && o.ReferenceTime != nil {
		return true
	}

	return false
}

// SetReferenceTime gets a reference to the given time.Time and assigns it to the ReferenceTime field.
func (o *IwotenantTenantStatusAllOf) SetReferenceTime(v time.Time) {
	o.ReferenceTime = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IwotenantTenantStatusAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantTenantStatusAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IwotenantTenantStatusAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IwotenantTenantStatusAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o IwotenantTenantStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeployStatus != nil {
		toSerialize["DeployStatus"] = o.DeployStatus
	}
	if o.IwoId != nil {
		toSerialize["IwoId"] = o.IwoId
	}
	if o.ReferenceTime != nil {
		toSerialize["ReferenceTime"] = o.ReferenceTime
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IwotenantTenantStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIwotenantTenantStatusAllOf := _IwotenantTenantStatusAllOf{}

	if err = json.Unmarshal(bytes, &varIwotenantTenantStatusAllOf); err == nil {
		*o = IwotenantTenantStatusAllOf(varIwotenantTenantStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeployStatus")
		delete(additionalProperties, "IwoId")
		delete(additionalProperties, "ReferenceTime")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIwotenantTenantStatusAllOf struct {
	value *IwotenantTenantStatusAllOf
	isSet bool
}

func (v NullableIwotenantTenantStatusAllOf) Get() *IwotenantTenantStatusAllOf {
	return v.value
}

func (v *NullableIwotenantTenantStatusAllOf) Set(val *IwotenantTenantStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIwotenantTenantStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIwotenantTenantStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIwotenantTenantStatusAllOf(val *IwotenantTenantStatusAllOf) *NullableIwotenantTenantStatusAllOf {
	return &NullableIwotenantTenantStatusAllOf{value: val, isSet: true}
}

func (v NullableIwotenantTenantStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIwotenantTenantStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


