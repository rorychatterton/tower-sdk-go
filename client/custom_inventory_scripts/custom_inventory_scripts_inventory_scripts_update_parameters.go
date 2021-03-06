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
)

// NewCustomInventoryScriptsInventoryScriptsUpdateParams creates a new CustomInventoryScriptsInventoryScriptsUpdateParams object
// with the default values initialized.
func NewCustomInventoryScriptsInventoryScriptsUpdateParams() *CustomInventoryScriptsInventoryScriptsUpdateParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomInventoryScriptsInventoryScriptsUpdateParamsWithTimeout creates a new CustomInventoryScriptsInventoryScriptsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomInventoryScriptsInventoryScriptsUpdateParamsWithTimeout(timeout time.Duration) *CustomInventoryScriptsInventoryScriptsUpdateParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsUpdateParams{

		timeout: timeout,
	}
}

// NewCustomInventoryScriptsInventoryScriptsUpdateParamsWithContext creates a new CustomInventoryScriptsInventoryScriptsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomInventoryScriptsInventoryScriptsUpdateParamsWithContext(ctx context.Context) *CustomInventoryScriptsInventoryScriptsUpdateParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsUpdateParams{

		Context: ctx,
	}
}

// NewCustomInventoryScriptsInventoryScriptsUpdateParamsWithHTTPClient creates a new CustomInventoryScriptsInventoryScriptsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomInventoryScriptsInventoryScriptsUpdateParamsWithHTTPClient(client *http.Client) *CustomInventoryScriptsInventoryScriptsUpdateParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsUpdateParams{
		HTTPClient: client,
	}
}

/*CustomInventoryScriptsInventoryScriptsUpdateParams contains all the parameters to send to the API endpoint
for the custom inventory scripts inventory scripts update operation typically these are written to a http.Request
*/
type CustomInventoryScriptsInventoryScriptsUpdateParams struct {

	/*Data*/
	Data CustomInventoryScriptsInventoryScriptsUpdateBody
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

// WithTimeout adds the timeout to the custom inventory scripts inventory scripts update params
func (o *CustomInventoryScriptsInventoryScriptsUpdateParams) WithTimeout(timeout time.Duration) *CustomInventoryScriptsInventoryScriptsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom inventory scripts inventory scripts update params
func (o *CustomInventoryScriptsInventoryScriptsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom inventory scripts inventory scripts update params
func (o *CustomInventoryScriptsInventoryScriptsUpdateParams) WithContext(ctx context.Context) *CustomInventoryScriptsInventoryScriptsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom inventory scripts inventory scripts update params
func (o *CustomInventoryScriptsInventoryScriptsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom inventory scripts inventory scripts update params
func (o *CustomInventoryScriptsInventoryScriptsUpdateParams) WithHTTPClient(client *http.Client) *CustomInventoryScriptsInventoryScriptsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom inventory scripts inventory scripts update params
func (o *CustomInventoryScriptsInventoryScriptsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the custom inventory scripts inventory scripts update params
func (o *CustomInventoryScriptsInventoryScriptsUpdateParams) WithData(data CustomInventoryScriptsInventoryScriptsUpdateBody) *CustomInventoryScriptsInventoryScriptsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the custom inventory scripts inventory scripts update params
func (o *CustomInventoryScriptsInventoryScriptsUpdateParams) SetData(data CustomInventoryScriptsInventoryScriptsUpdateBody) {
	o.Data = data
}

// WithID adds the id to the custom inventory scripts inventory scripts update params
func (o *CustomInventoryScriptsInventoryScriptsUpdateParams) WithID(id string) *CustomInventoryScriptsInventoryScriptsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom inventory scripts inventory scripts update params
func (o *CustomInventoryScriptsInventoryScriptsUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the custom inventory scripts inventory scripts update params
func (o *CustomInventoryScriptsInventoryScriptsUpdateParams) WithSearch(search *string) *CustomInventoryScriptsInventoryScriptsUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the custom inventory scripts inventory scripts update params
func (o *CustomInventoryScriptsInventoryScriptsUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *CustomInventoryScriptsInventoryScriptsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
		return err
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
