// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// NotificationsNotificationsListReader is a Reader for the NotificationsNotificationsList structure.
type NotificationsNotificationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationsNotificationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationsNotificationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNotificationsNotificationsListOK creates a NotificationsNotificationsListOK with default headers values
func NewNotificationsNotificationsListOK() *NotificationsNotificationsListOK {
	return &NotificationsNotificationsListOK{}
}

/*NotificationsNotificationsListOK handles this case with default header values.

OK
*/
type NotificationsNotificationsListOK struct {
}

func (o *NotificationsNotificationsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/][%d] notificationsNotificationsListOK ", 200)
}

func (o *NotificationsNotificationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
