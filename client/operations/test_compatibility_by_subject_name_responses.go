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

// TestCompatibilityBySubjectNameReader is a Reader for the TestCompatibilityBySubjectName structure.
type TestCompatibilityBySubjectNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestCompatibilityBySubjectNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestCompatibilityBySubjectNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewTestCompatibilityBySubjectNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewTestCompatibilityBySubjectNameUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTestCompatibilityBySubjectNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTestCompatibilityBySubjectNameOK creates a TestCompatibilityBySubjectNameOK with default headers values
func NewTestCompatibilityBySubjectNameOK() *TestCompatibilityBySubjectNameOK {
	return &TestCompatibilityBySubjectNameOK{}
}

/*TestCompatibilityBySubjectNameOK handles this case with default header values.

successful operation
*/
type TestCompatibilityBySubjectNameOK struct {
	Payload *models.CompatibilityCheckResponse
}

func (o *TestCompatibilityBySubjectNameOK) Error() string {
	return fmt.Sprintf("[POST /compatibility/subjects/{subject}/versions/{version}][%d] testCompatibilityBySubjectNameOK  %+v", 200, o.Payload)
}

func (o *TestCompatibilityBySubjectNameOK) GetPayload() *models.CompatibilityCheckResponse {
	return o.Payload
}

func (o *TestCompatibilityBySubjectNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CompatibilityCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestCompatibilityBySubjectNameNotFound creates a TestCompatibilityBySubjectNameNotFound with default headers values
func NewTestCompatibilityBySubjectNameNotFound() *TestCompatibilityBySubjectNameNotFound {
	return &TestCompatibilityBySubjectNameNotFound{}
}

/*TestCompatibilityBySubjectNameNotFound handles this case with default header values.

Error code 40401 -- Subject not found
Error code 40402 -- Version not found
*/
type TestCompatibilityBySubjectNameNotFound struct {
}

func (o *TestCompatibilityBySubjectNameNotFound) Error() string {
	return fmt.Sprintf("[POST /compatibility/subjects/{subject}/versions/{version}][%d] testCompatibilityBySubjectNameNotFound ", 404)
}

func (o *TestCompatibilityBySubjectNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTestCompatibilityBySubjectNameUnprocessableEntity creates a TestCompatibilityBySubjectNameUnprocessableEntity with default headers values
func NewTestCompatibilityBySubjectNameUnprocessableEntity() *TestCompatibilityBySubjectNameUnprocessableEntity {
	return &TestCompatibilityBySubjectNameUnprocessableEntity{}
}

/*TestCompatibilityBySubjectNameUnprocessableEntity handles this case with default header values.

Error code 42201 -- Invalid schema or schema type
Error code 42202 -- Invalid version
*/
type TestCompatibilityBySubjectNameUnprocessableEntity struct {
}

func (o *TestCompatibilityBySubjectNameUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /compatibility/subjects/{subject}/versions/{version}][%d] testCompatibilityBySubjectNameUnprocessableEntity ", 422)
}

func (o *TestCompatibilityBySubjectNameUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTestCompatibilityBySubjectNameInternalServerError creates a TestCompatibilityBySubjectNameInternalServerError with default headers values
func NewTestCompatibilityBySubjectNameInternalServerError() *TestCompatibilityBySubjectNameInternalServerError {
	return &TestCompatibilityBySubjectNameInternalServerError{}
}

/*TestCompatibilityBySubjectNameInternalServerError handles this case with default header values.

Error code 50001 -- Error in the backend data store
*/
type TestCompatibilityBySubjectNameInternalServerError struct {
}

func (o *TestCompatibilityBySubjectNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /compatibility/subjects/{subject}/versions/{version}][%d] testCompatibilityBySubjectNameInternalServerError ", 500)
}

func (o *TestCompatibilityBySubjectNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}