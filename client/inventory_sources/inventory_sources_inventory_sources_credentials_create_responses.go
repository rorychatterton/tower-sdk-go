// Code generated by go-swagger; DO NOT EDIT.

package inventory_sources

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

// InventorySourcesInventorySourcesCredentialsCreateReader is a Reader for the InventorySourcesInventorySourcesCredentialsCreate structure.
type InventorySourcesInventorySourcesCredentialsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventorySourcesInventorySourcesCredentialsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInventorySourcesInventorySourcesCredentialsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInventorySourcesInventorySourcesCredentialsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventorySourcesInventorySourcesCredentialsCreateCreated creates a InventorySourcesInventorySourcesCredentialsCreateCreated with default headers values
func NewInventorySourcesInventorySourcesCredentialsCreateCreated() *InventorySourcesInventorySourcesCredentialsCreateCreated {
	return &InventorySourcesInventorySourcesCredentialsCreateCreated{}
}

/*InventorySourcesInventorySourcesCredentialsCreateCreated handles this case with default header values.

InventorySourcesInventorySourcesCredentialsCreateCreated inventory sources inventory sources credentials create created
*/
type InventorySourcesInventorySourcesCredentialsCreateCreated struct {
}

func (o *InventorySourcesInventorySourcesCredentialsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/inventory_sources/{id}/credentials/][%d] inventorySourcesInventorySourcesCredentialsCreateCreated ", 201)
}

func (o *InventorySourcesInventorySourcesCredentialsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInventorySourcesInventorySourcesCredentialsCreateBadRequest creates a InventorySourcesInventorySourcesCredentialsCreateBadRequest with default headers values
func NewInventorySourcesInventorySourcesCredentialsCreateBadRequest() *InventorySourcesInventorySourcesCredentialsCreateBadRequest {
	return &InventorySourcesInventorySourcesCredentialsCreateBadRequest{}
}

/*InventorySourcesInventorySourcesCredentialsCreateBadRequest handles this case with default header values.

Bad Request
*/
type InventorySourcesInventorySourcesCredentialsCreateBadRequest struct {
}

func (o *InventorySourcesInventorySourcesCredentialsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/inventory_sources/{id}/credentials/][%d] inventorySourcesInventorySourcesCredentialsCreateBadRequest ", 400)
}

func (o *InventorySourcesInventorySourcesCredentialsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*InventorySourcesInventorySourcesCredentialsCreateBody inventory sources inventory sources credentials create body
swagger:model InventorySourcesInventorySourcesCredentialsCreateBody
*/
type InventorySourcesInventorySourcesCredentialsCreateBody struct {

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

// Validate validates this inventory sources inventory sources credentials create body
func (o *InventorySourcesInventorySourcesCredentialsCreateBody) Validate(formats strfmt.Registry) error {
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

func (o *InventorySourcesInventorySourcesCredentialsCreateBody) validateCredentialType(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"credential_type", "body", o.CredentialType); err != nil {
		return err
	}

	return nil
}

func (o *InventorySourcesInventorySourcesCredentialsCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InventorySourcesInventorySourcesCredentialsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InventorySourcesInventorySourcesCredentialsCreateBody) UnmarshalBinary(b []byte) error {
	var res InventorySourcesInventorySourcesCredentialsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
