// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CredentialsCredentialsTestCreateReader is a Reader for the CredentialsCredentialsTestCreate structure.
type CredentialsCredentialsTestCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialsCredentialsTestCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCredentialsCredentialsTestCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCredentialsCredentialsTestCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCredentialsCredentialsTestCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCredentialsCredentialsTestCreateCreated creates a CredentialsCredentialsTestCreateCreated with default headers values
func NewCredentialsCredentialsTestCreateCreated() *CredentialsCredentialsTestCreateCreated {
	return &CredentialsCredentialsTestCreateCreated{}
}

/*CredentialsCredentialsTestCreateCreated handles this case with default header values.

CredentialsCredentialsTestCreateCreated credentials credentials test create created
*/
type CredentialsCredentialsTestCreateCreated struct {
}

func (o *CredentialsCredentialsTestCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/credentials/{id}/test/][%d] credentialsCredentialsTestCreateCreated ", 201)
}

func (o *CredentialsCredentialsTestCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCredentialsCredentialsTestCreateAccepted creates a CredentialsCredentialsTestCreateAccepted with default headers values
func NewCredentialsCredentialsTestCreateAccepted() *CredentialsCredentialsTestCreateAccepted {
	return &CredentialsCredentialsTestCreateAccepted{}
}

/*CredentialsCredentialsTestCreateAccepted handles this case with default header values.

Accepted
*/
type CredentialsCredentialsTestCreateAccepted struct {
}

func (o *CredentialsCredentialsTestCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/credentials/{id}/test/][%d] credentialsCredentialsTestCreateAccepted ", 202)
}

func (o *CredentialsCredentialsTestCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCredentialsCredentialsTestCreateForbidden creates a CredentialsCredentialsTestCreateForbidden with default headers values
func NewCredentialsCredentialsTestCreateForbidden() *CredentialsCredentialsTestCreateForbidden {
	return &CredentialsCredentialsTestCreateForbidden{}
}

/*CredentialsCredentialsTestCreateForbidden handles this case with default header values.

No Permission Response
*/
type CredentialsCredentialsTestCreateForbidden struct {
}

func (o *CredentialsCredentialsTestCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/credentials/{id}/test/][%d] credentialsCredentialsTestCreateForbidden ", 403)
}

func (o *CredentialsCredentialsTestCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
