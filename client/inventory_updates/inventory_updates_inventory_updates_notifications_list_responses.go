// Code generated by go-swagger; DO NOT EDIT.

package inventory_updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoryUpdatesInventoryUpdatesNotificationsListReader is a Reader for the InventoryUpdatesInventoryUpdatesNotificationsList structure.
type InventoryUpdatesInventoryUpdatesNotificationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryUpdatesInventoryUpdatesNotificationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoryUpdatesInventoryUpdatesNotificationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventoryUpdatesInventoryUpdatesNotificationsListOK creates a InventoryUpdatesInventoryUpdatesNotificationsListOK with default headers values
func NewInventoryUpdatesInventoryUpdatesNotificationsListOK() *InventoryUpdatesInventoryUpdatesNotificationsListOK {
	return &InventoryUpdatesInventoryUpdatesNotificationsListOK{}
}

/*InventoryUpdatesInventoryUpdatesNotificationsListOK handles this case with default header values.

OK
*/
type InventoryUpdatesInventoryUpdatesNotificationsListOK struct {
}

func (o *InventoryUpdatesInventoryUpdatesNotificationsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventory_updates/{id}/notifications/][%d] inventoryUpdatesInventoryUpdatesNotificationsListOK ", 200)
}

func (o *InventoryUpdatesInventoryUpdatesNotificationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}