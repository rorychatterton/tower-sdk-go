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

// NewJobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams creates a new JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams object
// with the default values initialized.
func NewJobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams() *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams {
	var ()
	return &JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJobTemplatesJobTemplatesNotificationTemplatesStartedCreateParamsWithTimeout creates a new JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobTemplatesJobTemplatesNotificationTemplatesStartedCreateParamsWithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams {
	var ()
	return &JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams{

		timeout: timeout,
	}
}

// NewJobTemplatesJobTemplatesNotificationTemplatesStartedCreateParamsWithContext creates a new JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobTemplatesJobTemplatesNotificationTemplatesStartedCreateParamsWithContext(ctx context.Context) *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams {
	var ()
	return &JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams{

		Context: ctx,
	}
}

// NewJobTemplatesJobTemplatesNotificationTemplatesStartedCreateParamsWithHTTPClient creates a new JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJobTemplatesJobTemplatesNotificationTemplatesStartedCreateParamsWithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams {
	var ()
	return &JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams{
		HTTPClient: client,
	}
}

/*JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams contains all the parameters to send to the API endpoint
for the job templates job templates notification templates started create operation typically these are written to a http.Request
*/
type JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams struct {

	/*Data*/
	Data interface{}
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the job templates job templates notification templates started create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams) WithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job templates job templates notification templates started create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job templates job templates notification templates started create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams) WithContext(ctx context.Context) *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job templates job templates notification templates started create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job templates job templates notification templates started create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams) WithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job templates job templates notification templates started create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the job templates job templates notification templates started create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams) WithData(data interface{}) *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the job templates job templates notification templates started create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the job templates job templates notification templates started create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams) WithID(id string) *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the job templates job templates notification templates started create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JobTemplatesJobTemplatesNotificationTemplatesStartedCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
