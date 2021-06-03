// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDReader is a Reader for the DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventID structure.
type DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDOK creates a DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDOK with default headers values
func NewDeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDOK() *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDOK {
	return &DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDOK{}
}

/* DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDOK describes a response with status code 200, with default header values.

Delete a report event by its incident ID
*/
type DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDOK struct {
	Payload *models.EventEntity
}

func (o *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/post_mortems/reports/{report_id}/events/from_incident/{incident_event_id}][%d] deleteV1PostMortemsReportsReportIdEventsFromIncidentIncidentEventIdOK  %+v", 200, o.Payload)
}
func (o *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDOK) GetPayload() *models.EventEntity {
	return o.Payload
}

func (o *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}