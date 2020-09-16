// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationsOrganizationsActivityStreamListReader is a Reader for the OrganizationsOrganizationsActivityStreamList structure.
type OrganizationsOrganizationsActivityStreamListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsActivityStreamListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsOrganizationsActivityStreamListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrganizationsOrganizationsActivityStreamListOK creates a OrganizationsOrganizationsActivityStreamListOK with default headers values
func NewOrganizationsOrganizationsActivityStreamListOK() *OrganizationsOrganizationsActivityStreamListOK {
	return &OrganizationsOrganizationsActivityStreamListOK{}
}

/*OrganizationsOrganizationsActivityStreamListOK handles this case with default header values.

OK
*/
type OrganizationsOrganizationsActivityStreamListOK struct {
}

func (o *OrganizationsOrganizationsActivityStreamListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/activity_stream/][%d] organizationsOrganizationsActivityStreamListOK ", 200)
}

func (o *OrganizationsOrganizationsActivityStreamListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}