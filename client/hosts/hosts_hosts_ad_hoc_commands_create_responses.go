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

// HostsHostsAdHocCommandsCreateReader is a Reader for the HostsHostsAdHocCommandsCreate structure.
type HostsHostsAdHocCommandsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsHostsAdHocCommandsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewHostsHostsAdHocCommandsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHostsHostsAdHocCommandsCreateCreated creates a HostsHostsAdHocCommandsCreateCreated with default headers values
func NewHostsHostsAdHocCommandsCreateCreated() *HostsHostsAdHocCommandsCreateCreated {
	return &HostsHostsAdHocCommandsCreateCreated{}
}

/*HostsHostsAdHocCommandsCreateCreated handles this case with default header values.

HostsHostsAdHocCommandsCreateCreated hosts hosts ad hoc commands create created
*/
type HostsHostsAdHocCommandsCreateCreated struct {
}

func (o *HostsHostsAdHocCommandsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/hosts/{id}/ad_hoc_commands/][%d] hostsHostsAdHocCommandsCreateCreated ", 201)
}

func (o *HostsHostsAdHocCommandsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*HostsHostsAdHocCommandsCreateBody hosts hosts ad hoc commands create body
swagger:model HostsHostsAdHocCommandsCreateBody
*/
type HostsHostsAdHocCommandsCreateBody struct {

	// become enabled
	BecomeEnabled string `json:"become_enabled,omitempty"`

	// credential
	Credential int64 `json:"credential,omitempty"`

	// diff mode
	DiffMode string `json:"diff_mode,omitempty"`

	// extra vars
	ExtraVars string `json:"extra_vars,omitempty"`

	// forks
	Forks int64 `json:"forks,omitempty"`

	// inventory
	Inventory int64 `json:"inventory,omitempty"`

	// job type
	JobType string `json:"job_type,omitempty"`

	// limit
	Limit string `json:"limit,omitempty"`

	// module args
	ModuleArgs string `json:"module_args,omitempty"`

	// module name
	ModuleName string `json:"module_name,omitempty"`

	// verbosity
	Verbosity string `json:"verbosity,omitempty"`
}

// Validate validates this hosts hosts ad hoc commands create body
func (o *HostsHostsAdHocCommandsCreateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *HostsHostsAdHocCommandsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *HostsHostsAdHocCommandsCreateBody) UnmarshalBinary(b []byte) error {
	var res HostsHostsAdHocCommandsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
