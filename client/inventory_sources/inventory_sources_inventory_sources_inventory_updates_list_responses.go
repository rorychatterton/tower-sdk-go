// Code generated by go-swagger; DO NOT EDIT.

package inventory_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventorySourcesInventorySourcesInventoryUpdatesListReader is a Reader for the InventorySourcesInventorySourcesInventoryUpdatesList structure.
type InventorySourcesInventorySourcesInventoryUpdatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventorySourcesInventorySourcesInventoryUpdatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventorySourcesInventorySourcesInventoryUpdatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventorySourcesInventorySourcesInventoryUpdatesListOK creates a InventorySourcesInventorySourcesInventoryUpdatesListOK with default headers values
func NewInventorySourcesInventorySourcesInventoryUpdatesListOK() *InventorySourcesInventorySourcesInventoryUpdatesListOK {
	return &InventorySourcesInventorySourcesInventoryUpdatesListOK{}
}

/*InventorySourcesInventorySourcesInventoryUpdatesListOK handles this case with default header values.

OK
*/
type InventorySourcesInventorySourcesInventoryUpdatesListOK struct {
}

func (o *InventorySourcesInventorySourcesInventoryUpdatesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventory_sources/{id}/inventory_updates/][%d] inventorySourcesInventorySourcesInventoryUpdatesListOK ", 200)
}

func (o *InventorySourcesInventorySourcesInventoryUpdatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}