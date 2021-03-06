// Code generated by go-swagger; DO NOT EDIT.

package system_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SystemJobsSystemJobsCancelCreateReader is a Reader for the SystemJobsSystemJobsCancelCreate structure.
type SystemJobsSystemJobsCancelCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemJobsSystemJobsCancelCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSystemJobsSystemJobsCancelCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemJobsSystemJobsCancelCreateCreated creates a SystemJobsSystemJobsCancelCreateCreated with default headers values
func NewSystemJobsSystemJobsCancelCreateCreated() *SystemJobsSystemJobsCancelCreateCreated {
	return &SystemJobsSystemJobsCancelCreateCreated{}
}

/*SystemJobsSystemJobsCancelCreateCreated handles this case with default header values.

SystemJobsSystemJobsCancelCreateCreated system jobs system jobs cancel create created
*/
type SystemJobsSystemJobsCancelCreateCreated struct {
}

func (o *SystemJobsSystemJobsCancelCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/system_jobs/{id}/cancel/][%d] systemJobsSystemJobsCancelCreateCreated ", 201)
}

func (o *SystemJobsSystemJobsCancelCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
