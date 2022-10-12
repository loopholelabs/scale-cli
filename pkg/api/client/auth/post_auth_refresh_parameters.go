// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// NewPostAuthRefreshParams creates a new PostAuthRefreshParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAuthRefreshParams() *PostAuthRefreshParams {
	return &PostAuthRefreshParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAuthRefreshParamsWithTimeout creates a new PostAuthRefreshParams object
// with the ability to set a timeout on a request.
func NewPostAuthRefreshParamsWithTimeout(timeout time.Duration) *PostAuthRefreshParams {
	return &PostAuthRefreshParams{
		timeout: timeout,
	}
}

// NewPostAuthRefreshParamsWithContext creates a new PostAuthRefreshParams object
// with the ability to set a context for a request.
func NewPostAuthRefreshParamsWithContext(ctx context.Context) *PostAuthRefreshParams {
	return &PostAuthRefreshParams{
		Context: ctx,
	}
}

// NewPostAuthRefreshParamsWithHTTPClient creates a new PostAuthRefreshParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAuthRefreshParamsWithHTTPClient(client *http.Client) *PostAuthRefreshParams {
	return &PostAuthRefreshParams{
		HTTPClient: client,
	}
}

/* PostAuthRefreshParams contains all the parameters to send to the API endpoint
   for the post auth refresh operation.

   Typically these are written to a http.Request.
*/
type PostAuthRefreshParams struct {

	/* GrantType.

	   Grant Type
	*/
	GrantType string

	/* RefreshToken.

	   Refresh Token
	*/
	RefreshToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post auth refresh params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAuthRefreshParams) WithDefaults() *PostAuthRefreshParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post auth refresh params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAuthRefreshParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post auth refresh params
func (o *PostAuthRefreshParams) WithTimeout(timeout time.Duration) *PostAuthRefreshParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post auth refresh params
func (o *PostAuthRefreshParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post auth refresh params
func (o *PostAuthRefreshParams) WithContext(ctx context.Context) *PostAuthRefreshParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post auth refresh params
func (o *PostAuthRefreshParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post auth refresh params
func (o *PostAuthRefreshParams) WithHTTPClient(client *http.Client) *PostAuthRefreshParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post auth refresh params
func (o *PostAuthRefreshParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGrantType adds the grantType to the post auth refresh params
func (o *PostAuthRefreshParams) WithGrantType(grantType string) *PostAuthRefreshParams {
	o.SetGrantType(grantType)
	return o
}

// SetGrantType adds the grantType to the post auth refresh params
func (o *PostAuthRefreshParams) SetGrantType(grantType string) {
	o.GrantType = grantType
}

// WithRefreshToken adds the refreshToken to the post auth refresh params
func (o *PostAuthRefreshParams) WithRefreshToken(refreshToken string) *PostAuthRefreshParams {
	o.SetRefreshToken(refreshToken)
	return o
}

// SetRefreshToken adds the refreshToken to the post auth refresh params
func (o *PostAuthRefreshParams) SetRefreshToken(refreshToken string) {
	o.RefreshToken = refreshToken
}

// WriteToRequest writes these params to a swagger request
func (o *PostAuthRefreshParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param grantType
	frGrantType := o.GrantType
	fGrantType := frGrantType
	if fGrantType != "" {
		if err := r.SetFormParam("grantType", fGrantType); err != nil {
			return err
		}
	}

	// form param refreshToken
	frRefreshToken := o.RefreshToken
	fRefreshToken := frRefreshToken
	if fRefreshToken != "" {
		if err := r.SetFormParam("refreshToken", fRefreshToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
