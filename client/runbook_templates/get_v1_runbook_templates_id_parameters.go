// Code generated by go-swagger; DO NOT EDIT.

package runbook_templates

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

// NewGetV1RunbookTemplatesIDParams creates a new GetV1RunbookTemplatesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1RunbookTemplatesIDParams() *GetV1RunbookTemplatesIDParams {
	return &GetV1RunbookTemplatesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1RunbookTemplatesIDParamsWithTimeout creates a new GetV1RunbookTemplatesIDParams object
// with the ability to set a timeout on a request.
func NewGetV1RunbookTemplatesIDParamsWithTimeout(timeout time.Duration) *GetV1RunbookTemplatesIDParams {
	return &GetV1RunbookTemplatesIDParams{
		timeout: timeout,
	}
}

// NewGetV1RunbookTemplatesIDParamsWithContext creates a new GetV1RunbookTemplatesIDParams object
// with the ability to set a context for a request.
func NewGetV1RunbookTemplatesIDParamsWithContext(ctx context.Context) *GetV1RunbookTemplatesIDParams {
	return &GetV1RunbookTemplatesIDParams{
		Context: ctx,
	}
}

// NewGetV1RunbookTemplatesIDParamsWithHTTPClient creates a new GetV1RunbookTemplatesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1RunbookTemplatesIDParamsWithHTTPClient(client *http.Client) *GetV1RunbookTemplatesIDParams {
	return &GetV1RunbookTemplatesIDParams{
		HTTPClient: client,
	}
}

/* GetV1RunbookTemplatesIDParams contains all the parameters to send to the API endpoint
   for the get v1 runbook templates Id operation.

   Typically these are written to a http.Request.
*/
type GetV1RunbookTemplatesIDParams struct {

	/* ID.

	   UUID of the Runbook template
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 runbook templates Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RunbookTemplatesIDParams) WithDefaults() *GetV1RunbookTemplatesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 runbook templates Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RunbookTemplatesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 runbook templates Id params
func (o *GetV1RunbookTemplatesIDParams) WithTimeout(timeout time.Duration) *GetV1RunbookTemplatesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 runbook templates Id params
func (o *GetV1RunbookTemplatesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 runbook templates Id params
func (o *GetV1RunbookTemplatesIDParams) WithContext(ctx context.Context) *GetV1RunbookTemplatesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 runbook templates Id params
func (o *GetV1RunbookTemplatesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 runbook templates Id params
func (o *GetV1RunbookTemplatesIDParams) WithHTTPClient(client *http.Client) *GetV1RunbookTemplatesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 runbook templates Id params
func (o *GetV1RunbookTemplatesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get v1 runbook templates Id params
func (o *GetV1RunbookTemplatesIDParams) WithID(id string) *GetV1RunbookTemplatesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 runbook templates Id params
func (o *GetV1RunbookTemplatesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1RunbookTemplatesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
