// Code generated by go-swagger; DO NOT EDIT.

package instances

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

// NewInstancesInstancesJobsListParams creates a new InstancesInstancesJobsListParams object
// with the default values initialized.
func NewInstancesInstancesJobsListParams() *InstancesInstancesJobsListParams {
	var ()
	return &InstancesInstancesJobsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInstancesInstancesJobsListParamsWithTimeout creates a new InstancesInstancesJobsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInstancesInstancesJobsListParamsWithTimeout(timeout time.Duration) *InstancesInstancesJobsListParams {
	var ()
	return &InstancesInstancesJobsListParams{

		timeout: timeout,
	}
}

// NewInstancesInstancesJobsListParamsWithContext creates a new InstancesInstancesJobsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewInstancesInstancesJobsListParamsWithContext(ctx context.Context) *InstancesInstancesJobsListParams {
	var ()
	return &InstancesInstancesJobsListParams{

		Context: ctx,
	}
}

// NewInstancesInstancesJobsListParamsWithHTTPClient creates a new InstancesInstancesJobsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInstancesInstancesJobsListParamsWithHTTPClient(client *http.Client) *InstancesInstancesJobsListParams {
	var ()
	return &InstancesInstancesJobsListParams{
		HTTPClient: client,
	}
}

/*InstancesInstancesJobsListParams contains all the parameters to send to the API endpoint
for the instances instances jobs list operation typically these are written to a http.Request
*/
type InstancesInstancesJobsListParams struct {

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

// WithTimeout adds the timeout to the instances instances jobs list params
func (o *InstancesInstancesJobsListParams) WithTimeout(timeout time.Duration) *InstancesInstancesJobsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the instances instances jobs list params
func (o *InstancesInstancesJobsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the instances instances jobs list params
func (o *InstancesInstancesJobsListParams) WithContext(ctx context.Context) *InstancesInstancesJobsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the instances instances jobs list params
func (o *InstancesInstancesJobsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the instances instances jobs list params
func (o *InstancesInstancesJobsListParams) WithHTTPClient(client *http.Client) *InstancesInstancesJobsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the instances instances jobs list params
func (o *InstancesInstancesJobsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the instances instances jobs list params
func (o *InstancesInstancesJobsListParams) WithID(id string) *InstancesInstancesJobsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the instances instances jobs list params
func (o *InstancesInstancesJobsListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the instances instances jobs list params
func (o *InstancesInstancesJobsListParams) WithPage(page *int64) *InstancesInstancesJobsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the instances instances jobs list params
func (o *InstancesInstancesJobsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the instances instances jobs list params
func (o *InstancesInstancesJobsListParams) WithPageSize(pageSize *int64) *InstancesInstancesJobsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the instances instances jobs list params
func (o *InstancesInstancesJobsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the instances instances jobs list params
func (o *InstancesInstancesJobsListParams) WithSearch(search *string) *InstancesInstancesJobsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the instances instances jobs list params
func (o *InstancesInstancesJobsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InstancesInstancesJobsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
