// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// VRouterSize v router size
//
// swagger:model VRouterSize
type VRouterSize string

func NewVRouterSize(value VRouterSize) *VRouterSize {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VRouterSize.
func (m VRouterSize) Pointer() *VRouterSize {
	return &m
}

const (

	// VRouterSizeVROUTERINSTANCEUNKNOWN captures enum value "VROUTER_INSTANCE_UNKNOWN"
	VRouterSizeVROUTERINSTANCEUNKNOWN VRouterSize = "VROUTER_INSTANCE_UNKNOWN"

	// VRouterSizeVROUTERINSTANCESMALL captures enum value "VROUTER_INSTANCE_SMALL"
	VRouterSizeVROUTERINSTANCESMALL VRouterSize = "VROUTER_INSTANCE_SMALL"

	// VRouterSizeVROUTERINSTANCEMEDIUM captures enum value "VROUTER_INSTANCE_MEDIUM"
	VRouterSizeVROUTERINSTANCEMEDIUM VRouterSize = "VROUTER_INSTANCE_MEDIUM"

	// VRouterSizeVROUTERINSTANCELARGE captures enum value "VROUTER_INSTANCE_LARGE"
	VRouterSizeVROUTERINSTANCELARGE VRouterSize = "VROUTER_INSTANCE_LARGE"
)

// for schema
var vRouterSizeEnum []interface{}

func init() {
	var res []VRouterSize
	if err := json.Unmarshal([]byte(`["VROUTER_INSTANCE_UNKNOWN","VROUTER_INSTANCE_SMALL","VROUTER_INSTANCE_MEDIUM","VROUTER_INSTANCE_LARGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vRouterSizeEnum = append(vRouterSizeEnum, v)
	}
}

func (m VRouterSize) validateVRouterSizeEnum(path, location string, value VRouterSize) error {
	if err := validate.EnumCase(path, location, value, vRouterSizeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v router size
func (m VRouterSize) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVRouterSizeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v router size based on context it is used
func (m VRouterSize) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}