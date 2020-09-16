// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessListReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessList structure.
type WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessListOK creates a WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessListOK with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessListOK() *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessListOK {
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessListOK{}
}

/*WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessListOK handles this case with default header values.

OK
*/
type WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessListOK struct {
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_templates/{id}/notification_templates_success/][%d] workflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessListOK ", 200)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
