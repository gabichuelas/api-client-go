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

// DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDReader is a Reader for the DeleteV1PostMortemsReportsReportIDParticipantsParticipantID structure.
type DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1PostMortemsReportsReportIDParticipantsParticipantIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteV1PostMortemsReportsReportIDParticipantsParticipantIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1PostMortemsReportsReportIDParticipantsParticipantIDNoContent creates a DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDNoContent with default headers values
func NewDeleteV1PostMortemsReportsReportIDParticipantsParticipantIDNoContent() *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDNoContent {
	return &DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDNoContent{}
}

/* DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDNoContent describes a response with status code 204, with default header values.

Remove a participant from a post mortem report
*/
type DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDNoContent struct {
}

func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/post_mortems/reports/{report_id}/participants/{participant_id}][%d] deleteV1PostMortemsReportsReportIdParticipantsParticipantIdNoContent ", 204)
}

func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteV1PostMortemsReportsReportIDParticipantsParticipantIDBadRequest creates a DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDBadRequest with default headers values
func NewDeleteV1PostMortemsReportsReportIDParticipantsParticipantIDBadRequest() *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDBadRequest {
	return &DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDBadRequest{}
}

/* DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDBadRequest struct {
	Payload *models.ErrorEntity
}

func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/post_mortems/reports/{report_id}/participants/{participant_id}][%d] deleteV1PostMortemsReportsReportIdParticipantsParticipantIdBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDBadRequest) GetPayload() *models.ErrorEntity {
	return o.Payload
}

func (o *DeleteV1PostMortemsReportsReportIDParticipantsParticipantIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
