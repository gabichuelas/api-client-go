// Code generated by go-swagger; DO NOT EDIT.

package runbooks

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
)

// NewGetV1RunbooksActionsParams creates a new GetV1RunbooksActionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1RunbooksActionsParams() *GetV1RunbooksActionsParams {
	return &GetV1RunbooksActionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1RunbooksActionsParamsWithTimeout creates a new GetV1RunbooksActionsParams object
// with the ability to set a timeout on a request.
func NewGetV1RunbooksActionsParamsWithTimeout(timeout time.Duration) *GetV1RunbooksActionsParams {
	return &GetV1RunbooksActionsParams{
		timeout: timeout,
	}
}

// NewGetV1RunbooksActionsParamsWithContext creates a new GetV1RunbooksActionsParams object
// with the ability to set a context for a request.
func NewGetV1RunbooksActionsParamsWithContext(ctx context.Context) *GetV1RunbooksActionsParams {
	return &GetV1RunbooksActionsParams{
		Context: ctx,
	}
}

// NewGetV1RunbooksActionsParamsWithHTTPClient creates a new GetV1RunbooksActionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1RunbooksActionsParamsWithHTTPClient(client *http.Client) *GetV1RunbooksActionsParams {
	return &GetV1RunbooksActionsParams{
		HTTPClient: client,
	}
}

/* GetV1RunbooksActionsParams contains all the parameters to send to the API endpoint
   for the get v1 runbooks actions operation.

   Typically these are written to a http.Request.
*/
type GetV1RunbooksActionsParams struct {

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	/* Type.

	   List actions supporting this specific Runbook type
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 runbooks actions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RunbooksActionsParams) WithDefaults() *GetV1RunbooksActionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 runbooks actions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RunbooksActionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 runbooks actions params
func (o *GetV1RunbooksActionsParams) WithTimeout(timeout time.Duration) *GetV1RunbooksActionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 runbooks actions params
func (o *GetV1RunbooksActionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 runbooks actions params
func (o *GetV1RunbooksActionsParams) WithContext(ctx context.Context) *GetV1RunbooksActionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 runbooks actions params
func (o *GetV1RunbooksActionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 runbooks actions params
func (o *GetV1RunbooksActionsParams) WithHTTPClient(client *http.Client) *GetV1RunbooksActionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 runbooks actions params
func (o *GetV1RunbooksActionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 runbooks actions params
func (o *GetV1RunbooksActionsParams) WithPage(page *int32) *GetV1RunbooksActionsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 runbooks actions params
func (o *GetV1RunbooksActionsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 runbooks actions params
func (o *GetV1RunbooksActionsParams) WithPerPage(perPage *int32) *GetV1RunbooksActionsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 runbooks actions params
func (o *GetV1RunbooksActionsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithType adds the typeVar to the get v1 runbooks actions params
func (o *GetV1RunbooksActionsParams) WithType(typeVar *string) *GetV1RunbooksActionsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get v1 runbooks actions params
func (o *GetV1RunbooksActionsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1RunbooksActionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
