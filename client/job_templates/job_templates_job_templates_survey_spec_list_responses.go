// Code generated by go-swagger; DO NOT EDIT.

package job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobTemplatesJobTemplatesSurveySpecListReader is a Reader for the JobTemplatesJobTemplatesSurveySpecList structure.
type JobTemplatesJobTemplatesSurveySpecListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobTemplatesJobTemplatesSurveySpecListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobTemplatesJobTemplatesSurveySpecListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJobTemplatesJobTemplatesSurveySpecListOK creates a JobTemplatesJobTemplatesSurveySpecListOK with default headers values
func NewJobTemplatesJobTemplatesSurveySpecListOK() *JobTemplatesJobTemplatesSurveySpecListOK {
	return &JobTemplatesJobTemplatesSurveySpecListOK{}
}

/*JobTemplatesJobTemplatesSurveySpecListOK handles this case with default header values.

OK
*/
type JobTemplatesJobTemplatesSurveySpecListOK struct {
}

func (o *JobTemplatesJobTemplatesSurveySpecListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/job_templates/{id}/survey_spec/][%d] jobTemplatesJobTemplatesSurveySpecListOK ", 200)
}

func (o *JobTemplatesJobTemplatesSurveySpecListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}