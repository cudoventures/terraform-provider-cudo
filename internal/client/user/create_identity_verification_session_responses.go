// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// CreateIdentityVerificationSessionReader is a Reader for the CreateIdentityVerificationSession structure.
type CreateIdentityVerificationSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIdentityVerificationSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateIdentityVerificationSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateIdentityVerificationSessionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateIdentityVerificationSessionOK creates a CreateIdentityVerificationSessionOK with default headers values
func NewCreateIdentityVerificationSessionOK() *CreateIdentityVerificationSessionOK {
	return &CreateIdentityVerificationSessionOK{}
}

/*
CreateIdentityVerificationSessionOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateIdentityVerificationSessionOK struct {
	Payload *models.IdentityVerificationSession
}

// IsSuccess returns true when this create identity verification session o k response has a 2xx status code
func (o *CreateIdentityVerificationSessionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create identity verification session o k response has a 3xx status code
func (o *CreateIdentityVerificationSessionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create identity verification session o k response has a 4xx status code
func (o *CreateIdentityVerificationSessionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create identity verification session o k response has a 5xx status code
func (o *CreateIdentityVerificationSessionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create identity verification session o k response a status code equal to that given
func (o *CreateIdentityVerificationSessionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create identity verification session o k response
func (o *CreateIdentityVerificationSessionOK) Code() int {
	return 200
}

func (o *CreateIdentityVerificationSessionOK) Error() string {
	return fmt.Sprintf("[GET /v1/auth/create-identity-verification-session][%d] createIdentityVerificationSessionOK  %+v", 200, o.Payload)
}

func (o *CreateIdentityVerificationSessionOK) String() string {
	return fmt.Sprintf("[GET /v1/auth/create-identity-verification-session][%d] createIdentityVerificationSessionOK  %+v", 200, o.Payload)
}

func (o *CreateIdentityVerificationSessionOK) GetPayload() *models.IdentityVerificationSession {
	return o.Payload
}

func (o *CreateIdentityVerificationSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdentityVerificationSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIdentityVerificationSessionDefault creates a CreateIdentityVerificationSessionDefault with default headers values
func NewCreateIdentityVerificationSessionDefault(code int) *CreateIdentityVerificationSessionDefault {
	return &CreateIdentityVerificationSessionDefault{
		_statusCode: code,
	}
}

/*
CreateIdentityVerificationSessionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateIdentityVerificationSessionDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this create identity verification session default response has a 2xx status code
func (o *CreateIdentityVerificationSessionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create identity verification session default response has a 3xx status code
func (o *CreateIdentityVerificationSessionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create identity verification session default response has a 4xx status code
func (o *CreateIdentityVerificationSessionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create identity verification session default response has a 5xx status code
func (o *CreateIdentityVerificationSessionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create identity verification session default response a status code equal to that given
func (o *CreateIdentityVerificationSessionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create identity verification session default response
func (o *CreateIdentityVerificationSessionDefault) Code() int {
	return o._statusCode
}

func (o *CreateIdentityVerificationSessionDefault) Error() string {
	return fmt.Sprintf("[GET /v1/auth/create-identity-verification-session][%d] CreateIdentityVerificationSession default  %+v", o._statusCode, o.Payload)
}

func (o *CreateIdentityVerificationSessionDefault) String() string {
	return fmt.Sprintf("[GET /v1/auth/create-identity-verification-session][%d] CreateIdentityVerificationSession default  %+v", o._statusCode, o.Payload)
}

func (o *CreateIdentityVerificationSessionDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *CreateIdentityVerificationSessionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}