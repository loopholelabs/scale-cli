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

// PostRegistrySignatureReader is a Reader for the PostRegistrySignature structure.
type PostRegistrySignatureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRegistrySignatureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRegistrySignatureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRegistrySignatureBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRegistrySignatureUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRegistrySignatureNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewPostRegistrySignaturePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRegistrySignatureInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /registry/signature] PostRegistrySignature", response, response.Code())
	}
}

// NewPostRegistrySignatureOK creates a PostRegistrySignatureOK with default headers values
func NewPostRegistrySignatureOK() *PostRegistrySignatureOK {
	return &PostRegistrySignatureOK{}
}

/*
PostRegistrySignatureOK describes a response with status code 200, with default header values.

OK
*/
type PostRegistrySignatureOK struct {
	Payload *models.ModelsSignatureResponse
}

// IsSuccess returns true when this post registry signature o k response has a 2xx status code
func (o *PostRegistrySignatureOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post registry signature o k response has a 3xx status code
func (o *PostRegistrySignatureOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post registry signature o k response has a 4xx status code
func (o *PostRegistrySignatureOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post registry signature o k response has a 5xx status code
func (o *PostRegistrySignatureOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post registry signature o k response a status code equal to that given
func (o *PostRegistrySignatureOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post registry signature o k response
func (o *PostRegistrySignatureOK) Code() int {
	return 200
}

func (o *PostRegistrySignatureOK) Error() string {
	return fmt.Sprintf("[POST /registry/signature][%d] postRegistrySignatureOK  %+v", 200, o.Payload)
}

func (o *PostRegistrySignatureOK) String() string {
	return fmt.Sprintf("[POST /registry/signature][%d] postRegistrySignatureOK  %+v", 200, o.Payload)
}

func (o *PostRegistrySignatureOK) GetPayload() *models.ModelsSignatureResponse {
	return o.Payload
}

func (o *PostRegistrySignatureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsSignatureResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRegistrySignatureBadRequest creates a PostRegistrySignatureBadRequest with default headers values
func NewPostRegistrySignatureBadRequest() *PostRegistrySignatureBadRequest {
	return &PostRegistrySignatureBadRequest{}
}

/*
PostRegistrySignatureBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostRegistrySignatureBadRequest struct {
	Payload string
}

// IsSuccess returns true when this post registry signature bad request response has a 2xx status code
func (o *PostRegistrySignatureBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post registry signature bad request response has a 3xx status code
func (o *PostRegistrySignatureBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post registry signature bad request response has a 4xx status code
func (o *PostRegistrySignatureBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post registry signature bad request response has a 5xx status code
func (o *PostRegistrySignatureBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post registry signature bad request response a status code equal to that given
func (o *PostRegistrySignatureBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post registry signature bad request response
func (o *PostRegistrySignatureBadRequest) Code() int {
	return 400
}

func (o *PostRegistrySignatureBadRequest) Error() string {
	return fmt.Sprintf("[POST /registry/signature][%d] postRegistrySignatureBadRequest  %+v", 400, o.Payload)
}

func (o *PostRegistrySignatureBadRequest) String() string {
	return fmt.Sprintf("[POST /registry/signature][%d] postRegistrySignatureBadRequest  %+v", 400, o.Payload)
}

func (o *PostRegistrySignatureBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PostRegistrySignatureBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRegistrySignatureUnauthorized creates a PostRegistrySignatureUnauthorized with default headers values
func NewPostRegistrySignatureUnauthorized() *PostRegistrySignatureUnauthorized {
	return &PostRegistrySignatureUnauthorized{}
}

/*
PostRegistrySignatureUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostRegistrySignatureUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this post registry signature unauthorized response has a 2xx status code
func (o *PostRegistrySignatureUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post registry signature unauthorized response has a 3xx status code
func (o *PostRegistrySignatureUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post registry signature unauthorized response has a 4xx status code
func (o *PostRegistrySignatureUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post registry signature unauthorized response has a 5xx status code
func (o *PostRegistrySignatureUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post registry signature unauthorized response a status code equal to that given
func (o *PostRegistrySignatureUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post registry signature unauthorized response
func (o *PostRegistrySignatureUnauthorized) Code() int {
	return 401
}

func (o *PostRegistrySignatureUnauthorized) Error() string {
	return fmt.Sprintf("[POST /registry/signature][%d] postRegistrySignatureUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRegistrySignatureUnauthorized) String() string {
	return fmt.Sprintf("[POST /registry/signature][%d] postRegistrySignatureUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRegistrySignatureUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *PostRegistrySignatureUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRegistrySignatureNotFound creates a PostRegistrySignatureNotFound with default headers values
func NewPostRegistrySignatureNotFound() *PostRegistrySignatureNotFound {
	return &PostRegistrySignatureNotFound{}
}

/*
PostRegistrySignatureNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostRegistrySignatureNotFound struct {
	Payload string
}

// IsSuccess returns true when this post registry signature not found response has a 2xx status code
func (o *PostRegistrySignatureNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post registry signature not found response has a 3xx status code
func (o *PostRegistrySignatureNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post registry signature not found response has a 4xx status code
func (o *PostRegistrySignatureNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post registry signature not found response has a 5xx status code
func (o *PostRegistrySignatureNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post registry signature not found response a status code equal to that given
func (o *PostRegistrySignatureNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post registry signature not found response
func (o *PostRegistrySignatureNotFound) Code() int {
	return 404
}

func (o *PostRegistrySignatureNotFound) Error() string {
	return fmt.Sprintf("[POST /registry/signature][%d] postRegistrySignatureNotFound  %+v", 404, o.Payload)
}

func (o *PostRegistrySignatureNotFound) String() string {
	return fmt.Sprintf("[POST /registry/signature][%d] postRegistrySignatureNotFound  %+v", 404, o.Payload)
}

func (o *PostRegistrySignatureNotFound) GetPayload() string {
	return o.Payload
}

func (o *PostRegistrySignatureNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRegistrySignaturePreconditionFailed creates a PostRegistrySignaturePreconditionFailed with default headers values
func NewPostRegistrySignaturePreconditionFailed() *PostRegistrySignaturePreconditionFailed {
	return &PostRegistrySignaturePreconditionFailed{}
}

/*
PostRegistrySignaturePreconditionFailed describes a response with status code 412, with default header values.

Precondition Failed
*/
type PostRegistrySignaturePreconditionFailed struct {
	Payload string
}

// IsSuccess returns true when this post registry signature precondition failed response has a 2xx status code
func (o *PostRegistrySignaturePreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post registry signature precondition failed response has a 3xx status code
func (o *PostRegistrySignaturePreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post registry signature precondition failed response has a 4xx status code
func (o *PostRegistrySignaturePreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this post registry signature precondition failed response has a 5xx status code
func (o *PostRegistrySignaturePreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this post registry signature precondition failed response a status code equal to that given
func (o *PostRegistrySignaturePreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the post registry signature precondition failed response
func (o *PostRegistrySignaturePreconditionFailed) Code() int {
	return 412
}

func (o *PostRegistrySignaturePreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /registry/signature][%d] postRegistrySignaturePreconditionFailed  %+v", 412, o.Payload)
}

func (o *PostRegistrySignaturePreconditionFailed) String() string {
	return fmt.Sprintf("[POST /registry/signature][%d] postRegistrySignaturePreconditionFailed  %+v", 412, o.Payload)
}

func (o *PostRegistrySignaturePreconditionFailed) GetPayload() string {
	return o.Payload
}

func (o *PostRegistrySignaturePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRegistrySignatureInternalServerError creates a PostRegistrySignatureInternalServerError with default headers values
func NewPostRegistrySignatureInternalServerError() *PostRegistrySignatureInternalServerError {
	return &PostRegistrySignatureInternalServerError{}
}

/*
PostRegistrySignatureInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostRegistrySignatureInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this post registry signature internal server error response has a 2xx status code
func (o *PostRegistrySignatureInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post registry signature internal server error response has a 3xx status code
func (o *PostRegistrySignatureInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post registry signature internal server error response has a 4xx status code
func (o *PostRegistrySignatureInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post registry signature internal server error response has a 5xx status code
func (o *PostRegistrySignatureInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post registry signature internal server error response a status code equal to that given
func (o *PostRegistrySignatureInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post registry signature internal server error response
func (o *PostRegistrySignatureInternalServerError) Code() int {
	return 500
}

func (o *PostRegistrySignatureInternalServerError) Error() string {
	return fmt.Sprintf("[POST /registry/signature][%d] postRegistrySignatureInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRegistrySignatureInternalServerError) String() string {
	return fmt.Sprintf("[POST /registry/signature][%d] postRegistrySignatureInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRegistrySignatureInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *PostRegistrySignatureInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
