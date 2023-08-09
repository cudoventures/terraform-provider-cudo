// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// CreateNetworkReader is a Reader for the CreateNetwork structure.
type CreateNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateNetworkOK creates a CreateNetworkOK with default headers values
func NewCreateNetworkOK() *CreateNetworkOK {
	return &CreateNetworkOK{}
}

/*
CreateNetworkOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateNetworkOK struct {
	Payload models.CreateNetworkResponse
}

// IsSuccess returns true when this create network o k response has a 2xx status code
func (o *CreateNetworkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network o k response has a 3xx status code
func (o *CreateNetworkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network o k response has a 4xx status code
func (o *CreateNetworkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network o k response has a 5xx status code
func (o *CreateNetworkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create network o k response a status code equal to that given
func (o *CreateNetworkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create network o k response
func (o *CreateNetworkOK) Code() int {
	return 200
}

func (o *CreateNetworkOK) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/networks][%d] createNetworkOK  %+v", 200, o.Payload)
}

func (o *CreateNetworkOK) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/networks][%d] createNetworkOK  %+v", 200, o.Payload)
}

func (o *CreateNetworkOK) GetPayload() models.CreateNetworkResponse {
	return o.Payload
}

func (o *CreateNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNetworkDefault creates a CreateNetworkDefault with default headers values
func NewCreateNetworkDefault(code int) *CreateNetworkDefault {
	return &CreateNetworkDefault{
		_statusCode: code,
	}
}

/*
CreateNetworkDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateNetworkDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this create network default response has a 2xx status code
func (o *CreateNetworkDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create network default response has a 3xx status code
func (o *CreateNetworkDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create network default response has a 4xx status code
func (o *CreateNetworkDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create network default response has a 5xx status code
func (o *CreateNetworkDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create network default response a status code equal to that given
func (o *CreateNetworkDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create network default response
func (o *CreateNetworkDefault) Code() int {
	return o._statusCode
}

func (o *CreateNetworkDefault) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/networks][%d] CreateNetwork default  %+v", o._statusCode, o.Payload)
}

func (o *CreateNetworkDefault) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/networks][%d] CreateNetwork default  %+v", o._statusCode, o.Payload)
}

func (o *CreateNetworkDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *CreateNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateNetworkBody create network body
swagger:model CreateNetworkBody
*/
type CreateNetworkBody struct {

	// cidr prefix
	// Required: true
	CidrPrefix *string `json:"cidrPrefix"`

	// data center Id
	// Required: true
	DataCenterID *string `json:"dataCenterId"`

	// network Id
	// Required: true
	NetworkID *string `json:"networkId"`

	// vrouter size
	VrouterSize *models.VRouterSize `json:"vrouterSize,omitempty"`
}

// Validate validates this create network body
func (o *CreateNetworkBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCidrPrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDataCenterID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetworkID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVrouterSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkBody) validateCidrPrefix(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"cidrPrefix", "body", o.CidrPrefix); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkBody) validateDataCenterID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"dataCenterId", "body", o.DataCenterID); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkBody) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"networkId", "body", o.NetworkID); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkBody) validateVrouterSize(formats strfmt.Registry) error {
	if swag.IsZero(o.VrouterSize) { // not required
		return nil
	}

	if o.VrouterSize != nil {
		if err := o.VrouterSize.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "vrouterSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "vrouterSize")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create network body based on the context it is used
func (o *CreateNetworkBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateVrouterSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkBody) contextValidateVrouterSize(ctx context.Context, formats strfmt.Registry) error {

	if o.VrouterSize != nil {
		if err := o.VrouterSize.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "vrouterSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "vrouterSize")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
