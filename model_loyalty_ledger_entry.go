/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put
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

// LoyaltyLedgerEntry A single row of the ledger, describing one addition or deduction.
type LoyaltyLedgerEntry struct {
	Created           time.Time `json:"created"`
	ProgramID         int32     `json:"programID"`
	CustomerProfileID string    `json:"customerProfileID"`
	CustomerSessionID *string   `json:"customerSessionID,omitempty"`
	EventID           *int32    `json:"eventID,omitempty"`
	// The type of the ledger transaction. Possible values are addition, subtraction, expire or expiring (for expiring points ledgers)
	Type       string     `json:"type"`
	Amount     float32    `json:"amount"`
	StartDate  *time.Time `json:"startDate,omitempty"`
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// A name referencing the condition or effect that added this entry, or the specific name provided in an API call.
	Name string `json:"name"`
	// This specifies if we are adding loyalty points to the main ledger or a subledger
	SubLedgerID string `json:"subLedgerID"`
	// This is the ID of the user who created this entry, if the addition or subtraction was done manually.
	UserID *int32 `json:"userID,omitempty"`
}

// GetCreated returns the Created field value
func (o *LoyaltyLedgerEntry) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *LoyaltyLedgerEntry) SetCreated(v time.Time) {
	o.Created = v
}

// GetProgramID returns the ProgramID field value
func (o *LoyaltyLedgerEntry) GetProgramID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProgramID
}

// SetProgramID sets field value
func (o *LoyaltyLedgerEntry) SetProgramID(v int32) {
	o.ProgramID = v
}

// GetCustomerProfileID returns the CustomerProfileID field value
func (o *LoyaltyLedgerEntry) GetCustomerProfileID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerProfileID
}

// SetCustomerProfileID sets field value
func (o *LoyaltyLedgerEntry) SetCustomerProfileID(v string) {
	o.CustomerProfileID = v
}

// GetCustomerSessionID returns the CustomerSessionID field value if set, zero value otherwise.
func (o *LoyaltyLedgerEntry) GetCustomerSessionID() string {
	if o == nil || o.CustomerSessionID == nil {
		var ret string
		return ret
	}
	return *o.CustomerSessionID
}

// GetCustomerSessionIDOk returns a tuple with the CustomerSessionID field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetCustomerSessionIDOk() (string, bool) {
	if o == nil || o.CustomerSessionID == nil {
		var ret string
		return ret, false
	}
	return *o.CustomerSessionID, true
}

// HasCustomerSessionID returns a boolean if a field has been set.
func (o *LoyaltyLedgerEntry) HasCustomerSessionID() bool {
	if o != nil && o.CustomerSessionID != nil {
		return true
	}

	return false
}

// SetCustomerSessionID gets a reference to the given string and assigns it to the CustomerSessionID field.
func (o *LoyaltyLedgerEntry) SetCustomerSessionID(v string) {
	o.CustomerSessionID = &v
}

// GetEventID returns the EventID field value if set, zero value otherwise.
func (o *LoyaltyLedgerEntry) GetEventID() int32 {
	if o == nil || o.EventID == nil {
		var ret int32
		return ret
	}
	return *o.EventID
}

// GetEventIDOk returns a tuple with the EventID field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetEventIDOk() (int32, bool) {
	if o == nil || o.EventID == nil {
		var ret int32
		return ret, false
	}
	return *o.EventID, true
}

// HasEventID returns a boolean if a field has been set.
func (o *LoyaltyLedgerEntry) HasEventID() bool {
	if o != nil && o.EventID != nil {
		return true
	}

	return false
}

// SetEventID gets a reference to the given int32 and assigns it to the EventID field.
func (o *LoyaltyLedgerEntry) SetEventID(v int32) {
	o.EventID = &v
}

// GetType returns the Type field value
func (o *LoyaltyLedgerEntry) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *LoyaltyLedgerEntry) SetType(v string) {
	o.Type = v
}

// GetAmount returns the Amount field value
func (o *LoyaltyLedgerEntry) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// SetAmount sets field value
func (o *LoyaltyLedgerEntry) SetAmount(v float32) {
	o.Amount = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *LoyaltyLedgerEntry) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetStartDateOk() (time.Time, bool) {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *LoyaltyLedgerEntry) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *LoyaltyLedgerEntry) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *LoyaltyLedgerEntry) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetExpiryDateOk() (time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *LoyaltyLedgerEntry) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *LoyaltyLedgerEntry) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetName returns the Name field value
func (o *LoyaltyLedgerEntry) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *LoyaltyLedgerEntry) SetName(v string) {
	o.Name = v
}

// GetSubLedgerID returns the SubLedgerID field value
func (o *LoyaltyLedgerEntry) GetSubLedgerID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubLedgerID
}

// SetSubLedgerID sets field value
func (o *LoyaltyLedgerEntry) SetSubLedgerID(v string) {
	o.SubLedgerID = v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *LoyaltyLedgerEntry) GetUserID() int32 {
	if o == nil || o.UserID == nil {
		var ret int32
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerEntry) GetUserIDOk() (int32, bool) {
	if o == nil || o.UserID == nil {
		var ret int32
		return ret, false
	}
	return *o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *LoyaltyLedgerEntry) HasUserID() bool {
	if o != nil && o.UserID != nil {
		return true
	}

	return false
}

// SetUserID gets a reference to the given int32 and assigns it to the UserID field.
func (o *LoyaltyLedgerEntry) SetUserID(v int32) {
	o.UserID = &v
}

type NullableLoyaltyLedgerEntry struct {
	Value        LoyaltyLedgerEntry
	ExplicitNull bool
}

func (v NullableLoyaltyLedgerEntry) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLoyaltyLedgerEntry) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
