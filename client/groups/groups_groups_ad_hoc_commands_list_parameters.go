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

// NewGroupsGroupsAdHocCommandsListParams creates a new GroupsGroupsAdHocCommandsListParams object
// with the default values initialized.
func NewGroupsGroupsAdHocCommandsListParams() *GroupsGroupsAdHocCommandsListParams {
	var ()
	return &GroupsGroupsAdHocCommandsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupsGroupsAdHocCommandsListParamsWithTimeout creates a new GroupsGroupsAdHocCommandsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupsGroupsAdHocCommandsListParamsWithTimeout(timeout time.Duration) *GroupsGroupsAdHocCommandsListParams {
	var ()
	return &GroupsGroupsAdHocCommandsListParams{

		timeout: timeout,
	}
}

// NewGroupsGroupsAdHocCommandsListParamsWithContext creates a new GroupsGroupsAdHocCommandsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupsGroupsAdHocCommandsListParamsWithContext(ctx context.Context) *GroupsGroupsAdHocCommandsListParams {
	var ()
	return &GroupsGroupsAdHocCommandsListParams{

		Context: ctx,
	}
}

// NewGroupsGroupsAdHocCommandsListParamsWithHTTPClient creates a new GroupsGroupsAdHocCommandsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupsGroupsAdHocCommandsListParamsWithHTTPClient(client *http.Client) *GroupsGroupsAdHocCommandsListParams {
	var ()
	return &GroupsGroupsAdHocCommandsListParams{
		HTTPClient: client,
	}
}

/*GroupsGroupsAdHocCommandsListParams contains all the parameters to send to the API endpoint
for the groups groups ad hoc commands list operation typically these are written to a http.Request
*/
type GroupsGroupsAdHocCommandsListParams struct {

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

// WithTimeout adds the timeout to the groups groups ad hoc commands list params
func (o *GroupsGroupsAdHocCommandsListParams) WithTimeout(timeout time.Duration) *GroupsGroupsAdHocCommandsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the groups groups ad hoc commands list params
func (o *GroupsGroupsAdHocCommandsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the groups groups ad hoc commands list params
func (o *GroupsGroupsAdHocCommandsListParams) WithContext(ctx context.Context) *GroupsGroupsAdHocCommandsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the groups groups ad hoc commands list params
func (o *GroupsGroupsAdHocCommandsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the groups groups ad hoc commands list params
func (o *GroupsGroupsAdHocCommandsListParams) WithHTTPClient(client *http.Client) *GroupsGroupsAdHocCommandsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the groups groups ad hoc commands list params
func (o *GroupsGroupsAdHocCommandsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the groups groups ad hoc commands list params
func (o *GroupsGroupsAdHocCommandsListParams) WithID(id string) *GroupsGroupsAdHocCommandsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the groups groups ad hoc commands list params
func (o *GroupsGroupsAdHocCommandsListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the groups groups ad hoc commands list params
func (o *GroupsGroupsAdHocCommandsListParams) WithPage(page *int64) *GroupsGroupsAdHocCommandsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the groups groups ad hoc commands list params
func (o *GroupsGroupsAdHocCommandsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the groups groups ad hoc commands list params
func (o *GroupsGroupsAdHocCommandsListParams) WithPageSize(pageSize *int64) *GroupsGroupsAdHocCommandsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the groups groups ad hoc commands list params
func (o *GroupsGroupsAdHocCommandsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the groups groups ad hoc commands list params
func (o *GroupsGroupsAdHocCommandsListParams) WithSearch(search *string) *GroupsGroupsAdHocCommandsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the groups groups ad hoc commands list params
func (o *GroupsGroupsAdHocCommandsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GroupsGroupsAdHocCommandsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
