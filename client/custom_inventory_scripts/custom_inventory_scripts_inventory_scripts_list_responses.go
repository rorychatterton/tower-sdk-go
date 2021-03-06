// Code generated by go-swagger; DO NOT EDIT.

package custom_inventory_scripts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CustomInventoryScriptsInventoryScriptsListReader is a Reader for the CustomInventoryScriptsInventoryScriptsList structure.
type CustomInventoryScriptsInventoryScriptsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomInventoryScriptsInventoryScriptsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomInventoryScriptsInventoryScriptsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCustomInventoryScriptsInventoryScriptsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomInventoryScriptsInventoryScriptsListOK creates a CustomInventoryScriptsInventoryScriptsListOK with default headers values
func NewCustomInventoryScriptsInventoryScriptsListOK() *CustomInventoryScriptsInventoryScriptsListOK {
	return &CustomInventoryScriptsInventoryScriptsListOK{}
}

/*CustomInventoryScriptsInventoryScriptsListOK handles this case with default header values.

OK
*/
type CustomInventoryScriptsInventoryScriptsListOK struct {
}

func (o *CustomInventoryScriptsInventoryScriptsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventory_scripts/][%d] customInventoryScriptsInventoryScriptsListOK ", 200)
}

func (o *CustomInventoryScriptsInventoryScriptsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomInventoryScriptsInventoryScriptsListForbidden creates a CustomInventoryScriptsInventoryScriptsListForbidden with default headers values
func NewCustomInventoryScriptsInventoryScriptsListForbidden() *CustomInventoryScriptsInventoryScriptsListForbidden {
	return &CustomInventoryScriptsInventoryScriptsListForbidden{}
}

/*CustomInventoryScriptsInventoryScriptsListForbidden handles this case with default header values.

No Permission Response
*/
type CustomInventoryScriptsInventoryScriptsListForbidden struct {
}

func (o *CustomInventoryScriptsInventoryScriptsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventory_scripts/][%d] customInventoryScriptsInventoryScriptsListForbidden ", 403)
}

func (o *CustomInventoryScriptsInventoryScriptsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
