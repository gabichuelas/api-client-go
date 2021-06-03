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

// MilestoneEntity milestone entity
//
// swagger:model MilestoneEntity
type MilestoneEntity struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// duration
	Duration string `json:"duration,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// occurred at
	// Format: date-time
	OccurredAt strfmt.DateTime `json:"occurred_at,omitempty"`

	// type
	// Enum: [started detected acknowledged first_action mitigated resolved postmortem_started postmortem_completed closed]
	Type string `json:"type,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this milestone entity
func (m *MilestoneEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOccurredAt(formats); err != nil {
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

func (m *MilestoneEntity) validateOccurredAt(formats strfmt.Registry) error {
	if swag.IsZero(m.OccurredAt) { // not required
		return nil
	}

	if err := validate.FormatOf("occurred_at", "body", "date-time", m.OccurredAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var milestoneEntityTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["started","detected","acknowledged","first_action","mitigated","resolved","postmortem_started","postmortem_completed","closed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		milestoneEntityTypeTypePropEnum = append(milestoneEntityTypeTypePropEnum, v)
	}
}

const (

	// MilestoneEntityTypeStarted captures enum value "started"
	MilestoneEntityTypeStarted string = "started"

	// MilestoneEntityTypeDetected captures enum value "detected"
	MilestoneEntityTypeDetected string = "detected"

	// MilestoneEntityTypeAcknowledged captures enum value "acknowledged"
	MilestoneEntityTypeAcknowledged string = "acknowledged"

	// MilestoneEntityTypeFirstAction captures enum value "first_action"
	MilestoneEntityTypeFirstAction string = "first_action"

	// MilestoneEntityTypeMitigated captures enum value "mitigated"
	MilestoneEntityTypeMitigated string = "mitigated"

	// MilestoneEntityTypeResolved captures enum value "resolved"
	MilestoneEntityTypeResolved string = "resolved"

	// MilestoneEntityTypePostmortemStarted captures enum value "postmortem_started"
	MilestoneEntityTypePostmortemStarted string = "postmortem_started"

	// MilestoneEntityTypePostmortemCompleted captures enum value "postmortem_completed"
	MilestoneEntityTypePostmortemCompleted string = "postmortem_completed"

	// MilestoneEntityTypeClosed captures enum value "closed"
	MilestoneEntityTypeClosed string = "closed"
)

// prop value enum
func (m *MilestoneEntity) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, milestoneEntityTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MilestoneEntity) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this milestone entity based on context it is used
func (m *MilestoneEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MilestoneEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MilestoneEntity) UnmarshalBinary(b []byte) error {
	var res MilestoneEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}