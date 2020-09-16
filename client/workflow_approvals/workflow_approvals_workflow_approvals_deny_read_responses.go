// Code generated by go-swagger; DO NOT EDIT.

package workflow_approvals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowApprovalsWorkflowApprovalsDenyReadReader is a Reader for the WorkflowApprovalsWorkflowApprovalsDenyRead structure.
type WorkflowApprovalsWorkflowApprovalsDenyReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowApprovalsWorkflowApprovalsDenyReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowApprovalsWorkflowApprovalsDenyReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowApprovalsWorkflowApprovalsDenyReadOK creates a WorkflowApprovalsWorkflowApprovalsDenyReadOK with default headers values
func NewWorkflowApprovalsWorkflowApprovalsDenyReadOK() *WorkflowApprovalsWorkflowApprovalsDenyReadOK {
	return &WorkflowApprovalsWorkflowApprovalsDenyReadOK{}
}

/*WorkflowApprovalsWorkflowApprovalsDenyReadOK handles this case with default header values.

OK
*/
type WorkflowApprovalsWorkflowApprovalsDenyReadOK struct {
}

func (o *WorkflowApprovalsWorkflowApprovalsDenyReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_approvals/{id}/deny/][%d] workflowApprovalsWorkflowApprovalsDenyReadOK ", 200)
}

func (o *WorkflowApprovalsWorkflowApprovalsDenyReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
