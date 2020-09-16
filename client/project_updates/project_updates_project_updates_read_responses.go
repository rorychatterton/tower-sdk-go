// Code generated by go-swagger; DO NOT EDIT.

package project_updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectUpdatesProjectUpdatesReadReader is a Reader for the ProjectUpdatesProjectUpdatesRead structure.
type ProjectUpdatesProjectUpdatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectUpdatesProjectUpdatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectUpdatesProjectUpdatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectUpdatesProjectUpdatesReadOK creates a ProjectUpdatesProjectUpdatesReadOK with default headers values
func NewProjectUpdatesProjectUpdatesReadOK() *ProjectUpdatesProjectUpdatesReadOK {
	return &ProjectUpdatesProjectUpdatesReadOK{}
}

/*ProjectUpdatesProjectUpdatesReadOK handles this case with default header values.

OK
*/
type ProjectUpdatesProjectUpdatesReadOK struct {
}

func (o *ProjectUpdatesProjectUpdatesReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/project_updates/{id}/][%d] projectUpdatesProjectUpdatesReadOK ", 200)
}

func (o *ProjectUpdatesProjectUpdatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
