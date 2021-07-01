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

// ExecutionEntity Updates the execution's step.
//
// swagger:model ExecutionEntity
type ExecutionEntity struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// executed for
	ExecutedFor string `json:"executed_for,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// runbook
	Runbook *SlimRunbookEntity `json:"runbook,omitempty"`

	// steps
	Steps *ExecutionStepEntity `json:"steps,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this execution entity
func (m *ExecutionEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunbook(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExecutionEntity) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExecutionEntity) validateRunbook(formats strfmt.Registry) error {
	if swag.IsZero(m.Runbook) { // not required
		return nil
	}

	if m.Runbook != nil {
		if err := m.Runbook.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runbook")
			}
			return err
		}
	}

	return nil
}

func (m *ExecutionEntity) validateSteps(formats strfmt.Registry) error {
	if swag.IsZero(m.Steps) { // not required
		return nil
	}

	if m.Steps != nil {
		if err := m.Steps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("steps")
			}
			return err
		}
	}

	return nil
}

func (m *ExecutionEntity) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this execution entity based on the context it is used
func (m *ExecutionEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRunbook(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSteps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExecutionEntity) contextValidateRunbook(ctx context.Context, formats strfmt.Registry) error {

	if m.Runbook != nil {
		if err := m.Runbook.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runbook")
			}
			return err
		}
	}

	return nil
}

func (m *ExecutionEntity) contextValidateSteps(ctx context.Context, formats strfmt.Registry) error {

	if m.Steps != nil {
		if err := m.Steps.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("steps")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExecutionEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExecutionEntity) UnmarshalBinary(b []byte) error {
	var res ExecutionEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
