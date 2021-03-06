// Code generated by go-swagger; DO NOT EDIT.

package labels

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

// LabelsLabelsUpdateReader is a Reader for the LabelsLabelsUpdate structure.
type LabelsLabelsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LabelsLabelsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLabelsLabelsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLabelsLabelsUpdateOK creates a LabelsLabelsUpdateOK with default headers values
func NewLabelsLabelsUpdateOK() *LabelsLabelsUpdateOK {
	return &LabelsLabelsUpdateOK{}
}

/*LabelsLabelsUpdateOK handles this case with default header values.

OK
*/
type LabelsLabelsUpdateOK struct {
}

func (o *LabelsLabelsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/labels/{id}/][%d] labelsLabelsUpdateOK ", 200)
}

func (o *LabelsLabelsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*LabelsLabelsUpdateBody labels labels update body
swagger:model LabelsLabelsUpdateBody
*/
type LabelsLabelsUpdateBody struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// Organization this label belongs to.
	// Required: true
	Organization *int64 `json:"organization"`
}

// Validate validates this labels labels update body
func (o *LabelsLabelsUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

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

func (o *LabelsLabelsUpdateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *LabelsLabelsUpdateBody) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organization", "body", o.Organization); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LabelsLabelsUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LabelsLabelsUpdateBody) UnmarshalBinary(b []byte) error {
	var res LabelsLabelsUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
