// Code generated by go-swagger; DO NOT EDIT.

package inventory_updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoryUpdatesInventoryUpdatesEventsListReader is a Reader for the InventoryUpdatesInventoryUpdatesEventsList structure.
type InventoryUpdatesInventoryUpdatesEventsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryUpdatesInventoryUpdatesEventsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoryUpdatesInventoryUpdatesEventsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventoryUpdatesInventoryUpdatesEventsListOK creates a InventoryUpdatesInventoryUpdatesEventsListOK with default headers values
func NewInventoryUpdatesInventoryUpdatesEventsListOK() *InventoryUpdatesInventoryUpdatesEventsListOK {
	return &InventoryUpdatesInventoryUpdatesEventsListOK{}
}

/*InventoryUpdatesInventoryUpdatesEventsListOK handles this case with default header values.

OK
*/
type InventoryUpdatesInventoryUpdatesEventsListOK struct {
}

func (o *InventoryUpdatesInventoryUpdatesEventsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventory_updates/{id}/events/][%d] inventoryUpdatesInventoryUpdatesEventsListOK ", 200)
}

func (o *InventoryUpdatesInventoryUpdatesEventsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}