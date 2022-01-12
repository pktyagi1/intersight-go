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
)

// AssetDeviceInformationAllOf Definition of the list of properties defined in 'asset.DeviceInformation', excluding properties defined in parent classes.
type AssetDeviceInformationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Application name reported by Cisco Install Base.
	ApplicationName *string `json:"ApplicationName,omitempty"`
	DeviceTransactions []AssetDeviceTransaction `json:"DeviceTransactions,omitempty"`
	EndCustomer NullableAssetCustomerInformation `json:"EndCustomer,omitempty"`
	// Instance number of the device. example \"917280220\".
	InstanceId *string `json:"InstanceId,omitempty"`
	// Item type flag. example ATO, Child, Standalone. ATO - refers to Cisco Block based major device. Child - refers to Child device part of an ATO block. Standalone - refers to a device that is managed and configured as an individual entity with limited capacity.
	ItemType *string `json:"ItemType,omitempty"`
	// Identifier for consumption based pricing. MLB refers to MultiLine Bundle.
	MlbOfferType *string `json:"MlbOfferType,omitempty"`
	// Identifier corresponding to MLB Product Name. MLB refers to MultiLine Bundle.
	MlbProductId *int64 `json:"MlbProductId,omitempty"`
	// Product Name for the device. It is used to determine if the server is to be billed as a UCS compute device or a storage cluster. MLB refers to MultiLine Bundle.
	MlbProductName *string `json:"MlbProductName,omitempty"`
	// Unique identifier of old Cisco device. It is the device ID of old Cisco device, which got replaced by the new device.
	OldDeviceId *string `json:"OldDeviceId,omitempty"`
	// Description of status of old Cisco device, which got replaced by the new device.
	OldDeviceStatusDescription *string `json:"OldDeviceStatusDescription,omitempty"`
	// Status ID of old Cisco device, which got replaced by the new device. * `0` - A default value to catch cases where device status is not correctly detected. * `10000` - Device is installed successfully. * `1010041` - Device is currently in Return Material Authorization process. * `10005` - Device is replaced successfully with another device. * `10007` - Device is returned succcessfuly. * `10009` - Device is terminated at customer's end.
	OldDeviceStatusId *int32 `json:"OldDeviceStatusId,omitempty"`
	// Instance number of the old device, which got replaced by the new device.
	OldInstanceId *string `json:"OldInstanceId,omitempty"`
	// Product Family is the field used to identify the hypervisor type. example \"ESXi\".
	ProductFamily *string `json:"ProductFamily,omitempty"`
	// Product type helps to determine if device has to be billed using consumption metering. example \"SERVER\".
	ProductType *string `json:"ProductType,omitempty"`
	// Unit of Measure is flag used to identify the type of metric being pushed. example - \"STORAGE\" for hardware metrics , \"VM\" - for hypervisor metrics. * `None` - A default value to catch cases where unit of measure is not correctly detected. * `STORAGE` - The metric type of the device is a storage metric. * `NODE` - The metric type of the device is a hardware metric. * `VM` - The metric type of the device is a hypervisor metric.
	UnitOfMeasure *string `json:"UnitOfMeasure,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetDeviceInformationAllOf AssetDeviceInformationAllOf

// NewAssetDeviceInformationAllOf instantiates a new AssetDeviceInformationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeviceInformationAllOf(classId string, objectType string) *AssetDeviceInformationAllOf {
	this := AssetDeviceInformationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetDeviceInformationAllOfWithDefaults instantiates a new AssetDeviceInformationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeviceInformationAllOfWithDefaults() *AssetDeviceInformationAllOf {
	this := AssetDeviceInformationAllOf{}
	var classId string = "asset.DeviceInformation"
	this.ClassId = classId
	var objectType string = "asset.DeviceInformation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetDeviceInformationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetDeviceInformationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetDeviceInformationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetDeviceInformationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *AssetDeviceInformationAllOf) GetApplicationName() string {
	if o == nil || o.ApplicationName == nil {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformationAllOf) GetApplicationNameOk() (*string, bool) {
	if o == nil || o.ApplicationName == nil {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *AssetDeviceInformationAllOf) HasApplicationName() bool {
	if o != nil && o.ApplicationName != nil {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *AssetDeviceInformationAllOf) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetDeviceTransactions returns the DeviceTransactions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceInformationAllOf) GetDeviceTransactions() []AssetDeviceTransaction {
	if o == nil  {
		var ret []AssetDeviceTransaction
		return ret
	}
	return o.DeviceTransactions
}

// GetDeviceTransactionsOk returns a tuple with the DeviceTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceInformationAllOf) GetDeviceTransactionsOk() (*[]AssetDeviceTransaction, bool) {
	if o == nil || o.DeviceTransactions == nil {
		return nil, false
	}
	return &o.DeviceTransactions, true
}

// HasDeviceTransactions returns a boolean if a field has been set.
func (o *AssetDeviceInformationAllOf) HasDeviceTransactions() bool {
	if o != nil && o.DeviceTransactions != nil {
		return true
	}

	return false
}

// SetDeviceTransactions gets a reference to the given []AssetDeviceTransaction and assigns it to the DeviceTransactions field.
func (o *AssetDeviceInformationAllOf) SetDeviceTransactions(v []AssetDeviceTransaction) {
	o.DeviceTransactions = v
}

// GetEndCustomer returns the EndCustomer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceInformationAllOf) GetEndCustomer() AssetCustomerInformation {
	if o == nil || o.EndCustomer.Get() == nil {
		var ret AssetCustomerInformation
		return ret
	}
	return *o.EndCustomer.Get()
}

// GetEndCustomerOk returns a tuple with the EndCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceInformationAllOf) GetEndCustomerOk() (*AssetCustomerInformation, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndCustomer.Get(), o.EndCustomer.IsSet()
}

// HasEndCustomer returns a boolean if a field has been set.
func (o *AssetDeviceInformationAllOf) HasEndCustomer() bool {
	if o != nil && o.EndCustomer.IsSet() {
		return true
	}

	return false
}

// SetEndCustomer gets a reference to the given NullableAssetCustomerInformation and assigns it to the EndCustomer field.
func (o *AssetDeviceInformationAllOf) SetEndCustomer(v AssetCustomerInformation) {
	o.EndCustomer.Set(&v)
}
// SetEndCustomerNil sets the value for EndCustomer to be an explicit nil
func (o *AssetDeviceInformationAllOf) SetEndCustomerNil() {
	o.EndCustomer.Set(nil)
}

// UnsetEndCustomer ensures that no value is present for EndCustomer, not even an explicit nil
func (o *AssetDeviceInformationAllOf) UnsetEndCustomer() {
	o.EndCustomer.Unset()
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *AssetDeviceInformationAllOf) GetInstanceId() string {
	if o == nil || o.InstanceId == nil {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformationAllOf) GetInstanceIdOk() (*string, bool) {
	if o == nil || o.InstanceId == nil {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *AssetDeviceInformationAllOf) HasInstanceId() bool {
	if o != nil && o.InstanceId != nil {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *AssetDeviceInformationAllOf) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *AssetDeviceInformationAllOf) GetItemType() string {
	if o == nil || o.ItemType == nil {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformationAllOf) GetItemTypeOk() (*string, bool) {
	if o == nil || o.ItemType == nil {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *AssetDeviceInformationAllOf) HasItemType() bool {
	if o != nil && o.ItemType != nil {
		return true
	}

	return false
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *AssetDeviceInformationAllOf) SetItemType(v string) {
	o.ItemType = &v
}

// GetMlbOfferType returns the MlbOfferType field value if set, zero value otherwise.
func (o *AssetDeviceInformationAllOf) GetMlbOfferType() string {
	if o == nil || o.MlbOfferType == nil {
		var ret string
		return ret
	}
	return *o.MlbOfferType
}

// GetMlbOfferTypeOk returns a tuple with the MlbOfferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformationAllOf) GetMlbOfferTypeOk() (*string, bool) {
	if o == nil || o.MlbOfferType == nil {
		return nil, false
	}
	return o.MlbOfferType, true
}

// HasMlbOfferType returns a boolean if a field has been set.
func (o *AssetDeviceInformationAllOf) HasMlbOfferType() bool {
	if o != nil && o.MlbOfferType != nil {
		return true
	}

	return false
}

// SetMlbOfferType gets a reference to the given string and assigns it to the MlbOfferType field.
func (o *AssetDeviceInformationAllOf) SetMlbOfferType(v string) {
	o.MlbOfferType = &v
}

// GetMlbProductId returns the MlbProductId field value if set, zero value otherwise.
func (o *AssetDeviceInformationAllOf) GetMlbProductId() int64 {
	if o == nil || o.MlbProductId == nil {
		var ret int64
		return ret
	}
	return *o.MlbProductId
}

// GetMlbProductIdOk returns a tuple with the MlbProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformationAllOf) GetMlbProductIdOk() (*int64, bool) {
	if o == nil || o.MlbProductId == nil {
		return nil, false
	}
	return o.MlbProductId, true
}

// HasMlbProductId returns a boolean if a field has been set.
func (o *AssetDeviceInformationAllOf) HasMlbProductId() bool {
	if o != nil && o.MlbProductId != nil {
		return true
	}

	return false
}

// SetMlbProductId gets a reference to the given int64 and assigns it to the MlbProductId field.
func (o *AssetDeviceInformationAllOf) SetMlbProductId(v int64) {
	o.MlbProductId = &v
}

// GetMlbProductName returns the MlbProductName field value if set, zero value otherwise.
func (o *AssetDeviceInformationAllOf) GetMlbProductName() string {
	if o == nil || o.MlbProductName == nil {
		var ret string
		return ret
	}
	return *o.MlbProductName
}

// GetMlbProductNameOk returns a tuple with the MlbProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformationAllOf) GetMlbProductNameOk() (*string, bool) {
	if o == nil || o.MlbProductName == nil {
		return nil, false
	}
	return o.MlbProductName, true
}

// HasMlbProductName returns a boolean if a field has been set.
func (o *AssetDeviceInformationAllOf) HasMlbProductName() bool {
	if o != nil && o.MlbProductName != nil {
		return true
	}

	return false
}

// SetMlbProductName gets a reference to the given string and assigns it to the MlbProductName field.
func (o *AssetDeviceInformationAllOf) SetMlbProductName(v string) {
	o.MlbProductName = &v
}

// GetOldDeviceId returns the OldDeviceId field value if set, zero value otherwise.
func (o *AssetDeviceInformationAllOf) GetOldDeviceId() string {
	if o == nil || o.OldDeviceId == nil {
		var ret string
		return ret
	}
	return *o.OldDeviceId
}

// GetOldDeviceIdOk returns a tuple with the OldDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformationAllOf) GetOldDeviceIdOk() (*string, bool) {
	if o == nil || o.OldDeviceId == nil {
		return nil, false
	}
	return o.OldDeviceId, true
}

// HasOldDeviceId returns a boolean if a field has been set.
func (o *AssetDeviceInformationAllOf) HasOldDeviceId() bool {
	if o != nil && o.OldDeviceId != nil {
		return true
	}

	return false
}

// SetOldDeviceId gets a reference to the given string and assigns it to the OldDeviceId field.
func (o *AssetDeviceInformationAllOf) SetOldDeviceId(v string) {
	o.OldDeviceId = &v
}

// GetOldDeviceStatusDescription returns the OldDeviceStatusDescription field value if set, zero value otherwise.
func (o *AssetDeviceInformationAllOf) GetOldDeviceStatusDescription() string {
	if o == nil || o.OldDeviceStatusDescription == nil {
		var ret string
		return ret
	}
	return *o.OldDeviceStatusDescription
}

// GetOldDeviceStatusDescriptionOk returns a tuple with the OldDeviceStatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformationAllOf) GetOldDeviceStatusDescriptionOk() (*string, bool) {
	if o == nil || o.OldDeviceStatusDescription == nil {
		return nil, false
	}
	return o.OldDeviceStatusDescription, true
}

// HasOldDeviceStatusDescription returns a boolean if a field has been set.
func (o *AssetDeviceInformationAllOf) HasOldDeviceStatusDescription() bool {
	if o != nil && o.OldDeviceStatusDescription != nil {
		return true
	}

	return false
}

// SetOldDeviceStatusDescription gets a reference to the given string and assigns it to the OldDeviceStatusDescription field.
func (o *AssetDeviceInformationAllOf) SetOldDeviceStatusDescription(v string) {
	o.OldDeviceStatusDescription = &v
}

// GetOldDeviceStatusId returns the OldDeviceStatusId field value if set, zero value otherwise.
func (o *AssetDeviceInformationAllOf) GetOldDeviceStatusId() int32 {
	if o == nil || o.OldDeviceStatusId == nil {
		var ret int32
		return ret
	}
	return *o.OldDeviceStatusId
}

// GetOldDeviceStatusIdOk returns a tuple with the OldDeviceStatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformationAllOf) GetOldDeviceStatusIdOk() (*int32, bool) {
	if o == nil || o.OldDeviceStatusId == nil {
		return nil, false
	}
	return o.OldDeviceStatusId, true
}

// HasOldDeviceStatusId returns a boolean if a field has been set.
func (o *AssetDeviceInformationAllOf) HasOldDeviceStatusId() bool {
	if o != nil && o.OldDeviceStatusId != nil {
		return true
	}

	return false
}

// SetOldDeviceStatusId gets a reference to the given int32 and assigns it to the OldDeviceStatusId field.
func (o *AssetDeviceInformationAllOf) SetOldDeviceStatusId(v int32) {
	o.OldDeviceStatusId = &v
}

// GetOldInstanceId returns the OldInstanceId field value if set, zero value otherwise.
func (o *AssetDeviceInformationAllOf) GetOldInstanceId() string {
	if o == nil || o.OldInstanceId == nil {
		var ret string
		return ret
	}
	return *o.OldInstanceId
}

// GetOldInstanceIdOk returns a tuple with the OldInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformationAllOf) GetOldInstanceIdOk() (*string, bool) {
	if o == nil || o.OldInstanceId == nil {
		return nil, false
	}
	return o.OldInstanceId, true
}

// HasOldInstanceId returns a boolean if a field has been set.
func (o *AssetDeviceInformationAllOf) HasOldInstanceId() bool {
	if o != nil && o.OldInstanceId != nil {
		return true
	}

	return false
}

// SetOldInstanceId gets a reference to the given string and assigns it to the OldInstanceId field.
func (o *AssetDeviceInformationAllOf) SetOldInstanceId(v string) {
	o.OldInstanceId = &v
}

// GetProductFamily returns the ProductFamily field value if set, zero value otherwise.
func (o *AssetDeviceInformationAllOf) GetProductFamily() string {
	if o == nil || o.ProductFamily == nil {
		var ret string
		return ret
	}
	return *o.ProductFamily
}

// GetProductFamilyOk returns a tuple with the ProductFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformationAllOf) GetProductFamilyOk() (*string, bool) {
	if o == nil || o.ProductFamily == nil {
		return nil, false
	}
	return o.ProductFamily, true
}

// HasProductFamily returns a boolean if a field has been set.
func (o *AssetDeviceInformationAllOf) HasProductFamily() bool {
	if o != nil && o.ProductFamily != nil {
		return true
	}

	return false
}

// SetProductFamily gets a reference to the given string and assigns it to the ProductFamily field.
func (o *AssetDeviceInformationAllOf) SetProductFamily(v string) {
	o.ProductFamily = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *AssetDeviceInformationAllOf) GetProductType() string {
	if o == nil || o.ProductType == nil {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformationAllOf) GetProductTypeOk() (*string, bool) {
	if o == nil || o.ProductType == nil {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *AssetDeviceInformationAllOf) HasProductType() bool {
	if o != nil && o.ProductType != nil {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *AssetDeviceInformationAllOf) SetProductType(v string) {
	o.ProductType = &v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value if set, zero value otherwise.
func (o *AssetDeviceInformationAllOf) GetUnitOfMeasure() string {
	if o == nil || o.UnitOfMeasure == nil {
		var ret string
		return ret
	}
	return *o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformationAllOf) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil || o.UnitOfMeasure == nil {
		return nil, false
	}
	return o.UnitOfMeasure, true
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *AssetDeviceInformationAllOf) HasUnitOfMeasure() bool {
	if o != nil && o.UnitOfMeasure != nil {
		return true
	}

	return false
}

// SetUnitOfMeasure gets a reference to the given string and assigns it to the UnitOfMeasure field.
func (o *AssetDeviceInformationAllOf) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = &v
}

func (o AssetDeviceInformationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ApplicationName != nil {
		toSerialize["ApplicationName"] = o.ApplicationName
	}
	if o.DeviceTransactions != nil {
		toSerialize["DeviceTransactions"] = o.DeviceTransactions
	}
	if o.EndCustomer.IsSet() {
		toSerialize["EndCustomer"] = o.EndCustomer.Get()
	}
	if o.InstanceId != nil {
		toSerialize["InstanceId"] = o.InstanceId
	}
	if o.ItemType != nil {
		toSerialize["ItemType"] = o.ItemType
	}
	if o.MlbOfferType != nil {
		toSerialize["MlbOfferType"] = o.MlbOfferType
	}
	if o.MlbProductId != nil {
		toSerialize["MlbProductId"] = o.MlbProductId
	}
	if o.MlbProductName != nil {
		toSerialize["MlbProductName"] = o.MlbProductName
	}
	if o.OldDeviceId != nil {
		toSerialize["OldDeviceId"] = o.OldDeviceId
	}
	if o.OldDeviceStatusDescription != nil {
		toSerialize["OldDeviceStatusDescription"] = o.OldDeviceStatusDescription
	}
	if o.OldDeviceStatusId != nil {
		toSerialize["OldDeviceStatusId"] = o.OldDeviceStatusId
	}
	if o.OldInstanceId != nil {
		toSerialize["OldInstanceId"] = o.OldInstanceId
	}
	if o.ProductFamily != nil {
		toSerialize["ProductFamily"] = o.ProductFamily
	}
	if o.ProductType != nil {
		toSerialize["ProductType"] = o.ProductType
	}
	if o.UnitOfMeasure != nil {
		toSerialize["UnitOfMeasure"] = o.UnitOfMeasure
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetDeviceInformationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetDeviceInformationAllOf := _AssetDeviceInformationAllOf{}

	if err = json.Unmarshal(bytes, &varAssetDeviceInformationAllOf); err == nil {
		*o = AssetDeviceInformationAllOf(varAssetDeviceInformationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApplicationName")
		delete(additionalProperties, "DeviceTransactions")
		delete(additionalProperties, "EndCustomer")
		delete(additionalProperties, "InstanceId")
		delete(additionalProperties, "ItemType")
		delete(additionalProperties, "MlbOfferType")
		delete(additionalProperties, "MlbProductId")
		delete(additionalProperties, "MlbProductName")
		delete(additionalProperties, "OldDeviceId")
		delete(additionalProperties, "OldDeviceStatusDescription")
		delete(additionalProperties, "OldDeviceStatusId")
		delete(additionalProperties, "OldInstanceId")
		delete(additionalProperties, "ProductFamily")
		delete(additionalProperties, "ProductType")
		delete(additionalProperties, "UnitOfMeasure")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetDeviceInformationAllOf struct {
	value *AssetDeviceInformationAllOf
	isSet bool
}

func (v NullableAssetDeviceInformationAllOf) Get() *AssetDeviceInformationAllOf {
	return v.value
}

func (v *NullableAssetDeviceInformationAllOf) Set(val *AssetDeviceInformationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeviceInformationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeviceInformationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeviceInformationAllOf(val *AssetDeviceInformationAllOf) *NullableAssetDeviceInformationAllOf {
	return &NullableAssetDeviceInformationAllOf{value: val, isSet: true}
}

func (v NullableAssetDeviceInformationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeviceInformationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

