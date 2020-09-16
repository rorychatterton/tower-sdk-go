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

// SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateReader is a Reader for the SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreate structure.
type SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateCreated creates a SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateCreated with default headers values
func NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateCreated() *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateCreated {
	return &SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateCreated{}
}

/*SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateCreated handles this case with default header values.

SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateCreated system job templates system job templates notification templates success create created
*/
type SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateCreated struct {
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/system_job_templates/{id}/notification_templates_success/][%d] systemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateCreated ", 201)
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateBody system job templates system job templates notification templates success create body
swagger:model SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateBody
*/
type SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateBody struct {

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

// Validate validates this system job templates system job templates notification templates success create body
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateBody) Validate(formats strfmt.Registry) error {
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

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateBody) validateNotificationType(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"notification_type", "body", o.NotificationType); err != nil {
		return err
	}

	return nil
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateBody) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organization", "body", o.Organization); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateBody) UnmarshalBinary(b []byte) error {
	var res SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
