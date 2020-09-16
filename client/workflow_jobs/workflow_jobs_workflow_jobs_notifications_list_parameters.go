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

// NewWorkflowJobsWorkflowJobsNotificationsListParams creates a new WorkflowJobsWorkflowJobsNotificationsListParams object
// with the default values initialized.
func NewWorkflowJobsWorkflowJobsNotificationsListParams() *WorkflowJobsWorkflowJobsNotificationsListParams {
	var ()
	return &WorkflowJobsWorkflowJobsNotificationsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobsWorkflowJobsNotificationsListParamsWithTimeout creates a new WorkflowJobsWorkflowJobsNotificationsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWorkflowJobsWorkflowJobsNotificationsListParamsWithTimeout(timeout time.Duration) *WorkflowJobsWorkflowJobsNotificationsListParams {
	var ()
	return &WorkflowJobsWorkflowJobsNotificationsListParams{

		timeout: timeout,
	}
}

// NewWorkflowJobsWorkflowJobsNotificationsListParamsWithContext creates a new WorkflowJobsWorkflowJobsNotificationsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewWorkflowJobsWorkflowJobsNotificationsListParamsWithContext(ctx context.Context) *WorkflowJobsWorkflowJobsNotificationsListParams {
	var ()
	return &WorkflowJobsWorkflowJobsNotificationsListParams{

		Context: ctx,
	}
}

// NewWorkflowJobsWorkflowJobsNotificationsListParamsWithHTTPClient creates a new WorkflowJobsWorkflowJobsNotificationsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWorkflowJobsWorkflowJobsNotificationsListParamsWithHTTPClient(client *http.Client) *WorkflowJobsWorkflowJobsNotificationsListParams {
	var ()
	return &WorkflowJobsWorkflowJobsNotificationsListParams{
		HTTPClient: client,
	}
}

/*WorkflowJobsWorkflowJobsNotificationsListParams contains all the parameters to send to the API endpoint
for the workflow jobs workflow jobs notifications list operation typically these are written to a http.Request
*/
type WorkflowJobsWorkflowJobsNotificationsListParams struct {

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

// WithTimeout adds the timeout to the workflow jobs workflow jobs notifications list params
func (o *WorkflowJobsWorkflowJobsNotificationsListParams) WithTimeout(timeout time.Duration) *WorkflowJobsWorkflowJobsNotificationsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow jobs workflow jobs notifications list params
func (o *WorkflowJobsWorkflowJobsNotificationsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow jobs workflow jobs notifications list params
func (o *WorkflowJobsWorkflowJobsNotificationsListParams) WithContext(ctx context.Context) *WorkflowJobsWorkflowJobsNotificationsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow jobs workflow jobs notifications list params
func (o *WorkflowJobsWorkflowJobsNotificationsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow jobs workflow jobs notifications list params
func (o *WorkflowJobsWorkflowJobsNotificationsListParams) WithHTTPClient(client *http.Client) *WorkflowJobsWorkflowJobsNotificationsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow jobs workflow jobs notifications list params
func (o *WorkflowJobsWorkflowJobsNotificationsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the workflow jobs workflow jobs notifications list params
func (o *WorkflowJobsWorkflowJobsNotificationsListParams) WithID(id string) *WorkflowJobsWorkflowJobsNotificationsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow jobs workflow jobs notifications list params
func (o *WorkflowJobsWorkflowJobsNotificationsListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the workflow jobs workflow jobs notifications list params
func (o *WorkflowJobsWorkflowJobsNotificationsListParams) WithPage(page *int64) *WorkflowJobsWorkflowJobsNotificationsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the workflow jobs workflow jobs notifications list params
func (o *WorkflowJobsWorkflowJobsNotificationsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the workflow jobs workflow jobs notifications list params
func (o *WorkflowJobsWorkflowJobsNotificationsListParams) WithPageSize(pageSize *int64) *WorkflowJobsWorkflowJobsNotificationsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the workflow jobs workflow jobs notifications list params
func (o *WorkflowJobsWorkflowJobsNotificationsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the workflow jobs workflow jobs notifications list params
func (o *WorkflowJobsWorkflowJobsNotificationsListParams) WithSearch(search *string) *WorkflowJobsWorkflowJobsNotificationsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the workflow jobs workflow jobs notifications list params
func (o *WorkflowJobsWorkflowJobsNotificationsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobsWorkflowJobsNotificationsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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