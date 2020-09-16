// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersMeListReader is a Reader for the UsersMeList structure.
type UsersMeListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersMeListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersMeListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersMeListOK creates a UsersMeListOK with default headers values
func NewUsersMeListOK() *UsersMeListOK {
	return &UsersMeListOK{}
}

/*UsersMeListOK handles this case with default header values.

OK
*/
type UsersMeListOK struct {
}

func (o *UsersMeListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/me/][%d] usersMeListOK ", 200)
}

func (o *UsersMeListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
