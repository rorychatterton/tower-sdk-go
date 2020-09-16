// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HostsHostsInventorySourcesListReader is a Reader for the HostsHostsInventorySourcesList structure.
type HostsHostsInventorySourcesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsHostsInventorySourcesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHostsHostsInventorySourcesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHostsHostsInventorySourcesListOK creates a HostsHostsInventorySourcesListOK with default headers values
func NewHostsHostsInventorySourcesListOK() *HostsHostsInventorySourcesListOK {
	return &HostsHostsInventorySourcesListOK{}
}

/*HostsHostsInventorySourcesListOK handles this case with default header values.

OK
*/
type HostsHostsInventorySourcesListOK struct {
}

func (o *HostsHostsInventorySourcesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/hosts/{id}/inventory_sources/][%d] hostsHostsInventorySourcesListOK ", 200)
}

func (o *HostsHostsInventorySourcesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}