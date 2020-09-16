// Code generated by go-swagger; DO NOT EDIT.

package custom_inventory_scripts

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

// CustomInventoryScriptsInventoryScriptsCreateReader is a Reader for the CustomInventoryScriptsInventoryScriptsCreate structure.
type CustomInventoryScriptsInventoryScriptsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomInventoryScriptsInventoryScriptsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCustomInventoryScriptsInventoryScriptsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomInventoryScriptsInventoryScriptsCreateCreated creates a CustomInventoryScriptsInventoryScriptsCreateCreated with default headers values
func NewCustomInventoryScriptsInventoryScriptsCreateCreated() *CustomInventoryScriptsInventoryScriptsCreateCreated {
	return &CustomInventoryScriptsInventoryScriptsCreateCreated{}
}

/*CustomInventoryScriptsInventoryScriptsCreateCreated handles this case with default header values.

CustomInventoryScriptsInventoryScriptsCreateCreated custom inventory scripts inventory scripts create created
*/
type CustomInventoryScriptsInventoryScriptsCreateCreated struct {
}

func (o *CustomInventoryScriptsInventoryScriptsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/inventory_scripts/][%d] customInventoryScriptsInventoryScriptsCreateCreated ", 201)
}

func (o *CustomInventoryScriptsInventoryScriptsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CustomInventoryScriptsInventoryScriptsCreateBody custom inventory scripts inventory scripts create body
swagger:model CustomInventoryScriptsInventoryScriptsCreateBody
*/
type CustomInventoryScriptsInventoryScriptsCreateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// Organization owning this inventory script
	// Required: true
	Organization *int64 `json:"organization"`

	// script
	// Required: true
	Script *string `json:"script"`
}

// Validate validates this custom inventory scripts inventory scripts create body
func (o *CustomInventoryScriptsInventoryScriptsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScript(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CustomInventoryScriptsInventoryScriptsCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CustomInventoryScriptsInventoryScriptsCreateBody) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organization", "body", o.Organization); err != nil {
		return err
	}

	return nil
}

func (o *CustomInventoryScriptsInventoryScriptsCreateBody) validateScript(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"script", "body", o.Script); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CustomInventoryScriptsInventoryScriptsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CustomInventoryScriptsInventoryScriptsCreateBody) UnmarshalBinary(b []byte) error {
	var res CustomInventoryScriptsInventoryScriptsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}