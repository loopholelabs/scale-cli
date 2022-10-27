// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/loopholelabs/scale-cli/pkg/client/models"
)

// GetConfigBuildReader is a Reader for the GetConfigBuild structure.
type GetConfigBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetConfigBuildInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConfigBuildOK creates a GetConfigBuildOK with default headers values
func NewGetConfigBuildOK() *GetConfigBuildOK {
	return &GetConfigBuildOK{}
}

/*
	GetConfigBuildOK describes a response with status code 200, with default header values.

OK
*/
type GetConfigBuildOK struct {
	Payload *models.ModelsBuildResponse
}

// IsSuccess returns true when this get config build o k response has a 2xx status code
func (o *GetConfigBuildOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get config build o k response has a 3xx status code
func (o *GetConfigBuildOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get config build o k response has a 4xx status code
func (o *GetConfigBuildOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get config build o k response has a 5xx status code
func (o *GetConfigBuildOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get config build o k response a status code equal to that given
func (o *GetConfigBuildOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConfigBuildOK) Error() string {
	return fmt.Sprintf("[GET /config/build][%d] getConfigBuildOK  %+v", 200, o.Payload)
}

func (o *GetConfigBuildOK) String() string {
	return fmt.Sprintf("[GET /config/build][%d] getConfigBuildOK  %+v", 200, o.Payload)
}

func (o *GetConfigBuildOK) GetPayload() *models.ModelsBuildResponse {
	return o.Payload
}

func (o *GetConfigBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsBuildResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigBuildInternalServerError creates a GetConfigBuildInternalServerError with default headers values
func NewGetConfigBuildInternalServerError() *GetConfigBuildInternalServerError {
	return &GetConfigBuildInternalServerError{}
}

/*
	GetConfigBuildInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetConfigBuildInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get config build internal server error response has a 2xx status code
func (o *GetConfigBuildInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get config build internal server error response has a 3xx status code
func (o *GetConfigBuildInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get config build internal server error response has a 4xx status code
func (o *GetConfigBuildInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get config build internal server error response has a 5xx status code
func (o *GetConfigBuildInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get config build internal server error response a status code equal to that given
func (o *GetConfigBuildInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConfigBuildInternalServerError) Error() string {
	return fmt.Sprintf("[GET /config/build][%d] getConfigBuildInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConfigBuildInternalServerError) String() string {
	return fmt.Sprintf("[GET /config/build][%d] getConfigBuildInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConfigBuildInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetConfigBuildInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
