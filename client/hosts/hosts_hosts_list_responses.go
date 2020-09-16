// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HostsHostsListReader is a Reader for the HostsHostsList structure.
type HostsHostsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsHostsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHostsHostsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHostsHostsListOK creates a HostsHostsListOK with default headers values
func NewHostsHostsListOK() *HostsHostsListOK {
	return &HostsHostsListOK{}
}

/*HostsHostsListOK handles this case with default header values.

OK
*/
type HostsHostsListOK struct {
}

func (o *HostsHostsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/hosts/][%d] hostsHostsListOK ", 200)
}

func (o *HostsHostsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
