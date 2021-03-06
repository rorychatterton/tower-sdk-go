// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationsOrganizationsNotificationTemplatesListReader is a Reader for the OrganizationsOrganizationsNotificationTemplatesList structure.
type OrganizationsOrganizationsNotificationTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsNotificationTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsOrganizationsNotificationTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrganizationsOrganizationsNotificationTemplatesListOK creates a OrganizationsOrganizationsNotificationTemplatesListOK with default headers values
func NewOrganizationsOrganizationsNotificationTemplatesListOK() *OrganizationsOrganizationsNotificationTemplatesListOK {
	return &OrganizationsOrganizationsNotificationTemplatesListOK{}
}

/*OrganizationsOrganizationsNotificationTemplatesListOK handles this case with default header values.

OK
*/
type OrganizationsOrganizationsNotificationTemplatesListOK struct {
}

func (o *OrganizationsOrganizationsNotificationTemplatesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/notification_templates/][%d] organizationsOrganizationsNotificationTemplatesListOK ", 200)
}

func (o *OrganizationsOrganizationsNotificationTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
