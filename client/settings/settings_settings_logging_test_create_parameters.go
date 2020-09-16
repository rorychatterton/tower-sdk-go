// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewSettingsSettingsLoggingTestCreateParams creates a new SettingsSettingsLoggingTestCreateParams object
// with the default values initialized.
func NewSettingsSettingsLoggingTestCreateParams() *SettingsSettingsLoggingTestCreateParams {
	var ()
	return &SettingsSettingsLoggingTestCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSettingsSettingsLoggingTestCreateParamsWithTimeout creates a new SettingsSettingsLoggingTestCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSettingsSettingsLoggingTestCreateParamsWithTimeout(timeout time.Duration) *SettingsSettingsLoggingTestCreateParams {
	var ()
	return &SettingsSettingsLoggingTestCreateParams{

		timeout: timeout,
	}
}

// NewSettingsSettingsLoggingTestCreateParamsWithContext creates a new SettingsSettingsLoggingTestCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSettingsSettingsLoggingTestCreateParamsWithContext(ctx context.Context) *SettingsSettingsLoggingTestCreateParams {
	var ()
	return &SettingsSettingsLoggingTestCreateParams{

		Context: ctx,
	}
}

// NewSettingsSettingsLoggingTestCreateParamsWithHTTPClient creates a new SettingsSettingsLoggingTestCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSettingsSettingsLoggingTestCreateParamsWithHTTPClient(client *http.Client) *SettingsSettingsLoggingTestCreateParams {
	var ()
	return &SettingsSettingsLoggingTestCreateParams{
		HTTPClient: client,
	}
}

/*SettingsSettingsLoggingTestCreateParams contains all the parameters to send to the API endpoint
for the settings settings logging test create operation typically these are written to a http.Request
*/
type SettingsSettingsLoggingTestCreateParams struct {

	/*Data*/
	Data SettingsSettingsLoggingTestCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the settings settings logging test create params
func (o *SettingsSettingsLoggingTestCreateParams) WithTimeout(timeout time.Duration) *SettingsSettingsLoggingTestCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the settings settings logging test create params
func (o *SettingsSettingsLoggingTestCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the settings settings logging test create params
func (o *SettingsSettingsLoggingTestCreateParams) WithContext(ctx context.Context) *SettingsSettingsLoggingTestCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the settings settings logging test create params
func (o *SettingsSettingsLoggingTestCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the settings settings logging test create params
func (o *SettingsSettingsLoggingTestCreateParams) WithHTTPClient(client *http.Client) *SettingsSettingsLoggingTestCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the settings settings logging test create params
func (o *SettingsSettingsLoggingTestCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the settings settings logging test create params
func (o *SettingsSettingsLoggingTestCreateParams) WithData(data SettingsSettingsLoggingTestCreateBody) *SettingsSettingsLoggingTestCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the settings settings logging test create params
func (o *SettingsSettingsLoggingTestCreateParams) SetData(data SettingsSettingsLoggingTestCreateBody) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *SettingsSettingsLoggingTestCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
