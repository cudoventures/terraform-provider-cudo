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

// ListVMDataCentersResponse list VM data centers response
//
// swagger:model ListVMDataCentersResponse
type ListVMDataCentersResponse struct {

	// data centers
	// Required: true
	DataCenters []*VMDataCenter `json:"dataCenters"`
}

// Validate validates this list VM data centers response
func (m *ListVMDataCentersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataCenters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListVMDataCentersResponse) validateDataCenters(formats strfmt.Registry) error {

	if err := validate.Required("dataCenters", "body", m.DataCenters); err != nil {
		return err
	}

	for i := 0; i < len(m.DataCenters); i++ {
		if swag.IsZero(m.DataCenters[i]) { // not required
			continue
		}

		if m.DataCenters[i] != nil {
			if err := m.DataCenters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataCenters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataCenters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list VM data centers response based on the context it is used
func (m *ListVMDataCentersResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataCenters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListVMDataCentersResponse) contextValidateDataCenters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DataCenters); i++ {

		if m.DataCenters[i] != nil {

			if swag.IsZero(m.DataCenters[i]) { // not required
				return nil
			}

			if err := m.DataCenters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataCenters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataCenters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListVMDataCentersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListVMDataCentersResponse) UnmarshalBinary(b []byte) error {
	var res ListVMDataCentersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}