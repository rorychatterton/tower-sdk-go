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

// CustomInventoryScriptsInventoryScriptsUpdateReader is a Reader for the CustomInventoryScriptsInventoryScriptsUpdate structure.
type CustomInventoryScriptsInventoryScriptsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomInventoryScriptsInventoryScriptsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomInventoryScriptsInventoryScriptsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomInventoryScriptsInventoryScriptsUpdateOK creates a CustomInventoryScriptsInventoryScriptsUpdateOK with default headers values
func NewCustomInventoryScriptsInventoryScriptsUpdateOK() *CustomInventoryScriptsInventoryScriptsUpdateOK {
	return &CustomInventoryScriptsInventoryScriptsUpdateOK{}
}

/*CustomInventoryScriptsInventoryScriptsUpdateOK handles this case with default header values.

OK
*/
type CustomInventoryScriptsInventoryScriptsUpdateOK struct {
}

func (o *CustomInventoryScriptsInventoryScriptsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/inventory_scripts/{id}/][%d] customInventoryScriptsInventoryScriptsUpdateOK ", 200)
}

func (o *CustomInventoryScriptsInventoryScriptsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CustomInventoryScriptsInventoryScriptsUpdateBody custom inventory scripts inventory scripts update body
swagger:model CustomInventoryScriptsInventoryScriptsUpdateBody
*/
type CustomInventoryScriptsInventoryScriptsUpdateBody struct {

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

// Validate validates this custom inventory scripts inventory scripts update body
func (o *CustomInventoryScriptsInventoryScriptsUpdateBody) Validate(formats strfmt.Registry) error {
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

func (o *CustomInventoryScriptsInventoryScriptsUpdateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CustomInventoryScriptsInventoryScriptsUpdateBody) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organization", "body", o.Organization); err != nil {
		return err
	}

	return nil
}

func (o *CustomInventoryScriptsInventoryScriptsUpdateBody) validateScript(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"script", "body", o.Script); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CustomInventoryScriptsInventoryScriptsUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CustomInventoryScriptsInventoryScriptsUpdateBody) UnmarshalBinary(b []byte) error {
	var res CustomInventoryScriptsInventoryScriptsUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}