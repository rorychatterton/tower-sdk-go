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
)

// NewInventorySourcesInventorySourcesUpdateReadParams creates a new InventorySourcesInventorySourcesUpdateReadParams object
// with the default values initialized.
func NewInventorySourcesInventorySourcesUpdateReadParams() *InventorySourcesInventorySourcesUpdateReadParams {
	var ()
	return &InventorySourcesInventorySourcesUpdateReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInventorySourcesInventorySourcesUpdateReadParamsWithTimeout creates a new InventorySourcesInventorySourcesUpdateReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInventorySourcesInventorySourcesUpdateReadParamsWithTimeout(timeout time.Duration) *InventorySourcesInventorySourcesUpdateReadParams {
	var ()
	return &InventorySourcesInventorySourcesUpdateReadParams{

		timeout: timeout,
	}
}

// NewInventorySourcesInventorySourcesUpdateReadParamsWithContext creates a new InventorySourcesInventorySourcesUpdateReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewInventorySourcesInventorySourcesUpdateReadParamsWithContext(ctx context.Context) *InventorySourcesInventorySourcesUpdateReadParams {
	var ()
	return &InventorySourcesInventorySourcesUpdateReadParams{

		Context: ctx,
	}
}

// NewInventorySourcesInventorySourcesUpdateReadParamsWithHTTPClient creates a new InventorySourcesInventorySourcesUpdateReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInventorySourcesInventorySourcesUpdateReadParamsWithHTTPClient(client *http.Client) *InventorySourcesInventorySourcesUpdateReadParams {
	var ()
	return &InventorySourcesInventorySourcesUpdateReadParams{
		HTTPClient: client,
	}
}

/*InventorySourcesInventorySourcesUpdateReadParams contains all the parameters to send to the API endpoint
for the inventory sources inventory sources update read operation typically these are written to a http.Request
*/
type InventorySourcesInventorySourcesUpdateReadParams struct {

	/*ID*/
	ID string
	/*Search
	  A search term.

	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the inventory sources inventory sources update read params
func (o *InventorySourcesInventorySourcesUpdateReadParams) WithTimeout(timeout time.Duration) *InventorySourcesInventorySourcesUpdateReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory sources inventory sources update read params
func (o *InventorySourcesInventorySourcesUpdateReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory sources inventory sources update read params
func (o *InventorySourcesInventorySourcesUpdateReadParams) WithContext(ctx context.Context) *InventorySourcesInventorySourcesUpdateReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory sources inventory sources update read params
func (o *InventorySourcesInventorySourcesUpdateReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory sources inventory sources update read params
func (o *InventorySourcesInventorySourcesUpdateReadParams) WithHTTPClient(client *http.Client) *InventorySourcesInventorySourcesUpdateReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory sources inventory sources update read params
func (o *InventorySourcesInventorySourcesUpdateReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the inventory sources inventory sources update read params
func (o *InventorySourcesInventorySourcesUpdateReadParams) WithID(id string) *InventorySourcesInventorySourcesUpdateReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventory sources inventory sources update read params
func (o *InventorySourcesInventorySourcesUpdateReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the inventory sources inventory sources update read params
func (o *InventorySourcesInventorySourcesUpdateReadParams) WithSearch(search *string) *InventorySourcesInventorySourcesUpdateReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the inventory sources inventory sources update read params
func (o *InventorySourcesInventorySourcesUpdateReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InventorySourcesInventorySourcesUpdateReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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