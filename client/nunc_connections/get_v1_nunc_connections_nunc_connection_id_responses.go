// Code generated by go-swagger; DO NOT EDIT.

package nunc_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1NuncConnectionsNuncConnectionIDReader is a Reader for the GetV1NuncConnectionsNuncConnectionID structure.
type GetV1NuncConnectionsNuncConnectionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1NuncConnectionsNuncConnectionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1NuncConnectionsNuncConnectionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1NuncConnectionsNuncConnectionIDOK creates a GetV1NuncConnectionsNuncConnectionIDOK with default headers values
func NewGetV1NuncConnectionsNuncConnectionIDOK() *GetV1NuncConnectionsNuncConnectionIDOK {
	return &GetV1NuncConnectionsNuncConnectionIDOK{}
}

/* GetV1NuncConnectionsNuncConnectionIDOK describes a response with status code 200, with default header values.

Retrieve the information displayed as part of your FireHydrant hosted status page.
*/
type GetV1NuncConnectionsNuncConnectionIDOK struct {
	Payload *models.NuncConnectionEntity
}

func (o *GetV1NuncConnectionsNuncConnectionIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/nunc_connections/{nunc_connection_id}][%d] getV1NuncConnectionsNuncConnectionIdOK  %+v", 200, o.Payload)
}
func (o *GetV1NuncConnectionsNuncConnectionIDOK) GetPayload() *models.NuncConnectionEntity {
	return o.Payload
}

func (o *GetV1NuncConnectionsNuncConnectionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NuncConnectionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
