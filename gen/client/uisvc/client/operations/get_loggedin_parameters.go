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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLoggedinParams creates a new GetLoggedinParams object
// with the default values initialized.
func NewGetLoggedinParams() *GetLoggedinParams {

	return &GetLoggedinParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoggedinParamsWithTimeout creates a new GetLoggedinParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLoggedinParamsWithTimeout(timeout time.Duration) *GetLoggedinParams {

	return &GetLoggedinParams{

		timeout: timeout,
	}
}

// NewGetLoggedinParamsWithContext creates a new GetLoggedinParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLoggedinParamsWithContext(ctx context.Context) *GetLoggedinParams {

	return &GetLoggedinParams{

		Context: ctx,
	}
}

// NewGetLoggedinParamsWithHTTPClient creates a new GetLoggedinParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLoggedinParamsWithHTTPClient(client *http.Client) *GetLoggedinParams {

	return &GetLoggedinParams{
		HTTPClient: client,
	}
}

/*GetLoggedinParams contains all the parameters to send to the API endpoint
for the get loggedin operation typically these are written to a http.Request
*/
type GetLoggedinParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get loggedin params
func (o *GetLoggedinParams) WithTimeout(timeout time.Duration) *GetLoggedinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get loggedin params
func (o *GetLoggedinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get loggedin params
func (o *GetLoggedinParams) WithContext(ctx context.Context) *GetLoggedinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get loggedin params
func (o *GetLoggedinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get loggedin params
func (o *GetLoggedinParams) WithHTTPClient(client *http.Client) *GetLoggedinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get loggedin params
func (o *GetLoggedinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoggedinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
