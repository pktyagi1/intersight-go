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

// KubernetesAddonRepository Docker registry or helm repository which hosts helm charts and docker images.
type KubernetesAddonRepository struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	CaCert NullableX509Certificate `json:"CaCert,omitempty"`
	// Allow connecting to http registries or https registries which do not have certificate signed by a well known CA.
	InsecureSkipVerification *bool `json:"InsecureSkipVerification,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Indicates whether the value of the 'token' property has been set.
	IsTokenSet *bool `json:"IsTokenSet,omitempty"`
	// Name of the addon repository or registry.
	Name *string `json:"Name,omitempty"`
	// URL for the repository where the addon is hosted.
	RepositoryUrl *string `json:"RepositoryUrl,omitempty"`
	// Username to authenticate to the addon registry.
	Username *string `json:"Username,omitempty"`
	Catalog *WorkflowCatalogRelationship `json:"Catalog,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAddonRepository KubernetesAddonRepository

// NewKubernetesAddonRepository instantiates a new KubernetesAddonRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAddonRepository(classId string, objectType string) *KubernetesAddonRepository {
	this := KubernetesAddonRepository{}
	this.ClassId = classId
	this.ObjectType = objectType
	var insecureSkipVerification bool = false
	this.InsecureSkipVerification = &insecureSkipVerification
	return &this
}

// NewKubernetesAddonRepositoryWithDefaults instantiates a new KubernetesAddonRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAddonRepositoryWithDefaults() *KubernetesAddonRepository {
	this := KubernetesAddonRepository{}
	var classId string = "kubernetes.AddonRepository"
	this.ClassId = classId
	var objectType string = "kubernetes.AddonRepository"
	this.ObjectType = objectType
	var insecureSkipVerification bool = false
	this.InsecureSkipVerification = &insecureSkipVerification
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAddonRepository) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonRepository) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAddonRepository) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAddonRepository) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonRepository) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAddonRepository) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCaCert returns the CaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAddonRepository) GetCaCert() X509Certificate {
	if o == nil || o.CaCert.Get() == nil {
		var ret X509Certificate
		return ret
	}
	return *o.CaCert.Get()
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAddonRepository) GetCaCertOk() (*X509Certificate, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CaCert.Get(), o.CaCert.IsSet()
}

// HasCaCert returns a boolean if a field has been set.
func (o *KubernetesAddonRepository) HasCaCert() bool {
	if o != nil && o.CaCert.IsSet() {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given NullableX509Certificate and assigns it to the CaCert field.
func (o *KubernetesAddonRepository) SetCaCert(v X509Certificate) {
	o.CaCert.Set(&v)
}
// SetCaCertNil sets the value for CaCert to be an explicit nil
func (o *KubernetesAddonRepository) SetCaCertNil() {
	o.CaCert.Set(nil)
}

// UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
func (o *KubernetesAddonRepository) UnsetCaCert() {
	o.CaCert.Unset()
}

// GetInsecureSkipVerification returns the InsecureSkipVerification field value if set, zero value otherwise.
func (o *KubernetesAddonRepository) GetInsecureSkipVerification() bool {
	if o == nil || o.InsecureSkipVerification == nil {
		var ret bool
		return ret
	}
	return *o.InsecureSkipVerification
}

// GetInsecureSkipVerificationOk returns a tuple with the InsecureSkipVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonRepository) GetInsecureSkipVerificationOk() (*bool, bool) {
	if o == nil || o.InsecureSkipVerification == nil {
		return nil, false
	}
	return o.InsecureSkipVerification, true
}

// HasInsecureSkipVerification returns a boolean if a field has been set.
func (o *KubernetesAddonRepository) HasInsecureSkipVerification() bool {
	if o != nil && o.InsecureSkipVerification != nil {
		return true
	}

	return false
}

// SetInsecureSkipVerification gets a reference to the given bool and assigns it to the InsecureSkipVerification field.
func (o *KubernetesAddonRepository) SetInsecureSkipVerification(v bool) {
	o.InsecureSkipVerification = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *KubernetesAddonRepository) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonRepository) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *KubernetesAddonRepository) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *KubernetesAddonRepository) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetIsTokenSet returns the IsTokenSet field value if set, zero value otherwise.
func (o *KubernetesAddonRepository) GetIsTokenSet() bool {
	if o == nil || o.IsTokenSet == nil {
		var ret bool
		return ret
	}
	return *o.IsTokenSet
}

// GetIsTokenSetOk returns a tuple with the IsTokenSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonRepository) GetIsTokenSetOk() (*bool, bool) {
	if o == nil || o.IsTokenSet == nil {
		return nil, false
	}
	return o.IsTokenSet, true
}

// HasIsTokenSet returns a boolean if a field has been set.
func (o *KubernetesAddonRepository) HasIsTokenSet() bool {
	if o != nil && o.IsTokenSet != nil {
		return true
	}

	return false
}

// SetIsTokenSet gets a reference to the given bool and assigns it to the IsTokenSet field.
func (o *KubernetesAddonRepository) SetIsTokenSet(v bool) {
	o.IsTokenSet = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesAddonRepository) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonRepository) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesAddonRepository) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesAddonRepository) SetName(v string) {
	o.Name = &v
}

// GetRepositoryUrl returns the RepositoryUrl field value if set, zero value otherwise.
func (o *KubernetesAddonRepository) GetRepositoryUrl() string {
	if o == nil || o.RepositoryUrl == nil {
		var ret string
		return ret
	}
	return *o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonRepository) GetRepositoryUrlOk() (*string, bool) {
	if o == nil || o.RepositoryUrl == nil {
		return nil, false
	}
	return o.RepositoryUrl, true
}

