// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUsersUsersAdminOfOrganizationsListParams creates a new UsersUsersAdminOfOrganizationsListParams object
// with the default values initialized.
func NewUsersUsersAdminOfOrganizationsListParams() *UsersUsersAdminOfOrganizationsListParams {
	var ()
	return &UsersUsersAdminOfOrganizationsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersUsersAdminOfOrganizationsListParamsWithTimeout creates a new UsersUsersAdminOfOrganizationsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersUsersAdminOfOrganizationsListParamsWithTimeout(timeout time.Duration) *UsersUsersAdminOfOrganizationsListParams {
	var ()
	return &UsersUsersAdminOfOrganizationsListParams{

		timeout: timeout,
	}
}

// NewUsersUsersAdminOfOrganizationsListParamsWithContext creates a new UsersUsersAdminOfOrganizationsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersUsersAdminOfOrganizationsListParamsWithContext(ctx context.Context) *UsersUsersAdminOfOrganizationsListParams {
	var ()
	return &UsersUsersAdminOfOrganizationsListParams{

		Context: ctx,
	}
}

// NewUsersUsersAdminOfOrganizationsListParamsWithHTTPClient creates a new UsersUsersAdminOfOrganizationsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersUsersAdminOfOrganizationsListParamsWithHTTPClient(client *http.Client) *UsersUsersAdminOfOrganizationsListParams {
	var ()
	return &UsersUsersAdminOfOrganizationsListParams{
		HTTPClient: client,
	}
}

/*UsersUsersAdminOfOrganizationsListParams contains all the parameters to send to the API endpoint
for the users users admin of organizations list operation typically these are written to a http.Request
*/
type UsersUsersAdminOfOrganizationsListParams struct {

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

// WithTimeout adds the timeout to the users users admin of organizations list params
func (o *UsersUsersAdminOfOrganizationsListParams) WithTimeout(timeout time.Duration) *UsersUsersAdminOfOrganizationsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users users admin of organizations list params
func (o *UsersUsersAdminOfOrganizationsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users users admin of organizations list params
func (o *UsersUsersAdminOfOrganizationsListParams) WithContext(ctx context.Context) *UsersUsersAdminOfOrganizationsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users users admin of organizations list params
func (o *UsersUsersAdminOfOrganizationsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users users admin of organizations list params
func (o *UsersUsersAdminOfOrganizationsListParams) WithHTTPClient(client *http.Client) *UsersUsersAdminOfOrganizationsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users users admin of organizations list params
func (o *UsersUsersAdminOfOrganizationsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users users admin of organizations list params
func (o *UsersUsersAdminOfOrganizationsListParams) WithID(id string) *UsersUsersAdminOfOrganizationsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users users admin of organizations list params
func (o *UsersUsersAdminOfOrganizationsListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the users users admin of organizations list params
func (o *UsersUsersAdminOfOrganizationsListParams) WithPage(page *int64) *UsersUsersAdminOfOrganizationsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the users users admin of organizations list params
func (o *UsersUsersAdminOfOrganizationsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the users users admin of organizations list params
func (o *UsersUsersAdminOfOrganizationsListParams) WithPageSize(pageSize *int64) *UsersUsersAdminOfOrganizationsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the users users admin of organizations list params
func (o *UsersUsersAdminOfOrganizationsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the users users admin of organizations list params
func (o *UsersUsersAdminOfOrganizationsListParams) WithSearch(search *string) *UsersUsersAdminOfOrganizationsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the users users admin of organizations list params
func (o *UsersUsersAdminOfOrganizationsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *UsersUsersAdminOfOrganizationsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
