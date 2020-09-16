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

// NewWorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams creates a new WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams object
// with the default values initialized.
func NewWorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams() *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams {
	var ()
	return &WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobNodesWorkflowJobNodesAlwaysNodesListParamsWithTimeout creates a new WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWorkflowJobNodesWorkflowJobNodesAlwaysNodesListParamsWithTimeout(timeout time.Duration) *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams {
	var ()
	return &WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams{

		timeout: timeout,
	}
}

// NewWorkflowJobNodesWorkflowJobNodesAlwaysNodesListParamsWithContext creates a new WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewWorkflowJobNodesWorkflowJobNodesAlwaysNodesListParamsWithContext(ctx context.Context) *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams {
	var ()
	return &WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams{

		Context: ctx,
	}
}

// NewWorkflowJobNodesWorkflowJobNodesAlwaysNodesListParamsWithHTTPClient creates a new WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWorkflowJobNodesWorkflowJobNodesAlwaysNodesListParamsWithHTTPClient(client *http.Client) *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams {
	var ()
	return &WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams{
		HTTPClient: client,
	}
}

/*WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams contains all the parameters to send to the API endpoint
for the workflow job nodes workflow job nodes always nodes list operation typically these are written to a http.Request
*/
type WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams struct {

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

// WithTimeout adds the timeout to the workflow job nodes workflow job nodes always nodes list params
func (o *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams) WithTimeout(timeout time.Duration) *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job nodes workflow job nodes always nodes list params
func (o *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job nodes workflow job nodes always nodes list params
func (o *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams) WithContext(ctx context.Context) *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job nodes workflow job nodes always nodes list params
func (o *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job nodes workflow job nodes always nodes list params
func (o *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams) WithHTTPClient(client *http.Client) *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job nodes workflow job nodes always nodes list params
func (o *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the workflow job nodes workflow job nodes always nodes list params
func (o *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams) WithID(id string) *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job nodes workflow job nodes always nodes list params
func (o *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the workflow job nodes workflow job nodes always nodes list params
func (o *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams) WithPage(page *int64) *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the workflow job nodes workflow job nodes always nodes list params
func (o *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the workflow job nodes workflow job nodes always nodes list params
func (o *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams) WithPageSize(pageSize *int64) *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the workflow job nodes workflow job nodes always nodes list params
func (o *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the workflow job nodes workflow job nodes always nodes list params
func (o *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams) WithSearch(search *string) *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the workflow job nodes workflow job nodes always nodes list params
func (o *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobNodesWorkflowJobNodesAlwaysNodesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
