// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobNodesWorkflowJobNodesListReader is a Reader for the WorkflowJobNodesWorkflowJobNodesList structure.
type WorkflowJobNodesWorkflowJobNodesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobNodesWorkflowJobNodesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobNodesWorkflowJobNodesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowJobNodesWorkflowJobNodesListOK creates a WorkflowJobNodesWorkflowJobNodesListOK with default headers values
func NewWorkflowJobNodesWorkflowJobNodesListOK() *WorkflowJobNodesWorkflowJobNodesListOK {
	return &WorkflowJobNodesWorkflowJobNodesListOK{}
}

/*WorkflowJobNodesWorkflowJobNodesListOK handles this case with default header values.

OK
*/
type WorkflowJobNodesWorkflowJobNodesListOK struct {
}

func (o *WorkflowJobNodesWorkflowJobNodesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_nodes/][%d] workflowJobNodesWorkflowJobNodesListOK ", 200)
}

func (o *WorkflowJobNodesWorkflowJobNodesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
