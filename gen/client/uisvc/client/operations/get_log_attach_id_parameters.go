// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLogAttachIDParams creates a new GetLogAttachIDParams object
// with the default values initialized.
func NewGetLogAttachIDParams() *GetLogAttachIDParams {
	var ()
	return &GetLogAttachIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogAttachIDParamsWithTimeout creates a new GetLogAttachIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLogAttachIDParamsWithTimeout(timeout time.Duration) *GetLogAttachIDParams {
	var ()
	return &GetLogAttachIDParams{

		timeout: timeout,
	}
}

// NewGetLogAttachIDParamsWithContext creates a new GetLogAttachIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLogAttachIDParamsWithContext(ctx context.Context) *GetLogAttachIDParams {
	var ()
	return &GetLogAttachIDParams{

		Context: ctx,
	}
}

// NewGetLogAttachIDParamsWithHTTPClient creates a new GetLogAttachIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLogAttachIDParamsWithHTTPClient(client *http.Client) *GetLogAttachIDParams {
	var ()
	return &GetLogAttachIDParams{
		HTTPClient: client,
	}
}

/*GetLogAttachIDParams contains all the parameters to send to the API endpoint
for the get log attach ID operation typically these are written to a http.Request
*/
type GetLogAttachIDParams struct {

	/*ID
	  The ID of the run to retrieve the log for.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get log attach ID params
func (o *GetLogAttachIDParams) WithTimeout(timeout time.Duration) *GetLogAttachIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get log attach ID params
func (o *GetLogAttachIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get log attach ID params
func (o *GetLogAttachIDParams) WithContext(ctx context.Context) *GetLogAttachIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get log attach ID params
func (o *GetLogAttachIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get log attach ID params
func (o *GetLogAttachIDParams) WithHTTPClient(client *http.Client) *GetLogAttachIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get log attach ID params
func (o *GetLogAttachIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get log attach ID params
func (o *GetLogAttachIDParams) WithID(id int64) *GetLogAttachIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get log attach ID params
func (o *GetLogAttachIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogAttachIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
