// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dblooman/schema-registry-client/models"
)

// UpdateSubjectLevelConfigReader is a Reader for the UpdateSubjectLevelConfig structure.
type UpdateSubjectLevelConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSubjectLevelConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSubjectLevelConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewUpdateSubjectLevelConfigUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateSubjectLevelConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSubjectLevelConfigOK creates a UpdateSubjectLevelConfigOK with default headers values
func NewUpdateSubjectLevelConfigOK() *UpdateSubjectLevelConfigOK {
	return &UpdateSubjectLevelConfigOK{}
}

/*UpdateSubjectLevelConfigOK handles this case with default header values.

successful operation
*/
type UpdateSubjectLevelConfigOK struct {
	Payload *models.ConfigUpdateRequest
}

func (o *UpdateSubjectLevelConfigOK) Error() string {
	return fmt.Sprintf("[PUT /config/{subject}][%d] updateSubjectLevelConfigOK  %+v", 200, o.Payload)
}

func (o *UpdateSubjectLevelConfigOK) GetPayload() *models.ConfigUpdateRequest {
	return o.Payload
}

func (o *UpdateSubjectLevelConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigUpdateRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubjectLevelConfigUnprocessableEntity creates a UpdateSubjectLevelConfigUnprocessableEntity with default headers values
func NewUpdateSubjectLevelConfigUnprocessableEntity() *UpdateSubjectLevelConfigUnprocessableEntity {
	return &UpdateSubjectLevelConfigUnprocessableEntity{}
}

/*UpdateSubjectLevelConfigUnprocessableEntity handles this case with default header values.

Error code 42203 -- Invalid compatibility level
Error code 40402 -- Version not found
*/
type UpdateSubjectLevelConfigUnprocessableEntity struct {
}

func (o *UpdateSubjectLevelConfigUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /config/{subject}][%d] updateSubjectLevelConfigUnprocessableEntity ", 422)
}

func (o *UpdateSubjectLevelConfigUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSubjectLevelConfigInternalServerError creates a UpdateSubjectLevelConfigInternalServerError with default headers values
func NewUpdateSubjectLevelConfigInternalServerError() *UpdateSubjectLevelConfigInternalServerError {
	return &UpdateSubjectLevelConfigInternalServerError{}
}

/*UpdateSubjectLevelConfigInternalServerError handles this case with default header values.

Error code 50001 -- Error in the backend data store
Error code 50003 -- Error while forwarding the request to the primary
*/
type UpdateSubjectLevelConfigInternalServerError struct {
}

func (o *UpdateSubjectLevelConfigInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /config/{subject}][%d] updateSubjectLevelConfigInternalServerError ", 500)
}

func (o *UpdateSubjectLevelConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
