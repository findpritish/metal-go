// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1MachineNicExtended v1 machine nic extended
// swagger:model v1.MachineNicExtended
type V1MachineNicExtended struct {

	// the mac address of this network interface
	// Required: true
	Mac *string `json:"mac"`

	// the name of this network interface
	// Required: true
	Name *string `json:"name"`

	// the neighbors visible to this network interface
	// Required: true
	Neighbors []*V1MachineNicExtended `json:"neighbors"`
}

// Validate validates this v1 machine nic extended
func (m *V1MachineNicExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMac(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNeighbors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineNicExtended) validateMac(formats strfmt.Registry) error {

	if err := validate.Required("mac", "body", m.Mac); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineNicExtended) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineNicExtended) validateNeighbors(formats strfmt.Registry) error {

	if err := validate.Required("neighbors", "body", m.Neighbors); err != nil {
		return err
	}

	for i := 0; i < len(m.Neighbors); i++ {
		if swag.IsZero(m.Neighbors[i]) { // not required
			continue
		}

		if m.Neighbors[i] != nil {
			if err := m.Neighbors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("neighbors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineNicExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineNicExtended) UnmarshalBinary(b []byte) error {
	var res V1MachineNicExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
