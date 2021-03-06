// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

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

// WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreate structure.
type WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateCreated creates a WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateCreated with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateCreated() *WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateCreated {
	return &WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateCreated{}
}

/*WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateCreated handles this case with default header values.

WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateCreated workflow job templates workflow job templates labels create created
*/
type WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateCreated struct {
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/labels/][%d] workflowJobTemplatesWorkflowJobTemplatesLabelsCreateCreated ", 201)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateBody workflow job templates workflow job templates labels create body
swagger:model WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateBody
*/
type WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateBody struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// Organization this label belongs to.
	// Required: true
	Organization *int64 `json:"organization"`
}

// Validate validates this workflow job templates workflow job templates labels create body
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateBody) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organization", "body", o.Organization); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateBody) UnmarshalBinary(b []byte) error {
	var res WorkflowJobTemplatesWorkflowJobTemplatesLabelsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
