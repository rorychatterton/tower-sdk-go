// Code generated by go-swagger; DO NOT EDIT.

package job_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobEventsJobEventsHostsListReader is a Reader for the JobEventsJobEventsHostsList structure.
type JobEventsJobEventsHostsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobEventsJobEventsHostsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobEventsJobEventsHostsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobEventsJobEventsHostsListOK creates a JobEventsJobEventsHostsListOK with default headers values
func NewJobEventsJobEventsHostsListOK() *JobEventsJobEventsHostsListOK {
	return &JobEventsJobEventsHostsListOK{}
}

/*JobEventsJobEventsHostsListOK handles this case with default header values.

OK
*/
type JobEventsJobEventsHostsListOK struct {
}

func (o *JobEventsJobEventsHostsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/job_events/{id}/hosts/][%d] jobEventsJobEventsHostsListOK ", 200)
}

func (o *JobEventsJobEventsHostsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}