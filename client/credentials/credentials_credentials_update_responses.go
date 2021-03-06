// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CredentialsCredentialsUpdateReader is a Reader for the CredentialsCredentialsUpdate structure.
type CredentialsCredentialsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialsCredentialsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCredentialsCredentialsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCredentialsCredentialsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCredentialsCredentialsUpdateOK creates a CredentialsCredentialsUpdateOK with default headers values
func NewCredentialsCredentialsUpdateOK() *CredentialsCredentialsUpdateOK {
	return &CredentialsCredentialsUpdateOK{}
}

/*CredentialsCredentialsUpdateOK handles this case with default header values.

OK
*/
type CredentialsCredentialsUpdateOK struct {
}

func (o *CredentialsCredentialsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/credentials/{id}/][%d] credentialsCredentialsUpdateOK ", 200)
}

func (o *CredentialsCredentialsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCredentialsCredentialsUpdateBadRequest creates a CredentialsCredentialsUpdateBadRequest with default headers values
func NewCredentialsCredentialsUpdateBadRequest() *CredentialsCredentialsUpdateBadRequest {
	return &CredentialsCredentialsUpdateBadRequest{}
}

/*CredentialsCredentialsUpdateBadRequest handles this case with default header values.

Bad Request
*/
type CredentialsCredentialsUpdateBadRequest struct {
}

func (o *CredentialsCredentialsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/credentials/{id}/][%d] credentialsCredentialsUpdateBadRequest ", 400)
}

func (o *CredentialsCredentialsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
