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
	"github.com/go-openapi/swag"
)

// NewJobTemplatesJobTemplatesSliceWorkflowJobsListParams creates a new JobTemplatesJobTemplatesSliceWorkflowJobsListParams object
// with the default values initialized.
func NewJobTemplatesJobTemplatesSliceWorkflowJobsListParams() *JobTemplatesJobTemplatesSliceWorkflowJobsListParams {
	var ()
	return &JobTemplatesJobTemplatesSliceWorkflowJobsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJobTemplatesJobTemplatesSliceWorkflowJobsListParamsWithTimeout creates a new JobTemplatesJobTemplatesSliceWorkflowJobsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobTemplatesJobTemplatesSliceWorkflowJobsListParamsWithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesSliceWorkflowJobsListParams {
	var ()
	return &JobTemplatesJobTemplatesSliceWorkflowJobsListParams{

		timeout: timeout,
	}
}

// NewJobTemplatesJobTemplatesSliceWorkflowJobsListParamsWithContext creates a new JobTemplatesJobTemplatesSliceWorkflowJobsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobTemplatesJobTemplatesSliceWorkflowJobsListParamsWithContext(ctx context.Context) *JobTemplatesJobTemplatesSliceWorkflowJobsListParams {
	var ()
	return &JobTemplatesJobTemplatesSliceWorkflowJobsListParams{

		Context: ctx,
	}
}

// NewJobTemplatesJobTemplatesSliceWorkflowJobsListParamsWithHTTPClient creates a new JobTemplatesJobTemplatesSliceWorkflowJobsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJobTemplatesJobTemplatesSliceWorkflowJobsListParamsWithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesSliceWorkflowJobsListParams {
	var ()
	return &JobTemplatesJobTemplatesSliceWorkflowJobsListParams{
		HTTPClient: client,
	}
}

/*JobTemplatesJobTemplatesSliceWorkflowJobsListParams contains all the parameters to send to the API endpoint
for the job templates job templates slice workflow jobs list operation typically these are written to a http.Request
*/
type JobTemplatesJobTemplatesSliceWorkflowJobsListParams struct {

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

// WithTimeout adds the timeout to the job templates job templates slice workflow jobs list params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListParams) WithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesSliceWorkflowJobsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job templates job templates slice workflow jobs list params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job templates job templates slice workflow jobs list params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListParams) WithContext(ctx context.Context) *JobTemplatesJobTemplatesSliceWorkflowJobsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job templates job templates slice workflow jobs list params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job templates job templates slice workflow jobs list params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListParams) WithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesSliceWorkflowJobsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job templates job templates slice workflow jobs list params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the job templates job templates slice workflow jobs list params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListParams) WithID(id string) *JobTemplatesJobTemplatesSliceWorkflowJobsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the job templates job templates slice workflow jobs list params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the job templates job templates slice workflow jobs list params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListParams) WithPage(page *int64) *JobTemplatesJobTemplatesSliceWorkflowJobsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the job templates job templates slice workflow jobs list params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the job templates job templates slice workflow jobs list params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListParams) WithPageSize(pageSize *int64) *JobTemplatesJobTemplatesSliceWorkflowJobsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the job templates job templates slice workflow jobs list params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the job templates job templates slice workflow jobs list params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListParams) WithSearch(search *string) *JobTemplatesJobTemplatesSliceWorkflowJobsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the job templates job templates slice workflow jobs list params
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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