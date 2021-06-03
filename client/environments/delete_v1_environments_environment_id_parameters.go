// Code generated by go-swagger; DO NOT EDIT.

package environments

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

// NewDeleteV1EnvironmentsEnvironmentIDParams creates a new DeleteV1EnvironmentsEnvironmentIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1EnvironmentsEnvironmentIDParams() *DeleteV1EnvironmentsEnvironmentIDParams {
	return &DeleteV1EnvironmentsEnvironmentIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1EnvironmentsEnvironmentIDParamsWithTimeout creates a new DeleteV1EnvironmentsEnvironmentIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1EnvironmentsEnvironmentIDParamsWithTimeout(timeout time.Duration) *DeleteV1EnvironmentsEnvironmentIDParams {
	return &DeleteV1EnvironmentsEnvironmentIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1EnvironmentsEnvironmentIDParamsWithContext creates a new DeleteV1EnvironmentsEnvironmentIDParams object
// with the ability to set a context for a request.
func NewDeleteV1EnvironmentsEnvironmentIDParamsWithContext(ctx context.Context) *DeleteV1EnvironmentsEnvironmentIDParams {
	return &DeleteV1EnvironmentsEnvironmentIDParams{
		Context: ctx,
	}
}

// NewDeleteV1EnvironmentsEnvironmentIDParamsWithHTTPClient creates a new DeleteV1EnvironmentsEnvironmentIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1EnvironmentsEnvironmentIDParamsWithHTTPClient(client *http.Client) *DeleteV1EnvironmentsEnvironmentIDParams {
	return &DeleteV1EnvironmentsEnvironmentIDParams{
		HTTPClient: client,
	}
}

/* DeleteV1EnvironmentsEnvironmentIDParams contains all the parameters to send to the API endpoint
   for the delete v1 environments environment Id operation.

   Typically these are written to a http.Request.
*/
type DeleteV1EnvironmentsEnvironmentIDParams struct {

	/* EnvironmentID.

	   Environment UUID
	*/
	EnvironmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 environments environment Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1EnvironmentsEnvironmentIDParams) WithDefaults() *DeleteV1EnvironmentsEnvironmentIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 environments environment Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1EnvironmentsEnvironmentIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 environments environment Id params
func (o *DeleteV1EnvironmentsEnvironmentIDParams) WithTimeout(timeout time.Duration) *DeleteV1EnvironmentsEnvironmentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 environments environment Id params
func (o *DeleteV1EnvironmentsEnvironmentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 environments environment Id params
func (o *DeleteV1EnvironmentsEnvironmentIDParams) WithContext(ctx context.Context) *DeleteV1EnvironmentsEnvironmentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 environments environment Id params
func (o *DeleteV1EnvironmentsEnvironmentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 environments environment Id params
func (o *DeleteV1EnvironmentsEnvironmentIDParams) WithHTTPClient(client *http.Client) *DeleteV1EnvironmentsEnvironmentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 environments environment Id params
func (o *DeleteV1EnvironmentsEnvironmentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the delete v1 environments environment Id params
func (o *DeleteV1EnvironmentsEnvironmentIDParams) WithEnvironmentID(environmentID string) *DeleteV1EnvironmentsEnvironmentIDParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the delete v1 environments environment Id params
func (o *DeleteV1EnvironmentsEnvironmentIDParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1EnvironmentsEnvironmentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environment_id
	if err := r.SetPathParam("environment_id", o.EnvironmentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}