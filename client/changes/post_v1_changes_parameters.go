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

// NewPostV1ChangesParams creates a new PostV1ChangesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1ChangesParams() *PostV1ChangesParams {
	return &PostV1ChangesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1ChangesParamsWithTimeout creates a new PostV1ChangesParams object
// with the ability to set a timeout on a request.
func NewPostV1ChangesParamsWithTimeout(timeout time.Duration) *PostV1ChangesParams {
	return &PostV1ChangesParams{
		timeout: timeout,
	}
}

// NewPostV1ChangesParamsWithContext creates a new PostV1ChangesParams object
// with the ability to set a context for a request.
func NewPostV1ChangesParamsWithContext(ctx context.Context) *PostV1ChangesParams {
	return &PostV1ChangesParams{
		Context: ctx,
	}
}

// NewPostV1ChangesParamsWithHTTPClient creates a new PostV1ChangesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1ChangesParamsWithHTTPClient(client *http.Client) *PostV1ChangesParams {
	return &PostV1ChangesParams{
		HTTPClient: client,
	}
}

/* PostV1ChangesParams contains all the parameters to send to the API endpoint
   for the post v1 changes operation.

   Typically these are written to a http.Request.
*/
type PostV1ChangesParams struct {

	// V1Changes.
	V1Changes *models.PostV1Changes

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 changes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ChangesParams) WithDefaults() *PostV1ChangesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 changes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ChangesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 changes params
func (o *PostV1ChangesParams) WithTimeout(timeout time.Duration) *PostV1ChangesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 changes params
func (o *PostV1ChangesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 changes params
func (o *PostV1ChangesParams) WithContext(ctx context.Context) *PostV1ChangesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 changes params
func (o *PostV1ChangesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 changes params
func (o *PostV1ChangesParams) WithHTTPClient(client *http.Client) *PostV1ChangesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 changes params
func (o *PostV1ChangesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1Changes adds the v1Changes to the post v1 changes params
func (o *PostV1ChangesParams) WithV1Changes(v1Changes *models.PostV1Changes) *PostV1ChangesParams {
	o.SetV1Changes(v1Changes)
	return o
}

// SetV1Changes adds the v1Changes to the post v1 changes params
func (o *PostV1ChangesParams) SetV1Changes(v1Changes *models.PostV1Changes) {
	o.V1Changes = v1Changes
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1ChangesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1Changes != nil {
		if err := r.SetBodyParam(o.V1Changes); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
