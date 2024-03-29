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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/loopholelabs/scale-cli/client/models"
)

// GetDomainReader is a Reader for the GetDomain structure.
type GetDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDomainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDomainUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDomainNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDomainInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /domain] GetDomain", response, response.Code())
	}
}

// NewGetDomainOK creates a GetDomainOK with default headers values
func NewGetDomainOK() *GetDomainOK {
	return &GetDomainOK{}
}

/*
GetDomainOK describes a response with status code 200, with default header values.

OK
*/
type GetDomainOK struct {
	Payload []*models.ModelsDomainResponse
}

// IsSuccess returns true when this get domain o k response has a 2xx status code
func (o *GetDomainOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get domain o k response has a 3xx status code
func (o *GetDomainOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain o k response has a 4xx status code
func (o *GetDomainOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get domain o k response has a 5xx status code
func (o *GetDomainOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain o k response a status code equal to that given
func (o *GetDomainOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get domain o k response
func (o *GetDomainOK) Code() int {
	return 200
}

func (o *GetDomainOK) Error() string {
	return fmt.Sprintf("[GET /domain][%d] getDomainOK  %+v", 200, o.Payload)
}

func (o *GetDomainOK) String() string {
	return fmt.Sprintf("[GET /domain][%d] getDomainOK  %+v", 200, o.Payload)
}

func (o *GetDomainOK) GetPayload() []*models.ModelsDomainResponse {
	return o.Payload
}

func (o *GetDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainBadRequest creates a GetDomainBadRequest with default headers values
func NewGetDomainBadRequest() *GetDomainBadRequest {
	return &GetDomainBadRequest{}
}

/*
GetDomainBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDomainBadRequest struct {
	Payload string
}

// IsSuccess returns true when this get domain bad request response has a 2xx status code
func (o *GetDomainBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain bad request response has a 3xx status code
func (o *GetDomainBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain bad request response has a 4xx status code
func (o *GetDomainBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get domain bad request response has a 5xx status code
func (o *GetDomainBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain bad request response a status code equal to that given
func (o *GetDomainBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get domain bad request response
func (o *GetDomainBadRequest) Code() int {
	return 400
}

func (o *GetDomainBadRequest) Error() string {
	return fmt.Sprintf("[GET /domain][%d] getDomainBadRequest  %+v", 400, o.Payload)
}

func (o *GetDomainBadRequest) String() string {
	return fmt.Sprintf("[GET /domain][%d] getDomainBadRequest  %+v", 400, o.Payload)
}

func (o *GetDomainBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetDomainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainUnauthorized creates a GetDomainUnauthorized with default headers values
func NewGetDomainUnauthorized() *GetDomainUnauthorized {
	return &GetDomainUnauthorized{}
}

/*
GetDomainUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDomainUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this get domain unauthorized response has a 2xx status code
func (o *GetDomainUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain unauthorized response has a 3xx status code
func (o *GetDomainUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain unauthorized response has a 4xx status code
func (o *GetDomainUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get domain unauthorized response has a 5xx status code
func (o *GetDomainUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain unauthorized response a status code equal to that given
func (o *GetDomainUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get domain unauthorized response
func (o *GetDomainUnauthorized) Code() int {
	return 401
}

func (o *GetDomainUnauthorized) Error() string {
	return fmt.Sprintf("[GET /domain][%d] getDomainUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDomainUnauthorized) String() string {
	return fmt.Sprintf("[GET /domain][%d] getDomainUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDomainUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *GetDomainUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainNotFound creates a GetDomainNotFound with default headers values
func NewGetDomainNotFound() *GetDomainNotFound {
	return &GetDomainNotFound{}
}

/*
GetDomainNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDomainNotFound struct {
	Payload string
}

// IsSuccess returns true when this get domain not found response has a 2xx status code
func (o *GetDomainNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain not found response has a 3xx status code
func (o *GetDomainNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain not found response has a 4xx status code
func (o *GetDomainNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get domain not found response has a 5xx status code
func (o *GetDomainNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain not found response a status code equal to that given
func (o *GetDomainNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get domain not found response
func (o *GetDomainNotFound) Code() int {
	return 404
}

func (o *GetDomainNotFound) Error() string {
	return fmt.Sprintf("[GET /domain][%d] getDomainNotFound  %+v", 404, o.Payload)
}

func (o *GetDomainNotFound) String() string {
	return fmt.Sprintf("[GET /domain][%d] getDomainNotFound  %+v", 404, o.Payload)
}

func (o *GetDomainNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetDomainNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainInternalServerError creates a GetDomainInternalServerError with default headers values
func NewGetDomainInternalServerError() *GetDomainInternalServerError {
	return &GetDomainInternalServerError{}
}

/*
GetDomainInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDomainInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get domain internal server error response has a 2xx status code
func (o *GetDomainInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain internal server error response has a 3xx status code
func (o *GetDomainInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain internal server error response has a 4xx status code
func (o *GetDomainInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get domain internal server error response has a 5xx status code
func (o *GetDomainInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get domain internal server error response a status code equal to that given
func (o *GetDomainInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get domain internal server error response
func (o *GetDomainInternalServerError) Code() int {
	return 500
}

func (o *GetDomainInternalServerError) Error() string {
	return fmt.Sprintf("[GET /domain][%d] getDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDomainInternalServerError) String() string {
	return fmt.Sprintf("[GET /domain][%d] getDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDomainInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetDomainInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
