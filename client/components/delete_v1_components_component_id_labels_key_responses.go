// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteV1ComponentsComponentIDLabelsKeyReader is a Reader for the DeleteV1ComponentsComponentIDLabelsKey structure.
type DeleteV1ComponentsComponentIDLabelsKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1ComponentsComponentIDLabelsKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteV1ComponentsComponentIDLabelsKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteV1ComponentsComponentIDLabelsKeyNoContent creates a DeleteV1ComponentsComponentIDLabelsKeyNoContent with default headers values
func NewDeleteV1ComponentsComponentIDLabelsKeyNoContent() *DeleteV1ComponentsComponentIDLabelsKeyNoContent {
	return &DeleteV1ComponentsComponentIDLabelsKeyNoContent{}
}

/*DeleteV1ComponentsComponentIDLabelsKeyNoContent handles this case with default header values.

Remove a label from a component
*/
type DeleteV1ComponentsComponentIDLabelsKeyNoContent struct {
}

func (o *DeleteV1ComponentsComponentIDLabelsKeyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/components/{component_id}/labels/{key}][%d] deleteV1ComponentsComponentIdLabelsKeyNoContent ", 204)
}

func (o *DeleteV1ComponentsComponentIDLabelsKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}