// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firehydrant/api-client-go/models"
)

// PostV1ComponentsReader is a Reader for the PostV1Components structure.
type PostV1ComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostV1ComponentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostV1ComponentsCreated creates a PostV1ComponentsCreated with default headers values
func NewPostV1ComponentsCreated() *PostV1ComponentsCreated {
	return &PostV1ComponentsCreated{}
}

/*PostV1ComponentsCreated handles this case with default header values.

Creates an component
*/
type PostV1ComponentsCreated struct {
	Payload *models.ComponentEntity
}

func (o *PostV1ComponentsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/components][%d] postV1ComponentsCreated  %+v", 201, o.Payload)
}

func (o *PostV1ComponentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComponentEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}