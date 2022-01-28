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
	"reflect"
	"strings"
)

// KubernetesPodStatus The current status of the Pod.
type KubernetesPodStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The IP of the host that the pod is running on.
	HostIp *string `json:"HostIp,omitempty"`
	// The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle.
	Phase *string `json:"Phase,omitempty"`
	// The IP of the pod. The IP is allocated by the Pod CIDR from the kubernetes cluster itself.
	PodIp *string `json:"PodIp,omitempty"`
	// The Quality of Service (QOS) classification assigned to the pod based on resource requirements.
	QosClass *string `json:"QosClass,omitempty"`
	// The time that the pod was started.
	StartTime *string `json:"StartTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesPodStatus KubernetesPodStatus

// NewKubernetesPodStatus instantiates a new KubernetesPodStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesPodStatus(classId string, objectType string) *KubernetesPodStatus {
	this := KubernetesPodStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesPodStatusWithDefaults instantiates a new KubernetesPodStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesPodStatusWithDefaults() *KubernetesPodStatus {
	this := KubernetesPodStatus{}
	var classId string = "kubernetes.PodStatus"
	this.ClassId = classId
	var objectType string = "kubernetes.PodStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesPodStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesPodStatus) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesPodStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesPodStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesPodStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesPodStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHostIp returns the HostIp field value if set, zero value otherwise.
func (o *KubernetesPodStatus) GetHostIp() string {
	if o == nil || o.HostIp == nil {
		var ret string
		return ret
	}
	return *o.HostIp
}

// GetHostIpOk returns a tuple with the HostIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesPodStatus) GetHostIpOk() (*string, bool) {
	if o == nil || o.HostIp == nil {
		return nil, false
	}
	return o.HostIp, true
}

// HasHostIp returns a boolean if a field has been set.
func (o *KubernetesPodStatus) HasHostIp() bool {
	if o != nil && o.HostIp != nil {
		return true
	}

	return false
}

// SetHostIp gets a reference to the given string and assigns it to the HostIp field.
func (o *KubernetesPodStatus) SetHostIp(v string) {
	o.HostIp = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *KubernetesPodStatus) GetPhase() string {
	if o == nil || o.Phase == nil {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesPodStatus) GetPhaseOk() (*string, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *KubernetesPodStatus) HasPhase() bool {
	if o != nil && o.Phase != nil {
		return true
	}

	return false
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *KubernetesPodStatus) SetPhase(v string) {
	o.Phase = &v
}

// GetPodIp returns the PodIp field value if set, zero value otherwise.
func (o *KubernetesPodStatus) GetPodIp() string {
	if o == nil || o.PodIp == nil {
		var ret string
		return ret
	}
	return *o.PodIp
}

// GetPodIpOk returns a tuple with the PodIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesPodStatus) GetPodIpOk() (*string, bool) {
	if o == nil || o.PodIp == nil {
		return nil, false
	}
	return o.PodIp, true
}

// HasPodIp returns a boolean if a field has been set.
func (o *KubernetesPodStatus) HasPodIp() bool {
	if o != nil && o.PodIp != nil {
		return true
	}

	return false
}

// SetPodIp gets a reference to the given string and assigns it to the PodIp field.
func (o *KubernetesPodStatus) SetPodIp(v string) {
	o.PodIp = &v
}

// GetQosClass returns the QosClass field value if set, zero value otherwise.
func (o *KubernetesPodStatus) GetQosClass() string {
	if o == nil || o.QosClass == nil {
		var ret string
		return ret
	}
	return *o.QosClass
}

// GetQosClassOk returns a tuple with the QosClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesPodStatus) GetQosClassOk() (*string, bool) {
	if o == nil || o.QosClass == nil {
		return nil, false
	}
	return o.QosClass, true
}

// HasQosClass returns a boolean if a field has been set.
func (o *KubernetesPodStatus) HasQosClass() bool {
	if o != nil && o.QosClass != nil {
		return true
	}

	return false
}

// SetQosClass gets a reference to the given string and assigns it to the QosClass field.
func (o *KubernetesPodStatus) SetQosClass(v string) {
	o.QosClass = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *KubernetesPodStatus) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesPodStatus) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *KubernetesPodStatus) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *KubernetesPodStatus) SetStartTime(v string) {
	o.StartTime = &v
}

func (o KubernetesPodStatus) MarshalJSON() ([]byte, error) {
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
	if o.HostIp != nil {
		toSerialize["HostIp"] = o.HostIp
	}
	if o.Phase != nil {
		toSerialize["Phase"] = o.Phase
	}
	if o.PodIp != nil {
		toSerialize["PodIp"] = o.PodIp
	}
	if o.QosClass != nil {
		toSerialize["QosClass"] = o.QosClass
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesPodStatus) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesPodStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The IP of the host that the pod is running on.
		HostIp *string `json:"HostIp,omitempty"`
		// The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle.
		Phase *string `json:"Phase,omitempty"`
		// The IP of the pod. The IP is allocated by the Pod CIDR from the kubernetes cluster itself.
		PodIp *string `json:"PodIp,omitempty"`
		// The Quality of Service (QOS) classification assigned to the pod based on resource requirements.
		QosClass *string `json:"QosClass,omitempty"`
		// The time that the pod was started.
		StartTime *string `json:"StartTime,omitempty"`
	}

	varKubernetesPodStatusWithoutEmbeddedStruct := KubernetesPodStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesPodStatusWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesPodStatus := _KubernetesPodStatus{}
		varKubernetesPodStatus.ClassId = varKubernetesPodStatusWithoutEmbeddedStruct.ClassId
		varKubernetesPodStatus.ObjectType = varKubernetesPodStatusWithoutEmbeddedStruct.ObjectType
		varKubernetesPodStatus.HostIp = varKubernetesPodStatusWithoutEmbeddedStruct.HostIp
		varKubernetesPodStatus.Phase = varKubernetesPodStatusWithoutEmbeddedStruct.Phase
		varKubernetesPodStatus.PodIp = varKubernetesPodStatusWithoutEmbeddedStruct.PodIp
		varKubernetesPodStatus.QosClass = varKubernetesPodStatusWithoutEmbeddedStruct.QosClass
		varKubernetesPodStatus.StartTime = varKubernetesPodStatusWithoutEmbeddedStruct.StartTime
		*o = KubernetesPodStatus(varKubernetesPodStatus)
	} else {
		return err
	}

	varKubernetesPodStatus := _KubernetesPodStatus{}

	err = json.Unmarshal(bytes, &varKubernetesPodStatus)
	if err == nil {
		o.MoBaseComplexType = varKubernetesPodStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HostIp")
		delete(additionalProperties, "Phase")
		delete(additionalProperties, "PodIp")
		delete(additionalProperties, "QosClass")
		delete(additionalProperties, "StartTime")

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

type NullableKubernetesPodStatus struct {
	value *KubernetesPodStatus
	isSet bool
}

func (v NullableKubernetesPodStatus) Get() *KubernetesPodStatus {
	return v.value
}

func (v *NullableKubernetesPodStatus) Set(val *KubernetesPodStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesPodStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesPodStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesPodStatus(val *KubernetesPodStatus) *NullableKubernetesPodStatus {
	return &NullableKubernetesPodStatus{value: val, isSet: true}
}

func (v NullableKubernetesPodStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesPodStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


