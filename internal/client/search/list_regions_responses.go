// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// ListRegionsReader is a Reader for the ListRegions structure.
type ListRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListRegionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRegionsOK creates a ListRegionsOK with default headers values
func NewListRegionsOK() *ListRegionsOK {
	return &ListRegionsOK{}
}

/*
ListRegionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListRegionsOK struct {
	Payload *models.ListRegionsResponse
}

// IsSuccess returns true when this list regions o k response has a 2xx status code
func (o *ListRegionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list regions o k response has a 3xx status code
func (o *ListRegionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list regions o k response has a 4xx status code
func (o *ListRegionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list regions o k response has a 5xx status code
func (o *ListRegionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list regions o k response a status code equal to that given
func (o *ListRegionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list regions o k response
func (o *ListRegionsOK) Code() int {
	return 200
}

func (o *ListRegionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/regions][%d] listRegionsOK  %+v", 200, o.Payload)
}

func (o *ListRegionsOK) String() string {
	return fmt.Sprintf("[GET /v1/regions][%d] listRegionsOK  %+v", 200, o.Payload)
}

func (o *ListRegionsOK) GetPayload() *models.ListRegionsResponse {
	return o.Payload
}

func (o *ListRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListRegionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegionsDefault creates a ListRegionsDefault with default headers values
func NewListRegionsDefault(code int) *ListRegionsDefault {
	return &ListRegionsDefault{
		_statusCode: code,
	}
}

/*
ListRegionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListRegionsDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this list regions default response has a 2xx status code
func (o *ListRegionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list regions default response has a 3xx status code
func (o *ListRegionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list regions default response has a 4xx status code
func (o *ListRegionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list regions default response has a 5xx status code
func (o *ListRegionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list regions default response a status code equal to that given
func (o *ListRegionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list regions default response
func (o *ListRegionsDefault) Code() int {
	return o._statusCode
}

func (o *ListRegionsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/regions][%d] ListRegions default  %+v", o._statusCode, o.Payload)
}

func (o *ListRegionsDefault) String() string {
	return fmt.Sprintf("[GET /v1/regions][%d] ListRegions default  %+v", o._statusCode, o.Payload)
}

func (o *ListRegionsDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *ListRegionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}