// Code generated by go-swagger; DO NOT EDIT.

package credentials

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

// CredentialsCredentialsCopyCreateReader is a Reader for the CredentialsCredentialsCopyCreate structure.
type CredentialsCredentialsCopyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialsCredentialsCopyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCredentialsCredentialsCopyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCredentialsCredentialsCopyCreateCreated creates a CredentialsCredentialsCopyCreateCreated with default headers values
func NewCredentialsCredentialsCopyCreateCreated() *CredentialsCredentialsCopyCreateCreated {
	return &CredentialsCredentialsCopyCreateCreated{}
}

/*CredentialsCredentialsCopyCreateCreated handles this case with default header values.

CredentialsCredentialsCopyCreateCreated credentials credentials copy create created
*/
type CredentialsCredentialsCopyCreateCreated struct {
}

func (o *CredentialsCredentialsCopyCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/credentials/{id}/copy/][%d] credentialsCredentialsCopyCreateCreated ", 201)
}

func (o *CredentialsCredentialsCopyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CredentialsCredentialsCopyCreateBody credentials credentials copy create body
swagger:model CredentialsCredentialsCopyCreateBody
*/
type CredentialsCredentialsCopyCreateBody struct {

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this credentials credentials copy create body
func (o *CredentialsCredentialsCopyCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CredentialsCredentialsCopyCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CredentialsCredentialsCopyCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CredentialsCredentialsCopyCreateBody) UnmarshalBinary(b []byte) error {
	var res CredentialsCredentialsCopyCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}