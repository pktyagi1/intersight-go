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

// IamPrivateKeySpecAllOf Definition of the list of properties defined in 'iam.PrivateKeySpec', excluding properties defined in parent classes.
type IamPrivateKeySpecAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	Algorithm NullablePkixKeyGenerationSpec `json:"Algorithm,omitempty"`
	CertificateRequest *IamCertificateRequestRelationship `json:"CertificateRequest,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamPrivateKeySpecAllOf IamPrivateKeySpecAllOf

// NewIamPrivateKeySpecAllOf instantiates a new IamPrivateKeySpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamPrivateKeySpecAllOf(classId string, objectType string) *IamPrivateKeySpecAllOf {
	this := IamPrivateKeySpecAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamPrivateKeySpecAllOfWithDefaults instantiates a new IamPrivateKeySpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamPrivateKeySpecAllOfWithDefaults() *IamPrivateKeySpecAllOf {
	this := IamPrivateKeySpecAllOf{}
	var classId string = "iam.PrivateKeySpec"
	this.ClassId = classId
	var objectType string = "iam.PrivateKeySpec"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamPrivateKeySpecAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamPrivateKeySpecAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamPrivateKeySpecAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamPrivateKeySpecAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamPrivateKeySpecAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamPrivateKeySpecAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPrivateKeySpecAllOf) GetAlgorithm() PkixKeyGenerationSpec {
	if o == nil || o.Algorithm.Get() == nil {
		var ret PkixKeyGenerationSpec
		return ret
	}
	return *o.Algorithm.Get()
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPrivateKeySpecAllOf) GetAlgorithmOk() (*PkixKeyGenerationSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Algorithm.Get(), o.Algorithm.IsSet()
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *IamPrivateKeySpecAllOf) HasAlgorithm() bool {
	if o != nil && o.Algorithm.IsSet() {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given NullablePkixKeyGenerationSpec and assigns it to the Algorithm field.
func (o *IamPrivateKeySpecAllOf) SetAlgorithm(v PkixKeyGenerationSpec) {
	o.Algorithm.Set(&v)
}
// SetAlgorithmNil sets the value for Algorithm to be an explicit nil
func (o *IamPrivateKeySpecAllOf) SetAlgorithmNil() {
	o.Algorithm.Set(nil)
}

// UnsetAlgorithm ensures that no value is present for Algorithm, not even an explicit nil
func (o *IamPrivateKeySpecAllOf) UnsetAlgorithm() {
	o.Algorithm.Unset()
}

// GetCertificateRequest returns the CertificateRequest field value if set, zero value otherwise.
func (o *IamPrivateKeySpecAllOf) GetCertificateRequest() IamCertificateRequestRelationship {
	if o == nil || o.CertificateRequest == nil {
		var ret IamCertificateRequestRelationship
		return ret
	}
	return *o.CertificateRequest
}

// GetCertificateRequestOk returns a tuple with the CertificateRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivateKeySpecAllOf) GetCertificateRequestOk() (*IamCertificateRequestRelationship, bool) {
	if o == nil || o.CertificateRequest == nil {
		return nil, false
	}
	return o.CertificateRequest, true
}

// HasCertificateRequest returns a boolean if a field has been set.
func (o *IamPrivateKeySpecAllOf) HasCertificateRequest() bool {
	if o != nil && o.CertificateRequest != nil {
		return true
	}

	return false
}

// SetCertificateRequest gets a reference to the given IamCertificateRequestRelationship and assigns it to the CertificateRequest field.
func (o *IamPrivateKeySpecAllOf) SetCertificateRequest(v IamCertificateRequestRelationship) {
	o.CertificateRequest = &v
}

func (o IamPrivateKeySpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Algorithm.IsSet() {
		toSerialize["Algorithm"] = o.Algorithm.Get()
	}
	if o.CertificateRequest != nil {
		toSerialize["CertificateRequest"] = o.CertificateRequest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamPrivateKeySpecAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamPrivateKeySpecAllOf := _IamPrivateKeySpecAllOf{}

	if err = json.Unmarshal(bytes, &varIamPrivateKeySpecAllOf); err == nil {
		*o = IamPrivateKeySpecAllOf(varIamPrivateKeySpecAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Algorithm")
		delete(additionalProperties, "CertificateRequest")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamPrivateKeySpecAllOf struct {
	value *IamPrivateKeySpecAllOf
	isSet bool
}

func (v NullableIamPrivateKeySpecAllOf) Get() *IamPrivateKeySpecAllOf {
	return v.value
}

func (v *NullableIamPrivateKeySpecAllOf) Set(val *IamPrivateKeySpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamPrivateKeySpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamPrivateKeySpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamPrivateKeySpecAllOf(val *IamPrivateKeySpecAllOf) *NullableIamPrivateKeySpecAllOf {
	return &NullableIamPrivateKeySpecAllOf{value: val, isSet: true}
}

func (v NullableIamPrivateKeySpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamPrivateKeySpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


