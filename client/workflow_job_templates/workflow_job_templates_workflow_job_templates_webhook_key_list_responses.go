// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyList structure.
type WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListOK creates a WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListOK with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListOK() *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListOK {
	return &WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListOK{}
}

/*WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListOK handles this case with default header values.

OK
*/
type WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListOK struct {
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_templates/{id}/webhook_key/][%d] workflowJobTemplatesWorkflowJobTemplatesWebhookKeyListOK ", 200)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListForbidden creates a WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListForbidden with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListForbidden() *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListForbidden {
	return &WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListForbidden{}
}

/*WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListForbidden handles this case with default header values.

No Permission Response
*/
type WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListForbidden struct {
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_templates/{id}/webhook_key/][%d] workflowJobTemplatesWorkflowJobTemplatesWebhookKeyListForbidden ", 403)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}