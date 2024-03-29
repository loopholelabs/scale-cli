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
)

// DeleteDomainDomainReader is a Reader for the DeleteDomainDomain structure.
type DeleteDomainDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDomainDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDomainDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDomainDomainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteDomainDomainUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDomainDomainNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewDeleteDomainDomainPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDomainDomainInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /domain/{domain}] DeleteDomainDomain", response, response.Code())
	}
}

// NewDeleteDomainDomainOK creates a DeleteDomainDomainOK with default headers values
func NewDeleteDomainDomainOK() *DeleteDomainDomainOK {
	return &DeleteDomainDomainOK{}
}

/*
DeleteDomainDomainOK describes a response with status code 200, with default header values.

OK
*/
type DeleteDomainDomainOK struct {
	Payload string
}

// IsSuccess returns true when this delete domain domain o k response has a 2xx status code
func (o *DeleteDomainDomainOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete domain domain o k response has a 3xx status code
func (o *DeleteDomainDomainOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete domain domain o k response has a 4xx status code
func (o *DeleteDomainDomainOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete domain domain o k response has a 5xx status code
func (o *DeleteDomainDomainOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete domain domain o k response a status code equal to that given
func (o *DeleteDomainDomainOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete domain domain o k response
func (o *DeleteDomainDomainOK) Code() int {
	return 200
}

func (o *DeleteDomainDomainOK) Error() string {
	return fmt.Sprintf("[DELETE /domain/{domain}][%d] deleteDomainDomainOK  %+v", 200, o.Payload)
}

func (o *DeleteDomainDomainOK) String() string {
	return fmt.Sprintf("[DELETE /domain/{domain}][%d] deleteDomainDomainOK  %+v", 200, o.Payload)
}

func (o *DeleteDomainDomainOK) GetPayload() string {
	return o.Payload
}

func (o *DeleteDomainDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDomainDomainBadRequest creates a DeleteDomainDomainBadRequest with default headers values
func NewDeleteDomainDomainBadRequest() *DeleteDomainDomainBadRequest {
	return &DeleteDomainDomainBadRequest{}
}

/*
DeleteDomainDomainBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteDomainDomainBadRequest struct {
	Payload string
}

// IsSuccess returns true when this delete domain domain bad request response has a 2xx status code
func (o *DeleteDomainDomainBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete domain domain bad request response has a 3xx status code
func (o *DeleteDomainDomainBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete domain domain bad request response has a 4xx status code
func (o *DeleteDomainDomainBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete domain domain bad request response has a 5xx status code
func (o *DeleteDomainDomainBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete domain domain bad request response a status code equal to that given
func (o *DeleteDomainDomainBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete domain domain bad request response
func (o *DeleteDomainDomainBadRequest) Code() int {
	return 400
}

func (o *DeleteDomainDomainBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /domain/{domain}][%d] deleteDomainDomainBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDomainDomainBadRequest) String() string {
	return fmt.Sprintf("[DELETE /domain/{domain}][%d] deleteDomainDomainBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDomainDomainBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteDomainDomainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDomainDomainUnauthorized creates a DeleteDomainDomainUnauthorized with default headers values
func NewDeleteDomainDomainUnauthorized() *DeleteDomainDomainUnauthorized {
	return &DeleteDomainDomainUnauthorized{}
}

/*
DeleteDomainDomainUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteDomainDomainUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this delete domain domain unauthorized response has a 2xx status code
func (o *DeleteDomainDomainUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete domain domain unauthorized response has a 3xx status code
func (o *DeleteDomainDomainUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete domain domain unauthorized response has a 4xx status code
func (o *DeleteDomainDomainUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete domain domain unauthorized response has a 5xx status code
func (o *DeleteDomainDomainUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete domain domain unauthorized response a status code equal to that given
func (o *DeleteDomainDomainUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete domain domain unauthorized response
func (o *DeleteDomainDomainUnauthorized) Code() int {
	return 401
}

func (o *DeleteDomainDomainUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /domain/{domain}][%d] deleteDomainDomainUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteDomainDomainUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /domain/{domain}][%d] deleteDomainDomainUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteDomainDomainUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *DeleteDomainDomainUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDomainDomainNotFound creates a DeleteDomainDomainNotFound with default headers values
func NewDeleteDomainDomainNotFound() *DeleteDomainDomainNotFound {
	return &DeleteDomainDomainNotFound{}
}

/*
DeleteDomainDomainNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteDomainDomainNotFound struct {
	Payload string
}

// IsSuccess returns true when this delete domain domain not found response has a 2xx status code
func (o *DeleteDomainDomainNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete domain domain not found response has a 3xx status code
func (o *DeleteDomainDomainNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete domain domain not found response has a 4xx status code
func (o *DeleteDomainDomainNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete domain domain not found response has a 5xx status code
func (o *DeleteDomainDomainNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete domain domain not found response a status code equal to that given
func (o *DeleteDomainDomainNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete domain domain not found response
func (o *DeleteDomainDomainNotFound) Code() int {
	return 404
}

func (o *DeleteDomainDomainNotFound) Error() string {
	return fmt.Sprintf("[DELETE /domain/{domain}][%d] deleteDomainDomainNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDomainDomainNotFound) String() string {
	return fmt.Sprintf("[DELETE /domain/{domain}][%d] deleteDomainDomainNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDomainDomainNotFound) GetPayload() string {
	return o.Payload
}

func (o *DeleteDomainDomainNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDomainDomainPreconditionFailed creates a DeleteDomainDomainPreconditionFailed with default headers values
func NewDeleteDomainDomainPreconditionFailed() *DeleteDomainDomainPreconditionFailed {
	return &DeleteDomainDomainPreconditionFailed{}
}

/*
DeleteDomainDomainPreconditionFailed describes a response with status code 412, with default header values.

Precondition Failed
*/
type DeleteDomainDomainPreconditionFailed struct {
	Payload string
}

// IsSuccess returns true when this delete domain domain precondition failed response has a 2xx status code
func (o *DeleteDomainDomainPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete domain domain precondition failed response has a 3xx status code
func (o *DeleteDomainDomainPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete domain domain precondition failed response has a 4xx status code
func (o *DeleteDomainDomainPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete domain domain precondition failed response has a 5xx status code
func (o *DeleteDomainDomainPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this delete domain domain precondition failed response a status code equal to that given
func (o *DeleteDomainDomainPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the delete domain domain precondition failed response
func (o *DeleteDomainDomainPreconditionFailed) Code() int {
	return 412
}

func (o *DeleteDomainDomainPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /domain/{domain}][%d] deleteDomainDomainPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteDomainDomainPreconditionFailed) String() string {
	return fmt.Sprintf("[DELETE /domain/{domain}][%d] deleteDomainDomainPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteDomainDomainPreconditionFailed) GetPayload() string {
	return o.Payload
}

func (o *DeleteDomainDomainPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDomainDomainInternalServerError creates a DeleteDomainDomainInternalServerError with default headers values
func NewDeleteDomainDomainInternalServerError() *DeleteDomainDomainInternalServerError {
	return &DeleteDomainDomainInternalServerError{}
}

/*
DeleteDomainDomainInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteDomainDomainInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this delete domain domain internal server error response has a 2xx status code
func (o *DeleteDomainDomainInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete domain domain internal server error response has a 3xx status code
func (o *DeleteDomainDomainInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete domain domain internal server error response has a 4xx status code
func (o *DeleteDomainDomainInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete domain domain internal server error response has a 5xx status code
func (o *DeleteDomainDomainInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete domain domain internal server error response a status code equal to that given
func (o *DeleteDomainDomainInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete domain domain internal server error response
func (o *DeleteDomainDomainInternalServerError) Code() int {
	return 500
}

func (o *DeleteDomainDomainInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /domain/{domain}][%d] deleteDomainDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDomainDomainInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /domain/{domain}][%d] deleteDomainDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDomainDomainInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *DeleteDomainDomainInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
