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

// BaseLoyaltyProgram struct for BaseLoyaltyProgram
type BaseLoyaltyProgram struct {
	// The display title for the Loyalty Program.
	Title *string `json:"title,omitempty"`
	// Description of our Loyalty Program.
	Description *string `json:"description,omitempty"`
	// A list containing the IDs of all applications that are subscribed to this Loyalty Program.
	SubscribedApplications *[]int32 `json:"subscribedApplications,omitempty"`
	// Indicates the default duration after which new loyalty points should expire. The format is a number, followed by one letter indicating the unit; like '1h' or '40m'.
	DefaultValidity *string `json:"defaultValidity,omitempty"`
	// Indicates the default duration for the pending time, after which points will be valid. The format is a number followed by a duration unit, like '1h' or '40m'.
	DefaultPending *string `json:"defaultPending,omitempty"`
	// Indicates if this program supports subledgers inside the program.
	AllowSubledger *bool `json:"allowSubledger,omitempty"`
	// The max amount of user profiles with whom a card can be shared. This can be set to 0 for no limit. This property is only used when `cardBased` is `true`.
	UsersPerCardLimit *int32 `json:"usersPerCardLimit,omitempty"`
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BaseLoyaltyProgram) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseLoyaltyProgram) GetTitleOk() (string, bool) {
	if o == nil || o.Title == nil {
		var ret string
		return ret, false
	}
	return *o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BaseLoyaltyProgram) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BaseLoyaltyProgram) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseLoyaltyProgram) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseLoyaltyProgram) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseLoyaltyProgram) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseLoyaltyProgram) SetDescription(v string) {
	o.Description = &v
}

// GetSubscribedApplications returns the SubscribedApplications field value if set, zero value otherwise.
func (o *BaseLoyaltyProgram) GetSubscribedApplications() []int32 {
	if o == nil || o.SubscribedApplications == nil {
		var ret []int32
		return ret
	}
	return *o.SubscribedApplications
}

// GetSubscribedApplicationsOk returns a tuple with the SubscribedApplications field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseLoyaltyProgram) GetSubscribedApplicationsOk() ([]int32, bool) {
	if o == nil || o.SubscribedApplications == nil {
		var ret []int32
		return ret, false
	}
	return *o.SubscribedApplications, true
}

// HasSubscribedApplications returns a boolean if a field has been set.
func (o *BaseLoyaltyProgram) HasSubscribedApplications() bool {
	if o != nil && o.SubscribedApplications != nil {
		return true
	}

	return false
}

// SetSubscribedApplications gets a reference to the given []int32 and assigns it to the SubscribedApplications field.
func (o *BaseLoyaltyProgram) SetSubscribedApplications(v []int32) {
	o.SubscribedApplications = &v
}

// GetDefaultValidity returns the DefaultValidity field value if set, zero value otherwise.
func (o *BaseLoyaltyProgram) GetDefaultValidity() string {
	if o == nil || o.DefaultValidity == nil {
		var ret string
		return ret
	}
	return *o.DefaultValidity
}

// GetDefaultValidityOk returns a tuple with the DefaultValidity field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseLoyaltyProgram) GetDefaultValidityOk() (string, bool) {
	if o == nil || o.DefaultValidity == nil {
		var ret string
		return ret, false
	}
	return *o.DefaultValidity, true
}

// HasDefaultValidity returns a boolean if a field has been set.
func (o *BaseLoyaltyProgram) HasDefaultValidity() bool {
	if o != nil && o.DefaultValidity != nil {
		return true
	}

	return false
}

// SetDefaultValidity gets a reference to the given string and assigns it to the DefaultValidity field.
func (o *BaseLoyaltyProgram) SetDefaultValidity(v string) {
	o.DefaultValidity = &v
}

// GetDefaultPending returns the DefaultPending field value if set, zero value otherwise.
func (o *BaseLoyaltyProgram) GetDefaultPending() string {
	if o == nil || o.DefaultPending == nil {
		var ret string
		return ret
	}
	return *o.DefaultPending
}

// GetDefaultPendingOk returns a tuple with the DefaultPending field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseLoyaltyProgram) GetDefaultPendingOk() (string, bool) {
	if o == nil || o.DefaultPending == nil {
		var ret string
		return ret, false
	}
	return *o.DefaultPending, true
}

// HasDefaultPending returns a boolean if a field has been set.
func (o *BaseLoyaltyProgram) HasDefaultPending() bool {
	if o != nil && o.DefaultPending != nil {
		return true
	}

	return false
}

// SetDefaultPending gets a reference to the given string and assigns it to the DefaultPending field.
func (o *BaseLoyaltyProgram) SetDefaultPending(v string) {
	o.DefaultPending = &v
}

// GetAllowSubledger returns the AllowSubledger field value if set, zero value otherwise.
func (o *BaseLoyaltyProgram) GetAllowSubledger() bool {
	if o == nil || o.AllowSubledger == nil {
		var ret bool
		return ret
	}
	return *o.AllowSubledger
}

// GetAllowSubledgerOk returns a tuple with the AllowSubledger field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseLoyaltyProgram) GetAllowSubledgerOk() (bool, bool) {
	if o == nil || o.AllowSubledger == nil {
		var ret bool
		return ret, false
	}
	return *o.AllowSubledger, true
}

// HasAllowSubledger returns a boolean if a field has been set.
func (o *BaseLoyaltyProgram) HasAllowSubledger() bool {
	if o != nil && o.AllowSubledger != nil {
		return true
	}

	return false
}

// SetAllowSubledger gets a reference to the given bool and assigns it to the AllowSubledger field.
func (o *BaseLoyaltyProgram) SetAllowSubledger(v bool) {
	o.AllowSubledger = &v
}

// GetUsersPerCardLimit returns the UsersPerCardLimit field value if set, zero value otherwise.
func (o *BaseLoyaltyProgram) GetUsersPerCardLimit() int32 {
	if o == nil || o.UsersPerCardLimit == nil {
		var ret int32
		return ret
	}
	return *o.UsersPerCardLimit
}

// GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseLoyaltyProgram) GetUsersPerCardLimitOk() (int32, bool) {
	if o == nil || o.UsersPerCardLimit == nil {
		var ret int32
		return ret, false
	}
	return *o.UsersPerCardLimit, true
}

// HasUsersPerCardLimit returns a boolean if a field has been set.
func (o *BaseLoyaltyProgram) HasUsersPerCardLimit() bool {
	if o != nil && o.UsersPerCardLimit != nil {
		return true
	}

	return false
}

// SetUsersPerCardLimit gets a reference to the given int32 and assigns it to the UsersPerCardLimit field.
func (o *BaseLoyaltyProgram) SetUsersPerCardLimit(v int32) {
	o.UsersPerCardLimit = &v
}

type NullableBaseLoyaltyProgram struct {
	Value        BaseLoyaltyProgram
	ExplicitNull bool
}

func (v NullableBaseLoyaltyProgram) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableBaseLoyaltyProgram) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
