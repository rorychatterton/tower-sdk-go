// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoriesInventoriesAdHocCommandsListReader is a Reader for the InventoriesInventoriesAdHocCommandsList structure.
type InventoriesInventoriesAdHocCommandsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesAdHocCommandsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoriesInventoriesAdHocCommandsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewInventoriesInventoriesAdHocCommandsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventoriesInventoriesAdHocCommandsListOK creates a InventoriesInventoriesAdHocCommandsListOK with default headers values
func NewInventoriesInventoriesAdHocCommandsListOK() *InventoriesInventoriesAdHocCommandsListOK {
	return &InventoriesInventoriesAdHocCommandsListOK{}
}

/*InventoriesInventoriesAdHocCommandsListOK handles this case with default header values.

OK
*/
type InventoriesInventoriesAdHocCommandsListOK struct {
}

func (o *InventoriesInventoriesAdHocCommandsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/ad_hoc_commands/][%d] inventoriesInventoriesAdHocCommandsListOK ", 200)
}

func (o *InventoriesInventoriesAdHocCommandsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInventoriesInventoriesAdHocCommandsListForbidden creates a InventoriesInventoriesAdHocCommandsListForbidden with default headers values
func NewInventoriesInventoriesAdHocCommandsListForbidden() *InventoriesInventoriesAdHocCommandsListForbidden {
	return &InventoriesInventoriesAdHocCommandsListForbidden{}
}

/*InventoriesInventoriesAdHocCommandsListForbidden handles this case with default header values.

No Permission Response
*/
type InventoriesInventoriesAdHocCommandsListForbidden struct {
}

func (o *InventoriesInventoriesAdHocCommandsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/ad_hoc_commands/][%d] inventoriesInventoriesAdHocCommandsListForbidden ", 403)
}

func (o *InventoriesInventoriesAdHocCommandsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
