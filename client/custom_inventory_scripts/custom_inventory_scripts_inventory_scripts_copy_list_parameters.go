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

// NewCustomInventoryScriptsInventoryScriptsCopyListParams creates a new CustomInventoryScriptsInventoryScriptsCopyListParams object
// with the default values initialized.
func NewCustomInventoryScriptsInventoryScriptsCopyListParams() *CustomInventoryScriptsInventoryScriptsCopyListParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsCopyListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomInventoryScriptsInventoryScriptsCopyListParamsWithTimeout creates a new CustomInventoryScriptsInventoryScriptsCopyListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomInventoryScriptsInventoryScriptsCopyListParamsWithTimeout(timeout time.Duration) *CustomInventoryScriptsInventoryScriptsCopyListParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsCopyListParams{

		timeout: timeout,
	}
}

// NewCustomInventoryScriptsInventoryScriptsCopyListParamsWithContext creates a new CustomInventoryScriptsInventoryScriptsCopyListParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomInventoryScriptsInventoryScriptsCopyListParamsWithContext(ctx context.Context) *CustomInventoryScriptsInventoryScriptsCopyListParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsCopyListParams{

		Context: ctx,
	}
}

// NewCustomInventoryScriptsInventoryScriptsCopyListParamsWithHTTPClient creates a new CustomInventoryScriptsInventoryScriptsCopyListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomInventoryScriptsInventoryScriptsCopyListParamsWithHTTPClient(client *http.Client) *CustomInventoryScriptsInventoryScriptsCopyListParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsCopyListParams{
		HTTPClient: client,
	}
}

/*CustomInventoryScriptsInventoryScriptsCopyListParams contains all the parameters to send to the API endpoint
for the custom inventory scripts inventory scripts copy list operation typically these are written to a http.Request
*/
type CustomInventoryScriptsInventoryScriptsCopyListParams struct {

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

// WithTimeout adds the timeout to the custom inventory scripts inventory scripts copy list params
func (o *CustomInventoryScriptsInventoryScriptsCopyListParams) WithTimeout(timeout time.Duration) *CustomInventoryScriptsInventoryScriptsCopyListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom inventory scripts inventory scripts copy list params
func (o *CustomInventoryScriptsInventoryScriptsCopyListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom inventory scripts inventory scripts copy list params
func (o *CustomInventoryScriptsInventoryScriptsCopyListParams) WithContext(ctx context.Context) *CustomInventoryScriptsInventoryScriptsCopyListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom inventory scripts inventory scripts copy list params
func (o *CustomInventoryScriptsInventoryScriptsCopyListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom inventory scripts inventory scripts copy list params
func (o *CustomInventoryScriptsInventoryScriptsCopyListParams) WithHTTPClient(client *http.Client) *CustomInventoryScriptsInventoryScriptsCopyListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom inventory scripts inventory scripts copy list params
func (o *CustomInventoryScriptsInventoryScriptsCopyListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom inventory scripts inventory scripts copy list params
func (o *CustomInventoryScriptsInventoryScriptsCopyListParams) WithID(id string) *CustomInventoryScriptsInventoryScriptsCopyListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom inventory scripts inventory scripts copy list params
func (o *CustomInventoryScriptsInventoryScriptsCopyListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the custom inventory scripts inventory scripts copy list params
func (o *CustomInventoryScriptsInventoryScriptsCopyListParams) WithPage(page *int64) *CustomInventoryScriptsInventoryScriptsCopyListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the custom inventory scripts inventory scripts copy list params
func (o *CustomInventoryScriptsInventoryScriptsCopyListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the custom inventory scripts inventory scripts copy list params
func (o *CustomInventoryScriptsInventoryScriptsCopyListParams) WithPageSize(pageSize *int64) *CustomInventoryScriptsInventoryScriptsCopyListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the custom inventory scripts inventory scripts copy list params
func (o *CustomInventoryScriptsInventoryScriptsCopyListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the custom inventory scripts inventory scripts copy list params
func (o *CustomInventoryScriptsInventoryScriptsCopyListParams) WithSearch(search *string) *CustomInventoryScriptsInventoryScriptsCopyListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the custom inventory scripts inventory scripts copy list params
func (o *CustomInventoryScriptsInventoryScriptsCopyListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *CustomInventoryScriptsInventoryScriptsCopyListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
