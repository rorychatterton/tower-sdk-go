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

// NewJobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams creates a new JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams object
// with the default values initialized.
func NewJobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams() *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams {
	var ()
	return &JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJobTemplatesJobTemplatesNotificationTemplatesErrorCreateParamsWithTimeout creates a new JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobTemplatesJobTemplatesNotificationTemplatesErrorCreateParamsWithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams {
	var ()
	return &JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams{

		timeout: timeout,
	}
}

// NewJobTemplatesJobTemplatesNotificationTemplatesErrorCreateParamsWithContext creates a new JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobTemplatesJobTemplatesNotificationTemplatesErrorCreateParamsWithContext(ctx context.Context) *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams {
	var ()
	return &JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams{

		Context: ctx,
	}
}

// NewJobTemplatesJobTemplatesNotificationTemplatesErrorCreateParamsWithHTTPClient creates a new JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJobTemplatesJobTemplatesNotificationTemplatesErrorCreateParamsWithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams {
	var ()
	return &JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams{
		HTTPClient: client,
	}
}

/*JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams contains all the parameters to send to the API endpoint
for the job templates job templates notification templates error create operation typically these are written to a http.Request
*/
type JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams struct {

	/*Data*/
	Data JobTemplatesJobTemplatesNotificationTemplatesErrorCreateBody
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the job templates job templates notification templates error create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams) WithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job templates job templates notification templates error create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job templates job templates notification templates error create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams) WithContext(ctx context.Context) *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job templates job templates notification templates error create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job templates job templates notification templates error create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams) WithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job templates job templates notification templates error create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the job templates job templates notification templates error create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams) WithData(data JobTemplatesJobTemplatesNotificationTemplatesErrorCreateBody) *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the job templates job templates notification templates error create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams) SetData(data JobTemplatesJobTemplatesNotificationTemplatesErrorCreateBody) {
	o.Data = data
}

// WithID adds the id to the job templates job templates notification templates error create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams) WithID(id string) *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the job templates job templates notification templates error create params
func (o *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JobTemplatesJobTemplatesNotificationTemplatesErrorCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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