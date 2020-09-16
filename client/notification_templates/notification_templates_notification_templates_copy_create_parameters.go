// Code generated by go-swagger; DO NOT EDIT.

package notification_templates

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

// NewNotificationTemplatesNotificationTemplatesCopyCreateParams creates a new NotificationTemplatesNotificationTemplatesCopyCreateParams object
// with the default values initialized.
func NewNotificationTemplatesNotificationTemplatesCopyCreateParams() *NotificationTemplatesNotificationTemplatesCopyCreateParams {
	var ()
	return &NotificationTemplatesNotificationTemplatesCopyCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNotificationTemplatesNotificationTemplatesCopyCreateParamsWithTimeout creates a new NotificationTemplatesNotificationTemplatesCopyCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNotificationTemplatesNotificationTemplatesCopyCreateParamsWithTimeout(timeout time.Duration) *NotificationTemplatesNotificationTemplatesCopyCreateParams {
	var ()
	return &NotificationTemplatesNotificationTemplatesCopyCreateParams{

		timeout: timeout,
	}
}

// NewNotificationTemplatesNotificationTemplatesCopyCreateParamsWithContext creates a new NotificationTemplatesNotificationTemplatesCopyCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewNotificationTemplatesNotificationTemplatesCopyCreateParamsWithContext(ctx context.Context) *NotificationTemplatesNotificationTemplatesCopyCreateParams {
	var ()
	return &NotificationTemplatesNotificationTemplatesCopyCreateParams{

		Context: ctx,
	}
}

// NewNotificationTemplatesNotificationTemplatesCopyCreateParamsWithHTTPClient creates a new NotificationTemplatesNotificationTemplatesCopyCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNotificationTemplatesNotificationTemplatesCopyCreateParamsWithHTTPClient(client *http.Client) *NotificationTemplatesNotificationTemplatesCopyCreateParams {
	var ()
	return &NotificationTemplatesNotificationTemplatesCopyCreateParams{
		HTTPClient: client,
	}
}

/*NotificationTemplatesNotificationTemplatesCopyCreateParams contains all the parameters to send to the API endpoint
for the notification templates notification templates copy create operation typically these are written to a http.Request
*/
type NotificationTemplatesNotificationTemplatesCopyCreateParams struct {

	/*Data*/
	Data NotificationTemplatesNotificationTemplatesCopyCreateBody
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the notification templates notification templates copy create params
func (o *NotificationTemplatesNotificationTemplatesCopyCreateParams) WithTimeout(timeout time.Duration) *NotificationTemplatesNotificationTemplatesCopyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notification templates notification templates copy create params
func (o *NotificationTemplatesNotificationTemplatesCopyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notification templates notification templates copy create params
func (o *NotificationTemplatesNotificationTemplatesCopyCreateParams) WithContext(ctx context.Context) *NotificationTemplatesNotificationTemplatesCopyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notification templates notification templates copy create params
func (o *NotificationTemplatesNotificationTemplatesCopyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notification templates notification templates copy create params
func (o *NotificationTemplatesNotificationTemplatesCopyCreateParams) WithHTTPClient(client *http.Client) *NotificationTemplatesNotificationTemplatesCopyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notification templates notification templates copy create params
func (o *NotificationTemplatesNotificationTemplatesCopyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the notification templates notification templates copy create params
func (o *NotificationTemplatesNotificationTemplatesCopyCreateParams) WithData(data NotificationTemplatesNotificationTemplatesCopyCreateBody) *NotificationTemplatesNotificationTemplatesCopyCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the notification templates notification templates copy create params
func (o *NotificationTemplatesNotificationTemplatesCopyCreateParams) SetData(data NotificationTemplatesNotificationTemplatesCopyCreateBody) {
	o.Data = data
}

// WithID adds the id to the notification templates notification templates copy create params
func (o *NotificationTemplatesNotificationTemplatesCopyCreateParams) WithID(id string) *NotificationTemplatesNotificationTemplatesCopyCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the notification templates notification templates copy create params
func (o *NotificationTemplatesNotificationTemplatesCopyCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *NotificationTemplatesNotificationTemplatesCopyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
