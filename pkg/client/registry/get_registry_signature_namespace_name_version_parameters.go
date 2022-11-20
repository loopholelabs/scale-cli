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

// NewGetRegistrySignatureNamespaceNameVersionParams creates a new GetRegistrySignatureNamespaceNameVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRegistrySignatureNamespaceNameVersionParams() *GetRegistrySignatureNamespaceNameVersionParams {
	return &GetRegistrySignatureNamespaceNameVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegistrySignatureNamespaceNameVersionParamsWithTimeout creates a new GetRegistrySignatureNamespaceNameVersionParams object
// with the ability to set a timeout on a request.
func NewGetRegistrySignatureNamespaceNameVersionParamsWithTimeout(timeout time.Duration) *GetRegistrySignatureNamespaceNameVersionParams {
	return &GetRegistrySignatureNamespaceNameVersionParams{
		timeout: timeout,
	}
}

// NewGetRegistrySignatureNamespaceNameVersionParamsWithContext creates a new GetRegistrySignatureNamespaceNameVersionParams object
// with the ability to set a context for a request.
func NewGetRegistrySignatureNamespaceNameVersionParamsWithContext(ctx context.Context) *GetRegistrySignatureNamespaceNameVersionParams {
	return &GetRegistrySignatureNamespaceNameVersionParams{
		Context: ctx,
	}
}

// NewGetRegistrySignatureNamespaceNameVersionParamsWithHTTPClient creates a new GetRegistrySignatureNamespaceNameVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRegistrySignatureNamespaceNameVersionParamsWithHTTPClient(client *http.Client) *GetRegistrySignatureNamespaceNameVersionParams {
	return &GetRegistrySignatureNamespaceNameVersionParams{
		HTTPClient: client,
	}
}

/*
GetRegistrySignatureNamespaceNameVersionParams contains all the parameters to send to the API endpoint

	for the get registry signature namespace name version operation.

	Typically these are written to a http.Request.
*/
type GetRegistrySignatureNamespaceNameVersionParams struct {

	/* Name.

	   name
	*/
	Name string

	/* Namespace.

	   namespace
	*/
	Namespace string

	/* Version.

	   version
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get registry signature namespace name version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegistrySignatureNamespaceNameVersionParams) WithDefaults() *GetRegistrySignatureNamespaceNameVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get registry signature namespace name version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegistrySignatureNamespaceNameVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get registry signature namespace name version params
func (o *GetRegistrySignatureNamespaceNameVersionParams) WithTimeout(timeout time.Duration) *GetRegistrySignatureNamespaceNameVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get registry signature namespace name version params
func (o *GetRegistrySignatureNamespaceNameVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get registry signature namespace name version params
func (o *GetRegistrySignatureNamespaceNameVersionParams) WithContext(ctx context.Context) *GetRegistrySignatureNamespaceNameVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get registry signature namespace name version params
func (o *GetRegistrySignatureNamespaceNameVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get registry signature namespace name version params
func (o *GetRegistrySignatureNamespaceNameVersionParams) WithHTTPClient(client *http.Client) *GetRegistrySignatureNamespaceNameVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get registry signature namespace name version params
func (o *GetRegistrySignatureNamespaceNameVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get registry signature namespace name version params
func (o *GetRegistrySignatureNamespaceNameVersionParams) WithName(name string) *GetRegistrySignatureNamespaceNameVersionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get registry signature namespace name version params
func (o *GetRegistrySignatureNamespaceNameVersionParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the get registry signature namespace name version params
func (o *GetRegistrySignatureNamespaceNameVersionParams) WithNamespace(namespace string) *GetRegistrySignatureNamespaceNameVersionParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get registry signature namespace name version params
func (o *GetRegistrySignatureNamespaceNameVersionParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithVersion adds the version to the get registry signature namespace name version params
func (o *GetRegistrySignatureNamespaceNameVersionParams) WithVersion(version string) *GetRegistrySignatureNamespaceNameVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get registry signature namespace name version params
func (o *GetRegistrySignatureNamespaceNameVersionParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegistrySignatureNamespaceNameVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}