// Code generated by go-swagger; DO NOT EDIT.

package job_templates

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

// NewJobTemplatesJobTemplatesCredentialsCreateParams creates a new JobTemplatesJobTemplatesCredentialsCreateParams object
// with the default values initialized.
func NewJobTemplatesJobTemplatesCredentialsCreateParams() *JobTemplatesJobTemplatesCredentialsCreateParams {
	var ()
	return &JobTemplatesJobTemplatesCredentialsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJobTemplatesJobTemplatesCredentialsCreateParamsWithTimeout creates a new JobTemplatesJobTemplatesCredentialsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobTemplatesJobTemplatesCredentialsCreateParamsWithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesCredentialsCreateParams {
	var ()
	return &JobTemplatesJobTemplatesCredentialsCreateParams{

		timeout: timeout,
	}
}

// NewJobTemplatesJobTemplatesCredentialsCreateParamsWithContext creates a new JobTemplatesJobTemplatesCredentialsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobTemplatesJobTemplatesCredentialsCreateParamsWithContext(ctx context.Context) *JobTemplatesJobTemplatesCredentialsCreateParams {
	var ()
	return &JobTemplatesJobTemplatesCredentialsCreateParams{

		Context: ctx,
	}
}

// NewJobTemplatesJobTemplatesCredentialsCreateParamsWithHTTPClient creates a new JobTemplatesJobTemplatesCredentialsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJobTemplatesJobTemplatesCredentialsCreateParamsWithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesCredentialsCreateParams {
	var ()
	return &JobTemplatesJobTemplatesCredentialsCreateParams{
		HTTPClient: client,
	}
}

/*JobTemplatesJobTemplatesCredentialsCreateParams contains all the parameters to send to the API endpoint
for the job templates job templates credentials create operation typically these are written to a http.Request
*/
type JobTemplatesJobTemplatesCredentialsCreateParams struct {

	/*Data*/
	Data interface{}
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the job templates job templates credentials create params
func (o *JobTemplatesJobTemplatesCredentialsCreateParams) WithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesCredentialsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job templates job templates credentials create params
func (o *JobTemplatesJobTemplatesCredentialsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job templates job templates credentials create params
func (o *JobTemplatesJobTemplatesCredentialsCreateParams) WithContext(ctx context.Context) *JobTemplatesJobTemplatesCredentialsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job templates job templates credentials create params
func (o *JobTemplatesJobTemplatesCredentialsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job templates job templates credentials create params
func (o *JobTemplatesJobTemplatesCredentialsCreateParams) WithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesCredentialsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job templates job templates credentials create params
func (o *JobTemplatesJobTemplatesCredentialsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the job templates job templates credentials create params
func (o *JobTemplatesJobTemplatesCredentialsCreateParams) WithData(data interface{}) *JobTemplatesJobTemplatesCredentialsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the job templates job templates credentials create params
func (o *JobTemplatesJobTemplatesCredentialsCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the job templates job templates credentials create params
func (o *JobTemplatesJobTemplatesCredentialsCreateParams) WithID(id string) *JobTemplatesJobTemplatesCredentialsCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the job templates job templates credentials create params
func (o *JobTemplatesJobTemplatesCredentialsCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JobTemplatesJobTemplatesCredentialsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
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
