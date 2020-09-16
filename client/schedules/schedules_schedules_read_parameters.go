// Code generated by go-swagger; DO NOT EDIT.

package schedules

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

// NewSchedulesSchedulesReadParams creates a new SchedulesSchedulesReadParams object
// with the default values initialized.
func NewSchedulesSchedulesReadParams() *SchedulesSchedulesReadParams {
	var ()
	return &SchedulesSchedulesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSchedulesSchedulesReadParamsWithTimeout creates a new SchedulesSchedulesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSchedulesSchedulesReadParamsWithTimeout(timeout time.Duration) *SchedulesSchedulesReadParams {
	var ()
	return &SchedulesSchedulesReadParams{

		timeout: timeout,
	}
}

// NewSchedulesSchedulesReadParamsWithContext creates a new SchedulesSchedulesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewSchedulesSchedulesReadParamsWithContext(ctx context.Context) *SchedulesSchedulesReadParams {
	var ()
	return &SchedulesSchedulesReadParams{

		Context: ctx,
	}
}

// NewSchedulesSchedulesReadParamsWithHTTPClient creates a new SchedulesSchedulesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSchedulesSchedulesReadParamsWithHTTPClient(client *http.Client) *SchedulesSchedulesReadParams {
	var ()
	return &SchedulesSchedulesReadParams{
		HTTPClient: client,
	}
}

/*SchedulesSchedulesReadParams contains all the parameters to send to the API endpoint
for the schedules schedules read operation typically these are written to a http.Request
*/
type SchedulesSchedulesReadParams struct {

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

// WithTimeout adds the timeout to the schedules schedules read params
func (o *SchedulesSchedulesReadParams) WithTimeout(timeout time.Duration) *SchedulesSchedulesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schedules schedules read params
func (o *SchedulesSchedulesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schedules schedules read params
func (o *SchedulesSchedulesReadParams) WithContext(ctx context.Context) *SchedulesSchedulesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schedules schedules read params
func (o *SchedulesSchedulesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schedules schedules read params
func (o *SchedulesSchedulesReadParams) WithHTTPClient(client *http.Client) *SchedulesSchedulesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schedules schedules read params
func (o *SchedulesSchedulesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the schedules schedules read params
func (o *SchedulesSchedulesReadParams) WithID(id string) *SchedulesSchedulesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the schedules schedules read params
func (o *SchedulesSchedulesReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the schedules schedules read params
func (o *SchedulesSchedulesReadParams) WithSearch(search *string) *SchedulesSchedulesReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the schedules schedules read params
func (o *SchedulesSchedulesReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *SchedulesSchedulesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
