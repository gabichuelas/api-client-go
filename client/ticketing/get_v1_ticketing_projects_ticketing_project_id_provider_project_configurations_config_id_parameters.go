// Code generated by go-swagger; DO NOT EDIT.

package ticketing

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

// NewGetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams creates a new GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams() *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams {
	return &GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParamsWithTimeout creates a new GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams object
// with the ability to set a timeout on a request.
func NewGetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParamsWithTimeout(timeout time.Duration) *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams {
	return &GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams{
		timeout: timeout,
	}
}

// NewGetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParamsWithContext creates a new GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams object
// with the ability to set a context for a request.
func NewGetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParamsWithContext(ctx context.Context) *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams {
	return &GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams{
		Context: ctx,
	}
}

// NewGetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParamsWithHTTPClient creates a new GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParamsWithHTTPClient(client *http.Client) *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams {
	return &GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams{
		HTTPClient: client,
	}
}

/* GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams contains all the parameters to send to the API endpoint
   for the get v1 ticketing projects ticketing project Id provider project configurations config Id operation.

   Typically these are written to a http.Request.
*/
type GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams struct {

	// ConfigID.
	ConfigID string

	// TicketingProjectID.
	TicketingProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 ticketing projects ticketing project Id provider project configurations config Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams) WithDefaults() *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 ticketing projects ticketing project Id provider project configurations config Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 ticketing projects ticketing project Id provider project configurations config Id params
func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams) WithTimeout(timeout time.Duration) *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 ticketing projects ticketing project Id provider project configurations config Id params
func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 ticketing projects ticketing project Id provider project configurations config Id params
func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams) WithContext(ctx context.Context) *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 ticketing projects ticketing project Id provider project configurations config Id params
func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 ticketing projects ticketing project Id provider project configurations config Id params
func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams) WithHTTPClient(client *http.Client) *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 ticketing projects ticketing project Id provider project configurations config Id params
func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigID adds the configID to the get v1 ticketing projects ticketing project Id provider project configurations config Id params
func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams) WithConfigID(configID string) *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams {
	o.SetConfigID(configID)
	return o
}

// SetConfigID adds the configId to the get v1 ticketing projects ticketing project Id provider project configurations config Id params
func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams) SetConfigID(configID string) {
	o.ConfigID = configID
}

// WithTicketingProjectID adds the ticketingProjectID to the get v1 ticketing projects ticketing project Id provider project configurations config Id params
func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams) WithTicketingProjectID(ticketingProjectID string) *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams {
	o.SetTicketingProjectID(ticketingProjectID)
	return o
}

// SetTicketingProjectID adds the ticketingProjectId to the get v1 ticketing projects ticketing project Id provider project configurations config Id params
func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams) SetTicketingProjectID(ticketingProjectID string) {
	o.TicketingProjectID = ticketingProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param config_id
	if err := r.SetPathParam("config_id", o.ConfigID); err != nil {
		return err
	}

	// path param ticketing_project_id
	if err := r.SetPathParam("ticketing_project_id", o.TicketingProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}