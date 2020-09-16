// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TeamsTeamsPartialUpdateReader is a Reader for the TeamsTeamsPartialUpdate structure.
type TeamsTeamsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamsTeamsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTeamsTeamsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamsTeamsPartialUpdateOK creates a TeamsTeamsPartialUpdateOK with default headers values
func NewTeamsTeamsPartialUpdateOK() *TeamsTeamsPartialUpdateOK {
	return &TeamsTeamsPartialUpdateOK{}
}

/*TeamsTeamsPartialUpdateOK handles this case with default header values.

OK
*/
type TeamsTeamsPartialUpdateOK struct {
}

func (o *TeamsTeamsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/teams/{id}/][%d] teamsTeamsPartialUpdateOK ", 200)
}

func (o *TeamsTeamsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*TeamsTeamsPartialUpdateBody teams teams partial update body
swagger:model TeamsTeamsPartialUpdateBody
*/
type TeamsTeamsPartialUpdateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization
	Organization int64 `json:"organization,omitempty"`
}

// Validate validates this teams teams partial update body
func (o *TeamsTeamsPartialUpdateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TeamsTeamsPartialUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TeamsTeamsPartialUpdateBody) UnmarshalBinary(b []byte) error {
	var res TeamsTeamsPartialUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
