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

// NewCredentialTypesCredentialTypesPartialUpdateParams creates a new CredentialTypesCredentialTypesPartialUpdateParams object
// with the default values initialized.
func NewCredentialTypesCredentialTypesPartialUpdateParams() *CredentialTypesCredentialTypesPartialUpdateParams {
	var ()
	return &CredentialTypesCredentialTypesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCredentialTypesCredentialTypesPartialUpdateParamsWithTimeout creates a new CredentialTypesCredentialTypesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCredentialTypesCredentialTypesPartialUpdateParamsWithTimeout(timeout time.Duration) *CredentialTypesCredentialTypesPartialUpdateParams {
	var ()
	return &CredentialTypesCredentialTypesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewCredentialTypesCredentialTypesPartialUpdateParamsWithContext creates a new CredentialTypesCredentialTypesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCredentialTypesCredentialTypesPartialUpdateParamsWithContext(ctx context.Context) *CredentialTypesCredentialTypesPartialUpdateParams {
	var ()
	return &CredentialTypesCredentialTypesPartialUpdateParams{

		Context: ctx,
	}
}

// NewCredentialTypesCredentialTypesPartialUpdateParamsWithHTTPClient creates a new CredentialTypesCredentialTypesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCredentialTypesCredentialTypesPartialUpdateParamsWithHTTPClient(client *http.Client) *CredentialTypesCredentialTypesPartialUpdateParams {
	var ()
	return &CredentialTypesCredentialTypesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*CredentialTypesCredentialTypesPartialUpdateParams contains all the parameters to send to the API endpoint
for the credential types credential types partial update operation typically these are written to a http.Request
*/
type CredentialTypesCredentialTypesPartialUpdateParams struct {

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

// WithTimeout adds the timeout to the credential types credential types partial update params
func (o *CredentialTypesCredentialTypesPartialUpdateParams) WithTimeout(timeout time.Duration) *CredentialTypesCredentialTypesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the credential types credential types partial update params
func (o *CredentialTypesCredentialTypesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the credential types credential types partial update params
func (o *CredentialTypesCredentialTypesPartialUpdateParams) WithContext(ctx context.Context) *CredentialTypesCredentialTypesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the credential types credential types partial update params
func (o *CredentialTypesCredentialTypesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the credential types credential types partial update params
func (o *CredentialTypesCredentialTypesPartialUpdateParams) WithHTTPClient(client *http.Client) *CredentialTypesCredentialTypesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the credential types credential types partial update params
func (o *CredentialTypesCredentialTypesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the credential types credential types partial update params
func (o *CredentialTypesCredentialTypesPartialUpdateParams) WithData(data interface{}) *CredentialTypesCredentialTypesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the credential types credential types partial update params
func (o *CredentialTypesCredentialTypesPartialUpdateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the credential types credential types partial update params
func (o *CredentialTypesCredentialTypesPartialUpdateParams) WithID(id string) *CredentialTypesCredentialTypesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the credential types credential types partial update params
func (o *CredentialTypesCredentialTypesPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the credential types credential types partial update params
func (o *CredentialTypesCredentialTypesPartialUpdateParams) WithSearch(search *string) *CredentialTypesCredentialTypesPartialUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the credential types credential types partial update params
func (o *CredentialTypesCredentialTypesPartialUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *CredentialTypesCredentialTypesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
