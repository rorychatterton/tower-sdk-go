// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoriesInventoriesInventorySourcesCreateReader is a Reader for the InventoriesInventoriesInventorySourcesCreate structure.
type InventoriesInventoriesInventorySourcesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesInventorySourcesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInventoriesInventoriesInventorySourcesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInventoriesInventoriesInventorySourcesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInventoriesInventoriesInventorySourcesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventoriesInventoriesInventorySourcesCreateCreated creates a InventoriesInventoriesInventorySourcesCreateCreated with default headers values
func NewInventoriesInventoriesInventorySourcesCreateCreated() *InventoriesInventoriesInventorySourcesCreateCreated {
	return &InventoriesInventoriesInventorySourcesCreateCreated{}
}

/*InventoriesInventoriesInventorySourcesCreateCreated handles this case with default header values.

InventoriesInventoriesInventorySourcesCreateCreated inventories inventories inventory sources create created
*/
type InventoriesInventoriesInventorySourcesCreateCreated struct {
}

func (o *InventoriesInventoriesInventorySourcesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/inventories/{id}/inventory_sources/][%d] inventoriesInventoriesInventorySourcesCreateCreated ", 201)
}

func (o *InventoriesInventoriesInventorySourcesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInventoriesInventoriesInventorySourcesCreateBadRequest creates a InventoriesInventoriesInventorySourcesCreateBadRequest with default headers values
func NewInventoriesInventoriesInventorySourcesCreateBadRequest() *InventoriesInventoriesInventorySourcesCreateBadRequest {
	return &InventoriesInventoriesInventorySourcesCreateBadRequest{}
}

/*InventoriesInventoriesInventorySourcesCreateBadRequest handles this case with default header values.

Bad Request
*/
type InventoriesInventoriesInventorySourcesCreateBadRequest struct {
}

func (o *InventoriesInventoriesInventorySourcesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/inventories/{id}/inventory_sources/][%d] inventoriesInventoriesInventorySourcesCreateBadRequest ", 400)
}

func (o *InventoriesInventoriesInventorySourcesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInventoriesInventoriesInventorySourcesCreateForbidden creates a InventoriesInventoriesInventorySourcesCreateForbidden with default headers values
func NewInventoriesInventoriesInventorySourcesCreateForbidden() *InventoriesInventoriesInventorySourcesCreateForbidden {
	return &InventoriesInventoriesInventorySourcesCreateForbidden{}
}

/*InventoriesInventoriesInventorySourcesCreateForbidden handles this case with default header values.

No Permission Response
*/
type InventoriesInventoriesInventorySourcesCreateForbidden struct {
}

func (o *InventoriesInventoriesInventorySourcesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/inventories/{id}/inventory_sources/][%d] inventoriesInventoriesInventorySourcesCreateForbidden ", 403)
}

func (o *InventoriesInventoriesInventorySourcesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
