// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CredentialsCredentialsTestReadReader is a Reader for the CredentialsCredentialsTestRead structure.
type CredentialsCredentialsTestReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialsCredentialsTestReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCredentialsCredentialsTestReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCredentialsCredentialsTestReadOK creates a CredentialsCredentialsTestReadOK with default headers values
func NewCredentialsCredentialsTestReadOK() *CredentialsCredentialsTestReadOK {
	return &CredentialsCredentialsTestReadOK{}
}

/*CredentialsCredentialsTestReadOK handles this case with default header values.

OK
*/
type CredentialsCredentialsTestReadOK struct {
}

func (o *CredentialsCredentialsTestReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/credentials/{id}/test/][%d] credentialsCredentialsTestReadOK ", 200)
}

func (o *CredentialsCredentialsTestReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}