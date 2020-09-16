// Code generated by go-swagger; DO NOT EDIT.

package metrics

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

// NewMetricsMetricsListParams creates a new MetricsMetricsListParams object
// with the default values initialized.
func NewMetricsMetricsListParams() *MetricsMetricsListParams {

	return &MetricsMetricsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMetricsMetricsListParamsWithTimeout creates a new MetricsMetricsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMetricsMetricsListParamsWithTimeout(timeout time.Duration) *MetricsMetricsListParams {

	return &MetricsMetricsListParams{

		timeout: timeout,
	}
}

// NewMetricsMetricsListParamsWithContext creates a new MetricsMetricsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewMetricsMetricsListParamsWithContext(ctx context.Context) *MetricsMetricsListParams {

	return &MetricsMetricsListParams{

		Context: ctx,
	}
}

// NewMetricsMetricsListParamsWithHTTPClient creates a new MetricsMetricsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMetricsMetricsListParamsWithHTTPClient(client *http.Client) *MetricsMetricsListParams {

	return &MetricsMetricsListParams{
		HTTPClient: client,
	}
}

/*MetricsMetricsListParams contains all the parameters to send to the API endpoint
for the metrics metrics list operation typically these are written to a http.Request
*/
type MetricsMetricsListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the metrics metrics list params
func (o *MetricsMetricsListParams) WithTimeout(timeout time.Duration) *MetricsMetricsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the metrics metrics list params
func (o *MetricsMetricsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the metrics metrics list params
func (o *MetricsMetricsListParams) WithContext(ctx context.Context) *MetricsMetricsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the metrics metrics list params
func (o *MetricsMetricsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the metrics metrics list params
func (o *MetricsMetricsListParams) WithHTTPClient(client *http.Client) *MetricsMetricsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the metrics metrics list params
func (o *MetricsMetricsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MetricsMetricsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
