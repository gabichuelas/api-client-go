// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1PostMortemsReportsReportIDReasonsReasonIDReader is a Reader for the DeleteV1PostMortemsReportsReportIDReasonsReasonID structure.
type DeleteV1PostMortemsReportsReportIDReasonsReasonIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1PostMortemsReportsReportIDReasonsReasonIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1PostMortemsReportsReportIDReasonsReasonIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1PostMortemsReportsReportIDReasonsReasonIDNoContent creates a DeleteV1PostMortemsReportsReportIDReasonsReasonIDNoContent with default headers values
func NewDeleteV1PostMortemsReportsReportIDReasonsReasonIDNoContent() *DeleteV1PostMortemsReportsReportIDReasonsReasonIDNoContent {
	return &DeleteV1PostMortemsReportsReportIDReasonsReasonIDNoContent{}
}

/* DeleteV1PostMortemsReportsReportIDReasonsReasonIDNoContent describes a response with status code 204, with default header values.

Deletes a reason from a report
*/
type DeleteV1PostMortemsReportsReportIDReasonsReasonIDNoContent struct {
}

func (o *DeleteV1PostMortemsReportsReportIDReasonsReasonIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/post_mortems/reports/{report_id}/reasons/{reason_id}][%d] deleteV1PostMortemsReportsReportIdReasonsReasonIdNoContent ", 204)
}

func (o *DeleteV1PostMortemsReportsReportIDReasonsReasonIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
