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

// NewInstancesInstancesListParams creates a new InstancesInstancesListParams object
// with the default values initialized.
func NewInstancesInstancesListParams() *InstancesInstancesListParams {
	var ()
	return &InstancesInstancesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInstancesInstancesListParamsWithTimeout creates a new InstancesInstancesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInstancesInstancesListParamsWithTimeout(timeout time.Duration) *InstancesInstancesListParams {
	var ()
	return &InstancesInstancesListParams{

		timeout: timeout,
	}
}

// NewInstancesInstancesListParamsWithContext creates a new InstancesInstancesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewInstancesInstancesListParamsWithContext(ctx context.Context) *InstancesInstancesListParams {
	var ()
	return &InstancesInstancesListParams{

		Context: ctx,
	}
}

// NewInstancesInstancesListParamsWithHTTPClient creates a new InstancesInstancesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInstancesInstancesListParamsWithHTTPClient(client *http.Client) *InstancesInstancesListParams {
	var ()
	return &InstancesInstancesListParams{
		HTTPClient: client,
	}
}

/*InstancesInstancesListParams contains all the parameters to send to the API endpoint
for the instances instances list operation typically these are written to a http.Request
*/
type InstancesInstancesListParams struct {

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

// WithTimeout adds the timeout to the instances instances list params
func (o *InstancesInstancesListParams) WithTimeout(timeout time.Duration) *InstancesInstancesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the instances instances list params
func (o *InstancesInstancesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the instances instances list params
func (o *InstancesInstancesListParams) WithContext(ctx context.Context) *InstancesInstancesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the instances instances list params
func (o *InstancesInstancesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the instances instances list params
func (o *InstancesInstancesListParams) WithHTTPClient(client *http.Client) *InstancesInstancesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the instances instances list params
func (o *InstancesInstancesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the instances instances list params
func (o *InstancesInstancesListParams) WithPage(page *int64) *InstancesInstancesListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the instances instances list params
func (o *InstancesInstancesListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the instances instances list params
func (o *InstancesInstancesListParams) WithPageSize(pageSize *int64) *InstancesInstancesListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the instances instances list params
func (o *InstancesInstancesListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the instances instances list params
func (o *InstancesInstancesListParams) WithSearch(search *string) *InstancesInstancesListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the instances instances list params
func (o *InstancesInstancesListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InstancesInstancesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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