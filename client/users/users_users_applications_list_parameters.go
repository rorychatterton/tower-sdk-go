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

// NewUsersUsersApplicationsListParams creates a new UsersUsersApplicationsListParams object
// with the default values initialized.
func NewUsersUsersApplicationsListParams() *UsersUsersApplicationsListParams {
	var ()
	return &UsersUsersApplicationsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersUsersApplicationsListParamsWithTimeout creates a new UsersUsersApplicationsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersUsersApplicationsListParamsWithTimeout(timeout time.Duration) *UsersUsersApplicationsListParams {
	var ()
	return &UsersUsersApplicationsListParams{

		timeout: timeout,
	}
}

// NewUsersUsersApplicationsListParamsWithContext creates a new UsersUsersApplicationsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersUsersApplicationsListParamsWithContext(ctx context.Context) *UsersUsersApplicationsListParams {
	var ()
	return &UsersUsersApplicationsListParams{

		Context: ctx,
	}
}

// NewUsersUsersApplicationsListParamsWithHTTPClient creates a new UsersUsersApplicationsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersUsersApplicationsListParamsWithHTTPClient(client *http.Client) *UsersUsersApplicationsListParams {
	var ()
	return &UsersUsersApplicationsListParams{
		HTTPClient: client,
	}
}

/*UsersUsersApplicationsListParams contains all the parameters to send to the API endpoint
for the users users applications list operation typically these are written to a http.Request
*/
type UsersUsersApplicationsListParams struct {

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

// WithTimeout adds the timeout to the users users applications list params
func (o *UsersUsersApplicationsListParams) WithTimeout(timeout time.Duration) *UsersUsersApplicationsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users users applications list params
func (o *UsersUsersApplicationsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users users applications list params
func (o *UsersUsersApplicationsListParams) WithContext(ctx context.Context) *UsersUsersApplicationsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users users applications list params
func (o *UsersUsersApplicationsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users users applications list params
func (o *UsersUsersApplicationsListParams) WithHTTPClient(client *http.Client) *UsersUsersApplicationsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users users applications list params
func (o *UsersUsersApplicationsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users users applications list params
func (o *UsersUsersApplicationsListParams) WithID(id string) *UsersUsersApplicationsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users users applications list params
func (o *UsersUsersApplicationsListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the users users applications list params
func (o *UsersUsersApplicationsListParams) WithPage(page *int64) *UsersUsersApplicationsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the users users applications list params
func (o *UsersUsersApplicationsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the users users applications list params
func (o *UsersUsersApplicationsListParams) WithPageSize(pageSize *int64) *UsersUsersApplicationsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the users users applications list params
func (o *UsersUsersApplicationsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the users users applications list params
func (o *UsersUsersApplicationsListParams) WithSearch(search *string) *UsersUsersApplicationsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the users users applications list params
func (o *UsersUsersApplicationsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *UsersUsersApplicationsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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