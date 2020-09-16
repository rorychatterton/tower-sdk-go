// Code generated by go-swagger; DO NOT EDIT.

package versioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VersioningReadReader is a Reader for the VersioningRead structure.
type VersioningReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VersioningReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVersioningReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVersioningReadOK creates a VersioningReadOK with default headers values
func NewVersioningReadOK() *VersioningReadOK {
	return &VersioningReadOK{}
}

/*VersioningReadOK handles this case with default header values.

OK
*/
type VersioningReadOK struct {
}

func (o *VersioningReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/][%d] versioningReadOK ", 200)
}

func (o *VersioningReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
