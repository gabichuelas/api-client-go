// Code generated by go-swagger; DO NOT EDIT.

package changes

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

	"github.com/firehydrant/api-client-go/models"
)

// NewPostV1ChangesChangeIDIdentitiesParams creates a new PostV1ChangesChangeIDIdentitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1ChangesChangeIDIdentitiesParams() *PostV1ChangesChangeIDIdentitiesParams {
	return &PostV1ChangesChangeIDIdentitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1ChangesChangeIDIdentitiesParamsWithTimeout creates a new PostV1ChangesChangeIDIdentitiesParams object
// with the ability to set a timeout on a request.
func NewPostV1ChangesChangeIDIdentitiesParamsWithTimeout(timeout time.Duration) *PostV1ChangesChangeIDIdentitiesParams {
	return &PostV1ChangesChangeIDIdentitiesParams{
		timeout: timeout,
	}
}

// NewPostV1ChangesChangeIDIdentitiesParamsWithContext creates a new PostV1ChangesChangeIDIdentitiesParams object
// with the ability to set a context for a request.
func NewPostV1ChangesChangeIDIdentitiesParamsWithContext(ctx context.Context) *PostV1ChangesChangeIDIdentitiesParams {
	return &PostV1ChangesChangeIDIdentitiesParams{
		Context: ctx,
	}
}

// NewPostV1ChangesChangeIDIdentitiesParamsWithHTTPClient creates a new PostV1ChangesChangeIDIdentitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1ChangesChangeIDIdentitiesParamsWithHTTPClient(client *http.Client) *PostV1ChangesChangeIDIdentitiesParams {
	return &PostV1ChangesChangeIDIdentitiesParams{
		HTTPClient: client,
	}
}

/* PostV1ChangesChangeIDIdentitiesParams contains all the parameters to send to the API endpoint
   for the post v1 changes change Id identities operation.

   Typically these are written to a http.Request.
*/
type PostV1ChangesChangeIDIdentitiesParams struct {

	// V1ChangesChangeIDIdentities.
	V1ChangesChangeIDIdentities *models.PostV1ChangesChangeIDIdentities

	// ChangeID.
	ChangeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 changes change Id identities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ChangesChangeIDIdentitiesParams) WithDefaults() *PostV1ChangesChangeIDIdentitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 changes change Id identities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ChangesChangeIDIdentitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 changes change Id identities params
func (o *PostV1ChangesChangeIDIdentitiesParams) WithTimeout(timeout time.Duration) *PostV1ChangesChangeIDIdentitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 changes change Id identities params
func (o *PostV1ChangesChangeIDIdentitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 changes change Id identities params
func (o *PostV1ChangesChangeIDIdentitiesParams) WithContext(ctx context.Context) *PostV1ChangesChangeIDIdentitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 changes change Id identities params
func (o *PostV1ChangesChangeIDIdentitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 changes change Id identities params
func (o *PostV1ChangesChangeIDIdentitiesParams) WithHTTPClient(client *http.Client) *PostV1ChangesChangeIDIdentitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 changes change Id identities params
func (o *PostV1ChangesChangeIDIdentitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1ChangesChangeIDIdentities adds the v1ChangesChangeIDIdentities to the post v1 changes change Id identities params
func (o *PostV1ChangesChangeIDIdentitiesParams) WithV1ChangesChangeIDIdentities(v1ChangesChangeIDIdentities *models.PostV1ChangesChangeIDIdentities) *PostV1ChangesChangeIDIdentitiesParams {
	o.SetV1ChangesChangeIDIdentities(v1ChangesChangeIDIdentities)
	return o
}

// SetV1ChangesChangeIDIdentities adds the v1ChangesChangeIdIdentities to the post v1 changes change Id identities params
func (o *PostV1ChangesChangeIDIdentitiesParams) SetV1ChangesChangeIDIdentities(v1ChangesChangeIDIdentities *models.PostV1ChangesChangeIDIdentities) {
	o.V1ChangesChangeIDIdentities = v1ChangesChangeIDIdentities
}

// WithChangeID adds the changeID to the post v1 changes change Id identities params
func (o *PostV1ChangesChangeIDIdentitiesParams) WithChangeID(changeID string) *PostV1ChangesChangeIDIdentitiesParams {
	o.SetChangeID(changeID)
	return o
}

// SetChangeID adds the changeId to the post v1 changes change Id identities params
func (o *PostV1ChangesChangeIDIdentitiesParams) SetChangeID(changeID string) {
	o.ChangeID = changeID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1ChangesChangeIDIdentitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1ChangesChangeIDIdentities != nil {
		if err := r.SetBodyParam(o.V1ChangesChangeIDIdentities); err != nil {
			return err
		}
	}

	// path param change_id
	if err := r.SetPathParam("change_id", o.ChangeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}