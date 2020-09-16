// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_nodes

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

// NewWorkflowJobNodesWorkflowJobNodesCredentialsListParams creates a new WorkflowJobNodesWorkflowJobNodesCredentialsListParams object
// with the default values initialized.
func NewWorkflowJobNodesWorkflowJobNodesCredentialsListParams() *WorkflowJobNodesWorkflowJobNodesCredentialsListParams {
	var ()
	return &WorkflowJobNodesWorkflowJobNodesCredentialsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobNodesWorkflowJobNodesCredentialsListParamsWithTimeout creates a new WorkflowJobNodesWorkflowJobNodesCredentialsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWorkflowJobNodesWorkflowJobNodesCredentialsListParamsWithTimeout(timeout time.Duration) *WorkflowJobNodesWorkflowJobNodesCredentialsListParams {
	var ()
	return &WorkflowJobNodesWorkflowJobNodesCredentialsListParams{

		timeout: timeout,
	}
}

// NewWorkflowJobNodesWorkflowJobNodesCredentialsListParamsWithContext creates a new WorkflowJobNodesWorkflowJobNodesCredentialsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewWorkflowJobNodesWorkflowJobNodesCredentialsListParamsWithContext(ctx context.Context) *WorkflowJobNodesWorkflowJobNodesCredentialsListParams {
	var ()
	return &WorkflowJobNodesWorkflowJobNodesCredentialsListParams{

		Context: ctx,
	}
}

// NewWorkflowJobNodesWorkflowJobNodesCredentialsListParamsWithHTTPClient creates a new WorkflowJobNodesWorkflowJobNodesCredentialsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWorkflowJobNodesWorkflowJobNodesCredentialsListParamsWithHTTPClient(client *http.Client) *WorkflowJobNodesWorkflowJobNodesCredentialsListParams {
	var ()
	return &WorkflowJobNodesWorkflowJobNodesCredentialsListParams{
		HTTPClient: client,
	}
}

/*WorkflowJobNodesWorkflowJobNodesCredentialsListParams contains all the parameters to send to the API endpoint
for the workflow job nodes workflow job nodes credentials list operation typically these are written to a http.Request
*/
type WorkflowJobNodesWorkflowJobNodesCredentialsListParams struct {

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

// WithTimeout adds the timeout to the workflow job nodes workflow job nodes credentials list params
func (o *WorkflowJobNodesWorkflowJobNodesCredentialsListParams) WithTimeout(timeout time.Duration) *WorkflowJobNodesWorkflowJobNodesCredentialsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job nodes workflow job nodes credentials list params
func (o *WorkflowJobNodesWorkflowJobNodesCredentialsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job nodes workflow job nodes credentials list params
func (o *WorkflowJobNodesWorkflowJobNodesCredentialsListParams) WithContext(ctx context.Context) *WorkflowJobNodesWorkflowJobNodesCredentialsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job nodes workflow job nodes credentials list params
func (o *WorkflowJobNodesWorkflowJobNodesCredentialsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job nodes workflow job nodes credentials list params
func (o *WorkflowJobNodesWorkflowJobNodesCredentialsListParams) WithHTTPClient(client *http.Client) *WorkflowJobNodesWorkflowJobNodesCredentialsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job nodes workflow job nodes credentials list params
func (o *WorkflowJobNodesWorkflowJobNodesCredentialsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the workflow job nodes workflow job nodes credentials list params
func (o *WorkflowJobNodesWorkflowJobNodesCredentialsListParams) WithID(id string) *WorkflowJobNodesWorkflowJobNodesCredentialsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job nodes workflow job nodes credentials list params
func (o *WorkflowJobNodesWorkflowJobNodesCredentialsListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the workflow job nodes workflow job nodes credentials list params
func (o *WorkflowJobNodesWorkflowJobNodesCredentialsListParams) WithPage(page *int64) *WorkflowJobNodesWorkflowJobNodesCredentialsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the workflow job nodes workflow job nodes credentials list params
func (o *WorkflowJobNodesWorkflowJobNodesCredentialsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the workflow job nodes workflow job nodes credentials list params
func (o *WorkflowJobNodesWorkflowJobNodesCredentialsListParams) WithPageSize(pageSize *int64) *WorkflowJobNodesWorkflowJobNodesCredentialsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the workflow job nodes workflow job nodes credentials list params
func (o *WorkflowJobNodesWorkflowJobNodesCredentialsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the workflow job nodes workflow job nodes credentials list params
func (o *WorkflowJobNodesWorkflowJobNodesCredentialsListParams) WithSearch(search *string) *WorkflowJobNodesWorkflowJobNodesCredentialsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the workflow job nodes workflow job nodes credentials list params
func (o *WorkflowJobNodesWorkflowJobNodesCredentialsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobNodesWorkflowJobNodesCredentialsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
