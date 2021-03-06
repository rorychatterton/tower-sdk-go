// Code generated by go-swagger; DO NOT EDIT.

package credential_input_sources

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

// NewCredentialInputSourcesCredentialInputSourcesUpdateParams creates a new CredentialInputSourcesCredentialInputSourcesUpdateParams object
// with the default values initialized.
func NewCredentialInputSourcesCredentialInputSourcesUpdateParams() *CredentialInputSourcesCredentialInputSourcesUpdateParams {
	var ()
	return &CredentialInputSourcesCredentialInputSourcesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCredentialInputSourcesCredentialInputSourcesUpdateParamsWithTimeout creates a new CredentialInputSourcesCredentialInputSourcesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCredentialInputSourcesCredentialInputSourcesUpdateParamsWithTimeout(timeout time.Duration) *CredentialInputSourcesCredentialInputSourcesUpdateParams {
	var ()
	return &CredentialInputSourcesCredentialInputSourcesUpdateParams{

		timeout: timeout,
	}
}

// NewCredentialInputSourcesCredentialInputSourcesUpdateParamsWithContext creates a new CredentialInputSourcesCredentialInputSourcesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCredentialInputSourcesCredentialInputSourcesUpdateParamsWithContext(ctx context.Context) *CredentialInputSourcesCredentialInputSourcesUpdateParams {
	var ()
	return &CredentialInputSourcesCredentialInputSourcesUpdateParams{

		Context: ctx,
	}
}

// NewCredentialInputSourcesCredentialInputSourcesUpdateParamsWithHTTPClient creates a new CredentialInputSourcesCredentialInputSourcesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCredentialInputSourcesCredentialInputSourcesUpdateParamsWithHTTPClient(client *http.Client) *CredentialInputSourcesCredentialInputSourcesUpdateParams {
	var ()
	return &CredentialInputSourcesCredentialInputSourcesUpdateParams{
		HTTPClient: client,
	}
}

/*CredentialInputSourcesCredentialInputSourcesUpdateParams contains all the parameters to send to the API endpoint
for the credential input sources credential input sources update operation typically these are written to a http.Request
*/
type CredentialInputSourcesCredentialInputSourcesUpdateParams struct {

	/*Data*/
	Data CredentialInputSourcesCredentialInputSourcesUpdateBody
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

// WithTimeout adds the timeout to the credential input sources credential input sources update params
func (o *CredentialInputSourcesCredentialInputSourcesUpdateParams) WithTimeout(timeout time.Duration) *CredentialInputSourcesCredentialInputSourcesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the credential input sources credential input sources update params
func (o *CredentialInputSourcesCredentialInputSourcesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the credential input sources credential input sources update params
func (o *CredentialInputSourcesCredentialInputSourcesUpdateParams) WithContext(ctx context.Context) *CredentialInputSourcesCredentialInputSourcesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the credential input sources credential input sources update params
func (o *CredentialInputSourcesCredentialInputSourcesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the credential input sources credential input sources update params
func (o *CredentialInputSourcesCredentialInputSourcesUpdateParams) WithHTTPClient(client *http.Client) *CredentialInputSourcesCredentialInputSourcesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the credential input sources credential input sources update params
func (o *CredentialInputSourcesCredentialInputSourcesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the credential input sources credential input sources update params
func (o *CredentialInputSourcesCredentialInputSourcesUpdateParams) WithData(data CredentialInputSourcesCredentialInputSourcesUpdateBody) *CredentialInputSourcesCredentialInputSourcesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the credential input sources credential input sources update params
func (o *CredentialInputSourcesCredentialInputSourcesUpdateParams) SetData(data CredentialInputSourcesCredentialInputSourcesUpdateBody) {
	o.Data = data
}

// WithID adds the id to the credential input sources credential input sources update params
func (o *CredentialInputSourcesCredentialInputSourcesUpdateParams) WithID(id string) *CredentialInputSourcesCredentialInputSourcesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the credential input sources credential input sources update params
func (o *CredentialInputSourcesCredentialInputSourcesUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the credential input sources credential input sources update params
func (o *CredentialInputSourcesCredentialInputSourcesUpdateParams) WithSearch(search *string) *CredentialInputSourcesCredentialInputSourcesUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the credential input sources credential input sources update params
func (o *CredentialInputSourcesCredentialInputSourcesUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *CredentialInputSourcesCredentialInputSourcesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
