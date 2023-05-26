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

// SearchCompute2Reader is a Reader for the SearchCompute2 structure.
type SearchCompute2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchCompute2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchCompute2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchCompute2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchCompute2OK creates a SearchCompute2OK with default headers values
func NewSearchCompute2OK() *SearchCompute2OK {
	return &SearchCompute2OK{}
}

/*
	SearchCompute2OK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchCompute2OK struct {
	Payload *models.SearchComputeResponse
}

// IsSuccess returns true when this search compute2 o k response has a 2xx status code
func (o *SearchCompute2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search compute2 o k response has a 3xx status code
func (o *SearchCompute2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search compute2 o k response has a 4xx status code
func (o *SearchCompute2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search compute2 o k response has a 5xx status code
func (o *SearchCompute2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this search compute2 o k response a status code equal to that given
func (o *SearchCompute2OK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchCompute2OK) Error() string {
	return fmt.Sprintf("[GET /v1/compute/search_v2][%d] searchCompute2OK  %+v", 200, o.Payload)
}

func (o *SearchCompute2OK) String() string {
	return fmt.Sprintf("[GET /v1/compute/search_v2][%d] searchCompute2OK  %+v", 200, o.Payload)
}

func (o *SearchCompute2OK) GetPayload() *models.SearchComputeResponse {
	return o.Payload
}

func (o *SearchCompute2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SearchComputeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchCompute2Default creates a SearchCompute2Default with default headers values
func NewSearchCompute2Default(code int) *SearchCompute2Default {
	return &SearchCompute2Default{
		_statusCode: code,
	}
}

/*
	SearchCompute2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SearchCompute2Default struct {
	_statusCode int

	Payload *models.Status
}

// Code gets the status code for the search compute2 default response
func (o *SearchCompute2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this search compute2 default response has a 2xx status code
func (o *SearchCompute2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search compute2 default response has a 3xx status code
func (o *SearchCompute2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search compute2 default response has a 4xx status code
func (o *SearchCompute2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search compute2 default response has a 5xx status code
func (o *SearchCompute2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search compute2 default response a status code equal to that given
func (o *SearchCompute2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SearchCompute2Default) Error() string {
	return fmt.Sprintf("[GET /v1/compute/search_v2][%d] SearchCompute2 default  %+v", o._statusCode, o.Payload)
}

func (o *SearchCompute2Default) String() string {
	return fmt.Sprintf("[GET /v1/compute/search_v2][%d] SearchCompute2 default  %+v", o._statusCode, o.Payload)
}

func (o *SearchCompute2Default) GetPayload() *models.Status {
	return o.Payload
}

func (o *SearchCompute2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}