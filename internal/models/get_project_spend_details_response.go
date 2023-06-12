// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetProjectSpendDetailsResponse get project spend details response
//
// swagger:model GetProjectSpendDetailsResponse
type GetProjectSpendDetailsResponse struct {

	// orders
	// Required: true
	Orders []*Order `json:"orders"`

	// spend
	// Required: true
	Spend *ProjectSpend `json:"spend"`
}

// Validate validates this get project spend details response
func (m *GetProjectSpendDetailsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpend(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProjectSpendDetailsResponse) validateOrders(formats strfmt.Registry) error {

	if err := validate.Required("orders", "body", m.Orders); err != nil {
		return err
	}

	for i := 0; i < len(m.Orders); i++ {
		if swag.IsZero(m.Orders[i]) { // not required
			continue
		}

		if m.Orders[i] != nil {
			if err := m.Orders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("orders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetProjectSpendDetailsResponse) validateSpend(formats strfmt.Registry) error {

	if err := validate.Required("spend", "body", m.Spend); err != nil {
		return err
	}

	if m.Spend != nil {
		if err := m.Spend.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spend")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spend")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get project spend details response based on the context it is used
func (m *GetProjectSpendDetailsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpend(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProjectSpendDetailsResponse) contextValidateOrders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Orders); i++ {

		if m.Orders[i] != nil {

			if swag.IsZero(m.Orders[i]) { // not required
				return nil
			}

			if err := m.Orders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("orders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetProjectSpendDetailsResponse) contextValidateSpend(ctx context.Context, formats strfmt.Registry) error {

	if m.Spend != nil {

		if err := m.Spend.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spend")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spend")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetProjectSpendDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetProjectSpendDetailsResponse) UnmarshalBinary(b []byte) error {
	var res GetProjectSpendDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
