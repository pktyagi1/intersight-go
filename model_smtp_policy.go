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

// SmtpPolicy Name that identifies the SMTP Policy.
type SmtpPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If enabled, controls the state of the SMTP client service on the managed device.
	Enabled *bool `json:"Enabled,omitempty"`
	// Minimum fault severity level to receive email notifications. Email notifications are sent for all faults whose severity is equal to or greater than the chosen level. * `critical` - Minimum severity to report is critical. * `condition` - Minimum severity to report is informational. * `warning` - Minimum severity to report is warning. * `minor` - Minimum severity to report is minor. * `major` - Minimum severity to report is major.
	MinSeverity *string `json:"MinSeverity,omitempty"`
	// The email address entered here will be displayed as the from address (mail received from address) of all the SMTP mail alerts that are received. If not configured, the hostname of the server is used in the from address field.
	SenderEmail *string `json:"SenderEmail,omitempty"`
	// Port number used by the SMTP server for outgoing SMTP communication.
	SmtpPort *int64 `json:"SmtpPort,omitempty"`
	SmtpRecipients []string `json:"SmtpRecipients,omitempty"`
	// IP address or hostname of the SMTP server. The SMTP server is used by the managed device to send email notifications.
	SmtpServer *string `json:"SmtpServer,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SmtpPolicy SmtpPolicy

// NewSmtpPolicy instantiates a new SmtpPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmtpPolicy(classId string, objectType string) *SmtpPolicy {
	this := SmtpPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var minSeverity string = "critical"
	this.MinSeverity = &minSeverity
	var smtpPort int64 = 25
	this.SmtpPort = &smtpPort
	return &this
}

// NewSmtpPolicyWithDefaults instantiates a new SmtpPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmtpPolicyWithDefaults() *SmtpPolicy {
	this := SmtpPolicy{}
	var classId string = "smtp.Policy"
	this.ClassId = classId
	var objectType string = "smtp.Policy"
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var minSeverity string = "critical"
	this.MinSeverity = &minSeverity
	var smtpPort int64 = 25
	this.SmtpPort = &smtpPort
	return &this
}

// GetClassId returns the ClassId field value
func (o *SmtpPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SmtpPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SmtpPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SmtpPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SmtpPolicy) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SmtpPolicy) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SmtpPolicy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMinSeverity returns the MinSeverity field value if set, zero value otherwise.
func (o *SmtpPolicy) GetMinSeverity() string {
	if o == nil || o.MinSeverity == nil {
		var ret string
		return ret
	}
	return *o.MinSeverity
}

// GetMinSeverityOk returns a tuple with the MinSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetMinSeverityOk() (*string, bool) {
	if o == nil || o.MinSeverity == nil {
		return nil, false
	}
	return o.MinSeverity, true
}

// HasMinSeverity returns a boolean if a field has been set.
func (o *SmtpPolicy) HasMinSeverity() bool {
	if o != nil && o.MinSeverity != nil {
		return true
	}

	return false
}

// SetMinSeverity gets a reference to the given string and assigns it to the MinSeverity field.
func (o *SmtpPolicy) SetMinSeverity(v string) {
	o.MinSeverity = &v
}

// GetSenderEmail returns the SenderEmail field value if set, zero value otherwise.
func (o *SmtpPolicy) GetSenderEmail() string {
	if o == nil || o.SenderEmail == nil {
		var ret string
		return ret
	}
	return *o.SenderEmail
}

// GetSenderEmailOk returns a tuple with the SenderEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetSenderEmailOk() (*string, bool) {
	if o == nil || o.SenderEmail == nil {
		return nil, false
	}
	return o.SenderEmail, true
}

// HasSenderEmail returns a boolean if a field has been set.
func (o *SmtpPolicy) HasSenderEmail() bool {
	if o != nil && o.SenderEmail != nil {
		return true
	}

	return false
}

// SetSenderEmail gets a reference to the given string and assigns it to the SenderEmail field.
func (o *SmtpPolicy) SetSenderEmail(v string) {
	o.SenderEmail = &v
}

// GetSmtpPort returns the SmtpPort field value if set, zero value otherwise.
func (o *SmtpPolicy) GetSmtpPort() int64 {
	if o == nil || o.SmtpPort == nil {
		var ret int64
		return ret
	}
	return *o.SmtpPort
}

// GetSmtpPortOk returns a tuple with the SmtpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetSmtpPortOk() (*int64, bool) {
	if o == nil || o.SmtpPort == nil {
		return nil, false
	}
	return o.SmtpPort, true
}

// HasSmtpPort returns a boolean if a field has been set.
func (o *SmtpPolicy) HasSmtpPort() bool {
	if o != nil && o.SmtpPort != nil {
		return true
	}

	return false
}

// SetSmtpPort gets a reference to the given int64 and assigns it to the SmtpPort field.
func (o *SmtpPolicy) SetSmtpPort(v int64) {
	o.SmtpPort = &v
}

// GetSmtpRecipients returns the SmtpRecipients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpPolicy) GetSmtpRecipients() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.SmtpRecipients
}

// GetSmtpRecipientsOk returns a tuple with the SmtpRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpPolicy) GetSmtpRecipientsOk() (*[]string, bool) {
	if o == nil || o.SmtpRecipients == nil {
		return nil, false
	}
	return &o.SmtpRecipients, true
}

// HasSmtpRecipients returns a boolean if a field has been set.
func (o *SmtpPolicy) HasSmtpRecipients() bool {
	if o != nil && o.SmtpRecipients != nil {
		return true
	}

	return false
}

// SetSmtpRecipients gets a reference to the given []string and assigns it to the SmtpRecipients field.
func (o *SmtpPolicy) SetSmtpRecipients(v []string) {
	o.SmtpRecipients = v
}

// GetSmtpServer returns the SmtpServer field value if set, zero value otherwise.
func (o *SmtpPolicy) GetSmtpServer() string {
	if o == nil || o.SmtpServer == nil {
		var ret string
		return ret
	}
	return *o.SmtpServer
}

// GetSmtpServerOk returns a tuple with the SmtpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetSmtpServerOk() (*string, bool) {
	if o == nil || o.SmtpServer == nil {
		return nil, false
	}
	return o.SmtpServer, true
}

// HasSmtpServer returns a boolean if a field has been set.
func (o *SmtpPolicy) HasSmtpServer() bool {
	if o != nil && o.SmtpServer != nil {
		return true
	}

	return false
}

// SetSmtpServer gets a reference to the given string and assigns it to the SmtpServer field.
func (o *SmtpPolicy) SetSmtpServer(v string) {
	o.SmtpServer = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *SmtpPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *SmtpPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *SmtpPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil  {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *SmtpPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *SmtpPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o SmtpPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.MinSeverity != nil {
		toSerialize["MinSeverity"] = o.MinSeverity
	}
	if o.SenderEmail != nil {
		toSerialize["SenderEmail"] = o.SenderEmail
	}
	if o.SmtpPort != nil {
		toSerialize["SmtpPort"] = o.SmtpPort
	}
	if o.SmtpRecipients != nil {
		toSerialize["SmtpRecipients"] = o.SmtpRecipients
	}
	if o.SmtpServer != nil {
		toSerialize["SmtpServer"] = o.SmtpServer
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SmtpPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type SmtpPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// If enabled, controls the state of the SMTP client service on the managed device.
		Enabled *bool `json:"Enabled,omitempty"`
		// Minimum fault severity level to receive email notifications. Email notifications are sent for all faults whose severity is equal to or greater than the chosen level. * `critical` - Minimum severity to report is critical. * `condition` - Minimum severity to report is informational. * `warning` - Minimum severity to report is warning. * `minor` - Minimum severity to report is minor. * `major` - Minimum severity to report is major.
		MinSeverity *string `json:"MinSeverity,omitempty"`
		// The email address entered here will be displayed as the from address (mail received from address) of all the SMTP mail alerts that are received. If not configured, the hostname of the server is used in the from address field.
		SenderEmail *string `json:"SenderEmail,omitempty"`
		// Port number used by the SMTP server for outgoing SMTP communication.
		SmtpPort *int64 `json:"SmtpPort,omitempty"`
		SmtpRecipients []string `json:"SmtpRecipients,omitempty"`
		// IP address or hostname of the SMTP server. The SMTP server is used by the managed device to send email notifications.
		SmtpServer *string `json:"SmtpServer,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varSmtpPolicyWithoutEmbeddedStruct := SmtpPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSmtpPolicyWithoutEmbeddedStruct)
	if err == nil {
		varSmtpPolicy := _SmtpPolicy{}
		varSmtpPolicy.ClassId = varSmtpPolicyWithoutEmbeddedStruct.ClassId
		varSmtpPolicy.ObjectType = varSmtpPolicyWithoutEmbeddedStruct.ObjectType
		varSmtpPolicy.Enabled = varSmtpPolicyWithoutEmbeddedStruct.Enabled
		varSmtpPolicy.MinSeverity = varSmtpPolicyWithoutEmbeddedStruct.MinSeverity
		varSmtpPolicy.SenderEmail = varSmtpPolicyWithoutEmbeddedStruct.SenderEmail
		varSmtpPolicy.SmtpPort = varSmtpPolicyWithoutEmbeddedStruct.SmtpPort
		varSmtpPolicy.SmtpRecipients = varSmtpPolicyWithoutEmbeddedStruct.SmtpRecipients
		varSmtpPolicy.SmtpServer = varSmtpPolicyWithoutEmbeddedStruct.SmtpServer
		varSmtpPolicy.Organization = varSmtpPolicyWithoutEmbeddedStruct.Organization
		varSmtpPolicy.Profiles = varSmtpPolicyWithoutEmbeddedStruct.Profiles
		*o = SmtpPolicy(varSmtpPolicy)
	} else {
		return err
	}

	varSmtpPolicy := _SmtpPolicy{}

	err = json.Unmarshal(bytes, &varSmtpPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varSmtpPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "MinSeverity")
		delete(additionalProperties, "SenderEmail")
		delete(additionalProperties, "SmtpPort")
		delete(additionalProperties, "SmtpRecipients")
		delete(additionalProperties, "SmtpServer")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableSmtpPolicy struct {
	value *SmtpPolicy
	isSet bool
}

func (v NullableSmtpPolicy) Get() *SmtpPolicy {
	return v.value
}

func (v *NullableSmtpPolicy) Set(val *SmtpPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableSmtpPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableSmtpPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmtpPolicy(val *SmtpPolicy) *NullableSmtpPolicy {
	return &NullableSmtpPolicy{value: val, isSet: true}
}

func (v NullableSmtpPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmtpPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


