// Code generated by go-swagger; DO NOT EDIT.

package job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobTemplatesJobTemplatesGithubCreateReader is a Reader for the JobTemplatesJobTemplatesGithubCreate structure.
type JobTemplatesJobTemplatesGithubCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobTemplatesJobTemplatesGithubCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewJobTemplatesJobTemplatesGithubCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobTemplatesJobTemplatesGithubCreateCreated creates a JobTemplatesJobTemplatesGithubCreateCreated with default headers values
func NewJobTemplatesJobTemplatesGithubCreateCreated() *JobTemplatesJobTemplatesGithubCreateCreated {
	return &JobTemplatesJobTemplatesGithubCreateCreated{}
}

/*JobTemplatesJobTemplatesGithubCreateCreated handles this case with default header values.

JobTemplatesJobTemplatesGithubCreateCreated job templates job templates github create created
*/
type JobTemplatesJobTemplatesGithubCreateCreated struct {
}

func (o *JobTemplatesJobTemplatesGithubCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/job_templates/{id}/github/][%d] jobTemplatesJobTemplatesGithubCreateCreated ", 201)
}

func (o *JobTemplatesJobTemplatesGithubCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
