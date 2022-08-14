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

// RedeemReferralEffectProps This effect is **deprecated**. The properties specific to the \"redeemReferral\" effect. This gets triggered whenever the referral code is valid, and a rule was triggered that contains a \"redeem referral\" effect.
type RedeemReferralEffectProps struct {
	// The id of the referral code that was redeemed.
	Id int32 `json:"id"`
	// The referral code that was redeemed.
	Value string `json:"value"`
}

// GetId returns the Id field value
func (o *RedeemReferralEffectProps) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *RedeemReferralEffectProps) SetId(v int32) {
	o.Id = v
}

// GetValue returns the Value field value
func (o *RedeemReferralEffectProps) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *RedeemReferralEffectProps) SetValue(v string) {
	o.Value = v
}

type NullableRedeemReferralEffectProps struct {
	Value        RedeemReferralEffectProps
	ExplicitNull bool
}

func (v NullableRedeemReferralEffectProps) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRedeemReferralEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
