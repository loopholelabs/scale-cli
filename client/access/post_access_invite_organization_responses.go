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

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostAccessInviteOrganizationReader is a Reader for the PostAccessInviteOrganization structure.
type PostAccessInviteOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAccessInviteOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAccessInviteOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAccessInviteOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAccessInviteOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAccessInviteOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /access/invite/{organization}] PostAccessInviteOrganization", response, response.Code())
	}
}

// NewPostAccessInviteOrganizationOK creates a PostAccessInviteOrganizationOK with default headers values
func NewPostAccessInviteOrganizationOK() *PostAccessInviteOrganizationOK {
	return &PostAccessInviteOrganizationOK{}
}

/*
PostAccessInviteOrganizationOK describes a response with status code 200, with default header values.

OK
*/
type PostAccessInviteOrganizationOK struct {
	Payload string
}

// IsSuccess returns true when this post access invite organization o k response has a 2xx status code
func (o *PostAccessInviteOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post access invite organization o k response has a 3xx status code
func (o *PostAccessInviteOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post access invite organization o k response has a 4xx status code
func (o *PostAccessInviteOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post access invite organization o k response has a 5xx status code
func (o *PostAccessInviteOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post access invite organization o k response a status code equal to that given
func (o *PostAccessInviteOrganizationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post access invite organization o k response
func (o *PostAccessInviteOrganizationOK) Code() int {
	return 200
}

func (o *PostAccessInviteOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /access/invite/{organization}][%d] postAccessInviteOrganizationOK  %+v", 200, o.Payload)
}

func (o *PostAccessInviteOrganizationOK) String() string {
	return fmt.Sprintf("[POST /access/invite/{organization}][%d] postAccessInviteOrganizationOK  %+v", 200, o.Payload)
}

func (o *PostAccessInviteOrganizationOK) GetPayload() string {
	return o.Payload
}

func (o *PostAccessInviteOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccessInviteOrganizationBadRequest creates a PostAccessInviteOrganizationBadRequest with default headers values
func NewPostAccessInviteOrganizationBadRequest() *PostAccessInviteOrganizationBadRequest {
	return &PostAccessInviteOrganizationBadRequest{}
}

/*
PostAccessInviteOrganizationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAccessInviteOrganizationBadRequest struct {
	Payload string
}

// IsSuccess returns true when this post access invite organization bad request response has a 2xx status code
func (o *PostAccessInviteOrganizationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post access invite organization bad request response has a 3xx status code
func (o *PostAccessInviteOrganizationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post access invite organization bad request response has a 4xx status code
func (o *PostAccessInviteOrganizationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post access invite organization bad request response has a 5xx status code
func (o *PostAccessInviteOrganizationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post access invite organization bad request response a status code equal to that given
func (o *PostAccessInviteOrganizationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post access invite organization bad request response
func (o *PostAccessInviteOrganizationBadRequest) Code() int {
	return 400
}

func (o *PostAccessInviteOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[POST /access/invite/{organization}][%d] postAccessInviteOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *PostAccessInviteOrganizationBadRequest) String() string {
	return fmt.Sprintf("[POST /access/invite/{organization}][%d] postAccessInviteOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *PostAccessInviteOrganizationBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PostAccessInviteOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccessInviteOrganizationUnauthorized creates a PostAccessInviteOrganizationUnauthorized with default headers values
func NewPostAccessInviteOrganizationUnauthorized() *PostAccessInviteOrganizationUnauthorized {
	return &PostAccessInviteOrganizationUnauthorized{}
}

/*
PostAccessInviteOrganizationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostAccessInviteOrganizationUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this post access invite organization unauthorized response has a 2xx status code
func (o *PostAccessInviteOrganizationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post access invite organization unauthorized response has a 3xx status code
func (o *PostAccessInviteOrganizationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post access invite organization unauthorized response has a 4xx status code
func (o *PostAccessInviteOrganizationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post access invite organization unauthorized response has a 5xx status code
func (o *PostAccessInviteOrganizationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post access invite organization unauthorized response a status code equal to that given
func (o *PostAccessInviteOrganizationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post access invite organization unauthorized response
func (o *PostAccessInviteOrganizationUnauthorized) Code() int {
	return 401
}

func (o *PostAccessInviteOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /access/invite/{organization}][%d] postAccessInviteOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAccessInviteOrganizationUnauthorized) String() string {
	return fmt.Sprintf("[POST /access/invite/{organization}][%d] postAccessInviteOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAccessInviteOrganizationUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *PostAccessInviteOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccessInviteOrganizationInternalServerError creates a PostAccessInviteOrganizationInternalServerError with default headers values
func NewPostAccessInviteOrganizationInternalServerError() *PostAccessInviteOrganizationInternalServerError {
	return &PostAccessInviteOrganizationInternalServerError{}
}

/*
PostAccessInviteOrganizationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAccessInviteOrganizationInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this post access invite organization internal server error response has a 2xx status code
func (o *PostAccessInviteOrganizationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post access invite organization internal server error response has a 3xx status code
func (o *PostAccessInviteOrganizationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post access invite organization internal server error response has a 4xx status code
func (o *PostAccessInviteOrganizationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post access invite organization internal server error response has a 5xx status code
func (o *PostAccessInviteOrganizationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post access invite organization internal server error response a status code equal to that given
func (o *PostAccessInviteOrganizationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post access invite organization internal server error response
func (o *PostAccessInviteOrganizationInternalServerError) Code() int {
	return 500
}

func (o *PostAccessInviteOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /access/invite/{organization}][%d] postAccessInviteOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAccessInviteOrganizationInternalServerError) String() string {
	return fmt.Sprintf("[POST /access/invite/{organization}][%d] postAccessInviteOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAccessInviteOrganizationInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *PostAccessInviteOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
