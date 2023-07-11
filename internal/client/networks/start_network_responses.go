// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// StartNetworkReader is a Reader for the StartNetwork structure.
type StartNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartNetworkOK creates a StartNetworkOK with default headers values
func NewStartNetworkOK() *StartNetworkOK {
	return &StartNetworkOK{}
}

/*
StartNetworkOK describes a response with status code 200, with default header values.

A successful response.
*/
type StartNetworkOK struct {
	Payload models.StartNetworkResponse
}

// IsSuccess returns true when this start network o k response has a 2xx status code
func (o *StartNetworkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start network o k response has a 3xx status code
func (o *StartNetworkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start network o k response has a 4xx status code
func (o *StartNetworkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start network o k response has a 5xx status code
func (o *StartNetworkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start network o k response a status code equal to that given
func (o *StartNetworkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the start network o k response
func (o *StartNetworkOK) Code() int {
	return 200
}

func (o *StartNetworkOK) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/networks/{networkId}/start][%d] startNetworkOK  %+v", 200, o.Payload)
}

func (o *StartNetworkOK) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/networks/{networkId}/start][%d] startNetworkOK  %+v", 200, o.Payload)
}

func (o *StartNetworkOK) GetPayload() models.StartNetworkResponse {
	return o.Payload
}

func (o *StartNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartNetworkDefault creates a StartNetworkDefault with default headers values
func NewStartNetworkDefault(code int) *StartNetworkDefault {
	return &StartNetworkDefault{
		_statusCode: code,
	}
}

/*
StartNetworkDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StartNetworkDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this start network default response has a 2xx status code
func (o *StartNetworkDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this start network default response has a 3xx status code
func (o *StartNetworkDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this start network default response has a 4xx status code
func (o *StartNetworkDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this start network default response has a 5xx status code
func (o *StartNetworkDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this start network default response a status code equal to that given
func (o *StartNetworkDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the start network default response
func (o *StartNetworkDefault) Code() int {
	return o._statusCode
}

func (o *StartNetworkDefault) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/networks/{networkId}/start][%d] StartNetwork default  %+v", o._statusCode, o.Payload)
}

func (o *StartNetworkDefault) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/networks/{networkId}/start][%d] StartNetwork default  %+v", o._statusCode, o.Payload)
}

func (o *StartNetworkDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *StartNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
