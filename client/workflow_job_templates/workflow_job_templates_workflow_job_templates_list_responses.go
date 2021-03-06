// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplatesWorkflowJobTemplatesListReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesList structure.
type WorkflowJobTemplatesWorkflowJobTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesListOK creates a WorkflowJobTemplatesWorkflowJobTemplatesListOK with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesListOK() *WorkflowJobTemplatesWorkflowJobTemplatesListOK {
	return &WorkflowJobTemplatesWorkflowJobTemplatesListOK{}
}

/*WorkflowJobTemplatesWorkflowJobTemplatesListOK handles this case with default header values.

OK
*/
type WorkflowJobTemplatesWorkflowJobTemplatesListOK struct {
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_templates/][%d] workflowJobTemplatesWorkflowJobTemplatesListOK ", 200)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
