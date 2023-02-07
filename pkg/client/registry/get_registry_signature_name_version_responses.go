// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/loopholelabs/scale-cli/pkg/client/models"
)

// GetRegistrySignatureNameVersionReader is a Reader for the GetRegistrySignatureNameVersion structure.
type GetRegistrySignatureNameVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegistrySignatureNameVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegistrySignatureNameVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRegistrySignatureNameVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRegistrySignatureNameVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRegistrySignatureNameVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRegistrySignatureNameVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRegistrySignatureNameVersionOK creates a GetRegistrySignatureNameVersionOK with default headers values
func NewGetRegistrySignatureNameVersionOK() *GetRegistrySignatureNameVersionOK {
	return &GetRegistrySignatureNameVersionOK{}
}

/*
	GetRegistrySignatureNameVersionOK describes a response with status code 200, with default header values.

OK
*/
type GetRegistrySignatureNameVersionOK struct {
	Payload *models.ModelsGetSignatureResponse
}

// IsSuccess returns true when this get registry signature name version o k response has a 2xx status code
func (o *GetRegistrySignatureNameVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get registry signature name version o k response has a 3xx status code
func (o *GetRegistrySignatureNameVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry signature name version o k response has a 4xx status code
func (o *GetRegistrySignatureNameVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get registry signature name version o k response has a 5xx status code
func (o *GetRegistrySignatureNameVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry signature name version o k response a status code equal to that given
func (o *GetRegistrySignatureNameVersionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRegistrySignatureNameVersionOK) Error() string {
	return fmt.Sprintf("[GET /registry/signature/{name}/{version}][%d] getRegistrySignatureNameVersionOK  %+v", 200, o.Payload)
}

func (o *GetRegistrySignatureNameVersionOK) String() string {
	return fmt.Sprintf("[GET /registry/signature/{name}/{version}][%d] getRegistrySignatureNameVersionOK  %+v", 200, o.Payload)
}

func (o *GetRegistrySignatureNameVersionOK) GetPayload() *models.ModelsGetSignatureResponse {
	return o.Payload
}

func (o *GetRegistrySignatureNameVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsGetSignatureResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistrySignatureNameVersionBadRequest creates a GetRegistrySignatureNameVersionBadRequest with default headers values
func NewGetRegistrySignatureNameVersionBadRequest() *GetRegistrySignatureNameVersionBadRequest {
	return &GetRegistrySignatureNameVersionBadRequest{}
}

/*
	GetRegistrySignatureNameVersionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetRegistrySignatureNameVersionBadRequest struct {
	Payload string
}

// IsSuccess returns true when this get registry signature name version bad request response has a 2xx status code
func (o *GetRegistrySignatureNameVersionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry signature name version bad request response has a 3xx status code
func (o *GetRegistrySignatureNameVersionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry signature name version bad request response has a 4xx status code
func (o *GetRegistrySignatureNameVersionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry signature name version bad request response has a 5xx status code
func (o *GetRegistrySignatureNameVersionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry signature name version bad request response a status code equal to that given
func (o *GetRegistrySignatureNameVersionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRegistrySignatureNameVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /registry/signature/{name}/{version}][%d] getRegistrySignatureNameVersionBadRequest  %+v", 400, o.Payload)
}

func (o *GetRegistrySignatureNameVersionBadRequest) String() string {
	return fmt.Sprintf("[GET /registry/signature/{name}/{version}][%d] getRegistrySignatureNameVersionBadRequest  %+v", 400, o.Payload)
}

func (o *GetRegistrySignatureNameVersionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetRegistrySignatureNameVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistrySignatureNameVersionUnauthorized creates a GetRegistrySignatureNameVersionUnauthorized with default headers values
func NewGetRegistrySignatureNameVersionUnauthorized() *GetRegistrySignatureNameVersionUnauthorized {
	return &GetRegistrySignatureNameVersionUnauthorized{}
}

/*
	GetRegistrySignatureNameVersionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRegistrySignatureNameVersionUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this get registry signature name version unauthorized response has a 2xx status code
func (o *GetRegistrySignatureNameVersionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry signature name version unauthorized response has a 3xx status code
func (o *GetRegistrySignatureNameVersionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry signature name version unauthorized response has a 4xx status code
func (o *GetRegistrySignatureNameVersionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry signature name version unauthorized response has a 5xx status code
func (o *GetRegistrySignatureNameVersionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry signature name version unauthorized response a status code equal to that given
func (o *GetRegistrySignatureNameVersionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRegistrySignatureNameVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /registry/signature/{name}/{version}][%d] getRegistrySignatureNameVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRegistrySignatureNameVersionUnauthorized) String() string {
	return fmt.Sprintf("[GET /registry/signature/{name}/{version}][%d] getRegistrySignatureNameVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRegistrySignatureNameVersionUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *GetRegistrySignatureNameVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistrySignatureNameVersionNotFound creates a GetRegistrySignatureNameVersionNotFound with default headers values
func NewGetRegistrySignatureNameVersionNotFound() *GetRegistrySignatureNameVersionNotFound {
	return &GetRegistrySignatureNameVersionNotFound{}
}

/*
	GetRegistrySignatureNameVersionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRegistrySignatureNameVersionNotFound struct {
	Payload string
}

// IsSuccess returns true when this get registry signature name version not found response has a 2xx status code
func (o *GetRegistrySignatureNameVersionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry signature name version not found response has a 3xx status code
func (o *GetRegistrySignatureNameVersionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry signature name version not found response has a 4xx status code
func (o *GetRegistrySignatureNameVersionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry signature name version not found response has a 5xx status code
func (o *GetRegistrySignatureNameVersionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry signature name version not found response a status code equal to that given
func (o *GetRegistrySignatureNameVersionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRegistrySignatureNameVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /registry/signature/{name}/{version}][%d] getRegistrySignatureNameVersionNotFound  %+v", 404, o.Payload)
}

func (o *GetRegistrySignatureNameVersionNotFound) String() string {
	return fmt.Sprintf("[GET /registry/signature/{name}/{version}][%d] getRegistrySignatureNameVersionNotFound  %+v", 404, o.Payload)
}

func (o *GetRegistrySignatureNameVersionNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetRegistrySignatureNameVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistrySignatureNameVersionInternalServerError creates a GetRegistrySignatureNameVersionInternalServerError with default headers values
func NewGetRegistrySignatureNameVersionInternalServerError() *GetRegistrySignatureNameVersionInternalServerError {
	return &GetRegistrySignatureNameVersionInternalServerError{}
}

/*
	GetRegistrySignatureNameVersionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRegistrySignatureNameVersionInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get registry signature name version internal server error response has a 2xx status code
func (o *GetRegistrySignatureNameVersionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry signature name version internal server error response has a 3xx status code
func (o *GetRegistrySignatureNameVersionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry signature name version internal server error response has a 4xx status code
func (o *GetRegistrySignatureNameVersionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get registry signature name version internal server error response has a 5xx status code
func (o *GetRegistrySignatureNameVersionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get registry signature name version internal server error response a status code equal to that given
func (o *GetRegistrySignatureNameVersionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRegistrySignatureNameVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /registry/signature/{name}/{version}][%d] getRegistrySignatureNameVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRegistrySignatureNameVersionInternalServerError) String() string {
	return fmt.Sprintf("[GET /registry/signature/{name}/{version}][%d] getRegistrySignatureNameVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRegistrySignatureNameVersionInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetRegistrySignatureNameVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}