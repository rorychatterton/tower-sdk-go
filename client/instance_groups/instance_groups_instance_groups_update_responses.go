// Code generated by go-swagger; DO NOT EDIT.

package instance_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InstanceGroupsInstanceGroupsUpdateReader is a Reader for the InstanceGroupsInstanceGroupsUpdate structure.
type InstanceGroupsInstanceGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstanceGroupsInstanceGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstanceGroupsInstanceGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInstanceGroupsInstanceGroupsUpdateOK creates a InstanceGroupsInstanceGroupsUpdateOK with default headers values
func NewInstanceGroupsInstanceGroupsUpdateOK() *InstanceGroupsInstanceGroupsUpdateOK {
	return &InstanceGroupsInstanceGroupsUpdateOK{}
}

/*InstanceGroupsInstanceGroupsUpdateOK handles this case with default header values.

OK
*/
type InstanceGroupsInstanceGroupsUpdateOK struct {
}

func (o *InstanceGroupsInstanceGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/instance_groups/{id}/][%d] instanceGroupsInstanceGroupsUpdateOK ", 200)
}

func (o *InstanceGroupsInstanceGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*InstanceGroupsInstanceGroupsUpdateBody instance groups instance groups update body
swagger:model InstanceGroupsInstanceGroupsUpdateBody
*/
type InstanceGroupsInstanceGroupsUpdateBody struct {

	// credential
	Credential int64 `json:"credential,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// pod spec override
	PodSpecOverride string `json:"pod_spec_override,omitempty"`

	// List of exact-match Instances that will be assigned to this group
	PolicyInstanceList []string `json:"policy_instance_list"`

	// Static minimum number of Instances that will be automatically assign to this group when new instances come online.
	PolicyInstanceMinimum int64 `json:"policy_instance_minimum,omitempty"`

	// Minimum percentage of all instances that will be automatically assigned to this group when new instances come online.
	PolicyInstancePercentage int64 `json:"policy_instance_percentage,omitempty"`
}

// Validate validates this instance groups instance groups update body
func (o *InstanceGroupsInstanceGroupsUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InstanceGroupsInstanceGroupsUpdateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InstanceGroupsInstanceGroupsUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InstanceGroupsInstanceGroupsUpdateBody) UnmarshalBinary(b []byte) error {
	var res InstanceGroupsInstanceGroupsUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
