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

// NotificationTemplatesNotificationTemplatesCreateReader is a Reader for the NotificationTemplatesNotificationTemplatesCreate structure.
type NotificationTemplatesNotificationTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationTemplatesNotificationTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNotificationTemplatesNotificationTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNotificationTemplatesNotificationTemplatesCreateCreated creates a NotificationTemplatesNotificationTemplatesCreateCreated with default headers values
func NewNotificationTemplatesNotificationTemplatesCreateCreated() *NotificationTemplatesNotificationTemplatesCreateCreated {
	return &NotificationTemplatesNotificationTemplatesCreateCreated{}
}

/*NotificationTemplatesNotificationTemplatesCreateCreated handles this case with default header values.

NotificationTemplatesNotificationTemplatesCreateCreated notification templates notification templates create created
*/
type NotificationTemplatesNotificationTemplatesCreateCreated struct {
}

func (o *NotificationTemplatesNotificationTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/notification_templates/][%d] notificationTemplatesNotificationTemplatesCreateCreated ", 201)
}

func (o *NotificationTemplatesNotificationTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*NotificationTemplatesNotificationTemplatesCreateBody notification templates notification templates create body
swagger:model NotificationTemplatesNotificationTemplatesCreateBody
*/
type NotificationTemplatesNotificationTemplatesCreateBody struct {

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

// Validate validates this notification templates notification templates create body
func (o *NotificationTemplatesNotificationTemplatesCreateBody) Validate(formats strfmt.Registry) error {
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

func (o *NotificationTemplatesNotificationTemplatesCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *NotificationTemplatesNotificationTemplatesCreateBody) validateNotificationType(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"notification_type", "body", o.NotificationType); err != nil {
		return err
	}

	return nil
}

func (o *NotificationTemplatesNotificationTemplatesCreateBody) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organization", "body", o.Organization); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NotificationTemplatesNotificationTemplatesCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NotificationTemplatesNotificationTemplatesCreateBody) UnmarshalBinary(b []byte) error {
	var res NotificationTemplatesNotificationTemplatesCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
