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

// NewCustomInventoryScriptsInventoryScriptsCreateParams creates a new CustomInventoryScriptsInventoryScriptsCreateParams object
// with the default values initialized.
func NewCustomInventoryScriptsInventoryScriptsCreateParams() *CustomInventoryScriptsInventoryScriptsCreateParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomInventoryScriptsInventoryScriptsCreateParamsWithTimeout creates a new CustomInventoryScriptsInventoryScriptsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomInventoryScriptsInventoryScriptsCreateParamsWithTimeout(timeout time.Duration) *CustomInventoryScriptsInventoryScriptsCreateParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsCreateParams{

		timeout: timeout,
	}
}

// NewCustomInventoryScriptsInventoryScriptsCreateParamsWithContext creates a new CustomInventoryScriptsInventoryScriptsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomInventoryScriptsInventoryScriptsCreateParamsWithContext(ctx context.Context) *CustomInventoryScriptsInventoryScriptsCreateParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsCreateParams{

		Context: ctx,
	}
}

// NewCustomInventoryScriptsInventoryScriptsCreateParamsWithHTTPClient creates a new CustomInventoryScriptsInventoryScriptsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomInventoryScriptsInventoryScriptsCreateParamsWithHTTPClient(client *http.Client) *CustomInventoryScriptsInventoryScriptsCreateParams {
	var ()
	return &CustomInventoryScriptsInventoryScriptsCreateParams{
		HTTPClient: client,
	}
}

/*CustomInventoryScriptsInventoryScriptsCreateParams contains all the parameters to send to the API endpoint
for the custom inventory scripts inventory scripts create operation typically these are written to a http.Request
*/
type CustomInventoryScriptsInventoryScriptsCreateParams struct {

	/*Data*/
	Data CustomInventoryScriptsInventoryScriptsCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom inventory scripts inventory scripts create params
func (o *CustomInventoryScriptsInventoryScriptsCreateParams) WithTimeout(timeout time.Duration) *CustomInventoryScriptsInventoryScriptsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom inventory scripts inventory scripts create params
func (o *CustomInventoryScriptsInventoryScriptsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom inventory scripts inventory scripts create params
func (o *CustomInventoryScriptsInventoryScriptsCreateParams) WithContext(ctx context.Context) *CustomInventoryScriptsInventoryScriptsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom inventory scripts inventory scripts create params
func (o *CustomInventoryScriptsInventoryScriptsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom inventory scripts inventory scripts create params
func (o *CustomInventoryScriptsInventoryScriptsCreateParams) WithHTTPClient(client *http.Client) *CustomInventoryScriptsInventoryScriptsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom inventory scripts inventory scripts create params
func (o *CustomInventoryScriptsInventoryScriptsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the custom inventory scripts inventory scripts create params
func (o *CustomInventoryScriptsInventoryScriptsCreateParams) WithData(data CustomInventoryScriptsInventoryScriptsCreateBody) *CustomInventoryScriptsInventoryScriptsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the custom inventory scripts inventory scripts create params
func (o *CustomInventoryScriptsInventoryScriptsCreateParams) SetData(data CustomInventoryScriptsInventoryScriptsCreateBody) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CustomInventoryScriptsInventoryScriptsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
