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

// GetTokenReader is a Reader for the GetToken structure.
type GetTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTokenOK creates a GetTokenOK with default headers values
func NewGetTokenOK() *GetTokenOK {
	return &GetTokenOK{}
}

/*GetTokenOK handles this case with default header values.

Returns the token which you can pass as a bearer token in headers for further requests as this user.

*/
type GetTokenOK struct {
	Payload string
}

func (o *GetTokenOK) Error() string {
	return fmt.Sprintf("[GET /token][%d] getTokenOK  %+v", 200, o.Payload) // #nosec
}

func (o *GetTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokenInternalServerError creates a GetTokenInternalServerError with default headers values
func NewGetTokenInternalServerError() *GetTokenInternalServerError {
	return &GetTokenInternalServerError{}
}

/*GetTokenInternalServerError handles this case with default header values.

An error occurred.
*/
type GetTokenInternalServerError struct {
	Payload *models.Error
}

func (o *GetTokenInternalServerError) Error() string {
	return fmt.Sprintf("[GET /token][%d] getTokenInternalServerError  %+v", 500, o.Payload) // #nosec
}

func (o *GetTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
