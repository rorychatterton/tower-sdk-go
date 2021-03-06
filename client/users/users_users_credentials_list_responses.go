// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersUsersCredentialsListReader is a Reader for the UsersUsersCredentialsList structure.
type UsersUsersCredentialsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersCredentialsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersCredentialsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersUsersCredentialsListOK creates a UsersUsersCredentialsListOK with default headers values
func NewUsersUsersCredentialsListOK() *UsersUsersCredentialsListOK {
	return &UsersUsersCredentialsListOK{}
}

/*UsersUsersCredentialsListOK handles this case with default header values.

OK
*/
type UsersUsersCredentialsListOK struct {
}

func (o *UsersUsersCredentialsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{id}/credentials/][%d] usersUsersCredentialsListOK ", 200)
}

func (o *UsersUsersCredentialsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
