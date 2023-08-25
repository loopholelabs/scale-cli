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

package cloud

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

// NewGetCloudDeploymentParams creates a new GetCloudDeploymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCloudDeploymentParams() *GetCloudDeploymentParams {
	return &GetCloudDeploymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudDeploymentParamsWithTimeout creates a new GetCloudDeploymentParams object
// with the ability to set a timeout on a request.
func NewGetCloudDeploymentParamsWithTimeout(timeout time.Duration) *GetCloudDeploymentParams {
	return &GetCloudDeploymentParams{
		timeout: timeout,
	}
}

// NewGetCloudDeploymentParamsWithContext creates a new GetCloudDeploymentParams object
// with the ability to set a context for a request.
func NewGetCloudDeploymentParamsWithContext(ctx context.Context) *GetCloudDeploymentParams {
	return &GetCloudDeploymentParams{
		Context: ctx,
	}
}

// NewGetCloudDeploymentParamsWithHTTPClient creates a new GetCloudDeploymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCloudDeploymentParamsWithHTTPClient(client *http.Client) *GetCloudDeploymentParams {
	return &GetCloudDeploymentParams{
		HTTPClient: client,
	}
}

/*
GetCloudDeploymentParams contains all the parameters to send to the API endpoint

	for the get cloud deployment operation.

	Typically these are written to a http.Request.
*/
type GetCloudDeploymentParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cloud deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCloudDeploymentParams) WithDefaults() *GetCloudDeploymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cloud deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCloudDeploymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cloud deployment params
func (o *GetCloudDeploymentParams) WithTimeout(timeout time.Duration) *GetCloudDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud deployment params
func (o *GetCloudDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud deployment params
func (o *GetCloudDeploymentParams) WithContext(ctx context.Context) *GetCloudDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud deployment params
func (o *GetCloudDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud deployment params
func (o *GetCloudDeploymentParams) WithHTTPClient(client *http.Client) *GetCloudDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud deployment params
func (o *GetCloudDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}