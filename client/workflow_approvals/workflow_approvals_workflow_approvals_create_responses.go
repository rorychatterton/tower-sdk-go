// Code generated by go-swagger; DO NOT EDIT.

package workflow_approvals

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

// WorkflowApprovalsWorkflowApprovalsCreateReader is a Reader for the WorkflowApprovalsWorkflowApprovalsCreate structure.
type WorkflowApprovalsWorkflowApprovalsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowApprovalsWorkflowApprovalsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWorkflowApprovalsWorkflowApprovalsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWorkflowApprovalsWorkflowApprovalsCreateCreated creates a WorkflowApprovalsWorkflowApprovalsCreateCreated with default headers values
func NewWorkflowApprovalsWorkflowApprovalsCreateCreated() *WorkflowApprovalsWorkflowApprovalsCreateCreated {
	return &WorkflowApprovalsWorkflowApprovalsCreateCreated{}
}

/*WorkflowApprovalsWorkflowApprovalsCreateCreated handles this case with default header values.

WorkflowApprovalsWorkflowApprovalsCreateCreated workflow approvals workflow approvals create created
*/
type WorkflowApprovalsWorkflowApprovalsCreateCreated struct {
}

func (o *WorkflowApprovalsWorkflowApprovalsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workflow_approvals/][%d] workflowApprovalsWorkflowApprovalsCreateCreated ", 201)
}

func (o *WorkflowApprovalsWorkflowApprovalsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*WorkflowApprovalsWorkflowApprovalsCreateBody workflow approvals workflow approvals create body
swagger:model WorkflowApprovalsWorkflowApprovalsCreateBody
*/
type WorkflowApprovalsWorkflowApprovalsCreateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this workflow approvals workflow approvals create body
func (o *WorkflowApprovalsWorkflowApprovalsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WorkflowApprovalsWorkflowApprovalsCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *WorkflowApprovalsWorkflowApprovalsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WorkflowApprovalsWorkflowApprovalsCreateBody) UnmarshalBinary(b []byte) error {
	var res WorkflowApprovalsWorkflowApprovalsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}