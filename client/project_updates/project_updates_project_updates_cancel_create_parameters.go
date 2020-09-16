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

// NewProjectUpdatesProjectUpdatesCancelCreateParams creates a new ProjectUpdatesProjectUpdatesCancelCreateParams object
// with the default values initialized.
func NewProjectUpdatesProjectUpdatesCancelCreateParams() *ProjectUpdatesProjectUpdatesCancelCreateParams {
	var ()
	return &ProjectUpdatesProjectUpdatesCancelCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectUpdatesProjectUpdatesCancelCreateParamsWithTimeout creates a new ProjectUpdatesProjectUpdatesCancelCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectUpdatesProjectUpdatesCancelCreateParamsWithTimeout(timeout time.Duration) *ProjectUpdatesProjectUpdatesCancelCreateParams {
	var ()
	return &ProjectUpdatesProjectUpdatesCancelCreateParams{

		timeout: timeout,
	}
}

// NewProjectUpdatesProjectUpdatesCancelCreateParamsWithContext creates a new ProjectUpdatesProjectUpdatesCancelCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectUpdatesProjectUpdatesCancelCreateParamsWithContext(ctx context.Context) *ProjectUpdatesProjectUpdatesCancelCreateParams {
	var ()
	return &ProjectUpdatesProjectUpdatesCancelCreateParams{

		Context: ctx,
	}
}

// NewProjectUpdatesProjectUpdatesCancelCreateParamsWithHTTPClient creates a new ProjectUpdatesProjectUpdatesCancelCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectUpdatesProjectUpdatesCancelCreateParamsWithHTTPClient(client *http.Client) *ProjectUpdatesProjectUpdatesCancelCreateParams {
	var ()
	return &ProjectUpdatesProjectUpdatesCancelCreateParams{
		HTTPClient: client,
	}
}

/*ProjectUpdatesProjectUpdatesCancelCreateParams contains all the parameters to send to the API endpoint
for the project updates project updates cancel create operation typically these are written to a http.Request
*/
type ProjectUpdatesProjectUpdatesCancelCreateParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project updates project updates cancel create params
func (o *ProjectUpdatesProjectUpdatesCancelCreateParams) WithTimeout(timeout time.Duration) *ProjectUpdatesProjectUpdatesCancelCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project updates project updates cancel create params
func (o *ProjectUpdatesProjectUpdatesCancelCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project updates project updates cancel create params
func (o *ProjectUpdatesProjectUpdatesCancelCreateParams) WithContext(ctx context.Context) *ProjectUpdatesProjectUpdatesCancelCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project updates project updates cancel create params
func (o *ProjectUpdatesProjectUpdatesCancelCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project updates project updates cancel create params
func (o *ProjectUpdatesProjectUpdatesCancelCreateParams) WithHTTPClient(client *http.Client) *ProjectUpdatesProjectUpdatesCancelCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project updates project updates cancel create params
func (o *ProjectUpdatesProjectUpdatesCancelCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the project updates project updates cancel create params
func (o *ProjectUpdatesProjectUpdatesCancelCreateParams) WithID(id string) *ProjectUpdatesProjectUpdatesCancelCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the project updates project updates cancel create params
func (o *ProjectUpdatesProjectUpdatesCancelCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectUpdatesProjectUpdatesCancelCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
