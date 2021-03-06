// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// AuthenticationApplicationsUpdate0Reader is a Reader for the AuthenticationApplicationsUpdate0 structure.
type AuthenticationApplicationsUpdate0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthenticationApplicationsUpdate0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthenticationApplicationsUpdate0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAuthenticationApplicationsUpdate0OK creates a AuthenticationApplicationsUpdate0OK with default headers values
func NewAuthenticationApplicationsUpdate0OK() *AuthenticationApplicationsUpdate0OK {
	return &AuthenticationApplicationsUpdate0OK{}
}

/*AuthenticationApplicationsUpdate0OK handles this case with default header values.

OK
*/
type AuthenticationApplicationsUpdate0OK struct {
}

func (o *AuthenticationApplicationsUpdate0OK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/applications/{id}/][%d] authenticationApplicationsUpdate0OK ", 200)
}

func (o *AuthenticationApplicationsUpdate0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*AuthenticationApplicationsUpdate0Body authentication applications update0 body
swagger:model AuthenticationApplicationsUpdate0Body
*/
type AuthenticationApplicationsUpdate0Body struct {

	// The Grant type the user must use for acquire tokens for this application.
	// Required: true
	AuthorizationGrantType *string `json:"authorization_grant_type"`

	// Set to Public or Confidential depending on how secure the client device is.
	// Required: true
	ClientType *string `json:"client_type"`

	// description
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// Organization containing this application.
	// Required: true
	Organization *int64 `json:"organization"`

	// Allowed URIs list, space separated
	RedirectUris string `json:"redirect_uris,omitempty"`

	// Set True to skip authorization step for completely trusted applications.
	SkipAuthorization string `json:"skip_authorization,omitempty"`
}

// Validate validates this authentication applications update0 body
func (o *AuthenticationApplicationsUpdate0Body) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAuthorizationGrantType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateClientType(formats); err != nil {
		res = append(res, err)
	}

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

func (o *AuthenticationApplicationsUpdate0Body) validateAuthorizationGrantType(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"authorization_grant_type", "body", o.AuthorizationGrantType); err != nil {
		return err
	}

	return nil
}

func (o *AuthenticationApplicationsUpdate0Body) validateClientType(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"client_type", "body", o.ClientType); err != nil {
		return err
	}

	return nil
}

func (o *AuthenticationApplicationsUpdate0Body) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *AuthenticationApplicationsUpdate0Body) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organization", "body", o.Organization); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AuthenticationApplicationsUpdate0Body) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AuthenticationApplicationsUpdate0Body) UnmarshalBinary(b []byte) error {
	var res AuthenticationApplicationsUpdate0Body
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
