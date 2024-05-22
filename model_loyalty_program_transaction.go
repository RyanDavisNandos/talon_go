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

// LoyaltyProgramTransaction struct for LoyaltyProgramTransaction
type LoyaltyProgramTransaction struct {
	// ID of the loyalty ledger transaction.
	Id int32 `json:"id"`
	// ID of the loyalty program.
	ProgramId int32 `json:"programId"`
	// ID of the campaign.
	CampaignId *int32 `json:"campaignId,omitempty"`
	// Date and time the loyalty transaction occurred.
	Created time.Time `json:"created"`
	// Type of transaction. Possible values:   - `addition`: Signifies added points.   - `subtraction`: Signifies deducted points.
	Type string `json:"type"`
	// Amount of loyalty points added or deducted in the transaction.
	Amount float32 `json:"amount"`
	// Name or reason for the loyalty ledger transaction.
	Name string `json:"name"`
	// When points become active. Possible values:   - `immediate`: Points are immediately active.   - a timestamp value: Points become active at a given date and time.
	StartDate string `json:"startDate"`
	// When points expire. Possible values:   - `unlimited`: Points have no expiration date.   - a timestamp value: Points expire at a given date and time.
	ExpiryDate string `json:"expiryDate"`
	// Customer profile integration ID used in the loyalty program.
	CustomerProfileId *string `json:"customerProfileId,omitempty"`
	// The alphanumeric identifier of the loyalty card.
	CardIdentifier *string `json:"cardIdentifier,omitempty"`
	// ID of the subledger.
	SubledgerId string `json:"subledgerId"`
	// ID of the customer session where the transaction occurred.
	CustomerSessionId *string `json:"customerSessionId,omitempty"`
	// ID of the import where the transaction occurred.
	ImportId *int32 `json:"importId,omitempty"`
	// ID of the user who manually added or deducted points. Applies only to manual transactions.
	UserId *int32 `json:"userId,omitempty"`
	// The email of the Campaign Manager account that manually added or deducted points. Applies only to manual transactions.
	UserEmail *string `json:"userEmail,omitempty"`
	// ID of the ruleset containing the rule that triggered the effect. Applies only for transactions that resulted from a customer session.
	RulesetId *int32 `json:"rulesetId,omitempty"`
	// Name of the rule that triggered the effect. Applies only for transactions that resulted from a customer session.
	RuleName *string `json:"ruleName,omitempty"`
}

// GetId returns the Id field value
func (o *LoyaltyProgramTransaction) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *LoyaltyProgramTransaction) SetId(v int32) {
	o.Id = v
}

// GetProgramId returns the ProgramId field value
func (o *LoyaltyProgramTransaction) GetProgramId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProgramId
}

// SetProgramId sets field value
func (o *LoyaltyProgramTransaction) SetProgramId(v int32) {
	o.ProgramId = v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *LoyaltyProgramTransaction) GetCampaignId() int32 {
	if o == nil || o.CampaignId == nil {
		var ret int32
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyProgramTransaction) GetCampaignIdOk() (int32, bool) {
	if o == nil || o.CampaignId == nil {
		var ret int32
		return ret, false
	}
	return *o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *LoyaltyProgramTransaction) HasCampaignId() bool {
	if o != nil && o.CampaignId != nil {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.
func (o *LoyaltyProgramTransaction) SetCampaignId(v int32) {
	o.CampaignId = &v
}

// GetCreated returns the Created field value
func (o *LoyaltyProgramTransaction) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *LoyaltyProgramTransaction) SetCreated(v time.Time) {
	o.Created = v
}

// GetType returns the Type field value
func (o *LoyaltyProgramTransaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *LoyaltyProgramTransaction) SetType(v string) {
	o.Type = v
}

// GetAmount returns the Amount field value
func (o *LoyaltyProgramTransaction) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// SetAmount sets field value
func (o *LoyaltyProgramTransaction) SetAmount(v float32) {
	o.Amount = v
}

// GetName returns the Name field value
func (o *LoyaltyProgramTransaction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *LoyaltyProgramTransaction) SetName(v string) {
	o.Name = v
}

// GetStartDate returns the StartDate field value
func (o *LoyaltyProgramTransaction) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// SetStartDate sets field value
func (o *LoyaltyProgramTransaction) SetStartDate(v string) {
	o.StartDate = v
}

// GetExpiryDate returns the ExpiryDate field value
func (o *LoyaltyProgramTransaction) GetExpiryDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiryDate
}

// SetExpiryDate sets field value
func (o *LoyaltyProgramTransaction) SetExpiryDate(v string) {
	o.ExpiryDate = v
}

// GetCustomerProfileId returns the CustomerProfileId field value if set, zero value otherwise.
func (o *LoyaltyProgramTransaction) GetCustomerProfileId() string {
	if o == nil || o.CustomerProfileId == nil {
		var ret string
		return ret
	}
	return *o.CustomerProfileId
}

// GetCustomerProfileIdOk returns a tuple with the CustomerProfileId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyProgramTransaction) GetCustomerProfileIdOk() (string, bool) {
	if o == nil || o.CustomerProfileId == nil {
		var ret string
		return ret, false
	}
	return *o.CustomerProfileId, true
}

// HasCustomerProfileId returns a boolean if a field has been set.
func (o *LoyaltyProgramTransaction) HasCustomerProfileId() bool {
	if o != nil && o.CustomerProfileId != nil {
		return true
	}

	return false
}

// SetCustomerProfileId gets a reference to the given string and assigns it to the CustomerProfileId field.
func (o *LoyaltyProgramTransaction) SetCustomerProfileId(v string) {
	o.CustomerProfileId = &v
}

