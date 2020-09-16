// Code generated by go-swagger; DO NOT EDIT.

package unified_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new unified jobs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for unified jobs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	UnifiedJobsUnifiedJobsList(params *UnifiedJobsUnifiedJobsListParams) (*UnifiedJobsUnifiedJobsListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  UnifiedJobsUnifiedJobsList lists unified jobs


Make a GET request to this resource to retrieve the list of
unified jobs.

The resulting data structure contains:

    {
        "count": 99,
        "next": null,
        "previous": null,
        "results": [
            ...
        ]
    }

The `count` field indicates the total number of unified jobs
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more unified job records.

## Results

Each unified job data structure includes the following fields:

* `id`: Database ID for this unified job. (integer)
* `type`: Data type for this unified job. (choice)
* `url`: URL for this unified job. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this unified job was created. (datetime)
* `modified`: Timestamp when this unified job was last modified. (datetime)
* `name`: Name of this unified job. (string)
* `description`: Optional description of this unified job. (string)
* `unified_job_template`:  (id)
* `launch_type`:  (choice)
    - `manual`: Manual
    - `relaunch`: Relaunch
    - `callback`: Callback
    - `scheduled`: Scheduled
    - `dependency`: Dependency
    - `workflow`: Workflow
    - `webhook`: Webhook
    - `sync`: Sync
    - `scm`: SCM Update
* `status`:  (choice)
    - `new`: New
    - `pending`: Pending
    - `waiting`: Waiting
    - `running`: Running
    - `successful`: Successful
    - `failed`: Failed
    - `error`: Error
    - `canceled`: Canceled
* `failed`:  (boolean)
* `started`: The date and time the job was queued for starting. (datetime)
* `finished`: The date and time the job finished execution. (datetime)
* `canceled_on`: The date and time when the cancel request was sent. (datetime)
* `elapsed`: Elapsed time in seconds that the job ran. (decimal)
* `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string)
* `execution_node`: The node the job executed on. (string)
* `controller_node`: The instance that managed the isolated execution environment. (string)



## Sorting

To specify that unified jobs are returned in a particular
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
func (a *Client) UnifiedJobsUnifiedJobsList(params *UnifiedJobsUnifiedJobsListParams) (*UnifiedJobsUnifiedJobsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnifiedJobsUnifiedJobsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Unified Jobs_unified_jobs_list",
		Method:             "GET",
		PathPattern:        "/api/v2/unified_jobs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnifiedJobsUnifiedJobsListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnifiedJobsUnifiedJobsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Unified Jobs_unified_jobs_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
