// Code generated by go-swagger; DO NOT EDIT.

package schedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SchedulesSchedulesCreateReader is a Reader for the SchedulesSchedulesCreate structure.
type SchedulesSchedulesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchedulesSchedulesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSchedulesSchedulesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSchedulesSchedulesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSchedulesSchedulesCreateCreated creates a SchedulesSchedulesCreateCreated with default headers values
func NewSchedulesSchedulesCreateCreated() *SchedulesSchedulesCreateCreated {
	return &SchedulesSchedulesCreateCreated{}
}

/*SchedulesSchedulesCreateCreated handles this case with default header values.

SchedulesSchedulesCreateCreated schedules schedules create created
*/
type SchedulesSchedulesCreateCreated struct {
}

func (o *SchedulesSchedulesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/schedules/][%d] schedulesSchedulesCreateCreated ", 201)
}

func (o *SchedulesSchedulesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchedulesSchedulesCreateForbidden creates a SchedulesSchedulesCreateForbidden with default headers values
func NewSchedulesSchedulesCreateForbidden() *SchedulesSchedulesCreateForbidden {
	return &SchedulesSchedulesCreateForbidden{}
}

/*SchedulesSchedulesCreateForbidden handles this case with default header values.

No Permission Response
*/
type SchedulesSchedulesCreateForbidden struct {
}

func (o *SchedulesSchedulesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/schedules/][%d] schedulesSchedulesCreateForbidden ", 403)
}

func (o *SchedulesSchedulesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
