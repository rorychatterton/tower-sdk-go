// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HostsHostsVariableDataReadReader is a Reader for the HostsHostsVariableDataRead structure.
type HostsHostsVariableDataReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsHostsVariableDataReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHostsHostsVariableDataReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHostsHostsVariableDataReadOK creates a HostsHostsVariableDataReadOK with default headers values
func NewHostsHostsVariableDataReadOK() *HostsHostsVariableDataReadOK {
	return &HostsHostsVariableDataReadOK{}
}

/*HostsHostsVariableDataReadOK handles this case with default header values.

OK
*/
type HostsHostsVariableDataReadOK struct {
}

func (o *HostsHostsVariableDataReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/hosts/{id}/variable_data/][%d] hostsHostsVariableDataReadOK ", 200)
}

func (o *HostsHostsVariableDataReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
