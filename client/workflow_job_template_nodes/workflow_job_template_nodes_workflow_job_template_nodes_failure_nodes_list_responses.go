// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_template_nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListReader is a Reader for the WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesList structure.
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListOK creates a WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListOK with default headers values
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListOK() *WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListOK {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListOK{}
}

/*WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListOK handles this case with default header values.

OK
*/
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListOK struct {
}

func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_template_nodes/{id}/failure_nodes/][%d] workflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListOK ", 200)
}

func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}