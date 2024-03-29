// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package healthy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new healthy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for healthy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetHealthyLiveness(params *GetHealthyLivenessParams, opts ...ClientOption) (*GetHealthyLivenessOK, error)

	GetHealthyReadiness(params *GetHealthyReadinessParams, opts ...ClientOption) (*GetHealthyReadinessOK, error)

	GetHealthyStartup(params *GetHealthyStartupParams, opts ...ClientOption) (*GetHealthyStartupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetHealthyLiveness livenesses probe

pod liveness probe for agent and controller pod
*/
func (a *Client) GetHealthyLiveness(params *GetHealthyLivenessParams, opts ...ClientOption) (*GetHealthyLivenessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHealthyLivenessParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHealthyLiveness",
		Method:             "GET",
		PathPattern:        "/healthy/liveness",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHealthyLivenessReader{formats: a.formats},
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
	success, ok := result.(*GetHealthyLivenessOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHealthyLiveness: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHealthyReadiness readinesses probe

pod readiness probe for agent and controller pod
*/
func (a *Client) GetHealthyReadiness(params *GetHealthyReadinessParams, opts ...ClientOption) (*GetHealthyReadinessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHealthyReadinessParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHealthyReadiness",
		Method:             "GET",
		PathPattern:        "/healthy/readiness",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHealthyReadinessReader{formats: a.formats},
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
	success, ok := result.(*GetHealthyReadinessOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHealthyReadiness: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHealthyStartup startups probe

pod startup probe for agent and controller pod
*/
func (a *Client) GetHealthyStartup(params *GetHealthyStartupParams, opts ...ClientOption) (*GetHealthyStartupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHealthyStartupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHealthyStartup",
		Method:             "GET",
		PathPattern:        "/healthy/startup",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHealthyStartupReader{formats: a.formats},
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
	success, ok := result.(*GetHealthyStartupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHealthyStartup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
