// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchV1TicketingTickets Update a ticket's attributes
//
// swagger:model patchV1TicketingTickets
type PatchV1TicketingTickets struct {

	// description
	Description string `json:"description,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this patch v1 ticketing tickets
func (m *PatchV1TicketingTickets) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch v1 ticketing tickets based on context it is used
func (m *PatchV1TicketingTickets) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1TicketingTickets) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1TicketingTickets) UnmarshalBinary(b []byte) error {
	var res PatchV1TicketingTickets
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
