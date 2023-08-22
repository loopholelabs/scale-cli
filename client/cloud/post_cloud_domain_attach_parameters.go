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

// NewPostCloudDomainAttachParams creates a new PostCloudDomainAttachParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostCloudDomainAttachParams() *PostCloudDomainAttachParams {
	return &PostCloudDomainAttachParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostCloudDomainAttachParamsWithTimeout creates a new PostCloudDomainAttachParams object
// with the ability to set a timeout on a request.
func NewPostCloudDomainAttachParamsWithTimeout(timeout time.Duration) *PostCloudDomainAttachParams {
	return &PostCloudDomainAttachParams{
		timeout: timeout,
	}
}

// NewPostCloudDomainAttachParamsWithContext creates a new PostCloudDomainAttachParams object
// with the ability to set a context for a request.
func NewPostCloudDomainAttachParamsWithContext(ctx context.Context) *PostCloudDomainAttachParams {
	return &PostCloudDomainAttachParams{
		Context: ctx,
	}
}

// NewPostCloudDomainAttachParamsWithHTTPClient creates a new PostCloudDomainAttachParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostCloudDomainAttachParamsWithHTTPClient(client *http.Client) *PostCloudDomainAttachParams {
	return &PostCloudDomainAttachParams{
		HTTPClient: client,
	}
}

/*
PostCloudDomainAttachParams contains all the parameters to send to the API endpoint

	for the post cloud domain attach operation.

	Typically these are written to a http.Request.
*/
type PostCloudDomainAttachParams struct {

	/* Request.

	   Attach Domain Request
	*/
	Request *models.ModelsAttachDomainRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post cloud domain attach params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCloudDomainAttachParams) WithDefaults() *PostCloudDomainAttachParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post cloud domain attach params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCloudDomainAttachParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post cloud domain attach params
func (o *PostCloudDomainAttachParams) WithTimeout(timeout time.Duration) *PostCloudDomainAttachParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cloud domain attach params
func (o *PostCloudDomainAttachParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cloud domain attach params
func (o *PostCloudDomainAttachParams) WithContext(ctx context.Context) *PostCloudDomainAttachParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cloud domain attach params
func (o *PostCloudDomainAttachParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cloud domain attach params
func (o *PostCloudDomainAttachParams) WithHTTPClient(client *http.Client) *PostCloudDomainAttachParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cloud domain attach params
func (o *PostCloudDomainAttachParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post cloud domain attach params
func (o *PostCloudDomainAttachParams) WithRequest(request *models.ModelsAttachDomainRequest) *PostCloudDomainAttachParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post cloud domain attach params
func (o *PostCloudDomainAttachParams) SetRequest(request *models.ModelsAttachDomainRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostCloudDomainAttachParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
