// Code generated by go-swagger; DO NOT EDIT.

package job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobTemplatesJobTemplatesNotificationTemplatesStartedListReader is a Reader for the JobTemplatesJobTemplatesNotificationTemplatesStartedList structure.
type JobTemplatesJobTemplatesNotificationTemplatesStartedListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobTemplatesJobTemplatesNotificationTemplatesStartedListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobTemplatesJobTemplatesNotificationTemplatesStartedListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobTemplatesJobTemplatesNotificationTemplatesStartedListOK creates a JobTemplatesJobTemplatesNotificationTemplatesStartedListOK with default headers values
func NewJobTemplatesJobTemplatesNotificationTemplatesStartedListOK() *JobTemplatesJobTemplatesNotificationTemplatesStartedListOK {
	return &JobTemplatesJobTemplatesNotificationTemplatesStartedListOK{}
}

/*JobTemplatesJobTemplatesNotificationTemplatesStartedListOK handles this case with default header values.

OK
*/
type JobTemplatesJobTemplatesNotificationTemplatesStartedListOK struct {
}

func (o *JobTemplatesJobTemplatesNotificationTemplatesStartedListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/job_templates/{id}/notification_templates_started/][%d] jobTemplatesJobTemplatesNotificationTemplatesStartedListOK ", 200)
}

func (o *JobTemplatesJobTemplatesNotificationTemplatesStartedListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}