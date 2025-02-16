// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchV1IncidentRoles Update an incident role
//
// swagger:model patchV1IncidentRoles
type PatchV1IncidentRoles struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`
}

// Validate validates this patch v1 incident roles
func (m *PatchV1IncidentRoles) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch v1 incident roles based on context it is used
func (m *PatchV1IncidentRoles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1IncidentRoles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1IncidentRoles) UnmarshalBinary(b []byte) error {
	var res PatchV1IncidentRoles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
