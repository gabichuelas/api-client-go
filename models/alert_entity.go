// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AlertEntity Assign an alert a primary status
//
// swagger:model AlertEntity
type AlertEntity struct {

	// alert
	Alert *AlertEntity `json:"alert,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// whether or not this is the primary alert for this incident
	Primary bool `json:"primary,omitempty"`
}

// Validate validates this alert entity
func (m *AlertEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlert(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertEntity) validateAlert(formats strfmt.Registry) error {
	if swag.IsZero(m.Alert) { // not required
		return nil
	}

	if m.Alert != nil {
		if err := m.Alert.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alert")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this alert entity based on the context it is used
func (m *AlertEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlert(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertEntity) contextValidateAlert(ctx context.Context, formats strfmt.Registry) error {

	if m.Alert != nil {
		if err := m.Alert.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alert")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertEntity) UnmarshalBinary(b []byte) error {
	var res AlertEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}