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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutV1IncidentsIncidentIDEventsEventIDStarsParams creates a new PutV1IncidentsIncidentIDEventsEventIDStarsParams object
// with the default values initialized.
func NewPutV1IncidentsIncidentIDEventsEventIDStarsParams() *PutV1IncidentsIncidentIDEventsEventIDStarsParams {
	var ()
	return &PutV1IncidentsIncidentIDEventsEventIDStarsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutV1IncidentsIncidentIDEventsEventIDStarsParamsWithTimeout creates a new PutV1IncidentsIncidentIDEventsEventIDStarsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutV1IncidentsIncidentIDEventsEventIDStarsParamsWithTimeout(timeout time.Duration) *PutV1IncidentsIncidentIDEventsEventIDStarsParams {
	var ()
	return &PutV1IncidentsIncidentIDEventsEventIDStarsParams{

		timeout: timeout,
	}
}

// NewPutV1IncidentsIncidentIDEventsEventIDStarsParamsWithContext creates a new PutV1IncidentsIncidentIDEventsEventIDStarsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutV1IncidentsIncidentIDEventsEventIDStarsParamsWithContext(ctx context.Context) *PutV1IncidentsIncidentIDEventsEventIDStarsParams {
	var ()
	return &PutV1IncidentsIncidentIDEventsEventIDStarsParams{

		Context: ctx,
	}
}

// NewPutV1IncidentsIncidentIDEventsEventIDStarsParamsWithHTTPClient creates a new PutV1IncidentsIncidentIDEventsEventIDStarsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutV1IncidentsIncidentIDEventsEventIDStarsParamsWithHTTPClient(client *http.Client) *PutV1IncidentsIncidentIDEventsEventIDStarsParams {
	var ()
	return &PutV1IncidentsIncidentIDEventsEventIDStarsParams{
		HTTPClient: client,
	}
}

/*PutV1IncidentsIncidentIDEventsEventIDStarsParams contains all the parameters to send to the API endpoint
for the put v1 incidents incident Id events event Id stars operation typically these are written to a http.Request
*/
type PutV1IncidentsIncidentIDEventsEventIDStarsParams struct {

	/*EventID*/
	EventID int32
	/*IncidentID*/
	IncidentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put v1 incidents incident Id events event Id stars params
func (o *PutV1IncidentsIncidentIDEventsEventIDStarsParams) WithTimeout(timeout time.Duration) *PutV1IncidentsIncidentIDEventsEventIDStarsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put v1 incidents incident Id events event Id stars params
func (o *PutV1IncidentsIncidentIDEventsEventIDStarsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put v1 incidents incident Id events event Id stars params
func (o *PutV1IncidentsIncidentIDEventsEventIDStarsParams) WithContext(ctx context.Context) *PutV1IncidentsIncidentIDEventsEventIDStarsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put v1 incidents incident Id events event Id stars params
func (o *PutV1IncidentsIncidentIDEventsEventIDStarsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put v1 incidents incident Id events event Id stars params
func (o *PutV1IncidentsIncidentIDEventsEventIDStarsParams) WithHTTPClient(client *http.Client) *PutV1IncidentsIncidentIDEventsEventIDStarsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put v1 incidents incident Id events event Id stars params
func (o *PutV1IncidentsIncidentIDEventsEventIDStarsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEventID adds the eventID to the put v1 incidents incident Id events event Id stars params
func (o *PutV1IncidentsIncidentIDEventsEventIDStarsParams) WithEventID(eventID int32) *PutV1IncidentsIncidentIDEventsEventIDStarsParams {
	o.SetEventID(eventID)
	return o
}

// SetEventID adds the eventId to the put v1 incidents incident Id events event Id stars params
func (o *PutV1IncidentsIncidentIDEventsEventIDStarsParams) SetEventID(eventID int32) {
	o.EventID = eventID
}

// WithIncidentID adds the incidentID to the put v1 incidents incident Id events event Id stars params
func (o *PutV1IncidentsIncidentIDEventsEventIDStarsParams) WithIncidentID(incidentID string) *PutV1IncidentsIncidentIDEventsEventIDStarsParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the put v1 incidents incident Id events event Id stars params
func (o *PutV1IncidentsIncidentIDEventsEventIDStarsParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1IncidentsIncidentIDEventsEventIDStarsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param event_id
	if err := r.SetPathParam("event_id", swag.FormatInt32(o.EventID)); err != nil {
		return err
	}

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}