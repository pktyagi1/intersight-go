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

// FirmwareBaseImpact The abstract entity depicting impact on the endpoint during the upgrade operation.
type FirmwareBaseImpact struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Details for the error that occurred during the reboot validation analysis.
	ComputationError *string `json:"ComputationError,omitempty"`
	// The computation status of the estimate operation for a component. * `Inprogress` - Upgrade impact calculation is in progress. * `Completed` - Upgrade impact calculation is completed. * `Unavailable` - Upgrade impact is not available since the image is not present in the Fabric Interconnect. * `Failed` - Upgrade impact is not available due to an unknown error.
	ComputationStatusDetail *string `json:"ComputationStatusDetail,omitempty"`
	// The endpoint type or name.
	DomainName *string `json:"DomainName,omitempty"`
	// A reference to a REST resource, uniquely identified by object type and MOID.
	EndPoint *string `json:"EndPoint,omitempty"`
	// The current firmware version of the component.
	FirmwareVersion *string `json:"FirmwareVersion,omitempty"`
	// The impact type of the endpoint, whether a reboot is required or not. * `NoReboot` - A reboot is not required for the endpoint after upgrade. * `Reboot` - A reboot is required to the endpoint after upgrade.
	ImpactType *string `json:"ImpactType,omitempty"`
	// The model details of the component.
	Model *string `json:"Model,omitempty"`
	// The target firmware version of the component.
	TargetFirmwareVersion *string `json:"TargetFirmwareVersion,omitempty"`
	// The change of version impact for the endpoint. * `None` - No change in version for the component. * `Upgrade` - The component will be upgraded. * `Downgrade` - The component will be downgraded.
	VersionImpact *string `json:"VersionImpact,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareBaseImpact FirmwareBaseImpact

// NewFirmwareBaseImpact instantiates a new FirmwareBaseImpact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareBaseImpact(classId string, objectType string) *FirmwareBaseImpact {
	this := FirmwareBaseImpact{}
	this.ClassId = classId
	this.ObjectType = objectType
	var computationStatusDetail string = "Inprogress"
	this.ComputationStatusDetail = &computationStatusDetail
	var impactType string = "NoReboot"
	this.ImpactType = &impactType
	var versionImpact string = "None"
	this.VersionImpact = &versionImpact
	return &this
}

// NewFirmwareBaseImpactWithDefaults instantiates a new FirmwareBaseImpact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareBaseImpactWithDefaults() *FirmwareBaseImpact {
	this := FirmwareBaseImpact{}
	var computationStatusDetail string = "Inprogress"
	this.ComputationStatusDetail = &computationStatusDetail
	var impactType string = "NoReboot"
	this.ImpactType = &impactType
	var versionImpact string = "None"
	this.VersionImpact = &versionImpact
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareBaseImpact) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareBaseImpact) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareBaseImpact) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareBaseImpact) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareBaseImpact) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareBaseImpact) SetObjectType(v string) {
	o.ObjectType = v
}

// GetComputationError returns the ComputationError field value if set, zero value otherwise.
func (o *FirmwareBaseImpact) GetComputationError() string {
	if o == nil || o.ComputationError == nil {
		var ret string
		return ret
	}
	return *o.ComputationError
}

// GetComputationErrorOk returns a tuple with the ComputationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseImpact) GetComputationErrorOk() (*string, bool) {
	if o == nil || o.ComputationError == nil {
		return nil, false
	}
	return o.ComputationError, true
}

// HasComputationError returns a boolean if a field has been set.
func (o *FirmwareBaseImpact) HasComputationError() bool {
	if o != nil && o.ComputationError != nil {
		return true
	}

	return false
}

// SetComputationError gets a reference to the given string and assigns it to the ComputationError field.
func (o *FirmwareBaseImpact) SetComputationError(v string) {
	o.ComputationError = &v
}

// GetComputationStatusDetail returns the ComputationStatusDetail field value if set, zero value otherwise.
func (o *FirmwareBaseImpact) GetComputationStatusDetail() string {
	if o == nil || o.ComputationStatusDetail == nil {
		var ret string
		return ret
	}
	return *o.ComputationStatusDetail
}

// GetComputationStatusDetailOk returns a tuple with the ComputationStatusDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseImpact) GetComputationStatusDetailOk() (*string, bool) {
	if o == nil || o.ComputationStatusDetail == nil {
		return nil, false
	}
	return o.ComputationStatusDetail, true
}

// HasComputationStatusDetail returns a boolean if a field has been set.
func (o *FirmwareBaseImpact) HasComputationStatusDetail() bool {
	if o != nil && o.ComputationStatusDetail != nil {
		return true
	}

	return false
}

// SetComputationStatusDetail gets a reference to the given string and assigns it to the ComputationStatusDetail field.
func (o *FirmwareBaseImpact) SetComputationStatusDetail(v string) {
	o.ComputationStatusDetail = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *FirmwareBaseImpact) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseImpact) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *FirmwareBaseImpact) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *FirmwareBaseImpact) SetDomainName(v string) {
	o.DomainName = &v
}

// GetEndPoint returns the EndPoint field value if set, zero value otherwise.
func (o *FirmwareBaseImpact) GetEndPoint() string {
	if o == nil || o.EndPoint == nil {
		var ret string
		return ret
	}
	return *o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseImpact) GetEndPointOk() (*string, bool) {
	if o == nil || o.EndPoint == nil {
		return nil, false
	}
	return o.EndPoint, true
}

// HasEndPoint returns a boolean if a field has been set.
func (o *FirmwareBaseImpact) HasEndPoint() bool {
	if o != nil && o.EndPoint != nil {
		return true
	}

	return false
}

// SetEndPoint gets a reference to the given string and assigns it to the EndPoint field.
func (o *FirmwareBaseImpact) SetEndPoint(v string) {
	o.EndPoint = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *FirmwareBaseImpact) GetFirmwareVersion() string {
	if o == nil || o.FirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseImpact) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || o.FirmwareVersion == nil {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *FirmwareBaseImpact) HasFirmwareVersion() bool {
	if o != nil && o.FirmwareVersion != nil {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *FirmwareBaseImpact) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetImpactType returns the ImpactType field value if set, zero value otherwise.
func (o *FirmwareBaseImpact) GetImpactType() string {
	if o == nil || o.ImpactType == nil {
		var ret string
		return ret
	}
	return *o.ImpactType
}

// GetImpactTypeOk returns a tuple with the ImpactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseImpact) GetImpactTypeOk() (*string, bool) {
	if o == nil || o.ImpactType == nil {
		return nil, false
	}
	return o.ImpactType, true
}

// HasImpactType returns a boolean if a field has been set.
func (o *FirmwareBaseImpact) HasImpactType() bool {
	if o != nil && o.ImpactType != nil {
		return true
	}

	return false
}

// SetImpactType gets a reference to the given string and assigns it to the ImpactType field.
func (o *FirmwareBaseImpact) SetImpactType(v string) {
	o.ImpactType = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *FirmwareBaseImpact) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseImpact) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *FirmwareBaseImpact) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *FirmwareBaseImpact) SetModel(v string) {
	o.Model = &v
}

// GetTargetFirmwareVersion returns the TargetFirmwareVersion field value if set, zero value otherwise.
func (o *FirmwareBaseImpact) GetTargetFirmwareVersion() string {
	if o == nil || o.TargetFirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.TargetFirmwareVersion
}

// GetTargetFirmwareVersionOk returns a tuple with the TargetFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseImpact) GetTargetFirmwareVersionOk() (*string, bool) {
	if o == nil || o.TargetFirmwareVersion == nil {
		return nil, false
	}
	return o.TargetFirmwareVersion, true
}

// HasTargetFirmwareVersion returns a boolean if a field has been set.
func (o *FirmwareBaseImpact) HasTargetFirmwareVersion() bool {
	if o != nil && o.TargetFirmwareVersion != nil {
		return true
	}

	return false
}

// SetTargetFirmwareVersion gets a reference to the given string and assigns it to the TargetFirmwareVersion field.
func (o *FirmwareBaseImpact) SetTargetFirmwareVersion(v string) {
	o.TargetFirmwareVersion = &v
}

// GetVersionImpact returns the VersionImpact field value if set, zero value otherwise.
func (o *FirmwareBaseImpact) GetVersionImpact() string {
	if o == nil || o.VersionImpact == nil {
		var ret string
		return ret
	}
	return *o.VersionImpact
}

// GetVersionImpactOk returns a tuple with the VersionImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseImpact) GetVersionImpactOk() (*string, bool) {
	if o == nil || o.VersionImpact == nil {
		return nil, false
	}
	return o.VersionImpact, true
}

// HasVersionImpact returns a boolean if a field has been set.
func (o *FirmwareBaseImpact) HasVersionImpact() bool {
	if o != nil && o.VersionImpact != nil {
		return true
	}

	return false
}

// SetVersionImpact gets a reference to the given string and assigns it to the VersionImpact field.
func (o *FirmwareBaseImpact) SetVersionImpact(v string) {
	o.VersionImpact = &v
}

func (o FirmwareBaseImpact) MarshalJSON() ([]byte, error) {
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
	if o.ComputationError != nil {
		toSerialize["ComputationError"] = o.ComputationError
	}
	if o.ComputationStatusDetail != nil {
		toSerialize["ComputationStatusDetail"] = o.ComputationStatusDetail
	}
	if o.DomainName != nil {
		toSerialize["DomainName"] = o.DomainName
	}
	if o.EndPoint != nil {
		toSerialize["EndPoint"] = o.EndPoint
	}
	if o.FirmwareVersion != nil {
		toSerialize["FirmwareVersion"] = o.FirmwareVersion
	}
	if o.ImpactType != nil {
		toSerialize["ImpactType"] = o.ImpactType
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.TargetFirmwareVersion != nil {
		toSerialize["TargetFirmwareVersion"] = o.TargetFirmwareVersion
	}
	if o.VersionImpact != nil {
		toSerialize["VersionImpact"] = o.VersionImpact
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareBaseImpact) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareBaseImpactWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Details for the error that occurred during the reboot validation analysis.
		ComputationError *string `json:"ComputationError,omitempty"`
		// The computation status of the estimate operation for a component. * `Inprogress` - Upgrade impact calculation is in progress. * `Completed` - Upgrade impact calculation is completed. * `Unavailable` - Upgrade impact is not available since the image is not present in the Fabric Interconnect. * `Failed` - Upgrade impact is not available due to an unknown error.
		ComputationStatusDetail *string `json:"ComputationStatusDetail,omitempty"`
		// The endpoint type or name.
		DomainName *string `json:"DomainName,omitempty"`
		// A reference to a REST resource, uniquely identified by object type and MOID.
		EndPoint *string `json:"EndPoint,omitempty"`
		// The current firmware version of the component.
		FirmwareVersion *string `json:"FirmwareVersion,omitempty"`
		// The impact type of the endpoint, whether a reboot is required or not. * `NoReboot` - A reboot is not required for the endpoint after upgrade. * `Reboot` - A reboot is required to the endpoint after upgrade.
		ImpactType *string `json:"ImpactType,omitempty"`
		// The model details of the component.
		Model *string `json:"Model,omitempty"`
		// The target firmware version of the component.
		TargetFirmwareVersion *string `json:"TargetFirmwareVersion,omitempty"`
		// The change of version impact for the endpoint. * `None` - No change in version for the component. * `Upgrade` - The component will be upgraded. * `Downgrade` - The component will be downgraded.
		VersionImpact *string `json:"VersionImpact,omitempty"`
	}

	varFirmwareBaseImpactWithoutEmbeddedStruct := FirmwareBaseImpactWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareBaseImpactWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareBaseImpact := _FirmwareBaseImpact{}
		varFirmwareBaseImpact.ClassId = varFirmwareBaseImpactWithoutEmbeddedStruct.ClassId
		varFirmwareBaseImpact.ObjectType = varFirmwareBaseImpactWithoutEmbeddedStruct.ObjectType
		varFirmwareBaseImpact.ComputationError = varFirmwareBaseImpactWithoutEmbeddedStruct.ComputationError
		varFirmwareBaseImpact.ComputationStatusDetail = varFirmwareBaseImpactWithoutEmbeddedStruct.ComputationStatusDetail
		varFirmwareBaseImpact.DomainName = varFirmwareBaseImpactWithoutEmbeddedStruct.DomainName
		varFirmwareBaseImpact.EndPoint = varFirmwareBaseImpactWithoutEmbeddedStruct.EndPoint
		varFirmwareBaseImpact.FirmwareVersion = varFirmwareBaseImpactWithoutEmbeddedStruct.FirmwareVersion
		varFirmwareBaseImpact.ImpactType = varFirmwareBaseImpactWithoutEmbeddedStruct.ImpactType
		varFirmwareBaseImpact.Model = varFirmwareBaseImpactWithoutEmbeddedStruct.Model
		varFirmwareBaseImpact.TargetFirmwareVersion = varFirmwareBaseImpactWithoutEmbeddedStruct.TargetFirmwareVersion
		varFirmwareBaseImpact.VersionImpact = varFirmwareBaseImpactWithoutEmbeddedStruct.VersionImpact
		*o = FirmwareBaseImpact(varFirmwareBaseImpact)
	} else {
		return err
	}

	varFirmwareBaseImpact := _FirmwareBaseImpact{}

	err = json.Unmarshal(bytes, &varFirmwareBaseImpact)
	if err == nil {
		o.MoBaseComplexType = varFirmwareBaseImpact.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ComputationError")
		delete(additionalProperties, "ComputationStatusDetail")
		delete(additionalProperties, "DomainName")
		delete(additionalProperties, "EndPoint")
		delete(additionalProperties, "FirmwareVersion")
		delete(additionalProperties, "ImpactType")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "TargetFirmwareVersion")
		delete(additionalProperties, "VersionImpact")

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

type NullableFirmwareBaseImpact struct {
	value *FirmwareBaseImpact
	isSet bool
}

func (v NullableFirmwareBaseImpact) Get() *FirmwareBaseImpact {
	return v.value
}

func (v *NullableFirmwareBaseImpact) Set(val *FirmwareBaseImpact) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareBaseImpact) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareBaseImpact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareBaseImpact(val *FirmwareBaseImpact) *NullableFirmwareBaseImpact {
	return &NullableFirmwareBaseImpact{value: val, isSet: true}
}

func (v NullableFirmwareBaseImpact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareBaseImpact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


