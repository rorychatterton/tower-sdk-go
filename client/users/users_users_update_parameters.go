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
)

// NewUsersUsersUpdateParams creates a new UsersUsersUpdateParams object
// with the default values initialized.
func NewUsersUsersUpdateParams() *UsersUsersUpdateParams {
	var ()
	return &UsersUsersUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersUsersUpdateParamsWithTimeout creates a new UsersUsersUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersUsersUpdateParamsWithTimeout(timeout time.Duration) *UsersUsersUpdateParams {
	var ()
	return &UsersUsersUpdateParams{

		timeout: timeout,
	}
}

// NewUsersUsersUpdateParamsWithContext creates a new UsersUsersUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersUsersUpdateParamsWithContext(ctx context.Context) *UsersUsersUpdateParams {
	var ()
	return &UsersUsersUpdateParams{

		Context: ctx,
	}
}

// NewUsersUsersUpdateParamsWithHTTPClient creates a new UsersUsersUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersUsersUpdateParamsWithHTTPClient(client *http.Client) *UsersUsersUpdateParams {
	var ()
	return &UsersUsersUpdateParams{
		HTTPClient: client,
	}
}

/*UsersUsersUpdateParams contains all the parameters to send to the API endpoint
for the users users update operation typically these are written to a http.Request
*/
type UsersUsersUpdateParams struct {

	/*Data*/
	Data UsersUsersUpdateBody
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

// WithTimeout adds the timeout to the users users update params
func (o *UsersUsersUpdateParams) WithTimeout(timeout time.Duration) *UsersUsersUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users users update params
func (o *UsersUsersUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users users update params
func (o *UsersUsersUpdateParams) WithContext(ctx context.Context) *UsersUsersUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users users update params
func (o *UsersUsersUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users users update params
func (o *UsersUsersUpdateParams) WithHTTPClient(client *http.Client) *UsersUsersUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users users update params
func (o *UsersUsersUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users users update params
func (o *UsersUsersUpdateParams) WithData(data UsersUsersUpdateBody) *UsersUsersUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users users update params
func (o *UsersUsersUpdateParams) SetData(data UsersUsersUpdateBody) {
	o.Data = data
}

// WithID adds the id to the users users update params
func (o *UsersUsersUpdateParams) WithID(id string) *UsersUsersUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users users update params
func (o *UsersUsersUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the users users update params
func (o *UsersUsersUpdateParams) WithSearch(search *string) *UsersUsersUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the users users update params
func (o *UsersUsersUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *UsersUsersUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

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
