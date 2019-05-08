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

// NewGetUserPropertiesParams creates a new GetUserPropertiesParams object
// with the default values initialized.
func NewGetUserPropertiesParams() *GetUserPropertiesParams {

	return &GetUserPropertiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserPropertiesParamsWithTimeout creates a new GetUserPropertiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserPropertiesParamsWithTimeout(timeout time.Duration) *GetUserPropertiesParams {

	return &GetUserPropertiesParams{

		timeout: timeout,
	}
}

// NewGetUserPropertiesParamsWithContext creates a new GetUserPropertiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserPropertiesParamsWithContext(ctx context.Context) *GetUserPropertiesParams {

	return &GetUserPropertiesParams{

		Context: ctx,
	}
}

// NewGetUserPropertiesParamsWithHTTPClient creates a new GetUserPropertiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserPropertiesParamsWithHTTPClient(client *http.Client) *GetUserPropertiesParams {

	return &GetUserPropertiesParams{
		HTTPClient: client,
	}
}

/*GetUserPropertiesParams contains all the parameters to send to the API endpoint
for the get user properties operation typically these are written to a http.Request
*/
type GetUserPropertiesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user properties params
func (o *GetUserPropertiesParams) WithTimeout(timeout time.Duration) *GetUserPropertiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user properties params
func (o *GetUserPropertiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user properties params
func (o *GetUserPropertiesParams) WithContext(ctx context.Context) *GetUserPropertiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user properties params
func (o *GetUserPropertiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user properties params
func (o *GetUserPropertiesParams) WithHTTPClient(client *http.Client) *GetUserPropertiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user properties params
func (o *GetUserPropertiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserPropertiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserPropertiesParams) HandleQueryParam(r runtime.ClientRequest, reg strfmt.Registry, name string, param interface{}, formatter func(interface{}) string) error {
	if (reflect.TypeOf(param).Kind() == reflect.Ptr) && param == nil {
		return nil
	}

	if formatter == nil {
		formatter = func(i interface{}) string { return i.(string) }
	}

	return r.SetQueryParam(name, formatter(param))
}
