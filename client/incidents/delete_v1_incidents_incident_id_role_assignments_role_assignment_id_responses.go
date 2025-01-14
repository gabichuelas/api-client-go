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

// DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDReader is a Reader for the DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentID structure.
type DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDOK creates a DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDOK with default headers values
func NewDeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDOK() *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDOK {
	return &DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDOK{}
}

/* DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDOK describes a response with status code 200, with default header values.

Unassign a role from a user. Any tasks that were created on the incident will remain in whatever state they are in
*/
type DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDOK struct {
	Payload *models.RoleAssignmentEntity
}

func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/incidents/{incident_id}/role_assignments/{role_assignment_id}][%d] deleteV1IncidentsIncidentIdRoleAssignmentsRoleAssignmentIdOK  %+v", 200, o.Payload)
}
func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDOK) GetPayload() *models.RoleAssignmentEntity {
	return o.Payload
}

func (o *DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleAssignmentEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
