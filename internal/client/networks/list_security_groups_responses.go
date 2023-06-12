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

// ListSecurityGroupsReader is a Reader for the ListSecurityGroups structure.
type ListSecurityGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSecurityGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSecurityGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListSecurityGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSecurityGroupsOK creates a ListSecurityGroupsOK with default headers values
func NewListSecurityGroupsOK() *ListSecurityGroupsOK {
	return &ListSecurityGroupsOK{}
}

/*
ListSecurityGroupsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListSecurityGroupsOK struct {
	Payload *models.ListSecurityGroupsResponse
}

// IsSuccess returns true when this list security groups o k response has a 2xx status code
func (o *ListSecurityGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list security groups o k response has a 3xx status code
func (o *ListSecurityGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list security groups o k response has a 4xx status code
func (o *ListSecurityGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list security groups o k response has a 5xx status code
func (o *ListSecurityGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list security groups o k response a status code equal to that given
func (o *ListSecurityGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list security groups o k response
func (o *ListSecurityGroupsOK) Code() int {
	return 200
}

func (o *ListSecurityGroupsOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/networks/security-groups][%d] listSecurityGroupsOK  %+v", 200, o.Payload)
}

func (o *ListSecurityGroupsOK) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/networks/security-groups][%d] listSecurityGroupsOK  %+v", 200, o.Payload)
}

func (o *ListSecurityGroupsOK) GetPayload() *models.ListSecurityGroupsResponse {
	return o.Payload
}

func (o *ListSecurityGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListSecurityGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSecurityGroupsDefault creates a ListSecurityGroupsDefault with default headers values
func NewListSecurityGroupsDefault(code int) *ListSecurityGroupsDefault {
	return &ListSecurityGroupsDefault{
		_statusCode: code,
	}
}

/*
ListSecurityGroupsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListSecurityGroupsDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this list security groups default response has a 2xx status code
func (o *ListSecurityGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list security groups default response has a 3xx status code
func (o *ListSecurityGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list security groups default response has a 4xx status code
func (o *ListSecurityGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list security groups default response has a 5xx status code
func (o *ListSecurityGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list security groups default response a status code equal to that given
func (o *ListSecurityGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list security groups default response
func (o *ListSecurityGroupsDefault) Code() int {
	return o._statusCode
}

func (o *ListSecurityGroupsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/networks/security-groups][%d] ListSecurityGroups default  %+v", o._statusCode, o.Payload)
}

func (o *ListSecurityGroupsDefault) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/networks/security-groups][%d] ListSecurityGroups default  %+v", o._statusCode, o.Payload)
}

func (o *ListSecurityGroupsDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *ListSecurityGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
