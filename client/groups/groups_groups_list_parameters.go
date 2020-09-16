// Code generated by go-swagger; DO NOT EDIT.

package groups

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

// NewGroupsGroupsListParams creates a new GroupsGroupsListParams object
// with the default values initialized.
func NewGroupsGroupsListParams() *GroupsGroupsListParams {
	var ()
	return &GroupsGroupsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupsGroupsListParamsWithTimeout creates a new GroupsGroupsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupsGroupsListParamsWithTimeout(timeout time.Duration) *GroupsGroupsListParams {
	var ()
	return &GroupsGroupsListParams{

		timeout: timeout,
	}
}

// NewGroupsGroupsListParamsWithContext creates a new GroupsGroupsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupsGroupsListParamsWithContext(ctx context.Context) *GroupsGroupsListParams {
	var ()
	return &GroupsGroupsListParams{

		Context: ctx,
	}
}

// NewGroupsGroupsListParamsWithHTTPClient creates a new GroupsGroupsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupsGroupsListParamsWithHTTPClient(client *http.Client) *GroupsGroupsListParams {
	var ()
	return &GroupsGroupsListParams{
		HTTPClient: client,
	}
}

/*GroupsGroupsListParams contains all the parameters to send to the API endpoint
for the groups groups list operation typically these are written to a http.Request
*/
type GroupsGroupsListParams struct {

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

// WithTimeout adds the timeout to the groups groups list params
func (o *GroupsGroupsListParams) WithTimeout(timeout time.Duration) *GroupsGroupsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the groups groups list params
func (o *GroupsGroupsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the groups groups list params
func (o *GroupsGroupsListParams) WithContext(ctx context.Context) *GroupsGroupsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the groups groups list params
func (o *GroupsGroupsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the groups groups list params
func (o *GroupsGroupsListParams) WithHTTPClient(client *http.Client) *GroupsGroupsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the groups groups list params
func (o *GroupsGroupsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the groups groups list params
func (o *GroupsGroupsListParams) WithPage(page *int64) *GroupsGroupsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the groups groups list params
func (o *GroupsGroupsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the groups groups list params
func (o *GroupsGroupsListParams) WithPageSize(pageSize *int64) *GroupsGroupsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the groups groups list params
func (o *GroupsGroupsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the groups groups list params
func (o *GroupsGroupsListParams) WithSearch(search *string) *GroupsGroupsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the groups groups list params
func (o *GroupsGroupsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GroupsGroupsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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