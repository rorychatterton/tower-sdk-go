// Code generated by go-swagger; DO NOT EDIT.

package inventory_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventorySourcesInventorySourcesCredentialsListReader is a Reader for the InventorySourcesInventorySourcesCredentialsList structure.
type InventorySourcesInventorySourcesCredentialsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventorySourcesInventorySourcesCredentialsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventorySourcesInventorySourcesCredentialsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventorySourcesInventorySourcesCredentialsListOK creates a InventorySourcesInventorySourcesCredentialsListOK with default headers values
func NewInventorySourcesInventorySourcesCredentialsListOK() *InventorySourcesInventorySourcesCredentialsListOK {
	return &InventorySourcesInventorySourcesCredentialsListOK{}
}

/*InventorySourcesInventorySourcesCredentialsListOK handles this case with default header values.

OK
*/
type InventorySourcesInventorySourcesCredentialsListOK struct {
}

func (o *InventorySourcesInventorySourcesCredentialsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventory_sources/{id}/credentials/][%d] inventorySourcesInventorySourcesCredentialsListOK ", 200)
}

func (o *InventorySourcesInventorySourcesCredentialsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
