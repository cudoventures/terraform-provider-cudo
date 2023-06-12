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

// ListPublicVMImagesReader is a Reader for the ListPublicVMImages structure.
type ListPublicVMImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPublicVMImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPublicVMImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListPublicVMImagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListPublicVMImagesOK creates a ListPublicVMImagesOK with default headers values
func NewListPublicVMImagesOK() *ListPublicVMImagesOK {
	return &ListPublicVMImagesOK{}
}

/*
ListPublicVMImagesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListPublicVMImagesOK struct {
	Payload *models.ListPublicVMImagesResponse
}

// IsSuccess returns true when this list public Vm images o k response has a 2xx status code
func (o *ListPublicVMImagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list public Vm images o k response has a 3xx status code
func (o *ListPublicVMImagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list public Vm images o k response has a 4xx status code
func (o *ListPublicVMImagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list public Vm images o k response has a 5xx status code
func (o *ListPublicVMImagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list public Vm images o k response a status code equal to that given
func (o *ListPublicVMImagesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list public Vm images o k response
func (o *ListPublicVMImagesOK) Code() int {
	return 200
}

func (o *ListPublicVMImagesOK) Error() string {
	return fmt.Sprintf("[GET /v1/vms/public-images][%d] listPublicVmImagesOK  %+v", 200, o.Payload)
}

func (o *ListPublicVMImagesOK) String() string {
	return fmt.Sprintf("[GET /v1/vms/public-images][%d] listPublicVmImagesOK  %+v", 200, o.Payload)
}

func (o *ListPublicVMImagesOK) GetPayload() *models.ListPublicVMImagesResponse {
	return o.Payload
}

func (o *ListPublicVMImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListPublicVMImagesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPublicVMImagesDefault creates a ListPublicVMImagesDefault with default headers values
func NewListPublicVMImagesDefault(code int) *ListPublicVMImagesDefault {
	return &ListPublicVMImagesDefault{
		_statusCode: code,
	}
}

/*
ListPublicVMImagesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListPublicVMImagesDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this list public VM images default response has a 2xx status code
func (o *ListPublicVMImagesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list public VM images default response has a 3xx status code
func (o *ListPublicVMImagesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list public VM images default response has a 4xx status code
func (o *ListPublicVMImagesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list public VM images default response has a 5xx status code
func (o *ListPublicVMImagesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list public VM images default response a status code equal to that given
func (o *ListPublicVMImagesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list public VM images default response
func (o *ListPublicVMImagesDefault) Code() int {
	return o._statusCode
}

func (o *ListPublicVMImagesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/vms/public-images][%d] ListPublicVMImages default  %+v", o._statusCode, o.Payload)
}

func (o *ListPublicVMImagesDefault) String() string {
	return fmt.Sprintf("[GET /v1/vms/public-images][%d] ListPublicVMImages default  %+v", o._statusCode, o.Payload)
}

func (o *ListPublicVMImagesDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *ListPublicVMImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
