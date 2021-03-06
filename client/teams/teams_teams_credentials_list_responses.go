// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TeamsTeamsCredentialsListReader is a Reader for the TeamsTeamsCredentialsList structure.
type TeamsTeamsCredentialsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamsTeamsCredentialsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTeamsTeamsCredentialsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamsTeamsCredentialsListOK creates a TeamsTeamsCredentialsListOK with default headers values
func NewTeamsTeamsCredentialsListOK() *TeamsTeamsCredentialsListOK {
	return &TeamsTeamsCredentialsListOK{}
}

/*TeamsTeamsCredentialsListOK handles this case with default header values.

OK
*/
type TeamsTeamsCredentialsListOK struct {
}

func (o *TeamsTeamsCredentialsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/teams/{id}/credentials/][%d] teamsTeamsCredentialsListOK ", 200)
}

func (o *TeamsTeamsCredentialsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
