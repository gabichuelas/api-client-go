// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImpactEntity Update a specific impact
//
// swagger:model ImpactEntity
type ImpactEntity struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// position
	Position string `json:"position,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this impact entity
func (m *ImpactEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this impact entity based on context it is used
func (m *ImpactEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImpactEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImpactEntity) UnmarshalBinary(b []byte) error {
	var res ImpactEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
