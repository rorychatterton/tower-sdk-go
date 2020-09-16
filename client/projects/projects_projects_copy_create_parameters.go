// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewProjectsProjectsCopyCreateParams creates a new ProjectsProjectsCopyCreateParams object
// with the default values initialized.
func NewProjectsProjectsCopyCreateParams() *ProjectsProjectsCopyCreateParams {
	var ()
	return &ProjectsProjectsCopyCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsProjectsCopyCreateParamsWithTimeout creates a new ProjectsProjectsCopyCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsProjectsCopyCreateParamsWithTimeout(timeout time.Duration) *ProjectsProjectsCopyCreateParams {
	var ()
	return &ProjectsProjectsCopyCreateParams{

		timeout: timeout,
	}
}

// NewProjectsProjectsCopyCreateParamsWithContext creates a new ProjectsProjectsCopyCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsProjectsCopyCreateParamsWithContext(ctx context.Context) *ProjectsProjectsCopyCreateParams {
	var ()
	return &ProjectsProjectsCopyCreateParams{

		Context: ctx,
	}
}

// NewProjectsProjectsCopyCreateParamsWithHTTPClient creates a new ProjectsProjectsCopyCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsProjectsCopyCreateParamsWithHTTPClient(client *http.Client) *ProjectsProjectsCopyCreateParams {
	var ()
	return &ProjectsProjectsCopyCreateParams{
		HTTPClient: client,
	}
}

/*ProjectsProjectsCopyCreateParams contains all the parameters to send to the API endpoint
for the projects projects copy create operation typically these are written to a http.Request
*/
type ProjectsProjectsCopyCreateParams struct {

	/*Data*/
	Data ProjectsProjectsCopyCreateBody
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects projects copy create params
func (o *ProjectsProjectsCopyCreateParams) WithTimeout(timeout time.Duration) *ProjectsProjectsCopyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects projects copy create params
func (o *ProjectsProjectsCopyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects projects copy create params
func (o *ProjectsProjectsCopyCreateParams) WithContext(ctx context.Context) *ProjectsProjectsCopyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects projects copy create params
func (o *ProjectsProjectsCopyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects projects copy create params
func (o *ProjectsProjectsCopyCreateParams) WithHTTPClient(client *http.Client) *ProjectsProjectsCopyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects projects copy create params
func (o *ProjectsProjectsCopyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects projects copy create params
func (o *ProjectsProjectsCopyCreateParams) WithData(data ProjectsProjectsCopyCreateBody) *ProjectsProjectsCopyCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects projects copy create params
func (o *ProjectsProjectsCopyCreateParams) SetData(data ProjectsProjectsCopyCreateBody) {
	o.Data = data
}

// WithID adds the id to the projects projects copy create params
func (o *ProjectsProjectsCopyCreateParams) WithID(id string) *ProjectsProjectsCopyCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects projects copy create params
func (o *ProjectsProjectsCopyCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsProjectsCopyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
