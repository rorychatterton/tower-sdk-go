// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoriesInventoriesVariableDataPartialUpdateReader is a Reader for the InventoriesInventoriesVariableDataPartialUpdate structure.
type InventoriesInventoriesVariableDataPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesVariableDataPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoriesInventoriesVariableDataPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewInventoriesInventoriesVariableDataPartialUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventoriesInventoriesVariableDataPartialUpdateOK creates a InventoriesInventoriesVariableDataPartialUpdateOK with default headers values
func NewInventoriesInventoriesVariableDataPartialUpdateOK() *InventoriesInventoriesVariableDataPartialUpdateOK {
	return &InventoriesInventoriesVariableDataPartialUpdateOK{}
}

/*InventoriesInventoriesVariableDataPartialUpdateOK handles this case with default header values.

OK
*/
type InventoriesInventoriesVariableDataPartialUpdateOK struct {
}

func (o *InventoriesInventoriesVariableDataPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/inventories/{id}/variable_data/][%d] inventoriesInventoriesVariableDataPartialUpdateOK ", 200)
}

func (o *InventoriesInventoriesVariableDataPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInventoriesInventoriesVariableDataPartialUpdateForbidden creates a InventoriesInventoriesVariableDataPartialUpdateForbidden with default headers values
func NewInventoriesInventoriesVariableDataPartialUpdateForbidden() *InventoriesInventoriesVariableDataPartialUpdateForbidden {
	return &InventoriesInventoriesVariableDataPartialUpdateForbidden{}
}

/*InventoriesInventoriesVariableDataPartialUpdateForbidden handles this case with default header values.

No Permission Response
*/
type InventoriesInventoriesVariableDataPartialUpdateForbidden struct {
}

func (o *InventoriesInventoriesVariableDataPartialUpdateForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/inventories/{id}/variable_data/][%d] inventoriesInventoriesVariableDataPartialUpdateForbidden ", 403)
}

func (o *InventoriesInventoriesVariableDataPartialUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
