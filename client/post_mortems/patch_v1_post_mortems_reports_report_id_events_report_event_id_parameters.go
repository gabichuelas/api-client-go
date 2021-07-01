// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

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

// NewPatchV1PostMortemsReportsReportIDEventsReportEventIDParams creates a new PatchV1PostMortemsReportsReportIDEventsReportEventIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1PostMortemsReportsReportIDEventsReportEventIDParams() *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams {
	return &PatchV1PostMortemsReportsReportIDEventsReportEventIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1PostMortemsReportsReportIDEventsReportEventIDParamsWithTimeout creates a new PatchV1PostMortemsReportsReportIDEventsReportEventIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1PostMortemsReportsReportIDEventsReportEventIDParamsWithTimeout(timeout time.Duration) *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams {
	return &PatchV1PostMortemsReportsReportIDEventsReportEventIDParams{
		timeout: timeout,
	}
}

// NewPatchV1PostMortemsReportsReportIDEventsReportEventIDParamsWithContext creates a new PatchV1PostMortemsReportsReportIDEventsReportEventIDParams object
// with the ability to set a context for a request.
func NewPatchV1PostMortemsReportsReportIDEventsReportEventIDParamsWithContext(ctx context.Context) *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams {
	return &PatchV1PostMortemsReportsReportIDEventsReportEventIDParams{
		Context: ctx,
	}
}

// NewPatchV1PostMortemsReportsReportIDEventsReportEventIDParamsWithHTTPClient creates a new PatchV1PostMortemsReportsReportIDEventsReportEventIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1PostMortemsReportsReportIDEventsReportEventIDParamsWithHTTPClient(client *http.Client) *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams {
	return &PatchV1PostMortemsReportsReportIDEventsReportEventIDParams{
		HTTPClient: client,
	}
}

/* PatchV1PostMortemsReportsReportIDEventsReportEventIDParams contains all the parameters to send to the API endpoint
   for the patch v1 post mortems reports report Id events report event Id operation.

   Typically these are written to a http.Request.
*/
type PatchV1PostMortemsReportsReportIDEventsReportEventIDParams struct {

	// V1PostMortemsReportsReportIDEvents.
	V1PostMortemsReportsReportIDEvents *models.PatchV1PostMortemsReportsReportIDEvents

	// ReportEventID.
	ReportEventID string

	// ReportID.
	//
	// Format: int32
	ReportID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 post mortems reports report Id events report event Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams) WithDefaults() *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 post mortems reports report Id events report event Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 post mortems reports report Id events report event Id params
func (o *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams) WithTimeout(timeout time.Duration) *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 post mortems reports report Id events report event Id params
func (o *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 post mortems reports report Id events report event Id params
func (o *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams) WithContext(ctx context.Context) *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 post mortems reports report Id events report event Id params
func (o *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 post mortems reports report Id events report event Id params
func (o *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams) WithHTTPClient(client *http.Client) *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 post mortems reports report Id events report event Id params
func (o *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1PostMortemsReportsReportIDEvents adds the v1PostMortemsReportsReportIDEvents to the patch v1 post mortems reports report Id events report event Id params
func (o *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams) WithV1PostMortemsReportsReportIDEvents(v1PostMortemsReportsReportIDEvents *models.PatchV1PostMortemsReportsReportIDEvents) *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams {
	o.SetV1PostMortemsReportsReportIDEvents(v1PostMortemsReportsReportIDEvents)
	return o
}

// SetV1PostMortemsReportsReportIDEvents adds the v1PostMortemsReportsReportIdEvents to the patch v1 post mortems reports report Id events report event Id params
func (o *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams) SetV1PostMortemsReportsReportIDEvents(v1PostMortemsReportsReportIDEvents *models.PatchV1PostMortemsReportsReportIDEvents) {
	o.V1PostMortemsReportsReportIDEvents = v1PostMortemsReportsReportIDEvents
}

// WithReportEventID adds the reportEventID to the patch v1 post mortems reports report Id events report event Id params
func (o *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams) WithReportEventID(reportEventID string) *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams {
	o.SetReportEventID(reportEventID)
	return o
}

// SetReportEventID adds the reportEventId to the patch v1 post mortems reports report Id events report event Id params
func (o *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams) SetReportEventID(reportEventID string) {
	o.ReportEventID = reportEventID
}

// WithReportID adds the reportID to the patch v1 post mortems reports report Id events report event Id params
func (o *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams) WithReportID(reportID int32) *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the patch v1 post mortems reports report Id events report event Id params
func (o *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams) SetReportID(reportID int32) {
	o.ReportID = reportID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1PostMortemsReportsReportIDEventsReportEventIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1PostMortemsReportsReportIDEvents != nil {
		if err := r.SetBodyParam(o.V1PostMortemsReportsReportIDEvents); err != nil {
			return err
		}
	}

	// path param report_event_id
	if err := r.SetPathParam("report_event_id", o.ReportEventID); err != nil {
		return err
	}

	// path param report_id
	if err := r.SetPathParam("report_id", swag.FormatInt32(o.ReportID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
