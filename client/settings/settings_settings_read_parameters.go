// Code generated by go-swagger; DO NOT EDIT.

package settings

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

// NewSettingsSettingsReadParams creates a new SettingsSettingsReadParams object
// with the default values initialized.
func NewSettingsSettingsReadParams() *SettingsSettingsReadParams {
	var ()
	return &SettingsSettingsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSettingsSettingsReadParamsWithTimeout creates a new SettingsSettingsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSettingsSettingsReadParamsWithTimeout(timeout time.Duration) *SettingsSettingsReadParams {
	var ()
	return &SettingsSettingsReadParams{

		timeout: timeout,
	}
}

// NewSettingsSettingsReadParamsWithContext creates a new SettingsSettingsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewSettingsSettingsReadParamsWithContext(ctx context.Context) *SettingsSettingsReadParams {
	var ()
	return &SettingsSettingsReadParams{

		Context: ctx,
	}
}

// NewSettingsSettingsReadParamsWithHTTPClient creates a new SettingsSettingsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSettingsSettingsReadParamsWithHTTPClient(client *http.Client) *SettingsSettingsReadParams {
	var ()
	return &SettingsSettingsReadParams{
		HTTPClient: client,
	}
}

/*SettingsSettingsReadParams contains all the parameters to send to the API endpoint
for the settings settings read operation typically these are written to a http.Request
*/
type SettingsSettingsReadParams struct {

	/*CategorySlug*/
	CategorySlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the settings settings read params
func (o *SettingsSettingsReadParams) WithTimeout(timeout time.Duration) *SettingsSettingsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the settings settings read params
func (o *SettingsSettingsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the settings settings read params
func (o *SettingsSettingsReadParams) WithContext(ctx context.Context) *SettingsSettingsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the settings settings read params
func (o *SettingsSettingsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the settings settings read params
func (o *SettingsSettingsReadParams) WithHTTPClient(client *http.Client) *SettingsSettingsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the settings settings read params
func (o *SettingsSettingsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategorySlug adds the categorySlug to the settings settings read params
func (o *SettingsSettingsReadParams) WithCategorySlug(categorySlug string) *SettingsSettingsReadParams {
	o.SetCategorySlug(categorySlug)
	return o
}

// SetCategorySlug adds the categorySlug to the settings settings read params
func (o *SettingsSettingsReadParams) SetCategorySlug(categorySlug string) {
	o.CategorySlug = categorySlug
}

// WriteToRequest writes these params to a swagger request
func (o *SettingsSettingsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param category_slug
	if err := r.SetPathParam("category_slug", o.CategorySlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
