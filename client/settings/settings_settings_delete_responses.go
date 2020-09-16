// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SettingsSettingsDeleteReader is a Reader for the SettingsSettingsDelete structure.
type SettingsSettingsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SettingsSettingsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSettingsSettingsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSettingsSettingsDeleteNoContent creates a SettingsSettingsDeleteNoContent with default headers values
func NewSettingsSettingsDeleteNoContent() *SettingsSettingsDeleteNoContent {
	return &SettingsSettingsDeleteNoContent{}
}

/*SettingsSettingsDeleteNoContent handles this case with default header values.

SettingsSettingsDeleteNoContent settings settings delete no content
*/
type SettingsSettingsDeleteNoContent struct {
}

func (o *SettingsSettingsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/settings/{category_slug}/][%d] settingsSettingsDeleteNoContent ", 204)
}

func (o *SettingsSettingsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
