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

// RuleFailureReason Details about why a rule failed.
type RuleFailureReason struct {
	// The ID of the campaign that contains the rule that failed.
	CampaignID int32 `json:"campaignID"`
	// The name of the campaign that contains the rule that failed.
	CampaignName string `json:"campaignName"`
	// The ID of the ruleset that contains the rule that failed.
	RulesetID int32 `json:"rulesetID"`
	// The ID of the coupon that was being evaluated at the time of the rule failure.
	CouponID *int32 `json:"couponID,omitempty"`
	// The code of the coupon that was being evaluated at the time of the rule failure.
	CouponValue *string `json:"couponValue,omitempty"`
	// The ID of the referral that was being evaluated at the time of the rule failure.
	ReferralID *int32 `json:"referralID,omitempty"`
	// The code of the referral that was being evaluated at the time of the rule failure.
	ReferralValue *string `json:"referralValue,omitempty"`
	// The index of the rule that failed within the ruleset.
	RuleIndex int32 `json:"ruleIndex"`
	// The name of the rule that failed within the ruleset.
	RuleName string `json:"ruleName"`
	// The index of the condition that failed.
	ConditionIndex *int32 `json:"conditionIndex,omitempty"`
	// The index of the effect that failed.
	EffectIndex *int32 `json:"effectIndex,omitempty"`
	// More details about the failure.
	Details *string `json:"details,omitempty"`
}

// GetCampaignID returns the CampaignID field value
func (o *RuleFailureReason) GetCampaignID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignID
}

// SetCampaignID sets field value
func (o *RuleFailureReason) SetCampaignID(v int32) {
	o.CampaignID = v
}

// GetCampaignName returns the CampaignName field value
func (o *RuleFailureReason) GetCampaignName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignName
}

// SetCampaignName sets field value
func (o *RuleFailureReason) SetCampaignName(v string) {
	o.CampaignName = v
}

// GetRulesetID returns the RulesetID field value
func (o *RuleFailureReason) GetRulesetID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RulesetID
}

// SetRulesetID sets field value
func (o *RuleFailureReason) SetRulesetID(v int32) {
	o.RulesetID = v
}

// GetCouponID returns the CouponID field value if set, zero value otherwise.
func (o *RuleFailureReason) GetCouponID() int32 {
	if o == nil || o.CouponID == nil {
		var ret int32
		return ret
	}
	return *o.CouponID
}

// GetCouponIDOk returns a tuple with the CouponID field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RuleFailureReason) GetCouponIDOk() (int32, bool) {
	if o == nil || o.CouponID == nil {
		var ret int32
		return ret, false
	}
	return *o.CouponID, true
}

// HasCouponID returns a boolean if a field has been set.
func (o *RuleFailureReason) HasCouponID() bool {
	if o != nil && o.CouponID != nil {
		return true
	}

	return false
}

// SetCouponID gets a reference to the given int32 and assigns it to the CouponID field.
func (o *RuleFailureReason) SetCouponID(v int32) {
	o.CouponID = &v
}

// GetCouponValue returns the CouponValue field value if set, zero value otherwise.
func (o *RuleFailureReason) GetCouponValue() string {
	if o == nil || o.CouponValue == nil {
		var ret string
		return ret
	}
	return *o.CouponValue
}

// GetCouponValueOk returns a tuple with the CouponValue field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RuleFailureReason) GetCouponValueOk() (string, bool) {
	if o == nil || o.CouponValue == nil {
		var ret string
		return ret, false
	}
	return *o.CouponValue, true
}

// HasCouponValue returns a boolean if a field has been set.
func (o *RuleFailureReason) HasCouponValue() bool {
	if o != nil && o.CouponValue != nil {
		return true
	}

	return false
}

// SetCouponValue gets a reference to the given string and assigns it to the CouponValue field.
func (o *RuleFailureReason) SetCouponValue(v string) {
	o.CouponValue = &v
}

// GetReferralID returns the ReferralID field value if set, zero value otherwise.
func (o *RuleFailureReason) GetReferralID() int32 {
	if o == nil || o.ReferralID == nil {
		var ret int32
		return ret
	}
	return *o.ReferralID
}

// GetReferralIDOk returns a tuple with the ReferralID field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RuleFailureReason) GetReferralIDOk() (int32, bool) {
	if o == nil || o.ReferralID == nil {
		var ret int32
		return ret, false
	}
	return *o.ReferralID, true
}

