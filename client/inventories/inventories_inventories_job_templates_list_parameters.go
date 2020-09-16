// Code generated by go-swagger; DO NOT EDIT.

package inventories

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

// NewInventoriesInventoriesJobTemplatesListParams creates a new InventoriesInventoriesJobTemplatesListParams object
// with the default values initialized.
func NewInventoriesInventoriesJobTemplatesListParams() *InventoriesInventoriesJobTemplatesListParams {
	var ()
	return &InventoriesInventoriesJobTemplatesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInventoriesInventoriesJobTemplatesListParamsWithTimeout creates a new InventoriesInventoriesJobTemplatesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInventoriesInventoriesJobTemplatesListParamsWithTimeout(timeout time.Duration) *InventoriesInventoriesJobTemplatesListParams {
	var ()
	return &InventoriesInventoriesJobTemplatesListParams{

		timeout: timeout,
	}
}

// NewInventoriesInventoriesJobTemplatesListParamsWithContext creates a new InventoriesInventoriesJobTemplatesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewInventoriesInventoriesJobTemplatesListParamsWithContext(ctx context.Context) *InventoriesInventoriesJobTemplatesListParams {
	var ()
	return &InventoriesInventoriesJobTemplatesListParams{

		Context: ctx,
	}
}

// NewInventoriesInventoriesJobTemplatesListParamsWithHTTPClient creates a new InventoriesInventoriesJobTemplatesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInventoriesInventoriesJobTemplatesListParamsWithHTTPClient(client *http.Client) *InventoriesInventoriesJobTemplatesListParams {
	var ()
	return &InventoriesInventoriesJobTemplatesListParams{
		HTTPClient: client,
	}
}

/*InventoriesInventoriesJobTemplatesListParams contains all the parameters to send to the API endpoint
for the inventories inventories job templates list operation typically these are written to a http.Request
*/
type InventoriesInventoriesJobTemplatesListParams struct {

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

// WithTimeout adds the timeout to the inventories inventories job templates list params
func (o *InventoriesInventoriesJobTemplatesListParams) WithTimeout(timeout time.Duration) *InventoriesInventoriesJobTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventories inventories job templates list params
func (o *InventoriesInventoriesJobTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventories inventories job templates list params
func (o *InventoriesInventoriesJobTemplatesListParams) WithContext(ctx context.Context) *InventoriesInventoriesJobTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventories inventories job templates list params
func (o *InventoriesInventoriesJobTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventories inventories job templates list params
func (o *InventoriesInventoriesJobTemplatesListParams) WithHTTPClient(client *http.Client) *InventoriesInventoriesJobTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventories inventories job templates list params
func (o *InventoriesInventoriesJobTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the inventories inventories job templates list params
func (o *InventoriesInventoriesJobTemplatesListParams) WithID(id string) *InventoriesInventoriesJobTemplatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventories inventories job templates list params
func (o *InventoriesInventoriesJobTemplatesListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the inventories inventories job templates list params
func (o *InventoriesInventoriesJobTemplatesListParams) WithPage(page *int64) *InventoriesInventoriesJobTemplatesListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the inventories inventories job templates list params
func (o *InventoriesInventoriesJobTemplatesListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the inventories inventories job templates list params
func (o *InventoriesInventoriesJobTemplatesListParams) WithPageSize(pageSize *int64) *InventoriesInventoriesJobTemplatesListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the inventories inventories job templates list params
func (o *InventoriesInventoriesJobTemplatesListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the inventories inventories job templates list params
func (o *InventoriesInventoriesJobTemplatesListParams) WithSearch(search *string) *InventoriesInventoriesJobTemplatesListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the inventories inventories job templates list params
func (o *InventoriesInventoriesJobTemplatesListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InventoriesInventoriesJobTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
