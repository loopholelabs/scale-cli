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

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetRegistryFunctionOrgNameTagParams creates a new GetRegistryFunctionOrgNameTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRegistryFunctionOrgNameTagParams() *GetRegistryFunctionOrgNameTagParams {
	return &GetRegistryFunctionOrgNameTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegistryFunctionOrgNameTagParamsWithTimeout creates a new GetRegistryFunctionOrgNameTagParams object
// with the ability to set a timeout on a request.
func NewGetRegistryFunctionOrgNameTagParamsWithTimeout(timeout time.Duration) *GetRegistryFunctionOrgNameTagParams {
	return &GetRegistryFunctionOrgNameTagParams{
		timeout: timeout,
	}
}

// NewGetRegistryFunctionOrgNameTagParamsWithContext creates a new GetRegistryFunctionOrgNameTagParams object
// with the ability to set a context for a request.
func NewGetRegistryFunctionOrgNameTagParamsWithContext(ctx context.Context) *GetRegistryFunctionOrgNameTagParams {
	return &GetRegistryFunctionOrgNameTagParams{
		Context: ctx,
	}
}

// NewGetRegistryFunctionOrgNameTagParamsWithHTTPClient creates a new GetRegistryFunctionOrgNameTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRegistryFunctionOrgNameTagParamsWithHTTPClient(client *http.Client) *GetRegistryFunctionOrgNameTagParams {
	return &GetRegistryFunctionOrgNameTagParams{
		HTTPClient: client,
	}
}

/*
GetRegistryFunctionOrgNameTagParams contains all the parameters to send to the API endpoint

	for the get registry function org name tag operation.

	Typically these are written to a http.Request.
*/
type GetRegistryFunctionOrgNameTagParams struct {

	/* Name.

	   name
	*/
	Name string

	/* Org.

	   org
	*/
	Org string

	/* Tag.

	   tag
	*/
	Tag string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get registry function org name tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegistryFunctionOrgNameTagParams) WithDefaults() *GetRegistryFunctionOrgNameTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get registry function org name tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegistryFunctionOrgNameTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get registry function org name tag params
func (o *GetRegistryFunctionOrgNameTagParams) WithTimeout(timeout time.Duration) *GetRegistryFunctionOrgNameTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get registry function org name tag params
func (o *GetRegistryFunctionOrgNameTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get registry function org name tag params
func (o *GetRegistryFunctionOrgNameTagParams) WithContext(ctx context.Context) *GetRegistryFunctionOrgNameTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get registry function org name tag params
func (o *GetRegistryFunctionOrgNameTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get registry function org name tag params
func (o *GetRegistryFunctionOrgNameTagParams) WithHTTPClient(client *http.Client) *GetRegistryFunctionOrgNameTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get registry function org name tag params
func (o *GetRegistryFunctionOrgNameTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get registry function org name tag params
func (o *GetRegistryFunctionOrgNameTagParams) WithName(name string) *GetRegistryFunctionOrgNameTagParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get registry function org name tag params
func (o *GetRegistryFunctionOrgNameTagParams) SetName(name string) {
	o.Name = name
}

// WithOrg adds the org to the get registry function org name tag params
func (o *GetRegistryFunctionOrgNameTagParams) WithOrg(org string) *GetRegistryFunctionOrgNameTagParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the get registry function org name tag params
func (o *GetRegistryFunctionOrgNameTagParams) SetOrg(org string) {
	o.Org = org
}

// WithTag adds the tag to the get registry function org name tag params
func (o *GetRegistryFunctionOrgNameTagParams) WithTag(tag string) *GetRegistryFunctionOrgNameTagParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the get registry function org name tag params
func (o *GetRegistryFunctionOrgNameTagParams) SetTag(tag string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegistryFunctionOrgNameTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param org
	if err := r.SetPathParam("org", o.Org); err != nil {
		return err
	}

	// path param tag
	if err := r.SetPathParam("tag", o.Tag); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
