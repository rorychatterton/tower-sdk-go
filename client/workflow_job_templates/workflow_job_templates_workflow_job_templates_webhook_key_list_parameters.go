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
	"github.com/go-openapi/swag"
)

// NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams creates a new WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams object
// with the default values initialized.
func NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams() *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams {
	var ()
	return &WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParamsWithTimeout creates a new WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParamsWithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams {
	var ()
	return &WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams{

		timeout: timeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParamsWithContext creates a new WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams object
// with the default values initialized, and the ability to set a context for a request
func NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParamsWithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams {
	var ()
	return &WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams{

		Context: ctx,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParamsWithHTTPClient creates a new WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParamsWithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams {
	var ()
	return &WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams{
		HTTPClient: client,
	}
}

/*WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams contains all the parameters to send to the API endpoint
for the workflow job templates workflow job templates webhook key list operation typically these are written to a http.Request
*/
type WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams struct {

	/*ID*/
	ID string
	/*Page
	  A page number within the paginated result set.

	*/
	Page *int64
	/*PageSize
	  Number of results to return per page.

	*/
	PageSize *int64
	/*Search
	  A search term.

	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the workflow job templates workflow job templates webhook key list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams) WithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job templates workflow job templates webhook key list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job templates workflow job templates webhook key list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams) WithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job templates workflow job templates webhook key list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job templates workflow job templates webhook key list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams) WithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job templates workflow job templates webhook key list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the workflow job templates workflow job templates webhook key list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams) WithID(id string) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job templates workflow job templates webhook key list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the workflow job templates workflow job templates webhook key list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams) WithPage(page *int64) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the workflow job templates workflow job templates webhook key list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the workflow job templates workflow job templates webhook key list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams) WithPageSize(pageSize *int64) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the workflow job templates workflow job templates webhook key list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the workflow job templates workflow job templates webhook key list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams) WithSearch(search *string) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the workflow job templates workflow job templates webhook key list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}

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
