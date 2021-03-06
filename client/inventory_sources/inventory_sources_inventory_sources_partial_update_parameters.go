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

// NewInventorySourcesInventorySourcesPartialUpdateParams creates a new InventorySourcesInventorySourcesPartialUpdateParams object
// with the default values initialized.
func NewInventorySourcesInventorySourcesPartialUpdateParams() *InventorySourcesInventorySourcesPartialUpdateParams {
	var ()
	return &InventorySourcesInventorySourcesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInventorySourcesInventorySourcesPartialUpdateParamsWithTimeout creates a new InventorySourcesInventorySourcesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInventorySourcesInventorySourcesPartialUpdateParamsWithTimeout(timeout time.Duration) *InventorySourcesInventorySourcesPartialUpdateParams {
	var ()
	return &InventorySourcesInventorySourcesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewInventorySourcesInventorySourcesPartialUpdateParamsWithContext creates a new InventorySourcesInventorySourcesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewInventorySourcesInventorySourcesPartialUpdateParamsWithContext(ctx context.Context) *InventorySourcesInventorySourcesPartialUpdateParams {
	var ()
	return &InventorySourcesInventorySourcesPartialUpdateParams{

		Context: ctx,
	}
}

// NewInventorySourcesInventorySourcesPartialUpdateParamsWithHTTPClient creates a new InventorySourcesInventorySourcesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInventorySourcesInventorySourcesPartialUpdateParamsWithHTTPClient(client *http.Client) *InventorySourcesInventorySourcesPartialUpdateParams {
	var ()
	return &InventorySourcesInventorySourcesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*InventorySourcesInventorySourcesPartialUpdateParams contains all the parameters to send to the API endpoint
for the inventory sources inventory sources partial update operation typically these are written to a http.Request
*/
type InventorySourcesInventorySourcesPartialUpdateParams struct {

	/*Data*/
	Data interface{}
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

// WithTimeout adds the timeout to the inventory sources inventory sources partial update params
func (o *InventorySourcesInventorySourcesPartialUpdateParams) WithTimeout(timeout time.Duration) *InventorySourcesInventorySourcesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory sources inventory sources partial update params
func (o *InventorySourcesInventorySourcesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory sources inventory sources partial update params
func (o *InventorySourcesInventorySourcesPartialUpdateParams) WithContext(ctx context.Context) *InventorySourcesInventorySourcesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory sources inventory sources partial update params
func (o *InventorySourcesInventorySourcesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory sources inventory sources partial update params
func (o *InventorySourcesInventorySourcesPartialUpdateParams) WithHTTPClient(client *http.Client) *InventorySourcesInventorySourcesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory sources inventory sources partial update params
func (o *InventorySourcesInventorySourcesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the inventory sources inventory sources partial update params
func (o *InventorySourcesInventorySourcesPartialUpdateParams) WithData(data interface{}) *InventorySourcesInventorySourcesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the inventory sources inventory sources partial update params
func (o *InventorySourcesInventorySourcesPartialUpdateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the inventory sources inventory sources partial update params
func (o *InventorySourcesInventorySourcesPartialUpdateParams) WithID(id string) *InventorySourcesInventorySourcesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventory sources inventory sources partial update params
func (o *InventorySourcesInventorySourcesPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the inventory sources inventory sources partial update params
func (o *InventorySourcesInventorySourcesPartialUpdateParams) WithSearch(search *string) *InventorySourcesInventorySourcesPartialUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the inventory sources inventory sources partial update params
func (o *InventorySourcesInventorySourcesPartialUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InventorySourcesInventorySourcesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

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
