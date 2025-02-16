// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PatchV1IntegrationsAwsCloudtrailBatchesIDReader is a Reader for the PatchV1IntegrationsAwsCloudtrailBatchesID structure.
type PatchV1IntegrationsAwsCloudtrailBatchesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1IntegrationsAwsCloudtrailBatchesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1IntegrationsAwsCloudtrailBatchesIDOK creates a PatchV1IntegrationsAwsCloudtrailBatchesIDOK with default headers values
func NewPatchV1IntegrationsAwsCloudtrailBatchesIDOK() *PatchV1IntegrationsAwsCloudtrailBatchesIDOK {
	return &PatchV1IntegrationsAwsCloudtrailBatchesIDOK{}
}

/* PatchV1IntegrationsAwsCloudtrailBatchesIDOK describes a response with status code 200, with default header values.

Update a CloudTrail batch with new information.
*/
type PatchV1IntegrationsAwsCloudtrailBatchesIDOK struct {
	Payload *models.CloudtrailBatchEntity
}

func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/integrations/aws/cloudtrail_batches/{id}][%d] patchV1IntegrationsAwsCloudtrailBatchesIdOK  %+v", 200, o.Payload)
}
func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDOK) GetPayload() *models.CloudtrailBatchEntity {
	return o.Payload
}

func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudtrailBatchEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
