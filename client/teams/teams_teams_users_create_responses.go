// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// TeamsTeamsUsersCreateReader is a Reader for the TeamsTeamsUsersCreate structure.
type TeamsTeamsUsersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamsTeamsUsersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTeamsTeamsUsersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamsTeamsUsersCreateCreated creates a TeamsTeamsUsersCreateCreated with default headers values
func NewTeamsTeamsUsersCreateCreated() *TeamsTeamsUsersCreateCreated {
	return &TeamsTeamsUsersCreateCreated{}
}

/*TeamsTeamsUsersCreateCreated handles this case with default header values.

TeamsTeamsUsersCreateCreated teams teams users create created
*/
type TeamsTeamsUsersCreateCreated struct {
}

func (o *TeamsTeamsUsersCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams/{id}/users/][%d] teamsTeamsUsersCreateCreated ", 201)
}

func (o *TeamsTeamsUsersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*TeamsTeamsUsersCreateBody teams teams users create body
swagger:model TeamsTeamsUsersCreateBody
*/
type TeamsTeamsUsersCreateBody struct {

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// Designates that this user has all permissions without explicitly assigning them.
	IsSuperuser string `json:"is_superuser,omitempty"`

	// is system auditor
	IsSystemAuditor bool `json:"is_system_auditor,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// Write-only field used to change the password.
	Password string `json:"password,omitempty"`

	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this teams teams users create body
func (o *TeamsTeamsUsersCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TeamsTeamsUsersCreateBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TeamsTeamsUsersCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TeamsTeamsUsersCreateBody) UnmarshalBinary(b []byte) error {
	var res TeamsTeamsUsersCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
