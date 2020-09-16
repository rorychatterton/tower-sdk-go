// Code generated by go-swagger; DO NOT EDIT.

package project_updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectUpdatesProjectUpdatesScmInventoryUpdatesListReader is a Reader for the ProjectUpdatesProjectUpdatesScmInventoryUpdatesList structure.
type ProjectUpdatesProjectUpdatesScmInventoryUpdatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectUpdatesProjectUpdatesScmInventoryUpdatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK creates a ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK with default headers values
func NewProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK() *ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK {
	return &ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK{}
}

/*ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK handles this case with default header values.

OK
*/
type ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK struct {
}

func (o *ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/project_updates/{id}/scm_inventory_updates/][%d] projectUpdatesProjectUpdatesScmInventoryUpdatesListOK ", 200)
}

func (o *ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}