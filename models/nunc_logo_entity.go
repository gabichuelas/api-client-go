// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NuncLogoEntity nunc logo entity
//
// swagger:model NuncLogoEntity
type NuncLogoEntity struct {

	// original url
	OriginalURL string `json:"original_url,omitempty"`

	// versions urls
	VersionsUrls string `json:"versions_urls,omitempty"`
}

// Validate validates this nunc logo entity
func (m *NuncLogoEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nunc logo entity based on context it is used
func (m *NuncLogoEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NuncLogoEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NuncLogoEntity) UnmarshalBinary(b []byte) error {
	var res NuncLogoEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
