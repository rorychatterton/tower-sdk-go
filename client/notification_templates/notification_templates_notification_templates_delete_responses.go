// Code generated by go-swagger; DO NOT EDIT.

package notification_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// NotificationTemplatesNotificationTemplatesDeleteReader is a Reader for the NotificationTemplatesNotificationTemplatesDelete structure.
type NotificationTemplatesNotificationTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationTemplatesNotificationTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewNotificationTemplatesNotificationTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNotificationTemplatesNotificationTemplatesDeleteNoContent creates a NotificationTemplatesNotificationTemplatesDeleteNoContent with default headers values
func NewNotificationTemplatesNotificationTemplatesDeleteNoContent() *NotificationTemplatesNotificationTemplatesDeleteNoContent {
	return &NotificationTemplatesNotificationTemplatesDeleteNoContent{}
}

/*NotificationTemplatesNotificationTemplatesDeleteNoContent handles this case with default header values.

NotificationTemplatesNotificationTemplatesDeleteNoContent notification templates notification templates delete no content
*/
type NotificationTemplatesNotificationTemplatesDeleteNoContent struct {
}

func (o *NotificationTemplatesNotificationTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/notification_templates/{id}/][%d] notificationTemplatesNotificationTemplatesDeleteNoContent ", 204)
}

func (o *NotificationTemplatesNotificationTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
