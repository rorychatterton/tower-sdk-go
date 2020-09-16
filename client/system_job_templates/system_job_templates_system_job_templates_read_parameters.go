// Code generated by go-swagger; DO NOT EDIT.

package system_job_templates

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

// NewSystemJobTemplatesSystemJobTemplatesReadParams creates a new SystemJobTemplatesSystemJobTemplatesReadParams object
// with the default values initialized.
func NewSystemJobTemplatesSystemJobTemplatesReadParams() *SystemJobTemplatesSystemJobTemplatesReadParams {
	var ()
	return &SystemJobTemplatesSystemJobTemplatesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSystemJobTemplatesSystemJobTemplatesReadParamsWithTimeout creates a new SystemJobTemplatesSystemJobTemplatesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSystemJobTemplatesSystemJobTemplatesReadParamsWithTimeout(timeout time.Duration) *SystemJobTemplatesSystemJobTemplatesReadParams {
	var ()
	return &SystemJobTemplatesSystemJobTemplatesReadParams{

		timeout: timeout,
	}
}

// NewSystemJobTemplatesSystemJobTemplatesReadParamsWithContext creates a new SystemJobTemplatesSystemJobTemplatesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewSystemJobTemplatesSystemJobTemplatesReadParamsWithContext(ctx context.Context) *SystemJobTemplatesSystemJobTemplatesReadParams {
	var ()
	return &SystemJobTemplatesSystemJobTemplatesReadParams{

		Context: ctx,
	}
}

// NewSystemJobTemplatesSystemJobTemplatesReadParamsWithHTTPClient creates a new SystemJobTemplatesSystemJobTemplatesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSystemJobTemplatesSystemJobTemplatesReadParamsWithHTTPClient(client *http.Client) *SystemJobTemplatesSystemJobTemplatesReadParams {
	var ()
	return &SystemJobTemplatesSystemJobTemplatesReadParams{
		HTTPClient: client,
	}
}

/*SystemJobTemplatesSystemJobTemplatesReadParams contains all the parameters to send to the API endpoint
for the system job templates system job templates read operation typically these are written to a http.Request
*/
type SystemJobTemplatesSystemJobTemplatesReadParams struct {

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

// WithTimeout adds the timeout to the system job templates system job templates read params
func (o *SystemJobTemplatesSystemJobTemplatesReadParams) WithTimeout(timeout time.Duration) *SystemJobTemplatesSystemJobTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system job templates system job templates read params
func (o *SystemJobTemplatesSystemJobTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system job templates system job templates read params
func (o *SystemJobTemplatesSystemJobTemplatesReadParams) WithContext(ctx context.Context) *SystemJobTemplatesSystemJobTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system job templates system job templates read params
func (o *SystemJobTemplatesSystemJobTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system job templates system job templates read params
func (o *SystemJobTemplatesSystemJobTemplatesReadParams) WithHTTPClient(client *http.Client) *SystemJobTemplatesSystemJobTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system job templates system job templates read params
func (o *SystemJobTemplatesSystemJobTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the system job templates system job templates read params
func (o *SystemJobTemplatesSystemJobTemplatesReadParams) WithID(id string) *SystemJobTemplatesSystemJobTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the system job templates system job templates read params
func (o *SystemJobTemplatesSystemJobTemplatesReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the system job templates system job templates read params
func (o *SystemJobTemplatesSystemJobTemplatesReadParams) WithSearch(search *string) *SystemJobTemplatesSystemJobTemplatesReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the system job templates system job templates read params
func (o *SystemJobTemplatesSystemJobTemplatesReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *SystemJobTemplatesSystemJobTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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