/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5517
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// TerraformExecutor Executor is a ManagedObject, aka MO. This API is used to execute terraform scripts.
type TerraformExecutor struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	CloudResource []TerraformCloudResource `json:"CloudResource,omitempty"`
	// Command to be executed during update operation.
	Command *string `json:"Command,omitempty"`
	// Enum indicates what operation is being done. * `Create` - Creating a Terraform resource. * `Update` - Updating a Terraform resource. * `Delete` - Deleting a Terraform resource.
	Operation *string `json:"Operation,omitempty"`
	// Terraform output of the entire execution.
	Output interface{} `json:"Output,omitempty"`
	// The Platform type used in conjunction with 'sourceFolderPath' and 'sourceFolderName' determines unique path for a Terraform workflow.
	PlatformType *string `json:"PlatformType,omitempty"`
	RunState []TerraformRunstate `json:"RunState,omitempty"`
	// Folder Name where Terraform workflows are stored.
	SourceFolderName *string `json:"SourceFolderName,omitempty"`
	// Relative folder Path where 'sourceFolderName' is located.
	SourceFolderPath *string `json:"SourceFolderPath,omitempty"`
	// Flag indicates whether workflow is internal/external.
	SourceLocation *string `json:"SourceLocation,omitempty"`
	// Status of the terraform execution.
	Status *string `json:"Status,omitempty"`
	// Stderr of the terraform execution will be captured here.
	Stderr interface{} `json:"Stderr,omitempty"`
	// Stdout of the terraform execution will be captured here.
	Stdout interface{} `json:"Stdout,omitempty"`
	// TaskId of a pontem workflow is same as the MO.
	TaskId *string `json:"TaskId,omitempty"`
	// Variables needed by the terraform configuration as a JSON object.
	Variables interface{} `json:"Variables,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	WorkflowInfo *WorkflowWorkflowInfoRelationship `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TerraformExecutor TerraformExecutor

// NewTerraformExecutor instantiates a new TerraformExecutor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerraformExecutor(classId string, objectType string) *TerraformExecutor {
	this := TerraformExecutor{}
	this.ClassId = classId
	this.ObjectType = objectType
	var operation string = "Create"
	this.Operation = &operation
	return &this
}

// NewTerraformExecutorWithDefaults instantiates a new TerraformExecutor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerraformExecutorWithDefaults() *TerraformExecutor {
	this := TerraformExecutor{}
	var classId string = "terraform.Executor"
	this.ClassId = classId
	var objectType string = "terraform.Executor"
	this.ObjectType = objectType
	var operation string = "Create"
	this.Operation = &operation
	return &this
}

// GetClassId returns the ClassId field value
func (o *TerraformExecutor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TerraformExecutor) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TerraformExecutor) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TerraformExecutor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TerraformExecutor) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TerraformExecutor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCloudResource returns the CloudResource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TerraformExecutor) GetCloudResource() []TerraformCloudResource {
	if o == nil  {
		var ret []TerraformCloudResource
		return ret
	}
	return o.CloudResource
}

// GetCloudResourceOk returns a tuple with the CloudResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerraformExecutor) GetCloudResourceOk() (*[]TerraformCloudResource, bool) {
	if o == nil || o.CloudResource == nil {
		return nil, false
	}
	return &o.CloudResource, true
}

// HasCloudResource returns a boolean if a field has been set.
func (o *TerraformExecutor) HasCloudResource() bool {
	if o != nil && o.CloudResource != nil {
		return true
	}

	return false
}

// SetCloudResource gets a reference to the given []TerraformCloudResource and assigns it to the CloudResource field.
func (o *TerraformExecutor) SetCloudResource(v []TerraformCloudResource) {
	o.CloudResource = v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *TerraformExecutor) GetCommand() string {
	if o == nil || o.Command == nil {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformExecutor) GetCommandOk() (*string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *TerraformExecutor) HasCommand() bool {
	if o != nil && o.Command != nil {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *TerraformExecutor) SetCommand(v string) {
	o.Command = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *TerraformExecutor) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformExecutor) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *TerraformExecutor) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *TerraformExecutor) SetOperation(v string) {
	o.Operation = &v
}

