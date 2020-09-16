// Code generated by go-swagger; DO NOT EDIT.

package system_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SystemJobTemplatesSystemJobTemplatesLaunchListReader is a Reader for the SystemJobTemplatesSystemJobTemplatesLaunchList structure.
type SystemJobTemplatesSystemJobTemplatesLaunchListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemJobTemplatesSystemJobTemplatesLaunchListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemJobTemplatesSystemJobTemplatesLaunchListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemJobTemplatesSystemJobTemplatesLaunchListOK creates a SystemJobTemplatesSystemJobTemplatesLaunchListOK with default headers values
func NewSystemJobTemplatesSystemJobTemplatesLaunchListOK() *SystemJobTemplatesSystemJobTemplatesLaunchListOK {
	return &SystemJobTemplatesSystemJobTemplatesLaunchListOK{}
}

/*SystemJobTemplatesSystemJobTemplatesLaunchListOK handles this case with default header values.

OK
*/
type SystemJobTemplatesSystemJobTemplatesLaunchListOK struct {
}

func (o *SystemJobTemplatesSystemJobTemplatesLaunchListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/system_job_templates/{id}/launch/][%d] systemJobTemplatesSystemJobTemplatesLaunchListOK ", 200)
}

func (o *SystemJobTemplatesSystemJobTemplatesLaunchListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
