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

// IncreaseAchievementProgressEffectProps The properties specific to the \"increaseAchievementProgress\" effect. This gets triggered whenever a validated rule contained an \"increase customer progress\" effect.
type IncreaseAchievementProgressEffectProps struct {
	// The internal ID of the achievement.
	AchievementId int32 `json:"achievementId"`
	// The name of the achievement.
	AchievementName string `json:"achievementName"`
	// The internal ID of the achievement progress tracker.
	ProgressTrackerId *int32 `json:"progressTrackerId,omitempty"`
	// The value by which the customer's current progress in the achievement is increased.
	Delta float32 `json:"delta"`
	// The current progress of the customer in the achievement.
	Value float32 `json:"value"`
	// The required number of actions or the transactional milestone to complete the achievement.
	Target float32 `json:"target"`
	// Indicates if the customer has completed the achievement in the current session.
	IsJustCompleted bool `json:"isJustCompleted"`
}

// GetAchievementId returns the AchievementId field value
func (o *IncreaseAchievementProgressEffectProps) GetAchievementId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AchievementId
}

// SetAchievementId sets field value
func (o *IncreaseAchievementProgressEffectProps) SetAchievementId(v int32) {
	o.AchievementId = v
}

// GetAchievementName returns the AchievementName field value
func (o *IncreaseAchievementProgressEffectProps) GetAchievementName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AchievementName
}

// SetAchievementName sets field value
func (o *IncreaseAchievementProgressEffectProps) SetAchievementName(v string) {
	o.AchievementName = v
}

// GetProgressTrackerId returns the ProgressTrackerId field value if set, zero value otherwise.
func (o *IncreaseAchievementProgressEffectProps) GetProgressTrackerId() int32 {
	if o == nil || o.ProgressTrackerId == nil {
		var ret int32
		return ret
	}
	return *o.ProgressTrackerId
}

// GetProgressTrackerIdOk returns a tuple with the ProgressTrackerId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IncreaseAchievementProgressEffectProps) GetProgressTrackerIdOk() (int32, bool) {
	if o == nil || o.ProgressTrackerId == nil {
		var ret int32
		return ret, false
	}
	return *o.ProgressTrackerId, true
}

// HasProgressTrackerId returns a boolean if a field has been set.
func (o *IncreaseAchievementProgressEffectProps) HasProgressTrackerId() bool {
	if o != nil && o.ProgressTrackerId != nil {
		return true
	}

	return false
}

// SetProgressTrackerId gets a reference to the given int32 and assigns it to the ProgressTrackerId field.
func (o *IncreaseAchievementProgressEffectProps) SetProgressTrackerId(v int32) {
	o.ProgressTrackerId = &v
}

// GetDelta returns the Delta field value
func (o *IncreaseAchievementProgressEffectProps) GetDelta() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Delta
}

// SetDelta sets field value
func (o *IncreaseAchievementProgressEffectProps) SetDelta(v float32) {
	o.Delta = v
}

// GetValue returns the Value field value
func (o *IncreaseAchievementProgressEffectProps) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *IncreaseAchievementProgressEffectProps) SetValue(v float32) {
	o.Value = v
}

// GetTarget returns the Target field value
func (o *IncreaseAchievementProgressEffectProps) GetTarget() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Target
}

// SetTarget sets field value
func (o *IncreaseAchievementProgressEffectProps) SetTarget(v float32) {
	o.Target = v
}

// GetIsJustCompleted returns the IsJustCompleted field value
func (o *IncreaseAchievementProgressEffectProps) GetIsJustCompleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsJustCompleted
}

// SetIsJustCompleted sets field value
func (o *IncreaseAchievementProgressEffectProps) SetIsJustCompleted(v bool) {
	o.IsJustCompleted = v
}

type NullableIncreaseAchievementProgressEffectProps struct {
	Value        IncreaseAchievementProgressEffectProps
	ExplicitNull bool
}

func (v NullableIncreaseAchievementProgressEffectProps) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableIncreaseAchievementProgressEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
