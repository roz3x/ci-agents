// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tinyci/ci-agents/gen/client/uisvc/models"
)

// GetRunRunIDReader is a Reader for the GetRunRunID structure.
type GetRunRunIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunRunIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRunRunIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetRunRunIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRunRunIDOK creates a GetRunRunIDOK with default headers values
func NewGetRunRunIDOK() *GetRunRunIDOK {
	return &GetRunRunIDOK{}
}

/*GetRunRunIDOK handles this case with default header values.

OK
*/
type GetRunRunIDOK struct {
	Payload *models.Run
}

func (o *GetRunRunIDOK) Error() string {
	return fmt.Sprintf("[GET /run/{run_id}][%d] getRunRunIdOK  %+v", 200, o.Payload) // #nosec
}

func (o *GetRunRunIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Run)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunRunIDInternalServerError creates a GetRunRunIDInternalServerError with default headers values
func NewGetRunRunIDInternalServerError() *GetRunRunIDInternalServerError {
	return &GetRunRunIDInternalServerError{}
}

/*GetRunRunIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetRunRunIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetRunRunIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /run/{run_id}][%d] getRunRunIdInternalServerError  %+v", 500, o.Payload) // #nosec
}

func (o *GetRunRunIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
