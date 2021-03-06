// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HostsHostsPartialUpdateReader is a Reader for the HostsHostsPartialUpdate structure.
type HostsHostsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsHostsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHostsHostsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHostsHostsPartialUpdateOK creates a HostsHostsPartialUpdateOK with default headers values
func NewHostsHostsPartialUpdateOK() *HostsHostsPartialUpdateOK {
	return &HostsHostsPartialUpdateOK{}
}

/*HostsHostsPartialUpdateOK handles this case with default header values.

OK
*/
type HostsHostsPartialUpdateOK struct {
}

func (o *HostsHostsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/hosts/{id}/][%d] hostsHostsPartialUpdateOK ", 200)
}

func (o *HostsHostsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*HostsHostsPartialUpdateBody hosts hosts partial update body
swagger:model HostsHostsPartialUpdateBody
*/
type HostsHostsPartialUpdateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// Is this host online and available for running jobs?
	Enabled string `json:"enabled,omitempty"`

	// The value used by the remote inventory source to uniquely identify the host
	InstanceID string `json:"instance_id,omitempty"`

	// inventory
	Inventory int64 `json:"inventory,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Host variables in JSON or YAML format.
	Variables string `json:"variables,omitempty"`
}

// Validate validates this hosts hosts partial update body
func (o *HostsHostsPartialUpdateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *HostsHostsPartialUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *HostsHostsPartialUpdateBody) UnmarshalBinary(b []byte) error {
	var res HostsHostsPartialUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
