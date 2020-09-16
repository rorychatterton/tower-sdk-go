// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// ProjectsProjectsNotificationTemplatesSuccessCreateReader is a Reader for the ProjectsProjectsNotificationTemplatesSuccessCreate structure.
type ProjectsProjectsNotificationTemplatesSuccessCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsProjectsNotificationTemplatesSuccessCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewProjectsProjectsNotificationTemplatesSuccessCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsProjectsNotificationTemplatesSuccessCreateCreated creates a ProjectsProjectsNotificationTemplatesSuccessCreateCreated with default headers values
func NewProjectsProjectsNotificationTemplatesSuccessCreateCreated() *ProjectsProjectsNotificationTemplatesSuccessCreateCreated {
	return &ProjectsProjectsNotificationTemplatesSuccessCreateCreated{}
}

/*ProjectsProjectsNotificationTemplatesSuccessCreateCreated handles this case with default header values.

ProjectsProjectsNotificationTemplatesSuccessCreateCreated projects projects notification templates success create created
*/
type ProjectsProjectsNotificationTemplatesSuccessCreateCreated struct {
}

func (o *ProjectsProjectsNotificationTemplatesSuccessCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{id}/notification_templates_success/][%d] projectsProjectsNotificationTemplatesSuccessCreateCreated ", 201)
}

func (o *ProjectsProjectsNotificationTemplatesSuccessCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsProjectsNotificationTemplatesSuccessCreateBody projects projects notification templates success create body
swagger:model ProjectsProjectsNotificationTemplatesSuccessCreateBody
*/
type ProjectsProjectsNotificationTemplatesSuccessCreateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// Optional custom messages for notification template.
	Messages string `json:"messages,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// notification configuration
	NotificationConfiguration string `json:"notification_configuration,omitempty"`

	// notification type
	// Required: true
	NotificationType *string `json:"notification_type"`

	// organization
	// Required: true
	Organization *int64 `json:"organization"`
}

// Validate validates this projects projects notification templates success create body
func (o *ProjectsProjectsNotificationTemplatesSuccessCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNotificationType(formats); err != nil {
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

func (o *ProjectsProjectsNotificationTemplatesSuccessCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsProjectsNotificationTemplatesSuccessCreateBody) validateNotificationType(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"notification_type", "body", o.NotificationType); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsProjectsNotificationTemplatesSuccessCreateBody) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organization", "body", o.Organization); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ProjectsProjectsNotificationTemplatesSuccessCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectsProjectsNotificationTemplatesSuccessCreateBody) UnmarshalBinary(b []byte) error {
	var res ProjectsProjectsNotificationTemplatesSuccessCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}