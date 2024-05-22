/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}`
 *
 * API version:
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
	"time"
)

// Store
type Store struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created. The time this entity was created.
	Created time.Time `json:"created"`
	// The name of the store.
	Name string `json:"name"`
	// The description of the store.
	Description string `json:"description"`
	// The attributes of the store.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// The integration ID of the store. You choose this ID when you create a store.  **Note**: You cannot edit the `integrationId` after the store has been created.
	IntegrationId string `json:"integrationId"`
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// Timestamp of the most recent update on this entity.
	Updated time.Time `json:"updated"`
	// A list of IDs of the campaigns that are linked with current store.
	LinkedCampaignIds *[]int32 `json:"linkedCampaignIds,omitempty"`
}

// GetId returns the Id field value
func (o *Store) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Store) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Store) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *Store) SetCreated(v time.Time) {
	o.Created = v
}

// GetName returns the Name field value
func (o *Store) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *Store) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Store) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// SetDescription sets field value
func (o *Store) SetDescription(v string) {
	o.Description = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Store) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Store) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *Store) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetIntegrationId returns the IntegrationId field value
func (o *Store) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// SetIntegrationId sets field value
func (o *Store) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetApplicationId returns the ApplicationId field value
func (o *Store) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// SetApplicationId sets field value
func (o *Store) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetUpdated returns the Updated field value
func (o *Store) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// SetUpdated sets field value
func (o *Store) SetUpdated(v time.Time) {
	o.Updated = v
}

// GetLinkedCampaignIds returns the LinkedCampaignIds field value if set, zero value otherwise.
func (o *Store) GetLinkedCampaignIds() []int32 {
	if o == nil || o.LinkedCampaignIds == nil {
		var ret []int32
		return ret
	}
	return *o.LinkedCampaignIds
}

// GetLinkedCampaignIdsOk returns a tuple with the LinkedCampaignIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetLinkedCampaignIdsOk() ([]int32, bool) {
	if o == nil || o.LinkedCampaignIds == nil {
		var ret []int32
		return ret, false
	}
	return *o.LinkedCampaignIds, true
}

// HasLinkedCampaignIds returns a boolean if a field has been set.
func (o *Store) HasLinkedCampaignIds() bool {
	if o != nil && o.LinkedCampaignIds != nil {
		return true
	}

	return false
}

// SetLinkedCampaignIds gets a reference to the given []int32 and assigns it to the LinkedCampaignIds field.
func (o *Store) SetLinkedCampaignIds(v []int32) {
	o.LinkedCampaignIds = &v
}

type NullableStore struct {
	Value        Store
	ExplicitNull bool
}

func (v NullableStore) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableStore) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
