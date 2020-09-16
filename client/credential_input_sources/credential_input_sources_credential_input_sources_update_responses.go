// Code generated by go-swagger; DO NOT EDIT.

package credential_input_sources

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

// CredentialInputSourcesCredentialInputSourcesUpdateReader is a Reader for the CredentialInputSourcesCredentialInputSourcesUpdate structure.
type CredentialInputSourcesCredentialInputSourcesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialInputSourcesCredentialInputSourcesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCredentialInputSourcesCredentialInputSourcesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCredentialInputSourcesCredentialInputSourcesUpdateOK creates a CredentialInputSourcesCredentialInputSourcesUpdateOK with default headers values
func NewCredentialInputSourcesCredentialInputSourcesUpdateOK() *CredentialInputSourcesCredentialInputSourcesUpdateOK {
	return &CredentialInputSourcesCredentialInputSourcesUpdateOK{}
}

/*CredentialInputSourcesCredentialInputSourcesUpdateOK handles this case with default header values.

OK
*/
type CredentialInputSourcesCredentialInputSourcesUpdateOK struct {
}

func (o *CredentialInputSourcesCredentialInputSourcesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/credential_input_sources/{id}/][%d] credentialInputSourcesCredentialInputSourcesUpdateOK ", 200)
}

func (o *CredentialInputSourcesCredentialInputSourcesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CredentialInputSourcesCredentialInputSourcesUpdateBody credential input sources credential input sources update body
swagger:model CredentialInputSourcesCredentialInputSourcesUpdateBody
*/
type CredentialInputSourcesCredentialInputSourcesUpdateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// input field name
	// Required: true
	InputFieldName *string `json:"input_field_name"`

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// source credential
	// Required: true
	SourceCredential *int64 `json:"source_credential"`

	// target credential
	// Required: true
	TargetCredential *int64 `json:"target_credential"`
}

// Validate validates this credential input sources credential input sources update body
func (o *CredentialInputSourcesCredentialInputSourcesUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInputFieldName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSourceCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTargetCredential(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CredentialInputSourcesCredentialInputSourcesUpdateBody) validateInputFieldName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"input_field_name", "body", o.InputFieldName); err != nil {
		return err
	}

	return nil
}

func (o *CredentialInputSourcesCredentialInputSourcesUpdateBody) validateSourceCredential(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"source_credential", "body", o.SourceCredential); err != nil {
		return err
	}

	return nil
}

func (o *CredentialInputSourcesCredentialInputSourcesUpdateBody) validateTargetCredential(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"target_credential", "body", o.TargetCredential); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CredentialInputSourcesCredentialInputSourcesUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CredentialInputSourcesCredentialInputSourcesUpdateBody) UnmarshalBinary(b []byte) error {
	var res CredentialInputSourcesCredentialInputSourcesUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}