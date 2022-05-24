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

// PVMInstanceVolumeUpdate p VM instance volume update
//
// swagger:model PVMInstanceVolumeUpdate
type PVMInstanceVolumeUpdate struct {

	// Indicates if the volume should be deleted when the PVMInstance is terminated
	// Required: true
	DeleteOnTermination *bool `json:"deleteOnTermination"`
}

// Validate validates this p VM instance volume update
func (m *PVMInstanceVolumeUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeleteOnTermination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstanceVolumeUpdate) validateDeleteOnTermination(formats strfmt.Registry) error {

	if err := validate.Required("deleteOnTermination", "body", m.DeleteOnTermination); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p VM instance volume update based on context it is used
func (m *PVMInstanceVolumeUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstanceVolumeUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstanceVolumeUpdate) UnmarshalBinary(b []byte) error {
	var res PVMInstanceVolumeUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}