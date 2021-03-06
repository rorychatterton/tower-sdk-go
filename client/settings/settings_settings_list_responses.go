// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SettingsSettingsListReader is a Reader for the SettingsSettingsList structure.
type SettingsSettingsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SettingsSettingsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSettingsSettingsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSettingsSettingsListOK creates a SettingsSettingsListOK with default headers values
func NewSettingsSettingsListOK() *SettingsSettingsListOK {
	return &SettingsSettingsListOK{}
}

/*SettingsSettingsListOK handles this case with default header values.

OK
*/
type SettingsSettingsListOK struct {
}

func (o *SettingsSettingsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/settings/][%d] settingsSettingsListOK ", 200)
}

func (o *SettingsSettingsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
