// Code generated by go-swagger; DO NOT EDIT.

package ad_hoc_commands

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

// NewAdHocCommandsAdHocCommandsRelaunchCreateParams creates a new AdHocCommandsAdHocCommandsRelaunchCreateParams object
// with the default values initialized.
func NewAdHocCommandsAdHocCommandsRelaunchCreateParams() *AdHocCommandsAdHocCommandsRelaunchCreateParams {
	var ()
	return &AdHocCommandsAdHocCommandsRelaunchCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdHocCommandsAdHocCommandsRelaunchCreateParamsWithTimeout creates a new AdHocCommandsAdHocCommandsRelaunchCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdHocCommandsAdHocCommandsRelaunchCreateParamsWithTimeout(timeout time.Duration) *AdHocCommandsAdHocCommandsRelaunchCreateParams {
	var ()
	return &AdHocCommandsAdHocCommandsRelaunchCreateParams{

		timeout: timeout,
	}
}

// NewAdHocCommandsAdHocCommandsRelaunchCreateParamsWithContext creates a new AdHocCommandsAdHocCommandsRelaunchCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdHocCommandsAdHocCommandsRelaunchCreateParamsWithContext(ctx context.Context) *AdHocCommandsAdHocCommandsRelaunchCreateParams {
	var ()
	return &AdHocCommandsAdHocCommandsRelaunchCreateParams{

		Context: ctx,
	}
}

// NewAdHocCommandsAdHocCommandsRelaunchCreateParamsWithHTTPClient creates a new AdHocCommandsAdHocCommandsRelaunchCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdHocCommandsAdHocCommandsRelaunchCreateParamsWithHTTPClient(client *http.Client) *AdHocCommandsAdHocCommandsRelaunchCreateParams {
	var ()
	return &AdHocCommandsAdHocCommandsRelaunchCreateParams{
		HTTPClient: client,
	}
}

/*AdHocCommandsAdHocCommandsRelaunchCreateParams contains all the parameters to send to the API endpoint
for the ad hoc commands ad hoc commands relaunch create operation typically these are written to a http.Request
*/
type AdHocCommandsAdHocCommandsRelaunchCreateParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ad hoc commands ad hoc commands relaunch create params
func (o *AdHocCommandsAdHocCommandsRelaunchCreateParams) WithTimeout(timeout time.Duration) *AdHocCommandsAdHocCommandsRelaunchCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ad hoc commands ad hoc commands relaunch create params
func (o *AdHocCommandsAdHocCommandsRelaunchCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ad hoc commands ad hoc commands relaunch create params
func (o *AdHocCommandsAdHocCommandsRelaunchCreateParams) WithContext(ctx context.Context) *AdHocCommandsAdHocCommandsRelaunchCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ad hoc commands ad hoc commands relaunch create params
func (o *AdHocCommandsAdHocCommandsRelaunchCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ad hoc commands ad hoc commands relaunch create params
func (o *AdHocCommandsAdHocCommandsRelaunchCreateParams) WithHTTPClient(client *http.Client) *AdHocCommandsAdHocCommandsRelaunchCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ad hoc commands ad hoc commands relaunch create params
func (o *AdHocCommandsAdHocCommandsRelaunchCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ad hoc commands ad hoc commands relaunch create params
func (o *AdHocCommandsAdHocCommandsRelaunchCreateParams) WithID(id string) *AdHocCommandsAdHocCommandsRelaunchCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ad hoc commands ad hoc commands relaunch create params
func (o *AdHocCommandsAdHocCommandsRelaunchCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AdHocCommandsAdHocCommandsRelaunchCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
