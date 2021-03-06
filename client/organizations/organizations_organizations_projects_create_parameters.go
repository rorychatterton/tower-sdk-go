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

// NewOrganizationsOrganizationsProjectsCreateParams creates a new OrganizationsOrganizationsProjectsCreateParams object
// with the default values initialized.
func NewOrganizationsOrganizationsProjectsCreateParams() *OrganizationsOrganizationsProjectsCreateParams {
	var ()
	return &OrganizationsOrganizationsProjectsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationsOrganizationsProjectsCreateParamsWithTimeout creates a new OrganizationsOrganizationsProjectsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrganizationsOrganizationsProjectsCreateParamsWithTimeout(timeout time.Duration) *OrganizationsOrganizationsProjectsCreateParams {
	var ()
	return &OrganizationsOrganizationsProjectsCreateParams{

		timeout: timeout,
	}
}

// NewOrganizationsOrganizationsProjectsCreateParamsWithContext creates a new OrganizationsOrganizationsProjectsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrganizationsOrganizationsProjectsCreateParamsWithContext(ctx context.Context) *OrganizationsOrganizationsProjectsCreateParams {
	var ()
	return &OrganizationsOrganizationsProjectsCreateParams{

		Context: ctx,
	}
}

// NewOrganizationsOrganizationsProjectsCreateParamsWithHTTPClient creates a new OrganizationsOrganizationsProjectsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrganizationsOrganizationsProjectsCreateParamsWithHTTPClient(client *http.Client) *OrganizationsOrganizationsProjectsCreateParams {
	var ()
	return &OrganizationsOrganizationsProjectsCreateParams{
		HTTPClient: client,
	}
}

/*OrganizationsOrganizationsProjectsCreateParams contains all the parameters to send to the API endpoint
for the organizations organizations projects create operation typically these are written to a http.Request
*/
type OrganizationsOrganizationsProjectsCreateParams struct {

	/*Data*/
	Data OrganizationsOrganizationsProjectsCreateBody
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the organizations organizations projects create params
func (o *OrganizationsOrganizationsProjectsCreateParams) WithTimeout(timeout time.Duration) *OrganizationsOrganizationsProjectsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organizations organizations projects create params
func (o *OrganizationsOrganizationsProjectsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organizations organizations projects create params
func (o *OrganizationsOrganizationsProjectsCreateParams) WithContext(ctx context.Context) *OrganizationsOrganizationsProjectsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organizations organizations projects create params
func (o *OrganizationsOrganizationsProjectsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organizations organizations projects create params
func (o *OrganizationsOrganizationsProjectsCreateParams) WithHTTPClient(client *http.Client) *OrganizationsOrganizationsProjectsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organizations organizations projects create params
func (o *OrganizationsOrganizationsProjectsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the organizations organizations projects create params
func (o *OrganizationsOrganizationsProjectsCreateParams) WithData(data OrganizationsOrganizationsProjectsCreateBody) *OrganizationsOrganizationsProjectsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the organizations organizations projects create params
func (o *OrganizationsOrganizationsProjectsCreateParams) SetData(data OrganizationsOrganizationsProjectsCreateBody) {
	o.Data = data
}

// WithID adds the id to the organizations organizations projects create params
func (o *OrganizationsOrganizationsProjectsCreateParams) WithID(id string) *OrganizationsOrganizationsProjectsCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the organizations organizations projects create params
func (o *OrganizationsOrganizationsProjectsCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationsOrganizationsProjectsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
