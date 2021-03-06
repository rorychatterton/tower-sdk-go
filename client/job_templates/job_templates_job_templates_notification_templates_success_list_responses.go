// Code generated by go-swagger; DO NOT EDIT.

package job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobTemplatesJobTemplatesNotificationTemplatesSuccessListReader is a Reader for the JobTemplatesJobTemplatesNotificationTemplatesSuccessList structure.
type JobTemplatesJobTemplatesNotificationTemplatesSuccessListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobTemplatesJobTemplatesNotificationTemplatesSuccessListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobTemplatesJobTemplatesNotificationTemplatesSuccessListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobTemplatesJobTemplatesNotificationTemplatesSuccessListOK creates a JobTemplatesJobTemplatesNotificationTemplatesSuccessListOK with default headers values
func NewJobTemplatesJobTemplatesNotificationTemplatesSuccessListOK() *JobTemplatesJobTemplatesNotificationTemplatesSuccessListOK {
	return &JobTemplatesJobTemplatesNotificationTemplatesSuccessListOK{}
}

/*JobTemplatesJobTemplatesNotificationTemplatesSuccessListOK handles this case with default header values.

OK
*/
type JobTemplatesJobTemplatesNotificationTemplatesSuccessListOK struct {
}

func (o *JobTemplatesJobTemplatesNotificationTemplatesSuccessListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/job_templates/{id}/notification_templates_success/][%d] jobTemplatesJobTemplatesNotificationTemplatesSuccessListOK ", 200)
}

func (o *JobTemplatesJobTemplatesNotificationTemplatesSuccessListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
