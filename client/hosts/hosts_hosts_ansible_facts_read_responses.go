// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HostsHostsAnsibleFactsReadReader is a Reader for the HostsHostsAnsibleFactsRead structure.
type HostsHostsAnsibleFactsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsHostsAnsibleFactsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHostsHostsAnsibleFactsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHostsHostsAnsibleFactsReadOK creates a HostsHostsAnsibleFactsReadOK with default headers values
func NewHostsHostsAnsibleFactsReadOK() *HostsHostsAnsibleFactsReadOK {
	return &HostsHostsAnsibleFactsReadOK{}
}

/*HostsHostsAnsibleFactsReadOK handles this case with default header values.

OK
*/
type HostsHostsAnsibleFactsReadOK struct {
}

func (o *HostsHostsAnsibleFactsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/hosts/{id}/ansible_facts/][%d] hostsHostsAnsibleFactsReadOK ", 200)
}

func (o *HostsHostsAnsibleFactsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
