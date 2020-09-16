// Code generated by go-swagger; DO NOT EDIT.

package notification_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NotificationTemplatesNotificationTemplatesUpdateReader is a Reader for the NotificationTemplatesNotificationTemplatesUpdate structure.
type NotificationTemplatesNotificationTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationTemplatesNotificationTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationTemplatesNotificationTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNotificationTemplatesNotificationTemplatesUpdateOK creates a NotificationTemplatesNotificationTemplatesUpdateOK with default headers values
func NewNotificationTemplatesNotificationTemplatesUpdateOK() *NotificationTemplatesNotificationTemplatesUpdateOK {
	return &NotificationTemplatesNotificationTemplatesUpdateOK{}
}

/*NotificationTemplatesNotificationTemplatesUpdateOK handles this case with default header values.

OK
*/
type NotificationTemplatesNotificationTemplatesUpdateOK struct {
}

func (o *NotificationTemplatesNotificationTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/notification_templates/{id}/][%d] notificationTemplatesNotificationTemplatesUpdateOK ", 200)
}

func (o *NotificationTemplatesNotificationTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*NotificationTemplatesNotificationTemplatesUpdateBody notification templates notification templates update body
swagger:model NotificationTemplatesNotificationTemplatesUpdateBody
*/
type NotificationTemplatesNotificationTemplatesUpdateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// Optional custom messages for notification template.
	Messages string `json:"messages,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// notification configuration
	NotificationConfiguration string `json:"notification_configuration,omitempty"`

	// notification type
	// Required: true
	NotificationType *string `json:"notification_type"`

	// organization
	// Required: true
	Organization *int64 `json:"organization"`
}

// Validate validates this notification templates notification templates update body
func (o *NotificationTemplatesNotificationTemplatesUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNotificationType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NotificationTemplatesNotificationTemplatesUpdateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *NotificationTemplatesNotificationTemplatesUpdateBody) validateNotificationType(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"notification_type", "body", o.NotificationType); err != nil {
		return err
	}

	return nil
}

func (o *NotificationTemplatesNotificationTemplatesUpdateBody) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organization", "body", o.Organization); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NotificationTemplatesNotificationTemplatesUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NotificationTemplatesNotificationTemplatesUpdateBody) UnmarshalBinary(b []byte) error {
	var res NotificationTemplatesNotificationTemplatesUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}