// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PatchV1IncidentsIncidentIDActionItemsActionItemIDReader is a Reader for the PatchV1IncidentsIncidentIDActionItemsActionItemID structure.
type PatchV1IncidentsIncidentIDActionItemsActionItemIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1IncidentsIncidentIDActionItemsActionItemIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1IncidentsIncidentIDActionItemsActionItemIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1IncidentsIncidentIDActionItemsActionItemIDOK creates a PatchV1IncidentsIncidentIDActionItemsActionItemIDOK with default headers values
func NewPatchV1IncidentsIncidentIDActionItemsActionItemIDOK() *PatchV1IncidentsIncidentIDActionItemsActionItemIDOK {
	return &PatchV1IncidentsIncidentIDActionItemsActionItemIDOK{}
}

/* PatchV1IncidentsIncidentIDActionItemsActionItemIDOK describes a response with status code 200, with default header values.

patched ActionItem
*/
type PatchV1IncidentsIncidentIDActionItemsActionItemIDOK struct {
}

func (o *PatchV1IncidentsIncidentIDActionItemsActionItemIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/incidents/{incident_id}/action_items/{action_item_id}][%d] patchV1IncidentsIncidentIdActionItemsActionItemIdOK ", 200)
}

func (o *PatchV1IncidentsIncidentIDActionItemsActionItemIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