// HasRepositoryUrl returns a boolean if a field has been set.
func (o *KubernetesAddonRepository) HasRepositoryUrl() bool {
	if o != nil && o.RepositoryUrl != nil {
		return true
	}

	return false
}

// SetRepositoryUrl gets a reference to the given string and assigns it to the RepositoryUrl field.
func (o *KubernetesAddonRepository) SetRepositoryUrl(v string) {
	o.RepositoryUrl = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *KubernetesAddonRepository) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonRepository) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *KubernetesAddonRepository) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *KubernetesAddonRepository) SetUsername(v string) {
	o.Username = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *KubernetesAddonRepository) GetCatalog() WorkflowCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret WorkflowCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonRepository) GetCatalogOk() (*WorkflowCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *KubernetesAddonRepository) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given WorkflowCatalogRelationship and assigns it to the Catalog field.
func (o *KubernetesAddonRepository) SetCatalog(v WorkflowCatalogRelationship) {
	o.Catalog = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *KubernetesAddonRepository) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonRepository) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *KubernetesAddonRepository) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *KubernetesAddonRepository) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o KubernetesAddonRepository) MarshalJSON() ([]byte, error) {
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
	if o.CaCert.IsSet() {
		toSerialize["CaCert"] = o.CaCert.Get()
	}
	if o.InsecureSkipVerification != nil {
		toSerialize["InsecureSkipVerification"] = o.InsecureSkipVerification
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.IsTokenSet != nil {
		toSerialize["IsTokenSet"] = o.IsTokenSet
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.RepositoryUrl != nil {
		toSerialize["RepositoryUrl"] = o.RepositoryUrl
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesAddonRepository) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesAddonRepositoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		CaCert NullableX509Certificate `json:"CaCert,omitempty"`
		// Allow connecting to http registries or https registries which do not have certificate signed by a well known CA.
		InsecureSkipVerification *bool `json:"InsecureSkipVerification,omitempty"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// Indicates whether the value of the 'token' property has been set.
		IsTokenSet *bool `json:"IsTokenSet,omitempty"`
		// Name of the addon repository or registry.
		Name *string `json:"Name,omitempty"`
		// URL for the repository where the addon is hosted.
		RepositoryUrl *string `json:"RepositoryUrl,omitempty"`
		// Username to authenticate to the addon registry.
		Username *string `json:"Username,omitempty"`
		Catalog *WorkflowCatalogRelationship `json:"Catalog,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varKubernetesAddonRepositoryWithoutEmbeddedStruct := KubernetesAddonRepositoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesAddonRepositoryWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesAddonRepository := _KubernetesAddonRepository{}
		varKubernetesAddonRepository.ClassId = varKubernetesAddonRepositoryWithoutEmbeddedStruct.ClassId
		varKubernetesAddonRepository.ObjectType = varKubernetesAddonRepositoryWithoutEmbeddedStruct.ObjectType
		varKubernetesAddonRepository.CaCert = varKubernetesAddonRepositoryWithoutEmbeddedStruct.CaCert
		varKubernetesAddonRepository.InsecureSkipVerification = varKubernetesAddonRepositoryWithoutEmbeddedStruct.InsecureSkipVerification
		varKubernetesAddonRepository.IsPasswordSet = varKubernetesAddonRepositoryWithoutEmbeddedStruct.IsPasswordSet
		varKubernetesAddonRepository.IsTokenSet = varKubernetesAddonRepositoryWithoutEmbeddedStruct.IsTokenSet
		varKubernetesAddonRepository.Name = varKubernetesAddonRepositoryWithoutEmbeddedStruct.Name
		varKubernetesAddonRepository.RepositoryUrl = varKubernetesAddonRepositoryWithoutEmbeddedStruct.RepositoryUrl
		varKubernetesAddonRepository.Username = varKubernetesAddonRepositoryWithoutEmbeddedStruct.Username
		varKubernetesAddonRepository.Catalog = varKubernetesAddonRepositoryWithoutEmbeddedStruct.Catalog
		varKubernetesAddonRepository.RegisteredDevice = varKubernetesAddonRepositoryWithoutEmbeddedStruct.RegisteredDevice
		*o = KubernetesAddonRepository(varKubernetesAddonRepository)
	} else {
		return err
	}

	varKubernetesAddonRepository := _KubernetesAddonRepository{}

	err = json.Unmarshal(bytes, &varKubernetesAddonRepository)
	if err == nil {
		o.MoBaseMo = varKubernetesAddonRepository.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CaCert")
		delete(additionalProperties, "InsecureSkipVerification")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "IsTokenSet")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RepositoryUrl")
		delete(additionalProperties, "Username")
		delete(additionalProperties, "Catalog")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableKubernetesAddonRepository struct {
	value *KubernetesAddonRepository
	isSet bool
}

func (v NullableKubernetesAddonRepository) Get() *KubernetesAddonRepository {
	return v.value
}

func (v *NullableKubernetesAddonRepository) Set(val *KubernetesAddonRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAddonRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAddonRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAddonRepository(val *KubernetesAddonRepository) *NullableKubernetesAddonRepository {
	return &NullableKubernetesAddonRepository{value: val, isSet: true}
}

func (v NullableKubernetesAddonRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAddonRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

