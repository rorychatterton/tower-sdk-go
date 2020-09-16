// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplatesWorkflowJobTemplatesActivityStreamListReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesActivityStreamList structure.
type WorkflowJobTemplatesWorkflowJobTemplatesActivityStreamListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesActivityStreamListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesActivityStreamListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesActivityStreamListOK creates a WorkflowJobTemplatesWorkflowJobTemplatesActivityStreamListOK with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesActivityStreamListOK() *WorkflowJobTemplatesWorkflowJobTemplatesActivityStreamListOK {
	return &WorkflowJobTemplatesWorkflowJobTemplatesActivityStreamListOK{}
}

/*WorkflowJobTemplatesWorkflowJobTemplatesActivityStreamListOK handles this case with default header values.

OK
*/
type WorkflowJobTemplatesWorkflowJobTemplatesActivityStreamListOK struct {
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesActivityStreamListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_templates/{id}/activity_stream/][%d] workflowJobTemplatesWorkflowJobTemplatesActivityStreamListOK ", 200)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesActivityStreamListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
