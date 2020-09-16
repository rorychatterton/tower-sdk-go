// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplatesWorkflowJobTemplatesLaunchReadReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesLaunchRead structure.
type WorkflowJobTemplatesWorkflowJobTemplatesLaunchReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesLaunchReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesLaunchReadOK creates a WorkflowJobTemplatesWorkflowJobTemplatesLaunchReadOK with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesLaunchReadOK() *WorkflowJobTemplatesWorkflowJobTemplatesLaunchReadOK {
	return &WorkflowJobTemplatesWorkflowJobTemplatesLaunchReadOK{}
}

/*WorkflowJobTemplatesWorkflowJobTemplatesLaunchReadOK handles this case with default header values.

OK
*/
type WorkflowJobTemplatesWorkflowJobTemplatesLaunchReadOK struct {
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_templates/{id}/launch/][%d] workflowJobTemplatesWorkflowJobTemplatesLaunchReadOK ", 200)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}