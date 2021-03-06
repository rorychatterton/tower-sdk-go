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

// LabelsLabelsCreateReader is a Reader for the LabelsLabelsCreate structure.
type LabelsLabelsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LabelsLabelsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewLabelsLabelsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLabelsLabelsCreateCreated creates a LabelsLabelsCreateCreated with default headers values
func NewLabelsLabelsCreateCreated() *LabelsLabelsCreateCreated {
	return &LabelsLabelsCreateCreated{}
}

/*LabelsLabelsCreateCreated handles this case with default header values.

LabelsLabelsCreateCreated labels labels create created
*/
type LabelsLabelsCreateCreated struct {
}

func (o *LabelsLabelsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/labels/][%d] labelsLabelsCreateCreated ", 201)
}

func (o *LabelsLabelsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*LabelsLabelsCreateBody labels labels create body
swagger:model LabelsLabelsCreateBody
*/
type LabelsLabelsCreateBody struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// Organization this label belongs to.
	// Required: true
	Organization *int64 `json:"organization"`
}

// Validate validates this labels labels create body
func (o *LabelsLabelsCreateBody) Validate(formats strfmt.Registry) error {
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

func (o *LabelsLabelsCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *LabelsLabelsCreateBody) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organization", "body", o.Organization); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LabelsLabelsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LabelsLabelsCreateBody) UnmarshalBinary(b []byte) error {
	var res LabelsLabelsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
