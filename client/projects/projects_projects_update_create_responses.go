// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectsProjectsUpdateCreateReader is a Reader for the ProjectsProjectsUpdateCreate structure.
type ProjectsProjectsUpdateCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsProjectsUpdateCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewProjectsProjectsUpdateCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsProjectsUpdateCreateCreated creates a ProjectsProjectsUpdateCreateCreated with default headers values
func NewProjectsProjectsUpdateCreateCreated() *ProjectsProjectsUpdateCreateCreated {
	return &ProjectsProjectsUpdateCreateCreated{}
}

/*ProjectsProjectsUpdateCreateCreated handles this case with default header values.

ProjectsProjectsUpdateCreateCreated projects projects update create created
*/
type ProjectsProjectsUpdateCreateCreated struct {
}

func (o *ProjectsProjectsUpdateCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{id}/update/][%d] projectsProjectsUpdateCreateCreated ", 201)
}

func (o *ProjectsProjectsUpdateCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
