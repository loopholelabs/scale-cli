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

// DeleteAccessApikeyIDReader is a Reader for the DeleteAccessApikeyID structure.
type DeleteAccessApikeyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAccessApikeyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAccessApikeyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAccessApikeyIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAccessApikeyIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAccessApikeyIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAccessApikeyIDOK creates a DeleteAccessApikeyIDOK with default headers values
func NewDeleteAccessApikeyIDOK() *DeleteAccessApikeyIDOK {
	return &DeleteAccessApikeyIDOK{}
}

/*
	DeleteAccessApikeyIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteAccessApikeyIDOK struct {
	Payload string
}

// IsSuccess returns true when this delete access apikey Id o k response has a 2xx status code
func (o *DeleteAccessApikeyIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete access apikey Id o k response has a 3xx status code
func (o *DeleteAccessApikeyIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete access apikey Id o k response has a 4xx status code
func (o *DeleteAccessApikeyIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete access apikey Id o k response has a 5xx status code
func (o *DeleteAccessApikeyIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete access apikey Id o k response a status code equal to that given
func (o *DeleteAccessApikeyIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteAccessApikeyIDOK) Error() string {
	return fmt.Sprintf("[DELETE /access/apikey/{id}][%d] deleteAccessApikeyIdOK  %+v", 200, o.Payload)
}

func (o *DeleteAccessApikeyIDOK) String() string {
	return fmt.Sprintf("[DELETE /access/apikey/{id}][%d] deleteAccessApikeyIdOK  %+v", 200, o.Payload)
}

func (o *DeleteAccessApikeyIDOK) GetPayload() string {
	return o.Payload
}

func (o *DeleteAccessApikeyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccessApikeyIDUnauthorized creates a DeleteAccessApikeyIDUnauthorized with default headers values
func NewDeleteAccessApikeyIDUnauthorized() *DeleteAccessApikeyIDUnauthorized {
	return &DeleteAccessApikeyIDUnauthorized{}
}

/*
	DeleteAccessApikeyIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteAccessApikeyIDUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this delete access apikey Id unauthorized response has a 2xx status code
func (o *DeleteAccessApikeyIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete access apikey Id unauthorized response has a 3xx status code
func (o *DeleteAccessApikeyIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete access apikey Id unauthorized response has a 4xx status code
func (o *DeleteAccessApikeyIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete access apikey Id unauthorized response has a 5xx status code
func (o *DeleteAccessApikeyIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete access apikey Id unauthorized response a status code equal to that given
func (o *DeleteAccessApikeyIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteAccessApikeyIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /access/apikey/{id}][%d] deleteAccessApikeyIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAccessApikeyIDUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /access/apikey/{id}][%d] deleteAccessApikeyIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAccessApikeyIDUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *DeleteAccessApikeyIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccessApikeyIDNotFound creates a DeleteAccessApikeyIDNotFound with default headers values
func NewDeleteAccessApikeyIDNotFound() *DeleteAccessApikeyIDNotFound {
	return &DeleteAccessApikeyIDNotFound{}
}

/*
	DeleteAccessApikeyIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteAccessApikeyIDNotFound struct {
	Payload string
}

// IsSuccess returns true when this delete access apikey Id not found response has a 2xx status code
func (o *DeleteAccessApikeyIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete access apikey Id not found response has a 3xx status code
func (o *DeleteAccessApikeyIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete access apikey Id not found response has a 4xx status code
func (o *DeleteAccessApikeyIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete access apikey Id not found response has a 5xx status code
func (o *DeleteAccessApikeyIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete access apikey Id not found response a status code equal to that given
func (o *DeleteAccessApikeyIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteAccessApikeyIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /access/apikey/{id}][%d] deleteAccessApikeyIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAccessApikeyIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /access/apikey/{id}][%d] deleteAccessApikeyIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAccessApikeyIDNotFound) GetPayload() string {
	return o.Payload
}

func (o *DeleteAccessApikeyIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccessApikeyIDInternalServerError creates a DeleteAccessApikeyIDInternalServerError with default headers values
func NewDeleteAccessApikeyIDInternalServerError() *DeleteAccessApikeyIDInternalServerError {
	return &DeleteAccessApikeyIDInternalServerError{}
}

/*
	DeleteAccessApikeyIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteAccessApikeyIDInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this delete access apikey Id internal server error response has a 2xx status code
func (o *DeleteAccessApikeyIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete access apikey Id internal server error response has a 3xx status code
func (o *DeleteAccessApikeyIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete access apikey Id internal server error response has a 4xx status code
func (o *DeleteAccessApikeyIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete access apikey Id internal server error response has a 5xx status code
func (o *DeleteAccessApikeyIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete access apikey Id internal server error response a status code equal to that given
func (o *DeleteAccessApikeyIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteAccessApikeyIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /access/apikey/{id}][%d] deleteAccessApikeyIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAccessApikeyIDInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /access/apikey/{id}][%d] deleteAccessApikeyIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAccessApikeyIDInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *DeleteAccessApikeyIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}