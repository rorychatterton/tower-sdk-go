// Code generated by go-swagger; DO NOT EDIT.

package project_updates

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

// NewProjectUpdatesProjectUpdatesStdoutReadParams creates a new ProjectUpdatesProjectUpdatesStdoutReadParams object
// with the default values initialized.
func NewProjectUpdatesProjectUpdatesStdoutReadParams() *ProjectUpdatesProjectUpdatesStdoutReadParams {
	var ()
	return &ProjectUpdatesProjectUpdatesStdoutReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectUpdatesProjectUpdatesStdoutReadParamsWithTimeout creates a new ProjectUpdatesProjectUpdatesStdoutReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectUpdatesProjectUpdatesStdoutReadParamsWithTimeout(timeout time.Duration) *ProjectUpdatesProjectUpdatesStdoutReadParams {
	var ()
	return &ProjectUpdatesProjectUpdatesStdoutReadParams{

		timeout: timeout,
	}
}

// NewProjectUpdatesProjectUpdatesStdoutReadParamsWithContext creates a new ProjectUpdatesProjectUpdatesStdoutReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectUpdatesProjectUpdatesStdoutReadParamsWithContext(ctx context.Context) *ProjectUpdatesProjectUpdatesStdoutReadParams {
	var ()
	return &ProjectUpdatesProjectUpdatesStdoutReadParams{

		Context: ctx,
	}
}

// NewProjectUpdatesProjectUpdatesStdoutReadParamsWithHTTPClient creates a new ProjectUpdatesProjectUpdatesStdoutReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectUpdatesProjectUpdatesStdoutReadParamsWithHTTPClient(client *http.Client) *ProjectUpdatesProjectUpdatesStdoutReadParams {
	var ()
	return &ProjectUpdatesProjectUpdatesStdoutReadParams{
		HTTPClient: client,
	}
}

/*ProjectUpdatesProjectUpdatesStdoutReadParams contains all the parameters to send to the API endpoint
for the project updates project updates stdout read operation typically these are written to a http.Request
*/
type ProjectUpdatesProjectUpdatesStdoutReadParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project updates project updates stdout read params
func (o *ProjectUpdatesProjectUpdatesStdoutReadParams) WithTimeout(timeout time.Duration) *ProjectUpdatesProjectUpdatesStdoutReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project updates project updates stdout read params
func (o *ProjectUpdatesProjectUpdatesStdoutReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project updates project updates stdout read params
func (o *ProjectUpdatesProjectUpdatesStdoutReadParams) WithContext(ctx context.Context) *ProjectUpdatesProjectUpdatesStdoutReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project updates project updates stdout read params
func (o *ProjectUpdatesProjectUpdatesStdoutReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project updates project updates stdout read params
func (o *ProjectUpdatesProjectUpdatesStdoutReadParams) WithHTTPClient(client *http.Client) *ProjectUpdatesProjectUpdatesStdoutReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project updates project updates stdout read params
func (o *ProjectUpdatesProjectUpdatesStdoutReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the project updates project updates stdout read params
func (o *ProjectUpdatesProjectUpdatesStdoutReadParams) WithID(id string) *ProjectUpdatesProjectUpdatesStdoutReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the project updates project updates stdout read params
func (o *ProjectUpdatesProjectUpdatesStdoutReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectUpdatesProjectUpdatesStdoutReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
