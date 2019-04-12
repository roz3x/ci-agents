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

// GetRepositoriesSubAddOwnerRepoReader is a Reader for the GetRepositoriesSubAddOwnerRepo structure.
type GetRepositoriesSubAddOwnerRepoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesSubAddOwnerRepoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRepositoriesSubAddOwnerRepoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetRepositoriesSubAddOwnerRepoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRepositoriesSubAddOwnerRepoOK creates a GetRepositoriesSubAddOwnerRepoOK with default headers values
func NewGetRepositoriesSubAddOwnerRepoOK() *GetRepositoriesSubAddOwnerRepoOK {
	return &GetRepositoriesSubAddOwnerRepoOK{}
}

/*GetRepositoriesSubAddOwnerRepoOK handles this case with default header values.

OK
*/
type GetRepositoriesSubAddOwnerRepoOK struct {
}

func (o *GetRepositoriesSubAddOwnerRepoOK) Error() string {
	return fmt.Sprintf("[GET /repositories/sub/add/{owner}/{repo}][%d] getRepositoriesSubAddOwnerRepoOK ", 200) // #nosec
}

func (o *GetRepositoriesSubAddOwnerRepoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepositoriesSubAddOwnerRepoInternalServerError creates a GetRepositoriesSubAddOwnerRepoInternalServerError with default headers values
func NewGetRepositoriesSubAddOwnerRepoInternalServerError() *GetRepositoriesSubAddOwnerRepoInternalServerError {
	return &GetRepositoriesSubAddOwnerRepoInternalServerError{}
}

/*GetRepositoriesSubAddOwnerRepoInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetRepositoriesSubAddOwnerRepoInternalServerError struct {
	Payload *models.Error
}

func (o *GetRepositoriesSubAddOwnerRepoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /repositories/sub/add/{owner}/{repo}][%d] getRepositoriesSubAddOwnerRepoInternalServerError  %+v", 500, o.Payload) // #nosec
}

func (o *GetRepositoriesSubAddOwnerRepoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
