// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// OrganizationsOrganizationsNotificationTemplatesErrorCreateReader is a Reader for the OrganizationsOrganizationsNotificationTemplatesErrorCreate structure.
type OrganizationsOrganizationsNotificationTemplatesErrorCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsNotificationTemplatesErrorCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewOrganizationsOrganizationsNotificationTemplatesErrorCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrganizationsOrganizationsNotificationTemplatesErrorCreateCreated creates a OrganizationsOrganizationsNotificationTemplatesErrorCreateCreated with default headers values
func NewOrganizationsOrganizationsNotificationTemplatesErrorCreateCreated() *OrganizationsOrganizationsNotificationTemplatesErrorCreateCreated {
	return &OrganizationsOrganizationsNotificationTemplatesErrorCreateCreated{}
}

/*OrganizationsOrganizationsNotificationTemplatesErrorCreateCreated handles this case with default header values.

OrganizationsOrganizationsNotificationTemplatesErrorCreateCreated organizations organizations notification templates error create created
*/
type OrganizationsOrganizationsNotificationTemplatesErrorCreateCreated struct {
}

func (o *OrganizationsOrganizationsNotificationTemplatesErrorCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/organizations/{id}/notification_templates_error/][%d] organizationsOrganizationsNotificationTemplatesErrorCreateCreated ", 201)
}

func (o *OrganizationsOrganizationsNotificationTemplatesErrorCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*OrganizationsOrganizationsNotificationTemplatesErrorCreateBody organizations organizations notification templates error create body
swagger:model OrganizationsOrganizationsNotificationTemplatesErrorCreateBody
*/
type OrganizationsOrganizationsNotificationTemplatesErrorCreateBody struct {

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

// Validate validates this organizations organizations notification templates error create body
func (o *OrganizationsOrganizationsNotificationTemplatesErrorCreateBody) Validate(formats strfmt.Registry) error {
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

func (o *OrganizationsOrganizationsNotificationTemplatesErrorCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *OrganizationsOrganizationsNotificationTemplatesErrorCreateBody) validateNotificationType(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"notification_type", "body", o.NotificationType); err != nil {
		return err
	}

	return nil
}

func (o *OrganizationsOrganizationsNotificationTemplatesErrorCreateBody) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organization", "body", o.Organization); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *OrganizationsOrganizationsNotificationTemplatesErrorCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OrganizationsOrganizationsNotificationTemplatesErrorCreateBody) UnmarshalBinary(b []byte) error {
	var res OrganizationsOrganizationsNotificationTemplatesErrorCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
