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

// HyperflexClusterBackupPolicyInventory Response to Backup Policy requests and queries.
type HyperflexClusterBackupPolicyInventory struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Validate, Deploy, Prepare, Commit, or Abort the Backup Policy. Allowed values are \"VALIDATE\", \"DEPLOY\", \"PREPARE\", \"COMMIT\", \"ABORT\". * `VALIDATE` - Check for problems in policy request without deploying. * `DEPLOY` - Remove the policy.  Only allowed when cleanup field is true. * `PREPARE` - Prepare the policy for subsequent \"COMMIT\" or \"ABORT\".  Only allowed when cleanup field is false. * `COMMIT` - Commit the prepared policy.  Only allowed when cleanup field is false. * `ABORT` - Abort the prepared policy.  Only allowed when cleanup field is false.
	Action *string `json:"Action,omitempty"`
	// If true, remove the policy. Otherwise proceed with the specified policy action.
	Cleanup *bool `json:"Cleanup,omitempty"`
	// Indicates if the HyperFlex Cluster is the source cluster or the target cluster. When set to true, the HyperFlex Cluster is the source for VM backups. When set to false, the HyperFlex Cluster is the target cluster where VM backups are saved.
	IsSource *bool `json:"IsSource,omitempty"`
	// Details from associated HyperFlex job execution.
	JobDetails *string `json:"JobDetails,omitempty"`
	// Job Exception message from associated HyperFlex job execution.
	JobExceptionMessage *string `json:"JobExceptionMessage,omitempty"`
	// Job ID of associated HyperFlex job.
	JobId *string `json:"JobId,omitempty"`
	// State of the associated HyperFlex job. When present possible values are \"RUNNING\", \"COMPLETED\" or \"EXCEPTION\". * `RUNNING` - HyperFlex policy job is currently \"RUNNING\". * `COMPLETED` - HyperFlex policy job completed successfully. * `EXCEPTION` - HyperFlex policy job failed.
	JobState *string `json:"JobState,omitempty"`
	// Intersight HyperFlex Cluster Backup Policy MOID.
	PolicyMoid *string `json:"PolicyMoid,omitempty"`
	// Unique request ID allowing retry of the same logical request following a transient communication failure.
	RequestId *string `json:"RequestId,omitempty"`
	Settings NullableHyperflexBackupPolicySettings `json:"Settings,omitempty"`
	// UUID of the source HyperFlex Cluster.
	SourceUuid *string `json:"SourceUuid,omitempty"`
	// UUID of the target HyperFlex Cluster.
	TargetUuid *string `json:"TargetUuid,omitempty"`
	// Version of the Backup Policy.
	Version *int64 `json:"Version,omitempty"`
	SrcCluster *HyperflexClusterRelationship `json:"SrcCluster,omitempty"`
	TgtCluster *HyperflexClusterRelationship `json:"TgtCluster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexClusterBackupPolicyInventory HyperflexClusterBackupPolicyInventory

// NewHyperflexClusterBackupPolicyInventory instantiates a new HyperflexClusterBackupPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexClusterBackupPolicyInventory(classId string, objectType string) *HyperflexClusterBackupPolicyInventory {
	this := HyperflexClusterBackupPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexClusterBackupPolicyInventoryWithDefaults instantiates a new HyperflexClusterBackupPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterBackupPolicyInventoryWithDefaults() *HyperflexClusterBackupPolicyInventory {
	this := HyperflexClusterBackupPolicyInventory{}
	var classId string = "hyperflex.ClusterBackupPolicyInventory"
	this.ClassId = classId
	var objectType string = "hyperflex.ClusterBackupPolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexClusterBackupPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexClusterBackupPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexClusterBackupPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexClusterBackupPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyInventory) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyInventory) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *HyperflexClusterBackupPolicyInventory) SetAction(v string) {
	o.Action = &v
}

// GetCleanup returns the Cleanup field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyInventory) GetCleanup() bool {
	if o == nil || o.Cleanup == nil {
		var ret bool
		return ret
	}
	return *o.Cleanup
}

// GetCleanupOk returns a tuple with the Cleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetCleanupOk() (*bool, bool) {
	if o == nil || o.Cleanup == nil {
		return nil, false
	}
	return o.Cleanup, true
}

// HasCleanup returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyInventory) HasCleanup() bool {
	if o != nil && o.Cleanup != nil {
		return true
	}

	return false
}

// SetCleanup gets a reference to the given bool and assigns it to the Cleanup field.
func (o *HyperflexClusterBackupPolicyInventory) SetCleanup(v bool) {
	o.Cleanup = &v
}

// GetIsSource returns the IsSource field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyInventory) GetIsSource() bool {
	if o == nil || o.IsSource == nil {
		var ret bool
		return ret
	}
	return *o.IsSource
}

// GetIsSourceOk returns a tuple with the IsSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetIsSourceOk() (*bool, bool) {
	if o == nil || o.IsSource == nil {
		return nil, false
	}
	return o.IsSource, true
}

// HasIsSource returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyInventory) HasIsSource() bool {
	if o != nil && o.IsSource != nil {
		return true
	}

	return false
}

// SetIsSource gets a reference to the given bool and assigns it to the IsSource field.
func (o *HyperflexClusterBackupPolicyInventory) SetIsSource(v bool) {
	o.IsSource = &v
}

// GetJobDetails returns the JobDetails field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyInventory) GetJobDetails() string {
	if o == nil || o.JobDetails == nil {
		var ret string
		return ret
	}
	return *o.JobDetails
}

// GetJobDetailsOk returns a tuple with the JobDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetJobDetailsOk() (*string, bool) {
	if o == nil || o.JobDetails == nil {
		return nil, false
	}
	return o.JobDetails, true
}

// HasJobDetails returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyInventory) HasJobDetails() bool {
	if o != nil && o.JobDetails != nil {
		return true
	}

	return false
}

// SetJobDetails gets a reference to the given string and assigns it to the JobDetails field.
func (o *HyperflexClusterBackupPolicyInventory) SetJobDetails(v string) {
	o.JobDetails = &v
}

// GetJobExceptionMessage returns the JobExceptionMessage field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyInventory) GetJobExceptionMessage() string {
	if o == nil || o.JobExceptionMessage == nil {
		var ret string
		return ret
	}
	return *o.JobExceptionMessage
}

// GetJobExceptionMessageOk returns a tuple with the JobExceptionMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetJobExceptionMessageOk() (*string, bool) {
	if o == nil || o.JobExceptionMessage == nil {
		return nil, false
	}
	return o.JobExceptionMessage, true
}

// HasJobExceptionMessage returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyInventory) HasJobExceptionMessage() bool {
	if o != nil && o.JobExceptionMessage != nil {
		return true
	}

	return false
}

// SetJobExceptionMessage gets a reference to the given string and assigns it to the JobExceptionMessage field.
func (o *HyperflexClusterBackupPolicyInventory) SetJobExceptionMessage(v string) {
	o.JobExceptionMessage = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyInventory) GetJobId() string {
	if o == nil || o.JobId == nil {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetJobIdOk() (*string, bool) {
	if o == nil || o.JobId == nil {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyInventory) HasJobId() bool {
	if o != nil && o.JobId != nil {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *HyperflexClusterBackupPolicyInventory) SetJobId(v string) {
	o.JobId = &v
}

// GetJobState returns the JobState field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyInventory) GetJobState() string {
	if o == nil || o.JobState == nil {
		var ret string
		return ret
	}
	return *o.JobState
}

// GetJobStateOk returns a tuple with the JobState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetJobStateOk() (*string, bool) {
	if o == nil || o.JobState == nil {
		return nil, false
	}
	return o.JobState, true
}

// HasJobState returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyInventory) HasJobState() bool {
	if o != nil && o.JobState != nil {
		return true
	}

	return false
}

// SetJobState gets a reference to the given string and assigns it to the JobState field.
func (o *HyperflexClusterBackupPolicyInventory) SetJobState(v string) {
	o.JobState = &v
}

// GetPolicyMoid returns the PolicyMoid field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyInventory) GetPolicyMoid() string {
	if o == nil || o.PolicyMoid == nil {
		var ret string
		return ret
	}
	return *o.PolicyMoid
}

// GetPolicyMoidOk returns a tuple with the PolicyMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetPolicyMoidOk() (*string, bool) {
	if o == nil || o.PolicyMoid == nil {
		return nil, false
	}
	return o.PolicyMoid, true
}

// HasPolicyMoid returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyInventory) HasPolicyMoid() bool {
	if o != nil && o.PolicyMoid != nil {
		return true
	}

	return false
}

// SetPolicyMoid gets a reference to the given string and assigns it to the PolicyMoid field.
func (o *HyperflexClusterBackupPolicyInventory) SetPolicyMoid(v string) {
	o.PolicyMoid = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyInventory) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyInventory) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *HyperflexClusterBackupPolicyInventory) SetRequestId(v string) {
	o.RequestId = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterBackupPolicyInventory) GetSettings() HyperflexBackupPolicySettings {
	if o == nil || o.Settings.Get() == nil {
		var ret HyperflexBackupPolicySettings
		return ret
	}
	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterBackupPolicyInventory) GetSettingsOk() (*HyperflexBackupPolicySettings, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// HasSettings returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyInventory) HasSettings() bool {
	if o != nil && o.Settings.IsSet() {
		return true
	}

	return false
}

// SetSettings gets a reference to the given NullableHyperflexBackupPolicySettings and assigns it to the Settings field.
func (o *HyperflexClusterBackupPolicyInventory) SetSettings(v HyperflexBackupPolicySettings) {
	o.Settings.Set(&v)
}
// SetSettingsNil sets the value for Settings to be an explicit nil
func (o *HyperflexClusterBackupPolicyInventory) SetSettingsNil() {
	o.Settings.Set(nil)
}

// UnsetSettings ensures that no value is present for Settings, not even an explicit nil
func (o *HyperflexClusterBackupPolicyInventory) UnsetSettings() {
	o.Settings.Unset()
}

// GetSourceUuid returns the SourceUuid field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyInventory) GetSourceUuid() string {
	if o == nil || o.SourceUuid == nil {
		var ret string
		return ret
	}
	return *o.SourceUuid
}

// GetSourceUuidOk returns a tuple with the SourceUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetSourceUuidOk() (*string, bool) {
	if o == nil || o.SourceUuid == nil {
		return nil, false
	}
	return o.SourceUuid, true
}

// HasSourceUuid returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyInventory) HasSourceUuid() bool {
	if o != nil && o.SourceUuid != nil {
		return true
	}

	return false
}

// SetSourceUuid gets a reference to the given string and assigns it to the SourceUuid field.
func (o *HyperflexClusterBackupPolicyInventory) SetSourceUuid(v string) {
	o.SourceUuid = &v
}

// GetTargetUuid returns the TargetUuid field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyInventory) GetTargetUuid() string {
	if o == nil || o.TargetUuid == nil {
		var ret string
		return ret
	}
	return *o.TargetUuid
}

// GetTargetUuidOk returns a tuple with the TargetUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetTargetUuidOk() (*string, bool) {
	if o == nil || o.TargetUuid == nil {
		return nil, false
	}
	return o.TargetUuid, true
}

// HasTargetUuid returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyInventory) HasTargetUuid() bool {
	if o != nil && o.TargetUuid != nil {
		return true
	}

	return false
}

// SetTargetUuid gets a reference to the given string and assigns it to the TargetUuid field.
func (o *HyperflexClusterBackupPolicyInventory) SetTargetUuid(v string) {
	o.TargetUuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyInventory) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyInventory) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *HyperflexClusterBackupPolicyInventory) SetVersion(v int64) {
	o.Version = &v
}

// GetSrcCluster returns the SrcCluster field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyInventory) GetSrcCluster() HyperflexClusterRelationship {
	if o == nil || o.SrcCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.SrcCluster
}

// GetSrcClusterOk returns a tuple with the SrcCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetSrcClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.SrcCluster == nil {
		return nil, false
	}
	return o.SrcCluster, true
}

// HasSrcCluster returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyInventory) HasSrcCluster() bool {
	if o != nil && o.SrcCluster != nil {
		return true
	}

	return false
}

// SetSrcCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the SrcCluster field.
func (o *HyperflexClusterBackupPolicyInventory) SetSrcCluster(v HyperflexClusterRelationship) {
	o.SrcCluster = &v
}

// GetTgtCluster returns the TgtCluster field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyInventory) GetTgtCluster() HyperflexClusterRelationship {
	if o == nil || o.TgtCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.TgtCluster
}

// GetTgtClusterOk returns a tuple with the TgtCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyInventory) GetTgtClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.TgtCluster == nil {
		return nil, false
	}
	return o.TgtCluster, true
}

// HasTgtCluster returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyInventory) HasTgtCluster() bool {
	if o != nil && o.TgtCluster != nil {
		return true
	}

	return false
}

// SetTgtCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the TgtCluster field.
func (o *HyperflexClusterBackupPolicyInventory) SetTgtCluster(v HyperflexClusterRelationship) {
	o.TgtCluster = &v
}

func (o HyperflexClusterBackupPolicyInventory) MarshalJSON() ([]byte, error) {
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
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.Cleanup != nil {
		toSerialize["Cleanup"] = o.Cleanup
	}
	if o.IsSource != nil {
		toSerialize["IsSource"] = o.IsSource
	}
	if o.JobDetails != nil {
		toSerialize["JobDetails"] = o.JobDetails
	}
	if o.JobExceptionMessage != nil {
		toSerialize["JobExceptionMessage"] = o.JobExceptionMessage
	}
	if o.JobId != nil {
		toSerialize["JobId"] = o.JobId
	}
	if o.JobState != nil {
		toSerialize["JobState"] = o.JobState
	}
	if o.PolicyMoid != nil {
		toSerialize["PolicyMoid"] = o.PolicyMoid
	}
	if o.RequestId != nil {
		toSerialize["RequestId"] = o.RequestId
	}
	if o.Settings.IsSet() {
		toSerialize["Settings"] = o.Settings.Get()
	}
	if o.SourceUuid != nil {
		toSerialize["SourceUuid"] = o.SourceUuid
	}
	if o.TargetUuid != nil {
		toSerialize["TargetUuid"] = o.TargetUuid
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.SrcCluster != nil {
		toSerialize["SrcCluster"] = o.SrcCluster
	}
	if o.TgtCluster != nil {
		toSerialize["TgtCluster"] = o.TgtCluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexClusterBackupPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Validate, Deploy, Prepare, Commit, or Abort the Backup Policy. Allowed values are \"VALIDATE\", \"DEPLOY\", \"PREPARE\", \"COMMIT\", \"ABORT\". * `VALIDATE` - Check for problems in policy request without deploying. * `DEPLOY` - Remove the policy.  Only allowed when cleanup field is true. * `PREPARE` - Prepare the policy for subsequent \"COMMIT\" or \"ABORT\".  Only allowed when cleanup field is false. * `COMMIT` - Commit the prepared policy.  Only allowed when cleanup field is false. * `ABORT` - Abort the prepared policy.  Only allowed when cleanup field is false.
		Action *string `json:"Action,omitempty"`
		// If true, remove the policy. Otherwise proceed with the specified policy action.
		Cleanup *bool `json:"Cleanup,omitempty"`
		// Indicates if the HyperFlex Cluster is the source cluster or the target cluster. When set to true, the HyperFlex Cluster is the source for VM backups. When set to false, the HyperFlex Cluster is the target cluster where VM backups are saved.
		IsSource *bool `json:"IsSource,omitempty"`
		// Details from associated HyperFlex job execution.
		JobDetails *string `json:"JobDetails,omitempty"`
		// Job Exception message from associated HyperFlex job execution.
		JobExceptionMessage *string `json:"JobExceptionMessage,omitempty"`
		// Job ID of associated HyperFlex job.
		JobId *string `json:"JobId,omitempty"`
		// State of the associated HyperFlex job. When present possible values are \"RUNNING\", \"COMPLETED\" or \"EXCEPTION\". * `RUNNING` - HyperFlex policy job is currently \"RUNNING\". * `COMPLETED` - HyperFlex policy job completed successfully. * `EXCEPTION` - HyperFlex policy job failed.
		JobState *string `json:"JobState,omitempty"`
		// Intersight HyperFlex Cluster Backup Policy MOID.
		PolicyMoid *string `json:"PolicyMoid,omitempty"`
		// Unique request ID allowing retry of the same logical request following a transient communication failure.
		RequestId *string `json:"RequestId,omitempty"`
		Settings NullableHyperflexBackupPolicySettings `json:"Settings,omitempty"`
		// UUID of the source HyperFlex Cluster.
		SourceUuid *string `json:"SourceUuid,omitempty"`
		// UUID of the target HyperFlex Cluster.
		TargetUuid *string `json:"TargetUuid,omitempty"`
		// Version of the Backup Policy.
		Version *int64 `json:"Version,omitempty"`
		SrcCluster *HyperflexClusterRelationship `json:"SrcCluster,omitempty"`
		TgtCluster *HyperflexClusterRelationship `json:"TgtCluster,omitempty"`
	}

	varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct := HyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexClusterBackupPolicyInventory := _HyperflexClusterBackupPolicyInventory{}
		varHyperflexClusterBackupPolicyInventory.ClassId = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.ClassId
		varHyperflexClusterBackupPolicyInventory.ObjectType = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varHyperflexClusterBackupPolicyInventory.Action = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.Action
		varHyperflexClusterBackupPolicyInventory.Cleanup = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.Cleanup
		varHyperflexClusterBackupPolicyInventory.IsSource = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.IsSource
		varHyperflexClusterBackupPolicyInventory.JobDetails = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.JobDetails
		varHyperflexClusterBackupPolicyInventory.JobExceptionMessage = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.JobExceptionMessage
		varHyperflexClusterBackupPolicyInventory.JobId = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.JobId
		varHyperflexClusterBackupPolicyInventory.JobState = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.JobState
		varHyperflexClusterBackupPolicyInventory.PolicyMoid = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.PolicyMoid
		varHyperflexClusterBackupPolicyInventory.RequestId = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.RequestId
		varHyperflexClusterBackupPolicyInventory.Settings = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.Settings
		varHyperflexClusterBackupPolicyInventory.SourceUuid = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.SourceUuid
		varHyperflexClusterBackupPolicyInventory.TargetUuid = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.TargetUuid
		varHyperflexClusterBackupPolicyInventory.Version = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.Version
		varHyperflexClusterBackupPolicyInventory.SrcCluster = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.SrcCluster
		varHyperflexClusterBackupPolicyInventory.TgtCluster = varHyperflexClusterBackupPolicyInventoryWithoutEmbeddedStruct.TgtCluster
		*o = HyperflexClusterBackupPolicyInventory(varHyperflexClusterBackupPolicyInventory)
	} else {
		return err
	}

	varHyperflexClusterBackupPolicyInventory := _HyperflexClusterBackupPolicyInventory{}

	err = json.Unmarshal(bytes, &varHyperflexClusterBackupPolicyInventory)
	if err == nil {
		o.MoBaseMo = varHyperflexClusterBackupPolicyInventory.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "Cleanup")
		delete(additionalProperties, "IsSource")
		delete(additionalProperties, "JobDetails")
		delete(additionalProperties, "JobExceptionMessage")
		delete(additionalProperties, "JobId")
		delete(additionalProperties, "JobState")
		delete(additionalProperties, "PolicyMoid")
		delete(additionalProperties, "RequestId")
		delete(additionalProperties, "Settings")
		delete(additionalProperties, "SourceUuid")
		delete(additionalProperties, "TargetUuid")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "SrcCluster")
		delete(additionalProperties, "TgtCluster")

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

type NullableHyperflexClusterBackupPolicyInventory struct {
	value *HyperflexClusterBackupPolicyInventory
	isSet bool
}

func (v NullableHyperflexClusterBackupPolicyInventory) Get() *HyperflexClusterBackupPolicyInventory {
	return v.value
}

func (v *NullableHyperflexClusterBackupPolicyInventory) Set(val *HyperflexClusterBackupPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterBackupPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterBackupPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterBackupPolicyInventory(val *HyperflexClusterBackupPolicyInventory) *NullableHyperflexClusterBackupPolicyInventory {
	return &NullableHyperflexClusterBackupPolicyInventory{value: val, isSet: true}
}

func (v NullableHyperflexClusterBackupPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterBackupPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


