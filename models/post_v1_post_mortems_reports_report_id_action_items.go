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

// PostV1PostMortemsReportsReportIDActionItems Create an action item on a report
//
// swagger:model postV1PostMortemsReportsReportIdActionItems
type PostV1PostMortemsReportsReportIDActionItems struct {

	// description
	Description string `json:"description,omitempty"`

	// summary
	// Required: true
	Summary *string `json:"summary"`
}

// Validate validates this post v1 post mortems reports report Id action items
func (m *PostV1PostMortemsReportsReportIDActionItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1PostMortemsReportsReportIDActionItems) validateSummary(formats strfmt.Registry) error {

	if err := validate.Required("summary", "body", m.Summary); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 post mortems reports report Id action items based on context it is used
func (m *PostV1PostMortemsReportsReportIDActionItems) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1PostMortemsReportsReportIDActionItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1PostMortemsReportsReportIDActionItems) UnmarshalBinary(b []byte) error {
	var res PostV1PostMortemsReportsReportIDActionItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}