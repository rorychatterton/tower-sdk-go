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

// NewInventorySourcesInventorySourcesGroupsDeleteParams creates a new InventorySourcesInventorySourcesGroupsDeleteParams object
// with the default values initialized.
func NewInventorySourcesInventorySourcesGroupsDeleteParams() *InventorySourcesInventorySourcesGroupsDeleteParams {
	var ()
	return &InventorySourcesInventorySourcesGroupsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInventorySourcesInventorySourcesGroupsDeleteParamsWithTimeout creates a new InventorySourcesInventorySourcesGroupsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInventorySourcesInventorySourcesGroupsDeleteParamsWithTimeout(timeout time.Duration) *InventorySourcesInventorySourcesGroupsDeleteParams {
	var ()
	return &InventorySourcesInventorySourcesGroupsDeleteParams{

		timeout: timeout,
	}
}

// NewInventorySourcesInventorySourcesGroupsDeleteParamsWithContext creates a new InventorySourcesInventorySourcesGroupsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewInventorySourcesInventorySourcesGroupsDeleteParamsWithContext(ctx context.Context) *InventorySourcesInventorySourcesGroupsDeleteParams {
	var ()
	return &InventorySourcesInventorySourcesGroupsDeleteParams{

		Context: ctx,
	}
}

// NewInventorySourcesInventorySourcesGroupsDeleteParamsWithHTTPClient creates a new InventorySourcesInventorySourcesGroupsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInventorySourcesInventorySourcesGroupsDeleteParamsWithHTTPClient(client *http.Client) *InventorySourcesInventorySourcesGroupsDeleteParams {
	var ()
	return &InventorySourcesInventorySourcesGroupsDeleteParams{
		HTTPClient: client,
	}
}

/*InventorySourcesInventorySourcesGroupsDeleteParams contains all the parameters to send to the API endpoint
for the inventory sources inventory sources groups delete operation typically these are written to a http.Request
*/
type InventorySourcesInventorySourcesGroupsDeleteParams struct {

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

// WithTimeout adds the timeout to the inventory sources inventory sources groups delete params
func (o *InventorySourcesInventorySourcesGroupsDeleteParams) WithTimeout(timeout time.Duration) *InventorySourcesInventorySourcesGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory sources inventory sources groups delete params
func (o *InventorySourcesInventorySourcesGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory sources inventory sources groups delete params
func (o *InventorySourcesInventorySourcesGroupsDeleteParams) WithContext(ctx context.Context) *InventorySourcesInventorySourcesGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory sources inventory sources groups delete params
func (o *InventorySourcesInventorySourcesGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory sources inventory sources groups delete params
func (o *InventorySourcesInventorySourcesGroupsDeleteParams) WithHTTPClient(client *http.Client) *InventorySourcesInventorySourcesGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory sources inventory sources groups delete params
func (o *InventorySourcesInventorySourcesGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the inventory sources inventory sources groups delete params
func (o *InventorySourcesInventorySourcesGroupsDeleteParams) WithID(id string) *InventorySourcesInventorySourcesGroupsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventory sources inventory sources groups delete params
func (o *InventorySourcesInventorySourcesGroupsDeleteParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the inventory sources inventory sources groups delete params
func (o *InventorySourcesInventorySourcesGroupsDeleteParams) WithSearch(search *string) *InventorySourcesInventorySourcesGroupsDeleteParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the inventory sources inventory sources groups delete params
func (o *InventorySourcesInventorySourcesGroupsDeleteParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InventorySourcesInventorySourcesGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
