// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams creates a new DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams() *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams {
	return &DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParamsWithTimeout creates a new DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParamsWithTimeout(timeout time.Duration) *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams {
	return &DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParamsWithContext creates a new DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams object
// with the ability to set a context for a request.
func NewDeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParamsWithContext(ctx context.Context) *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams {
	return &DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams{
		Context: ctx,
	}
}

// NewDeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParamsWithHTTPClient creates a new DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParamsWithHTTPClient(client *http.Client) *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams {
	return &DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams{
		HTTPClient: client,
	}
}

/* DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams contains all the parameters to send to the API endpoint
   for the delete v1 incidents incident Id role assignments role assignment Id operation.

   Typically these are written to a http.Request.
*/
type DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams struct {

	// IncidentID.
	IncidentID string

	// RoleAssignmentID.
	RoleAssignmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 incidents incident Id role assignments role assignment Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams) WithDefaults() *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 incidents incident Id role assignments role assignment Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 incidents incident Id role assignments role assignment Id params
func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams) WithTimeout(timeout time.Duration) *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 incidents incident Id role assignments role assignment Id params
func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 incidents incident Id role assignments role assignment Id params
func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams) WithContext(ctx context.Context) *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 incidents incident Id role assignments role assignment Id params
func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 incidents incident Id role assignments role assignment Id params
func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams) WithHTTPClient(client *http.Client) *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 incidents incident Id role assignments role assignment Id params
func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the delete v1 incidents incident Id role assignments role assignment Id params
func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams) WithIncidentID(incidentID string) *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the delete v1 incidents incident Id role assignments role assignment Id params
func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithRoleAssignmentID adds the roleAssignmentID to the delete v1 incidents incident Id role assignments role assignment Id params
func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams) WithRoleAssignmentID(roleAssignmentID string) *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams {
	o.SetRoleAssignmentID(roleAssignmentID)
	return o
}

// SetRoleAssignmentID adds the roleAssignmentId to the delete v1 incidents incident Id role assignments role assignment Id params
func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams) SetRoleAssignmentID(roleAssignmentID string) {
	o.RoleAssignmentID = roleAssignmentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	// path param role_assignment_id
	if err := r.SetPathParam("role_assignment_id", o.RoleAssignmentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
