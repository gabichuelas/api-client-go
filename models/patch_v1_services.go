// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchV1Services Update a services attributes, you may also add or remove functionalities from the service as well.
// Note: You may not remove or add individual label key/value pairs. You must include the entire object to override label values.
//
//
// swagger:model patchV1Services
type PatchV1Services struct {

	// description
	Description string `json:"description,omitempty"`

	// An array of functionalities
	Functionalities []*PatchV1ServicesFunctionalitiesItems0 `json:"functionalities"`

	// A hash of label keys and values
	Labels map[string]string `json:"labels,omitempty"`

	// An array of links to associate with this service. This will remove all links not present in the patch. Only acts if 'links' key is included in the payload.
	Links []*PatchV1ServicesLinksItems0 `json:"links"`

	// name
	// Required: true
	Name *string `json:"name"`

	// owner
	Owner *PatchV1ServicesOwner `json:"owner,omitempty"`

	// If set to true, any functionality objects tagged on the service that are not included in the given array will be removed. Set this if you want to do a replacement operation for the functionalities
	RemoveRemainingFunctionalities bool `json:"remove_remaining_functionalities,omitempty"`

	// If set to true, any team objects tagged on the service that are not included in the given teams array will be removed. Set this if you want to do a replacement operation for the teams
	RemoveRemainingTeams bool `json:"remove_remaining_teams,omitempty"`

	// Integer representing service tier
	// Enum: [1 2 3 4 5]
	ServiceTier int32 `json:"service_tier,omitempty"`

	// An array of teams to attach to this service.
	Teams []*PatchV1ServicesTeamsItems0 `json:"teams"`
}

// Validate validates this patch v1 services
func (m *PatchV1Services) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFunctionalities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceTier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1Services) validateFunctionalities(formats strfmt.Registry) error {
	if swag.IsZero(m.Functionalities) { // not required
		return nil
	}

	for i := 0; i < len(m.Functionalities); i++ {
		if swag.IsZero(m.Functionalities[i]) { // not required
			continue
		}

		if m.Functionalities[i] != nil {
			if err := m.Functionalities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("functionalities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PatchV1Services) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PatchV1Services) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PatchV1Services) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

var patchV1ServicesTypeServiceTierPropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[1,2,3,4,5]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchV1ServicesTypeServiceTierPropEnum = append(patchV1ServicesTypeServiceTierPropEnum, v)
	}
}

// prop value enum
func (m *PatchV1Services) validateServiceTierEnum(path, location string, value int32) error {
	if err := validate.EnumCase(path, location, value, patchV1ServicesTypeServiceTierPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PatchV1Services) validateServiceTier(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceTier) { // not required
		return nil
	}

	// value enum
	if err := m.validateServiceTierEnum("service_tier", "body", m.ServiceTier); err != nil {
		return err
	}

	return nil
}

func (m *PatchV1Services) validateTeams(formats strfmt.Registry) error {
	if swag.IsZero(m.Teams) { // not required
		return nil
	}

	for i := 0; i < len(m.Teams); i++ {
		if swag.IsZero(m.Teams[i]) { // not required
			continue
		}

		if m.Teams[i] != nil {
			if err := m.Teams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this patch v1 services based on the context it is used
func (m *PatchV1Services) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFunctionalities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1Services) contextValidateFunctionalities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Functionalities); i++ {

		if m.Functionalities[i] != nil {
			if err := m.Functionalities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("functionalities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PatchV1Services) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Links); i++ {

		if m.Links[i] != nil {
			if err := m.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PatchV1Services) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {
		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *PatchV1Services) contextValidateTeams(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Teams); i++ {

		if m.Teams[i] != nil {
			if err := m.Teams[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1Services) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1Services) UnmarshalBinary(b []byte) error {
	var res PatchV1Services
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1ServicesFunctionalitiesItems0 patch v1 services functionalities items0
//
// swagger:model PatchV1ServicesFunctionalitiesItems0
type PatchV1ServicesFunctionalitiesItems0 struct {

	// If you are trying to reuse a functionality, you may set the ID to attach it to the service
	ID string `json:"id,omitempty"`

	// If you are trying to remove a functionality from a service, set this to "true"
	Remove bool `json:"remove,omitempty"`

	// If you are trying to create a new functionality and attach it to this service, set the summary key
	Summary string `json:"summary,omitempty"`
}

// Validate validates this patch v1 services functionalities items0
func (m *PatchV1ServicesFunctionalitiesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch v1 services functionalities items0 based on context it is used
func (m *PatchV1ServicesFunctionalitiesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1ServicesFunctionalitiesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1ServicesFunctionalitiesItems0) UnmarshalBinary(b []byte) error {
	var res PatchV1ServicesFunctionalitiesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1ServicesLinksItems0 patch v1 services links items0
//
// swagger:model PatchV1ServicesLinksItems0
type PatchV1ServicesLinksItems0 struct {

	// URL
	// Required: true
	HrefURL *string `json:"href_url"`

	// An optional URL to an icon representing this link
	IconURL string `json:"icon_url,omitempty"`

	// If updating an existing link, specify it's id.
	ID string `json:"id,omitempty"`

	// Short name used to display and identify this link
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this patch v1 services links items0
func (m *PatchV1ServicesLinksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHrefURL(formats); err != nil {
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

func (m *PatchV1ServicesLinksItems0) validateHrefURL(formats strfmt.Registry) error {

	if err := validate.Required("href_url", "body", m.HrefURL); err != nil {
		return err
	}

	return nil
}

func (m *PatchV1ServicesLinksItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 services links items0 based on context it is used
func (m *PatchV1ServicesLinksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1ServicesLinksItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1ServicesLinksItems0) UnmarshalBinary(b []byte) error {
	var res PatchV1ServicesLinksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1ServicesOwner An object representing a Team that owns the service
//
// swagger:model PatchV1ServicesOwner
type PatchV1ServicesOwner struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this patch v1 services owner
func (m *PatchV1ServicesOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1ServicesOwner) validateID(formats strfmt.Registry) error {

	if err := validate.Required("owner"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 services owner based on context it is used
func (m *PatchV1ServicesOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1ServicesOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1ServicesOwner) UnmarshalBinary(b []byte) error {
	var res PatchV1ServicesOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1ServicesTeamsItems0 patch v1 services teams items0
//
// swagger:model PatchV1ServicesTeamsItems0
type PatchV1ServicesTeamsItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// If you are trying to remove a team from a service, set this to "true"
	Remove bool `json:"remove,omitempty"`
}

// Validate validates this patch v1 services teams items0
func (m *PatchV1ServicesTeamsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1ServicesTeamsItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 services teams items0 based on context it is used
func (m *PatchV1ServicesTeamsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1ServicesTeamsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1ServicesTeamsItems0) UnmarshalBinary(b []byte) error {
	var res PatchV1ServicesTeamsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}