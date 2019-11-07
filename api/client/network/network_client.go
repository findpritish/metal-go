// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new network API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for network API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AllocateNetwork allocates a child network from a partition s private super network
*/
func (a *Client) AllocateNetwork(params *AllocateNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*AllocateNetworkCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllocateNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "allocateNetwork",
		Method:             "POST",
		PathPattern:        "/v1/network/allocate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AllocateNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AllocateNetworkCreated), nil

}

/*
CreateNetwork creates a network if the given ID already exists a conflict is returned
*/
func (a *Client) CreateNetwork(params *CreateNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*CreateNetworkCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createNetwork",
		Method:             "PUT",
		PathPattern:        "/v1/network",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateNetworkCreated), nil

}

/*
DeleteNetwork deletes a network and returns the deleted entity
*/
func (a *Client) DeleteNetwork(params *DeleteNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNetwork",
		Method:             "DELETE",
		PathPattern:        "/v1/network/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNetworkOK), nil

}

/*
FindNetwork gets network by id
*/
func (a *Client) FindNetwork(params *FindNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*FindNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findNetwork",
		Method:             "GET",
		PathPattern:        "/v1/network/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindNetworkOK), nil

}

/*
FindNetworks gets all networks that match given properties
*/
func (a *Client) FindNetworks(params *FindNetworksParams, authInfo runtime.ClientAuthInfoWriter) (*FindNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindNetworksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findNetworks",
		Method:             "POST",
		PathPattern:        "/v1/network/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindNetworksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindNetworksOK), nil

}

/*
FreeNetwork frees a network
*/
func (a *Client) FreeNetwork(params *FreeNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*FreeNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFreeNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "freeNetwork",
		Method:             "POST",
		PathPattern:        "/v1/network/free/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FreeNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FreeNetworkOK), nil

}

/*
ListNetworks gets all networks
*/
func (a *Client) ListNetworks(params *ListNetworksParams, authInfo runtime.ClientAuthInfoWriter) (*ListNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNetworksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listNetworks",
		Method:             "GET",
		PathPattern:        "/v1/network",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListNetworksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListNetworksOK), nil

}

/*
UpdateNetwork updates a network if the network was changed since this one was read a conflict is returned
*/
func (a *Client) UpdateNetwork(params *UpdateNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetwork",
		Method:             "POST",
		PathPattern:        "/v1/network",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateNetworkOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
