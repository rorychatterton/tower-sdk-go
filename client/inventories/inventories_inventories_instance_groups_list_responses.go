// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoriesInventoriesInstanceGroupsListReader is a Reader for the InventoriesInventoriesInstanceGroupsList structure.
type InventoriesInventoriesInstanceGroupsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesInstanceGroupsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoriesInventoriesInstanceGroupsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventoriesInventoriesInstanceGroupsListOK creates a InventoriesInventoriesInstanceGroupsListOK with default headers values
func NewInventoriesInventoriesInstanceGroupsListOK() *InventoriesInventoriesInstanceGroupsListOK {
	return &InventoriesInventoriesInstanceGroupsListOK{}
}

/*InventoriesInventoriesInstanceGroupsListOK handles this case with default header values.

OK
*/
type InventoriesInventoriesInstanceGroupsListOK struct {
}

func (o *InventoriesInventoriesInstanceGroupsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/instance_groups/][%d] inventoriesInventoriesInstanceGroupsListOK ", 200)
}

func (o *InventoriesInventoriesInstanceGroupsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
