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

// GetRegistryFunctionOrgReader is a Reader for the GetRegistryFunctionOrg structure.
type GetRegistryFunctionOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegistryFunctionOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegistryFunctionOrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRegistryFunctionOrgBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRegistryFunctionOrgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRegistryFunctionOrgNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRegistryFunctionOrgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /registry/function/{org}] GetRegistryFunctionOrg", response, response.Code())
	}
}

// NewGetRegistryFunctionOrgOK creates a GetRegistryFunctionOrgOK with default headers values
func NewGetRegistryFunctionOrgOK() *GetRegistryFunctionOrgOK {
	return &GetRegistryFunctionOrgOK{}
}

/*
GetRegistryFunctionOrgOK describes a response with status code 200, with default header values.

OK
*/
type GetRegistryFunctionOrgOK struct {
	Payload []*models.ModelsFunctionResponse
}

// IsSuccess returns true when this get registry function org o k response has a 2xx status code
func (o *GetRegistryFunctionOrgOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get registry function org o k response has a 3xx status code
func (o *GetRegistryFunctionOrgOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry function org o k response has a 4xx status code
func (o *GetRegistryFunctionOrgOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get registry function org o k response has a 5xx status code
func (o *GetRegistryFunctionOrgOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry function org o k response a status code equal to that given
func (o *GetRegistryFunctionOrgOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get registry function org o k response
func (o *GetRegistryFunctionOrgOK) Code() int {
	return 200
}

func (o *GetRegistryFunctionOrgOK) Error() string {
	return fmt.Sprintf("[GET /registry/function/{org}][%d] getRegistryFunctionOrgOK  %+v", 200, o.Payload)
}

func (o *GetRegistryFunctionOrgOK) String() string {
	return fmt.Sprintf("[GET /registry/function/{org}][%d] getRegistryFunctionOrgOK  %+v", 200, o.Payload)
}

func (o *GetRegistryFunctionOrgOK) GetPayload() []*models.ModelsFunctionResponse {
	return o.Payload
}

func (o *GetRegistryFunctionOrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryFunctionOrgBadRequest creates a GetRegistryFunctionOrgBadRequest with default headers values
func NewGetRegistryFunctionOrgBadRequest() *GetRegistryFunctionOrgBadRequest {
	return &GetRegistryFunctionOrgBadRequest{}
}

/*
GetRegistryFunctionOrgBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetRegistryFunctionOrgBadRequest struct {
	Payload string
}

// IsSuccess returns true when this get registry function org bad request response has a 2xx status code
func (o *GetRegistryFunctionOrgBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry function org bad request response has a 3xx status code
func (o *GetRegistryFunctionOrgBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry function org bad request response has a 4xx status code
func (o *GetRegistryFunctionOrgBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry function org bad request response has a 5xx status code
func (o *GetRegistryFunctionOrgBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry function org bad request response a status code equal to that given
func (o *GetRegistryFunctionOrgBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get registry function org bad request response
func (o *GetRegistryFunctionOrgBadRequest) Code() int {
	return 400
}

func (o *GetRegistryFunctionOrgBadRequest) Error() string {
	return fmt.Sprintf("[GET /registry/function/{org}][%d] getRegistryFunctionOrgBadRequest  %+v", 400, o.Payload)
}

func (o *GetRegistryFunctionOrgBadRequest) String() string {
	return fmt.Sprintf("[GET /registry/function/{org}][%d] getRegistryFunctionOrgBadRequest  %+v", 400, o.Payload)
}

func (o *GetRegistryFunctionOrgBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetRegistryFunctionOrgBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryFunctionOrgUnauthorized creates a GetRegistryFunctionOrgUnauthorized with default headers values
func NewGetRegistryFunctionOrgUnauthorized() *GetRegistryFunctionOrgUnauthorized {
	return &GetRegistryFunctionOrgUnauthorized{}
}

/*
GetRegistryFunctionOrgUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRegistryFunctionOrgUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this get registry function org unauthorized response has a 2xx status code
func (o *GetRegistryFunctionOrgUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry function org unauthorized response has a 3xx status code
func (o *GetRegistryFunctionOrgUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry function org unauthorized response has a 4xx status code
func (o *GetRegistryFunctionOrgUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry function org unauthorized response has a 5xx status code
func (o *GetRegistryFunctionOrgUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry function org unauthorized response a status code equal to that given
func (o *GetRegistryFunctionOrgUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get registry function org unauthorized response
func (o *GetRegistryFunctionOrgUnauthorized) Code() int {
	return 401
}

func (o *GetRegistryFunctionOrgUnauthorized) Error() string {
	return fmt.Sprintf("[GET /registry/function/{org}][%d] getRegistryFunctionOrgUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRegistryFunctionOrgUnauthorized) String() string {
	return fmt.Sprintf("[GET /registry/function/{org}][%d] getRegistryFunctionOrgUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRegistryFunctionOrgUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *GetRegistryFunctionOrgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryFunctionOrgNotFound creates a GetRegistryFunctionOrgNotFound with default headers values
func NewGetRegistryFunctionOrgNotFound() *GetRegistryFunctionOrgNotFound {
	return &GetRegistryFunctionOrgNotFound{}
}

/*
GetRegistryFunctionOrgNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRegistryFunctionOrgNotFound struct {
	Payload string
}

// IsSuccess returns true when this get registry function org not found response has a 2xx status code
func (o *GetRegistryFunctionOrgNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry function org not found response has a 3xx status code
func (o *GetRegistryFunctionOrgNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry function org not found response has a 4xx status code
func (o *GetRegistryFunctionOrgNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry function org not found response has a 5xx status code
func (o *GetRegistryFunctionOrgNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry function org not found response a status code equal to that given
func (o *GetRegistryFunctionOrgNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get registry function org not found response
func (o *GetRegistryFunctionOrgNotFound) Code() int {
	return 404
}

func (o *GetRegistryFunctionOrgNotFound) Error() string {
	return fmt.Sprintf("[GET /registry/function/{org}][%d] getRegistryFunctionOrgNotFound  %+v", 404, o.Payload)
}

func (o *GetRegistryFunctionOrgNotFound) String() string {
	return fmt.Sprintf("[GET /registry/function/{org}][%d] getRegistryFunctionOrgNotFound  %+v", 404, o.Payload)
}

func (o *GetRegistryFunctionOrgNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetRegistryFunctionOrgNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryFunctionOrgInternalServerError creates a GetRegistryFunctionOrgInternalServerError with default headers values
func NewGetRegistryFunctionOrgInternalServerError() *GetRegistryFunctionOrgInternalServerError {
	return &GetRegistryFunctionOrgInternalServerError{}
}

/*
GetRegistryFunctionOrgInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRegistryFunctionOrgInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get registry function org internal server error response has a 2xx status code
func (o *GetRegistryFunctionOrgInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry function org internal server error response has a 3xx status code
func (o *GetRegistryFunctionOrgInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry function org internal server error response has a 4xx status code
func (o *GetRegistryFunctionOrgInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get registry function org internal server error response has a 5xx status code
func (o *GetRegistryFunctionOrgInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get registry function org internal server error response a status code equal to that given
func (o *GetRegistryFunctionOrgInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get registry function org internal server error response
func (o *GetRegistryFunctionOrgInternalServerError) Code() int {
	return 500
}

func (o *GetRegistryFunctionOrgInternalServerError) Error() string {
	return fmt.Sprintf("[GET /registry/function/{org}][%d] getRegistryFunctionOrgInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRegistryFunctionOrgInternalServerError) String() string {
	return fmt.Sprintf("[GET /registry/function/{org}][%d] getRegistryFunctionOrgInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRegistryFunctionOrgInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetRegistryFunctionOrgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
