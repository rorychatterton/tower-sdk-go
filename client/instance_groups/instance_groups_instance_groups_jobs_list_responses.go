// Code generated by go-swagger; DO NOT EDIT.

package instance_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InstanceGroupsInstanceGroupsJobsListReader is a Reader for the InstanceGroupsInstanceGroupsJobsList structure.
type InstanceGroupsInstanceGroupsJobsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstanceGroupsInstanceGroupsJobsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstanceGroupsInstanceGroupsJobsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInstanceGroupsInstanceGroupsJobsListOK creates a InstanceGroupsInstanceGroupsJobsListOK with default headers values
func NewInstanceGroupsInstanceGroupsJobsListOK() *InstanceGroupsInstanceGroupsJobsListOK {
	return &InstanceGroupsInstanceGroupsJobsListOK{}
}

/*InstanceGroupsInstanceGroupsJobsListOK handles this case with default header values.

OK
*/
type InstanceGroupsInstanceGroupsJobsListOK struct {
}

func (o *InstanceGroupsInstanceGroupsJobsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/instance_groups/{id}/jobs/][%d] instanceGroupsInstanceGroupsJobsListOK ", 200)
}

func (o *InstanceGroupsInstanceGroupsJobsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
