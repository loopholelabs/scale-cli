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

	"github.com/loopholelabs/scale-cli/client/models"
)

// GetRegistryFunctionOrgNameTagReader is a Reader for the GetRegistryFunctionOrgNameTag structure.
type GetRegistryFunctionOrgNameTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegistryFunctionOrgNameTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegistryFunctionOrgNameTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRegistryFunctionOrgNameTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRegistryFunctionOrgNameTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRegistryFunctionOrgNameTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRegistryFunctionOrgNameTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /registry/function/{org}/{name}/{tag}] GetRegistryFunctionOrgNameTag", response, response.Code())
	}
}

// NewGetRegistryFunctionOrgNameTagOK creates a GetRegistryFunctionOrgNameTagOK with default headers values
func NewGetRegistryFunctionOrgNameTagOK() *GetRegistryFunctionOrgNameTagOK {
	return &GetRegistryFunctionOrgNameTagOK{}
}

/*
GetRegistryFunctionOrgNameTagOK describes a response with status code 200, with default header values.

OK
*/
type GetRegistryFunctionOrgNameTagOK struct {
	Payload *models.ModelsGetFunctionResponse
}

// IsSuccess returns true when this get registry function org name tag o k response has a 2xx status code
func (o *GetRegistryFunctionOrgNameTagOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get registry function org name tag o k response has a 3xx status code
func (o *GetRegistryFunctionOrgNameTagOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry function org name tag o k response has a 4xx status code
func (o *GetRegistryFunctionOrgNameTagOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get registry function org name tag o k response has a 5xx status code
func (o *GetRegistryFunctionOrgNameTagOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry function org name tag o k response a status code equal to that given
func (o *GetRegistryFunctionOrgNameTagOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get registry function org name tag o k response
func (o *GetRegistryFunctionOrgNameTagOK) Code() int {
	return 200
}

func (o *GetRegistryFunctionOrgNameTagOK) Error() string {
	return fmt.Sprintf("[GET /registry/function/{org}/{name}/{tag}][%d] getRegistryFunctionOrgNameTagOK  %+v", 200, o.Payload)
}

func (o *GetRegistryFunctionOrgNameTagOK) String() string {
	return fmt.Sprintf("[GET /registry/function/{org}/{name}/{tag}][%d] getRegistryFunctionOrgNameTagOK  %+v", 200, o.Payload)
}

func (o *GetRegistryFunctionOrgNameTagOK) GetPayload() *models.ModelsGetFunctionResponse {
	return o.Payload
}

func (o *GetRegistryFunctionOrgNameTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsGetFunctionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryFunctionOrgNameTagBadRequest creates a GetRegistryFunctionOrgNameTagBadRequest with default headers values
func NewGetRegistryFunctionOrgNameTagBadRequest() *GetRegistryFunctionOrgNameTagBadRequest {
	return &GetRegistryFunctionOrgNameTagBadRequest{}
}

/*
GetRegistryFunctionOrgNameTagBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetRegistryFunctionOrgNameTagBadRequest struct {
	Payload string
}

// IsSuccess returns true when this get registry function org name tag bad request response has a 2xx status code
func (o *GetRegistryFunctionOrgNameTagBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry function org name tag bad request response has a 3xx status code
func (o *GetRegistryFunctionOrgNameTagBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry function org name tag bad request response has a 4xx status code
func (o *GetRegistryFunctionOrgNameTagBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry function org name tag bad request response has a 5xx status code
func (o *GetRegistryFunctionOrgNameTagBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry function org name tag bad request response a status code equal to that given
func (o *GetRegistryFunctionOrgNameTagBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get registry function org name tag bad request response
func (o *GetRegistryFunctionOrgNameTagBadRequest) Code() int {
	return 400
}

func (o *GetRegistryFunctionOrgNameTagBadRequest) Error() string {
	return fmt.Sprintf("[GET /registry/function/{org}/{name}/{tag}][%d] getRegistryFunctionOrgNameTagBadRequest  %+v", 400, o.Payload)
}

func (o *GetRegistryFunctionOrgNameTagBadRequest) String() string {
	return fmt.Sprintf("[GET /registry/function/{org}/{name}/{tag}][%d] getRegistryFunctionOrgNameTagBadRequest  %+v", 400, o.Payload)
}

func (o *GetRegistryFunctionOrgNameTagBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetRegistryFunctionOrgNameTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryFunctionOrgNameTagUnauthorized creates a GetRegistryFunctionOrgNameTagUnauthorized with default headers values
func NewGetRegistryFunctionOrgNameTagUnauthorized() *GetRegistryFunctionOrgNameTagUnauthorized {
	return &GetRegistryFunctionOrgNameTagUnauthorized{}
}

/*
GetRegistryFunctionOrgNameTagUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRegistryFunctionOrgNameTagUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this get registry function org name tag unauthorized response has a 2xx status code
func (o *GetRegistryFunctionOrgNameTagUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry function org name tag unauthorized response has a 3xx status code
func (o *GetRegistryFunctionOrgNameTagUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry function org name tag unauthorized response has a 4xx status code
func (o *GetRegistryFunctionOrgNameTagUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry function org name tag unauthorized response has a 5xx status code
func (o *GetRegistryFunctionOrgNameTagUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry function org name tag unauthorized response a status code equal to that given
func (o *GetRegistryFunctionOrgNameTagUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get registry function org name tag unauthorized response
func (o *GetRegistryFunctionOrgNameTagUnauthorized) Code() int {
	return 401
}

func (o *GetRegistryFunctionOrgNameTagUnauthorized) Error() string {
	return fmt.Sprintf("[GET /registry/function/{org}/{name}/{tag}][%d] getRegistryFunctionOrgNameTagUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRegistryFunctionOrgNameTagUnauthorized) String() string {
	return fmt.Sprintf("[GET /registry/function/{org}/{name}/{tag}][%d] getRegistryFunctionOrgNameTagUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRegistryFunctionOrgNameTagUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *GetRegistryFunctionOrgNameTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryFunctionOrgNameTagNotFound creates a GetRegistryFunctionOrgNameTagNotFound with default headers values
func NewGetRegistryFunctionOrgNameTagNotFound() *GetRegistryFunctionOrgNameTagNotFound {
	return &GetRegistryFunctionOrgNameTagNotFound{}
}

/*
GetRegistryFunctionOrgNameTagNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRegistryFunctionOrgNameTagNotFound struct {
	Payload string
}

// IsSuccess returns true when this get registry function org name tag not found response has a 2xx status code
func (o *GetRegistryFunctionOrgNameTagNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry function org name tag not found response has a 3xx status code
func (o *GetRegistryFunctionOrgNameTagNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry function org name tag not found response has a 4xx status code
func (o *GetRegistryFunctionOrgNameTagNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry function org name tag not found response has a 5xx status code
func (o *GetRegistryFunctionOrgNameTagNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry function org name tag not found response a status code equal to that given
func (o *GetRegistryFunctionOrgNameTagNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get registry function org name tag not found response
func (o *GetRegistryFunctionOrgNameTagNotFound) Code() int {
	return 404
}

func (o *GetRegistryFunctionOrgNameTagNotFound) Error() string {
	return fmt.Sprintf("[GET /registry/function/{org}/{name}/{tag}][%d] getRegistryFunctionOrgNameTagNotFound  %+v", 404, o.Payload)
}

func (o *GetRegistryFunctionOrgNameTagNotFound) String() string {
	return fmt.Sprintf("[GET /registry/function/{org}/{name}/{tag}][%d] getRegistryFunctionOrgNameTagNotFound  %+v", 404, o.Payload)
}

func (o *GetRegistryFunctionOrgNameTagNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetRegistryFunctionOrgNameTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryFunctionOrgNameTagInternalServerError creates a GetRegistryFunctionOrgNameTagInternalServerError with default headers values
func NewGetRegistryFunctionOrgNameTagInternalServerError() *GetRegistryFunctionOrgNameTagInternalServerError {
	return &GetRegistryFunctionOrgNameTagInternalServerError{}
}

/*
GetRegistryFunctionOrgNameTagInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRegistryFunctionOrgNameTagInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get registry function org name tag internal server error response has a 2xx status code
func (o *GetRegistryFunctionOrgNameTagInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry function org name tag internal server error response has a 3xx status code
func (o *GetRegistryFunctionOrgNameTagInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry function org name tag internal server error response has a 4xx status code
func (o *GetRegistryFunctionOrgNameTagInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get registry function org name tag internal server error response has a 5xx status code
func (o *GetRegistryFunctionOrgNameTagInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get registry function org name tag internal server error response a status code equal to that given
func (o *GetRegistryFunctionOrgNameTagInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get registry function org name tag internal server error response
func (o *GetRegistryFunctionOrgNameTagInternalServerError) Code() int {
	return 500
}

func (o *GetRegistryFunctionOrgNameTagInternalServerError) Error() string {
	return fmt.Sprintf("[GET /registry/function/{org}/{name}/{tag}][%d] getRegistryFunctionOrgNameTagInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRegistryFunctionOrgNameTagInternalServerError) String() string {
	return fmt.Sprintf("[GET /registry/function/{org}/{name}/{tag}][%d] getRegistryFunctionOrgNameTagInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRegistryFunctionOrgNameTagInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetRegistryFunctionOrgNameTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
