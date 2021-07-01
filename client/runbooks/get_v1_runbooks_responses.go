// Code generated by go-swagger; DO NOT EDIT.

package runbooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1RunbooksReader is a Reader for the GetV1Runbooks structure.
type GetV1RunbooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1RunbooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1RunbooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1RunbooksOK creates a GetV1RunbooksOK with default headers values
func NewGetV1RunbooksOK() *GetV1RunbooksOK {
	return &GetV1RunbooksOK{}
}

/* GetV1RunbooksOK describes a response with status code 200, with default header values.

Lists all available runbooks.
*/
type GetV1RunbooksOK struct {
	Payload *models.RunbookEntity
}

func (o *GetV1RunbooksOK) Error() string {
	return fmt.Sprintf("[GET /v1/runbooks][%d] getV1RunbooksOK  %+v", 200, o.Payload)
}
func (o *GetV1RunbooksOK) GetPayload() *models.RunbookEntity {
	return o.Payload
}

func (o *GetV1RunbooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
