// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AuthenticationTokensActivityStreamListReader is a Reader for the AuthenticationTokensActivityStreamList structure.
type AuthenticationTokensActivityStreamListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthenticationTokensActivityStreamListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthenticationTokensActivityStreamListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAuthenticationTokensActivityStreamListOK creates a AuthenticationTokensActivityStreamListOK with default headers values
func NewAuthenticationTokensActivityStreamListOK() *AuthenticationTokensActivityStreamListOK {
	return &AuthenticationTokensActivityStreamListOK{}
}

/*AuthenticationTokensActivityStreamListOK handles this case with default header values.

OK
*/
type AuthenticationTokensActivityStreamListOK struct {
}

func (o *AuthenticationTokensActivityStreamListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/tokens/{id}/activity_stream/][%d] authenticationTokensActivityStreamListOK ", 200)
}

func (o *AuthenticationTokensActivityStreamListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}