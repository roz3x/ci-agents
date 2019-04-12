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

// GetLoggedinReader is a Reader for the GetLoggedin structure.
type GetLoggedinReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoggedinReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLoggedinOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetLoggedinInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLoggedinOK creates a GetLoggedinOK with default headers values
func NewGetLoggedinOK() *GetLoggedinOK {
	return &GetLoggedinOK{}
}

/*GetLoggedinOK handles this case with default header values.

Either "true" or the URL to the oauth challenge.
*/
type GetLoggedinOK struct {
	Payload string
}

func (o *GetLoggedinOK) Error() string {
	return fmt.Sprintf("[GET /loggedin][%d] getLoggedinOK  %+v", 200, o.Payload) // #nosec
}

func (o *GetLoggedinOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoggedinInternalServerError creates a GetLoggedinInternalServerError with default headers values
func NewGetLoggedinInternalServerError() *GetLoggedinInternalServerError {
	return &GetLoggedinInternalServerError{}
}

/*GetLoggedinInternalServerError handles this case with default header values.

An error occurred. Body has error result.
*/
type GetLoggedinInternalServerError struct {
	Payload *models.Error
}

func (o *GetLoggedinInternalServerError) Error() string {
	return fmt.Sprintf("[GET /loggedin][%d] getLoggedinInternalServerError  %+v", 500, o.Payload) // #nosec
}

func (o *GetLoggedinInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
