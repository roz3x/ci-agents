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

// DeleteTokenReader is a Reader for the DeleteToken structure.
type DeleteTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewDeleteTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteTokenOK creates a DeleteTokenOK with default headers values
func NewDeleteTokenOK() *DeleteTokenOK {
	return &DeleteTokenOK{}
}

/*DeleteTokenOK handles this case with default header values.

OK
*/
type DeleteTokenOK struct {
}

func (o *DeleteTokenOK) Error() string {
	return fmt.Sprintf("[DELETE /token][%d] deleteTokenOK ", 200) // #nosec
}

func (o *DeleteTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTokenInternalServerError creates a DeleteTokenInternalServerError with default headers values
func NewDeleteTokenInternalServerError() *DeleteTokenInternalServerError {
	return &DeleteTokenInternalServerError{}
}

/*DeleteTokenInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteTokenInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteTokenInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /token][%d] deleteTokenInternalServerError  %+v", 500, o.Payload) // #nosec
}

func (o *DeleteTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
