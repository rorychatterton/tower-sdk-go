// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesGithubCreate structure.
type WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated creates a WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated() *WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated {
	return &WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated{}
}

/*WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated handles this case with default header values.

WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated workflow job templates workflow job templates github create created
*/
type WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated struct {
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/github/][%d] workflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated ", 201)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
