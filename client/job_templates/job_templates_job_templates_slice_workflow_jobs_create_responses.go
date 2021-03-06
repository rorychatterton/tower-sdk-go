// Code generated by go-swagger; DO NOT EDIT.

package job_templates

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

// JobTemplatesJobTemplatesSliceWorkflowJobsCreateReader is a Reader for the JobTemplatesJobTemplatesSliceWorkflowJobsCreate structure.
type JobTemplatesJobTemplatesSliceWorkflowJobsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewJobTemplatesJobTemplatesSliceWorkflowJobsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobTemplatesJobTemplatesSliceWorkflowJobsCreateCreated creates a JobTemplatesJobTemplatesSliceWorkflowJobsCreateCreated with default headers values
func NewJobTemplatesJobTemplatesSliceWorkflowJobsCreateCreated() *JobTemplatesJobTemplatesSliceWorkflowJobsCreateCreated {
	return &JobTemplatesJobTemplatesSliceWorkflowJobsCreateCreated{}
}

/*JobTemplatesJobTemplatesSliceWorkflowJobsCreateCreated handles this case with default header values.

JobTemplatesJobTemplatesSliceWorkflowJobsCreateCreated job templates job templates slice workflow jobs create created
*/
type JobTemplatesJobTemplatesSliceWorkflowJobsCreateCreated struct {
}

func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/job_templates/{id}/slice_workflow_jobs/][%d] jobTemplatesJobTemplatesSliceWorkflowJobsCreateCreated ", 201)
}

func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*JobTemplatesJobTemplatesSliceWorkflowJobsCreateBody job templates job templates slice workflow jobs create body
swagger:model JobTemplatesJobTemplatesSliceWorkflowJobsCreateBody
*/
type JobTemplatesJobTemplatesSliceWorkflowJobsCreateBody struct {

	// allow simultaneous
	AllowSimultaneous string `json:"allow_simultaneous,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// extra vars
	ExtraVars string `json:"extra_vars,omitempty"`

	// Inventory applied as a prompt, assuming job template prompts for inventory
	Inventory int64 `json:"inventory,omitempty"`

	// is sliced job
	IsSlicedJob string `json:"is_sliced_job,omitempty"`

	// If automatically created for a sliced job run, the job template the workflow job was created from.
	JobTemplate string `json:"job_template,omitempty"`

	// limit
	Limit string `json:"limit,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// scm branch
	ScmBranch string `json:"scm_branch,omitempty"`

	// Personal Access Token for posting back the status to the service API
	WebhookCredential int64 `json:"webhook_credential,omitempty"`

	// Unique identifier of the event that triggered this webhook
	WebhookGUID string `json:"webhook_guid,omitempty"`

	// Service that webhook requests will be accepted from
	WebhookService string `json:"webhook_service,omitempty"`

	// workflow job template
	WorkflowJobTemplate string `json:"workflow_job_template,omitempty"`
}

// Validate validates this job templates job templates slice workflow jobs create body
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *JobTemplatesJobTemplatesSliceWorkflowJobsCreateBody) UnmarshalBinary(b []byte) error {
	var res JobTemplatesJobTemplatesSliceWorkflowJobsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
