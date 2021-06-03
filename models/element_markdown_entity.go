// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ElementMarkdownEntity element markdown entity
//
// swagger:model ElementMarkdownEntity
type ElementMarkdownEntity struct {

	// text
	Text string `json:"text,omitempty"`
}

// Validate validates this element markdown entity
func (m *ElementMarkdownEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this element markdown entity based on context it is used
func (m *ElementMarkdownEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElementMarkdownEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElementMarkdownEntity) UnmarshalBinary(b []byte) error {
	var res ElementMarkdownEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}