// Code generated by go-swagger; DO NOT EDIT.

package credential_input_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CredentialInputSourcesCredentialInputSourcesPartialUpdateReader is a Reader for the CredentialInputSourcesCredentialInputSourcesPartialUpdate structure.
type CredentialInputSourcesCredentialInputSourcesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCredentialInputSourcesCredentialInputSourcesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCredentialInputSourcesCredentialInputSourcesPartialUpdateOK creates a CredentialInputSourcesCredentialInputSourcesPartialUpdateOK with default headers values
func NewCredentialInputSourcesCredentialInputSourcesPartialUpdateOK() *CredentialInputSourcesCredentialInputSourcesPartialUpdateOK {
	return &CredentialInputSourcesCredentialInputSourcesPartialUpdateOK{}
}

/*CredentialInputSourcesCredentialInputSourcesPartialUpdateOK handles this case with default header values.

OK
*/
type CredentialInputSourcesCredentialInputSourcesPartialUpdateOK struct {
}

func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/credential_input_sources/{id}/][%d] credentialInputSourcesCredentialInputSourcesPartialUpdateOK ", 200)
}

func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden creates a CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden with default headers values
func NewCredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden() *CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden {
	return &CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden{}
}

/*CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden handles this case with default header values.

No Permission Response
*/
type CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden struct {
}

func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/credential_input_sources/{id}/][%d] credentialInputSourcesCredentialInputSourcesPartialUpdateForbidden ", 403)
}

func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
