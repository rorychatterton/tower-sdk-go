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
)

// NewOrganizationsOrganizationsAdminsCreateParams creates a new OrganizationsOrganizationsAdminsCreateParams object
// with the default values initialized.
func NewOrganizationsOrganizationsAdminsCreateParams() *OrganizationsOrganizationsAdminsCreateParams {
	var ()
	return &OrganizationsOrganizationsAdminsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationsOrganizationsAdminsCreateParamsWithTimeout creates a new OrganizationsOrganizationsAdminsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrganizationsOrganizationsAdminsCreateParamsWithTimeout(timeout time.Duration) *OrganizationsOrganizationsAdminsCreateParams {
	var ()
	return &OrganizationsOrganizationsAdminsCreateParams{

		timeout: timeout,
	}
}

// NewOrganizationsOrganizationsAdminsCreateParamsWithContext creates a new OrganizationsOrganizationsAdminsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrganizationsOrganizationsAdminsCreateParamsWithContext(ctx context.Context) *OrganizationsOrganizationsAdminsCreateParams {
	var ()
	return &OrganizationsOrganizationsAdminsCreateParams{

		Context: ctx,
	}
}

// NewOrganizationsOrganizationsAdminsCreateParamsWithHTTPClient creates a new OrganizationsOrganizationsAdminsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrganizationsOrganizationsAdminsCreateParamsWithHTTPClient(client *http.Client) *OrganizationsOrganizationsAdminsCreateParams {
	var ()
	return &OrganizationsOrganizationsAdminsCreateParams{
		HTTPClient: client,
	}
}

/*OrganizationsOrganizationsAdminsCreateParams contains all the parameters to send to the API endpoint
for the organizations organizations admins create operation typically these are written to a http.Request
*/
type OrganizationsOrganizationsAdminsCreateParams struct {

	/*Data*/
	Data interface{}
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the organizations organizations admins create params
func (o *OrganizationsOrganizationsAdminsCreateParams) WithTimeout(timeout time.Duration) *OrganizationsOrganizationsAdminsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organizations organizations admins create params
func (o *OrganizationsOrganizationsAdminsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organizations organizations admins create params
func (o *OrganizationsOrganizationsAdminsCreateParams) WithContext(ctx context.Context) *OrganizationsOrganizationsAdminsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organizations organizations admins create params
func (o *OrganizationsOrganizationsAdminsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organizations organizations admins create params
func (o *OrganizationsOrganizationsAdminsCreateParams) WithHTTPClient(client *http.Client) *OrganizationsOrganizationsAdminsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organizations organizations admins create params
func (o *OrganizationsOrganizationsAdminsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the organizations organizations admins create params
func (o *OrganizationsOrganizationsAdminsCreateParams) WithData(data interface{}) *OrganizationsOrganizationsAdminsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the organizations organizations admins create params
func (o *OrganizationsOrganizationsAdminsCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the organizations organizations admins create params
func (o *OrganizationsOrganizationsAdminsCreateParams) WithID(id string) *OrganizationsOrganizationsAdminsCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the organizations organizations admins create params
func (o *OrganizationsOrganizationsAdminsCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationsOrganizationsAdminsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}