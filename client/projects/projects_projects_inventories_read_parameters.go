// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewProjectsProjectsInventoriesReadParams creates a new ProjectsProjectsInventoriesReadParams object
// with the default values initialized.
func NewProjectsProjectsInventoriesReadParams() *ProjectsProjectsInventoriesReadParams {
	var ()
	return &ProjectsProjectsInventoriesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsProjectsInventoriesReadParamsWithTimeout creates a new ProjectsProjectsInventoriesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsProjectsInventoriesReadParamsWithTimeout(timeout time.Duration) *ProjectsProjectsInventoriesReadParams {
	var ()
	return &ProjectsProjectsInventoriesReadParams{

		timeout: timeout,
	}
}

// NewProjectsProjectsInventoriesReadParamsWithContext creates a new ProjectsProjectsInventoriesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsProjectsInventoriesReadParamsWithContext(ctx context.Context) *ProjectsProjectsInventoriesReadParams {
	var ()
	return &ProjectsProjectsInventoriesReadParams{

		Context: ctx,
	}
}

// NewProjectsProjectsInventoriesReadParamsWithHTTPClient creates a new ProjectsProjectsInventoriesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsProjectsInventoriesReadParamsWithHTTPClient(client *http.Client) *ProjectsProjectsInventoriesReadParams {
	var ()
	return &ProjectsProjectsInventoriesReadParams{
		HTTPClient: client,
	}
}

/*ProjectsProjectsInventoriesReadParams contains all the parameters to send to the API endpoint
for the projects projects inventories read operation typically these are written to a http.Request
*/
type ProjectsProjectsInventoriesReadParams struct {

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

// WithTimeout adds the timeout to the projects projects inventories read params
func (o *ProjectsProjectsInventoriesReadParams) WithTimeout(timeout time.Duration) *ProjectsProjectsInventoriesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects projects inventories read params
func (o *ProjectsProjectsInventoriesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects projects inventories read params
func (o *ProjectsProjectsInventoriesReadParams) WithContext(ctx context.Context) *ProjectsProjectsInventoriesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects projects inventories read params
func (o *ProjectsProjectsInventoriesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects projects inventories read params
func (o *ProjectsProjectsInventoriesReadParams) WithHTTPClient(client *http.Client) *ProjectsProjectsInventoriesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects projects inventories read params
func (o *ProjectsProjectsInventoriesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the projects projects inventories read params
func (o *ProjectsProjectsInventoriesReadParams) WithID(id string) *ProjectsProjectsInventoriesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects projects inventories read params
func (o *ProjectsProjectsInventoriesReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the projects projects inventories read params
func (o *ProjectsProjectsInventoriesReadParams) WithSearch(search *string) *ProjectsProjectsInventoriesReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the projects projects inventories read params
func (o *ProjectsProjectsInventoriesReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsProjectsInventoriesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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