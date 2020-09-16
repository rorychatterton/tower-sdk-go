// Code generated by go-swagger; DO NOT EDIT.

package job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobTemplatesJobTemplatesSchedulesCreateReader is a Reader for the JobTemplatesJobTemplatesSchedulesCreate structure.
type JobTemplatesJobTemplatesSchedulesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobTemplatesJobTemplatesSchedulesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewJobTemplatesJobTemplatesSchedulesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewJobTemplatesJobTemplatesSchedulesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobTemplatesJobTemplatesSchedulesCreateCreated creates a JobTemplatesJobTemplatesSchedulesCreateCreated with default headers values
func NewJobTemplatesJobTemplatesSchedulesCreateCreated() *JobTemplatesJobTemplatesSchedulesCreateCreated {
	return &JobTemplatesJobTemplatesSchedulesCreateCreated{}
}

/*JobTemplatesJobTemplatesSchedulesCreateCreated handles this case with default header values.

JobTemplatesJobTemplatesSchedulesCreateCreated job templates job templates schedules create created
*/
type JobTemplatesJobTemplatesSchedulesCreateCreated struct {
}

func (o *JobTemplatesJobTemplatesSchedulesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/job_templates/{id}/schedules/][%d] jobTemplatesJobTemplatesSchedulesCreateCreated ", 201)
}

func (o *JobTemplatesJobTemplatesSchedulesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewJobTemplatesJobTemplatesSchedulesCreateBadRequest creates a JobTemplatesJobTemplatesSchedulesCreateBadRequest with default headers values
func NewJobTemplatesJobTemplatesSchedulesCreateBadRequest() *JobTemplatesJobTemplatesSchedulesCreateBadRequest {
	return &JobTemplatesJobTemplatesSchedulesCreateBadRequest{}
}

/*JobTemplatesJobTemplatesSchedulesCreateBadRequest handles this case with default header values.

Bad Request
*/
type JobTemplatesJobTemplatesSchedulesCreateBadRequest struct {
}

func (o *JobTemplatesJobTemplatesSchedulesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/job_templates/{id}/schedules/][%d] jobTemplatesJobTemplatesSchedulesCreateBadRequest ", 400)
}

func (o *JobTemplatesJobTemplatesSchedulesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}