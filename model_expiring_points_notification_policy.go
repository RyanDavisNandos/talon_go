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
)

// ExpiringPointsNotificationPolicy struct for ExpiringPointsNotificationPolicy
type ExpiringPointsNotificationPolicy struct {
	// Notification name.
	Name     string                              `json:"name"`
	Triggers []ExpiringPointsNotificationTrigger `json:"triggers"`
	// Indicates whether batching is activated.
	BatchingEnabled *bool `json:"batchingEnabled,omitempty"`
}

// GetName returns the Name field value
func (o *ExpiringPointsNotificationPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *ExpiringPointsNotificationPolicy) SetName(v string) {
	o.Name = v
}

// GetTriggers returns the Triggers field value
func (o *ExpiringPointsNotificationPolicy) GetTriggers() []ExpiringPointsNotificationTrigger {
	if o == nil {
		var ret []ExpiringPointsNotificationTrigger
		return ret
	}

	return o.Triggers
}

// SetTriggers sets field value
func (o *ExpiringPointsNotificationPolicy) SetTriggers(v []ExpiringPointsNotificationTrigger) {
	o.Triggers = v
}

// GetBatchingEnabled returns the BatchingEnabled field value if set, zero value otherwise.
func (o *ExpiringPointsNotificationPolicy) GetBatchingEnabled() bool {
	if o == nil || o.BatchingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BatchingEnabled
}

// GetBatchingEnabledOk returns a tuple with the BatchingEnabled field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ExpiringPointsNotificationPolicy) GetBatchingEnabledOk() (bool, bool) {
	if o == nil || o.BatchingEnabled == nil {
		var ret bool
		return ret, false
	}
	return *o.BatchingEnabled, true
}

// HasBatchingEnabled returns a boolean if a field has been set.
func (o *ExpiringPointsNotificationPolicy) HasBatchingEnabled() bool {
	if o != nil && o.BatchingEnabled != nil {
		return true
	}

	return false
}

// SetBatchingEnabled gets a reference to the given bool and assigns it to the BatchingEnabled field.
func (o *ExpiringPointsNotificationPolicy) SetBatchingEnabled(v bool) {
	o.BatchingEnabled = &v
}

type NullableExpiringPointsNotificationPolicy struct {
	Value        ExpiringPointsNotificationPolicy
	ExplicitNull bool
}

func (v NullableExpiringPointsNotificationPolicy) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableExpiringPointsNotificationPolicy) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
