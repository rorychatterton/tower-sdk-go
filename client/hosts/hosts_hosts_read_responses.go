// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HostsHostsReadReader is a Reader for the HostsHostsRead structure.
type HostsHostsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsHostsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHostsHostsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHostsHostsReadOK creates a HostsHostsReadOK with default headers values
func NewHostsHostsReadOK() *HostsHostsReadOK {
	return &HostsHostsReadOK{}
}

/*HostsHostsReadOK handles this case with default header values.

OK
*/
type HostsHostsReadOK struct {
}

func (o *HostsHostsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/hosts/{id}/][%d] hostsHostsReadOK ", 200)
}

func (o *HostsHostsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
