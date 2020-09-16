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
	"github.com/go-openapi/swag"
)

// NewProjectsProjectsObjectRolesListParams creates a new ProjectsProjectsObjectRolesListParams object
// with the default values initialized.
func NewProjectsProjectsObjectRolesListParams() *ProjectsProjectsObjectRolesListParams {
	var ()
	return &ProjectsProjectsObjectRolesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsProjectsObjectRolesListParamsWithTimeout creates a new ProjectsProjectsObjectRolesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsProjectsObjectRolesListParamsWithTimeout(timeout time.Duration) *ProjectsProjectsObjectRolesListParams {
	var ()
	return &ProjectsProjectsObjectRolesListParams{

		timeout: timeout,
	}
}

// NewProjectsProjectsObjectRolesListParamsWithContext creates a new ProjectsProjectsObjectRolesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsProjectsObjectRolesListParamsWithContext(ctx context.Context) *ProjectsProjectsObjectRolesListParams {
	var ()
	return &ProjectsProjectsObjectRolesListParams{

		Context: ctx,
	}
}

// NewProjectsProjectsObjectRolesListParamsWithHTTPClient creates a new ProjectsProjectsObjectRolesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsProjectsObjectRolesListParamsWithHTTPClient(client *http.Client) *ProjectsProjectsObjectRolesListParams {
	var ()
	return &ProjectsProjectsObjectRolesListParams{
		HTTPClient: client,
	}
}

/*ProjectsProjectsObjectRolesListParams contains all the parameters to send to the API endpoint
for the projects projects object roles list operation typically these are written to a http.Request
*/
type ProjectsProjectsObjectRolesListParams struct {

	/*ID*/
	ID string
	/*Page
	  A page number within the paginated result set.

	*/
	Page *int64
	/*PageSize
	  Number of results to return per page.

	*/
	PageSize *int64
	/*Search
	  A search term.

	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects projects object roles list params
func (o *ProjectsProjectsObjectRolesListParams) WithTimeout(timeout time.Duration) *ProjectsProjectsObjectRolesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects projects object roles list params
func (o *ProjectsProjectsObjectRolesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects projects object roles list params
func (o *ProjectsProjectsObjectRolesListParams) WithContext(ctx context.Context) *ProjectsProjectsObjectRolesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects projects object roles list params
func (o *ProjectsProjectsObjectRolesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects projects object roles list params
func (o *ProjectsProjectsObjectRolesListParams) WithHTTPClient(client *http.Client) *ProjectsProjectsObjectRolesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects projects object roles list params
func (o *ProjectsProjectsObjectRolesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the projects projects object roles list params
func (o *ProjectsProjectsObjectRolesListParams) WithID(id string) *ProjectsProjectsObjectRolesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects projects object roles list params
func (o *ProjectsProjectsObjectRolesListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the projects projects object roles list params
func (o *ProjectsProjectsObjectRolesListParams) WithPage(page *int64) *ProjectsProjectsObjectRolesListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the projects projects object roles list params
func (o *ProjectsProjectsObjectRolesListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the projects projects object roles list params
func (o *ProjectsProjectsObjectRolesListParams) WithPageSize(pageSize *int64) *ProjectsProjectsObjectRolesListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the projects projects object roles list params
func (o *ProjectsProjectsObjectRolesListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the projects projects object roles list params
func (o *ProjectsProjectsObjectRolesListParams) WithSearch(search *string) *ProjectsProjectsObjectRolesListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the projects projects object roles list params
func (o *ProjectsProjectsObjectRolesListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsProjectsObjectRolesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}

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
