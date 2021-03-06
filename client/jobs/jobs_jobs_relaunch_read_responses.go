// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobsJobsRelaunchReadReader is a Reader for the JobsJobsRelaunchRead structure.
type JobsJobsRelaunchReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobsJobsRelaunchReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobsJobsRelaunchReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobsJobsRelaunchReadOK creates a JobsJobsRelaunchReadOK with default headers values
func NewJobsJobsRelaunchReadOK() *JobsJobsRelaunchReadOK {
	return &JobsJobsRelaunchReadOK{}
}

/*JobsJobsRelaunchReadOK handles this case with default header values.

OK
*/
type JobsJobsRelaunchReadOK struct {
}

func (o *JobsJobsRelaunchReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/jobs/{id}/relaunch/][%d] jobsJobsRelaunchReadOK ", 200)
}

func (o *JobsJobsRelaunchReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
