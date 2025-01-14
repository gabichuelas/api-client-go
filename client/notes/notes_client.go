// Code generated by go-swagger; DO NOT EDIT.

package notes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new notes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for notes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PatchV1NotesNoteID(params *PatchV1NotesNoteIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1NotesNoteIDOK, error)

	PostV1Notes(params *PostV1NotesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1NotesCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PatchV1NotesNoteID Update a note
*/
func (a *Client) PatchV1NotesNoteID(params *PatchV1NotesNoteIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1NotesNoteIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1NotesNoteIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1NotesNoteId",
		Method:             "PATCH",
		PathPattern:        "/v1/notes/{note_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1NotesNoteIDReader{formats: a.formats},
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
	success, ok := result.(*PatchV1NotesNoteIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1NotesNoteId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostV1Notes Add a note to an incident
*/
func (a *Client) PostV1Notes(params *PostV1NotesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1NotesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1NotesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1Notes",
		Method:             "POST",
		PathPattern:        "/v1/notes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1NotesReader{formats: a.formats},
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
	success, ok := result.(*PostV1NotesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1Notes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
