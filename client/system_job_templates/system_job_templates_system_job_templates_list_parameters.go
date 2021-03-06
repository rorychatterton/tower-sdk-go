// Code generated by go-swagger; DO NOT EDIT.

package system_job_templates

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

// NewSystemJobTemplatesSystemJobTemplatesListParams creates a new SystemJobTemplatesSystemJobTemplatesListParams object
// with the default values initialized.
func NewSystemJobTemplatesSystemJobTemplatesListParams() *SystemJobTemplatesSystemJobTemplatesListParams {
	var ()
	return &SystemJobTemplatesSystemJobTemplatesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSystemJobTemplatesSystemJobTemplatesListParamsWithTimeout creates a new SystemJobTemplatesSystemJobTemplatesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSystemJobTemplatesSystemJobTemplatesListParamsWithTimeout(timeout time.Duration) *SystemJobTemplatesSystemJobTemplatesListParams {
	var ()
	return &SystemJobTemplatesSystemJobTemplatesListParams{

		timeout: timeout,
	}
}

// NewSystemJobTemplatesSystemJobTemplatesListParamsWithContext creates a new SystemJobTemplatesSystemJobTemplatesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewSystemJobTemplatesSystemJobTemplatesListParamsWithContext(ctx context.Context) *SystemJobTemplatesSystemJobTemplatesListParams {
	var ()
	return &SystemJobTemplatesSystemJobTemplatesListParams{

		Context: ctx,
	}
}

// NewSystemJobTemplatesSystemJobTemplatesListParamsWithHTTPClient creates a new SystemJobTemplatesSystemJobTemplatesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSystemJobTemplatesSystemJobTemplatesListParamsWithHTTPClient(client *http.Client) *SystemJobTemplatesSystemJobTemplatesListParams {
	var ()
	return &SystemJobTemplatesSystemJobTemplatesListParams{
		HTTPClient: client,
	}
}

/*SystemJobTemplatesSystemJobTemplatesListParams contains all the parameters to send to the API endpoint
for the system job templates system job templates list operation typically these are written to a http.Request
*/
type SystemJobTemplatesSystemJobTemplatesListParams struct {

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

// WithTimeout adds the timeout to the system job templates system job templates list params
func (o *SystemJobTemplatesSystemJobTemplatesListParams) WithTimeout(timeout time.Duration) *SystemJobTemplatesSystemJobTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system job templates system job templates list params
func (o *SystemJobTemplatesSystemJobTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system job templates system job templates list params
func (o *SystemJobTemplatesSystemJobTemplatesListParams) WithContext(ctx context.Context) *SystemJobTemplatesSystemJobTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system job templates system job templates list params
func (o *SystemJobTemplatesSystemJobTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system job templates system job templates list params
func (o *SystemJobTemplatesSystemJobTemplatesListParams) WithHTTPClient(client *http.Client) *SystemJobTemplatesSystemJobTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system job templates system job templates list params
func (o *SystemJobTemplatesSystemJobTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the system job templates system job templates list params
func (o *SystemJobTemplatesSystemJobTemplatesListParams) WithPage(page *int64) *SystemJobTemplatesSystemJobTemplatesListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the system job templates system job templates list params
func (o *SystemJobTemplatesSystemJobTemplatesListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the system job templates system job templates list params
func (o *SystemJobTemplatesSystemJobTemplatesListParams) WithPageSize(pageSize *int64) *SystemJobTemplatesSystemJobTemplatesListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the system job templates system job templates list params
func (o *SystemJobTemplatesSystemJobTemplatesListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the system job templates system job templates list params
func (o *SystemJobTemplatesSystemJobTemplatesListParams) WithSearch(search *string) *SystemJobTemplatesSystemJobTemplatesListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the system job templates system job templates list params
func (o *SystemJobTemplatesSystemJobTemplatesListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *SystemJobTemplatesSystemJobTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
