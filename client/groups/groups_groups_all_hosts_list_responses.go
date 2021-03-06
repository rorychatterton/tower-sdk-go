// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GroupsGroupsAllHostsListReader is a Reader for the GroupsGroupsAllHostsList structure.
type GroupsGroupsAllHostsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsGroupsAllHostsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupsGroupsAllHostsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGroupsGroupsAllHostsListOK creates a GroupsGroupsAllHostsListOK with default headers values
func NewGroupsGroupsAllHostsListOK() *GroupsGroupsAllHostsListOK {
	return &GroupsGroupsAllHostsListOK{}
}

/*GroupsGroupsAllHostsListOK handles this case with default header values.

OK
*/
type GroupsGroupsAllHostsListOK struct {
}

func (o *GroupsGroupsAllHostsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{id}/all_hosts/][%d] groupsGroupsAllHostsListOK ", 200)
}

func (o *GroupsGroupsAllHostsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
