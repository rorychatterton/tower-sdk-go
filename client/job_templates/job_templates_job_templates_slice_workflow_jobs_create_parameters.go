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

// NewJobTemplatesJobTemplatesSliceWorkflowJobsCreateParams creates a new JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams object
// with the default values initialized.
func NewJobTemplatesJobTemplatesSliceWorkflowJobsCreateParams() *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams {
	var ()
	return &JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJobTemplatesJobTemplatesSliceWorkflowJobsCreateParamsWithTimeout creates a new JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobTemplatesJobTemplatesSliceWorkflowJobsCreateParamsWithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams {
	var ()
	return &JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams{

		timeout: timeout,
	}
}

// NewJobTemplatesJobTemplatesSliceWorkflowJobsCreateParamsWithContext creates a new JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobTemplatesJobTemplatesSliceWorkflowJobsCreateParamsWithContext(ctx context.Context) *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams {
	var ()
	return &JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams{

		Context: ctx,
	}
}

// NewJobTemplatesJobTemplatesSliceWorkflowJobsCreateParamsWithHTTPClient creates a new JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJobTemplatesJobTemplatesSliceWorkflowJobsCreateParamsWithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams {
	var ()
	return &JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams{
		HTTPClient: client,
	}
}

/*JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams contains all the parameters to send to the API endpoint
for the job templates job templates slice workflow jobs create operation typically these are written to a http.Request
*/
type JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams struct {

	/*Data*/
	Data JobTemplatesJobTemplatesSliceWorkflowJobsCreateBody
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the job templates job templates slice workflow jobs create params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams) WithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job templates job templates slice workflow jobs create params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job templates job templates slice workflow jobs create params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams) WithContext(ctx context.Context) *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job templates job templates slice workflow jobs create params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job templates job templates slice workflow jobs create params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams) WithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job templates job templates slice workflow jobs create params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the job templates job templates slice workflow jobs create params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams) WithData(data JobTemplatesJobTemplatesSliceWorkflowJobsCreateBody) *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the job templates job templates slice workflow jobs create params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams) SetData(data JobTemplatesJobTemplatesSliceWorkflowJobsCreateBody) {
	o.Data = data
}

// WithID adds the id to the job templates job templates slice workflow jobs create params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams) WithID(id string) *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the job templates job templates slice workflow jobs create params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
