// Code generated by go-swagger; DO NOT EDIT.

package runbook_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostV1RunbookTemplatesIDRunbooksReader is a Reader for the PostV1RunbookTemplatesIDRunbooks structure.
type PostV1RunbookTemplatesIDRunbooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1RunbookTemplatesIDRunbooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1RunbookTemplatesIDRunbooksCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1RunbookTemplatesIDRunbooksCreated creates a PostV1RunbookTemplatesIDRunbooksCreated with default headers values
func NewPostV1RunbookTemplatesIDRunbooksCreated() *PostV1RunbookTemplatesIDRunbooksCreated {
	return &PostV1RunbookTemplatesIDRunbooksCreated{}
}

/* PostV1RunbookTemplatesIDRunbooksCreated describes a response with status code 201, with default header values.

created Runbook
*/
type PostV1RunbookTemplatesIDRunbooksCreated struct {
}

func (o *PostV1RunbookTemplatesIDRunbooksCreated) Error() string {
	return fmt.Sprintf("[POST /v1/runbook_templates/{id}/runbooks][%d] postV1RunbookTemplatesIdRunbooksCreated ", 201)
}

func (o *PostV1RunbookTemplatesIDRunbooksCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
