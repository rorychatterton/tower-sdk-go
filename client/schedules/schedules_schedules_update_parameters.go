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

// NewSchedulesSchedulesUpdateParams creates a new SchedulesSchedulesUpdateParams object
// with the default values initialized.
func NewSchedulesSchedulesUpdateParams() *SchedulesSchedulesUpdateParams {
	var ()
	return &SchedulesSchedulesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSchedulesSchedulesUpdateParamsWithTimeout creates a new SchedulesSchedulesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSchedulesSchedulesUpdateParamsWithTimeout(timeout time.Duration) *SchedulesSchedulesUpdateParams {
	var ()
	return &SchedulesSchedulesUpdateParams{

		timeout: timeout,
	}
}

// NewSchedulesSchedulesUpdateParamsWithContext creates a new SchedulesSchedulesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSchedulesSchedulesUpdateParamsWithContext(ctx context.Context) *SchedulesSchedulesUpdateParams {
	var ()
	return &SchedulesSchedulesUpdateParams{

		Context: ctx,
	}
}

// NewSchedulesSchedulesUpdateParamsWithHTTPClient creates a new SchedulesSchedulesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSchedulesSchedulesUpdateParamsWithHTTPClient(client *http.Client) *SchedulesSchedulesUpdateParams {
	var ()
	return &SchedulesSchedulesUpdateParams{
		HTTPClient: client,
	}
}

/*SchedulesSchedulesUpdateParams contains all the parameters to send to the API endpoint
for the schedules schedules update operation typically these are written to a http.Request
*/
type SchedulesSchedulesUpdateParams struct {

	/*Data*/
	Data SchedulesSchedulesUpdateBody
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

// WithTimeout adds the timeout to the schedules schedules update params
func (o *SchedulesSchedulesUpdateParams) WithTimeout(timeout time.Duration) *SchedulesSchedulesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schedules schedules update params
func (o *SchedulesSchedulesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schedules schedules update params
func (o *SchedulesSchedulesUpdateParams) WithContext(ctx context.Context) *SchedulesSchedulesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schedules schedules update params
func (o *SchedulesSchedulesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schedules schedules update params
func (o *SchedulesSchedulesUpdateParams) WithHTTPClient(client *http.Client) *SchedulesSchedulesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schedules schedules update params
func (o *SchedulesSchedulesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the schedules schedules update params
func (o *SchedulesSchedulesUpdateParams) WithData(data SchedulesSchedulesUpdateBody) *SchedulesSchedulesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the schedules schedules update params
func (o *SchedulesSchedulesUpdateParams) SetData(data SchedulesSchedulesUpdateBody) {
	o.Data = data
}

// WithID adds the id to the schedules schedules update params
func (o *SchedulesSchedulesUpdateParams) WithID(id string) *SchedulesSchedulesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the schedules schedules update params
func (o *SchedulesSchedulesUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the schedules schedules update params
func (o *SchedulesSchedulesUpdateParams) WithSearch(search *string) *SchedulesSchedulesUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the schedules schedules update params
func (o *SchedulesSchedulesUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *SchedulesSchedulesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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