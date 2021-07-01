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
)

// NewDeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams creates a new DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams() *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams {
	return &DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParamsWithTimeout creates a new DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParamsWithTimeout(timeout time.Duration) *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams {
	return &DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParamsWithContext creates a new DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams object
// with the ability to set a context for a request.
func NewDeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParamsWithContext(ctx context.Context) *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams {
	return &DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams{
		Context: ctx,
	}
}

// NewDeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParamsWithHTTPClient creates a new DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParamsWithHTTPClient(client *http.Client) *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams {
	return &DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams{
		HTTPClient: client,
	}
}

/* DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams contains all the parameters to send to the API endpoint
   for the delete v1 post mortems reports report Id participants participant Id operation.

   Typically these are written to a http.Request.
*/
type DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams struct {

	// ParticipantID.
	ParticipantID string

	// ReportID.
	ReportID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 post mortems reports report Id participants participant Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams) WithDefaults() *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 post mortems reports report Id participants participant Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 post mortems reports report Id participants participant Id params
func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams) WithTimeout(timeout time.Duration) *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 post mortems reports report Id participants participant Id params
func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 post mortems reports report Id participants participant Id params
func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams) WithContext(ctx context.Context) *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 post mortems reports report Id participants participant Id params
func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 post mortems reports report Id participants participant Id params
func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams) WithHTTPClient(client *http.Client) *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 post mortems reports report Id participants participant Id params
func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParticipantID adds the participantID to the delete v1 post mortems reports report Id participants participant Id params
func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams) WithParticipantID(participantID string) *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the delete v1 post mortems reports report Id participants participant Id params
func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WithReportID adds the reportID to the delete v1 post mortems reports report Id participants participant Id params
func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams) WithReportID(reportID string) *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the delete v1 post mortems reports report Id participants participant Id params
func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams) SetReportID(reportID string) {
	o.ReportID = reportID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param participant_id
	if err := r.SetPathParam("participant_id", o.ParticipantID); err != nil {
		return err
	}

	// path param report_id
	if err := r.SetPathParam("report_id", o.ReportID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
