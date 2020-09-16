// Code generated by go-swagger; DO NOT EDIT.

package groups

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

// GroupsGroupsCreateReader is a Reader for the GroupsGroupsCreate structure.
type GroupsGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGroupsGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGroupsGroupsCreateCreated creates a GroupsGroupsCreateCreated with default headers values
func NewGroupsGroupsCreateCreated() *GroupsGroupsCreateCreated {
	return &GroupsGroupsCreateCreated{}
}

/*GroupsGroupsCreateCreated handles this case with default header values.

GroupsGroupsCreateCreated groups groups create created
*/
type GroupsGroupsCreateCreated struct {
}

func (o *GroupsGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/][%d] groupsGroupsCreateCreated ", 201)
}

func (o *GroupsGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GroupsGroupsCreateBody groups groups create body
swagger:model GroupsGroupsCreateBody
*/
type GroupsGroupsCreateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// inventory
	// Required: true
	Inventory *int64 `json:"inventory"`

	// name
	// Required: true
	Name *string `json:"name"`

	// Group variables in JSON or YAML format.
	Variables string `json:"variables,omitempty"`
}

// Validate validates this groups groups create body
func (o *GroupsGroupsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInventory(formats); err != nil {
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

func (o *GroupsGroupsCreateBody) validateInventory(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"inventory", "body", o.Inventory); err != nil {
		return err
	}

	return nil
}

func (o *GroupsGroupsCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GroupsGroupsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GroupsGroupsCreateBody) UnmarshalBinary(b []byte) error {
	var res GroupsGroupsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
