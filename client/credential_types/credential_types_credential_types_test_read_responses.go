// Code generated by go-swagger; DO NOT EDIT.

package credential_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CredentialTypesCredentialTypesTestReadReader is a Reader for the CredentialTypesCredentialTypesTestRead structure.
type CredentialTypesCredentialTypesTestReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialTypesCredentialTypesTestReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCredentialTypesCredentialTypesTestReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCredentialTypesCredentialTypesTestReadOK creates a CredentialTypesCredentialTypesTestReadOK with default headers values
func NewCredentialTypesCredentialTypesTestReadOK() *CredentialTypesCredentialTypesTestReadOK {
	return &CredentialTypesCredentialTypesTestReadOK{}
}

/*CredentialTypesCredentialTypesTestReadOK handles this case with default header values.

OK
*/
type CredentialTypesCredentialTypesTestReadOK struct {
}

func (o *CredentialTypesCredentialTypesTestReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/credential_types/{id}/test/][%d] credentialTypesCredentialTypesTestReadOK ", 200)
}

func (o *CredentialTypesCredentialTypesTestReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
