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
)

// AssetSubscriptionAllOf Definition of the list of properties defined in 'asset.Subscription', excluding properties defined in parent classes.
type AssetSubscriptionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Application name reported by Cisco Install Base.
	ApplicationName *string `json:"ApplicationName,omitempty"`
	// Identifies the consumption-based subscription.
	SubscriptionRefId *string `json:"SubscriptionRefId,omitempty"`
	// An array of relationships to assetDeployment resources.
	Deployments []AssetDeploymentRelationship `json:"Deployments,omitempty"`
	SubscriptionAccount *AssetSubscriptionAccountRelationship `json:"SubscriptionAccount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetSubscriptionAllOf AssetSubscriptionAllOf

// NewAssetSubscriptionAllOf instantiates a new AssetSubscriptionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetSubscriptionAllOf(classId string, objectType string) *AssetSubscriptionAllOf {
	this := AssetSubscriptionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetSubscriptionAllOfWithDefaults instantiates a new AssetSubscriptionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetSubscriptionAllOfWithDefaults() *AssetSubscriptionAllOf {
	this := AssetSubscriptionAllOf{}
	var classId string = "asset.Subscription"
	this.ClassId = classId
	var objectType string = "asset.Subscription"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetSubscriptionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetSubscriptionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetSubscriptionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetSubscriptionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetSubscriptionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetSubscriptionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *AssetSubscriptionAllOf) GetApplicationName() string {
	if o == nil || o.ApplicationName == nil {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSubscriptionAllOf) GetApplicationNameOk() (*string, bool) {
	if o == nil || o.ApplicationName == nil {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *AssetSubscriptionAllOf) HasApplicationName() bool {
	if o != nil && o.ApplicationName != nil {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *AssetSubscriptionAllOf) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetSubscriptionRefId returns the SubscriptionRefId field value if set, zero value otherwise.
func (o *AssetSubscriptionAllOf) GetSubscriptionRefId() string {
	if o == nil || o.SubscriptionRefId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionRefId
}

// GetSubscriptionRefIdOk returns a tuple with the SubscriptionRefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSubscriptionAllOf) GetSubscriptionRefIdOk() (*string, bool) {
	if o == nil || o.SubscriptionRefId == nil {
		return nil, false
	}
	return o.SubscriptionRefId, true
}

// HasSubscriptionRefId returns a boolean if a field has been set.
func (o *AssetSubscriptionAllOf) HasSubscriptionRefId() bool {
	if o != nil && o.SubscriptionRefId != nil {
		return true
	}

	return false
}

// SetSubscriptionRefId gets a reference to the given string and assigns it to the SubscriptionRefId field.
func (o *AssetSubscriptionAllOf) SetSubscriptionRefId(v string) {
	o.SubscriptionRefId = &v
}

// GetDeployments returns the Deployments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetSubscriptionAllOf) GetDeployments() []AssetDeploymentRelationship {
	if o == nil  {
		var ret []AssetDeploymentRelationship
		return ret
	}
	return o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetSubscriptionAllOf) GetDeploymentsOk() (*[]AssetDeploymentRelationship, bool) {
	if o == nil || o.Deployments == nil {
		return nil, false
	}
	return &o.Deployments, true
}

// HasDeployments returns a boolean if a field has been set.
func (o *AssetSubscriptionAllOf) HasDeployments() bool {
	if o != nil && o.Deployments != nil {
		return true
	}

	return false
}

// SetDeployments gets a reference to the given []AssetDeploymentRelationship and assigns it to the Deployments field.
func (o *AssetSubscriptionAllOf) SetDeployments(v []AssetDeploymentRelationship) {
	o.Deployments = v
}

// GetSubscriptionAccount returns the SubscriptionAccount field value if set, zero value otherwise.
func (o *AssetSubscriptionAllOf) GetSubscriptionAccount() AssetSubscriptionAccountRelationship {
	if o == nil || o.SubscriptionAccount == nil {
		var ret AssetSubscriptionAccountRelationship
		return ret
	}
	return *o.SubscriptionAccount
}

// GetSubscriptionAccountOk returns a tuple with the SubscriptionAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSubscriptionAllOf) GetSubscriptionAccountOk() (*AssetSubscriptionAccountRelationship, bool) {
	if o == nil || o.SubscriptionAccount == nil {
		return nil, false
	}
	return o.SubscriptionAccount, true
}

// HasSubscriptionAccount returns a boolean if a field has been set.
func (o *AssetSubscriptionAllOf) HasSubscriptionAccount() bool {
	if o != nil && o.SubscriptionAccount != nil {
		return true
	}

	return false
}

// SetSubscriptionAccount gets a reference to the given AssetSubscriptionAccountRelationship and assigns it to the SubscriptionAccount field.
func (o *AssetSubscriptionAllOf) SetSubscriptionAccount(v AssetSubscriptionAccountRelationship) {
	o.SubscriptionAccount = &v
}

func (o AssetSubscriptionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ApplicationName != nil {
		toSerialize["ApplicationName"] = o.ApplicationName
	}
	if o.SubscriptionRefId != nil {
		toSerialize["SubscriptionRefId"] = o.SubscriptionRefId
	}
	if o.Deployments != nil {
		toSerialize["Deployments"] = o.Deployments
	}
	if o.SubscriptionAccount != nil {
		toSerialize["SubscriptionAccount"] = o.SubscriptionAccount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetSubscriptionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetSubscriptionAllOf := _AssetSubscriptionAllOf{}

	if err = json.Unmarshal(bytes, &varAssetSubscriptionAllOf); err == nil {
		*o = AssetSubscriptionAllOf(varAssetSubscriptionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApplicationName")
		delete(additionalProperties, "SubscriptionRefId")
		delete(additionalProperties, "Deployments")
		delete(additionalProperties, "SubscriptionAccount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetSubscriptionAllOf struct {
	value *AssetSubscriptionAllOf
	isSet bool
}

func (v NullableAssetSubscriptionAllOf) Get() *AssetSubscriptionAllOf {
	return v.value
}

func (v *NullableAssetSubscriptionAllOf) Set(val *AssetSubscriptionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetSubscriptionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetSubscriptionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetSubscriptionAllOf(val *AssetSubscriptionAllOf) *NullableAssetSubscriptionAllOf {
	return &NullableAssetSubscriptionAllOf{value: val, isSet: true}
}

func (v NullableAssetSubscriptionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetSubscriptionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


