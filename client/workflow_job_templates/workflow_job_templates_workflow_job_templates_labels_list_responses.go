// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplatesWorkflowJobTemplatesLabelsListReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesLabelsList structure.
type WorkflowJobTemplatesWorkflowJobTemplatesLabelsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK creates a WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK() *WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK {
	return &WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK{}
}

/*WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK handles this case with default header values.

OK
*/
type WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK struct {
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_templates/{id}/labels/][%d] workflowJobTemplatesWorkflowJobTemplatesLabelsListOK ", 200)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
