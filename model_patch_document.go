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
)

// PatchDocument A JSONPatch document as defined by RFC 6902. A JSONPatch document can be used in a request body when the 'Content-Type' HTTP header is set to 'application/json-patch+json'.
type PatchDocument struct {
	// The PATCH operation to be performed. 'move' and 'copy' are defined in RFC 6902 but are not supported in the Intersight API.
	Op string `json:"op"`
	// A JSON-Pointer to a property in a REST resource.
	Path string `json:"path"`
	// The value to be used within the operations.
	Value *map[string]interface{} `json:"value,omitempty"`
	// A string containing a JSON Pointer value.
	From *string `json:"from,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchDocument PatchDocument

// NewPatchDocument instantiates a new PatchDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchDocument(op string, path string) *PatchDocument {
	this := PatchDocument{}
	this.Op = op
	this.Path = path
	return &this
}

// NewPatchDocumentWithDefaults instantiates a new PatchDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchDocumentWithDefaults() *PatchDocument {
	this := PatchDocument{}
	return &this
}

// GetOp returns the Op field value
func (o *PatchDocument) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *PatchDocument) GetOpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *PatchDocument) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *PatchDocument) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PatchDocument) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PatchDocument) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PatchDocument) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDocument) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PatchDocument) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *PatchDocument) SetValue(v map[string]interface{}) {
	o.Value = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *PatchDocument) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDocument) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *PatchDocument) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *PatchDocument) SetFrom(v string) {
	o.From = &v
}

func (o PatchDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["op"] = o.Op
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PatchDocument) UnmarshalJSON(bytes []byte) (err error) {
	varPatchDocument := _PatchDocument{}

	if err = json.Unmarshal(bytes, &varPatchDocument); err == nil {
		*o = PatchDocument(varPatchDocument)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		delete(additionalProperties, "from")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchDocument struct {
	value *PatchDocument
	isSet bool
}

func (v NullablePatchDocument) Get() *PatchDocument {
	return v.value
}

func (v *NullablePatchDocument) Set(val *PatchDocument) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchDocument) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchDocument(val *PatchDocument) *NullablePatchDocument {
	return &NullablePatchDocument{value: val, isSet: true}
}

func (v NullablePatchDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


