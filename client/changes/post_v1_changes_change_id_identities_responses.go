// Code generated by go-swagger; DO NOT EDIT.

package changes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1ChangesChangeIDIdentitiesReader is a Reader for the PostV1ChangesChangeIDIdentities structure.
type PostV1ChangesChangeIDIdentitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ChangesChangeIDIdentitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1ChangesChangeIDIdentitiesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1ChangesChangeIDIdentitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1ChangesChangeIDIdentitiesCreated creates a PostV1ChangesChangeIDIdentitiesCreated with default headers values
func NewPostV1ChangesChangeIDIdentitiesCreated() *PostV1ChangesChangeIDIdentitiesCreated {
	return &PostV1ChangesChangeIDIdentitiesCreated{}
}

/* PostV1ChangesChangeIDIdentitiesCreated describes a response with status code 201, with default header values.

Create an identity for this change
*/
type PostV1ChangesChangeIDIdentitiesCreated struct {
	Payload *models.ChangeIdentityEntity
}

func (o *PostV1ChangesChangeIDIdentitiesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/changes/{change_id}/identities][%d] postV1ChangesChangeIdIdentitiesCreated  %+v", 201, o.Payload)
}
func (o *PostV1ChangesChangeIDIdentitiesCreated) GetPayload() *models.ChangeIdentityEntity {
	return o.Payload
}

func (o *PostV1ChangesChangeIDIdentitiesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChangeIdentityEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ChangesChangeIDIdentitiesBadRequest creates a PostV1ChangesChangeIDIdentitiesBadRequest with default headers values
func NewPostV1ChangesChangeIDIdentitiesBadRequest() *PostV1ChangesChangeIDIdentitiesBadRequest {
	return &PostV1ChangesChangeIDIdentitiesBadRequest{}
}

/* PostV1ChangesChangeIDIdentitiesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostV1ChangesChangeIDIdentitiesBadRequest struct {
	Payload *models.ErrorEntity
}

func (o *PostV1ChangesChangeIDIdentitiesBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/changes/{change_id}/identities][%d] postV1ChangesChangeIdIdentitiesBadRequest  %+v", 400, o.Payload)
}
func (o *PostV1ChangesChangeIDIdentitiesBadRequest) GetPayload() *models.ErrorEntity {
	return o.Payload
}

func (o *PostV1ChangesChangeIDIdentitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
