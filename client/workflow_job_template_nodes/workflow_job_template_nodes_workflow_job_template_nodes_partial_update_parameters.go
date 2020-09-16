// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_template_nodes

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

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams object
// with the default values initialized.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams() *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams {
	var ()
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParamsWithTimeout creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParamsWithTimeout(timeout time.Duration) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams {
	var ()
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParamsWithContext creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParamsWithContext(ctx context.Context) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams {
	var ()
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams{

		Context: ctx,
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParamsWithHTTPClient creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParamsWithHTTPClient(client *http.Client) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams {
	var ()
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams contains all the parameters to send to the API endpoint
for the workflow job template nodes workflow job template nodes partial update operation typically these are written to a http.Request
*/
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams struct {

	/*Data*/
	Data WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateBody
	/*ID*/
	ID string
	/*Search
	  A search term.

	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the workflow job template nodes workflow job template nodes partial update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams) WithTimeout(timeout time.Duration) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job template nodes workflow job template nodes partial update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job template nodes workflow job template nodes partial update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams) WithContext(ctx context.Context) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job template nodes workflow job template nodes partial update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job template nodes workflow job template nodes partial update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams) WithHTTPClient(client *http.Client) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job template nodes workflow job template nodes partial update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the workflow job template nodes workflow job template nodes partial update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams) WithData(data WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateBody) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the workflow job template nodes workflow job template nodes partial update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams) SetData(data WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateBody) {
	o.Data = data
}

// WithID adds the id to the workflow job template nodes workflow job template nodes partial update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams) WithID(id string) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job template nodes workflow job template nodes partial update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the workflow job template nodes workflow job template nodes partial update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams) WithSearch(search *string) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the workflow job template nodes workflow job template nodes partial update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Search != nil {

		// query param search
		var qrSearch string
		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {
			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}