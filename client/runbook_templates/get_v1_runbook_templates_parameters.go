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

// NewGetV1RunbookTemplatesParams creates a new GetV1RunbookTemplatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1RunbookTemplatesParams() *GetV1RunbookTemplatesParams {
	return &GetV1RunbookTemplatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1RunbookTemplatesParamsWithTimeout creates a new GetV1RunbookTemplatesParams object
// with the ability to set a timeout on a request.
func NewGetV1RunbookTemplatesParamsWithTimeout(timeout time.Duration) *GetV1RunbookTemplatesParams {
	return &GetV1RunbookTemplatesParams{
		timeout: timeout,
	}
}

// NewGetV1RunbookTemplatesParamsWithContext creates a new GetV1RunbookTemplatesParams object
// with the ability to set a context for a request.
func NewGetV1RunbookTemplatesParamsWithContext(ctx context.Context) *GetV1RunbookTemplatesParams {
	return &GetV1RunbookTemplatesParams{
		Context: ctx,
	}
}

// NewGetV1RunbookTemplatesParamsWithHTTPClient creates a new GetV1RunbookTemplatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1RunbookTemplatesParamsWithHTTPClient(client *http.Client) *GetV1RunbookTemplatesParams {
	return &GetV1RunbookTemplatesParams{
		HTTPClient: client,
	}
}

/* GetV1RunbookTemplatesParams contains all the parameters to send to the API endpoint
   for the get v1 runbook templates operation.

   Typically these are written to a http.Request.
*/
type GetV1RunbookTemplatesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 runbook templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RunbookTemplatesParams) WithDefaults() *GetV1RunbookTemplatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 runbook templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RunbookTemplatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 runbook templates params
func (o *GetV1RunbookTemplatesParams) WithTimeout(timeout time.Duration) *GetV1RunbookTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 runbook templates params
func (o *GetV1RunbookTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 runbook templates params
func (o *GetV1RunbookTemplatesParams) WithContext(ctx context.Context) *GetV1RunbookTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 runbook templates params
func (o *GetV1RunbookTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 runbook templates params
func (o *GetV1RunbookTemplatesParams) WithHTTPClient(client *http.Client) *GetV1RunbookTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 runbook templates params
func (o *GetV1RunbookTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1RunbookTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
