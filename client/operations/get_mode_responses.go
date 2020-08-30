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

// GetModeReader is a Reader for the GetMode structure.
type GetModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetModeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetModeOK creates a GetModeOK with default headers values
func NewGetModeOK() *GetModeOK {
	return &GetModeOK{}
}

/*GetModeOK handles this case with default header values.

successful operation
*/
type GetModeOK struct {
	Payload *models.ModeGetResponse
}

func (o *GetModeOK) Error() string {
	return fmt.Sprintf("[GET /mode/{subject}][%d] getModeOK  %+v", 200, o.Payload)
}

func (o *GetModeOK) GetPayload() *models.ModeGetResponse {
	return o.Payload
}

func (o *GetModeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModeGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}