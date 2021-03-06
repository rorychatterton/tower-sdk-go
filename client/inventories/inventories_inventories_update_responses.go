// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoriesInventoriesUpdateReader is a Reader for the InventoriesInventoriesUpdate structure.
type InventoriesInventoriesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoriesInventoriesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewInventoriesInventoriesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventoriesInventoriesUpdateOK creates a InventoriesInventoriesUpdateOK with default headers values
func NewInventoriesInventoriesUpdateOK() *InventoriesInventoriesUpdateOK {
	return &InventoriesInventoriesUpdateOK{}
}

/*InventoriesInventoriesUpdateOK handles this case with default header values.

OK
*/
type InventoriesInventoriesUpdateOK struct {
}

func (o *InventoriesInventoriesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/inventories/{id}/][%d] inventoriesInventoriesUpdateOK ", 200)
}

func (o *InventoriesInventoriesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInventoriesInventoriesUpdateForbidden creates a InventoriesInventoriesUpdateForbidden with default headers values
func NewInventoriesInventoriesUpdateForbidden() *InventoriesInventoriesUpdateForbidden {
	return &InventoriesInventoriesUpdateForbidden{}
}

/*InventoriesInventoriesUpdateForbidden handles this case with default header values.

No Permission Response
*/
type InventoriesInventoriesUpdateForbidden struct {
}

func (o *InventoriesInventoriesUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/inventories/{id}/][%d] inventoriesInventoriesUpdateForbidden ", 403)
}

func (o *InventoriesInventoriesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