// GetCardIdentifier returns the CardIdentifier field value if set, zero value otherwise.
func (o *LoyaltyProgramTransaction) GetCardIdentifier() string {
	if o == nil || o.CardIdentifier == nil {
		var ret string
		return ret
	}
	return *o.CardIdentifier
}

// GetCardIdentifierOk returns a tuple with the CardIdentifier field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyProgramTransaction) GetCardIdentifierOk() (string, bool) {
	if o == nil || o.CardIdentifier == nil {
		var ret string
		return ret, false
	}
	return *o.CardIdentifier, true
}

// HasCardIdentifier returns a boolean if a field has been set.
func (o *LoyaltyProgramTransaction) HasCardIdentifier() bool {
	if o != nil && o.CardIdentifier != nil {
		return true
	}

	return false
}

// SetCardIdentifier gets a reference to the given string and assigns it to the CardIdentifier field.
func (o *LoyaltyProgramTransaction) SetCardIdentifier(v string) {
	o.CardIdentifier = &v
}

// GetSubledgerId returns the SubledgerId field value
func (o *LoyaltyProgramTransaction) GetSubledgerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubledgerId
}

// SetSubledgerId sets field value
func (o *LoyaltyProgramTransaction) SetSubledgerId(v string) {
	o.SubledgerId = v
}

// GetCustomerSessionId returns the CustomerSessionId field value if set, zero value otherwise.
func (o *LoyaltyProgramTransaction) GetCustomerSessionId() string {
	if o == nil || o.CustomerSessionId == nil {
		var ret string
		return ret
	}
	return *o.CustomerSessionId
}

// GetCustomerSessionIdOk returns a tuple with the CustomerSessionId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyProgramTransaction) GetCustomerSessionIdOk() (string, bool) {
	if o == nil || o.CustomerSessionId == nil {
		var ret string
		return ret, false
	}
	return *o.CustomerSessionId, true
}

// HasCustomerSessionId returns a boolean if a field has been set.
func (o *LoyaltyProgramTransaction) HasCustomerSessionId() bool {
	if o != nil && o.CustomerSessionId != nil {
		return true
	}

	return false
}

// SetCustomerSessionId gets a reference to the given string and assigns it to the CustomerSessionId field.
func (o *LoyaltyProgramTransaction) SetCustomerSessionId(v string) {
	o.CustomerSessionId = &v
}

// GetImportId returns the ImportId field value if set, zero value otherwise.
func (o *LoyaltyProgramTransaction) GetImportId() int32 {
	if o == nil || o.ImportId == nil {
		var ret int32
		return ret
	}
	return *o.ImportId
}

// GetImportIdOk returns a tuple with the ImportId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyProgramTransaction) GetImportIdOk() (int32, bool) {
	if o == nil || o.ImportId == nil {
		var ret int32
		return ret, false
	}
	return *o.ImportId, true
}

// HasImportId returns a boolean if a field has been set.
func (o *LoyaltyProgramTransaction) HasImportId() bool {
	if o != nil && o.ImportId != nil {
		return true
	}

	return false
}

// SetImportId gets a reference to the given int32 and assigns it to the ImportId field.
func (o *LoyaltyProgramTransaction) SetImportId(v int32) {
	o.ImportId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *LoyaltyProgramTransaction) GetUserId() int32 {
	if o == nil || o.UserId == nil {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyProgramTransaction) GetUserIdOk() (int32, bool) {
	if o == nil || o.UserId == nil {
		var ret int32
		return ret, false
	}
	return *o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *LoyaltyProgramTransaction) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *LoyaltyProgramTransaction) SetUserId(v int32) {
	o.UserId = &v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise.
func (o *LoyaltyProgramTransaction) GetUserEmail() string {
	if o == nil || o.UserEmail == nil {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyProgramTransaction) GetUserEmailOk() (string, bool) {
	if o == nil || o.UserEmail == nil {
		var ret string
		return ret, false
	}
	return *o.UserEmail, true
}

// HasUserEmail returns a boolean if a field has been set.
func (o *LoyaltyProgramTransaction) HasUserEmail() bool {
	if o != nil && o.UserEmail != nil {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *LoyaltyProgramTransaction) SetUserEmail(v string) {
	o.UserEmail = &v
}

// GetRulesetId returns the RulesetId field value if set, zero value otherwise.
func (o *LoyaltyProgramTransaction) GetRulesetId() int32 {
	if o == nil || o.RulesetId == nil {
		var ret int32
		return ret
	}
	return *o.RulesetId
}

// GetRulesetIdOk returns a tuple with the RulesetId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyProgramTransaction) GetRulesetIdOk() (int32, bool) {
	if o == nil || o.RulesetId == nil {
		var ret int32
		return ret, false
	}
	return *o.RulesetId, true
}

// HasRulesetId returns a boolean if a field has been set.
func (o *LoyaltyProgramTransaction) HasRulesetId() bool {
	if o != nil && o.RulesetId != nil {
		return true
	}

	return false
}

// SetRulesetId gets a reference to the given int32 and assigns it to the RulesetId field.
func (o *LoyaltyProgramTransaction) SetRulesetId(v int32) {
	o.RulesetId = &v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *LoyaltyProgramTransaction) GetRuleName() string {
	if o == nil || o.RuleName == nil {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyProgramTransaction) GetRuleNameOk() (string, bool) {
	if o == nil || o.RuleName == nil {
		var ret string
		return ret, false
	}
	return *o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *LoyaltyProgramTransaction) HasRuleName() bool {
	if o != nil && o.RuleName != nil {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *LoyaltyProgramTransaction) SetRuleName(v string) {
	o.RuleName = &v
}

type NullableLoyaltyProgramTransaction struct {
	Value        LoyaltyProgramTransaction
	ExplicitNull bool
}

func (v NullableLoyaltyProgramTransaction) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLoyaltyProgramTransaction) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
