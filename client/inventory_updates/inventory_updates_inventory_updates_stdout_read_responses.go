// Code generated by go-swagger; DO NOT EDIT.

package inventory_updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoryUpdatesInventoryUpdatesStdoutReadReader is a Reader for the InventoryUpdatesInventoryUpdatesStdoutRead structure.
type InventoryUpdatesInventoryUpdatesStdoutReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryUpdatesInventoryUpdatesStdoutReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoryUpdatesInventoryUpdatesStdoutReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventoryUpdatesInventoryUpdatesStdoutReadOK creates a InventoryUpdatesInventoryUpdatesStdoutReadOK with default headers values
func NewInventoryUpdatesInventoryUpdatesStdoutReadOK() *InventoryUpdatesInventoryUpdatesStdoutReadOK {
	return &InventoryUpdatesInventoryUpdatesStdoutReadOK{}
}

/*InventoryUpdatesInventoryUpdatesStdoutReadOK handles this case with default header values.

OK
*/
type InventoryUpdatesInventoryUpdatesStdoutReadOK struct {
}

func (o *InventoryUpdatesInventoryUpdatesStdoutReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventory_updates/{id}/stdout/][%d] inventoryUpdatesInventoryUpdatesStdoutReadOK ", 200)
}

func (o *InventoryUpdatesInventoryUpdatesStdoutReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
