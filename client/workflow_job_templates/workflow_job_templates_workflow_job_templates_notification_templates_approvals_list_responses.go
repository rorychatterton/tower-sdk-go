// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsListReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsList structure.
type WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsListOK creates a WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsListOK with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsListOK() *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsListOK {
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsListOK{}
}

/*WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsListOK handles this case with default header values.

OK
*/
type WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsListOK struct {
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_templates/{id}/notification_templates_approvals/][%d] workflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsListOK ", 200)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
