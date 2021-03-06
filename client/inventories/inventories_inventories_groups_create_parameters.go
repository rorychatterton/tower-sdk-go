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

// NewInventoriesInventoriesGroupsCreateParams creates a new InventoriesInventoriesGroupsCreateParams object
// with the default values initialized.
func NewInventoriesInventoriesGroupsCreateParams() *InventoriesInventoriesGroupsCreateParams {
	var ()
	return &InventoriesInventoriesGroupsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInventoriesInventoriesGroupsCreateParamsWithTimeout creates a new InventoriesInventoriesGroupsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInventoriesInventoriesGroupsCreateParamsWithTimeout(timeout time.Duration) *InventoriesInventoriesGroupsCreateParams {
	var ()
	return &InventoriesInventoriesGroupsCreateParams{

		timeout: timeout,
	}
}

// NewInventoriesInventoriesGroupsCreateParamsWithContext creates a new InventoriesInventoriesGroupsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewInventoriesInventoriesGroupsCreateParamsWithContext(ctx context.Context) *InventoriesInventoriesGroupsCreateParams {
	var ()
	return &InventoriesInventoriesGroupsCreateParams{

		Context: ctx,
	}
}

// NewInventoriesInventoriesGroupsCreateParamsWithHTTPClient creates a new InventoriesInventoriesGroupsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInventoriesInventoriesGroupsCreateParamsWithHTTPClient(client *http.Client) *InventoriesInventoriesGroupsCreateParams {
	var ()
	return &InventoriesInventoriesGroupsCreateParams{
		HTTPClient: client,
	}
}

/*InventoriesInventoriesGroupsCreateParams contains all the parameters to send to the API endpoint
for the inventories inventories groups create operation typically these are written to a http.Request
*/
type InventoriesInventoriesGroupsCreateParams struct {

	/*Data*/
	Data interface{}
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the inventories inventories groups create params
func (o *InventoriesInventoriesGroupsCreateParams) WithTimeout(timeout time.Duration) *InventoriesInventoriesGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventories inventories groups create params
func (o *InventoriesInventoriesGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventories inventories groups create params
func (o *InventoriesInventoriesGroupsCreateParams) WithContext(ctx context.Context) *InventoriesInventoriesGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventories inventories groups create params
func (o *InventoriesInventoriesGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventories inventories groups create params
func (o *InventoriesInventoriesGroupsCreateParams) WithHTTPClient(client *http.Client) *InventoriesInventoriesGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventories inventories groups create params
func (o *InventoriesInventoriesGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the inventories inventories groups create params
func (o *InventoriesInventoriesGroupsCreateParams) WithData(data interface{}) *InventoriesInventoriesGroupsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the inventories inventories groups create params
func (o *InventoriesInventoriesGroupsCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the inventories inventories groups create params
func (o *InventoriesInventoriesGroupsCreateParams) WithID(id string) *InventoriesInventoriesGroupsCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventories inventories groups create params
func (o *InventoriesInventoriesGroupsCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InventoriesInventoriesGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
