// Code generated by go-swagger; DO NOT EDIT.

package project_updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectUpdatesProjectUpdatesStdoutReadReader is a Reader for the ProjectUpdatesProjectUpdatesStdoutRead structure.
type ProjectUpdatesProjectUpdatesStdoutReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectUpdatesProjectUpdatesStdoutReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectUpdatesProjectUpdatesStdoutReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectUpdatesProjectUpdatesStdoutReadOK creates a ProjectUpdatesProjectUpdatesStdoutReadOK with default headers values
func NewProjectUpdatesProjectUpdatesStdoutReadOK() *ProjectUpdatesProjectUpdatesStdoutReadOK {
	return &ProjectUpdatesProjectUpdatesStdoutReadOK{}
}

/*ProjectUpdatesProjectUpdatesStdoutReadOK handles this case with default header values.

OK
*/
type ProjectUpdatesProjectUpdatesStdoutReadOK struct {
}

func (o *ProjectUpdatesProjectUpdatesStdoutReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/project_updates/{id}/stdout/][%d] projectUpdatesProjectUpdatesStdoutReadOK ", 200)
}

func (o *ProjectUpdatesProjectUpdatesStdoutReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
