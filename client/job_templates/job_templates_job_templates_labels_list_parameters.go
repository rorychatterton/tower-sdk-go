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

// NewJobTemplatesJobTemplatesLabelsListParams creates a new JobTemplatesJobTemplatesLabelsListParams object
// with the default values initialized.
func NewJobTemplatesJobTemplatesLabelsListParams() *JobTemplatesJobTemplatesLabelsListParams {
	var ()
	return &JobTemplatesJobTemplatesLabelsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJobTemplatesJobTemplatesLabelsListParamsWithTimeout creates a new JobTemplatesJobTemplatesLabelsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobTemplatesJobTemplatesLabelsListParamsWithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesLabelsListParams {
	var ()
	return &JobTemplatesJobTemplatesLabelsListParams{

		timeout: timeout,
	}
}

// NewJobTemplatesJobTemplatesLabelsListParamsWithContext creates a new JobTemplatesJobTemplatesLabelsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobTemplatesJobTemplatesLabelsListParamsWithContext(ctx context.Context) *JobTemplatesJobTemplatesLabelsListParams {
	var ()
	return &JobTemplatesJobTemplatesLabelsListParams{

		Context: ctx,
	}
}

// NewJobTemplatesJobTemplatesLabelsListParamsWithHTTPClient creates a new JobTemplatesJobTemplatesLabelsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJobTemplatesJobTemplatesLabelsListParamsWithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesLabelsListParams {
	var ()
	return &JobTemplatesJobTemplatesLabelsListParams{
		HTTPClient: client,
	}
}

/*JobTemplatesJobTemplatesLabelsListParams contains all the parameters to send to the API endpoint
for the job templates job templates labels list operation typically these are written to a http.Request
*/
type JobTemplatesJobTemplatesLabelsListParams struct {

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

// WithTimeout adds the timeout to the job templates job templates labels list params
func (o *JobTemplatesJobTemplatesLabelsListParams) WithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesLabelsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job templates job templates labels list params
func (o *JobTemplatesJobTemplatesLabelsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job templates job templates labels list params
func (o *JobTemplatesJobTemplatesLabelsListParams) WithContext(ctx context.Context) *JobTemplatesJobTemplatesLabelsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job templates job templates labels list params
func (o *JobTemplatesJobTemplatesLabelsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job templates job templates labels list params
func (o *JobTemplatesJobTemplatesLabelsListParams) WithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesLabelsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job templates job templates labels list params
func (o *JobTemplatesJobTemplatesLabelsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the job templates job templates labels list params
func (o *JobTemplatesJobTemplatesLabelsListParams) WithID(id string) *JobTemplatesJobTemplatesLabelsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the job templates job templates labels list params
func (o *JobTemplatesJobTemplatesLabelsListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the job templates job templates labels list params
func (o *JobTemplatesJobTemplatesLabelsListParams) WithPage(page *int64) *JobTemplatesJobTemplatesLabelsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the job templates job templates labels list params
func (o *JobTemplatesJobTemplatesLabelsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the job templates job templates labels list params
func (o *JobTemplatesJobTemplatesLabelsListParams) WithPageSize(pageSize *int64) *JobTemplatesJobTemplatesLabelsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the job templates job templates labels list params
func (o *JobTemplatesJobTemplatesLabelsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the job templates job templates labels list params
func (o *JobTemplatesJobTemplatesLabelsListParams) WithSearch(search *string) *JobTemplatesJobTemplatesLabelsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the job templates job templates labels list params
func (o *JobTemplatesJobTemplatesLabelsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *JobTemplatesJobTemplatesLabelsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
