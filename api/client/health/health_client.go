// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new health API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for health API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Health performs a healthcheck
*/
func (a *Client) Health(params *HealthParams, authInfo runtime.ClientAuthInfoWriter) (*HealthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHealthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "health",
		Method:             "GET",
		PathPattern:        "/v1/health",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HealthReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HealthOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
