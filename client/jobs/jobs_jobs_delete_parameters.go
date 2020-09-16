// Code generated by go-swagger; DO NOT EDIT.

package jobs

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

// NewJobsJobsDeleteParams creates a new JobsJobsDeleteParams object
// with the default values initialized.
func NewJobsJobsDeleteParams() *JobsJobsDeleteParams {
	var ()
	return &JobsJobsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJobsJobsDeleteParamsWithTimeout creates a new JobsJobsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobsJobsDeleteParamsWithTimeout(timeout time.Duration) *JobsJobsDeleteParams {
	var ()
	return &JobsJobsDeleteParams{

		timeout: timeout,
	}
}

// NewJobsJobsDeleteParamsWithContext creates a new JobsJobsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobsJobsDeleteParamsWithContext(ctx context.Context) *JobsJobsDeleteParams {
	var ()
	return &JobsJobsDeleteParams{

		Context: ctx,
	}
}

// NewJobsJobsDeleteParamsWithHTTPClient creates a new JobsJobsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJobsJobsDeleteParamsWithHTTPClient(client *http.Client) *JobsJobsDeleteParams {
	var ()
	return &JobsJobsDeleteParams{
		HTTPClient: client,
	}
}

/*JobsJobsDeleteParams contains all the parameters to send to the API endpoint
for the jobs jobs delete operation typically these are written to a http.Request
*/
type JobsJobsDeleteParams struct {

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

// WithTimeout adds the timeout to the jobs jobs delete params
func (o *JobsJobsDeleteParams) WithTimeout(timeout time.Duration) *JobsJobsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the jobs jobs delete params
func (o *JobsJobsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the jobs jobs delete params
func (o *JobsJobsDeleteParams) WithContext(ctx context.Context) *JobsJobsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the jobs jobs delete params
func (o *JobsJobsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the jobs jobs delete params
func (o *JobsJobsDeleteParams) WithHTTPClient(client *http.Client) *JobsJobsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the jobs jobs delete params
func (o *JobsJobsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the jobs jobs delete params
func (o *JobsJobsDeleteParams) WithID(id string) *JobsJobsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the jobs jobs delete params
func (o *JobsJobsDeleteParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the jobs jobs delete params
func (o *JobsJobsDeleteParams) WithSearch(search *string) *JobsJobsDeleteParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the jobs jobs delete params
func (o *JobsJobsDeleteParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *JobsJobsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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