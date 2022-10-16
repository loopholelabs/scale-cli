// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelsRefreshResponse models refresh response
//
// swagger:model models.RefreshResponse
type ModelsRefreshResponse struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`

	// expires in
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// refresh token
	RefreshToken string `json:"refresh_token,omitempty"`

	// token type
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this models refresh response
func (m *ModelsRefreshResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this models refresh response based on context it is used
func (m *ModelsRefreshResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsRefreshResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsRefreshResponse) UnmarshalBinary(b []byte) error {
	var res ModelsRefreshResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
