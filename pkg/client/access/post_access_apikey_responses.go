// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/loopholelabs/scale-cli/pkg/client/models"
)

// PostAccessApikeyReader is a Reader for the PostAccessApikey structure.
type PostAccessApikeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAccessApikeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAccessApikeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostAccessApikeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAccessApikeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAccessApikeyOK creates a PostAccessApikeyOK with default headers values
func NewPostAccessApikeyOK() *PostAccessApikeyOK {
	return &PostAccessApikeyOK{}
}

/*
	PostAccessApikeyOK describes a response with status code 200, with default header values.

OK
*/
type PostAccessApikeyOK struct {
	Payload *models.ModelsCreateAPIKeyResponse
}

// IsSuccess returns true when this post access apikey o k response has a 2xx status code
func (o *PostAccessApikeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post access apikey o k response has a 3xx status code
func (o *PostAccessApikeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post access apikey o k response has a 4xx status code
func (o *PostAccessApikeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post access apikey o k response has a 5xx status code
func (o *PostAccessApikeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post access apikey o k response a status code equal to that given
func (o *PostAccessApikeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostAccessApikeyOK) Error() string {
	return fmt.Sprintf("[POST /access/apikey][%d] postAccessApikeyOK  %+v", 200, o.Payload)
}

func (o *PostAccessApikeyOK) String() string {
	return fmt.Sprintf("[POST /access/apikey][%d] postAccessApikeyOK  %+v", 200, o.Payload)
}

func (o *PostAccessApikeyOK) GetPayload() *models.ModelsCreateAPIKeyResponse {
	return o.Payload
}

func (o *PostAccessApikeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsCreateAPIKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccessApikeyUnauthorized creates a PostAccessApikeyUnauthorized with default headers values
func NewPostAccessApikeyUnauthorized() *PostAccessApikeyUnauthorized {
	return &PostAccessApikeyUnauthorized{}
}

/*
	PostAccessApikeyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostAccessApikeyUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this post access apikey unauthorized response has a 2xx status code
func (o *PostAccessApikeyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post access apikey unauthorized response has a 3xx status code
func (o *PostAccessApikeyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post access apikey unauthorized response has a 4xx status code
func (o *PostAccessApikeyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post access apikey unauthorized response has a 5xx status code
func (o *PostAccessApikeyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post access apikey unauthorized response a status code equal to that given
func (o *PostAccessApikeyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostAccessApikeyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /access/apikey][%d] postAccessApikeyUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAccessApikeyUnauthorized) String() string {
	return fmt.Sprintf("[POST /access/apikey][%d] postAccessApikeyUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAccessApikeyUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *PostAccessApikeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccessApikeyInternalServerError creates a PostAccessApikeyInternalServerError with default headers values
func NewPostAccessApikeyInternalServerError() *PostAccessApikeyInternalServerError {
	return &PostAccessApikeyInternalServerError{}
}

/*
	PostAccessApikeyInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAccessApikeyInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this post access apikey internal server error response has a 2xx status code
func (o *PostAccessApikeyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post access apikey internal server error response has a 3xx status code
func (o *PostAccessApikeyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post access apikey internal server error response has a 4xx status code
func (o *PostAccessApikeyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post access apikey internal server error response has a 5xx status code
func (o *PostAccessApikeyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post access apikey internal server error response a status code equal to that given
func (o *PostAccessApikeyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostAccessApikeyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /access/apikey][%d] postAccessApikeyInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAccessApikeyInternalServerError) String() string {
	return fmt.Sprintf("[POST /access/apikey][%d] postAccessApikeyInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAccessApikeyInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *PostAccessApikeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
