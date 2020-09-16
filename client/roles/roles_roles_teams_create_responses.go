// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RolesRolesTeamsCreateReader is a Reader for the RolesRolesTeamsCreate structure.
type RolesRolesTeamsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RolesRolesTeamsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRolesRolesTeamsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRolesRolesTeamsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRolesRolesTeamsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRolesRolesTeamsCreateCreated creates a RolesRolesTeamsCreateCreated with default headers values
func NewRolesRolesTeamsCreateCreated() *RolesRolesTeamsCreateCreated {
	return &RolesRolesTeamsCreateCreated{}
}

/*RolesRolesTeamsCreateCreated handles this case with default header values.

RolesRolesTeamsCreateCreated roles roles teams create created
*/
type RolesRolesTeamsCreateCreated struct {
}

func (o *RolesRolesTeamsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/roles/{id}/teams/][%d] rolesRolesTeamsCreateCreated ", 201)
}

func (o *RolesRolesTeamsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRolesRolesTeamsCreateBadRequest creates a RolesRolesTeamsCreateBadRequest with default headers values
func NewRolesRolesTeamsCreateBadRequest() *RolesRolesTeamsCreateBadRequest {
	return &RolesRolesTeamsCreateBadRequest{}
}

/*RolesRolesTeamsCreateBadRequest handles this case with default header values.

Bad Request
*/
type RolesRolesTeamsCreateBadRequest struct {
}

func (o *RolesRolesTeamsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/roles/{id}/teams/][%d] rolesRolesTeamsCreateBadRequest ", 400)
}

func (o *RolesRolesTeamsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRolesRolesTeamsCreateForbidden creates a RolesRolesTeamsCreateForbidden with default headers values
func NewRolesRolesTeamsCreateForbidden() *RolesRolesTeamsCreateForbidden {
	return &RolesRolesTeamsCreateForbidden{}
}

/*RolesRolesTeamsCreateForbidden handles this case with default header values.

No Permission Response
*/
type RolesRolesTeamsCreateForbidden struct {
}

func (o *RolesRolesTeamsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/roles/{id}/teams/][%d] rolesRolesTeamsCreateForbidden ", 403)
}

func (o *RolesRolesTeamsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}