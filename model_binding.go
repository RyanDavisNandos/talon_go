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

// Binding struct for Binding
type Binding struct {
	// A descriptive name for the value to be bound.
	Name string `json:"name"`
	// The kind of binding. Possible values are: - `bundle` - `cartItemFilter` - `subledgerBalance` - `templateParameter`
	Type *string `json:"type,omitempty"`
	// A Talang expression that will be evaluated and its result attached to the name of the binding.
	Expression []interface{} `json:"expression"`
	// Can be one of the following: - `string` - `number` - `boolean`
	ValueType *string `json:"valueType,omitempty"`
}

// GetName returns the Name field value
func (o *Binding) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *Binding) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Binding) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Binding) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Binding) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Binding) SetType(v string) {
	o.Type = &v
}

// GetExpression returns the Expression field value
func (o *Binding) GetExpression() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Expression
}

// SetExpression sets field value
func (o *Binding) SetExpression(v []interface{}) {
	o.Expression = v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *Binding) GetValueType() string {
	if o == nil || o.ValueType == nil {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Binding) GetValueTypeOk() (string, bool) {
	if o == nil || o.ValueType == nil {
		var ret string
		return ret, false
	}
	return *o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *Binding) HasValueType() bool {
	if o != nil && o.ValueType != nil {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *Binding) SetValueType(v string) {
	o.ValueType = &v
}

type NullableBinding struct {
	Value        Binding
	ExplicitNull bool
}

func (v NullableBinding) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableBinding) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
