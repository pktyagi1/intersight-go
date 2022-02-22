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
	"reflect"
	"strings"
)

// KubernetesNodeSpec Kubernetes Node Specification.
type KubernetesNodeSpec struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Node Pod CIDR. In Kubernetes, the workload (Pod) is allocated to an IP address by Kubernetes. The IP address is from a Pod CIDR of the cluster. Each node will (mostly) evenly be distributed the Pod CIDR.
	PodCidr *string `json:"PodCidr,omitempty"`
	// Kubernetes can be running on a specific cloud provider such as Openstack, Amazon Web Services, vCenter etc. Each cloud provider will have a specific cloud provider ID. This field is to represent that ID for the corresponding Kubernetes cluster.
	ProviderId *string `json:"ProviderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNodeSpec KubernetesNodeSpec

// NewKubernetesNodeSpec instantiates a new KubernetesNodeSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeSpec(classId string, objectType string) *KubernetesNodeSpec {
	this := KubernetesNodeSpec{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesNodeSpecWithDefaults instantiates a new KubernetesNodeSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeSpecWithDefaults() *KubernetesNodeSpec {
	this := KubernetesNodeSpec{}
	var classId string = "kubernetes.NodeSpec"
	this.ClassId = classId
	var objectType string = "kubernetes.NodeSpec"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNodeSpec) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeSpec) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNodeSpec) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNodeSpec) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeSpec) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNodeSpec) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPodCidr returns the PodCidr field value if set, zero value otherwise.
func (o *KubernetesNodeSpec) GetPodCidr() string {
	if o == nil || o.PodCidr == nil {
		var ret string
		return ret
	}
	return *o.PodCidr
}

// GetPodCidrOk returns a tuple with the PodCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeSpec) GetPodCidrOk() (*string, bool) {
	if o == nil || o.PodCidr == nil {
		return nil, false
	}
	return o.PodCidr, true
}

// HasPodCidr returns a boolean if a field has been set.
func (o *KubernetesNodeSpec) HasPodCidr() bool {
	if o != nil && o.PodCidr != nil {
		return true
	}

	return false
}

// SetPodCidr gets a reference to the given string and assigns it to the PodCidr field.
func (o *KubernetesNodeSpec) SetPodCidr(v string) {
	o.PodCidr = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *KubernetesNodeSpec) GetProviderId() string {
	if o == nil || o.ProviderId == nil {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeSpec) GetProviderIdOk() (*string, bool) {
	if o == nil || o.ProviderId == nil {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *KubernetesNodeSpec) HasProviderId() bool {
	if o != nil && o.ProviderId != nil {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *KubernetesNodeSpec) SetProviderId(v string) {
	o.ProviderId = &v
}

func (o KubernetesNodeSpec) MarshalJSON() ([]byte, error) {
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
	if o.PodCidr != nil {
		toSerialize["PodCidr"] = o.PodCidr
	}
	if o.ProviderId != nil {
		toSerialize["ProviderId"] = o.ProviderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesNodeSpec) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesNodeSpecWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Node Pod CIDR. In Kubernetes, the workload (Pod) is allocated to an IP address by Kubernetes. The IP address is from a Pod CIDR of the cluster. Each node will (mostly) evenly be distributed the Pod CIDR.
		PodCidr *string `json:"PodCidr,omitempty"`
		// Kubernetes can be running on a specific cloud provider such as Openstack, Amazon Web Services, vCenter etc. Each cloud provider will have a specific cloud provider ID. This field is to represent that ID for the corresponding Kubernetes cluster.
		ProviderId *string `json:"ProviderId,omitempty"`
	}

	varKubernetesNodeSpecWithoutEmbeddedStruct := KubernetesNodeSpecWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesNodeSpecWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesNodeSpec := _KubernetesNodeSpec{}
		varKubernetesNodeSpec.ClassId = varKubernetesNodeSpecWithoutEmbeddedStruct.ClassId
		varKubernetesNodeSpec.ObjectType = varKubernetesNodeSpecWithoutEmbeddedStruct.ObjectType
		varKubernetesNodeSpec.PodCidr = varKubernetesNodeSpecWithoutEmbeddedStruct.PodCidr
		varKubernetesNodeSpec.ProviderId = varKubernetesNodeSpecWithoutEmbeddedStruct.ProviderId
		*o = KubernetesNodeSpec(varKubernetesNodeSpec)
	} else {
		return err
	}

	varKubernetesNodeSpec := _KubernetesNodeSpec{}

	err = json.Unmarshal(bytes, &varKubernetesNodeSpec)
	if err == nil {
		o.MoBaseComplexType = varKubernetesNodeSpec.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PodCidr")
		delete(additionalProperties, "ProviderId")

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

type NullableKubernetesNodeSpec struct {
	value *KubernetesNodeSpec
	isSet bool
}

func (v NullableKubernetesNodeSpec) Get() *KubernetesNodeSpec {
	return v.value
}

func (v *NullableKubernetesNodeSpec) Set(val *KubernetesNodeSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeSpec(val *KubernetesNodeSpec) *NullableKubernetesNodeSpec {
	return &NullableKubernetesNodeSpec{value: val, isSet: true}
}

func (v NullableKubernetesNodeSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


