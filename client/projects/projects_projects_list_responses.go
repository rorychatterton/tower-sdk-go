// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectsProjectsListReader is a Reader for the ProjectsProjectsList structure.
type ProjectsProjectsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsProjectsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsProjectsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsProjectsListOK creates a ProjectsProjectsListOK with default headers values
func NewProjectsProjectsListOK() *ProjectsProjectsListOK {
	return &ProjectsProjectsListOK{}
}

/*ProjectsProjectsListOK handles this case with default header values.

OK
*/
type ProjectsProjectsListOK struct {
}

func (o *ProjectsProjectsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/][%d] projectsProjectsListOK ", 200)
}

func (o *ProjectsProjectsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
