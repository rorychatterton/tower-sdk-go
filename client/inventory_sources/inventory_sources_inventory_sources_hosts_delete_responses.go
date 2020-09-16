// Code generated by go-swagger; DO NOT EDIT.

package inventory_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventorySourcesInventorySourcesHostsDeleteReader is a Reader for the InventorySourcesInventorySourcesHostsDelete structure.
type InventorySourcesInventorySourcesHostsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventorySourcesInventorySourcesHostsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewInventorySourcesInventorySourcesHostsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewInventorySourcesInventorySourcesHostsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventorySourcesInventorySourcesHostsDeleteNoContent creates a InventorySourcesInventorySourcesHostsDeleteNoContent with default headers values
func NewInventorySourcesInventorySourcesHostsDeleteNoContent() *InventorySourcesInventorySourcesHostsDeleteNoContent {
	return &InventorySourcesInventorySourcesHostsDeleteNoContent{}
}

/*InventorySourcesInventorySourcesHostsDeleteNoContent handles this case with default header values.

InventorySourcesInventorySourcesHostsDeleteNoContent inventory sources inventory sources hosts delete no content
*/
type InventorySourcesInventorySourcesHostsDeleteNoContent struct {
}

func (o *InventorySourcesInventorySourcesHostsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/inventory_sources/{id}/hosts/][%d] inventorySourcesInventorySourcesHostsDeleteNoContent ", 204)
}

func (o *InventorySourcesInventorySourcesHostsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInventorySourcesInventorySourcesHostsDeleteForbidden creates a InventorySourcesInventorySourcesHostsDeleteForbidden with default headers values
func NewInventorySourcesInventorySourcesHostsDeleteForbidden() *InventorySourcesInventorySourcesHostsDeleteForbidden {
	return &InventorySourcesInventorySourcesHostsDeleteForbidden{}
}

/*InventorySourcesInventorySourcesHostsDeleteForbidden handles this case with default header values.

No Permission Response
*/
type InventorySourcesInventorySourcesHostsDeleteForbidden struct {
}

func (o *InventorySourcesInventorySourcesHostsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/inventory_sources/{id}/hosts/][%d] inventorySourcesInventorySourcesHostsDeleteForbidden ", 403)
}

func (o *InventorySourcesInventorySourcesHostsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
