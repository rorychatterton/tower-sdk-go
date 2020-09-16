// Code generated by go-swagger; DO NOT EDIT.

package notification_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// NotificationTemplatesNotificationTemplatesCopyListReader is a Reader for the NotificationTemplatesNotificationTemplatesCopyList structure.
type NotificationTemplatesNotificationTemplatesCopyListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationTemplatesNotificationTemplatesCopyListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationTemplatesNotificationTemplatesCopyListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNotificationTemplatesNotificationTemplatesCopyListOK creates a NotificationTemplatesNotificationTemplatesCopyListOK with default headers values
func NewNotificationTemplatesNotificationTemplatesCopyListOK() *NotificationTemplatesNotificationTemplatesCopyListOK {
	return &NotificationTemplatesNotificationTemplatesCopyListOK{}
}

/*NotificationTemplatesNotificationTemplatesCopyListOK handles this case with default header values.

OK
*/
type NotificationTemplatesNotificationTemplatesCopyListOK struct {
}

func (o *NotificationTemplatesNotificationTemplatesCopyListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/notification_templates/{id}/copy/][%d] notificationTemplatesNotificationTemplatesCopyListOK ", 200)
}

func (o *NotificationTemplatesNotificationTemplatesCopyListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
