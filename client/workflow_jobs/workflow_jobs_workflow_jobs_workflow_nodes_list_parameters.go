// Code generated by go-swagger; DO NOT EDIT.

package workflow_jobs

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

// NewWorkflowJobsWorkflowJobsWorkflowNodesListParams creates a new WorkflowJobsWorkflowJobsWorkflowNodesListParams object
// with the default values initialized.
func NewWorkflowJobsWorkflowJobsWorkflowNodesListParams() *WorkflowJobsWorkflowJobsWorkflowNodesListParams {
	var ()
	return &WorkflowJobsWorkflowJobsWorkflowNodesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobsWorkflowJobsWorkflowNodesListParamsWithTimeout creates a new WorkflowJobsWorkflowJobsWorkflowNodesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWorkflowJobsWorkflowJobsWorkflowNodesListParamsWithTimeout(timeout time.Duration) *WorkflowJobsWorkflowJobsWorkflowNodesListParams {
	var ()
	return &WorkflowJobsWorkflowJobsWorkflowNodesListParams{

		timeout: timeout,
	}
}

// NewWorkflowJobsWorkflowJobsWorkflowNodesListParamsWithContext creates a new WorkflowJobsWorkflowJobsWorkflowNodesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewWorkflowJobsWorkflowJobsWorkflowNodesListParamsWithContext(ctx context.Context) *WorkflowJobsWorkflowJobsWorkflowNodesListParams {
	var ()
	return &WorkflowJobsWorkflowJobsWorkflowNodesListParams{

		Context: ctx,
	}
}

// NewWorkflowJobsWorkflowJobsWorkflowNodesListParamsWithHTTPClient creates a new WorkflowJobsWorkflowJobsWorkflowNodesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWorkflowJobsWorkflowJobsWorkflowNodesListParamsWithHTTPClient(client *http.Client) *WorkflowJobsWorkflowJobsWorkflowNodesListParams {
	var ()
	return &WorkflowJobsWorkflowJobsWorkflowNodesListParams{
		HTTPClient: client,
	}
}

/*WorkflowJobsWorkflowJobsWorkflowNodesListParams contains all the parameters to send to the API endpoint
for the workflow jobs workflow jobs workflow nodes list operation typically these are written to a http.Request
*/
type WorkflowJobsWorkflowJobsWorkflowNodesListParams struct {

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

// WithTimeout adds the timeout to the workflow jobs workflow jobs workflow nodes list params
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListParams) WithTimeout(timeout time.Duration) *WorkflowJobsWorkflowJobsWorkflowNodesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow jobs workflow jobs workflow nodes list params
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow jobs workflow jobs workflow nodes list params
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListParams) WithContext(ctx context.Context) *WorkflowJobsWorkflowJobsWorkflowNodesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow jobs workflow jobs workflow nodes list params
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow jobs workflow jobs workflow nodes list params
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListParams) WithHTTPClient(client *http.Client) *WorkflowJobsWorkflowJobsWorkflowNodesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow jobs workflow jobs workflow nodes list params
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the workflow jobs workflow jobs workflow nodes list params
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListParams) WithID(id string) *WorkflowJobsWorkflowJobsWorkflowNodesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow jobs workflow jobs workflow nodes list params
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the workflow jobs workflow jobs workflow nodes list params
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListParams) WithPage(page *int64) *WorkflowJobsWorkflowJobsWorkflowNodesListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the workflow jobs workflow jobs workflow nodes list params
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the workflow jobs workflow jobs workflow nodes list params
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListParams) WithPageSize(pageSize *int64) *WorkflowJobsWorkflowJobsWorkflowNodesListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the workflow jobs workflow jobs workflow nodes list params
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the workflow jobs workflow jobs workflow nodes list params
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListParams) WithSearch(search *string) *WorkflowJobsWorkflowJobsWorkflowNodesListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the workflow jobs workflow jobs workflow nodes list params
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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