// GetOutput returns the Output field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TerraformExecutor) GetOutput() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerraformExecutor) GetOutputOk() (*interface{}, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return &o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *TerraformExecutor) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given interface{} and assigns it to the Output field.
func (o *TerraformExecutor) SetOutput(v interface{}) {
	o.Output = v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *TerraformExecutor) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformExecutor) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *TerraformExecutor) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *TerraformExecutor) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetRunState returns the RunState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TerraformExecutor) GetRunState() []TerraformRunstate {
	if o == nil  {
		var ret []TerraformRunstate
		return ret
	}
	return o.RunState
}

// GetRunStateOk returns a tuple with the RunState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerraformExecutor) GetRunStateOk() (*[]TerraformRunstate, bool) {
	if o == nil || o.RunState == nil {
		return nil, false
	}
	return &o.RunState, true
}

// HasRunState returns a boolean if a field has been set.
func (o *TerraformExecutor) HasRunState() bool {
	if o != nil && o.RunState != nil {
		return true
	}

	return false
}

// SetRunState gets a reference to the given []TerraformRunstate and assigns it to the RunState field.
func (o *TerraformExecutor) SetRunState(v []TerraformRunstate) {
	o.RunState = v
}

// GetSourceFolderName returns the SourceFolderName field value if set, zero value otherwise.
func (o *TerraformExecutor) GetSourceFolderName() string {
	if o == nil || o.SourceFolderName == nil {
		var ret string
		return ret
	}
	return *o.SourceFolderName
}

// GetSourceFolderNameOk returns a tuple with the SourceFolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformExecutor) GetSourceFolderNameOk() (*string, bool) {
	if o == nil || o.SourceFolderName == nil {
		return nil, false
	}
	return o.SourceFolderName, true
}

// HasSourceFolderName returns a boolean if a field has been set.
func (o *TerraformExecutor) HasSourceFolderName() bool {
	if o != nil && o.SourceFolderName != nil {
		return true
	}

	return false
}

// SetSourceFolderName gets a reference to the given string and assigns it to the SourceFolderName field.
func (o *TerraformExecutor) SetSourceFolderName(v string) {
	o.SourceFolderName = &v
}

// GetSourceFolderPath returns the SourceFolderPath field value if set, zero value otherwise.
func (o *TerraformExecutor) GetSourceFolderPath() string {
	if o == nil || o.SourceFolderPath == nil {
		var ret string
		return ret
	}
	return *o.SourceFolderPath
}

// GetSourceFolderPathOk returns a tuple with the SourceFolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformExecutor) GetSourceFolderPathOk() (*string, bool) {
	if o == nil || o.SourceFolderPath == nil {
		return nil, false
	}
	return o.SourceFolderPath, true
}

// HasSourceFolderPath returns a boolean if a field has been set.
func (o *TerraformExecutor) HasSourceFolderPath() bool {
	if o != nil && o.SourceFolderPath != nil {
		return true
	}

	return false
}

// SetSourceFolderPath gets a reference to the given string and assigns it to the SourceFolderPath field.
func (o *TerraformExecutor) SetSourceFolderPath(v string) {
	o.SourceFolderPath = &v
}

// GetSourceLocation returns the SourceLocation field value if set, zero value otherwise.
func (o *TerraformExecutor) GetSourceLocation() string {
	if o == nil || o.SourceLocation == nil {
		var ret string
		return ret
	}
	return *o.SourceLocation
}

// GetSourceLocationOk returns a tuple with the SourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformExecutor) GetSourceLocationOk() (*string, bool) {
	if o == nil || o.SourceLocation == nil {
		return nil, false
	}
	return o.SourceLocation, true
}

// HasSourceLocation returns a boolean if a field has been set.
func (o *TerraformExecutor) HasSourceLocation() bool {
	if o != nil && o.SourceLocation != nil {
		return true
	}

	return false
}

// SetSourceLocation gets a reference to the given string and assigns it to the SourceLocation field.
func (o *TerraformExecutor) SetSourceLocation(v string) {
	o.SourceLocation = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TerraformExecutor) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformExecutor) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TerraformExecutor) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TerraformExecutor) SetStatus(v string) {
	o.Status = &v
}

// GetStderr returns the Stderr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TerraformExecutor) GetStderr() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Stderr
}

// GetStderrOk returns a tuple with the Stderr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerraformExecutor) GetStderrOk() (*interface{}, bool) {
	if o == nil || o.Stderr == nil {
		return nil, false
	}
	return &o.Stderr, true
}

// HasStderr returns a boolean if a field has been set.
func (o *TerraformExecutor) HasStderr() bool {
	if o != nil && o.Stderr != nil {
		return true
	}

	return false
}

