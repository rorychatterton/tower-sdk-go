// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GroupsGroupsInventorySourcesListReader is a Reader for the GroupsGroupsInventorySourcesList structure.
type GroupsGroupsInventorySourcesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsGroupsInventorySourcesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupsGroupsInventorySourcesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGroupsGroupsInventorySourcesListOK creates a GroupsGroupsInventorySourcesListOK with default headers values
func NewGroupsGroupsInventorySourcesListOK() *GroupsGroupsInventorySourcesListOK {
	return &GroupsGroupsInventorySourcesListOK{}
}

/*GroupsGroupsInventorySourcesListOK handles this case with default header values.

OK
*/
type GroupsGroupsInventorySourcesListOK struct {
}

func (o *GroupsGroupsInventorySourcesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{id}/inventory_sources/][%d] groupsGroupsInventorySourcesListOK ", 200)
}

func (o *GroupsGroupsInventorySourcesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}