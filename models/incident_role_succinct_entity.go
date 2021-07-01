// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IncidentRoleSuccinctEntity incident role succinct entity
//
// swagger:model IncidentRoleSuccinctEntity
type IncidentRoleSuccinctEntity struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// discarded at
	DiscardedAt string `json:"discarded_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this incident role succinct entity
func (m *IncidentRoleSuccinctEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this incident role succinct entity based on context it is used
func (m *IncidentRoleSuccinctEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IncidentRoleSuccinctEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IncidentRoleSuccinctEntity) UnmarshalBinary(b []byte) error {
	var res IncidentRoleSuccinctEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
