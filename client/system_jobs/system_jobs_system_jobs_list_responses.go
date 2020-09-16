// Code generated by go-swagger; DO NOT EDIT.

package system_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SystemJobsSystemJobsListReader is a Reader for the SystemJobsSystemJobsList structure.
type SystemJobsSystemJobsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemJobsSystemJobsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemJobsSystemJobsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemJobsSystemJobsListOK creates a SystemJobsSystemJobsListOK with default headers values
func NewSystemJobsSystemJobsListOK() *SystemJobsSystemJobsListOK {
	return &SystemJobsSystemJobsListOK{}
}

/*SystemJobsSystemJobsListOK handles this case with default header values.

OK
*/
type SystemJobsSystemJobsListOK struct {
}

func (o *SystemJobsSystemJobsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/system_jobs/][%d] systemJobsSystemJobsListOK ", 200)
}

func (o *SystemJobsSystemJobsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}