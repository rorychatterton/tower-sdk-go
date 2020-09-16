// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

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

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams creates a new WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams object
// with the default values initialized.
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams() *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams {
	var ()
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParamsWithTimeout creates a new WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParamsWithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams {
	var ()
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams{

		timeout: timeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParamsWithContext creates a new WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParamsWithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams {
	var ()
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams{

		Context: ctx,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParamsWithHTTPClient creates a new WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParamsWithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams {
	var ()
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams{
		HTTPClient: client,
	}
}

/*WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams contains all the parameters to send to the API endpoint
for the workflow job templates workflow job templates notification templates started create operation typically these are written to a http.Request
*/
type WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams struct {

	/*Data*/
	Data interface{}
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the workflow job templates workflow job templates notification templates started create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams) WithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job templates workflow job templates notification templates started create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job templates workflow job templates notification templates started create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams) WithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job templates workflow job templates notification templates started create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job templates workflow job templates notification templates started create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams) WithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job templates workflow job templates notification templates started create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the workflow job templates workflow job templates notification templates started create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams) WithData(data interface{}) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the workflow job templates workflow job templates notification templates started create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the workflow job templates workflow job templates notification templates started create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams) WithID(id string) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job templates workflow job templates notification templates started create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
