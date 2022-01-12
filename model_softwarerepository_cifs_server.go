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

// SoftwarerepositoryCifsServer An external file repository accessible through the CIFS protocol.
type SoftwarerepositoryCifsServer struct {
	SoftwarerepositoryFileServer
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile.
	FileLocation *string `json:"FileLocation,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// For CIFS, leave the field blank or enter one or more comma seperated options from the following. For Example, \" \" , \" soft \" , \" soft , nounix \" . * soft. * nounix. * noserviceino. * guest. * USERNAME=VALUE. * PASSWORD=VALUE. * sec=VALUE (VALUE could be None, Ntlm, Ntlmi, Ntlmssp, Ntlmsspi, Ntlmv2, Ntlmv2i).
	MountOption *string `json:"MountOption,omitempty"`
	// Password configured on the CIFS server.
	Password *string `json:"Password,omitempty"`
	// Filename of the image in the CIFS server. For example:ucs-c220m5-huu-3.1.2c.iso.
	RemoteFile *string `json:"RemoteFile,omitempty"`
	// Hostname or IP Address of the CIFS server.
	RemoteIp *string `json:"RemoteIp,omitempty"`
	// Remote directory where the image is present. For example:/share/subfolder.
	RemoteShare *string `json:"RemoteShare,omitempty"`
	// Username configured on the CIFS server.
	Username *string `json:"Username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryCifsServer SoftwarerepositoryCifsServer

// NewSoftwarerepositoryCifsServer instantiates a new SoftwarerepositoryCifsServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryCifsServer(classId string, objectType string) *SoftwarerepositoryCifsServer {
	this := SoftwarerepositoryCifsServer{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwarerepositoryCifsServerWithDefaults instantiates a new SoftwarerepositoryCifsServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryCifsServerWithDefaults() *SoftwarerepositoryCifsServer {
	this := SoftwarerepositoryCifsServer{}
	var classId string = "softwarerepository.CifsServer"
	this.ClassId = classId
	var objectType string = "softwarerepository.CifsServer"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryCifsServer) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServer) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryCifsServer) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryCifsServer) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServer) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryCifsServer) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileLocation returns the FileLocation field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServer) GetFileLocation() string {
	if o == nil || o.FileLocation == nil {
		var ret string
		return ret
	}
	return *o.FileLocation
}

// GetFileLocationOk returns a tuple with the FileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServer) GetFileLocationOk() (*string, bool) {
	if o == nil || o.FileLocation == nil {
		return nil, false
	}
	return o.FileLocation, true
}

// HasFileLocation returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServer) HasFileLocation() bool {
	if o != nil && o.FileLocation != nil {
		return true
	}

	return false
}

// SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.
func (o *SoftwarerepositoryCifsServer) SetFileLocation(v string) {
	o.FileLocation = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServer) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServer) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServer) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *SoftwarerepositoryCifsServer) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetMountOption returns the MountOption field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServer) GetMountOption() string {
	if o == nil || o.MountOption == nil {
		var ret string
		return ret
	}
	return *o.MountOption
}

// GetMountOptionOk returns a tuple with the MountOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServer) GetMountOptionOk() (*string, bool) {
	if o == nil || o.MountOption == nil {
		return nil, false
	}
	return o.MountOption, true
}

// HasMountOption returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServer) HasMountOption() bool {
	if o != nil && o.MountOption != nil {
		return true
	}

	return false
}

// SetMountOption gets a reference to the given string and assigns it to the MountOption field.
func (o *SoftwarerepositoryCifsServer) SetMountOption(v string) {
	o.MountOption = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServer) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServer) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServer) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SoftwarerepositoryCifsServer) SetPassword(v string) {
	o.Password = &v
}

// GetRemoteFile returns the RemoteFile field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServer) GetRemoteFile() string {
	if o == nil || o.RemoteFile == nil {
		var ret string
		return ret
	}
	return *o.RemoteFile
}

