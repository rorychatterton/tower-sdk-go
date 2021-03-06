// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobsJobsJobEventsListReader is a Reader for the JobsJobsJobEventsList structure.
type JobsJobsJobEventsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobsJobsJobEventsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobsJobsJobEventsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobsJobsJobEventsListOK creates a JobsJobsJobEventsListOK with default headers values
func NewJobsJobsJobEventsListOK() *JobsJobsJobEventsListOK {
	return &JobsJobsJobEventsListOK{}
}

/*JobsJobsJobEventsListOK handles this case with default header values.

OK
*/
type JobsJobsJobEventsListOK struct {
}

func (o *JobsJobsJobEventsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/jobs/{id}/job_events/][%d] jobsJobsJobEventsListOK ", 200)
}

func (o *JobsJobsJobEventsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
