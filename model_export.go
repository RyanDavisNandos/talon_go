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
	"time"
)

// Export
type Export struct {
	// Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// The ID of the account that owns this entity.
	UserId int32 `json:"userId"`
	// The name of the entity that was exported.
	Entity string `json:"entity"`
	// Map of keys and values that were used to filter the exported rows.
	Filter map[string]interface{} `json:"filter"`
}

// GetId returns the Id field value
func (o *Export) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Export) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Export) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *Export) SetCreated(v time.Time) {
	o.Created = v
}

// GetAccountId returns the AccountId field value
func (o *Export) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *Export) SetAccountId(v int32) {
	o.AccountId = v
}

// GetUserId returns the UserId field value
func (o *Export) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// SetUserId sets field value
func (o *Export) SetUserId(v int32) {
	o.UserId = v
}

// GetEntity returns the Entity field value
func (o *Export) GetEntity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Entity
}

// SetEntity sets field value
func (o *Export) SetEntity(v string) {
	o.Entity = v
}

// GetFilter returns the Filter field value
func (o *Export) GetFilter() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Filter
}

// SetFilter sets field value
func (o *Export) SetFilter(v map[string]interface{}) {
	o.Filter = v
}

type NullableExport struct {
	Value        Export
	ExplicitNull bool
}

func (v NullableExport) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableExport) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
