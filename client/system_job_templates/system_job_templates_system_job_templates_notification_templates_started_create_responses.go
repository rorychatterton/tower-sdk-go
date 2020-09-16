// Code generated by go-swagger; DO NOT EDIT.

package system_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateReader is a Reader for the SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreate structure.
type SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated creates a SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated with default headers values
func NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated() *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated {
	return &SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated{}
}

/*SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated handles this case with default header values.

SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated system job templates system job templates notification templates started create created
*/
type SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated struct {
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/system_job_templates/{id}/notification_templates_started/][%d] systemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated ", 201)
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}