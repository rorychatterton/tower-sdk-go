// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CredentialsCredentialsInputSourcesListReader is a Reader for the CredentialsCredentialsInputSourcesList structure.
type CredentialsCredentialsInputSourcesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialsCredentialsInputSourcesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCredentialsCredentialsInputSourcesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCredentialsCredentialsInputSourcesListOK creates a CredentialsCredentialsInputSourcesListOK with default headers values
func NewCredentialsCredentialsInputSourcesListOK() *CredentialsCredentialsInputSourcesListOK {
	return &CredentialsCredentialsInputSourcesListOK{}
}

/*CredentialsCredentialsInputSourcesListOK handles this case with default header values.

OK
*/
type CredentialsCredentialsInputSourcesListOK struct {
}

func (o *CredentialsCredentialsInputSourcesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/credentials/{id}/input_sources/][%d] credentialsCredentialsInputSourcesListOK ", 200)
}

func (o *CredentialsCredentialsInputSourcesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
