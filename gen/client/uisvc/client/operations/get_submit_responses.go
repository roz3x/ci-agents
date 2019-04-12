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

// GetSubmitReader is a Reader for the GetSubmit structure.
type GetSubmitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubmitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubmitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetSubmitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSubmitOK creates a GetSubmitOK with default headers values
func NewGetSubmitOK() *GetSubmitOK {
	return &GetSubmitOK{}
}

/*GetSubmitOK handles this case with default header values.

OK
*/
type GetSubmitOK struct {
}

func (o *GetSubmitOK) Error() string {
	return fmt.Sprintf("[GET /submit][%d] getSubmitOK ", 200) // #nosec
}

func (o *GetSubmitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSubmitInternalServerError creates a GetSubmitInternalServerError with default headers values
func NewGetSubmitInternalServerError() *GetSubmitInternalServerError {
	return &GetSubmitInternalServerError{}
}

/*GetSubmitInternalServerError handles this case with default header values.

An error occurred. Body has error result.
*/
type GetSubmitInternalServerError struct {
	Payload *models.Error
}

func (o *GetSubmitInternalServerError) Error() string {
	return fmt.Sprintf("[GET /submit][%d] getSubmitInternalServerError  %+v", 500, o.Payload) // #nosec
}

func (o *GetSubmitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
