// Code generated by go-swagger; DO NOT EDIT.

package inventory_sources

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

// NewInventorySourcesInventorySourcesNotificationTemplatesStartedListParams creates a new InventorySourcesInventorySourcesNotificationTemplatesStartedListParams object
// with the default values initialized.
func NewInventorySourcesInventorySourcesNotificationTemplatesStartedListParams() *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams {
	var ()
	return &InventorySourcesInventorySourcesNotificationTemplatesStartedListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInventorySourcesInventorySourcesNotificationTemplatesStartedListParamsWithTimeout creates a new InventorySourcesInventorySourcesNotificationTemplatesStartedListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInventorySourcesInventorySourcesNotificationTemplatesStartedListParamsWithTimeout(timeout time.Duration) *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams {
	var ()
	return &InventorySourcesInventorySourcesNotificationTemplatesStartedListParams{

		timeout: timeout,
	}
}

// NewInventorySourcesInventorySourcesNotificationTemplatesStartedListParamsWithContext creates a new InventorySourcesInventorySourcesNotificationTemplatesStartedListParams object
// with the default values initialized, and the ability to set a context for a request
func NewInventorySourcesInventorySourcesNotificationTemplatesStartedListParamsWithContext(ctx context.Context) *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams {
	var ()
	return &InventorySourcesInventorySourcesNotificationTemplatesStartedListParams{

		Context: ctx,
	}
}

// NewInventorySourcesInventorySourcesNotificationTemplatesStartedListParamsWithHTTPClient creates a new InventorySourcesInventorySourcesNotificationTemplatesStartedListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInventorySourcesInventorySourcesNotificationTemplatesStartedListParamsWithHTTPClient(client *http.Client) *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams {
	var ()
	return &InventorySourcesInventorySourcesNotificationTemplatesStartedListParams{
		HTTPClient: client,
	}
}

/*InventorySourcesInventorySourcesNotificationTemplatesStartedListParams contains all the parameters to send to the API endpoint
for the inventory sources inventory sources notification templates started list operation typically these are written to a http.Request
*/
type InventorySourcesInventorySourcesNotificationTemplatesStartedListParams struct {

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

// WithTimeout adds the timeout to the inventory sources inventory sources notification templates started list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams) WithTimeout(timeout time.Duration) *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory sources inventory sources notification templates started list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory sources inventory sources notification templates started list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams) WithContext(ctx context.Context) *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory sources inventory sources notification templates started list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory sources inventory sources notification templates started list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams) WithHTTPClient(client *http.Client) *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory sources inventory sources notification templates started list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the inventory sources inventory sources notification templates started list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams) WithID(id string) *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventory sources inventory sources notification templates started list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the inventory sources inventory sources notification templates started list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams) WithPage(page *int64) *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the inventory sources inventory sources notification templates started list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the inventory sources inventory sources notification templates started list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams) WithPageSize(pageSize *int64) *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the inventory sources inventory sources notification templates started list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the inventory sources inventory sources notification templates started list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams) WithSearch(search *string) *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the inventory sources inventory sources notification templates started list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
