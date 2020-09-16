// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobsJobsListReader is a Reader for the JobsJobsList structure.
type JobsJobsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobsJobsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobsJobsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobsJobsListOK creates a JobsJobsListOK with default headers values
func NewJobsJobsListOK() *JobsJobsListOK {
	return &JobsJobsListOK{}
}

/*JobsJobsListOK handles this case with default header values.

OK
*/
type JobsJobsListOK struct {
}

func (o *JobsJobsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/jobs/][%d] jobsJobsListOK ", 200)
}

func (o *JobsJobsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
