// Code generated by go-swagger; DO NOT EDIT.

package ssh_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// CreateSSHKeyReader is a Reader for the CreateSSHKey structure.
type CreateSSHKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSSHKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSSHKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateSSHKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSSHKeyOK creates a CreateSSHKeyOK with default headers values
func NewCreateSSHKeyOK() *CreateSSHKeyOK {
	return &CreateSSHKeyOK{}
}

/*
CreateSSHKeyOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateSSHKeyOK struct {
	Payload *models.SSHKey
}

// IsSuccess returns true when this create Ssh key o k response has a 2xx status code
func (o *CreateSSHKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create Ssh key o k response has a 3xx status code
func (o *CreateSSHKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Ssh key o k response has a 4xx status code
func (o *CreateSSHKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Ssh key o k response has a 5xx status code
func (o *CreateSSHKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create Ssh key o k response a status code equal to that given
func (o *CreateSSHKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create Ssh key o k response
func (o *CreateSSHKeyOK) Code() int {
	return 200
}

func (o *CreateSSHKeyOK) Error() string {
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] createSshKeyOK  %+v", 200, o.Payload)
}

func (o *CreateSSHKeyOK) String() string {
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] createSshKeyOK  %+v", 200, o.Payload)
}

func (o *CreateSSHKeyOK) GetPayload() *models.SSHKey {
	return o.Payload
}

func (o *CreateSSHKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SSHKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSSHKeyDefault creates a CreateSSHKeyDefault with default headers values
func NewCreateSSHKeyDefault(code int) *CreateSSHKeyDefault {
	return &CreateSSHKeyDefault{
		_statusCode: code,
	}
}

/*
CreateSSHKeyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateSSHKeyDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this create Ssh key default response has a 2xx status code
func (o *CreateSSHKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create Ssh key default response has a 3xx status code
func (o *CreateSSHKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create Ssh key default response has a 4xx status code
func (o *CreateSSHKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create Ssh key default response has a 5xx status code
func (o *CreateSSHKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create Ssh key default response a status code equal to that given
func (o *CreateSSHKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create Ssh key default response
func (o *CreateSSHKeyDefault) Code() int {
	return o._statusCode
}

func (o *CreateSSHKeyDefault) Error() string {
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] CreateSshKey default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSSHKeyDefault) String() string {
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] CreateSshKey default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSSHKeyDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *CreateSSHKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
