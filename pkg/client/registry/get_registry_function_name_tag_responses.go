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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/loopholelabs/scale-cli/pkg/client/models"
)

// GetRegistryFunctionNameTagReader is a Reader for the GetRegistryFunctionNameTag structure.
type GetRegistryFunctionNameTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegistryFunctionNameTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegistryFunctionNameTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRegistryFunctionNameTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRegistryFunctionNameTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRegistryFunctionNameTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRegistryFunctionNameTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRegistryFunctionNameTagOK creates a GetRegistryFunctionNameTagOK with default headers values
func NewGetRegistryFunctionNameTagOK() *GetRegistryFunctionNameTagOK {
	return &GetRegistryFunctionNameTagOK{}
}

/*
GetRegistryFunctionNameTagOK describes a response with status code 200, with default header values.

OK
*/
type GetRegistryFunctionNameTagOK struct {
	Payload *models.ModelsGetFunctionResponse
}

// IsSuccess returns true when this get registry function name tag o k response has a 2xx status code
func (o *GetRegistryFunctionNameTagOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get registry function name tag o k response has a 3xx status code
func (o *GetRegistryFunctionNameTagOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry function name tag o k response has a 4xx status code
func (o *GetRegistryFunctionNameTagOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get registry function name tag o k response has a 5xx status code
func (o *GetRegistryFunctionNameTagOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry function name tag o k response a status code equal to that given
func (o *GetRegistryFunctionNameTagOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get registry function name tag o k response
func (o *GetRegistryFunctionNameTagOK) Code() int {
	return 200
}

func (o *GetRegistryFunctionNameTagOK) Error() string {
	return fmt.Sprintf("[GET /registry/function/{name}/{tag}][%d] getRegistryFunctionNameTagOK  %+v", 200, o.Payload)
}

func (o *GetRegistryFunctionNameTagOK) String() string {
	return fmt.Sprintf("[GET /registry/function/{name}/{tag}][%d] getRegistryFunctionNameTagOK  %+v", 200, o.Payload)
}

func (o *GetRegistryFunctionNameTagOK) GetPayload() *models.ModelsGetFunctionResponse {
	return o.Payload
}

func (o *GetRegistryFunctionNameTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsGetFunctionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryFunctionNameTagBadRequest creates a GetRegistryFunctionNameTagBadRequest with default headers values
func NewGetRegistryFunctionNameTagBadRequest() *GetRegistryFunctionNameTagBadRequest {
	return &GetRegistryFunctionNameTagBadRequest{}
}

/*
GetRegistryFunctionNameTagBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetRegistryFunctionNameTagBadRequest struct {
	Payload string
}

// IsSuccess returns true when this get registry function name tag bad request response has a 2xx status code
func (o *GetRegistryFunctionNameTagBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry function name tag bad request response has a 3xx status code
func (o *GetRegistryFunctionNameTagBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry function name tag bad request response has a 4xx status code
func (o *GetRegistryFunctionNameTagBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry function name tag bad request response has a 5xx status code
func (o *GetRegistryFunctionNameTagBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry function name tag bad request response a status code equal to that given
func (o *GetRegistryFunctionNameTagBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get registry function name tag bad request response
func (o *GetRegistryFunctionNameTagBadRequest) Code() int {
	return 400
}

func (o *GetRegistryFunctionNameTagBadRequest) Error() string {
	return fmt.Sprintf("[GET /registry/function/{name}/{tag}][%d] getRegistryFunctionNameTagBadRequest  %+v", 400, o.Payload)
}

func (o *GetRegistryFunctionNameTagBadRequest) String() string {
	return fmt.Sprintf("[GET /registry/function/{name}/{tag}][%d] getRegistryFunctionNameTagBadRequest  %+v", 400, o.Payload)
}

func (o *GetRegistryFunctionNameTagBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetRegistryFunctionNameTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryFunctionNameTagUnauthorized creates a GetRegistryFunctionNameTagUnauthorized with default headers values
func NewGetRegistryFunctionNameTagUnauthorized() *GetRegistryFunctionNameTagUnauthorized {
	return &GetRegistryFunctionNameTagUnauthorized{}
}

/*
GetRegistryFunctionNameTagUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRegistryFunctionNameTagUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this get registry function name tag unauthorized response has a 2xx status code
func (o *GetRegistryFunctionNameTagUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry function name tag unauthorized response has a 3xx status code
func (o *GetRegistryFunctionNameTagUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry function name tag unauthorized response has a 4xx status code
func (o *GetRegistryFunctionNameTagUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry function name tag unauthorized response has a 5xx status code
func (o *GetRegistryFunctionNameTagUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry function name tag unauthorized response a status code equal to that given
func (o *GetRegistryFunctionNameTagUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get registry function name tag unauthorized response
func (o *GetRegistryFunctionNameTagUnauthorized) Code() int {
	return 401
}

func (o *GetRegistryFunctionNameTagUnauthorized) Error() string {
	return fmt.Sprintf("[GET /registry/function/{name}/{tag}][%d] getRegistryFunctionNameTagUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRegistryFunctionNameTagUnauthorized) String() string {
	return fmt.Sprintf("[GET /registry/function/{name}/{tag}][%d] getRegistryFunctionNameTagUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRegistryFunctionNameTagUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *GetRegistryFunctionNameTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryFunctionNameTagNotFound creates a GetRegistryFunctionNameTagNotFound with default headers values
func NewGetRegistryFunctionNameTagNotFound() *GetRegistryFunctionNameTagNotFound {
	return &GetRegistryFunctionNameTagNotFound{}
}

/*
GetRegistryFunctionNameTagNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRegistryFunctionNameTagNotFound struct {
	Payload string
}

// IsSuccess returns true when this get registry function name tag not found response has a 2xx status code
func (o *GetRegistryFunctionNameTagNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry function name tag not found response has a 3xx status code
func (o *GetRegistryFunctionNameTagNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry function name tag not found response has a 4xx status code
func (o *GetRegistryFunctionNameTagNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry function name tag not found response has a 5xx status code
func (o *GetRegistryFunctionNameTagNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry function name tag not found response a status code equal to that given
func (o *GetRegistryFunctionNameTagNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get registry function name tag not found response
func (o *GetRegistryFunctionNameTagNotFound) Code() int {
	return 404
}

func (o *GetRegistryFunctionNameTagNotFound) Error() string {
	return fmt.Sprintf("[GET /registry/function/{name}/{tag}][%d] getRegistryFunctionNameTagNotFound  %+v", 404, o.Payload)
}

func (o *GetRegistryFunctionNameTagNotFound) String() string {
	return fmt.Sprintf("[GET /registry/function/{name}/{tag}][%d] getRegistryFunctionNameTagNotFound  %+v", 404, o.Payload)
}

func (o *GetRegistryFunctionNameTagNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetRegistryFunctionNameTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryFunctionNameTagInternalServerError creates a GetRegistryFunctionNameTagInternalServerError with default headers values
func NewGetRegistryFunctionNameTagInternalServerError() *GetRegistryFunctionNameTagInternalServerError {
	return &GetRegistryFunctionNameTagInternalServerError{}
}

/*
GetRegistryFunctionNameTagInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRegistryFunctionNameTagInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get registry function name tag internal server error response has a 2xx status code
func (o *GetRegistryFunctionNameTagInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry function name tag internal server error response has a 3xx status code
func (o *GetRegistryFunctionNameTagInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry function name tag internal server error response has a 4xx status code
func (o *GetRegistryFunctionNameTagInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get registry function name tag internal server error response has a 5xx status code
func (o *GetRegistryFunctionNameTagInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get registry function name tag internal server error response a status code equal to that given
func (o *GetRegistryFunctionNameTagInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get registry function name tag internal server error response
func (o *GetRegistryFunctionNameTagInternalServerError) Code() int {
	return 500
}

func (o *GetRegistryFunctionNameTagInternalServerError) Error() string {
	return fmt.Sprintf("[GET /registry/function/{name}/{tag}][%d] getRegistryFunctionNameTagInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRegistryFunctionNameTagInternalServerError) String() string {
	return fmt.Sprintf("[GET /registry/function/{name}/{tag}][%d] getRegistryFunctionNameTagInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRegistryFunctionNameTagInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetRegistryFunctionNameTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}