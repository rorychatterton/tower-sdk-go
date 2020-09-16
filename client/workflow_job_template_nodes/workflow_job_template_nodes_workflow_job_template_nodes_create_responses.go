// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_template_nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateReader is a Reader for the WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreate structure.
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateCreated creates a WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateCreated with default headers values
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateCreated() *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateCreated {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateCreated{}
}

/*WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateCreated handles this case with default header values.

WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateCreated workflow job template nodes workflow job template nodes create created
*/
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateCreated struct {
}

func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_template_nodes/][%d] workflowJobTemplateNodesWorkflowJobTemplateNodesCreateCreated ", 201)
}

func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateBody workflow job template nodes workflow job template nodes create body
swagger:model WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateBody
*/
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateBody struct {

	// If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node
	AllParentsMustConverge string `json:"all_parents_must_converge,omitempty"`

	// diff mode
	DiffMode string `json:"diff_mode,omitempty"`

	// extra data
	ExtraData string `json:"extra_data,omitempty"`

	// An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node.
	Identifier string `json:"identifier,omitempty"`

	// Inventory applied as a prompt, assuming job template prompts for inventory
	Inventory int64 `json:"inventory,omitempty"`

	// job tags
	JobTags string `json:"job_tags,omitempty"`

	// job type
	JobType string `json:"job_type,omitempty"`

	// limit
	Limit string `json:"limit,omitempty"`

	// scm branch
	ScmBranch string `json:"scm_branch,omitempty"`

	// skip tags
	SkipTags string `json:"skip_tags,omitempty"`

	// unified job template
	UnifiedJobTemplate int64 `json:"unified_job_template,omitempty"`

	// verbosity
	Verbosity string `json:"verbosity,omitempty"`

	// workflow job template
	// Required: true
	WorkflowJobTemplate *string `json:"workflow_job_template"`
}

// Validate validates this workflow job template nodes workflow job template nodes create body
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateWorkflowJobTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateBody) validateWorkflowJobTemplate(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"workflow_job_template", "body", o.WorkflowJobTemplate); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateBody) UnmarshalBinary(b []byte) error {
	var res WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}