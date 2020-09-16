// Code generated by go-swagger; DO NOT EDIT.

package credential_types

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

// CredentialTypesCredentialTypesCredentialsCreateReader is a Reader for the CredentialTypesCredentialTypesCredentialsCreate structure.
type CredentialTypesCredentialTypesCredentialsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialTypesCredentialTypesCredentialsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCredentialTypesCredentialTypesCredentialsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCredentialTypesCredentialTypesCredentialsCreateCreated creates a CredentialTypesCredentialTypesCredentialsCreateCreated with default headers values
func NewCredentialTypesCredentialTypesCredentialsCreateCreated() *CredentialTypesCredentialTypesCredentialsCreateCreated {
	return &CredentialTypesCredentialTypesCredentialsCreateCreated{}
}

/*CredentialTypesCredentialTypesCredentialsCreateCreated handles this case with default header values.

CredentialTypesCredentialTypesCredentialsCreateCreated credential types credential types credentials create created
*/
type CredentialTypesCredentialTypesCredentialsCreateCreated struct {
}

func (o *CredentialTypesCredentialTypesCredentialsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/credential_types/{id}/credentials/][%d] credentialTypesCredentialTypesCredentialsCreateCreated ", 201)
}

func (o *CredentialTypesCredentialTypesCredentialsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CredentialTypesCredentialTypesCredentialsCreateBody credential types credential types credentials create body
swagger:model CredentialTypesCredentialTypesCredentialsCreateBody
*/
type CredentialTypesCredentialTypesCredentialsCreateBody struct {

	// Specify the type of credential you want to create. Refer to the Ansible Tower documentation for details on each type.
	// Required: true
	CredentialType *int64 `json:"credential_type"`

	// description
	Description string `json:"description,omitempty"`

	// Enter inputs using either JSON or YAML syntax. Refer to the Ansible Tower documentation for example syntax.
	Inputs interface{} `json:"inputs,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// organization
	Organization int64 `json:"organization,omitempty"`
}

// Validate validates this credential types credential types credentials create body
func (o *CredentialTypesCredentialTypesCredentialsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCredentialType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CredentialTypesCredentialTypesCredentialsCreateBody) validateCredentialType(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"credential_type", "body", o.CredentialType); err != nil {
		return err
	}

	return nil
}

func (o *CredentialTypesCredentialTypesCredentialsCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CredentialTypesCredentialTypesCredentialsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CredentialTypesCredentialTypesCredentialsCreateBody) UnmarshalBinary(b []byte) error {
	var res CredentialTypesCredentialTypesCredentialsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}