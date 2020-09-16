// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationsOrganizationsListReader is a Reader for the OrganizationsOrganizationsList structure.
type OrganizationsOrganizationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsOrganizationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewOrganizationsOrganizationsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrganizationsOrganizationsListOK creates a OrganizationsOrganizationsListOK with default headers values
func NewOrganizationsOrganizationsListOK() *OrganizationsOrganizationsListOK {
	return &OrganizationsOrganizationsListOK{}
}

/*OrganizationsOrganizationsListOK handles this case with default header values.

OK
*/
type OrganizationsOrganizationsListOK struct {
}

func (o *OrganizationsOrganizationsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/][%d] organizationsOrganizationsListOK ", 200)
}

func (o *OrganizationsOrganizationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrganizationsOrganizationsListUnauthorized creates a OrganizationsOrganizationsListUnauthorized with default headers values
func NewOrganizationsOrganizationsListUnauthorized() *OrganizationsOrganizationsListUnauthorized {
	return &OrganizationsOrganizationsListUnauthorized{}
}

/*OrganizationsOrganizationsListUnauthorized handles this case with default header values.

Unauthorised
*/
type OrganizationsOrganizationsListUnauthorized struct {
}

func (o *OrganizationsOrganizationsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/][%d] organizationsOrganizationsListUnauthorized ", 401)
}

func (o *OrganizationsOrganizationsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}