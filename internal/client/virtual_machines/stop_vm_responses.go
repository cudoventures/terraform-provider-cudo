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

// StopVMReader is a Reader for the StopVM structure.
type StopVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStopVMDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStopVMOK creates a StopVMOK with default headers values
func NewStopVMOK() *StopVMOK {
	return &StopVMOK{}
}

/*
StopVMOK describes a response with status code 200, with default header values.

A successful response.
*/
type StopVMOK struct {
	Payload models.StopVMResponse
}

// IsSuccess returns true when this stop Vm o k response has a 2xx status code
func (o *StopVMOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop Vm o k response has a 3xx status code
func (o *StopVMOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop Vm o k response has a 4xx status code
func (o *StopVMOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop Vm o k response has a 5xx status code
func (o *StopVMOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop Vm o k response a status code equal to that given
func (o *StopVMOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stop Vm o k response
func (o *StopVMOK) Code() int {
	return 200
}

func (o *StopVMOK) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/vms/{id}/stop][%d] stopVmOK  %+v", 200, o.Payload)
}

func (o *StopVMOK) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/vms/{id}/stop][%d] stopVmOK  %+v", 200, o.Payload)
}

func (o *StopVMOK) GetPayload() models.StopVMResponse {
	return o.Payload
}

func (o *StopVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopVMDefault creates a StopVMDefault with default headers values
func NewStopVMDefault(code int) *StopVMDefault {
	return &StopVMDefault{
		_statusCode: code,
	}
}

/*
StopVMDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StopVMDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this stop VM default response has a 2xx status code
func (o *StopVMDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this stop VM default response has a 3xx status code
func (o *StopVMDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this stop VM default response has a 4xx status code
func (o *StopVMDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this stop VM default response has a 5xx status code
func (o *StopVMDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this stop VM default response a status code equal to that given
func (o *StopVMDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the stop VM default response
func (o *StopVMDefault) Code() int {
	return o._statusCode
}

func (o *StopVMDefault) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/vms/{id}/stop][%d] StopVM default  %+v", o._statusCode, o.Payload)
}

func (o *StopVMDefault) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/vms/{id}/stop][%d] StopVM default  %+v", o._statusCode, o.Payload)
}

func (o *StopVMDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *StopVMDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
