// Code generated by go-swagger; DO NOT EDIT.

package severities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1SeveritiesReader is a Reader for the PostV1Severities structure.
type PostV1SeveritiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1SeveritiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1SeveritiesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1SeveritiesCreated creates a PostV1SeveritiesCreated with default headers values
func NewPostV1SeveritiesCreated() *PostV1SeveritiesCreated {
	return &PostV1SeveritiesCreated{}
}

/* PostV1SeveritiesCreated describes a response with status code 201, with default header values.

Create a new severity
*/
type PostV1SeveritiesCreated struct {
	Payload *models.SeverityEntity
}

func (o *PostV1SeveritiesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/severities][%d] postV1SeveritiesCreated  %+v", 201, o.Payload)
}
func (o *PostV1SeveritiesCreated) GetPayload() *models.SeverityEntity {
	return o.Payload
}

func (o *PostV1SeveritiesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SeverityEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
