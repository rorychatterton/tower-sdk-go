// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AuthenticationApplicationsTokensList0Reader is a Reader for the AuthenticationApplicationsTokensList0 structure.
type AuthenticationApplicationsTokensList0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthenticationApplicationsTokensList0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthenticationApplicationsTokensList0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAuthenticationApplicationsTokensList0OK creates a AuthenticationApplicationsTokensList0OK with default headers values
func NewAuthenticationApplicationsTokensList0OK() *AuthenticationApplicationsTokensList0OK {
	return &AuthenticationApplicationsTokensList0OK{}
}

/*AuthenticationApplicationsTokensList0OK handles this case with default header values.

OK
*/
type AuthenticationApplicationsTokensList0OK struct {
}

func (o *AuthenticationApplicationsTokensList0OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/applications/{id}/tokens/][%d] authenticationApplicationsTokensList0OK ", 200)
}

func (o *AuthenticationApplicationsTokensList0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
