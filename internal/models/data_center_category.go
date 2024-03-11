// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataCenterCategory data center category
//
// swagger:model DataCenterCategory
type DataCenterCategory struct {

	// count Vm available
	// Required: true
	CountVMAvailable *int32 `json:"countVmAvailable"`

	// id
	// Required: true
	ID *string `json:"id"`

	// min price hr
	// Required: true
	MinPriceHr *Decimal `json:"minPriceHr"`

	// renewable energy
	// Required: true
	RenewableEnergy *bool `json:"renewableEnergy"`
}

// Validate validates this data center category
func (m *DataCenterCategory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountVMAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinPriceHr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRenewableEnergy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataCenterCategory) validateCountVMAvailable(formats strfmt.Registry) error {

	if err := validate.Required("countVmAvailable", "body", m.CountVMAvailable); err != nil {
		return err
	}

	return nil
}

func (m *DataCenterCategory) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DataCenterCategory) validateMinPriceHr(formats strfmt.Registry) error {

	if err := validate.Required("minPriceHr", "body", m.MinPriceHr); err != nil {
		return err
	}

	if m.MinPriceHr != nil {
		if err := m.MinPriceHr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("minPriceHr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("minPriceHr")
			}
			return err
		}
	}

	return nil
}

func (m *DataCenterCategory) validateRenewableEnergy(formats strfmt.Registry) error {

	if err := validate.Required("renewableEnergy", "body", m.RenewableEnergy); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this data center category based on the context it is used
func (m *DataCenterCategory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMinPriceHr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataCenterCategory) contextValidateMinPriceHr(ctx context.Context, formats strfmt.Registry) error {

	if m.MinPriceHr != nil {

		if err := m.MinPriceHr.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("minPriceHr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("minPriceHr")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataCenterCategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataCenterCategory) UnmarshalBinary(b []byte) error {
	var res DataCenterCategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}