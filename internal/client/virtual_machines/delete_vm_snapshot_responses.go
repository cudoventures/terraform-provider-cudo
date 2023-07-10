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

// DeleteVMSnapshotReader is a Reader for the DeleteVMSnapshot structure.
type DeleteVMSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVMSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVMSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteVMSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteVMSnapshotOK creates a DeleteVMSnapshotOK with default headers values
func NewDeleteVMSnapshotOK() *DeleteVMSnapshotOK {
	return &DeleteVMSnapshotOK{}
}

/*
DeleteVMSnapshotOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteVMSnapshotOK struct {
	Payload models.DeleteVMSnapshotResponse
}

// IsSuccess returns true when this delete Vm snapshot o k response has a 2xx status code
func (o *DeleteVMSnapshotOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Vm snapshot o k response has a 3xx status code
func (o *DeleteVMSnapshotOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Vm snapshot o k response has a 4xx status code
func (o *DeleteVMSnapshotOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Vm snapshot o k response has a 5xx status code
func (o *DeleteVMSnapshotOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Vm snapshot o k response a status code equal to that given
func (o *DeleteVMSnapshotOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete Vm snapshot o k response
func (o *DeleteVMSnapshotOK) Code() int {
	return 200
}

func (o *DeleteVMSnapshotOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/vms/{id}/snapshots][%d] deleteVmSnapshotOK  %+v", 200, o.Payload)
}

func (o *DeleteVMSnapshotOK) String() string {
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/vms/{id}/snapshots][%d] deleteVmSnapshotOK  %+v", 200, o.Payload)
}

func (o *DeleteVMSnapshotOK) GetPayload() models.DeleteVMSnapshotResponse {
	return o.Payload
}

func (o *DeleteVMSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVMSnapshotDefault creates a DeleteVMSnapshotDefault with default headers values
func NewDeleteVMSnapshotDefault(code int) *DeleteVMSnapshotDefault {
	return &DeleteVMSnapshotDefault{
		_statusCode: code,
	}
}

/*
DeleteVMSnapshotDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteVMSnapshotDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this delete VM snapshot default response has a 2xx status code
func (o *DeleteVMSnapshotDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete VM snapshot default response has a 3xx status code
func (o *DeleteVMSnapshotDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete VM snapshot default response has a 4xx status code
func (o *DeleteVMSnapshotDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete VM snapshot default response has a 5xx status code
func (o *DeleteVMSnapshotDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete VM snapshot default response a status code equal to that given
func (o *DeleteVMSnapshotDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete VM snapshot default response
func (o *DeleteVMSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVMSnapshotDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/vms/{id}/snapshots][%d] DeleteVMSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVMSnapshotDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/vms/{id}/snapshots][%d] DeleteVMSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVMSnapshotDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *DeleteVMSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
