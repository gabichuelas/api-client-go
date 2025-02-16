// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostV1Runbooks Create a new runbook for use with incidents.
//
// swagger:model postV1Runbooks
type PostV1Runbooks struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// summary
	Summary string `json:"summary,omitempty"`

	// type
	// Enum: [incident general infrastructure incident_role]
	Type string `json:"type,omitempty"`
}

// Validate validates this post v1 runbooks
func (m *PostV1Runbooks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
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

func (m *PostV1Runbooks) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var postV1RunbooksTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["incident","general","infrastructure","incident_role"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postV1RunbooksTypeTypePropEnum = append(postV1RunbooksTypeTypePropEnum, v)
	}
}

const (

	// PostV1RunbooksTypeIncident captures enum value "incident"
	PostV1RunbooksTypeIncident string = "incident"

	// PostV1RunbooksTypeGeneral captures enum value "general"
	PostV1RunbooksTypeGeneral string = "general"

	// PostV1RunbooksTypeInfrastructure captures enum value "infrastructure"
	PostV1RunbooksTypeInfrastructure string = "infrastructure"

	// PostV1RunbooksTypeIncidentRole captures enum value "incident_role"
	PostV1RunbooksTypeIncidentRole string = "incident_role"
)

// prop value enum
func (m *PostV1Runbooks) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postV1RunbooksTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostV1Runbooks) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 runbooks based on context it is used
func (m *PostV1Runbooks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1Runbooks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1Runbooks) UnmarshalBinary(b []byte) error {
	var res PostV1Runbooks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
