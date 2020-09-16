// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationsOrganizationsJobTemplatesListReader is a Reader for the OrganizationsOrganizationsJobTemplatesList structure.
type OrganizationsOrganizationsJobTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsJobTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsOrganizationsJobTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrganizationsOrganizationsJobTemplatesListOK creates a OrganizationsOrganizationsJobTemplatesListOK with default headers values
func NewOrganizationsOrganizationsJobTemplatesListOK() *OrganizationsOrganizationsJobTemplatesListOK {
	return &OrganizationsOrganizationsJobTemplatesListOK{}
}

/*OrganizationsOrganizationsJobTemplatesListOK handles this case with default header values.

OK
*/
type OrganizationsOrganizationsJobTemplatesListOK struct {
}

func (o *OrganizationsOrganizationsJobTemplatesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/job_templates/][%d] organizationsOrganizationsJobTemplatesListOK ", 200)
}

func (o *OrganizationsOrganizationsJobTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}