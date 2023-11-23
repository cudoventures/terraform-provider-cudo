// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// RemoveDataCenterUserPermissionReader is a Reader for the RemoveDataCenterUserPermission structure.
type RemoveDataCenterUserPermissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveDataCenterUserPermissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveDataCenterUserPermissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveDataCenterUserPermissionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveDataCenterUserPermissionOK creates a RemoveDataCenterUserPermissionOK with default headers values
func NewRemoveDataCenterUserPermissionOK() *RemoveDataCenterUserPermissionOK {
	return &RemoveDataCenterUserPermissionOK{}
}

/*
RemoveDataCenterUserPermissionOK describes a response with status code 200, with default header values.

A successful response.
*/
type RemoveDataCenterUserPermissionOK struct {
	Payload interface{}
}

// IsSuccess returns true when this remove data center user permission o k response has a 2xx status code
func (o *RemoveDataCenterUserPermissionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove data center user permission o k response has a 3xx status code
func (o *RemoveDataCenterUserPermissionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove data center user permission o k response has a 4xx status code
func (o *RemoveDataCenterUserPermissionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove data center user permission o k response has a 5xx status code
func (o *RemoveDataCenterUserPermissionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove data center user permission o k response a status code equal to that given
func (o *RemoveDataCenterUserPermissionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove data center user permission o k response
func (o *RemoveDataCenterUserPermissionOK) Code() int {
	return 200
}

func (o *RemoveDataCenterUserPermissionOK) Error() string {
	return fmt.Sprintf("[POST /v1/data-centers/{dataCenterId}/remove-user-permission][%d] removeDataCenterUserPermissionOK  %+v", 200, o.Payload)
}

func (o *RemoveDataCenterUserPermissionOK) String() string {
	return fmt.Sprintf("[POST /v1/data-centers/{dataCenterId}/remove-user-permission][%d] removeDataCenterUserPermissionOK  %+v", 200, o.Payload)
}

func (o *RemoveDataCenterUserPermissionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RemoveDataCenterUserPermissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveDataCenterUserPermissionDefault creates a RemoveDataCenterUserPermissionDefault with default headers values
func NewRemoveDataCenterUserPermissionDefault(code int) *RemoveDataCenterUserPermissionDefault {
	return &RemoveDataCenterUserPermissionDefault{
		_statusCode: code,
	}
}

/*
RemoveDataCenterUserPermissionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RemoveDataCenterUserPermissionDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this remove data center user permission default response has a 2xx status code
func (o *RemoveDataCenterUserPermissionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this remove data center user permission default response has a 3xx status code
func (o *RemoveDataCenterUserPermissionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this remove data center user permission default response has a 4xx status code
func (o *RemoveDataCenterUserPermissionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this remove data center user permission default response has a 5xx status code
func (o *RemoveDataCenterUserPermissionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this remove data center user permission default response a status code equal to that given
func (o *RemoveDataCenterUserPermissionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the remove data center user permission default response
func (o *RemoveDataCenterUserPermissionDefault) Code() int {
	return o._statusCode
}

func (o *RemoveDataCenterUserPermissionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/data-centers/{dataCenterId}/remove-user-permission][%d] RemoveDataCenterUserPermission default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveDataCenterUserPermissionDefault) String() string {
	return fmt.Sprintf("[POST /v1/data-centers/{dataCenterId}/remove-user-permission][%d] RemoveDataCenterUserPermission default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveDataCenterUserPermissionDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *RemoveDataCenterUserPermissionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
RemoveDataCenterUserPermissionBody remove data center user permission body
swagger:model RemoveDataCenterUserPermissionBody
*/
type RemoveDataCenterUserPermissionBody struct {

	// billing account Id
	BillingAccountID string `json:"billingAccountId,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// role
	// Required: true
	Role *models.Role `json:"role"`

	// user Id
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this remove data center user permission body
func (o *RemoveDataCenterUserPermissionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveDataCenterUserPermissionBody) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"role", "body", o.Role); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"role", "body", o.Role); err != nil {
		return err
	}

	if o.Role != nil {
		if err := o.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "role")
			}
			return err
		}
	}

	return nil
}

func (o *RemoveDataCenterUserPermissionBody) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"userId", "body", o.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this remove data center user permission body based on the context it is used
func (o *RemoveDataCenterUserPermissionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveDataCenterUserPermissionBody) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if o.Role != nil {

		if err := o.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "role")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RemoveDataCenterUserPermissionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveDataCenterUserPermissionBody) UnmarshalBinary(b []byte) error {
	var res RemoveDataCenterUserPermissionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}