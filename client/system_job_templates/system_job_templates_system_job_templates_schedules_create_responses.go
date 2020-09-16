// Code generated by go-swagger; DO NOT EDIT.

package system_job_templates

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

// SystemJobTemplatesSystemJobTemplatesSchedulesCreateReader is a Reader for the SystemJobTemplatesSystemJobTemplatesSchedulesCreate structure.
type SystemJobTemplatesSystemJobTemplatesSchedulesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSystemJobTemplatesSystemJobTemplatesSchedulesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemJobTemplatesSystemJobTemplatesSchedulesCreateCreated creates a SystemJobTemplatesSystemJobTemplatesSchedulesCreateCreated with default headers values
func NewSystemJobTemplatesSystemJobTemplatesSchedulesCreateCreated() *SystemJobTemplatesSystemJobTemplatesSchedulesCreateCreated {
	return &SystemJobTemplatesSystemJobTemplatesSchedulesCreateCreated{}
}

/*SystemJobTemplatesSystemJobTemplatesSchedulesCreateCreated handles this case with default header values.

SystemJobTemplatesSystemJobTemplatesSchedulesCreateCreated system job templates system job templates schedules create created
*/
type SystemJobTemplatesSystemJobTemplatesSchedulesCreateCreated struct {
}

func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/system_job_templates/{id}/schedules/][%d] systemJobTemplatesSystemJobTemplatesSchedulesCreateCreated ", 201)
}

func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*SystemJobTemplatesSystemJobTemplatesSchedulesCreateBody system job templates system job templates schedules create body
swagger:model SystemJobTemplatesSystemJobTemplatesSchedulesCreateBody
*/
type SystemJobTemplatesSystemJobTemplatesSchedulesCreateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// diff mode
	DiffMode string `json:"diff_mode,omitempty"`

	// Enables processing of this schedule.
	Enabled string `json:"enabled,omitempty"`

	// extra data
	ExtraData string `json:"extra_data,omitempty"`

	// Inventory applied as a prompt, assuming job template prompts for inventory
	Inventory int64 `json:"inventory,omitempty"`

	// job tags
	JobTags string `json:"job_tags,omitempty"`

	// job type
	JobType string `json:"job_type,omitempty"`

	// limit
	Limit string `json:"limit,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// A value representing the schedules iCal recurrence rule.
	// Required: true
	Rrule *string `json:"rrule"`

	// scm branch
	ScmBranch string `json:"scm_branch,omitempty"`

	// skip tags
	SkipTags string `json:"skip_tags,omitempty"`

	// unified job template
	// Required: true
	UnifiedJobTemplate *int64 `json:"unified_job_template"`

	// verbosity
	Verbosity string `json:"verbosity,omitempty"`
}

// Validate validates this system job templates system job templates schedules create body
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRrule(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUnifiedJobTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateBody) validateRrule(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"rrule", "body", o.Rrule); err != nil {
		return err
	}

	return nil
}

func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateBody) validateUnifiedJobTemplate(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"unified_job_template", "body", o.UnifiedJobTemplate); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateBody) UnmarshalBinary(b []byte) error {
	var res SystemJobTemplatesSystemJobTemplatesSchedulesCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
