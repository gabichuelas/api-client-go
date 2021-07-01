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
	"github.com/go-openapi/swag"

	"github.com/firehydrant/api-client-go/models"
)

// NewPatchV1TicketingTicketsTicketIDParams creates a new PatchV1TicketingTicketsTicketIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1TicketingTicketsTicketIDParams() *PatchV1TicketingTicketsTicketIDParams {
	return &PatchV1TicketingTicketsTicketIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1TicketingTicketsTicketIDParamsWithTimeout creates a new PatchV1TicketingTicketsTicketIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1TicketingTicketsTicketIDParamsWithTimeout(timeout time.Duration) *PatchV1TicketingTicketsTicketIDParams {
	return &PatchV1TicketingTicketsTicketIDParams{
		timeout: timeout,
	}
}

// NewPatchV1TicketingTicketsTicketIDParamsWithContext creates a new PatchV1TicketingTicketsTicketIDParams object
// with the ability to set a context for a request.
func NewPatchV1TicketingTicketsTicketIDParamsWithContext(ctx context.Context) *PatchV1TicketingTicketsTicketIDParams {
	return &PatchV1TicketingTicketsTicketIDParams{
		Context: ctx,
	}
}

// NewPatchV1TicketingTicketsTicketIDParamsWithHTTPClient creates a new PatchV1TicketingTicketsTicketIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1TicketingTicketsTicketIDParamsWithHTTPClient(client *http.Client) *PatchV1TicketingTicketsTicketIDParams {
	return &PatchV1TicketingTicketsTicketIDParams{
		HTTPClient: client,
	}
}

/* PatchV1TicketingTicketsTicketIDParams contains all the parameters to send to the API endpoint
   for the patch v1 ticketing tickets ticket Id operation.

   Typically these are written to a http.Request.
*/
type PatchV1TicketingTicketsTicketIDParams struct {

	// V1TicketingTickets.
	V1TicketingTickets *models.PatchV1TicketingTickets

	// TicketID.
	//
	// Format: int32
	TicketID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 ticketing tickets ticket Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1TicketingTicketsTicketIDParams) WithDefaults() *PatchV1TicketingTicketsTicketIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 ticketing tickets ticket Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1TicketingTicketsTicketIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 ticketing tickets ticket Id params
func (o *PatchV1TicketingTicketsTicketIDParams) WithTimeout(timeout time.Duration) *PatchV1TicketingTicketsTicketIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 ticketing tickets ticket Id params
func (o *PatchV1TicketingTicketsTicketIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 ticketing tickets ticket Id params
func (o *PatchV1TicketingTicketsTicketIDParams) WithContext(ctx context.Context) *PatchV1TicketingTicketsTicketIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 ticketing tickets ticket Id params
func (o *PatchV1TicketingTicketsTicketIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 ticketing tickets ticket Id params
func (o *PatchV1TicketingTicketsTicketIDParams) WithHTTPClient(client *http.Client) *PatchV1TicketingTicketsTicketIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 ticketing tickets ticket Id params
func (o *PatchV1TicketingTicketsTicketIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1TicketingTickets adds the v1TicketingTickets to the patch v1 ticketing tickets ticket Id params
func (o *PatchV1TicketingTicketsTicketIDParams) WithV1TicketingTickets(v1TicketingTickets *models.PatchV1TicketingTickets) *PatchV1TicketingTicketsTicketIDParams {
	o.SetV1TicketingTickets(v1TicketingTickets)
	return o
}

// SetV1TicketingTickets adds the v1TicketingTickets to the patch v1 ticketing tickets ticket Id params
func (o *PatchV1TicketingTicketsTicketIDParams) SetV1TicketingTickets(v1TicketingTickets *models.PatchV1TicketingTickets) {
	o.V1TicketingTickets = v1TicketingTickets
}

// WithTicketID adds the ticketID to the patch v1 ticketing tickets ticket Id params
func (o *PatchV1TicketingTicketsTicketIDParams) WithTicketID(ticketID int32) *PatchV1TicketingTicketsTicketIDParams {
	o.SetTicketID(ticketID)
	return o
}

// SetTicketID adds the ticketId to the patch v1 ticketing tickets ticket Id params
func (o *PatchV1TicketingTicketsTicketIDParams) SetTicketID(ticketID int32) {
	o.TicketID = ticketID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1TicketingTicketsTicketIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1TicketingTickets != nil {
		if err := r.SetBodyParam(o.V1TicketingTickets); err != nil {
			return err
		}
	}

	// path param ticket_id
	if err := r.SetPathParam("ticket_id", swag.FormatInt32(o.TicketID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
