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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetAccessOrganizationParams creates a new GetAccessOrganizationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccessOrganizationParams() *GetAccessOrganizationParams {
	return &GetAccessOrganizationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccessOrganizationParamsWithTimeout creates a new GetAccessOrganizationParams object
// with the ability to set a timeout on a request.
func NewGetAccessOrganizationParamsWithTimeout(timeout time.Duration) *GetAccessOrganizationParams {
	return &GetAccessOrganizationParams{
		timeout: timeout,
	}
}

// NewGetAccessOrganizationParamsWithContext creates a new GetAccessOrganizationParams object
// with the ability to set a context for a request.
func NewGetAccessOrganizationParamsWithContext(ctx context.Context) *GetAccessOrganizationParams {
	return &GetAccessOrganizationParams{
		Context: ctx,
	}
}

// NewGetAccessOrganizationParamsWithHTTPClient creates a new GetAccessOrganizationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccessOrganizationParamsWithHTTPClient(client *http.Client) *GetAccessOrganizationParams {
	return &GetAccessOrganizationParams{
		HTTPClient: client,
	}
}

/*
GetAccessOrganizationParams contains all the parameters to send to the API endpoint

	for the get access organization operation.

	Typically these are written to a http.Request.
*/
type GetAccessOrganizationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get access organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccessOrganizationParams) WithDefaults() *GetAccessOrganizationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get access organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccessOrganizationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get access organization params
func (o *GetAccessOrganizationParams) WithTimeout(timeout time.Duration) *GetAccessOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get access organization params
func (o *GetAccessOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get access organization params
func (o *GetAccessOrganizationParams) WithContext(ctx context.Context) *GetAccessOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get access organization params
func (o *GetAccessOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get access organization params
func (o *GetAccessOrganizationParams) WithHTTPClient(client *http.Client) *GetAccessOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get access organization params
func (o *GetAccessOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccessOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
