// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobsJobsRelaunchCreateReader is a Reader for the JobsJobsRelaunchCreate structure.
type JobsJobsRelaunchCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobsJobsRelaunchCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewJobsJobsRelaunchCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewJobsJobsRelaunchCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobsJobsRelaunchCreateCreated creates a JobsJobsRelaunchCreateCreated with default headers values
func NewJobsJobsRelaunchCreateCreated() *JobsJobsRelaunchCreateCreated {
	return &JobsJobsRelaunchCreateCreated{}
}

/*JobsJobsRelaunchCreateCreated handles this case with default header values.

JobsJobsRelaunchCreateCreated jobs jobs relaunch create created
*/
type JobsJobsRelaunchCreateCreated struct {
}

func (o *JobsJobsRelaunchCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/jobs/{id}/relaunch/][%d] jobsJobsRelaunchCreateCreated ", 201)
}

func (o *JobsJobsRelaunchCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewJobsJobsRelaunchCreateForbidden creates a JobsJobsRelaunchCreateForbidden with default headers values
func NewJobsJobsRelaunchCreateForbidden() *JobsJobsRelaunchCreateForbidden {
	return &JobsJobsRelaunchCreateForbidden{}
}

/*JobsJobsRelaunchCreateForbidden handles this case with default header values.

No Permission Response
*/
type JobsJobsRelaunchCreateForbidden struct {
}

func (o *JobsJobsRelaunchCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/jobs/{id}/relaunch/][%d] jobsJobsRelaunchCreateForbidden ", 403)
}

func (o *JobsJobsRelaunchCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
