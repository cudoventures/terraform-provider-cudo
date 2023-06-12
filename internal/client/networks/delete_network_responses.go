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

// DeleteNetworkReader is a Reader for the DeleteNetwork structure.
type DeleteNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNetworkOK creates a DeleteNetworkOK with default headers values
func NewDeleteNetworkOK() *DeleteNetworkOK {
	return &DeleteNetworkOK{}
}

/*
DeleteNetworkOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteNetworkOK struct {
	Payload models.DeleteNetworkResponse
}

// IsSuccess returns true when this delete network o k response has a 2xx status code
func (o *DeleteNetworkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete network o k response has a 3xx status code
func (o *DeleteNetworkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network o k response has a 4xx status code
func (o *DeleteNetworkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete network o k response has a 5xx status code
func (o *DeleteNetworkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network o k response a status code equal to that given
func (o *DeleteNetworkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete network o k response
func (o *DeleteNetworkOK) Code() int {
	return 200
}

func (o *DeleteNetworkOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/networks/{networkId}][%d] deleteNetworkOK  %+v", 200, o.Payload)
}

func (o *DeleteNetworkOK) String() string {
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/networks/{networkId}][%d] deleteNetworkOK  %+v", 200, o.Payload)
}

func (o *DeleteNetworkOK) GetPayload() models.DeleteNetworkResponse {
	return o.Payload
}

func (o *DeleteNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkDefault creates a DeleteNetworkDefault with default headers values
func NewDeleteNetworkDefault(code int) *DeleteNetworkDefault {
	return &DeleteNetworkDefault{
		_statusCode: code,
	}
}

/*
DeleteNetworkDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteNetworkDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this delete network default response has a 2xx status code
func (o *DeleteNetworkDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete network default response has a 3xx status code
func (o *DeleteNetworkDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete network default response has a 4xx status code
func (o *DeleteNetworkDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete network default response has a 5xx status code
func (o *DeleteNetworkDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete network default response a status code equal to that given
func (o *DeleteNetworkDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete network default response
func (o *DeleteNetworkDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworkDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/networks/{networkId}][%d] DeleteNetwork default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNetworkDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/networks/{networkId}][%d] DeleteNetwork default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNetworkDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *DeleteNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
