// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceReprovisionInstallRequest service reprovision install request
//
// swagger:model service.ReprovisionInstallRequest
type ServiceReprovisionInstallRequest struct {

	// plan only
	PlanOnly bool `json:"plan_only,omitempty"`
}

// Validate validates this service reprovision install request
func (m *ServiceReprovisionInstallRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service reprovision install request based on context it is used
func (m *ServiceReprovisionInstallRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceReprovisionInstallRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceReprovisionInstallRequest) UnmarshalBinary(b []byte) error {
	var res ServiceReprovisionInstallRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