// SetStderr gets a reference to the given interface{} and assigns it to the Stderr field.
func (o *TerraformExecutor) SetStderr(v interface{}) {
	o.Stderr = v
}

// GetStdout returns the Stdout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TerraformExecutor) GetStdout() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Stdout
}

// GetStdoutOk returns a tuple with the Stdout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerraformExecutor) GetStdoutOk() (*interface{}, bool) {
	if o == nil || o.Stdout == nil {
		return nil, false
	}
	return &o.Stdout, true
}

// HasStdout returns a boolean if a field has been set.
func (o *TerraformExecutor) HasStdout() bool {
	if o != nil && o.Stdout != nil {
		return true
	}

	return false
}

// SetStdout gets a reference to the given interface{} and assigns it to the Stdout field.
func (o *TerraformExecutor) SetStdout(v interface{}) {
	o.Stdout = v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *TerraformExecutor) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformExecutor) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *TerraformExecutor) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *TerraformExecutor) SetTaskId(v string) {
	o.TaskId = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TerraformExecutor) GetVariables() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerraformExecutor) GetVariablesOk() (*interface{}, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return &o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *TerraformExecutor) HasVariables() bool {
	if o != nil && o.Variables != nil {
		return true
	}

	return false
}

// SetVariables gets a reference to the given interface{} and assigns it to the Variables field.
func (o *TerraformExecutor) SetVariables(v interface{}) {
	o.Variables = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *TerraformExecutor) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformExecutor) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *TerraformExecutor) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *TerraformExecutor) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *TerraformExecutor) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformExecutor) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *TerraformExecutor) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *TerraformExecutor) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *TerraformExecutor) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformExecutor) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *TerraformExecutor) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *TerraformExecutor) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o TerraformExecutor) MarshalJSON() ([]byte, error) {
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
	if o.CloudResource != nil {
		toSerialize["CloudResource"] = o.CloudResource
	}
	if o.Command != nil {
		toSerialize["Command"] = o.Command
	}
	if o.Operation != nil {
		toSerialize["Operation"] = o.Operation
	}
	if o.Output != nil {
		toSerialize["Output"] = o.Output
	}
	if o.PlatformType != nil {
		toSerialize["PlatformType"] = o.PlatformType
	}
	if o.RunState != nil {
		toSerialize["RunState"] = o.RunState
	}
	if o.SourceFolderName != nil {
		toSerialize["SourceFolderName"] = o.SourceFolderName
	}
	if o.SourceFolderPath != nil {
		toSerialize["SourceFolderPath"] = o.SourceFolderPath
	}
	if o.SourceLocation != nil {
		toSerialize["SourceLocation"] = o.SourceLocation
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Stderr != nil {
		toSerialize["Stderr"] = o.Stderr
	}
	if o.Stdout != nil {
		toSerialize["Stdout"] = o.Stdout
	}
	if o.TaskId != nil {
		toSerialize["TaskId"] = o.TaskId
	}
	if o.Variables != nil {
		toSerialize["Variables"] = o.Variables
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TerraformExecutor) UnmarshalJSON(bytes []byte) (err error) {
	type TerraformExecutorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		CloudResource []TerraformCloudResource `json:"CloudResource,omitempty"`
		// Command to be executed during update operation.
		Command *string `json:"Command,omitempty"`
		// Enum indicates what operation is being done. * `Create` - Creating a Terraform resource. * `Update` - Updating a Terraform resource. * `Delete` - Deleting a Terraform resource.
		Operation *string `json:"Operation,omitempty"`
		// Terraform output of the entire execution.
		Output interface{} `json:"Output,omitempty"`
		// The Platform type used in conjunction with 'sourceFolderPath' and 'sourceFolderName' determines unique path for a Terraform workflow.
		PlatformType *string `json:"PlatformType,omitempty"`
		RunState []TerraformRunstate `json:"RunState,omitempty"`
		// Folder Name where Terraform workflows are stored.
		SourceFolderName *string `json:"SourceFolderName,omitempty"`
		// Relative folder Path where 'sourceFolderName' is located.
		SourceFolderPath *string `json:"SourceFolderPath,omitempty"`
		// Flag indicates whether workflow is internal/external.
		SourceLocation *string `json:"SourceLocation,omitempty"`
		// Status of the terraform execution.
		Status *string `json:"Status,omitempty"`
		// Stderr of the terraform execution will be captured here.
		Stderr interface{} `json:"Stderr,omitempty"`
		// Stdout of the terraform execution will be captured here.
		Stdout interface{} `json:"Stdout,omitempty"`
		// TaskId of a pontem workflow is same as the MO.
		TaskId *string `json:"TaskId,omitempty"`
		// Variables needed by the terraform configuration as a JSON object.
		Variables interface{} `json:"Variables,omitempty"`
		Account *IamAccountRelationship `json:"Account,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		WorkflowInfo *WorkflowWorkflowInfoRelationship `json:"WorkflowInfo,omitempty"`
	}

	varTerraformExecutorWithoutEmbeddedStruct := TerraformExecutorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTerraformExecutorWithoutEmbeddedStruct)
	if err == nil {
		varTerraformExecutor := _TerraformExecutor{}
		varTerraformExecutor.ClassId = varTerraformExecutorWithoutEmbeddedStruct.ClassId
		varTerraformExecutor.ObjectType = varTerraformExecutorWithoutEmbeddedStruct.ObjectType
		varTerraformExecutor.CloudResource = varTerraformExecutorWithoutEmbeddedStruct.CloudResource
		varTerraformExecutor.Command = varTerraformExecutorWithoutEmbeddedStruct.Command
		varTerraformExecutor.Operation = varTerraformExecutorWithoutEmbeddedStruct.Operation
		varTerraformExecutor.Output = varTerraformExecutorWithoutEmbeddedStruct.Output
		varTerraformExecutor.PlatformType = varTerraformExecutorWithoutEmbeddedStruct.PlatformType
		varTerraformExecutor.RunState = varTerraformExecutorWithoutEmbeddedStruct.RunState
		varTerraformExecutor.SourceFolderName = varTerraformExecutorWithoutEmbeddedStruct.SourceFolderName
		varTerraformExecutor.SourceFolderPath = varTerraformExecutorWithoutEmbeddedStruct.SourceFolderPath
		varTerraformExecutor.SourceLocation = varTerraformExecutorWithoutEmbeddedStruct.SourceLocation
		varTerraformExecutor.Status = varTerraformExecutorWithoutEmbeddedStruct.Status
		varTerraformExecutor.Stderr = varTerraformExecutorWithoutEmbeddedStruct.Stderr
		varTerraformExecutor.Stdout = varTerraformExecutorWithoutEmbeddedStruct.Stdout
		varTerraformExecutor.TaskId = varTerraformExecutorWithoutEmbeddedStruct.TaskId
		varTerraformExecutor.Variables = varTerraformExecutorWithoutEmbeddedStruct.Variables
		varTerraformExecutor.Account = varTerraformExecutorWithoutEmbeddedStruct.Account
		varTerraformExecutor.RegisteredDevice = varTerraformExecutorWithoutEmbeddedStruct.RegisteredDevice
		varTerraformExecutor.WorkflowInfo = varTerraformExecutorWithoutEmbeddedStruct.WorkflowInfo
		*o = TerraformExecutor(varTerraformExecutor)
	} else {
		return err
	}

	varTerraformExecutor := _TerraformExecutor{}

	err = json.Unmarshal(bytes, &varTerraformExecutor)
	if err == nil {
		o.MoBaseMo = varTerraformExecutor.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CloudResource")
		delete(additionalProperties, "Command")
		delete(additionalProperties, "Operation")
		delete(additionalProperties, "Output")
		delete(additionalProperties, "PlatformType")
		delete(additionalProperties, "RunState")
		delete(additionalProperties, "SourceFolderName")
		delete(additionalProperties, "SourceFolderPath")
		delete(additionalProperties, "SourceLocation")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Stderr")
		delete(additionalProperties, "Stdout")
		delete(additionalProperties, "TaskId")
		delete(additionalProperties, "Variables")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "WorkflowInfo")

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

type NullableTerraformExecutor struct {
	value *TerraformExecutor
	isSet bool
}

func (v NullableTerraformExecutor) Get() *TerraformExecutor {
	return v.value
}

func (v *NullableTerraformExecutor) Set(val *TerraformExecutor) {
	v.value = val
	v.isSet = true
}

func (v NullableTerraformExecutor) IsSet() bool {
	return v.isSet
}

func (v *NullableTerraformExecutor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerraformExecutor(val *TerraformExecutor) *NullableTerraformExecutor {
	return &NullableTerraformExecutor{value: val, isSet: true}
}

func (v NullableTerraformExecutor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerraformExecutor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


