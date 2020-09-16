// Code generated by go-swagger; DO NOT EDIT.

package job_events

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

// NewJobEventsJobEventsListParams creates a new JobEventsJobEventsListParams object
// with the default values initialized.
func NewJobEventsJobEventsListParams() *JobEventsJobEventsListParams {
	var ()
	return &JobEventsJobEventsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJobEventsJobEventsListParamsWithTimeout creates a new JobEventsJobEventsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobEventsJobEventsListParamsWithTimeout(timeout time.Duration) *JobEventsJobEventsListParams {
	var ()
	return &JobEventsJobEventsListParams{

		timeout: timeout,
	}
}

// NewJobEventsJobEventsListParamsWithContext creates a new JobEventsJobEventsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobEventsJobEventsListParamsWithContext(ctx context.Context) *JobEventsJobEventsListParams {
	var ()
	return &JobEventsJobEventsListParams{

		Context: ctx,
	}
}

// NewJobEventsJobEventsListParamsWithHTTPClient creates a new JobEventsJobEventsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJobEventsJobEventsListParamsWithHTTPClient(client *http.Client) *JobEventsJobEventsListParams {
	var ()
	return &JobEventsJobEventsListParams{
		HTTPClient: client,
	}
}

/*JobEventsJobEventsListParams contains all the parameters to send to the API endpoint
for the job events job events list operation typically these are written to a http.Request
*/
type JobEventsJobEventsListParams struct {

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

// WithTimeout adds the timeout to the job events job events list params
func (o *JobEventsJobEventsListParams) WithTimeout(timeout time.Duration) *JobEventsJobEventsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job events job events list params
func (o *JobEventsJobEventsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job events job events list params
func (o *JobEventsJobEventsListParams) WithContext(ctx context.Context) *JobEventsJobEventsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job events job events list params
func (o *JobEventsJobEventsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job events job events list params
func (o *JobEventsJobEventsListParams) WithHTTPClient(client *http.Client) *JobEventsJobEventsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job events job events list params
func (o *JobEventsJobEventsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the job events job events list params
func (o *JobEventsJobEventsListParams) WithPage(page *int64) *JobEventsJobEventsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the job events job events list params
func (o *JobEventsJobEventsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the job events job events list params
func (o *JobEventsJobEventsListParams) WithPageSize(pageSize *int64) *JobEventsJobEventsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the job events job events list params
func (o *JobEventsJobEventsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the job events job events list params
func (o *JobEventsJobEventsListParams) WithSearch(search *string) *JobEventsJobEventsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the job events job events list params
func (o *JobEventsJobEventsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *JobEventsJobEventsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
