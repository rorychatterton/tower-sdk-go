// Code generated by go-swagger; DO NOT EDIT.

package labels

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

// NewLabelsLabelsListParams creates a new LabelsLabelsListParams object
// with the default values initialized.
func NewLabelsLabelsListParams() *LabelsLabelsListParams {
	var ()
	return &LabelsLabelsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLabelsLabelsListParamsWithTimeout creates a new LabelsLabelsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLabelsLabelsListParamsWithTimeout(timeout time.Duration) *LabelsLabelsListParams {
	var ()
	return &LabelsLabelsListParams{

		timeout: timeout,
	}
}

// NewLabelsLabelsListParamsWithContext creates a new LabelsLabelsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewLabelsLabelsListParamsWithContext(ctx context.Context) *LabelsLabelsListParams {
	var ()
	return &LabelsLabelsListParams{

		Context: ctx,
	}
}

// NewLabelsLabelsListParamsWithHTTPClient creates a new LabelsLabelsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLabelsLabelsListParamsWithHTTPClient(client *http.Client) *LabelsLabelsListParams {
	var ()
	return &LabelsLabelsListParams{
		HTTPClient: client,
	}
}

/*LabelsLabelsListParams contains all the parameters to send to the API endpoint
for the labels labels list operation typically these are written to a http.Request
*/
type LabelsLabelsListParams struct {

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

// WithTimeout adds the timeout to the labels labels list params
func (o *LabelsLabelsListParams) WithTimeout(timeout time.Duration) *LabelsLabelsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the labels labels list params
func (o *LabelsLabelsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the labels labels list params
func (o *LabelsLabelsListParams) WithContext(ctx context.Context) *LabelsLabelsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the labels labels list params
func (o *LabelsLabelsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the labels labels list params
func (o *LabelsLabelsListParams) WithHTTPClient(client *http.Client) *LabelsLabelsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the labels labels list params
func (o *LabelsLabelsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the labels labels list params
func (o *LabelsLabelsListParams) WithPage(page *int64) *LabelsLabelsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the labels labels list params
func (o *LabelsLabelsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the labels labels list params
func (o *LabelsLabelsListParams) WithPageSize(pageSize *int64) *LabelsLabelsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the labels labels list params
func (o *LabelsLabelsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the labels labels list params
func (o *LabelsLabelsListParams) WithSearch(search *string) *LabelsLabelsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the labels labels list params
func (o *LabelsLabelsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *LabelsLabelsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
