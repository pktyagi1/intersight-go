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

// CloudTfcOrganization Organizations are a shared space for teams to collaborate on workspaces in Terraform Cloud.
type CloudTfcOrganization struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The max number of active agents allowed in this organization.
	AgentCeiling *int64 `json:"AgentCeiling,omitempty"`
	// The admin email address associated with a user under this organization.
	Email *string `json:"Email,omitempty"`
	// The ID of the organization.
	Identity *string `json:"Identity,omitempty"`
	// The name of the organization.
	Name *string `json:"Name,omitempty"`
	// The number of teams under this organization.
	NumTeams *int64 `json:"NumTeams,omitempty"`
	// The number of users in this organization.
	NumUsers *int64 `json:"NumUsers,omitempty"`
	// The max number of simultaneous runs allowed in this organization.
	RunCeiling *int64 `json:"RunCeiling,omitempty"`
	// Total number of VCS providers in the organization.
	VcsProviders *int64 `json:"VcsProviders,omitempty"`
	Target *AssetTargetRelationship `json:"Target,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudTfcOrganization CloudTfcOrganization

// NewCloudTfcOrganization instantiates a new CloudTfcOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudTfcOrganization(classId string, objectType string) *CloudTfcOrganization {
	this := CloudTfcOrganization{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudTfcOrganizationWithDefaults instantiates a new CloudTfcOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudTfcOrganizationWithDefaults() *CloudTfcOrganization {
	this := CloudTfcOrganization{}
	var classId string = "cloud.TfcOrganization"
	this.ClassId = classId
	var objectType string = "cloud.TfcOrganization"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudTfcOrganization) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudTfcOrganization) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudTfcOrganization) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudTfcOrganization) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAgentCeiling returns the AgentCeiling field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetAgentCeiling() int64 {
	if o == nil || o.AgentCeiling == nil {
		var ret int64
		return ret
	}
	return *o.AgentCeiling
}

// GetAgentCeilingOk returns a tuple with the AgentCeiling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetAgentCeilingOk() (*int64, bool) {
	if o == nil || o.AgentCeiling == nil {
		return nil, false
	}
	return o.AgentCeiling, true
}

// HasAgentCeiling returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasAgentCeiling() bool {
	if o != nil && o.AgentCeiling != nil {
		return true
	}

	return false
}

// SetAgentCeiling gets a reference to the given int64 and assigns it to the AgentCeiling field.
func (o *CloudTfcOrganization) SetAgentCeiling(v int64) {
	o.AgentCeiling = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CloudTfcOrganization) SetEmail(v string) {
	o.Email = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *CloudTfcOrganization) SetIdentity(v string) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudTfcOrganization) SetName(v string) {
	o.Name = &v
}

// GetNumTeams returns the NumTeams field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetNumTeams() int64 {
	if o == nil || o.NumTeams == nil {
		var ret int64
		return ret
	}
	return *o.NumTeams
}

// GetNumTeamsOk returns a tuple with the NumTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetNumTeamsOk() (*int64, bool) {
	if o == nil || o.NumTeams == nil {
		return nil, false
	}
	return o.NumTeams, true
}

// HasNumTeams returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasNumTeams() bool {
	if o != nil && o.NumTeams != nil {
		return true
	}

	return false
}

// SetNumTeams gets a reference to the given int64 and assigns it to the NumTeams field.
func (o *CloudTfcOrganization) SetNumTeams(v int64) {
	o.NumTeams = &v
}

// GetNumUsers returns the NumUsers field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetNumUsers() int64 {
	if o == nil || o.NumUsers == nil {
		var ret int64
		return ret
	}
	return *o.NumUsers
}

// GetNumUsersOk returns a tuple with the NumUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetNumUsersOk() (*int64, bool) {
	if o == nil || o.NumUsers == nil {
		return nil, false
	}
	return o.NumUsers, true
}

// HasNumUsers returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasNumUsers() bool {
	if o != nil && o.NumUsers != nil {
		return true
	}

	return false
}

// SetNumUsers gets a reference to the given int64 and assigns it to the NumUsers field.
func (o *CloudTfcOrganization) SetNumUsers(v int64) {
	o.NumUsers = &v
}

// GetRunCeiling returns the RunCeiling field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetRunCeiling() int64 {
	if o == nil || o.RunCeiling == nil {
		var ret int64
		return ret
	}
	return *o.RunCeiling
}

// GetRunCeilingOk returns a tuple with the RunCeiling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetRunCeilingOk() (*int64, bool) {
	if o == nil || o.RunCeiling == nil {
		return nil, false
	}
	return o.RunCeiling, true
}

// HasRunCeiling returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasRunCeiling() bool {
	if o != nil && o.RunCeiling != nil {
		return true
	}

	return false
}

// SetRunCeiling gets a reference to the given int64 and assigns it to the RunCeiling field.
func (o *CloudTfcOrganization) SetRunCeiling(v int64) {
	o.RunCeiling = &v
}

// GetVcsProviders returns the VcsProviders field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetVcsProviders() int64 {
	if o == nil || o.VcsProviders == nil {
		var ret int64
		return ret
	}
	return *o.VcsProviders
}

// GetVcsProvidersOk returns a tuple with the VcsProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetVcsProvidersOk() (*int64, bool) {
	if o == nil || o.VcsProviders == nil {
		return nil, false
	}
	return o.VcsProviders, true
}

// HasVcsProviders returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasVcsProviders() bool {
	if o != nil && o.VcsProviders != nil {
		return true
	}

	return false
}

// SetVcsProviders gets a reference to the given int64 and assigns it to the VcsProviders field.
func (o *CloudTfcOrganization) SetVcsProviders(v int64) {
	o.VcsProviders = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetTarget() AssetTargetRelationship {
	if o == nil || o.Target == nil {
		var ret AssetTargetRelationship
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetTargetOk() (*AssetTargetRelationship, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AssetTargetRelationship and assigns it to the Target field.
func (o *CloudTfcOrganization) SetTarget(v AssetTargetRelationship) {
	o.Target = &v
}

func (o CloudTfcOrganization) MarshalJSON() ([]byte, error) {
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
	if o.AgentCeiling != nil {
		toSerialize["AgentCeiling"] = o.AgentCeiling
	}
	if o.Email != nil {
		toSerialize["Email"] = o.Email
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NumTeams != nil {
		toSerialize["NumTeams"] = o.NumTeams
	}
	if o.NumUsers != nil {
		toSerialize["NumUsers"] = o.NumUsers
	}
	if o.RunCeiling != nil {
		toSerialize["RunCeiling"] = o.RunCeiling
	}
	if o.VcsProviders != nil {
		toSerialize["VcsProviders"] = o.VcsProviders
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudTfcOrganization) UnmarshalJSON(bytes []byte) (err error) {
	type CloudTfcOrganizationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The max number of active agents allowed in this organization.
		AgentCeiling *int64 `json:"AgentCeiling,omitempty"`
		// The admin email address associated with a user under this organization.
		Email *string `json:"Email,omitempty"`
		// The ID of the organization.
		Identity *string `json:"Identity,omitempty"`
		// The name of the organization.
		Name *string `json:"Name,omitempty"`
		// The number of teams under this organization.
		NumTeams *int64 `json:"NumTeams,omitempty"`
		// The number of users in this organization.
		NumUsers *int64 `json:"NumUsers,omitempty"`
		// The max number of simultaneous runs allowed in this organization.
		RunCeiling *int64 `json:"RunCeiling,omitempty"`
		// Total number of VCS providers in the organization.
		VcsProviders *int64 `json:"VcsProviders,omitempty"`
		Target *AssetTargetRelationship `json:"Target,omitempty"`
	}

	varCloudTfcOrganizationWithoutEmbeddedStruct := CloudTfcOrganizationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudTfcOrganizationWithoutEmbeddedStruct)
	if err == nil {
		varCloudTfcOrganization := _CloudTfcOrganization{}
		varCloudTfcOrganization.ClassId = varCloudTfcOrganizationWithoutEmbeddedStruct.ClassId
		varCloudTfcOrganization.ObjectType = varCloudTfcOrganizationWithoutEmbeddedStruct.ObjectType
		varCloudTfcOrganization.AgentCeiling = varCloudTfcOrganizationWithoutEmbeddedStruct.AgentCeiling
		varCloudTfcOrganization.Email = varCloudTfcOrganizationWithoutEmbeddedStruct.Email
		varCloudTfcOrganization.Identity = varCloudTfcOrganizationWithoutEmbeddedStruct.Identity
		varCloudTfcOrganization.Name = varCloudTfcOrganizationWithoutEmbeddedStruct.Name
		varCloudTfcOrganization.NumTeams = varCloudTfcOrganizationWithoutEmbeddedStruct.NumTeams
		varCloudTfcOrganization.NumUsers = varCloudTfcOrganizationWithoutEmbeddedStruct.NumUsers
		varCloudTfcOrganization.RunCeiling = varCloudTfcOrganizationWithoutEmbeddedStruct.RunCeiling
		varCloudTfcOrganization.VcsProviders = varCloudTfcOrganizationWithoutEmbeddedStruct.VcsProviders
		varCloudTfcOrganization.Target = varCloudTfcOrganizationWithoutEmbeddedStruct.Target
		*o = CloudTfcOrganization(varCloudTfcOrganization)
	} else {
		return err
	}

	varCloudTfcOrganization := _CloudTfcOrganization{}

	err = json.Unmarshal(bytes, &varCloudTfcOrganization)
	if err == nil {
		o.MoBaseMo = varCloudTfcOrganization.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AgentCeiling")
		delete(additionalProperties, "Email")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NumTeams")
		delete(additionalProperties, "NumUsers")
		delete(additionalProperties, "RunCeiling")
		delete(additionalProperties, "VcsProviders")
		delete(additionalProperties, "Target")

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

type NullableCloudTfcOrganization struct {
	value *CloudTfcOrganization
	isSet bool
}

func (v NullableCloudTfcOrganization) Get() *CloudTfcOrganization {
	return v.value
}

func (v *NullableCloudTfcOrganization) Set(val *CloudTfcOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudTfcOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudTfcOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudTfcOrganization(val *CloudTfcOrganization) *NullableCloudTfcOrganization {
	return &NullableCloudTfcOrganization{value: val, isSet: true}
}

func (v NullableCloudTfcOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudTfcOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


