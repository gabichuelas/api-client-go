// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PutV1IncidentsIncidentIDImpact put v1 incidents incident Id impact
//
// swagger:model putV1IncidentsIncidentIdImpact
type PutV1IncidentsIncidentIDImpact struct {

	// impact
	Impact []*PutV1IncidentsIncidentIDImpactImpactItems0 `json:"impact"`

	// milestone
	Milestone string `json:"milestone,omitempty"`

	// note
	Note string `json:"note,omitempty"`

	// status pages
	StatusPages []*PutV1IncidentsIncidentIDImpactStatusPagesItems0 `json:"status_pages"`
}

// Validate validates this put v1 incidents incident Id impact
func (m *PutV1IncidentsIncidentIDImpact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImpact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusPages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1IncidentsIncidentIDImpact) validateImpact(formats strfmt.Registry) error {
	if swag.IsZero(m.Impact) { // not required
		return nil
	}

	for i := 0; i < len(m.Impact); i++ {
		if swag.IsZero(m.Impact[i]) { // not required
			continue
		}

		if m.Impact[i] != nil {
			if err := m.Impact[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("impact" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PutV1IncidentsIncidentIDImpact) validateStatusPages(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusPages) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusPages); i++ {
		if swag.IsZero(m.StatusPages[i]) { // not required
			continue
		}

		if m.StatusPages[i] != nil {
			if err := m.StatusPages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_pages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this put v1 incidents incident Id impact based on the context it is used
func (m *PutV1IncidentsIncidentIDImpact) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImpact(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusPages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1IncidentsIncidentIDImpact) contextValidateImpact(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Impact); i++ {

		if m.Impact[i] != nil {
			if err := m.Impact[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("impact" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PutV1IncidentsIncidentIDImpact) contextValidateStatusPages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusPages); i++ {

		if m.StatusPages[i] != nil {
			if err := m.StatusPages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_pages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutV1IncidentsIncidentIDImpact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1IncidentsIncidentIDImpact) UnmarshalBinary(b []byte) error {
	var res PutV1IncidentsIncidentIDImpact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutV1IncidentsIncidentIDImpactImpactItems0 put v1 incidents incident ID impact impact items0
//
// swagger:model PutV1IncidentsIncidentIDImpactImpactItems0
type PutV1IncidentsIncidentIDImpactImpactItems0 struct {

	// condition id
	// Required: true
	ConditionID *string `json:"condition_id"`

	// id
	// Required: true
	ID *string `json:"id"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this put v1 incidents incident ID impact impact items0
func (m *PutV1IncidentsIncidentIDImpactImpactItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1IncidentsIncidentIDImpactImpactItems0) validateConditionID(formats strfmt.Registry) error {

	if err := validate.Required("condition_id", "body", m.ConditionID); err != nil {
		return err
	}

	return nil
}

func (m *PutV1IncidentsIncidentIDImpactImpactItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PutV1IncidentsIncidentIDImpactImpactItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put v1 incidents incident ID impact impact items0 based on context it is used
func (m *PutV1IncidentsIncidentIDImpactImpactItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutV1IncidentsIncidentIDImpactImpactItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1IncidentsIncidentIDImpactImpactItems0) UnmarshalBinary(b []byte) error {
	var res PutV1IncidentsIncidentIDImpactImpactItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutV1IncidentsIncidentIDImpactStatusPagesItems0 put v1 incidents incident ID impact status pages items0
//
// swagger:model PutV1IncidentsIncidentIDImpactStatusPagesItems0
type PutV1IncidentsIncidentIDImpactStatusPagesItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// integration slug
	// Required: true
	IntegrationSlug *string `json:"integration_slug"`
}

// Validate validates this put v1 incidents incident ID impact status pages items0
func (m *PutV1IncidentsIncidentIDImpactStatusPagesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegrationSlug(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1IncidentsIncidentIDImpactStatusPagesItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PutV1IncidentsIncidentIDImpactStatusPagesItems0) validateIntegrationSlug(formats strfmt.Registry) error {

	if err := validate.Required("integration_slug", "body", m.IntegrationSlug); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put v1 incidents incident ID impact status pages items0 based on context it is used
func (m *PutV1IncidentsIncidentIDImpactStatusPagesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutV1IncidentsIncidentIDImpactStatusPagesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1IncidentsIncidentIDImpactStatusPagesItems0) UnmarshalBinary(b []byte) error {
	var res PutV1IncidentsIncidentIDImpactStatusPagesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
