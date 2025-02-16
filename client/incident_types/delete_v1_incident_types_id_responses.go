// Code generated by go-swagger; DO NOT EDIT.

package incident_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1IncidentTypesIDReader is a Reader for the DeleteV1IncidentTypesID structure.
type DeleteV1IncidentTypesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1IncidentTypesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1IncidentTypesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1IncidentTypesIDNoContent creates a DeleteV1IncidentTypesIDNoContent with default headers values
func NewDeleteV1IncidentTypesIDNoContent() *DeleteV1IncidentTypesIDNoContent {
	return &DeleteV1IncidentTypesIDNoContent{}
}

/* DeleteV1IncidentTypesIDNoContent describes a response with status code 204, with default header values.

Archive an incident type
*/
type DeleteV1IncidentTypesIDNoContent struct {
}

func (o *DeleteV1IncidentTypesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/incident_types/{id}][%d] deleteV1IncidentTypesIdNoContent ", 204)
}

func (o *DeleteV1IncidentTypesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
