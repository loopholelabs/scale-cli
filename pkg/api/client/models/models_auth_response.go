// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelsAuthResponse models auth response
//
// swagger:model models.AuthResponse
type ModelsAuthResponse struct {

	// auth url
	AuthURL string `json:"auth_url,omitempty"`

	// expiry
	Expiry int64 `json:"expiry,omitempty"`
}

// Validate validates this models auth response
func (m *ModelsAuthResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this models auth response based on context it is used
func (m *ModelsAuthResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsAuthResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsAuthResponse) UnmarshalBinary(b []byte) error {
	var res ModelsAuthResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
