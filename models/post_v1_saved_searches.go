// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostV1SavedSearches Create a new saved search for a particular resource type
//
// swagger:model postV1SavedSearches
type PostV1SavedSearches struct {

	// filter values
	// Required: true
	FilterValues interface{} `json:"filter_values"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this post v1 saved searches
func (m *PostV1SavedSearches) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilterValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1SavedSearches) validateFilterValues(formats strfmt.Registry) error {

	if m.FilterValues == nil {
		return errors.Required("filter_values", "body", nil)
	}

	return nil
}

func (m *PostV1SavedSearches) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 saved searches based on context it is used
func (m *PostV1SavedSearches) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1SavedSearches) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1SavedSearches) UnmarshalBinary(b []byte) error {
	var res PostV1SavedSearches
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
