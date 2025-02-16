// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReleaseNoteEntity Retrieve a single release note set
//
// swagger:model ReleaseNoteEntity
type ReleaseNoteEntity struct {

	// note
	Note string `json:"note,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`
}

// Validate validates this release note entity
func (m *ReleaseNoteEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this release note entity based on context it is used
func (m *ReleaseNoteEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseNoteEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseNoteEntity) UnmarshalBinary(b []byte) error {
	var res ReleaseNoteEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
