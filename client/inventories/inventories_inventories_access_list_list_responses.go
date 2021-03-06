// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoriesInventoriesAccessListListReader is a Reader for the InventoriesInventoriesAccessListList structure.
type InventoriesInventoriesAccessListListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesAccessListListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoriesInventoriesAccessListListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventoriesInventoriesAccessListListOK creates a InventoriesInventoriesAccessListListOK with default headers values
func NewInventoriesInventoriesAccessListListOK() *InventoriesInventoriesAccessListListOK {
	return &InventoriesInventoriesAccessListListOK{}
}

/*InventoriesInventoriesAccessListListOK handles this case with default header values.

OK
*/
type InventoriesInventoriesAccessListListOK struct {
}

func (o *InventoriesInventoriesAccessListListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/access_list/][%d] inventoriesInventoriesAccessListListOK ", 200)
}

func (o *InventoriesInventoriesAccessListListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
