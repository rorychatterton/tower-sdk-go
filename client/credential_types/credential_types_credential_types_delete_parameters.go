// Code generated by go-swagger; DO NOT EDIT.

package credential_types

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

// NewCredentialTypesCredentialTypesDeleteParams creates a new CredentialTypesCredentialTypesDeleteParams object
// with the default values initialized.
func NewCredentialTypesCredentialTypesDeleteParams() *CredentialTypesCredentialTypesDeleteParams {
	var ()
	return &CredentialTypesCredentialTypesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCredentialTypesCredentialTypesDeleteParamsWithTimeout creates a new CredentialTypesCredentialTypesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCredentialTypesCredentialTypesDeleteParamsWithTimeout(timeout time.Duration) *CredentialTypesCredentialTypesDeleteParams {
	var ()
	return &CredentialTypesCredentialTypesDeleteParams{

		timeout: timeout,
	}
}

// NewCredentialTypesCredentialTypesDeleteParamsWithContext creates a new CredentialTypesCredentialTypesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewCredentialTypesCredentialTypesDeleteParamsWithContext(ctx context.Context) *CredentialTypesCredentialTypesDeleteParams {
	var ()
	return &CredentialTypesCredentialTypesDeleteParams{

		Context: ctx,
	}
}

// NewCredentialTypesCredentialTypesDeleteParamsWithHTTPClient creates a new CredentialTypesCredentialTypesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCredentialTypesCredentialTypesDeleteParamsWithHTTPClient(client *http.Client) *CredentialTypesCredentialTypesDeleteParams {
	var ()
	return &CredentialTypesCredentialTypesDeleteParams{
		HTTPClient: client,
	}
}

/*CredentialTypesCredentialTypesDeleteParams contains all the parameters to send to the API endpoint
for the credential types credential types delete operation typically these are written to a http.Request
*/
type CredentialTypesCredentialTypesDeleteParams struct {

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

// WithTimeout adds the timeout to the credential types credential types delete params
func (o *CredentialTypesCredentialTypesDeleteParams) WithTimeout(timeout time.Duration) *CredentialTypesCredentialTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the credential types credential types delete params
func (o *CredentialTypesCredentialTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the credential types credential types delete params
func (o *CredentialTypesCredentialTypesDeleteParams) WithContext(ctx context.Context) *CredentialTypesCredentialTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the credential types credential types delete params
func (o *CredentialTypesCredentialTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the credential types credential types delete params
func (o *CredentialTypesCredentialTypesDeleteParams) WithHTTPClient(client *http.Client) *CredentialTypesCredentialTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the credential types credential types delete params
func (o *CredentialTypesCredentialTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the credential types credential types delete params
func (o *CredentialTypesCredentialTypesDeleteParams) WithID(id string) *CredentialTypesCredentialTypesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the credential types credential types delete params
func (o *CredentialTypesCredentialTypesDeleteParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the credential types credential types delete params
func (o *CredentialTypesCredentialTypesDeleteParams) WithSearch(search *string) *CredentialTypesCredentialTypesDeleteParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the credential types credential types delete params
func (o *CredentialTypesCredentialTypesDeleteParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *CredentialTypesCredentialTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
