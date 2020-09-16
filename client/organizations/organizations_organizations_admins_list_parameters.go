// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewOrganizationsOrganizationsAdminsListParams creates a new OrganizationsOrganizationsAdminsListParams object
// with the default values initialized.
func NewOrganizationsOrganizationsAdminsListParams() *OrganizationsOrganizationsAdminsListParams {
	var ()
	return &OrganizationsOrganizationsAdminsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationsOrganizationsAdminsListParamsWithTimeout creates a new OrganizationsOrganizationsAdminsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrganizationsOrganizationsAdminsListParamsWithTimeout(timeout time.Duration) *OrganizationsOrganizationsAdminsListParams {
	var ()
	return &OrganizationsOrganizationsAdminsListParams{

		timeout: timeout,
	}
}

// NewOrganizationsOrganizationsAdminsListParamsWithContext creates a new OrganizationsOrganizationsAdminsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrganizationsOrganizationsAdminsListParamsWithContext(ctx context.Context) *OrganizationsOrganizationsAdminsListParams {
	var ()
	return &OrganizationsOrganizationsAdminsListParams{

		Context: ctx,
	}
}

// NewOrganizationsOrganizationsAdminsListParamsWithHTTPClient creates a new OrganizationsOrganizationsAdminsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrganizationsOrganizationsAdminsListParamsWithHTTPClient(client *http.Client) *OrganizationsOrganizationsAdminsListParams {
	var ()
	return &OrganizationsOrganizationsAdminsListParams{
		HTTPClient: client,
	}
}

/*OrganizationsOrganizationsAdminsListParams contains all the parameters to send to the API endpoint
for the organizations organizations admins list operation typically these are written to a http.Request
*/
type OrganizationsOrganizationsAdminsListParams struct {

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

// WithTimeout adds the timeout to the organizations organizations admins list params
func (o *OrganizationsOrganizationsAdminsListParams) WithTimeout(timeout time.Duration) *OrganizationsOrganizationsAdminsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organizations organizations admins list params
func (o *OrganizationsOrganizationsAdminsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organizations organizations admins list params
func (o *OrganizationsOrganizationsAdminsListParams) WithContext(ctx context.Context) *OrganizationsOrganizationsAdminsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organizations organizations admins list params
func (o *OrganizationsOrganizationsAdminsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organizations organizations admins list params
func (o *OrganizationsOrganizationsAdminsListParams) WithHTTPClient(client *http.Client) *OrganizationsOrganizationsAdminsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organizations organizations admins list params
func (o *OrganizationsOrganizationsAdminsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the organizations organizations admins list params
func (o *OrganizationsOrganizationsAdminsListParams) WithID(id string) *OrganizationsOrganizationsAdminsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the organizations organizations admins list params
func (o *OrganizationsOrganizationsAdminsListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the organizations organizations admins list params
func (o *OrganizationsOrganizationsAdminsListParams) WithPage(page *int64) *OrganizationsOrganizationsAdminsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the organizations organizations admins list params
func (o *OrganizationsOrganizationsAdminsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the organizations organizations admins list params
func (o *OrganizationsOrganizationsAdminsListParams) WithPageSize(pageSize *int64) *OrganizationsOrganizationsAdminsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the organizations organizations admins list params
func (o *OrganizationsOrganizationsAdminsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the organizations organizations admins list params
func (o *OrganizationsOrganizationsAdminsListParams) WithSearch(search *string) *OrganizationsOrganizationsAdminsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the organizations organizations admins list params
func (o *OrganizationsOrganizationsAdminsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationsOrganizationsAdminsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
