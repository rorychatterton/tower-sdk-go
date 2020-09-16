// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CredentialsCredentialsOwnerTeamsListReader is a Reader for the CredentialsCredentialsOwnerTeamsList structure.
type CredentialsCredentialsOwnerTeamsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialsCredentialsOwnerTeamsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCredentialsCredentialsOwnerTeamsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCredentialsCredentialsOwnerTeamsListOK creates a CredentialsCredentialsOwnerTeamsListOK with default headers values
func NewCredentialsCredentialsOwnerTeamsListOK() *CredentialsCredentialsOwnerTeamsListOK {
	return &CredentialsCredentialsOwnerTeamsListOK{}
}

/*CredentialsCredentialsOwnerTeamsListOK handles this case with default header values.

OK
*/
type CredentialsCredentialsOwnerTeamsListOK struct {
}

func (o *CredentialsCredentialsOwnerTeamsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/credentials/{id}/owner_teams/][%d] credentialsCredentialsOwnerTeamsListOK ", 200)
}

func (o *CredentialsCredentialsOwnerTeamsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
