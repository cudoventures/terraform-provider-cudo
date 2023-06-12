// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new projects API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for projects API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateProject(params *CreateProjectParams, opts ...ClientOption) (*CreateProjectOK, error)

	DeleteProject(params *DeleteProjectParams, opts ...ClientOption) (*DeleteProjectOK, error)

	GetProject(params *GetProjectParams, opts ...ClientOption) (*GetProjectOK, error)

	GetProjectCurrentSpend(params *GetProjectCurrentSpendParams, opts ...ClientOption) (*GetProjectCurrentSpendOK, error)

	GetProjectSpendDetails(params *GetProjectSpendDetailsParams, opts ...ClientOption) (*GetProjectSpendDetailsOK, error)

	GetProjectSpendHistory(params *GetProjectSpendHistoryParams, opts ...ClientOption) (*GetProjectSpendHistoryOK, error)

	ListProjectSSHKeys(params *ListProjectSSHKeysParams, opts ...ClientOption) (*ListProjectSSHKeysOK, error)

	ListProjects(params *ListProjectsParams, opts ...ClientOption) (*ListProjectsOK, error)

	UpdateProject(params *UpdateProjectParams, opts ...ClientOption) (*UpdateProjectOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateProject creates
*/
func (a *Client) CreateProject(params *CreateProjectParams, opts ...ClientOption) (*CreateProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateProject",
		Method:             "POST",
		PathPattern:        "/v1/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateProjectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateProjectDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteProject deletes
*/
func (a *Client) DeleteProject(params *DeleteProjectParams, opts ...ClientOption) (*DeleteProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteProject",
		Method:             "DELETE",
		PathPattern:        "/v1/projects/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteProjectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteProjectDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProject gets
*/
func (a *Client) GetProject(params *GetProjectParams, opts ...ClientOption) (*GetProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProject",
		Method:             "GET",
		PathPattern:        "/v1/projects/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProjectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProjectDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProjectCurrentSpend spends
*/
func (a *Client) GetProjectCurrentSpend(params *GetProjectCurrentSpendParams, opts ...ClientOption) (*GetProjectCurrentSpendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectCurrentSpendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProjectCurrentSpend",
		Method:             "GET",
		PathPattern:        "/v1/projects/{id}/spend/current",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProjectCurrentSpendReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProjectCurrentSpendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProjectCurrentSpendDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProjectSpendDetails spends details
*/
func (a *Client) GetProjectSpendDetails(params *GetProjectSpendDetailsParams, opts ...ClientOption) (*GetProjectSpendDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectSpendDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProjectSpendDetails",
		Method:             "GET",
		PathPattern:        "/v1/projects/{projectId}/spend/{spendId}/details",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProjectSpendDetailsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProjectSpendDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProjectSpendDetailsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProjectSpendHistory spends history
*/
func (a *Client) GetProjectSpendHistory(params *GetProjectSpendHistoryParams, opts ...ClientOption) (*GetProjectSpendHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectSpendHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProjectSpendHistory",
		Method:             "GET",
		PathPattern:        "/v1/projects/{id}/spend",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProjectSpendHistoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProjectSpendHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProjectSpendHistoryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectSSHKeys lists SSH keys
*/
func (a *Client) ListProjectSSHKeys(params *ListProjectSSHKeysParams, opts ...ClientOption) (*ListProjectSSHKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectSSHKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListProjectSshKeys",
		Method:             "GET",
		PathPattern:        "/v1/project/{projectId}/ssh-keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectSSHKeysReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProjectSSHKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectSSHKeysDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjects lists
*/
func (a *Client) ListProjects(params *ListProjectsParams, opts ...ClientOption) (*ListProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListProjects",
		Method:             "GET",
		PathPattern:        "/v1/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateProject updates
*/
func (a *Client) UpdateProject(params *UpdateProjectParams, opts ...ClientOption) (*UpdateProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateProject",
		Method:             "PATCH",
		PathPattern:        "/v1/projects/{project.id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateProjectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateProjectDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
