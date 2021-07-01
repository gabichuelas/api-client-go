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

// GetV1ChangesChangeIDIdentitiesReader is a Reader for the GetV1ChangesChangeIDIdentities structure.
type GetV1ChangesChangeIDIdentitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ChangesChangeIDIdentitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ChangesChangeIDIdentitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1ChangesChangeIDIdentitiesOK creates a GetV1ChangesChangeIDIdentitiesOK with default headers values
func NewGetV1ChangesChangeIDIdentitiesOK() *GetV1ChangesChangeIDIdentitiesOK {
	return &GetV1ChangesChangeIDIdentitiesOK{}
}

/* GetV1ChangesChangeIDIdentitiesOK describes a response with status code 200, with default header values.

Retrieve all identities for the change
*/
type GetV1ChangesChangeIDIdentitiesOK struct {
	Payload *models.ChangeIdentityEntityPaginated
}

func (o *GetV1ChangesChangeIDIdentitiesOK) Error() string {
	return fmt.Sprintf("[GET /v1/changes/{change_id}/identities][%d] getV1ChangesChangeIdIdentitiesOK  %+v", 200, o.Payload)
}
func (o *GetV1ChangesChangeIDIdentitiesOK) GetPayload() *models.ChangeIdentityEntityPaginated {
	return o.Payload
}

func (o *GetV1ChangesChangeIDIdentitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChangeIdentityEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
