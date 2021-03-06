// Code generated by go-swagger; DO NOT EDIT.

package job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobTemplatesJobTemplatesSurveySpecDeleteReader is a Reader for the JobTemplatesJobTemplatesSurveySpecDelete structure.
type JobTemplatesJobTemplatesSurveySpecDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobTemplatesJobTemplatesSurveySpecDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewJobTemplatesJobTemplatesSurveySpecDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobTemplatesJobTemplatesSurveySpecDeleteNoContent creates a JobTemplatesJobTemplatesSurveySpecDeleteNoContent with default headers values
func NewJobTemplatesJobTemplatesSurveySpecDeleteNoContent() *JobTemplatesJobTemplatesSurveySpecDeleteNoContent {
	return &JobTemplatesJobTemplatesSurveySpecDeleteNoContent{}
}

/*JobTemplatesJobTemplatesSurveySpecDeleteNoContent handles this case with default header values.

JobTemplatesJobTemplatesSurveySpecDeleteNoContent job templates job templates survey spec delete no content
*/
type JobTemplatesJobTemplatesSurveySpecDeleteNoContent struct {
}

func (o *JobTemplatesJobTemplatesSurveySpecDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/job_templates/{id}/survey_spec/][%d] jobTemplatesJobTemplatesSurveySpecDeleteNoContent ", 204)
}

func (o *JobTemplatesJobTemplatesSurveySpecDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
