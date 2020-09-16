// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InstancesInstancesListReader is a Reader for the InstancesInstancesList structure.
type InstancesInstancesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstancesInstancesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstancesInstancesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInstancesInstancesListOK creates a InstancesInstancesListOK with default headers values
func NewInstancesInstancesListOK() *InstancesInstancesListOK {
	return &InstancesInstancesListOK{}
}

/*InstancesInstancesListOK handles this case with default header values.

OK
*/
type InstancesInstancesListOK struct {
}

func (o *InstancesInstancesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/instances/][%d] instancesInstancesListOK ", 200)
}

func (o *InstancesInstancesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}