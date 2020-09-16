// Code generated by go-swagger; DO NOT EDIT.

package custom_inventory_scripts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new custom inventory scripts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for custom inventory scripts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CustomInventoryScriptsInventoryScriptsCopyCreate(params *CustomInventoryScriptsInventoryScriptsCopyCreateParams) (*CustomInventoryScriptsInventoryScriptsCopyCreateCreated, error)

	CustomInventoryScriptsInventoryScriptsCopyList(params *CustomInventoryScriptsInventoryScriptsCopyListParams) (*CustomInventoryScriptsInventoryScriptsCopyListOK, error)

	CustomInventoryScriptsInventoryScriptsCreate(params *CustomInventoryScriptsInventoryScriptsCreateParams) (*CustomInventoryScriptsInventoryScriptsCreateCreated, error)

	CustomInventoryScriptsInventoryScriptsDelete(params *CustomInventoryScriptsInventoryScriptsDeleteParams) (*CustomInventoryScriptsInventoryScriptsDeleteNoContent, error)

	CustomInventoryScriptsInventoryScriptsList(params *CustomInventoryScriptsInventoryScriptsListParams) (*CustomInventoryScriptsInventoryScriptsListOK, error)

	CustomInventoryScriptsInventoryScriptsObjectRolesList(params *CustomInventoryScriptsInventoryScriptsObjectRolesListParams) (*CustomInventoryScriptsInventoryScriptsObjectRolesListOK, error)

	CustomInventoryScriptsInventoryScriptsPartialUpdate(params *CustomInventoryScriptsInventoryScriptsPartialUpdateParams) (*CustomInventoryScriptsInventoryScriptsPartialUpdateOK, error)

	CustomInventoryScriptsInventoryScriptsRead(params *CustomInventoryScriptsInventoryScriptsReadParams) (*CustomInventoryScriptsInventoryScriptsReadOK, error)

	CustomInventoryScriptsInventoryScriptsUpdate(params *CustomInventoryScriptsInventoryScriptsUpdateParams) (*CustomInventoryScriptsInventoryScriptsUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CustomInventoryScriptsInventoryScriptsCopyCreate custom inventory scripts inventory scripts copy create API
*/
func (a *Client) CustomInventoryScriptsInventoryScriptsCopyCreate(params *CustomInventoryScriptsInventoryScriptsCopyCreateParams) (*CustomInventoryScriptsInventoryScriptsCopyCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomInventoryScriptsInventoryScriptsCopyCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Custom Inventory Scripts_inventory_scripts_copy_create",
		Method:             "POST",
		PathPattern:        "/api/v2/inventory_scripts/{id}/copy/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomInventoryScriptsInventoryScriptsCopyCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomInventoryScriptsInventoryScriptsCopyCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Custom Inventory Scripts_inventory_scripts_copy_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CustomInventoryScriptsInventoryScriptsCopyList custom inventory scripts inventory scripts copy list API
*/
func (a *Client) CustomInventoryScriptsInventoryScriptsCopyList(params *CustomInventoryScriptsInventoryScriptsCopyListParams) (*CustomInventoryScriptsInventoryScriptsCopyListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomInventoryScriptsInventoryScriptsCopyListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Custom Inventory Scripts_inventory_scripts_copy_list",
		Method:             "GET",
		PathPattern:        "/api/v2/inventory_scripts/{id}/copy/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomInventoryScriptsInventoryScriptsCopyListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomInventoryScriptsInventoryScriptsCopyListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Custom Inventory Scripts_inventory_scripts_copy_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CustomInventoryScriptsInventoryScriptsCreate creates a custom inventory script


Make a POST request to this resource with the following custom inventory script
fields to create a new custom inventory script:









* `name`: Name of this custom inventory script. (string, required)
* `description`: Optional description of this custom inventory script. (string, default=`""`)
* `script`:  (string, required)
* `organization`: Organization owning this inventory script (id, required)
*/
func (a *Client) CustomInventoryScriptsInventoryScriptsCreate(params *CustomInventoryScriptsInventoryScriptsCreateParams) (*CustomInventoryScriptsInventoryScriptsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomInventoryScriptsInventoryScriptsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Custom Inventory Scripts_inventory_scripts_create",
		Method:             "POST",
		PathPattern:        "/api/v2/inventory_scripts/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomInventoryScriptsInventoryScriptsCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomInventoryScriptsInventoryScriptsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Custom Inventory Scripts_inventory_scripts_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CustomInventoryScriptsInventoryScriptsDelete deletes a custom inventory script


Make a DELETE request to this resource to delete this custom inventory script.
*/
func (a *Client) CustomInventoryScriptsInventoryScriptsDelete(params *CustomInventoryScriptsInventoryScriptsDeleteParams) (*CustomInventoryScriptsInventoryScriptsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomInventoryScriptsInventoryScriptsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Custom Inventory Scripts_inventory_scripts_delete",
		Method:             "DELETE",
		PathPattern:        "/api/v2/inventory_scripts/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomInventoryScriptsInventoryScriptsDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomInventoryScriptsInventoryScriptsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Custom Inventory Scripts_inventory_scripts_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CustomInventoryScriptsInventoryScriptsList lists custom inventory scripts


Make a GET request to this resource to retrieve the list of
custom inventory scripts.

The resulting data structure contains:

    {
        "count": 99,
        "next": null,
        "previous": null,
        "results": [
            ...
        ]
    }

The `count` field indicates the total number of custom inventory scripts
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more custom inventory script records.

## Results

Each custom inventory script data structure includes the following fields:

* `id`: Database ID for this custom inventory script. (integer)
* `type`: Data type for this custom inventory script. (choice)
* `url`: URL for this custom inventory script. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this custom inventory script was created. (datetime)
* `modified`: Timestamp when this custom inventory script was last modified. (datetime)
* `name`: Name of this custom inventory script. (string)
* `description`: Optional description of this custom inventory script. (string)
* `script`:  (string)
* `organization`: Organization owning this inventory script (id)



## Sorting

To specify that custom inventory scripts are returned in a particular
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
func (a *Client) CustomInventoryScriptsInventoryScriptsList(params *CustomInventoryScriptsInventoryScriptsListParams) (*CustomInventoryScriptsInventoryScriptsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomInventoryScriptsInventoryScriptsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Custom Inventory Scripts_inventory_scripts_list",
		Method:             "GET",
		PathPattern:        "/api/v2/inventory_scripts/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomInventoryScriptsInventoryScriptsListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomInventoryScriptsInventoryScriptsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Custom Inventory Scripts_inventory_scripts_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CustomInventoryScriptsInventoryScriptsObjectRolesList lists roles for a custom inventory script


Make a GET request to this resource to retrieve a list of
roles associated with the selected
custom inventory script.

The resulting data structure contains:

    {
        "count": 99,
        "next": null,
        "previous": null,
        "results": [
            ...
        ]
    }

The `count` field indicates the total number of roles
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more role records.

## Results

Each role data structure includes the following fields:

* `id`: Database ID for this role. (integer)
* `type`: Data type for this role. (choice)
* `url`: URL for this role. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `name`: Name of this role. (field)
* `description`: Optional description of this role. (field)



## Sorting

To specify that roles are returned in a particular
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
func (a *Client) CustomInventoryScriptsInventoryScriptsObjectRolesList(params *CustomInventoryScriptsInventoryScriptsObjectRolesListParams) (*CustomInventoryScriptsInventoryScriptsObjectRolesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomInventoryScriptsInventoryScriptsObjectRolesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Custom Inventory Scripts_inventory_scripts_object_roles_list",
		Method:             "GET",
		PathPattern:        "/api/v2/inventory_scripts/{id}/object_roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomInventoryScriptsInventoryScriptsObjectRolesListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomInventoryScriptsInventoryScriptsObjectRolesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Custom Inventory Scripts_inventory_scripts_object_roles_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CustomInventoryScriptsInventoryScriptsPartialUpdate updates a custom inventory script


Make a PUT or PATCH request to this resource to update this
custom inventory script.  The following fields may be modified:









* `name`: Name of this custom inventory script. (string, required)
* `description`: Optional description of this custom inventory script. (string, default=`""`)
* `script`:  (string, required)
* `organization`: Organization owning this inventory script (id, required)








For a PATCH request, include only the fields that are being modified.
*/
func (a *Client) CustomInventoryScriptsInventoryScriptsPartialUpdate(params *CustomInventoryScriptsInventoryScriptsPartialUpdateParams) (*CustomInventoryScriptsInventoryScriptsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomInventoryScriptsInventoryScriptsPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Custom Inventory Scripts_inventory_scripts_partial_update",
		Method:             "PATCH",
		PathPattern:        "/api/v2/inventory_scripts/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomInventoryScriptsInventoryScriptsPartialUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomInventoryScriptsInventoryScriptsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Custom Inventory Scripts_inventory_scripts_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CustomInventoryScriptsInventoryScriptsRead retrieves a custom inventory script


Make GET request to this resource to retrieve a single custom inventory script
record containing the following fields:

* `id`: Database ID for this custom inventory script. (integer)
* `type`: Data type for this custom inventory script. (choice)
* `url`: URL for this custom inventory script. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this custom inventory script was created. (datetime)
* `modified`: Timestamp when this custom inventory script was last modified. (datetime)
* `name`: Name of this custom inventory script. (string)
* `description`: Optional description of this custom inventory script. (string)
* `script`:  (string)
* `organization`: Organization owning this inventory script (id)
*/
func (a *Client) CustomInventoryScriptsInventoryScriptsRead(params *CustomInventoryScriptsInventoryScriptsReadParams) (*CustomInventoryScriptsInventoryScriptsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomInventoryScriptsInventoryScriptsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Custom Inventory Scripts_inventory_scripts_read",
		Method:             "GET",
		PathPattern:        "/api/v2/inventory_scripts/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomInventoryScriptsInventoryScriptsReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomInventoryScriptsInventoryScriptsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Custom Inventory Scripts_inventory_scripts_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CustomInventoryScriptsInventoryScriptsUpdate updates a custom inventory script


Make a PUT or PATCH request to this resource to update this
custom inventory script.  The following fields may be modified:









* `name`: Name of this custom inventory script. (string, required)
* `description`: Optional description of this custom inventory script. (string, default=`""`)
* `script`:  (string, required)
* `organization`: Organization owning this inventory script (id, required)






For a PUT request, include **all** fields in the request.
*/
func (a *Client) CustomInventoryScriptsInventoryScriptsUpdate(params *CustomInventoryScriptsInventoryScriptsUpdateParams) (*CustomInventoryScriptsInventoryScriptsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomInventoryScriptsInventoryScriptsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Custom Inventory Scripts_inventory_scripts_update",
		Method:             "PUT",
		PathPattern:        "/api/v2/inventory_scripts/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomInventoryScriptsInventoryScriptsUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomInventoryScriptsInventoryScriptsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Custom Inventory Scripts_inventory_scripts_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}