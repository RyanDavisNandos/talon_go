/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerSession](https://docs.talon.one/integration-api/#operation/updateCustomerSessionV2) endpoint is `https://mycompany.talon.one/v2/customer_sessions/{Id}`
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// IntegrationEventV2Request
type IntegrationEventV2Request struct {
	// ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known `profileId`, we recommend you use a guest `profileId`.
	ProfileId *string `json:"profileId,omitempty"`
	// A string representing the event. Must not be a reserved event name.
	Type string `json:"type"`
	// Arbitrary additional JSON data associated with the event.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// Optional list of requested information to be present on the response related to the tracking custom event.
	ResponseContent *[]string `json:"responseContent,omitempty"`
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *IntegrationEventV2Request) GetProfileId() string {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationEventV2Request) GetProfileIdOk() (string, bool) {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret, false
	}
	return *o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *IntegrationEventV2Request) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *IntegrationEventV2Request) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetType returns the Type field value
func (o *IntegrationEventV2Request) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *IntegrationEventV2Request) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IntegrationEventV2Request) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationEventV2Request) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IntegrationEventV2Request) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *IntegrationEventV2Request) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *IntegrationEventV2Request) GetResponseContent() []string {
	if o == nil || o.ResponseContent == nil {
		var ret []string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationEventV2Request) GetResponseContentOk() ([]string, bool) {
	if o == nil || o.ResponseContent == nil {
		var ret []string
		return ret, false
	}
	return *o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *IntegrationEventV2Request) HasResponseContent() bool {
	if o != nil && o.ResponseContent != nil {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given []string and assigns it to the ResponseContent field.
func (o *IntegrationEventV2Request) SetResponseContent(v []string) {
	o.ResponseContent = &v
}

type NullableIntegrationEventV2Request struct {
	Value        IntegrationEventV2Request
	ExplicitNull bool
}

func (v NullableIntegrationEventV2Request) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableIntegrationEventV2Request) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
