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
	"time"
	"reflect"
	"strings"
)

// AssetDeviceConnection An abstract class that holds the details of a device connectors connection to Intersight.
type AssetDeviceConnection struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The version of the connector API, describes the capability of the connector's framework. If the version is lower than the current minimum supported version defined in the service managing the connection, the device connector will be connected with limited capabilities until the device connector is upgraded to a fully supported version. For example if a device connector that was released without delta inventory capabilities registers and connects to Intersight, inventory collection may be disabled until it has been upgraded.
	ApiVersion *int64 `json:"ApiVersion,omitempty"`
	// The partition number corresponding to the instance of the Proxy App which is managing the web-socket to the device connector.
	AppPartitionNumber *int64 `json:"AppPartitionNumber,omitempty"`
	// The unique identifier for the current connection. The identifier persists across network connectivity loss and is reset on device connector process restart or platform administrator toggle of the Intersight connectivity. The connectionId can be used by services that need to interact with stateful plugins running in the device connector process. For example if a service schedules an inventory in a devices job scheduler plugin at registration it is not necessary to reschedule the job if the device loses network connectivity due to an Intersight service upgrade or intermittent network issues in the devices datacenter.
	ConnectionId *string `json:"ConnectionId,omitempty"`
	// If 'connectionStatus' is not equal to Connected, connectionReason provides further details about why the device is not connected with Intersight.
	ConnectionReason *string `json:"ConnectionReason,omitempty"`
	// The status of the persistent connection between the device connector and Intersight. * `` - The target details have been persisted but Intersight has not yet attempted to connect to the target. * `Connected` - Intersight is able to establish a connection to the target and initiate management activities. * `NotConnected` - Intersight is unable to establish a connection to the target. * `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established. * `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight. * `Claimed` - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// The last time at which the 'connectionStatus' property value changed. If connectionStatus is Connected, this time can be interpreted as the starting time since which a persistent connection has been maintained between Intersight and Device Connector. If connectionStatus is NotConnected, this time can be interpreted as the last time the device connector was connected with Intersight.
	ConnectionStatusLastChangeTime *time.Time `json:"ConnectionStatusLastChangeTime,omitempty"`
	// The version of the device connector running on the managed device.
	ConnectorVersion *string `json:"ConnectorVersion,omitempty"`
	// The IP Address of the managed device as seen from Intersight at the time of registration. This could be the IP address of the managed device's interface which has a route to the internet or a NAT IP addresss when the managed device is deployed in a private network.
	DeviceExternalIpAddress *string `json:"DeviceExternalIpAddress,omitempty"`
	// The name of the app which will proxy the messages to the device connector.
	ProxyApp *string `json:"ProxyApp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetDeviceConnection AssetDeviceConnection

// NewAssetDeviceConnection instantiates a new AssetDeviceConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeviceConnection(classId string, objectType string) *AssetDeviceConnection {
	this := AssetDeviceConnection{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetDeviceConnectionWithDefaults instantiates a new AssetDeviceConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeviceConnectionWithDefaults() *AssetDeviceConnection {
	this := AssetDeviceConnection{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetDeviceConnection) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceConnection) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetDeviceConnection) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetDeviceConnection) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceConnection) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetDeviceConnection) SetObjectType(v string) {
	o.ObjectType = v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *AssetDeviceConnection) GetApiVersion() int64 {
	if o == nil || o.ApiVersion == nil {
		var ret int64
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConnection) GetApiVersionOk() (*int64, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *AssetDeviceConnection) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given int64 and assigns it to the ApiVersion field.
func (o *AssetDeviceConnection) SetApiVersion(v int64) {
	o.ApiVersion = &v
}

// GetAppPartitionNumber returns the AppPartitionNumber field value if set, zero value otherwise.
func (o *AssetDeviceConnection) GetAppPartitionNumber() int64 {
	if o == nil || o.AppPartitionNumber == nil {
		var ret int64
		return ret
	}
	return *o.AppPartitionNumber
}

// GetAppPartitionNumberOk returns a tuple with the AppPartitionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConnection) GetAppPartitionNumberOk() (*int64, bool) {
	if o == nil || o.AppPartitionNumber == nil {
		return nil, false
	}
	return o.AppPartitionNumber, true
}

// HasAppPartitionNumber returns a boolean if a field has been set.
func (o *AssetDeviceConnection) HasAppPartitionNumber() bool {
	if o != nil && o.AppPartitionNumber != nil {
		return true
	}

	return false
}

// SetAppPartitionNumber gets a reference to the given int64 and assigns it to the AppPartitionNumber field.
func (o *AssetDeviceConnection) SetAppPartitionNumber(v int64) {
	o.AppPartitionNumber = &v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *AssetDeviceConnection) GetConnectionId() string {
	if o == nil || o.ConnectionId == nil {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConnection) GetConnectionIdOk() (*string, bool) {
	if o == nil || o.ConnectionId == nil {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *AssetDeviceConnection) HasConnectionId() bool {
	if o != nil && o.ConnectionId != nil {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *AssetDeviceConnection) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetConnectionReason returns the ConnectionReason field value if set, zero value otherwise.
func (o *AssetDeviceConnection) GetConnectionReason() string {
	if o == nil || o.ConnectionReason == nil {
		var ret string
		return ret
	}
	return *o.ConnectionReason
}

// GetConnectionReasonOk returns a tuple with the ConnectionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConnection) GetConnectionReasonOk() (*string, bool) {
	if o == nil || o.ConnectionReason == nil {
		return nil, false
	}
	return o.ConnectionReason, true
}

// HasConnectionReason returns a boolean if a field has been set.
func (o *AssetDeviceConnection) HasConnectionReason() bool {
	if o != nil && o.ConnectionReason != nil {
		return true
	}

	return false
}

// SetConnectionReason gets a reference to the given string and assigns it to the ConnectionReason field.
func (o *AssetDeviceConnection) SetConnectionReason(v string) {
	o.ConnectionReason = &v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *AssetDeviceConnection) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConnection) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *AssetDeviceConnection) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *AssetDeviceConnection) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetConnectionStatusLastChangeTime returns the ConnectionStatusLastChangeTime field value if set, zero value otherwise.
func (o *AssetDeviceConnection) GetConnectionStatusLastChangeTime() time.Time {
	if o == nil || o.ConnectionStatusLastChangeTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ConnectionStatusLastChangeTime
}

// GetConnectionStatusLastChangeTimeOk returns a tuple with the ConnectionStatusLastChangeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConnection) GetConnectionStatusLastChangeTimeOk() (*time.Time, bool) {
	if o == nil || o.ConnectionStatusLastChangeTime == nil {
		return nil, false
	}
	return o.ConnectionStatusLastChangeTime, true
}

// HasConnectionStatusLastChangeTime returns a boolean if a field has been set.
func (o *AssetDeviceConnection) HasConnectionStatusLastChangeTime() bool {
	if o != nil && o.ConnectionStatusLastChangeTime != nil {
		return true
	}

	return false
}

// SetConnectionStatusLastChangeTime gets a reference to the given time.Time and assigns it to the ConnectionStatusLastChangeTime field.
func (o *AssetDeviceConnection) SetConnectionStatusLastChangeTime(v time.Time) {
	o.ConnectionStatusLastChangeTime = &v
}

// GetConnectorVersion returns the ConnectorVersion field value if set, zero value otherwise.
func (o *AssetDeviceConnection) GetConnectorVersion() string {
	if o == nil || o.ConnectorVersion == nil {
		var ret string
		return ret
	}
	return *o.ConnectorVersion
}

// GetConnectorVersionOk returns a tuple with the ConnectorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConnection) GetConnectorVersionOk() (*string, bool) {
	if o == nil || o.ConnectorVersion == nil {
		return nil, false
	}
	return o.ConnectorVersion, true
}

// HasConnectorVersion returns a boolean if a field has been set.
func (o *AssetDeviceConnection) HasConnectorVersion() bool {
	if o != nil && o.ConnectorVersion != nil {
		return true
	}

	return false
}

// SetConnectorVersion gets a reference to the given string and assigns it to the ConnectorVersion field.
func (o *AssetDeviceConnection) SetConnectorVersion(v string) {
	o.ConnectorVersion = &v
}

// GetDeviceExternalIpAddress returns the DeviceExternalIpAddress field value if set, zero value otherwise.
func (o *AssetDeviceConnection) GetDeviceExternalIpAddress() string {
	if o == nil || o.DeviceExternalIpAddress == nil {
		var ret string
		return ret
	}
	return *o.DeviceExternalIpAddress
}

// GetDeviceExternalIpAddressOk returns a tuple with the DeviceExternalIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConnection) GetDeviceExternalIpAddressOk() (*string, bool) {
	if o == nil || o.DeviceExternalIpAddress == nil {
		return nil, false
	}
	return o.DeviceExternalIpAddress, true
}

// HasDeviceExternalIpAddress returns a boolean if a field has been set.
func (o *AssetDeviceConnection) HasDeviceExternalIpAddress() bool {
	if o != nil && o.DeviceExternalIpAddress != nil {
		return true
	}

	return false
}

// SetDeviceExternalIpAddress gets a reference to the given string and assigns it to the DeviceExternalIpAddress field.
func (o *AssetDeviceConnection) SetDeviceExternalIpAddress(v string) {
	o.DeviceExternalIpAddress = &v
}

// GetProxyApp returns the ProxyApp field value if set, zero value otherwise.
func (o *AssetDeviceConnection) GetProxyApp() string {
	if o == nil || o.ProxyApp == nil {
		var ret string
		return ret
	}
	return *o.ProxyApp
}

// GetProxyAppOk returns a tuple with the ProxyApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConnection) GetProxyAppOk() (*string, bool) {
	if o == nil || o.ProxyApp == nil {
		return nil, false
	}
	return o.ProxyApp, true
}

// HasProxyApp returns a boolean if a field has been set.
func (o *AssetDeviceConnection) HasProxyApp() bool {
	if o != nil && o.ProxyApp != nil {
		return true
	}

	return false
}

// SetProxyApp gets a reference to the given string and assigns it to the ProxyApp field.
func (o *AssetDeviceConnection) SetProxyApp(v string) {
	o.ProxyApp = &v
}

func (o AssetDeviceConnection) MarshalJSON() ([]byte, error) {
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
	if o.ApiVersion != nil {
		toSerialize["ApiVersion"] = o.ApiVersion
	}
	if o.AppPartitionNumber != nil {
		toSerialize["AppPartitionNumber"] = o.AppPartitionNumber
	}
	if o.ConnectionId != nil {
		toSerialize["ConnectionId"] = o.ConnectionId
	}
	if o.ConnectionReason != nil {
		toSerialize["ConnectionReason"] = o.ConnectionReason
	}
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.ConnectionStatusLastChangeTime != nil {
		toSerialize["ConnectionStatusLastChangeTime"] = o.ConnectionStatusLastChangeTime
	}
	if o.ConnectorVersion != nil {
		toSerialize["ConnectorVersion"] = o.ConnectorVersion
	}
	if o.DeviceExternalIpAddress != nil {
		toSerialize["DeviceExternalIpAddress"] = o.DeviceExternalIpAddress
	}
	if o.ProxyApp != nil {
		toSerialize["ProxyApp"] = o.ProxyApp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetDeviceConnection) UnmarshalJSON(bytes []byte) (err error) {
	type AssetDeviceConnectionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The version of the connector API, describes the capability of the connector's framework. If the version is lower than the current minimum supported version defined in the service managing the connection, the device connector will be connected with limited capabilities until the device connector is upgraded to a fully supported version. For example if a device connector that was released without delta inventory capabilities registers and connects to Intersight, inventory collection may be disabled until it has been upgraded.
		ApiVersion *int64 `json:"ApiVersion,omitempty"`
		// The partition number corresponding to the instance of the Proxy App which is managing the web-socket to the device connector.
		AppPartitionNumber *int64 `json:"AppPartitionNumber,omitempty"`
		// The unique identifier for the current connection. The identifier persists across network connectivity loss and is reset on device connector process restart or platform administrator toggle of the Intersight connectivity. The connectionId can be used by services that need to interact with stateful plugins running in the device connector process. For example if a service schedules an inventory in a devices job scheduler plugin at registration it is not necessary to reschedule the job if the device loses network connectivity due to an Intersight service upgrade or intermittent network issues in the devices datacenter.
		ConnectionId *string `json:"ConnectionId,omitempty"`
		// If 'connectionStatus' is not equal to Connected, connectionReason provides further details about why the device is not connected with Intersight.
		ConnectionReason *string `json:"ConnectionReason,omitempty"`
		// The status of the persistent connection between the device connector and Intersight. * `` - The target details have been persisted but Intersight has not yet attempted to connect to the target. * `Connected` - Intersight is able to establish a connection to the target and initiate management activities. * `NotConnected` - Intersight is unable to establish a connection to the target. * `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established. * `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight. * `Claimed` - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.
		ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
		// The last time at which the 'connectionStatus' property value changed. If connectionStatus is Connected, this time can be interpreted as the starting time since which a persistent connection has been maintained between Intersight and Device Connector. If connectionStatus is NotConnected, this time can be interpreted as the last time the device connector was connected with Intersight.
		ConnectionStatusLastChangeTime *time.Time `json:"ConnectionStatusLastChangeTime,omitempty"`
		// The version of the device connector running on the managed device.
		ConnectorVersion *string `json:"ConnectorVersion,omitempty"`
		// The IP Address of the managed device as seen from Intersight at the time of registration. This could be the IP address of the managed device's interface which has a route to the internet or a NAT IP addresss when the managed device is deployed in a private network.
		DeviceExternalIpAddress *string `json:"DeviceExternalIpAddress,omitempty"`
		// The name of the app which will proxy the messages to the device connector.
		ProxyApp *string `json:"ProxyApp,omitempty"`
	}

	varAssetDeviceConnectionWithoutEmbeddedStruct := AssetDeviceConnectionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetDeviceConnectionWithoutEmbeddedStruct)
	if err == nil {
		varAssetDeviceConnection := _AssetDeviceConnection{}
		varAssetDeviceConnection.ClassId = varAssetDeviceConnectionWithoutEmbeddedStruct.ClassId
		varAssetDeviceConnection.ObjectType = varAssetDeviceConnectionWithoutEmbeddedStruct.ObjectType
		varAssetDeviceConnection.ApiVersion = varAssetDeviceConnectionWithoutEmbeddedStruct.ApiVersion
		varAssetDeviceConnection.AppPartitionNumber = varAssetDeviceConnectionWithoutEmbeddedStruct.AppPartitionNumber
		varAssetDeviceConnection.ConnectionId = varAssetDeviceConnectionWithoutEmbeddedStruct.ConnectionId
		varAssetDeviceConnection.ConnectionReason = varAssetDeviceConnectionWithoutEmbeddedStruct.ConnectionReason
		varAssetDeviceConnection.ConnectionStatus = varAssetDeviceConnectionWithoutEmbeddedStruct.ConnectionStatus
		varAssetDeviceConnection.ConnectionStatusLastChangeTime = varAssetDeviceConnectionWithoutEmbeddedStruct.ConnectionStatusLastChangeTime
		varAssetDeviceConnection.ConnectorVersion = varAssetDeviceConnectionWithoutEmbeddedStruct.ConnectorVersion
		varAssetDeviceConnection.DeviceExternalIpAddress = varAssetDeviceConnectionWithoutEmbeddedStruct.DeviceExternalIpAddress
		varAssetDeviceConnection.ProxyApp = varAssetDeviceConnectionWithoutEmbeddedStruct.ProxyApp
		*o = AssetDeviceConnection(varAssetDeviceConnection)
	} else {
		return err
	}

	varAssetDeviceConnection := _AssetDeviceConnection{}

	err = json.Unmarshal(bytes, &varAssetDeviceConnection)
	if err == nil {
		o.MoBaseMo = varAssetDeviceConnection.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApiVersion")
		delete(additionalProperties, "AppPartitionNumber")
		delete(additionalProperties, "ConnectionId")
		delete(additionalProperties, "ConnectionReason")
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "ConnectionStatusLastChangeTime")
		delete(additionalProperties, "ConnectorVersion")
		delete(additionalProperties, "DeviceExternalIpAddress")
		delete(additionalProperties, "ProxyApp")

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

type NullableAssetDeviceConnection struct {
	value *AssetDeviceConnection
	isSet bool
}

func (v NullableAssetDeviceConnection) Get() *AssetDeviceConnection {
	return v.value
}

func (v *NullableAssetDeviceConnection) Set(val *AssetDeviceConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeviceConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeviceConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeviceConnection(val *AssetDeviceConnection) *NullableAssetDeviceConnection {
	return &NullableAssetDeviceConnection{value: val, isSet: true}
}

func (v NullableAssetDeviceConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeviceConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


