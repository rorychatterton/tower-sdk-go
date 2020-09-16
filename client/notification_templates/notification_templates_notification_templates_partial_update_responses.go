// Code generated by go-swagger; DO NOT EDIT.

package notification_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotificationTemplatesNotificationTemplatesPartialUpdateReader is a Reader for the NotificationTemplatesNotificationTemplatesPartialUpdate structure.
type NotificationTemplatesNotificationTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationTemplatesNotificationTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationTemplatesNotificationTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNotificationTemplatesNotificationTemplatesPartialUpdateOK creates a NotificationTemplatesNotificationTemplatesPartialUpdateOK with default headers values
func NewNotificationTemplatesNotificationTemplatesPartialUpdateOK() *NotificationTemplatesNotificationTemplatesPartialUpdateOK {
	return &NotificationTemplatesNotificationTemplatesPartialUpdateOK{}
}

/*NotificationTemplatesNotificationTemplatesPartialUpdateOK handles this case with default header values.

OK
*/
type NotificationTemplatesNotificationTemplatesPartialUpdateOK struct {
}

func (o *NotificationTemplatesNotificationTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/notification_templates/{id}/][%d] notificationTemplatesNotificationTemplatesPartialUpdateOK ", 200)
}

func (o *NotificationTemplatesNotificationTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*NotificationTemplatesNotificationTemplatesPartialUpdateBody notification templates notification templates partial update body
swagger:model NotificationTemplatesNotificationTemplatesPartialUpdateBody
*/
type NotificationTemplatesNotificationTemplatesPartialUpdateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// Optional custom messages for notification template.
	Messages string `json:"messages,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// notification configuration
	NotificationConfiguration string `json:"notification_configuration,omitempty"`

	// notification type
	NotificationType string `json:"notification_type,omitempty"`

	// organization
	Organization int64 `json:"organization,omitempty"`
}

// Validate validates this notification templates notification templates partial update body
func (o *NotificationTemplatesNotificationTemplatesPartialUpdateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NotificationTemplatesNotificationTemplatesPartialUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NotificationTemplatesNotificationTemplatesPartialUpdateBody) UnmarshalBinary(b []byte) error {
	var res NotificationTemplatesNotificationTemplatesPartialUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}