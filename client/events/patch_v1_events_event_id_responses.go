// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PatchV1EventsEventIDReader is a Reader for the PatchV1EventsEventID structure.
type PatchV1EventsEventIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1EventsEventIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1EventsEventIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1EventsEventIDOK creates a PatchV1EventsEventIDOK with default headers values
func NewPatchV1EventsEventIDOK() *PatchV1EventsEventIDOK {
	return &PatchV1EventsEventIDOK{}
}

/* PatchV1EventsEventIDOK describes a response with status code 200, with default header values.

Update an incident event
*/
type PatchV1EventsEventIDOK struct {
	Payload *models.IncidentEventEntity
}

func (o *PatchV1EventsEventIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/events/{event_id}][%d] patchV1EventsEventIdOK  %+v", 200, o.Payload)
}
func (o *PatchV1EventsEventIDOK) GetPayload() *models.IncidentEventEntity {
	return o.Payload
}

func (o *PatchV1EventsEventIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncidentEventEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
