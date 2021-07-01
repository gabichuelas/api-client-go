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

// GetV1IncidentsIncidentIDUsersUserIDReader is a Reader for the GetV1IncidentsIncidentIDUsersUserID structure.
type GetV1IncidentsIncidentIDUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IncidentsIncidentIDUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1IncidentsIncidentIDUsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1IncidentsIncidentIDUsersUserIDOK creates a GetV1IncidentsIncidentIDUsersUserIDOK with default headers values
func NewGetV1IncidentsIncidentIDUsersUserIDOK() *GetV1IncidentsIncidentIDUsersUserIDOK {
	return &GetV1IncidentsIncidentIDUsersUserIDOK{}
}

/* GetV1IncidentsIncidentIDUsersUserIDOK describes a response with status code 200, with default header values.

Retrieve a user with current roles for an incident
*/
type GetV1IncidentsIncidentIDUsersUserIDOK struct {
	Payload *models.RoleAssignmentEntity
}

func (o *GetV1IncidentsIncidentIDUsersUserIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/incidents/{incident_id}/users/{user_id}][%d] getV1IncidentsIncidentIdUsersUserIdOK  %+v", 200, o.Payload)
}
func (o *GetV1IncidentsIncidentIDUsersUserIDOK) GetPayload() *models.RoleAssignmentEntity {
	return o.Payload
}

func (o *GetV1IncidentsIncidentIDUsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleAssignmentEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
