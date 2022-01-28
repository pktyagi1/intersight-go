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

// KubernetesNodeInfoAllOf Definition of the list of properties defined in 'kubernetes.NodeInfo', excluding properties defined in parent classes.
type KubernetesNodeInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Node Operating System architecture (amd64, arm64).
	Architecture *string `json:"Architecture,omitempty"`
	// A Boot ID is an Identifier generated by the host everytimes it gets reboot.
	BootId *string `json:"BootId,omitempty"`
	// To run containers in Pods, Kubernetes uses a container runtime. This field describes Container Runtime Version.
	ContainerRuntimeVersion *string `json:"ContainerRuntimeVersion,omitempty"`
	// Node Operating System kernel version.
	KernelVersion *string `json:"KernelVersion,omitempty"`
	// The Kubernetes network proxy runs on each node. This reflects services as defined in the Kubernetes API on each node and can do simple TCP, UDP, and SCTP stream forwarding or round robin TCP, UDP, and SCTP forwarding across a set of backends. This field describes the kube-proxy version.
	KubeProxyVersion *string `json:"KubeProxyVersion,omitempty"`
	// The kubelet is the primary \"node agent\" that runs on each node. It can register the node with the apiserver using one of such as the hostname; a flag to override the hostname; or specific logic for a cloud provider. This field describes the kubelet version the node currently using.
	KubeletVersion *string `json:"KubeletVersion,omitempty"`
	// It is a node identifier in Kubernetes. When the node joins a kubernetes cluster, Kubernetes will assign a machine ID to that node. Learn more from man machine-id from http://man7.org/linux/man-pages/man5/machine-id.5.html.
	MachineId *string `json:"MachineId,omitempty"`
	// Operating System installed on this Node.
	OperatingSystem *string `json:"OperatingSystem,omitempty"`
	// Node current Operating System image.
	OsImage *string `json:"OsImage,omitempty"`
	// SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts https://access.redhat.com/documentation/en-US/Red_Hat_Subscription_Management/1/html/RHSM/getting-system-uuid.html.
	SystemUuid *string `json:"SystemUuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNodeInfoAllOf KubernetesNodeInfoAllOf

// NewKubernetesNodeInfoAllOf instantiates a new KubernetesNodeInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeInfoAllOf(classId string, objectType string) *KubernetesNodeInfoAllOf {
	this := KubernetesNodeInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesNodeInfoAllOfWithDefaults instantiates a new KubernetesNodeInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeInfoAllOfWithDefaults() *KubernetesNodeInfoAllOf {
	this := KubernetesNodeInfoAllOf{}
	var classId string = "kubernetes.NodeInfo"
	this.ClassId = classId
	var objectType string = "kubernetes.NodeInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNodeInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNodeInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNodeInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNodeInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *KubernetesNodeInfoAllOf) GetArchitecture() string {
	if o == nil || o.Architecture == nil {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfoAllOf) GetArchitectureOk() (*string, bool) {
	if o == nil || o.Architecture == nil {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *KubernetesNodeInfoAllOf) HasArchitecture() bool {
	if o != nil && o.Architecture != nil {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *KubernetesNodeInfoAllOf) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetBootId returns the BootId field value if set, zero value otherwise.
func (o *KubernetesNodeInfoAllOf) GetBootId() string {
	if o == nil || o.BootId == nil {
		var ret string
		return ret
	}
	return *o.BootId
}

// GetBootIdOk returns a tuple with the BootId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfoAllOf) GetBootIdOk() (*string, bool) {
	if o == nil || o.BootId == nil {
		return nil, false
	}
	return o.BootId, true
}

// HasBootId returns a boolean if a field has been set.
func (o *KubernetesNodeInfoAllOf) HasBootId() bool {
	if o != nil && o.BootId != nil {
		return true
	}

	return false
}

// SetBootId gets a reference to the given string and assigns it to the BootId field.
func (o *KubernetesNodeInfoAllOf) SetBootId(v string) {
	o.BootId = &v
}

// GetContainerRuntimeVersion returns the ContainerRuntimeVersion field value if set, zero value otherwise.
func (o *KubernetesNodeInfoAllOf) GetContainerRuntimeVersion() string {
	if o == nil || o.ContainerRuntimeVersion == nil {
		var ret string
		return ret
	}
	return *o.ContainerRuntimeVersion
}

// GetContainerRuntimeVersionOk returns a tuple with the ContainerRuntimeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfoAllOf) GetContainerRuntimeVersionOk() (*string, bool) {
	if o == nil || o.ContainerRuntimeVersion == nil {
		return nil, false
	}
	return o.ContainerRuntimeVersion, true
}

// HasContainerRuntimeVersion returns a boolean if a field has been set.
func (o *KubernetesNodeInfoAllOf) HasContainerRuntimeVersion() bool {
	if o != nil && o.ContainerRuntimeVersion != nil {
		return true
	}

	return false
}

// SetContainerRuntimeVersion gets a reference to the given string and assigns it to the ContainerRuntimeVersion field.
func (o *KubernetesNodeInfoAllOf) SetContainerRuntimeVersion(v string) {
	o.ContainerRuntimeVersion = &v
}

// GetKernelVersion returns the KernelVersion field value if set, zero value otherwise.
func (o *KubernetesNodeInfoAllOf) GetKernelVersion() string {
	if o == nil || o.KernelVersion == nil {
		var ret string
		return ret
	}
	return *o.KernelVersion
}

// GetKernelVersionOk returns a tuple with the KernelVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfoAllOf) GetKernelVersionOk() (*string, bool) {
	if o == nil || o.KernelVersion == nil {
		return nil, false
	}
	return o.KernelVersion, true
}

// HasKernelVersion returns a boolean if a field has been set.
func (o *KubernetesNodeInfoAllOf) HasKernelVersion() bool {
	if o != nil && o.KernelVersion != nil {
		return true
	}

	return false
}

// SetKernelVersion gets a reference to the given string and assigns it to the KernelVersion field.
func (o *KubernetesNodeInfoAllOf) SetKernelVersion(v string) {
	o.KernelVersion = &v
}

// GetKubeProxyVersion returns the KubeProxyVersion field value if set, zero value otherwise.
func (o *KubernetesNodeInfoAllOf) GetKubeProxyVersion() string {
	if o == nil || o.KubeProxyVersion == nil {
		var ret string
		return ret
	}
	return *o.KubeProxyVersion
}

// GetKubeProxyVersionOk returns a tuple with the KubeProxyVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfoAllOf) GetKubeProxyVersionOk() (*string, bool) {
	if o == nil || o.KubeProxyVersion == nil {
		return nil, false
	}
	return o.KubeProxyVersion, true
}

// HasKubeProxyVersion returns a boolean if a field has been set.
func (o *KubernetesNodeInfoAllOf) HasKubeProxyVersion() bool {
	if o != nil && o.KubeProxyVersion != nil {
		return true
	}

	return false
}

// SetKubeProxyVersion gets a reference to the given string and assigns it to the KubeProxyVersion field.
func (o *KubernetesNodeInfoAllOf) SetKubeProxyVersion(v string) {
	o.KubeProxyVersion = &v
}

// GetKubeletVersion returns the KubeletVersion field value if set, zero value otherwise.
func (o *KubernetesNodeInfoAllOf) GetKubeletVersion() string {
	if o == nil || o.KubeletVersion == nil {
		var ret string
		return ret
	}
	return *o.KubeletVersion
}

// GetKubeletVersionOk returns a tuple with the KubeletVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfoAllOf) GetKubeletVersionOk() (*string, bool) {
	if o == nil || o.KubeletVersion == nil {
		return nil, false
	}
	return o.KubeletVersion, true
}

// HasKubeletVersion returns a boolean if a field has been set.
func (o *KubernetesNodeInfoAllOf) HasKubeletVersion() bool {
	if o != nil && o.KubeletVersion != nil {
		return true
	}

	return false
}

// SetKubeletVersion gets a reference to the given string and assigns it to the KubeletVersion field.
func (o *KubernetesNodeInfoAllOf) SetKubeletVersion(v string) {
	o.KubeletVersion = &v
}

// GetMachineId returns the MachineId field value if set, zero value otherwise.
func (o *KubernetesNodeInfoAllOf) GetMachineId() string {
	if o == nil || o.MachineId == nil {
		var ret string
		return ret
	}
	return *o.MachineId
}

// GetMachineIdOk returns a tuple with the MachineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfoAllOf) GetMachineIdOk() (*string, bool) {
	if o == nil || o.MachineId == nil {
		return nil, false
	}
	return o.MachineId, true
}

// HasMachineId returns a boolean if a field has been set.
func (o *KubernetesNodeInfoAllOf) HasMachineId() bool {
	if o != nil && o.MachineId != nil {
		return true
	}

	return false
}

// SetMachineId gets a reference to the given string and assigns it to the MachineId field.
func (o *KubernetesNodeInfoAllOf) SetMachineId(v string) {
	o.MachineId = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *KubernetesNodeInfoAllOf) GetOperatingSystem() string {
	if o == nil || o.OperatingSystem == nil {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfoAllOf) GetOperatingSystemOk() (*string, bool) {
	if o == nil || o.OperatingSystem == nil {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *KubernetesNodeInfoAllOf) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem != nil {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *KubernetesNodeInfoAllOf) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

// GetOsImage returns the OsImage field value if set, zero value otherwise.
func (o *KubernetesNodeInfoAllOf) GetOsImage() string {
	if o == nil || o.OsImage == nil {
		var ret string
		return ret
	}
	return *o.OsImage
}

// GetOsImageOk returns a tuple with the OsImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfoAllOf) GetOsImageOk() (*string, bool) {
	if o == nil || o.OsImage == nil {
		return nil, false
	}
	return o.OsImage, true
}

// HasOsImage returns a boolean if a field has been set.
func (o *KubernetesNodeInfoAllOf) HasOsImage() bool {
	if o != nil && o.OsImage != nil {
		return true
	}

	return false
}

// SetOsImage gets a reference to the given string and assigns it to the OsImage field.
func (o *KubernetesNodeInfoAllOf) SetOsImage(v string) {
	o.OsImage = &v
}

// GetSystemUuid returns the SystemUuid field value if set, zero value otherwise.
func (o *KubernetesNodeInfoAllOf) GetSystemUuid() string {
	if o == nil || o.SystemUuid == nil {
		var ret string
		return ret
	}
	return *o.SystemUuid
}

// GetSystemUuidOk returns a tuple with the SystemUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeInfoAllOf) GetSystemUuidOk() (*string, bool) {
	if o == nil || o.SystemUuid == nil {
		return nil, false
	}
	return o.SystemUuid, true
}

// HasSystemUuid returns a boolean if a field has been set.
func (o *KubernetesNodeInfoAllOf) HasSystemUuid() bool {
	if o != nil && o.SystemUuid != nil {
		return true
	}

	return false
}

// SetSystemUuid gets a reference to the given string and assigns it to the SystemUuid field.
func (o *KubernetesNodeInfoAllOf) SetSystemUuid(v string) {
	o.SystemUuid = &v
}

func (o KubernetesNodeInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Architecture != nil {
		toSerialize["Architecture"] = o.Architecture
	}
	if o.BootId != nil {
		toSerialize["BootId"] = o.BootId
	}
	if o.ContainerRuntimeVersion != nil {
		toSerialize["ContainerRuntimeVersion"] = o.ContainerRuntimeVersion
	}
	if o.KernelVersion != nil {
		toSerialize["KernelVersion"] = o.KernelVersion
	}
	if o.KubeProxyVersion != nil {
		toSerialize["KubeProxyVersion"] = o.KubeProxyVersion
	}
	if o.KubeletVersion != nil {
		toSerialize["KubeletVersion"] = o.KubeletVersion
	}
	if o.MachineId != nil {
		toSerialize["MachineId"] = o.MachineId
	}
	if o.OperatingSystem != nil {
		toSerialize["OperatingSystem"] = o.OperatingSystem
	}
	if o.OsImage != nil {
		toSerialize["OsImage"] = o.OsImage
	}
	if o.SystemUuid != nil {
		toSerialize["SystemUuid"] = o.SystemUuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesNodeInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesNodeInfoAllOf := _KubernetesNodeInfoAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesNodeInfoAllOf); err == nil {
		*o = KubernetesNodeInfoAllOf(varKubernetesNodeInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Architecture")
		delete(additionalProperties, "BootId")
		delete(additionalProperties, "ContainerRuntimeVersion")
		delete(additionalProperties, "KernelVersion")
		delete(additionalProperties, "KubeProxyVersion")
		delete(additionalProperties, "KubeletVersion")
		delete(additionalProperties, "MachineId")
		delete(additionalProperties, "OperatingSystem")
		delete(additionalProperties, "OsImage")
		delete(additionalProperties, "SystemUuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesNodeInfoAllOf struct {
	value *KubernetesNodeInfoAllOf
	isSet bool
}

func (v NullableKubernetesNodeInfoAllOf) Get() *KubernetesNodeInfoAllOf {
	return v.value
}

func (v *NullableKubernetesNodeInfoAllOf) Set(val *KubernetesNodeInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeInfoAllOf(val *KubernetesNodeInfoAllOf) *NullableKubernetesNodeInfoAllOf {
	return &NullableKubernetesNodeInfoAllOf{value: val, isSet: true}
}

func (v NullableKubernetesNodeInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


