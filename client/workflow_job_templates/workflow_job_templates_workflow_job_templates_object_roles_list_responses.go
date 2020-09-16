// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplatesWorkflowJobTemplatesObjectRolesListReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesObjectRolesList structure.
type WorkflowJobTemplatesWorkflowJobTemplatesObjectRolesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesObjectRolesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesObjectRolesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesObjectRolesListOK creates a WorkflowJobTemplatesWorkflowJobTemplatesObjectRolesListOK with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesObjectRolesListOK() *WorkflowJobTemplatesWorkflowJobTemplatesObjectRolesListOK {
	return &WorkflowJobTemplatesWorkflowJobTemplatesObjectRolesListOK{}
}

/*WorkflowJobTemplatesWorkflowJobTemplatesObjectRolesListOK handles this case with default header values.

OK
*/
type WorkflowJobTemplatesWorkflowJobTemplatesObjectRolesListOK struct {
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesObjectRolesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_templates/{id}/object_roles/][%d] workflowJobTemplatesWorkflowJobTemplatesObjectRolesListOK ", 200)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesObjectRolesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}