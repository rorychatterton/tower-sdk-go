// Code generated by go-swagger; DO NOT EDIT.

package job_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobEventsJobEventsListReader is a Reader for the JobEventsJobEventsList structure.
type JobEventsJobEventsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobEventsJobEventsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobEventsJobEventsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobEventsJobEventsListOK creates a JobEventsJobEventsListOK with default headers values
func NewJobEventsJobEventsListOK() *JobEventsJobEventsListOK {
	return &JobEventsJobEventsListOK{}
}

/*JobEventsJobEventsListOK handles this case with default header values.

OK
*/
type JobEventsJobEventsListOK struct {
}

func (o *JobEventsJobEventsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/job_events/][%d] jobEventsJobEventsListOK ", 200)
}

func (o *JobEventsJobEventsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
