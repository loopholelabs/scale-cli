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

	"github.com/loopholelabs/scale-cli/client/models"
)

// NewPostCloudDeploymentParams creates a new PostCloudDeploymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostCloudDeploymentParams() *PostCloudDeploymentParams {
	return &PostCloudDeploymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostCloudDeploymentParamsWithTimeout creates a new PostCloudDeploymentParams object
// with the ability to set a timeout on a request.
func NewPostCloudDeploymentParamsWithTimeout(timeout time.Duration) *PostCloudDeploymentParams {
	return &PostCloudDeploymentParams{
		timeout: timeout,
	}
}

// NewPostCloudDeploymentParamsWithContext creates a new PostCloudDeploymentParams object
// with the ability to set a context for a request.
func NewPostCloudDeploymentParamsWithContext(ctx context.Context) *PostCloudDeploymentParams {
	return &PostCloudDeploymentParams{
		Context: ctx,
	}
}

// NewPostCloudDeploymentParamsWithHTTPClient creates a new PostCloudDeploymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostCloudDeploymentParamsWithHTTPClient(client *http.Client) *PostCloudDeploymentParams {
	return &PostCloudDeploymentParams{
		HTTPClient: client,
	}
}

/*
PostCloudDeploymentParams contains all the parameters to send to the API endpoint

	for the post cloud deployment operation.

	Typically these are written to a http.Request.
*/
type PostCloudDeploymentParams struct {

	/* Request.

	   Create Deployment Request
	*/
	Request *models.ModelsCreateDeploymentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post cloud deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCloudDeploymentParams) WithDefaults() *PostCloudDeploymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post cloud deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCloudDeploymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post cloud deployment params
func (o *PostCloudDeploymentParams) WithTimeout(timeout time.Duration) *PostCloudDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cloud deployment params
func (o *PostCloudDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cloud deployment params
func (o *PostCloudDeploymentParams) WithContext(ctx context.Context) *PostCloudDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cloud deployment params
func (o *PostCloudDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cloud deployment params
func (o *PostCloudDeploymentParams) WithHTTPClient(client *http.Client) *PostCloudDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cloud deployment params
func (o *PostCloudDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post cloud deployment params
func (o *PostCloudDeploymentParams) WithRequest(request *models.ModelsCreateDeploymentRequest) *PostCloudDeploymentParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post cloud deployment params
func (o *PostCloudDeploymentParams) SetRequest(request *models.ModelsCreateDeploymentRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostCloudDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
