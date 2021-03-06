// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// TeamsTeamsCreateReader is a Reader for the TeamsTeamsCreate structure.
type TeamsTeamsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamsTeamsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTeamsTeamsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamsTeamsCreateCreated creates a TeamsTeamsCreateCreated with default headers values
func NewTeamsTeamsCreateCreated() *TeamsTeamsCreateCreated {
	return &TeamsTeamsCreateCreated{}
}

/*TeamsTeamsCreateCreated handles this case with default header values.

TeamsTeamsCreateCreated teams teams create created
*/
type TeamsTeamsCreateCreated struct {
}

func (o *TeamsTeamsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams/][%d] teamsTeamsCreateCreated ", 201)
}

func (o *TeamsTeamsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*TeamsTeamsCreateBody teams teams create body
swagger:model TeamsTeamsCreateBody
*/
type TeamsTeamsCreateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// organization
	// Required: true
	Organization *int64 `json:"organization"`
}

// Validate validates this teams teams create body
func (o *TeamsTeamsCreateBody) Validate(formats strfmt.Registry) error {
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

func (o *TeamsTeamsCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *TeamsTeamsCreateBody) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organization", "body", o.Organization); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TeamsTeamsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TeamsTeamsCreateBody) UnmarshalBinary(b []byte) error {
	var res TeamsTeamsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
