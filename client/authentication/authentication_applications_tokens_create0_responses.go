// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AuthenticationApplicationsTokensCreate0Reader is a Reader for the AuthenticationApplicationsTokensCreate0 structure.
type AuthenticationApplicationsTokensCreate0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthenticationApplicationsTokensCreate0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAuthenticationApplicationsTokensCreate0Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAuthenticationApplicationsTokensCreate0Created creates a AuthenticationApplicationsTokensCreate0Created with default headers values
func NewAuthenticationApplicationsTokensCreate0Created() *AuthenticationApplicationsTokensCreate0Created {
	return &AuthenticationApplicationsTokensCreate0Created{}
}

/*AuthenticationApplicationsTokensCreate0Created handles this case with default header values.

AuthenticationApplicationsTokensCreate0Created authentication applications tokens create0 created
*/
type AuthenticationApplicationsTokensCreate0Created struct {
}

func (o *AuthenticationApplicationsTokensCreate0Created) Error() string {
	return fmt.Sprintf("[POST /api/v2/applications/{id}/tokens/][%d] authenticationApplicationsTokensCreate0Created ", 201)
}

func (o *AuthenticationApplicationsTokensCreate0Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
