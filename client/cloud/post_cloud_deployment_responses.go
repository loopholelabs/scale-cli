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

// PostCloudDeploymentReader is a Reader for the PostCloudDeployment structure.
type PostCloudDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCloudDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCloudDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostCloudDeploymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostCloudDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostCloudDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostCloudDeploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cloud/deployment] PostCloudDeployment", response, response.Code())
	}
}

// NewPostCloudDeploymentOK creates a PostCloudDeploymentOK with default headers values
func NewPostCloudDeploymentOK() *PostCloudDeploymentOK {
	return &PostCloudDeploymentOK{}
}

/*
PostCloudDeploymentOK describes a response with status code 200, with default header values.

OK
*/
type PostCloudDeploymentOK struct {
	Payload *models.ModelsDeploymentResponse
}

// IsSuccess returns true when this post cloud deployment o k response has a 2xx status code
func (o *PostCloudDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post cloud deployment o k response has a 3xx status code
func (o *PostCloudDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post cloud deployment o k response has a 4xx status code
func (o *PostCloudDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post cloud deployment o k response has a 5xx status code
func (o *PostCloudDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post cloud deployment o k response a status code equal to that given
func (o *PostCloudDeploymentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post cloud deployment o k response
func (o *PostCloudDeploymentOK) Code() int {
	return 200
}

func (o *PostCloudDeploymentOK) Error() string {
	return fmt.Sprintf("[POST /cloud/deployment][%d] postCloudDeploymentOK  %+v", 200, o.Payload)
}

func (o *PostCloudDeploymentOK) String() string {
	return fmt.Sprintf("[POST /cloud/deployment][%d] postCloudDeploymentOK  %+v", 200, o.Payload)
}

func (o *PostCloudDeploymentOK) GetPayload() *models.ModelsDeploymentResponse {
	return o.Payload
}

func (o *PostCloudDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsDeploymentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCloudDeploymentBadRequest creates a PostCloudDeploymentBadRequest with default headers values
func NewPostCloudDeploymentBadRequest() *PostCloudDeploymentBadRequest {
	return &PostCloudDeploymentBadRequest{}
}

/*
PostCloudDeploymentBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostCloudDeploymentBadRequest struct {
	Payload string
}

// IsSuccess returns true when this post cloud deployment bad request response has a 2xx status code
func (o *PostCloudDeploymentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post cloud deployment bad request response has a 3xx status code
func (o *PostCloudDeploymentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post cloud deployment bad request response has a 4xx status code
func (o *PostCloudDeploymentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post cloud deployment bad request response has a 5xx status code
func (o *PostCloudDeploymentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post cloud deployment bad request response a status code equal to that given
func (o *PostCloudDeploymentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post cloud deployment bad request response
func (o *PostCloudDeploymentBadRequest) Code() int {
	return 400
}

func (o *PostCloudDeploymentBadRequest) Error() string {
	return fmt.Sprintf("[POST /cloud/deployment][%d] postCloudDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *PostCloudDeploymentBadRequest) String() string {
	return fmt.Sprintf("[POST /cloud/deployment][%d] postCloudDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *PostCloudDeploymentBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PostCloudDeploymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCloudDeploymentUnauthorized creates a PostCloudDeploymentUnauthorized with default headers values
func NewPostCloudDeploymentUnauthorized() *PostCloudDeploymentUnauthorized {
	return &PostCloudDeploymentUnauthorized{}
}

/*
PostCloudDeploymentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostCloudDeploymentUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this post cloud deployment unauthorized response has a 2xx status code
func (o *PostCloudDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post cloud deployment unauthorized response has a 3xx status code
func (o *PostCloudDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post cloud deployment unauthorized response has a 4xx status code
func (o *PostCloudDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post cloud deployment unauthorized response has a 5xx status code
func (o *PostCloudDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post cloud deployment unauthorized response a status code equal to that given
func (o *PostCloudDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post cloud deployment unauthorized response
func (o *PostCloudDeploymentUnauthorized) Code() int {
	return 401
}

func (o *PostCloudDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cloud/deployment][%d] postCloudDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *PostCloudDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[POST /cloud/deployment][%d] postCloudDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *PostCloudDeploymentUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *PostCloudDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCloudDeploymentNotFound creates a PostCloudDeploymentNotFound with default headers values
func NewPostCloudDeploymentNotFound() *PostCloudDeploymentNotFound {
	return &PostCloudDeploymentNotFound{}
}

/*
PostCloudDeploymentNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostCloudDeploymentNotFound struct {
	Payload string
}

// IsSuccess returns true when this post cloud deployment not found response has a 2xx status code
func (o *PostCloudDeploymentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post cloud deployment not found response has a 3xx status code
func (o *PostCloudDeploymentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post cloud deployment not found response has a 4xx status code
func (o *PostCloudDeploymentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post cloud deployment not found response has a 5xx status code
func (o *PostCloudDeploymentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post cloud deployment not found response a status code equal to that given
func (o *PostCloudDeploymentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post cloud deployment not found response
func (o *PostCloudDeploymentNotFound) Code() int {
	return 404
}

func (o *PostCloudDeploymentNotFound) Error() string {
	return fmt.Sprintf("[POST /cloud/deployment][%d] postCloudDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *PostCloudDeploymentNotFound) String() string {
	return fmt.Sprintf("[POST /cloud/deployment][%d] postCloudDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *PostCloudDeploymentNotFound) GetPayload() string {
	return o.Payload
}

func (o *PostCloudDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCloudDeploymentInternalServerError creates a PostCloudDeploymentInternalServerError with default headers values
func NewPostCloudDeploymentInternalServerError() *PostCloudDeploymentInternalServerError {
	return &PostCloudDeploymentInternalServerError{}
}

/*
PostCloudDeploymentInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostCloudDeploymentInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this post cloud deployment internal server error response has a 2xx status code
func (o *PostCloudDeploymentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post cloud deployment internal server error response has a 3xx status code
func (o *PostCloudDeploymentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post cloud deployment internal server error response has a 4xx status code
func (o *PostCloudDeploymentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post cloud deployment internal server error response has a 5xx status code
func (o *PostCloudDeploymentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post cloud deployment internal server error response a status code equal to that given
func (o *PostCloudDeploymentInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post cloud deployment internal server error response
func (o *PostCloudDeploymentInternalServerError) Code() int {
	return 500
}

func (o *PostCloudDeploymentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cloud/deployment][%d] postCloudDeploymentInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCloudDeploymentInternalServerError) String() string {
	return fmt.Sprintf("[POST /cloud/deployment][%d] postCloudDeploymentInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCloudDeploymentInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *PostCloudDeploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
