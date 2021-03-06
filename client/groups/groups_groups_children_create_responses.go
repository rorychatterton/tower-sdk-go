// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GroupsGroupsChildrenCreateReader is a Reader for the GroupsGroupsChildrenCreate structure.
type GroupsGroupsChildrenCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsGroupsChildrenCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGroupsGroupsChildrenCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGroupsGroupsChildrenCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGroupsGroupsChildrenCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGroupsGroupsChildrenCreateCreated creates a GroupsGroupsChildrenCreateCreated with default headers values
func NewGroupsGroupsChildrenCreateCreated() *GroupsGroupsChildrenCreateCreated {
	return &GroupsGroupsChildrenCreateCreated{}
}

/*GroupsGroupsChildrenCreateCreated handles this case with default header values.

GroupsGroupsChildrenCreateCreated groups groups children create created
*/
type GroupsGroupsChildrenCreateCreated struct {
}

func (o *GroupsGroupsChildrenCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{id}/children/][%d] groupsGroupsChildrenCreateCreated ", 201)
}

func (o *GroupsGroupsChildrenCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGroupsGroupsChildrenCreateBadRequest creates a GroupsGroupsChildrenCreateBadRequest with default headers values
func NewGroupsGroupsChildrenCreateBadRequest() *GroupsGroupsChildrenCreateBadRequest {
	return &GroupsGroupsChildrenCreateBadRequest{}
}

/*GroupsGroupsChildrenCreateBadRequest handles this case with default header values.

Bad Request
*/
type GroupsGroupsChildrenCreateBadRequest struct {
}

func (o *GroupsGroupsChildrenCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{id}/children/][%d] groupsGroupsChildrenCreateBadRequest ", 400)
}

func (o *GroupsGroupsChildrenCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGroupsGroupsChildrenCreateForbidden creates a GroupsGroupsChildrenCreateForbidden with default headers values
func NewGroupsGroupsChildrenCreateForbidden() *GroupsGroupsChildrenCreateForbidden {
	return &GroupsGroupsChildrenCreateForbidden{}
}

/*GroupsGroupsChildrenCreateForbidden handles this case with default header values.

No Permission Response
*/
type GroupsGroupsChildrenCreateForbidden struct {
}

func (o *GroupsGroupsChildrenCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{id}/children/][%d] groupsGroupsChildrenCreateForbidden ", 403)
}

func (o *GroupsGroupsChildrenCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
