// Code generated by go-swagger; DO NOT EDIT.

package system_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListReader is a Reader for the SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessList structure.
type SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListOK creates a SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListOK with default headers values
func NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListOK() *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListOK {
	return &SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListOK{}
}

/*SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListOK handles this case with default header values.

OK
*/
type SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListOK struct {
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/system_job_templates/{id}/notification_templates_success/][%d] systemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListOK ", 200)
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}