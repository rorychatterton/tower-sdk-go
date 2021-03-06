// Code generated by go-swagger; DO NOT EDIT.

package custom_inventory_scripts

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

// NewCustomInventoryScriptsInventoryScriptsListParams creates a new CustomInventoryScriptsInventoryScriptsListParams object
// with the default values initialized.
func NewCustomInventoryScriptsInventoryScriptsListParams() *CustomInventoryScriptsInventoryScriptsListParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomInventoryScriptsInventoryScriptsListParamsWithTimeout creates a new CustomInventoryScriptsInventoryScriptsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomInventoryScriptsInventoryScriptsListParamsWithTimeout(timeout time.Duration) *CustomInventoryScriptsInventoryScriptsListParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsListParams{

		timeout: timeout,
	}
}

// NewCustomInventoryScriptsInventoryScriptsListParamsWithContext creates a new CustomInventoryScriptsInventoryScriptsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomInventoryScriptsInventoryScriptsListParamsWithContext(ctx context.Context) *CustomInventoryScriptsInventoryScriptsListParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsListParams{

		Context: ctx,
	}
}

// NewCustomInventoryScriptsInventoryScriptsListParamsWithHTTPClient creates a new CustomInventoryScriptsInventoryScriptsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomInventoryScriptsInventoryScriptsListParamsWithHTTPClient(client *http.Client) *CustomInventoryScriptsInventoryScriptsListParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsListParams{
		HTTPClient: client,
	}
}

/*CustomInventoryScriptsInventoryScriptsListParams contains all the parameters to send to the API endpoint
for the custom inventory scripts inventory scripts list operation typically these are written to a http.Request
*/
type CustomInventoryScriptsInventoryScriptsListParams struct {

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

// WithTimeout adds the timeout to the custom inventory scripts inventory scripts list params
func (o *CustomInventoryScriptsInventoryScriptsListParams) WithTimeout(timeout time.Duration) *CustomInventoryScriptsInventoryScriptsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom inventory scripts inventory scripts list params
func (o *CustomInventoryScriptsInventoryScriptsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom inventory scripts inventory scripts list params
func (o *CustomInventoryScriptsInventoryScriptsListParams) WithContext(ctx context.Context) *CustomInventoryScriptsInventoryScriptsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom inventory scripts inventory scripts list params
func (o *CustomInventoryScriptsInventoryScriptsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom inventory scripts inventory scripts list params
func (o *CustomInventoryScriptsInventoryScriptsListParams) WithHTTPClient(client *http.Client) *CustomInventoryScriptsInventoryScriptsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom inventory scripts inventory scripts list params
func (o *CustomInventoryScriptsInventoryScriptsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the custom inventory scripts inventory scripts list params
func (o *CustomInventoryScriptsInventoryScriptsListParams) WithPage(page *int64) *CustomInventoryScriptsInventoryScriptsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the custom inventory scripts inventory scripts list params
func (o *CustomInventoryScriptsInventoryScriptsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the custom inventory scripts inventory scripts list params
func (o *CustomInventoryScriptsInventoryScriptsListParams) WithPageSize(pageSize *int64) *CustomInventoryScriptsInventoryScriptsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the custom inventory scripts inventory scripts list params
func (o *CustomInventoryScriptsInventoryScriptsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the custom inventory scripts inventory scripts list params
func (o *CustomInventoryScriptsInventoryScriptsListParams) WithSearch(search *string) *CustomInventoryScriptsInventoryScriptsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the custom inventory scripts inventory scripts list params
func (o *CustomInventoryScriptsInventoryScriptsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *CustomInventoryScriptsInventoryScriptsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
