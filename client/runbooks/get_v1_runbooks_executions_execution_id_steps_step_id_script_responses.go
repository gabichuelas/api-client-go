// Code generated by go-swagger; DO NOT EDIT.

package runbooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptReader is a Reader for the GetV1RunbooksExecutionsExecutionIDStepsStepIDScript structure.
type GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1RunbooksExecutionsExecutionIDStepsStepIDScriptOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1RunbooksExecutionsExecutionIDStepsStepIDScriptOK creates a GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptOK with default headers values
func NewGetV1RunbooksExecutionsExecutionIDStepsStepIDScriptOK() *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptOK {
	return &GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptOK{}
}

/* GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptOK describes a response with status code 200, with default header values.

Retrieves the bash script from a "script" step.
*/
type GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptOK struct {
	Payload *models.ExecutionEntity
}

func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptOK) Error() string {
	return fmt.Sprintf("[GET /v1/runbooks/executions/{execution_id}/steps/{step_id}/script][%d] getV1RunbooksExecutionsExecutionIdStepsStepIdScriptOK  %+v", 200, o.Payload)
}
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptOK) GetPayload() *models.ExecutionEntity {
	return o.Payload
}

func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExecutionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
