// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoriesInventoriesVariableDataReadReader is a Reader for the InventoriesInventoriesVariableDataRead structure.
type InventoriesInventoriesVariableDataReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesVariableDataReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoriesInventoriesVariableDataReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventoriesInventoriesVariableDataReadOK creates a InventoriesInventoriesVariableDataReadOK with default headers values
func NewInventoriesInventoriesVariableDataReadOK() *InventoriesInventoriesVariableDataReadOK {
	return &InventoriesInventoriesVariableDataReadOK{}
}

/*InventoriesInventoriesVariableDataReadOK handles this case with default header values.

OK
*/
type InventoriesInventoriesVariableDataReadOK struct {
}

func (o *InventoriesInventoriesVariableDataReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/variable_data/][%d] inventoriesInventoriesVariableDataReadOK ", 200)
}

func (o *InventoriesInventoriesVariableDataReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
