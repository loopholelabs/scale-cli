/*
	Copyright 2022 Loophole Labs

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

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAuthGithubDeviceReader is a Reader for the GetAuthGithubDevice structure.
type GetAuthGithubDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthGithubDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewGetAuthGithubDeviceFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthGithubDeviceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthGithubDeviceFound creates a GetAuthGithubDeviceFound with default headers values
func NewGetAuthGithubDeviceFound() *GetAuthGithubDeviceFound {
	return &GetAuthGithubDeviceFound{}
}

/*
	GetAuthGithubDeviceFound describes a response with status code 302, with default header values.

Found
*/
type GetAuthGithubDeviceFound struct {
	Location string

	Payload string
}

// IsSuccess returns true when this get auth github device found response has a 2xx status code
func (o *GetAuthGithubDeviceFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get auth github device found response has a 3xx status code
func (o *GetAuthGithubDeviceFound) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get auth github device found response has a 4xx status code
func (o *GetAuthGithubDeviceFound) IsClientError() bool {
	return false
}

// IsServerError returns true when this get auth github device found response has a 5xx status code
func (o *GetAuthGithubDeviceFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get auth github device found response a status code equal to that given
func (o *GetAuthGithubDeviceFound) IsCode(code int) bool {
	return code == 302
}

func (o *GetAuthGithubDeviceFound) Error() string {
	return fmt.Sprintf("[GET /auth/github/device][%d] getAuthGithubDeviceFound  %+v", 302, o.Payload)
}

func (o *GetAuthGithubDeviceFound) String() string {
	return fmt.Sprintf("[GET /auth/github/device][%d] getAuthGithubDeviceFound  %+v", 302, o.Payload)
}

func (o *GetAuthGithubDeviceFound) GetPayload() string {
	return o.Payload
}

func (o *GetAuthGithubDeviceFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthGithubDeviceInternalServerError creates a GetAuthGithubDeviceInternalServerError with default headers values
func NewGetAuthGithubDeviceInternalServerError() *GetAuthGithubDeviceInternalServerError {
	return &GetAuthGithubDeviceInternalServerError{}
}

/*
	GetAuthGithubDeviceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAuthGithubDeviceInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get auth github device internal server error response has a 2xx status code
func (o *GetAuthGithubDeviceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get auth github device internal server error response has a 3xx status code
func (o *GetAuthGithubDeviceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get auth github device internal server error response has a 4xx status code
func (o *GetAuthGithubDeviceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get auth github device internal server error response has a 5xx status code
func (o *GetAuthGithubDeviceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get auth github device internal server error response a status code equal to that given
func (o *GetAuthGithubDeviceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAuthGithubDeviceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /auth/github/device][%d] getAuthGithubDeviceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthGithubDeviceInternalServerError) String() string {
	return fmt.Sprintf("[GET /auth/github/device][%d] getAuthGithubDeviceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthGithubDeviceInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetAuthGithubDeviceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
