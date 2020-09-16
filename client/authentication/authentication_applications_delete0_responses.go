// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AuthenticationApplicationsDelete0Reader is a Reader for the AuthenticationApplicationsDelete0 structure.
type AuthenticationApplicationsDelete0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthenticationApplicationsDelete0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAuthenticationApplicationsDelete0NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAuthenticationApplicationsDelete0NoContent creates a AuthenticationApplicationsDelete0NoContent with default headers values
func NewAuthenticationApplicationsDelete0NoContent() *AuthenticationApplicationsDelete0NoContent {
	return &AuthenticationApplicationsDelete0NoContent{}
}

/*AuthenticationApplicationsDelete0NoContent handles this case with default header values.

AuthenticationApplicationsDelete0NoContent authentication applications delete0 no content
*/
type AuthenticationApplicationsDelete0NoContent struct {
}

func (o *AuthenticationApplicationsDelete0NoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/applications/{id}/][%d] authenticationApplicationsDelete0NoContent ", 204)
}

func (o *AuthenticationApplicationsDelete0NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
