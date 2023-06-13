// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// CountVMsReader is a Reader for the CountVMs structure.
type CountVMsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CountVMsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCountVMsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCountVMsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCountVMsOK creates a CountVMsOK with default headers values
func NewCountVMsOK() *CountVMsOK {
	return &CountVMsOK{}
}

/*
CountVMsOK describes a response with status code 200, with default header values.

A successful response.
*/
type CountVMsOK struct {
	Payload *models.CountVMsResponse
}

// IsSuccess returns true when this count v ms o k response has a 2xx status code
func (o *CountVMsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this count v ms o k response has a 3xx status code
func (o *CountVMsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this count v ms o k response has a 4xx status code
func (o *CountVMsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this count v ms o k response has a 5xx status code
func (o *CountVMsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this count v ms o k response a status code equal to that given
func (o *CountVMsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the count v ms o k response
func (o *CountVMsOK) Code() int {
	return 200
}

func (o *CountVMsOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/count-vms][%d] countVMsOK  %+v", 200, o.Payload)
}

func (o *CountVMsOK) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/count-vms][%d] countVMsOK  %+v", 200, o.Payload)
}

func (o *CountVMsOK) GetPayload() *models.CountVMsResponse {
	return o.Payload
}

func (o *CountVMsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CountVMsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCountVMsDefault creates a CountVMsDefault with default headers values
func NewCountVMsDefault(code int) *CountVMsDefault {
	return &CountVMsDefault{
		_statusCode: code,
	}
}

/*
CountVMsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CountVMsDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this count v ms default response has a 2xx status code
func (o *CountVMsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this count v ms default response has a 3xx status code
func (o *CountVMsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this count v ms default response has a 4xx status code
func (o *CountVMsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this count v ms default response has a 5xx status code
func (o *CountVMsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this count v ms default response a status code equal to that given
func (o *CountVMsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the count v ms default response
func (o *CountVMsDefault) Code() int {
	return o._statusCode
}

func (o *CountVMsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/count-vms][%d] CountVMs default  %+v", o._statusCode, o.Payload)
}

func (o *CountVMsDefault) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/count-vms][%d] CountVMs default  %+v", o._statusCode, o.Payload)
}

func (o *CountVMsDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *CountVMsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}