// GetRemoteFileOk returns a tuple with the RemoteFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServer) GetRemoteFileOk() (*string, bool) {
	if o == nil || o.RemoteFile == nil {
		return nil, false
	}
	return o.RemoteFile, true
}

// HasRemoteFile returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServer) HasRemoteFile() bool {
	if o != nil && o.RemoteFile != nil {
		return true
	}

	return false
}

// SetRemoteFile gets a reference to the given string and assigns it to the RemoteFile field.
func (o *SoftwarerepositoryCifsServer) SetRemoteFile(v string) {
	o.RemoteFile = &v
}

// GetRemoteIp returns the RemoteIp field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServer) GetRemoteIp() string {
	if o == nil || o.RemoteIp == nil {
		var ret string
		return ret
	}
	return *o.RemoteIp
}

// GetRemoteIpOk returns a tuple with the RemoteIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServer) GetRemoteIpOk() (*string, bool) {
	if o == nil || o.RemoteIp == nil {
		return nil, false
	}
	return o.RemoteIp, true
}

// HasRemoteIp returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServer) HasRemoteIp() bool {
	if o != nil && o.RemoteIp != nil {
		return true
	}

	return false
}

// SetRemoteIp gets a reference to the given string and assigns it to the RemoteIp field.
func (o *SoftwarerepositoryCifsServer) SetRemoteIp(v string) {
	o.RemoteIp = &v
}

// GetRemoteShare returns the RemoteShare field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServer) GetRemoteShare() string {
	if o == nil || o.RemoteShare == nil {
		var ret string
		return ret
	}
	return *o.RemoteShare
}

// GetRemoteShareOk returns a tuple with the RemoteShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServer) GetRemoteShareOk() (*string, bool) {
	if o == nil || o.RemoteShare == nil {
		return nil, false
	}
	return o.RemoteShare, true
}

// HasRemoteShare returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServer) HasRemoteShare() bool {
	if o != nil && o.RemoteShare != nil {
		return true
	}

	return false
}

// SetRemoteShare gets a reference to the given string and assigns it to the RemoteShare field.
func (o *SoftwarerepositoryCifsServer) SetRemoteShare(v string) {
	o.RemoteShare = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SoftwarerepositoryCifsServer) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCifsServer) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SoftwarerepositoryCifsServer) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SoftwarerepositoryCifsServer) SetUsername(v string) {
	o.Username = &v
}

func (o SoftwarerepositoryCifsServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSoftwarerepositoryFileServer, errSoftwarerepositoryFileServer := json.Marshal(o.SoftwarerepositoryFileServer)
	if errSoftwarerepositoryFileServer != nil {
		return []byte{}, errSoftwarerepositoryFileServer
	}
	errSoftwarerepositoryFileServer = json.Unmarshal([]byte(serializedSoftwarerepositoryFileServer), &toSerialize)
	if errSoftwarerepositoryFileServer != nil {
		return []byte{}, errSoftwarerepositoryFileServer
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FileLocation != nil {
		toSerialize["FileLocation"] = o.FileLocation
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.MountOption != nil {
		toSerialize["MountOption"] = o.MountOption
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.RemoteFile != nil {
		toSerialize["RemoteFile"] = o.RemoteFile
	}
	if o.RemoteIp != nil {
		toSerialize["RemoteIp"] = o.RemoteIp
	}
	if o.RemoteShare != nil {
		toSerialize["RemoteShare"] = o.RemoteShare
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryCifsServer) UnmarshalJSON(bytes []byte) (err error) {
	type SoftwarerepositoryCifsServerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile.
		FileLocation *string `json:"FileLocation,omitempty"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// For CIFS, leave the field blank or enter one or more comma seperated options from the following. For Example, \" \" , \" soft \" , \" soft , nounix \" . * soft. * nounix. * noserviceino. * guest. * USERNAME=VALUE. * PASSWORD=VALUE. * sec=VALUE (VALUE could be None, Ntlm, Ntlmi, Ntlmssp, Ntlmsspi, Ntlmv2, Ntlmv2i).
		MountOption *string `json:"MountOption,omitempty"`
		// Password configured on the CIFS server.
		Password *string `json:"Password,omitempty"`
		// Filename of the image in the CIFS server. For example:ucs-c220m5-huu-3.1.2c.iso.
		RemoteFile *string `json:"RemoteFile,omitempty"`
		// Hostname or IP Address of the CIFS server.
		RemoteIp *string `json:"RemoteIp,omitempty"`
		// Remote directory where the image is present. For example:/share/subfolder.
		RemoteShare *string `json:"RemoteShare,omitempty"`
		// Username configured on the CIFS server.
		Username *string `json:"Username,omitempty"`
	}

	varSoftwarerepositoryCifsServerWithoutEmbeddedStruct := SoftwarerepositoryCifsServerWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryCifsServerWithoutEmbeddedStruct)
	if err == nil {
		varSoftwarerepositoryCifsServer := _SoftwarerepositoryCifsServer{}
		varSoftwarerepositoryCifsServer.ClassId = varSoftwarerepositoryCifsServerWithoutEmbeddedStruct.ClassId
		varSoftwarerepositoryCifsServer.ObjectType = varSoftwarerepositoryCifsServerWithoutEmbeddedStruct.ObjectType
		varSoftwarerepositoryCifsServer.FileLocation = varSoftwarerepositoryCifsServerWithoutEmbeddedStruct.FileLocation
		varSoftwarerepositoryCifsServer.IsPasswordSet = varSoftwarerepositoryCifsServerWithoutEmbeddedStruct.IsPasswordSet
		varSoftwarerepositoryCifsServer.MountOption = varSoftwarerepositoryCifsServerWithoutEmbeddedStruct.MountOption
		varSoftwarerepositoryCifsServer.Password = varSoftwarerepositoryCifsServerWithoutEmbeddedStruct.Password
		varSoftwarerepositoryCifsServer.RemoteFile = varSoftwarerepositoryCifsServerWithoutEmbeddedStruct.RemoteFile
		varSoftwarerepositoryCifsServer.RemoteIp = varSoftwarerepositoryCifsServerWithoutEmbeddedStruct.RemoteIp
		varSoftwarerepositoryCifsServer.RemoteShare = varSoftwarerepositoryCifsServerWithoutEmbeddedStruct.RemoteShare
		varSoftwarerepositoryCifsServer.Username = varSoftwarerepositoryCifsServerWithoutEmbeddedStruct.Username
		*o = SoftwarerepositoryCifsServer(varSoftwarerepositoryCifsServer)
	} else {
		return err
	}

	varSoftwarerepositoryCifsServer := _SoftwarerepositoryCifsServer{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryCifsServer)
	if err == nil {
		o.SoftwarerepositoryFileServer = varSoftwarerepositoryCifsServer.SoftwarerepositoryFileServer
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileLocation")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "MountOption")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "RemoteFile")
		delete(additionalProperties, "RemoteIp")
		delete(additionalProperties, "RemoteShare")
		delete(additionalProperties, "Username")

		// remove fields from embedded structs
		reflectSoftwarerepositoryFileServer := reflect.ValueOf(o.SoftwarerepositoryFileServer)
		for i := 0; i < reflectSoftwarerepositoryFileServer.Type().NumField(); i++ {
			t := reflectSoftwarerepositoryFileServer.Type().Field(i)

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

type NullableSoftwarerepositoryCifsServer struct {
	value *SoftwarerepositoryCifsServer
	isSet bool
}

func (v NullableSoftwarerepositoryCifsServer) Get() *SoftwarerepositoryCifsServer {
	return v.value
}

func (v *NullableSoftwarerepositoryCifsServer) Set(val *SoftwarerepositoryCifsServer) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryCifsServer) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryCifsServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryCifsServer(val *SoftwarerepositoryCifsServer) *NullableSoftwarerepositoryCifsServer {
	return &NullableSoftwarerepositoryCifsServer{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryCifsServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryCifsServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

