// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationsOrganizationsUsersListReader is a Reader for the OrganizationsOrganizationsUsersList structure.
type OrganizationsOrganizationsUsersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsUsersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsOrganizationsUsersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrganizationsOrganizationsUsersListOK creates a OrganizationsOrganizationsUsersListOK with default headers values
func NewOrganizationsOrganizationsUsersListOK() *OrganizationsOrganizationsUsersListOK {
	return &OrganizationsOrganizationsUsersListOK{}
}

/*OrganizationsOrganizationsUsersListOK handles this case with default header values.

OK
*/
type OrganizationsOrganizationsUsersListOK struct {
}

func (o *OrganizationsOrganizationsUsersListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/users/][%d] organizationsOrganizationsUsersListOK ", 200)
}

func (o *OrganizationsOrganizationsUsersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
