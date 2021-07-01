// Code generated by go-swagger; DO NOT EDIT.

package functionalities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// DeleteV1FunctionalitiesFunctionalityIDReader is a Reader for the DeleteV1FunctionalitiesFunctionalityID structure.
type DeleteV1FunctionalitiesFunctionalityIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1FunctionalitiesFunctionalityIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1FunctionalitiesFunctionalityIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1FunctionalitiesFunctionalityIDOK creates a DeleteV1FunctionalitiesFunctionalityIDOK with default headers values
func NewDeleteV1FunctionalitiesFunctionalityIDOK() *DeleteV1FunctionalitiesFunctionalityIDOK {
	return &DeleteV1FunctionalitiesFunctionalityIDOK{}
}

/* DeleteV1FunctionalitiesFunctionalityIDOK describes a response with status code 200, with default header values.

Archive a functionality
*/
type DeleteV1FunctionalitiesFunctionalityIDOK struct {
	Payload *models.FunctionalityEntity
}

func (o *DeleteV1FunctionalitiesFunctionalityIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/functionalities/{functionality_id}][%d] deleteV1FunctionalitiesFunctionalityIdOK  %+v", 200, o.Payload)
}
func (o *DeleteV1FunctionalitiesFunctionalityIDOK) GetPayload() *models.FunctionalityEntity {
	return o.Payload
}

func (o *DeleteV1FunctionalitiesFunctionalityIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FunctionalityEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
