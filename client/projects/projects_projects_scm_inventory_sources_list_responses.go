// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectsProjectsScmInventorySourcesListReader is a Reader for the ProjectsProjectsScmInventorySourcesList structure.
type ProjectsProjectsScmInventorySourcesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsProjectsScmInventorySourcesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsProjectsScmInventorySourcesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsProjectsScmInventorySourcesListOK creates a ProjectsProjectsScmInventorySourcesListOK with default headers values
func NewProjectsProjectsScmInventorySourcesListOK() *ProjectsProjectsScmInventorySourcesListOK {
	return &ProjectsProjectsScmInventorySourcesListOK{}
}

/*ProjectsProjectsScmInventorySourcesListOK handles this case with default header values.

OK
*/
type ProjectsProjectsScmInventorySourcesListOK struct {
}

func (o *ProjectsProjectsScmInventorySourcesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{id}/scm_inventory_sources/][%d] projectsProjectsScmInventorySourcesListOK ", 200)
}

func (o *ProjectsProjectsScmInventorySourcesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}