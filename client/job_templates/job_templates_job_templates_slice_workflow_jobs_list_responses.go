// Code generated by go-swagger; DO NOT EDIT.

package job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobTemplatesJobTemplatesSliceWorkflowJobsListReader is a Reader for the JobTemplatesJobTemplatesSliceWorkflowJobsList structure.
type JobTemplatesJobTemplatesSliceWorkflowJobsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobTemplatesJobTemplatesSliceWorkflowJobsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobTemplatesJobTemplatesSliceWorkflowJobsListOK creates a JobTemplatesJobTemplatesSliceWorkflowJobsListOK with default headers values
func NewJobTemplatesJobTemplatesSliceWorkflowJobsListOK() *JobTemplatesJobTemplatesSliceWorkflowJobsListOK {
	return &JobTemplatesJobTemplatesSliceWorkflowJobsListOK{}
}

/*JobTemplatesJobTemplatesSliceWorkflowJobsListOK handles this case with default header values.

OK
*/
type JobTemplatesJobTemplatesSliceWorkflowJobsListOK struct {
}

func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/job_templates/{id}/slice_workflow_jobs/][%d] jobTemplatesJobTemplatesSliceWorkflowJobsListOK ", 200)
}

func (o *JobTemplatesJobTemplatesSliceWorkflowJobsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
