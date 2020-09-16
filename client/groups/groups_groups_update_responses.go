// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GroupsGroupsUpdateReader is a Reader for the GroupsGroupsUpdate structure.
type GroupsGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupsGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGroupsGroupsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGroupsGroupsUpdateOK creates a GroupsGroupsUpdateOK with default headers values
func NewGroupsGroupsUpdateOK() *GroupsGroupsUpdateOK {
	return &GroupsGroupsUpdateOK{}
}

/*GroupsGroupsUpdateOK handles this case with default header values.

OK
*/
type GroupsGroupsUpdateOK struct {
}

func (o *GroupsGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{id}/][%d] groupsGroupsUpdateOK ", 200)
}

func (o *GroupsGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGroupsGroupsUpdateForbidden creates a GroupsGroupsUpdateForbidden with default headers values
func NewGroupsGroupsUpdateForbidden() *GroupsGroupsUpdateForbidden {
	return &GroupsGroupsUpdateForbidden{}
}

/*GroupsGroupsUpdateForbidden handles this case with default header values.

No Permission Response
*/
type GroupsGroupsUpdateForbidden struct {
}

func (o *GroupsGroupsUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{id}/][%d] groupsGroupsUpdateForbidden ", 403)
}

func (o *GroupsGroupsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
