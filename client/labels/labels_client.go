// Code generated by go-swagger; DO NOT EDIT.

package labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new labels API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for labels API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	LabelsLabelsCreate(params *LabelsLabelsCreateParams) (*LabelsLabelsCreateCreated, error)

	LabelsLabelsList(params *LabelsLabelsListParams) (*LabelsLabelsListOK, error)

	LabelsLabelsPartialUpdate(params *LabelsLabelsPartialUpdateParams) (*LabelsLabelsPartialUpdateOK, error)

	LabelsLabelsRead(params *LabelsLabelsReadParams) (*LabelsLabelsReadOK, error)

	LabelsLabelsUpdate(params *LabelsLabelsUpdateParams) (*LabelsLabelsUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  LabelsLabelsCreate creates a label


Make a POST request to this resource with the following label
fields to create a new label:









* `name`: Name of this label. (string, required)
* `organization`: Organization this label belongs to. (id, required)
*/
func (a *Client) LabelsLabelsCreate(params *LabelsLabelsCreateParams) (*LabelsLabelsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLabelsLabelsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Labels_labels_create",
		Method:             "POST",
		PathPattern:        "/api/v2/labels/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LabelsLabelsCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LabelsLabelsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Labels_labels_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LabelsLabelsList lists labels


Make a GET request to this resource to retrieve the list of
labels.

The resulting data structure contains:

    {
        "count": 99,
        "next": null,
        "previous": null,
        "results": [
            ...
        ]
    }

The `count` field indicates the total number of labels
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more label records.

## Results

Each label data structure includes the following fields:

* `id`: Database ID for this label. (integer)
* `type`: Data type for this label. (choice)
* `url`: URL for this label. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this label was created. (datetime)
* `modified`: Timestamp when this label was last modified. (datetime)
* `name`: Name of this label. (string)
* `organization`: Organization this label belongs to. (id)



## Sorting

To specify that labels are returned in a particular
order, use the `order_by` query string parameter on the GET request.

    ?order_by=name

Prefix the field name with a dash `-` to sort in reverse:

    ?order_by=-name

Multiple sorting fields may be specified by separating the field names with a
comma `,`:

    ?order_by=name,some_other_field

## Pagination

Use the `page_size` query string parameter to change the number of results
returned for each request.  Use the `page` query string parameter to retrieve
a particular page of results.

    ?page_size=100&page=2

The `previous` and `next` links returned with the results will set these query
string parameters automatically.

## Searching

Use the `search` query string parameter to perform a case-insensitive search
within all designated text fields of a model.

    ?search=findme

(_Added in Ansible Tower 3.1.0_) Search across related fields:

    ?related__search=findme
*/
func (a *Client) LabelsLabelsList(params *LabelsLabelsListParams) (*LabelsLabelsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLabelsLabelsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Labels_labels_list",
		Method:             "GET",
		PathPattern:        "/api/v2/labels/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LabelsLabelsListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LabelsLabelsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Labels_labels_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LabelsLabelsPartialUpdate updates a label


Make a PUT or PATCH request to this resource to update this
label.  The following fields may be modified:









* `name`: Name of this label. (string, required)
* `organization`: Organization this label belongs to. (id, required)








For a PATCH request, include only the fields that are being modified.
*/
func (a *Client) LabelsLabelsPartialUpdate(params *LabelsLabelsPartialUpdateParams) (*LabelsLabelsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLabelsLabelsPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Labels_labels_partial_update",
		Method:             "PATCH",
		PathPattern:        "/api/v2/labels/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LabelsLabelsPartialUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LabelsLabelsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Labels_labels_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LabelsLabelsRead retrieves a label


Make GET request to this resource to retrieve a single label
record containing the following fields:

* `id`: Database ID for this label. (integer)
* `type`: Data type for this label. (choice)
* `url`: URL for this label. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this label was created. (datetime)
* `modified`: Timestamp when this label was last modified. (datetime)
* `name`: Name of this label. (string)
* `organization`: Organization this label belongs to. (id)
*/
func (a *Client) LabelsLabelsRead(params *LabelsLabelsReadParams) (*LabelsLabelsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLabelsLabelsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Labels_labels_read",
		Method:             "GET",
		PathPattern:        "/api/v2/labels/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LabelsLabelsReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LabelsLabelsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Labels_labels_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LabelsLabelsUpdate updates a label


Make a PUT or PATCH request to this resource to update this
label.  The following fields may be modified:









* `name`: Name of this label. (string, required)
* `organization`: Organization this label belongs to. (id, required)






For a PUT request, include **all** fields in the request.
*/
func (a *Client) LabelsLabelsUpdate(params *LabelsLabelsUpdateParams) (*LabelsLabelsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLabelsLabelsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Labels_labels_update",
		Method:             "PUT",
		PathPattern:        "/api/v2/labels/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LabelsLabelsUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LabelsLabelsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Labels_labels_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
