// Code generated by go-swagger; DO NOT EDIT.

package ticketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1TicketingProjectsReader is a Reader for the GetV1TicketingProjects structure.
type GetV1TicketingProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1TicketingProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1TicketingProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1TicketingProjectsOK creates a GetV1TicketingProjectsOK with default headers values
func NewGetV1TicketingProjectsOK() *GetV1TicketingProjectsOK {
	return &GetV1TicketingProjectsOK{}
}

/* GetV1TicketingProjectsOK describes a response with status code 200, with default header values.

List all ticketing projects available to the organization
*/
type GetV1TicketingProjectsOK struct {
	Payload *models.ProjectEntity
}

func (o *GetV1TicketingProjectsOK) Error() string {
	return fmt.Sprintf("[GET /v1/ticketing/projects][%d] getV1TicketingProjectsOK  %+v", 200, o.Payload)
}
func (o *GetV1TicketingProjectsOK) GetPayload() *models.ProjectEntity {
	return o.Payload
}

func (o *GetV1TicketingProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}