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
)

// NewInventoriesInventoriesTreeReadParams creates a new InventoriesInventoriesTreeReadParams object
// with the default values initialized.
func NewInventoriesInventoriesTreeReadParams() *InventoriesInventoriesTreeReadParams {
	var ()
	return &InventoriesInventoriesTreeReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInventoriesInventoriesTreeReadParamsWithTimeout creates a new InventoriesInventoriesTreeReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInventoriesInventoriesTreeReadParamsWithTimeout(timeout time.Duration) *InventoriesInventoriesTreeReadParams {
	var ()
	return &InventoriesInventoriesTreeReadParams{

		timeout: timeout,
	}
}

// NewInventoriesInventoriesTreeReadParamsWithContext creates a new InventoriesInventoriesTreeReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewInventoriesInventoriesTreeReadParamsWithContext(ctx context.Context) *InventoriesInventoriesTreeReadParams {
	var ()
	return &InventoriesInventoriesTreeReadParams{

		Context: ctx,
	}
}

// NewInventoriesInventoriesTreeReadParamsWithHTTPClient creates a new InventoriesInventoriesTreeReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInventoriesInventoriesTreeReadParamsWithHTTPClient(client *http.Client) *InventoriesInventoriesTreeReadParams {
	var ()
	return &InventoriesInventoriesTreeReadParams{
		HTTPClient: client,
	}
}

/*InventoriesInventoriesTreeReadParams contains all the parameters to send to the API endpoint
for the inventories inventories tree read operation typically these are written to a http.Request
*/
type InventoriesInventoriesTreeReadParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the inventories inventories tree read params
func (o *InventoriesInventoriesTreeReadParams) WithTimeout(timeout time.Duration) *InventoriesInventoriesTreeReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventories inventories tree read params
func (o *InventoriesInventoriesTreeReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventories inventories tree read params
func (o *InventoriesInventoriesTreeReadParams) WithContext(ctx context.Context) *InventoriesInventoriesTreeReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventories inventories tree read params
func (o *InventoriesInventoriesTreeReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventories inventories tree read params
func (o *InventoriesInventoriesTreeReadParams) WithHTTPClient(client *http.Client) *InventoriesInventoriesTreeReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventories inventories tree read params
func (o *InventoriesInventoriesTreeReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the inventories inventories tree read params
func (o *InventoriesInventoriesTreeReadParams) WithID(id string) *InventoriesInventoriesTreeReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventories inventories tree read params
func (o *InventoriesInventoriesTreeReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InventoriesInventoriesTreeReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
