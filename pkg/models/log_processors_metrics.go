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

// LogProcessorsMetrics LogProcessorsMetrics
//
// swagger:model LogProcessorsMetrics
type LogProcessorsMetrics []*LogProcessorsMetricsItems0

// Validate validates this log processors metrics
func (m LogProcessorsMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	iLogProcessorsMetricsSize := int64(len(m))

	if err := validate.MaxItems("", "body", iLogProcessorsMetricsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this log processors metrics based on the context it is used
func (m LogProcessorsMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if m[i] != nil {

			if swag.IsZero(m[i]) { // not required
				return nil
			}

			if err := m[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// LogProcessorsMetricsItems0 log processors metrics items0
//
// swagger:model LogProcessorsMetricsItems0
type LogProcessorsMetricsItems0 struct {
	BaseMetrics

	// console options
	ConsoleOptions ConsoleOptions `json:"console_options,omitempty"`

	// Number of datasources per type
	Datasources map[string]int64 `json:"datasources,omitempty"`

	// hub items
	HubItems HubItems `json:"hub_items,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *LogProcessorsMetricsItems0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseMetrics
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseMetrics = aO0

	// AO1
	var dataAO1 struct {
		ConsoleOptions ConsoleOptions `json:"console_options,omitempty"`

		Datasources map[string]int64 `json:"datasources,omitempty"`

		HubItems HubItems `json:"hub_items,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ConsoleOptions = dataAO1.ConsoleOptions

	m.Datasources = dataAO1.Datasources

	m.HubItems = dataAO1.HubItems

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m LogProcessorsMetricsItems0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseMetrics)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ConsoleOptions ConsoleOptions `json:"console_options,omitempty"`

		Datasources map[string]int64 `json:"datasources,omitempty"`

		HubItems HubItems `json:"hub_items,omitempty"`
	}

	dataAO1.ConsoleOptions = m.ConsoleOptions

	dataAO1.Datasources = m.Datasources

	dataAO1.HubItems = m.HubItems

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this log processors metrics items0
func (m *LogProcessorsMetricsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseMetrics
	if err := m.BaseMetrics.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsoleOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHubItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogProcessorsMetricsItems0) validateConsoleOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.ConsoleOptions) { // not required
		return nil
	}

	if err := m.ConsoleOptions.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("console_options")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("console_options")
		}
		return err
	}

	return nil
}

func (m *LogProcessorsMetricsItems0) validateHubItems(formats strfmt.Registry) error {

	if swag.IsZero(m.HubItems) { // not required
		return nil
	}

	if m.HubItems != nil {
		if err := m.HubItems.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hub_items")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hub_items")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this log processors metrics items0 based on the context it is used
func (m *LogProcessorsMetricsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseMetrics
	if err := m.BaseMetrics.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsoleOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHubItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogProcessorsMetricsItems0) contextValidateConsoleOptions(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ConsoleOptions.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("console_options")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("console_options")
		}
		return err
	}

	return nil
}

func (m *LogProcessorsMetricsItems0) contextValidateHubItems(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.HubItems) { // not required
		return nil
	}

	if err := m.HubItems.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hub_items")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("hub_items")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogProcessorsMetricsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogProcessorsMetricsItems0) UnmarshalBinary(b []byte) error {
	var res LogProcessorsMetricsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
