// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tinyci/ci-agents/ci-gen/gen/client/uisvc/models"
)

// GetLogoutReader is a Reader for the GetLogout structure.
type GetLogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 302:
		result := NewGetLogoutFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetLogoutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLogoutFound creates a GetLogoutFound with default headers values
func NewGetLogoutFound() *GetLogoutFound {
	return &GetLogoutFound{}
}

/*GetLogoutFound handles this case with default header values.

Redirection to another page indicates success.
*/
type GetLogoutFound struct {
}

func (o *GetLogoutFound) Error() string {
	return fmt.Sprintf("[GET /logout][%d] getLogoutFound ", 302) // #nosec
}

func (o *GetLogoutFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogoutInternalServerError creates a GetLogoutInternalServerError with default headers values
func NewGetLogoutInternalServerError() *GetLogoutInternalServerError {
	return &GetLogoutInternalServerError{}
}

/*GetLogoutInternalServerError handles this case with default header values.

An error occurred. Body has error result.
*/
type GetLogoutInternalServerError struct {
	Payload *models.Error
}

func (o *GetLogoutInternalServerError) Error() string {
	return fmt.Sprintf("[GET /logout][%d] getLogoutInternalServerError  %+v", 500, o.Payload) // #nosec
}

func (o *GetLogoutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
