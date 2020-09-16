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

// NewJobTemplatesJobTemplatesSurveySpecCreateParams creates a new JobTemplatesJobTemplatesSurveySpecCreateParams object
// with the default values initialized.
func NewJobTemplatesJobTemplatesSurveySpecCreateParams() *JobTemplatesJobTemplatesSurveySpecCreateParams {
	var ()
	return &JobTemplatesJobTemplatesSurveySpecCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJobTemplatesJobTemplatesSurveySpecCreateParamsWithTimeout creates a new JobTemplatesJobTemplatesSurveySpecCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobTemplatesJobTemplatesSurveySpecCreateParamsWithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesSurveySpecCreateParams {
	var ()
	return &JobTemplatesJobTemplatesSurveySpecCreateParams{

		timeout: timeout,
	}
}

// NewJobTemplatesJobTemplatesSurveySpecCreateParamsWithContext creates a new JobTemplatesJobTemplatesSurveySpecCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobTemplatesJobTemplatesSurveySpecCreateParamsWithContext(ctx context.Context) *JobTemplatesJobTemplatesSurveySpecCreateParams {
	var ()
	return &JobTemplatesJobTemplatesSurveySpecCreateParams{

		Context: ctx,
	}
}

// NewJobTemplatesJobTemplatesSurveySpecCreateParamsWithHTTPClient creates a new JobTemplatesJobTemplatesSurveySpecCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJobTemplatesJobTemplatesSurveySpecCreateParamsWithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesSurveySpecCreateParams {
	var ()
	return &JobTemplatesJobTemplatesSurveySpecCreateParams{
		HTTPClient: client,
	}
}

/*JobTemplatesJobTemplatesSurveySpecCreateParams contains all the parameters to send to the API endpoint
for the job templates job templates survey spec create operation typically these are written to a http.Request
*/
type JobTemplatesJobTemplatesSurveySpecCreateParams struct {

	/*Data*/
	Data interface{}
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the job templates job templates survey spec create params
func (o *JobTemplatesJobTemplatesSurveySpecCreateParams) WithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesSurveySpecCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job templates job templates survey spec create params
func (o *JobTemplatesJobTemplatesSurveySpecCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job templates job templates survey spec create params
func (o *JobTemplatesJobTemplatesSurveySpecCreateParams) WithContext(ctx context.Context) *JobTemplatesJobTemplatesSurveySpecCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job templates job templates survey spec create params
func (o *JobTemplatesJobTemplatesSurveySpecCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job templates job templates survey spec create params
func (o *JobTemplatesJobTemplatesSurveySpecCreateParams) WithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesSurveySpecCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job templates job templates survey spec create params
func (o *JobTemplatesJobTemplatesSurveySpecCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the job templates job templates survey spec create params
func (o *JobTemplatesJobTemplatesSurveySpecCreateParams) WithData(data interface{}) *JobTemplatesJobTemplatesSurveySpecCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the job templates job templates survey spec create params
func (o *JobTemplatesJobTemplatesSurveySpecCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the job templates job templates survey spec create params
func (o *JobTemplatesJobTemplatesSurveySpecCreateParams) WithID(id string) *JobTemplatesJobTemplatesSurveySpecCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the job templates job templates survey spec create params
func (o *JobTemplatesJobTemplatesSurveySpecCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JobTemplatesJobTemplatesSurveySpecCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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