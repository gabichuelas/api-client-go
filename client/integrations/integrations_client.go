// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new integrations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for integrations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1IntegrationsStatuspageConnectionsID(params *DeleteV1IntegrationsStatuspageConnectionsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1IntegrationsStatuspageConnectionsIDOK, error)

	GetV1Integrations(params *GetV1IntegrationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsOK, error)

	GetV1IntegrationsAwsCloudtrailBatches(params *GetV1IntegrationsAwsCloudtrailBatchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsAwsCloudtrailBatchesOK, error)

	GetV1IntegrationsAwsCloudtrailBatchesID(params *GetV1IntegrationsAwsCloudtrailBatchesIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsAwsCloudtrailBatchesIDOK, error)

	GetV1IntegrationsAwsCloudtrailBatchesIDEvents(params *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK, error)

	GetV1IntegrationsAwsConnections(params *GetV1IntegrationsAwsConnectionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsAwsConnectionsOK, error)

	GetV1IntegrationsAwsConnectionsID(params *GetV1IntegrationsAwsConnectionsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsAwsConnectionsIDOK, error)

	GetV1IntegrationsConnections(params *GetV1IntegrationsConnectionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsConnectionsOK, error)

	GetV1IntegrationsStatusesSlug(params *GetV1IntegrationsStatusesSlugParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsStatusesSlugOK, error)

	GetV1IntegrationsStatuspageConnections(params *GetV1IntegrationsStatuspageConnectionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsStatuspageConnectionsOK, error)

	GetV1IntegrationsStatuspageConnectionsID(params *GetV1IntegrationsStatuspageConnectionsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsStatuspageConnectionsIDOK, error)

	GetV1IntegrationsStatuspageConnectionsIDPages(params *GetV1IntegrationsStatuspageConnectionsIDPagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsStatuspageConnectionsIDPagesOK, error)

	PatchV1IntegrationsAwsCloudtrailBatchesID(params *PatchV1IntegrationsAwsCloudtrailBatchesIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1IntegrationsAwsCloudtrailBatchesIDOK, error)

	PatchV1IntegrationsAwsConnectionsID(params *PatchV1IntegrationsAwsConnectionsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1IntegrationsAwsConnectionsIDOK, error)

	PatchV1IntegrationsStatuspageConnectionsID(params *PatchV1IntegrationsStatuspageConnectionsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1IntegrationsStatuspageConnectionsIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteV1IntegrationsStatuspageConnectionsID deletes a statuspage connection

  Deletes the given Statuspage integration connection.
*/
func (a *Client) DeleteV1IntegrationsStatuspageConnectionsID(params *DeleteV1IntegrationsStatuspageConnectionsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1IntegrationsStatuspageConnectionsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1IntegrationsStatuspageConnectionsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1IntegrationsStatuspageConnectionsId",
		Method:             "DELETE",
		PathPattern:        "/v1/integrations/statuspage/connections/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1IntegrationsStatuspageConnectionsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteV1IntegrationsStatuspageConnectionsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1IntegrationsStatuspageConnectionsId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1Integrations get v1 integrations API
*/
func (a *Client) GetV1Integrations(params *GetV1IntegrationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1IntegrationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1Integrations",
		Method:             "GET",
		PathPattern:        "/v1/integrations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1IntegrationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1IntegrationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1Integrations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1IntegrationsAwsCloudtrailBatches lists cloud trail batches

  Lists CloudTrail batches for the authenticated organization.
*/
func (a *Client) GetV1IntegrationsAwsCloudtrailBatches(params *GetV1IntegrationsAwsCloudtrailBatchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsAwsCloudtrailBatchesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1IntegrationsAwsCloudtrailBatchesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1IntegrationsAwsCloudtrailBatches",
		Method:             "GET",
		PathPattern:        "/v1/integrations/aws/cloudtrail_batches",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1IntegrationsAwsCloudtrailBatchesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1IntegrationsAwsCloudtrailBatchesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1IntegrationsAwsCloudtrailBatches: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1IntegrationsAwsCloudtrailBatchesID retrieves a cloud trail batch

  Retrieve a single CloudTrail batch.
*/
func (a *Client) GetV1IntegrationsAwsCloudtrailBatchesID(params *GetV1IntegrationsAwsCloudtrailBatchesIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsAwsCloudtrailBatchesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1IntegrationsAwsCloudtrailBatchesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1IntegrationsAwsCloudtrailBatchesId",
		Method:             "GET",
		PathPattern:        "/v1/integrations/aws/cloudtrail_batches/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1IntegrationsAwsCloudtrailBatchesIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1IntegrationsAwsCloudtrailBatchesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1IntegrationsAwsCloudtrailBatchesId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1IntegrationsAwsCloudtrailBatchesIDEvents get v1 integrations aws cloudtrail batches Id events API
*/
func (a *Client) GetV1IntegrationsAwsCloudtrailBatchesIDEvents(params *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1IntegrationsAwsCloudtrailBatchesIDEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1IntegrationsAwsCloudtrailBatchesIdEvents",
		Method:             "GET",
		PathPattern:        "/v1/integrations/aws/cloudtrail_batches/{id}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1IntegrationsAwsCloudtrailBatchesIDEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1IntegrationsAwsCloudtrailBatchesIdEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1IntegrationsAwsConnections lists a w s connections

  Lists the available and configured AWS integration connections for the authenticated organization.
*/
func (a *Client) GetV1IntegrationsAwsConnections(params *GetV1IntegrationsAwsConnectionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsAwsConnectionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1IntegrationsAwsConnectionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1IntegrationsAwsConnections",
		Method:             "GET",
		PathPattern:        "/v1/integrations/aws/connections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1IntegrationsAwsConnectionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1IntegrationsAwsConnectionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1IntegrationsAwsConnections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1IntegrationsAwsConnectionsID retrieves an a w s connection

  Retrieves the information about the AWS connection.
*/
func (a *Client) GetV1IntegrationsAwsConnectionsID(params *GetV1IntegrationsAwsConnectionsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsAwsConnectionsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1IntegrationsAwsConnectionsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1IntegrationsAwsConnectionsId",
		Method:             "GET",
		PathPattern:        "/v1/integrations/aws/connections/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1IntegrationsAwsConnectionsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1IntegrationsAwsConnectionsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1IntegrationsAwsConnectionsId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1IntegrationsConnections get v1 integrations connections API
*/
func (a *Client) GetV1IntegrationsConnections(params *GetV1IntegrationsConnectionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsConnectionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1IntegrationsConnectionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1IntegrationsConnections",
		Method:             "GET",
		PathPattern:        "/v1/integrations/connections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1IntegrationsConnectionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1IntegrationsConnectionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1IntegrationsConnections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1IntegrationsStatusesSlug get v1 integrations statuses slug API
*/
func (a *Client) GetV1IntegrationsStatusesSlug(params *GetV1IntegrationsStatusesSlugParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsStatusesSlugOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1IntegrationsStatusesSlugParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1IntegrationsStatusesSlug",
		Method:             "GET",
		PathPattern:        "/v1/integrations/statuses/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1IntegrationsStatusesSlugReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1IntegrationsStatusesSlugOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1IntegrationsStatusesSlug: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1IntegrationsStatuspageConnections lists statuspage connections

  Lists the available and configured Statuspage integrations connections for the authenticated organization.
*/
func (a *Client) GetV1IntegrationsStatuspageConnections(params *GetV1IntegrationsStatuspageConnectionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsStatuspageConnectionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1IntegrationsStatuspageConnectionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1IntegrationsStatuspageConnections",
		Method:             "GET",
		PathPattern:        "/v1/integrations/statuspage/connections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1IntegrationsStatuspageConnectionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1IntegrationsStatuspageConnectionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1IntegrationsStatuspageConnections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1IntegrationsStatuspageConnectionsID retrieves a statuspage connection

  Retrieve the information about the Statuspage connection.
*/
func (a *Client) GetV1IntegrationsStatuspageConnectionsID(params *GetV1IntegrationsStatuspageConnectionsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsStatuspageConnectionsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1IntegrationsStatuspageConnectionsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1IntegrationsStatuspageConnectionsId",
		Method:             "GET",
		PathPattern:        "/v1/integrations/statuspage/connections/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1IntegrationsStatuspageConnectionsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1IntegrationsStatuspageConnectionsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1IntegrationsStatuspageConnectionsId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1IntegrationsStatuspageConnectionsIDPages get v1 integrations statuspage connections Id pages API
*/
func (a *Client) GetV1IntegrationsStatuspageConnectionsIDPages(params *GetV1IntegrationsStatuspageConnectionsIDPagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IntegrationsStatuspageConnectionsIDPagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1IntegrationsStatuspageConnectionsIDPagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1IntegrationsStatuspageConnectionsIdPages",
		Method:             "GET",
		PathPattern:        "/v1/integrations/statuspage/connections/{id}/pages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1IntegrationsStatuspageConnectionsIDPagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1IntegrationsStatuspageConnectionsIDPagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1IntegrationsStatuspageConnectionsIdPages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchV1IntegrationsAwsCloudtrailBatchesID updates a cloud trail batch

  Update a CloudTrail batch with new information.
*/
func (a *Client) PatchV1IntegrationsAwsCloudtrailBatchesID(params *PatchV1IntegrationsAwsCloudtrailBatchesIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1IntegrationsAwsCloudtrailBatchesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1IntegrationsAwsCloudtrailBatchesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1IntegrationsAwsCloudtrailBatchesId",
		Method:             "PATCH",
		PathPattern:        "/v1/integrations/aws/cloudtrail_batches/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1IntegrationsAwsCloudtrailBatchesIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchV1IntegrationsAwsCloudtrailBatchesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1IntegrationsAwsCloudtrailBatchesId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchV1IntegrationsAwsConnectionsID updates an a w s connection

  Update the AWS connection with the provided data.
*/
func (a *Client) PatchV1IntegrationsAwsConnectionsID(params *PatchV1IntegrationsAwsConnectionsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1IntegrationsAwsConnectionsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1IntegrationsAwsConnectionsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1IntegrationsAwsConnectionsId",
		Method:             "PATCH",
		PathPattern:        "/v1/integrations/aws/connections/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1IntegrationsAwsConnectionsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchV1IntegrationsAwsConnectionsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1IntegrationsAwsConnectionsId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchV1IntegrationsStatuspageConnectionsID updates a statuspage connection

  Update the given Statuspage integration connection.
*/
func (a *Client) PatchV1IntegrationsStatuspageConnectionsID(params *PatchV1IntegrationsStatuspageConnectionsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1IntegrationsStatuspageConnectionsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1IntegrationsStatuspageConnectionsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1IntegrationsStatuspageConnectionsId",
		Method:             "PATCH",
		PathPattern:        "/v1/integrations/statuspage/connections/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1IntegrationsStatuspageConnectionsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchV1IntegrationsStatuspageConnectionsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1IntegrationsStatuspageConnectionsId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
