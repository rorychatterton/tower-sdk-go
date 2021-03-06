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

// NewCustomInventoryScriptsInventoryScriptsReadParams creates a new CustomInventoryScriptsInventoryScriptsReadParams object
// with the default values initialized.
func NewCustomInventoryScriptsInventoryScriptsReadParams() *CustomInventoryScriptsInventoryScriptsReadParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomInventoryScriptsInventoryScriptsReadParamsWithTimeout creates a new CustomInventoryScriptsInventoryScriptsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomInventoryScriptsInventoryScriptsReadParamsWithTimeout(timeout time.Duration) *CustomInventoryScriptsInventoryScriptsReadParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsReadParams{

		timeout: timeout,
	}
}

// NewCustomInventoryScriptsInventoryScriptsReadParamsWithContext creates a new CustomInventoryScriptsInventoryScriptsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomInventoryScriptsInventoryScriptsReadParamsWithContext(ctx context.Context) *CustomInventoryScriptsInventoryScriptsReadParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsReadParams{

		Context: ctx,
	}
}

// NewCustomInventoryScriptsInventoryScriptsReadParamsWithHTTPClient creates a new CustomInventoryScriptsInventoryScriptsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomInventoryScriptsInventoryScriptsReadParamsWithHTTPClient(client *http.Client) *CustomInventoryScriptsInventoryScriptsReadParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsReadParams{
		HTTPClient: client,
	}
}

/*CustomInventoryScriptsInventoryScriptsReadParams contains all the parameters to send to the API endpoint
for the custom inventory scripts inventory scripts read operation typically these are written to a http.Request
*/
type CustomInventoryScriptsInventoryScriptsReadParams struct {

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

// WithTimeout adds the timeout to the custom inventory scripts inventory scripts read params
func (o *CustomInventoryScriptsInventoryScriptsReadParams) WithTimeout(timeout time.Duration) *CustomInventoryScriptsInventoryScriptsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom inventory scripts inventory scripts read params
func (o *CustomInventoryScriptsInventoryScriptsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom inventory scripts inventory scripts read params
func (o *CustomInventoryScriptsInventoryScriptsReadParams) WithContext(ctx context.Context) *CustomInventoryScriptsInventoryScriptsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom inventory scripts inventory scripts read params
func (o *CustomInventoryScriptsInventoryScriptsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom inventory scripts inventory scripts read params
func (o *CustomInventoryScriptsInventoryScriptsReadParams) WithHTTPClient(client *http.Client) *CustomInventoryScriptsInventoryScriptsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom inventory scripts inventory scripts read params
func (o *CustomInventoryScriptsInventoryScriptsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom inventory scripts inventory scripts read params
func (o *CustomInventoryScriptsInventoryScriptsReadParams) WithID(id string) *CustomInventoryScriptsInventoryScriptsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom inventory scripts inventory scripts read params
func (o *CustomInventoryScriptsInventoryScriptsReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the custom inventory scripts inventory scripts read params
func (o *CustomInventoryScriptsInventoryScriptsReadParams) WithSearch(search *string) *CustomInventoryScriptsInventoryScriptsReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the custom inventory scripts inventory scripts read params
func (o *CustomInventoryScriptsInventoryScriptsReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *CustomInventoryScriptsInventoryScriptsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
