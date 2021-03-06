// Code generated by go-swagger; DO NOT EDIT.

package system_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SystemJobsSystemJobsReadReader is a Reader for the SystemJobsSystemJobsRead structure.
type SystemJobsSystemJobsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemJobsSystemJobsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemJobsSystemJobsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemJobsSystemJobsReadOK creates a SystemJobsSystemJobsReadOK with default headers values
func NewSystemJobsSystemJobsReadOK() *SystemJobsSystemJobsReadOK {
	return &SystemJobsSystemJobsReadOK{}
}

/*SystemJobsSystemJobsReadOK handles this case with default header values.

OK
*/
type SystemJobsSystemJobsReadOK struct {
}

func (o *SystemJobsSystemJobsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/system_jobs/{id}/][%d] systemJobsSystemJobsReadOK ", 200)
}

func (o *SystemJobsSystemJobsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
