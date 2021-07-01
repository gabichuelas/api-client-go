// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SeverityEntity Retrieve a specific severity
//
// swagger:model SeverityEntity
type SeverityEntity struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// position
	Position string `json:"position,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// system record
	SystemRecord string `json:"system_record,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this severity entity
func (m *SeverityEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this severity entity based on context it is used
func (m *SeverityEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SeverityEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SeverityEntity) UnmarshalBinary(b []byte) error {
	var res SeverityEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
