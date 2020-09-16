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

// NewOrganizationsOrganizationsWorkflowJobTemplatesCreateParams creates a new OrganizationsOrganizationsWorkflowJobTemplatesCreateParams object
// with the default values initialized.
func NewOrganizationsOrganizationsWorkflowJobTemplatesCreateParams() *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams {
	var ()
	return &OrganizationsOrganizationsWorkflowJobTemplatesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationsOrganizationsWorkflowJobTemplatesCreateParamsWithTimeout creates a new OrganizationsOrganizationsWorkflowJobTemplatesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrganizationsOrganizationsWorkflowJobTemplatesCreateParamsWithTimeout(timeout time.Duration) *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams {
	var ()
	return &OrganizationsOrganizationsWorkflowJobTemplatesCreateParams{

		timeout: timeout,
	}
}

// NewOrganizationsOrganizationsWorkflowJobTemplatesCreateParamsWithContext creates a new OrganizationsOrganizationsWorkflowJobTemplatesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrganizationsOrganizationsWorkflowJobTemplatesCreateParamsWithContext(ctx context.Context) *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams {
	var ()
	return &OrganizationsOrganizationsWorkflowJobTemplatesCreateParams{

		Context: ctx,
	}
}

// NewOrganizationsOrganizationsWorkflowJobTemplatesCreateParamsWithHTTPClient creates a new OrganizationsOrganizationsWorkflowJobTemplatesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrganizationsOrganizationsWorkflowJobTemplatesCreateParamsWithHTTPClient(client *http.Client) *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams {
	var ()
	return &OrganizationsOrganizationsWorkflowJobTemplatesCreateParams{
		HTTPClient: client,
	}
}

/*OrganizationsOrganizationsWorkflowJobTemplatesCreateParams contains all the parameters to send to the API endpoint
for the organizations organizations workflow job templates create operation typically these are written to a http.Request
*/
type OrganizationsOrganizationsWorkflowJobTemplatesCreateParams struct {

	/*Data*/
	Data OrganizationsOrganizationsWorkflowJobTemplatesCreateBody
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the organizations organizations workflow job templates create params
func (o *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams) WithTimeout(timeout time.Duration) *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organizations organizations workflow job templates create params
func (o *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organizations organizations workflow job templates create params
func (o *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams) WithContext(ctx context.Context) *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organizations organizations workflow job templates create params
func (o *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organizations organizations workflow job templates create params
func (o *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams) WithHTTPClient(client *http.Client) *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organizations organizations workflow job templates create params
func (o *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the organizations organizations workflow job templates create params
func (o *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams) WithData(data OrganizationsOrganizationsWorkflowJobTemplatesCreateBody) *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the organizations organizations workflow job templates create params
func (o *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams) SetData(data OrganizationsOrganizationsWorkflowJobTemplatesCreateBody) {
	o.Data = data
}

// WithID adds the id to the organizations organizations workflow job templates create params
func (o *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams) WithID(id string) *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the organizations organizations workflow job templates create params
func (o *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationsOrganizationsWorkflowJobTemplatesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
