// Code generated by go-swagger; DO NOT EDIT.

package instance_groups

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

// NewInstanceGroupsInstanceGroupsReadParams creates a new InstanceGroupsInstanceGroupsReadParams object
// with the default values initialized.
func NewInstanceGroupsInstanceGroupsReadParams() *InstanceGroupsInstanceGroupsReadParams {
	var ()
	return &InstanceGroupsInstanceGroupsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInstanceGroupsInstanceGroupsReadParamsWithTimeout creates a new InstanceGroupsInstanceGroupsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInstanceGroupsInstanceGroupsReadParamsWithTimeout(timeout time.Duration) *InstanceGroupsInstanceGroupsReadParams {
	var ()
	return &InstanceGroupsInstanceGroupsReadParams{

		timeout: timeout,
	}
}

// NewInstanceGroupsInstanceGroupsReadParamsWithContext creates a new InstanceGroupsInstanceGroupsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewInstanceGroupsInstanceGroupsReadParamsWithContext(ctx context.Context) *InstanceGroupsInstanceGroupsReadParams {
	var ()
	return &InstanceGroupsInstanceGroupsReadParams{

		Context: ctx,
	}
}

// NewInstanceGroupsInstanceGroupsReadParamsWithHTTPClient creates a new InstanceGroupsInstanceGroupsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInstanceGroupsInstanceGroupsReadParamsWithHTTPClient(client *http.Client) *InstanceGroupsInstanceGroupsReadParams {
	var ()
	return &InstanceGroupsInstanceGroupsReadParams{
		HTTPClient: client,
	}
}

/*InstanceGroupsInstanceGroupsReadParams contains all the parameters to send to the API endpoint
for the instance groups instance groups read operation typically these are written to a http.Request
*/
type InstanceGroupsInstanceGroupsReadParams struct {

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

// WithTimeout adds the timeout to the instance groups instance groups read params
func (o *InstanceGroupsInstanceGroupsReadParams) WithTimeout(timeout time.Duration) *InstanceGroupsInstanceGroupsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the instance groups instance groups read params
func (o *InstanceGroupsInstanceGroupsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the instance groups instance groups read params
func (o *InstanceGroupsInstanceGroupsReadParams) WithContext(ctx context.Context) *InstanceGroupsInstanceGroupsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the instance groups instance groups read params
func (o *InstanceGroupsInstanceGroupsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the instance groups instance groups read params
func (o *InstanceGroupsInstanceGroupsReadParams) WithHTTPClient(client *http.Client) *InstanceGroupsInstanceGroupsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the instance groups instance groups read params
func (o *InstanceGroupsInstanceGroupsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the instance groups instance groups read params
func (o *InstanceGroupsInstanceGroupsReadParams) WithID(id string) *InstanceGroupsInstanceGroupsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the instance groups instance groups read params
func (o *InstanceGroupsInstanceGroupsReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the instance groups instance groups read params
func (o *InstanceGroupsInstanceGroupsReadParams) WithSearch(search *string) *InstanceGroupsInstanceGroupsReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the instance groups instance groups read params
func (o *InstanceGroupsInstanceGroupsReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InstanceGroupsInstanceGroupsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
