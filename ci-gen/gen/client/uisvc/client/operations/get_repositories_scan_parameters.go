// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"reflect"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRepositoriesScanParams creates a new GetRepositoriesScanParams object
// with the default values initialized.
func NewGetRepositoriesScanParams() *GetRepositoriesScanParams {

	return &GetRepositoriesScanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesScanParamsWithTimeout creates a new GetRepositoriesScanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesScanParamsWithTimeout(timeout time.Duration) *GetRepositoriesScanParams {

	return &GetRepositoriesScanParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesScanParamsWithContext creates a new GetRepositoriesScanParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesScanParamsWithContext(ctx context.Context) *GetRepositoriesScanParams {

	return &GetRepositoriesScanParams{

		Context: ctx,
	}
}

// NewGetRepositoriesScanParamsWithHTTPClient creates a new GetRepositoriesScanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesScanParamsWithHTTPClient(client *http.Client) *GetRepositoriesScanParams {

	return &GetRepositoriesScanParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesScanParams contains all the parameters to send to the API endpoint
for the get repositories scan operation typically these are written to a http.Request
*/
type GetRepositoriesScanParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get repositories scan params
func (o *GetRepositoriesScanParams) WithTimeout(timeout time.Duration) *GetRepositoriesScanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories scan params
func (o *GetRepositoriesScanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories scan params
func (o *GetRepositoriesScanParams) WithContext(ctx context.Context) *GetRepositoriesScanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories scan params
func (o *GetRepositoriesScanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories scan params
func (o *GetRepositoriesScanParams) WithHTTPClient(client *http.Client) *GetRepositoriesScanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories scan params
func (o *GetRepositoriesScanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesScanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRepositoriesScanParams) HandleQueryParam(r runtime.ClientRequest, reg strfmt.Registry, name string, param interface{}, formatter func(interface{}) string) error {
	if (reflect.TypeOf(param).Kind() == reflect.Ptr) && param == nil {
		return nil
	}

	if formatter == nil {
		formatter = func(i interface{}) string { return i.(string) }
	}

	return r.SetQueryParam(name, formatter(param))
}
