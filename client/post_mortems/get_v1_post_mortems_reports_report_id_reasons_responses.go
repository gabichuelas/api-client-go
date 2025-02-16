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

// GetV1PostMortemsReportsReportIDReasonsReader is a Reader for the GetV1PostMortemsReportsReportIDReasons structure.
type GetV1PostMortemsReportsReportIDReasonsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1PostMortemsReportsReportIDReasonsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1PostMortemsReportsReportIDReasonsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1PostMortemsReportsReportIDReasonsOK creates a GetV1PostMortemsReportsReportIDReasonsOK with default headers values
func NewGetV1PostMortemsReportsReportIDReasonsOK() *GetV1PostMortemsReportsReportIDReasonsOK {
	return &GetV1PostMortemsReportsReportIDReasonsOK{}
}

/* GetV1PostMortemsReportsReportIDReasonsOK describes a response with status code 200, with default header values.

Retrieve post mortem report reasons
*/
type GetV1PostMortemsReportsReportIDReasonsOK struct {
	Payload *models.ReasonEntityPaginated
}

func (o *GetV1PostMortemsReportsReportIDReasonsOK) Error() string {
	return fmt.Sprintf("[GET /v1/post_mortems/reports/{report_id}/reasons][%d] getV1PostMortemsReportsReportIdReasonsOK  %+v", 200, o.Payload)
}
func (o *GetV1PostMortemsReportsReportIDReasonsOK) GetPayload() *models.ReasonEntityPaginated {
	return o.Payload
}

func (o *GetV1PostMortemsReportsReportIDReasonsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReasonEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
