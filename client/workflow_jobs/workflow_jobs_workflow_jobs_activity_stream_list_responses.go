// Code generated by go-swagger; DO NOT EDIT.

package workflow_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobsWorkflowJobsActivityStreamListReader is a Reader for the WorkflowJobsWorkflowJobsActivityStreamList structure.
type WorkflowJobsWorkflowJobsActivityStreamListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobsWorkflowJobsActivityStreamListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobsWorkflowJobsActivityStreamListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowJobsWorkflowJobsActivityStreamListOK creates a WorkflowJobsWorkflowJobsActivityStreamListOK with default headers values
func NewWorkflowJobsWorkflowJobsActivityStreamListOK() *WorkflowJobsWorkflowJobsActivityStreamListOK {
	return &WorkflowJobsWorkflowJobsActivityStreamListOK{}
}

/*WorkflowJobsWorkflowJobsActivityStreamListOK handles this case with default header values.

OK
*/
type WorkflowJobsWorkflowJobsActivityStreamListOK struct {
}

func (o *WorkflowJobsWorkflowJobsActivityStreamListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_jobs/{id}/activity_stream/][%d] workflowJobsWorkflowJobsActivityStreamListOK ", 200)
}

func (o *WorkflowJobsWorkflowJobsActivityStreamListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
