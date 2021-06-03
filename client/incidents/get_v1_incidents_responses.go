// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1IncidentsReader is a Reader for the GetV1Incidents structure.
type GetV1IncidentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IncidentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1IncidentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1IncidentsOK creates a GetV1IncidentsOK with default headers values
func NewGetV1IncidentsOK() *GetV1IncidentsOK {
	return &GetV1IncidentsOK{}
}

/* GetV1IncidentsOK describes a response with status code 200, with default header values.

Retrieve incidents
*/
type GetV1IncidentsOK struct {
	Payload *models.IncidentEntityPaginated
}

func (o *GetV1IncidentsOK) Error() string {
	return fmt.Sprintf("[GET /v1/incidents][%d] getV1IncidentsOK  %+v", 200, o.Payload)
}
func (o *GetV1IncidentsOK) GetPayload() *models.IncidentEntityPaginated {
	return o.Payload
}

func (o *GetV1IncidentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncidentEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}