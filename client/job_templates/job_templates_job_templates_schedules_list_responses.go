// Code generated by go-swagger; DO NOT EDIT.

package job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobTemplatesJobTemplatesSchedulesListReader is a Reader for the JobTemplatesJobTemplatesSchedulesList structure.
type JobTemplatesJobTemplatesSchedulesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobTemplatesJobTemplatesSchedulesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobTemplatesJobTemplatesSchedulesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobTemplatesJobTemplatesSchedulesListOK creates a JobTemplatesJobTemplatesSchedulesListOK with default headers values
func NewJobTemplatesJobTemplatesSchedulesListOK() *JobTemplatesJobTemplatesSchedulesListOK {
	return &JobTemplatesJobTemplatesSchedulesListOK{}
}

/*JobTemplatesJobTemplatesSchedulesListOK handles this case with default header values.

OK
*/
type JobTemplatesJobTemplatesSchedulesListOK struct {
}

func (o *JobTemplatesJobTemplatesSchedulesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/job_templates/{id}/schedules/][%d] jobTemplatesJobTemplatesSchedulesListOK ", 200)
}

func (o *JobTemplatesJobTemplatesSchedulesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