// HasReferralID returns a boolean if a field has been set.
func (o *RuleFailureReason) HasReferralID() bool {
	if o != nil && o.ReferralID != nil {
		return true
	}

	return false
}

// SetReferralID gets a reference to the given int32 and assigns it to the ReferralID field.
func (o *RuleFailureReason) SetReferralID(v int32) {
	o.ReferralID = &v
}

// GetReferralValue returns the ReferralValue field value if set, zero value otherwise.
func (o *RuleFailureReason) GetReferralValue() string {
	if o == nil || o.ReferralValue == nil {
		var ret string
		return ret
	}
	return *o.ReferralValue
}

// GetReferralValueOk returns a tuple with the ReferralValue field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RuleFailureReason) GetReferralValueOk() (string, bool) {
	if o == nil || o.ReferralValue == nil {
		var ret string
		return ret, false
	}
	return *o.ReferralValue, true
}

// HasReferralValue returns a boolean if a field has been set.
func (o *RuleFailureReason) HasReferralValue() bool {
	if o != nil && o.ReferralValue != nil {
		return true
	}

	return false
}

// SetReferralValue gets a reference to the given string and assigns it to the ReferralValue field.
func (o *RuleFailureReason) SetReferralValue(v string) {
	o.ReferralValue = &v
}

// GetRuleIndex returns the RuleIndex field value
func (o *RuleFailureReason) GetRuleIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RuleIndex
}

// SetRuleIndex sets field value
func (o *RuleFailureReason) SetRuleIndex(v int32) {
	o.RuleIndex = v
}

// GetRuleName returns the RuleName field value
func (o *RuleFailureReason) GetRuleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleName
}

// SetRuleName sets field value
func (o *RuleFailureReason) SetRuleName(v string) {
	o.RuleName = v
}

// GetConditionIndex returns the ConditionIndex field value if set, zero value otherwise.
func (o *RuleFailureReason) GetConditionIndex() int32 {
	if o == nil || o.ConditionIndex == nil {
		var ret int32
		return ret
	}
	return *o.ConditionIndex
}

// GetConditionIndexOk returns a tuple with the ConditionIndex field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RuleFailureReason) GetConditionIndexOk() (int32, bool) {
	if o == nil || o.ConditionIndex == nil {
		var ret int32
		return ret, false
	}
	return *o.ConditionIndex, true
}

// HasConditionIndex returns a boolean if a field has been set.
func (o *RuleFailureReason) HasConditionIndex() bool {
	if o != nil && o.ConditionIndex != nil {
		return true
	}

	return false
}

// SetConditionIndex gets a reference to the given int32 and assigns it to the ConditionIndex field.
func (o *RuleFailureReason) SetConditionIndex(v int32) {
	o.ConditionIndex = &v
}

// GetEffectIndex returns the EffectIndex field value if set, zero value otherwise.
func (o *RuleFailureReason) GetEffectIndex() int32 {
	if o == nil || o.EffectIndex == nil {
		var ret int32
		return ret
	}
	return *o.EffectIndex
}

// GetEffectIndexOk returns a tuple with the EffectIndex field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RuleFailureReason) GetEffectIndexOk() (int32, bool) {
	if o == nil || o.EffectIndex == nil {
		var ret int32
		return ret, false
	}
	return *o.EffectIndex, true
}

// HasEffectIndex returns a boolean if a field has been set.
func (o *RuleFailureReason) HasEffectIndex() bool {
	if o != nil && o.EffectIndex != nil {
		return true
	}

	return false
}

// SetEffectIndex gets a reference to the given int32 and assigns it to the EffectIndex field.
func (o *RuleFailureReason) SetEffectIndex(v int32) {
	o.EffectIndex = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *RuleFailureReason) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RuleFailureReason) GetDetailsOk() (string, bool) {
	if o == nil || o.Details == nil {
		var ret string
		return ret, false
	}
	return *o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *RuleFailureReason) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *RuleFailureReason) SetDetails(v string) {
	o.Details = &v
}

type NullableRuleFailureReason struct {
	Value        RuleFailureReason
	ExplicitNull bool
}

func (v NullableRuleFailureReason) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRuleFailureReason) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
