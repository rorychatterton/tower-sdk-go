// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// NewAuthenticationApplicationsPartialUpdate0Params creates a new AuthenticationApplicationsPartialUpdate0Params object
// with the default values initialized.
func NewAuthenticationApplicationsPartialUpdate0Params() *AuthenticationApplicationsPartialUpdate0Params {
	var ()
	return &AuthenticationApplicationsPartialUpdate0Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAuthenticationApplicationsPartialUpdate0ParamsWithTimeout creates a new AuthenticationApplicationsPartialUpdate0Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAuthenticationApplicationsPartialUpdate0ParamsWithTimeout(timeout time.Duration) *AuthenticationApplicationsPartialUpdate0Params {
	var ()
	return &AuthenticationApplicationsPartialUpdate0Params{

		timeout: timeout,
	}
}

// NewAuthenticationApplicationsPartialUpdate0ParamsWithContext creates a new AuthenticationApplicationsPartialUpdate0Params object
// with the default values initialized, and the ability to set a context for a request
func NewAuthenticationApplicationsPartialUpdate0ParamsWithContext(ctx context.Context) *AuthenticationApplicationsPartialUpdate0Params {
	var ()
	return &AuthenticationApplicationsPartialUpdate0Params{

		Context: ctx,
	}
}

// NewAuthenticationApplicationsPartialUpdate0ParamsWithHTTPClient creates a new AuthenticationApplicationsPartialUpdate0Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAuthenticationApplicationsPartialUpdate0ParamsWithHTTPClient(client *http.Client) *AuthenticationApplicationsPartialUpdate0Params {
	var ()
	return &AuthenticationApplicationsPartialUpdate0Params{
		HTTPClient: client,
	}
}

/*AuthenticationApplicationsPartialUpdate0Params contains all the parameters to send to the API endpoint
for the authentication applications partial update 0 operation typically these are written to a http.Request
*/
type AuthenticationApplicationsPartialUpdate0Params struct {

	/*Data*/
	Data interface{}
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

// WithTimeout adds the timeout to the authentication applications partial update 0 params
func (o *AuthenticationApplicationsPartialUpdate0Params) WithTimeout(timeout time.Duration) *AuthenticationApplicationsPartialUpdate0Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authentication applications partial update 0 params
func (o *AuthenticationApplicationsPartialUpdate0Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authentication applications partial update 0 params
func (o *AuthenticationApplicationsPartialUpdate0Params) WithContext(ctx context.Context) *AuthenticationApplicationsPartialUpdate0Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authentication applications partial update 0 params
func (o *AuthenticationApplicationsPartialUpdate0Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authentication applications partial update 0 params
func (o *AuthenticationApplicationsPartialUpdate0Params) WithHTTPClient(client *http.Client) *AuthenticationApplicationsPartialUpdate0Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authentication applications partial update 0 params
func (o *AuthenticationApplicationsPartialUpdate0Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the authentication applications partial update 0 params
func (o *AuthenticationApplicationsPartialUpdate0Params) WithData(data interface{}) *AuthenticationApplicationsPartialUpdate0Params {
	o.SetData(data)
	return o
}

// SetData adds the data to the authentication applications partial update 0 params
func (o *AuthenticationApplicationsPartialUpdate0Params) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the authentication applications partial update 0 params
func (o *AuthenticationApplicationsPartialUpdate0Params) WithID(id string) *AuthenticationApplicationsPartialUpdate0Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the authentication applications partial update 0 params
func (o *AuthenticationApplicationsPartialUpdate0Params) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the authentication applications partial update 0 params
func (o *AuthenticationApplicationsPartialUpdate0Params) WithSearch(search *string) *AuthenticationApplicationsPartialUpdate0Params {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the authentication applications partial update 0 params
func (o *AuthenticationApplicationsPartialUpdate0Params) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *AuthenticationApplicationsPartialUpdate0Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
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