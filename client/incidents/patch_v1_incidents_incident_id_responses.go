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

// PatchV1IncidentsIncidentIDReader is a Reader for the PatchV1IncidentsIncidentID structure.
type PatchV1IncidentsIncidentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1IncidentsIncidentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1IncidentsIncidentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1IncidentsIncidentIDOK creates a PatchV1IncidentsIncidentIDOK with default headers values
func NewPatchV1IncidentsIncidentIDOK() *PatchV1IncidentsIncidentIDOK {
	return &PatchV1IncidentsIncidentIDOK{}
}

/* PatchV1IncidentsIncidentIDOK describes a response with status code 200, with default header values.

Update an incident
*/
type PatchV1IncidentsIncidentIDOK struct {
	Payload *models.IncidentEntity
}

func (o *PatchV1IncidentsIncidentIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/incidents/{incident_id}][%d] patchV1IncidentsIncidentIdOK  %+v", 200, o.Payload)
}
func (o *PatchV1IncidentsIncidentIDOK) GetPayload() *models.IncidentEntity {
	return o.Payload
}

func (o *PatchV1IncidentsIncidentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncidentEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
