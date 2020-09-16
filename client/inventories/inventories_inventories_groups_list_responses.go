// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoriesInventoriesGroupsListReader is a Reader for the InventoriesInventoriesGroupsList structure.
type InventoriesInventoriesGroupsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesGroupsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoriesInventoriesGroupsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventoriesInventoriesGroupsListOK creates a InventoriesInventoriesGroupsListOK with default headers values
func NewInventoriesInventoriesGroupsListOK() *InventoriesInventoriesGroupsListOK {
	return &InventoriesInventoriesGroupsListOK{}
}

/*InventoriesInventoriesGroupsListOK handles this case with default header values.

OK
*/
type InventoriesInventoriesGroupsListOK struct {
}

func (o *InventoriesInventoriesGroupsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/groups/][%d] inventoriesInventoriesGroupsListOK ", 200)
}

func (o *InventoriesInventoriesGroupsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}