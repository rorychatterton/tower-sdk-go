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

// CredentialTypesCredentialTypesUpdateReader is a Reader for the CredentialTypesCredentialTypesUpdate structure.
type CredentialTypesCredentialTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialTypesCredentialTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCredentialTypesCredentialTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCredentialTypesCredentialTypesUpdateOK creates a CredentialTypesCredentialTypesUpdateOK with default headers values
func NewCredentialTypesCredentialTypesUpdateOK() *CredentialTypesCredentialTypesUpdateOK {
	return &CredentialTypesCredentialTypesUpdateOK{}
}

/*CredentialTypesCredentialTypesUpdateOK handles this case with default header values.

OK
*/
type CredentialTypesCredentialTypesUpdateOK struct {
}

func (o *CredentialTypesCredentialTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/credential_types/{id}/][%d] credentialTypesCredentialTypesUpdateOK ", 200)
}

func (o *CredentialTypesCredentialTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CredentialTypesCredentialTypesUpdateBody credential types credential types update body
swagger:model CredentialTypesCredentialTypesUpdateBody
*/
type CredentialTypesCredentialTypesUpdateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// Enter injectors using either JSON or YAML syntax. Refer to the Ansible Tower documentation for example syntax.
	Injectors interface{} `json:"injectors,omitempty"`

	// Enter inputs using either JSON or YAML syntax. Refer to the Ansible Tower documentation for example syntax.
	Inputs interface{} `json:"inputs,omitempty"`

	// kind
	// Required: true
	Kind *string `json:"kind"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this credential types credential types update body
func (o *CredentialTypesCredentialTypesUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKind(formats); err != nil {
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

func (o *CredentialTypesCredentialTypesUpdateBody) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"kind", "body", o.Kind); err != nil {
		return err
	}

	return nil
}

func (o *CredentialTypesCredentialTypesUpdateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CredentialTypesCredentialTypesUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CredentialTypesCredentialTypesUpdateBody) UnmarshalBinary(b []byte) error {
	var res CredentialTypesCredentialTypesUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
