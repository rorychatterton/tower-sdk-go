// Code generated by go-swagger; DO NOT EDIT.

package inventory_updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoryUpdatesInventoryUpdatesCancelCreateReader is a Reader for the InventoryUpdatesInventoryUpdatesCancelCreate structure.
type InventoryUpdatesInventoryUpdatesCancelCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryUpdatesInventoryUpdatesCancelCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInventoryUpdatesInventoryUpdatesCancelCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventoryUpdatesInventoryUpdatesCancelCreateCreated creates a InventoryUpdatesInventoryUpdatesCancelCreateCreated with default headers values
func NewInventoryUpdatesInventoryUpdatesCancelCreateCreated() *InventoryUpdatesInventoryUpdatesCancelCreateCreated {
	return &InventoryUpdatesInventoryUpdatesCancelCreateCreated{}
}

/*InventoryUpdatesInventoryUpdatesCancelCreateCreated handles this case with default header values.

InventoryUpdatesInventoryUpdatesCancelCreateCreated inventory updates inventory updates cancel create created
*/
type InventoryUpdatesInventoryUpdatesCancelCreateCreated struct {
}

func (o *InventoryUpdatesInventoryUpdatesCancelCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/inventory_updates/{id}/cancel/][%d] inventoryUpdatesInventoryUpdatesCancelCreateCreated ", 201)
}

func (o *InventoryUpdatesInventoryUpdatesCancelCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
