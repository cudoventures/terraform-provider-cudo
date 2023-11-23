// Code generated by go-swagger; DO NOT EDIT.

package object_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// GenerateObjectStorageKeyReader is a Reader for the GenerateObjectStorageKey structure.
type GenerateObjectStorageKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateObjectStorageKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateObjectStorageKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGenerateObjectStorageKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGenerateObjectStorageKeyOK creates a GenerateObjectStorageKeyOK with default headers values
func NewGenerateObjectStorageKeyOK() *GenerateObjectStorageKeyOK {
	return &GenerateObjectStorageKeyOK{}
}

/*
GenerateObjectStorageKeyOK describes a response with status code 200, with default header values.

A successful response.
*/
type GenerateObjectStorageKeyOK struct {
	Payload *models.ObjectStorageKey
}

// IsSuccess returns true when this generate object storage key o k response has a 2xx status code
func (o *GenerateObjectStorageKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate object storage key o k response has a 3xx status code
func (o *GenerateObjectStorageKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate object storage key o k response has a 4xx status code
func (o *GenerateObjectStorageKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate object storage key o k response has a 5xx status code
func (o *GenerateObjectStorageKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this generate object storage key o k response a status code equal to that given
func (o *GenerateObjectStorageKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the generate object storage key o k response
func (o *GenerateObjectStorageKeyOK) Code() int {
	return 200
}

func (o *GenerateObjectStorageKeyOK) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/object-storage/users/{id}][%d] generateObjectStorageKeyOK  %+v", 200, o.Payload)
}

func (o *GenerateObjectStorageKeyOK) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/object-storage/users/{id}][%d] generateObjectStorageKeyOK  %+v", 200, o.Payload)
}

func (o *GenerateObjectStorageKeyOK) GetPayload() *models.ObjectStorageKey {
	return o.Payload
}

func (o *GenerateObjectStorageKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectStorageKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateObjectStorageKeyDefault creates a GenerateObjectStorageKeyDefault with default headers values
func NewGenerateObjectStorageKeyDefault(code int) *GenerateObjectStorageKeyDefault {
	return &GenerateObjectStorageKeyDefault{
		_statusCode: code,
	}
}

/*
GenerateObjectStorageKeyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GenerateObjectStorageKeyDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this generate object storage key default response has a 2xx status code
func (o *GenerateObjectStorageKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this generate object storage key default response has a 3xx status code
func (o *GenerateObjectStorageKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this generate object storage key default response has a 4xx status code
func (o *GenerateObjectStorageKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this generate object storage key default response has a 5xx status code
func (o *GenerateObjectStorageKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this generate object storage key default response a status code equal to that given
func (o *GenerateObjectStorageKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the generate object storage key default response
func (o *GenerateObjectStorageKeyDefault) Code() int {
	return o._statusCode
}

func (o *GenerateObjectStorageKeyDefault) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/object-storage/users/{id}][%d] GenerateObjectStorageKey default  %+v", o._statusCode, o.Payload)
}

func (o *GenerateObjectStorageKeyDefault) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/object-storage/users/{id}][%d] GenerateObjectStorageKey default  %+v", o._statusCode, o.Payload)
}

func (o *GenerateObjectStorageKeyDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *GenerateObjectStorageKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}