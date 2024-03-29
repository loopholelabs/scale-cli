/*
 	Copyright 2023 Loophole Labs

 	Licensed under the Apache License, Version 2.0 (the "License");
 	you may not use this file except in compliance with the License.
 	You may obtain a copy of the License at

 		   http://www.apache.org/licenses/LICENSE-2.0

 	Unless required by applicable law or agreed to in writing, software
 	distributed under the License is distributed on an "AS IS" BASIS,
 	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 	See the License for the specific language governing permissions and
 	limitations under the License.
*/

// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new access API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for access API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAccessApikeyNameorid(params *DeleteAccessApikeyNameoridParams, opts ...ClientOption) (*DeleteAccessApikeyNameoridOK, error)

	DeleteAccessInviteOrganization(params *DeleteAccessInviteOrganizationParams, opts ...ClientOption) (*DeleteAccessInviteOrganizationOK, error)

	DeleteAccessOrganizationInviteEmail(params *DeleteAccessOrganizationInviteEmailParams, opts ...ClientOption) (*DeleteAccessOrganizationInviteEmailOK, error)

	DeleteAccessOrganizationName(params *DeleteAccessOrganizationNameParams, opts ...ClientOption) (*DeleteAccessOrganizationNameOK, error)

	GetAccessApikey(params *GetAccessApikeyParams, opts ...ClientOption) (*GetAccessApikeyOK, error)

	GetAccessApikeyNameorid(params *GetAccessApikeyNameoridParams, opts ...ClientOption) (*GetAccessApikeyNameoridOK, error)

	GetAccessInvite(params *GetAccessInviteParams, opts ...ClientOption) (*GetAccessInviteOK, error)

	GetAccessOrganization(params *GetAccessOrganizationParams, opts ...ClientOption) (*GetAccessOrganizationOK, error)

	GetAccessOrganizationInvite(params *GetAccessOrganizationInviteParams, opts ...ClientOption) (*GetAccessOrganizationInviteOK, error)

	PostAccessApikey(params *PostAccessApikeyParams, opts ...ClientOption) (*PostAccessApikeyOK, error)

	PostAccessInviteOrganization(params *PostAccessInviteOrganizationParams, opts ...ClientOption) (*PostAccessInviteOrganizationOK, error)

	PostAccessOrganization(params *PostAccessOrganizationParams, opts ...ClientOption) (*PostAccessOrganizationOK, error)

	PostAccessOrganizationInvite(params *PostAccessOrganizationInviteParams, opts ...ClientOption) (*PostAccessOrganizationInviteOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteAccessApikeyNameorid Deletes an API Key given its `name` or `id`. The API Key must be part of the organization that the current session is scoped to.
*/
func (a *Client) DeleteAccessApikeyNameorid(params *DeleteAccessApikeyNameoridParams, opts ...ClientOption) (*DeleteAccessApikeyNameoridOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAccessApikeyNameoridParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAccessApikeyNameorid",
		Method:             "DELETE",
		PathPattern:        "/access/apikey/{nameorid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAccessApikeyNameoridReader{formats: a.formats},
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
	success, ok := result.(*DeleteAccessApikeyNameoridOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAccessApikeyNameorid: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAccessInviteOrganization Declines an Organization Invite given its `organization`.
*/
func (a *Client) DeleteAccessInviteOrganization(params *DeleteAccessInviteOrganizationParams, opts ...ClientOption) (*DeleteAccessInviteOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAccessInviteOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAccessInviteOrganization",
		Method:             "DELETE",
		PathPattern:        "/access/invite/{organization}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAccessInviteOrganizationReader{formats: a.formats},
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
	success, ok := result.(*DeleteAccessInviteOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAccessInviteOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAccessOrganizationInviteEmail Deletes an Organization Invite given its `email`. The current session must be scoped to the Organization.
*/
func (a *Client) DeleteAccessOrganizationInviteEmail(params *DeleteAccessOrganizationInviteEmailParams, opts ...ClientOption) (*DeleteAccessOrganizationInviteEmailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAccessOrganizationInviteEmailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAccessOrganizationInviteEmail",
		Method:             "DELETE",
		PathPattern:        "/access/organization/invite/{email}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAccessOrganizationInviteEmailReader{formats: a.formats},
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
	success, ok := result.(*DeleteAccessOrganizationInviteEmailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAccessOrganizationInviteEmail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAccessOrganizationName Deletes an Organization given its `name`. The user must be a member of the Organization.
*/
func (a *Client) DeleteAccessOrganizationName(params *DeleteAccessOrganizationNameParams, opts ...ClientOption) (*DeleteAccessOrganizationNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAccessOrganizationNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAccessOrganizationName",
		Method:             "DELETE",
		PathPattern:        "/access/organization/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAccessOrganizationNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteAccessOrganizationNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAccessOrganizationName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccessApikey Lists all the API Keys for the authenticated user. Only the API Keys for the organization that the current session is scoped to will be returned.
*/
func (a *Client) GetAccessApikey(params *GetAccessApikeyParams, opts ...ClientOption) (*GetAccessApikeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessApikeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAccessApikey",
		Method:             "GET",
		PathPattern:        "/access/apikey",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccessApikeyReader{formats: a.formats},
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
	success, ok := result.(*GetAccessApikeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccessApikey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccessApikeyNameorid Gets information about a specific API Key given its `name` or `id`. The API Key must be part of the organization that the current session is scoped to.
*/
func (a *Client) GetAccessApikeyNameorid(params *GetAccessApikeyNameoridParams, opts ...ClientOption) (*GetAccessApikeyNameoridOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessApikeyNameoridParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAccessApikeyNameorid",
		Method:             "GET",
		PathPattern:        "/access/apikey/{nameorid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccessApikeyNameoridReader{formats: a.formats},
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
	success, ok := result.(*GetAccessApikeyNameoridOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccessApikeyNameorid: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccessInvite Lists all the Organization Invites for the authenticated user.
*/
func (a *Client) GetAccessInvite(params *GetAccessInviteParams, opts ...ClientOption) (*GetAccessInviteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessInviteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAccessInvite",
		Method:             "GET",
		PathPattern:        "/access/invite",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccessInviteReader{formats: a.formats},
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
	success, ok := result.(*GetAccessInviteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccessInvite: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccessOrganization Lists all the Organizations for the authenticated user. Only the Organizations that the user is a member of will be returned.
*/
func (a *Client) GetAccessOrganization(params *GetAccessOrganizationParams, opts ...ClientOption) (*GetAccessOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAccessOrganization",
		Method:             "GET",
		PathPattern:        "/access/organization",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccessOrganizationReader{formats: a.formats},
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
	success, ok := result.(*GetAccessOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccessOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccessOrganizationInvite Lists all the Organization Invites for the authenticated user. Only the Organizations Invites for the Organization the session is scoped to will be returned.
*/
func (a *Client) GetAccessOrganizationInvite(params *GetAccessOrganizationInviteParams, opts ...ClientOption) (*GetAccessOrganizationInviteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessOrganizationInviteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAccessOrganizationInvite",
		Method:             "GET",
		PathPattern:        "/access/organization/invite",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccessOrganizationInviteReader{formats: a.formats},
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
	success, ok := result.(*GetAccessOrganizationInviteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccessOrganizationInvite: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAccessApikey Creates a new API Key with the given `name` scoped to all the organizations the user is a member or owner of. If the user's session is already tied to an organization, the new API Key will be scoped to that organization.
*/
func (a *Client) PostAccessApikey(params *PostAccessApikeyParams, opts ...ClientOption) (*PostAccessApikeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAccessApikeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAccessApikey",
		Method:             "POST",
		PathPattern:        "/access/apikey",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAccessApikeyReader{formats: a.formats},
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
	success, ok := result.(*PostAccessApikeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAccessApikey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAccessInviteOrganization Accept an Organization Invite given its `organization`.
*/
func (a *Client) PostAccessInviteOrganization(params *PostAccessInviteOrganizationParams, opts ...ClientOption) (*PostAccessInviteOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAccessInviteOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAccessInviteOrganization",
		Method:             "POST",
		PathPattern:        "/access/invite/{organization}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAccessInviteOrganizationReader{formats: a.formats},
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
	success, ok := result.(*PostAccessInviteOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAccessInviteOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAccessOrganization Creates a new Organization with the given `name`, and adds the user to it.
*/
func (a *Client) PostAccessOrganization(params *PostAccessOrganizationParams, opts ...ClientOption) (*PostAccessOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAccessOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAccessOrganization",
		Method:             "POST",
		PathPattern:        "/access/organization",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAccessOrganizationReader{formats: a.formats},
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
	success, ok := result.(*PostAccessOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAccessOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAccessOrganizationInvite Creates a new Organization Invite for the user with the given `email` for the Organization the session is scoped to.
*/
func (a *Client) PostAccessOrganizationInvite(params *PostAccessOrganizationInviteParams, opts ...ClientOption) (*PostAccessOrganizationInviteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAccessOrganizationInviteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAccessOrganizationInvite",
		Method:             "POST",
		PathPattern:        "/access/organization/invite",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAccessOrganizationInviteReader{formats: a.formats},
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
	success, ok := result.(*PostAccessOrganizationInviteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAccessOrganizationInvite: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
