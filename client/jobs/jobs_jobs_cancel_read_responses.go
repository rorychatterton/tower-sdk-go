// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobsJobsCancelReadReader is a Reader for the JobsJobsCancelRead structure.
type JobsJobsCancelReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobsJobsCancelReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobsJobsCancelReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobsJobsCancelReadOK creates a JobsJobsCancelReadOK with default headers values
func NewJobsJobsCancelReadOK() *JobsJobsCancelReadOK {
	return &JobsJobsCancelReadOK{}
}

/*JobsJobsCancelReadOK handles this case with default header values.

OK
*/
type JobsJobsCancelReadOK struct {
}

func (o *JobsJobsCancelReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/jobs/{id}/cancel/][%d] jobsJobsCancelReadOK ", 200)
}

func (o *JobsJobsCancelReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
