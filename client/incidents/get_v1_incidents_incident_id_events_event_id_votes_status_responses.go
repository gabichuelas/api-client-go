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

// GetV1IncidentsIncidentIDEventsEventIDVotesStatusReader is a Reader for the GetV1IncidentsIncidentIDEventsEventIDVotesStatus structure.
type GetV1IncidentsIncidentIDEventsEventIDVotesStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1IncidentsIncidentIDEventsEventIDVotesStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1IncidentsIncidentIDEventsEventIDVotesStatusOK creates a GetV1IncidentsIncidentIDEventsEventIDVotesStatusOK with default headers values
func NewGetV1IncidentsIncidentIDEventsEventIDVotesStatusOK() *GetV1IncidentsIncidentIDEventsEventIDVotesStatusOK {
	return &GetV1IncidentsIncidentIDEventsEventIDVotesStatusOK{}
}

/* GetV1IncidentsIncidentIDEventsEventIDVotesStatusOK describes a response with status code 200, with default header values.

Returns the status of the votes
*/
type GetV1IncidentsIncidentIDEventsEventIDVotesStatusOK struct {
	Payload *models.VotesEntity
}

func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/incidents/{incident_id}/events/{event_id}/votes/status][%d] getV1IncidentsIncidentIdEventsEventIdVotesStatusOK  %+v", 200, o.Payload)
}
func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusOK) GetPayload() *models.VotesEntity {
	return o.Payload
}

func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VotesEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}