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
)

// DeleteRegistrySignatureOrgNameTagReader is a Reader for the DeleteRegistrySignatureOrgNameTag structure.
type DeleteRegistrySignatureOrgNameTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRegistrySignatureOrgNameTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRegistrySignatureOrgNameTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRegistrySignatureOrgNameTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRegistrySignatureOrgNameTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRegistrySignatureOrgNameTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewDeleteRegistrySignatureOrgNameTagPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRegistrySignatureOrgNameTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /registry/signature/{org}/{name}/{tag}] DeleteRegistrySignatureOrgNameTag", response, response.Code())
	}
}

// NewDeleteRegistrySignatureOrgNameTagOK creates a DeleteRegistrySignatureOrgNameTagOK with default headers values
func NewDeleteRegistrySignatureOrgNameTagOK() *DeleteRegistrySignatureOrgNameTagOK {
	return &DeleteRegistrySignatureOrgNameTagOK{}
}

/*
DeleteRegistrySignatureOrgNameTagOK describes a response with status code 200, with default header values.

OK
*/
type DeleteRegistrySignatureOrgNameTagOK struct {
	Payload string
}

// IsSuccess returns true when this delete registry signature org name tag o k response has a 2xx status code
func (o *DeleteRegistrySignatureOrgNameTagOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete registry signature org name tag o k response has a 3xx status code
func (o *DeleteRegistrySignatureOrgNameTagOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry signature org name tag o k response has a 4xx status code
func (o *DeleteRegistrySignatureOrgNameTagOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete registry signature org name tag o k response has a 5xx status code
func (o *DeleteRegistrySignatureOrgNameTagOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry signature org name tag o k response a status code equal to that given
func (o *DeleteRegistrySignatureOrgNameTagOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete registry signature org name tag o k response
func (o *DeleteRegistrySignatureOrgNameTagOK) Code() int {
	return 200
}

func (o *DeleteRegistrySignatureOrgNameTagOK) Error() string {
	return fmt.Sprintf("[DELETE /registry/signature/{org}/{name}/{tag}][%d] deleteRegistrySignatureOrgNameTagOK  %+v", 200, o.Payload)
}

func (o *DeleteRegistrySignatureOrgNameTagOK) String() string {
	return fmt.Sprintf("[DELETE /registry/signature/{org}/{name}/{tag}][%d] deleteRegistrySignatureOrgNameTagOK  %+v", 200, o.Payload)
}

func (o *DeleteRegistrySignatureOrgNameTagOK) GetPayload() string {
	return o.Payload
}

func (o *DeleteRegistrySignatureOrgNameTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegistrySignatureOrgNameTagBadRequest creates a DeleteRegistrySignatureOrgNameTagBadRequest with default headers values
func NewDeleteRegistrySignatureOrgNameTagBadRequest() *DeleteRegistrySignatureOrgNameTagBadRequest {
	return &DeleteRegistrySignatureOrgNameTagBadRequest{}
}

/*
DeleteRegistrySignatureOrgNameTagBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteRegistrySignatureOrgNameTagBadRequest struct {
	Payload string
}

// IsSuccess returns true when this delete registry signature org name tag bad request response has a 2xx status code
func (o *DeleteRegistrySignatureOrgNameTagBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry signature org name tag bad request response has a 3xx status code
func (o *DeleteRegistrySignatureOrgNameTagBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry signature org name tag bad request response has a 4xx status code
func (o *DeleteRegistrySignatureOrgNameTagBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry signature org name tag bad request response has a 5xx status code
func (o *DeleteRegistrySignatureOrgNameTagBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry signature org name tag bad request response a status code equal to that given
func (o *DeleteRegistrySignatureOrgNameTagBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete registry signature org name tag bad request response
func (o *DeleteRegistrySignatureOrgNameTagBadRequest) Code() int {
	return 400
}

func (o *DeleteRegistrySignatureOrgNameTagBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /registry/signature/{org}/{name}/{tag}][%d] deleteRegistrySignatureOrgNameTagBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRegistrySignatureOrgNameTagBadRequest) String() string {
	return fmt.Sprintf("[DELETE /registry/signature/{org}/{name}/{tag}][%d] deleteRegistrySignatureOrgNameTagBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRegistrySignatureOrgNameTagBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteRegistrySignatureOrgNameTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegistrySignatureOrgNameTagUnauthorized creates a DeleteRegistrySignatureOrgNameTagUnauthorized with default headers values
func NewDeleteRegistrySignatureOrgNameTagUnauthorized() *DeleteRegistrySignatureOrgNameTagUnauthorized {
	return &DeleteRegistrySignatureOrgNameTagUnauthorized{}
}

/*
DeleteRegistrySignatureOrgNameTagUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteRegistrySignatureOrgNameTagUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this delete registry signature org name tag unauthorized response has a 2xx status code
func (o *DeleteRegistrySignatureOrgNameTagUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry signature org name tag unauthorized response has a 3xx status code
func (o *DeleteRegistrySignatureOrgNameTagUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry signature org name tag unauthorized response has a 4xx status code
func (o *DeleteRegistrySignatureOrgNameTagUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry signature org name tag unauthorized response has a 5xx status code
func (o *DeleteRegistrySignatureOrgNameTagUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry signature org name tag unauthorized response a status code equal to that given
func (o *DeleteRegistrySignatureOrgNameTagUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete registry signature org name tag unauthorized response
func (o *DeleteRegistrySignatureOrgNameTagUnauthorized) Code() int {
	return 401
}

func (o *DeleteRegistrySignatureOrgNameTagUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /registry/signature/{org}/{name}/{tag}][%d] deleteRegistrySignatureOrgNameTagUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRegistrySignatureOrgNameTagUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /registry/signature/{org}/{name}/{tag}][%d] deleteRegistrySignatureOrgNameTagUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRegistrySignatureOrgNameTagUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *DeleteRegistrySignatureOrgNameTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegistrySignatureOrgNameTagNotFound creates a DeleteRegistrySignatureOrgNameTagNotFound with default headers values
func NewDeleteRegistrySignatureOrgNameTagNotFound() *DeleteRegistrySignatureOrgNameTagNotFound {
	return &DeleteRegistrySignatureOrgNameTagNotFound{}
}

/*
DeleteRegistrySignatureOrgNameTagNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteRegistrySignatureOrgNameTagNotFound struct {
	Payload string
}

// IsSuccess returns true when this delete registry signature org name tag not found response has a 2xx status code
func (o *DeleteRegistrySignatureOrgNameTagNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry signature org name tag not found response has a 3xx status code
func (o *DeleteRegistrySignatureOrgNameTagNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry signature org name tag not found response has a 4xx status code
func (o *DeleteRegistrySignatureOrgNameTagNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry signature org name tag not found response has a 5xx status code
func (o *DeleteRegistrySignatureOrgNameTagNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry signature org name tag not found response a status code equal to that given
func (o *DeleteRegistrySignatureOrgNameTagNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete registry signature org name tag not found response
func (o *DeleteRegistrySignatureOrgNameTagNotFound) Code() int {
	return 404
}

func (o *DeleteRegistrySignatureOrgNameTagNotFound) Error() string {
	return fmt.Sprintf("[DELETE /registry/signature/{org}/{name}/{tag}][%d] deleteRegistrySignatureOrgNameTagNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRegistrySignatureOrgNameTagNotFound) String() string {
	return fmt.Sprintf("[DELETE /registry/signature/{org}/{name}/{tag}][%d] deleteRegistrySignatureOrgNameTagNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRegistrySignatureOrgNameTagNotFound) GetPayload() string {
	return o.Payload
}

func (o *DeleteRegistrySignatureOrgNameTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegistrySignatureOrgNameTagPreconditionFailed creates a DeleteRegistrySignatureOrgNameTagPreconditionFailed with default headers values
func NewDeleteRegistrySignatureOrgNameTagPreconditionFailed() *DeleteRegistrySignatureOrgNameTagPreconditionFailed {
	return &DeleteRegistrySignatureOrgNameTagPreconditionFailed{}
}

/*
DeleteRegistrySignatureOrgNameTagPreconditionFailed describes a response with status code 412, with default header values.

Precondition Failed
*/
type DeleteRegistrySignatureOrgNameTagPreconditionFailed struct {
	Payload string
}

// IsSuccess returns true when this delete registry signature org name tag precondition failed response has a 2xx status code
func (o *DeleteRegistrySignatureOrgNameTagPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry signature org name tag precondition failed response has a 3xx status code
func (o *DeleteRegistrySignatureOrgNameTagPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry signature org name tag precondition failed response has a 4xx status code
func (o *DeleteRegistrySignatureOrgNameTagPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry signature org name tag precondition failed response has a 5xx status code
func (o *DeleteRegistrySignatureOrgNameTagPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry signature org name tag precondition failed response a status code equal to that given
func (o *DeleteRegistrySignatureOrgNameTagPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the delete registry signature org name tag precondition failed response
func (o *DeleteRegistrySignatureOrgNameTagPreconditionFailed) Code() int {
	return 412
}

func (o *DeleteRegistrySignatureOrgNameTagPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /registry/signature/{org}/{name}/{tag}][%d] deleteRegistrySignatureOrgNameTagPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteRegistrySignatureOrgNameTagPreconditionFailed) String() string {
	return fmt.Sprintf("[DELETE /registry/signature/{org}/{name}/{tag}][%d] deleteRegistrySignatureOrgNameTagPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteRegistrySignatureOrgNameTagPreconditionFailed) GetPayload() string {
	return o.Payload
}

func (o *DeleteRegistrySignatureOrgNameTagPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegistrySignatureOrgNameTagInternalServerError creates a DeleteRegistrySignatureOrgNameTagInternalServerError with default headers values
func NewDeleteRegistrySignatureOrgNameTagInternalServerError() *DeleteRegistrySignatureOrgNameTagInternalServerError {
	return &DeleteRegistrySignatureOrgNameTagInternalServerError{}
}

/*
DeleteRegistrySignatureOrgNameTagInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteRegistrySignatureOrgNameTagInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this delete registry signature org name tag internal server error response has a 2xx status code
func (o *DeleteRegistrySignatureOrgNameTagInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry signature org name tag internal server error response has a 3xx status code
func (o *DeleteRegistrySignatureOrgNameTagInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry signature org name tag internal server error response has a 4xx status code
func (o *DeleteRegistrySignatureOrgNameTagInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete registry signature org name tag internal server error response has a 5xx status code
func (o *DeleteRegistrySignatureOrgNameTagInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete registry signature org name tag internal server error response a status code equal to that given
func (o *DeleteRegistrySignatureOrgNameTagInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete registry signature org name tag internal server error response
func (o *DeleteRegistrySignatureOrgNameTagInternalServerError) Code() int {
	return 500
}

func (o *DeleteRegistrySignatureOrgNameTagInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /registry/signature/{org}/{name}/{tag}][%d] deleteRegistrySignatureOrgNameTagInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRegistrySignatureOrgNameTagInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /registry/signature/{org}/{name}/{tag}][%d] deleteRegistrySignatureOrgNameTagInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRegistrySignatureOrgNameTagInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *DeleteRegistrySignatureOrgNameTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
