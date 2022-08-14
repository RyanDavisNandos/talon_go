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

// LoyaltyBalances List of loyalty balances for a ledger and its subledgers.
type LoyaltyBalances struct {
	Balance *LoyaltyBalance `json:"balance,omitempty"`
	// Map of the loyalty balances of the subledgers of a ledger.
	SubledgerBalances *map[string]LoyaltyBalance `json:"subledgerBalances,omitempty"`
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *LoyaltyBalances) GetBalance() LoyaltyBalance {
	if o == nil || o.Balance == nil {
		var ret LoyaltyBalance
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalances) GetBalanceOk() (LoyaltyBalance, bool) {
	if o == nil || o.Balance == nil {
		var ret LoyaltyBalance
		return ret, false
	}
	return *o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *LoyaltyBalances) HasBalance() bool {
	if o != nil && o.Balance != nil {
		return true
	}

	return false
}

// SetBalance gets a reference to the given LoyaltyBalance and assigns it to the Balance field.
func (o *LoyaltyBalances) SetBalance(v LoyaltyBalance) {
	o.Balance = &v
}

// GetSubledgerBalances returns the SubledgerBalances field value if set, zero value otherwise.
func (o *LoyaltyBalances) GetSubledgerBalances() map[string]LoyaltyBalance {
	if o == nil || o.SubledgerBalances == nil {
		var ret map[string]LoyaltyBalance
		return ret
	}
	return *o.SubledgerBalances
}

// GetSubledgerBalancesOk returns a tuple with the SubledgerBalances field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalances) GetSubledgerBalancesOk() (map[string]LoyaltyBalance, bool) {
	if o == nil || o.SubledgerBalances == nil {
		var ret map[string]LoyaltyBalance
		return ret, false
	}
	return *o.SubledgerBalances, true
}

// HasSubledgerBalances returns a boolean if a field has been set.
func (o *LoyaltyBalances) HasSubledgerBalances() bool {
	if o != nil && o.SubledgerBalances != nil {
		return true
	}

	return false
}

// SetSubledgerBalances gets a reference to the given map[string]LoyaltyBalance and assigns it to the SubledgerBalances field.
func (o *LoyaltyBalances) SetSubledgerBalances(v map[string]LoyaltyBalance) {
	o.SubledgerBalances = &v
}

type NullableLoyaltyBalances struct {
	Value        LoyaltyBalances
	ExplicitNull bool
}

func (v NullableLoyaltyBalances) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLoyaltyBalances) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
