// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationsOrganizationsNotificationTemplatesApprovalsListReader is a Reader for the OrganizationsOrganizationsNotificationTemplatesApprovalsList structure.
type OrganizationsOrganizationsNotificationTemplatesApprovalsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsOrganizationsNotificationTemplatesApprovalsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrganizationsOrganizationsNotificationTemplatesApprovalsListOK creates a OrganizationsOrganizationsNotificationTemplatesApprovalsListOK with default headers values
func NewOrganizationsOrganizationsNotificationTemplatesApprovalsListOK() *OrganizationsOrganizationsNotificationTemplatesApprovalsListOK {
	return &OrganizationsOrganizationsNotificationTemplatesApprovalsListOK{}
}

/*OrganizationsOrganizationsNotificationTemplatesApprovalsListOK handles this case with default header values.

OK
*/
type OrganizationsOrganizationsNotificationTemplatesApprovalsListOK struct {
}

func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/notification_templates_approvals/][%d] organizationsOrganizationsNotificationTemplatesApprovalsListOK ", 200)
}

func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}